package common

import (
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ExtraSeal = crypto.SignatureLength // Fixed number of extra-data suffix bytes reserved for signer seal
)

var (
	// ErrMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	ErrMissingSignature = errors.New("extra-data 65 byte signature suffix missing")
)
