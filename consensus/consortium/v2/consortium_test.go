package v2

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	mrand "math/rand"
	"sort"
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
	"github.com/hashicorp/golang-lru/arc/v2"
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
		finalityVotedValidators finality.BitSet
		aggregatedFinalityVotes blsCommon.Signature
		checkpointValidators    []finality.ValidatorWithBlsPub
		seal                    = make([]byte, finality.ExtraSeal)
		ret                     = &finality.HeaderExtraData{}
	)

	bits = bits % 64
	for i := 0; i < 6; i++ {
		if bits&(1<<i) != 0 {
			switch i {
			case 0:
				ret.HasFinalityVote = 1
				finalityVotedValidators = finality.BitSet(uint64(8))
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

			// cases 3,4 are used to test RLP encoding in both Shillin and Tripp blocks
			// as before Tripp, StakedAmount and BlockProducers are empty.
			case 3:
				for i := range ret.CheckpointValidators {
					ret.CheckpointValidators[i].Weight = uint16(333)
				}
			case 4:
				ret.BlockProducers = []common.Address{
					common.Address{0x11},
					common.Address{0x22},
					common.Address{0x33},
				}
			case 5:
				ret.BlockProducersBitSet = finality.BitSet(mrand.Uint64())
			}
		}
	}
	return ret
}

