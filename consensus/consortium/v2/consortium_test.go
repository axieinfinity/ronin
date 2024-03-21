package v2

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	consortiumCommon "github.com/ethereum/go-ethereum/consensus/consortium/common"
	"github.com/ethereum/go-ethereum/consensus/consortium/v2/finality"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
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

func mockExtraData(nVal int, bits uint32) *finality.HeaderExtraData {
	var (
		finalityVotedValidators finality.FinalityVoteBitSet
		aggregatedFinalityVotes blsCommon.Signature
		checkpointValidators    []finality.ValidatorWithBlsPub
		seal                    = make([]byte, finality.ExtraSeal)
		ret                     = &finality.HeaderExtraData{}
	)

	bits = bits % 32
	for i := 0; i < 5; i++ {
		if bits&(1<<i) != 0 {
			switch i {
			case 0:
				ret.HasFinalityVote = 1
				finalityVotedValidators = finality.FinalityVoteBitSet(uint64(8))
				ret.FinalityVotedValidators = finalityVotedValidators

				delegated, _ := blst.RandKey()
				msg := make([]byte, 64)
				rand.Read(msg)
				aggregatedFinalityVotes = delegated.Sign(msg)
				ret.AggregatedFinalityVotes = aggregatedFinalityVotes
			case 1:
				for i := 0; i < nVal; i++ {
					s, _ := blst.RandKey()
					sk, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
					addr := crypto.PubkeyToAddress(sk.PublicKey)
					val := finality.ValidatorWithBlsPub{
						Address:      addr,
						BlsPublicKey: s.PublicKey(),
					}
					checkpointValidators = append(checkpointValidators, val)
				}
				ret.CheckpointValidators = checkpointValidators
			case 2:
				// Even seal does not get assigned a random value, seal
				// is still be zero-filled byte array of size ExtraSeal
				rand.Read(seal)
				ret.Seal = [finality.ExtraSeal]byte(seal)
			}
		}
	}
	return ret
}

func TestExtraDataEncodeRLP(t *testing.T) {
	nVal := 22
	for i := 0; i < 7; i++ {
		ext := mockExtraData(nVal, uint32(i))
		enc, err := ext.EncodeRLP()
		if err != nil {
			t.Errorf("encode rlp error: %v", err)
		}
		if len(enc) < finality.ExtraSeal {
			t.Error("encode rlp error: invalid length of encoded data")
		}
	}

	var extraData finality.HeaderExtraData
	extraData.HasFinalityVote = 2
	_, err := extraData.EncodeRLP()
	if !errors.Is(err, finality.ErrInvalidHasFinalityVote) {
		t.Fatalf("Expect error: %s, got: %s", finality.ErrInvalidHasFinalityVote, err)
	}
}

func TestExtraDataDecodeRLP(t *testing.T) {
	nVals := 22
	// loop 64 times, equivalent to 64 combinations of 6 bits
	for i := 0; i < 7; i++ {
		ext := mockExtraData(nVals, uint32(i))
		enc, err := ext.EncodeRLP()
		if err != nil {
			t.Errorf("encode rlp error: %v", err)
		}
		dec, err := finality.DecodeExtraRLP(enc)
		if err != nil {
			t.Errorf("decode rlp error: %v", err)
		}
		if !bytes.Equal(dec.Vanity[:], ext.Vanity[:]) {
			t.Errorf("Mismatched decoded data")
		}
		if dec.FinalityVotedValidators != ext.FinalityVotedValidators {
			t.Errorf("Mismatch decoded data")
		}
		if (dec.AggregatedFinalityVotes != nil && ext.AggregatedFinalityVotes == nil) ||
			(dec.AggregatedFinalityVotes == nil && ext.AggregatedFinalityVotes != nil) {
			t.Errorf("Mismatch decoded data")
		}
		if dec.AggregatedFinalityVotes != nil &&
			ext.AggregatedFinalityVotes != nil &&
			!bytes.Equal(dec.AggregatedFinalityVotes.Marshal(), ext.AggregatedFinalityVotes.Marshal()) {
			t.Errorf("Mismatch decoded data")
		}
		if len(dec.CheckpointValidators) != len(ext.CheckpointValidators) {
			t.Errorf("Mismatch decoded data")
		}
		for i := 0; i < len(ext.CheckpointValidators); i++ {
			if dec.CheckpointValidators[i].Address.Hex() != ext.CheckpointValidators[i].Address.Hex() {
				t.Errorf("Mismatch decoded data")
			}
			if (dec.CheckpointValidators[i].BlsPublicKey == nil && ext.CheckpointValidators[i].BlsPublicKey != nil) ||
				(dec.CheckpointValidators[i].BlsPublicKey != nil && ext.CheckpointValidators[i].BlsPublicKey == nil) {
				t.Errorf("Mismatch decoded data")
			}
			if dec.CheckpointValidators[i].BlsPublicKey != nil &&
				ext.CheckpointValidators[i].BlsPublicKey != nil &&
				!bytes.Equal(dec.CheckpointValidators[i].BlsPublicKey.Marshal(), ext.CheckpointValidators[i].BlsPublicKey.Marshal()) {
				t.Errorf("Mismatch decoded data")
			}
		}
		if !bytes.Equal(dec.Seal[:], ext.Seal[:]) {
			t.Errorf("Mismatch decoded data")
		}
	}

	_, err := finality.DecodeExtraRLP([]byte{})
	if !errors.Is(err, finality.ErrInvalidEncodedExtraData) {
		t.Fatalf("Expect error: %s, got: %s", finality.ErrInvalidEncodedExtraData, err)
	}

	encodedData := [finality.ExtraSeal]byte{}
	_, err = finality.DecodeExtraRLP(encodedData[:])
	if !errors.Is(err, io.EOF) {
		t.Fatalf("Expect error: %s, got: %s", io.EOF, err)
	}
}

