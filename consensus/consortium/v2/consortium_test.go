package v2

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func TestSealableValidators(t *testing.T) {
	const NUM_OF_VALIDATORS = 21

	validators := make([]common.Address, NUM_OF_VALIDATORS)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		validators = append(validators, common.BigToAddress(big.NewInt(int64(i))))
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil)
	for i := 0; i <= 10; i++ {
		snap.Recents[uint64(i)] = common.BigToAddress(big.NewInt(int64(i)))
	}

	for i := 1; i <= 10; i++ {
		position, _ := snap.sealableValidators(common.BigToAddress(big.NewInt(int64(i))))
		if position != -1 {
			t.Errorf("Validator that is not allowed to seal is in sealable list, position: %d", position)
		}
	}

	// Validator 0 is allowed to seal block again, current block (block 11) shifts it out of recent list
	position, numOfSealableValidators := snap.sealableValidators(common.BigToAddress(common.Big0))
	if position < 0 || position >= numOfSealableValidators {
		t.Errorf("Sealable validator has invalid position, position: %d", position)
	}

	for i := 11; i < NUM_OF_VALIDATORS; i++ {
		position, numOfSealableValidators := snap.sealableValidators(common.BigToAddress(big.NewInt(int64(i))))
		if position < 0 || position >= numOfSealableValidators {
			t.Errorf("Sealable validator has invalid position, position: %d", position)
		}

		if numOfSealableValidators != 11 {
			t.Errorf("Invalid number of sealable validators, got %d exp %d", numOfSealableValidators, 11)
		}
	}
}

// This test assumes the wiggleTime is 1 second so the delay
// ranges from [0, 6]
func TestBackoffTime(t *testing.T) {
	const NUM_OF_VALIDATORS = 21
	const MAX_DELAY = 6

	c := Consortium{
		chainConfig: &params.ChainConfig{
			BubaBlock: big.NewInt(0),
		},
	}

	validators := make([]common.Address, NUM_OF_VALIDATORS)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		validators = append(validators, common.BigToAddress(big.NewInt(int64(i))))
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil)
	for i := 0; i <= 10; i++ {
		snap.Recents[uint64(i)] = common.BigToAddress(big.NewInt(int64(i)))
	}

	delayMapping := make(map[uint64]int)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		val := common.BigToAddress(big.NewInt(int64(i)))
		header := &types.Header{
			Coinbase: val,
			Number:   new(big.Int).SetUint64(snap.Number + 1),
		}
		delay := backOffTime(header, snap, c.chainConfig)
		if delay == 0 {
			// Validator in recent sign list is not able to seal block
			// and has 0 backOffTime
			inRecent := false
			for _, recent := range snap.Recents {
				if recent == val {
					inRecent = true
					break
				}
			}
			if !inRecent && !snap.inturn(val) {
				t.Error("Out of turn validator has no delay")
			}
		} else if delay > MAX_DELAY {
			t.Errorf("Validator's delay exceeds max limit, delay: %d", delay)
		} else if delayMapping[delay] > 2 {
			t.Errorf("More than 2 validators have the same delay, delay %d", delay)
		}

		delayMapping[delay]++
	}
}

// This test assumes the wiggleTime is 1 second so the delay
// ranges from [0, 11]
func TestBackoffTimeOlek(t *testing.T) {
	const NUM_OF_VALIDATORS = 21
	const MAX_DELAY = 11

	c := Consortium{
		chainConfig: &params.ChainConfig{
			BubaBlock: big.NewInt(0),
			OlekBlock: big.NewInt(0),
		},
	}

	validators := make([]common.Address, NUM_OF_VALIDATORS)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		validators = append(validators, common.BigToAddress(big.NewInt(int64(i))))
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil)
	for i := 0; i <= 10; i++ {
		snap.Recents[uint64(i)] = common.BigToAddress(big.NewInt(int64(i)))
	}

	delayMapping := make(map[uint64]int)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		val := common.BigToAddress(big.NewInt(int64(i)))
		header := &types.Header{
			Coinbase: val,
			Number:   new(big.Int).SetUint64(snap.Number + 1),
		}
		delay := backOffTime(header, snap, c.chainConfig)
		if delay == 0 {
			// Validator in recent sign list is not able to seal block
			// and has 0 backOffTime
			inRecent := false
			for _, recent := range snap.Recents {
				if recent == val {
					inRecent = true
					break
				}
			}
			if !inRecent && !snap.inturn(val) {
				t.Error("Out of turn validator has no delay")
			}
		} else if delay > MAX_DELAY {
			t.Errorf("Validator's delay exceeds max limit, delay: %d", delay)
		} else if delayMapping[delay] > 1 {
			t.Errorf("More than 1 validator have the same delay, delay %d", delay)
		}

		delayMapping[delay]++
	}
}

