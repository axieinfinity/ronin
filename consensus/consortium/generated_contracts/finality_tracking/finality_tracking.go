// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package finalityTracking

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
	_ = abi.ConvertType
)

// FinalityTrackingMetaData contains all meta data concerning the FinalityTracking contract.
var FinalityTrackingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"}],\"name\":\"recordFinality\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FinalityTrackingABI is the input ABI used to generate the binding from.
// Deprecated: Use FinalityTrackingMetaData.ABI instead.
var FinalityTrackingABI = FinalityTrackingMetaData.ABI

// FinalityTracking is an auto generated Go binding around an Ethereum contract.
type FinalityTracking struct {
	FinalityTrackingCaller     // Read-only binding to the contract
	FinalityTrackingTransactor // Write-only binding to the contract
	FinalityTrackingFilterer   // Log filterer for contract events
}

// FinalityTrackingCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalityTrackingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityTrackingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalityTrackingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityTrackingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalityTrackingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityTrackingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalityTrackingSession struct {
	Contract     *FinalityTracking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FinalityTrackingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalityTrackingCallerSession struct {
	Contract *FinalityTrackingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FinalityTrackingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalityTrackingTransactorSession struct {
	Contract     *FinalityTrackingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FinalityTrackingRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalityTrackingRaw struct {
	Contract *FinalityTracking // Generic contract binding to access the raw methods on
}

// FinalityTrackingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalityTrackingCallerRaw struct {
	Contract *FinalityTrackingCaller // Generic read-only contract binding to access the raw methods on
}

// FinalityTrackingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalityTrackingTransactorRaw struct {
	Contract *FinalityTrackingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalityTracking creates a new instance of FinalityTracking, bound to a specific deployed contract.
func NewFinalityTracking(address common.Address, backend bind.ContractBackend) (*FinalityTracking, error) {
	contract, err := bindFinalityTracking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalityTracking{FinalityTrackingCaller: FinalityTrackingCaller{contract: contract}, FinalityTrackingTransactor: FinalityTrackingTransactor{contract: contract}, FinalityTrackingFilterer: FinalityTrackingFilterer{contract: contract}}, nil
}

// NewFinalityTrackingCaller creates a new read-only instance of FinalityTracking, bound to a specific deployed contract.
func NewFinalityTrackingCaller(address common.Address, caller bind.ContractCaller) (*FinalityTrackingCaller, error) {
	contract, err := bindFinalityTracking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalityTrackingCaller{contract: contract}, nil
}

// NewFinalityTrackingTransactor creates a new write-only instance of FinalityTracking, bound to a specific deployed contract.
func NewFinalityTrackingTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalityTrackingTransactor, error) {
	contract, err := bindFinalityTracking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalityTrackingTransactor{contract: contract}, nil
}

// NewFinalityTrackingFilterer creates a new log filterer instance of FinalityTracking, bound to a specific deployed contract.
func NewFinalityTrackingFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalityTrackingFilterer, error) {
	contract, err := bindFinalityTracking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalityTrackingFilterer{contract: contract}, nil
}

// bindFinalityTracking binds a generic wrapper to an already deployed contract.
func bindFinalityTracking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FinalityTrackingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalityTracking *FinalityTrackingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalityTracking.Contract.FinalityTrackingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalityTracking *FinalityTrackingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityTracking.Contract.FinalityTrackingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalityTracking *FinalityTrackingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalityTracking.Contract.FinalityTrackingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalityTracking *FinalityTrackingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalityTracking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalityTracking *FinalityTrackingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityTracking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalityTracking *FinalityTrackingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalityTracking.Contract.contract.Transact(opts, method, params...)
}

// RecordFinality is a paid mutator transaction binding the contract method 0xc245db0f.
//
// Solidity: function recordFinality(address[] voters) returns()
func (_FinalityTracking *FinalityTrackingTransactor) RecordFinality(opts *bind.TransactOpts, voters []common.Address) (*types.Transaction, error) {
	return _FinalityTracking.contract.Transact(opts, "recordFinality", voters)
}

// RecordFinality is a paid mutator transaction binding the contract method 0xc245db0f.
//
// Solidity: function recordFinality(address[] voters) returns()
func (_FinalityTracking *FinalityTrackingSession) RecordFinality(voters []common.Address) (*types.Transaction, error) {
	return _FinalityTracking.Contract.RecordFinality(&_FinalityTracking.TransactOpts, voters)
}

// RecordFinality is a paid mutator transaction binding the contract method 0xc245db0f.
//
// Solidity: function recordFinality(address[] voters) returns()
func (_FinalityTracking *FinalityTrackingTransactorSession) RecordFinality(voters []common.Address) (*types.Transaction, error) {
	return _FinalityTracking.Contract.RecordFinality(&_FinalityTracking.TransactOpts, voters)
}