func BenchmarkEncodeRLP(b *testing.B) {
	nVal := 22
	ext := mockExtraData(nVal, 7)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ext.EncodeRLP()
	}
}

func BenchmarkEncode(b *testing.B) {
	nVal := 22
	ext := mockExtraData(nVal, 7)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ext.Encode(true)
	}
}

func BenchmarkDecodeRLP(b *testing.B) {
	nVal := 22
	ext := mockExtraData(nVal, 7)
	dec, _ := ext.EncodeRLP()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		finality.DecodeExtraRLP(dec)
	}
}

func BenchmarkDecode(b *testing.B) {
	nVal := 22
	ext := mockExtraData(nVal, 7)
	dec := ext.Encode(true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		finality.DecodeExtra(dec, true)
	}
}

func TestVerifyFinalitySignature(t *testing.T) {
	const numValidator = 3
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
			Address:      common.BigToAddress(big.NewInt(int64(i))),
			BlsPublicKey: secretKey[i].PublicKey(),
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
	votedBitSet.SetBit(3)
	err = c.verifyFinalitySignatures(nil, votedBitSet, nil, 0, snap.Hash, nil)
	if !errors.Is(err, finality.ErrInvalidFinalityVotedBitSet) {
		t.Errorf("Expect error %v have %v", finality.ErrInvalidFinalityVotedBitSet, err)
	}

	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature := blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[3],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if !errors.Is(err, finality.ErrFinalitySignatureVerificationFailed) {
		t.Errorf("Expect error %v have %v", finality.ErrFinalitySignatureVerificationFailed, err)
	}

	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
		signature[3],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if !errors.Is(err, finality.ErrFinalitySignatureVerificationFailed) {
		t.Errorf("Expect error %v have %v", finality.ErrFinalitySignatureVerificationFailed, err)
	}

	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if err != nil {
		t.Errorf("Expect successful verification have %v", err)
	}
}

func TestVerifyFinalitySignatureTripp(t *testing.T) {
	const numValidator = 3
	var err error

	secretKey := make([]blsCommon.SecretKey, numValidator)
	for i := 0; i < len(secretKey); i++ {
		secretKey[i], err = blst.RandKey()
		if err != nil {
			t.Fatalf("Failed to generate secret key, err %s", err)
		}
	}

	valWithBlsPub := make([]finality.ValidatorWithBlsPub, numValidator)
	for i := 0; i < len(valWithBlsPub); i++ {
		valWithBlsPub[i] = finality.ValidatorWithBlsPub{
			Address:      common.BigToAddress(big.NewInt(int64(i))),
			BlsPublicKey: secretKey[i].PublicKey(),
		}
	}

	blockNumber := uint64(0)
	blockHash := common.Hash{0x1}
	vote := types.VoteData{
		TargetNumber: blockNumber,
		TargetHash:   blockHash,
	}

	digest := vote.Hash()
	signature := make([]blsCommon.Signature, numValidator)
	for i := 0; i < len(signature); i++ {
		signature[i] = secretKey[i].Sign(digest[:])
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, nil, valWithBlsPub, nil)
	snap.FinalityVoteWeight = make([]uint16, numValidator)
	snap.FinalityVoteWeight[0] = 6666
	snap.FinalityVoteWeight[1] = 1
	snap.FinalityVoteWeight[2] = 3333

	recents, _ := lru.NewARC(inmemorySnapshots)
	c := Consortium{
		chainConfig: &params.ChainConfig{
			ShillinBlock: big.NewInt(0),
			TrippBlock:   big.NewInt(0),
		},
		config: &params.ConsortiumConfig{
			EpochV2: 300,
		},
		recents: recents,
	}
	snap.Hash = blockHash
	c.recents.Add(snap.Hash, snap)

	// 1 voter with vote weight 6666 does not reach the threshold
	votedBitSet := finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	aggregatedSignature := blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if !errors.Is(err, finality.ErrNotEnoughFinalityVote) {
		t.Errorf("Expect error %v have %v", finality.ErrNotEnoughFinalityVote, err)
	}

	// 2 voters with total vote weight 3333 + 1 does not reach the threshold
	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[1],
		signature[2],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if !errors.Is(err, finality.ErrNotEnoughFinalityVote) {
		t.Errorf("Expect error %v have %v", finality.ErrNotEnoughFinalityVote, err)
	}

	// 2 voters with total vote weight 6666 + 1 reach the threshold
	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if err != nil {
		t.Errorf("Expect successful verification have %v", err)
	}

	// All voters vote
	votedBitSet = finality.FinalityVoteBitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, 0, snap.Hash, nil)
	if err != nil {
		t.Errorf("Expect successful verification have %v", err)
	}
}

