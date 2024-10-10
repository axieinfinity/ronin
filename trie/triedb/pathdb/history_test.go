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
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>

package pathdb

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie/testutil"
	"github.com/ethereum/go-ethereum/trie/triestate"
)

const (
	SeedingHistory = 10
)

// randomStateSet generates a random state change set.
func randomStateSet(n int) *triestate.Set {
	var (
		accounts = make(map[common.Address][]byte)
		storages = make(map[common.Address]map[common.Hash][]byte)
	)
	for i := 0; i < n; i++ {
		addr := testutil.RandomAddress()
		storages[addr] = make(map[common.Hash][]byte)
		for j := 0; j < 3; j++ {
			v, _ := rlp.EncodeToBytes(common.TrimLeftZeroes(testutil.RandBytes(32)))
			storages[addr][testutil.RandomHash()] = v
		}
		account := generateAccount(types.EmptyRootHash)
		accounts[addr] = types.SlimAccountRLP(account)
	}
	return triestate.New(accounts, storages, nil)
}

func makeHistory() *history {
	return newHistory(testutil.RandomHash(), types.EmptyRootHash, 0, randomStateSet(3))
}

func makeHistories(n int) []*history {
	var (
		parent = types.EmptyRootHash
		result []*history
	)
	for i := 0; i < n; i++ {
		root := testutil.RandomHash()
		h := newHistory(root, parent, uint64(i), randomStateSet(3))
		parent = root
		result = append(result, h)
	}
	return result
}

func TestEncodeDecodeHistory(t *testing.T) {
	var (
		m   meta
		dec history
		obj = makeHistory()
	)
	// check if meta data can be correctly encode/decode
	blob := obj.meta.encode()
	if err := m.decode(blob); err != nil {
		t.Fatalf("Failed to decode %v", err)
	}
	if !reflect.DeepEqual(&m, obj.meta) {
		t.Fatal("meta is mismatched")
	}

	// check if account/storage data can be correctly encode/decode
	accountData, storageData, accountIndexes, storageIndexes := obj.encode()
	if err := dec.decode(accountData, storageData, accountIndexes, storageIndexes); err != nil {
		t.Fatalf("Failed to decode, err: %v", err)
	}
	if !compareSet(dec.accounts, obj.accounts) {
		t.Fatal("account data is mismatched")
	}
	if !compareStorages(dec.storages, obj.storages) {
		t.Fatal("storage data is mismatched")
	}
	if !compareList(dec.accountList, obj.accountList) {
		t.Fatal("account list is mismatched")
	}
	if !compareStorageList(dec.storageList, obj.storageList) {
		t.Fatal("storage list is mismatched")
	}
}

func checkHistory(t *testing.T, db ethdb.KeyValueReader, freezer *rawdb.ResettableFreezer, id uint64, root common.Hash, exist bool) {
	blob := rawdb.ReadStateHistoryMeta(freezer, id)
	if exist && len(blob) == 0 {
		t.Fatalf("Failed to load trie history, %d", id)
	}
	if !exist && len(blob) != 0 {
		t.Fatalf("Unexpected trie history, %d", id)
	}
	if exist && rawdb.ReadStateID(db, root) == nil {
		t.Fatalf("Root->ID mapping is not found, %d", id)
	}
	if !exist && rawdb.ReadStateID(db, root) != nil {
		t.Fatalf("Unexpected root->ID mapping, %d", id)
	}
}

func checkHistoriesInRange(t *testing.T, db ethdb.KeyValueReader, freezer *rawdb.ResettableFreezer, from, to uint64, roots []common.Hash, exist bool) {
	for i, j := from, 0; i <= to; i, j = i+1, j+1 {
		checkHistory(t, db, freezer, i, roots[j], exist)
	}
}

func TestTruncateHeadHistory(t *testing.T) {
	var (
		roots      []common.Hash
		hs         = makeHistories(SeedingHistory)
		db         = rawdb.NewMemoryDatabase()
		freezer, _ = openFreezer(t.TempDir(), false)
	)
	defer freezer.Close()

	for i := 0; i < len(hs); i++ {
		accountData, storageData, accountIndex, storageIndex := hs[i].encode()
		rawdb.WriteStateHistory(freezer, uint64(i+1), hs[i].meta.encode(), accountIndex, storageIndex, accountData, storageData)
		rawdb.WriteStateID(db, hs[i].meta.root, uint64(i+1))
		roots = append(roots, hs[i].meta.root)
	}
	for size := len(hs); size > 0; size-- {
		pruned, err := truncateFromHead(db, freezer, uint64(size-1))
		if err != nil {
			t.Fatalf("Failed to truncate from head %v", err)
		}
		if pruned != 1 {
			t.Error("Unexpected pruned items", "want", 1, "got", pruned)
		}
		checkHistoriesInRange(t, db, freezer, uint64(size), uint64(SeedingHistory), roots[size-1:], false)
		checkHistoriesInRange(t, db, freezer, uint64(1), uint64(size-1), roots[:size-1], true)
	}
}

