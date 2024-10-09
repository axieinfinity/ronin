// Copyright 2022 The go-ethereum Authors
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

package pathdb

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/trie/triestate"
)

// maxDiffLayers is the maximum diff layers allowed in the layer tree.
const maxDiffLayers = 128

// layer is the interface implemented by all state layers which includes some
// public methods and some additional methods for internal usage.
type layer interface {
	// Node retrieves the trie node with the node info. An error will be returned
	// if the read operation exits abnormally. For example, if the layer is already
	// stale, or the associated state is regarded as corrupted. Notably, no error
	// will be returned if the requested node is not found in database.
	Node(owner common.Hash, path []byte, hash common.Hash) ([]byte, error)

	// rootHash returns the root hash for which this layer was made.
	rootHash() common.Hash

	// stateID returns the associated state id of layer.
	stateID() uint64

	// parentLayer returns the subsequent layer of it, or nil if the disk was reached.
	parentLayer() layer

	// update creates a new layer on top of the existing layer diff tree with
	// the provided dirty trie nodes along with the state change set.
	//
	// Note, the maps are retained by the method to avoid copying everything.
	update(root common.Hash, id uint64, block uint64, nodes map[common.Hash]map[string]*trienode.Node, states *triestate.Set) *diffLayer
}

// Config contains the settings for database.
type Config struct {
	StateLimit uint64 // Number of recent blocks to maintain state history for
	CleanSize  int    // Maximum memory allowance (in bytes) for caching clean nodes
	DirtySize  int    // Maximum memory allowance (in bytes) for caching dirty nodes
	ReadOnly   bool   // Flag whether the database is opened in read only mode.
}

var (
	// defaultCleanSize is the default memory allowance of clean cache.
	defaultCleanSize = 16 * 1024 * 1024

	// defaultBufferSize is the default memory allowance of node buffer
	// that aggregates the writes from above until it's flushed into the
	// disk. Do not increase the buffer size arbitrarily, otherwise the
	// system pause time will increase when the database writes happen.
	defaultBufferSize = 128 * 1024 * 1024
)

// Defaults contains default settings for Ethereum mainnet.
var Defaults = &Config{
	StateLimit: params.FullImmutabilityThreshold,
	CleanSize:  defaultCleanSize,
	DirtySize:  defaultBufferSize,
}

// Database is a multiple-layered structure for maintaining in-memory trie nodes.
// It consists of one persistent base layer backed by a key-value store, on top
// of which arbitrarily many in-memory diff layers are stacked. The memory diffs
// can form a tree with branching, but the disk layer is singleton and common to
// all. If a reorg goes deeper than the disk layer, a batch of reverse diffs can
// be applied to rollback. The deepest reorg that can be handled depends on the
// amount of state histories tracked in the disk.
//
// At most one readable and writable database can be opened at the same time in
// the whole system which ensures that only one database writer can operate disk
// state. Unexpected open operations can cause the system to panic.
type Database struct {
	// readOnly is the flag whether the mutation is allowed to be applied.
	// It will be set automatically when the database is journaled during
	// the shutdown to reject all following unexpected mutations.
	readOnly   bool                     // Indicator if database is opened in read only mode
	bufferSize int                      // Memory allowance (in bytes) for caching dirty nodes
	config     *Config                  // Configuration for database
	diskdb     ethdb.Database           // Persistent storage for matured trie nodes
	tree       *layerTree               // The group for all known layers
	freezer    *rawdb.ResettableFreezer // Freezer for storing trie histories, nil possible in tests
	lock       sync.RWMutex             // Lock to prevent mutations from happening at the same time
}