// When validator is in recent list we expect the minimum delay is
// 1s before Olek and 0s after Olek
func TestBackoffTimeInturnValidatorInRecentList(t *testing.T) {
	const NUM_OF_VALIDATORS = 21

	c := Consortium{
		chainConfig: &params.ChainConfig{
			OlekBlock: big.NewInt(12),
		},
	}

	validators := make([]common.Address, NUM_OF_VALIDATORS)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		validators = append(validators, common.BigToAddress(big.NewInt(int64(i))))
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil)
	for i := 0; i <= 9; i++ {
		snap.Recents[uint64(i)] = common.BigToAddress(big.NewInt(int64(i)))
	}
	snap.Recents[10] = common.BigToAddress(big.NewInt(int64(11)))

	var minDelay uint64 = 10000
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		val := common.BigToAddress(big.NewInt(int64(i)))
		header := &types.Header{
			Coinbase: val,
			Number:   new(big.Int).SetUint64(snap.Number + 1),
		}
		// This validator is not in recent list
		if position, _ := snap.sealableValidators(val); position != -1 {
			delay := backOffTime(header, snap, c.chainConfig)
			if delay < minDelay {
				minDelay = delay
			}
		}
	}

	if minDelay != 1 {
		t.Errorf("Expect min delay is 1s before Olek, get %ds", minDelay)
	}

	c.chainConfig.OlekBlock = big.NewInt(0)
	minDelay = 10000
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		val := common.BigToAddress(big.NewInt(int64(i)))
		header := &types.Header{
			Coinbase: val,
			Number:   new(big.Int).SetUint64(snap.Number + 1),
		}
		// This validator is not in recent list
		if position, _ := snap.sealableValidators(val); position != -1 {
			delay := backOffTime(header, snap, c.chainConfig)
			if delay < minDelay {
				minDelay = delay
			}
		}
	}

	if minDelay != 0 {
		t.Errorf("Expect min delay is 0s before Olek, get %ds", minDelay/uint64(time.Second))
	}
}

func TestVerifyBlockHeaderTime(t *testing.T) {
	const NUM_OF_VALIDATORS = 21
	const BLOCK_PERIOD = 3

	validators := make([]common.Address, NUM_OF_VALIDATORS)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		validators = append(validators, common.BigToAddress(big.NewInt(int64(i))))
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil)
	for i := 0; i <= 10; i++ {
		snap.Recents[uint64(i)] = common.BigToAddress(big.NewInt(int64(i)))
	}

	c := Consortium{
		chainConfig: &params.ChainConfig{
			BubaBlock: big.NewInt(12),
		},
		config: &params.ConsortiumConfig{
			Period: BLOCK_PERIOD,
		},
	}

	now := uint64(time.Now().Unix())
	header := &types.Header{
		Coinbase: common.BigToAddress(big.NewInt(18)),
		Number:   big.NewInt(11),
		Time:     now + 100 + BLOCK_PERIOD,
	}
	parentHeader := &types.Header{
		Number: big.NewInt(10),
		Time:   now + 100,
	}
	if err := c.verifyHeaderTime(header, parentHeader, snap); !errors.Is(err, consensus.ErrFutureBlock) {
		t.Error("Expect future block error when block's timestamp is higher than current timestamp")
	}

	parentHeader.Time = now - 10
	header.Time = now - 9
	if err := c.verifyHeaderTime(header, parentHeader, snap); err != nil {
		t.Errorf("Expect successful verification, got %s", err)
	}

	c.chainConfig.BubaBlock = big.NewInt(11)
	if err := c.verifyHeaderTime(header, parentHeader, snap); !errors.Is(err, consensus.ErrFutureBlock) {
		t.Errorf("Expect future block error when block's timestamp is lower than minimum requirements")
	}

	header.Time = parentHeader.Time + BLOCK_PERIOD + backOffTime(header, snap, c.chainConfig)
	if err := c.verifyHeaderTime(header, parentHeader, snap); err != nil {
		t.Errorf("Expect successful verification, got %s", err)
	}
}
