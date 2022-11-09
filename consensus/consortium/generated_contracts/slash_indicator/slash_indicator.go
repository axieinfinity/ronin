// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package slashIndicator

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

// SlashIndicatorMetaData contains all meta data concerning the SlashIndicator contract.
var SlashIndicatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"BailedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"missingVotesRatioTier1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"missingVotesRatioTier2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jailDurationForMissingVotesRatioTier2\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorSlashingConfigsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgeVotingThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgeVotingSlashAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeVotingSlashingConfigsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gainCreditScore\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxCreditScore\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bailOutCostMultiplier\",\"type\":\"uint256\"}],\"name\":\"CreditScoreConfigsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"creditScores\",\"type\":\"uint256[]\"}],\"name\":\"CreditScoresUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashDoubleSignAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"doubleSigningJailUntilBlock\",\"type\":\"uint256\"}],\"name\":\"DoubleSignSlashingConfigsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MaintenanceContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RoninGovernanceAdminContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RoninTrustedOrganizationContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIBaseSlash.SlashType\",\"name\":\"slashType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unavailabilityTier1Threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unavailabilityTier2Threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmountForUnavailabilityTier2Threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jailDurationForUnavailabilityTier2Threshold\",\"type\":\"uint256\"}],\"name\":\"UnavailabilitySlashingConfigsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"bailOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bailOutCostMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"currentUnavailabilityIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainCreditScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeOperatorSlashingConfigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeVotingSlashingConfigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"getBulkCreditScore\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_resultList\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getCreditScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreditScoreConfigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_gainCreditScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCreditScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bailOutCostMultiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDoubleSignSlashingConfigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"getUnavailabilityIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnavailabilitySlashingConfigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__validatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__maintenanceContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__roninTrustedOrganizationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__roninGovernanceAdminContract\",\"type\":\"address\"},{\"internalType\":\"uint256[3]\",\"name\":\"_bridgeOperatorSlashingConfigs\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_bridgeVotingSlashingConfigs\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_doubleSignSlashingConfigs\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_unavailabilitySlashingConfigs\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_creditScoreConfigs\",\"type\":\"uint256[3]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUnavailabilitySlashedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maintenanceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCreditScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompileValidateDoubleSignAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninGovernanceAdminContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninTrustedOrganizationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ratioTier1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ratioTier2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_jailDurationTier2\",\"type\":\"uint256\"}],\"name\":\"setBridgeOperatorSlashingConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"}],\"name\":\"setBridgeVotingSlashingConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gainCreditScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCreditScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bailOutCostMultiplier\",\"type\":\"uint256\"}],\"name\":\"setCreditScoreConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_jailUntilBlock\",\"type\":\"uint256\"}],\"name\":\"setDoubleSignSlashingConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMaintenanceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRoninGovernanceAdminContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRoninTrustedOrganizationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tier1Threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tier2Threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashAmountForTier2Threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_jailDurationForTier2Threshold\",\"type\":\"uint256\"}],\"name\":\"setUnavailabilitySlashingConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"slashBridgeVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensuAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_header1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_header2\",\"type\":\"bytes\"}],\"name\":\"slashDoubleSign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"slashUnavailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"updateCreditScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SlashIndicatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashIndicatorMetaData.ABI instead.
var SlashIndicatorABI = SlashIndicatorMetaData.ABI

// SlashIndicator is an auto generated Go binding around an Ethereum contract.
type SlashIndicator struct {
	SlashIndicatorCaller     // Read-only binding to the contract
	SlashIndicatorTransactor // Write-only binding to the contract
	SlashIndicatorFilterer   // Log filterer for contract events
}

// SlashIndicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashIndicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashIndicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashIndicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashIndicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashIndicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashIndicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashIndicatorSession struct {
	Contract     *SlashIndicator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashIndicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashIndicatorCallerSession struct {
	Contract *SlashIndicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SlashIndicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashIndicatorTransactorSession struct {
	Contract     *SlashIndicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SlashIndicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashIndicatorRaw struct {
	Contract *SlashIndicator // Generic contract binding to access the raw methods on
}

// SlashIndicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashIndicatorCallerRaw struct {
	Contract *SlashIndicatorCaller // Generic read-only contract binding to access the raw methods on
}

// SlashIndicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashIndicatorTransactorRaw struct {
	Contract *SlashIndicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashIndicator creates a new instance of SlashIndicator, bound to a specific deployed contract.
func NewSlashIndicator(address common.Address, backend bind.ContractBackend) (*SlashIndicator, error) {
	contract, err := bindSlashIndicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashIndicator{SlashIndicatorCaller: SlashIndicatorCaller{contract: contract}, SlashIndicatorTransactor: SlashIndicatorTransactor{contract: contract}, SlashIndicatorFilterer: SlashIndicatorFilterer{contract: contract}}, nil
}

// NewSlashIndicatorCaller creates a new read-only instance of SlashIndicator, bound to a specific deployed contract.
func NewSlashIndicatorCaller(address common.Address, caller bind.ContractCaller) (*SlashIndicatorCaller, error) {
	contract, err := bindSlashIndicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorCaller{contract: contract}, nil
}

// NewSlashIndicatorTransactor creates a new write-only instance of SlashIndicator, bound to a specific deployed contract.
func NewSlashIndicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashIndicatorTransactor, error) {
	contract, err := bindSlashIndicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorTransactor{contract: contract}, nil
}

// NewSlashIndicatorFilterer creates a new log filterer instance of SlashIndicator, bound to a specific deployed contract.
func NewSlashIndicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashIndicatorFilterer, error) {
	contract, err := bindSlashIndicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorFilterer{contract: contract}, nil
}

// bindSlashIndicator binds a generic wrapper to an already deployed contract.
func bindSlashIndicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SlashIndicatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashIndicator *SlashIndicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashIndicator.Contract.SlashIndicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashIndicator *SlashIndicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashIndicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashIndicator *SlashIndicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashIndicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashIndicator *SlashIndicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashIndicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashIndicator *SlashIndicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashIndicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashIndicator *SlashIndicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashIndicator.Contract.contract.Transact(opts, method, params...)
}

// BailOutCostMultiplier is a free data retrieval call binding the contract method 0x37c597ea.
//
// Solidity: function bailOutCostMultiplier() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) BailOutCostMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "bailOutCostMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BailOutCostMultiplier is a free data retrieval call binding the contract method 0x37c597ea.
//
// Solidity: function bailOutCostMultiplier() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) BailOutCostMultiplier() (*big.Int, error) {
	return _SlashIndicator.Contract.BailOutCostMultiplier(&_SlashIndicator.CallOpts)
}

// BailOutCostMultiplier is a free data retrieval call binding the contract method 0x37c597ea.
//
// Solidity: function bailOutCostMultiplier() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) BailOutCostMultiplier() (*big.Int, error) {
	return _SlashIndicator.Contract.BailOutCostMultiplier(&_SlashIndicator.CallOpts)
}

// CurrentUnavailabilityIndicator is a free data retrieval call binding the contract method 0x082e7420.
//
// Solidity: function currentUnavailabilityIndicator(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) CurrentUnavailabilityIndicator(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "currentUnavailabilityIndicator", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentUnavailabilityIndicator is a free data retrieval call binding the contract method 0x082e7420.
//
// Solidity: function currentUnavailabilityIndicator(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) CurrentUnavailabilityIndicator(_validator common.Address) (*big.Int, error) {
	return _SlashIndicator.Contract.CurrentUnavailabilityIndicator(&_SlashIndicator.CallOpts, _validator)
}

// CurrentUnavailabilityIndicator is a free data retrieval call binding the contract method 0x082e7420.
//
// Solidity: function currentUnavailabilityIndicator(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) CurrentUnavailabilityIndicator(_validator common.Address) (*big.Int, error) {
	return _SlashIndicator.Contract.CurrentUnavailabilityIndicator(&_SlashIndicator.CallOpts, _validator)
}

// GainCreditScore is a free data retrieval call binding the contract method 0x3be2ed8a.
//
// Solidity: function gainCreditScore() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) GainCreditScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "gainCreditScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GainCreditScore is a free data retrieval call binding the contract method 0x3be2ed8a.
//
// Solidity: function gainCreditScore() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) GainCreditScore() (*big.Int, error) {
	return _SlashIndicator.Contract.GainCreditScore(&_SlashIndicator.CallOpts)
}

// GainCreditScore is a free data retrieval call binding the contract method 0x3be2ed8a.
//
// Solidity: function gainCreditScore() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GainCreditScore() (*big.Int, error) {
	return _SlashIndicator.Contract.GainCreditScore(&_SlashIndicator.CallOpts)
}