func TestExtraDataEncodeRLP(t *testing.T) {
	nVal := 22
	for i := 0; i < 32; i++ {
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
	for i := 0; i < 32; i++ {
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
				!dec.CheckpointValidators[i].BlsPublicKey.Equals(ext.CheckpointValidators[i].BlsPublicKey) {
				t.Errorf("Mismatch decoded data")
			}
			if dec.CheckpointValidators[i].Weight != ext.CheckpointValidators[i].Weight {
				t.Error("Mismatch decoded data")
			}
		}
		if len(dec.BlockProducers) != len(ext.BlockProducers) {
			t.Errorf("Mismatch decoded data")
		}
		for i := 0; i < len(ext.BlockProducers); i++ {
			if dec.BlockProducers[i].Hex() != ext.BlockProducers[i].Hex() {
				t.Errorf("Mismatch decoded data")
			}
		}
		if !bytes.Equal(dec.Seal[:], ext.Seal[:]) {
			t.Errorf("Mismatch decoded data")
		}
		if dec.BlockProducersBitSet != ext.BlockProducersBitSet {
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

	// test case extra encoding with rlp optional
	enc := common.Hex2Bytes("f905fd821fdfb860b84d08d4923f9833f1217bdaab39ae7210200ea854429e2a0324a278639fa5e61cfe81ad05d2058cb2256ecb674fbc6b0a78772372afc2361e7476e83f2829cf4a312bb9a44b12b8539df408d15448b5feb06ede96871e6ce8147a9f590f03fbf904acf84994847c2b1f0138e82c0e12c23d9b1f58bffbe8e43bb0afa23456ca3bd535b4308ec0496eb98f60d0ecf332263cf92337c56f7468cf824a5f3078cf9d727a483754e546a40b958202b1f849949f1abc67bea4db5560371ff3089f4bfe934c36bcb0b5902528277a835bd7d779d43aacbedaa5266f7cbbebe783950c5832ccdad1c98e00408276677711e57814da4bb10091820344f84994e9bf2a788c27dadc6b169d52408b710d267b9bffb08aeb082b66e80ee32e3f29a787ed41914f9ddd37041108fbc207801dfb7207628ecdb2524d8f8f1e0863dfdf19e93b888202b3f84994a85ddddceeab43dccaa259dd4936ac104386f9aab09256ab3792329b85dc7b633a3f7f99d8f84a8924a27576d89323988f09871deaeb82a18248cd02af3e7837c91d38b62982033ff84994d086d2e3fac052a3f695a4e8905ce1722531163cb0a548cc15b37218e9b402465ccfc4f7f9bff8c9bc552286287c261c6818f2e811653bafb9fb921b0e3cb1883261a60a708202fff84894f4ff69528abdf88ec509fdd950497f54d7890bc8b091850ea30f0d24e5b458d72820d5b2f1b0c9a88239f990779231531ccee857ea20d96da1ff6c60bf3ee65280ca21547681edf84994a325fd3a2f4f5cafe2c151ee428b5ceded628193b099210a94511cacc37ba3c972618ef8f805dcdb484a09e2d3660c3a468b39021b21093d2ad62244d698de69ebaf951c79820320f849949422d990acdc3f2b3aa3b97303ad3060f09d7ffcb0ab948584bc2b98314144e78361b684fe1b14fa05cf38fbb549e988522ca1ab97dc9593460c2d2b2afd0a2a31f56557ff82038df84994c3c97512421bf3e339e9fd412f18584e53138bfab0811c1cb10f63e1ad0fd0a8e5ab9a535f78ad8b12de3761ac7e70df6ca9768ee046b0b78bf722c700ab61e984a622c40f82038df84994a49541ee1bbdd6aca7aadcdf5b591cce0f460795b0a54c52c9ac032fa63d2c5a892ccb28f94ce23a1af9044f2e5086bef79e62e778c81d8a6d9ebd78c6dcef021fda31bcd782017bf84994614381cbb8afebd58a55937e6a47b678adb0c2c6b084704f4c348684cf80be54a63c557bc0b63b9bfe8d1351f279f8d1347fc507121d6cf8bdfd48a2925508fa826b37c64a8202b1f849948b5608c77cb1309f2e06a3473bf4bf43aff5144cb0b30585ee4a72a987dfed0e3b7390b52d4351120d40f8566b1dc59117987acbea48e165510209c0ec25d12b1a316a05ce8202b1f84994726f02987863f4aeeacc94a81ef6755a58ed676bb0966f03c7be65bc0e771901857f517d67f11615090030cab29a4dfdcf440b31819e7e0aa763be68028b60afa082614a3a820179f84894786e3c84a3a8ccc38a6994d4fa7f37219ba6a98bb0a8e75014e6c7b4d09b3af3aa73fdd9f8ede55d916b535893584aa9b2c29d1a2a3886424e0ee429e6e414a0518b71e3c981edf84894bb79150e2774dc627869f750aa59246d4d8f3a63b0b8eea2bb0567c225d0dec7abe2d2e49a86a15099e034b5c46d43cdb723cc17f0d93e04a5cfc34597eccfd4f6655cb56481edf84894d0fa4a759b94aae2767d7bfcbbfe739b6d6f20c6b0a8d544a46f35384348128403a76cca5c437c20bc922766918017993e10afa439b4bd049c28bc1a8d415f27c253e371f181edf8e794726f02987863f4aeeacc94a81ef6755a58ed676b94847c2b1f0138e82c0e12c23d9b1f58bffbe8e43b948b5608c77cb1309f2e06a3473bf4bf43aff5144c949422d990acdc3f2b3aa3b97303ad3060f09d7ffc949f1abc67bea4db5560371ff3089f4bfe934c36bc94a325fd3a2f4f5cafe2c151ee428b5ceded62819394a49541ee1bbdd6aca7aadcdf5b591cce0f46079594a85ddddceeab43dccaa259dd4936ac104386f9aa94c3c97512421bf3e339e9fd412f18584e53138bfa94d086d2e3fac052a3f695a4e8905ce1722531163c94e9bf2a788c27dadc6b169d52408b710d267b9bffa81f644242489cf62023766d1e0768d0471740d034b37d4d36c77efb36e1e6f576fbdfe968389adb3b05e5edaf0f4ea070649102bd0cb6dbb784f382bbe84b9600")
	_, err = finality.DecodeExtraRLP(enc[:])
	if err != nil {
		t.Fatal(err)
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
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
	c := Consortium{
		chainConfig: &params.ChainConfig{
			ShillinBlock: big.NewInt(0),
		},
		config: &params.ConsortiumConfig{
			EpochV2: 300,
		},
		recents:            recents,
		testTrippEffective: true,
	}
	snap.Hash = blockHash
	c.recents.Add(snap.Hash, snap)

	header := types.Header{Number: big.NewInt(int64(blockNumber + 1)), ParentHash: blockHash}
	var votedBitSet finality.BitSet
	votedBitSet.SetBit(0)
	err = c.verifyFinalitySignatures(nil, votedBitSet, nil, &header, nil)
	if !errors.Is(err, finality.ErrNotEnoughFinalityVote) {
		t.Errorf("Expect error %v have %v", finality.ErrNotEnoughFinalityVote, err)
	}

	votedBitSet = finality.BitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(3)
	err = c.verifyFinalitySignatures(nil, votedBitSet, nil, &header, nil)
	if !errors.Is(err, finality.ErrInvalidFinalityVotedBitSet) {
		t.Errorf("Expect error %v have %v", finality.ErrInvalidFinalityVotedBitSet, err)
	}

	votedBitSet = finality.BitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature := blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[3],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, &header, nil)
	if !errors.Is(err, finality.ErrFinalitySignatureVerificationFailed) {
		t.Errorf("Expect error %v have %v", finality.ErrFinalitySignatureVerificationFailed, err)
	}

	votedBitSet = finality.BitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
		signature[3],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, &header, nil)
	if !errors.Is(err, finality.ErrFinalitySignatureVerificationFailed) {
		t.Errorf("Expect error %v have %v", finality.ErrFinalitySignatureVerificationFailed, err)
	}

	votedBitSet = finality.BitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, &header, nil)
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
	valWithBlsPub[0].Weight = 6666
	valWithBlsPub[1].Weight = 1
	valWithBlsPub[2].Weight = 3333

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
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
	c := Consortium{
		chainConfig: &params.ChainConfig{
			ShillinBlock: big.NewInt(0),
			TrippBlock:   big.NewInt(0),
		},
		config: &params.ConsortiumConfig{
			EpochV2: 300,
		},
		recents:            recents,
		isTest:             true,
		testTrippEffective: true,
	}
	snap.Hash = blockHash
	c.recents.Add(snap.Hash, snap)

	header := types.Header{Number: big.NewInt(int64(blockNumber + 1)), ParentHash: blockHash}
	// 1 voter with vote weight 6666 does not reach the threshold
	votedBitSet := finality.BitSet(0)
	votedBitSet.SetBit(0)
	aggregatedSignature := blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, &header, nil)
	if !errors.Is(err, finality.ErrNotEnoughFinalityVote) {
		t.Errorf("Expect error %v have %v", finality.ErrNotEnoughFinalityVote, err)
	}

	// 2 voters with total vote weight 3333 + 1 does not reach the threshold
	votedBitSet = finality.BitSet(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[1],
		signature[2],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, &header, nil)
	if !errors.Is(err, finality.ErrNotEnoughFinalityVote) {
		t.Errorf("Expect error %v have %v", finality.ErrNotEnoughFinalityVote, err)
	}

	// 2 voters with total vote weight 6666 + 1 reach the threshold
	votedBitSet = finality.BitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, &header, nil)
	if err != nil {
		t.Errorf("Expect successful verification have %v", err)
	}

	// All voters vote
	votedBitSet = finality.BitSet(0)
	votedBitSet.SetBit(0)
	votedBitSet.SetBit(1)
	votedBitSet.SetBit(2)
	aggregatedSignature = blst.AggregateSignatures([]blsCommon.Signature{
		signature[0],
		signature[1],
		signature[2],
	})
	err = c.verifyFinalitySignatures(nil, votedBitSet, aggregatedSignature, &header, nil)
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

type validatorWithBlsWeight struct {
	Address      common.Address
	BlsPublicKey blsCommon.PublicKey
	StakedAmount *big.Int
}