func TestSnapshotValidatorWithBlsKey(t *testing.T) {
	secretKey, err := blst.RandKey()
	if err != nil {
		t.Fatalf("Failed to generate secret key, err: %s", err)
	}

	validators := []finality.ValidatorWithBlsPub{
		{
			Address:      common.Address{0x1},
			BlsPublicKey: secretKey.PublicKey(),
		},
	}
	snap := newSnapshot(nil, nil, nil, 10, common.Hash{0x2}, nil, validators, nil)
	db := rawdb.NewMemoryDatabase()
	err = snap.store(db)
	if err != nil {
		t.Fatalf("Failed to store snapshot, err: %s", err)
	}

	savedSnap, err := loadSnapshot(nil, nil, db, common.Hash{0x2}, nil, nil)
	if err != nil {
		t.Fatalf("Failed to load snapshot, err: %s", err)
	}

	savedValidators := savedSnap.ValidatorsWithBlsPub
	if len(savedValidators) != len(validators) {
		t.Fatalf("Saved snapshot is corrupted")
	}

	for i := range validators {
		if validators[i].Address != savedValidators[i].Address {
			t.Fatalf("Saved snapshot is corrupted")
		}

		if !validators[i].BlsPublicKey.Equals(savedValidators[i].BlsPublicKey) {
			t.Fatalf("Saved snapshot is corrupted")
		}
	}
}

type mockContract struct {
	validators map[common.Address]blsCommon.PublicKey
}

func (contract *mockContract) WrapUpEpoch(opts *consortiumCommon.ApplyTransactOpts) error {
	if opts.ReceivedTxs != nil && len(*opts.ReceivedTxs) != 0 {
		*opts.ReceivedTxs = (*opts.ReceivedTxs)[1:]
	}
	return nil
}

func (contract *mockContract) SubmitBlockReward(opts *consortiumCommon.ApplyTransactOpts) error {
	if opts.ReceivedTxs != nil && len(*opts.ReceivedTxs) != 0 {
		*opts.ReceivedTxs = (*opts.ReceivedTxs)[1:]
	}
	return nil
}

func (contract *mockContract) Slash(opts *consortiumCommon.ApplyTransactOpts, spoiledValidator common.Address) error {
	if opts.ReceivedTxs != nil && len(*opts.ReceivedTxs) != 0 {
		*opts.ReceivedTxs = (*opts.ReceivedTxs)[1:]
	}
	return nil
}

func (contract *mockContract) FinalityReward(opts *consortiumCommon.ApplyTransactOpts, votedValidators []common.Address) error {
	if opts.ReceivedTxs != nil && len(*opts.ReceivedTxs) != 0 {
		*opts.ReceivedTxs = (*opts.ReceivedTxs)[1:]
	}
	return nil
}

func (contract *mockContract) GetValidators(*big.Int) ([]common.Address, error) {
	var validatorAddresses []common.Address
	for address := range contract.validators {
		validatorAddresses = append(validatorAddresses, address)
	}
	return validatorAddresses, nil
}

func (contract *mockContract) GetValidatorCandidates(blockNumber *big.Int) ([]common.Address, error) {
	return nil, nil
}

func (contract *mockContract) GetBlsPublicKey(_ *big.Int, address common.Address) (blsCommon.PublicKey, error) {
	if key, ok := contract.validators[address]; ok {
		if key != nil {
			return key, nil
		} else {
			return nil, errors.New("no BLS public key found")
		}
	} else {
		return nil, errors.New("address is not a validator")
	}
}

func (contract *mockContract) GetStakedAmount(_ *big.Int, _ []common.Address) ([]*big.Int, error) {
	return nil, nil
}

func TestGetCheckpointValidatorFromContract(t *testing.T) {
	var err error
	secretKeys := make([]blsCommon.SecretKey, 3)
	for i := 0; i < len(secretKeys); i++ {
		secretKeys[i], err = blst.RandKey()
		if err != nil {
			t.Fatalf("Failed to generate secret key, err: %s", err)
		}
	}

	mock := &mockContract{
		validators: map[common.Address]blsCommon.PublicKey{
			common.Address{0x1}: secretKeys[1].PublicKey(),
			common.Address{0x2}: nil,
			common.Address{0x5}: secretKeys[0].PublicKey(),
			common.Address{0x3}: secretKeys[2].PublicKey(),
		},
	}
	c := Consortium{
		chainConfig: &params.ChainConfig{
			ShillinBlock: big.NewInt(0),
		},
		contract: mock,
	}

	validatorWithPubs, err := c.getCheckpointValidatorsFromContract(&types.Header{Number: big.NewInt(3)})
	if err != nil {
		t.Fatalf("Failed to get checkpoint validators from contract, err: %s", err)
	}

	if len(validatorWithPubs) != 3 {
		t.Fatalf("Expect returned list, length: %d have: %d", 3, len(validatorWithPubs))
	}
	if validatorWithPubs[0].Address != (common.Address{0x1}) {
		t.Fatalf("Wrong returned list")
	}
	if !validatorWithPubs[0].BlsPublicKey.Equals(secretKeys[1].PublicKey()) {
		t.Fatalf("Wrong returned list")
	}
	if validatorWithPubs[1].Address != (common.Address{0x3}) {
		t.Fatalf("Wrong returned list")
	}
	if !validatorWithPubs[1].BlsPublicKey.Equals(secretKeys[2].PublicKey()) {
		t.Fatalf("Wrong returned list")
	}
	if validatorWithPubs[2].Address != (common.Address{0x5}) {
		t.Fatalf("Wrong returned list")
	}
	if !validatorWithPubs[2].BlsPublicKey.Equals(secretKeys[0].PublicKey()) {
		t.Fatalf("Wrong returned list")
	}
}

