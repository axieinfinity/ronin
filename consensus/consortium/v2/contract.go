package v2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

// Interaction with contract/account
func (c *Consortium) getCurrentValidators() ([]common.Address, error) {
	return c.validatorSC.Validators()
}

// slash spoiled validators
func (c *Consortium) distributeIncoming(
	state *state.StateDB,
	header *types.Header,
) error {
	coinbase := header.Coinbase
	balance := state.GetBalance(consensus.SystemAddress)
	if balance.Cmp(common.Big0) <= 0 {
		return nil
	}
	state.SetBalance(consensus.SystemAddress, big.NewInt(0))
	state.AddBalance(coinbase, balance)

	log.Trace("distribute to validator contract", "block hash", header.Hash(), "amount", balance)
	return c.validatorSC.Deposit(c.val, c.val, balance)
}

// slash spoiled validators
func (c *Consortium) slash(spoiledVal common.Address) error {
	return c.validatorSC.Slash(spoiledVal)
}
