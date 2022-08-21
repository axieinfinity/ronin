// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validators

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorsMetaData contains all meta data concerning the Validators contract.
var ValidatorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentValidatorSetMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"}],\"name\":\"depositReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ValidatorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorsMetaData.ABI instead.
var ValidatorsABI = ValidatorsMetaData.ABI

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
	ValidatorsFilterer   // Log filterer for contract events
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsSession struct {
	Contract     *Validators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsCallerSession struct {
	Contract *ValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsTransactorSession struct {
	Contract     *ValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsRaw struct {
	Contract *Validators // Generic contract binding to access the raw methods on
}

// ValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsCallerRaw struct {
	Contract *ValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsTransactorRaw struct {
	Contract *ValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidators creates a new instance of Validators, bound to a specific deployed contract.
func NewValidators(address common.Address, backend bind.ContractBackend) (*Validators, error) {
	contract, err := bindValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// NewValidatorsFilterer creates a new log filterer instance of Validators, bound to a specific deployed contract.
func NewValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsFilterer, error) {
	contract, err := bindValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsFilterer{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.ValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transact(opts, method, params...)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, uint256 totalAmount, uint256 jailTime, uint256 reward)
func (_Validators *ValidatorsCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	TotalAmount      *big.Int
	JailTime         *big.Int
	Reward           *big.Int
}, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "currentValidatorSet", arg0)

	outstruct := new(struct {
		ConsensusAddress common.Address
		FeeAddress       common.Address
		TotalAmount      *big.Int
		JailTime         *big.Int
		Reward           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FeeAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.JailTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, uint256 totalAmount, uint256 jailTime, uint256 reward)
func (_Validators *ValidatorsSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	TotalAmount      *big.Int
	JailTime         *big.Int
	Reward           *big.Int
}, error) {
	return _Validators.Contract.CurrentValidatorSet(&_Validators.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, uint256 totalAmount, uint256 jailTime, uint256 reward)
func (_Validators *ValidatorsCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	TotalAmount      *big.Int
	JailTime         *big.Int
	Reward           *big.Int
}, error) {
	return _Validators.Contract.CurrentValidatorSet(&_Validators.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validators *ValidatorsCaller) CurrentValidatorSetMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "currentValidatorSetMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validators *ValidatorsSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _Validators.Contract.CurrentValidatorSetMap(&_Validators.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validators *ValidatorsCallerSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _Validators.Contract.CurrentValidatorSetMap(&_Validators.CallOpts, arg0)
}

// GetLastUpdated is a free data retrieval call binding the contract method 0x78121dd4.
//
// Solidity: function getLastUpdated() view returns(uint256)
func (_Validators *ValidatorsCaller) GetLastUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getLastUpdated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdated is a free data retrieval call binding the contract method 0x78121dd4.
//
// Solidity: function getLastUpdated() view returns(uint256)
func (_Validators *ValidatorsSession) GetLastUpdated() (*big.Int, error) {
	return _Validators.Contract.GetLastUpdated(&_Validators.CallOpts)
}

// GetLastUpdated is a free data retrieval call binding the contract method 0x78121dd4.
//
// Solidity: function getLastUpdated() view returns(uint256)
func (_Validators *ValidatorsCallerSession) GetLastUpdated() (*big.Int, error) {
	return _Validators.Contract.GetLastUpdated(&_Validators.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validators *ValidatorsCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validators *ValidatorsSession) GetValidators() ([]common.Address, error) {
	return _Validators.Contract.GetValidators(&_Validators.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetValidators() ([]common.Address, error) {
	return _Validators.Contract.GetValidators(&_Validators.CallOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0xdb8ec38c.
//
// Solidity: function addNode(address consensusAddress, address feeAddress) returns()
func (_Validators *ValidatorsTransactor) AddNode(opts *bind.TransactOpts, consensusAddress common.Address, feeAddress common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addNode", consensusAddress, feeAddress)
}

// AddNode is a paid mutator transaction binding the contract method 0xdb8ec38c.
//
// Solidity: function addNode(address consensusAddress, address feeAddress) returns()
func (_Validators *ValidatorsSession) AddNode(consensusAddress common.Address, feeAddress common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddNode(&_Validators.TransactOpts, consensusAddress, feeAddress)
}

// AddNode is a paid mutator transaction binding the contract method 0xdb8ec38c.
//
// Solidity: function addNode(address consensusAddress, address feeAddress) returns()
func (_Validators *ValidatorsTransactorSession) AddNode(consensusAddress common.Address, feeAddress common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddNode(&_Validators.TransactOpts, consensusAddress, feeAddress)
}

// DepositReward is a paid mutator transaction binding the contract method 0x6ffa4dc1.
//
// Solidity: function depositReward(address valAddr) payable returns()
func (_Validators *ValidatorsTransactor) DepositReward(opts *bind.TransactOpts, valAddr common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "depositReward", valAddr)
}

// DepositReward is a paid mutator transaction binding the contract method 0x6ffa4dc1.
//
// Solidity: function depositReward(address valAddr) payable returns()
func (_Validators *ValidatorsSession) DepositReward(valAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.DepositReward(&_Validators.TransactOpts, valAddr)
}

// DepositReward is a paid mutator transaction binding the contract method 0x6ffa4dc1.
//
// Solidity: function depositReward(address valAddr) payable returns()
func (_Validators *ValidatorsTransactorSession) DepositReward(valAddr common.Address) (*types.Transaction, error) {
	return _Validators.Contract.DepositReward(&_Validators.TransactOpts, valAddr)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x2d497ba2.
//
// Solidity: function updateValidators() returns()
func (_Validators *ValidatorsTransactor) UpdateValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateValidators")
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x2d497ba2.
//
// Solidity: function updateValidators() returns()
func (_Validators *ValidatorsSession) UpdateValidators() (*types.Transaction, error) {
	return _Validators.Contract.UpdateValidators(&_Validators.TransactOpts)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x2d497ba2.
//
// Solidity: function updateValidators() returns()
func (_Validators *ValidatorsTransactorSession) UpdateValidators() (*types.Transaction, error) {
	return _Validators.Contract.UpdateValidators(&_Validators.TransactOpts)
}
