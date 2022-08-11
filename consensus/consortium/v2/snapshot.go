package v2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

type Snapshot struct {
}

func newSnapshot(config *params.ConsortiumConfig, sigcache *lru.ARCCache, number uint64, hash common.Hash, signers []common.Address) *Snapshot {
	snap := &Snapshot{}

	return snap
}

func loadSnapshot(config *params.ConsortiumConfig, sigcache *lru.ARCCache, db ethdb.Database, hash common.Hash) (*Snapshot, error) {
	return nil, nil
}

func (s *Snapshot) store(db ethdb.Database) error {
	return nil
}

func (s *Snapshot) copy() *Snapshot {
	return nil
}

func (s *Snapshot) apply(chain consensus.ChainHeaderReader, c *Consortium, headers []*types.Header, parents []*types.Header) (*Snapshot, error) {
	return nil, nil
}
