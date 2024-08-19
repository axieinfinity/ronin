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

package trie

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
)

// preimageStore is the store for caching preimages of node key.
type preimageStore struct {
	lock          sync.RWMutex
	disk          ethdb.KeyValueStore
	preimages     map[common.Hash][]byte // Preimages of nodes from the secure trie
	preimagesSize common.StorageSize     // Storage size of the preimages cache
}

// newPreimageStore initializes the store for caching preimages.
func newPreimageStore(disk ethdb.KeyValueStore) *preimageStore {
	return &preimageStore{
		disk:      disk,
		preimages: make(map[common.Hash][]byte),
	}
}

func (store *preimageStore) insertPreimage(preimages map[common.Hash][]byte) {
	store.lock.Lock()
	defer store.lock.Unlock()

	for hash, preimage := range preimages {
		if _, ok := store.preimages[hash]; ok {
			continue
		}
		store.preimages[hash] = preimage
		store.preimagesSize += common.StorageSize(common.HashLength + len(preimage))
	}
}

func (store *preimageStore) preimage(hash common.Hash) []byte {
	// Lock the store for reading
	store.lock.RLock()
	preimage := store.preimages[hash]
	store.lock.RUnlock()
	if preimage != nil {
		return preimage
	}
	// Incase preimage is not existed in memory, then read from disk.
	return rawdb.ReadPreimage(store.disk, hash)
}

func (store *preimageStore) commit(force bool) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	// If preimages size is less than 4MB and not forced to commit, then return.
	if store.preimagesSize <= 4*1024*1024 && !force {
		return nil
	}

	batch := store.disk.NewBatch()
	rawdb.WritePreimages(batch, store.preimages)
	if err := batch.Write(); err != nil {
		return err
	}
	store.preimages, store.preimagesSize = make(map[common.Hash][]byte), 0
	return nil
}

func (store *preimageStore) size() common.StorageSize {
	store.lock.RLock()
	defer store.lock.RUnlock()

	return store.preimagesSize
}
