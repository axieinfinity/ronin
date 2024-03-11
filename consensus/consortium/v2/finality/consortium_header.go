package finality

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	blsCommon "github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	ExtraSeal   = crypto.SignatureLength
	ExtraVanity = 32
)

var (
	// ErrInvalidHasFinalityVote is returned if a block's extra-data contains invalid
	// has finality vote byte
	ErrInvalidHasFinalityVote = errors.New("invalid has finality vote byte")

	// ErrMissingHasFinalityVote is returned if a block's extra-data section does not seem
	// to include 1 byte to determine if the extra data has the finality votes
	ErrMissingHasFinalityVote = errors.New("extra-data 1 byte has finality votes missing")

	// ErrMissingFinalityVoteBitSet is returned if a block's extra-data section does not seem
	// to include 8 bytes of finality vote bitset
	ErrMissingFinalityVoteBitSet = errors.New("extra-data 8 bytes finality votes bitset missing")

	// ErrMissingFinalitySignature is returned if a block's extra-data section does not seem
	// to include finality signature
	ErrMissingFinalitySignature = errors.New("extra-data finality signature missing")

	// ErrMissingFinalitySignature is returned if the number of finality votes is under
	// the threshold
	ErrNotEnoughFinalityVote = errors.New("not enough finality vote")

	// ErrFinalitySignatureVerificationFailed is returned if the finality signature verification
	// failed
	ErrFinalitySignatureVerificationFailed = errors.New("failed to verify finality signature")

	// ErrInvalidFinalityVotedBitSet is returned if the voted validator in bit set is not in
	// snapshot validator set
	ErrInvalidFinalityVotedBitSet = errors.New("invalid finality voted bit set")

	// ErrUnauthorizedFinalityVoter is returned if finality voter is not in validator set
	ErrUnauthorizedFinalityVoter = errors.New("unauthorized finality voter")

	// ErrMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	ErrMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// ErrMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	ErrMissingSignature = errors.New("extra-data 65 byte signature suffix missing")

	// ErrInvalidSpanValidators is returned if a block contains an
	// invalid list of validators (i.e. non divisible by 20 bytes).
	ErrInvalidSpanValidators = errors.New("invalid validator list on sprint end block")

	// ErrInvalidTargetNumber is returned if the vote contains invalid
	// target number
	ErrInvalidTargetNumber = errors.New("invalid target number in vote")

	// ErrNilAggregatedFinalityVotes is returned if the aggregated votes is nil
	ErrNilAggregatedFinalityVotes = errors.New("aggregated finality votes is nil")

	// ErrNilBlsPublicKey is returned if the bls public key is nil
	ErrNilBlsPublicKey = errors.New("bls public key is nil")
)

type ValidatorWithBlsPub struct {
	Address      common.Address
	BlsPublicKey blsCommon.PublicKey
}

type savedValidatorWithBlsPub struct {
	Address      common.Address `json:"address"`
	BlsPublicKey string         `json:"blsPublicKey,omitempty"`
}

func (validator *ValidatorWithBlsPub) UnmarshalJSON(input []byte) error {
	var (
		savedValidator savedValidatorWithBlsPub
		err            error
	)

	if err = json.Unmarshal(input, &savedValidator); err != nil {
		return err
	}

	validator.Address = savedValidator.Address
	rawPublicKey, err := hex.DecodeString(savedValidator.BlsPublicKey)
	if err != nil {
		return err
	}
	validator.BlsPublicKey, err = blst.PublicKeyFromBytes(rawPublicKey)
	if err != nil {
		return err
	}
	return nil
}

func (validator *ValidatorWithBlsPub) MarshalJSON() ([]byte, error) {
	savedValidator := savedValidatorWithBlsPub{
		Address: validator.Address,
	}

	if validator.BlsPublicKey != nil {
		savedValidator.BlsPublicKey = hex.EncodeToString(validator.BlsPublicKey.Marshal())
	}

	return json.Marshal(&savedValidator)
}

// CheckpointValidatorAscending implements the sort interface to allow sorting a list
// of checkpoint validator
type CheckpointValidatorAscending []ValidatorWithBlsPub

func (validator CheckpointValidatorAscending) Len() int { return len(validator) }
func (validator CheckpointValidatorAscending) Less(i, j int) bool {
	return bytes.Compare(validator[i].Address[:], validator[j].Address[:]) < 0
}
func (validator CheckpointValidatorAscending) Swap(i, j int) {
	validator[i], validator[j] = validator[j], validator[i]
}

