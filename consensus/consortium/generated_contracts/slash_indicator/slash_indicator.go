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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRoninValidatorSet\",\"name\":\"_validatorSetContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"UnavailabilityIndicatorsReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumISlashIndicator.SlashType\",\"name\":\"slashType\",\"type\":\"uint8\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"felonyThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getSlashIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlashThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSlashedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"misdemeanorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validatorAddrs\",\"type\":\"address[]\"}],\"name\":\"resetCounters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_evidence\",\"type\":\"bytes\"}],\"name\":\"slashDoubleSign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"contractIRoninValidatorSet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200104d3803806200104d83398181016040528101906200003791906200010d565b6032600281905550609660038190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200013f565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000c18262000094565b9050919050565b6000620000d582620000b4565b9050919050565b620000e781620000c8565b8114620000f357600080fd5b50565b6000815190506200010781620000dc565b92915050565b6000602082840312156200012657620001256200008f565b5b60006200013684828501620000f6565b91505092915050565b610efe806200014f6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638256ace6116100665780638256ace614610120578063994390891461013f578063aef250be1461015d578063c96be4cb1461017b578063ec35fe721461019757610093565b806337c8dab914610098578063389f4f71146100c8578063518e463a146100e6578063567a372d14610102575b600080fd5b6100b260048036038101906100ad91906107e0565b6101b3565b6040516100bf9190610826565b60405180910390f35b6100d06101fb565b6040516100dd9190610826565b60405180910390f35b61010060048036038101906100fb91906108a6565b610201565b005b61010a6102aa565b6040516101179190610826565b60405180910390f35b6101286102b0565b604051610136929190610906565b60405180910390f35b6101476102c1565b604051610154919061098e565b60405180910390f35b6101656102e7565b6040516101729190610826565b60405180910390f35b610195600480360381019061019091906107e0565b6102ed565b005b6101b160048036038101906101ac91906109ff565b610605565b005b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60035481565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461026f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026690610acf565b60405180910390fd5b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a190610b3b565b60405180910390fd5b60025481565b600080600254600354915091509091565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461035b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035290610acf565b60405180910390fd5b600154431161039f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039690610bcd565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603156105fb5760008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815461042090610c1c565b9190508190559050600354810361051257600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663722d298f836040518263ffffffff1660e01b815260040161048c9190610c73565b600060405180830381600087803b1580156104a657600080fd5b505af11580156104ba573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff167f72b4cbc6db714fd31d4f3a8686c34df7b11819178ea38a792798f38a99d15e4a60026040516105059190610d05565b60405180910390a26105f9565b60025481036105f857600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633dc56e09836040518263ffffffff1660e01b81526004016105769190610c73565b600060405180830381600087803b15801561059057600080fd5b505af11580156105a4573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff167f72b4cbc6db714fd31d4f3a8686c34df7b11819178ea38a792798f38a99d15e4a60016040516105ef9190610d05565b60405180910390a25b5b505b4360018190555050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610695576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068c90610d92565b60405180910390fd5b61069f82826106a3565b5050565b60008282905003156107745760005b82829050811015610739576000808484848181106106d3576106d2610db2565b5b90506020020160208101906106e891906107e0565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055808061073190610c1c565b9150506106b2565b507f9ae71e4f83d6019cd988ca4c38a696d015d5f6e41f3fd708fb746d2b672b4996828260405161076b929190610ea4565b60405180910390a15b5050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107ad82610782565b9050919050565b6107bd816107a2565b81146107c857600080fd5b50565b6000813590506107da816107b4565b92915050565b6000602082840312156107f6576107f5610778565b5b6000610804848285016107cb565b91505092915050565b6000819050919050565b6108208161080d565b82525050565b600060208201905061083b6000830184610817565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261086657610865610841565b5b8235905067ffffffffffffffff81111561088357610882610846565b5b60208301915083600182028301111561089f5761089e61084b565b5b9250929050565b6000806000604084860312156108bf576108be610778565b5b60006108cd868287016107cb565b935050602084013567ffffffffffffffff8111156108ee576108ed61077d565b5b6108fa86828701610850565b92509250509250925092565b600060408201905061091b6000830185610817565b6109286020830184610817565b9392505050565b6000819050919050565b600061095461094f61094a84610782565b61092f565b610782565b9050919050565b600061096682610939565b9050919050565b60006109788261095b565b9050919050565b6109888161096d565b82525050565b60006020820190506109a3600083018461097f565b92915050565b60008083601f8401126109bf576109be610841565b5b8235905067ffffffffffffffff8111156109dc576109db610846565b5b6020830191508360208202830111156109f8576109f761084b565b5b9250929050565b60008060208385031215610a1657610a15610778565b5b600083013567ffffffffffffffff811115610a3457610a3361077d565b5b610a40858286016109a9565b92509250509250929050565b600082825260208201905092915050565b7f536c617368496e64696361746f723a206d6574686f642063616c6c657220697360008201527f206e6f742074686520636f696e62617365000000000000000000000000000000602082015250565b6000610ab9603183610a4c565b9150610ac482610a5d565b604082019050919050565b60006020820190508181036000830152610ae881610aac565b9050919050565b7f4e6f7420696d706c656d656e7465640000000000000000000000000000000000600082015250565b6000610b25600f83610a4c565b9150610b3082610aef565b602082019050919050565b60006020820190508181036000830152610b5481610b18565b9050919050565b7f536c617368496e64696361746f723a2063616e6e6f7420736c6173682074776960008201527f636520696e206f6e6520626c6f636b0000000000000000000000000000000000602082015250565b6000610bb7602f83610a4c565b9150610bc282610b5b565b604082019050919050565b60006020820190508181036000830152610be681610baa565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c278261080d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c5957610c58610bed565b5b600182019050919050565b610c6d816107a2565b82525050565b6000602082019050610c886000830184610c64565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110610cce57610ccd610c8e565b5b50565b6000819050610cdf82610cbd565b919050565b6000610cef82610cd1565b9050919050565b610cff81610ce4565b82525050565b6000602082019050610d1a6000830184610cf6565b92915050565b7f536c617368496e64696361746f723a206d6574686f642063616c6c657220697360008201527f206e6f74207468652076616c696461746f7220636f6e74726163740000000000602082015250565b6000610d7c603b83610a4c565b9150610d8782610d20565b604082019050919050565b60006020820190508181036000830152610dab81610d6f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b6000819050919050565b610e05816107a2565b82525050565b6000610e178383610dfc565b60208301905092915050565b6000610e3260208401846107cb565b905092915050565b6000602082019050919050565b6000610e538385610de1565b9350610e5e82610df2565b8060005b85811015610e9757610e748284610e23565b610e7e8882610e0b565b9750610e8983610e3a565b925050600181019050610e62565b5085925050509392505050565b60006020820190508181036000830152610ebf818486610e47565b9050939250505056fea264697066735822122091242cd808a850e5ec6d87e9be207e3ce6e1d7f067c64bbea66e01a4beca6a7b64736f6c63430008100033",
}

// SlashIndicatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashIndicatorMetaData.ABI instead.
var SlashIndicatorABI = SlashIndicatorMetaData.ABI

// SlashIndicatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashIndicatorMetaData.Bin instead.
var SlashIndicatorBin = SlashIndicatorMetaData.Bin

// DeploySlashIndicator deploys a new Ethereum contract, binding an instance of SlashIndicator to it.
func DeploySlashIndicator(auth *bind.TransactOpts, backend bind.ContractBackend, _validatorSetContract common.Address) (common.Address, *types.Transaction, *SlashIndicator, error) {
	parsed, err := SlashIndicatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashIndicatorBin), backend, _validatorSetContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashIndicator{SlashIndicatorCaller: SlashIndicatorCaller{contract: contract}, SlashIndicatorTransactor: SlashIndicatorTransactor{contract: contract}, SlashIndicatorFilterer: SlashIndicatorFilterer{contract: contract}}, nil
}

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

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) FelonyThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "felonyThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) FelonyThreshold() (*big.Int, error) {
	return _SlashIndicator.Contract.FelonyThreshold(&_SlashIndicator.CallOpts)
}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) FelonyThreshold() (*big.Int, error) {
	return _SlashIndicator.Contract.FelonyThreshold(&_SlashIndicator.CallOpts)
}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetSlashIndicator(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getSlashIndicator", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) GetSlashIndicator(validator common.Address) (*big.Int, error) {
	return _SlashIndicator.Contract.GetSlashIndicator(&_SlashIndicator.CallOpts, validator)
}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetSlashIndicator(validator common.Address) (*big.Int, error) {
	return _SlashIndicator.Contract.GetSlashIndicator(&_SlashIndicator.CallOpts, validator)
}