type mockVotePool struct {
	vote []*types.VoteEnvelope
}

func (votePool *mockVotePool) FetchVoteByBlockHash(hash common.Hash) []*types.VoteEnvelope {
	return votePool.vote
}

func TestAssembleFinalityVote(t *testing.T) {
	var err error
	secretKeys := make([]blsCommon.SecretKey, 10)
	for i := 0; i < len(secretKeys); i++ {
		secretKeys[i], err = blst.RandKey()
		if err != nil {
			t.Fatalf("Failed to generate secret key, err: %s", err)
		}
	}

	voteData := types.VoteData{
		TargetNumber: 4,
		TargetHash:   common.Hash{0x1},
	}
	digest := voteData.Hash()

	signatures := make([]blsCommon.Signature, 10)
	for i := 0; i < len(signatures); i++ {
		signatures[i] = secretKeys[i].Sign(digest[:])
	}

	var votes []*types.VoteEnvelope
	for i := 0; i < 10; i++ {
		votes = append(votes, &types.VoteEnvelope{
			RawVoteEnvelope: types.RawVoteEnvelope{
				PublicKey: types.BLSPublicKey(secretKeys[i].PublicKey().Marshal()),
				Signature: types.BLSSignature(signatures[i].Marshal()),
				Data:      &voteData,
			},
		})
	}

	mock := mockVotePool{
		vote: votes,
	}
	c := Consortium{
		chainConfig: &params.ChainConfig{
			ShillinBlock: big.NewInt(0),
		},
		votePool: &mock,
	}

	var validators []finality.ValidatorWithBlsPub
	for i := 0; i < 9; i++ {
		validators = append(validators, finality.ValidatorWithBlsPub{
			Address:      common.BigToAddress(big.NewInt(int64(i))),
			BlsPublicKey: secretKeys[i].PublicKey(),
		})
	}

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, nil, validators, nil)

	header := types.Header{Number: big.NewInt(5)}
	extraData := &finality.HeaderExtraData{}
	header.Extra = extraData.Encode(true)
	c.assembleFinalityVote(&header, snap)

	extraData, err = finality.DecodeExtra(header.Extra, true)
	if err != nil {
		t.Fatalf("Failed to decode extra data, err: %s", err)
	}

	if extraData.HasFinalityVote != 1 {
		t.Fatal("Missing finality vote in header")
	}

	bitSet := finality.FinalityVoteBitSet(0)
	for i := 0; i < 9; i++ {
		bitSet.SetBit(i)
	}

	if uint64(bitSet) != uint64(extraData.FinalityVotedValidators) {
		t.Fatalf(
			"Mismatch voted validator, expect %d have %d",
			uint64(bitSet),
			uint64(extraData.FinalityVotedValidators),
		)
	}

	var includedSignatures []blsCommon.Signature
	for i := 0; i < 9; i++ {
		includedSignatures = append(includedSignatures, signatures[i])
	}

	aggregatedSignature := blst.AggregateSignatures(includedSignatures)

	if !bytes.Equal(aggregatedSignature.Marshal(), extraData.AggregatedFinalityVotes.Marshal()) {
		t.Fatal("Mismatch signature")
	}
}

// TODO: Add AssembleFinalityVoteTripp test
func TestAssembleFinalityVoteTripp(t *testing.T) {

}

