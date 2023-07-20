package v2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	consortiumCommon "github.com/ethereum/go-ethereum/consensus/consortium/common"
	"github.com/ethereum/go-ethereum/consensus/consortium/v2/finality"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

func TestSealableValidators(t *testing.T) {
	const NUM_OF_VALIDATORS = 21

	validators := make([]common.Address, NUM_OF_VALIDATORS)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		validators = append(validators, common.BigToAddress(big.NewInt(int64(i))))
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil, nil)
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

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil, nil)
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

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil, nil)
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

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil, nil)
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

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, validators, nil, nil)
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

func TestExtraDataEncode(t *testing.T) {
	extraData := finality.HeaderExtraData{}
	data := extraData.Encode(false)
	expectedLen := consortiumCommon.ExtraSeal + consortiumCommon.ExtraVanity
	if len(data) != expectedLen {
		t.Errorf(
			"Mismatch header extra data length before hardfork, have %v expect %v",
			len(data), expectedLen,
		)
	}

	extraData = finality.HeaderExtraData{
		CheckpointValidators: []finality.ValidatorWithBlsPub{
			{
				Address: common.Address{0x1},
			},
			{
				Address: common.Address{0x2},
			},
		},
	}
	expectedLen = consortiumCommon.ExtraSeal + consortiumCommon.ExtraVanity + common.AddressLength*2
	data = extraData.Encode(false)
	if len(data) != expectedLen {
		t.Errorf(
			"Mismatch header extra data length before hardfork, have %v expect %v",
			len(data), expectedLen,
		)
	}

	expectedLen = consortiumCommon.ExtraSeal + consortiumCommon.ExtraVanity + 1
	extraData = finality.HeaderExtraData{}
	data = extraData.Encode(true)
	if len(data) != expectedLen {
		t.Errorf(
			"Mismatch header extra data length before hardfork, have %v expect %v",
			len(data), expectedLen,
		)
	}

	secretKey, err := blst.RandKey()
	if err != nil {
		t.Fatalf("Failed to generate secret key, err %s", err)
	}
	dummyDigest := [32]byte{}
	signature := secretKey.Sign(dummyDigest[:])

	extraData = finality.HeaderExtraData{
		HasFinalityVote:         1,
		AggregatedFinalityVotes: signature,
	}
	expectedLen = consortiumCommon.ExtraSeal + consortiumCommon.ExtraVanity + 1 + 8 + params.BLSSignatureLength
	data = extraData.Encode(true)
	if len(data) != expectedLen {
		t.Errorf(
			"Mismatch header extra data length after hardfork, have %v expect %v",
			len(data), expectedLen,
		)
	}

	extraData = finality.HeaderExtraData{
		HasFinalityVote:         1,
		AggregatedFinalityVotes: signature,
		CheckpointValidators: []finality.ValidatorWithBlsPub{
			{
				Address:      common.Address{0x1},
				BlsPublicKey: secretKey.PublicKey(),
			},
			{
				Address:      common.Address{0x2},
				BlsPublicKey: secretKey.PublicKey(),
			},
		},
	}
	expectedLen = consortiumCommon.ExtraSeal + consortiumCommon.ExtraVanity + 1 + 8 + params.BLSSignatureLength + 2*(common.AddressLength+params.BLSPubkeyLength)
	data = extraData.Encode(true)
	if len(data) != expectedLen {
		t.Errorf(
			"Mismatch header extra data length after hardfork, have %v expect %v",
			len(data), expectedLen,
		)
	}
}

