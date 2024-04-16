// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executionservice

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

// AccessControlUpgradeableMetaData contains all meta data concerning the AccessControlUpgradeable contract.
var AccessControlUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// AccessControlUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlUpgradeableMetaData.ABI instead.
var AccessControlUpgradeableABI = AccessControlUpgradeableMetaData.ABI

// Deprecated: Use AccessControlUpgradeableMetaData.Sigs instead.
// AccessControlUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlUpgradeableFuncSigs = AccessControlUpgradeableMetaData.Sigs

// AccessControlUpgradeable is an auto generated Go binding around an Ethereum contract.
type AccessControlUpgradeable struct {
	AccessControlUpgradeableCaller     // Read-only binding to the contract
	AccessControlUpgradeableTransactor // Write-only binding to the contract
	AccessControlUpgradeableFilterer   // Log filterer for contract events
}

// AccessControlUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlUpgradeableSession struct {
	Contract     *AccessControlUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccessControlUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlUpgradeableCallerSession struct {
	Contract *AccessControlUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AccessControlUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlUpgradeableTransactorSession struct {
	Contract     *AccessControlUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AccessControlUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlUpgradeableRaw struct {
	Contract *AccessControlUpgradeable // Generic contract binding to access the raw methods on
}

// AccessControlUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlUpgradeableCallerRaw struct {
	Contract *AccessControlUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlUpgradeableTransactorRaw struct {
	Contract *AccessControlUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlUpgradeable creates a new instance of AccessControlUpgradeable, bound to a specific deployed contract.
func NewAccessControlUpgradeable(address common.Address, backend bind.ContractBackend) (*AccessControlUpgradeable, error) {
	contract, err := bindAccessControlUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeable{AccessControlUpgradeableCaller: AccessControlUpgradeableCaller{contract: contract}, AccessControlUpgradeableTransactor: AccessControlUpgradeableTransactor{contract: contract}, AccessControlUpgradeableFilterer: AccessControlUpgradeableFilterer{contract: contract}}, nil
}

// NewAccessControlUpgradeableCaller creates a new read-only instance of AccessControlUpgradeable, bound to a specific deployed contract.
func NewAccessControlUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AccessControlUpgradeableCaller, error) {
	contract, err := bindAccessControlUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableCaller{contract: contract}, nil
}

// NewAccessControlUpgradeableTransactor creates a new write-only instance of AccessControlUpgradeable, bound to a specific deployed contract.
func NewAccessControlUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlUpgradeableTransactor, error) {
	contract, err := bindAccessControlUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableTransactor{contract: contract}, nil
}

// NewAccessControlUpgradeableFilterer creates a new log filterer instance of AccessControlUpgradeable, bound to a specific deployed contract.
func NewAccessControlUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlUpgradeableFilterer, error) {
	contract, err := bindAccessControlUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableFilterer{contract: contract}, nil
}

