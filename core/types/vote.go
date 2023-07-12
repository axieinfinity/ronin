package types

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

type BLSPublicKey [params.BLSPubkeyLength]byte
type BLSSignature [params.BLSSignatureLength]byte
type ValidatorsBitSet uint64

// VoteData represents the vote range that validator voted for fast finality.
type VoteData struct {
	TargetNumber uint64      // The target block number which validator wants to vote for.
	TargetHash   common.Hash // The block hash of the target block.
}

// Hash returns the hash of the vote data.
func (d *VoteData) Hash() common.Hash { return rlpHash(d) }

// RawVoteEnvelope is VoteEnvelop without cached hash
type RawVoteEnvelope struct {
	PublicKey BLSPublicKey // The BLS public key of the validator.
	Signature BLSSignature // Validator's signature for the vote data.
	Data      *VoteData    // The vote data for fast finality.
}

// VoteEnvelope represents the vote of a single validator.
type VoteEnvelope struct {
	RawVoteEnvelope

	// caches
	hash atomic.Value
}

func (v *VoteEnvelope) Raw() *RawVoteEnvelope {
	return &v.RawVoteEnvelope
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
		PublicKey BLSPublicKey
		Signature BLSSignature
		Data      *VoteData
	}{v.PublicKey, v.Signature, v.Data}
	return rlpHash(vote)
}

func (b BLSPublicKey) Bytes() []byte { return b[:] }

// Verify vote using BLS.
func (vote *VoteEnvelope) Verify() error {
	blsPubKey, err := bls.PublicKeyFromBytes(vote.PublicKey[:])
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
