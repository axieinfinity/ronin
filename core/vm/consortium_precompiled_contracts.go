package vm

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"math/rand"
	"os"
	"strings"
)

var (
	consortiumLogAbi           = `[{"inputs":[{"internalType":"string","name":"message","type":"string"}],"name":"log","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	consortiumSortValidatorAbi = `[{"inputs":[],"name":"getValidators","outputs":[{"internalType":"address[]","name":"_validatorList","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"limit","type":"uint256"}],"name":"sortValidators","outputs":[{"internalType":"address[]","name":"validators","type":"address[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address[]","name":"_poolList","type":"address[]"}],"name":"totalBalances","outputs":[{"internalType":"uint256[]","name":"_balances","type":"uint256[]"}],"stateMutability":"view","type":"function"}]`
)

const (
	sortValidatorsMethod = "sortValidators"
	logMethod            = "log"
	getValidatorsMethod  = "getValidators"
	totalBalancesMethod  = "totalBalances"
)

func PrecompiledContractsConsortium(caller ContractRef, evm *EVM) map[common.Address]PrecompiledContract {
	return map[common.Address]PrecompiledContract{
		common.BytesToAddress([]byte{101}): &consortiumLog{},
		common.BytesToAddress([]byte{102}): &consortiumValidatorSorting{caller: caller, evm: evm},
	}
}

type consortiumLog struct{}

func (c *consortiumLog) RequiredGas(input []byte) uint64 {
	return 0
}

func (c *consortiumLog) Run(input []byte) ([]byte, error) {
	if os.Getenv("DEBUG") != "true" {
		return input, nil
	}
	_, method, args, err := loadMethodAndArgs(consortiumLogAbi, input)
	if err != nil {
		return nil, err
	}
	switch method.Name {
	case logMethod:
		if len(args) == 0 {
			return input, nil
		}
		if _, ok := args[0].(string); ok {
			log.Info("[consortiumLog] log message from smart contract", "message", args[0].(string))
		}
	}
	return input, nil
}

type consortiumValidatorSorting struct {
	caller ContractRef
	evm    *EVM
}

func (c *consortiumValidatorSorting) RequiredGas(input []byte) uint64 {
	return 0
}

func (c *consortiumValidatorSorting) Run(input []byte) ([]byte, error) {
	if c.evm.ChainConfig().ConsortiumV2Contracts == nil {
		return nil, errors.New("cannot find consortium v2 contracts")
	}
	if !c.evm.ChainConfig().ConsortiumV2Contracts.IsSystemContract(c.caller.Address()) {
		return nil, errors.New("unauthorized sender")
	}
	// get method, args from abi
	smcAbi, method, args, err := loadMethodAndArgs(consortiumSortValidatorAbi, input)
	if err != nil {
		return nil, err
	}
	if method.Name != sortValidatorsMethod {
		return nil, errors.New("invalid method")
	}
	if len(args) != 1 {
		return nil, errors.New(fmt.Sprintf("invalid arguments, expected 1 got %d", len(args)))
	}
	// cast args[0] to number
	limit, ok := args[0].(*big.Int)
	if !ok {
		return nil, errors.New("invalid argument type")
	}
	validators, err := loadValidators(c.evm, smcAbi, c.caller.Address())
	if err != nil {
		return nil, err
	}
	totalBalances, err := loadTotalBalances(c.evm, smcAbi, c.caller.Address(), validators)
	if err != nil {
		return nil, err
	}
	if len(validators) != len(totalBalances) {
		return nil, errors.New("balances and validators length mismatched")
	}
	sortValidators(validators, totalBalances)

	if limit.Int64() > int64(len(validators)) {
		limit = big.NewInt(int64(len(validators)))
	}

	return method.Outputs.Pack(validators[:limit.Int64()])
}

func sortValidators(validators []common.Address, balances []*big.Int) {
	if len(validators) < 2 {
		return
	}

	left, right := 0, len(validators)-1

	pivot := rand.Int() % len(validators)

	validators[pivot], validators[right] = validators[right], validators[pivot]
	balances[pivot], balances[right] = balances[right], balances[pivot]

	for i, _ := range validators {
		cmp := balances[i].Cmp(balances[right])
		addrsCmp := big.NewInt(0).SetBytes(validators[i].Bytes()).Cmp(big.NewInt(0).SetBytes(validators[right].Bytes())) > 0
		if cmp > 0 || (cmp == 0 && addrsCmp) {
			validators[left], validators[i] = validators[i], validators[left]
			balances[left], balances[i] = balances[i], balances[left]
			left++
		}
	}

	validators[left], validators[right] = validators[right], validators[left]
	balances[left], balances[right] = balances[right], balances[left]

	sortValidators(validators[:left], balances[:left])
	sortValidators(validators[left+1:], balances[left+1:])

	return
}

func loadValidators(evm *EVM, smcAbi abi.ABI, sender common.Address) ([]common.Address, error) {
	res, err := staticCall(evm, smcAbi, getValidatorsMethod, evm.ChainConfig().ConsortiumV2Contracts.RoninValidatorSet, sender)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(res[0], new([]common.Address)).(*[]common.Address), nil
}

func loadTotalBalances(evm *EVM, smcAbi abi.ABI, sender common.Address, validators []common.Address) ([]*big.Int, error) {
	res, err := staticCall(evm, smcAbi, totalBalancesMethod, evm.ChainConfig().ConsortiumV2Contracts.StakingContract, sender, validators)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(res[0], new([]*big.Int)).(*[]*big.Int), nil
}

func staticCall(evm *EVM, smcAbi abi.ABI, method string, contract, sender common.Address, args ...interface{}) ([]interface{}, error) {
	inputParams, err := smcAbi.Pack(method, args...)
	if err != nil {
		return nil, err
	}
	ret, _, err := evm.StaticCall(AccountRef(sender), contract, inputParams, math.MaxUint64/2)
	if err != nil {
		return nil, err
	}
	out, err := smcAbi.Unpack(method, ret)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, errors.New("data not found")
	}
	return out, nil
}

func loadMethodAndArgs(smcAbi string, input []byte) (abi.ABI, *abi.Method, []interface{}, error) {
	var (
		pAbi   abi.ABI
		err    error
		method *abi.Method
		args   []interface{}
	)
	if pAbi, err = abi.JSON(strings.NewReader(smcAbi)); err != nil {
		return abi.ABI{}, nil, nil, err
	}
	if method, err = pAbi.MethodById(input); err != nil {
		return abi.ABI{}, nil, nil, err
	}
	if args, err = method.Inputs.Unpack(input[4:]); err != nil {
		return abi.ABI{}, nil, nil, err
	}
	return pAbi, method, args, nil
}