// GetBridgeOperatorSlashingConfigs is a free data retrieval call binding the contract method 0x1079402a.
//
// Solidity: function getBridgeOperatorSlashingConfigs() view returns(uint256, uint256, uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetBridgeOperatorSlashingConfigs(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getBridgeOperatorSlashingConfigs")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetBridgeOperatorSlashingConfigs is a free data retrieval call binding the contract method 0x1079402a.
//
// Solidity: function getBridgeOperatorSlashingConfigs() view returns(uint256, uint256, uint256)
func (_SlashIndicator *SlashIndicatorSession) GetBridgeOperatorSlashingConfigs() (*big.Int, *big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetBridgeOperatorSlashingConfigs(&_SlashIndicator.CallOpts)
}

// GetBridgeOperatorSlashingConfigs is a free data retrieval call binding the contract method 0x1079402a.
//
// Solidity: function getBridgeOperatorSlashingConfigs() view returns(uint256, uint256, uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetBridgeOperatorSlashingConfigs() (*big.Int, *big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetBridgeOperatorSlashingConfigs(&_SlashIndicator.CallOpts)
}

// GetBridgeVotingSlashingConfigs is a free data retrieval call binding the contract method 0xc2e524dc.
//
// Solidity: function getBridgeVotingSlashingConfigs() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetBridgeVotingSlashingConfigs(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getBridgeVotingSlashingConfigs")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBridgeVotingSlashingConfigs is a free data retrieval call binding the contract method 0xc2e524dc.
//
// Solidity: function getBridgeVotingSlashingConfigs() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorSession) GetBridgeVotingSlashingConfigs() (*big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetBridgeVotingSlashingConfigs(&_SlashIndicator.CallOpts)
}

// GetBridgeVotingSlashingConfigs is a free data retrieval call binding the contract method 0xc2e524dc.
//
// Solidity: function getBridgeVotingSlashingConfigs() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetBridgeVotingSlashingConfigs() (*big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetBridgeVotingSlashingConfigs(&_SlashIndicator.CallOpts)
}

// GetBulkCreditScore is a free data retrieval call binding the contract method 0x9c0b57d8.
//
// Solidity: function getBulkCreditScore(address[] _validators) view returns(uint256[] _resultList)
func (_SlashIndicator *SlashIndicatorCaller) GetBulkCreditScore(opts *bind.CallOpts, _validators []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getBulkCreditScore", _validators)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBulkCreditScore is a free data retrieval call binding the contract method 0x9c0b57d8.
//
// Solidity: function getBulkCreditScore(address[] _validators) view returns(uint256[] _resultList)
func (_SlashIndicator *SlashIndicatorSession) GetBulkCreditScore(_validators []common.Address) ([]*big.Int, error) {
	return _SlashIndicator.Contract.GetBulkCreditScore(&_SlashIndicator.CallOpts, _validators)
}

// GetBulkCreditScore is a free data retrieval call binding the contract method 0x9c0b57d8.
//
// Solidity: function getBulkCreditScore(address[] _validators) view returns(uint256[] _resultList)
func (_SlashIndicator *SlashIndicatorCallerSession) GetBulkCreditScore(_validators []common.Address) ([]*big.Int, error) {
	return _SlashIndicator.Contract.GetBulkCreditScore(&_SlashIndicator.CallOpts, _validators)
}

// GetCreditScore is a free data retrieval call binding the contract method 0xd3dd2bdf.
//
// Solidity: function getCreditScore(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetCreditScore(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getCreditScore", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreditScore is a free data retrieval call binding the contract method 0xd3dd2bdf.
//
// Solidity: function getCreditScore(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) GetCreditScore(_validator common.Address) (*big.Int, error) {
	return _SlashIndicator.Contract.GetCreditScore(&_SlashIndicator.CallOpts, _validator)
}

// GetCreditScore is a free data retrieval call binding the contract method 0xd3dd2bdf.
//
// Solidity: function getCreditScore(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetCreditScore(_validator common.Address) (*big.Int, error) {
	return _SlashIndicator.Contract.GetCreditScore(&_SlashIndicator.CallOpts, _validator)
}

// GetCreditScoreConfigs is a free data retrieval call binding the contract method 0xc6391fa2.
//
// Solidity: function getCreditScoreConfigs() view returns(uint256 _gainCreditScore, uint256 _maxCreditScore, uint256 _bailOutCostMultiplier)
func (_SlashIndicator *SlashIndicatorCaller) GetCreditScoreConfigs(opts *bind.CallOpts) (struct {
	GainCreditScore       *big.Int
	MaxCreditScore        *big.Int
	BailOutCostMultiplier *big.Int
}, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getCreditScoreConfigs")

	outstruct := new(struct {
		GainCreditScore       *big.Int
		MaxCreditScore        *big.Int
		BailOutCostMultiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GainCreditScore = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxCreditScore = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BailOutCostMultiplier = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCreditScoreConfigs is a free data retrieval call binding the contract method 0xc6391fa2.
//
// Solidity: function getCreditScoreConfigs() view returns(uint256 _gainCreditScore, uint256 _maxCreditScore, uint256 _bailOutCostMultiplier)
func (_SlashIndicator *SlashIndicatorSession) GetCreditScoreConfigs() (struct {
	GainCreditScore       *big.Int
	MaxCreditScore        *big.Int
	BailOutCostMultiplier *big.Int
}, error) {
	return _SlashIndicator.Contract.GetCreditScoreConfigs(&_SlashIndicator.CallOpts)
}

// GetCreditScoreConfigs is a free data retrieval call binding the contract method 0xc6391fa2.
//
// Solidity: function getCreditScoreConfigs() view returns(uint256 _gainCreditScore, uint256 _maxCreditScore, uint256 _bailOutCostMultiplier)
func (_SlashIndicator *SlashIndicatorCallerSession) GetCreditScoreConfigs() (struct {
	GainCreditScore       *big.Int
	MaxCreditScore        *big.Int
	BailOutCostMultiplier *big.Int
}, error) {
	return _SlashIndicator.Contract.GetCreditScoreConfigs(&_SlashIndicator.CallOpts)
}

// GetDoubleSignSlashingConfigs is a free data retrieval call binding the contract method 0xdf4b6ee0.
//
// Solidity: function getDoubleSignSlashingConfigs() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetDoubleSignSlashingConfigs(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getDoubleSignSlashingConfigs")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDoubleSignSlashingConfigs is a free data retrieval call binding the contract method 0xdf4b6ee0.
//
// Solidity: function getDoubleSignSlashingConfigs() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorSession) GetDoubleSignSlashingConfigs() (*big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetDoubleSignSlashingConfigs(&_SlashIndicator.CallOpts)
}

// GetDoubleSignSlashingConfigs is a free data retrieval call binding the contract method 0xdf4b6ee0.
//
// Solidity: function getDoubleSignSlashingConfigs() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetDoubleSignSlashingConfigs() (*big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetDoubleSignSlashingConfigs(&_SlashIndicator.CallOpts)
}

// GetUnavailabilityIndicator is a free data retrieval call binding the contract method 0x62ffe6cb.
//
// Solidity: function getUnavailabilityIndicator(address _validator, uint256 _period) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetUnavailabilityIndicator(opts *bind.CallOpts, _validator common.Address, _period *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getUnavailabilityIndicator", _validator, _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnavailabilityIndicator is a free data retrieval call binding the contract method 0x62ffe6cb.
//
// Solidity: function getUnavailabilityIndicator(address _validator, uint256 _period) view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) GetUnavailabilityIndicator(_validator common.Address, _period *big.Int) (*big.Int, error) {
	return _SlashIndicator.Contract.GetUnavailabilityIndicator(&_SlashIndicator.CallOpts, _validator, _period)
}

// GetUnavailabilityIndicator is a free data retrieval call binding the contract method 0x62ffe6cb.
//
// Solidity: function getUnavailabilityIndicator(address _validator, uint256 _period) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetUnavailabilityIndicator(_validator common.Address, _period *big.Int) (*big.Int, error) {
	return _SlashIndicator.Contract.GetUnavailabilityIndicator(&_SlashIndicator.CallOpts, _validator, _period)
}

// GetUnavailabilitySlashingConfigs is a free data retrieval call binding the contract method 0x3d48fd7d.
//
// Solidity: function getUnavailabilitySlashingConfigs() view returns(uint256, uint256, uint256, uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetUnavailabilitySlashingConfigs(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getUnavailabilitySlashingConfigs")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetUnavailabilitySlashingConfigs is a free data retrieval call binding the contract method 0x3d48fd7d.
//
// Solidity: function getUnavailabilitySlashingConfigs() view returns(uint256, uint256, uint256, uint256)
func (_SlashIndicator *SlashIndicatorSession) GetUnavailabilitySlashingConfigs() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetUnavailabilitySlashingConfigs(&_SlashIndicator.CallOpts)
}

