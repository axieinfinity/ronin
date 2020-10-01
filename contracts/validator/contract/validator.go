// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_acknowledgement\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denom\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"acknowledgement\",\"outputs\":[{\"internalType\":\"contractAcknowledgement\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteCount\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"num\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"updateQuorum\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Acknowledgement is a free data retrieval call binding the contract method 0xfc819750.
//
// Solidity: function acknowledgement() view returns(address)
func (_Contract *ContractCaller) Acknowledgement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "acknowledgement")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acknowledgement is a free data retrieval call binding the contract method 0xfc819750.
//
// Solidity: function acknowledgement() view returns(address)
func (_Contract *ContractSession) Acknowledgement() (common.Address, error) {
	return _Contract.Contract.Acknowledgement(&_Contract.CallOpts)
}

// Acknowledgement is a free data retrieval call binding the contract method 0xfc819750.
//
// Solidity: function acknowledgement() view returns(address)
func (_Contract *ContractCallerSession) Acknowledgement() (common.Address, error) {
	return _Contract.Contract.Acknowledgement(&_Contract.CallOpts)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteCount) view returns(bool)
func (_Contract *ContractCaller) CheckThreshold(opts *bind.CallOpts, _voteCount *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "checkThreshold", _voteCount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteCount) view returns(bool)
func (_Contract *ContractSession) CheckThreshold(_voteCount *big.Int) (bool, error) {
	return _Contract.Contract.CheckThreshold(&_Contract.CallOpts, _voteCount)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteCount) view returns(bool)