func TestVerifyVote(t *testing.T) {
	const numValidator = 3
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
			Address:      common.BigToAddress(big.NewInt(int64(i))),
			BlsPublicKey: secretKey[i].PublicKey(),
		}
	}

	db := rawdb.NewMemoryDatabase()
	genesis := (&core.Genesis{
		Config:  params.TestChainConfig,
		BaseFee: big.NewInt(params.InitialBaseFee),
	}).MustCommit(db)
	chain, _ := core.NewBlockChain(db, nil, params.TestChainConfig, ethash.NewFullFaker(), vm.Config{}, nil, nil)

	bs, _ := core.GenerateChain(params.TestChainConfig, genesis, ethash.NewFaker(), db, 1, nil, true)
	if _, err := chain.InsertChain(bs[:]); err != nil {
		panic(err)
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
	snap.Hash = bs[0].Hash()
	c.recents.Add(snap.Hash, snap)

	// invalid vote number
	voteData := types.VoteData{
		TargetNumber: 2,
		TargetHash:   bs[0].Hash(),
	}
	signature := secretKey[0].Sign(voteData.Hash().Bytes())

	vote := types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			PublicKey: types.BLSPublicKey(secretKey[0].PublicKey().Marshal()),
			Signature: types.BLSSignature(signature.Marshal()),
			Data:      &voteData,
		},
	}

	err = c.VerifyVote(chain, &vote)
	if !errors.Is(err, finality.ErrInvalidTargetNumber) {
		t.Errorf("Expect error %v have %v", finality.ErrInvalidTargetNumber, err)
	}

	// invalid public key
	voteData = types.VoteData{
		TargetNumber: 1,
		TargetHash:   bs[0].Hash(),
	}
	signature = secretKey[numValidator].Sign(voteData.Hash().Bytes())

	vote = types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			PublicKey: types.BLSPublicKey(secretKey[numValidator].PublicKey().Marshal()),
			Signature: types.BLSSignature(signature.Marshal()),
			Data:      &voteData,
		},
	}

	err = c.VerifyVote(chain, &vote)
	if !errors.Is(err, finality.ErrUnauthorizedFinalityVoter) {
		t.Errorf("Expect error %v have %v", finality.ErrUnauthorizedFinalityVoter, err)
	}

	// sucessful case
	voteData = types.VoteData{
		TargetNumber: 1,
		TargetHash:   bs[0].Hash(),
	}
	signature = secretKey[0].Sign(voteData.Hash().Bytes())

	vote = types.VoteEnvelope{
		RawVoteEnvelope: types.RawVoteEnvelope{
			PublicKey: types.BLSPublicKey(secretKey[0].PublicKey().Marshal()),
			Signature: types.BLSSignature(signature.Marshal()),
			Data:      &voteData,
		},
	}

	err = c.VerifyVote(chain, &vote)
	if err != nil {
		t.Errorf("Expect sucessful verification have %s", err)
	}
}

