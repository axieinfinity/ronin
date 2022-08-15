package v2

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

type Snapshot struct {
	config   *params.ConsortiumConfig // Consensus engine parameters to fine tune behavior
	ethAPI   *ethapi.PublicBlockChainAPI
	sigCache *lru.ARCCache // Cache of recent block signatures to speed up ecrecover

	Number     uint64                      `json:"number"`     // Block number where the snapshot was created
	Hash       common.Hash                 `json:"hash"`       // Block hash where the snapshot was created
	Validators map[common.Address]struct{} `json:"validators"` // Set of authorized validators at this moment
	Recents    map[uint64]common.Address   `json:"recents"`    // Set of recent validators for spam protections
}

// validatorsAscending implements the sort interface to allow sorting a list of addresses
type validatorsAscending []common.Address

func (s validatorsAscending) Len() int           { return len(s) }
func (s validatorsAscending) Less(i, j int) bool { return bytes.Compare(s[i][:], s[j][:]) < 0 }
func (s validatorsAscending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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

// validators retrieves the list of validators in ascending order.
func (s *Snapshot) validators() []common.Address {
	validators := make([]common.Address, 0, len(s.Validators))
	return validators
}

func (s *Snapshot) enoughDistance(validator common.Address, header *types.Header) bool {
	return false
}

func (s *Snapshot) indexOfVal(validator common.Address) int {
	return -1
}
