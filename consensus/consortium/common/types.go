package common

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// SignerFn is a signer callback function to request the wallet to sign the hash of the given data
type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

// SignerTxFn is a signer callback function to request the wallet to sign the given transaction.
type SignerTxFn func(accounts.Account, *types.Transaction, *big.Int) (*types.Transaction, error)

// ConsortiumAdapter defines a small collection of methods needed to access the private
// methods between consensus engines
type ConsortiumAdapter interface {
	GetRecents(chain consensus.ChainHeaderReader, number uint64) map[uint64]common.Address
}
