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

package rawdb

import "path/filepath"

// The list of table names of chain freezer. (headers, hashes, bodies, difficulties)

const (
	// chainFreezerHeaderTable indicates the name of the freezer header table.
	chainFreezerHeaderTable = "headers"

	// chainFreezerHashTable indicates the name of the freezer canonical hash table.
	chainFreezerHashTable = "hashes"

	// chainFreezerBodiesTable indicates the name of the freezer block body table.
	chainFreezerBodiesTable = "bodies"

	// chainFreezerReceiptTable indicates the name of the freezer receipts table.
	chainFreezerReceiptTable = "receipts"

	// chainFreezerDifficultyTable indicates the name of the freezer total difficulty table.
	chainFreezerDifficultyTable = "diffs"
)

const (
	// stateHistoryTableSize defines the maximum size of freezer data files.
	stateHistoryTableSize = 2 * 1000 * 1000 * 1000 // 2GB

	// stateHistoryAccountIndex indicates the name of the freezer state history table (Account + Storage).
	stateHistoryMeta         = "history.meta"
	stateHistoryAccountIndex = "account.index"
	stateHistoryStorageIndex = "storage.index"
	stateHistoryAccountData  = "account.data"
	stateHistoryStorageData  = "storage.data"

	namespace = "eth/db/state"
)

// stateFreezerNoSnappy configures whether compression is disabled for the stateHistory.
// https://github.com/golang/snappy, Reason for splititng files for looking up in archive mode easily.
var stateFreezerNoSnappy = map[string]bool{
	stateHistoryMeta:         true,
	stateHistoryAccountIndex: false,
	stateHistoryStorageIndex: false,
	stateHistoryAccountData:  false,
	stateHistoryStorageData:  false,
}

// chainFreezerNoSnappy configures whether compression is disabled for the ancient-tables.
// Hashes and difficulties don't compress well.
var chainFreezerNoSnappy = map[string]bool{
	chainFreezerHeaderTable:     false,
	chainFreezerHashTable:       true,
	chainFreezerBodiesTable:     false,
	chainFreezerReceiptTable:    false,
	chainFreezerDifficultyTable: true,
}

// The list of identifiers of ancient stores. It can split more in the futures.
var (
	chainFreezerName = "chain" // the folder name of chain segment ancient store.
	stateFreezerName = "state" // the folder name of reverse diff ancient store.
)

// freezers the collections of all builtin freezers.
var freezers = []string{chainFreezerName, stateFreezerName}

// NewStateFreezer initializes the freezer for state history.
func NewStateFreezer(ancientDir string, readOnly bool) (*ResettableFreezer, error) {
	return NewResettableFreezer(
		filepath.Join(ancientDir, stateFreezerName), namespace, readOnly,
		stateHistoryTableSize, stateFreezerNoSnappy)
}
