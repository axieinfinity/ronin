package v2

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Consortium) blockTimeForConsortiumV2Fork(snap *Snapshot, header, parent *types.Header) uint64 {
	blockTime := parent.Time + c.config.Period
	//if c.chainConfig.IsConsortiumV2(header.Number) {
	//	blockTime = blockTime + backOffTime(snap, c.val)
	//}
	return blockTime
}