// bindAccessControlUpgradeable binds a generic wrapper to an already deployed contract.
func bindAccessControlUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlUpgradeable *AccessControlUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlUpgradeable.Contract.AccessControlUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlUpgradeable *AccessControlUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.AccessControlUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlUpgradeable *AccessControlUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.AccessControlUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlUpgradeable.Contract.DEFAULTADMINROLE(&_AccessControlUpgradeable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlUpgradeable.Contract.DEFAULTADMINROLE(&_AccessControlUpgradeable.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlUpgradeable.Contract.GetRoleAdmin(&_AccessControlUpgradeable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlUpgradeable.Contract.GetRoleAdmin(&_AccessControlUpgradeable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlUpgradeable.Contract.HasRole(&_AccessControlUpgradeable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlUpgradeable.Contract.HasRole(&_AccessControlUpgradeable.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControlUpgradeable.Contract.SupportsInterface(&_AccessControlUpgradeable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControlUpgradeable.Contract.SupportsInterface(&_AccessControlUpgradeable.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.GrantRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.GrantRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RenounceRole(&_AccessControlUpgradeable.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RenounceRole(&_AccessControlUpgradeable.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RevokeRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RevokeRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// AccessControlUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableInitializedIterator struct {
	Event *AccessControlUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeableInitialized)
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
		it.Event = new(AccessControlUpgradeableInitialized)
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
func (it *AccessControlUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeableInitialized represents a Initialized event raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccessControlUpgradeableInitializedIterator, error) {

	logs, sub, err := _AccessControlUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableInitializedIterator{contract: _AccessControlUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _AccessControlUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeableInitialized)
				if err := _AccessControlUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) ParseInitialized(log types.Log) (*AccessControlUpgradeableInitialized, error) {
	event := new(AccessControlUpgradeableInitialized)
	if err := _AccessControlUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlUpgradeableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleAdminChangedIterator struct {
	Event *AccessControlUpgradeableRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeableRoleAdminChanged)
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
		it.Event = new(AccessControlUpgradeableRoleAdminChanged)
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
func (it *AccessControlUpgradeableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeableRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlUpgradeableRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControlUpgradeable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableRoleAdminChangedIterator{contract: _AccessControlUpgradeable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControlUpgradeable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeableRoleAdminChanged)
				if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlUpgradeableRoleAdminChanged, error) {
	event := new(AccessControlUpgradeableRoleAdminChanged)
	if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlUpgradeableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleGrantedIterator struct {
	Event *AccessControlUpgradeableRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeableRoleGranted)
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
		it.Event = new(AccessControlUpgradeableRoleGranted)
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
func (it *AccessControlUpgradeableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeableRoleGranted represents a RoleGranted event raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlUpgradeableRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlUpgradeable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableRoleGrantedIterator{contract: _AccessControlUpgradeable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlUpgradeable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeableRoleGranted)
				if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) ParseRoleGranted(log types.Log) (*AccessControlUpgradeableRoleGranted, error) {
	event := new(AccessControlUpgradeableRoleGranted)
	if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlUpgradeableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleRevokedIterator struct {
	Event *AccessControlUpgradeableRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeableRoleRevoked)
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
		it.Event = new(AccessControlUpgradeableRoleRevoked)
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
func (it *AccessControlUpgradeableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeableRoleRevoked represents a RoleRevoked event raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlUpgradeableRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlUpgradeable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableRoleRevokedIterator{contract: _AccessControlUpgradeable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlUpgradeable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeableRoleRevoked)
				if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) ParseRoleRevoked(log types.Log) (*AccessControlUpgradeableRoleRevoked, error) {
	event := new(AccessControlUpgradeableRoleRevoked)
	if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContextUpgradeable contract.
type ContextUpgradeableInitializedIterator struct {
	Event *ContextUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ContextUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContextUpgradeableInitialized)
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
		it.Event = new(ContextUpgradeableInitialized)
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
func (it *ContextUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContextUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContextUpgradeableInitialized represents a Initialized event raised by the ContextUpgradeable contract.
type ContextUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContextUpgradeableInitializedIterator, error) {

	logs, sub, err := _ContextUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableInitializedIterator{contract: _ContextUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContextUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ContextUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContextUpgradeableInitialized)
				if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) ParseInitialized(log types.Log) (*ContextUpgradeableInitialized, error) {
	event := new(ContextUpgradeableInitialized)
	if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC165UpgradeableMetaData contains all meta data concerning the ERC165Upgradeable contract.
var ERC165UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// ERC165UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165UpgradeableMetaData.ABI instead.
var ERC165UpgradeableABI = ERC165UpgradeableMetaData.ABI

// Deprecated: Use ERC165UpgradeableMetaData.Sigs instead.
// ERC165UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var ERC165UpgradeableFuncSigs = ERC165UpgradeableMetaData.Sigs

// ERC165Upgradeable is an auto generated Go binding around an Ethereum contract.
type ERC165Upgradeable struct {
	ERC165UpgradeableCaller     // Read-only binding to the contract
	ERC165UpgradeableTransactor // Write-only binding to the contract
	ERC165UpgradeableFilterer   // Log filterer for contract events
}

// ERC165UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165UpgradeableSession struct {
	Contract     *ERC165Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC165UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165UpgradeableCallerSession struct {
	Contract *ERC165UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ERC165UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165UpgradeableTransactorSession struct {
	Contract     *ERC165UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC165UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165UpgradeableRaw struct {
	Contract *ERC165Upgradeable // Generic contract binding to access the raw methods on
}

// ERC165UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165UpgradeableCallerRaw struct {
	Contract *ERC165UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC165UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165UpgradeableTransactorRaw struct {
	Contract *ERC165UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165Upgradeable creates a new instance of ERC165Upgradeable, bound to a specific deployed contract.
func NewERC165Upgradeable(address common.Address, backend bind.ContractBackend) (*ERC165Upgradeable, error) {
	contract, err := bindERC165Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165Upgradeable{ERC165UpgradeableCaller: ERC165UpgradeableCaller{contract: contract}, ERC165UpgradeableTransactor: ERC165UpgradeableTransactor{contract: contract}, ERC165UpgradeableFilterer: ERC165UpgradeableFilterer{contract: contract}}, nil
}

// NewERC165UpgradeableCaller creates a new read-only instance of ERC165Upgradeable, bound to a specific deployed contract.
func NewERC165UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC165UpgradeableCaller, error) {
	contract, err := bindERC165Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165UpgradeableCaller{contract: contract}, nil
}

// NewERC165UpgradeableTransactor creates a new write-only instance of ERC165Upgradeable, bound to a specific deployed contract.
func NewERC165UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC165UpgradeableTransactor, error) {
	contract, err := bindERC165Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165UpgradeableTransactor{contract: contract}, nil
}

// NewERC165UpgradeableFilterer creates a new log filterer instance of ERC165Upgradeable, bound to a specific deployed contract.
func NewERC165UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC165UpgradeableFilterer, error) {
	contract, err := bindERC165Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165UpgradeableFilterer{contract: contract}, nil
}

// bindERC165Upgradeable binds a generic wrapper to an already deployed contract.
func bindERC165Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC165UpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Upgradeable *ERC165UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165Upgradeable.Contract.ERC165UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Upgradeable *ERC165UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Upgradeable.Contract.ERC165UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Upgradeable *ERC165UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Upgradeable.Contract.ERC165UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Upgradeable *ERC165UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Upgradeable *ERC165UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Upgradeable *ERC165UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165Upgradeable *ERC165UpgradeableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165Upgradeable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165Upgradeable *ERC165UpgradeableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165Upgradeable.Contract.SupportsInterface(&_ERC165Upgradeable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165Upgradeable *ERC165UpgradeableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165Upgradeable.Contract.SupportsInterface(&_ERC165Upgradeable.CallOpts, interfaceId)
}

// ERC165UpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC165Upgradeable contract.
type ERC165UpgradeableInitializedIterator struct {
	Event *ERC165UpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ERC165UpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC165UpgradeableInitialized)
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
		it.Event = new(ERC165UpgradeableInitialized)
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
func (it *ERC165UpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC165UpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC165UpgradeableInitialized represents a Initialized event raised by the ERC165Upgradeable contract.
type ERC165UpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC165Upgradeable *ERC165UpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC165UpgradeableInitializedIterator, error) {

	logs, sub, err := _ERC165Upgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC165UpgradeableInitializedIterator{contract: _ERC165Upgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC165Upgradeable *ERC165UpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC165UpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC165Upgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC165UpgradeableInitialized)
				if err := _ERC165Upgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC165Upgradeable *ERC165UpgradeableFilterer) ParseInitialized(log types.Log) (*ERC165UpgradeableInitialized, error) {
	event := new(ERC165UpgradeableInitialized)
	if err := _ERC165Upgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlMetaData contains all meta data concerning the IAccessControl contract.
var IAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
	},
}

// IAccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlMetaData.ABI instead.
var IAccessControlABI = IAccessControlMetaData.ABI

// Deprecated: Use IAccessControlMetaData.Sigs instead.
// IAccessControlFuncSigs maps the 4-byte function signature to its string representation.
var IAccessControlFuncSigs = IAccessControlMetaData.Sigs

// IAccessControl is an auto generated Go binding around an Ethereum contract.
type IAccessControl struct {
	IAccessControlCaller     // Read-only binding to the contract
	IAccessControlTransactor // Write-only binding to the contract
	IAccessControlFilterer   // Log filterer for contract events
}

// IAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlSession struct {
	Contract     *IAccessControl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlCallerSession struct {
	Contract *IAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlTransactorSession struct {
	Contract     *IAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlRaw struct {
	Contract *IAccessControl // Generic contract binding to access the raw methods on
}

// IAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlCallerRaw struct {
	Contract *IAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlTransactorRaw struct {
	Contract *IAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControl creates a new instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControl(address common.Address, backend bind.ContractBackend) (*IAccessControl, error) {
	contract, err := bindIAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControl{IAccessControlCaller: IAccessControlCaller{contract: contract}, IAccessControlTransactor: IAccessControlTransactor{contract: contract}, IAccessControlFilterer: IAccessControlFilterer{contract: contract}}, nil
}

// NewIAccessControlCaller creates a new read-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlCaller, error) {
	contract, err := bindIAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlCaller{contract: contract}, nil
}

// NewIAccessControlTransactor creates a new write-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlTransactor, error) {
	contract, err := bindIAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlTransactor{contract: contract}, nil
}

// NewIAccessControlFilterer creates a new log filterer instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlFilterer, error) {
	contract, err := bindIAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlFilterer{contract: contract}, nil
}

// bindIAccessControl binds a generic wrapper to an already deployed contract.
func bindIAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.IAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// IAccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IAccessControl contract.
type IAccessControlRoleAdminChangedIterator struct {
	Event *IAccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleAdminChanged)
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
		it.Event = new(IAccessControlRoleAdminChanged)
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
func (it *IAccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the IAccessControl contract.
type IAccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IAccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleAdminChangedIterator{contract: _IAccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleAdminChanged)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*IAccessControlRoleAdminChanged, error) {
	event := new(IAccessControlRoleAdminChanged)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IAccessControl contract.
type IAccessControlRoleGrantedIterator struct {
	Event *IAccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleGranted)
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
		it.Event = new(IAccessControlRoleGranted)
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
func (it *IAccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleGranted represents a RoleGranted event raised by the IAccessControl contract.
type IAccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleGrantedIterator{contract: _IAccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleGranted)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) ParseRoleGranted(log types.Log) (*IAccessControlRoleGranted, error) {
	event := new(IAccessControlRoleGranted)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IAccessControl contract.
type IAccessControlRoleRevokedIterator struct {
	Event *IAccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleRevoked)
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
		it.Event = new(IAccessControlRoleRevoked)
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
func (it *IAccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleRevoked represents a RoleRevoked event raised by the IAccessControl contract.
type IAccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleRevokedIterator{contract: _IAccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleRevoked)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) ParseRoleRevoked(log types.Log) (*IAccessControlRoleRevoked, error) {
	event := new(IAccessControlRoleRevoked)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// Deprecated: Use IERC165MetaData.Sigs instead.
// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = IERC165MetaData.Sigs

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IExecutionServiceMetaData contains all meta data concerning the IExecutionService contract.
var IExecutionServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62014bad": "executorEOA()",
		"96fda4da": "getExecutionFee(uint64,uint256,bytes)",
		"592a8799": "requestExecution(uint64,uint256,bytes32,uint256,bytes)",
	},
}

// IExecutionServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionServiceMetaData.ABI instead.
var IExecutionServiceABI = IExecutionServiceMetaData.ABI

// Deprecated: Use IExecutionServiceMetaData.Sigs instead.
// IExecutionServiceFuncSigs maps the 4-byte function signature to its string representation.
var IExecutionServiceFuncSigs = IExecutionServiceMetaData.Sigs

// IExecutionService is an auto generated Go binding around an Ethereum contract.
type IExecutionService struct {
	IExecutionServiceCaller     // Read-only binding to the contract
	IExecutionServiceTransactor // Write-only binding to the contract
	IExecutionServiceFilterer   // Log filterer for contract events
}

// IExecutionServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionServiceSession struct {
	Contract     *IExecutionService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IExecutionServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionServiceCallerSession struct {
	Contract *IExecutionServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IExecutionServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionServiceTransactorSession struct {
	Contract     *IExecutionServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IExecutionServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionServiceRaw struct {
	Contract *IExecutionService // Generic contract binding to access the raw methods on
}

// IExecutionServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionServiceCallerRaw struct {
	Contract *IExecutionServiceCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionServiceTransactorRaw struct {
	Contract *IExecutionServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecutionService creates a new instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionService(address common.Address, backend bind.ContractBackend) (*IExecutionService, error) {
	contract, err := bindIExecutionService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecutionService{IExecutionServiceCaller: IExecutionServiceCaller{contract: contract}, IExecutionServiceTransactor: IExecutionServiceTransactor{contract: contract}, IExecutionServiceFilterer: IExecutionServiceFilterer{contract: contract}}, nil
}

// NewIExecutionServiceCaller creates a new read-only instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceCaller(address common.Address, caller bind.ContractCaller) (*IExecutionServiceCaller, error) {
	contract, err := bindIExecutionService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceCaller{contract: contract}, nil
}

// NewIExecutionServiceTransactor creates a new write-only instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionServiceTransactor, error) {
	contract, err := bindIExecutionService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceTransactor{contract: contract}, nil
}

// NewIExecutionServiceFilterer creates a new log filterer instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionServiceFilterer, error) {
	contract, err := bindIExecutionService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceFilterer{contract: contract}, nil
}

// bindIExecutionService binds a generic wrapper to an already deployed contract.
func bindIExecutionService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExecutionServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionService *IExecutionServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionService.Contract.IExecutionServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionService *IExecutionServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionService.Contract.IExecutionServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionService *IExecutionServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionService.Contract.IExecutionServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionService *IExecutionServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionService *IExecutionServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionService *IExecutionServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionService.Contract.contract.Transact(opts, method, params...)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceCaller) ExecutorEOA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IExecutionService.contract.Call(opts, &out, "executorEOA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceSession) ExecutorEOA() (common.Address, error) {
	return _IExecutionService.Contract.ExecutorEOA(&_IExecutionService.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceCallerSession) ExecutorEOA() (common.Address, error) {
	return _IExecutionService.Contract.ExecutorEOA(&_IExecutionService.CallOpts)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceCaller) GetExecutionFee(opts *bind.CallOpts, dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionService.contract.Call(opts, &out, "getExecutionFee", dstChainId, txPayloadSize, options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceSession) GetExecutionFee(dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _IExecutionService.Contract.GetExecutionFee(&_IExecutionService.CallOpts, dstChainId, txPayloadSize, options)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceCallerSession) GetExecutionFee(dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _IExecutionService.Contract.GetExecutionFee(&_IExecutionService.CallOpts, dstChainId, txPayloadSize, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceTransactor) RequestExecution(opts *bind.TransactOpts, dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.contract.Transact(opts, "requestExecution", dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceSession) RequestExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.Contract.RequestExecution(&_IExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceTransactorSession) RequestExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.Contract.RequestExecution(&_IExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// IGasOracleMetaData contains all meta data concerning the IGasOracle contract.
var IGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40658a74": "convertRemoteValueToLocalUnits(uint64,uint256)",
		"bf495c88": "estimateTxCostInLocalUnits(uint64,uint256,uint256)",
		"b376a688": "estimateTxCostInRemoteUnits(uint64,uint256,uint256)",
	},
}

// IGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IGasOracleMetaData.ABI instead.
var IGasOracleABI = IGasOracleMetaData.ABI

// Deprecated: Use IGasOracleMetaData.Sigs instead.
// IGasOracleFuncSigs maps the 4-byte function signature to its string representation.
var IGasOracleFuncSigs = IGasOracleMetaData.Sigs

// IGasOracle is an auto generated Go binding around an Ethereum contract.
type IGasOracle struct {
	IGasOracleCaller     // Read-only binding to the contract
	IGasOracleTransactor // Write-only binding to the contract
	IGasOracleFilterer   // Log filterer for contract events
}

// IGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGasOracleSession struct {
	Contract     *IGasOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGasOracleCallerSession struct {
	Contract *IGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGasOracleTransactorSession struct {
	Contract     *IGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGasOracleRaw struct {
	Contract *IGasOracle // Generic contract binding to access the raw methods on
}

// IGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGasOracleCallerRaw struct {
	Contract *IGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGasOracleTransactorRaw struct {
	Contract *IGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGasOracle creates a new instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracle(address common.Address, backend bind.ContractBackend) (*IGasOracle, error) {
	contract, err := bindIGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGasOracle{IGasOracleCaller: IGasOracleCaller{contract: contract}, IGasOracleTransactor: IGasOracleTransactor{contract: contract}, IGasOracleFilterer: IGasOracleFilterer{contract: contract}}, nil
}

// NewIGasOracleCaller creates a new read-only instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleCaller(address common.Address, caller bind.ContractCaller) (*IGasOracleCaller, error) {
	contract, err := bindIGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGasOracleCaller{contract: contract}, nil
}

// NewIGasOracleTransactor creates a new write-only instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IGasOracleTransactor, error) {
	contract, err := bindIGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGasOracleTransactor{contract: contract}, nil
}

// NewIGasOracleFilterer creates a new log filterer instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IGasOracleFilterer, error) {
	contract, err := bindIGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGasOracleFilterer{contract: contract}, nil
}

// bindIGasOracle binds a generic wrapper to an already deployed contract.
func bindIGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGasOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasOracle *IGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasOracle.Contract.IGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasOracle *IGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasOracle.Contract.IGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasOracle *IGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasOracle.Contract.IGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasOracle *IGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasOracle *IGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasOracle *IGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasOracle.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId uint64, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// ISynapseExecutionServiceV1MetaData contains all meta data concerning the ISynapseExecutionServiceV1 contract.
var ISynapseExecutionServiceV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"SynapseExecutionService__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__GasOracleNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"SynapseExecutionService__OptionsVersionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalMarkup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executorEOA_\",\"type\":\"address\"}],\"name\":\"setExecutorEOA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalMarkup_\",\"type\":\"uint256\"}],\"name\":\"setGlobalMarkup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62014bad": "executorEOA()",
		"5d62a8dd": "gasOracle()",
		"96fda4da": "getExecutionFee(uint64,uint256,bytes)",
		"efd07ec2": "globalMarkup()",
		"592a8799": "requestExecution(uint64,uint256,bytes32,uint256,bytes)",
		"2d54566c": "setExecutorEOA(address)",
		"a87b8152": "setGasOracle(address)",
		"cf4f578f": "setGlobalMarkup(uint256)",
	},
}

// ISynapseExecutionServiceV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseExecutionServiceV1MetaData.ABI instead.
var ISynapseExecutionServiceV1ABI = ISynapseExecutionServiceV1MetaData.ABI

// Deprecated: Use ISynapseExecutionServiceV1MetaData.Sigs instead.
// ISynapseExecutionServiceV1FuncSigs maps the 4-byte function signature to its string representation.
var ISynapseExecutionServiceV1FuncSigs = ISynapseExecutionServiceV1MetaData.Sigs

// ISynapseExecutionServiceV1 is an auto generated Go binding around an Ethereum contract.
type ISynapseExecutionServiceV1 struct {
	ISynapseExecutionServiceV1Caller     // Read-only binding to the contract
	ISynapseExecutionServiceV1Transactor // Write-only binding to the contract
	ISynapseExecutionServiceV1Filterer   // Log filterer for contract events
}

// ISynapseExecutionServiceV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseExecutionServiceV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseExecutionServiceV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseExecutionServiceV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseExecutionServiceV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseExecutionServiceV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseExecutionServiceV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseExecutionServiceV1Session struct {
	Contract     *ISynapseExecutionServiceV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ISynapseExecutionServiceV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseExecutionServiceV1CallerSession struct {
	Contract *ISynapseExecutionServiceV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ISynapseExecutionServiceV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseExecutionServiceV1TransactorSession struct {
	Contract     *ISynapseExecutionServiceV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ISynapseExecutionServiceV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseExecutionServiceV1Raw struct {
	Contract *ISynapseExecutionServiceV1 // Generic contract binding to access the raw methods on
}

// ISynapseExecutionServiceV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseExecutionServiceV1CallerRaw struct {
	Contract *ISynapseExecutionServiceV1Caller // Generic read-only contract binding to access the raw methods on
}

// ISynapseExecutionServiceV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseExecutionServiceV1TransactorRaw struct {
	Contract *ISynapseExecutionServiceV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseExecutionServiceV1 creates a new instance of ISynapseExecutionServiceV1, bound to a specific deployed contract.
func NewISynapseExecutionServiceV1(address common.Address, backend bind.ContractBackend) (*ISynapseExecutionServiceV1, error) {
	contract, err := bindISynapseExecutionServiceV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseExecutionServiceV1{ISynapseExecutionServiceV1Caller: ISynapseExecutionServiceV1Caller{contract: contract}, ISynapseExecutionServiceV1Transactor: ISynapseExecutionServiceV1Transactor{contract: contract}, ISynapseExecutionServiceV1Filterer: ISynapseExecutionServiceV1Filterer{contract: contract}}, nil
}

// NewISynapseExecutionServiceV1Caller creates a new read-only instance of ISynapseExecutionServiceV1, bound to a specific deployed contract.
func NewISynapseExecutionServiceV1Caller(address common.Address, caller bind.ContractCaller) (*ISynapseExecutionServiceV1Caller, error) {
	contract, err := bindISynapseExecutionServiceV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseExecutionServiceV1Caller{contract: contract}, nil
}

// NewISynapseExecutionServiceV1Transactor creates a new write-only instance of ISynapseExecutionServiceV1, bound to a specific deployed contract.
func NewISynapseExecutionServiceV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseExecutionServiceV1Transactor, error) {
	contract, err := bindISynapseExecutionServiceV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseExecutionServiceV1Transactor{contract: contract}, nil
}

// NewISynapseExecutionServiceV1Filterer creates a new log filterer instance of ISynapseExecutionServiceV1, bound to a specific deployed contract.
func NewISynapseExecutionServiceV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseExecutionServiceV1Filterer, error) {
	contract, err := bindISynapseExecutionServiceV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseExecutionServiceV1Filterer{contract: contract}, nil
}

// bindISynapseExecutionServiceV1 binds a generic wrapper to an already deployed contract.
func bindISynapseExecutionServiceV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISynapseExecutionServiceV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseExecutionServiceV1.Contract.ISynapseExecutionServiceV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.ISynapseExecutionServiceV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.ISynapseExecutionServiceV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseExecutionServiceV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.contract.Transact(opts, method, params...)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Caller) ExecutorEOA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISynapseExecutionServiceV1.contract.Call(opts, &out, "executorEOA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) ExecutorEOA() (common.Address, error) {
	return _ISynapseExecutionServiceV1.Contract.ExecutorEOA(&_ISynapseExecutionServiceV1.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1CallerSession) ExecutorEOA() (common.Address, error) {
	return _ISynapseExecutionServiceV1.Contract.ExecutorEOA(&_ISynapseExecutionServiceV1.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Caller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISynapseExecutionServiceV1.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) GasOracle() (common.Address, error) {
	return _ISynapseExecutionServiceV1.Contract.GasOracle(&_ISynapseExecutionServiceV1.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1CallerSession) GasOracle() (common.Address, error) {
	return _ISynapseExecutionServiceV1.Contract.GasOracle(&_ISynapseExecutionServiceV1.CallOpts)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Caller) GetExecutionFee(opts *bind.CallOpts, dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseExecutionServiceV1.contract.Call(opts, &out, "getExecutionFee", dstChainId, txPayloadSize, options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) GetExecutionFee(dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _ISynapseExecutionServiceV1.Contract.GetExecutionFee(&_ISynapseExecutionServiceV1.CallOpts, dstChainId, txPayloadSize, options)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1CallerSession) GetExecutionFee(dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _ISynapseExecutionServiceV1.Contract.GetExecutionFee(&_ISynapseExecutionServiceV1.CallOpts, dstChainId, txPayloadSize, options)
}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Caller) GlobalMarkup(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseExecutionServiceV1.contract.Call(opts, &out, "globalMarkup")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) GlobalMarkup() (*big.Int, error) {
	return _ISynapseExecutionServiceV1.Contract.GlobalMarkup(&_ISynapseExecutionServiceV1.CallOpts)
}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1CallerSession) GlobalMarkup() (*big.Int, error) {
	return _ISynapseExecutionServiceV1.Contract.GlobalMarkup(&_ISynapseExecutionServiceV1.CallOpts)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Transactor) RequestExecution(opts *bind.TransactOpts, dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.contract.Transact(opts, "requestExecution", dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) RequestExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.RequestExecution(&_ISynapseExecutionServiceV1.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1TransactorSession) RequestExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.RequestExecution(&_ISynapseExecutionServiceV1.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Transactor) SetExecutorEOA(opts *bind.TransactOpts, executorEOA_ common.Address) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.contract.Transact(opts, "setExecutorEOA", executorEOA_)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) SetExecutorEOA(executorEOA_ common.Address) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.SetExecutorEOA(&_ISynapseExecutionServiceV1.TransactOpts, executorEOA_)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1TransactorSession) SetExecutorEOA(executorEOA_ common.Address) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.SetExecutorEOA(&_ISynapseExecutionServiceV1.TransactOpts, executorEOA_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Transactor) SetGasOracle(opts *bind.TransactOpts, gasOracle_ common.Address) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.contract.Transact(opts, "setGasOracle", gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.SetGasOracle(&_ISynapseExecutionServiceV1.TransactOpts, gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1TransactorSession) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.SetGasOracle(&_ISynapseExecutionServiceV1.TransactOpts, gasOracle_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Transactor) SetGlobalMarkup(opts *bind.TransactOpts, globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.contract.Transact(opts, "setGlobalMarkup", globalMarkup_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) SetGlobalMarkup(globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.SetGlobalMarkup(&_ISynapseExecutionServiceV1.TransactOpts, globalMarkup_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1TransactorSession) SetGlobalMarkup(globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.SetGlobalMarkup(&_ISynapseExecutionServiceV1.TransactOpts, globalMarkup_)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InitializableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

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
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
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
		it.Event = new(InitializableInitialized)
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
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {

	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {

	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptionsLibMetaData contains all meta data concerning the OptionsLib contract.
var OptionsLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c78452e4f1ebd2dbe8c4b9372a915af6ede1682b9db9f7b197b0ee1738074bcf64736f6c63430008140033",
}

// OptionsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OptionsLibMetaData.ABI instead.
var OptionsLibABI = OptionsLibMetaData.ABI

// OptionsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptionsLibMetaData.Bin instead.
var OptionsLibBin = OptionsLibMetaData.Bin

// DeployOptionsLib deploys a new Ethereum contract, binding an instance of OptionsLib to it.
func DeployOptionsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptionsLib, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptionsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// OptionsLib is an auto generated Go binding around an Ethereum contract.
type OptionsLib struct {
	OptionsLibCaller     // Read-only binding to the contract
	OptionsLibTransactor // Write-only binding to the contract
	OptionsLibFilterer   // Log filterer for contract events
}

// OptionsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptionsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptionsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptionsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptionsLibSession struct {
	Contract     *OptionsLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptionsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptionsLibCallerSession struct {
	Contract *OptionsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OptionsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptionsLibTransactorSession struct {
	Contract     *OptionsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OptionsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptionsLibRaw struct {
	Contract *OptionsLib // Generic contract binding to access the raw methods on
}

// OptionsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptionsLibCallerRaw struct {
	Contract *OptionsLibCaller // Generic read-only contract binding to access the raw methods on
}

// OptionsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptionsLibTransactorRaw struct {
	Contract *OptionsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptionsLib creates a new instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLib(address common.Address, backend bind.ContractBackend) (*OptionsLib, error) {
	contract, err := bindOptionsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// NewOptionsLibCaller creates a new read-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibCaller(address common.Address, caller bind.ContractCaller) (*OptionsLibCaller, error) {
	contract, err := bindOptionsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibCaller{contract: contract}, nil
}

// NewOptionsLibTransactor creates a new write-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OptionsLibTransactor, error) {
	contract, err := bindOptionsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibTransactor{contract: contract}, nil
}

// NewOptionsLibFilterer creates a new log filterer instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OptionsLibFilterer, error) {
	contract, err := bindOptionsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptionsLibFilterer{contract: contract}, nil
}

// bindOptionsLib binds a generic wrapper to an already deployed contract.
func bindOptionsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.OptionsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transact(opts, method, params...)
}

// SynapseExecutionServiceEventsMetaData contains all meta data concerning the SynapseExecutionServiceEvents contract.
var SynapseExecutionServiceEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executorEOA\",\"type\":\"address\"}],\"name\":\"ExecutorEOASet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalMarkup\",\"type\":\"uint256\"}],\"name\":\"GlobalMarkupSet\",\"type\":\"event\"}]",
}

// SynapseExecutionServiceEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseExecutionServiceEventsMetaData.ABI instead.
var SynapseExecutionServiceEventsABI = SynapseExecutionServiceEventsMetaData.ABI

// SynapseExecutionServiceEvents is an auto generated Go binding around an Ethereum contract.
type SynapseExecutionServiceEvents struct {
	SynapseExecutionServiceEventsCaller     // Read-only binding to the contract
	SynapseExecutionServiceEventsTransactor // Write-only binding to the contract
	SynapseExecutionServiceEventsFilterer   // Log filterer for contract events
}

// SynapseExecutionServiceEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseExecutionServiceEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseExecutionServiceEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseExecutionServiceEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseExecutionServiceEventsSession struct {
	Contract     *SynapseExecutionServiceEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SynapseExecutionServiceEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseExecutionServiceEventsCallerSession struct {
	Contract *SynapseExecutionServiceEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// SynapseExecutionServiceEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseExecutionServiceEventsTransactorSession struct {
	Contract     *SynapseExecutionServiceEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// SynapseExecutionServiceEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseExecutionServiceEventsRaw struct {
	Contract *SynapseExecutionServiceEvents // Generic contract binding to access the raw methods on
}

// SynapseExecutionServiceEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseExecutionServiceEventsCallerRaw struct {
	Contract *SynapseExecutionServiceEventsCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseExecutionServiceEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseExecutionServiceEventsTransactorRaw struct {
	Contract *SynapseExecutionServiceEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseExecutionServiceEvents creates a new instance of SynapseExecutionServiceEvents, bound to a specific deployed contract.
func NewSynapseExecutionServiceEvents(address common.Address, backend bind.ContractBackend) (*SynapseExecutionServiceEvents, error) {
	contract, err := bindSynapseExecutionServiceEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceEvents{SynapseExecutionServiceEventsCaller: SynapseExecutionServiceEventsCaller{contract: contract}, SynapseExecutionServiceEventsTransactor: SynapseExecutionServiceEventsTransactor{contract: contract}, SynapseExecutionServiceEventsFilterer: SynapseExecutionServiceEventsFilterer{contract: contract}}, nil
}

// NewSynapseExecutionServiceEventsCaller creates a new read-only instance of SynapseExecutionServiceEvents, bound to a specific deployed contract.
func NewSynapseExecutionServiceEventsCaller(address common.Address, caller bind.ContractCaller) (*SynapseExecutionServiceEventsCaller, error) {
	contract, err := bindSynapseExecutionServiceEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceEventsCaller{contract: contract}, nil
}

// NewSynapseExecutionServiceEventsTransactor creates a new write-only instance of SynapseExecutionServiceEvents, bound to a specific deployed contract.
func NewSynapseExecutionServiceEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseExecutionServiceEventsTransactor, error) {
	contract, err := bindSynapseExecutionServiceEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceEventsTransactor{contract: contract}, nil
}

// NewSynapseExecutionServiceEventsFilterer creates a new log filterer instance of SynapseExecutionServiceEvents, bound to a specific deployed contract.
func NewSynapseExecutionServiceEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseExecutionServiceEventsFilterer, error) {
	contract, err := bindSynapseExecutionServiceEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceEventsFilterer{contract: contract}, nil
}

// bindSynapseExecutionServiceEvents binds a generic wrapper to an already deployed contract.
func bindSynapseExecutionServiceEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseExecutionServiceEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseExecutionServiceEvents.Contract.SynapseExecutionServiceEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseExecutionServiceEvents.Contract.SynapseExecutionServiceEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseExecutionServiceEvents.Contract.SynapseExecutionServiceEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseExecutionServiceEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseExecutionServiceEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseExecutionServiceEvents.Contract.contract.Transact(opts, method, params...)
}

// SynapseExecutionServiceEventsExecutionRequestedIterator is returned from FilterExecutionRequested and is used to iterate over the raw logs and unpacked data for ExecutionRequested events raised by the SynapseExecutionServiceEvents contract.
type SynapseExecutionServiceEventsExecutionRequestedIterator struct {
	Event *SynapseExecutionServiceEventsExecutionRequested // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceEventsExecutionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceEventsExecutionRequested)
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
		it.Event = new(SynapseExecutionServiceEventsExecutionRequested)
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
func (it *SynapseExecutionServiceEventsExecutionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceEventsExecutionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceEventsExecutionRequested represents a ExecutionRequested event raised by the SynapseExecutionServiceEvents contract.
type SynapseExecutionServiceEventsExecutionRequested struct {
	TransactionId [32]byte
	Client        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) FilterExecutionRequested(opts *bind.FilterOpts, transactionId [][32]byte) (*SynapseExecutionServiceEventsExecutionRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SynapseExecutionServiceEvents.contract.FilterLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceEventsExecutionRequestedIterator{contract: _SynapseExecutionServiceEvents.contract, event: "ExecutionRequested", logs: logs, sub: sub}, nil
}

// WatchExecutionRequested is a free log subscription operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) WatchExecutionRequested(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceEventsExecutionRequested, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SynapseExecutionServiceEvents.contract.WatchLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceEventsExecutionRequested)
				if err := _SynapseExecutionServiceEvents.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
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

// ParseExecutionRequested is a log parse operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) ParseExecutionRequested(log types.Log) (*SynapseExecutionServiceEventsExecutionRequested, error) {
	event := new(SynapseExecutionServiceEventsExecutionRequested)
	if err := _SynapseExecutionServiceEvents.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceEventsExecutorEOASetIterator is returned from FilterExecutorEOASet and is used to iterate over the raw logs and unpacked data for ExecutorEOASet events raised by the SynapseExecutionServiceEvents contract.
type SynapseExecutionServiceEventsExecutorEOASetIterator struct {
	Event *SynapseExecutionServiceEventsExecutorEOASet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceEventsExecutorEOASetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceEventsExecutorEOASet)
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
		it.Event = new(SynapseExecutionServiceEventsExecutorEOASet)
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
func (it *SynapseExecutionServiceEventsExecutorEOASetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceEventsExecutorEOASetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceEventsExecutorEOASet represents a ExecutorEOASet event raised by the SynapseExecutionServiceEvents contract.
type SynapseExecutionServiceEventsExecutorEOASet struct {
	ExecutorEOA common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutorEOASet is a free log retrieval operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) FilterExecutorEOASet(opts *bind.FilterOpts) (*SynapseExecutionServiceEventsExecutorEOASetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceEvents.contract.FilterLogs(opts, "ExecutorEOASet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceEventsExecutorEOASetIterator{contract: _SynapseExecutionServiceEvents.contract, event: "ExecutorEOASet", logs: logs, sub: sub}, nil
}

// WatchExecutorEOASet is a free log subscription operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) WatchExecutorEOASet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceEventsExecutorEOASet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceEvents.contract.WatchLogs(opts, "ExecutorEOASet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceEventsExecutorEOASet)
				if err := _SynapseExecutionServiceEvents.contract.UnpackLog(event, "ExecutorEOASet", log); err != nil {
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

// ParseExecutorEOASet is a log parse operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) ParseExecutorEOASet(log types.Log) (*SynapseExecutionServiceEventsExecutorEOASet, error) {
	event := new(SynapseExecutionServiceEventsExecutorEOASet)
	if err := _SynapseExecutionServiceEvents.contract.UnpackLog(event, "ExecutorEOASet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceEventsGasOracleSetIterator is returned from FilterGasOracleSet and is used to iterate over the raw logs and unpacked data for GasOracleSet events raised by the SynapseExecutionServiceEvents contract.
type SynapseExecutionServiceEventsGasOracleSetIterator struct {
	Event *SynapseExecutionServiceEventsGasOracleSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceEventsGasOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceEventsGasOracleSet)
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
		it.Event = new(SynapseExecutionServiceEventsGasOracleSet)
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
func (it *SynapseExecutionServiceEventsGasOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceEventsGasOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceEventsGasOracleSet represents a GasOracleSet event raised by the SynapseExecutionServiceEvents contract.
type SynapseExecutionServiceEventsGasOracleSet struct {
	GasOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOracleSet is a free log retrieval operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) FilterGasOracleSet(opts *bind.FilterOpts) (*SynapseExecutionServiceEventsGasOracleSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceEvents.contract.FilterLogs(opts, "GasOracleSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceEventsGasOracleSetIterator{contract: _SynapseExecutionServiceEvents.contract, event: "GasOracleSet", logs: logs, sub: sub}, nil
}

// WatchGasOracleSet is a free log subscription operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) WatchGasOracleSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceEventsGasOracleSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceEvents.contract.WatchLogs(opts, "GasOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceEventsGasOracleSet)
				if err := _SynapseExecutionServiceEvents.contract.UnpackLog(event, "GasOracleSet", log); err != nil {
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

// ParseGasOracleSet is a log parse operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) ParseGasOracleSet(log types.Log) (*SynapseExecutionServiceEventsGasOracleSet, error) {
	event := new(SynapseExecutionServiceEventsGasOracleSet)
	if err := _SynapseExecutionServiceEvents.contract.UnpackLog(event, "GasOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceEventsGlobalMarkupSetIterator is returned from FilterGlobalMarkupSet and is used to iterate over the raw logs and unpacked data for GlobalMarkupSet events raised by the SynapseExecutionServiceEvents contract.
type SynapseExecutionServiceEventsGlobalMarkupSetIterator struct {
	Event *SynapseExecutionServiceEventsGlobalMarkupSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceEventsGlobalMarkupSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceEventsGlobalMarkupSet)
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
		it.Event = new(SynapseExecutionServiceEventsGlobalMarkupSet)
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
func (it *SynapseExecutionServiceEventsGlobalMarkupSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceEventsGlobalMarkupSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceEventsGlobalMarkupSet represents a GlobalMarkupSet event raised by the SynapseExecutionServiceEvents contract.
type SynapseExecutionServiceEventsGlobalMarkupSet struct {
	GlobalMarkup *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGlobalMarkupSet is a free log retrieval operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) FilterGlobalMarkupSet(opts *bind.FilterOpts) (*SynapseExecutionServiceEventsGlobalMarkupSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceEvents.contract.FilterLogs(opts, "GlobalMarkupSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceEventsGlobalMarkupSetIterator{contract: _SynapseExecutionServiceEvents.contract, event: "GlobalMarkupSet", logs: logs, sub: sub}, nil
}

// WatchGlobalMarkupSet is a free log subscription operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) WatchGlobalMarkupSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceEventsGlobalMarkupSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceEvents.contract.WatchLogs(opts, "GlobalMarkupSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceEventsGlobalMarkupSet)
				if err := _SynapseExecutionServiceEvents.contract.UnpackLog(event, "GlobalMarkupSet", log); err != nil {
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

// ParseGlobalMarkupSet is a log parse operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceEvents *SynapseExecutionServiceEventsFilterer) ParseGlobalMarkupSet(log types.Log) (*SynapseExecutionServiceEventsGlobalMarkupSet, error) {
	event := new(SynapseExecutionServiceEventsGlobalMarkupSet)
	if err := _SynapseExecutionServiceEvents.contract.UnpackLog(event, "GlobalMarkupSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1MetaData contains all meta data concerning the SynapseExecutionServiceV1 contract.
var SynapseExecutionServiceV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"SynapseExecutionService__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__GasOracleNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"SynapseExecutionService__OptionsVersionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__TooShort\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executorEOA\",\"type\":\"address\"}],\"name\":\"ExecutorEOASet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalMarkup\",\"type\":\"uint256\"}],\"name\":\"GlobalMarkupSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IC_CLIENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalMarkup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executorEOA_\",\"type\":\"address\"}],\"name\":\"setExecutorEOA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalMarkup_\",\"type\":\"uint256\"}],\"name\":\"setGlobalMarkup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ccc57490": "GOVERNOR_ROLE()",
		"08c5c0db": "IC_CLIENT_ROLE()",
		"62014bad": "executorEOA()",
		"5d62a8dd": "gasOracle()",
		"96fda4da": "getExecutionFee(uint64,uint256,bytes)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"efd07ec2": "globalMarkup()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"c4d66de8": "initialize(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"592a8799": "requestExecution(uint64,uint256,bytes32,uint256,bytes)",
		"d547741f": "revokeRole(bytes32,address)",
		"2d54566c": "setExecutorEOA(address)",
		"a87b8152": "setGasOracle(address)",
		"cf4f578f": "setGlobalMarkup(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611674806100df6000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c806391d14854116100b2578063c4d66de811610081578063cf4f578f11610066578063cf4f578f14610390578063d547741f146103a3578063efd07ec2146103b657600080fd5b8063c4d66de814610356578063ccc574901461036957600080fd5b806391d14854146102c357806396fda4da14610328578063a217fddf1461033b578063a87b81521461034357600080fd5b80632f2ff15d11610109578063592a8799116100ee578063592a8799146102155780635d62a8dd1461022857806362014bad1461028657600080fd5b80632f2ff15d146101ef57806336568abe1461020257600080fd5b806301ffc9a71461013b57806308c5c0db14610163578063248a9ca3146101985780632d54566c146101da575b600080fd5b61014e61014936600461123c565b6103dd565b60405190151581526020015b60405180910390f35b61018a7f506033f42d439a89b8dbacb157256b8ef7e613d9e48db1be101b85411778abfb81565b60405190815260200161015a565b61018a6101a636600461127e565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6101ed6101e83660046112c0565b610476565b005b6101ed6101fd3660046112db565b610587565b6101ed6102103660046112db565b6105d1565b6101ed610223366004611368565b61062f565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31005473ffffffffffffffffffffffffffffffffffffffff16610261565b61014e6102d13660046112db565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61018a6103363660046113d9565b6106f1565b61018a600081565b6101ed6103513660046112c0565b6109db565b6101ed6103643660046112c0565b610b06565b61018a7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b6101ed61039e36600461127e565b610c89565b6101ed6103b13660046112db565b610d29565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31025461018a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061047057507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f556104a081610d6d565b73ffffffffffffffffffffffffffffffffffffffff82166104ed576040517f0a0c2d7000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e310080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841690811782556040519081527f4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71906020015b60405180910390a1505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546105c181610d6d565b6105cb8383610d7a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610620576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61062a8282610e9b565b505050565b7f506033f42d439a89b8dbacb157256b8ef7e613d9e48db1be101b85411778abfb61065981610d6d565b6000610667888886866106f1565b9050808510156106b2576040517f4641246a00000000000000000000000000000000000000000000000000000000815260048101869052602481018290526044015b60405180910390fd5b60405133815286907f507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd13829060200160405180910390a25050505050505050565b6000806107327fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31015473ffffffffffffffffffffffffffffffffffffffff1690565b905073ffffffffffffffffffffffffffffffffffffffff8116610781576040517f38d2192500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061078d8585610f79565b9050600161ffff821611156107d4576040517f05e98f3a00000000000000000000000000000000000000000000000000000000815261ffff821660048201526024016106a9565b600061081586868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fc392505050565b80516040517fbf495c8800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b16600482015260248101919091526044810189905290915073ffffffffffffffffffffffffffffffffffffffff84169063bf495c8890606401602060405180830381865afa15801561089d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c19190611433565b6020820151909450156109805760208101516040517f40658a7400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8a166004820152602481019190915273ffffffffffffffffffffffffffffffffffffffff8416906340658a7490604401602060405180830381865afa15801561094f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109739190611433565b61097d908561147b565b93505b670de0b6b3a76400006109b17fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31025490565b6109bb908661148e565b6109c591906114a5565b6109cf908561147b565b98975050505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610a0581610d6d565b73ffffffffffffffffffffffffffffffffffffffff8216610a52576040517f0a0c2d7000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e310180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040519081527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa9060200161057a565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610b515750825b905060008267ffffffffffffffff166001148015610b6e5750303b155b905081158015610b7c575080155b15610bb3576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610c145784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610c1f600087610d7a565b508315610c815784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610cb381610d6d565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31028290556040518281527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d9060200161057a565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d6381610d6d565b6105cb8383610e9b565b610d77813361104c565b50565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16610e915760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610e2d3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610470565b6000915050610470565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615610e915760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610470565b60006002821015610fba5782826040517f659cf9fa0000000000000000000000000000000000000000000000000000000081526004016106a99291906114e0565b50503560f01c90565b60408051808201909152600080825260208201526000610fe2836110f7565b9050600161ffff82161015611029576040517fb94fa72500000000000000000000000000000000000000000000000000000000815261ffff821660048201526024016106a9565b61103283611142565b806020019051810190611045919061155c565b9392505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166110f3576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044016106a9565b5050565b600060028251101561113757816040517f659cf9fa0000000000000000000000000000000000000000000000000000000081526004016106a991906115d2565b506020015160f01c90565b606060028251101561118257816040517f659cf9fa0000000000000000000000000000000000000000000000000000000081526004016106a991906115d2565b81517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe018067ffffffffffffffff8111156111bf576111bf61152d565b6040519080825280601f01601f1916602001820160405280156111e9576020820181803683370190505b50915060008160208401836022870160045afa905080611235576040517f101e44fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050919050565b60006020828403121561124e57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461104557600080fd5b60006020828403121561129057600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146112bb57600080fd5b919050565b6000602082840312156112d257600080fd5b61104582611297565b600080604083850312156112ee57600080fd5b823591506112fe60208401611297565b90509250929050565b803567ffffffffffffffff811681146112bb57600080fd5b60008083601f84011261133157600080fd5b50813567ffffffffffffffff81111561134957600080fd5b60208301915083602082850101111561136157600080fd5b9250929050565b60008060008060008060a0878903121561138157600080fd5b61138a87611307565b9550602087013594506040870135935060608701359250608087013567ffffffffffffffff8111156113bb57600080fd5b6113c789828a0161131f565b979a9699509497509295939492505050565b600080600080606085870312156113ef57600080fd5b6113f885611307565b935060208501359250604085013567ffffffffffffffff81111561141b57600080fd5b6114278782880161131f565b95989497509550505050565b60006020828403121561144557600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104705761047061144c565b80820281158282048414176104705761047061144c565b6000826114db577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006040828403121561156e57600080fd5b6040516040810181811067ffffffffffffffff821117156115b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052825181526020928301519281019290925250919050565b600060208083528351808285015260005b818110156115ff578581018301518582016040015282016115e3565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509291505056fea264697066735822122014817ae42636063453c8c6be0662cbd3f8de3afa1038f76b995f39e5ee6518fa64736f6c63430008140033",
}

// SynapseExecutionServiceV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseExecutionServiceV1MetaData.ABI instead.
var SynapseExecutionServiceV1ABI = SynapseExecutionServiceV1MetaData.ABI

// Deprecated: Use SynapseExecutionServiceV1MetaData.Sigs instead.
// SynapseExecutionServiceV1FuncSigs maps the 4-byte function signature to its string representation.
var SynapseExecutionServiceV1FuncSigs = SynapseExecutionServiceV1MetaData.Sigs

// SynapseExecutionServiceV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseExecutionServiceV1MetaData.Bin instead.
var SynapseExecutionServiceV1Bin = SynapseExecutionServiceV1MetaData.Bin

// DeploySynapseExecutionServiceV1 deploys a new Ethereum contract, binding an instance of SynapseExecutionServiceV1 to it.
func DeploySynapseExecutionServiceV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseExecutionServiceV1, error) {
	parsed, err := SynapseExecutionServiceV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseExecutionServiceV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseExecutionServiceV1{SynapseExecutionServiceV1Caller: SynapseExecutionServiceV1Caller{contract: contract}, SynapseExecutionServiceV1Transactor: SynapseExecutionServiceV1Transactor{contract: contract}, SynapseExecutionServiceV1Filterer: SynapseExecutionServiceV1Filterer{contract: contract}}, nil
}

// SynapseExecutionServiceV1 is an auto generated Go binding around an Ethereum contract.
type SynapseExecutionServiceV1 struct {
	SynapseExecutionServiceV1Caller     // Read-only binding to the contract
	SynapseExecutionServiceV1Transactor // Write-only binding to the contract
	SynapseExecutionServiceV1Filterer   // Log filterer for contract events
}

// SynapseExecutionServiceV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseExecutionServiceV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseExecutionServiceV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseExecutionServiceV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseExecutionServiceV1Session struct {
	Contract     *SynapseExecutionServiceV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SynapseExecutionServiceV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseExecutionServiceV1CallerSession struct {
	Contract *SynapseExecutionServiceV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// SynapseExecutionServiceV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseExecutionServiceV1TransactorSession struct {
	Contract     *SynapseExecutionServiceV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// SynapseExecutionServiceV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseExecutionServiceV1Raw struct {
	Contract *SynapseExecutionServiceV1 // Generic contract binding to access the raw methods on
}

// SynapseExecutionServiceV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseExecutionServiceV1CallerRaw struct {
	Contract *SynapseExecutionServiceV1Caller // Generic read-only contract binding to access the raw methods on
}

// SynapseExecutionServiceV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseExecutionServiceV1TransactorRaw struct {
	Contract *SynapseExecutionServiceV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseExecutionServiceV1 creates a new instance of SynapseExecutionServiceV1, bound to a specific deployed contract.
func NewSynapseExecutionServiceV1(address common.Address, backend bind.ContractBackend) (*SynapseExecutionServiceV1, error) {
	contract, err := bindSynapseExecutionServiceV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1{SynapseExecutionServiceV1Caller: SynapseExecutionServiceV1Caller{contract: contract}, SynapseExecutionServiceV1Transactor: SynapseExecutionServiceV1Transactor{contract: contract}, SynapseExecutionServiceV1Filterer: SynapseExecutionServiceV1Filterer{contract: contract}}, nil
}

// NewSynapseExecutionServiceV1Caller creates a new read-only instance of SynapseExecutionServiceV1, bound to a specific deployed contract.
func NewSynapseExecutionServiceV1Caller(address common.Address, caller bind.ContractCaller) (*SynapseExecutionServiceV1Caller, error) {
	contract, err := bindSynapseExecutionServiceV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1Caller{contract: contract}, nil
}

// NewSynapseExecutionServiceV1Transactor creates a new write-only instance of SynapseExecutionServiceV1, bound to a specific deployed contract.
func NewSynapseExecutionServiceV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SynapseExecutionServiceV1Transactor, error) {
	contract, err := bindSynapseExecutionServiceV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1Transactor{contract: contract}, nil
}

// NewSynapseExecutionServiceV1Filterer creates a new log filterer instance of SynapseExecutionServiceV1, bound to a specific deployed contract.
func NewSynapseExecutionServiceV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SynapseExecutionServiceV1Filterer, error) {
	contract, err := bindSynapseExecutionServiceV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1Filterer{contract: contract}, nil
}

// bindSynapseExecutionServiceV1 binds a generic wrapper to an already deployed contract.
func bindSynapseExecutionServiceV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseExecutionServiceV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseExecutionServiceV1.Contract.SynapseExecutionServiceV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SynapseExecutionServiceV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SynapseExecutionServiceV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseExecutionServiceV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1.Contract.DEFAULTADMINROLE(&_SynapseExecutionServiceV1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1.Contract.DEFAULTADMINROLE(&_SynapseExecutionServiceV1.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GOVERNORROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1.Contract.GOVERNORROLE(&_SynapseExecutionServiceV1.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GOVERNORROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1.Contract.GOVERNORROLE(&_SynapseExecutionServiceV1.CallOpts)
}

// ICCLIENTROLE is a free data retrieval call binding the contract method 0x08c5c0db.
//
// Solidity: function IC_CLIENT_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) ICCLIENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "IC_CLIENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ICCLIENTROLE is a free data retrieval call binding the contract method 0x08c5c0db.
//
// Solidity: function IC_CLIENT_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) ICCLIENTROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1.Contract.ICCLIENTROLE(&_SynapseExecutionServiceV1.CallOpts)
}

// ICCLIENTROLE is a free data retrieval call binding the contract method 0x08c5c0db.
//
// Solidity: function IC_CLIENT_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) ICCLIENTROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1.Contract.ICCLIENTROLE(&_SynapseExecutionServiceV1.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) ExecutorEOA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "executorEOA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) ExecutorEOA() (common.Address, error) {
	return _SynapseExecutionServiceV1.Contract.ExecutorEOA(&_SynapseExecutionServiceV1.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) ExecutorEOA() (common.Address, error) {
	return _SynapseExecutionServiceV1.Contract.ExecutorEOA(&_SynapseExecutionServiceV1.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GasOracle() (common.Address, error) {
	return _SynapseExecutionServiceV1.Contract.GasOracle(&_SynapseExecutionServiceV1.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GasOracle() (common.Address, error) {
	return _SynapseExecutionServiceV1.Contract.GasOracle(&_SynapseExecutionServiceV1.CallOpts)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256 executionFee)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GetExecutionFee(opts *bind.CallOpts, dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "getExecutionFee", dstChainId, txPayloadSize, options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256 executionFee)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GetExecutionFee(dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GetExecutionFee(&_SynapseExecutionServiceV1.CallOpts, dstChainId, txPayloadSize, options)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256 executionFee)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GetExecutionFee(dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GetExecutionFee(&_SynapseExecutionServiceV1.CallOpts, dstChainId, txPayloadSize, options)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SynapseExecutionServiceV1.Contract.GetRoleAdmin(&_SynapseExecutionServiceV1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SynapseExecutionServiceV1.Contract.GetRoleAdmin(&_SynapseExecutionServiceV1.CallOpts, role)
}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GlobalMarkup(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "globalMarkup")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GlobalMarkup() (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GlobalMarkup(&_SynapseExecutionServiceV1.CallOpts)
}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GlobalMarkup() (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GlobalMarkup(&_SynapseExecutionServiceV1.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SynapseExecutionServiceV1.Contract.HasRole(&_SynapseExecutionServiceV1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SynapseExecutionServiceV1.Contract.HasRole(&_SynapseExecutionServiceV1.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SynapseExecutionServiceV1.Contract.SupportsInterface(&_SynapseExecutionServiceV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SynapseExecutionServiceV1.Contract.SupportsInterface(&_SynapseExecutionServiceV1.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.GrantRole(&_SynapseExecutionServiceV1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.GrantRole(&_SynapseExecutionServiceV1.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) Initialize(admin common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.Initialize(&_SynapseExecutionServiceV1.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.Initialize(&_SynapseExecutionServiceV1.TransactOpts, admin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.RenounceRole(&_SynapseExecutionServiceV1.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.RenounceRole(&_SynapseExecutionServiceV1.TransactOpts, role, callerConfirmation)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) RequestExecution(opts *bind.TransactOpts, dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "requestExecution", dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) RequestExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.RequestExecution(&_SynapseExecutionServiceV1.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) RequestExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.RequestExecution(&_SynapseExecutionServiceV1.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.RevokeRole(&_SynapseExecutionServiceV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.RevokeRole(&_SynapseExecutionServiceV1.TransactOpts, role, account)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) SetExecutorEOA(opts *bind.TransactOpts, executorEOA_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "setExecutorEOA", executorEOA_)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) SetExecutorEOA(executorEOA_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SetExecutorEOA(&_SynapseExecutionServiceV1.TransactOpts, executorEOA_)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) SetExecutorEOA(executorEOA_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SetExecutorEOA(&_SynapseExecutionServiceV1.TransactOpts, executorEOA_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) SetGasOracle(opts *bind.TransactOpts, gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "setGasOracle", gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SetGasOracle(&_SynapseExecutionServiceV1.TransactOpts, gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SetGasOracle(&_SynapseExecutionServiceV1.TransactOpts, gasOracle_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) SetGlobalMarkup(opts *bind.TransactOpts, globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "setGlobalMarkup", globalMarkup_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) SetGlobalMarkup(globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SetGlobalMarkup(&_SynapseExecutionServiceV1.TransactOpts, globalMarkup_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) SetGlobalMarkup(globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SetGlobalMarkup(&_SynapseExecutionServiceV1.TransactOpts, globalMarkup_)
}

// SynapseExecutionServiceV1ExecutionRequestedIterator is returned from FilterExecutionRequested and is used to iterate over the raw logs and unpacked data for ExecutionRequested events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1ExecutionRequestedIterator struct {
	Event *SynapseExecutionServiceV1ExecutionRequested // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1ExecutionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1ExecutionRequested)
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
		it.Event = new(SynapseExecutionServiceV1ExecutionRequested)
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
func (it *SynapseExecutionServiceV1ExecutionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1ExecutionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1ExecutionRequested represents a ExecutionRequested event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1ExecutionRequested struct {
	TransactionId [32]byte
	Client        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterExecutionRequested(opts *bind.FilterOpts, transactionId [][32]byte) (*SynapseExecutionServiceV1ExecutionRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1ExecutionRequestedIterator{contract: _SynapseExecutionServiceV1.contract, event: "ExecutionRequested", logs: logs, sub: sub}, nil
}

// WatchExecutionRequested is a free log subscription operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchExecutionRequested(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1ExecutionRequested, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1ExecutionRequested)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
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

// ParseExecutionRequested is a log parse operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseExecutionRequested(log types.Log) (*SynapseExecutionServiceV1ExecutionRequested, error) {
	event := new(SynapseExecutionServiceV1ExecutionRequested)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1ExecutorEOASetIterator is returned from FilterExecutorEOASet and is used to iterate over the raw logs and unpacked data for ExecutorEOASet events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1ExecutorEOASetIterator struct {
	Event *SynapseExecutionServiceV1ExecutorEOASet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1ExecutorEOASetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1ExecutorEOASet)
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
		it.Event = new(SynapseExecutionServiceV1ExecutorEOASet)
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
func (it *SynapseExecutionServiceV1ExecutorEOASetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1ExecutorEOASetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1ExecutorEOASet represents a ExecutorEOASet event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1ExecutorEOASet struct {
	ExecutorEOA common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutorEOASet is a free log retrieval operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterExecutorEOASet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1ExecutorEOASetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "ExecutorEOASet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1ExecutorEOASetIterator{contract: _SynapseExecutionServiceV1.contract, event: "ExecutorEOASet", logs: logs, sub: sub}, nil
}

// WatchExecutorEOASet is a free log subscription operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchExecutorEOASet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1ExecutorEOASet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "ExecutorEOASet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1ExecutorEOASet)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "ExecutorEOASet", log); err != nil {
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

// ParseExecutorEOASet is a log parse operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseExecutorEOASet(log types.Log) (*SynapseExecutionServiceV1ExecutorEOASet, error) {
	event := new(SynapseExecutionServiceV1ExecutorEOASet)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "ExecutorEOASet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1GasOracleSetIterator is returned from FilterGasOracleSet and is used to iterate over the raw logs and unpacked data for GasOracleSet events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1GasOracleSetIterator struct {
	Event *SynapseExecutionServiceV1GasOracleSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1GasOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1GasOracleSet)
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
		it.Event = new(SynapseExecutionServiceV1GasOracleSet)
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
func (it *SynapseExecutionServiceV1GasOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1GasOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1GasOracleSet represents a GasOracleSet event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1GasOracleSet struct {
	GasOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOracleSet is a free log retrieval operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterGasOracleSet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1GasOracleSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "GasOracleSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1GasOracleSetIterator{contract: _SynapseExecutionServiceV1.contract, event: "GasOracleSet", logs: logs, sub: sub}, nil
}

// WatchGasOracleSet is a free log subscription operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchGasOracleSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1GasOracleSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "GasOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1GasOracleSet)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "GasOracleSet", log); err != nil {
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

// ParseGasOracleSet is a log parse operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseGasOracleSet(log types.Log) (*SynapseExecutionServiceV1GasOracleSet, error) {
	event := new(SynapseExecutionServiceV1GasOracleSet)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "GasOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1GlobalMarkupSetIterator is returned from FilterGlobalMarkupSet and is used to iterate over the raw logs and unpacked data for GlobalMarkupSet events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1GlobalMarkupSetIterator struct {
	Event *SynapseExecutionServiceV1GlobalMarkupSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1GlobalMarkupSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1GlobalMarkupSet)
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
		it.Event = new(SynapseExecutionServiceV1GlobalMarkupSet)
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
func (it *SynapseExecutionServiceV1GlobalMarkupSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1GlobalMarkupSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1GlobalMarkupSet represents a GlobalMarkupSet event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1GlobalMarkupSet struct {
	GlobalMarkup *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGlobalMarkupSet is a free log retrieval operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterGlobalMarkupSet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1GlobalMarkupSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "GlobalMarkupSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1GlobalMarkupSetIterator{contract: _SynapseExecutionServiceV1.contract, event: "GlobalMarkupSet", logs: logs, sub: sub}, nil
}

// WatchGlobalMarkupSet is a free log subscription operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchGlobalMarkupSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1GlobalMarkupSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "GlobalMarkupSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1GlobalMarkupSet)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "GlobalMarkupSet", log); err != nil {
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

// ParseGlobalMarkupSet is a log parse operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseGlobalMarkupSet(log types.Log) (*SynapseExecutionServiceV1GlobalMarkupSet, error) {
	event := new(SynapseExecutionServiceV1GlobalMarkupSet)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "GlobalMarkupSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1InitializedIterator struct {
	Event *SynapseExecutionServiceV1Initialized // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1Initialized)
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
		it.Event = new(SynapseExecutionServiceV1Initialized)
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
func (it *SynapseExecutionServiceV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1Initialized represents a Initialized event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*SynapseExecutionServiceV1InitializedIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1InitializedIterator{contract: _SynapseExecutionServiceV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1Initialized) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1Initialized)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseInitialized(log types.Log) (*SynapseExecutionServiceV1Initialized, error) {
	event := new(SynapseExecutionServiceV1Initialized)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1RoleAdminChangedIterator struct {
	Event *SynapseExecutionServiceV1RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1RoleAdminChanged)
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
		it.Event = new(SynapseExecutionServiceV1RoleAdminChanged)
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
func (it *SynapseExecutionServiceV1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1RoleAdminChanged represents a RoleAdminChanged event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SynapseExecutionServiceV1RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1RoleAdminChangedIterator{contract: _SynapseExecutionServiceV1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1RoleAdminChanged)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseRoleAdminChanged(log types.Log) (*SynapseExecutionServiceV1RoleAdminChanged, error) {
	event := new(SynapseExecutionServiceV1RoleAdminChanged)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1RoleGrantedIterator struct {
	Event *SynapseExecutionServiceV1RoleGranted // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1RoleGranted)
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
		it.Event = new(SynapseExecutionServiceV1RoleGranted)
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
func (it *SynapseExecutionServiceV1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1RoleGranted represents a RoleGranted event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SynapseExecutionServiceV1RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1RoleGrantedIterator{contract: _SynapseExecutionServiceV1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1RoleGranted)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseRoleGranted(log types.Log) (*SynapseExecutionServiceV1RoleGranted, error) {
	event := new(SynapseExecutionServiceV1RoleGranted)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1RoleRevokedIterator struct {
	Event *SynapseExecutionServiceV1RoleRevoked // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1RoleRevoked)
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
		it.Event = new(SynapseExecutionServiceV1RoleRevoked)
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
func (it *SynapseExecutionServiceV1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1RoleRevoked represents a RoleRevoked event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SynapseExecutionServiceV1RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1RoleRevokedIterator{contract: _SynapseExecutionServiceV1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1RoleRevoked)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseRoleRevoked(log types.Log) (*SynapseExecutionServiceV1RoleRevoked, error) {
	event := new(SynapseExecutionServiceV1RoleRevoked)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessMetaData contains all meta data concerning the SynapseExecutionServiceV1Harness contract.
var SynapseExecutionServiceV1HarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"SynapseExecutionService__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__GasOracleNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"SynapseExecutionService__OptionsVersionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__TooShort\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executorEOA\",\"type\":\"address\"}],\"name\":\"ExecutorEOASet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalMarkup\",\"type\":\"uint256\"}],\"name\":\"GlobalMarkupSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IC_CLIENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalMarkup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executorEOA_\",\"type\":\"address\"}],\"name\":\"setExecutorEOA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalMarkup_\",\"type\":\"uint256\"}],\"name\":\"setGlobalMarkup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ccc57490": "GOVERNOR_ROLE()",
		"08c5c0db": "IC_CLIENT_ROLE()",
		"62014bad": "executorEOA()",
		"5d62a8dd": "gasOracle()",
		"96fda4da": "getExecutionFee(uint64,uint256,bytes)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"efd07ec2": "globalMarkup()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"c4d66de8": "initialize(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"592a8799": "requestExecution(uint64,uint256,bytes32,uint256,bytes)",
		"d547741f": "revokeRole(bytes32,address)",
		"2d54566c": "setExecutorEOA(address)",
		"a87b8152": "setGasOracle(address)",
		"cf4f578f": "setGlobalMarkup(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061001c60003361004d565b506100477f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f553361004d565b5061011e565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16610112576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556100c83390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610118565b60009150505b92915050565b6116748061012d6000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c806391d14854116100b2578063c4d66de811610081578063cf4f578f11610066578063cf4f578f14610390578063d547741f146103a3578063efd07ec2146103b657600080fd5b8063c4d66de814610356578063ccc574901461036957600080fd5b806391d14854146102c357806396fda4da14610328578063a217fddf1461033b578063a87b81521461034357600080fd5b80632f2ff15d11610109578063592a8799116100ee578063592a8799146102155780635d62a8dd1461022857806362014bad1461028657600080fd5b80632f2ff15d146101ef57806336568abe1461020257600080fd5b806301ffc9a71461013b57806308c5c0db14610163578063248a9ca3146101985780632d54566c146101da575b600080fd5b61014e61014936600461123c565b6103dd565b60405190151581526020015b60405180910390f35b61018a7f506033f42d439a89b8dbacb157256b8ef7e613d9e48db1be101b85411778abfb81565b60405190815260200161015a565b61018a6101a636600461127e565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6101ed6101e83660046112c0565b610476565b005b6101ed6101fd3660046112db565b610587565b6101ed6102103660046112db565b6105d1565b6101ed610223366004611368565b61062f565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31005473ffffffffffffffffffffffffffffffffffffffff16610261565b61014e6102d13660046112db565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61018a6103363660046113d9565b6106f1565b61018a600081565b6101ed6103513660046112c0565b6109db565b6101ed6103643660046112c0565b610b06565b61018a7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b6101ed61039e36600461127e565b610c89565b6101ed6103b13660046112db565b610d29565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31025461018a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061047057507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f556104a081610d6d565b73ffffffffffffffffffffffffffffffffffffffff82166104ed576040517f0a0c2d7000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e310080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841690811782556040519081527f4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71906020015b60405180910390a1505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546105c181610d6d565b6105cb8383610d7a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610620576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61062a8282610e9b565b505050565b7f506033f42d439a89b8dbacb157256b8ef7e613d9e48db1be101b85411778abfb61065981610d6d565b6000610667888886866106f1565b9050808510156106b2576040517f4641246a00000000000000000000000000000000000000000000000000000000815260048101869052602481018290526044015b60405180910390fd5b60405133815286907f507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd13829060200160405180910390a25050505050505050565b6000806107327fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31015473ffffffffffffffffffffffffffffffffffffffff1690565b905073ffffffffffffffffffffffffffffffffffffffff8116610781576040517f38d2192500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061078d8585610f79565b9050600161ffff821611156107d4576040517f05e98f3a00000000000000000000000000000000000000000000000000000000815261ffff821660048201526024016106a9565b600061081586868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fc392505050565b80516040517fbf495c8800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b16600482015260248101919091526044810189905290915073ffffffffffffffffffffffffffffffffffffffff84169063bf495c8890606401602060405180830381865afa15801561089d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c19190611433565b6020820151909450156109805760208101516040517f40658a7400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8a166004820152602481019190915273ffffffffffffffffffffffffffffffffffffffff8416906340658a7490604401602060405180830381865afa15801561094f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109739190611433565b61097d908561147b565b93505b670de0b6b3a76400006109b17fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31025490565b6109bb908661148e565b6109c591906114a5565b6109cf908561147b565b98975050505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610a0581610d6d565b73ffffffffffffffffffffffffffffffffffffffff8216610a52576040517f0a0c2d7000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e310180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040519081527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa9060200161057a565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610b515750825b905060008267ffffffffffffffff166001148015610b6e5750303b155b905081158015610b7c575080155b15610bb3576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610c145784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610c1f600087610d7a565b508315610c815784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610cb381610d6d565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31028290556040518281527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d9060200161057a565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d6381610d6d565b6105cb8383610e9b565b610d77813361104c565b50565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16610e915760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610e2d3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610470565b6000915050610470565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615610e915760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610470565b60006002821015610fba5782826040517f659cf9fa0000000000000000000000000000000000000000000000000000000081526004016106a99291906114e0565b50503560f01c90565b60408051808201909152600080825260208201526000610fe2836110f7565b9050600161ffff82161015611029576040517fb94fa72500000000000000000000000000000000000000000000000000000000815261ffff821660048201526024016106a9565b61103283611142565b806020019051810190611045919061155c565b9392505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166110f3576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044016106a9565b5050565b600060028251101561113757816040517f659cf9fa0000000000000000000000000000000000000000000000000000000081526004016106a991906115d2565b506020015160f01c90565b606060028251101561118257816040517f659cf9fa0000000000000000000000000000000000000000000000000000000081526004016106a991906115d2565b81517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe018067ffffffffffffffff8111156111bf576111bf61152d565b6040519080825280601f01601f1916602001820160405280156111e9576020820181803683370190505b50915060008160208401836022870160045afa905080611235576040517f101e44fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050919050565b60006020828403121561124e57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461104557600080fd5b60006020828403121561129057600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146112bb57600080fd5b919050565b6000602082840312156112d257600080fd5b61104582611297565b600080604083850312156112ee57600080fd5b823591506112fe60208401611297565b90509250929050565b803567ffffffffffffffff811681146112bb57600080fd5b60008083601f84011261133157600080fd5b50813567ffffffffffffffff81111561134957600080fd5b60208301915083602082850101111561136157600080fd5b9250929050565b60008060008060008060a0878903121561138157600080fd5b61138a87611307565b9550602087013594506040870135935060608701359250608087013567ffffffffffffffff8111156113bb57600080fd5b6113c789828a0161131f565b979a9699509497509295939492505050565b600080600080606085870312156113ef57600080fd5b6113f885611307565b935060208501359250604085013567ffffffffffffffff81111561141b57600080fd5b6114278782880161131f565b95989497509550505050565b60006020828403121561144557600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104705761047061144c565b80820281158282048414176104705761047061144c565b6000826114db577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006040828403121561156e57600080fd5b6040516040810181811067ffffffffffffffff821117156115b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052825181526020928301519281019290925250919050565b600060208083528351808285015260005b818110156115ff578581018301518582016040015282016115e3565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509291505056fea26469706673582212203fab957a73a91623176ea6c7ee0802610f9f073e009fe1160639cbbc842a684264736f6c63430008140033",
}

// SynapseExecutionServiceV1HarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseExecutionServiceV1HarnessMetaData.ABI instead.
var SynapseExecutionServiceV1HarnessABI = SynapseExecutionServiceV1HarnessMetaData.ABI

// Deprecated: Use SynapseExecutionServiceV1HarnessMetaData.Sigs instead.
// SynapseExecutionServiceV1HarnessFuncSigs maps the 4-byte function signature to its string representation.
var SynapseExecutionServiceV1HarnessFuncSigs = SynapseExecutionServiceV1HarnessMetaData.Sigs

// SynapseExecutionServiceV1HarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseExecutionServiceV1HarnessMetaData.Bin instead.
var SynapseExecutionServiceV1HarnessBin = SynapseExecutionServiceV1HarnessMetaData.Bin

// DeploySynapseExecutionServiceV1Harness deploys a new Ethereum contract, binding an instance of SynapseExecutionServiceV1Harness to it.
func DeploySynapseExecutionServiceV1Harness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseExecutionServiceV1Harness, error) {
	parsed, err := SynapseExecutionServiceV1HarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseExecutionServiceV1HarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseExecutionServiceV1Harness{SynapseExecutionServiceV1HarnessCaller: SynapseExecutionServiceV1HarnessCaller{contract: contract}, SynapseExecutionServiceV1HarnessTransactor: SynapseExecutionServiceV1HarnessTransactor{contract: contract}, SynapseExecutionServiceV1HarnessFilterer: SynapseExecutionServiceV1HarnessFilterer{contract: contract}}, nil
}

// SynapseExecutionServiceV1Harness is an auto generated Go binding around an Ethereum contract.
type SynapseExecutionServiceV1Harness struct {
	SynapseExecutionServiceV1HarnessCaller     // Read-only binding to the contract
	SynapseExecutionServiceV1HarnessTransactor // Write-only binding to the contract
	SynapseExecutionServiceV1HarnessFilterer   // Log filterer for contract events
}

// SynapseExecutionServiceV1HarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseExecutionServiceV1HarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceV1HarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseExecutionServiceV1HarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceV1HarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseExecutionServiceV1HarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseExecutionServiceV1HarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseExecutionServiceV1HarnessSession struct {
	Contract     *SynapseExecutionServiceV1Harness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// SynapseExecutionServiceV1HarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseExecutionServiceV1HarnessCallerSession struct {
	Contract *SynapseExecutionServiceV1HarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// SynapseExecutionServiceV1HarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseExecutionServiceV1HarnessTransactorSession struct {
	Contract     *SynapseExecutionServiceV1HarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// SynapseExecutionServiceV1HarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseExecutionServiceV1HarnessRaw struct {
	Contract *SynapseExecutionServiceV1Harness // Generic contract binding to access the raw methods on
}

// SynapseExecutionServiceV1HarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseExecutionServiceV1HarnessCallerRaw struct {
	Contract *SynapseExecutionServiceV1HarnessCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseExecutionServiceV1HarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseExecutionServiceV1HarnessTransactorRaw struct {
	Contract *SynapseExecutionServiceV1HarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseExecutionServiceV1Harness creates a new instance of SynapseExecutionServiceV1Harness, bound to a specific deployed contract.
func NewSynapseExecutionServiceV1Harness(address common.Address, backend bind.ContractBackend) (*SynapseExecutionServiceV1Harness, error) {
	contract, err := bindSynapseExecutionServiceV1Harness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1Harness{SynapseExecutionServiceV1HarnessCaller: SynapseExecutionServiceV1HarnessCaller{contract: contract}, SynapseExecutionServiceV1HarnessTransactor: SynapseExecutionServiceV1HarnessTransactor{contract: contract}, SynapseExecutionServiceV1HarnessFilterer: SynapseExecutionServiceV1HarnessFilterer{contract: contract}}, nil
}

// NewSynapseExecutionServiceV1HarnessCaller creates a new read-only instance of SynapseExecutionServiceV1Harness, bound to a specific deployed contract.
func NewSynapseExecutionServiceV1HarnessCaller(address common.Address, caller bind.ContractCaller) (*SynapseExecutionServiceV1HarnessCaller, error) {
	contract, err := bindSynapseExecutionServiceV1Harness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessCaller{contract: contract}, nil
}

// NewSynapseExecutionServiceV1HarnessTransactor creates a new write-only instance of SynapseExecutionServiceV1Harness, bound to a specific deployed contract.
func NewSynapseExecutionServiceV1HarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseExecutionServiceV1HarnessTransactor, error) {
	contract, err := bindSynapseExecutionServiceV1Harness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessTransactor{contract: contract}, nil
}

// NewSynapseExecutionServiceV1HarnessFilterer creates a new log filterer instance of SynapseExecutionServiceV1Harness, bound to a specific deployed contract.
func NewSynapseExecutionServiceV1HarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseExecutionServiceV1HarnessFilterer, error) {
	contract, err := bindSynapseExecutionServiceV1Harness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessFilterer{contract: contract}, nil
}

// bindSynapseExecutionServiceV1Harness binds a generic wrapper to an already deployed contract.
func bindSynapseExecutionServiceV1Harness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseExecutionServiceV1HarnessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseExecutionServiceV1Harness.Contract.SynapseExecutionServiceV1HarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SynapseExecutionServiceV1HarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SynapseExecutionServiceV1HarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseExecutionServiceV1Harness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1Harness.Contract.DEFAULTADMINROLE(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1Harness.Contract.DEFAULTADMINROLE(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GOVERNORROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GOVERNORROLE(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GOVERNORROLE(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// ICCLIENTROLE is a free data retrieval call binding the contract method 0x08c5c0db.
//
// Solidity: function IC_CLIENT_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) ICCLIENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "IC_CLIENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ICCLIENTROLE is a free data retrieval call binding the contract method 0x08c5c0db.
//
// Solidity: function IC_CLIENT_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) ICCLIENTROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1Harness.Contract.ICCLIENTROLE(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// ICCLIENTROLE is a free data retrieval call binding the contract method 0x08c5c0db.
//
// Solidity: function IC_CLIENT_ROLE() view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) ICCLIENTROLE() ([32]byte, error) {
	return _SynapseExecutionServiceV1Harness.Contract.ICCLIENTROLE(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) ExecutorEOA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "executorEOA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) ExecutorEOA() (common.Address, error) {
	return _SynapseExecutionServiceV1Harness.Contract.ExecutorEOA(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) ExecutorEOA() (common.Address, error) {
	return _SynapseExecutionServiceV1Harness.Contract.ExecutorEOA(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GasOracle() (common.Address, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GasOracle(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GasOracle() (common.Address, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GasOracle(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256 executionFee)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GetExecutionFee(opts *bind.CallOpts, dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "getExecutionFee", dstChainId, txPayloadSize, options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256 executionFee)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GetExecutionFee(dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetExecutionFee(&_SynapseExecutionServiceV1Harness.CallOpts, dstChainId, txPayloadSize, options)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0x96fda4da.
//
// Solidity: function getExecutionFee(uint64 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256 executionFee)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GetExecutionFee(dstChainId uint64, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetExecutionFee(&_SynapseExecutionServiceV1Harness.CallOpts, dstChainId, txPayloadSize, options)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetRoleAdmin(&_SynapseExecutionServiceV1Harness.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetRoleAdmin(&_SynapseExecutionServiceV1Harness.CallOpts, role)
}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GlobalMarkup(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "globalMarkup")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GlobalMarkup() (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GlobalMarkup(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GlobalMarkup is a free data retrieval call binding the contract method 0xefd07ec2.
//
// Solidity: function globalMarkup() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GlobalMarkup() (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GlobalMarkup(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SynapseExecutionServiceV1Harness.Contract.HasRole(&_SynapseExecutionServiceV1Harness.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SynapseExecutionServiceV1Harness.Contract.HasRole(&_SynapseExecutionServiceV1Harness.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SupportsInterface(&_SynapseExecutionServiceV1Harness.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SupportsInterface(&_SynapseExecutionServiceV1Harness.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GrantRole(&_SynapseExecutionServiceV1Harness.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GrantRole(&_SynapseExecutionServiceV1Harness.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.Initialize(&_SynapseExecutionServiceV1Harness.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.Initialize(&_SynapseExecutionServiceV1Harness.TransactOpts, admin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.RenounceRole(&_SynapseExecutionServiceV1Harness.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.RenounceRole(&_SynapseExecutionServiceV1Harness.TransactOpts, role, callerConfirmation)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) RequestExecution(opts *bind.TransactOpts, dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "requestExecution", dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) RequestExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.RequestExecution(&_SynapseExecutionServiceV1Harness.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0x592a8799.
//
// Solidity: function requestExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) RequestExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.RequestExecution(&_SynapseExecutionServiceV1Harness.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.RevokeRole(&_SynapseExecutionServiceV1Harness.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.RevokeRole(&_SynapseExecutionServiceV1Harness.TransactOpts, role, account)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) SetExecutorEOA(opts *bind.TransactOpts, executorEOA_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "setExecutorEOA", executorEOA_)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) SetExecutorEOA(executorEOA_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SetExecutorEOA(&_SynapseExecutionServiceV1Harness.TransactOpts, executorEOA_)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address executorEOA_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) SetExecutorEOA(executorEOA_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SetExecutorEOA(&_SynapseExecutionServiceV1Harness.TransactOpts, executorEOA_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) SetGasOracle(opts *bind.TransactOpts, gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "setGasOracle", gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SetGasOracle(&_SynapseExecutionServiceV1Harness.TransactOpts, gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SetGasOracle(&_SynapseExecutionServiceV1Harness.TransactOpts, gasOracle_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) SetGlobalMarkup(opts *bind.TransactOpts, globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "setGlobalMarkup", globalMarkup_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) SetGlobalMarkup(globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SetGlobalMarkup(&_SynapseExecutionServiceV1Harness.TransactOpts, globalMarkup_)
}

// SetGlobalMarkup is a paid mutator transaction binding the contract method 0xcf4f578f.
//
// Solidity: function setGlobalMarkup(uint256 globalMarkup_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) SetGlobalMarkup(globalMarkup_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SetGlobalMarkup(&_SynapseExecutionServiceV1Harness.TransactOpts, globalMarkup_)
}

// SynapseExecutionServiceV1HarnessExecutionRequestedIterator is returned from FilterExecutionRequested and is used to iterate over the raw logs and unpacked data for ExecutionRequested events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessExecutionRequestedIterator struct {
	Event *SynapseExecutionServiceV1HarnessExecutionRequested // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessExecutionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessExecutionRequested)
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
		it.Event = new(SynapseExecutionServiceV1HarnessExecutionRequested)
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
func (it *SynapseExecutionServiceV1HarnessExecutionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessExecutionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessExecutionRequested represents a ExecutionRequested event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessExecutionRequested struct {
	TransactionId [32]byte
	Client        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterExecutionRequested(opts *bind.FilterOpts, transactionId [][32]byte) (*SynapseExecutionServiceV1HarnessExecutionRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessExecutionRequestedIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "ExecutionRequested", logs: logs, sub: sub}, nil
}

// WatchExecutionRequested is a free log subscription operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchExecutionRequested(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessExecutionRequested, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessExecutionRequested)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
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

// ParseExecutionRequested is a log parse operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseExecutionRequested(log types.Log) (*SynapseExecutionServiceV1HarnessExecutionRequested, error) {
	event := new(SynapseExecutionServiceV1HarnessExecutionRequested)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessExecutorEOASetIterator is returned from FilterExecutorEOASet and is used to iterate over the raw logs and unpacked data for ExecutorEOASet events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessExecutorEOASetIterator struct {
	Event *SynapseExecutionServiceV1HarnessExecutorEOASet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessExecutorEOASetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessExecutorEOASet)
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
		it.Event = new(SynapseExecutionServiceV1HarnessExecutorEOASet)
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
func (it *SynapseExecutionServiceV1HarnessExecutorEOASetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessExecutorEOASetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessExecutorEOASet represents a ExecutorEOASet event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessExecutorEOASet struct {
	ExecutorEOA common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutorEOASet is a free log retrieval operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterExecutorEOASet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1HarnessExecutorEOASetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "ExecutorEOASet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessExecutorEOASetIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "ExecutorEOASet", logs: logs, sub: sub}, nil
}

// WatchExecutorEOASet is a free log subscription operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchExecutorEOASet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessExecutorEOASet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "ExecutorEOASet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessExecutorEOASet)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "ExecutorEOASet", log); err != nil {
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

// ParseExecutorEOASet is a log parse operation binding the contract event 0x4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f71.
//
// Solidity: event ExecutorEOASet(address executorEOA)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseExecutorEOASet(log types.Log) (*SynapseExecutionServiceV1HarnessExecutorEOASet, error) {
	event := new(SynapseExecutionServiceV1HarnessExecutorEOASet)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "ExecutorEOASet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessGasOracleSetIterator is returned from FilterGasOracleSet and is used to iterate over the raw logs and unpacked data for GasOracleSet events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessGasOracleSetIterator struct {
	Event *SynapseExecutionServiceV1HarnessGasOracleSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessGasOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessGasOracleSet)
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
		it.Event = new(SynapseExecutionServiceV1HarnessGasOracleSet)
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
func (it *SynapseExecutionServiceV1HarnessGasOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessGasOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessGasOracleSet represents a GasOracleSet event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessGasOracleSet struct {
	GasOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOracleSet is a free log retrieval operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterGasOracleSet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1HarnessGasOracleSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "GasOracleSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessGasOracleSetIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "GasOracleSet", logs: logs, sub: sub}, nil
}

// WatchGasOracleSet is a free log subscription operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchGasOracleSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessGasOracleSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "GasOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessGasOracleSet)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "GasOracleSet", log); err != nil {
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

// ParseGasOracleSet is a log parse operation binding the contract event 0x3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa.
//
// Solidity: event GasOracleSet(address gasOracle)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseGasOracleSet(log types.Log) (*SynapseExecutionServiceV1HarnessGasOracleSet, error) {
	event := new(SynapseExecutionServiceV1HarnessGasOracleSet)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "GasOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessGlobalMarkupSetIterator is returned from FilterGlobalMarkupSet and is used to iterate over the raw logs and unpacked data for GlobalMarkupSet events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessGlobalMarkupSetIterator struct {
	Event *SynapseExecutionServiceV1HarnessGlobalMarkupSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessGlobalMarkupSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessGlobalMarkupSet)
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
		it.Event = new(SynapseExecutionServiceV1HarnessGlobalMarkupSet)
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
func (it *SynapseExecutionServiceV1HarnessGlobalMarkupSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessGlobalMarkupSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessGlobalMarkupSet represents a GlobalMarkupSet event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessGlobalMarkupSet struct {
	GlobalMarkup *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGlobalMarkupSet is a free log retrieval operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterGlobalMarkupSet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1HarnessGlobalMarkupSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "GlobalMarkupSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessGlobalMarkupSetIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "GlobalMarkupSet", logs: logs, sub: sub}, nil
}

// WatchGlobalMarkupSet is a free log subscription operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchGlobalMarkupSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessGlobalMarkupSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "GlobalMarkupSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessGlobalMarkupSet)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "GlobalMarkupSet", log); err != nil {
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

// ParseGlobalMarkupSet is a log parse operation binding the contract event 0x1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d.
//
// Solidity: event GlobalMarkupSet(uint256 globalMarkup)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseGlobalMarkupSet(log types.Log) (*SynapseExecutionServiceV1HarnessGlobalMarkupSet, error) {
	event := new(SynapseExecutionServiceV1HarnessGlobalMarkupSet)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "GlobalMarkupSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessInitializedIterator struct {
	Event *SynapseExecutionServiceV1HarnessInitialized // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessInitialized)
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
		it.Event = new(SynapseExecutionServiceV1HarnessInitialized)
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
func (it *SynapseExecutionServiceV1HarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessInitialized represents a Initialized event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*SynapseExecutionServiceV1HarnessInitializedIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessInitializedIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessInitialized)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseInitialized(log types.Log) (*SynapseExecutionServiceV1HarnessInitialized, error) {
	event := new(SynapseExecutionServiceV1HarnessInitialized)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessRoleAdminChangedIterator struct {
	Event *SynapseExecutionServiceV1HarnessRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessRoleAdminChanged)
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
		it.Event = new(SynapseExecutionServiceV1HarnessRoleAdminChanged)
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
func (it *SynapseExecutionServiceV1HarnessRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessRoleAdminChanged represents a RoleAdminChanged event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SynapseExecutionServiceV1HarnessRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessRoleAdminChangedIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessRoleAdminChanged)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseRoleAdminChanged(log types.Log) (*SynapseExecutionServiceV1HarnessRoleAdminChanged, error) {
	event := new(SynapseExecutionServiceV1HarnessRoleAdminChanged)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessRoleGrantedIterator struct {
	Event *SynapseExecutionServiceV1HarnessRoleGranted // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessRoleGranted)
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
		it.Event = new(SynapseExecutionServiceV1HarnessRoleGranted)
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
func (it *SynapseExecutionServiceV1HarnessRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessRoleGranted represents a RoleGranted event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SynapseExecutionServiceV1HarnessRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessRoleGrantedIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessRoleGranted)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseRoleGranted(log types.Log) (*SynapseExecutionServiceV1HarnessRoleGranted, error) {
	event := new(SynapseExecutionServiceV1HarnessRoleGranted)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessRoleRevokedIterator struct {
	Event *SynapseExecutionServiceV1HarnessRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessRoleRevoked)
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
		it.Event = new(SynapseExecutionServiceV1HarnessRoleRevoked)
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
func (it *SynapseExecutionServiceV1HarnessRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessRoleRevoked represents a RoleRevoked event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SynapseExecutionServiceV1HarnessRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessRoleRevokedIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessRoleRevoked)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseRoleRevoked(log types.Log) (*SynapseExecutionServiceV1HarnessRoleRevoked, error) {
	event := new(SynapseExecutionServiceV1HarnessRoleRevoked)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VersionedPayloadLibMetaData contains all meta data concerning the VersionedPayloadLib contract.
var VersionedPayloadLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__TooShort\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122022e61c11ebef79fa94eb8e95509478ba8fca8bbecb1b3dd64c22ae37f59ec4f864736f6c63430008140033",
}

// VersionedPayloadLibABI is the input ABI used to generate the binding from.
// Deprecated: Use VersionedPayloadLibMetaData.ABI instead.
var VersionedPayloadLibABI = VersionedPayloadLibMetaData.ABI

// VersionedPayloadLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VersionedPayloadLibMetaData.Bin instead.
var VersionedPayloadLibBin = VersionedPayloadLibMetaData.Bin

// DeployVersionedPayloadLib deploys a new Ethereum contract, binding an instance of VersionedPayloadLib to it.
func DeployVersionedPayloadLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VersionedPayloadLib, error) {
	parsed, err := VersionedPayloadLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VersionedPayloadLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VersionedPayloadLib{VersionedPayloadLibCaller: VersionedPayloadLibCaller{contract: contract}, VersionedPayloadLibTransactor: VersionedPayloadLibTransactor{contract: contract}, VersionedPayloadLibFilterer: VersionedPayloadLibFilterer{contract: contract}}, nil
}

// VersionedPayloadLib is an auto generated Go binding around an Ethereum contract.
type VersionedPayloadLib struct {
	VersionedPayloadLibCaller     // Read-only binding to the contract
	VersionedPayloadLibTransactor // Write-only binding to the contract
	VersionedPayloadLibFilterer   // Log filterer for contract events
}

// VersionedPayloadLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type VersionedPayloadLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedPayloadLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VersionedPayloadLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedPayloadLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VersionedPayloadLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedPayloadLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VersionedPayloadLibSession struct {
	Contract     *VersionedPayloadLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VersionedPayloadLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VersionedPayloadLibCallerSession struct {
	Contract *VersionedPayloadLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VersionedPayloadLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VersionedPayloadLibTransactorSession struct {
	Contract     *VersionedPayloadLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VersionedPayloadLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type VersionedPayloadLibRaw struct {
	Contract *VersionedPayloadLib // Generic contract binding to access the raw methods on
}

// VersionedPayloadLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VersionedPayloadLibCallerRaw struct {
	Contract *VersionedPayloadLibCaller // Generic read-only contract binding to access the raw methods on
}

// VersionedPayloadLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VersionedPayloadLibTransactorRaw struct {
	Contract *VersionedPayloadLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVersionedPayloadLib creates a new instance of VersionedPayloadLib, bound to a specific deployed contract.
func NewVersionedPayloadLib(address common.Address, backend bind.ContractBackend) (*VersionedPayloadLib, error) {
	contract, err := bindVersionedPayloadLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VersionedPayloadLib{VersionedPayloadLibCaller: VersionedPayloadLibCaller{contract: contract}, VersionedPayloadLibTransactor: VersionedPayloadLibTransactor{contract: contract}, VersionedPayloadLibFilterer: VersionedPayloadLibFilterer{contract: contract}}, nil
}

// NewVersionedPayloadLibCaller creates a new read-only instance of VersionedPayloadLib, bound to a specific deployed contract.
func NewVersionedPayloadLibCaller(address common.Address, caller bind.ContractCaller) (*VersionedPayloadLibCaller, error) {
	contract, err := bindVersionedPayloadLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedPayloadLibCaller{contract: contract}, nil
}

// NewVersionedPayloadLibTransactor creates a new write-only instance of VersionedPayloadLib, bound to a specific deployed contract.
func NewVersionedPayloadLibTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionedPayloadLibTransactor, error) {
	contract, err := bindVersionedPayloadLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedPayloadLibTransactor{contract: contract}, nil
}

// NewVersionedPayloadLibFilterer creates a new log filterer instance of VersionedPayloadLib, bound to a specific deployed contract.
func NewVersionedPayloadLibFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionedPayloadLibFilterer, error) {
	contract, err := bindVersionedPayloadLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionedPayloadLibFilterer{contract: contract}, nil
}

// bindVersionedPayloadLib binds a generic wrapper to an already deployed contract.
func bindVersionedPayloadLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VersionedPayloadLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VersionedPayloadLib *VersionedPayloadLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VersionedPayloadLib.Contract.VersionedPayloadLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VersionedPayloadLib *VersionedPayloadLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedPayloadLib.Contract.VersionedPayloadLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VersionedPayloadLib *VersionedPayloadLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VersionedPayloadLib.Contract.VersionedPayloadLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VersionedPayloadLib *VersionedPayloadLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VersionedPayloadLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VersionedPayloadLib *VersionedPayloadLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedPayloadLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VersionedPayloadLib *VersionedPayloadLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VersionedPayloadLib.Contract.contract.Transact(opts, method, params...)
}
