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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WhitelistDeployerABI is the input ABI used to generate the binding from.
const WhitelistDeployerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"whitelistAllAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"whitelistAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"whitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"AddressWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"WhitelistAllChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"}]"

// WhitelistDeployer is an auto generated Go binding around an Ethereum contract.
type WhitelistDeployer struct {
	WhitelistDeployerCaller     // Read-only binding to the contract
	WhitelistDeployerTransactor // Write-only binding to the contract
	WhitelistDeployerFilterer   // Log filterer for contract events
}

// WhitelistDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistDeployerSession struct {
	Contract     *WhitelistDeployer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WhitelistDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistDeployerCallerSession struct {
	Contract *WhitelistDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// WhitelistDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistDeployerTransactorSession struct {
	Contract     *WhitelistDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// WhitelistDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistDeployerRaw struct {
	Contract *WhitelistDeployer // Generic contract binding to access the raw methods on
}

// WhitelistDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistDeployerCallerRaw struct {
	Contract *WhitelistDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistDeployerTransactorRaw struct {
	Contract *WhitelistDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistDeployer creates a new instance of WhitelistDeployer, bound to a specific deployed contract.
func NewWhitelistDeployer(address common.Address, backend bind.ContractBackend) (*WhitelistDeployer, error) {
	contract, err := bindWhitelistDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistDeployer{WhitelistDeployerCaller: WhitelistDeployerCaller{contract: contract}, WhitelistDeployerTransactor: WhitelistDeployerTransactor{contract: contract}, WhitelistDeployerFilterer: WhitelistDeployerFilterer{contract: contract}}, nil
}

// NewWhitelistDeployerCaller creates a new read-only instance of WhitelistDeployer, bound to a specific deployed contract.
func NewWhitelistDeployerCaller(address common.Address, caller bind.ContractCaller) (*WhitelistDeployerCaller, error) {
	contract, err := bindWhitelistDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistDeployerCaller{contract: contract}, nil
}

// NewWhitelistDeployerTransactor creates a new write-only instance of WhitelistDeployer, bound to a specific deployed contract.
func NewWhitelistDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistDeployerTransactor, error) {
	contract, err := bindWhitelistDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistDeployerTransactor{contract: contract}, nil
}

// NewWhitelistDeployerFilterer creates a new log filterer instance of WhitelistDeployer, bound to a specific deployed contract.
func NewWhitelistDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistDeployerFilterer, error) {
	contract, err := bindWhitelistDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistDeployerFilterer{contract: contract}, nil
}