type FinalityVoteBitSet uint64

const finalityVoteBitSetByteLength int = 8

func (bitSet *FinalityVoteBitSet) Indices() []int {
	var votedValidatorPositions []int

	for i := 0; i < finalityVoteBitSetByteLength*8; i++ {
		if uint64(*bitSet)&(1<<i) != 0 {
			votedValidatorPositions = append(votedValidatorPositions, i)
		}
	}
	return votedValidatorPositions
}

func (bitSet *FinalityVoteBitSet) SetBit(index int) {
	if index >= finalityVoteBitSetByteLength*8 {
		return
	}

	*bitSet = FinalityVoteBitSet(uint64(*bitSet) | (1 << index))
}

// HeaderExtraData represents the information in the extra data of header,
// this helps to make the code more readable
type HeaderExtraData struct {
	Vanity                  [ExtraVanity]byte     // unused in Consortium, filled with zero
	HasFinalityVote         uint8                 // determine if the header extra has the finality vote
	FinalityVotedValidators FinalityVoteBitSet    // the bit set of validators that vote for finality
	AggregatedFinalityVotes blsCommon.Signature   // aggregated BLS signatures for finality vote
	CheckpointValidators    []ValidatorWithBlsPub // validator addresses and BLS public key appended at checkpoint block
	Seal                    [ExtraSeal]byte       // the sealing block signature
}

func (extraData *HeaderExtraData) Encode(isShillin bool) []byte {
	var rawBytes []byte

	rawBytes = append(rawBytes, extraData.Vanity[:]...)
	if isShillin {
		rawBytes = append(rawBytes, extraData.HasFinalityVote)
		if extraData.HasFinalityVote == 1 {
			rawBytes = binary.LittleEndian.AppendUint64(rawBytes, uint64(extraData.FinalityVotedValidators))
			rawBytes = append(rawBytes, extraData.AggregatedFinalityVotes.Marshal()...)
		}
	}
	for _, validator := range extraData.CheckpointValidators {
		rawBytes = append(rawBytes, validator.Address.Bytes()...)
		if isShillin {
			rawBytes = append(rawBytes, validator.BlsPublicKey.Marshal()...)
		}
	}
	rawBytes = append(rawBytes, extraData.Seal[:]...)

	return rawBytes
}

type extraDataRLP struct {
	Vanity                  [ExtraVanity]byte
	HasFinalityVote         uint8
	FinalityVotedValidators FinalityVoteBitSet
	AggregatedFinalityVotes []byte
	CheckpointValidators    []validatorWithBlsPubRLP
	Seal                    [ExtraSeal]byte `rlp:"optional"`
}

type validatorWithBlsPubRLP struct {
	Address      common.Address
	BlsPublicKey []byte
}

func (extraData *HeaderExtraData) EncodeRLP() ([]byte, error) {
	ext := &extraDataRLP{
		Vanity:          extraData.Vanity,
		HasFinalityVote: extraData.HasFinalityVote,
		Seal:            extraData.Seal,
	}
	cp := make([]validatorWithBlsPubRLP, 0)
	for _, val := range extraData.CheckpointValidators {
		v := validatorWithBlsPubRLP{
			Address: val.Address,
		}
		if val.BlsPublicKey == nil {
			return nil, ErrNilBlsPublicKey
		}
		v.BlsPublicKey = val.BlsPublicKey.Marshal()
		cp = append(cp, v)
	}
	ext.CheckpointValidators = cp
	if extraData.HasFinalityVote != 0 && extraData.HasFinalityVote != 1 {
		return nil, ErrInvalidHasFinalityVote
	}
	if extraData.HasFinalityVote == 1 {
		ext.FinalityVotedValidators = extraData.FinalityVotedValidators
		if extraData.AggregatedFinalityVotes == nil {
			return nil, ErrNilAggregatedFinalityVotes
		}
		ext.AggregatedFinalityVotes = extraData.AggregatedFinalityVotes.Marshal()
	}
	return rlp.EncodeToBytes(ext)
}