func TestKnownBlockReorg(t *testing.T) {
	db := rawdb.NewMemoryDatabase()

	blsKeys := make([]blsCommon.SecretKey, 3)
	ecdsaKeys := make([]*ecdsa.PrivateKey, 3)
	validatorAddrs := make([]common.Address, 3)

	for i := range blsKeys {
		blsKey, err := blst.RandKey()
		if err != nil {
			t.Fatal(err)
		}
		blsKeys[i] = blsKey

		secretKey, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		ecdsaKeys[i] = secretKey
		validatorAddrs[i] = crypto.PubkeyToAddress(secretKey.PublicKey)
	}

	for i := 0; i < len(blsKeys)-1; i++ {
		for j := i; j < len(blsKeys); j++ {
			if bytes.Compare(validatorAddrs[i][:], validatorAddrs[j][:]) > 0 {
				validatorAddrs[i], validatorAddrs[j] = validatorAddrs[j], validatorAddrs[i]
				blsKeys[i], blsKeys[j] = blsKeys[j], blsKeys[i]
				ecdsaKeys[i], ecdsaKeys[j] = ecdsaKeys[j], ecdsaKeys[i]
			}
		}
	}

	chainConfig := params.ChainConfig{
		ChainID:           big.NewInt(2021),
		HomesteadBlock:    common.Big0,
		EIP150Block:       common.Big0,
		EIP155Block:       common.Big0,
		EIP158Block:       common.Big0,
		ConsortiumV2Block: common.Big0,
		ShillinBlock:      big.NewInt(10),
		Consortium: &params.ConsortiumConfig{
			EpochV2: 10,
		},
	}

	genesis := (&core.Genesis{
		Config: &chainConfig,
	}).MustCommit(db)

	mock := &mockContract{
		validators: make(map[common.Address]blsCommon.PublicKey),
	}
	mock.validators[validatorAddrs[0]] = blsKeys[0].PublicKey()
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	v2 := Consortium{
		chainConfig: &chainConfig,
		contract:    mock,
		recents:     recents,
		signatures:  signatures,
		config:      chainConfig.Consortium,
		db:          db,
	}

	chain, _ := core.NewBlockChain(db, nil, &chainConfig, &v2, vm.Config{}, nil, nil)
	extraData := [consortiumCommon.ExtraVanity + consortiumCommon.ExtraSeal]byte{}

	blocks, _ := core.GenerateConsortiumChain(
		&chainConfig,
		genesis,
		&v2,
		db,
		9,
		func(i int, bg *core.BlockGen) {
			bg.SetCoinbase(validatorAddrs[0])
			bg.SetExtra(extraData[:])
			bg.SetDifficulty(big.NewInt(7))
		},
		true,
		func(i int, bg *core.BlockGen) {
			header := bg.Header()
			hash := calculateSealHash(header, big.NewInt(2021))
			sig, err := crypto.Sign(hash[:], ecdsaKeys[0])
			if err != nil {
				t.Fatalf("Failed to sign block, err %s", err)
			}
			copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
			bg.SetExtra(header.Extra)
		},
	)

	_, err := chain.InsertChain(blocks)
	if err != nil {
		t.Fatalf("Failed to insert block, err %s", err)
	}

	for i := range validatorAddrs {
		mock.validators[validatorAddrs[i]] = blsKeys[i].PublicKey()
	}

	var checkpointValidators []finality.ValidatorWithBlsPub
	for i := range validatorAddrs {
		checkpointValidators = append(checkpointValidators, finality.ValidatorWithBlsPub{
			Address:      validatorAddrs[i],
			BlsPublicKey: blsKeys[i].PublicKey(),
		})
	}

	// Prepare checkpoint block
	blocks, _ = core.GenerateConsortiumChain(
		&chainConfig,
		blocks[len(blocks)-1],
		&v2,
		db,
		1,
		func(i int, bg *core.BlockGen) {
			var extra finality.HeaderExtraData

			bg.SetCoinbase(validatorAddrs[0])
			bg.SetDifficulty(big.NewInt(7))
			extra.CheckpointValidators = checkpointValidators
			bg.SetExtra(extra.Encode(true))
		},
		true,
		func(i int, bg *core.BlockGen) {
			header := bg.Header()
			hash := calculateSealHash(header, big.NewInt(2021))
			sig, err := crypto.Sign(hash[:], ecdsaKeys[0])
			if err != nil {
				t.Fatalf("Failed to sign block, err %s", err)
			}
			copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
			bg.SetExtra(header.Extra)
		},
	)

	_, err = chain.InsertChain(blocks)
	if err != nil {
		t.Fatalf("Failed to insert block, err %s", err)
	}

	extraDataShillin := [consortiumCommon.ExtraVanity + 1 + consortiumCommon.ExtraSeal]byte{}
	knownBlocks, _ := core.GenerateConsortiumChain(
		&chainConfig,
		blocks[len(blocks)-1],
		&v2,
		db,
		1,
		func(i int, bg *core.BlockGen) {
			bg.SetCoinbase(validatorAddrs[2])
			bg.SetExtra(extraDataShillin[:])
			bg.SetDifficulty(big.NewInt(7))
		},
		true,
		func(i int, bg *core.BlockGen) {
			header := bg.Header()
			hash := calculateSealHash(header, big.NewInt(2021))
			sig, err := crypto.Sign(hash[:], ecdsaKeys[2])
			if err != nil {
				t.Fatalf("Failed to sign block, err %s", err)
			}
			copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
			bg.SetExtra(header.Extra)
		},
	)

	_, err = chain.InsertChain(knownBlocks)
	if err != nil {
		t.Fatalf("Failed to insert block, err %s", err)
	}

	header := chain.CurrentHeader()
	if header.Number.Uint64() != 11 {
		t.Fatalf("Expect head header to be %d, got %d", 11, header.Number.Uint64())
	}
	if header.Difficulty.Cmp(big.NewInt(7)) != 0 {
		t.Fatalf("Expect header header to have difficulty %d, got %d", 7, header.Difficulty.Uint64())
	}

	justifiedBlocks, _ := core.GenerateConsortiumChain(
		&chainConfig,
		blocks[len(blocks)-1],
		&v2,
		db,
		2,
		func(i int, bg *core.BlockGen) {
			if bg.Number().Uint64() == 11 {
				bg.SetCoinbase(validatorAddrs[1])
				bg.SetExtra(extraDataShillin[:])
			} else {
				bg.SetCoinbase(validatorAddrs[2])

				var (
					extra      finality.HeaderExtraData
					voteBitset finality.FinalityVoteBitSet
					signatures []blsCommon.Signature
				)
				voteBitset.SetBit(0)
				voteBitset.SetBit(1)
				voteBitset.SetBit(2)
				extra.HasFinalityVote = 1
				extra.FinalityVotedValidators = voteBitset

				block := bg.PrevBlock(-1)
				voteData := types.VoteData{
					TargetNumber: block.NumberU64(),
					TargetHash:   block.Hash(),
				}
				for i := range blsKeys {
					signatures = append(signatures, blsKeys[i].Sign(voteData.Hash().Bytes()))
				}

				extra.AggregatedFinalityVotes = blst.AggregateSignatures(signatures)
				bg.SetExtra(extra.Encode(true))
			}

			bg.SetDifficulty(big.NewInt(3))
		},
		true,
		func(i int, bg *core.BlockGen) {
			header := bg.Header()
			hash := calculateSealHash(header, big.NewInt(2021))

			var ecdsaKey *ecdsa.PrivateKey
			if bg.Number().Uint64() == 11 {
				ecdsaKey = ecdsaKeys[1]
			} else {
				ecdsaKey = ecdsaKeys[2]
			}
			sig, err := crypto.Sign(hash[:], ecdsaKey)
			if err != nil {
				t.Fatalf("Failed to sign block, err %s", err)
			}
			copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
			bg.SetExtra(header.Extra)
		},
	)

	_, err = chain.InsertChain(justifiedBlocks)
	if err != nil {
		t.Fatalf("Failed to insert block, err %s", err)
	}

	header = chain.CurrentHeader()
	if header.Number.Uint64() != 12 {
		t.Fatalf("Expect head header to be %d, got %d", 12, header.Number.Uint64())
	}

	_, err = chain.InsertChain(knownBlocks)
	if err != nil {
		t.Fatalf("Failed to insert block, err %s", err)
	}
	header = chain.CurrentHeader()
	if header.Number.Uint64() != 12 {
		t.Fatalf("Expect head header to be %d, got %d", 12, header.Number.Uint64())
	}
	header = chain.GetHeaderByNumber(11)
	if header.Difficulty.Uint64() != 3 {
		t.Fatalf("Expect head header to have difficulty %d, got %d", 3, header.Difficulty.Uint64())
	}
}

