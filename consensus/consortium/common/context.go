package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
)

// ChainContext supports retrieving headers and consensus parameters from the
// current blockchain to be used during transaction processing.
type ChainContext struct {
	Chain      consensus.ChainHeaderReader
	Consortium consensus.Engine
}

// Engine retrieves the chain's consensus engine.
func (c ChainContext) Engine() consensus.Engine {
	return c.Consortium
}

// GetHeader returns the hash corresponding to their hash.
func (c ChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return c.Chain.GetHeader(hash, number)
}
