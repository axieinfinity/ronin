// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rawdb

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/params"
	"github.com/prometheus/tsdb/fileutil"
)

var (
	// errReadOnly is returned if the freezer is opened in read only mode. All the
	// mutations are disallowed.
	errReadOnly = errors.New("read only")

	// errUnknownTable is returned if the user attempts to read from a table that is
	// not tracked by the freezer.
	errUnknownTable = errors.New("unknown table")

	// errOutOrderInsertion is returned if the user attempts to inject out-of-order
	// binary blobs into the freezer.
	errOutOrderInsertion = errors.New("the append operation is out-order")

	// errSymlinkDatadir is returned if the ancient directory specified by user
	// is a symbolic link.
	errSymlinkDatadir = errors.New("symbolic link datadir is not supported")
)

const (

	// freezerTableSize defines the maximum size of freezer data files.
	freezerTableSize = 2 * 1000 * 1000 * 1000
)

// freezer is an memory mapped append-only database to store immutable chain data
// into flat files:
//
//   - The append only nature ensures that disk writes are minimized.
//   - The memory mapping ensures we can max out system memory for caching without
//     reserving it for go-ethereum. This would also reduce the memory requirements
//     of Geth, and thus also GC overhead.
type Freezer struct {
	frozen    atomic.Uint64 // Number of items already frozen
	threshold atomic.Uint64 // Number of recent blocks not to freeze (params.FullImmutabilityThreshold apart from tests)

	datadir string // Path of root directory of ancient store

	// This lock synchronizes writers and the truncate operation, as well as
	// the "atomic" (batched) read operations.
	writeLock  sync.RWMutex
	writeBatch *freezerBatch

	readonly     bool
	tables       map[string]*freezerTable // Data tables for storing everything
	instanceLock fileutil.Releaser        // File-system lock to prevent double opens

	trigger chan chan struct{} // Manual blocking freeze trigger, test determinism

	quit      chan struct{}
	wg        sync.WaitGroup
	closeOnce sync.Once
}

// newFreezer creates a chain freezer that moves ancient chain data into
// append-only flat file containers.
//
// The 'tables' argument defines the data tables. If the value of a map
// entry is true, snappy compression is disabled for the table.
func NewFreezer(datadir string, namespace string, readonly bool, maxTableSize uint32, tables map[string]bool) (*Freezer, error) {
	// Create the initial freezer object
	var (
		readMeter  = metrics.NewRegisteredMeter(namespace+"ancient/read", nil)
		writeMeter = metrics.NewRegisteredMeter(namespace+"ancient/write", nil)
		sizeGauge  = metrics.NewRegisteredGauge(namespace+"ancient/size", nil)
	)
	// Ensure the datadir is not a symbolic link if it exists.
	if info, err := os.Lstat(datadir); !os.IsNotExist(err) {
		if info.Mode()&os.ModeSymlink != 0 {
			log.Warn("Symbolic link ancient database is not supported", "path", datadir)
			return nil, errSymlinkDatadir
		}
	}
	// Leveldb uses LOCK as the filelock filename. To prevent the
	// name collision, we use FLOCK as the lock name.
	lock, _, err := fileutil.Flock(filepath.Join(datadir, "FLOCK"))
	if err != nil {
		return nil, err
	}
	// Open all the supported data tables
	freezer := &Freezer{
		readonly:     readonly,
		tables:       make(map[string]*freezerTable),
		instanceLock: lock,
		trigger:      make(chan chan struct{}),
		quit:         make(chan struct{}),
		datadir:      datadir,
	}
	freezer.threshold.Store(params.FullImmutabilityThreshold)

	// Create the tables.
	for name, disableSnappy := range tables {
		table, err := newTable(datadir, name, readMeter, writeMeter, sizeGauge, maxTableSize, disableSnappy)
		if err != nil {
			for _, table := range freezer.tables {
				table.Close()
			}
			lock.Release()
			return nil, err
		}
		freezer.tables[name] = table
	}

	// Truncate all tables to common length.
	if err := freezer.repair(); err != nil {
		for _, table := range freezer.tables {
			table.Close()
		}
		lock.Release()
		return nil, err
	}

	// Create the write batch.
	freezer.writeBatch = newFreezerBatch(freezer)

	log.Info("Opened ancient database", "database", datadir, "readonly", readonly)
	return freezer, nil
}