func TestUpgradeRoninTrustedOrg(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	blsSecretKey, err := blst.RandKey()
	if err != nil {
		t.Fatal(err)
	}
	secretKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	validatorAddr := crypto.PubkeyToAddress(secretKey.PublicKey)

	chainConfig := params.ChainConfig{
		ChainID:           big.NewInt(2021),
		HomesteadBlock:    common.Big0,
		EIP150Block:       common.Big0,
		EIP155Block:       common.Big0,
		EIP158Block:       common.Big0,
		ConsortiumV2Block: common.Big0,
		MikoBlock:         common.Big3,
		Consortium: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		RoninTrustedOrgUpgrade: &params.ContractUpgrade{
			ProxyAddress:          common.Address{0x10},
			ImplementationAddress: common.Address{0x20},
		},
	}

	genesis := (&core.Genesis{
		Config: &chainConfig,
		Alloc: core.GenesisAlloc{
			// Make proxy address non-empty to avoid being deleted
			common.Address{0x10}: core.GenesisAccount{Balance: common.Big1},
		},
	}).MustCommit(db)

	mock := &mockContract{
		validators: map[common.Address]blsCommon.PublicKey{
			validatorAddr: blsSecretKey.PublicKey(),
		},
	}
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	v2 := Consortium{
		chainConfig: &chainConfig,
		contract:    mock,
		recents:     recents,
		signatures:  signatures,
		config: &params.ConsortiumConfig{
			EpochV2: 200,
		},
	}

	chain, _ := core.NewBlockChain(db, nil, &chainConfig, &v2, vm.Config{}, nil, nil)
	extraData := [consortiumCommon.ExtraVanity + consortiumCommon.ExtraSeal]byte{}

	parent := genesis
	for i := 0; i < 5; i++ {
		block, _ := core.GenerateChain(
			&chainConfig,
			parent,
			&v2,
			db,
			1,
			func(i int, bg *core.BlockGen) {
				bg.SetCoinbase(validatorAddr)
				bg.SetExtra(extraData[:])
				bg.SetDifficulty(big.NewInt(7))
			},
			true,
		)

		header := block[0].Header()
		hash := calculateSealHash(header, big.NewInt(2021))
		sig, err := crypto.Sign(hash[:], secretKey)
		if err != nil {
			t.Fatalf("Failed to sign block, err %s", err)
		}

		copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
		block[0] = block[0].WithSeal(header)
		parent = block[0]

		if i == int(chainConfig.MikoBlock.Int64()-1) {
			statedb, err := chain.State()
			if err != nil {
				t.Fatalf("Failed to get statedb, err %s", err)
			}

			implementationAddr := statedb.GetState(v2.chainConfig.RoninTrustedOrgUpgrade.ProxyAddress, implementationSlot)
			if implementationAddr != (common.Hash{}) {
				t.Fatalf(
					"Implementation slot mismatches, exp: {%x} got {%x}",
					common.Hash{},
					implementationAddr,
				)
			}
		}

		_, err = chain.InsertChain(block)
		if err != nil {
			t.Fatalf("Failed to insert chain, err %s", err)
		}

		if i == int(chainConfig.MikoBlock.Int64()-1) {
			statedb, err := chain.State()
			if err != nil {
				t.Fatalf("Failed to get statedb, err %s", err)
			}

			implementationAddr := statedb.GetState(v2.chainConfig.RoninTrustedOrgUpgrade.ProxyAddress, implementationSlot)
			if implementationAddr != v2.chainConfig.RoninTrustedOrgUpgrade.ImplementationAddress.Hash() {
				t.Fatalf(
					"Implementation slot mismatches, exp: {%x} got {%x}",
					v2.chainConfig.RoninTrustedOrgUpgrade.ImplementationAddress.Hash(),
					implementationAddr,
				)
			}
		}
	}
}

