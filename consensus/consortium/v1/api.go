// Copyright 2017 The go-ethereum Authors
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

package v1

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the proof-of-authority scheme.
type API struct {
	chain      consensus.ChainHeaderReader
	consortium *Consortium
}

// GetSigners retrieves the list of authorized signers at the specified block.
func (api *API) GetSigners(number *rpc.BlockNumber) ([]common.Address, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return the signers from its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}

	validators, err := api.consortium.getValidatorsFromLastCheckpoint(api.chain, header.Number.Uint64(), nil)
	if err != nil {
		return nil, err
	}
	return validators, nil
}

func (api *API) GetDBValue(key string) (string, error) {
	value, err := api.chain.DB().Get(common.Hex2Bytes(key))
	if err != nil {
		log.Debug("Get value in DB failed, try to get it from trie node", "key", key)
		if value, err = api.chain.StateCache().TrieDB().Node(common.HexToHash(key)); err != nil {
			return common.Bytes2Hex([]byte{}), err
		}
	}
	return common.Bytes2Hex(value), nil
}

func (api *API) GetAncientValue(kind string, number rpc.BlockNumber) (string, error) {
	value, err := api.chain.DB().Ancient(kind, uint64(number.Int64()))
	if err != nil {
		return common.Bytes2Hex([]byte{}), err
	}
	return common.Bytes2Hex(value), nil
}