// GetUnavailabilitySlashingConfigs is a free data retrieval call binding the contract method 0x3d48fd7d.
//
// Solidity: function getUnavailabilitySlashingConfigs() view returns(uint256, uint256, uint256, uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetUnavailabilitySlashingConfigs() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetUnavailabilitySlashingConfigs(&_SlashIndicator.CallOpts)
}

// LastUnavailabilitySlashedBlock is a free data retrieval call binding the contract method 0xf562b3c4.
//
// Solidity: function lastUnavailabilitySlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) LastUnavailabilitySlashedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "lastUnavailabilitySlashedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUnavailabilitySlashedBlock is a free data retrieval call binding the contract method 0xf562b3c4.
//
// Solidity: function lastUnavailabilitySlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) LastUnavailabilitySlashedBlock() (*big.Int, error) {
	return _SlashIndicator.Contract.LastUnavailabilitySlashedBlock(&_SlashIndicator.CallOpts)
}

// LastUnavailabilitySlashedBlock is a free data retrieval call binding the contract method 0xf562b3c4.
//
// Solidity: function lastUnavailabilitySlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) LastUnavailabilitySlashedBlock() (*big.Int, error) {
	return _SlashIndicator.Contract.LastUnavailabilitySlashedBlock(&_SlashIndicator.CallOpts)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCaller) MaintenanceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "maintenanceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_SlashIndicator *SlashIndicatorSession) MaintenanceContract() (common.Address, error) {
	return _SlashIndicator.Contract.MaintenanceContract(&_SlashIndicator.CallOpts)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCallerSession) MaintenanceContract() (common.Address, error) {
	return _SlashIndicator.Contract.MaintenanceContract(&_SlashIndicator.CallOpts)
}

// MaxCreditScore is a free data retrieval call binding the contract method 0x4ee38230.
//
// Solidity: function maxCreditScore() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) MaxCreditScore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "maxCreditScore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCreditScore is a free data retrieval call binding the contract method 0x4ee38230.
//
// Solidity: function maxCreditScore() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) MaxCreditScore() (*big.Int, error) {
	return _SlashIndicator.Contract.MaxCreditScore(&_SlashIndicator.CallOpts)
}

// MaxCreditScore is a free data retrieval call binding the contract method 0x4ee38230.
//
// Solidity: function maxCreditScore() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) MaxCreditScore() (*big.Int, error) {
	return _SlashIndicator.Contract.MaxCreditScore(&_SlashIndicator.CallOpts)
}

// PrecompileValidateDoubleSignAddress is a free data retrieval call binding the contract method 0x7c2b55a0.
//
// Solidity: function precompileValidateDoubleSignAddress() view returns(address)
func (_SlashIndicator *SlashIndicatorCaller) PrecompileValidateDoubleSignAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "precompileValidateDoubleSignAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompileValidateDoubleSignAddress is a free data retrieval call binding the contract method 0x7c2b55a0.
//
// Solidity: function precompileValidateDoubleSignAddress() view returns(address)
func (_SlashIndicator *SlashIndicatorSession) PrecompileValidateDoubleSignAddress() (common.Address, error) {
	return _SlashIndicator.Contract.PrecompileValidateDoubleSignAddress(&_SlashIndicator.CallOpts)
}

// PrecompileValidateDoubleSignAddress is a free data retrieval call binding the contract method 0x7c2b55a0.
//
// Solidity: function precompileValidateDoubleSignAddress() view returns(address)
func (_SlashIndicator *SlashIndicatorCallerSession) PrecompileValidateDoubleSignAddress() (common.Address, error) {
	return _SlashIndicator.Contract.PrecompileValidateDoubleSignAddress(&_SlashIndicator.CallOpts)
}

// RoninGovernanceAdminContract is a free data retrieval call binding the contract method 0x23368e47.
//
// Solidity: function roninGovernanceAdminContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCaller) RoninGovernanceAdminContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "roninGovernanceAdminContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoninGovernanceAdminContract is a free data retrieval call binding the contract method 0x23368e47.
//
// Solidity: function roninGovernanceAdminContract() view returns(address)
func (_SlashIndicator *SlashIndicatorSession) RoninGovernanceAdminContract() (common.Address, error) {
	return _SlashIndicator.Contract.RoninGovernanceAdminContract(&_SlashIndicator.CallOpts)
}

// RoninGovernanceAdminContract is a free data retrieval call binding the contract method 0x23368e47.
//
// Solidity: function roninGovernanceAdminContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCallerSession) RoninGovernanceAdminContract() (common.Address, error) {
	return _SlashIndicator.Contract.RoninGovernanceAdminContract(&_SlashIndicator.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCaller) RoninTrustedOrganizationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "roninTrustedOrganizationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_SlashIndicator *SlashIndicatorSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _SlashIndicator.Contract.RoninTrustedOrganizationContract(&_SlashIndicator.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCallerSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _SlashIndicator.Contract.RoninTrustedOrganizationContract(&_SlashIndicator.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_SlashIndicator *SlashIndicatorSession) ValidatorContract() (common.Address, error) {
	return _SlashIndicator.Contract.ValidatorContract(&_SlashIndicator.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCallerSession) ValidatorContract() (common.Address, error) {
	return _SlashIndicator.Contract.ValidatorContract(&_SlashIndicator.CallOpts)
}

// BailOut is a paid mutator transaction binding the contract method 0xd1f992f7.
//
// Solidity: function bailOut(address _consensusAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) BailOut(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "bailOut", _consensusAddr)
}

// BailOut is a paid mutator transaction binding the contract method 0xd1f992f7.
//
// Solidity: function bailOut(address _consensusAddr) returns()
func (_SlashIndicator *SlashIndicatorSession) BailOut(_consensusAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.BailOut(&_SlashIndicator.TransactOpts, _consensusAddr)
}

// BailOut is a paid mutator transaction binding the contract method 0xd1f992f7.
//
// Solidity: function bailOut(address _consensusAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) BailOut(_consensusAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.BailOut(&_SlashIndicator.TransactOpts, _consensusAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x118c01a2.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __roninGovernanceAdminContract, uint256[3] _bridgeOperatorSlashingConfigs, uint256[2] _bridgeVotingSlashingConfigs, uint256[2] _doubleSignSlashingConfigs, uint256[4] _unavailabilitySlashingConfigs, uint256[3] _creditScoreConfigs) returns()
func (_SlashIndicator *SlashIndicatorTransactor) Initialize(opts *bind.TransactOpts, __validatorContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __roninGovernanceAdminContract common.Address, _bridgeOperatorSlashingConfigs [3]*big.Int, _bridgeVotingSlashingConfigs [2]*big.Int, _doubleSignSlashingConfigs [2]*big.Int, _unavailabilitySlashingConfigs [4]*big.Int, _creditScoreConfigs [3]*big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "initialize", __validatorContract, __maintenanceContract, __roninTrustedOrganizationContract, __roninGovernanceAdminContract, _bridgeOperatorSlashingConfigs, _bridgeVotingSlashingConfigs, _doubleSignSlashingConfigs, _unavailabilitySlashingConfigs, _creditScoreConfigs)
}

// Initialize is a paid mutator transaction binding the contract method 0x118c01a2.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __roninGovernanceAdminContract, uint256[3] _bridgeOperatorSlashingConfigs, uint256[2] _bridgeVotingSlashingConfigs, uint256[2] _doubleSignSlashingConfigs, uint256[4] _unavailabilitySlashingConfigs, uint256[3] _creditScoreConfigs) returns()
func (_SlashIndicator *SlashIndicatorSession) Initialize(__validatorContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __roninGovernanceAdminContract common.Address, _bridgeOperatorSlashingConfigs [3]*big.Int, _bridgeVotingSlashingConfigs [2]*big.Int, _doubleSignSlashingConfigs [2]*big.Int, _unavailabilitySlashingConfigs [4]*big.Int, _creditScoreConfigs [3]*big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Initialize(&_SlashIndicator.TransactOpts, __validatorContract, __maintenanceContract, __roninTrustedOrganizationContract, __roninGovernanceAdminContract, _bridgeOperatorSlashingConfigs, _bridgeVotingSlashingConfigs, _doubleSignSlashingConfigs, _unavailabilitySlashingConfigs, _creditScoreConfigs)
}

