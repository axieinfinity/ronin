package core

import (
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

const (
	BLSPublicKeyLength = 48
	BLSSignatureLength = 96

	MaxAttestationExtraLength = 256
)

type BLSPublicKey [BLSPublicKeyLength]byte
type BLSSignature [BLSSignatureLength]byte
type ValidatorsBitSet uint64

// VoteData represents the vote range that validator voted for fast finality.
type VoteData struct {
	BlockNumber uint64      // The target block number which validator wants to vote for.
	BlockHash   common.Hash // The block hash of the target block.
}

// Hash returns the hash of the vote data.
func (d *VoteData) Hash() common.Hash { return rlpHash(d) }

// VoteEnvelope represents the vote of a single validator.
type VoteEnvelope struct {
	VoteAddress BLSPublicKey // The BLS public key of the validator.
	Signature   BLSSignature // Validator's signature for the vote data.
	Data        *VoteData    // The vote data for fast finality.

	// caches
	hash atomic.Value
}

// VoteAttestation represents the votes of super majority validators.
type VoteAttestation struct {
	VoteAddressSet ValidatorsBitSet // The bitset marks the voted validators.
	AggSignature   BLSSignature     // The aggregated BLS signature of the voted validators' signatures.
	Data           *VoteData        // The vote data for fast finality.
	Extra          []byte           // Reserved for future usage.
}

// Hash returns the vote's hash.
func (v *VoteEnvelope) Hash() common.Hash {
	if hash := v.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}

	h := v.calcVoteHash()
	v.hash.Store(h)
	return h
}

func (v *VoteEnvelope) calcVoteHash() common.Hash {
	vote := struct {
		VoteAddress BLSPublicKey
		Signature   BLSSignature
		Data        *VoteData
	}{v.VoteAddress, v.Signature, v.Data}
	return rlpHash(vote)
}

func (b BLSPublicKey) Bytes() []byte { return b[:] }

// Verify vote using BLS.
func (vote *VoteEnvelope) Verify() error {
	blsPubKey, err := bls.PublicKeyFromBytes(vote.VoteAddress[:])
	if err != nil {
		return errors.Wrap(err, "convert public key from bytes to bls failed")
	}

	sig, err := bls.SignatureFromBytes(vote.Signature[:])
	if err != nil {
		return errors.Wrap(err, "invalid signature")
	}

	voteDataHash := vote.Data.Hash()
	if !sig.Verify(blsPubKey, voteDataHash[:]) {
		return errors.New("verify bls signature failed.")
	}
	return nil
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
