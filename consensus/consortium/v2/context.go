package v2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
)

// chain context
type chainContext struct {
	Chain      consensus.ChainHeaderReader
	consortium consensus.Engine
}

func (c chainContext) Engine() consensus.Engine {
	return c.consortium
}

func (c chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return c.Chain.GetHeader(hash, number)
}
