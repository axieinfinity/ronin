package systemcontracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/systemcontracts/generated_contracts/validators"
	"math/big"
)

type ValidatorSC struct {
	address  common.Address
	contract *validators.Validators
}

func NewValidatorSC(contractAddr common.Address, backend bind.ContractBackend) (*ValidatorSC, error) {
	c, err := validators.NewValidators(contractAddr, backend)
	if err != nil {
		return nil, err
	}

	return &ValidatorSC{
		address:  contractAddr,
		contract: c,
	}, nil
}

func (c *ValidatorSC) ContractAddr() common.Address {
	return c.address
}

func (c *ValidatorSC) Contract() *validators.Validators {
	return c.contract
}

func (c *ValidatorSC) Validators() ([]common.Address, error) {
	addresses, err := c.contract.GetValidators(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (c *ValidatorSC) Deposit(from, to common.Address, value *big.Int) error {
	_, err := c.contract.DepositReward(&bind.TransactOpts{
		From:     from,
		GasPrice: big.NewInt(0),
		Value:    value,
	}, to)
	if err != nil {
		return err
	}

	return nil
}

func (c *ValidatorSC) Slash(to common.Address) error {
	return nil
}