// bindWhitelistDeployer binds a generic wrapper to an already deployed contract.
func bindWhitelistDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistDeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistDeployer *WhitelistDeployerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistDeployer.Contract.WhitelistDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistDeployer *WhitelistDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.WhitelistDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistDeployer *WhitelistDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.WhitelistDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistDeployer *WhitelistDeployerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WhitelistDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistDeployer *WhitelistDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistDeployer *WhitelistDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_WhitelistDeployer *WhitelistDeployerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WhitelistDeployer.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_WhitelistDeployer *WhitelistDeployerSession) Admin() (common.Address, error) {
	return _WhitelistDeployer.Contract.Admin(&_WhitelistDeployer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_WhitelistDeployer *WhitelistDeployerCallerSession) Admin() (common.Address, error) {
	return _WhitelistDeployer.Contract.Admin(&_WhitelistDeployer.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistDeployer.contract.Call(opts, out, "isWhitelisted", _address)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _WhitelistDeployer.Contract.IsWhitelisted(&_WhitelistDeployer.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _WhitelistDeployer.Contract.IsWhitelisted(&_WhitelistDeployer.CallOpts, _address)
}

// WhitelistAll is a free data retrieval call binding the contract method 0x5d54e612.
//
// Solidity: function whitelistAll() constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerCaller) WhitelistAll(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistDeployer.contract.Call(opts, out, "whitelistAll")
	return *ret0, err
}

// WhitelistAll is a free data retrieval call binding the contract method 0x5d54e612.
//
// Solidity: function whitelistAll() constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerSession) WhitelistAll() (bool, error) {
	return _WhitelistDeployer.Contract.WhitelistAll(&_WhitelistDeployer.CallOpts)
}

// WhitelistAll is a free data retrieval call binding the contract method 0x5d54e612.
//
// Solidity: function whitelistAll() constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerCallerSession) WhitelistAll() (bool, error) {
	return _WhitelistDeployer.Contract.WhitelistAll(&_WhitelistDeployer.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WhitelistDeployer.contract.Call(opts, out, "whitelisted", arg0)
	return *ret0, err
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _WhitelistDeployer.Contract.Whitelisted(&_WhitelistDeployer.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) constant returns(bool)
func (_WhitelistDeployer *WhitelistDeployerCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _WhitelistDeployer.Contract.Whitelisted(&_WhitelistDeployer.CallOpts, arg0)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_WhitelistDeployer *WhitelistDeployerTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _WhitelistDeployer.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_WhitelistDeployer *WhitelistDeployerSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.ChangeAdmin(&_WhitelistDeployer.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_WhitelistDeployer *WhitelistDeployerTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.ChangeAdmin(&_WhitelistDeployer.TransactOpts, _newAdmin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_WhitelistDeployer *WhitelistDeployerTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistDeployer.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_WhitelistDeployer *WhitelistDeployerSession) RemoveAdmin() (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.RemoveAdmin(&_WhitelistDeployer.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_WhitelistDeployer *WhitelistDeployerTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.RemoveAdmin(&_WhitelistDeployer.TransactOpts)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address _address, bool _status) returns()
func (_WhitelistDeployer *WhitelistDeployerTransactor) Whitelist(opts *bind.TransactOpts, _address common.Address, _status bool) (*types.Transaction, error) {
	return _WhitelistDeployer.contract.Transact(opts, "whitelist", _address, _status)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address _address, bool _status) returns()
func (_WhitelistDeployer *WhitelistDeployerSession) Whitelist(_address common.Address, _status bool) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.Whitelist(&_WhitelistDeployer.TransactOpts, _address, _status)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address _address, bool _status) returns()
func (_WhitelistDeployer *WhitelistDeployerTransactorSession) Whitelist(_address common.Address, _status bool) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.Whitelist(&_WhitelistDeployer.TransactOpts, _address, _status)
}

// WhitelistAllAddresses is a paid mutator transaction binding the contract method 0x47a3ed0c.
//
// Solidity: function whitelistAllAddresses(bool _status) returns()
func (_WhitelistDeployer *WhitelistDeployerTransactor) WhitelistAllAddresses(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _WhitelistDeployer.contract.Transact(opts, "whitelistAllAddresses", _status)
}

// WhitelistAllAddresses is a paid mutator transaction binding the contract method 0x47a3ed0c.
//
// Solidity: function whitelistAllAddresses(bool _status) returns()
func (_WhitelistDeployer *WhitelistDeployerSession) WhitelistAllAddresses(_status bool) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.WhitelistAllAddresses(&_WhitelistDeployer.TransactOpts, _status)
}

// WhitelistAllAddresses is a paid mutator transaction binding the contract method 0x47a3ed0c.
//
// Solidity: function whitelistAllAddresses(bool _status) returns()
func (_WhitelistDeployer *WhitelistDeployerTransactorSession) WhitelistAllAddresses(_status bool) (*types.Transaction, error) {
	return _WhitelistDeployer.Contract.WhitelistAllAddresses(&_WhitelistDeployer.TransactOpts, _status)
}

// WhitelistDeployerAddressWhitelistedIterator is returned from FilterAddressWhitelisted and is used to iterate over the raw logs and unpacked data for AddressWhitelisted events raised by the WhitelistDeployer contract.
type WhitelistDeployerAddressWhitelistedIterator struct {
	Event *WhitelistDeployerAddressWhitelisted // Event containing the contract specifics and raw log

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
func (it *WhitelistDeployerAddressWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistDeployerAddressWhitelisted)
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
		it.Event = new(WhitelistDeployerAddressWhitelisted)
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
func (it *WhitelistDeployerAddressWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistDeployerAddressWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistDeployerAddressWhitelisted represents a AddressWhitelisted event raised by the WhitelistDeployer contract.
type WhitelistDeployerAddressWhitelisted struct {
	Address common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressWhitelisted is a free log retrieval operation binding the contract event 0xaf367c7d20ce5b2ab6da56afd0c9c39b00ba995263c60292a3e1ee3781fd4885.
//
// Solidity: event AddressWhitelisted(address indexed _address, bool indexed _status)
func (_WhitelistDeployer *WhitelistDeployerFilterer) FilterAddressWhitelisted(opts *bind.FilterOpts, _address []common.Address, _status []bool) (*WhitelistDeployerAddressWhitelistedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}
	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _WhitelistDeployer.contract.FilterLogs(opts, "AddressWhitelisted", _addressRule, _statusRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistDeployerAddressWhitelistedIterator{contract: _WhitelistDeployer.contract, event: "AddressWhitelisted", logs: logs, sub: sub}, nil
}

// WatchAddressWhitelisted is a free log subscription operation binding the contract event 0xaf367c7d20ce5b2ab6da56afd0c9c39b00ba995263c60292a3e1ee3781fd4885.
//
// Solidity: event AddressWhitelisted(address indexed _address, bool indexed _status)
func (_WhitelistDeployer *WhitelistDeployerFilterer) WatchAddressWhitelisted(opts *bind.WatchOpts, sink chan<- *WhitelistDeployerAddressWhitelisted, _address []common.Address, _status []bool) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}
	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _WhitelistDeployer.contract.WatchLogs(opts, "AddressWhitelisted", _addressRule, _statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistDeployerAddressWhitelisted)
				if err := _WhitelistDeployer.contract.UnpackLog(event, "AddressWhitelisted", log); err != nil {
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

// ParseAddressWhitelisted is a log parse operation binding the contract event 0xaf367c7d20ce5b2ab6da56afd0c9c39b00ba995263c60292a3e1ee3781fd4885.
//
// Solidity: event AddressWhitelisted(address indexed _address, bool indexed _status)
func (_WhitelistDeployer *WhitelistDeployerFilterer) ParseAddressWhitelisted(log types.Log) (*WhitelistDeployerAddressWhitelisted, error) {
	event := new(WhitelistDeployerAddressWhitelisted)
	if err := _WhitelistDeployer.contract.UnpackLog(event, "AddressWhitelisted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WhitelistDeployerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WhitelistDeployer contract.
type WhitelistDeployerAdminChangedIterator struct {
	Event *WhitelistDeployerAdminChanged // Event containing the contract specifics and raw log

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
func (it *WhitelistDeployerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistDeployerAdminChanged)
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
		it.Event = new(WhitelistDeployerAdminChanged)
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
func (it *WhitelistDeployerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistDeployerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistDeployerAdminChanged represents a AdminChanged event raised by the WhitelistDeployer contract.
type WhitelistDeployerAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_WhitelistDeployer *WhitelistDeployerFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*WhitelistDeployerAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _WhitelistDeployer.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistDeployerAdminChangedIterator{contract: _WhitelistDeployer.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_WhitelistDeployer *WhitelistDeployerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WhitelistDeployerAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _WhitelistDeployer.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistDeployerAdminChanged)
				if err := _WhitelistDeployer.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_WhitelistDeployer *WhitelistDeployerFilterer) ParseAdminChanged(log types.Log) (*WhitelistDeployerAdminChanged, error) {
	event := new(WhitelistDeployerAdminChanged)
	if err := _WhitelistDeployer.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WhitelistDeployerAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the WhitelistDeployer contract.
type WhitelistDeployerAdminRemovedIterator struct {
	Event *WhitelistDeployerAdminRemoved // Event containing the contract specifics and raw log

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
func (it *WhitelistDeployerAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistDeployerAdminRemoved)
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
		it.Event = new(WhitelistDeployerAdminRemoved)
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
func (it *WhitelistDeployerAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistDeployerAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistDeployerAdminRemoved represents a AdminRemoved event raised by the WhitelistDeployer contract.
type WhitelistDeployerAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_WhitelistDeployer *WhitelistDeployerFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*WhitelistDeployerAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _WhitelistDeployer.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistDeployerAdminRemovedIterator{contract: _WhitelistDeployer.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_WhitelistDeployer *WhitelistDeployerFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistDeployerAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _WhitelistDeployer.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistDeployerAdminRemoved)
				if err := _WhitelistDeployer.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_WhitelistDeployer *WhitelistDeployerFilterer) ParseAdminRemoved(log types.Log) (*WhitelistDeployerAdminRemoved, error) {
	event := new(WhitelistDeployerAdminRemoved)
	if err := _WhitelistDeployer.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WhitelistDeployerWhitelistAllChangeIterator is returned from FilterWhitelistAllChange and is used to iterate over the raw logs and unpacked data for WhitelistAllChange events raised by the WhitelistDeployer contract.
type WhitelistDeployerWhitelistAllChangeIterator struct {
	Event *WhitelistDeployerWhitelistAllChange // Event containing the contract specifics and raw log

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
func (it *WhitelistDeployerWhitelistAllChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistDeployerWhitelistAllChange)
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
		it.Event = new(WhitelistDeployerWhitelistAllChange)
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
func (it *WhitelistDeployerWhitelistAllChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistDeployerWhitelistAllChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistDeployerWhitelistAllChange represents a WhitelistAllChange event raised by the WhitelistDeployer contract.
type WhitelistDeployerWhitelistAllChange struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAllChange is a free log retrieval operation binding the contract event 0x01d0151a76b6ade3851c417c9a2511eefa3e319bc884bae10c3dc989e7e8d85e.
//
// Solidity: event WhitelistAllChange(bool indexed _status)
func (_WhitelistDeployer *WhitelistDeployerFilterer) FilterWhitelistAllChange(opts *bind.FilterOpts, _status []bool) (*WhitelistDeployerWhitelistAllChangeIterator, error) {

	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _WhitelistDeployer.contract.FilterLogs(opts, "WhitelistAllChange", _statusRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistDeployerWhitelistAllChangeIterator{contract: _WhitelistDeployer.contract, event: "WhitelistAllChange", logs: logs, sub: sub}, nil
}

// WatchWhitelistAllChange is a free log subscription operation binding the contract event 0x01d0151a76b6ade3851c417c9a2511eefa3e319bc884bae10c3dc989e7e8d85e.
//
// Solidity: event WhitelistAllChange(bool indexed _status)
func (_WhitelistDeployer *WhitelistDeployerFilterer) WatchWhitelistAllChange(opts *bind.WatchOpts, sink chan<- *WhitelistDeployerWhitelistAllChange, _status []bool) (event.Subscription, error) {

	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _WhitelistDeployer.contract.WatchLogs(opts, "WhitelistAllChange", _statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistDeployerWhitelistAllChange)
				if err := _WhitelistDeployer.contract.UnpackLog(event, "WhitelistAllChange", log); err != nil {
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

// ParseWhitelistAllChange is a log parse operation binding the contract event 0x01d0151a76b6ade3851c417c9a2511eefa3e319bc884bae10c3dc989e7e8d85e.
//
// Solidity: event WhitelistAllChange(bool indexed _status)
func (_WhitelistDeployer *WhitelistDeployerFilterer) ParseWhitelistAllChange(log types.Log) (*WhitelistDeployerWhitelistAllChange, error) {
	event := new(WhitelistDeployerWhitelistAllChange)
	if err := _WhitelistDeployer.contract.UnpackLog(event, "WhitelistAllChange", log); err != nil {
		return nil, err
	}
	return event, nil
}
