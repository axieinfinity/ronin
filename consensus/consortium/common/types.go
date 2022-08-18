package common

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// SignerFn is a Signer callback function to request a Header to be signed by a
// backing account.
type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

type SignerTxFn func(accounts.Account, *types.Transaction, *big.Int) (*types.Transaction, error)