func (_Contract *ContractCallerSession) CheckThreshold(_voteCount *big.Int) (bool, error) {
	return _Contract.Contract.CheckThreshold(&_Contract.CallOpts, _voteCount)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(uint256)
func (_Contract *ContractCaller) Denom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "denom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(uint256)
func (_Contract *ContractSession) Denom() (*big.Int, error) {
	return _Contract.Contract.Denom(&_Contract.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(uint256)
func (_Contract *ContractCallerSession) Denom() (*big.Int, error) {
	return _Contract.Contract.Denom(&_Contract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validators)
func (_Contract *ContractCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validators)
func (_Contract *ContractSession) GetValidators() ([]common.Address, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validators)
func (_Contract *ContractCallerSession) GetValidators() ([]common.Address, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Contract *ContractCaller) IsValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Contract *ContractSession) IsValidator(_addr common.Address) (bool, error) {
	return _Contract.Contract.IsValidator(&_Contract.CallOpts, _addr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Contract *ContractCallerSession) IsValidator(_addr common.Address) (bool, error) {
	return _Contract.Contract.IsValidator(&_Contract.CallOpts, _addr)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Contract *ContractCaller) Num(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "num")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Contract *ContractSession) Num() (*big.Int, error) {
	return _Contract.Contract.Num(&_Contract.CallOpts)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Contract *ContractCallerSession) Num() (*big.Int, error) {
	return _Contract.Contract.Num(&_Contract.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Contract *ContractCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Contract *ContractSession) ValidatorCount() (*big.Int, error) {
	return _Contract.Contract.ValidatorCount(&_Contract.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Contract *ContractCallerSession) ValidatorCount() (*big.Int, error) {
	return _Contract.Contract.ValidatorCount(&_Contract.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_Contract *ContractCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_Contract *ContractSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Validators(&_Contract.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Validators(&_Contract.CallOpts, arg0)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4b561753.
//
// Solidity: function addValidator(uint256 _id, address _validator) returns()
func (_Contract *ContractTransactor) AddValidator(opts *bind.TransactOpts, _id *big.Int, _validator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addValidator", _id, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4b561753.
//
// Solidity: function addValidator(uint256 _id, address _validator) returns()
func (_Contract *ContractSession) AddValidator(_id *big.Int, _validator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddValidator(&_Contract.TransactOpts, _id, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4b561753.
//
// Solidity: function addValidator(uint256 _id, address _validator) returns()
func (_Contract *ContractTransactorSession) AddValidator(_id *big.Int, _validator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddValidator(&_Contract.TransactOpts, _id, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x6a0cd1f5.
//
// Solidity: function removeValidator(uint256 _id, address _validator) returns()
func (_Contract *ContractTransactor) RemoveValidator(opts *bind.TransactOpts, _id *big.Int, _validator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeValidator", _id, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x6a0cd1f5.
//
// Solidity: function removeValidator(uint256 _id, address _validator) returns()
func (_Contract *ContractSession) RemoveValidator(_id *big.Int, _validator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveValidator(&_Contract.TransactOpts, _id, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x6a0cd1f5.
//
// Solidity: function removeValidator(uint256 _id, address _validator) returns()
func (_Contract *ContractTransactorSession) RemoveValidator(_id *big.Int, _validator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveValidator(&_Contract.TransactOpts, _id, _validator)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x53727d26.
//
// Solidity: function updateQuorum(uint256 _id, uint256 _numerator, uint256 _denominator) returns()
func (_Contract *ContractTransactor) UpdateQuorum(opts *bind.TransactOpts, _id *big.Int, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateQuorum", _id, _numerator, _denominator)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x53727d26.
//
// Solidity: function updateQuorum(uint256 _id, uint256 _numerator, uint256 _denominator) returns()
func (_Contract *ContractSession) UpdateQuorum(_id *big.Int, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateQuorum(&_Contract.TransactOpts, _id, _numerator, _denominator)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x53727d26.
//
// Solidity: function updateQuorum(uint256 _id, uint256 _numerator, uint256 _denominator) returns()
func (_Contract *ContractTransactorSession) UpdateQuorum(_id *big.Int, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateQuorum(&_Contract.TransactOpts, _id, _numerator, _denominator)
}

// ContractThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the Contract contract.
type ContractThresholdUpdatedIterator struct {
	Event *ContractThresholdUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractThresholdUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractThresholdUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractThresholdUpdated represents a ThresholdUpdated event raised by the Contract contract.
type ContractThresholdUpdated struct {
	Id                  *big.Int
	Numerator           *big.Int
	Denominator         *big.Int
	PreviousNumerator   *big.Int
	PreviousDenominator *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterThresholdUpdated is a free log retrieval operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 _id, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_Contract *ContractFilterer) FilterThresholdUpdated(opts *bind.FilterOpts, numerator []*big.Int, denominator []*big.Int) (*ContractThresholdUpdatedIterator, error) {

	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ThresholdUpdated", numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractThresholdUpdatedIterator{contract: _Contract.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 _id, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_Contract *ContractFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *ContractThresholdUpdated, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ThresholdUpdated", numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractThresholdUpdated)
				if err := _Contract.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThresholdUpdated is a log parse operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 _id, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_Contract *ContractFilterer) ParseThresholdUpdated(log types.Log) (*ContractThresholdUpdated, error) {
	event := new(ContractThresholdUpdated)
	if err := _Contract.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the Contract contract.
type ContractValidatorAddedIterator struct {
	Event *ContractValidatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorAdded represents a ValidatorAdded event raised by the Contract contract.
type ContractValidatorAdded struct {
	Id        *big.Int
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0x7429a06e9412e469f0d64f9d222640b0af359f556b709e2913588c227851b88d.
//
// Solidity: event ValidatorAdded(uint256 _id, address indexed validator)
func (_Contract *ContractFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*ContractValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorAddedIterator{contract: _Contract.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0x7429a06e9412e469f0d64f9d222640b0af359f556b709e2913588c227851b88d.
//
// Solidity: event ValidatorAdded(uint256 _id, address indexed validator)
func (_Contract *ContractFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ContractValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorAdded)
				if err := _Contract.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorAdded is a log parse operation binding the contract event 0x7429a06e9412e469f0d64f9d222640b0af359f556b709e2913588c227851b88d.
//
// Solidity: event ValidatorAdded(uint256 _id, address indexed validator)
func (_Contract *ContractFilterer) ParseValidatorAdded(log types.Log) (*ContractValidatorAdded, error) {
	event := new(ContractValidatorAdded)
	if err := _Contract.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Contract contract.
type ContractValidatorRemovedIterator struct {
	Event *ContractValidatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorRemoved represents a ValidatorRemoved event raised by the Contract contract.
type ContractValidatorRemoved struct {
	Id        *big.Int
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0x7126bef88d1149ccdff9681ed5aecd3ba5ae70c96517551de250af09cebd1a0b.
//
// Solidity: event ValidatorRemoved(uint256 _id, address indexed validator)
func (_Contract *ContractFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*ContractValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorRemovedIterator{contract: _Contract.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0x7126bef88d1149ccdff9681ed5aecd3ba5ae70c96517551de250af09cebd1a0b.
//
// Solidity: event ValidatorRemoved(uint256 _id, address indexed validator)
func (_Contract *ContractFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ContractValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorRemoved)
				if err := _Contract.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRemoved is a log parse operation binding the contract event 0x7126bef88d1149ccdff9681ed5aecd3ba5ae70c96517551de250af09cebd1a0b.
//
// Solidity: event ValidatorRemoved(uint256 _id, address indexed validator)
func (_Contract *ContractFilterer) ParseValidatorRemoved(log types.Log) (*ContractValidatorRemoved, error) {
	event := new(ContractValidatorRemoved)
	if err := _Contract.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