// GetSlashThresholds is a free data retrieval call binding the contract method 0x8256ace6.
//
// Solidity: function getSlashThresholds() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetSlashThresholds(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getSlashThresholds")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSlashThresholds is a free data retrieval call binding the contract method 0x8256ace6.
//
// Solidity: function getSlashThresholds() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorSession) GetSlashThresholds() (*big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetSlashThresholds(&_SlashIndicator.CallOpts)
}

// GetSlashThresholds is a free data retrieval call binding the contract method 0x8256ace6.
//
// Solidity: function getSlashThresholds() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetSlashThresholds() (*big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetSlashThresholds(&_SlashIndicator.CallOpts)
}

// LastSlashedBlock is a free data retrieval call binding the contract method 0xaef250be.
//
// Solidity: function lastSlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) LastSlashedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "lastSlashedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSlashedBlock is a free data retrieval call binding the contract method 0xaef250be.
//
// Solidity: function lastSlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) LastSlashedBlock() (*big.Int, error) {
	return _SlashIndicator.Contract.LastSlashedBlock(&_SlashIndicator.CallOpts)
}

// LastSlashedBlock is a free data retrieval call binding the contract method 0xaef250be.
//
// Solidity: function lastSlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) LastSlashedBlock() (*big.Int, error) {
	return _SlashIndicator.Contract.LastSlashedBlock(&_SlashIndicator.CallOpts)
}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) MisdemeanorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "misdemeanorThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) MisdemeanorThreshold() (*big.Int, error) {
	return _SlashIndicator.Contract.MisdemeanorThreshold(&_SlashIndicator.CallOpts)
}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) MisdemeanorThreshold() (*big.Int, error) {
	return _SlashIndicator.Contract.MisdemeanorThreshold(&_SlashIndicator.CallOpts)
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

// ResetCounters is a paid mutator transaction binding the contract method 0xec35fe72.
//
// Solidity: function resetCounters(address[] _validatorAddrs) returns()
func (_SlashIndicator *SlashIndicatorTransactor) ResetCounters(opts *bind.TransactOpts, _validatorAddrs []common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "resetCounters", _validatorAddrs)
}

// ResetCounters is a paid mutator transaction binding the contract method 0xec35fe72.
//
// Solidity: function resetCounters(address[] _validatorAddrs) returns()
func (_SlashIndicator *SlashIndicatorSession) ResetCounters(_validatorAddrs []common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.ResetCounters(&_SlashIndicator.TransactOpts, _validatorAddrs)
}

// ResetCounters is a paid mutator transaction binding the contract method 0xec35fe72.
//
// Solidity: function resetCounters(address[] _validatorAddrs) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) ResetCounters(_validatorAddrs []common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.ResetCounters(&_SlashIndicator.TransactOpts, _validatorAddrs)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) Slash(opts *bind.TransactOpts, _validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "slash", _validatorAddr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorSession) Slash(_validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Slash(&_SlashIndicator.TransactOpts, _validatorAddr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) Slash(_validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Slash(&_SlashIndicator.TransactOpts, _validatorAddr)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x518e463a.
//
// Solidity: function slashDoubleSign(address _valAddr, bytes _evidence) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SlashDoubleSign(opts *bind.TransactOpts, _valAddr common.Address, _evidence []byte) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "slashDoubleSign", _valAddr, _evidence)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x518e463a.
//
// Solidity: function slashDoubleSign(address _valAddr, bytes _evidence) returns()
func (_SlashIndicator *SlashIndicatorSession) SlashDoubleSign(_valAddr common.Address, _evidence []byte) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashDoubleSign(&_SlashIndicator.TransactOpts, _valAddr, _evidence)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x518e463a.
//
// Solidity: function slashDoubleSign(address _valAddr, bytes _evidence) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SlashDoubleSign(_valAddr common.Address, _evidence []byte) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashDoubleSign(&_SlashIndicator.TransactOpts, _valAddr, _evidence)
}

