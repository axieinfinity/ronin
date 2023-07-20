package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
)

// SignerFn is a signer callback function to request the wallet to sign the hash of the given data
type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

// SignerTxFn is a signer callback function to request the wallet to sign the given transaction.
type SignerTxFn func(accounts.Account, *types.Transaction, *big.Int) (*types.Transaction, error)

type BaseSnapshot struct {
	Number     uint64                      `json:"number"`     // Block number where the snapshot was created
	Hash       common.Hash                 `json:"hash"`       // Block hash where the snapshot was created
	SignerSet  map[common.Address]struct{} `json:"signerSet"`  // Set of authorized signers at this moment
	SignerList []common.Address            `json:"signerList"` // List of authorized signers at this moment
	Recents    map[uint64]common.Address   `json:"recents"`    // Set of recent signers for spam protections
}

// ConsortiumAdapter defines a small collection of methods needed to access the private
// methods between consensus engines
type ConsortiumAdapter interface {
	GetSnapshot(chain consensus.ChainHeaderReader, number uint64, parents []*types.Header) *BaseSnapshot
}