func DecodeExtraRLP(enc []byte) (*HeaderExtraData, error) {
	var err error
	dec := &extraDataRLP{}
	if err := rlp.DecodeBytes(enc, dec); err != nil {
		return nil, err
	}
	ret := &HeaderExtraData{
		Vanity:          dec.Vanity,
		HasFinalityVote: dec.HasFinalityVote,
		Seal:            dec.Seal,
	}
	cp := make([]ValidatorWithBlsPub, 0)
	for _, val := range dec.CheckpointValidators {
		v := ValidatorWithBlsPub{
			Address: val.Address,
		}
		v.BlsPublicKey, err = blst.PublicKeyFromBytes(val.BlsPublicKey)
		if err != nil {
			return nil, err
		}
		cp = append(cp, v)
	}
	ret.CheckpointValidators = cp
	if dec.HasFinalityVote != 1 && dec.HasFinalityVote != 0 {
		return nil, ErrInvalidHasFinalityVote
	}
	if dec.HasFinalityVote == 1 {
		ret.FinalityVotedValidators = dec.FinalityVotedValidators
		ret.AggregatedFinalityVotes, err = blst.SignatureFromBytes(dec.AggregatedFinalityVotes)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func DecodeExtra(rawBytes []byte, isShillin bool) (*HeaderExtraData, error) {
	var (
		extraData       HeaderExtraData
		currentPosition int
		err             error
	)

	rawBytesLength := len(rawBytes)
	if rawBytesLength < ExtraVanity {
		return nil, ErrMissingVanity
	}

	copy(extraData.Vanity[:], rawBytes[:ExtraVanity])
	currentPosition += ExtraVanity

	if isShillin {
		if rawBytesLength-currentPosition < 1 {
			return nil, ErrMissingHasFinalityVote
		}

		extraData.HasFinalityVote = rawBytes[currentPosition]
		currentPosition += 1

		if extraData.HasFinalityVote != 1 && extraData.HasFinalityVote != 0 {
			return nil, ErrInvalidHasFinalityVote
		}

		if extraData.HasFinalityVote == 1 {
			if rawBytesLength-currentPosition < finalityVoteBitSetByteLength {
				return nil, ErrMissingFinalityVoteBitSet
			}
			extraData.FinalityVotedValidators = FinalityVoteBitSet(
				binary.LittleEndian.Uint64(rawBytes[currentPosition : currentPosition+finalityVoteBitSetByteLength]),
			)
			currentPosition += finalityVoteBitSetByteLength

			if rawBytesLength-currentPosition < params.BLSSignatureLength {
				return nil, ErrMissingFinalitySignature
			}
			extraData.AggregatedFinalityVotes, err = blst.SignatureFromBytes(
				rawBytes[currentPosition : currentPosition+params.BLSSignatureLength],
			)
			if err != nil {
				return nil, err
			}
			currentPosition += params.BLSSignatureLength
		}
	}

	if rawBytesLength-currentPosition < ExtraSeal {
		return nil, ErrMissingSignature
	}

	checkpointValidatorsLength := rawBytesLength - currentPosition - ExtraSeal
	extraData.CheckpointValidators, err = ParseCheckpointData(
		rawBytes[currentPosition:currentPosition+checkpointValidatorsLength],
		isShillin,
	)
	if err != nil {
		return nil, err
	}
	currentPosition += checkpointValidatorsLength

	copy(extraData.Seal[:], rawBytes[currentPosition:])

	return &extraData, nil
}

// ParseCheckpointData retrieves the list of validator addresses and finality voter's public keys
// at the checkpoint block
func ParseCheckpointData(checkpointData []byte, isShillin bool) ([]ValidatorWithBlsPub, error) {
	var (
		lengthPerValidator int
		extraData          []ValidatorWithBlsPub
		currentPosition    int
		err                error
	)

	if isShillin {
		lengthPerValidator = common.AddressLength + params.BLSPubkeyLength
	} else {
		lengthPerValidator = common.AddressLength
	}

	if len(checkpointData)%lengthPerValidator != 0 {
		return nil, ErrInvalidSpanValidators
	}

	numValidators := len(checkpointData) / lengthPerValidator
	extraData = make([]ValidatorWithBlsPub, numValidators)
	for i := 0; i < numValidators; i++ {
		copy(
			extraData[i].Address[:],
			checkpointData[currentPosition:currentPosition+common.AddressLength],
		)
		currentPosition += common.AddressLength

		if isShillin {
			extraData[i].BlsPublicKey, err = blst.PublicKeyFromBytes(
				checkpointData[currentPosition : currentPosition+params.BLSPubkeyLength],
			)
			if err != nil {
				return nil, err
			}
			currentPosition += params.BLSPubkeyLength
		}
	}

	return extraData, nil
}