// SlashIndicatorUnavailabilityIndicatorsResetIterator is returned from FilterUnavailabilityIndicatorsReset and is used to iterate over the raw logs and unpacked data for UnavailabilityIndicatorsReset events raised by the SlashIndicator contract.
type SlashIndicatorUnavailabilityIndicatorsResetIterator struct {
	Event *SlashIndicatorUnavailabilityIndicatorsReset // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorUnavailabilityIndicatorsResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorUnavailabilityIndicatorsReset)
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
		it.Event = new(SlashIndicatorUnavailabilityIndicatorsReset)
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
func (it *SlashIndicatorUnavailabilityIndicatorsResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorUnavailabilityIndicatorsResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorUnavailabilityIndicatorsReset represents a UnavailabilityIndicatorsReset event raised by the SlashIndicator contract.
type SlashIndicatorUnavailabilityIndicatorsReset struct {
	Validators []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnavailabilityIndicatorsReset is a free log retrieval operation binding the contract event 0x9ae71e4f83d6019cd988ca4c38a696d015d5f6e41f3fd708fb746d2b672b4996.
//
// Solidity: event UnavailabilityIndicatorsReset(address[] validators)
func (_SlashIndicator *SlashIndicatorFilterer) FilterUnavailabilityIndicatorsReset(opts *bind.FilterOpts) (*SlashIndicatorUnavailabilityIndicatorsResetIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "UnavailabilityIndicatorsReset")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorUnavailabilityIndicatorsResetIterator{contract: _SlashIndicator.contract, event: "UnavailabilityIndicatorsReset", logs: logs, sub: sub}, nil
}

// WatchUnavailabilityIndicatorsReset is a free log subscription operation binding the contract event 0x9ae71e4f83d6019cd988ca4c38a696d015d5f6e41f3fd708fb746d2b672b4996.
//
// Solidity: event UnavailabilityIndicatorsReset(address[] validators)
func (_SlashIndicator *SlashIndicatorFilterer) WatchUnavailabilityIndicatorsReset(opts *bind.WatchOpts, sink chan<- *SlashIndicatorUnavailabilityIndicatorsReset) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "UnavailabilityIndicatorsReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorUnavailabilityIndicatorsReset)
				if err := _SlashIndicator.contract.UnpackLog(event, "UnavailabilityIndicatorsReset", log); err != nil {
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

// ParseUnavailabilityIndicatorsReset is a log parse operation binding the contract event 0x9ae71e4f83d6019cd988ca4c38a696d015d5f6e41f3fd708fb746d2b672b4996.
//
// Solidity: event UnavailabilityIndicatorsReset(address[] validators)
func (_SlashIndicator *SlashIndicatorFilterer) ParseUnavailabilityIndicatorsReset(log types.Log) (*SlashIndicatorUnavailabilityIndicatorsReset, error) {
	event := new(SlashIndicatorUnavailabilityIndicatorsReset)
	if err := _SlashIndicator.contract.UnpackLog(event, "UnavailabilityIndicatorsReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the SlashIndicator contract.
type SlashIndicatorValidatorSlashedIterator struct {
	Event *SlashIndicatorValidatorSlashed // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorValidatorSlashed)
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
		it.Event = new(SlashIndicatorValidatorSlashed)
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
func (it *SlashIndicatorValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorValidatorSlashed represents a ValidatorSlashed event raised by the SlashIndicator contract.
type SlashIndicatorValidatorSlashed struct {
	Validator common.Address
	SlashType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0x72b4cbc6db714fd31d4f3a8686c34df7b11819178ea38a792798f38a99d15e4a.
//
// Solidity: event ValidatorSlashed(address indexed validator, uint8 slashType)
func (_SlashIndicator *SlashIndicatorFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, validator []common.Address) (*SlashIndicatorValidatorSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "ValidatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorValidatorSlashedIterator{contract: _SlashIndicator.contract, event: "ValidatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0x72b4cbc6db714fd31d4f3a8686c34df7b11819178ea38a792798f38a99d15e4a.
//
// Solidity: event ValidatorSlashed(address indexed validator, uint8 slashType)
func (_SlashIndicator *SlashIndicatorFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *SlashIndicatorValidatorSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "ValidatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorValidatorSlashed)
				if err := _SlashIndicator.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
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

// ParseValidatorSlashed is a log parse operation binding the contract event 0x72b4cbc6db714fd31d4f3a8686c34df7b11819178ea38a792798f38a99d15e4a.
//
// Solidity: event ValidatorSlashed(address indexed validator, uint8 slashType)
func (_SlashIndicator *SlashIndicatorFilterer) ParseValidatorSlashed(log types.Log) (*SlashIndicatorValidatorSlashed, error) {
	event := new(SlashIndicatorValidatorSlashed)
	if err := _SlashIndicator.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