type mockTrippContract struct {
	blockProducers       []common.Address
	checkpointValidators []validatorWithBlsWeight
}

func (contract *mockTrippContract) WrapUpEpoch(opts *consortiumCommon.ApplyTransactOpts) error {
	return nil
}

func (contract *mockTrippContract) SubmitBlockReward(opts *consortiumCommon.ApplyTransactOpts) error {
	return nil
}

func (contract *mockTrippContract) Slash(opts *consortiumCommon.ApplyTransactOpts, spoiledValidator common.Address) error {
	return nil
}

func (contract *mockTrippContract) FinalityReward(opts *consortiumCommon.ApplyTransactOpts, votedValidators []common.Address) error {
	return nil
}

func (contract *mockTrippContract) GetBlockProducers(_ common.Hash, _ *big.Int) ([]common.Address, error) {
	return contract.blockProducers, nil
}

func (contract *mockTrippContract) GetValidatorCandidates(_ common.Hash, _ *big.Int) ([]common.Address, error) {
	validatorCandidates := make([]common.Address, 0)
	for _, val := range contract.checkpointValidators {
		validatorCandidates = append(validatorCandidates, val.Address)
	}
	return validatorCandidates, nil
}

func (contract *mockTrippContract) GetBlsPublicKey(_ common.Hash, _ *big.Int, address common.Address) (blsCommon.PublicKey, error) {
	for _, val := range contract.checkpointValidators {
		if val.Address == address {
			return val.BlsPublicKey, nil
		}
	}
	return nil, errors.New("address is not in validator candidate list")
}

func (contract *mockTrippContract) GetStakedAmount(_ common.Hash, _ *big.Int, addrs []common.Address) ([]*big.Int, error) {
	stakes := make([]*big.Int, 0)
	for _, addr := range addrs {
		for _, val := range contract.checkpointValidators {
			if val.Address == addr {
				stakes = append(stakes, val.StakedAmount)
			}
		}
	}
	return stakes, nil
}

func (contract *mockTrippContract) GetMaxValidatorNumber(blockHash common.Hash, blockNumber *big.Int) (*big.Int, error) {
	return big.NewInt(int64(len(contract.blockProducers))), nil
}