/*
Create n histories, write them to the freezer, and start truncate from the tail one by one.
*/
func TestTruncateTailHistory(t *testing.T) {
	var (
		roots        []common.Hash
		hs           = makeHistories(SeedingHistory)
		db           = rawdb.NewMemoryDatabase()
		freezer, err = openFreezer(t.TempDir(), false)
	)
	if err != nil {
		t.Fatalf("Failed to open freezer %v", err)
	}
	defer freezer.Close()

	for i := 0; i < len(hs); i++ {
		accountData, storageData, accountIndex, storageIndex := hs[i].encode()
		// append i-th history to the freezer.
		rawdb.WriteStateHistory(freezer, uint64(i+1), hs[i].meta.encode(), accountIndex, storageIndex, accountData, storageData)
		// Update the root->ID mapping in KeyValue database.
		rawdb.WriteStateID(db, hs[i].meta.root, uint64(i+1))
		roots = append(roots, hs[i].meta.root)
	}
	// truncate from the tail one by one, 1, 2, 3, ..., n-1.
	for newTail := 1; newTail < len(hs); newTail++ {
		pruned, _ := truncateFromTail(db, freezer, uint64(newTail))
		if pruned != 1 {
			t.Error("Unexpected pruned items", "want", 1, "got", pruned)
		}
		// Check this range should not be existed from 1 to newTail.
		checkHistoriesInRange(t, db, freezer, uint64(1), uint64(newTail), roots[:newTail], false)
		// Check this range should be existed from newTail+1 to SeedingHistory.
		checkHistoriesInRange(t, db, freezer, uint64(newTail+1), uint64(SeedingHistory), roots[newTail:], true)
	}
}

func TestTruncateTailHistories(t *testing.T) {
	var cases = []struct {
		limit             uint64
		expectedPruned    int
		maxPruned         uint64
		minUnprunedOffset uint64
		empty             bool
	}{
		{
			1, 9, 9, 10, false,
		},
		{
			0, 10, 10, 0 /* no meaning */, true,
		},
		{
			10, 0, 0, 1, false,
		},
	}

	for i, c := range cases {
		var (
			roots        []common.Hash
			hs           = makeHistories(SeedingHistory)
			db           = rawdb.NewMemoryDatabase()
			freezer, err = openFreezer(t.TempDir()+fmt.Sprintf("%d", i), false)
		)
		if err != nil {
			t.Fatalf("Failed to open freezer %v", err)
		}
		defer freezer.Close()

		// Write SeedingHistory histories to the freezer.
		for i := 0; i < len(hs); i++ {
			accountData, storageData, accountIndex, storageIndex := hs[i].encode()
			rawdb.WriteStateHistory(freezer, uint64(i+1), hs[i].meta.encode(), accountIndex, storageIndex, accountData, storageData)
			rawdb.WriteStateID(db, hs[i].meta.root, uint64(i+1))
			roots = append(roots, hs[i].meta.root)
		}
		// Truncate from the tail, In this case, we truncate a range of histories.
		tail := SeedingHistory - int(c.limit)
		pruned, _ := truncateFromTail(db, freezer, uint64(tail))
		if pruned != c.expectedPruned {
			t.Error("Unexpected pruned items", "want", c.expectedPruned, "got", pruned)
		}
		// In case of empty, jus make sure the range is truncated.
		if c.empty {
			checkHistoriesInRange(t, db, freezer, uint64(1), uint64(SeedingHistory), roots, false)
		} else {
			checkHistoriesInRange(t, db, freezer, uint64(1), c.maxPruned, roots[:tail], false)
			checkHistoriesInRange(t, db, freezer, c.minUnprunedOffset, uint64(SeedingHistory), roots[tail:], true)
		}
	}
}

// openFreezer initializes the freezer instance for storing state histories.
func openFreezer(datadir string, readOnly bool) (*rawdb.ResettableFreezer, error) {
	return rawdb.NewStateHistoryFreezer(datadir, readOnly)
}

func compareSet[k comparable](a, b map[k][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, ok := b[key]
		if !ok {
			return false
		}
		if !bytes.Equal(valA, valB) {
			return false
		}
	}
	return true
}

func compareList[k comparable](a, b []k) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compareStorages(a, b map[common.Address]map[common.Hash][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for h, subA := range a {
		subB, ok := b[h]
		if !ok {
			return false
		}
		if !compareSet(subA, subB) {
			return false
		}
	}
	return true
}

func compareStorageList(a, b map[common.Address][]common.Hash) bool {
	if len(a) != len(b) {
		return false
	}
	for h, la := range a {
		lb, ok := b[h]
		if !ok {
			return false
		}
		if !compareList(la, lb) {
			return false
		}
	}
	return true
}
