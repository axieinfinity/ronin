package common

import (
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ExtraSeal                        = crypto.SignatureLength
	ExtraVanity                      = 32
	MaxFinalityVotePercentage uint16 = 10_000
)

var (
	// ErrMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	ErrMissingSignature = errors.New("extra-data 65 byte signature suffix missing")

	// ErrUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	ErrUnknownBlock = errors.New("unknown block")

	// ErrMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	ErrMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// ErrExtraValidators is returned if non-sprint-end block contain validator data in
	// their extra-data fields.
	ErrExtraValidators = errors.New("non-sprint-end block contains extra validator list")

	// ErrInvalidSpanValidators is returned if a block contains an
	// invalid list of validators (i.e. non divisible by 20 bytes).
	ErrInvalidSpanValidators = errors.New("invalid validator list on sprint end block")

	// ErrInvalidMixDigest is returned if a block's mix digest is non-zero.
	ErrInvalidMixDigest = errors.New("non-zero mix digest")

	// ErrInvalidUncleHash is returned if a block contains an non-empty uncle list.
	ErrInvalidUncleHash = errors.New("non empty uncle hash")

	// ErrInvalidDifficulty is returned if the difficulty of a block neither 1 or 2.
	ErrInvalidDifficulty = errors.New("invalid difficulty")

	// ErrInvalidCheckpointSigners is returned if a checkpoint block contains an
	// invalid list of signers (i.e. non divisible by 20 bytes).
	ErrInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")

	// ErrRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	ErrRecentlySigned = errors.New("signed recently, must wait for others")

	// ErrWrongDifficulty is returned if the difficulty of a block doesn't match the
	// turn of the signer.
	ErrWrongDifficulty = errors.New("wrong difficulty")
)