// Close terminates the chain freezer, unmapping all the data files.
func (f *Freezer) Close() error {
	f.writeLock.Lock()
	defer f.writeLock.Unlock()

	var errs []error
	f.closeOnce.Do(func() {
		close(f.quit)
		// Wait for any background freezing to stop
		f.wg.Wait()
		for _, table := range f.tables {
			if err := table.Close(); err != nil {
				errs = append(errs, err)
			}
		}
		if err := f.instanceLock.Release(); err != nil {
			errs = append(errs, err)
		}
	})
	if errs != nil {
		return fmt.Errorf("%v", errs)
	}
	return nil
}

// HasAncient returns an indicator whether the specified ancient data exists
// in the freezer.
func (f *Freezer) HasAncient(kind string, number uint64) (bool, error) {
	if table := f.tables[kind]; table != nil {
		return table.has(number), nil
	}
	return false, nil
}

// Ancient retrieves an ancient binary blob from the append-only immutable files.
func (f *Freezer) Ancient(kind string, number uint64) ([]byte, error) {
	if table := f.tables[kind]; table != nil {
		return table.Retrieve(number)
	}
	return nil, errUnknownTable
}

// AncientRange retrieves multiple items in sequence, starting from the index 'start'.
// It will return
//   - at most 'max' items,
//   - at least 1 item (even if exceeding the maxByteSize), but will otherwise
//     return as many items as fit into maxByteSize.
func (f *Freezer) AncientRange(kind string, start, count, maxBytes uint64) ([][]byte, error) {
	if table := f.tables[kind]; table != nil {
		return table.RetrieveItems(start, count, maxBytes)
	}
	return nil, errUnknownTable
}

// Ancients returns the length of the frozen items.
func (f *Freezer) Ancients() (uint64, error) {
	return f.frozen.Load(), nil
}

// AncientSize returns the ancient size of the specified category.
func (f *Freezer) AncientSize(kind string) (uint64, error) {
	// This needs the write lock to avoid data races on table fields.
	// Speed doesn't matter here, AncientSize is for debugging.
	f.writeLock.RLock()
	defer f.writeLock.RUnlock()

	if table := f.tables[kind]; table != nil {
		return table.size()
	}
	return 0, errUnknownTable
}

// ReadAncients runs the given read operation while ensuring that no writes take place
// on the underlying freezer.
func (f *Freezer) ReadAncients(fn func(ethdb.AncientReaderOp) error) (err error) {
	f.writeLock.RLock()
	defer f.writeLock.RUnlock()
	return fn(f)
}

// ModifyAncients runs the given write operation.
func (f *Freezer) ModifyAncients(fn func(ethdb.AncientWriteOp) error) (writeSize int64, err error) {
	if f.readonly {
		return 0, errReadOnly
	}
	f.writeLock.Lock()
	defer f.writeLock.Unlock()

	// Roll back all tables to the starting position in case of error.
	prevItem := f.frozen.Load()
	defer func() {
		if err != nil {
			// The write operation has failed. Go back to the previous item position.
			for name, table := range f.tables {
				err := table.truncate(prevItem)
				if err != nil {
					log.Error("Freezer table roll-back failed", "table", name, "index", prevItem, "err", err)
				}
			}
		}
	}()

	f.writeBatch.reset()
	if err := fn(f.writeBatch); err != nil {
		return 0, err
	}
	item, writeSize, err := f.writeBatch.commit()
	if err != nil {
		return 0, err
	}
	f.frozen.Store(item)
	return writeSize, nil
}

// TruncateAncients discards any recent data above the provided threshold number.
func (f *Freezer) TruncateAncients(items uint64) error {
	if f.readonly {
		return errReadOnly
	}
	f.writeLock.Lock()
	defer f.writeLock.Unlock()

	if f.frozen.Load() <= items {
		return nil
	}
	for _, table := range f.tables {
		if err := table.truncate(items); err != nil {
			return err
		}
	}
	f.frozen.Store(items)
	return nil
}

// Sync flushes all data tables to disk.
func (f *Freezer) Sync() error {
	var errs []error
	for _, table := range f.tables {
		if err := table.Sync(); err != nil {
			errs = append(errs, err)
		}
	}
	if errs != nil {
		return fmt.Errorf("%v", errs)
	}
	return nil
}

// repair truncates all data tables to the same length.
func (f *Freezer) repair() error {
	min := uint64(math.MaxUint64)
	for _, table := range f.tables {
		items := table.items.Load()
		if min > items {
			min = items
		}
	}
	for _, table := range f.tables {
		if err := table.truncate(min); err != nil {
			return err
		}
	}
	f.frozen.Store(min)
	return nil
}

// AncientDatadir returns the root directory path of the ancient store.
func (f *Freezer) AncientDatadir() (string, error) {
	return f.datadir, nil
}