// Initialize is a paid mutator transaction binding the contract method 0x118c01a2.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __roninGovernanceAdminContract, uint256[3] _bridgeOperatorSlashingConfigs, uint256[2] _bridgeVotingSlashingConfigs, uint256[2] _doubleSignSlashingConfigs, uint256[4] _unavailabilitySlashingConfigs, uint256[3] _creditScoreConfigs) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) Initialize(__validatorContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __roninGovernanceAdminContract common.Address, _bridgeOperatorSlashingConfigs [3]*big.Int, _bridgeVotingSlashingConfigs [2]*big.Int, _doubleSignSlashingConfigs [2]*big.Int, _unavailabilitySlashingConfigs [4]*big.Int, _creditScoreConfigs [3]*big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Initialize(&_SlashIndicator.TransactOpts, __validatorContract, __maintenanceContract, __roninTrustedOrganizationContract, __roninGovernanceAdminContract, _bridgeOperatorSlashingConfigs, _bridgeVotingSlashingConfigs, _doubleSignSlashingConfigs, _unavailabilitySlashingConfigs, _creditScoreConfigs)
}

// SetBridgeOperatorSlashingConfigs is a paid mutator transaction binding the contract method 0xbe2be2b1.
//
// Solidity: function setBridgeOperatorSlashingConfigs(uint256 _ratioTier1, uint256 _ratioTier2, uint256 _jailDurationTier2) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetBridgeOperatorSlashingConfigs(opts *bind.TransactOpts, _ratioTier1 *big.Int, _ratioTier2 *big.Int, _jailDurationTier2 *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setBridgeOperatorSlashingConfigs", _ratioTier1, _ratioTier2, _jailDurationTier2)
}

// SetBridgeOperatorSlashingConfigs is a paid mutator transaction binding the contract method 0xbe2be2b1.
//
// Solidity: function setBridgeOperatorSlashingConfigs(uint256 _ratioTier1, uint256 _ratioTier2, uint256 _jailDurationTier2) returns()
func (_SlashIndicator *SlashIndicatorSession) SetBridgeOperatorSlashingConfigs(_ratioTier1 *big.Int, _ratioTier2 *big.Int, _jailDurationTier2 *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetBridgeOperatorSlashingConfigs(&_SlashIndicator.TransactOpts, _ratioTier1, _ratioTier2, _jailDurationTier2)
}

// SetBridgeOperatorSlashingConfigs is a paid mutator transaction binding the contract method 0xbe2be2b1.
//
// Solidity: function setBridgeOperatorSlashingConfigs(uint256 _ratioTier1, uint256 _ratioTier2, uint256 _jailDurationTier2) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetBridgeOperatorSlashingConfigs(_ratioTier1 *big.Int, _ratioTier2 *big.Int, _jailDurationTier2 *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetBridgeOperatorSlashingConfigs(&_SlashIndicator.TransactOpts, _ratioTier1, _ratioTier2, _jailDurationTier2)
}

// SetBridgeVotingSlashingConfigs is a paid mutator transaction binding the contract method 0x853af1b7.
//
// Solidity: function setBridgeVotingSlashingConfigs(uint256 _threshold, uint256 _slashAmount) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetBridgeVotingSlashingConfigs(opts *bind.TransactOpts, _threshold *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setBridgeVotingSlashingConfigs", _threshold, _slashAmount)
}

// SetBridgeVotingSlashingConfigs is a paid mutator transaction binding the contract method 0x853af1b7.
//
// Solidity: function setBridgeVotingSlashingConfigs(uint256 _threshold, uint256 _slashAmount) returns()
func (_SlashIndicator *SlashIndicatorSession) SetBridgeVotingSlashingConfigs(_threshold *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetBridgeVotingSlashingConfigs(&_SlashIndicator.TransactOpts, _threshold, _slashAmount)
}

// SetBridgeVotingSlashingConfigs is a paid mutator transaction binding the contract method 0x853af1b7.
//
// Solidity: function setBridgeVotingSlashingConfigs(uint256 _threshold, uint256 _slashAmount) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetBridgeVotingSlashingConfigs(_threshold *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetBridgeVotingSlashingConfigs(&_SlashIndicator.TransactOpts, _threshold, _slashAmount)
}

// SetCreditScoreConfigs is a paid mutator transaction binding the contract method 0x639ce29e.
//
// Solidity: function setCreditScoreConfigs(uint256 _gainCreditScore, uint256 _maxCreditScore, uint256 _bailOutCostMultiplier) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetCreditScoreConfigs(opts *bind.TransactOpts, _gainCreditScore *big.Int, _maxCreditScore *big.Int, _bailOutCostMultiplier *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setCreditScoreConfigs", _gainCreditScore, _maxCreditScore, _bailOutCostMultiplier)
}

// SetCreditScoreConfigs is a paid mutator transaction binding the contract method 0x639ce29e.
//
// Solidity: function setCreditScoreConfigs(uint256 _gainCreditScore, uint256 _maxCreditScore, uint256 _bailOutCostMultiplier) returns()
func (_SlashIndicator *SlashIndicatorSession) SetCreditScoreConfigs(_gainCreditScore *big.Int, _maxCreditScore *big.Int, _bailOutCostMultiplier *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetCreditScoreConfigs(&_SlashIndicator.TransactOpts, _gainCreditScore, _maxCreditScore, _bailOutCostMultiplier)
}

// SetCreditScoreConfigs is a paid mutator transaction binding the contract method 0x639ce29e.
//
// Solidity: function setCreditScoreConfigs(uint256 _gainCreditScore, uint256 _maxCreditScore, uint256 _bailOutCostMultiplier) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetCreditScoreConfigs(_gainCreditScore *big.Int, _maxCreditScore *big.Int, _bailOutCostMultiplier *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetCreditScoreConfigs(&_SlashIndicator.TransactOpts, _gainCreditScore, _maxCreditScore, _bailOutCostMultiplier)
}

// SetDoubleSignSlashingConfigs is a paid mutator transaction binding the contract method 0x61d3b60b.
//
// Solidity: function setDoubleSignSlashingConfigs(uint256 _slashAmount, uint256 _jailUntilBlock) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetDoubleSignSlashingConfigs(opts *bind.TransactOpts, _slashAmount *big.Int, _jailUntilBlock *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setDoubleSignSlashingConfigs", _slashAmount, _jailUntilBlock)
}

// SetDoubleSignSlashingConfigs is a paid mutator transaction binding the contract method 0x61d3b60b.
//
// Solidity: function setDoubleSignSlashingConfigs(uint256 _slashAmount, uint256 _jailUntilBlock) returns()
func (_SlashIndicator *SlashIndicatorSession) SetDoubleSignSlashingConfigs(_slashAmount *big.Int, _jailUntilBlock *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetDoubleSignSlashingConfigs(&_SlashIndicator.TransactOpts, _slashAmount, _jailUntilBlock)
}

// SetDoubleSignSlashingConfigs is a paid mutator transaction binding the contract method 0x61d3b60b.
//
// Solidity: function setDoubleSignSlashingConfigs(uint256 _slashAmount, uint256 _jailUntilBlock) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetDoubleSignSlashingConfigs(_slashAmount *big.Int, _jailUntilBlock *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetDoubleSignSlashingConfigs(&_SlashIndicator.TransactOpts, _slashAmount, _jailUntilBlock)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetMaintenanceContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setMaintenanceContract", _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetMaintenanceContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetMaintenanceContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetRoninGovernanceAdminContract is a paid mutator transaction binding the contract method 0xd73e81b8.
//
// Solidity: function setRoninGovernanceAdminContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetRoninGovernanceAdminContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setRoninGovernanceAdminContract", _addr)
}

// SetRoninGovernanceAdminContract is a paid mutator transaction binding the contract method 0xd73e81b8.
//
// Solidity: function setRoninGovernanceAdminContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorSession) SetRoninGovernanceAdminContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetRoninGovernanceAdminContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetRoninGovernanceAdminContract is a paid mutator transaction binding the contract method 0xd73e81b8.
//
// Solidity: function setRoninGovernanceAdminContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetRoninGovernanceAdminContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetRoninGovernanceAdminContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetRoninTrustedOrganizationContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setRoninTrustedOrganizationContract", _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetRoninTrustedOrganizationContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetRoninTrustedOrganizationContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetUnavailabilitySlashingConfigs is a paid mutator transaction binding the contract method 0xd1737e27.
//
// Solidity: function setUnavailabilitySlashingConfigs(uint256 _tier1Threshold, uint256 _tier2Threshold, uint256 _slashAmountForTier2Threshold, uint256 _jailDurationForTier2Threshold) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetUnavailabilitySlashingConfigs(opts *bind.TransactOpts, _tier1Threshold *big.Int, _tier2Threshold *big.Int, _slashAmountForTier2Threshold *big.Int, _jailDurationForTier2Threshold *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setUnavailabilitySlashingConfigs", _tier1Threshold, _tier2Threshold, _slashAmountForTier2Threshold, _jailDurationForTier2Threshold)
}

