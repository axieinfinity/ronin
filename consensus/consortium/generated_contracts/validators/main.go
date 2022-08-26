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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentValidatorSetMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"}],\"name\":\"depositReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526102586000806101000a81548161ffff021916908361ffff1602179055503480156200002f57600080fd5b5060405162000c3638038062000c3683398181016040528101906200005591906200019d565b80600190805190602001906200006d92919062000075565b505062000281565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620000ad5760008555620000f9565b82601f10620000c857805160ff1916838001178555620000f9565b82800160010185558215620000f9579182015b82811115620000f8578251825591602001919060010190620000db565b5b5090506200010891906200010c565b5090565b5b80821115620001275760008160009055506001016200010d565b5090565b6000620001426200013c8462000216565b620001e2565b9050828152602081018484840111156200015b57600080fd5b6200016884828562000249565b509392505050565b600082601f8301126200018257600080fd5b8151620001948482602086016200012b565b91505092915050565b600060208284031215620001b057600080fd5b600082015167ffffffffffffffff811115620001cb57600080fd5b620001d98482850162000170565b91505092915050565b6000604051905081810181811067ffffffffffffffff821117156200020c576200020b6200027f565b5b8060405250919050565b600067ffffffffffffffff8211156200023457620002336200027f565b5b601f19601f8301169050602081019050919050565b60005b83811015620002695780820151818401526020810190506200024c565b8381111562000279576000848401525b50505050565bfe5b6109a580620002916000396000f3fe6080604052600436106100705760003560e01c806378121dd41161004e57806378121dd4146100e9578063ad3c9da614610114578063b7ab4db514610151578063db8ec38c1461017c57610070565b80632d497ba2146100755780636969a25c1461008c5780636ffa4dc1146100cd575b600080fd5b34801561008157600080fd5b5061008a6101a5565b005b34801561009857600080fd5b506100b360048036038101906100ae91906106c7565b610260565b6040516100c49594939291906107e2565b60405180910390f35b6100e760048036038101906100e29190610662565b6102e6565b005b3480156100f557600080fd5b506100fe610438565b60405161010b9190610877565b60405180910390f35b34801561012057600080fd5b5061013b60048036038101906101369190610662565b610442565b6040516101489190610877565b60405180910390f35b34801561015d57600080fd5b5061016661045a565b6040516101739190610835565b60405180910390f35b34801561018857600080fd5b506101a3600480360381019061019e919061068b565b610574565b005b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610213576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020a90610857565b60405180910390fd5b6000439050600160008054906101000a900461ffff160361ffff1660008054906101000a900461ffff1661ffff16828161024957fe5b0614610255575061025e565b80600481905550505b565b6002818154811061027057600080fd5b90600052602060002090600502016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040154905085565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610354576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034b90610857565b60405180910390fd5b60005b6002805490508167ffffffffffffffff161015610433578173ffffffffffffffffffffffffffffffffffffffff1660028267ffffffffffffffff168154811061039c57fe5b906000526020600020906005020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610426573460028267ffffffffffffffff168154811061040257fe5b90600052602060002090600502016004016000828254019250508190555050610435565b8080600101915050610357565b505b50565b6000600454905090565b60036020528060005260406000206000915090505481565b6060600060028054905067ffffffffffffffff8111801561047a57600080fd5b506040519080825280602002602001820160405280156104a95781602001602082028036833780820191505090505b50905060005b6002805490508167ffffffffffffffff16101561056c5760028167ffffffffffffffff16815481106104dd57fe5b906000526020600020906005020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828267ffffffffffffffff168151811061052557fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806001019150506104af565b508091505090565b600060026001816001815401808255809150500390600052602060002090600502019050828160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550818160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b6000813590506106328161092a565b92915050565b60008135905061064781610941565b92915050565b60008135905061065c81610958565b92915050565b60006020828403121561067457600080fd5b600061068284828501610623565b91505092915050565b6000806040838503121561069e57600080fd5b60006106ac85828601610623565b92505060206106bd85828601610638565b9150509250929050565b6000602082840312156106d957600080fd5b60006106e78482850161064d565b91505092915050565b60006106fc8383610717565b60208301905092915050565b610711816108ee565b82525050565b610720816108dc565b82525050565b61072f816108dc565b82525050565b6000610740826108a2565b61074a81856108ba565b935061075583610892565b8060005b8381101561078657815161076d88826106f0565b9750610778836108ad565b925050600181019050610759565b5085935050505092915050565b60006107a06016836108cb565b91507f73656e646572206973206e6f7420636f696e62617365000000000000000000006000830152602082019050919050565b6107dc81610920565b82525050565b600060a0820190506107f76000830188610726565b6108046020830187610708565b61081160408301866107d3565b61081e60608301856107d3565b61082b60808301846107d3565b9695505050505050565b6000602082019050818103600083015261084f8184610735565b905092915050565b6000602082019050818103600083015261087081610793565b9050919050565b600060208201905061088c60008301846107d3565b92915050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006108e782610900565b9050919050565b60006108f982610900565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b610933816108dc565b811461093e57600080fd5b50565b61094a816108ee565b811461095557600080fd5b50565b61096181610920565b811461096c57600080fd5b5056fea264697066735822122015480192b8d148e024696a592a3fd5361e8c0abc4dc8ca3aebbc99c652f20abe64736f6c63430007060033",
}

// ValidatorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorsMetaData.ABI instead.
var ValidatorsABI = ValidatorsMetaData.ABI

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorsMetaData.Bin instead.
var ValidatorsBin = ValidatorsMetaData.Bin

// DeployValidators deploys a new Ethereum contract, binding an instance of Validators to it.
func DeployValidators(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Validators, error) {
	parsed, err := ValidatorsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorsBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

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