func TestSystemTransactionOrder(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	blsSecretKey, err := blst.RandKey()
	if err != nil {
		t.Fatal(err)
	}
	secretKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	validatorAddr := crypto.PubkeyToAddress(secretKey.PublicKey)

	userKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	chainConfig := params.ChainConfig{
		ChainID:           big.NewInt(2021),
		HomesteadBlock:    common.Big0,
		EIP150Block:       common.Big0,
		EIP155Block:       common.Big0,
		EIP158Block:       common.Big0,
		ConsortiumV2Block: common.Big0,
		MikoBlock:         common.Big0,
		Consortium: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		ConsortiumV2Contracts: &params.ConsortiumV2Contracts{
			RoninValidatorSet: common.HexToAddress("0xaa"),
		},
	}

	genesis := (&core.Genesis{
		Config: &chainConfig,
		Alloc: core.GenesisAlloc{
			// Make proxy address non-empty to avoid being deleted
			common.Address{0x10}: core.GenesisAccount{Balance: common.Big1},
		},
	}).MustCommit(db)

	mock := &mockContract{
		validators: map[common.Address]blsCommon.PublicKey{
			validatorAddr: blsSecretKey.PublicKey(),
		},
	}
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	v2 := Consortium{
		chainConfig: &chainConfig,
		contract:    mock,
		recents:     recents,
		signatures:  signatures,
		config: &params.ConsortiumConfig{
			EpochV2: 200,
		},
	}

	chain, _ := core.NewBlockChain(db, nil, &chainConfig, &v2, vm.Config{}, nil, nil)
	extraData := [consortiumCommon.ExtraVanity + consortiumCommon.ExtraSeal]byte{}

	signer := types.NewEIP155Signer(big.NewInt(2021))
	normalTx, err := types.SignTx(
		types.NewTransaction(
			0,
			common.Address{},
			new(big.Int),
			21000,
			new(big.Int),
			nil,
		),
		signer,
		userKey,
	)
	if err != nil {
		t.Fatalf("Failed to sign transaction, err %s", err)
	}

	systemTx, err := types.SignTx(
		types.NewTransaction(
			0,
			chainConfig.ConsortiumV2Contracts.RoninValidatorSet,
			new(big.Int),
			21000,
			new(big.Int),
			nil,
		),
		signer,
		secretKey,
	)
	if err != nil {
		t.Fatalf("Failed to sign transaction, err %s", err)
	}

	blocks, receipts := core.GenerateConsortiumChain(
		&chainConfig,
		genesis,
		&v2,
		db,
		1,
		func(i int, bg *core.BlockGen) {
			bg.SetCoinbase(validatorAddr)
			bg.SetExtra(extraData[:])
			bg.SetDifficulty(big.NewInt(7))
			bg.AddTx(normalTx)
		},
		true,
		func(i int, bg *core.BlockGen) {
			header := bg.Header()
			hash := calculateSealHash(header, big.NewInt(2021))
			sig, err := crypto.Sign(hash[:], secretKey)
			if err != nil {
				t.Fatalf("Failed to sign block, err %s", err)
			}
			copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
			bg.SetExtra(header.Extra)
		},
	)

	// Mock contract does not create system transaction so right now len(block.transactions) == 1.
	// Add the system transaction before normal transaction.
	block := types.NewBlock(blocks[0].Header(), []*types.Transaction{systemTx, normalTx}, nil, receipts[0], trie.NewStackTrie(nil))
	header := block.Header()
	hash := calculateSealHash(header, big.NewInt(2021))
	sig, err := crypto.Sign(hash[:], secretKey)
	if err != nil {
		t.Fatalf("Failed to sign block, err %s", err)
	}
	copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
	block = types.NewBlockWithHeader(header)
	block = types.NewBlock(block.Header(), []*types.Transaction{systemTx, normalTx}, nil, receipts[0], trie.NewStackTrie(nil))

	_, err = chain.InsertChain(types.Blocks{block})
	if !errors.Is(err, core.ErrOutOfOrderSystemTx) {
		t.Fatalf("Expected err: %s, got %s", core.ErrOutOfOrderSystemTx, err)
	}
}

func TestIsPeriodBlock(t *testing.T) {
	EpochV2 := uint64(200)
	db := rawdb.NewMemoryDatabase()
	genesis := (&core.Genesis{
		Config:  params.TestChainConfig,
		BaseFee: big.NewInt(params.InitialBaseFee),
	}).MustCommit(db)
	chain, _ := core.NewBlockChain(db, nil, params.TestChainConfig, ethash.NewFullFaker(), vm.Config{}, nil, nil)

	bs, _ := core.GenerateChain(params.TestChainConfig, genesis, ethash.NewFaker(), db, 10, nil, true)
	if _, err := chain.InsertChain(bs[:]); err != nil {
		panic(err)
	}

	header := &types.Header{Number: new(big.Int).SetUint64(5)}
	if IsPeriodBlock(chain, header, EpochV2) {
		t.Errorf("wrong period block")
	}
	header.Number = new(big.Int).SetUint64(0)
	if IsPeriodBlock(chain, header, EpochV2) {
		t.Errorf("wrong period block")
	}

	dateInSeconds := uint64(86400)
	now := uint64(time.Now().Unix())
	time := now - now%dateInSeconds
	ancient := &types.Header{
		Number: new(big.Int).SetUint64(1000),
		Time:   uint64(time - 1),
	}
	header = &types.Header{
		Number: new(big.Int).SetUint64(1200),
		Time:   uint64(time + 2),
	}
	if header.Time/dateInSeconds-1 != ancient.Time/dateInSeconds {
		t.Errorf("wrong period block logic")
	}

	ancient.Time = uint64(time+1)
	if header.Time/dateInSeconds != ancient.Time/dateInSeconds {
		t.Errorf("wrong period block logic")
	}
}