// SetUnavailabilitySlashingConfigs is a paid mutator transaction binding the contract method 0xd1737e27.
//
// Solidity: function setUnavailabilitySlashingConfigs(uint256 _tier1Threshold, uint256 _tier2Threshold, uint256 _slashAmountForTier2Threshold, uint256 _jailDurationForTier2Threshold) returns()
func (_SlashIndicator *SlashIndicatorSession) SetUnavailabilitySlashingConfigs(_tier1Threshold *big.Int, _tier2Threshold *big.Int, _slashAmountForTier2Threshold *big.Int, _jailDurationForTier2Threshold *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetUnavailabilitySlashingConfigs(&_SlashIndicator.TransactOpts, _tier1Threshold, _tier2Threshold, _slashAmountForTier2Threshold, _jailDurationForTier2Threshold)
}

// SetUnavailabilitySlashingConfigs is a paid mutator transaction binding the contract method 0xd1737e27.
//
// Solidity: function setUnavailabilitySlashingConfigs(uint256 _tier1Threshold, uint256 _tier2Threshold, uint256 _slashAmountForTier2Threshold, uint256 _jailDurationForTier2Threshold) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetUnavailabilitySlashingConfigs(_tier1Threshold *big.Int, _tier2Threshold *big.Int, _slashAmountForTier2Threshold *big.Int, _jailDurationForTier2Threshold *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetUnavailabilitySlashingConfigs(&_SlashIndicator.TransactOpts, _tier1Threshold, _tier2Threshold, _slashAmountForTier2Threshold, _jailDurationForTier2Threshold)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetValidatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setValidatorContract", _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetValidatorContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetValidatorContract(&_SlashIndicator.TransactOpts, _addr)
}

// SlashBridgeVoting is a paid mutator transaction binding the contract method 0x1a697341.
//
// Solidity: function slashBridgeVoting(address _consensusAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SlashBridgeVoting(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "slashBridgeVoting", _consensusAddr)
}

// SlashBridgeVoting is a paid mutator transaction binding the contract method 0x1a697341.
//
// Solidity: function slashBridgeVoting(address _consensusAddr) returns()
func (_SlashIndicator *SlashIndicatorSession) SlashBridgeVoting(_consensusAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashBridgeVoting(&_SlashIndicator.TransactOpts, _consensusAddr)
}

// SlashBridgeVoting is a paid mutator transaction binding the contract method 0x1a697341.
//
// Solidity: function slashBridgeVoting(address _consensusAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SlashBridgeVoting(_consensusAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashBridgeVoting(&_SlashIndicator.TransactOpts, _consensusAddr)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x1e90b2a0.
//
// Solidity: function slashDoubleSign(address _consensuAddr, bytes _header1, bytes _header2) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SlashDoubleSign(opts *bind.TransactOpts, _consensuAddr common.Address, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "slashDoubleSign", _consensuAddr, _header1, _header2)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x1e90b2a0.
//
// Solidity: function slashDoubleSign(address _consensuAddr, bytes _header1, bytes _header2) returns()
func (_SlashIndicator *SlashIndicatorSession) SlashDoubleSign(_consensuAddr common.Address, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashDoubleSign(&_SlashIndicator.TransactOpts, _consensuAddr, _header1, _header2)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x1e90b2a0.
//
// Solidity: function slashDoubleSign(address _consensuAddr, bytes _header1, bytes _header2) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SlashDoubleSign(_consensuAddr common.Address, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashDoubleSign(&_SlashIndicator.TransactOpts, _consensuAddr, _header1, _header2)
}

// SlashUnavailability is a paid mutator transaction binding the contract method 0xfd422cd0.
//
// Solidity: function slashUnavailability(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SlashUnavailability(opts *bind.TransactOpts, _validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "slashUnavailability", _validatorAddr)
}

// SlashUnavailability is a paid mutator transaction binding the contract method 0xfd422cd0.
//
// Solidity: function slashUnavailability(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorSession) SlashUnavailability(_validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashUnavailability(&_SlashIndicator.TransactOpts, _validatorAddr)
}

// SlashUnavailability is a paid mutator transaction binding the contract method 0xfd422cd0.
//
// Solidity: function slashUnavailability(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SlashUnavailability(_validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashUnavailability(&_SlashIndicator.TransactOpts, _validatorAddr)
}

// UpdateCreditScore is a paid mutator transaction binding the contract method 0x129fccc1.
//
// Solidity: function updateCreditScore(address[] _validators, uint256 _period) returns()
func (_SlashIndicator *SlashIndicatorTransactor) UpdateCreditScore(opts *bind.TransactOpts, _validators []common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "updateCreditScore", _validators, _period)
}

// UpdateCreditScore is a paid mutator transaction binding the contract method 0x129fccc1.
//
// Solidity: function updateCreditScore(address[] _validators, uint256 _period) returns()
func (_SlashIndicator *SlashIndicatorSession) UpdateCreditScore(_validators []common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.UpdateCreditScore(&_SlashIndicator.TransactOpts, _validators, _period)
}

// UpdateCreditScore is a paid mutator transaction binding the contract method 0x129fccc1.
//
// Solidity: function updateCreditScore(address[] _validators, uint256 _period) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) UpdateCreditScore(_validators []common.Address, _period *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.UpdateCreditScore(&_SlashIndicator.TransactOpts, _validators, _period)
}

