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
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

// Snapshot is the state of the authorization voting at a given point in time.
type Snapshot struct {
	config   *params.ConsortiumConfig // Consensus engine parameters to fine tune behavior
	sigcache *lru.ARCCache            // Cache of recent block signatures to speed up ecrecover

	Number     uint64                      `json:"number"`     // Block number where the snapshot was created
	Hash       common.Hash                 `json:"hash"`       // Block hash where the snapshot was created
	SignerSet  map[common.Address]struct{} `json:"signerSet"`  // Set of authorized signers at this moment
	SignerList []common.Address            `json:"signerList"` // List of authorized signers at this moment
	Recents    map[uint64]common.Address   `json:"recents"`    // Set of recent signers for spam protections
}

// newSnapshot creates a new snapshot with the specified startup parameters. This
// method does not initialize the set of recent signers, so only ever use if for
// the genesis block.
func newSnapshot(config *params.ConsortiumConfig, sigcache *lru.ARCCache, number uint64, hash common.Hash, signers []common.Address) *Snapshot {
	snap := &Snapshot{
		config:     config,
		sigcache:   sigcache,
		Number:     number,
		Hash:       hash,
		SignerSet:  make(map[common.Address]struct{}),
		SignerList: make([]common.Address, 0, len(signers)),
		Recents:    make(map[uint64]common.Address),
	}

	for _, signer := range signers {
		snap.SignerSet[signer] = struct{}{}
		snap.SignerList = append(snap.SignerList, signer)
	}

	return snap
}

// loadSnapshot loads an existing snapshot from the database.
func loadSnapshot(config *params.ConsortiumConfig, sigcache *lru.ARCCache, db ethdb.Database, hash common.Hash) (*Snapshot, error) {
	blob, err := rawdb.ReadSnapshotConsortium(db, hash)
	if err != nil {
		return nil, err
	}
	snap := new(Snapshot)
	if err := json.Unmarshal(blob, snap); err != nil {
		return nil, err
	}
	snap.config = config
	snap.sigcache = sigcache

	return snap, nil
}

// store inserts the snapshot into the database.
func (s *Snapshot) store(db ethdb.Database) error {
	blob, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return rawdb.WriteSnapshotConsortium(db, s.Hash, blob)
}

// copy creates a deep copy of the snapshot, though not the individual votes.
func (s *Snapshot) copy() *Snapshot {
	cpy := &Snapshot{
		config:     s.config,
		sigcache:   s.sigcache,
		Number:     s.Number,
		Hash:       s.Hash,
		SignerSet:  make(map[common.Address]struct{}),
		SignerList: make([]common.Address, 0, len(s.SignerList)),
		Recents:    make(map[uint64]common.Address),
	}

	for _, signer := range s.SignerList {
		cpy.SignerSet[signer] = struct{}{}
		cpy.SignerList = append(cpy.SignerList, signer)
	}

	for block, signer := range s.Recents {
		cpy.Recents[block] = signer
	}

	return cpy
}

// apply creates a new authorization snapshot by applying the given headers to
// the original one.
func (s *Snapshot) apply(chain consensus.ChainHeaderReader, c *Consortium, headers []*types.Header, parents []*types.Header) (*Snapshot, error) {
	// Allow passing in no headers for cleaner code
	if len(headers) == 0 {
		return s, nil
	}

	// Iterate through the headers and create a new snapshot
	snap := s.copy()

	var (
		start  = time.Now()
		logged = time.Now()
	)
	for i, header := range headers {
		number := header.Number.Uint64()
		// Delete the oldest signer from the recent list to allow it signing again
		if limit := uint64(len(snap.SignerSet)/2 + 1); number >= limit {
			delete(snap.Recents, number-limit)
		}

		// Resolve the authorization key and check against signers
		signer, err := Ecrecover(header, s.sigcache)
		if err != nil {
			return nil, err
		}
		if _, ok := snap.SignerSet[signer]; !ok {
			return nil, errUnauthorizedSigner
		}

		// If we're taking too much time (ecrecover), notify the user once a while
		if time.Since(logged) > 8*time.Second {
			log.Info("Reconstructing snapshot", "processed", i, "total", len(headers), "elapsed", common.PrettyDuration(time.Since(start)))
			logged = time.Now()
		}
		snap.Recents[number] = signer
	}

	snap.Number += uint64(len(headers))
	snap.Hash = headers[len(headers)-1].Hash()

	// Update the list of signers
	number := headers[len(headers)-1].Number.Uint64()
	validators, err := c.getValidatorsFromLastCheckpoint(chain, number-1, parents)
	if err != nil {
		return nil, err
	}
	snap.SignerSet = make(map[common.Address]struct{})
	snap.SignerList = make([]common.Address, 0, len(validators))
	for _, signer := range validators {
		snap.SignerSet[signer] = struct{}{}
		snap.SignerList = append(snap.SignerList, signer)
	}

	// If we're taking too much time, notify the user once a while
	if time.Since(start) > 8*time.Second {
		log.Info("Reconstructed snapshot", "processed", len(headers), "elapsed", common.PrettyDuration(time.Since(start)))
	}

	return snap, nil
}
