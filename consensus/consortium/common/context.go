package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
)

type ChainContext struct {
	Chain      consensus.ChainHeaderReader
	Consortium consensus.Engine
}

func (c ChainContext) Engine() consensus.Engine {
	return c.Consortium
}

func (c ChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return c.Chain.GetHeader(hash, number)
}