// SlashIndicatorBailedOutIterator is returned from FilterBailedOut and is used to iterate over the raw logs and unpacked data for BailedOut events raised by the SlashIndicator contract.
type SlashIndicatorBailedOutIterator struct {
	Event *SlashIndicatorBailedOut // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorBailedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorBailedOut)
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
		it.Event = new(SlashIndicatorBailedOut)
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
func (it *SlashIndicatorBailedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorBailedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorBailedOut represents a BailedOut event raised by the SlashIndicator contract.
type SlashIndicatorBailedOut struct {
	Validator common.Address
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBailedOut is a free log retrieval operation binding the contract event 0xf05ce5353bbe65ac592f5b335f37b5efe46dfa2035bb2e1e19795e73eb535627.
//
// Solidity: event BailedOut(address indexed validator, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) FilterBailedOut(opts *bind.FilterOpts, validator []common.Address) (*SlashIndicatorBailedOutIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "BailedOut", validatorRule)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorBailedOutIterator{contract: _SlashIndicator.contract, event: "BailedOut", logs: logs, sub: sub}, nil
}

// WatchBailedOut is a free log subscription operation binding the contract event 0xf05ce5353bbe65ac592f5b335f37b5efe46dfa2035bb2e1e19795e73eb535627.
//
// Solidity: event BailedOut(address indexed validator, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) WatchBailedOut(opts *bind.WatchOpts, sink chan<- *SlashIndicatorBailedOut, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "BailedOut", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorBailedOut)
				if err := _SlashIndicator.contract.UnpackLog(event, "BailedOut", log); err != nil {
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

// ParseBailedOut is a log parse operation binding the contract event 0xf05ce5353bbe65ac592f5b335f37b5efe46dfa2035bb2e1e19795e73eb535627.
//
// Solidity: event BailedOut(address indexed validator, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) ParseBailedOut(log types.Log) (*SlashIndicatorBailedOut, error) {
	event := new(SlashIndicatorBailedOut)
	if err := _SlashIndicator.contract.UnpackLog(event, "BailedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorBridgeOperatorSlashingConfigsUpdatedIterator is returned from FilterBridgeOperatorSlashingConfigsUpdated and is used to iterate over the raw logs and unpacked data for BridgeOperatorSlashingConfigsUpdated events raised by the SlashIndicator contract.
type SlashIndicatorBridgeOperatorSlashingConfigsUpdatedIterator struct {
	Event *SlashIndicatorBridgeOperatorSlashingConfigsUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorBridgeOperatorSlashingConfigsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorBridgeOperatorSlashingConfigsUpdated)
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
		it.Event = new(SlashIndicatorBridgeOperatorSlashingConfigsUpdated)
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
func (it *SlashIndicatorBridgeOperatorSlashingConfigsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorBridgeOperatorSlashingConfigsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorBridgeOperatorSlashingConfigsUpdated represents a BridgeOperatorSlashingConfigsUpdated event raised by the SlashIndicator contract.
type SlashIndicatorBridgeOperatorSlashingConfigsUpdated struct {
	MissingVotesRatioTier1                *big.Int
	MissingVotesRatioTier2                *big.Int
	JailDurationForMissingVotesRatioTier2 *big.Int
	Raw                                   types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorSlashingConfigsUpdated is a free log retrieval operation binding the contract event 0x48b79bd792893e8ce8de399a7c03796ca95a43d5782307ff7c775f492cdf7c82.
//
// Solidity: event BridgeOperatorSlashingConfigsUpdated(uint256 missingVotesRatioTier1, uint256 missingVotesRatioTier2, uint256 jailDurationForMissingVotesRatioTier2)
func (_SlashIndicator *SlashIndicatorFilterer) FilterBridgeOperatorSlashingConfigsUpdated(opts *bind.FilterOpts) (*SlashIndicatorBridgeOperatorSlashingConfigsUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "BridgeOperatorSlashingConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorBridgeOperatorSlashingConfigsUpdatedIterator{contract: _SlashIndicator.contract, event: "BridgeOperatorSlashingConfigsUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorSlashingConfigsUpdated is a free log subscription operation binding the contract event 0x48b79bd792893e8ce8de399a7c03796ca95a43d5782307ff7c775f492cdf7c82.
//
// Solidity: event BridgeOperatorSlashingConfigsUpdated(uint256 missingVotesRatioTier1, uint256 missingVotesRatioTier2, uint256 jailDurationForMissingVotesRatioTier2)
func (_SlashIndicator *SlashIndicatorFilterer) WatchBridgeOperatorSlashingConfigsUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorBridgeOperatorSlashingConfigsUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "BridgeOperatorSlashingConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorBridgeOperatorSlashingConfigsUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "BridgeOperatorSlashingConfigsUpdated", log); err != nil {
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

// ParseBridgeOperatorSlashingConfigsUpdated is a log parse operation binding the contract event 0x48b79bd792893e8ce8de399a7c03796ca95a43d5782307ff7c775f492cdf7c82.
//
// Solidity: event BridgeOperatorSlashingConfigsUpdated(uint256 missingVotesRatioTier1, uint256 missingVotesRatioTier2, uint256 jailDurationForMissingVotesRatioTier2)
func (_SlashIndicator *SlashIndicatorFilterer) ParseBridgeOperatorSlashingConfigsUpdated(log types.Log) (*SlashIndicatorBridgeOperatorSlashingConfigsUpdated, error) {
	event := new(SlashIndicatorBridgeOperatorSlashingConfigsUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "BridgeOperatorSlashingConfigsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorBridgeVotingSlashingConfigsUpdatedIterator is returned from FilterBridgeVotingSlashingConfigsUpdated and is used to iterate over the raw logs and unpacked data for BridgeVotingSlashingConfigsUpdated events raised by the SlashIndicator contract.
type SlashIndicatorBridgeVotingSlashingConfigsUpdatedIterator struct {
	Event *SlashIndicatorBridgeVotingSlashingConfigsUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorBridgeVotingSlashingConfigsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorBridgeVotingSlashingConfigsUpdated)
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
		it.Event = new(SlashIndicatorBridgeVotingSlashingConfigsUpdated)
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
func (it *SlashIndicatorBridgeVotingSlashingConfigsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorBridgeVotingSlashingConfigsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorBridgeVotingSlashingConfigsUpdated represents a BridgeVotingSlashingConfigsUpdated event raised by the SlashIndicator contract.
type SlashIndicatorBridgeVotingSlashingConfigsUpdated struct {
	BridgeVotingThreshold   *big.Int
	BridgeVotingSlashAmount *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBridgeVotingSlashingConfigsUpdated is a free log retrieval operation binding the contract event 0xbda9ec2980d7468ba6a9f363696315affca9f9770016396bdea2ac39c3e5d61a.
//
// Solidity: event BridgeVotingSlashingConfigsUpdated(uint256 bridgeVotingThreshold, uint256 bridgeVotingSlashAmount)
func (_SlashIndicator *SlashIndicatorFilterer) FilterBridgeVotingSlashingConfigsUpdated(opts *bind.FilterOpts) (*SlashIndicatorBridgeVotingSlashingConfigsUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "BridgeVotingSlashingConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorBridgeVotingSlashingConfigsUpdatedIterator{contract: _SlashIndicator.contract, event: "BridgeVotingSlashingConfigsUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeVotingSlashingConfigsUpdated is a free log subscription operation binding the contract event 0xbda9ec2980d7468ba6a9f363696315affca9f9770016396bdea2ac39c3e5d61a.
//
// Solidity: event BridgeVotingSlashingConfigsUpdated(uint256 bridgeVotingThreshold, uint256 bridgeVotingSlashAmount)
func (_SlashIndicator *SlashIndicatorFilterer) WatchBridgeVotingSlashingConfigsUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorBridgeVotingSlashingConfigsUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "BridgeVotingSlashingConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorBridgeVotingSlashingConfigsUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "BridgeVotingSlashingConfigsUpdated", log); err != nil {
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

// ParseBridgeVotingSlashingConfigsUpdated is a log parse operation binding the contract event 0xbda9ec2980d7468ba6a9f363696315affca9f9770016396bdea2ac39c3e5d61a.
//
// Solidity: event BridgeVotingSlashingConfigsUpdated(uint256 bridgeVotingThreshold, uint256 bridgeVotingSlashAmount)
func (_SlashIndicator *SlashIndicatorFilterer) ParseBridgeVotingSlashingConfigsUpdated(log types.Log) (*SlashIndicatorBridgeVotingSlashingConfigsUpdated, error) {
	event := new(SlashIndicatorBridgeVotingSlashingConfigsUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "BridgeVotingSlashingConfigsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorCreditScoreConfigsUpdatedIterator is returned from FilterCreditScoreConfigsUpdated and is used to iterate over the raw logs and unpacked data for CreditScoreConfigsUpdated events raised by the SlashIndicator contract.
type SlashIndicatorCreditScoreConfigsUpdatedIterator struct {
	Event *SlashIndicatorCreditScoreConfigsUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorCreditScoreConfigsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorCreditScoreConfigsUpdated)
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
		it.Event = new(SlashIndicatorCreditScoreConfigsUpdated)
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
func (it *SlashIndicatorCreditScoreConfigsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorCreditScoreConfigsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorCreditScoreConfigsUpdated represents a CreditScoreConfigsUpdated event raised by the SlashIndicator contract.
type SlashIndicatorCreditScoreConfigsUpdated struct {
	GainCreditScore       *big.Int
	MaxCreditScore        *big.Int
	BailOutCostMultiplier *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreditScoreConfigsUpdated is a free log retrieval operation binding the contract event 0xdfff174e3d0e322e127529251e09478f81f83b832d96f4d2b10c8416cc9e8525.
//
// Solidity: event CreditScoreConfigsUpdated(uint256 gainCreditScore, uint256 maxCreditScore, uint256 bailOutCostMultiplier)
func (_SlashIndicator *SlashIndicatorFilterer) FilterCreditScoreConfigsUpdated(opts *bind.FilterOpts) (*SlashIndicatorCreditScoreConfigsUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "CreditScoreConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorCreditScoreConfigsUpdatedIterator{contract: _SlashIndicator.contract, event: "CreditScoreConfigsUpdated", logs: logs, sub: sub}, nil
}

// WatchCreditScoreConfigsUpdated is a free log subscription operation binding the contract event 0xdfff174e3d0e322e127529251e09478f81f83b832d96f4d2b10c8416cc9e8525.
//
// Solidity: event CreditScoreConfigsUpdated(uint256 gainCreditScore, uint256 maxCreditScore, uint256 bailOutCostMultiplier)
func (_SlashIndicator *SlashIndicatorFilterer) WatchCreditScoreConfigsUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorCreditScoreConfigsUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "CreditScoreConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorCreditScoreConfigsUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "CreditScoreConfigsUpdated", log); err != nil {
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

// ParseCreditScoreConfigsUpdated is a log parse operation binding the contract event 0xdfff174e3d0e322e127529251e09478f81f83b832d96f4d2b10c8416cc9e8525.
//
// Solidity: event CreditScoreConfigsUpdated(uint256 gainCreditScore, uint256 maxCreditScore, uint256 bailOutCostMultiplier)
func (_SlashIndicator *SlashIndicatorFilterer) ParseCreditScoreConfigsUpdated(log types.Log) (*SlashIndicatorCreditScoreConfigsUpdated, error) {
	event := new(SlashIndicatorCreditScoreConfigsUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "CreditScoreConfigsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorCreditScoresUpdatedIterator is returned from FilterCreditScoresUpdated and is used to iterate over the raw logs and unpacked data for CreditScoresUpdated events raised by the SlashIndicator contract.
type SlashIndicatorCreditScoresUpdatedIterator struct {
	Event *SlashIndicatorCreditScoresUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorCreditScoresUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorCreditScoresUpdated)
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
		it.Event = new(SlashIndicatorCreditScoresUpdated)
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
func (it *SlashIndicatorCreditScoresUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorCreditScoresUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorCreditScoresUpdated represents a CreditScoresUpdated event raised by the SlashIndicator contract.
type SlashIndicatorCreditScoresUpdated struct {
	Validators   []common.Address
	CreditScores []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreditScoresUpdated is a free log retrieval operation binding the contract event 0x8c02b2ccee964dc50649a48bd0a0446f0f9e88ecdb72d099f4ace07c39c23480.
//
// Solidity: event CreditScoresUpdated(address[] validators, uint256[] creditScores)
func (_SlashIndicator *SlashIndicatorFilterer) FilterCreditScoresUpdated(opts *bind.FilterOpts) (*SlashIndicatorCreditScoresUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "CreditScoresUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorCreditScoresUpdatedIterator{contract: _SlashIndicator.contract, event: "CreditScoresUpdated", logs: logs, sub: sub}, nil
}

// WatchCreditScoresUpdated is a free log subscription operation binding the contract event 0x8c02b2ccee964dc50649a48bd0a0446f0f9e88ecdb72d099f4ace07c39c23480.
//
// Solidity: event CreditScoresUpdated(address[] validators, uint256[] creditScores)
func (_SlashIndicator *SlashIndicatorFilterer) WatchCreditScoresUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorCreditScoresUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "CreditScoresUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorCreditScoresUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "CreditScoresUpdated", log); err != nil {
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

// ParseCreditScoresUpdated is a log parse operation binding the contract event 0x8c02b2ccee964dc50649a48bd0a0446f0f9e88ecdb72d099f4ace07c39c23480.
//
// Solidity: event CreditScoresUpdated(address[] validators, uint256[] creditScores)
func (_SlashIndicator *SlashIndicatorFilterer) ParseCreditScoresUpdated(log types.Log) (*SlashIndicatorCreditScoresUpdated, error) {
	event := new(SlashIndicatorCreditScoresUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "CreditScoresUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorDoubleSignSlashingConfigsUpdatedIterator is returned from FilterDoubleSignSlashingConfigsUpdated and is used to iterate over the raw logs and unpacked data for DoubleSignSlashingConfigsUpdated events raised by the SlashIndicator contract.
type SlashIndicatorDoubleSignSlashingConfigsUpdatedIterator struct {
	Event *SlashIndicatorDoubleSignSlashingConfigsUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorDoubleSignSlashingConfigsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorDoubleSignSlashingConfigsUpdated)
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
		it.Event = new(SlashIndicatorDoubleSignSlashingConfigsUpdated)
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
func (it *SlashIndicatorDoubleSignSlashingConfigsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorDoubleSignSlashingConfigsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorDoubleSignSlashingConfigsUpdated represents a DoubleSignSlashingConfigsUpdated event raised by the SlashIndicator contract.
type SlashIndicatorDoubleSignSlashingConfigsUpdated struct {
	SlashDoubleSignAmount       *big.Int
	DoubleSigningJailUntilBlock *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterDoubleSignSlashingConfigsUpdated is a free log retrieval operation binding the contract event 0x2f551c9d5c16e8a5444109ee232c78ed055e4e5cefe25e162b3bae190af0dedc.
//
// Solidity: event DoubleSignSlashingConfigsUpdated(uint256 slashDoubleSignAmount, uint256 doubleSigningJailUntilBlock)
func (_SlashIndicator *SlashIndicatorFilterer) FilterDoubleSignSlashingConfigsUpdated(opts *bind.FilterOpts) (*SlashIndicatorDoubleSignSlashingConfigsUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "DoubleSignSlashingConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorDoubleSignSlashingConfigsUpdatedIterator{contract: _SlashIndicator.contract, event: "DoubleSignSlashingConfigsUpdated", logs: logs, sub: sub}, nil
}

// WatchDoubleSignSlashingConfigsUpdated is a free log subscription operation binding the contract event 0x2f551c9d5c16e8a5444109ee232c78ed055e4e5cefe25e162b3bae190af0dedc.
//
// Solidity: event DoubleSignSlashingConfigsUpdated(uint256 slashDoubleSignAmount, uint256 doubleSigningJailUntilBlock)
func (_SlashIndicator *SlashIndicatorFilterer) WatchDoubleSignSlashingConfigsUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorDoubleSignSlashingConfigsUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "DoubleSignSlashingConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorDoubleSignSlashingConfigsUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "DoubleSignSlashingConfigsUpdated", log); err != nil {
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

// ParseDoubleSignSlashingConfigsUpdated is a log parse operation binding the contract event 0x2f551c9d5c16e8a5444109ee232c78ed055e4e5cefe25e162b3bae190af0dedc.
//
// Solidity: event DoubleSignSlashingConfigsUpdated(uint256 slashDoubleSignAmount, uint256 doubleSigningJailUntilBlock)
func (_SlashIndicator *SlashIndicatorFilterer) ParseDoubleSignSlashingConfigsUpdated(log types.Log) (*SlashIndicatorDoubleSignSlashingConfigsUpdated, error) {
	event := new(SlashIndicatorDoubleSignSlashingConfigsUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "DoubleSignSlashingConfigsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SlashIndicator contract.
type SlashIndicatorInitializedIterator struct {
	Event *SlashIndicatorInitialized // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorInitialized)
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
		it.Event = new(SlashIndicatorInitialized)
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
func (it *SlashIndicatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorInitialized represents a Initialized event raised by the SlashIndicator contract.
type SlashIndicatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashIndicator *SlashIndicatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*SlashIndicatorInitializedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorInitializedIterator{contract: _SlashIndicator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashIndicator *SlashIndicatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SlashIndicatorInitialized) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorInitialized)
				if err := _SlashIndicator.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashIndicator *SlashIndicatorFilterer) ParseInitialized(log types.Log) (*SlashIndicatorInitialized, error) {
	event := new(SlashIndicatorInitialized)
	if err := _SlashIndicator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorMaintenanceContractUpdatedIterator is returned from FilterMaintenanceContractUpdated and is used to iterate over the raw logs and unpacked data for MaintenanceContractUpdated events raised by the SlashIndicator contract.
type SlashIndicatorMaintenanceContractUpdatedIterator struct {
	Event *SlashIndicatorMaintenanceContractUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorMaintenanceContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorMaintenanceContractUpdated)
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
		it.Event = new(SlashIndicatorMaintenanceContractUpdated)
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
func (it *SlashIndicatorMaintenanceContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorMaintenanceContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorMaintenanceContractUpdated represents a MaintenanceContractUpdated event raised by the SlashIndicator contract.
type SlashIndicatorMaintenanceContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceContractUpdated is a free log retrieval operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) FilterMaintenanceContractUpdated(opts *bind.FilterOpts) (*SlashIndicatorMaintenanceContractUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorMaintenanceContractUpdatedIterator{contract: _SlashIndicator.contract, event: "MaintenanceContractUpdated", logs: logs, sub: sub}, nil
}

// WatchMaintenanceContractUpdated is a free log subscription operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) WatchMaintenanceContractUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorMaintenanceContractUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorMaintenanceContractUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
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

// ParseMaintenanceContractUpdated is a log parse operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) ParseMaintenanceContractUpdated(log types.Log) (*SlashIndicatorMaintenanceContractUpdated, error) {
	event := new(SlashIndicatorMaintenanceContractUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorRoninGovernanceAdminContractUpdatedIterator is returned from FilterRoninGovernanceAdminContractUpdated and is used to iterate over the raw logs and unpacked data for RoninGovernanceAdminContractUpdated events raised by the SlashIndicator contract.
type SlashIndicatorRoninGovernanceAdminContractUpdatedIterator struct {
	Event *SlashIndicatorRoninGovernanceAdminContractUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorRoninGovernanceAdminContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorRoninGovernanceAdminContractUpdated)
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
		it.Event = new(SlashIndicatorRoninGovernanceAdminContractUpdated)
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
func (it *SlashIndicatorRoninGovernanceAdminContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorRoninGovernanceAdminContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorRoninGovernanceAdminContractUpdated represents a RoninGovernanceAdminContractUpdated event raised by the SlashIndicator contract.
type SlashIndicatorRoninGovernanceAdminContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRoninGovernanceAdminContractUpdated is a free log retrieval operation binding the contract event 0x9125df97e014f5cc4f107fd784acd35e8e2188ca7c2a0f7caa478365747c1c83.
//
// Solidity: event RoninGovernanceAdminContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) FilterRoninGovernanceAdminContractUpdated(opts *bind.FilterOpts) (*SlashIndicatorRoninGovernanceAdminContractUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "RoninGovernanceAdminContractUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorRoninGovernanceAdminContractUpdatedIterator{contract: _SlashIndicator.contract, event: "RoninGovernanceAdminContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRoninGovernanceAdminContractUpdated is a free log subscription operation binding the contract event 0x9125df97e014f5cc4f107fd784acd35e8e2188ca7c2a0f7caa478365747c1c83.
//
// Solidity: event RoninGovernanceAdminContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) WatchRoninGovernanceAdminContractUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorRoninGovernanceAdminContractUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "RoninGovernanceAdminContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorRoninGovernanceAdminContractUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "RoninGovernanceAdminContractUpdated", log); err != nil {
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

// ParseRoninGovernanceAdminContractUpdated is a log parse operation binding the contract event 0x9125df97e014f5cc4f107fd784acd35e8e2188ca7c2a0f7caa478365747c1c83.
//
// Solidity: event RoninGovernanceAdminContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) ParseRoninGovernanceAdminContractUpdated(log types.Log) (*SlashIndicatorRoninGovernanceAdminContractUpdated, error) {
	event := new(SlashIndicatorRoninGovernanceAdminContractUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "RoninGovernanceAdminContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorRoninTrustedOrganizationContractUpdatedIterator is returned from FilterRoninTrustedOrganizationContractUpdated and is used to iterate over the raw logs and unpacked data for RoninTrustedOrganizationContractUpdated events raised by the SlashIndicator contract.
type SlashIndicatorRoninTrustedOrganizationContractUpdatedIterator struct {
	Event *SlashIndicatorRoninTrustedOrganizationContractUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorRoninTrustedOrganizationContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorRoninTrustedOrganizationContractUpdated)
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
		it.Event = new(SlashIndicatorRoninTrustedOrganizationContractUpdated)
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
func (it *SlashIndicatorRoninTrustedOrganizationContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorRoninTrustedOrganizationContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorRoninTrustedOrganizationContractUpdated represents a RoninTrustedOrganizationContractUpdated event raised by the SlashIndicator contract.
type SlashIndicatorRoninTrustedOrganizationContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRoninTrustedOrganizationContractUpdated is a free log retrieval operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) FilterRoninTrustedOrganizationContractUpdated(opts *bind.FilterOpts) (*SlashIndicatorRoninTrustedOrganizationContractUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorRoninTrustedOrganizationContractUpdatedIterator{contract: _SlashIndicator.contract, event: "RoninTrustedOrganizationContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRoninTrustedOrganizationContractUpdated is a free log subscription operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) WatchRoninTrustedOrganizationContractUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorRoninTrustedOrganizationContractUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorRoninTrustedOrganizationContractUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
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

// ParseRoninTrustedOrganizationContractUpdated is a log parse operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) ParseRoninTrustedOrganizationContractUpdated(log types.Log) (*SlashIndicatorRoninTrustedOrganizationContractUpdated, error) {
	event := new(SlashIndicatorRoninTrustedOrganizationContractUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the SlashIndicator contract.
type SlashIndicatorSlashedIterator struct {
	Event *SlashIndicatorSlashed // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorSlashed)
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
		it.Event = new(SlashIndicatorSlashed)
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
func (it *SlashIndicatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorSlashed represents a Slashed event raised by the SlashIndicator contract.
type SlashIndicatorSlashed struct {
	Validator common.Address
	SlashType uint8
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x607adba66cff84b627e3537d1c17d088a98556bccd0536a2f3590c56329023d9.
//
// Solidity: event Slashed(address indexed validator, uint8 slashType, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) FilterSlashed(opts *bind.FilterOpts, validator []common.Address) (*SlashIndicatorSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "Slashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorSlashedIterator{contract: _SlashIndicator.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x607adba66cff84b627e3537d1c17d088a98556bccd0536a2f3590c56329023d9.
//
// Solidity: event Slashed(address indexed validator, uint8 slashType, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *SlashIndicatorSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "Slashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorSlashed)
				if err := _SlashIndicator.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x607adba66cff84b627e3537d1c17d088a98556bccd0536a2f3590c56329023d9.
//
// Solidity: event Slashed(address indexed validator, uint8 slashType, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) ParseSlashed(log types.Log) (*SlashIndicatorSlashed, error) {
	event := new(SlashIndicatorSlashed)
	if err := _SlashIndicator.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorUnavailabilitySlashingConfigsUpdatedIterator is returned from FilterUnavailabilitySlashingConfigsUpdated and is used to iterate over the raw logs and unpacked data for UnavailabilitySlashingConfigsUpdated events raised by the SlashIndicator contract.
type SlashIndicatorUnavailabilitySlashingConfigsUpdatedIterator struct {
	Event *SlashIndicatorUnavailabilitySlashingConfigsUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorUnavailabilitySlashingConfigsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorUnavailabilitySlashingConfigsUpdated)
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
		it.Event = new(SlashIndicatorUnavailabilitySlashingConfigsUpdated)
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
func (it *SlashIndicatorUnavailabilitySlashingConfigsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorUnavailabilitySlashingConfigsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorUnavailabilitySlashingConfigsUpdated represents a UnavailabilitySlashingConfigsUpdated event raised by the SlashIndicator contract.
type SlashIndicatorUnavailabilitySlashingConfigsUpdated struct {
	UnavailabilityTier1Threshold                *big.Int
	UnavailabilityTier2Threshold                *big.Int
	SlashAmountForUnavailabilityTier2Threshold  *big.Int
	JailDurationForUnavailabilityTier2Threshold *big.Int
	Raw                                         types.Log // Blockchain specific contextual infos
}

// FilterUnavailabilitySlashingConfigsUpdated is a free log retrieval operation binding the contract event 0x442862e6143ad95854e7c13ff4947ec6e43bc87160e3b193e7c1abaf6e3aaa98.
//
// Solidity: event UnavailabilitySlashingConfigsUpdated(uint256 unavailabilityTier1Threshold, uint256 unavailabilityTier2Threshold, uint256 slashAmountForUnavailabilityTier2Threshold, uint256 jailDurationForUnavailabilityTier2Threshold)
func (_SlashIndicator *SlashIndicatorFilterer) FilterUnavailabilitySlashingConfigsUpdated(opts *bind.FilterOpts) (*SlashIndicatorUnavailabilitySlashingConfigsUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "UnavailabilitySlashingConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorUnavailabilitySlashingConfigsUpdatedIterator{contract: _SlashIndicator.contract, event: "UnavailabilitySlashingConfigsUpdated", logs: logs, sub: sub}, nil
}

// WatchUnavailabilitySlashingConfigsUpdated is a free log subscription operation binding the contract event 0x442862e6143ad95854e7c13ff4947ec6e43bc87160e3b193e7c1abaf6e3aaa98.
//
// Solidity: event UnavailabilitySlashingConfigsUpdated(uint256 unavailabilityTier1Threshold, uint256 unavailabilityTier2Threshold, uint256 slashAmountForUnavailabilityTier2Threshold, uint256 jailDurationForUnavailabilityTier2Threshold)
func (_SlashIndicator *SlashIndicatorFilterer) WatchUnavailabilitySlashingConfigsUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorUnavailabilitySlashingConfigsUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "UnavailabilitySlashingConfigsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorUnavailabilitySlashingConfigsUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "UnavailabilitySlashingConfigsUpdated", log); err != nil {
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

// ParseUnavailabilitySlashingConfigsUpdated is a log parse operation binding the contract event 0x442862e6143ad95854e7c13ff4947ec6e43bc87160e3b193e7c1abaf6e3aaa98.
//
// Solidity: event UnavailabilitySlashingConfigsUpdated(uint256 unavailabilityTier1Threshold, uint256 unavailabilityTier2Threshold, uint256 slashAmountForUnavailabilityTier2Threshold, uint256 jailDurationForUnavailabilityTier2Threshold)
func (_SlashIndicator *SlashIndicatorFilterer) ParseUnavailabilitySlashingConfigsUpdated(log types.Log) (*SlashIndicatorUnavailabilitySlashingConfigsUpdated, error) {
	event := new(SlashIndicatorUnavailabilitySlashingConfigsUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "UnavailabilitySlashingConfigsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the SlashIndicator contract.
type SlashIndicatorValidatorContractUpdatedIterator struct {
	Event *SlashIndicatorValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorValidatorContractUpdated)
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
		it.Event = new(SlashIndicatorValidatorContractUpdated)
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
func (it *SlashIndicatorValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorValidatorContractUpdated represents a ValidatorContractUpdated event raised by the SlashIndicator contract.
type SlashIndicatorValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*SlashIndicatorValidatorContractUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorValidatorContractUpdatedIterator{contract: _SlashIndicator.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorValidatorContractUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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

// ParseValidatorContractUpdated is a log parse operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) ParseValidatorContractUpdated(log types.Log) (*SlashIndicatorValidatorContractUpdated, error) {
	event := new(SlashIndicatorValidatorContractUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
