package common

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/systemcontracts"
	"github.com/ethereum/go-ethereum/ethdb"
	chainParams "github.com/ethereum/go-ethereum/params"
)

// SignerFn is a signer callback function to request a header to be signed by a
// backing account.
type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

type ContractIntegrator struct {
	valSC systemcontracts.ValidatorSC
}

func NewContractIntegrator(bc consensus.ChainHeaderReader, config chainParams.ChainConfig, db ethdb.Database) *ContractIntegrator {
	simBackend := backends.NewSimulatedBackendWithBC(bc.(backends.BlockchainContext), db)
	return &ContractIntegrator{}
}

func (c *ContractIntegrator) Deposit() {

}
