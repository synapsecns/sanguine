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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220792fe26a604cfce1c4fb18c500672e03a11f5808a3a595a460a409f4e1b2821364736f6c63430008140033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ClaimableFeesMetaData contains all meta data concerning the ClaimableFees contract.
var ClaimableFeesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"name\":\"ClaimableFees__ClaimerFractionAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableFees__FeeAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableFees__FeeRecipientZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"}],\"name\":\"ClaimerFractionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerReward\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimerFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d294f093": "claimFees()",
		"c354bd6e": "getClaimableAmount()",
		"4f199114": "getClaimerFraction()",
		"26533fe9": "getClaimerReward()",
		"4ccb20c0": "getFeeRecipient()",
	},
}

// ClaimableFeesABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimableFeesMetaData.ABI instead.
var ClaimableFeesABI = ClaimableFeesMetaData.ABI

// Deprecated: Use ClaimableFeesMetaData.Sigs instead.
// ClaimableFeesFuncSigs maps the 4-byte function signature to its string representation.
var ClaimableFeesFuncSigs = ClaimableFeesMetaData.Sigs

// ClaimableFees is an auto generated Go binding around an Ethereum contract.
type ClaimableFees struct {
	ClaimableFeesCaller     // Read-only binding to the contract
	ClaimableFeesTransactor // Write-only binding to the contract
	ClaimableFeesFilterer   // Log filterer for contract events
}

// ClaimableFeesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimableFeesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableFeesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimableFeesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableFeesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimableFeesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableFeesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimableFeesSession struct {
	Contract     *ClaimableFees    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimableFeesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimableFeesCallerSession struct {
	Contract *ClaimableFeesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ClaimableFeesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimableFeesTransactorSession struct {
	Contract     *ClaimableFeesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ClaimableFeesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimableFeesRaw struct {
	Contract *ClaimableFees // Generic contract binding to access the raw methods on
}

// ClaimableFeesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimableFeesCallerRaw struct {
	Contract *ClaimableFeesCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimableFeesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimableFeesTransactorRaw struct {
	Contract *ClaimableFeesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimableFees creates a new instance of ClaimableFees, bound to a specific deployed contract.
func NewClaimableFees(address common.Address, backend bind.ContractBackend) (*ClaimableFees, error) {
	contract, err := bindClaimableFees(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimableFees{ClaimableFeesCaller: ClaimableFeesCaller{contract: contract}, ClaimableFeesTransactor: ClaimableFeesTransactor{contract: contract}, ClaimableFeesFilterer: ClaimableFeesFilterer{contract: contract}}, nil
}

// NewClaimableFeesCaller creates a new read-only instance of ClaimableFees, bound to a specific deployed contract.
func NewClaimableFeesCaller(address common.Address, caller bind.ContractCaller) (*ClaimableFeesCaller, error) {
	contract, err := bindClaimableFees(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesCaller{contract: contract}, nil
}

// NewClaimableFeesTransactor creates a new write-only instance of ClaimableFees, bound to a specific deployed contract.
func NewClaimableFeesTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimableFeesTransactor, error) {
	contract, err := bindClaimableFees(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesTransactor{contract: contract}, nil
}

// NewClaimableFeesFilterer creates a new log filterer instance of ClaimableFees, bound to a specific deployed contract.
func NewClaimableFeesFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimableFeesFilterer, error) {
	contract, err := bindClaimableFees(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesFilterer{contract: contract}, nil
}

// bindClaimableFees binds a generic wrapper to an already deployed contract.
func bindClaimableFees(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimableFeesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimableFees *ClaimableFeesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimableFees.Contract.ClaimableFeesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimableFees *ClaimableFeesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableFees.Contract.ClaimableFeesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimableFees *ClaimableFeesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimableFees.Contract.ClaimableFeesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimableFees *ClaimableFeesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimableFees.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimableFees *ClaimableFeesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableFees.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimableFees *ClaimableFeesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimableFees.Contract.contract.Transact(opts, method, params...)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_ClaimableFees *ClaimableFeesCaller) GetClaimableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimableFees.contract.Call(opts, &out, "getClaimableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_ClaimableFees *ClaimableFeesSession) GetClaimableAmount() (*big.Int, error) {
	return _ClaimableFees.Contract.GetClaimableAmount(&_ClaimableFees.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_ClaimableFees *ClaimableFeesCallerSession) GetClaimableAmount() (*big.Int, error) {
	return _ClaimableFees.Contract.GetClaimableAmount(&_ClaimableFees.CallOpts)
}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_ClaimableFees *ClaimableFeesCaller) GetClaimerFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimableFees.contract.Call(opts, &out, "getClaimerFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_ClaimableFees *ClaimableFeesSession) GetClaimerFraction() (*big.Int, error) {
	return _ClaimableFees.Contract.GetClaimerFraction(&_ClaimableFees.CallOpts)
}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_ClaimableFees *ClaimableFeesCallerSession) GetClaimerFraction() (*big.Int, error) {
	return _ClaimableFees.Contract.GetClaimerFraction(&_ClaimableFees.CallOpts)
}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_ClaimableFees *ClaimableFeesCaller) GetClaimerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimableFees.contract.Call(opts, &out, "getClaimerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_ClaimableFees *ClaimableFeesSession) GetClaimerReward() (*big.Int, error) {
	return _ClaimableFees.Contract.GetClaimerReward(&_ClaimableFees.CallOpts)
}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_ClaimableFees *ClaimableFeesCallerSession) GetClaimerReward() (*big.Int, error) {
	return _ClaimableFees.Contract.GetClaimerReward(&_ClaimableFees.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_ClaimableFees *ClaimableFeesCaller) GetFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimableFees.contract.Call(opts, &out, "getFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_ClaimableFees *ClaimableFeesSession) GetFeeRecipient() (common.Address, error) {
	return _ClaimableFees.Contract.GetFeeRecipient(&_ClaimableFees.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_ClaimableFees *ClaimableFeesCallerSession) GetFeeRecipient() (common.Address, error) {
	return _ClaimableFees.Contract.GetFeeRecipient(&_ClaimableFees.CallOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_ClaimableFees *ClaimableFeesTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableFees.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_ClaimableFees *ClaimableFeesSession) ClaimFees() (*types.Transaction, error) {
	return _ClaimableFees.Contract.ClaimFees(&_ClaimableFees.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_ClaimableFees *ClaimableFeesTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _ClaimableFees.Contract.ClaimFees(&_ClaimableFees.TransactOpts)
}

// ClaimableFeesClaimerFractionSetIterator is returned from FilterClaimerFractionSet and is used to iterate over the raw logs and unpacked data for ClaimerFractionSet events raised by the ClaimableFees contract.
type ClaimableFeesClaimerFractionSetIterator struct {
	Event *ClaimableFeesClaimerFractionSet // Event containing the contract specifics and raw log

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
func (it *ClaimableFeesClaimerFractionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableFeesClaimerFractionSet)
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
		it.Event = new(ClaimableFeesClaimerFractionSet)
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
func (it *ClaimableFeesClaimerFractionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableFeesClaimerFractionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableFeesClaimerFractionSet represents a ClaimerFractionSet event raised by the ClaimableFees contract.
type ClaimableFeesClaimerFractionSet struct {
	ClaimerFraction *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimerFractionSet is a free log retrieval operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_ClaimableFees *ClaimableFeesFilterer) FilterClaimerFractionSet(opts *bind.FilterOpts) (*ClaimableFeesClaimerFractionSetIterator, error) {

	logs, sub, err := _ClaimableFees.contract.FilterLogs(opts, "ClaimerFractionSet")
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesClaimerFractionSetIterator{contract: _ClaimableFees.contract, event: "ClaimerFractionSet", logs: logs, sub: sub}, nil
}

// WatchClaimerFractionSet is a free log subscription operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_ClaimableFees *ClaimableFeesFilterer) WatchClaimerFractionSet(opts *bind.WatchOpts, sink chan<- *ClaimableFeesClaimerFractionSet) (event.Subscription, error) {

	logs, sub, err := _ClaimableFees.contract.WatchLogs(opts, "ClaimerFractionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableFeesClaimerFractionSet)
				if err := _ClaimableFees.contract.UnpackLog(event, "ClaimerFractionSet", log); err != nil {
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

// ParseClaimerFractionSet is a log parse operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_ClaimableFees *ClaimableFeesFilterer) ParseClaimerFractionSet(log types.Log) (*ClaimableFeesClaimerFractionSet, error) {
	event := new(ClaimableFeesClaimerFractionSet)
	if err := _ClaimableFees.contract.UnpackLog(event, "ClaimerFractionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimableFeesFeeRecipientSetIterator is returned from FilterFeeRecipientSet and is used to iterate over the raw logs and unpacked data for FeeRecipientSet events raised by the ClaimableFees contract.
type ClaimableFeesFeeRecipientSetIterator struct {
	Event *ClaimableFeesFeeRecipientSet // Event containing the contract specifics and raw log

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
func (it *ClaimableFeesFeeRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableFeesFeeRecipientSet)
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
		it.Event = new(ClaimableFeesFeeRecipientSet)
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
func (it *ClaimableFeesFeeRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableFeesFeeRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableFeesFeeRecipientSet represents a FeeRecipientSet event raised by the ClaimableFees contract.
type ClaimableFeesFeeRecipientSet struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientSet is a free log retrieval operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_ClaimableFees *ClaimableFeesFilterer) FilterFeeRecipientSet(opts *bind.FilterOpts) (*ClaimableFeesFeeRecipientSetIterator, error) {

	logs, sub, err := _ClaimableFees.contract.FilterLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesFeeRecipientSetIterator{contract: _ClaimableFees.contract, event: "FeeRecipientSet", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientSet is a free log subscription operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_ClaimableFees *ClaimableFeesFilterer) WatchFeeRecipientSet(opts *bind.WatchOpts, sink chan<- *ClaimableFeesFeeRecipientSet) (event.Subscription, error) {

	logs, sub, err := _ClaimableFees.contract.WatchLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableFeesFeeRecipientSet)
				if err := _ClaimableFees.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
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

// ParseFeeRecipientSet is a log parse operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_ClaimableFees *ClaimableFeesFilterer) ParseFeeRecipientSet(log types.Log) (*ClaimableFeesFeeRecipientSet, error) {
	event := new(ClaimableFeesFeeRecipientSet)
	if err := _ClaimableFees.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimableFeesFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the ClaimableFees contract.
type ClaimableFeesFeesClaimedIterator struct {
	Event *ClaimableFeesFeesClaimed // Event containing the contract specifics and raw log

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
func (it *ClaimableFeesFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableFeesFeesClaimed)
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
		it.Event = new(ClaimableFeesFeesClaimed)
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
func (it *ClaimableFeesFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableFeesFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableFeesFeesClaimed represents a FeesClaimed event raised by the ClaimableFees contract.
type ClaimableFeesFeesClaimed struct {
	FeeRecipient  common.Address
	ClaimedFees   *big.Int
	Claimer       common.Address
	ClaimerReward *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_ClaimableFees *ClaimableFeesFilterer) FilterFeesClaimed(opts *bind.FilterOpts) (*ClaimableFeesFeesClaimedIterator, error) {

	logs, sub, err := _ClaimableFees.contract.FilterLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesFeesClaimedIterator{contract: _ClaimableFees.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_ClaimableFees *ClaimableFeesFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *ClaimableFeesFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _ClaimableFees.contract.WatchLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableFeesFeesClaimed)
				if err := _ClaimableFees.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_ClaimableFees *ClaimableFeesFilterer) ParseFeesClaimed(log types.Log) (*ClaimableFeesFeesClaimed, error) {
	event := new(ClaimableFeesFeesClaimed)
	if err := _ClaimableFees.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimableFeesEventsMetaData contains all meta data concerning the ClaimableFeesEvents contract.
var ClaimableFeesEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"}],\"name\":\"ClaimerFractionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerReward\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"}]",
}

// ClaimableFeesEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimableFeesEventsMetaData.ABI instead.
var ClaimableFeesEventsABI = ClaimableFeesEventsMetaData.ABI

// ClaimableFeesEvents is an auto generated Go binding around an Ethereum contract.
type ClaimableFeesEvents struct {
	ClaimableFeesEventsCaller     // Read-only binding to the contract
	ClaimableFeesEventsTransactor // Write-only binding to the contract
	ClaimableFeesEventsFilterer   // Log filterer for contract events
}

// ClaimableFeesEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimableFeesEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableFeesEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimableFeesEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableFeesEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimableFeesEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableFeesEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimableFeesEventsSession struct {
	Contract     *ClaimableFeesEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ClaimableFeesEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimableFeesEventsCallerSession struct {
	Contract *ClaimableFeesEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ClaimableFeesEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimableFeesEventsTransactorSession struct {
	Contract     *ClaimableFeesEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ClaimableFeesEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimableFeesEventsRaw struct {
	Contract *ClaimableFeesEvents // Generic contract binding to access the raw methods on
}

// ClaimableFeesEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimableFeesEventsCallerRaw struct {
	Contract *ClaimableFeesEventsCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimableFeesEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimableFeesEventsTransactorRaw struct {
	Contract *ClaimableFeesEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimableFeesEvents creates a new instance of ClaimableFeesEvents, bound to a specific deployed contract.
func NewClaimableFeesEvents(address common.Address, backend bind.ContractBackend) (*ClaimableFeesEvents, error) {
	contract, err := bindClaimableFeesEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesEvents{ClaimableFeesEventsCaller: ClaimableFeesEventsCaller{contract: contract}, ClaimableFeesEventsTransactor: ClaimableFeesEventsTransactor{contract: contract}, ClaimableFeesEventsFilterer: ClaimableFeesEventsFilterer{contract: contract}}, nil
}

// NewClaimableFeesEventsCaller creates a new read-only instance of ClaimableFeesEvents, bound to a specific deployed contract.
func NewClaimableFeesEventsCaller(address common.Address, caller bind.ContractCaller) (*ClaimableFeesEventsCaller, error) {
	contract, err := bindClaimableFeesEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesEventsCaller{contract: contract}, nil
}

// NewClaimableFeesEventsTransactor creates a new write-only instance of ClaimableFeesEvents, bound to a specific deployed contract.
func NewClaimableFeesEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimableFeesEventsTransactor, error) {
	contract, err := bindClaimableFeesEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesEventsTransactor{contract: contract}, nil
}

// NewClaimableFeesEventsFilterer creates a new log filterer instance of ClaimableFeesEvents, bound to a specific deployed contract.
func NewClaimableFeesEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimableFeesEventsFilterer, error) {
	contract, err := bindClaimableFeesEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesEventsFilterer{contract: contract}, nil
}

// bindClaimableFeesEvents binds a generic wrapper to an already deployed contract.
func bindClaimableFeesEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimableFeesEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimableFeesEvents *ClaimableFeesEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimableFeesEvents.Contract.ClaimableFeesEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimableFeesEvents *ClaimableFeesEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableFeesEvents.Contract.ClaimableFeesEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimableFeesEvents *ClaimableFeesEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimableFeesEvents.Contract.ClaimableFeesEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimableFeesEvents *ClaimableFeesEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimableFeesEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimableFeesEvents *ClaimableFeesEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableFeesEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimableFeesEvents *ClaimableFeesEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimableFeesEvents.Contract.contract.Transact(opts, method, params...)
}

// ClaimableFeesEventsClaimerFractionSetIterator is returned from FilterClaimerFractionSet and is used to iterate over the raw logs and unpacked data for ClaimerFractionSet events raised by the ClaimableFeesEvents contract.
type ClaimableFeesEventsClaimerFractionSetIterator struct {
	Event *ClaimableFeesEventsClaimerFractionSet // Event containing the contract specifics and raw log

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
func (it *ClaimableFeesEventsClaimerFractionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableFeesEventsClaimerFractionSet)
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
		it.Event = new(ClaimableFeesEventsClaimerFractionSet)
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
func (it *ClaimableFeesEventsClaimerFractionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableFeesEventsClaimerFractionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableFeesEventsClaimerFractionSet represents a ClaimerFractionSet event raised by the ClaimableFeesEvents contract.
type ClaimableFeesEventsClaimerFractionSet struct {
	ClaimerFraction *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimerFractionSet is a free log retrieval operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) FilterClaimerFractionSet(opts *bind.FilterOpts) (*ClaimableFeesEventsClaimerFractionSetIterator, error) {

	logs, sub, err := _ClaimableFeesEvents.contract.FilterLogs(opts, "ClaimerFractionSet")
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesEventsClaimerFractionSetIterator{contract: _ClaimableFeesEvents.contract, event: "ClaimerFractionSet", logs: logs, sub: sub}, nil
}

// WatchClaimerFractionSet is a free log subscription operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) WatchClaimerFractionSet(opts *bind.WatchOpts, sink chan<- *ClaimableFeesEventsClaimerFractionSet) (event.Subscription, error) {

	logs, sub, err := _ClaimableFeesEvents.contract.WatchLogs(opts, "ClaimerFractionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableFeesEventsClaimerFractionSet)
				if err := _ClaimableFeesEvents.contract.UnpackLog(event, "ClaimerFractionSet", log); err != nil {
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

// ParseClaimerFractionSet is a log parse operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) ParseClaimerFractionSet(log types.Log) (*ClaimableFeesEventsClaimerFractionSet, error) {
	event := new(ClaimableFeesEventsClaimerFractionSet)
	if err := _ClaimableFeesEvents.contract.UnpackLog(event, "ClaimerFractionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimableFeesEventsFeeRecipientSetIterator is returned from FilterFeeRecipientSet and is used to iterate over the raw logs and unpacked data for FeeRecipientSet events raised by the ClaimableFeesEvents contract.
type ClaimableFeesEventsFeeRecipientSetIterator struct {
	Event *ClaimableFeesEventsFeeRecipientSet // Event containing the contract specifics and raw log

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
func (it *ClaimableFeesEventsFeeRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableFeesEventsFeeRecipientSet)
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
		it.Event = new(ClaimableFeesEventsFeeRecipientSet)
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
func (it *ClaimableFeesEventsFeeRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableFeesEventsFeeRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableFeesEventsFeeRecipientSet represents a FeeRecipientSet event raised by the ClaimableFeesEvents contract.
type ClaimableFeesEventsFeeRecipientSet struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientSet is a free log retrieval operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) FilterFeeRecipientSet(opts *bind.FilterOpts) (*ClaimableFeesEventsFeeRecipientSetIterator, error) {

	logs, sub, err := _ClaimableFeesEvents.contract.FilterLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesEventsFeeRecipientSetIterator{contract: _ClaimableFeesEvents.contract, event: "FeeRecipientSet", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientSet is a free log subscription operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) WatchFeeRecipientSet(opts *bind.WatchOpts, sink chan<- *ClaimableFeesEventsFeeRecipientSet) (event.Subscription, error) {

	logs, sub, err := _ClaimableFeesEvents.contract.WatchLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableFeesEventsFeeRecipientSet)
				if err := _ClaimableFeesEvents.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
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

// ParseFeeRecipientSet is a log parse operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) ParseFeeRecipientSet(log types.Log) (*ClaimableFeesEventsFeeRecipientSet, error) {
	event := new(ClaimableFeesEventsFeeRecipientSet)
	if err := _ClaimableFeesEvents.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimableFeesEventsFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the ClaimableFeesEvents contract.
type ClaimableFeesEventsFeesClaimedIterator struct {
	Event *ClaimableFeesEventsFeesClaimed // Event containing the contract specifics and raw log

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
func (it *ClaimableFeesEventsFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableFeesEventsFeesClaimed)
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
		it.Event = new(ClaimableFeesEventsFeesClaimed)
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
func (it *ClaimableFeesEventsFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableFeesEventsFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableFeesEventsFeesClaimed represents a FeesClaimed event raised by the ClaimableFeesEvents contract.
type ClaimableFeesEventsFeesClaimed struct {
	FeeRecipient  common.Address
	ClaimedFees   *big.Int
	Claimer       common.Address
	ClaimerReward *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) FilterFeesClaimed(opts *bind.FilterOpts) (*ClaimableFeesEventsFeesClaimedIterator, error) {

	logs, sub, err := _ClaimableFeesEvents.contract.FilterLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return &ClaimableFeesEventsFeesClaimedIterator{contract: _ClaimableFeesEvents.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *ClaimableFeesEventsFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _ClaimableFeesEvents.contract.WatchLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableFeesEventsFeesClaimed)
				if err := _ClaimableFeesEvents.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_ClaimableFeesEvents *ClaimableFeesEventsFilterer) ParseFeesClaimed(log types.Log) (*ClaimableFeesEventsFeesClaimed, error) {
	event := new(ClaimableFeesEventsFeesClaimed)
	if err := _ClaimableFeesEvents.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// IClaimableFeesMetaData contains all meta data concerning the IClaimableFees contract.
var IClaimableFeesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"name\":\"ClaimableFees__ClaimerFractionAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableFees__FeeAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableFees__FeeRecipientZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimerFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d294f093": "claimFees()",
		"c354bd6e": "getClaimableAmount()",
		"4f199114": "getClaimerFraction()",
		"26533fe9": "getClaimerReward()",
		"4ccb20c0": "getFeeRecipient()",
	},
}

// IClaimableFeesABI is the input ABI used to generate the binding from.
// Deprecated: Use IClaimableFeesMetaData.ABI instead.
var IClaimableFeesABI = IClaimableFeesMetaData.ABI

// Deprecated: Use IClaimableFeesMetaData.Sigs instead.
// IClaimableFeesFuncSigs maps the 4-byte function signature to its string representation.
var IClaimableFeesFuncSigs = IClaimableFeesMetaData.Sigs

// IClaimableFees is an auto generated Go binding around an Ethereum contract.
type IClaimableFees struct {
	IClaimableFeesCaller     // Read-only binding to the contract
	IClaimableFeesTransactor // Write-only binding to the contract
	IClaimableFeesFilterer   // Log filterer for contract events
}

// IClaimableFeesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IClaimableFeesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClaimableFeesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IClaimableFeesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClaimableFeesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IClaimableFeesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClaimableFeesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IClaimableFeesSession struct {
	Contract     *IClaimableFees   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IClaimableFeesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IClaimableFeesCallerSession struct {
	Contract *IClaimableFeesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IClaimableFeesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IClaimableFeesTransactorSession struct {
	Contract     *IClaimableFeesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IClaimableFeesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IClaimableFeesRaw struct {
	Contract *IClaimableFees // Generic contract binding to access the raw methods on
}

// IClaimableFeesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IClaimableFeesCallerRaw struct {
	Contract *IClaimableFeesCaller // Generic read-only contract binding to access the raw methods on
}

// IClaimableFeesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IClaimableFeesTransactorRaw struct {
	Contract *IClaimableFeesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIClaimableFees creates a new instance of IClaimableFees, bound to a specific deployed contract.
func NewIClaimableFees(address common.Address, backend bind.ContractBackend) (*IClaimableFees, error) {
	contract, err := bindIClaimableFees(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IClaimableFees{IClaimableFeesCaller: IClaimableFeesCaller{contract: contract}, IClaimableFeesTransactor: IClaimableFeesTransactor{contract: contract}, IClaimableFeesFilterer: IClaimableFeesFilterer{contract: contract}}, nil
}

// NewIClaimableFeesCaller creates a new read-only instance of IClaimableFees, bound to a specific deployed contract.
func NewIClaimableFeesCaller(address common.Address, caller bind.ContractCaller) (*IClaimableFeesCaller, error) {
	contract, err := bindIClaimableFees(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IClaimableFeesCaller{contract: contract}, nil
}

// NewIClaimableFeesTransactor creates a new write-only instance of IClaimableFees, bound to a specific deployed contract.
func NewIClaimableFeesTransactor(address common.Address, transactor bind.ContractTransactor) (*IClaimableFeesTransactor, error) {
	contract, err := bindIClaimableFees(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IClaimableFeesTransactor{contract: contract}, nil
}

// NewIClaimableFeesFilterer creates a new log filterer instance of IClaimableFees, bound to a specific deployed contract.
func NewIClaimableFeesFilterer(address common.Address, filterer bind.ContractFilterer) (*IClaimableFeesFilterer, error) {
	contract, err := bindIClaimableFees(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IClaimableFeesFilterer{contract: contract}, nil
}

// bindIClaimableFees binds a generic wrapper to an already deployed contract.
func bindIClaimableFees(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IClaimableFeesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IClaimableFees *IClaimableFeesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IClaimableFees.Contract.IClaimableFeesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IClaimableFees *IClaimableFeesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IClaimableFees.Contract.IClaimableFeesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IClaimableFees *IClaimableFeesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IClaimableFees.Contract.IClaimableFeesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IClaimableFees *IClaimableFeesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IClaimableFees.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IClaimableFees *IClaimableFeesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IClaimableFees.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IClaimableFees *IClaimableFeesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IClaimableFees.Contract.contract.Transact(opts, method, params...)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_IClaimableFees *IClaimableFeesCaller) GetClaimableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IClaimableFees.contract.Call(opts, &out, "getClaimableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_IClaimableFees *IClaimableFeesSession) GetClaimableAmount() (*big.Int, error) {
	return _IClaimableFees.Contract.GetClaimableAmount(&_IClaimableFees.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_IClaimableFees *IClaimableFeesCallerSession) GetClaimableAmount() (*big.Int, error) {
	return _IClaimableFees.Contract.GetClaimableAmount(&_IClaimableFees.CallOpts)
}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_IClaimableFees *IClaimableFeesCaller) GetClaimerFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IClaimableFees.contract.Call(opts, &out, "getClaimerFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_IClaimableFees *IClaimableFeesSession) GetClaimerFraction() (*big.Int, error) {
	return _IClaimableFees.Contract.GetClaimerFraction(&_IClaimableFees.CallOpts)
}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_IClaimableFees *IClaimableFeesCallerSession) GetClaimerFraction() (*big.Int, error) {
	return _IClaimableFees.Contract.GetClaimerFraction(&_IClaimableFees.CallOpts)
}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_IClaimableFees *IClaimableFeesCaller) GetClaimerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IClaimableFees.contract.Call(opts, &out, "getClaimerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_IClaimableFees *IClaimableFeesSession) GetClaimerReward() (*big.Int, error) {
	return _IClaimableFees.Contract.GetClaimerReward(&_IClaimableFees.CallOpts)
}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_IClaimableFees *IClaimableFeesCallerSession) GetClaimerReward() (*big.Int, error) {
	return _IClaimableFees.Contract.GetClaimerReward(&_IClaimableFees.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_IClaimableFees *IClaimableFeesCaller) GetFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IClaimableFees.contract.Call(opts, &out, "getFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_IClaimableFees *IClaimableFeesSession) GetFeeRecipient() (common.Address, error) {
	return _IClaimableFees.Contract.GetFeeRecipient(&_IClaimableFees.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_IClaimableFees *IClaimableFeesCallerSession) GetFeeRecipient() (common.Address, error) {
	return _IClaimableFees.Contract.GetFeeRecipient(&_IClaimableFees.CallOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_IClaimableFees *IClaimableFeesTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IClaimableFees.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_IClaimableFees *IClaimableFeesSession) ClaimFees() (*types.Transaction, error) {
	return _IClaimableFees.Contract.ClaimFees(&_IClaimableFees.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_IClaimableFees *IClaimableFeesTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _IClaimableFees.Contract.ClaimFees(&_IClaimableFees.TransactOpts)
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
	ABI: "[{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestTxExecution\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62014bad": "executorEOA()",
		"96fda4da": "getExecutionFee(uint64,uint256,bytes)",
		"58efb47d": "requestTxExecution(uint64,uint256,bytes32,bytes)",
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

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_IExecutionService *IExecutionServiceTransactor) RequestTxExecution(opts *bind.TransactOpts, dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _IExecutionService.contract.Transact(opts, "requestTxExecution", dstChainId, txPayloadSize, transactionId, options)
}

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_IExecutionService *IExecutionServiceSession) RequestTxExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _IExecutionService.Contract.RequestTxExecution(&_IExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, options)
}

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_IExecutionService *IExecutionServiceTransactorSession) RequestTxExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _IExecutionService.Contract.RequestTxExecution(&_IExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, options)
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
	ABI: "[{\"inputs\":[],\"name\":\"SynapseExecutionService__ExecutorZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRequired\",\"type\":\"uint256\"}],\"name\":\"SynapseExecutionService__FeeAmountBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"SynapseExecutionService__GasOracleNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__GasOracleZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"SynapseExecutionService__OptionsVersionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalMarkup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestTxExecution\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"}],\"name\":\"setClaimerFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executorEOA_\",\"type\":\"address\"}],\"name\":\"setExecutorEOA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalMarkup_\",\"type\":\"uint256\"}],\"name\":\"setGlobalMarkup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62014bad": "executorEOA()",
		"5d62a8dd": "gasOracle()",
		"96fda4da": "getExecutionFee(uint64,uint256,bytes)",
		"efd07ec2": "globalMarkup()",
		"58efb47d": "requestTxExecution(uint64,uint256,bytes32,bytes)",
		"a9bc769b": "setClaimerFraction(uint256)",
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

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Transactor) RequestTxExecution(opts *bind.TransactOpts, dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.contract.Transact(opts, "requestTxExecution", dstChainId, txPayloadSize, transactionId, options)
}

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) RequestTxExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.RequestTxExecution(&_ISynapseExecutionServiceV1.TransactOpts, dstChainId, txPayloadSize, transactionId, options)
}

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1TransactorSession) RequestTxExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.RequestTxExecution(&_ISynapseExecutionServiceV1.TransactOpts, dstChainId, txPayloadSize, transactionId, options)
}

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Transactor) SetClaimerFraction(opts *bind.TransactOpts, claimerFraction *big.Int) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.contract.Transact(opts, "setClaimerFraction", claimerFraction)
}

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1Session) SetClaimerFraction(claimerFraction *big.Int) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.SetClaimerFraction(&_ISynapseExecutionServiceV1.TransactOpts, claimerFraction)
}

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction) returns()
func (_ISynapseExecutionServiceV1 *ISynapseExecutionServiceV1TransactorSession) SetClaimerFraction(claimerFraction *big.Int) (*types.Transaction, error) {
	return _ISynapseExecutionServiceV1.Contract.SetClaimerFraction(&_ISynapseExecutionServiceV1.TransactOpts, claimerFraction)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"OptionsLib__VersionInvalid\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f4b64ded52ede86acf8f97962c12b08467036aea41170f2556e0b6c071b6188364736f6c63430008140033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executorEOA\",\"type\":\"address\"}],\"name\":\"ExecutorEOASet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalMarkup\",\"type\":\"uint256\"}],\"name\":\"GlobalMarkupSet\",\"type\":\"event\"}]",
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
	ExecutionFee  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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

// WatchExecutionRequested is a free log subscription operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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

// ParseExecutionRequested is a log parse operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"name\":\"ClaimableFees__ClaimerFractionAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableFees__FeeAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableFees__FeeRecipientZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"OptionsLib__VersionInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__ExecutorZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRequired\",\"type\":\"uint256\"}],\"name\":\"SynapseExecutionService__FeeAmountBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"SynapseExecutionService__GasOracleNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__GasOracleZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"SynapseExecutionService__OptionsVersionNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__PayloadTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"}],\"name\":\"ClaimerFractionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executorEOA\",\"type\":\"address\"}],\"name\":\"ExecutorEOASet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerReward\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalMarkup\",\"type\":\"uint256\"}],\"name\":\"GlobalMarkupSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IC_CLIENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimerFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalMarkup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestTxExecution\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimerFraction_\",\"type\":\"uint256\"}],\"name\":\"setClaimerFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executorEOA_\",\"type\":\"address\"}],\"name\":\"setExecutorEOA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalMarkup_\",\"type\":\"uint256\"}],\"name\":\"setGlobalMarkup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ccc57490": "GOVERNOR_ROLE()",
		"08c5c0db": "IC_CLIENT_ROLE()",
		"d294f093": "claimFees()",
		"62014bad": "executorEOA()",
		"5d62a8dd": "gasOracle()",
		"c354bd6e": "getClaimableAmount()",
		"4f199114": "getClaimerFraction()",
		"26533fe9": "getClaimerReward()",
		"96fda4da": "getExecutionFee(uint64,uint256,bytes)",
		"4ccb20c0": "getFeeRecipient()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"efd07ec2": "globalMarkup()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"c4d66de8": "initialize(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"58efb47d": "requestTxExecution(uint64,uint256,bytes32,bytes)",
		"d547741f": "revokeRole(bytes32,address)",
		"a9bc769b": "setClaimerFraction(uint256)",
		"2d54566c": "setExecutorEOA(address)",
		"a87b8152": "setGasOracle(address)",
		"cf4f578f": "setGlobalMarkup(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611c81806100df6000396000f3fe60806040526004361061018b5760003560e01c806391d14854116100d6578063c4d66de81161007f578063d294f09311610059578063d294f09314610550578063d547741f14610565578063efd07ec21461058557600080fd5b8063c4d66de8146104dc578063ccc57490146104fc578063cf4f578f1461053057600080fd5b8063a87b8152116100b0578063a87b815214610489578063a9bc769b146104a9578063c354bd6e146104c957600080fd5b806391d14854146103e257806396fda4da14610454578063a217fddf1461047457600080fd5b806336568abe1161013857806358efb47d1161011257806358efb47d1461033b5780635d62a8dd1461034e57806362014bad1461039857600080fd5b806336568abe146102ad5780634ccb20c0146102cd5780634f1991141461030757600080fd5b806326533fe91161016957806326533fe9146102565780632d54566c1461026b5780632f2ff15d1461028d57600080fd5b806301ffc9a71461019057806308c5c0db146101c5578063248a9ca314610207575b600080fd5b34801561019c57600080fd5b506101b06101ab366004611852565b6105b9565b60405190151581526020015b60405180910390f35b3480156101d157600080fd5b506101f97f506033f42d439a89b8dbacb157256b8ef7e613d9e48db1be101b85411778abfb81565b6040519081526020016101bc565b34801561021357600080fd5b506101f9610222366004611894565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561026257600080fd5b506101f9610652565b34801561027757600080fd5b5061028b6102863660046118d6565b610664565b005b34801561029957600080fd5b5061028b6102a83660046118f1565b6107be565b3480156102b957600080fd5b5061028b6102c83660046118f1565b610808565b3480156102d957600080fd5b506102e2610866565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bc565b34801561031357600080fd5b507fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3103546101f9565b61028b61034936600461197e565b6108ab565b34801561035a57600080fd5b507fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31015473ffffffffffffffffffffffffffffffffffffffff166102e2565b3480156103a457600080fd5b507fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31005473ffffffffffffffffffffffffffffffffffffffff166102e2565b3480156103ee57600080fd5b506101b06103fd3660046118f1565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561046057600080fd5b506101f961046f3660046119e6565b610970565b34801561048057600080fd5b506101f9600081565b34801561049557600080fd5b5061028b6104a43660046118d6565b610c5a565b3480156104b557600080fd5b5061028b6104c4366004611894565b610da1565b3480156104d557600080fd5b50476101f9565b3480156104e857600080fd5b5061028b6104f73660046118d6565b610e92565b34801561050857600080fd5b506101f97f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b34801561053c57600080fd5b5061028b61054b366004611894565b611015565b34801561055c57600080fd5b5061028b6110b5565b34801561057157600080fd5b5061028b6105803660046118f1565b6111ce565b34801561059157600080fd5b507fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3102546101f9565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061064c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b60004761065e81611212565b91505090565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5561068e816112b4565b73ffffffffffffffffffffffffffffffffffffffff82166106db576040517f9e3a01ec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117825560408051918252517f4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f719181900360200190a160405173ffffffffffffffffffffffffffffffffffffffff841681527fbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323906020015b60405180910390a1505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546107f8816112b4565b61080283836112c1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610857576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61086182826113e2565b505050565b60006108a67fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31005473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b7f506033f42d439a89b8dbacb157256b8ef7e613d9e48db1be101b85411778abfb6108d5816112b4565b60006108e387878686610970565b90508034101561092d576040517f28c6ec70000000000000000000000000000000000000000000000000000000008152346004820152602481018290526044015b60405180910390fd5b6040805133815234602082015286917fc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c910160405180910390a250505050505050565b6000806109b17fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31015473ffffffffffffffffffffffffffffffffffffffff1690565b905073ffffffffffffffffffffffffffffffffffffffff8116610a00576040517f668604bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610a0c85856114c0565b9050600161ffff82161115610a53576040517f05e98f3a00000000000000000000000000000000000000000000000000000000815261ffff82166004820152602401610924565b6000610a9486868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061150a92505050565b80516040517fbf495c8800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b16600482015260248101919091526044810189905290915073ffffffffffffffffffffffffffffffffffffffff84169063bf495c8890606401602060405180830381865afa158015610b1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b409190611a40565b602082015190945015610bff5760208101516040517f40658a7400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8a166004820152602481019190915273ffffffffffffffffffffffffffffffffffffffff8416906340658a7490604401602060405180830381865afa158015610bce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bf29190611a40565b610bfc9085611a88565b93505b670de0b6b3a7640000610c307fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31025490565b610c3a9086611a9b565b610c449190611ab2565b610c4e9085611a88565b98975050505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610c84816112b4565b8173ffffffffffffffffffffffffffffffffffffffff163b600003610ced576040517fd7c25e1d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610924565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e310180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040519081527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa906020016107b1565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610dcb816112b4565b662386f26fc10000821115610e1c576040517f0ae993dd00000000000000000000000000000000000000000000000000000000815260048101839052662386f26fc100006024820152604401610924565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31038290556040518281527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9906020016107b1565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610edd5750825b905060008267ffffffffffffffff166001148015610efa5750303b155b905081158015610f08575080155b15610f3f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610fa05784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610fab6000876112c1565b50831561100d5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5561103f816112b4565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31028290556040518281527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d906020016107b1565b4760008190036110f1576040517f6e95c0a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006110fb610866565b905073ffffffffffffffffffffffffffffffffffffffff811661114a576040517f3c73eece00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061115583611212565b6040805173ffffffffffffffffffffffffffffffffffffffff851681529482900360208601819052338683015260608601839052905190949192507ff4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b19181900360800190a16111c48284611590565b6108613382611590565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611208816112b4565b61080283836113e2565b60008061123d7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31035490565b9050662386f26fc10000811115611290576040517f0ae993dd00000000000000000000000000000000000000000000000000000000815260048101829052662386f26fc100006024820152604401610924565b670de0b6b3a76400006112a38285611a9b565b6112ad9190611ab2565b9392505050565b6112be8133611666565b50565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166113d85760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556113743390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061064c565b600091505061064c565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156113d85760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061064c565b600060028210156115015782826040517fb0818b62000000000000000000000000000000000000000000000000000000008152600401610924929190611aed565b50503560f01c90565b604080518082019091526000808252602082015260006115298361170d565b9050600161ffff82161015611570576040517f2b346f3700000000000000000000000000000000000000000000000000000000815261ffff82166004820152602401610924565b61157983611758565b8060200190518101906112ad9190611b69565b5050565b804710156115cc576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610924565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114611626576040519150601f19603f3d011682016040523d82523d6000602084013e61162b565b606091505b5050905080610861576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661158c576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610924565b600060028251101561174d57816040517fb0818b620000000000000000000000000000000000000000000000000000000081526004016109249190611bdf565b506020015160f01c90565b606060028251101561179857816040517fb0818b620000000000000000000000000000000000000000000000000000000081526004016109249190611bdf565b81517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe018067ffffffffffffffff8111156117d5576117d5611b3a565b6040519080825280601f01601f1916602001820160405280156117ff576020820181803683370190505b50915060008160208401836022870160045afa90508061184b576040517f101e44fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050919050565b60006020828403121561186457600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146112ad57600080fd5b6000602082840312156118a657600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146118d157600080fd5b919050565b6000602082840312156118e857600080fd5b6112ad826118ad565b6000806040838503121561190457600080fd5b82359150611914602084016118ad565b90509250929050565b803567ffffffffffffffff811681146118d157600080fd5b60008083601f84011261194757600080fd5b50813567ffffffffffffffff81111561195f57600080fd5b60208301915083602082850101111561197757600080fd5b9250929050565b60008060008060006080868803121561199657600080fd5b61199f8661191d565b94506020860135935060408601359250606086013567ffffffffffffffff8111156119c957600080fd5b6119d588828901611935565b969995985093965092949392505050565b600080600080606085870312156119fc57600080fd5b611a058561191d565b935060208501359250604085013567ffffffffffffffff811115611a2857600080fd5b611a3487828801611935565b95989497509550505050565b600060208284031215611a5257600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561064c5761064c611a59565b808202811582820484141761064c5761064c611a59565b600082611ae8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060408284031215611b7b57600080fd5b6040516040810181811067ffffffffffffffff82111715611bc5577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052825181526020928301519281019290925250919050565b600060208083528351808285015260005b81811015611c0c57858101830151858201604001528201611bf0565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509291505056fea2646970667358221220f3b002ae15f7157e92b25dff0688b2db2fd2e13ab3a6d7c94ce030be0b3c70ad64736f6c63430008140033",
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

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GetClaimableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "getClaimableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GetClaimableAmount() (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GetClaimableAmount(&_SynapseExecutionServiceV1.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GetClaimableAmount() (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GetClaimableAmount(&_SynapseExecutionServiceV1.CallOpts)
}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GetClaimerFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "getClaimerFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GetClaimerFraction() (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GetClaimerFraction(&_SynapseExecutionServiceV1.CallOpts)
}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GetClaimerFraction() (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GetClaimerFraction(&_SynapseExecutionServiceV1.CallOpts)
}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GetClaimerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "getClaimerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GetClaimerReward() (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GetClaimerReward(&_SynapseExecutionServiceV1.CallOpts)
}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GetClaimerReward() (*big.Int, error) {
	return _SynapseExecutionServiceV1.Contract.GetClaimerReward(&_SynapseExecutionServiceV1.CallOpts)
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

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Caller) GetFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1.contract.Call(opts, &out, "getFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) GetFeeRecipient() (common.Address, error) {
	return _SynapseExecutionServiceV1.Contract.GetFeeRecipient(&_SynapseExecutionServiceV1.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1CallerSession) GetFeeRecipient() (common.Address, error) {
	return _SynapseExecutionServiceV1.Contract.GetFeeRecipient(&_SynapseExecutionServiceV1.CallOpts)
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

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) ClaimFees() (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.ClaimFees(&_SynapseExecutionServiceV1.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) ClaimFees() (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.ClaimFees(&_SynapseExecutionServiceV1.TransactOpts)
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

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) RequestTxExecution(opts *bind.TransactOpts, dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "requestTxExecution", dstChainId, txPayloadSize, transactionId, options)
}

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) RequestTxExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.RequestTxExecution(&_SynapseExecutionServiceV1.TransactOpts, dstChainId, txPayloadSize, transactionId, options)
}

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) RequestTxExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.RequestTxExecution(&_SynapseExecutionServiceV1.TransactOpts, dstChainId, txPayloadSize, transactionId, options)
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

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Transactor) SetClaimerFraction(opts *bind.TransactOpts, claimerFraction_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.contract.Transact(opts, "setClaimerFraction", claimerFraction_)
}

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Session) SetClaimerFraction(claimerFraction_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SetClaimerFraction(&_SynapseExecutionServiceV1.TransactOpts, claimerFraction_)
}

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction_) returns()
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1TransactorSession) SetClaimerFraction(claimerFraction_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1.Contract.SetClaimerFraction(&_SynapseExecutionServiceV1.TransactOpts, claimerFraction_)
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

// SynapseExecutionServiceV1ClaimerFractionSetIterator is returned from FilterClaimerFractionSet and is used to iterate over the raw logs and unpacked data for ClaimerFractionSet events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1ClaimerFractionSetIterator struct {
	Event *SynapseExecutionServiceV1ClaimerFractionSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1ClaimerFractionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1ClaimerFractionSet)
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
		it.Event = new(SynapseExecutionServiceV1ClaimerFractionSet)
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
func (it *SynapseExecutionServiceV1ClaimerFractionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1ClaimerFractionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1ClaimerFractionSet represents a ClaimerFractionSet event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1ClaimerFractionSet struct {
	ClaimerFraction *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimerFractionSet is a free log retrieval operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterClaimerFractionSet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1ClaimerFractionSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "ClaimerFractionSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1ClaimerFractionSetIterator{contract: _SynapseExecutionServiceV1.contract, event: "ClaimerFractionSet", logs: logs, sub: sub}, nil
}

// WatchClaimerFractionSet is a free log subscription operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchClaimerFractionSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1ClaimerFractionSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "ClaimerFractionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1ClaimerFractionSet)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "ClaimerFractionSet", log); err != nil {
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

// ParseClaimerFractionSet is a log parse operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseClaimerFractionSet(log types.Log) (*SynapseExecutionServiceV1ClaimerFractionSet, error) {
	event := new(SynapseExecutionServiceV1ClaimerFractionSet)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "ClaimerFractionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	ExecutionFee  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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

// WatchExecutionRequested is a free log subscription operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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

// ParseExecutionRequested is a log parse operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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

// SynapseExecutionServiceV1FeeRecipientSetIterator is returned from FilterFeeRecipientSet and is used to iterate over the raw logs and unpacked data for FeeRecipientSet events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1FeeRecipientSetIterator struct {
	Event *SynapseExecutionServiceV1FeeRecipientSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1FeeRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1FeeRecipientSet)
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
		it.Event = new(SynapseExecutionServiceV1FeeRecipientSet)
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
func (it *SynapseExecutionServiceV1FeeRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1FeeRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1FeeRecipientSet represents a FeeRecipientSet event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1FeeRecipientSet struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientSet is a free log retrieval operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterFeeRecipientSet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1FeeRecipientSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1FeeRecipientSetIterator{contract: _SynapseExecutionServiceV1.contract, event: "FeeRecipientSet", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientSet is a free log subscription operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchFeeRecipientSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1FeeRecipientSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1FeeRecipientSet)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
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

// ParseFeeRecipientSet is a log parse operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseFeeRecipientSet(log types.Log) (*SynapseExecutionServiceV1FeeRecipientSet, error) {
	event := new(SynapseExecutionServiceV1FeeRecipientSet)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1FeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1FeesClaimedIterator struct {
	Event *SynapseExecutionServiceV1FeesClaimed // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1FeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1FeesClaimed)
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
		it.Event = new(SynapseExecutionServiceV1FeesClaimed)
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
func (it *SynapseExecutionServiceV1FeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1FeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1FeesClaimed represents a FeesClaimed event raised by the SynapseExecutionServiceV1 contract.
type SynapseExecutionServiceV1FeesClaimed struct {
	FeeRecipient  common.Address
	ClaimedFees   *big.Int
	Claimer       common.Address
	ClaimerReward *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) FilterFeesClaimed(opts *bind.FilterOpts) (*SynapseExecutionServiceV1FeesClaimedIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.FilterLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1FeesClaimedIterator{contract: _SynapseExecutionServiceV1.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1FeesClaimed) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1.contract.WatchLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1FeesClaimed)
				if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_SynapseExecutionServiceV1 *SynapseExecutionServiceV1Filterer) ParseFeesClaimed(log types.Log) (*SynapseExecutionServiceV1FeesClaimed, error) {
	event := new(SynapseExecutionServiceV1FeesClaimed)
	if err := _SynapseExecutionServiceV1.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"name\":\"ClaimableFees__ClaimerFractionAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableFees__FeeAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableFees__FeeRecipientZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"OptionsLib__VersionInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__ExecutorZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRequired\",\"type\":\"uint256\"}],\"name\":\"SynapseExecutionService__FeeAmountBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"SynapseExecutionService__GasOracleNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseExecutionService__GasOracleZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"SynapseExecutionService__OptionsVersionNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__PayloadTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerFraction\",\"type\":\"uint256\"}],\"name\":\"ClaimerFractionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executorEOA\",\"type\":\"address\"}],\"name\":\"ExecutorEOASet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerReward\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalMarkup\",\"type\":\"uint256\"}],\"name\":\"GlobalMarkupSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IC_CLIENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimerFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalMarkup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestTxExecution\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimerFraction_\",\"type\":\"uint256\"}],\"name\":\"setClaimerFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executorEOA_\",\"type\":\"address\"}],\"name\":\"setExecutorEOA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalMarkup_\",\"type\":\"uint256\"}],\"name\":\"setGlobalMarkup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ccc57490": "GOVERNOR_ROLE()",
		"08c5c0db": "IC_CLIENT_ROLE()",
		"d294f093": "claimFees()",
		"62014bad": "executorEOA()",
		"5d62a8dd": "gasOracle()",
		"c354bd6e": "getClaimableAmount()",
		"4f199114": "getClaimerFraction()",
		"26533fe9": "getClaimerReward()",
		"96fda4da": "getExecutionFee(uint64,uint256,bytes)",
		"4ccb20c0": "getFeeRecipient()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"efd07ec2": "globalMarkup()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"c4d66de8": "initialize(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"58efb47d": "requestTxExecution(uint64,uint256,bytes32,bytes)",
		"d547741f": "revokeRole(bytes32,address)",
		"a9bc769b": "setClaimerFraction(uint256)",
		"2d54566c": "setExecutorEOA(address)",
		"a87b8152": "setGasOracle(address)",
		"cf4f578f": "setGlobalMarkup(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506200001f60003362000053565b506200004c7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f553362000053565b5062000127565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166200011b576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055620000d03390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505062000121565b60009150505b92915050565b611c8180620001376000396000f3fe60806040526004361061018b5760003560e01c806391d14854116100d6578063c4d66de81161007f578063d294f09311610059578063d294f09314610550578063d547741f14610565578063efd07ec21461058557600080fd5b8063c4d66de8146104dc578063ccc57490146104fc578063cf4f578f1461053057600080fd5b8063a87b8152116100b0578063a87b815214610489578063a9bc769b146104a9578063c354bd6e146104c957600080fd5b806391d14854146103e257806396fda4da14610454578063a217fddf1461047457600080fd5b806336568abe1161013857806358efb47d1161011257806358efb47d1461033b5780635d62a8dd1461034e57806362014bad1461039857600080fd5b806336568abe146102ad5780634ccb20c0146102cd5780634f1991141461030757600080fd5b806326533fe91161016957806326533fe9146102565780632d54566c1461026b5780632f2ff15d1461028d57600080fd5b806301ffc9a71461019057806308c5c0db146101c5578063248a9ca314610207575b600080fd5b34801561019c57600080fd5b506101b06101ab366004611852565b6105b9565b60405190151581526020015b60405180910390f35b3480156101d157600080fd5b506101f97f506033f42d439a89b8dbacb157256b8ef7e613d9e48db1be101b85411778abfb81565b6040519081526020016101bc565b34801561021357600080fd5b506101f9610222366004611894565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561026257600080fd5b506101f9610652565b34801561027757600080fd5b5061028b6102863660046118d6565b610664565b005b34801561029957600080fd5b5061028b6102a83660046118f1565b6107be565b3480156102b957600080fd5b5061028b6102c83660046118f1565b610808565b3480156102d957600080fd5b506102e2610866565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bc565b34801561031357600080fd5b507fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3103546101f9565b61028b61034936600461197e565b6108ab565b34801561035a57600080fd5b507fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31015473ffffffffffffffffffffffffffffffffffffffff166102e2565b3480156103a457600080fd5b507fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31005473ffffffffffffffffffffffffffffffffffffffff166102e2565b3480156103ee57600080fd5b506101b06103fd3660046118f1565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561046057600080fd5b506101f961046f3660046119e6565b610970565b34801561048057600080fd5b506101f9600081565b34801561049557600080fd5b5061028b6104a43660046118d6565b610c5a565b3480156104b557600080fd5b5061028b6104c4366004611894565b610da1565b3480156104d557600080fd5b50476101f9565b3480156104e857600080fd5b5061028b6104f73660046118d6565b610e92565b34801561050857600080fd5b506101f97f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b34801561053c57600080fd5b5061028b61054b366004611894565b611015565b34801561055c57600080fd5b5061028b6110b5565b34801561057157600080fd5b5061028b6105803660046118f1565b6111ce565b34801561059157600080fd5b507fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3102546101f9565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061064c57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b60004761065e81611212565b91505090565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5561068e816112b4565b73ffffffffffffffffffffffffffffffffffffffff82166106db576040517f9e3a01ec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117825560408051918252517f4ab11d24f4bb323219ce90846ba579a556c914e8587517e7c8c4264771cd9f719181900360200190a160405173ffffffffffffffffffffffffffffffffffffffff841681527fbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323906020015b60405180910390a1505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546107f8816112b4565b61080283836112c1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610857576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61086182826113e2565b505050565b60006108a67fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31005473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b7f506033f42d439a89b8dbacb157256b8ef7e613d9e48db1be101b85411778abfb6108d5816112b4565b60006108e387878686610970565b90508034101561092d576040517f28c6ec70000000000000000000000000000000000000000000000000000000008152346004820152602481018290526044015b60405180910390fd5b6040805133815234602082015286917fc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c910160405180910390a250505050505050565b6000806109b17fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31015473ffffffffffffffffffffffffffffffffffffffff1690565b905073ffffffffffffffffffffffffffffffffffffffff8116610a00576040517f668604bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610a0c85856114c0565b9050600161ffff82161115610a53576040517f05e98f3a00000000000000000000000000000000000000000000000000000000815261ffff82166004820152602401610924565b6000610a9486868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061150a92505050565b80516040517fbf495c8800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b16600482015260248101919091526044810189905290915073ffffffffffffffffffffffffffffffffffffffff84169063bf495c8890606401602060405180830381865afa158015610b1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b409190611a40565b602082015190945015610bff5760208101516040517f40658a7400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8a166004820152602481019190915273ffffffffffffffffffffffffffffffffffffffff8416906340658a7490604401602060405180830381865afa158015610bce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bf29190611a40565b610bfc9085611a88565b93505b670de0b6b3a7640000610c307fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31025490565b610c3a9086611a9b565b610c449190611ab2565b610c4e9085611a88565b98975050505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610c84816112b4565b8173ffffffffffffffffffffffffffffffffffffffff163b600003610ced576040517fd7c25e1d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610924565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e310180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040519081527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f3efbbb00c39812fb98647af6e9e2c3f4ec2b53d368cedd1e148330a05b652cfa906020016107b1565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610dcb816112b4565b662386f26fc10000821115610e1c576040517f0ae993dd00000000000000000000000000000000000000000000000000000000815260048101839052662386f26fc100006024820152604401610924565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31038290556040518281527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9906020016107b1565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610edd5750825b905060008267ffffffffffffffff166001148015610efa5750303b155b905081158015610f08575080155b15610f3f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610fa05784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610fab6000876112c1565b50831561100d5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5561103f816112b4565b7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31028290556040518281527fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e3100907f1957a4f563f2f13a7e7c1f9d8d6e719a1e6f687ac787704c33069f0a7997d75d906020016107b1565b4760008190036110f1576040517f6e95c0a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006110fb610866565b905073ffffffffffffffffffffffffffffffffffffffff811661114a576040517f3c73eece00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061115583611212565b6040805173ffffffffffffffffffffffffffffffffffffffff851681529482900360208601819052338683015260608601839052905190949192507ff4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b19181900360800190a16111c48284611590565b6108613382611590565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611208816112b4565b61080283836113e2565b60008061123d7fabc861e0f8da03757893d41bb54770e6953c799ce2884f80d6b14b66ba8e31035490565b9050662386f26fc10000811115611290576040517f0ae993dd00000000000000000000000000000000000000000000000000000000815260048101829052662386f26fc100006024820152604401610924565b670de0b6b3a76400006112a38285611a9b565b6112ad9190611ab2565b9392505050565b6112be8133611666565b50565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff166113d85760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556113743390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061064c565b600091505061064c565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16156113d85760008481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061064c565b600060028210156115015782826040517fb0818b62000000000000000000000000000000000000000000000000000000008152600401610924929190611aed565b50503560f01c90565b604080518082019091526000808252602082015260006115298361170d565b9050600161ffff82161015611570576040517f2b346f3700000000000000000000000000000000000000000000000000000000815261ffff82166004820152602401610924565b61157983611758565b8060200190518101906112ad9190611b69565b5050565b804710156115cc576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610924565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114611626576040519150601f19603f3d011682016040523d82523d6000602084013e61162b565b606091505b5050905080610861576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661158c576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610924565b600060028251101561174d57816040517fb0818b620000000000000000000000000000000000000000000000000000000081526004016109249190611bdf565b506020015160f01c90565b606060028251101561179857816040517fb0818b620000000000000000000000000000000000000000000000000000000081526004016109249190611bdf565b81517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe018067ffffffffffffffff8111156117d5576117d5611b3a565b6040519080825280601f01601f1916602001820160405280156117ff576020820181803683370190505b50915060008160208401836022870160045afa90508061184b576040517f101e44fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050919050565b60006020828403121561186457600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146112ad57600080fd5b6000602082840312156118a657600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146118d157600080fd5b919050565b6000602082840312156118e857600080fd5b6112ad826118ad565b6000806040838503121561190457600080fd5b82359150611914602084016118ad565b90509250929050565b803567ffffffffffffffff811681146118d157600080fd5b60008083601f84011261194757600080fd5b50813567ffffffffffffffff81111561195f57600080fd5b60208301915083602082850101111561197757600080fd5b9250929050565b60008060008060006080868803121561199657600080fd5b61199f8661191d565b94506020860135935060408601359250606086013567ffffffffffffffff8111156119c957600080fd5b6119d588828901611935565b969995985093965092949392505050565b600080600080606085870312156119fc57600080fd5b611a058561191d565b935060208501359250604085013567ffffffffffffffff811115611a2857600080fd5b611a3487828801611935565b95989497509550505050565b600060208284031215611a5257600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561064c5761064c611a59565b808202811582820484141761064c5761064c611a59565b600082611ae8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060408284031215611b7b57600080fd5b6040516040810181811067ffffffffffffffff82111715611bc5577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052825181526020928301519281019290925250919050565b600060208083528351808285015260005b81811015611c0c57858101830151858201604001528201611bf0565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509291505056fea2646970667358221220acc2a0fe7a2421f9293a36bd64e53e90660f4fc44511dc190483da941e09590764736f6c63430008140033",
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

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GetClaimableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "getClaimableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GetClaimableAmount() (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetClaimableAmount(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xc354bd6e.
//
// Solidity: function getClaimableAmount() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GetClaimableAmount() (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetClaimableAmount(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GetClaimerFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "getClaimerFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GetClaimerFraction() (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetClaimerFraction(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GetClaimerFraction is a free data retrieval call binding the contract method 0x4f199114.
//
// Solidity: function getClaimerFraction() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GetClaimerFraction() (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetClaimerFraction(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GetClaimerReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "getClaimerReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GetClaimerReward() (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetClaimerReward(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GetClaimerReward is a free data retrieval call binding the contract method 0x26533fe9.
//
// Solidity: function getClaimerReward() view returns(uint256)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GetClaimerReward() (*big.Int, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetClaimerReward(&_SynapseExecutionServiceV1Harness.CallOpts)
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

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCaller) GetFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseExecutionServiceV1Harness.contract.Call(opts, &out, "getFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) GetFeeRecipient() (common.Address, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetFeeRecipient(&_SynapseExecutionServiceV1Harness.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessCallerSession) GetFeeRecipient() (common.Address, error) {
	return _SynapseExecutionServiceV1Harness.Contract.GetFeeRecipient(&_SynapseExecutionServiceV1Harness.CallOpts)
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

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) ClaimFees() (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.ClaimFees(&_SynapseExecutionServiceV1Harness.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.ClaimFees(&_SynapseExecutionServiceV1Harness.TransactOpts)
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

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) RequestTxExecution(opts *bind.TransactOpts, dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "requestTxExecution", dstChainId, txPayloadSize, transactionId, options)
}

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) RequestTxExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.RequestTxExecution(&_SynapseExecutionServiceV1Harness.TransactOpts, dstChainId, txPayloadSize, transactionId, options)
}

// RequestTxExecution is a paid mutator transaction binding the contract method 0x58efb47d.
//
// Solidity: function requestTxExecution(uint64 dstChainId, uint256 txPayloadSize, bytes32 transactionId, bytes options) payable returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) RequestTxExecution(dstChainId uint64, txPayloadSize *big.Int, transactionId [32]byte, options []byte) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.RequestTxExecution(&_SynapseExecutionServiceV1Harness.TransactOpts, dstChainId, txPayloadSize, transactionId, options)
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

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactor) SetClaimerFraction(opts *bind.TransactOpts, claimerFraction_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.contract.Transact(opts, "setClaimerFraction", claimerFraction_)
}

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessSession) SetClaimerFraction(claimerFraction_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SetClaimerFraction(&_SynapseExecutionServiceV1Harness.TransactOpts, claimerFraction_)
}

// SetClaimerFraction is a paid mutator transaction binding the contract method 0xa9bc769b.
//
// Solidity: function setClaimerFraction(uint256 claimerFraction_) returns()
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessTransactorSession) SetClaimerFraction(claimerFraction_ *big.Int) (*types.Transaction, error) {
	return _SynapseExecutionServiceV1Harness.Contract.SetClaimerFraction(&_SynapseExecutionServiceV1Harness.TransactOpts, claimerFraction_)
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

// SynapseExecutionServiceV1HarnessClaimerFractionSetIterator is returned from FilterClaimerFractionSet and is used to iterate over the raw logs and unpacked data for ClaimerFractionSet events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessClaimerFractionSetIterator struct {
	Event *SynapseExecutionServiceV1HarnessClaimerFractionSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessClaimerFractionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessClaimerFractionSet)
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
		it.Event = new(SynapseExecutionServiceV1HarnessClaimerFractionSet)
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
func (it *SynapseExecutionServiceV1HarnessClaimerFractionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessClaimerFractionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessClaimerFractionSet represents a ClaimerFractionSet event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessClaimerFractionSet struct {
	ClaimerFraction *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimerFractionSet is a free log retrieval operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterClaimerFractionSet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1HarnessClaimerFractionSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "ClaimerFractionSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessClaimerFractionSetIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "ClaimerFractionSet", logs: logs, sub: sub}, nil
}

// WatchClaimerFractionSet is a free log subscription operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchClaimerFractionSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessClaimerFractionSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "ClaimerFractionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessClaimerFractionSet)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "ClaimerFractionSet", log); err != nil {
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

// ParseClaimerFractionSet is a log parse operation binding the contract event 0x2b76ed3837bd14c860020e473bce45e560d5bca9b5109ef2f08b2051d1cf6cc9.
//
// Solidity: event ClaimerFractionSet(uint256 claimerFraction)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseClaimerFractionSet(log types.Log) (*SynapseExecutionServiceV1HarnessClaimerFractionSet, error) {
	event := new(SynapseExecutionServiceV1HarnessClaimerFractionSet)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "ClaimerFractionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	ExecutionFee  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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

// WatchExecutionRequested is a free log subscription operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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

// ParseExecutionRequested is a log parse operation binding the contract event 0xc3afeef9d037dcadfc927cf1a2c10a5dccba06a26bac58e0e2adf916407f2a7c.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client, uint256 executionFee)
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

// SynapseExecutionServiceV1HarnessFeeRecipientSetIterator is returned from FilterFeeRecipientSet and is used to iterate over the raw logs and unpacked data for FeeRecipientSet events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessFeeRecipientSetIterator struct {
	Event *SynapseExecutionServiceV1HarnessFeeRecipientSet // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessFeeRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessFeeRecipientSet)
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
		it.Event = new(SynapseExecutionServiceV1HarnessFeeRecipientSet)
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
func (it *SynapseExecutionServiceV1HarnessFeeRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessFeeRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessFeeRecipientSet represents a FeeRecipientSet event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessFeeRecipientSet struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientSet is a free log retrieval operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterFeeRecipientSet(opts *bind.FilterOpts) (*SynapseExecutionServiceV1HarnessFeeRecipientSetIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessFeeRecipientSetIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "FeeRecipientSet", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientSet is a free log subscription operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchFeeRecipientSet(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessFeeRecipientSet) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessFeeRecipientSet)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
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

// ParseFeeRecipientSet is a log parse operation binding the contract event 0xbf9a9534339a9d6b81696e05dcfb614b7dc518a31d48be3cfb757988381fb323.
//
// Solidity: event FeeRecipientSet(address feeRecipient)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseFeeRecipientSet(log types.Log) (*SynapseExecutionServiceV1HarnessFeeRecipientSet, error) {
	event := new(SynapseExecutionServiceV1HarnessFeeRecipientSet)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseExecutionServiceV1HarnessFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessFeesClaimedIterator struct {
	Event *SynapseExecutionServiceV1HarnessFeesClaimed // Event containing the contract specifics and raw log

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
func (it *SynapseExecutionServiceV1HarnessFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseExecutionServiceV1HarnessFeesClaimed)
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
		it.Event = new(SynapseExecutionServiceV1HarnessFeesClaimed)
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
func (it *SynapseExecutionServiceV1HarnessFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseExecutionServiceV1HarnessFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseExecutionServiceV1HarnessFeesClaimed represents a FeesClaimed event raised by the SynapseExecutionServiceV1Harness contract.
type SynapseExecutionServiceV1HarnessFeesClaimed struct {
	FeeRecipient  common.Address
	ClaimedFees   *big.Int
	Claimer       common.Address
	ClaimerReward *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) FilterFeesClaimed(opts *bind.FilterOpts) (*SynapseExecutionServiceV1HarnessFeesClaimedIterator, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.FilterLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return &SynapseExecutionServiceV1HarnessFeesClaimedIterator{contract: _SynapseExecutionServiceV1Harness.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *SynapseExecutionServiceV1HarnessFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _SynapseExecutionServiceV1Harness.contract.WatchLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseExecutionServiceV1HarnessFeesClaimed)
				if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeRecipient, uint256 claimedFees, address claimer, uint256 claimerReward)
func (_SynapseExecutionServiceV1Harness *SynapseExecutionServiceV1HarnessFilterer) ParseFeesClaimed(log types.Log) (*SynapseExecutionServiceV1HarnessFeesClaimed, error) {
	event := new(SynapseExecutionServiceV1HarnessFeesClaimed)
	if err := _SynapseExecutionServiceV1Harness.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__PayloadTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205eb466bd41aa7b3b5a63b41d48a30506484be9f3e77d3812cc55bd6c4214620164736f6c63430008140033",
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
