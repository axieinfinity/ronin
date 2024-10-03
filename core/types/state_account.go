// Copyright 2021 The go-ethereum Authors
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

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var emptyCodeHash = crypto.Keccak256(nil)

// StateAccount is the Ethereum consensus representation of accounts.
// These objects are stored in the main account trie.
type StateAccount struct {
	Nonce    uint64
	Balance  *big.Int
	Root     common.Hash // merkle root of the storage trie
	CodeHash []byte
}

// NewEmptyStateAccount constructs an empty state account.
func NewEmptyStateAccount() *StateAccount {
	return &StateAccount{
		Balance:  new(big.Int),
		Root:     EmptyRootHash,
		CodeHash: emptyCodeHash,
	}
}

// Copy returns a deep-copied state account object.
func (acct *StateAccount) Copy() *StateAccount {
	var balance *big.Int
	if acct.Balance != nil {
		balance = new(big.Int).Set(acct.Balance)
	}
	return &StateAccount{
		Nonce:    acct.Nonce,
		Balance:  balance,
		Root:     acct.Root,
		CodeHash: common.CopyBytes(acct.CodeHash),
	}
}

type DirtyStateAccount struct {
	Address     common.Address `json:"address"`
	Nonce       uint64         `json:"nonce"`
	Balance     *big.Int       `json:"balance"`
	Root        common.Hash    `json:"root"` // merkle root of the storage trie
	CodeHash    common.Hash    `json:"codeHash"`
	BlockNumber uint64         `json:"blockNumber"`
	BlockHash   common.Hash    `json:"blockHash"`
	Deleted     bool           `json:"deleted"`
	Suicided    bool           `json:"suicided"`
	DirtyCode   bool           `json:"dirtyCode"`
}

type DirtyStateAccountsAndBlock struct {
	BlockHash     common.Hash
	DirtyAccounts []*DirtyStateAccount
}
