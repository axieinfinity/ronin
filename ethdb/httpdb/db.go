package httpdb

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"sync"
	"time"
)

const (
	GET = "consortium_getDBValue"
	ANCIENT = "consortium_ancient"
)
var notfoundErr = errors.New("not found")

func query(client *rpc.Client, method string, params ...interface{}) ([]byte, error) {
	var res string
	if err := client.Call(&res, method, params...); err != nil {
		return nil, err
	}
	return common.Hex2Bytes(res), nil
}

// DB is a read only database which is used to query data from other nodes by using consortium's rpc 'consortium_getDBValue'
// it also caches return values to memorydb and use them lately search without RPC calling everytime.
// DB only supports Get and Has function which query data from other nodes.
type DB struct {
	client *rpc.Client
	db     *memorydb.Database
	items  []string
	caches map[string]int64
	ttl    time.Duration
	interval time.Duration
	cancel chan struct{}
	lock sync.Mutex
}

func NewDB(rpcUrl string, interval, ttl time.Duration) *DB {
	client, err := rpc.DialHTTP(rpcUrl)
	if err != nil {
		log.Error("[httpdb] NewDB", "err", err)
		return nil
	}
	db := &DB{
		client: client,
		db:     memorydb.New(),
		items:  make([]string, 0),
		caches: make(map[string]int64),
		ttl: ttl,
		interval: interval,
		cancel: make(chan struct{}),
	}
	go db.cleanup()
	return db
}

func (db *DB) Close() error {
	db.cancel <- struct{}{}
	return db.db.Close()
}

func (db *DB) Has(key []byte) (bool, error) {
	res, err := db.db.Has(key)
	if err != nil {
		return false, err
	}
	// try to get data from rpc if res is false
	if !res {
		val, err := query(db.client, GET, common.Bytes2Hex(key))
		if err != nil {
			return false, err
		}
		if len(val) == 0 {
			return false, notfoundErr
		}
		// store val to memory db for later use
		db.store(key, val)
	}
	return true, nil
}

func (db *DB) store(key, value []byte) {
	// TODO: only store key in specific cases
	if err := db.db.Put(key, value); err != nil {
		log.Error("[httpDB] store error", "err", err, "key", common.Bytes2Hex(key))
		return
	}

	db.lock.Lock()
	defer db.lock.Unlock()

	db.items = append(db.items, string(key))
	db.caches[string(key)] = time.Now().Add(db.ttl).UnixNano()
}

func (db *DB) cleanup() {
	ticker := time.NewTicker(db.interval)
	for {
		 select {
		 case <-db.cancel:
			 return
		 case <-ticker.C:
			 db.lock.Lock()
			 for i, key := range db.items {
				 // delete until cached time in list is less than now
				 if db.caches[key] < time.Now().UnixNano() {
					 break
				 }
				 if i + 1 < len(db.items) {
					 db.items = db.items[i+1:]
				 } else {
					 db.items = make([]string, 0)
				 }
				 delete(db.caches, key)
				 db.db.Delete([]byte(key))
			 }
			 db.lock.Unlock()
		 }
	}
}

func (db *DB) Get(key []byte) (val []byte, err error) {
	res, err := db.db.Get(key)
	if err != nil && err.Error() != notfoundErr.Error() {
		return nil, err
	}
	// try to get data from rpc if res is nil
	if res == nil {
		val, err = query(db.client, GET, common.Bytes2Hex(key))
		if err != nil {
			return nil, err
		}
		if len(val) == 0 {
			return nil, notfoundErr
		}
		// store val to memory db for later use
		db.store(key, val)
	}
	return val, nil
}

func (db *DB) Put(key, value []byte) error {
	return nil
}

func (db *DB) Delete(key []byte) error {
	return nil
}

func (db *DB) Stats() {

}

func (db *DB) NewBatch() ethdb.Batch {
	return nil
}

func (db *DB) NewIterator(prefix []byte, start []byte) ethdb.Iterator {
	return nil
}

// Stat returns a particular internal stat of the database.
func (db *DB) Stat(property string) (string, error) {
	return "", errors.New("unknown property")
}

// Compact is not supported on a http db
func (db *DB) Compact(start []byte, limit []byte) error {
	return nil
}

// HasAncient returns an indicator whether the specified data exists in the
// ancient store.
func (db *DB) HasAncient(kind string, number uint64) (bool, error) {
	key := ancientKey(kind, number)
	has, err := db.db.Has([]byte(key))
	if err != nil {
		return false, err
	}
	if !has {
		data, err := query(db.client, ANCIENT, kind, number)
		if err != nil {
			return false, err
		}
		if len(data) == 0 {
			return false, notfoundErr
		}
		db.store([]byte(key), data)
	}
	return true, nil
}

// Ancient retrieves an ancient binary blob from the append-only immutable files.
func (db *DB) Ancient(kind string, number uint64) (data []byte, err error) {
	key := ancientKey(kind, number)
	if data, err = db.db.Get([]byte(key)); err != nil && err.Error() != notfoundErr.Error() {
		return nil, err
	}
	if data == nil {
		if data, err = query(db.client, ANCIENT, kind, number); err != nil {
			return nil, err
		}
		if len(data) == 0 {
			return nil, notfoundErr
		}
		db.store([]byte(key), data)
	}
	return data, nil
}

// Ancients returns the ancient item numbers in the ancient store.
func (db *DB) Ancients() (uint64, error) { return 0, nil }

// AncientSize returns the ancient size of the specified category.
func (db *DB) AncientSize(kind string) (uint64, error) { return 0, nil }

// AppendAncient injects all binary blobs belong to block at the end of the
// append-only immutable table files.
func (db *DB) AppendAncient(number uint64, hash, header, body, receipt, td []byte) error { return nil }

// TruncateAncients discards all but the first n ancient data from the ancient store.
func (db *DB) TruncateAncients(n uint64) error { return nil }

// Sync flushes all in-memory ancient store data to disk.
func (db *DB) Sync() error { return nil }

func ancientKey(kind string, number uint64) string {
	return fmt.Sprintf("ancient-%s%d", kind, number)
}