func TestExtraDataDecode(t *testing.T) {
	secretKey, err := blst.RandKey()
	if err != nil {
		t.Fatalf("Failed to generate secret key, err %s", err)
	}
	dummyDigest := [32]byte{}
	signature := secretKey.Sign(dummyDigest[:])

	rawBytes := []byte{'t', 'e', 's', 't'}
	_, err = finality.DecodeExtra(rawBytes, false)
	if !errors.Is(err, finality.ErrMissingVanity) {
		t.Errorf("Expect error %v have %v", finality.ErrMissingVanity, err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	_, err = finality.DecodeExtra(rawBytes, false)
	if !errors.Is(err, finality.ErrMissingSignature) {
		t.Errorf("Expect error %v have %v", finality.ErrMissingSignature, err)
	}

	rawBytes = append(rawBytes, byte(12))
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraSeal)...)
	_, err = finality.DecodeExtra(rawBytes, false)
	if !errors.Is(err, finality.ErrInvalidSpanValidators) {
		t.Errorf("Expect error %v have %v", finality.ErrInvalidSpanValidators, err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	_, err = finality.DecodeExtra(rawBytes, true)
	if !errors.Is(err, finality.ErrMissingHasFinalityVote) {
		t.Errorf("Expect error %v have %v", finality.ErrMissingHasFinalityVote, err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	rawBytes = append(rawBytes, byte(0x00))
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraSeal)...)
	_, err = finality.DecodeExtra(rawBytes, true)
	if err != nil {
		t.Errorf("Expect successful decode have %v", err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	rawBytes = append(rawBytes, byte(0x01))
	_, err = finality.DecodeExtra(rawBytes, true)
	if !errors.Is(err, finality.ErrMissingFinalityVoteBitSet) {
		t.Errorf("Expect error %v have %v", finality.ErrMissingFinalityVoteBitSet, err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	rawBytes = append(rawBytes, byte(0x01))
	rawBytes = binary.LittleEndian.AppendUint64(rawBytes, 0)
	_, err = finality.DecodeExtra(rawBytes, true)
	if !errors.Is(err, finality.ErrMissingFinalitySignature) {
		t.Errorf("Expect error %v have %v", finality.ErrMissingFinalitySignature, err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	rawBytes = append(rawBytes, byte(0x01))
	rawBytes = binary.LittleEndian.AppendUint64(rawBytes, 0)
	rawBytes = append(rawBytes, signature.Marshal()...)
	_, err = finality.DecodeExtra(rawBytes, true)
	if !errors.Is(err, finality.ErrMissingSignature) {
		t.Errorf("Expect error %v have %v", finality.ErrMissingSignature, err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	rawBytes = append(rawBytes, byte(0x01))
	rawBytes = binary.LittleEndian.AppendUint64(rawBytes, 0)
	rawBytes = append(rawBytes, signature.Marshal()...)
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraSeal)...)
	_, err = finality.DecodeExtra(rawBytes, true)
	if err != nil {
		t.Errorf("Expect successful decode have %v", err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	rawBytes = append(rawBytes, byte(0x01))
	rawBytes = binary.LittleEndian.AppendUint64(rawBytes, 0)
	rawBytes = append(rawBytes, signature.Marshal()...)
	rawBytes = append(rawBytes, common.Address{0x1}.Bytes()...)
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraSeal)...)
	_, err = finality.DecodeExtra(rawBytes, true)
	if !errors.Is(err, finality.ErrInvalidSpanValidators) {
		t.Errorf("Expect error %v have %v", finality.ErrInvalidSpanValidators, err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	rawBytes = append(rawBytes, byte(0x02))
	rawBytes = binary.LittleEndian.AppendUint64(rawBytes, 0)
	rawBytes = append(rawBytes, signature.Marshal()...)
	rawBytes = append(rawBytes, common.Address{0x1}.Bytes()...)
	rawBytes = append(rawBytes, secretKey.PublicKey().Marshal()...)
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraSeal)...)
	_, err = finality.DecodeExtra(rawBytes, true)
	if !errors.Is(err, finality.ErrInvalidHasFinalityVote) {
		t.Errorf("Expect error %v have %v", finality.ErrInvalidHasFinalityVote, err)
	}

	rawBytes = []byte{}
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraVanity)...)
	rawBytes = append(rawBytes, byte(0x01))
	rawBytes = binary.LittleEndian.AppendUint64(rawBytes, 0)
	rawBytes = append(rawBytes, signature.Marshal()...)
	rawBytes = append(rawBytes, common.Address{0x1}.Bytes()...)
	rawBytes = append(rawBytes, secretKey.PublicKey().Marshal()...)
	rawBytes = append(rawBytes, bytes.Repeat([]byte{0x00}, consortiumCommon.ExtraSeal)...)
	_, err = finality.DecodeExtra(rawBytes, true)
	if err != nil {
		t.Errorf("Expect successful decode have %v", err)
	}

	extraData := finality.HeaderExtraData{
		HasFinalityVote:         1,
		AggregatedFinalityVotes: signature,
		CheckpointValidators: []finality.ValidatorWithBlsPub{
			{
				Address:      common.Address{0x1},
				BlsPublicKey: secretKey.PublicKey(),
			},
			{
				Address:      common.Address{0x2},
				BlsPublicKey: secretKey.PublicKey(),
			},
		},
	}
	data := extraData.Encode(true)
	decodedData, err := finality.DecodeExtra(data, true)
	if err != nil {
		t.Errorf("Expect successful decode have %v", err)
	}

	// Do some sanity checks
	if !bytes.Equal(
		decodedData.AggregatedFinalityVotes.Marshal(),
		extraData.AggregatedFinalityVotes.Marshal(),
	) {
		t.Errorf("Mismatch decoded data")
	}

	if decodedData.CheckpointValidators[0].Address != extraData.CheckpointValidators[0].Address {
		t.Errorf("Mismatch decoded data")
	}

	if !decodedData.CheckpointValidators[0].BlsPublicKey.Equals(extraData.CheckpointValidators[0].BlsPublicKey) {
		t.Errorf("Mismatch decoded data")
	}
}

func TestVerifyFinalitySignature(t *testing.T) {
	const numValidator = 4
	var err error

	secretKey := make([]blsCommon.SecretKey, numValidator+1)
	for i := 0; i < len(secretKey); i++ {
		secretKey[i], err = blst.RandKey()
		if err != nil {
			t.Fatalf("Failed to generate secret key, err %s", err)
		}
	}

	valWithBlsPub := make([]finality.ValidatorWithBlsPub, numValidator)
	for i := 0; i < len(valWithBlsPub); i++ {
		valWithBlsPub[i] = finality.ValidatorWithBlsPub{
			common.BigToAddress(big.NewInt(int64(i))),
			secretKey[i].PublicKey(),
		}
	}

	blockNumber := uint64(0)
	blockHash := common.Hash{0x1}
	vote := types.VoteData{
		TargetNumber: blockNumber,
		TargetHash:   blockHash,
	}

	digest := vote.Hash()
	signature := make([]blsCommon.Signature, numValidator+1)
	for i := 0; i < len(signature); i++ {
		signature[i] = secretKey[i].Sign(digest[:])
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, nil, valWithBlsPub, nil)
	recents, _ := lru.NewARC(inmemorySnapshots)
	c := Consortium{
		chainConfig: &params.ChainConfig{
			ShillinBlock: big.NewInt(0),
		},
		config: &params.ConsortiumConfig{
			EpochV2: 300,
		},
		recents: recents,
	}
	snap.Hash = blockHash
	c.recents.Add(snap.Hash, snap)

	var votedBitSet finality.FinalityVoteBitSet
	votedBitSet.SetBit(0)
	err = c.verifyFinalitySignatures(nil, votedBitSet, nil, blockNumber, blockHash, nil)
	if !errors.Is(err, finality.ErrNotEnoughFinalityVote) {
		t.Errorf("Expect error %v have %v", finality.ErrNotEnoughFinalityVote, err)
	}

	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	votedBitSet.SetBit(4)
	err = c.verifyFinalitySignatures(nil, votedBitSet, nil, 0, snap.Hash, nil)
	if !errors.Is(err, finality.ErrInvalidFinalityVotedBitSet) {
		t.Errorf("Expect error %v have %v", finality.ErrInvalidFinalityVotedBitSet, err)
	}

	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	votedBitSet.SetBit(3)
	aggregatedSignature := blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
		signature[4],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if !errors.Is(err, finality.ErrFinalitySignatureVerificationFailed) {
		t.Errorf("Expect error %v have %v", finality.ErrFinalitySignatureVerificationFailed, err)
	}

	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	votedBitSet.SetBit(3)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
		signature[3],
		signature[4],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if !errors.Is(err, finality.ErrFinalitySignatureVerificationFailed) {
		t.Errorf("Expect error %v have %v", finality.ErrFinalitySignatureVerificationFailed, err)
	}

	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	votedBitSet.SetBit(3)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
		signature[3],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if err != nil {
		t.Errorf("Expect successful verification have %v", err)
	}
}
