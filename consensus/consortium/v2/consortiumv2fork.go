package v2

import (
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	wiggleTimeBeforeFork       = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers
	fixedBackOffTimeBeforeFork = 200 * time.Millisecond
)

func (c *Consortium) delayForConsortiumV2Fork(snap *Snapshot, header *types.Header) time.Duration {
	delay := time.Until(time.Unix(int64(header.Time), 0)) // nolint: gosimple
	if c.chainConfig.IsConsortiumV2(header.Number) {
		return delay
	}
	if header.Difficulty.Cmp(diffNoTurn) == 0 {
		// It's not our turn explicitly to sign, delay it a bit
		wiggle := time.Duration(len(snap.Validators)/2+1) * wiggleTimeBeforeFork
		delay += fixedBackOffTimeBeforeFork + time.Duration(rand.Int63n(int64(wiggle)))
	}
	return delay
}

// Ensure the timestamp has the correct delay
func (c *Consortium) blockTimeForConsortiumV2Fork(snap *Snapshot, header, parent *types.Header) uint64 {
	blockTime := parent.Time + c.config.Period
	return blockTime
}