func TestGetCheckpointValidatorFromContract(t *testing.T) {
	var err error
	secretKeys := make([]blsCommon.SecretKey, 4)
	for i := 0; i < len(secretKeys); i++ {
		secretKeys[i], err = blst.RandKey()
		if err != nil {
			t.Fatalf("Failed to generate secret key, err: %s", err)
		}
	}

	mock := &mockTrippContract{
		checkpointValidators: []validatorWithBlsWeight{
			validatorWithBlsWeight{
				Address:      common.Address{0x1},
				BlsPublicKey: secretKeys[0].PublicKey(),
				StakedAmount: new(big.Int).SetUint64(100),
			},
			validatorWithBlsWeight{
				Address:      common.Address{0x2},
				BlsPublicKey: secretKeys[1].PublicKey(),
				StakedAmount: new(big.Int).SetUint64(200),
			},
			validatorWithBlsWeight{
				Address:      common.Address{0x3},
				BlsPublicKey: secretKeys[2].PublicKey(),
				StakedAmount: new(big.Int).SetUint64(400),
			},
		},
		blockProducers: []common.Address{
			common.Address{0x11},
			common.Address{0x22},
		},
	}
	c := Consortium{
		chainConfig: &params.ChainConfig{
			ShillinBlock: big.NewInt(0),
			TrippBlock:   big.NewInt(1),
		},
		config: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		contract:           mock,
		isTest:             true,
		testTrippEffective: true,
		testTrippPeriod:    false,
	}

	validatorWithPubs, blockProducers, err := c.getCheckpointValidatorsFromContract(nil, &types.Header{Number: big.NewInt(3)})
	if err != nil {
		t.Fatalf("Failed to get checkpoint validators from contract, err: %s", err)
	}
	if validatorWithPubs != nil {
		t.Fatalf("Expect nil returned list")
	}
	if len(blockProducers) != 2 {
		t.Fatalf("Expect returned list, length: %d have: %d", 2, len(blockProducers))
	}
	if blockProducers[0] != (common.Address{0x11}) {
		t.Fatalf("Wrong returned list")
	}

	c.testTrippPeriod = true
	validatorWithPubs, blockProducers, err = c.getCheckpointValidatorsFromContract(nil, &types.Header{Number: big.NewInt(200)})
	if err != nil {
		t.Fatalf("Failed to get checkpoint validators from contract, err: %s", err)
	}
	if validatorWithPubs == nil {
		t.Fatalf("Expect returned list")
	}
	if len(validatorWithPubs) != 3 {
		t.Fatalf("Expect returned list, length: %d have: %d", 3, len(validatorWithPubs))
	}
	if validatorWithPubs[0].Address != (common.Address{0x1}) {
		t.Fatalf("Wrong returned list")
	}
	if !validatorWithPubs[0].BlsPublicKey.Equals(secretKeys[0].PublicKey()) {
		t.Fatalf("Wrong returned list")
	}
	if len(blockProducers) != 2 {
		t.Fatalf("Expect returned list, length: %d have: %d", 2, len(blockProducers))
	}
	if blockProducers[0] != (common.Address{0x11}) {
		t.Fatalf("Wrong returned list")
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

func (contract *mockContract) GetBlockProducers(_ common.Hash, _ *big.Int) ([]common.Address, error) {
	var validatorAddresses []common.Address
	for address := range contract.validators {
		validatorAddresses = append(validatorAddresses, address)
	}
	return validatorAddresses, nil
}

func (contract *mockContract) GetValidatorCandidates(_ common.Hash, _ *big.Int) ([]common.Address, error) {
	return nil, nil
}

func (contract *mockContract) GetBlsPublicKey(_ common.Hash, _ *big.Int, address common.Address) (blsCommon.PublicKey, error) {
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

func (contract *mockContract) GetStakedAmount(_ common.Hash, _ *big.Int, _ []common.Address) ([]*big.Int, error) {
	return nil, nil
}

func (contract *mockContract) GetMaxValidatorNumber(blockHash common.Hash, blockNumber *big.Int) (*big.Int, error) {
	return nil, nil
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
	c.assembleFinalityVote(nil, &header, snap)

	extraData, err = finality.DecodeExtra(header.Extra, true)
	if err != nil {
		t.Fatalf("Failed to decode extra data, err: %s", err)
	}

	if extraData.HasFinalityVote != 1 {
		t.Fatal("Missing finality vote in header")
	}

	bitSet := finality.BitSet(0)
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

func TestAssembleFinalityVoteTripp(t *testing.T) {
	var err error
	numValidators := 3
	secretKeys := make([]blsCommon.SecretKey, numValidators)
	for i := range secretKeys {
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

	signatures := make([]blsCommon.Signature, numValidators)
	for i := range signatures {
		signatures[i] = secretKeys[i].Sign(digest[:])
	}

	votes := make([]*types.VoteEnvelope, numValidators)
	for i := range votes {
		votes[i] = &types.VoteEnvelope{
			RawVoteEnvelope: types.RawVoteEnvelope{
				PublicKey: types.BLSPublicKey(secretKeys[i].PublicKey().Marshal()),
				Signature: types.BLSSignature(signatures[i].Marshal()),
				Data:      &voteData,
			},
		}
	}

	mock := mockVotePool{
		vote: votes,
	}

	chainConfig := params.ChainConfig{
		ShillinBlock: big.NewInt(0),
		TrippBlock:   big.NewInt(0),
	}
	c := Consortium{
		chainConfig:        &chainConfig,
		votePool:           &mock,
		isTest:             true,
		testTrippEffective: true,
	}

	validators := make([]finality.ValidatorWithBlsPub, numValidators)
	for i := range validators {
		validators[i] = finality.ValidatorWithBlsPub{
			Address:      common.BigToAddress(big.NewInt(int64(i))),
			BlsPublicKey: secretKeys[i].PublicKey(),
		}
	}
	validators[0].Weight = 6666
	validators[1].Weight = 1
	validators[2].Weight = 3333

	snap := newSnapshot(nil, nil, nil, 10, common.Hash{}, nil, validators, nil)

	header := types.Header{Number: big.NewInt(5)}
	extraData := &finality.HeaderExtraData{}
	header.Extra, err = extraData.EncodeV2(&chainConfig, header.Number)
	if err != nil {
		t.Fatalf("Failed to encode extradata, err %s", err)
	}

	// Case 1: only 1 validator with weight 6666 cannot reach the threshold
	mock.vote = mock.vote[:1]
	c.assembleFinalityVote(nil, &header, snap)
	extraData, err = finality.DecodeExtraV2(header.Extra, &chainConfig, header.Number)
	if err != nil {
		t.Fatalf("Failed to decode extradata, err %s", err)
	}
	if uint64(extraData.FinalityVotedValidators) != 0 {
		t.Fatalf("Expect vote bit set to be %d, got %d", 0, uint64(extraData.FinalityVotedValidators))
	}

	// Case 2: 1 validator with 2 votes cannot reach the threshold
	header = types.Header{Number: big.NewInt(5)}
	extraData = &finality.HeaderExtraData{}
	header.Extra, _ = extraData.EncodeV2(&chainConfig, header.Number)
	mock.vote = make([]*types.VoteEnvelope, 0)
	mock.vote = append(mock.vote, votes[0], votes[0])
	c.assembleFinalityVote(nil, &header, snap)
	extraData, err = finality.DecodeExtraV2(header.Extra, &chainConfig, header.Number)
	if err != nil {
		t.Fatalf("Failed to decode extradata, err %s", err)
	}
	if uint64(extraData.FinalityVotedValidators) != 0 {
		t.Fatalf("Expect vote bit set to be %d, got %d", 0, uint64(extraData.FinalityVotedValidators))
	}

	// Case 3: 2 validators with total vote weight 6667 can reach the threshold
	header = types.Header{Number: big.NewInt(5)}
	extraData = &finality.HeaderExtraData{}
	header.Extra, _ = extraData.EncodeV2(&chainConfig, header.Number)
	mock.vote = make([]*types.VoteEnvelope, 0)
	mock.vote = append(mock.vote, votes[0], votes[1])
	c.assembleFinalityVote(nil, &header, snap)
	extraData, err = finality.DecodeExtraV2(header.Extra, &chainConfig, header.Number)
	if err != nil {
		t.Fatalf("Failed to decode extradata, err %s", err)
	}
	bitSet := finality.BitSet(0)
	bitSet.SetBit(0)
	bitSet.SetBit(1)
	if uint64(extraData.FinalityVotedValidators) != uint64(bitSet) {
		t.Fatalf("Expect vote bit set to be %d, got %d", uint64(bitSet), uint64(extraData.FinalityVotedValidators))
	}
	var includedSignatures []blsCommon.Signature
	includedSignatures = append(includedSignatures, signatures[0], signatures[1])
	aggregatedSignature := blst.AggregateSignatures(includedSignatures)
	if !bytes.Equal(aggregatedSignature.Marshal(), extraData.AggregatedFinalityVotes.Marshal()) {
		t.Fatal("Mismatch signature")
	}

	// Case 4: 1 validator with vote weight 6667 can reach the threshold
	validators[0].Weight = 6667
	validators[1].Weight = 1
	validators[2].Weight = 3332

	snap = newSnapshot(nil, nil, nil, 10, common.Hash{}, nil, validators, nil)
	header = types.Header{Number: big.NewInt(5)}
	extraData = &finality.HeaderExtraData{}
	header.Extra, _ = extraData.EncodeV2(&chainConfig, header.Number)
	mock.vote = make([]*types.VoteEnvelope, 0)
	mock.vote = append(mock.vote, votes[0])
	c.assembleFinalityVote(nil, &header, snap)
	extraData, err = finality.DecodeExtraV2(header.Extra, &chainConfig, header.Number)
	if err != nil {
		t.Fatalf("Failed to decode extradata, err %s", err)
	}
	bitSet = finality.BitSet(0)
	bitSet.SetBit(0)
	if uint64(extraData.FinalityVotedValidators) != uint64(bitSet) {
		t.Fatalf("Expect vote bit set to be %d, got %d", uint64(bitSet), uint64(extraData.FinalityVotedValidators))
	}
	if !bytes.Equal(signatures[0].Marshal(), extraData.AggregatedFinalityVotes.Marshal()) {
		t.Fatal("Mismatch signature")
	}
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
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
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
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
	signatures, _ := arc.NewARC[common.Hash, common.Address](inmemorySignatures)

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
					voteBitset finality.BitSet
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
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
	signatures, _ := arc.NewARC[common.Hash, common.Address](inmemorySignatures)

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

func TestUpgradeAxieProxyCode(t *testing.T) {
	secretKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	validatorAddr := crypto.PubkeyToAddress(secretKey.PublicKey)
	blsSecret, err := blst.RandKey()
	if err != nil {
		t.Fatal(err)
	}

	db := rawdb.NewMemoryDatabase()
	data := map[string]string{
		"code": "0x608060405234801561001057600080fd5b50600436106100a95760003560e01c80636a0cd1f5116100715780636a0cd1f514610160578063b7ab4db51461018c578063c370b042146101e4578063dafae408146101ec578063facd743b1461021d578063fc81975014610243576100a9565b80630f43a677146100ae57806335aa2e44146100c85780634b561753146101015780634e70b1dc1461012f57806353727d2614610137575b600080fd5b6100b661024b565b60408051918252519081900360200190f35b6100e5600480360360208110156100de57600080fd5b5035610251565b604080516001600160a01b039092168252519081900360200190f35b61012d6004803603604081101561011757600080fd5b50803590602001356001600160a01b0316610278565b005b6100b6610408565b61012d6004803603606081101561014d57600080fd5b508035906020810135906040013561040e565b61012d6004803603604081101561017657600080fd5b50803590602001356001600160a01b03166105a4565b610194610743565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101d05781810151838201526020016101b8565b505050509050019250505060405180910390f35b6100b66107a5565b6102096004803603602081101561020257600080fd5b50356107ab565b604080519115158252519081900360200190f35b6102096004803603602081101561023357600080fd5b50356001600160a01b03166107e0565b6100e56107fe565b60025481565b6001818154811061025e57fe5b6000918252602090912001546001600160a01b0316905081565b610281336107e0565b61028a57600080fd5b604080516001600160a01b03808416828401526020808301849052600c60608401526b30b2322b30b634b230ba37b960a11b6080808501919091528451808503909101815260a09093019093528151919092012060055490916000911663ec9ab83c6102f461080d565b8685336040518563ffffffff1660e01b81526004018080602001858152602001848152602001836001600160a01b03166001600160a01b03168152602001828103825286818151815260200191508051906020019080838360005b8381101561036757818101518382015260200161034f565b50505050905090810190601f1680156103945780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156103b657600080fd5b505af11580156103ca573d6000803e3d6000fd5b505050506040513d60208110156103e057600080fd5b5051905060018160028111156103f257fe5b1415610402576104028484610944565b50505050565b60035481565b610417336107e0565b61042057600080fd5b604080518082018490526060808201849052602080830191909152600c60808301526b75706461746551756f72756d60a01b60a0808401919091528351808403909101815260c090920190925280519101206005546000906001600160a01b031663ec9ab83c61048e61080d565b8785336040518563ffffffff1660e01b81526004018080602001858152602001848152602001836001600160a01b03166001600160a01b03168152602001828103825286818151815260200191508051906020019080838360005b838110156105015781810151838201526020016104e9565b50505050905090810190601f16801561052e5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561055057600080fd5b505af1158015610564573d6000803e3d6000fd5b505050506040513d602081101561057a57600080fd5b50519050600181600281111561058c57fe5b141561059d5761059d858585610a02565b5050505050565b6105ad336107e0565b6105b657600080fd5b6105bf816107e0565b6105c857600080fd5b604080516001600160a01b03808416828401526020808301849052600f60608401526e3932b6b7bb32ab30b634b230ba37b960891b6080808501919091528451808503909101815260a09093019093528151919092012060055490916000911663ec9ab83c61063561080d565b8685336040518563ffffffff1660e01b81526004018080602001858152602001848152602001836001600160a01b03166001600160a01b03168152602001828103825286818151815260200191508051906020019080838360005b838110156106a8578181015183820152602001610690565b50505050905090810190601f1680156106d55780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156106f757600080fd5b505af115801561070b573d6000803e3d6000fd5b505050506040513d602081101561072157600080fd5b50519050600181600281111561073357fe5b1415610402576104028484610a69565b6060600180548060200260200160405190810160405280929190818152602001828054801561079b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161077d575b5050505050905090565b60045481565b60006107c4600254600354610bc890919063ffffffff16565b6004546107d890849063ffffffff610bc816565b101592915050565b6001600160a01b031660009081526020819052604090205460ff1690565b6005546001600160a01b031681565b60055460408051638e46684960e01b815290516060926001600160a01b031691638e466849916004808301926000929190829003018186803b15801561085257600080fd5b505afa158015610866573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561088f57600080fd5b81019080805160405193929190846401000000008211156108af57600080fd5b9083019060208201858111156108c457600080fd5b82516401000000008111828201881017156108de57600080fd5b82525081516020918201929091019080838360005b8381101561090b5781810151838201526020016108f3565b50505050905090810190601f1680156109385780820380516001836020036101000a031916815260200191505b50604052505050905090565b6001600160a01b03811660009081526020819052604090205460ff161561096a57600080fd5b6001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b038416908117909155600081815260208190526040808220805460ff191685179055600280549094019093559151909184917f7429a06e9412e469f0d64f9d222640b0af359f556b709e2913588c227851b88d9190a35050565b80821115610a0f57600080fd5b600380546004805492859055839055604080518281526020810184905281519293928592879289927f976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8929181900390910190a45050505050565b610a72816107e0565b610a7b57600080fd5b6000805b600254811015610acb57826001600160a01b031660018281548110610aa057fe5b6000918252602090912001546001600160a01b03161415610ac357809150610acb565b600101610a7f565b506001600160a01b0382166000908152602081905260409020805460ff1916905560025460018054909160001901908110610b0257fe5b600091825260209091200154600180546001600160a01b039092169183908110610b2857fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506001805480610b6157fe5b600082815260208120820160001990810180546001600160a01b03191690559182019092556002805490910190556040516001600160a01b0384169185917f7126bef88d1149ccdff9681ed5aecd3ba5ae70c96517551de250af09cebd1a0b9190a3505050565b600082610bd757506000610bf0565b5081810281838281610be557fe5b0414610bf057600080fd5b9291505056fea265627a7a72315820ee5f68147305f40cf9481c24f13db7dda3a5cbf9c93b41c4ead22f306768974b64736f6c63430005110032",
	}
	raw, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	var contractCodeUpgrade params.ContractCodeUpgrade
	if err := json.Unmarshal(raw, &contractCodeUpgrade); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(common.FromHex(data["code"]), contractCodeUpgrade.Code) {
		t.Fatal("mismatch code")
	}

	code := contractCodeUpgrade.Code
	chainConfig := &params.ChainConfig{
		ChainID:           big.NewInt(2021),
		HomesteadBlock:    common.Big0,
		EIP150Block:       common.Big0,
		EIP155Block:       common.Big0,
		EIP158Block:       common.Big0,
		ConsortiumV2Block: common.Big0,
		TrippBlock:        common.Big0,
		AaronBlock:        big.NewInt(3),
		Consortium: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		ConsortiumV2Contracts: &params.ConsortiumV2Contracts{
			RoninValidatorSet: common.HexToAddress("0xaa"),
		},
		TransparentProxyCodeUpgrade: &params.ContractCodeUpgrade{
			AxieAddress: common.Address{0x12},
			LandAddress: common.Address{0x13},
			Code:        code,
		},
	}
	genesis := (&core.Genesis{
		Config: chainConfig,
	}).MustCommit(db)
	mock := &mockTrippContract{
		checkpointValidators: []validatorWithBlsWeight{
			validatorWithBlsWeight{
				Address:      common.Address{0x1},
				BlsPublicKey: blsSecret.PublicKey(),
				StakedAmount: new(big.Int).SetUint64(100),
			},
		},
		blockProducers: []common.Address{
			validatorAddr,
		},
	}
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
	signatures, _ := arc.NewARC[common.Hash, common.Address](inmemorySignatures)
	v2 := &Consortium{
		chainConfig: chainConfig,
		contract:    mock,
		recents:     recents,
		signatures:  signatures,
		config: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		isTest:             true,
		testTrippEffective: true,
	}

	chain, _ := core.NewBlockChain(db, nil, chainConfig, v2, vm.Config{}, nil, nil)
	extraData := &finality.HeaderExtraData{}

	parent := genesis
	for i := 0; i < 5; i++ {
		blocks, _ := core.GenerateChain(
			chainConfig,
			parent,
			v2,
			db,
			1,
			func(i int, bg *core.BlockGen) {
				bg.SetCoinbase(validatorAddr)
				extra, err := extraData.EncodeV2(chainConfig, big.NewInt(int64(i)))
				if err != nil {
					t.Fatal(err)
				}
				bg.SetExtra(extra)
				bg.SetDifficulty(big.NewInt(7))
			},
			true)

		header := blocks[0].Header()
		hash := calculateSealHash(header, big.NewInt(2021))
		sig, err := crypto.Sign(hash[:], secretKey)
		if err != nil {
			t.Fatal(err)
		}

		copy(header.Extra[len(header.Extra)-consortiumCommon.ExtraSeal:], sig)
		blocks[0] = blocks[0].WithSeal(header)
		parent = blocks[0]

		_, err = chain.InsertChain(blocks)
		if err != nil {
			t.Fatalf("Failed to insert chain, err %s", err)
		}

		if header.Number.Uint64() == chainConfig.AaronBlock.Uint64()+1 {
			statedb, err := chain.State()
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(
				statedb.GetCode(v2.chainConfig.TransparentProxyCodeUpgrade.AxieAddress),
				chainConfig.TransparentProxyCodeUpgrade.Code,
			) {
				t.Fatal("Failed to set axie proxy code.")
			}
			if !bytes.Equal(
				statedb.GetCode(v2.chainConfig.TransparentProxyCodeUpgrade.LandAddress),
				chainConfig.TransparentProxyCodeUpgrade.Code,
			) {
				t.Fatal("Failed to set land proxy code.")
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
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
	signatures, _ := arc.NewARC[common.Hash, common.Address](inmemorySignatures)

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
	const NUM_OF_VALIDATORS = 21
	dateInSeconds := uint64(86400)
	now := uint64(time.Now().Unix())
	midnight := uint64(now / dateInSeconds * dateInSeconds)

	db := rawdb.NewMemoryDatabase()
	chainConfig := params.ChainConfig{
		ChainID:    big.NewInt(2021),
		TrippBlock: big.NewInt(30),
		Consortium: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		ConsortiumV2Contracts: &params.ConsortiumV2Contracts{
			RoninValidatorSet: common.HexToAddress("0xaa"),
		},
	}
	genesis := (&core.Genesis{
		Config:    &chainConfig,
		BaseFee:   big.NewInt(params.InitialBaseFee),
		Timestamp: midnight, // genesis at day 1
	}).MustCommit(db)
	chain, _ := core.NewBlockChain(db, nil, &chainConfig, ethash.NewFullFaker(), vm.Config{}, nil, nil)
	// create chain of up to 399 blocks, all of them are not period block
	bs, _ := core.GenerateChain(&chainConfig, genesis, ethash.NewFaker(), db, 399, nil, true) // create chain of up to 399 blocks
	if _, err := chain.InsertChain(bs[:]); err != nil {
		panic(err)
	}
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
	signatures, _ := arc.NewARC[common.Hash, common.Address](inmemorySignatures)
	mock := &mockContract{
		validators: map[common.Address]blsCommon.PublicKey{},
	}
	c := &Consortium{
		chainConfig: &chainConfig,
		recents:     recents,
		signatures:  signatures,
		config: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		db:       db,
		contract: mock,
	}
	validators := make([]common.Address, NUM_OF_VALIDATORS)
	for i := 0; i < NUM_OF_VALIDATORS; i++ {
		validators = append(validators, common.BigToAddress(big.NewInt(int64(i))))
	}

	var header = &types.Header{}

	// header of block 0
	// this must not a period block
	header = genesis.Header()
	if c.IsPeriodBlock(chain, header, nil) {
		t.Errorf("wrong period block")
	}

	// header of block 200
	// this must not a period block
	header = bs[199].Header()
	if c.IsPeriodBlock(chain, header, nil) {
		t.Error("wrong period block")
	}

	header = bs[351].Header()
	if c.IsPeriodBlock(chain, header, nil) {
		t.Error("wrong period block")
	}

	for i := 0; i < 210; i++ {
		callback := func(i int, bg *core.BlockGen) {
			if i == 0 {
				bg.OffsetTime(int64(dayInSeconds))
			}
		}
		block, _ := core.GenerateChain(&chainConfig, bs[len(bs)-1], ethash.NewFaker(), db, 1, callback, true)
		bs = append(bs, block...)
	}
	if _, err := chain.InsertChain(bs[:]); err != nil {
		panic(err)
	}

	// header of block 400
	// this must be a period block
	header = bs[399].Header()
	// this header must be period header
	if !c.IsPeriodBlock(chain, header, nil) {
		t.Errorf("wrong period block")
	}

	// header of block 500
	// this must not be a period block
	header = bs[499].Header()
	if c.IsPeriodBlock(chain, header, nil) {
		t.Errorf("wrong period block")
	}
}

func TestIsTrippEffective(t *testing.T) {
	now := uint64(time.Now().Unix())
	midnight := uint64(now / dayInSeconds * dayInSeconds)
	db := rawdb.NewMemoryDatabase()
	chainConfig := params.ChainConfig{
		ChainID:    big.NewInt(2021),
		TrippBlock: big.NewInt(30),
		Consortium: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		ConsortiumV2Contracts: &params.ConsortiumV2Contracts{
			RoninValidatorSet: common.HexToAddress("0xaa"),
		},
		TrippPeriod: new(big.Int).SetUint64(now / dayInSeconds),
	}
	genesis := (&core.Genesis{
		Config:    &chainConfig,
		BaseFee:   big.NewInt(params.InitialBaseFee),
		Timestamp: midnight, // genesis at day 1
	}).MustCommit(db)
	chain, _ := core.NewBlockChain(db, nil, &chainConfig, ethash.NewFullFaker(), vm.Config{}, nil, nil)
	// create chain of up to 399 blocks, all of them are not Tripp effective
	bs, _ := core.GenerateChain(&chainConfig, genesis, ethash.NewFaker(), db, 399, nil, true)
	if _, err := chain.InsertChain(bs[:]); err != nil {
		panic(err)
	}
	recents, _ := arc.NewARC[common.Hash, *Snapshot](inmemorySnapshots)
	signatures, _ := arc.NewARC[common.Hash, common.Address](inmemorySignatures)
	mock := &mockContract{
		validators: map[common.Address]blsCommon.PublicKey{},
	}
	c := &Consortium{
		chainConfig: &chainConfig,
		recents:     recents,
		signatures:  signatures,
		config: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		testTrippEffective: false,
		db:                 db,
		contract:           mock,
	}

	var header = &types.Header{}

	// header of block 30
	header = bs[29].Header()
	// this header must not be Tripp effective
	if c.IsTrippEffective(chain, header) {
		t.Error("fail test Tripp effective")
	}

	// header of block 201
	// this header must not be Tripp effective
	header = bs[201].Header()
	if c.IsTrippEffective(chain, header) {
		t.Error("fail test Tripp effective")
	}

	// header of block 200
	// this header must not be Tripp effective
	header = bs[200].Header()
	if c.IsTrippEffective(chain, header) {
		t.Error("fail test Tripp effective")
	}

	// header of block 399
	// this header must not be Tripp effective
	header = bs[398].Header()
	if c.IsTrippEffective(chain, header) {
		t.Error("fail test Tripp effective")
	}

	for i := 0; i < 210; i++ {
		callback := func(i int, bg *core.BlockGen) {
			if i == 0 {
				bg.OffsetTime(int64(dayInSeconds))
			}
		}
		block, _ := core.GenerateChain(&chainConfig, bs[len(bs)-1], ethash.NewFaker(), db, 1, callback, true)
		bs = append(bs, block...)
	}
	if _, err := chain.InsertChain(bs[:]); err != nil {
		panic(err)
	}

	// header of block 400
	// this header must be Tripp effective
	header = bs[399].Header()
	if !c.IsTrippEffective(nil, header) {
		t.Error("fail test Tripp effective")
	}

	// header of block 402
	// this header must be Tripp effective
	header = bs[401].Header()
	if !c.IsTrippEffective(chain, header) {
		t.Error("fail test Tripp effective")
	}

	header = bs[599].Header()
	// this header must be Tripp effective
	if !c.IsTrippEffective(chain, header) {
		t.Error("fail test Tripp effective")
	}
}

func TestHeaderExtraDataCheck(t *testing.T) {
	c := Consortium{
		chainConfig: &params.ChainConfig{
			TrippBlock: common.Big0,
		},
		config: &params.ConsortiumConfig{
			EpochV2: 200,
		},
		isTest:             true,
		testTrippEffective: true,
	}

	// Case 1: not an epoch block, every validator field must be empty
	// non-empty checkpoint validators
	header := types.Header{Number: big.NewInt(100)}
	extraData := finality.HeaderExtraData{
		CheckpointValidators: []finality.ValidatorWithBlsPub{
			{},
		},
	}
	err := c.verifyValidatorFieldsInExtraData(nil, &extraData, &header, nil)
	if !errors.Is(err, consortiumCommon.ErrNonEpochExtraData) {
		t.Fatalf("Expect err: %v got: %v", consortiumCommon.ErrNonEpochExtraData, err)
	}

	// non-empty block producers
	extraData = finality.HeaderExtraData{
		BlockProducers: []common.Address{
			{},
		},
	}
	err = c.verifyValidatorFieldsInExtraData(nil, &extraData, &header, nil)
	if !errors.Is(err, consortiumCommon.ErrNonEpochExtraData) {
		t.Fatalf("Expect err: %v got: %v", consortiumCommon.ErrNonEpochExtraData, err)
	}

	// non-empty block producer bitset
	extraData = finality.HeaderExtraData{
		BlockProducersBitSet: 10,
	}
	err = c.verifyValidatorFieldsInExtraData(nil, &extraData, &header, nil)
	if !errors.Is(err, consortiumCommon.ErrNonEpochExtraData) {
		t.Fatalf("Expect err: %v got: %v", consortiumCommon.ErrNonEpochExtraData, err)
	}

	// Case 2: Not a period block, checkpoint validators must be empty
	header = types.Header{Number: big.NewInt(200)}
	extraData = finality.HeaderExtraData{
		CheckpointValidators: []finality.ValidatorWithBlsPub{
			{},
		},
		BlockProducers: []common.Address{
			{},
		},
	}
	err = c.verifyValidatorFieldsInExtraData(nil, &extraData, &header, nil)
	if !errors.Is(err, consortiumCommon.ErrNonPeriodBlockExtraData) {
		t.Fatalf("Expect err: %v got: %v", consortiumCommon.ErrNonPeriodBlockExtraData, err)
	}

	// Case 3: Before Tripp effective, block producer, block producer bitset must be empty
	c.testTrippEffective = false
	header = types.Header{Number: big.NewInt(200)}
	extraData = finality.HeaderExtraData{
		CheckpointValidators: []finality.ValidatorWithBlsPub{
			{},
		},
		BlockProducers: []common.Address{
			{},
		},
	}
	err = c.verifyValidatorFieldsInExtraData(nil, &extraData, &header, nil)
	if !errors.Is(err, consortiumCommon.ErrPreTrippEpochProducerExtraData) {
		t.Fatalf("Expect err: %v got: %v", consortiumCommon.ErrPreTrippEpochProducerExtraData, err)
	}

	header = types.Header{Number: big.NewInt(200)}
	extraData = finality.HeaderExtraData{
		CheckpointValidators: []finality.ValidatorWithBlsPub{
			{},
		},
		BlockProducersBitSet: 5,
	}
	err = c.verifyValidatorFieldsInExtraData(nil, &extraData, &header, nil)
	if !errors.Is(err, consortiumCommon.ErrPreTrippEpochProducerExtraData) {
		t.Fatalf("Expect err: %v got: %v", consortiumCommon.ErrPreTrippEpochProducerExtraData, err)
	}

	// Case 4: A valid Aaron epoch block
	c.chainConfig.AaronBlock = common.Big0
	c.testTrippEffective = true
	header = types.Header{Number: big.NewInt(200)}
	extraData = finality.HeaderExtraData{
		BlockProducersBitSet: 5,
	}
	err = c.verifyValidatorFieldsInExtraData(nil, &extraData, &header, nil)
	if err != nil {
		t.Fatalf("Expect no error, got: %v", err)
	}
}

func TestEncodeDecodeValidatorBitSet(t *testing.T) {
	candidates := make([]finality.ValidatorWithBlsPub, 10)
	producers := make([]common.Address, 0)
	for i := 0; i < 10; i++ {
		secret, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			t.Fatal(err)
		}
		addr := crypto.PubkeyToAddress(secret.PublicKey)
		candidates[i] = finality.ValidatorWithBlsPub{Address: addr}
		if i%2 == 0 {
			producers = append(producers, addr)
		}
	}
	sort.Sort(finality.CheckpointValidatorAscending(candidates))
	enc := encodeValidatorBitSet(candidates, producers)

	// Test encode bit set
	sort.Sort(validatorsAscending(producers))
	indices := enc.Indices()
	if len(indices) != 5 {
		t.Fatalf("mismatch validator1, %v", indices)
	}
	var i int = 0
	for _, idx := range indices {
		if producers[i] != candidates[idx].Address {
			t.Fatal("mismatch validator")
		}
		i += 1
	}

	// Test decode bit set
	dec := decodeValidatorBitSet(enc, candidates)
	if len(dec) != 5 {
		t.Fatal("mismatch validator")
	}
	for i := 0; i < 5; i++ {
		if producers[i] != dec[i] {
			t.Fatal("mismatch validator")
		}
	}
}
