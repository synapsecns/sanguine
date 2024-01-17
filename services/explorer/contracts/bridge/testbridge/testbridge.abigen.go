// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testbridge

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
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

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControlUpgradeable.Contract.GetRoleMember(&_AccessControlUpgradeable.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControlUpgradeable.Contract.GetRoleMember(&_AccessControlUpgradeable.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControlUpgradeable.Contract.GetRoleMemberCount(&_AccessControlUpgradeable.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControlUpgradeable.Contract.GetRoleMemberCount(&_AccessControlUpgradeable.CallOpts, role)
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
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RenounceRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RenounceRole(&_AccessControlUpgradeable.TransactOpts, role, account)
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
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122017badb84445ebf628962405f5627fcbf36fcbb90c2f77f8e76257d61aef228e864736f6c634300060c0033",
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

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bdeb40f1dae52d38f5e88098b9f175f446b8da3780dfa8204722e9ce4724c19364736f6c634300060c0033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
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

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e8738038062000e87833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250508251620001b491506003906020850190620001e0565b508051620001ca906004906020840190620001e0565b50506005805460ff19166012179055506200027c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022357805160ff191683800117855562000253565b8280016001018555821562000253579182015b828111156200025357825182559160200191906001019062000236565b506200026192915062000265565b5090565b5b8082111562000261576000815560010162000266565b610bfb806200028c6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610287578063a9059cbb146102c0578063dd62ed3e146102f9576100c9565b8063395093511461021357806370a082311461024c57806395d89b411461027f576100c9565b806318160ddd116100b257806318160ddd1461019857806323b872dd146101b2578063313ce567146101f5576100c9565b806306fdde03146100ce578063095ea7b31461014b575b600080fd5b6100d6610334565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101105781810151838201526020016100f8565b50505050905090810190601f16801561013d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101846004803603604081101561016157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356103e8565b604080519115158252519081900360200190f35b6101a0610405565b60408051918252519081900360200190f35b610184600480360360608110156101c857600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135909116906040013561040b565b6101fd6104ac565b6040805160ff9092168252519081900360200190f35b6101846004803603604081101561022957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356104b5565b6101a06004803603602081101561026257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610510565b6100d6610538565b6101846004803603604081101561029d57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356105b7565b610184600480360360408110156102d657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020013561062c565b6101a06004803603604081101561030f57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610640565b60038054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103de5780601f106103b3576101008083540402835291602001916103de565b820191906000526020600020905b8154815290600101906020018083116103c157829003601f168201915b5050505050905090565b60006103fc6103f5610678565b848461067c565b50600192915050565b60025490565b60006104188484846107c3565b6104a284610424610678565b61049d85604051806060016040528060288152602001610b306028913973ffffffffffffffffffffffffffffffffffffffff8a1660009081526001602052604081209061046f610678565b73ffffffffffffffffffffffffffffffffffffffff1681526020810191909152604001600020549190610993565b61067c565b5060019392505050565b60055460ff1690565b60006103fc6104c2610678565b8461049d85600160006104d3610678565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918c168152925290205490610a44565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b60048054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103de5780601f106103b3576101008083540402835291602001916103de565b60006103fc6105c4610678565b8461049d85604051806060016040528060258152602001610ba160259139600160006105ee610678565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918d16815292529020549190610993565b60006103fc610639610678565b84846107c3565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b3390565b73ffffffffffffffffffffffffffffffffffffffff83166106e8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610b7d6024913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610754576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610ae86022913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661082f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610b586025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661089b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180610ac56023913960400191505060405180910390fd5b6108a6838383610abf565b6108f081604051806060016040528060268152602001610b0a6026913973ffffffffffffffffffffffffffffffffffffffff86166000908152602081905260409020549190610993565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220939093559084168152205461092c9082610a44565b73ffffffffffffffffffffffffffffffffffffffff8084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610a3c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a015781810151838201526020016109e9565b50505050905090810190601f168015610a2e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610ab857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220040bfb3441da59f33010ff945b3a9918d544b481e7c793a4db412ec0e87854e564736f6c634300060c0033",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// Deprecated: Use ERC20MetaData.Sigs instead.
// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = ERC20MetaData.Sigs

// ERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MetaData.Bin instead.
var ERC20Bin = ERC20MetaData.Bin

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableMetaData contains all meta data concerning the ERC20Burnable contract.
var ERC20BurnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"42966c68": "burn(uint256)",
		"79cc6790": "burnFrom(address,uint256)",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// ERC20BurnableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BurnableMetaData.ABI instead.
var ERC20BurnableABI = ERC20BurnableMetaData.ABI

// Deprecated: Use ERC20BurnableMetaData.Sigs instead.
// ERC20BurnableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BurnableFuncSigs = ERC20BurnableMetaData.Sigs

// ERC20Burnable is an auto generated Go binding around an Ethereum contract.
type ERC20Burnable struct {
	ERC20BurnableCaller     // Read-only binding to the contract
	ERC20BurnableTransactor // Write-only binding to the contract
	ERC20BurnableFilterer   // Log filterer for contract events
}

// ERC20BurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableSession struct {
	Contract     *ERC20Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableCallerSession struct {
	Contract *ERC20BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20BurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableTransactorSession struct {
	Contract     *ERC20BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20BurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableRaw struct {
	Contract *ERC20Burnable // Generic contract binding to access the raw methods on
}

// ERC20BurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableCallerRaw struct {
	Contract *ERC20BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactorRaw struct {
	Contract *ERC20BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Burnable creates a new instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20Burnable(address common.Address, backend bind.ContractBackend) (*ERC20Burnable, error) {
	contract, err := bindERC20Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// NewERC20BurnableCaller creates a new read-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableCaller, error) {
	contract, err := bindERC20Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableCaller{contract: contract}, nil
}

// NewERC20BurnableTransactor creates a new write-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableTransactor, error) {
	contract, err := bindERC20Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransactor{contract: contract}, nil
}

// NewERC20BurnableFilterer creates a new log filterer instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableFilterer, error) {
	contract, err := bindERC20Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableFilterer{contract: contract}, nil
}

// bindERC20Burnable binds a generic wrapper to an already deployed contract.
func bindERC20Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20BurnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.ERC20BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCallerSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// ERC20BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Burnable contract.
type ERC20BurnableApprovalIterator struct {
	Event *ERC20BurnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableApproval)
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
		it.Event = new(ERC20BurnableApproval)
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
func (it *ERC20BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableApproval represents a Approval event raised by the ERC20Burnable contract.
type ERC20BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableApprovalIterator{contract: _ERC20Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableApproval)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseApproval(log types.Log) (*ERC20BurnableApproval, error) {
	event := new(ERC20BurnableApproval)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Burnable contract.
type ERC20BurnableTransferIterator struct {
	Event *ERC20BurnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableTransfer)
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
		it.Event = new(ERC20BurnableTransfer)
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
func (it *ERC20BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableTransfer represents a Transfer event raised by the ERC20Burnable contract.
type ERC20BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransferIterator{contract: _ERC20Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableTransfer)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableTransfer, error) {
	event := new(ERC20BurnableTransfer)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnumerableSetUpgradeableMetaData contains all meta data concerning the EnumerableSetUpgradeable contract.
var EnumerableSetUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122043ac4cddaafd61c05eaf68f43278758f58160304c80b485c19be216c65ebad3d64736f6c634300060c0033",
}

// EnumerableSetUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetUpgradeableMetaData.ABI instead.
var EnumerableSetUpgradeableABI = EnumerableSetUpgradeableMetaData.ABI

// EnumerableSetUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetUpgradeableMetaData.Bin instead.
var EnumerableSetUpgradeableBin = EnumerableSetUpgradeableMetaData.Bin

// DeployEnumerableSetUpgradeable deploys a new Ethereum contract, binding an instance of EnumerableSetUpgradeable to it.
func DeployEnumerableSetUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSetUpgradeable, error) {
	parsed, err := EnumerableSetUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSetUpgradeable{EnumerableSetUpgradeableCaller: EnumerableSetUpgradeableCaller{contract: contract}, EnumerableSetUpgradeableTransactor: EnumerableSetUpgradeableTransactor{contract: contract}, EnumerableSetUpgradeableFilterer: EnumerableSetUpgradeableFilterer{contract: contract}}, nil
}

// EnumerableSetUpgradeable is an auto generated Go binding around an Ethereum contract.
type EnumerableSetUpgradeable struct {
	EnumerableSetUpgradeableCaller     // Read-only binding to the contract
	EnumerableSetUpgradeableTransactor // Write-only binding to the contract
	EnumerableSetUpgradeableFilterer   // Log filterer for contract events
}

// EnumerableSetUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetUpgradeableSession struct {
	Contract     *EnumerableSetUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EnumerableSetUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetUpgradeableCallerSession struct {
	Contract *EnumerableSetUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// EnumerableSetUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetUpgradeableTransactorSession struct {
	Contract     *EnumerableSetUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// EnumerableSetUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetUpgradeableRaw struct {
	Contract *EnumerableSetUpgradeable // Generic contract binding to access the raw methods on
}

// EnumerableSetUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetUpgradeableCallerRaw struct {
	Contract *EnumerableSetUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetUpgradeableTransactorRaw struct {
	Contract *EnumerableSetUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSetUpgradeable creates a new instance of EnumerableSetUpgradeable, bound to a specific deployed contract.
func NewEnumerableSetUpgradeable(address common.Address, backend bind.ContractBackend) (*EnumerableSetUpgradeable, error) {
	contract, err := bindEnumerableSetUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetUpgradeable{EnumerableSetUpgradeableCaller: EnumerableSetUpgradeableCaller{contract: contract}, EnumerableSetUpgradeableTransactor: EnumerableSetUpgradeableTransactor{contract: contract}, EnumerableSetUpgradeableFilterer: EnumerableSetUpgradeableFilterer{contract: contract}}, nil
}

// NewEnumerableSetUpgradeableCaller creates a new read-only instance of EnumerableSetUpgradeable, bound to a specific deployed contract.
func NewEnumerableSetUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetUpgradeableCaller, error) {
	contract, err := bindEnumerableSetUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetUpgradeableCaller{contract: contract}, nil
}

// NewEnumerableSetUpgradeableTransactor creates a new write-only instance of EnumerableSetUpgradeable, bound to a specific deployed contract.
func NewEnumerableSetUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetUpgradeableTransactor, error) {
	contract, err := bindEnumerableSetUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetUpgradeableTransactor{contract: contract}, nil
}

// NewEnumerableSetUpgradeableFilterer creates a new log filterer instance of EnumerableSetUpgradeable, bound to a specific deployed contract.
func NewEnumerableSetUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetUpgradeableFilterer, error) {
	contract, err := bindEnumerableSetUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetUpgradeableFilterer{contract: contract}, nil
}

// bindEnumerableSetUpgradeable binds a generic wrapper to an already deployed contract.
func bindEnumerableSetUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableSetUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSetUpgradeable *EnumerableSetUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSetUpgradeable.Contract.EnumerableSetUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSetUpgradeable *EnumerableSetUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSetUpgradeable.Contract.EnumerableSetUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSetUpgradeable *EnumerableSetUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSetUpgradeable.Contract.EnumerableSetUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSetUpgradeable *EnumerableSetUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSetUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSetUpgradeable *EnumerableSetUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSetUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSetUpgradeable *EnumerableSetUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSetUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MintableMetaData contains all meta data concerning the IERC20Mintable contract.
var IERC20MintableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"40c10f19": "mint(address,uint256)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20MintableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MintableMetaData.ABI instead.
var IERC20MintableABI = IERC20MintableMetaData.ABI

// Deprecated: Use IERC20MintableMetaData.Sigs instead.
// IERC20MintableFuncSigs maps the 4-byte function signature to its string representation.
var IERC20MintableFuncSigs = IERC20MintableMetaData.Sigs

// IERC20Mintable is an auto generated Go binding around an Ethereum contract.
type IERC20Mintable struct {
	IERC20MintableCaller     // Read-only binding to the contract
	IERC20MintableTransactor // Write-only binding to the contract
	IERC20MintableFilterer   // Log filterer for contract events
}

// IERC20MintableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MintableSession struct {
	Contract     *IERC20Mintable   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MintableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MintableCallerSession struct {
	Contract *IERC20MintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MintableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MintableTransactorSession struct {
	Contract     *IERC20MintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MintableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MintableRaw struct {
	Contract *IERC20Mintable // Generic contract binding to access the raw methods on
}

// IERC20MintableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MintableCallerRaw struct {
	Contract *IERC20MintableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MintableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MintableTransactorRaw struct {
	Contract *IERC20MintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Mintable creates a new instance of IERC20Mintable, bound to a specific deployed contract.
func NewIERC20Mintable(address common.Address, backend bind.ContractBackend) (*IERC20Mintable, error) {
	contract, err := bindIERC20Mintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Mintable{IERC20MintableCaller: IERC20MintableCaller{contract: contract}, IERC20MintableTransactor: IERC20MintableTransactor{contract: contract}, IERC20MintableFilterer: IERC20MintableFilterer{contract: contract}}, nil
}

// NewIERC20MintableCaller creates a new read-only instance of IERC20Mintable, bound to a specific deployed contract.
func NewIERC20MintableCaller(address common.Address, caller bind.ContractCaller) (*IERC20MintableCaller, error) {
	contract, err := bindIERC20Mintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableCaller{contract: contract}, nil
}

// NewIERC20MintableTransactor creates a new write-only instance of IERC20Mintable, bound to a specific deployed contract.
func NewIERC20MintableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MintableTransactor, error) {
	contract, err := bindIERC20Mintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableTransactor{contract: contract}, nil
}

// NewIERC20MintableFilterer creates a new log filterer instance of IERC20Mintable, bound to a specific deployed contract.
func NewIERC20MintableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MintableFilterer, error) {
	contract, err := bindIERC20Mintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableFilterer{contract: contract}, nil
}

// bindIERC20Mintable binds a generic wrapper to an already deployed contract.
func bindIERC20Mintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MintableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Mintable *IERC20MintableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Mintable.Contract.IERC20MintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Mintable *IERC20MintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.IERC20MintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Mintable *IERC20MintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.IERC20MintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Mintable *IERC20MintableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Mintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Mintable *IERC20MintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Mintable *IERC20MintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Mintable *IERC20MintableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Mintable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Mintable *IERC20MintableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Mintable.Contract.Allowance(&_IERC20Mintable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Mintable *IERC20MintableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Mintable.Contract.Allowance(&_IERC20Mintable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Mintable *IERC20MintableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Mintable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Mintable *IERC20MintableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Mintable.Contract.BalanceOf(&_IERC20Mintable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Mintable *IERC20MintableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Mintable.Contract.BalanceOf(&_IERC20Mintable.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Mintable *IERC20MintableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Mintable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Mintable *IERC20MintableSession) TotalSupply() (*big.Int, error) {
	return _IERC20Mintable.Contract.TotalSupply(&_IERC20Mintable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Mintable *IERC20MintableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Mintable.Contract.TotalSupply(&_IERC20Mintable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.Approve(&_IERC20Mintable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.Approve(&_IERC20Mintable.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Mintable *IERC20MintableTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Mintable *IERC20MintableSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.Mint(&_IERC20Mintable.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Mintable *IERC20MintableTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.Mint(&_IERC20Mintable.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.Transfer(&_IERC20Mintable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.Transfer(&_IERC20Mintable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.TransferFrom(&_IERC20Mintable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Mintable *IERC20MintableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Mintable.Contract.TransferFrom(&_IERC20Mintable.TransactOpts, sender, recipient, amount)
}

// IERC20MintableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Mintable contract.
type IERC20MintableApprovalIterator struct {
	Event *IERC20MintableApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MintableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MintableApproval)
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
		it.Event = new(IERC20MintableApproval)
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
func (it *IERC20MintableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MintableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MintableApproval represents a Approval event raised by the IERC20Mintable contract.
type IERC20MintableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Mintable *IERC20MintableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MintableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Mintable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableApprovalIterator{contract: _IERC20Mintable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Mintable *IERC20MintableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MintableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Mintable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MintableApproval)
				if err := _IERC20Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Mintable *IERC20MintableFilterer) ParseApproval(log types.Log) (*IERC20MintableApproval, error) {
	event := new(IERC20MintableApproval)
	if err := _IERC20Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MintableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Mintable contract.
type IERC20MintableTransferIterator struct {
	Event *IERC20MintableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MintableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MintableTransfer)
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
		it.Event = new(IERC20MintableTransfer)
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
func (it *IERC20MintableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MintableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MintableTransfer represents a Transfer event raised by the IERC20Mintable contract.
type IERC20MintableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Mintable *IERC20MintableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MintableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Mintable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MintableTransferIterator{contract: _IERC20Mintable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Mintable *IERC20MintableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MintableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Mintable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MintableTransfer)
				if err := _IERC20Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Mintable *IERC20MintableFilterer) ParseTransfer(log types.Log) (*IERC20MintableTransfer, error) {
	event := new(IERC20MintableTransfer)
	if err := _IERC20Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISwapMetaData contains all meta data concerning the ISwap contract.
var ISwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateRemoveLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"deposit\",\"type\":\"bool\"}],\"name\":\"calculateTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVirtualPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"pooledTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"decimals\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"lpTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lpTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adminFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpTokenTargetAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"maxBurnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4d49e87d": "addLiquidity(uint256[],uint256,uint256)",
		"f2fad2b6": "calculateRemoveLiquidity(uint256)",
		"342a87a1": "calculateRemoveLiquidityOneToken(uint256,uint8)",
		"a95b089f": "calculateSwap(uint8,uint8,uint256)",
		"e6ab2806": "calculateTokenAmount(uint256[],bool)",
		"d46300fd": "getA()",
		"82b86600": "getToken(uint8)",
		"91ceb3eb": "getTokenBalance(uint8)",
		"66c0bd24": "getTokenIndex(address)",
		"e25aa5fa": "getVirtualPrice()",
		"b28cb6dc": "initialize(address[],uint8[],string,string,uint256,uint256,uint256,address)",
		"31cd52b0": "removeLiquidity(uint256,uint256[],uint256)",
		"84cdd9bc": "removeLiquidityImbalance(uint256[],uint256,uint256)",
		"3e3a1560": "removeLiquidityOneToken(uint256,uint8,uint256,uint256)",
		"91695586": "swap(uint8,uint8,uint256,uint256,uint256)",
	},
}

// ISwapABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapMetaData.ABI instead.
var ISwapABI = ISwapMetaData.ABI

// Deprecated: Use ISwapMetaData.Sigs instead.
// ISwapFuncSigs maps the 4-byte function signature to its string representation.
var ISwapFuncSigs = ISwapMetaData.Sigs

// ISwap is an auto generated Go binding around an Ethereum contract.
type ISwap struct {
	ISwapCaller     // Read-only binding to the contract
	ISwapTransactor // Write-only binding to the contract
	ISwapFilterer   // Log filterer for contract events
}

// ISwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapSession struct {
	Contract     *ISwap            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapCallerSession struct {
	Contract *ISwapCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapTransactorSession struct {
	Contract     *ISwapTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapRaw struct {
	Contract *ISwap // Generic contract binding to access the raw methods on
}

// ISwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapCallerRaw struct {
	Contract *ISwapCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapTransactorRaw struct {
	Contract *ISwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwap creates a new instance of ISwap, bound to a specific deployed contract.
func NewISwap(address common.Address, backend bind.ContractBackend) (*ISwap, error) {
	contract, err := bindISwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwap{ISwapCaller: ISwapCaller{contract: contract}, ISwapTransactor: ISwapTransactor{contract: contract}, ISwapFilterer: ISwapFilterer{contract: contract}}, nil
}

// NewISwapCaller creates a new read-only instance of ISwap, bound to a specific deployed contract.
func NewISwapCaller(address common.Address, caller bind.ContractCaller) (*ISwapCaller, error) {
	contract, err := bindISwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapCaller{contract: contract}, nil
}

// NewISwapTransactor creates a new write-only instance of ISwap, bound to a specific deployed contract.
func NewISwapTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapTransactor, error) {
	contract, err := bindISwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapTransactor{contract: contract}, nil
}

// NewISwapFilterer creates a new log filterer instance of ISwap, bound to a specific deployed contract.
func NewISwapFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapFilterer, error) {
	contract, err := bindISwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapFilterer{contract: contract}, nil
}

// bindISwap binds a generic wrapper to an already deployed contract.
func bindISwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwap *ISwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwap.Contract.ISwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwap *ISwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwap.Contract.ISwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwap *ISwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwap.Contract.ISwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwap *ISwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwap *ISwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwap *ISwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwap.Contract.contract.Transact(opts, method, params...)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_ISwap *ISwapCaller) CalculateRemoveLiquidity(opts *bind.CallOpts, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calculateRemoveLiquidity", amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_ISwap *ISwapSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _ISwap.Contract.CalculateRemoveLiquidity(&_ISwap.CallOpts, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_ISwap *ISwapCallerSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _ISwap.Contract.CalculateRemoveLiquidity(&_ISwap.CallOpts, amount)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_ISwap *ISwapCaller) CalculateRemoveLiquidityOneToken(opts *bind.CallOpts, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calculateRemoveLiquidityOneToken", tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_ISwap *ISwapSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _ISwap.Contract.CalculateRemoveLiquidityOneToken(&_ISwap.CallOpts, tokenAmount, tokenIndex)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_ISwap *ISwapCallerSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _ISwap.Contract.CalculateRemoveLiquidityOneToken(&_ISwap.CallOpts, tokenAmount, tokenIndex)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_ISwap *ISwapCaller) CalculateSwap(opts *bind.CallOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calculateSwap", tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_ISwap *ISwapSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _ISwap.Contract.CalculateSwap(&_ISwap.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_ISwap *ISwapCallerSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _ISwap.Contract.CalculateSwap(&_ISwap.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_ISwap *ISwapCaller) CalculateTokenAmount(opts *bind.CallOpts, amounts []*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calculateTokenAmount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_ISwap *ISwapSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _ISwap.Contract.CalculateTokenAmount(&_ISwap.CallOpts, amounts, deposit)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_ISwap *ISwapCallerSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _ISwap.Contract.CalculateTokenAmount(&_ISwap.CallOpts, amounts, deposit)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_ISwap *ISwapCaller) GetA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_ISwap *ISwapSession) GetA() (*big.Int, error) {
	return _ISwap.Contract.GetA(&_ISwap.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_ISwap *ISwapCallerSession) GetA() (*big.Int, error) {
	return _ISwap.Contract.GetA(&_ISwap.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_ISwap *ISwapCaller) GetToken(opts *bind.CallOpts, index uint8) (common.Address, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_ISwap *ISwapSession) GetToken(index uint8) (common.Address, error) {
	return _ISwap.Contract.GetToken(&_ISwap.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_ISwap *ISwapCallerSession) GetToken(index uint8) (common.Address, error) {
	return _ISwap.Contract.GetToken(&_ISwap.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_ISwap *ISwapCaller) GetTokenBalance(opts *bind.CallOpts, index uint8) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getTokenBalance", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_ISwap *ISwapSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _ISwap.Contract.GetTokenBalance(&_ISwap.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_ISwap *ISwapCallerSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _ISwap.Contract.GetTokenBalance(&_ISwap.CallOpts, index)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_ISwap *ISwapCaller) GetTokenIndex(opts *bind.CallOpts, tokenAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getTokenIndex", tokenAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_ISwap *ISwapSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _ISwap.Contract.GetTokenIndex(&_ISwap.CallOpts, tokenAddress)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_ISwap *ISwapCallerSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _ISwap.Contract.GetTokenIndex(&_ISwap.CallOpts, tokenAddress)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_ISwap *ISwapCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getVirtualPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_ISwap *ISwapSession) GetVirtualPrice() (*big.Int, error) {
	return _ISwap.Contract.GetVirtualPrice(&_ISwap.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_ISwap *ISwapCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _ISwap.Contract.GetVirtualPrice(&_ISwap.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactor) AddLiquidity(opts *bind.TransactOpts, amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "addLiquidity", amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_ISwap *ISwapSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.AddLiquidity(&_ISwap.TransactOpts, amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactorSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.AddLiquidity(&_ISwap.TransactOpts, amounts, minToMint, deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 a, uint256 fee, uint256 adminFee, address lpTokenTargetAddress) returns()
func (_ISwap *ISwapTransactor) Initialize(opts *bind.TransactOpts, pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, a *big.Int, fee *big.Int, adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "initialize", pooledTokens, decimals, lpTokenName, lpTokenSymbol, a, fee, adminFee, lpTokenTargetAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 a, uint256 fee, uint256 adminFee, address lpTokenTargetAddress) returns()
func (_ISwap *ISwapSession) Initialize(pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, a *big.Int, fee *big.Int, adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _ISwap.Contract.Initialize(&_ISwap.TransactOpts, pooledTokens, decimals, lpTokenName, lpTokenSymbol, a, fee, adminFee, lpTokenTargetAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 a, uint256 fee, uint256 adminFee, address lpTokenTargetAddress) returns()
func (_ISwap *ISwapTransactorSession) Initialize(pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, a *big.Int, fee *big.Int, adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _ISwap.Contract.Initialize(&_ISwap.TransactOpts, pooledTokens, decimals, lpTokenName, lpTokenSymbol, a, fee, adminFee, lpTokenTargetAddress)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_ISwap *ISwapTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "removeLiquidity", amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_ISwap *ISwapSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidity(&_ISwap.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_ISwap *ISwapTransactorSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidity(&_ISwap.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "removeLiquidityImbalance", amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityImbalance(&_ISwap.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactorSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityImbalance(&_ISwap.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactor) RemoveLiquidityOneToken(opts *bind.TransactOpts, tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "removeLiquidityOneToken", tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityOneToken(&_ISwap.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactorSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityOneToken(&_ISwap.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactor) Swap(opts *bind.TransactOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "swap", tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_ISwap *ISwapSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.Swap(&_ISwap.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactorSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.Swap(&_ISwap.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// IWETH9MetaData contains all meta data concerning the IWETH9 contract.
var IWETH9MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"d0e30db0": "deposit()",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// IWETH9ABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETH9MetaData.ABI instead.
var IWETH9ABI = IWETH9MetaData.ABI

// Deprecated: Use IWETH9MetaData.Sigs instead.
// IWETH9FuncSigs maps the 4-byte function signature to its string representation.
var IWETH9FuncSigs = IWETH9MetaData.Sigs

// IWETH9 is an auto generated Go binding around an Ethereum contract.
type IWETH9 struct {
	IWETH9Caller     // Read-only binding to the contract
	IWETH9Transactor // Write-only binding to the contract
	IWETH9Filterer   // Log filterer for contract events
}

// IWETH9Caller is an auto generated read-only Go binding around an Ethereum contract.
type IWETH9Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETH9Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETH9Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETH9Session struct {
	Contract     *IWETH9           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETH9CallerSession struct {
	Contract *IWETH9Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETH9TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETH9TransactorSession struct {
	Contract     *IWETH9Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9Raw is an auto generated low-level Go binding around an Ethereum contract.
type IWETH9Raw struct {
	Contract *IWETH9 // Generic contract binding to access the raw methods on
}

// IWETH9CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETH9CallerRaw struct {
	Contract *IWETH9Caller // Generic read-only contract binding to access the raw methods on
}

// IWETH9TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETH9TransactorRaw struct {
	Contract *IWETH9Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH9 creates a new instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9(address common.Address, backend bind.ContractBackend) (*IWETH9, error) {
	contract, err := bindIWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH9{IWETH9Caller: IWETH9Caller{contract: contract}, IWETH9Transactor: IWETH9Transactor{contract: contract}, IWETH9Filterer: IWETH9Filterer{contract: contract}}, nil
}

// NewIWETH9Caller creates a new read-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Caller(address common.Address, caller bind.ContractCaller) (*IWETH9Caller, error) {
	contract, err := bindIWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Caller{contract: contract}, nil
}

// NewIWETH9Transactor creates a new write-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*IWETH9Transactor, error) {
	contract, err := bindIWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Transactor{contract: contract}, nil
}

// NewIWETH9Filterer creates a new log filterer instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*IWETH9Filterer, error) {
	contract, err := bindIWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETH9Filterer{contract: contract}, nil
}

// bindIWETH9 binds a generic wrapper to an already deployed contract.
func bindIWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWETH9MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.IWETH9Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IWETH9 *IWETH9Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IWETH9 *IWETH9Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IWETH9.Contract.Allowance(&_IWETH9.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IWETH9.Contract.Allowance(&_IWETH9.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IWETH9 *IWETH9Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IWETH9 *IWETH9Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _IWETH9.Contract.BalanceOf(&_IWETH9.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _IWETH9.Contract.BalanceOf(&_IWETH9.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH9 *IWETH9Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH9 *IWETH9Session) Decimals() (uint8, error) {
	return _IWETH9.Contract.Decimals(&_IWETH9.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH9 *IWETH9CallerSession) Decimals() (uint8, error) {
	return _IWETH9.Contract.Decimals(&_IWETH9.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH9 *IWETH9Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH9 *IWETH9Session) Name() (string, error) {
	return _IWETH9.Contract.Name(&_IWETH9.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH9 *IWETH9CallerSession) Name() (string, error) {
	return _IWETH9.Contract.Name(&_IWETH9.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH9 *IWETH9Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH9 *IWETH9Session) Symbol() (string, error) {
	return _IWETH9.Contract.Symbol(&_IWETH9.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH9 *IWETH9CallerSession) Symbol() (string, error) {
	return _IWETH9.Contract.Symbol(&_IWETH9.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9Session) TotalSupply() (*big.Int, error) {
	return _IWETH9.Contract.TotalSupply(&_IWETH9.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) TotalSupply() (*big.Int, error) {
	return _IWETH9.Contract.TotalSupply(&_IWETH9.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9Transactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9Session) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Approve(&_IWETH9.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Approve(&_IWETH9.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Session) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9TransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9Transactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9Session) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Transfer(&_IWETH9.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Transfer(&_IWETH9.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9Transactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9Session) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.TransferFrom(&_IWETH9.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.TransferFrom(&_IWETH9.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWETH9 *IWETH9Transactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWETH9 *IWETH9Session) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWETH9 *IWETH9TransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, wad)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IWETH9 *IWETH9Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IWETH9 *IWETH9Session) Receive() (*types.Transaction, error) {
	return _IWETH9.Contract.Receive(&_IWETH9.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IWETH9 *IWETH9TransactorSession) Receive() (*types.Transaction, error) {
	return _IWETH9.Contract.Receive(&_IWETH9.TransactOpts)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[]",
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

// PausableUpgradeableMetaData contains all meta data concerning the PausableUpgradeable contract.
var PausableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c975abb": "paused()",
	},
}

// PausableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableUpgradeableMetaData.ABI instead.
var PausableUpgradeableABI = PausableUpgradeableMetaData.ABI

// Deprecated: Use PausableUpgradeableMetaData.Sigs instead.
// PausableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var PausableUpgradeableFuncSigs = PausableUpgradeableMetaData.Sigs

// PausableUpgradeable is an auto generated Go binding around an Ethereum contract.
type PausableUpgradeable struct {
	PausableUpgradeableCaller     // Read-only binding to the contract
	PausableUpgradeableTransactor // Write-only binding to the contract
	PausableUpgradeableFilterer   // Log filterer for contract events
}

// PausableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableUpgradeableSession struct {
	Contract     *PausableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PausableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableUpgradeableCallerSession struct {
	Contract *PausableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PausableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableUpgradeableTransactorSession struct {
	Contract     *PausableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PausableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableUpgradeableRaw struct {
	Contract *PausableUpgradeable // Generic contract binding to access the raw methods on
}

// PausableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableUpgradeableCallerRaw struct {
	Contract *PausableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableUpgradeableTransactorRaw struct {
	Contract *PausableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableUpgradeable creates a new instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeable(address common.Address, backend bind.ContractBackend) (*PausableUpgradeable, error) {
	contract, err := bindPausableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeable{PausableUpgradeableCaller: PausableUpgradeableCaller{contract: contract}, PausableUpgradeableTransactor: PausableUpgradeableTransactor{contract: contract}, PausableUpgradeableFilterer: PausableUpgradeableFilterer{contract: contract}}, nil
}

// NewPausableUpgradeableCaller creates a new read-only instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*PausableUpgradeableCaller, error) {
	contract, err := bindPausableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableCaller{contract: contract}, nil
}

// NewPausableUpgradeableTransactor creates a new write-only instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableUpgradeableTransactor, error) {
	contract, err := bindPausableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableTransactor{contract: contract}, nil
}

// NewPausableUpgradeableFilterer creates a new log filterer instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableUpgradeableFilterer, error) {
	contract, err := bindPausableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableFilterer{contract: contract}, nil
}

// bindPausableUpgradeable binds a generic wrapper to an already deployed contract.
func bindPausableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PausableUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeable *PausableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeable.Contract.PausableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeable *PausableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.PausableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeable *PausableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.PausableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeable *PausableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeable *PausableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeable *PausableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PausableUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableSession) Paused() (bool, error) {
	return _PausableUpgradeable.Contract.Paused(&_PausableUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableCallerSession) Paused() (bool, error) {
	return _PausableUpgradeable.Contract.Paused(&_PausableUpgradeable.CallOpts)
}

// PausableUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PausableUpgradeable contract.
type PausableUpgradeablePausedIterator struct {
	Event *PausableUpgradeablePaused // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeablePaused)
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
		it.Event = new(PausableUpgradeablePaused)
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
func (it *PausableUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeablePaused represents a Paused event raised by the PausableUpgradeable contract.
type PausableUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausableUpgradeablePausedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeablePausedIterator{contract: _PausableUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeablePaused)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParsePaused(log types.Log) (*PausableUpgradeablePaused, error) {
	event := new(PausableUpgradeablePaused)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PausableUpgradeable contract.
type PausableUpgradeableUnpausedIterator struct {
	Event *PausableUpgradeableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeableUnpaused)
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
		it.Event = new(PausableUpgradeableUnpaused)
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
func (it *PausableUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeableUnpaused represents a Unpaused event raised by the PausableUpgradeable contract.
type PausableUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUpgradeableUnpausedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableUnpausedIterator{contract: _PausableUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeableUnpaused)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParseUnpaused(log types.Log) (*PausableUpgradeableUnpaused, error) {
	event := new(PausableUpgradeableUnpaused)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardUpgradeableMetaData contains all meta data concerning the ReentrancyGuardUpgradeable contract.
var ReentrancyGuardUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
}

// ReentrancyGuardUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardUpgradeableMetaData.ABI instead.
var ReentrancyGuardUpgradeableABI = ReentrancyGuardUpgradeableMetaData.ABI

// ReentrancyGuardUpgradeable is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeable struct {
	ReentrancyGuardUpgradeableCaller     // Read-only binding to the contract
	ReentrancyGuardUpgradeableTransactor // Write-only binding to the contract
	ReentrancyGuardUpgradeableFilterer   // Log filterer for contract events
}

// ReentrancyGuardUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardUpgradeableSession struct {
	Contract     *ReentrancyGuardUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardUpgradeableCallerSession struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ReentrancyGuardUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardUpgradeableTransactorSession struct {
	Contract     *ReentrancyGuardUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableRaw struct {
	Contract *ReentrancyGuardUpgradeable // Generic contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCallerRaw struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactorRaw struct {
	Contract *ReentrancyGuardUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuardUpgradeable creates a new instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeable(address common.Address, backend bind.ContractBackend) (*ReentrancyGuardUpgradeable, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeable{ReentrancyGuardUpgradeableCaller: ReentrancyGuardUpgradeableCaller{contract: contract}, ReentrancyGuardUpgradeableTransactor: ReentrancyGuardUpgradeableTransactor{contract: contract}, ReentrancyGuardUpgradeableFilterer: ReentrancyGuardUpgradeableFilterer{contract: contract}}, nil
}

// NewReentrancyGuardUpgradeableCaller creates a new read-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardUpgradeableCaller, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableCaller{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableTransactor creates a new write-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardUpgradeableTransactor, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableTransactor{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableFilterer creates a new log filterer instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardUpgradeableFilterer, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableFilterer{contract: contract}, nil
}

// bindReentrancyGuardUpgradeable binds a generic wrapper to an already deployed contract.
func bindReentrancyGuardUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220af1ffd449135a3159f6d6dd11b323c42c2527fb9c4fd60513fdb2a9f114eec0464736f6c634300060c0033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208fd2c64e6a9faa06b6bce3eb14400b0a013ef84dc15221323ef5f825557829e364736f6c634300060c0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SynapseBridgeMetaData contains all meta data concerning the SynapseBridge contract.
var SynapseBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TokenDepositAndSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"TokenMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"swapSuccess\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"TokenMintAndSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"}],\"name\":\"TokenRedeemAndRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TokenRedeemAndSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRedeemV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"swapSuccess\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"TokenWithdrawAndRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODEGROUP_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH_ADDRESS\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"kappas\",\"type\":\"bytes32[]\"}],\"name\":\"addKappas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"depositAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"kappaExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"contractISwap\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"mintAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"}],\"name\":\"redeemAndRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"redeemAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_wethAddress\",\"type\":\"address\"}],\"name\":\"setWethAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"contractISwap\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"f36c8f5c": "GOVERNANCE_ROLE()",
		"f3befd01": "NODEGROUP_ROLE()",
		"040141e5": "WETH_ADDRESS()",
		"e7a59998": "addKappas(bytes32[])",
		"ac865626": "bridgeVersion()",
		"e00a83e0": "chainGasAmount()",
		"90d25074": "deposit(address,uint256,address,uint256)",
		"a2a2af0b": "depositAndSwap(address,uint256,address,uint256,uint8,uint8,uint256,uint256)",
		"c78f6803": "getFeeBalance(address)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"8129fc1c": "initialize()",
		"2fe87b95": "kappaExists(bytes32)",
		"20d7b327": "mint(address,address,uint256,uint256,bytes32)",
		"17357892": "mintAndSwap(address,address,uint256,uint256,address,uint8,uint8,uint256,uint256,bytes32)",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"f3f094a1": "redeem(address,uint256,address,uint256)",
		"36e712ed": "redeemAndRemove(address,uint256,address,uint256,uint8,uint256,uint256)",
		"839ed90a": "redeemAndSwap(address,uint256,address,uint256,uint8,uint8,uint256,uint256)",
		"a07ed975": "redeemV2(bytes32,uint256,address,uint256)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"b250fe6b": "setChainGasAmount(uint256)",
		"a96e2423": "setWethAddress(address)",
		"498a4c2d": "startBlockNumber()",
		"3f4ba83a": "unpause()",
		"1cf5f07f": "withdraw(address,address,uint256,uint256,bytes32)",
		"d57eafac": "withdrawAndRemove(address,address,uint256,uint256,address,uint8,uint256,uint256,bytes32)",
		"f2555278": "withdrawFees(address,address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50613dcd806100206000396000f3fe60806040526004361061021d5760003560e01c806391d148541161011d578063ca15c873116100b0578063e7a599981161007f578063f36c8f5c11610064578063f36c8f5c1461093b578063f3befd0114610950578063f3f094a11461096557610224565b8063e7a5999814610883578063f25552781461090057610224565b8063ca15c8731461079f578063d547741f146107c9578063d57eafac14610802578063e00a83e01461086e57610224565b8063a96e2423116100ec578063a96e2423146106fa578063ac8656261461072d578063b250fe6b14610742578063c78f68031461076c57610224565b806391d1485414610600578063a07ed97514610639578063a217fddf1461067e578063a2a2af0b1461069357610224565b806336e712ed116101b05780638129fc1c1161017f5780638456cb59116101645780638456cb59146105745780639010d07c1461058957806390d25074146105b957610224565b80638129fc1c146104f8578063839ed90a1461050d57610224565b806336e712ed1461045d5780633f4ba83a146104b9578063498a4c2d146104ce5780635c975abb146104e357610224565b8063248a9ca3116101ec578063248a9ca3146103715780632f2ff15d146103ad5780632fe87b95146103e657806336568abe1461042457610224565b8063040141e514610229578063173578921461025a5780631cf5f07f146102d357806320d7b3271461032257610224565b3661022457005b600080fd5b34801561023557600080fd5b5061023e6109ac565b604080516001600160a01b039092168252519081900360200190f35b34801561026657600080fd5b506102d1600480360361014081101561027e57600080fd5b506001600160a01b03813581169160208101358216916040820135916060810135916080820135169060ff60a082013581169160c08101359091169060e0810135906101008101359061012001356109bb565b005b3480156102df57600080fd5b506102d1600480360360a08110156102f657600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060800135611315565b34801561032e57600080fd5b506102d1600480360360a081101561034557600080fd5b506001600160a01b0381358116916020810135909116906040810135906060810135906080013561175e565b34801561037d57600080fd5b5061039b6004803603602081101561039457600080fd5b5035611b04565b60408051918252519081900360200190f35b3480156103b957600080fd5b506102d1600480360360408110156103d057600080fd5b50803590602001356001600160a01b0316611b19565b3480156103f257600080fd5b506104106004803603602081101561040957600080fd5b5035611b85565b604080519115158252519081900360200190f35b34801561043057600080fd5b506102d16004803603604081101561044757600080fd5b50803590602001356001600160a01b0316611b9a565b34801561046957600080fd5b506102d1600480360360e081101561048057600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff6080820135169060a08101359060c00135611bfb565b3480156104c557600080fd5b506102d1611d91565b3480156104da57600080fd5b5061039b611e16565b3480156104ef57600080fd5b50610410611e1c565b34801561050457600080fd5b506102d1611e25565b34801561051957600080fd5b506102d1600480360361010081101561053157600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff608082013581169160a08101359091169060c08101359060e00135611f18565b34801561058057600080fd5b506102d16120c9565b34801561059557600080fd5b5061023e600480360360408110156105ac57600080fd5b508035906020013561214c565b3480156105c557600080fd5b506102d1600480360360808110156105dc57600080fd5b506001600160a01b038135811691602081013591604082013516906060013561216d565b34801561060c57600080fd5b506104106004803603604081101561062357600080fd5b50803590602001356001600160a01b0316612291565b34801561064557600080fd5b506102d16004803603608081101561065c57600080fd5b508035906020810135906001600160a01b0360408201351690606001356122a9565b34801561068a57600080fd5b5061039b612435565b34801561069f57600080fd5b506102d160048036036101008110156106b757600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff608082013581169160a08101359091169060c08101359060e0013561243a565b34801561070657600080fd5b506102d16004803603602081101561071d57600080fd5b50356001600160a01b0316612572565b34801561073957600080fd5b5061039b612608565b34801561074e57600080fd5b506102d16004803603602081101561076557600080fd5b503561260d565b34801561077857600080fd5b5061039b6004803603602081101561078f57600080fd5b50356001600160a01b031661268d565b3480156107ab57600080fd5b5061039b600480360360208110156107c257600080fd5b50356126a8565b3480156107d557600080fd5b506102d1600480360360408110156107ec57600080fd5b50803590602001356001600160a01b03166126bf565b34801561080e57600080fd5b506102d1600480360361012081101561082657600080fd5b506001600160a01b03813581169160208101358216916040820135916060810135916080820135169060ff60a0820135169060c08101359060e0810135906101000135612718565b34801561087a57600080fd5b5061039b612cd1565b34801561088f57600080fd5b506102d1600480360360208110156108a657600080fd5b8101906020810181356401000000008111156108c157600080fd5b8201836020820111156108d357600080fd5b803590602001918460208302840111640100000000831117156108f557600080fd5b509092509050612cd7565b34801561090c57600080fd5b506102d16004803603604081101561092357600080fd5b506001600160a01b0381358116916020013516612da4565b34801561094757600080fd5b5061039b612f32565b34801561095c57600080fd5b5061039b612f56565b34801561097157600080fd5b506102d16004803603608081101561098857600080fd5b506001600160a01b0381358116916020810135916040820135169060600135612f7a565b60cc546001600160a01b031681565b60026065541415610a13576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555610a20611e1c565b15610a72576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b610a9c7fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd33612291565b610aed576040805162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f742061206e6f64652067726f7570000000000000604482015290519081900360640190fd5b868811610b41576040805162461bcd60e51b815260206004820152601f60248201527f416d6f756e74206d7573742062652067726561746572207468616e2066656500604482015290519081900360640190fd5b600081815260cd602052604090205460ff1615610ba5576040805162461bcd60e51b815260206004820152601860248201527f4b6170706120697320616c72656164792070726573656e740000000000000000604482015290519081900360640190fd5b600081815260cd60209081526040808320805460ff191660011790556001600160a01b038c16835260c9909152902054610bdf90886130e5565b6001600160a01b038a16600090815260c9602052604090205560cb5415801590610c0a575060cb5447115b15610c635760cb546040516001600160a01b038c169190600081818185875af1925050503d8060008114610c5a576040519150601f19603f3d011682016040523d82523d6000602084013e610c5f565b606091505b5050505b60006001600160a01b03871663a95b089f8787610c808d8d61313f565b6040518463ffffffff1660e01b8152600401808460ff1681526020018360ff168152602001828152602001935050505060206040518083038186803b158015610cc857600080fd5b505afa158015610cdc573d6000803e3d6000fd5b505050506040513d6020811015610cf257600080fd5b505190508381106111e557604080517f40c10f19000000000000000000000000000000000000000000000000000000008152306004820152602481018b905290516001600160a01b038c16916340c10f1991604480830192600092919082900301818387803b158015610d6457600080fd5b505af1158015610d78573d6000803e3d6000fd5b50610d91925050506001600160a01b038b16888b61319c565b6001600160a01b03871663916955868787610dac8d8d61313f565b88886040518663ffffffff1660e01b8152600401808660ff1681526020018560ff16815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b158015610e0657600080fd5b505af1925050508015610e2b57506040513d6020811015610e2657600080fd5b505160015b610eda57610e4e8b610e3d8b8b61313f565b6001600160a01b038d1691906132bb565b816001600160a01b038c167f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c658c610e858d8d61313f565b604080516001600160a01b03909316835260208301919091528181018d905260ff808c1660608401528a16608083015260a0820189905260c08201889052600060e083015251908190036101000190a36111e0565b6000886001600160a01b03166382b86600886040518263ffffffff1660e01b8152600401808260ff16815260200191505060206040518083038186803b158015610f2357600080fd5b505afa158015610f37573d6000803e3d6000fd5b505050506040513d6020811015610f4d57600080fd5b505160cc549091506001600160a01b038083169116148015610f79575060cc546001600160a01b031615155b1561113c5760cc54604080517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526004810185905290516001600160a01b0390921691632e1a7d4d9160248082019260009290919082900301818387803b158015610fe457600080fd5b505af1158015610ff8573d6000803e3d6000fd5b5050505060008d6001600160a01b03168360405180600001905060006040518083038185875af1925050503d806000811461104f576040519150601f19603f3d011682016040523d82523d6000602084013e611054565b606091505b50509050806110aa576040805162461bcd60e51b815260206004820152601360248201527f4554485f5452414e534645525f4641494c454400000000000000000000000000604482015290519081900360640190fd5b848e6001600160a01b03167f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c658f868f8e8e8e8e600160405180896001600160a01b031681526020018881526020018781526020018660ff1681526020018560ff16815260200184815260200183815260200182151581526020019850505050505050505060405180910390a3506111dd565b6111506001600160a01b0382168e846132bb565b838d6001600160a01b03167f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c658e858e8d8d8d8d600160405180896001600160a01b031681526020018881526020018781526020018660ff1681526020018560ff16815260200184815260200183815260200182151581526020019850505050505050505060405180910390a35b50505b611303565b604080517f40c10f19000000000000000000000000000000000000000000000000000000008152306004820152602481018b905290516001600160a01b038c16916340c10f1991604480830192600092919082900301818387803b15801561124c57600080fd5b505af1158015611260573d6000803e3d6000fd5b5050505061127b8b610e3d8a8c61313f90919063ffffffff16565b816001600160a01b038c167f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c658c6112b28d8d61313f565b604080516001600160a01b03909316835260208301919091528181018d905260ff808c1660608401528a16608083015260a0820189905260c08201889052600060e083015251908190036101000190a35b50506001606555505050505050505050565b6002606554141561136d576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260655561137a611e1c565b156113cc576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6113f67fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd33612291565b611447576040805162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f742061206e6f64652067726f7570000000000000604482015290519081900360640190fd5b81831161149b576040805162461bcd60e51b815260206004820152601f60248201527f416d6f756e74206d7573742062652067726561746572207468616e2066656500604482015290519081900360640190fd5b600081815260cd602052604090205460ff16156114ff576040805162461bcd60e51b815260206004820152601860248201527f4b6170706120697320616c72656164792070726573656e740000000000000000604482015290519081900360640190fd5b600081815260cd60209081526040808320805460ff191660011790556001600160a01b038716835260c990915290205461153990836130e5565b6001600160a01b03808616600081815260c9602052604090209290925560cc5416148015611571575060cc546001600160a01b031615155b156116e35760cc546001600160a01b0316632e1a7d4d611591858561313f565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156115c757600080fd5b505af11580156115db573d6000803e3d6000fd5b506000925050506001600160a01b0386166115f6858561313f565b604051600081818185875af1925050503d8060008114611632576040519150601f19603f3d011682016040523d82523d6000602084013e611637565b606091505b505090508061168d576040805162461bcd60e51b815260206004820152601360248201527f4554485f5452414e534645525f4641494c454400000000000000000000000000604482015290519081900360640190fd5b604080516001600160a01b03878116825260208201879052818301869052915184928916917f8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26919081900360600190a350611752565b604080516001600160a01b03868116825260208201869052818301859052915183928816917f8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26919081900360600190a361175285611741858561313f565b6001600160a01b03871691906132bb565b50506001606555505050565b600260655414156117b6576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026065556117c3611e1c565b15611815576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b61183f7fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd33612291565b611890576040805162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f742061206e6f64652067726f7570000000000000604482015290519081900360640190fd5b8183116118e4576040805162461bcd60e51b815260206004820152601f60248201527f416d6f756e74206d7573742062652067726561746572207468616e2066656500604482015290519081900360640190fd5b600081815260cd602052604090205460ff1615611948576040805162461bcd60e51b815260206004820152601860248201527f4b6170706120697320616c72656164792070726573656e740000000000000000604482015290519081900360640190fd5b600081815260cd60209081526040808320805460ff191660011790556001600160a01b038716835260c990915290205461198290836130e5565b6001600160a01b03808616600090815260c96020526040902091909155819086167fbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd38866119cf878761313f565b604080516001600160a01b0390931683526020830191909152818101879052519081900360600190a3604080517f40c10f190000000000000000000000000000000000000000000000000000000081523060048201526024810185905290516001600160a01b038616916340c10f1991604480830192600092919082900301818387803b158015611a5f57600080fd5b505af1158015611a73573d6000803e3d6000fd5b50505050611a8e85611741848661313f90919063ffffffff16565b60cb5415801590611aa0575060cb5447115b156117525760cb546040516001600160a01b0387169190600081818185875af1925050503d8060008114611af0576040519150601f19603f3d011682016040523d82523d6000602084013e611af5565b606091505b50505050506001606555505050565b60009081526033602052604090206002015490565b600082815260336020526040902060020154611b3c90611b3761333b565b612291565b611b775760405162461bcd60e51b815260040180806020018281038252602f815260200180613c8c602f913960400191505060405180910390fd5b611b81828261333f565b5050565b600090815260cd602052604090205460ff1690565b611ba261333b565b6001600160a01b0316816001600160a01b031614611bf15760405162461bcd60e51b815260040180806020018281038252602f815260200180613d69602f913960400191505060405180910390fd5b611b8182826133a8565b60026065541415611c53576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555611c60611e1c565b15611cb2576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518781526001600160a01b03878116602083015281830187905260ff861660608301526080820185905260a082018490529151918916917f9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c9181900360c00190a2604080517f79cc67900000000000000000000000000000000000000000000000000000000081523360048201526024810186905290516001600160a01b038716916379cc679091604480830192600092919082900301818387803b158015611d7d57600080fd5b505af1158015611303573d6000803e3d6000fd5b611dbb7f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb133612291565b611e0c576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b611e14613411565b565b60ca5481565b60975460ff1690565b600054610100900460ff1680611e3e5750611e3e6134ba565b80611e4c575060005460ff16155b611e875760405162461bcd60e51b815260040180806020018281038252602e815260200180613d11602e913960400191505060405180910390fd5b600054610100900460ff16158015611ecf576000805460ff197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b4360ca55611ede600033611b77565b611ee66134cb565b8015611f1557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50565b60026065541415611f70576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555611f7d611e1c565b15611fcf576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518881526001600160a01b03888116602083015281830188905260ff80881660608401528616608083015260a0820185905260c082018490529151918a16917f91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc3194259181900360e00190a2604080517f79cc67900000000000000000000000000000000000000000000000000000000081523360048201526024810187905290516001600160a01b038816916379cc679091604480830192600092919082900301818387803b1580156120a257600080fd5b505af11580156120b6573d6000803e3d6000fd5b5050600160655550505050505050505050565b6120f37f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb133612291565b612144576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b611e14613585565b60008281526033602052604081206121649083613615565b90505b92915050565b600260655414156121c5576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026065556121d2611e1c565b15612224576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518481526001600160a01b0384811660208301528183018490529151918616917fda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b9181900360600190a26122866001600160a01b038316333084613621565b505060016065555050565b600082815260336020526040812061216490836136a9565b60026065541415612301576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260655561230e611e1c565b15612360576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518481526001600160a01b0384166020820152808201839052905185917f8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499919081900360600190a2604080517f79cc67900000000000000000000000000000000000000000000000000000000081523360048201526024810183905290516001600160a01b038416916379cc679091604480830192600092919082900301818387803b15801561241257600080fd5b505af1158015612426573d6000803e3d6000fd5b50506001606555505050505050565b600081565b60026065541415612492576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260655561249f611e1c565b156124f1576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518881526001600160a01b03888116602083015281830188905260ff80881660608401528616608083015260a0820185905260c082018490529151918a16917f79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f9181900360e00190a26124266001600160a01b038716333088613621565b61257d600033612291565b6125ce576040805162461bcd60e51b815260206004820152600960248201527f4e6f742061646d696e0000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60cc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600681565b6126377f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb133612291565b612688576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b60cb55565b6001600160a01b0316600090815260c9602052604090205490565b6000818152603360205260408120612167906136be565b6000828152603360205260409020600201546126dd90611b3761333b565b611bf15760405162461bcd60e51b8152600401808060200182810382526030815260200180613ce16030913960400191505060405180910390fd5b60026065541415612770576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260655561277d611e1c565b156127cf576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6127f97fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd33612291565b61284a576040805162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f742061206e6f64652067726f7570000000000000604482015290519081900360640190fd5b85871161289e576040805162461bcd60e51b815260206004820152601f60248201527f416d6f756e74206d7573742062652067726561746572207468616e2066656500604482015290519081900360640190fd5b600081815260cd602052604090205460ff1615612902576040805162461bcd60e51b815260206004820152601860248201527f4b6170706120697320616c72656164792070726573656e740000000000000000604482015290519081900360640190fd5b600081815260cd60209081526040808320805460ff191660011790556001600160a01b038b16835260c990915290205461293c90876130e5565b6001600160a01b03808a16600090815260c96020526040812092909255861663342a87a161296a8a8a61313f565b876040518363ffffffff1660e01b8152600401808381526020018260ff1681526020019250505060206040518083038186803b1580156129a957600080fd5b505afa1580156129bd573d6000803e3d6000fd5b505050506040513d60208110156129d357600080fd5b50519050838110612c33576129fd866129ec8a8a61313f565b6001600160a01b038c16919061319c565b6001600160a01b038616633e3a1560612a168a8a61313f565b8787876040518563ffffffff1660e01b8152600401808581526020018460ff168152602001838152602001828152602001945050505050602060405180830381600087803b158015612a6757600080fd5b505af1925050508015612a8c57506040513d6020811015612a8757600080fd5b505160015b612b3257612aaf8a612a9e8a8a61313f565b6001600160a01b038c1691906132bb565b816001600160a01b038b167fc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d8b612ae68c8c61313f565b604080516001600160a01b03909316835260208301919091528181018c905260ff8a1660608301526080820189905260a08201889052600060c0830152519081900360e00190a3612c2e565b6000876001600160a01b03166382b86600886040518263ffffffff1660e01b8152600401808260ff16815260200191505060206040518083038186803b158015612b7b57600080fd5b505afa158015612b8f573d6000803e3d6000fd5b505050506040513d6020811015612ba557600080fd5b50519050612bbd6001600160a01b0382168d846132bb565b604080516001600160a01b038d81168252602082018590528183018c905260ff8a1660608301526080820189905260a08201889052600160c0830152915186928f16917fc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d919081900360e00190a350505b612cc0565b612c418a612a9e8a8a61313f565b816001600160a01b038b167fc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d8b612c788c8c61313f565b604080516001600160a01b03909316835260208301919091528181018c905260ff8a1660608301526080820189905260a08201889052600060c0830152519081900360e00190a35b505060016065555050505050505050565b60cb5481565b612d017f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb133612291565b612d52576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b60005b81811015612d9f57600160cd6000858585818110612d6f57fe5b60209081029290920135835250810191909152604001600020805460ff1916911515919091179055600101612d55565b505050565b612dac611e1c565b15612dfe576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b612e287f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb133612291565b612e79576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038116612ed4576040805162461bcd60e51b815260206004820152601060248201527f4164647265737320697320307830303000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038216600090815260c9602052604090205415611b81576001600160a01b038216600081815260c96020526040902054612f17919083906132bb565b506001600160a01b0316600090815260c96020526040812055565b7f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb181565b7fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd81565b60026065541415612fd2576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555612fdf611e1c565b15613031576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518481526001600160a01b0384811660208301528183018490529151918616917fdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c9181900360600190a2604080517f79cc67900000000000000000000000000000000000000000000000000000000081523360048201526024810183905290516001600160a01b038416916379cc679091604480830192600092919082900301818387803b15801561241257600080fd5b600082820183811015612164576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600082821115613196576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600061323282856001600160a01b031663dd62ed3e30876040518363ffffffff1660e01b815260040180836001600160a01b03168152602001826001600160a01b031681526020019250505060206040518083038186803b15801561320057600080fd5b505afa158015613214573d6000803e3d6000fd5b505050506040513d602081101561322a57600080fd5b5051906130e5565b604080516001600160a01b038616602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b3000000000000000000000000000000000000000000000000000000001790529091506132b59085906136c9565b50505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052612d9f9084906136c9565b3390565b6000828152603360205260409020613357908261377a565b15611b815761336461333b565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526033602052604090206133c0908261378f565b15611b81576133cd61333b565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b613419611e1c565b61346a576040805162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015290519081900360640190fd5b6097805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61349d61333b565b604080516001600160a01b039092168252519081900360200190a1565b60006134c5306137a4565b15905090565b600054610100900460ff16806134e457506134e46134ba565b806134f2575060005460ff16155b61352d5760405162461bcd60e51b815260040180806020018281038252602e815260200180613d11602e913960400191505060405180910390fd5b600054610100900460ff16158015613575576000805460ff197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b61357d6137aa565b611ee66137aa565b61358d611e1c565b156135df576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6097805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861349d61333b565b60006121648383613884565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526132b59085906136c9565b6000612164836001600160a01b0384166138e8565b600061216782613900565b606061371e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166139049092919063ffffffff16565b805190915015612d9f5780806020019051602081101561373d57600080fd5b5051612d9f5760405162461bcd60e51b815260040180806020018281038252602a815260200180613d3f602a913960400191505060405180910390fd5b6000612164836001600160a01b03841661391d565b6000612164836001600160a01b038416613967565b3b151590565b600054610100900460ff16806137c357506137c36134ba565b806137d1575060005460ff16155b61380c5760405162461bcd60e51b815260040180806020018281038252602e815260200180613d11602e913960400191505060405180910390fd5b600054610100900460ff16158015611ee6576000805460ff197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790558015611f1557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905550565b815460009082106138c65760405162461bcd60e51b8152600401808060200182810382526022815260200180613c6a6022913960400191505060405180910390fd5b8260000182815481106138d557fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b60606139138484600085613a4b565b90505b9392505050565b600061392983836138e8565b61395f57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155612167565b506000612167565b60008181526001830160205260408120548015613a415783547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80830191908101906000908790839081106139b857fe5b90600052602060002001549050808760000184815481106139d557fe5b600091825260208083209091019290925582815260018981019092526040902090840190558654879080613a0557fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050612167565b6000915050612167565b606082471015613a8c5760405162461bcd60e51b8152600401808060200182810382526026815260200180613cbb6026913960400191505060405180910390fd5b613a95856137a4565b613ae6576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b60208310613b4357805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101613b06565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114613ba5576040519150601f19603f3d011682016040523d82523d6000602084013e613baa565b606091505b5091509150613bba828286613bc5565b979650505050505050565b60608315613bd4575081613916565b825115613be45782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613c2e578181015183820152602001613c16565b50505050905090810190601f168015613c5b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e74416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b65496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a264697066735822122075b560ce43ce934e445a8ec060c099ba76662debeab146eefbf44fc989f3919164736f6c634300060c0033",
}

// SynapseBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseBridgeMetaData.ABI instead.
var SynapseBridgeABI = SynapseBridgeMetaData.ABI

// Deprecated: Use SynapseBridgeMetaData.Sigs instead.
// SynapseBridgeFuncSigs maps the 4-byte function signature to its string representation.
var SynapseBridgeFuncSigs = SynapseBridgeMetaData.Sigs

// SynapseBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseBridgeMetaData.Bin instead.
var SynapseBridgeBin = SynapseBridgeMetaData.Bin

// DeploySynapseBridge deploys a new Ethereum contract, binding an instance of SynapseBridge to it.
func DeploySynapseBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseBridge, error) {
	parsed, err := SynapseBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseBridge{SynapseBridgeCaller: SynapseBridgeCaller{contract: contract}, SynapseBridgeTransactor: SynapseBridgeTransactor{contract: contract}, SynapseBridgeFilterer: SynapseBridgeFilterer{contract: contract}}, nil
}

// SynapseBridge is an auto generated Go binding around an Ethereum contract.
type SynapseBridge struct {
	SynapseBridgeCaller     // Read-only binding to the contract
	SynapseBridgeTransactor // Write-only binding to the contract
	SynapseBridgeFilterer   // Log filterer for contract events
}

// SynapseBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseBridgeSession struct {
	Contract     *SynapseBridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseBridgeCallerSession struct {
	Contract *SynapseBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SynapseBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseBridgeTransactorSession struct {
	Contract     *SynapseBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SynapseBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseBridgeRaw struct {
	Contract *SynapseBridge // Generic contract binding to access the raw methods on
}

// SynapseBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseBridgeCallerRaw struct {
	Contract *SynapseBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseBridgeTransactorRaw struct {
	Contract *SynapseBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseBridge creates a new instance of SynapseBridge, bound to a specific deployed contract.
func NewSynapseBridge(address common.Address, backend bind.ContractBackend) (*SynapseBridge, error) {
	contract, err := bindSynapseBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseBridge{SynapseBridgeCaller: SynapseBridgeCaller{contract: contract}, SynapseBridgeTransactor: SynapseBridgeTransactor{contract: contract}, SynapseBridgeFilterer: SynapseBridgeFilterer{contract: contract}}, nil
}

// NewSynapseBridgeCaller creates a new read-only instance of SynapseBridge, bound to a specific deployed contract.
func NewSynapseBridgeCaller(address common.Address, caller bind.ContractCaller) (*SynapseBridgeCaller, error) {
	contract, err := bindSynapseBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeCaller{contract: contract}, nil
}

// NewSynapseBridgeTransactor creates a new write-only instance of SynapseBridge, bound to a specific deployed contract.
func NewSynapseBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseBridgeTransactor, error) {
	contract, err := bindSynapseBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTransactor{contract: contract}, nil
}

// NewSynapseBridgeFilterer creates a new log filterer instance of SynapseBridge, bound to a specific deployed contract.
func NewSynapseBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseBridgeFilterer, error) {
	contract, err := bindSynapseBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeFilterer{contract: contract}, nil
}

// bindSynapseBridge binds a generic wrapper to an already deployed contract.
func bindSynapseBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseBridge *SynapseBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseBridge.Contract.SynapseBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseBridge *SynapseBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBridge.Contract.SynapseBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseBridge *SynapseBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseBridge.Contract.SynapseBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseBridge *SynapseBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseBridge *SynapseBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseBridge *SynapseBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SynapseBridge.Contract.DEFAULTADMINROLE(&_SynapseBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SynapseBridge.Contract.DEFAULTADMINROLE(&_SynapseBridge.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeSession) GOVERNANCEROLE() ([32]byte, error) {
	return _SynapseBridge.Contract.GOVERNANCEROLE(&_SynapseBridge.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _SynapseBridge.Contract.GOVERNANCEROLE(&_SynapseBridge.CallOpts)
}

// NODEGROUPROLE is a free data retrieval call binding the contract method 0xf3befd01.
//
// Solidity: function NODEGROUP_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeCaller) NODEGROUPROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "NODEGROUP_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NODEGROUPROLE is a free data retrieval call binding the contract method 0xf3befd01.
//
// Solidity: function NODEGROUP_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeSession) NODEGROUPROLE() ([32]byte, error) {
	return _SynapseBridge.Contract.NODEGROUPROLE(&_SynapseBridge.CallOpts)
}

// NODEGROUPROLE is a free data retrieval call binding the contract method 0xf3befd01.
//
// Solidity: function NODEGROUP_ROLE() view returns(bytes32)
func (_SynapseBridge *SynapseBridgeCallerSession) NODEGROUPROLE() ([32]byte, error) {
	return _SynapseBridge.Contract.NODEGROUPROLE(&_SynapseBridge.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SynapseBridge *SynapseBridgeCaller) WETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "WETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SynapseBridge *SynapseBridgeSession) WETHADDRESS() (common.Address, error) {
	return _SynapseBridge.Contract.WETHADDRESS(&_SynapseBridge.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SynapseBridge *SynapseBridgeCallerSession) WETHADDRESS() (common.Address, error) {
	return _SynapseBridge.Contract.WETHADDRESS(&_SynapseBridge.CallOpts)
}

// BridgeVersion is a free data retrieval call binding the contract method 0xac865626.
//
// Solidity: function bridgeVersion() view returns(uint256)
func (_SynapseBridge *SynapseBridgeCaller) BridgeVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "bridgeVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeVersion is a free data retrieval call binding the contract method 0xac865626.
//
// Solidity: function bridgeVersion() view returns(uint256)
func (_SynapseBridge *SynapseBridgeSession) BridgeVersion() (*big.Int, error) {
	return _SynapseBridge.Contract.BridgeVersion(&_SynapseBridge.CallOpts)
}

// BridgeVersion is a free data retrieval call binding the contract method 0xac865626.
//
// Solidity: function bridgeVersion() view returns(uint256)
func (_SynapseBridge *SynapseBridgeCallerSession) BridgeVersion() (*big.Int, error) {
	return _SynapseBridge.Contract.BridgeVersion(&_SynapseBridge.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseBridge *SynapseBridgeCaller) ChainGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "chainGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseBridge *SynapseBridgeSession) ChainGasAmount() (*big.Int, error) {
	return _SynapseBridge.Contract.ChainGasAmount(&_SynapseBridge.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseBridge *SynapseBridgeCallerSession) ChainGasAmount() (*big.Int, error) {
	return _SynapseBridge.Contract.ChainGasAmount(&_SynapseBridge.CallOpts)
}

// GetFeeBalance is a free data retrieval call binding the contract method 0xc78f6803.
//
// Solidity: function getFeeBalance(address tokenAddress) view returns(uint256)
func (_SynapseBridge *SynapseBridgeCaller) GetFeeBalance(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "getFeeBalance", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBalance is a free data retrieval call binding the contract method 0xc78f6803.
//
// Solidity: function getFeeBalance(address tokenAddress) view returns(uint256)
func (_SynapseBridge *SynapseBridgeSession) GetFeeBalance(tokenAddress common.Address) (*big.Int, error) {
	return _SynapseBridge.Contract.GetFeeBalance(&_SynapseBridge.CallOpts, tokenAddress)
}

// GetFeeBalance is a free data retrieval call binding the contract method 0xc78f6803.
//
// Solidity: function getFeeBalance(address tokenAddress) view returns(uint256)
func (_SynapseBridge *SynapseBridgeCallerSession) GetFeeBalance(tokenAddress common.Address) (*big.Int, error) {
	return _SynapseBridge.Contract.GetFeeBalance(&_SynapseBridge.CallOpts, tokenAddress)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseBridge *SynapseBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseBridge *SynapseBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SynapseBridge.Contract.GetRoleAdmin(&_SynapseBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SynapseBridge *SynapseBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SynapseBridge.Contract.GetRoleAdmin(&_SynapseBridge.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SynapseBridge *SynapseBridgeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SynapseBridge *SynapseBridgeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _SynapseBridge.Contract.GetRoleMember(&_SynapseBridge.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SynapseBridge *SynapseBridgeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _SynapseBridge.Contract.GetRoleMember(&_SynapseBridge.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SynapseBridge *SynapseBridgeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SynapseBridge *SynapseBridgeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _SynapseBridge.Contract.GetRoleMemberCount(&_SynapseBridge.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SynapseBridge *SynapseBridgeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _SynapseBridge.Contract.GetRoleMemberCount(&_SynapseBridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseBridge *SynapseBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseBridge *SynapseBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SynapseBridge.Contract.HasRole(&_SynapseBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SynapseBridge *SynapseBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SynapseBridge.Contract.HasRole(&_SynapseBridge.CallOpts, role, account)
}

// KappaExists is a free data retrieval call binding the contract method 0x2fe87b95.
//
// Solidity: function kappaExists(bytes32 kappa) view returns(bool)
func (_SynapseBridge *SynapseBridgeCaller) KappaExists(opts *bind.CallOpts, kappa [32]byte) (bool, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "kappaExists", kappa)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KappaExists is a free data retrieval call binding the contract method 0x2fe87b95.
//
// Solidity: function kappaExists(bytes32 kappa) view returns(bool)
func (_SynapseBridge *SynapseBridgeSession) KappaExists(kappa [32]byte) (bool, error) {
	return _SynapseBridge.Contract.KappaExists(&_SynapseBridge.CallOpts, kappa)
}

// KappaExists is a free data retrieval call binding the contract method 0x2fe87b95.
//
// Solidity: function kappaExists(bytes32 kappa) view returns(bool)
func (_SynapseBridge *SynapseBridgeCallerSession) KappaExists(kappa [32]byte) (bool, error) {
	return _SynapseBridge.Contract.KappaExists(&_SynapseBridge.CallOpts, kappa)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SynapseBridge *SynapseBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SynapseBridge *SynapseBridgeSession) Paused() (bool, error) {
	return _SynapseBridge.Contract.Paused(&_SynapseBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SynapseBridge *SynapseBridgeCallerSession) Paused() (bool, error) {
	return _SynapseBridge.Contract.Paused(&_SynapseBridge.CallOpts)
}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_SynapseBridge *SynapseBridgeCaller) StartBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseBridge.contract.Call(opts, &out, "startBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_SynapseBridge *SynapseBridgeSession) StartBlockNumber() (*big.Int, error) {
	return _SynapseBridge.Contract.StartBlockNumber(&_SynapseBridge.CallOpts)
}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_SynapseBridge *SynapseBridgeCallerSession) StartBlockNumber() (*big.Int, error) {
	return _SynapseBridge.Contract.StartBlockNumber(&_SynapseBridge.CallOpts)
}

// AddKappas is a paid mutator transaction binding the contract method 0xe7a59998.
//
// Solidity: function addKappas(bytes32[] kappas) returns()
func (_SynapseBridge *SynapseBridgeTransactor) AddKappas(opts *bind.TransactOpts, kappas [][32]byte) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "addKappas", kappas)
}

// AddKappas is a paid mutator transaction binding the contract method 0xe7a59998.
//
// Solidity: function addKappas(bytes32[] kappas) returns()
func (_SynapseBridge *SynapseBridgeSession) AddKappas(kappas [][32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.AddKappas(&_SynapseBridge.TransactOpts, kappas)
}

// AddKappas is a paid mutator transaction binding the contract method 0xe7a59998.
//
// Solidity: function addKappas(bytes32[] kappas) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) AddKappas(kappas [][32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.AddKappas(&_SynapseBridge.TransactOpts, kappas)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeTransactor) Deposit(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "deposit", to, chainId, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeSession) Deposit(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.Deposit(&_SynapseBridge.TransactOpts, to, chainId, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) Deposit(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.Deposit(&_SynapseBridge.TransactOpts, to, chainId, token, amount)
}

// DepositAndSwap is a paid mutator transaction binding the contract method 0xa2a2af0b.
//
// Solidity: function depositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_SynapseBridge *SynapseBridgeTransactor) DepositAndSwap(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "depositAndSwap", to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// DepositAndSwap is a paid mutator transaction binding the contract method 0xa2a2af0b.
//
// Solidity: function depositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_SynapseBridge *SynapseBridgeSession) DepositAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.DepositAndSwap(&_SynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// DepositAndSwap is a paid mutator transaction binding the contract method 0xa2a2af0b.
//
// Solidity: function depositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) DepositAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.DepositAndSwap(&_SynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.GrantRole(&_SynapseBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.GrantRole(&_SynapseBridge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SynapseBridge *SynapseBridgeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SynapseBridge *SynapseBridgeSession) Initialize() (*types.Transaction, error) {
	return _SynapseBridge.Contract.Initialize(&_SynapseBridge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) Initialize() (*types.Transaction, error) {
	return _SynapseBridge.Contract.Initialize(&_SynapseBridge.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x20d7b327.
//
// Solidity: function mint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeTransactor) Mint(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "mint", to, token, amount, fee, kappa)
}

// Mint is a paid mutator transaction binding the contract method 0x20d7b327.
//
// Solidity: function mint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeSession) Mint(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.Mint(&_SynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// Mint is a paid mutator transaction binding the contract method 0x20d7b327.
//
// Solidity: function mint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) Mint(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.Mint(&_SynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// MintAndSwap is a paid mutator transaction binding the contract method 0x17357892.
//
// Solidity: function mintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeTransactor) MintAndSwap(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "mintAndSwap", to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// MintAndSwap is a paid mutator transaction binding the contract method 0x17357892.
//
// Solidity: function mintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeSession) MintAndSwap(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.MintAndSwap(&_SynapseBridge.TransactOpts, to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// MintAndSwap is a paid mutator transaction binding the contract method 0x17357892.
//
// Solidity: function mintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) MintAndSwap(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.MintAndSwap(&_SynapseBridge.TransactOpts, to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SynapseBridge *SynapseBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SynapseBridge *SynapseBridgeSession) Pause() (*types.Transaction, error) {
	return _SynapseBridge.Contract.Pause(&_SynapseBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _SynapseBridge.Contract.Pause(&_SynapseBridge.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeTransactor) Redeem(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "redeem", to, chainId, token, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeSession) Redeem(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.Redeem(&_SynapseBridge.TransactOpts, to, chainId, token, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) Redeem(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.Redeem(&_SynapseBridge.TransactOpts, to, chainId, token, amount)
}

// RedeemAndRemove is a paid mutator transaction binding the contract method 0x36e712ed.
//
// Solidity: function redeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_SynapseBridge *SynapseBridgeTransactor) RedeemAndRemove(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "redeemAndRemove", to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// RedeemAndRemove is a paid mutator transaction binding the contract method 0x36e712ed.
//
// Solidity: function redeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_SynapseBridge *SynapseBridgeSession) RedeemAndRemove(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RedeemAndRemove(&_SynapseBridge.TransactOpts, to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// RedeemAndRemove is a paid mutator transaction binding the contract method 0x36e712ed.
//
// Solidity: function redeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) RedeemAndRemove(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RedeemAndRemove(&_SynapseBridge.TransactOpts, to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// RedeemAndSwap is a paid mutator transaction binding the contract method 0x839ed90a.
//
// Solidity: function redeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_SynapseBridge *SynapseBridgeTransactor) RedeemAndSwap(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "redeemAndSwap", to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// RedeemAndSwap is a paid mutator transaction binding the contract method 0x839ed90a.
//
// Solidity: function redeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_SynapseBridge *SynapseBridgeSession) RedeemAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RedeemAndSwap(&_SynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// RedeemAndSwap is a paid mutator transaction binding the contract method 0x839ed90a.
//
// Solidity: function redeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) RedeemAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RedeemAndSwap(&_SynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// RedeemV2 is a paid mutator transaction binding the contract method 0xa07ed975.
//
// Solidity: function redeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeTransactor) RedeemV2(opts *bind.TransactOpts, to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "redeemV2", to, chainId, token, amount)
}

// RedeemV2 is a paid mutator transaction binding the contract method 0xa07ed975.
//
// Solidity: function redeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeSession) RedeemV2(to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RedeemV2(&_SynapseBridge.TransactOpts, to, chainId, token, amount)
}

// RedeemV2 is a paid mutator transaction binding the contract method 0xa07ed975.
//
// Solidity: function redeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) RedeemV2(to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RedeemV2(&_SynapseBridge.TransactOpts, to, chainId, token, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RenounceRole(&_SynapseBridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RenounceRole(&_SynapseBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RevokeRole(&_SynapseBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.RevokeRole(&_SynapseBridge.TransactOpts, role, account)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeTransactor) SetChainGasAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "setChainGasAmount", amount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeSession) SetChainGasAmount(amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.SetChainGasAmount(&_SynapseBridge.TransactOpts, amount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 amount) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) SetChainGasAmount(amount *big.Int) (*types.Transaction, error) {
	return _SynapseBridge.Contract.SetChainGasAmount(&_SynapseBridge.TransactOpts, amount)
}

// SetWethAddress is a paid mutator transaction binding the contract method 0xa96e2423.
//
// Solidity: function setWethAddress(address _wethAddress) returns()
func (_SynapseBridge *SynapseBridgeTransactor) SetWethAddress(opts *bind.TransactOpts, _wethAddress common.Address) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "setWethAddress", _wethAddress)
}

// SetWethAddress is a paid mutator transaction binding the contract method 0xa96e2423.
//
// Solidity: function setWethAddress(address _wethAddress) returns()
func (_SynapseBridge *SynapseBridgeSession) SetWethAddress(_wethAddress common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.SetWethAddress(&_SynapseBridge.TransactOpts, _wethAddress)
}

// SetWethAddress is a paid mutator transaction binding the contract method 0xa96e2423.
//
// Solidity: function setWethAddress(address _wethAddress) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) SetWethAddress(_wethAddress common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.SetWethAddress(&_SynapseBridge.TransactOpts, _wethAddress)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SynapseBridge *SynapseBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SynapseBridge *SynapseBridgeSession) Unpause() (*types.Transaction, error) {
	return _SynapseBridge.Contract.Unpause(&_SynapseBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _SynapseBridge.Contract.Unpause(&_SynapseBridge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1cf5f07f.
//
// Solidity: function withdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "withdraw", to, token, amount, fee, kappa)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1cf5f07f.
//
// Solidity: function withdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeSession) Withdraw(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.Withdraw(&_SynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1cf5f07f.
//
// Solidity: function withdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) Withdraw(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.Withdraw(&_SynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// WithdrawAndRemove is a paid mutator transaction binding the contract method 0xd57eafac.
//
// Solidity: function withdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeTransactor) WithdrawAndRemove(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "withdrawAndRemove", to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// WithdrawAndRemove is a paid mutator transaction binding the contract method 0xd57eafac.
//
// Solidity: function withdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeSession) WithdrawAndRemove(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.WithdrawAndRemove(&_SynapseBridge.TransactOpts, to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// WithdrawAndRemove is a paid mutator transaction binding the contract method 0xd57eafac.
//
// Solidity: function withdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) WithdrawAndRemove(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _SynapseBridge.Contract.WithdrawAndRemove(&_SynapseBridge.TransactOpts, to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address token, address to) returns()
func (_SynapseBridge *SynapseBridgeTransactor) WithdrawFees(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _SynapseBridge.contract.Transact(opts, "withdrawFees", token, to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address token, address to) returns()
func (_SynapseBridge *SynapseBridgeSession) WithdrawFees(token common.Address, to common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.WithdrawFees(&_SynapseBridge.TransactOpts, token, to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address token, address to) returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) WithdrawFees(token common.Address, to common.Address) (*types.Transaction, error) {
	return _SynapseBridge.Contract.WithdrawFees(&_SynapseBridge.TransactOpts, token, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SynapseBridge *SynapseBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SynapseBridge *SynapseBridgeSession) Receive() (*types.Transaction, error) {
	return _SynapseBridge.Contract.Receive(&_SynapseBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SynapseBridge *SynapseBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _SynapseBridge.Contract.Receive(&_SynapseBridge.TransactOpts)
}

// SynapseBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SynapseBridge contract.
type SynapseBridgePausedIterator struct {
	Event *SynapseBridgePaused // Event containing the contract specifics and raw log

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
func (it *SynapseBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgePaused)
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
		it.Event = new(SynapseBridgePaused)
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
func (it *SynapseBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgePaused represents a Paused event raised by the SynapseBridge contract.
type SynapseBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SynapseBridge *SynapseBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*SynapseBridgePausedIterator, error) {

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SynapseBridgePausedIterator{contract: _SynapseBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SynapseBridge *SynapseBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SynapseBridgePaused) (event.Subscription, error) {

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgePaused)
				if err := _SynapseBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SynapseBridge *SynapseBridgeFilterer) ParsePaused(log types.Log) (*SynapseBridgePaused, error) {
	event := new(SynapseBridgePaused)
	if err := _SynapseBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SynapseBridge contract.
type SynapseBridgeRoleAdminChangedIterator struct {
	Event *SynapseBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeRoleAdminChanged)
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
		it.Event = new(SynapseBridgeRoleAdminChanged)
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
func (it *SynapseBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the SynapseBridge contract.
type SynapseBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SynapseBridge *SynapseBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SynapseBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeRoleAdminChangedIterator{contract: _SynapseBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SynapseBridge *SynapseBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SynapseBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeRoleAdminChanged)
				if err := _SynapseBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SynapseBridge *SynapseBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*SynapseBridgeRoleAdminChanged, error) {
	event := new(SynapseBridgeRoleAdminChanged)
	if err := _SynapseBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SynapseBridge contract.
type SynapseBridgeRoleGrantedIterator struct {
	Event *SynapseBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeRoleGranted)
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
		it.Event = new(SynapseBridgeRoleGranted)
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
func (it *SynapseBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeRoleGranted represents a RoleGranted event raised by the SynapseBridge contract.
type SynapseBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseBridge *SynapseBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SynapseBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeRoleGrantedIterator{contract: _SynapseBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseBridge *SynapseBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SynapseBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeRoleGranted)
				if err := _SynapseBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SynapseBridge *SynapseBridgeFilterer) ParseRoleGranted(log types.Log) (*SynapseBridgeRoleGranted, error) {
	event := new(SynapseBridgeRoleGranted)
	if err := _SynapseBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SynapseBridge contract.
type SynapseBridgeRoleRevokedIterator struct {
	Event *SynapseBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeRoleRevoked)
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
		it.Event = new(SynapseBridgeRoleRevoked)
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
func (it *SynapseBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeRoleRevoked represents a RoleRevoked event raised by the SynapseBridge contract.
type SynapseBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseBridge *SynapseBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SynapseBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeRoleRevokedIterator{contract: _SynapseBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SynapseBridge *SynapseBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SynapseBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeRoleRevoked)
				if err := _SynapseBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SynapseBridge *SynapseBridgeFilterer) ParseRoleRevoked(log types.Log) (*SynapseBridgeRoleRevoked, error) {
	event := new(SynapseBridgeRoleRevoked)
	if err := _SynapseBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenDepositIterator is returned from FilterTokenDeposit and is used to iterate over the raw logs and unpacked data for TokenDeposit events raised by the SynapseBridge contract.
type SynapseBridgeTokenDepositIterator struct {
	Event *SynapseBridgeTokenDeposit // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenDeposit)
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
		it.Event = new(SynapseBridgeTokenDeposit)
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
func (it *SynapseBridgeTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenDeposit represents a TokenDeposit event raised by the SynapseBridge contract.
type SynapseBridgeTokenDeposit struct {
	To      common.Address
	ChainId *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposit is a free log retrieval operation binding the contract event 0xda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b.
//
// Solidity: event TokenDeposit(address indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenDeposit(opts *bind.FilterOpts, to []common.Address) (*SynapseBridgeTokenDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenDeposit", toRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenDepositIterator{contract: _SynapseBridge.contract, event: "TokenDeposit", logs: logs, sub: sub}, nil
}

// WatchTokenDeposit is a free log subscription operation binding the contract event 0xda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b.
//
// Solidity: event TokenDeposit(address indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenDeposit(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenDeposit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenDeposit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenDeposit)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
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

// ParseTokenDeposit is a log parse operation binding the contract event 0xda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b.
//
// Solidity: event TokenDeposit(address indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenDeposit(log types.Log) (*SynapseBridgeTokenDeposit, error) {
	event := new(SynapseBridgeTokenDeposit)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenDepositAndSwapIterator is returned from FilterTokenDepositAndSwap and is used to iterate over the raw logs and unpacked data for TokenDepositAndSwap events raised by the SynapseBridge contract.
type SynapseBridgeTokenDepositAndSwapIterator struct {
	Event *SynapseBridgeTokenDepositAndSwap // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenDepositAndSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenDepositAndSwap)
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
		it.Event = new(SynapseBridgeTokenDepositAndSwap)
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
func (it *SynapseBridgeTokenDepositAndSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenDepositAndSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenDepositAndSwap represents a TokenDepositAndSwap event raised by the SynapseBridge contract.
type SynapseBridgeTokenDepositAndSwap struct {
	To             common.Address
	ChainId        *big.Int
	Token          common.Address
	Amount         *big.Int
	TokenIndexFrom uint8
	TokenIndexTo   uint8
	MinDy          *big.Int
	Deadline       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenDepositAndSwap is a free log retrieval operation binding the contract event 0x79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f.
//
// Solidity: event TokenDepositAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenDepositAndSwap(opts *bind.FilterOpts, to []common.Address) (*SynapseBridgeTokenDepositAndSwapIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenDepositAndSwap", toRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenDepositAndSwapIterator{contract: _SynapseBridge.contract, event: "TokenDepositAndSwap", logs: logs, sub: sub}, nil
}

// WatchTokenDepositAndSwap is a free log subscription operation binding the contract event 0x79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f.
//
// Solidity: event TokenDepositAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenDepositAndSwap(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenDepositAndSwap, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenDepositAndSwap", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenDepositAndSwap)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenDepositAndSwap", log); err != nil {
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

// ParseTokenDepositAndSwap is a log parse operation binding the contract event 0x79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f.
//
// Solidity: event TokenDepositAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenDepositAndSwap(log types.Log) (*SynapseBridgeTokenDepositAndSwap, error) {
	event := new(SynapseBridgeTokenDepositAndSwap)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenDepositAndSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenMintIterator is returned from FilterTokenMint and is used to iterate over the raw logs and unpacked data for TokenMint events raised by the SynapseBridge contract.
type SynapseBridgeTokenMintIterator struct {
	Event *SynapseBridgeTokenMint // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenMint)
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
		it.Event = new(SynapseBridgeTokenMint)
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
func (it *SynapseBridgeTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenMint represents a TokenMint event raised by the SynapseBridge contract.
type SynapseBridgeTokenMint struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Fee    *big.Int
	Kappa  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenMint is a free log retrieval operation binding the contract event 0xbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd38.
//
// Solidity: event TokenMint(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenMint(opts *bind.FilterOpts, to []common.Address, kappa [][32]byte) (*SynapseBridgeTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenMint", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenMintIterator{contract: _SynapseBridge.contract, event: "TokenMint", logs: logs, sub: sub}, nil
}

// WatchTokenMint is a free log subscription operation binding the contract event 0xbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd38.
//
// Solidity: event TokenMint(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenMint(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenMint, to []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenMint", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenMint)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenMint", log); err != nil {
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

// ParseTokenMint is a log parse operation binding the contract event 0xbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd38.
//
// Solidity: event TokenMint(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenMint(log types.Log) (*SynapseBridgeTokenMint, error) {
	event := new(SynapseBridgeTokenMint)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenMintAndSwapIterator is returned from FilterTokenMintAndSwap and is used to iterate over the raw logs and unpacked data for TokenMintAndSwap events raised by the SynapseBridge contract.
type SynapseBridgeTokenMintAndSwapIterator struct {
	Event *SynapseBridgeTokenMintAndSwap // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenMintAndSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenMintAndSwap)
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
		it.Event = new(SynapseBridgeTokenMintAndSwap)
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
func (it *SynapseBridgeTokenMintAndSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenMintAndSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenMintAndSwap represents a TokenMintAndSwap event raised by the SynapseBridge contract.
type SynapseBridgeTokenMintAndSwap struct {
	To             common.Address
	Token          common.Address
	Amount         *big.Int
	Fee            *big.Int
	TokenIndexFrom uint8
	TokenIndexTo   uint8
	MinDy          *big.Int
	Deadline       *big.Int
	SwapSuccess    bool
	Kappa          [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenMintAndSwap is a free log retrieval operation binding the contract event 0x4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c65.
//
// Solidity: event TokenMintAndSwap(address indexed to, address token, uint256 amount, uint256 fee, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bool swapSuccess, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenMintAndSwap(opts *bind.FilterOpts, to []common.Address, kappa [][32]byte) (*SynapseBridgeTokenMintAndSwapIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenMintAndSwap", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenMintAndSwapIterator{contract: _SynapseBridge.contract, event: "TokenMintAndSwap", logs: logs, sub: sub}, nil
}

// WatchTokenMintAndSwap is a free log subscription operation binding the contract event 0x4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c65.
//
// Solidity: event TokenMintAndSwap(address indexed to, address token, uint256 amount, uint256 fee, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bool swapSuccess, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenMintAndSwap(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenMintAndSwap, to []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenMintAndSwap", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenMintAndSwap)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenMintAndSwap", log); err != nil {
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

// ParseTokenMintAndSwap is a log parse operation binding the contract event 0x4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c65.
//
// Solidity: event TokenMintAndSwap(address indexed to, address token, uint256 amount, uint256 fee, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bool swapSuccess, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenMintAndSwap(log types.Log) (*SynapseBridgeTokenMintAndSwap, error) {
	event := new(SynapseBridgeTokenMintAndSwap)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenMintAndSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenRedeemIterator is returned from FilterTokenRedeem and is used to iterate over the raw logs and unpacked data for TokenRedeem events raised by the SynapseBridge contract.
type SynapseBridgeTokenRedeemIterator struct {
	Event *SynapseBridgeTokenRedeem // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenRedeem)
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
		it.Event = new(SynapseBridgeTokenRedeem)
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
func (it *SynapseBridgeTokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenRedeem represents a TokenRedeem event raised by the SynapseBridge contract.
type SynapseBridgeTokenRedeem struct {
	To      common.Address
	ChainId *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeem is a free log retrieval operation binding the contract event 0xdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c.
//
// Solidity: event TokenRedeem(address indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenRedeem(opts *bind.FilterOpts, to []common.Address) (*SynapseBridgeTokenRedeemIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenRedeem", toRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenRedeemIterator{contract: _SynapseBridge.contract, event: "TokenRedeem", logs: logs, sub: sub}, nil
}

// WatchTokenRedeem is a free log subscription operation binding the contract event 0xdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c.
//
// Solidity: event TokenRedeem(address indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenRedeem(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenRedeem, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenRedeem", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenRedeem)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenRedeem", log); err != nil {
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

// ParseTokenRedeem is a log parse operation binding the contract event 0xdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c.
//
// Solidity: event TokenRedeem(address indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenRedeem(log types.Log) (*SynapseBridgeTokenRedeem, error) {
	event := new(SynapseBridgeTokenRedeem)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenRedeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenRedeemAndRemoveIterator is returned from FilterTokenRedeemAndRemove and is used to iterate over the raw logs and unpacked data for TokenRedeemAndRemove events raised by the SynapseBridge contract.
type SynapseBridgeTokenRedeemAndRemoveIterator struct {
	Event *SynapseBridgeTokenRedeemAndRemove // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenRedeemAndRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenRedeemAndRemove)
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
		it.Event = new(SynapseBridgeTokenRedeemAndRemove)
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
func (it *SynapseBridgeTokenRedeemAndRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenRedeemAndRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenRedeemAndRemove represents a TokenRedeemAndRemove event raised by the SynapseBridge contract.
type SynapseBridgeTokenRedeemAndRemove struct {
	To             common.Address
	ChainId        *big.Int
	Token          common.Address
	Amount         *big.Int
	SwapTokenIndex uint8
	SwapMinAmount  *big.Int
	SwapDeadline   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeemAndRemove is a free log retrieval operation binding the contract event 0x9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c.
//
// Solidity: event TokenRedeemAndRemove(address indexed to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenRedeemAndRemove(opts *bind.FilterOpts, to []common.Address) (*SynapseBridgeTokenRedeemAndRemoveIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenRedeemAndRemove", toRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenRedeemAndRemoveIterator{contract: _SynapseBridge.contract, event: "TokenRedeemAndRemove", logs: logs, sub: sub}, nil
}

// WatchTokenRedeemAndRemove is a free log subscription operation binding the contract event 0x9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c.
//
// Solidity: event TokenRedeemAndRemove(address indexed to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenRedeemAndRemove(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenRedeemAndRemove, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenRedeemAndRemove", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenRedeemAndRemove)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenRedeemAndRemove", log); err != nil {
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

// ParseTokenRedeemAndRemove is a log parse operation binding the contract event 0x9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c.
//
// Solidity: event TokenRedeemAndRemove(address indexed to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenRedeemAndRemove(log types.Log) (*SynapseBridgeTokenRedeemAndRemove, error) {
	event := new(SynapseBridgeTokenRedeemAndRemove)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenRedeemAndRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenRedeemAndSwapIterator is returned from FilterTokenRedeemAndSwap and is used to iterate over the raw logs and unpacked data for TokenRedeemAndSwap events raised by the SynapseBridge contract.
type SynapseBridgeTokenRedeemAndSwapIterator struct {
	Event *SynapseBridgeTokenRedeemAndSwap // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenRedeemAndSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenRedeemAndSwap)
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
		it.Event = new(SynapseBridgeTokenRedeemAndSwap)
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
func (it *SynapseBridgeTokenRedeemAndSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenRedeemAndSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenRedeemAndSwap represents a TokenRedeemAndSwap event raised by the SynapseBridge contract.
type SynapseBridgeTokenRedeemAndSwap struct {
	To             common.Address
	ChainId        *big.Int
	Token          common.Address
	Amount         *big.Int
	TokenIndexFrom uint8
	TokenIndexTo   uint8
	MinDy          *big.Int
	Deadline       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeemAndSwap is a free log retrieval operation binding the contract event 0x91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc319425.
//
// Solidity: event TokenRedeemAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenRedeemAndSwap(opts *bind.FilterOpts, to []common.Address) (*SynapseBridgeTokenRedeemAndSwapIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenRedeemAndSwap", toRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenRedeemAndSwapIterator{contract: _SynapseBridge.contract, event: "TokenRedeemAndSwap", logs: logs, sub: sub}, nil
}

// WatchTokenRedeemAndSwap is a free log subscription operation binding the contract event 0x91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc319425.
//
// Solidity: event TokenRedeemAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenRedeemAndSwap(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenRedeemAndSwap, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenRedeemAndSwap", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenRedeemAndSwap)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenRedeemAndSwap", log); err != nil {
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

// ParseTokenRedeemAndSwap is a log parse operation binding the contract event 0x91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc319425.
//
// Solidity: event TokenRedeemAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenRedeemAndSwap(log types.Log) (*SynapseBridgeTokenRedeemAndSwap, error) {
	event := new(SynapseBridgeTokenRedeemAndSwap)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenRedeemAndSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenRedeemV2Iterator is returned from FilterTokenRedeemV2 and is used to iterate over the raw logs and unpacked data for TokenRedeemV2 events raised by the SynapseBridge contract.
type SynapseBridgeTokenRedeemV2Iterator struct {
	Event *SynapseBridgeTokenRedeemV2 // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenRedeemV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenRedeemV2)
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
		it.Event = new(SynapseBridgeTokenRedeemV2)
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
func (it *SynapseBridgeTokenRedeemV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenRedeemV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenRedeemV2 represents a TokenRedeemV2 event raised by the SynapseBridge contract.
type SynapseBridgeTokenRedeemV2 struct {
	To      [32]byte
	ChainId *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeemV2 is a free log retrieval operation binding the contract event 0x8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499.
//
// Solidity: event TokenRedeemV2(bytes32 indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenRedeemV2(opts *bind.FilterOpts, to [][32]byte) (*SynapseBridgeTokenRedeemV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenRedeemV2", toRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenRedeemV2Iterator{contract: _SynapseBridge.contract, event: "TokenRedeemV2", logs: logs, sub: sub}, nil
}

// WatchTokenRedeemV2 is a free log subscription operation binding the contract event 0x8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499.
//
// Solidity: event TokenRedeemV2(bytes32 indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenRedeemV2(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenRedeemV2, to [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenRedeemV2", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenRedeemV2)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenRedeemV2", log); err != nil {
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

// ParseTokenRedeemV2 is a log parse operation binding the contract event 0x8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499.
//
// Solidity: event TokenRedeemV2(bytes32 indexed to, uint256 chainId, address token, uint256 amount)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenRedeemV2(log types.Log) (*SynapseBridgeTokenRedeemV2, error) {
	event := new(SynapseBridgeTokenRedeemV2)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenRedeemV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the SynapseBridge contract.
type SynapseBridgeTokenWithdrawIterator struct {
	Event *SynapseBridgeTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenWithdraw)
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
		it.Event = new(SynapseBridgeTokenWithdraw)
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
func (it *SynapseBridgeTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenWithdraw represents a TokenWithdraw event raised by the SynapseBridge contract.
type SynapseBridgeTokenWithdraw struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Fee    *big.Int
	Kappa  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26.
//
// Solidity: event TokenWithdraw(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenWithdraw(opts *bind.FilterOpts, to []common.Address, kappa [][32]byte) (*SynapseBridgeTokenWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenWithdraw", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenWithdrawIterator{contract: _SynapseBridge.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26.
//
// Solidity: event TokenWithdraw(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenWithdraw, to []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenWithdraw", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenWithdraw)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26.
//
// Solidity: event TokenWithdraw(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenWithdraw(log types.Log) (*SynapseBridgeTokenWithdraw, error) {
	event := new(SynapseBridgeTokenWithdraw)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeTokenWithdrawAndRemoveIterator is returned from FilterTokenWithdrawAndRemove and is used to iterate over the raw logs and unpacked data for TokenWithdrawAndRemove events raised by the SynapseBridge contract.
type SynapseBridgeTokenWithdrawAndRemoveIterator struct {
	Event *SynapseBridgeTokenWithdrawAndRemove // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeTokenWithdrawAndRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeTokenWithdrawAndRemove)
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
		it.Event = new(SynapseBridgeTokenWithdrawAndRemove)
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
func (it *SynapseBridgeTokenWithdrawAndRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeTokenWithdrawAndRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeTokenWithdrawAndRemove represents a TokenWithdrawAndRemove event raised by the SynapseBridge contract.
type SynapseBridgeTokenWithdrawAndRemove struct {
	To             common.Address
	Token          common.Address
	Amount         *big.Int
	Fee            *big.Int
	SwapTokenIndex uint8
	SwapMinAmount  *big.Int
	SwapDeadline   *big.Int
	SwapSuccess    bool
	Kappa          [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawAndRemove is a free log retrieval operation binding the contract event 0xc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d.
//
// Solidity: event TokenWithdrawAndRemove(address indexed to, address token, uint256 amount, uint256 fee, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bool swapSuccess, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) FilterTokenWithdrawAndRemove(opts *bind.FilterOpts, to []common.Address, kappa [][32]byte) (*SynapseBridgeTokenWithdrawAndRemoveIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "TokenWithdrawAndRemove", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeTokenWithdrawAndRemoveIterator{contract: _SynapseBridge.contract, event: "TokenWithdrawAndRemove", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawAndRemove is a free log subscription operation binding the contract event 0xc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d.
//
// Solidity: event TokenWithdrawAndRemove(address indexed to, address token, uint256 amount, uint256 fee, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bool swapSuccess, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) WatchTokenWithdrawAndRemove(opts *bind.WatchOpts, sink chan<- *SynapseBridgeTokenWithdrawAndRemove, to []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "TokenWithdrawAndRemove", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeTokenWithdrawAndRemove)
				if err := _SynapseBridge.contract.UnpackLog(event, "TokenWithdrawAndRemove", log); err != nil {
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

// ParseTokenWithdrawAndRemove is a log parse operation binding the contract event 0xc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d.
//
// Solidity: event TokenWithdrawAndRemove(address indexed to, address token, uint256 amount, uint256 fee, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bool swapSuccess, bytes32 indexed kappa)
func (_SynapseBridge *SynapseBridgeFilterer) ParseTokenWithdrawAndRemove(log types.Log) (*SynapseBridgeTokenWithdrawAndRemove, error) {
	event := new(SynapseBridgeTokenWithdrawAndRemove)
	if err := _SynapseBridge.contract.UnpackLog(event, "TokenWithdrawAndRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SynapseBridge contract.
type SynapseBridgeUnpausedIterator struct {
	Event *SynapseBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *SynapseBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBridgeUnpaused)
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
		it.Event = new(SynapseBridgeUnpaused)
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
func (it *SynapseBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBridgeUnpaused represents a Unpaused event raised by the SynapseBridge contract.
type SynapseBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SynapseBridge *SynapseBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SynapseBridgeUnpausedIterator, error) {

	logs, sub, err := _SynapseBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SynapseBridgeUnpausedIterator{contract: _SynapseBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SynapseBridge *SynapseBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SynapseBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _SynapseBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBridgeUnpaused)
				if err := _SynapseBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SynapseBridge *SynapseBridgeFilterer) ParseUnpaused(log types.Log) (*SynapseBridgeUnpaused, error) {
	event := new(SynapseBridgeUnpaused)
	if err := _SynapseBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeMetaData contains all meta data concerning the TestSynapseBridge contract.
var TestSynapseBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TokenDepositAndSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"TokenMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"swapSuccess\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"TokenMintAndSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRedeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"}],\"name\":\"TokenRedeemAndRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"TokenRedeemAndSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRedeemV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"swapSuccess\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"TokenWithdrawAndRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODEGROUP_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH_ADDRESS\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"kappas\",\"type\":\"bytes32[]\"}],\"name\":\"addKappas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"depositAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"kappaExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"contractISwap\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"mintAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"}],\"name\":\"redeemAndRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"redeemAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_wethAddress\",\"type\":\"address\"}],\"name\":\"setWethAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"testDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"testDepositAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"testMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"contractISwap\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"testMintAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"testRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"}],\"name\":\"testRedeemAndRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"testRedeemAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20Burnable\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"testRedeemV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"testWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"contractISwap\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"testWithdrawAndRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"contractISwap\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"swapTokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"swapMinAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"f36c8f5c": "GOVERNANCE_ROLE()",
		"f3befd01": "NODEGROUP_ROLE()",
		"040141e5": "WETH_ADDRESS()",
		"e7a59998": "addKappas(bytes32[])",
		"ac865626": "bridgeVersion()",
		"e00a83e0": "chainGasAmount()",
		"90d25074": "deposit(address,uint256,address,uint256)",
		"a2a2af0b": "depositAndSwap(address,uint256,address,uint256,uint8,uint8,uint256,uint256)",
		"c78f6803": "getFeeBalance(address)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"8129fc1c": "initialize()",
		"2fe87b95": "kappaExists(bytes32)",
		"20d7b327": "mint(address,address,uint256,uint256,bytes32)",
		"17357892": "mintAndSwap(address,address,uint256,uint256,address,uint8,uint8,uint256,uint256,bytes32)",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"f3f094a1": "redeem(address,uint256,address,uint256)",
		"36e712ed": "redeemAndRemove(address,uint256,address,uint256,uint8,uint256,uint256)",
		"839ed90a": "redeemAndSwap(address,uint256,address,uint256,uint8,uint8,uint256,uint256)",
		"a07ed975": "redeemV2(bytes32,uint256,address,uint256)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"b250fe6b": "setChainGasAmount(uint256)",
		"a96e2423": "setWethAddress(address)",
		"498a4c2d": "startBlockNumber()",
		"ad863232": "testDeposit(address,uint256,address,uint256)",
		"09faf02b": "testDepositAndSwap(address,uint256,address,uint256,uint8,uint8,uint256,uint256)",
		"f3320221": "testMint(address,address,uint256,uint256,bytes32)",
		"328a9dc3": "testMintAndSwap(address,address,uint256,uint256,address,uint8,uint8,uint256,uint256,bytes32)",
		"e072f5cd": "testRedeem(address,uint256,address,uint256)",
		"b814ff5b": "testRedeemAndRemove(address,uint256,address,uint256,uint8,uint256,uint256)",
		"dbb176f4": "testRedeemAndSwap(address,uint256,address,uint256,uint8,uint8,uint256,uint256)",
		"98c3e142": "testRedeemV2(bytes32,uint256,address,uint256)",
		"ffafc49b": "testWithdraw(address,address,uint256,uint256,bytes32)",
		"e2176fbe": "testWithdrawAndRemove(address,address,uint256,uint256,address,uint8,uint256,uint256,bytes32)",
		"3f4ba83a": "unpause()",
		"1cf5f07f": "withdraw(address,address,uint256,uint256,bytes32)",
		"d57eafac": "withdrawAndRemove(address,address,uint256,uint256,address,uint8,uint256,uint256,bytes32)",
		"f2555278": "withdrawFees(address,address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50614613806100206000396000f3fe6080604052600436106102eb5760003560e01c8063a217fddf11610184578063dbb176f4116100d6578063f25552781161008a578063f3befd0111610064578063f3befd0114610d4d578063f3f094a114610d62578063ffafc49b14610da9576102f2565b8063f255527814610cae578063f332022114610ce9578063f36c8f5c14610d38576102f2565b8063e072f5cd116100bb578063e072f5cd14610b7e578063e2176fbe14610bc5578063e7a5999814610c31576102f2565b8063dbb176f414610b02578063e00a83e014610b69576102f2565b8063b250fe6b11610138578063ca15c87311610112578063ca15c87314610a33578063d547741f14610a5d578063d57eafac14610a96576102f2565b8063b250fe6b1461097a578063b814ff5b146109a4578063c78f680314610a00576102f2565b8063a96e242311610169578063a96e2423146108eb578063ac8656261461091e578063ad86323214610933576102f2565b8063a217fddf1461086f578063a2a2af0b14610884576102f2565b80633f4ba83a1161023d5780638456cb59116101f157806391d14854116101cb57806391d14854146107ac57806398c3e142146107e5578063a07ed9751461082a576102f2565b80638456cb59146107205780639010d07c1461073557806390d2507414610765576102f2565b80635c975abb116102225780635c975abb1461068f5780638129fc1c146106a4578063839ed90a146106b9576102f2565b80633f4ba83a14610665578063498a4c2d1461067a576102f2565b8063248a9ca31161029f578063328a9dc311610279578063328a9dc31461055957806336568abe146105d057806336e712ed14610609576102f2565b8063248a9ca3146104a65780632f2ff15d146104e25780632fe87b951461051b576102f2565b806317357892116102d057806317357892146103915780631cf5f07f1461040857806320d7b32714610457576102f2565b8063040141e5146102f757806309faf02b14610328576102f2565b366102f257005b600080fd5b34801561030357600080fd5b5061030c610df8565b604080516001600160a01b039092168252519081900360200190f35b34801561033457600080fd5b5061038f600480360361010081101561034c57600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff608082013581169160a08101359091169060c08101359060e00135610e07565b005b34801561039d57600080fd5b5061038f60048036036101408110156103b557600080fd5b506001600160a01b03813581169160208101358216916040820135916060810135916080820135169060ff60a082013581169160c08101359091169060e081013590610100810135906101200135610e7d565b34801561041457600080fd5b5061038f600480360360a081101561042b57600080fd5b506001600160a01b038135811691602081013590911690604081013590606081013590608001356117d7565b34801561046357600080fd5b5061038f600480360360a081101561047a57600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060800135611c20565b3480156104b257600080fd5b506104d0600480360360208110156104c957600080fd5b5035611fc6565b60408051918252519081900360200190f35b3480156104ee57600080fd5b5061038f6004803603604081101561050557600080fd5b50803590602001356001600160a01b0316611fdb565b34801561052757600080fd5b506105456004803603602081101561053e57600080fd5b5035612047565b604080519115158252519081900360200190f35b34801561056557600080fd5b5061038f600480360361014081101561057d57600080fd5b506001600160a01b03813581169160208101358216916040820135916060810135916080820135169060ff60a082013581169160c08101359091169060e08101359061010081013590610120013561205c565b3480156105dc57600080fd5b5061038f600480360360408110156105f357600080fd5b50803590602001356001600160a01b03166120df565b34801561061557600080fd5b5061038f600480360360e081101561062c57600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff6080820135169060a08101359060c00135612140565b34801561067157600080fd5b5061038f6122d6565b34801561068657600080fd5b506104d061235b565b34801561069b57600080fd5b50610545612361565b3480156106b057600080fd5b5061038f61236a565b3480156106c557600080fd5b5061038f60048036036101008110156106dd57600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff608082013581169160a08101359091169060c08101359060e0013561245d565b34801561072c57600080fd5b5061038f61260e565b34801561074157600080fd5b5061030c6004803603604081101561075857600080fd5b5080359060200135612691565b34801561077157600080fd5b5061038f6004803603608081101561078857600080fd5b506001600160a01b03813581169160208101359160408201351690606001356126b2565b3480156107b857600080fd5b50610545600480360360408110156107cf57600080fd5b50803590602001356001600160a01b03166127d6565b3480156107f157600080fd5b5061038f6004803603608081101561080857600080fd5b508035906020810135906001600160a01b0360408201351690606001356127ee565b34801561083657600080fd5b5061038f6004803603608081101561084d57600080fd5b508035906020810135906001600160a01b03604082013516906060013561283f565b34801561087b57600080fd5b506104d06129cb565b34801561089057600080fd5b5061038f60048036036101008110156108a857600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff608082013581169160a08101359091169060c08101359060e001356129d0565b3480156108f757600080fd5b5061038f6004803603602081101561090e57600080fd5b50356001600160a01b0316612b08565b34801561092a57600080fd5b506104d0612b9e565b34801561093f57600080fd5b5061038f6004803603608081101561095657600080fd5b506001600160a01b0381358116916020810135916040820135169060600135612ba3565b34801561098657600080fd5b5061038f6004803603602081101561099d57600080fd5b5035612bf6565b3480156109b057600080fd5b5061038f600480360360e08110156109c757600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff6080820135169060a08101359060c00135612c76565b348015610a0c57600080fd5b506104d060048036036020811015610a2357600080fd5b50356001600160a01b0316612ce3565b348015610a3f57600080fd5b506104d060048036036020811015610a5657600080fd5b5035612cfe565b348015610a6957600080fd5b5061038f60048036036040811015610a8057600080fd5b50803590602001356001600160a01b0316612d15565b348015610aa257600080fd5b5061038f6004803603610120811015610aba57600080fd5b506001600160a01b03813581169160208101358216916040820135916060810135916080820135169060ff60a0820135169060c08101359060e0810135906101000135612d6e565b348015610b0e57600080fd5b5061038f6004803603610100811015610b2657600080fd5b506001600160a01b0381358116916020810135916040820135169060608101359060ff608082013581169160a08101359091169060c08101359060e00135613327565b348015610b7557600080fd5b506104d061339d565b348015610b8a57600080fd5b5061038f60048036036080811015610ba157600080fd5b506001600160a01b03813581169160208101359160408201351690606001356133a3565b348015610bd157600080fd5b5061038f6004803603610120811015610be957600080fd5b506001600160a01b03813581169160208101358216916040820135916060810135916080820135169060ff60a0820135169060c08101359060e08101359061010001356133f6565b348015610c3d57600080fd5b5061038f60048036036020811015610c5457600080fd5b810190602081018135640100000000811115610c6f57600080fd5b820183602082011115610c8157600080fd5b80359060200191846020830284011164010000000083111715610ca357600080fd5b50909250905061346f565b348015610cba57600080fd5b5061038f60048036036040811015610cd157600080fd5b506001600160a01b038135811691602001351661353c565b348015610cf557600080fd5b5061038f600480360360a0811015610d0c57600080fd5b506001600160a01b038135811691602081013590911690604081013590606081013590608001356136ca565b348015610d4457600080fd5b506104d0613721565b348015610d5957600080fd5b506104d0613745565b348015610d6e57600080fd5b5061038f60048036036080811015610d8557600080fd5b506001600160a01b0381358116916020810135916040820135169060600135613769565b348015610db557600080fd5b5061038f600480360360a0811015610dcc57600080fd5b506001600160a01b038135811691602081013590911690604081013590606081013590608001356138d4565b60cc546001600160a01b031681565b604080518881526001600160a01b03888116602083015281830188905260ff80881660608401528616608083015260a0820185905260c082018490529151918a16917f79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f9181900360e00190a25050505050505050565b60026065541415610ed5576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555610ee2612361565b15610f34576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b610f5e7fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd336127d6565b610faf576040805162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f742061206e6f64652067726f7570000000000000604482015290519081900360640190fd5b868811611003576040805162461bcd60e51b815260206004820152601f60248201527f416d6f756e74206d7573742062652067726561746572207468616e2066656500604482015290519081900360640190fd5b600081815260cd602052604090205460ff1615611067576040805162461bcd60e51b815260206004820152601860248201527f4b6170706120697320616c72656164792070726573656e740000000000000000604482015290519081900360640190fd5b600081815260cd60209081526040808320805460ff191660011790556001600160a01b038c16835260c99091529020546110a1908861392b565b6001600160a01b038a16600090815260c9602052604090205560cb54158015906110cc575060cb5447115b156111255760cb546040516001600160a01b038c169190600081818185875af1925050503d806000811461111c576040519150601f19603f3d011682016040523d82523d6000602084013e611121565b606091505b5050505b60006001600160a01b03871663a95b089f87876111428d8d613985565b6040518463ffffffff1660e01b8152600401808460ff1681526020018360ff168152602001828152602001935050505060206040518083038186803b15801561118a57600080fd5b505afa15801561119e573d6000803e3d6000fd5b505050506040513d60208110156111b457600080fd5b505190508381106116a757604080517f40c10f19000000000000000000000000000000000000000000000000000000008152306004820152602481018b905290516001600160a01b038c16916340c10f1991604480830192600092919082900301818387803b15801561122657600080fd5b505af115801561123a573d6000803e3d6000fd5b50611253925050506001600160a01b038b16888b6139e2565b6001600160a01b0387166391695586878761126e8d8d613985565b88886040518663ffffffff1660e01b8152600401808660ff1681526020018560ff16815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b1580156112c857600080fd5b505af19250505080156112ed57506040513d60208110156112e857600080fd5b505160015b61139c576113108b6112ff8b8b613985565b6001600160a01b038d169190613b01565b816001600160a01b038c167f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c658c6113478d8d613985565b604080516001600160a01b03909316835260208301919091528181018d905260ff808c1660608401528a16608083015260a0820189905260c08201889052600060e083015251908190036101000190a36116a2565b6000886001600160a01b03166382b86600886040518263ffffffff1660e01b8152600401808260ff16815260200191505060206040518083038186803b1580156113e557600080fd5b505afa1580156113f9573d6000803e3d6000fd5b505050506040513d602081101561140f57600080fd5b505160cc549091506001600160a01b03808316911614801561143b575060cc546001600160a01b031615155b156115fe5760cc54604080517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526004810185905290516001600160a01b0390921691632e1a7d4d9160248082019260009290919082900301818387803b1580156114a657600080fd5b505af11580156114ba573d6000803e3d6000fd5b5050505060008d6001600160a01b03168360405180600001905060006040518083038185875af1925050503d8060008114611511576040519150601f19603f3d011682016040523d82523d6000602084013e611516565b606091505b505090508061156c576040805162461bcd60e51b815260206004820152601360248201527f4554485f5452414e534645525f4641494c454400000000000000000000000000604482015290519081900360640190fd5b848e6001600160a01b03167f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c658f868f8e8e8e8e600160405180896001600160a01b031681526020018881526020018781526020018660ff1681526020018560ff16815260200184815260200183815260200182151581526020019850505050505050505060405180910390a35061169f565b6116126001600160a01b0382168e84613b01565b838d6001600160a01b03167f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c658e858e8d8d8d8d600160405180896001600160a01b031681526020018881526020018781526020018660ff1681526020018560ff16815260200184815260200183815260200182151581526020019850505050505050505060405180910390a35b50505b6117c5565b604080517f40c10f19000000000000000000000000000000000000000000000000000000008152306004820152602481018b905290516001600160a01b038c16916340c10f1991604480830192600092919082900301818387803b15801561170e57600080fd5b505af1158015611722573d6000803e3d6000fd5b5050505061173d8b6112ff8a8c61398590919063ffffffff16565b816001600160a01b038c167f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c658c6117748d8d613985565b604080516001600160a01b03909316835260208301919091528181018d905260ff808c1660608401528a16608083015260a0820189905260c08201889052600060e083015251908190036101000190a35b50506001606555505050505050505050565b6002606554141561182f576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260655561183c612361565b1561188e576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6118b87fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd336127d6565b611909576040805162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f742061206e6f64652067726f7570000000000000604482015290519081900360640190fd5b81831161195d576040805162461bcd60e51b815260206004820152601f60248201527f416d6f756e74206d7573742062652067726561746572207468616e2066656500604482015290519081900360640190fd5b600081815260cd602052604090205460ff16156119c1576040805162461bcd60e51b815260206004820152601860248201527f4b6170706120697320616c72656164792070726573656e740000000000000000604482015290519081900360640190fd5b600081815260cd60209081526040808320805460ff191660011790556001600160a01b038716835260c99091529020546119fb908361392b565b6001600160a01b03808616600081815260c9602052604090209290925560cc5416148015611a33575060cc546001600160a01b031615155b15611ba55760cc546001600160a01b0316632e1a7d4d611a538585613985565b6040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015611a8957600080fd5b505af1158015611a9d573d6000803e3d6000fd5b506000925050506001600160a01b038616611ab88585613985565b604051600081818185875af1925050503d8060008114611af4576040519150601f19603f3d011682016040523d82523d6000602084013e611af9565b606091505b5050905080611b4f576040805162461bcd60e51b815260206004820152601360248201527f4554485f5452414e534645525f4641494c454400000000000000000000000000604482015290519081900360640190fd5b604080516001600160a01b03878116825260208201879052818301869052915184928916917f8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26919081900360600190a350611c14565b604080516001600160a01b03868116825260208201869052818301859052915183928816917f8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26919081900360600190a3611c1485611c038585613985565b6001600160a01b0387169190613b01565b50506001606555505050565b60026065541415611c78576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555611c85612361565b15611cd7576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b611d017fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd336127d6565b611d52576040805162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f742061206e6f64652067726f7570000000000000604482015290519081900360640190fd5b818311611da6576040805162461bcd60e51b815260206004820152601f60248201527f416d6f756e74206d7573742062652067726561746572207468616e2066656500604482015290519081900360640190fd5b600081815260cd602052604090205460ff1615611e0a576040805162461bcd60e51b815260206004820152601860248201527f4b6170706120697320616c72656164792070726573656e740000000000000000604482015290519081900360640190fd5b600081815260cd60209081526040808320805460ff191660011790556001600160a01b038716835260c9909152902054611e44908361392b565b6001600160a01b03808616600090815260c96020526040902091909155819086167fbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd3886611e918787613985565b604080516001600160a01b0390931683526020830191909152818101879052519081900360600190a3604080517f40c10f190000000000000000000000000000000000000000000000000000000081523060048201526024810185905290516001600160a01b038616916340c10f1991604480830192600092919082900301818387803b158015611f2157600080fd5b505af1158015611f35573d6000803e3d6000fd5b50505050611f5085611c03848661398590919063ffffffff16565b60cb5415801590611f62575060cb5447115b15611c145760cb546040516001600160a01b0387169190600081818185875af1925050503d8060008114611fb2576040519150601f19603f3d011682016040523d82523d6000602084013e611fb7565b606091505b50505050506001606555505050565b60009081526033602052604090206002015490565b600082815260336020526040902060020154611ffe90611ff9613b81565b6127d6565b6120395760405162461bcd60e51b815260040180806020018281038252602f8152602001806144d2602f913960400191505060405180910390fd5b6120438282613b85565b5050565b600090815260cd602052604090205460ff1690565b604080516001600160a01b038b81168252602082018b90528183018a905260ff80891660608401528716608083015260a0820186905260c08201859052600160e0830152915183928d16917f4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c6591908190036101000190a350505050505050505050565b6120e7613b81565b6001600160a01b0316816001600160a01b0316146121365760405162461bcd60e51b815260040180806020018281038252602f8152602001806145af602f913960400191505060405180910390fd5b6120438282613bee565b60026065541415612198576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026065556121a5612361565b156121f7576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518781526001600160a01b03878116602083015281830187905260ff861660608301526080820185905260a082018490529151918916917f9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c9181900360c00190a2604080517f79cc67900000000000000000000000000000000000000000000000000000000081523360048201526024810186905290516001600160a01b038716916379cc679091604480830192600092919082900301818387803b1580156122c257600080fd5b505af11580156117c5573d6000803e3d6000fd5b6123007f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb1336127d6565b612351576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b612359613c57565b565b60ca5481565b60975460ff1690565b600054610100900460ff16806123835750612383613d00565b80612391575060005460ff16155b6123cc5760405162461bcd60e51b815260040180806020018281038252602e815260200180614557602e913960400191505060405180910390fd5b600054610100900460ff16158015612414576000805460ff197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b4360ca55612423600033612039565b61242b613d11565b801561245a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50565b600260655414156124b5576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026065556124c2612361565b15612514576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518881526001600160a01b03888116602083015281830188905260ff80881660608401528616608083015260a0820185905260c082018490529151918a16917f91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc3194259181900360e00190a2604080517f79cc67900000000000000000000000000000000000000000000000000000000081523360048201526024810187905290516001600160a01b038816916379cc679091604480830192600092919082900301818387803b1580156125e757600080fd5b505af11580156125fb573d6000803e3d6000fd5b5050600160655550505050505050505050565b6126387f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb1336127d6565b612689576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b612359613dcb565b60008281526033602052604081206126a99083613e5b565b90505b92915050565b6002606554141561270a576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555612717612361565b15612769576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518481526001600160a01b0384811660208301528183018490529151918616917fda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b9181900360600190a26127cb6001600160a01b038316333084613e67565b505060016065555050565b60008281526033602052604081206126a99083613eef565b604080518481526001600160a01b0384166020820152808201839052905185917f8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499919081900360600190a250505050565b60026065541415612897576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026065556128a4612361565b156128f6576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518481526001600160a01b0384166020820152808201839052905185917f8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499919081900360600190a2604080517f79cc67900000000000000000000000000000000000000000000000000000000081523360048201526024810183905290516001600160a01b038416916379cc679091604480830192600092919082900301818387803b1580156129a857600080fd5b505af11580156129bc573d6000803e3d6000fd5b50506001606555505050505050565b600081565b60026065541415612a28576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555612a35612361565b15612a87576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518881526001600160a01b03888116602083015281830188905260ff80881660608401528616608083015260a0820185905260c082018490529151918a16917f79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f9181900360e00190a26129bc6001600160a01b038716333088613e67565b612b136000336127d6565b612b64576040805162461bcd60e51b815260206004820152600960248201527f4e6f742061646d696e0000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60cc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600681565b604080518481526001600160a01b0384811660208301528183018490529151918616917fda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b9181900360600190a250505050565b612c207f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb1336127d6565b612c71576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b60cb55565b604080518781526001600160a01b03878116602083015281830187905260ff861660608301526080820185905260a082018490529151918916917f9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c9181900360c00190a250505050505050565b6001600160a01b0316600090815260c9602052604090205490565b60008181526033602052604081206126ac90613f04565b600082815260336020526040902060020154612d3390611ff9613b81565b6121365760405162461bcd60e51b81526004018080602001828103825260308152602001806145276030913960400191505060405180910390fd5b60026065541415612dc6576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002606555612dd3612361565b15612e25576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b612e4f7fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd336127d6565b612ea0576040805162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f742061206e6f64652067726f7570000000000000604482015290519081900360640190fd5b858711612ef4576040805162461bcd60e51b815260206004820152601f60248201527f416d6f756e74206d7573742062652067726561746572207468616e2066656500604482015290519081900360640190fd5b600081815260cd602052604090205460ff1615612f58576040805162461bcd60e51b815260206004820152601860248201527f4b6170706120697320616c72656164792070726573656e740000000000000000604482015290519081900360640190fd5b600081815260cd60209081526040808320805460ff191660011790556001600160a01b038b16835260c9909152902054612f92908761392b565b6001600160a01b03808a16600090815260c96020526040812092909255861663342a87a1612fc08a8a613985565b876040518363ffffffff1660e01b8152600401808381526020018260ff1681526020019250505060206040518083038186803b158015612fff57600080fd5b505afa158015613013573d6000803e3d6000fd5b505050506040513d602081101561302957600080fd5b5051905083811061328957613053866130428a8a613985565b6001600160a01b038c1691906139e2565b6001600160a01b038616633e3a156061306c8a8a613985565b8787876040518563ffffffff1660e01b8152600401808581526020018460ff168152602001838152602001828152602001945050505050602060405180830381600087803b1580156130bd57600080fd5b505af19250505080156130e257506040513d60208110156130dd57600080fd5b505160015b613188576131058a6130f48a8a613985565b6001600160a01b038c169190613b01565b816001600160a01b038b167fc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d8b61313c8c8c613985565b604080516001600160a01b03909316835260208301919091528181018c905260ff8a1660608301526080820189905260a08201889052600060c0830152519081900360e00190a3613284565b6000876001600160a01b03166382b86600886040518263ffffffff1660e01b8152600401808260ff16815260200191505060206040518083038186803b1580156131d157600080fd5b505afa1580156131e5573d6000803e3d6000fd5b505050506040513d60208110156131fb57600080fd5b505190506132136001600160a01b0382168d84613b01565b604080516001600160a01b038d81168252602082018590528183018c905260ff8a1660608301526080820189905260a08201889052600160c0830152915186928f16917fc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d919081900360e00190a350505b613316565b6132978a6130f48a8a613985565b816001600160a01b038b167fc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d8b6132ce8c8c613985565b604080516001600160a01b03909316835260208301919091528181018c905260ff8a1660608301526080820189905260a08201889052600060c0830152519081900360e00190a35b505060016065555050505050505050565b604080518881526001600160a01b03888116602083015281830188905260ff80881660608401528616608083015260a0820185905260c082018490529151918a16917f91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc3194259181900360e00190a25050505050505050565b60cb5481565b604080518481526001600160a01b0384811660208301528183018490529151918616917fdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c9181900360600190a250505050565b604080516001600160a01b038a81168252602082018a905281830189905260ff871660608301526080820186905260a08201859052600160c0830152915183928c16917fc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d919081900360e00190a3505050505050505050565b6134997f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb1336127d6565b6134ea576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b60005b8181101561353757600160cd600085858581811061350757fe5b60209081029290920135835250810191909152604001600020805460ff19169115159190911790556001016134ed565b505050565b613544612361565b15613596576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6135c07f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb1336127d6565b613611576040805162461bcd60e51b815260206004820152600e60248201527f4e6f7420676f7665726e616e6365000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03811661366c576040805162461bcd60e51b815260206004820152601060248201527f4164647265737320697320307830303000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038216600090815260c9602052604090205415612043576001600160a01b038216600081815260c960205260409020546136af91908390613b01565b506001600160a01b0316600090815260c96020526040812055565b604080516001600160a01b03868116825260208201869052818301859052915183928816917fbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd38919081900360600190a35050505050565b7f71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb181565b7fb5c00e6706c3d213edd70ff33717fac657eacc5fe161f07180cf1fcab13cc4cd81565b600260655414156137c1576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026065556137ce612361565b15613820576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b604080518481526001600160a01b0384811660208301528183018490529151918616917fdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c9181900360600190a2604080517f79cc67900000000000000000000000000000000000000000000000000000000081523360048201526024810183905290516001600160a01b038416916379cc679091604480830192600092919082900301818387803b1580156129a857600080fd5b604080516001600160a01b03868116825260208201869052818301859052915183928816917f8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26919081900360600190a35050505050565b6000828201838110156126a9576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6000828211156139dc576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000613a7882856001600160a01b031663dd62ed3e30876040518363ffffffff1660e01b815260040180836001600160a01b03168152602001826001600160a01b031681526020019250505060206040518083038186803b158015613a4657600080fd5b505afa158015613a5a573d6000803e3d6000fd5b505050506040513d6020811015613a7057600080fd5b50519061392b565b604080516001600160a01b038616602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052909150613afb908590613f0f565b50505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052613537908490613f0f565b3390565b6000828152603360205260409020613b9d9082613fc0565b1561204357613baa613b81565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000828152603360205260409020613c069082613fd5565b1561204357613c13613b81565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b613c5f612361565b613cb0576040805162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015290519081900360640190fd5b6097805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613ce3613b81565b604080516001600160a01b039092168252519081900360200190a1565b6000613d0b30613fea565b15905090565b600054610100900460ff1680613d2a5750613d2a613d00565b80613d38575060005460ff16155b613d735760405162461bcd60e51b815260040180806020018281038252602e815260200180614557602e913960400191505060405180910390fd5b600054610100900460ff16158015613dbb576000805460ff197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b613dc3613ff0565b61242b613ff0565b613dd3612361565b15613e25576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6097805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613ce3613b81565b60006126a983836140ca565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052613afb908590613f0f565b60006126a9836001600160a01b03841661412e565b60006126ac82614146565b6060613f64826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661414a9092919063ffffffff16565b80519091501561353757808060200190516020811015613f8357600080fd5b50516135375760405162461bcd60e51b815260040180806020018281038252602a815260200180614585602a913960400191505060405180910390fd5b60006126a9836001600160a01b038416614163565b60006126a9836001600160a01b0384166141ad565b3b151590565b600054610100900460ff16806140095750614009613d00565b80614017575060005460ff16155b6140525760405162461bcd60e51b815260040180806020018281038252602e815260200180614557602e913960400191505060405180910390fd5b600054610100900460ff1615801561242b576000805460ff197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90911661010017166001179055801561245a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905550565b8154600090821061410c5760405162461bcd60e51b81526004018080602001828103825260228152602001806144b06022913960400191505060405180910390fd5b82600001828154811061411b57fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b60606141598484600085614291565b90505b9392505050565b600061416f838361412e565b6141a5575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556126ac565b5060006126ac565b600081815260018301602052604081205480156142875783547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80830191908101906000908790839081106141fe57fe5b906000526020600020015490508087600001848154811061421b57fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061424b57fe5b600190038181906000526020600020016000905590558660010160008781526020019081526020016000206000905560019450505050506126ac565b60009150506126ac565b6060824710156142d25760405162461bcd60e51b81526004018080602001828103825260268152602001806145016026913960400191505060405180910390fd5b6142db85613fea565b61432c576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b6020831061438957805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161434c565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146143eb576040519150601f19603f3d011682016040523d82523d6000602084013e6143f0565b606091505b509150915061440082828661440b565b979650505050505050565b6060831561441a57508161415c565b82511561442a5782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561447457818101518382015260200161445c565b50505050905090810190601f1680156144a15780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e6473416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e74416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b65496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a26469706673582212207ba281b3ad1f6d47ab31b714a2055ee4fab62ea8e4be64a30a74d7e6c6f6ac2e64736f6c634300060c0033",
}

// TestSynapseBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use TestSynapseBridgeMetaData.ABI instead.
var TestSynapseBridgeABI = TestSynapseBridgeMetaData.ABI

// Deprecated: Use TestSynapseBridgeMetaData.Sigs instead.
// TestSynapseBridgeFuncSigs maps the 4-byte function signature to its string representation.
var TestSynapseBridgeFuncSigs = TestSynapseBridgeMetaData.Sigs

// TestSynapseBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestSynapseBridgeMetaData.Bin instead.
var TestSynapseBridgeBin = TestSynapseBridgeMetaData.Bin

// DeployTestSynapseBridge deploys a new Ethereum contract, binding an instance of TestSynapseBridge to it.
func DeployTestSynapseBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestSynapseBridge, error) {
	parsed, err := TestSynapseBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestSynapseBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestSynapseBridge{TestSynapseBridgeCaller: TestSynapseBridgeCaller{contract: contract}, TestSynapseBridgeTransactor: TestSynapseBridgeTransactor{contract: contract}, TestSynapseBridgeFilterer: TestSynapseBridgeFilterer{contract: contract}}, nil
}

// TestSynapseBridge is an auto generated Go binding around an Ethereum contract.
type TestSynapseBridge struct {
	TestSynapseBridgeCaller     // Read-only binding to the contract
	TestSynapseBridgeTransactor // Write-only binding to the contract
	TestSynapseBridgeFilterer   // Log filterer for contract events
}

// TestSynapseBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestSynapseBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSynapseBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestSynapseBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSynapseBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestSynapseBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSynapseBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSynapseBridgeSession struct {
	Contract     *TestSynapseBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestSynapseBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestSynapseBridgeCallerSession struct {
	Contract *TestSynapseBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TestSynapseBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestSynapseBridgeTransactorSession struct {
	Contract     *TestSynapseBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TestSynapseBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestSynapseBridgeRaw struct {
	Contract *TestSynapseBridge // Generic contract binding to access the raw methods on
}

// TestSynapseBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestSynapseBridgeCallerRaw struct {
	Contract *TestSynapseBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// TestSynapseBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestSynapseBridgeTransactorRaw struct {
	Contract *TestSynapseBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestSynapseBridge creates a new instance of TestSynapseBridge, bound to a specific deployed contract.
func NewTestSynapseBridge(address common.Address, backend bind.ContractBackend) (*TestSynapseBridge, error) {
	contract, err := bindTestSynapseBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridge{TestSynapseBridgeCaller: TestSynapseBridgeCaller{contract: contract}, TestSynapseBridgeTransactor: TestSynapseBridgeTransactor{contract: contract}, TestSynapseBridgeFilterer: TestSynapseBridgeFilterer{contract: contract}}, nil
}

// NewTestSynapseBridgeCaller creates a new read-only instance of TestSynapseBridge, bound to a specific deployed contract.
func NewTestSynapseBridgeCaller(address common.Address, caller bind.ContractCaller) (*TestSynapseBridgeCaller, error) {
	contract, err := bindTestSynapseBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeCaller{contract: contract}, nil
}

// NewTestSynapseBridgeTransactor creates a new write-only instance of TestSynapseBridge, bound to a specific deployed contract.
func NewTestSynapseBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*TestSynapseBridgeTransactor, error) {
	contract, err := bindTestSynapseBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTransactor{contract: contract}, nil
}

// NewTestSynapseBridgeFilterer creates a new log filterer instance of TestSynapseBridge, bound to a specific deployed contract.
func NewTestSynapseBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*TestSynapseBridgeFilterer, error) {
	contract, err := bindTestSynapseBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeFilterer{contract: contract}, nil
}

// bindTestSynapseBridge binds a generic wrapper to an already deployed contract.
func bindTestSynapseBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestSynapseBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSynapseBridge *TestSynapseBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSynapseBridge.Contract.TestSynapseBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSynapseBridge *TestSynapseBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestSynapseBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSynapseBridge *TestSynapseBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestSynapseBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSynapseBridge *TestSynapseBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSynapseBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSynapseBridge *TestSynapseBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSynapseBridge *TestSynapseBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TestSynapseBridge.Contract.DEFAULTADMINROLE(&_TestSynapseBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TestSynapseBridge.Contract.DEFAULTADMINROLE(&_TestSynapseBridge.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeSession) GOVERNANCEROLE() ([32]byte, error) {
	return _TestSynapseBridge.Contract.GOVERNANCEROLE(&_TestSynapseBridge.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _TestSynapseBridge.Contract.GOVERNANCEROLE(&_TestSynapseBridge.CallOpts)
}

// NODEGROUPROLE is a free data retrieval call binding the contract method 0xf3befd01.
//
// Solidity: function NODEGROUP_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeCaller) NODEGROUPROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "NODEGROUP_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NODEGROUPROLE is a free data retrieval call binding the contract method 0xf3befd01.
//
// Solidity: function NODEGROUP_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeSession) NODEGROUPROLE() ([32]byte, error) {
	return _TestSynapseBridge.Contract.NODEGROUPROLE(&_TestSynapseBridge.CallOpts)
}

// NODEGROUPROLE is a free data retrieval call binding the contract method 0xf3befd01.
//
// Solidity: function NODEGROUP_ROLE() view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) NODEGROUPROLE() ([32]byte, error) {
	return _TestSynapseBridge.Contract.NODEGROUPROLE(&_TestSynapseBridge.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_TestSynapseBridge *TestSynapseBridgeCaller) WETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "WETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_TestSynapseBridge *TestSynapseBridgeSession) WETHADDRESS() (common.Address, error) {
	return _TestSynapseBridge.Contract.WETHADDRESS(&_TestSynapseBridge.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) WETHADDRESS() (common.Address, error) {
	return _TestSynapseBridge.Contract.WETHADDRESS(&_TestSynapseBridge.CallOpts)
}

// BridgeVersion is a free data retrieval call binding the contract method 0xac865626.
//
// Solidity: function bridgeVersion() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCaller) BridgeVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "bridgeVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeVersion is a free data retrieval call binding the contract method 0xac865626.
//
// Solidity: function bridgeVersion() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeSession) BridgeVersion() (*big.Int, error) {
	return _TestSynapseBridge.Contract.BridgeVersion(&_TestSynapseBridge.CallOpts)
}

// BridgeVersion is a free data retrieval call binding the contract method 0xac865626.
//
// Solidity: function bridgeVersion() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) BridgeVersion() (*big.Int, error) {
	return _TestSynapseBridge.Contract.BridgeVersion(&_TestSynapseBridge.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCaller) ChainGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "chainGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeSession) ChainGasAmount() (*big.Int, error) {
	return _TestSynapseBridge.Contract.ChainGasAmount(&_TestSynapseBridge.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) ChainGasAmount() (*big.Int, error) {
	return _TestSynapseBridge.Contract.ChainGasAmount(&_TestSynapseBridge.CallOpts)
}

// GetFeeBalance is a free data retrieval call binding the contract method 0xc78f6803.
//
// Solidity: function getFeeBalance(address tokenAddress) view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCaller) GetFeeBalance(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "getFeeBalance", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBalance is a free data retrieval call binding the contract method 0xc78f6803.
//
// Solidity: function getFeeBalance(address tokenAddress) view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeSession) GetFeeBalance(tokenAddress common.Address) (*big.Int, error) {
	return _TestSynapseBridge.Contract.GetFeeBalance(&_TestSynapseBridge.CallOpts, tokenAddress)
}

// GetFeeBalance is a free data retrieval call binding the contract method 0xc78f6803.
//
// Solidity: function getFeeBalance(address tokenAddress) view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) GetFeeBalance(tokenAddress common.Address) (*big.Int, error) {
	return _TestSynapseBridge.Contract.GetFeeBalance(&_TestSynapseBridge.CallOpts, tokenAddress)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TestSynapseBridge.Contract.GetRoleAdmin(&_TestSynapseBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TestSynapseBridge.Contract.GetRoleAdmin(&_TestSynapseBridge.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TestSynapseBridge *TestSynapseBridgeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TestSynapseBridge *TestSynapseBridgeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TestSynapseBridge.Contract.GetRoleMember(&_TestSynapseBridge.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TestSynapseBridge.Contract.GetRoleMember(&_TestSynapseBridge.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TestSynapseBridge.Contract.GetRoleMemberCount(&_TestSynapseBridge.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TestSynapseBridge.Contract.GetRoleMemberCount(&_TestSynapseBridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TestSynapseBridge.Contract.HasRole(&_TestSynapseBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TestSynapseBridge.Contract.HasRole(&_TestSynapseBridge.CallOpts, role, account)
}

// KappaExists is a free data retrieval call binding the contract method 0x2fe87b95.
//
// Solidity: function kappaExists(bytes32 kappa) view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeCaller) KappaExists(opts *bind.CallOpts, kappa [32]byte) (bool, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "kappaExists", kappa)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KappaExists is a free data retrieval call binding the contract method 0x2fe87b95.
//
// Solidity: function kappaExists(bytes32 kappa) view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeSession) KappaExists(kappa [32]byte) (bool, error) {
	return _TestSynapseBridge.Contract.KappaExists(&_TestSynapseBridge.CallOpts, kappa)
}

// KappaExists is a free data retrieval call binding the contract method 0x2fe87b95.
//
// Solidity: function kappaExists(bytes32 kappa) view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) KappaExists(kappa [32]byte) (bool, error) {
	return _TestSynapseBridge.Contract.KappaExists(&_TestSynapseBridge.CallOpts, kappa)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeSession) Paused() (bool, error) {
	return _TestSynapseBridge.Contract.Paused(&_TestSynapseBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) Paused() (bool, error) {
	return _TestSynapseBridge.Contract.Paused(&_TestSynapseBridge.CallOpts)
}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCaller) StartBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestSynapseBridge.contract.Call(opts, &out, "startBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeSession) StartBlockNumber() (*big.Int, error) {
	return _TestSynapseBridge.Contract.StartBlockNumber(&_TestSynapseBridge.CallOpts)
}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_TestSynapseBridge *TestSynapseBridgeCallerSession) StartBlockNumber() (*big.Int, error) {
	return _TestSynapseBridge.Contract.StartBlockNumber(&_TestSynapseBridge.CallOpts)
}

// AddKappas is a paid mutator transaction binding the contract method 0xe7a59998.
//
// Solidity: function addKappas(bytes32[] kappas) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) AddKappas(opts *bind.TransactOpts, kappas [][32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "addKappas", kappas)
}

// AddKappas is a paid mutator transaction binding the contract method 0xe7a59998.
//
// Solidity: function addKappas(bytes32[] kappas) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) AddKappas(kappas [][32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.AddKappas(&_TestSynapseBridge.TransactOpts, kappas)
}

// AddKappas is a paid mutator transaction binding the contract method 0xe7a59998.
//
// Solidity: function addKappas(bytes32[] kappas) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) AddKappas(kappas [][32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.AddKappas(&_TestSynapseBridge.TransactOpts, kappas)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) Deposit(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "deposit", to, chainId, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) Deposit(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Deposit(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) Deposit(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Deposit(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// DepositAndSwap is a paid mutator transaction binding the contract method 0xa2a2af0b.
//
// Solidity: function depositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) DepositAndSwap(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "depositAndSwap", to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// DepositAndSwap is a paid mutator transaction binding the contract method 0xa2a2af0b.
//
// Solidity: function depositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) DepositAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.DepositAndSwap(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// DepositAndSwap is a paid mutator transaction binding the contract method 0xa2a2af0b.
//
// Solidity: function depositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) DepositAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.DepositAndSwap(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.GrantRole(&_TestSynapseBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.GrantRole(&_TestSynapseBridge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) Initialize() (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Initialize(&_TestSynapseBridge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) Initialize() (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Initialize(&_TestSynapseBridge.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x20d7b327.
//
// Solidity: function mint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) Mint(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "mint", to, token, amount, fee, kappa)
}

// Mint is a paid mutator transaction binding the contract method 0x20d7b327.
//
// Solidity: function mint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) Mint(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Mint(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// Mint is a paid mutator transaction binding the contract method 0x20d7b327.
//
// Solidity: function mint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) Mint(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Mint(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// MintAndSwap is a paid mutator transaction binding the contract method 0x17357892.
//
// Solidity: function mintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) MintAndSwap(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "mintAndSwap", to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// MintAndSwap is a paid mutator transaction binding the contract method 0x17357892.
//
// Solidity: function mintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) MintAndSwap(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.MintAndSwap(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// MintAndSwap is a paid mutator transaction binding the contract method 0x17357892.
//
// Solidity: function mintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) MintAndSwap(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.MintAndSwap(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) Pause() (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Pause(&_TestSynapseBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Pause(&_TestSynapseBridge.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) Redeem(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "redeem", to, chainId, token, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) Redeem(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Redeem(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) Redeem(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Redeem(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// RedeemAndRemove is a paid mutator transaction binding the contract method 0x36e712ed.
//
// Solidity: function redeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) RedeemAndRemove(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "redeemAndRemove", to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// RedeemAndRemove is a paid mutator transaction binding the contract method 0x36e712ed.
//
// Solidity: function redeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) RedeemAndRemove(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RedeemAndRemove(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// RedeemAndRemove is a paid mutator transaction binding the contract method 0x36e712ed.
//
// Solidity: function redeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) RedeemAndRemove(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RedeemAndRemove(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// RedeemAndSwap is a paid mutator transaction binding the contract method 0x839ed90a.
//
// Solidity: function redeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) RedeemAndSwap(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "redeemAndSwap", to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// RedeemAndSwap is a paid mutator transaction binding the contract method 0x839ed90a.
//
// Solidity: function redeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) RedeemAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RedeemAndSwap(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// RedeemAndSwap is a paid mutator transaction binding the contract method 0x839ed90a.
//
// Solidity: function redeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) RedeemAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RedeemAndSwap(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// RedeemV2 is a paid mutator transaction binding the contract method 0xa07ed975.
//
// Solidity: function redeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) RedeemV2(opts *bind.TransactOpts, to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "redeemV2", to, chainId, token, amount)
}

// RedeemV2 is a paid mutator transaction binding the contract method 0xa07ed975.
//
// Solidity: function redeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) RedeemV2(to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RedeemV2(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// RedeemV2 is a paid mutator transaction binding the contract method 0xa07ed975.
//
// Solidity: function redeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) RedeemV2(to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RedeemV2(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RenounceRole(&_TestSynapseBridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RenounceRole(&_TestSynapseBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RevokeRole(&_TestSynapseBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.RevokeRole(&_TestSynapseBridge.TransactOpts, role, account)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) SetChainGasAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "setChainGasAmount", amount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) SetChainGasAmount(amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.SetChainGasAmount(&_TestSynapseBridge.TransactOpts, amount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) SetChainGasAmount(amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.SetChainGasAmount(&_TestSynapseBridge.TransactOpts, amount)
}

// SetWethAddress is a paid mutator transaction binding the contract method 0xa96e2423.
//
// Solidity: function setWethAddress(address _wethAddress) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) SetWethAddress(opts *bind.TransactOpts, _wethAddress common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "setWethAddress", _wethAddress)
}

// SetWethAddress is a paid mutator transaction binding the contract method 0xa96e2423.
//
// Solidity: function setWethAddress(address _wethAddress) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) SetWethAddress(_wethAddress common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.SetWethAddress(&_TestSynapseBridge.TransactOpts, _wethAddress)
}

// SetWethAddress is a paid mutator transaction binding the contract method 0xa96e2423.
//
// Solidity: function setWethAddress(address _wethAddress) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) SetWethAddress(_wethAddress common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.SetWethAddress(&_TestSynapseBridge.TransactOpts, _wethAddress)
}

// TestDeposit is a paid mutator transaction binding the contract method 0xad863232.
//
// Solidity: function testDeposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestDeposit(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testDeposit", to, chainId, token, amount)
}

// TestDeposit is a paid mutator transaction binding the contract method 0xad863232.
//
// Solidity: function testDeposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestDeposit(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestDeposit(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// TestDeposit is a paid mutator transaction binding the contract method 0xad863232.
//
// Solidity: function testDeposit(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestDeposit(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestDeposit(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// TestDepositAndSwap is a paid mutator transaction binding the contract method 0x09faf02b.
//
// Solidity: function testDepositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestDepositAndSwap(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testDepositAndSwap", to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// TestDepositAndSwap is a paid mutator transaction binding the contract method 0x09faf02b.
//
// Solidity: function testDepositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestDepositAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestDepositAndSwap(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// TestDepositAndSwap is a paid mutator transaction binding the contract method 0x09faf02b.
//
// Solidity: function testDepositAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestDepositAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestDepositAndSwap(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// TestMint is a paid mutator transaction binding the contract method 0xf3320221.
//
// Solidity: function testMint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestMint(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testMint", to, token, amount, fee, kappa)
}

// TestMint is a paid mutator transaction binding the contract method 0xf3320221.
//
// Solidity: function testMint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestMint(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestMint(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// TestMint is a paid mutator transaction binding the contract method 0xf3320221.
//
// Solidity: function testMint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestMint(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestMint(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// TestMintAndSwap is a paid mutator transaction binding the contract method 0x328a9dc3.
//
// Solidity: function testMintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestMintAndSwap(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testMintAndSwap", to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// TestMintAndSwap is a paid mutator transaction binding the contract method 0x328a9dc3.
//
// Solidity: function testMintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestMintAndSwap(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestMintAndSwap(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// TestMintAndSwap is a paid mutator transaction binding the contract method 0x328a9dc3.
//
// Solidity: function testMintAndSwap(address to, address token, uint256 amount, uint256 fee, address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestMintAndSwap(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestMintAndSwap(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, pool, tokenIndexFrom, tokenIndexTo, minDy, deadline, kappa)
}

// TestRedeem is a paid mutator transaction binding the contract method 0xe072f5cd.
//
// Solidity: function testRedeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestRedeem(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testRedeem", to, chainId, token, amount)
}

// TestRedeem is a paid mutator transaction binding the contract method 0xe072f5cd.
//
// Solidity: function testRedeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestRedeem(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestRedeem(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// TestRedeem is a paid mutator transaction binding the contract method 0xe072f5cd.
//
// Solidity: function testRedeem(address to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestRedeem(to common.Address, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestRedeem(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// TestRedeemAndRemove is a paid mutator transaction binding the contract method 0xb814ff5b.
//
// Solidity: function testRedeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestRedeemAndRemove(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testRedeemAndRemove", to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// TestRedeemAndRemove is a paid mutator transaction binding the contract method 0xb814ff5b.
//
// Solidity: function testRedeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestRedeemAndRemove(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestRedeemAndRemove(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// TestRedeemAndRemove is a paid mutator transaction binding the contract method 0xb814ff5b.
//
// Solidity: function testRedeemAndRemove(address to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestRedeemAndRemove(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestRedeemAndRemove(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, swapTokenIndex, swapMinAmount, swapDeadline)
}

// TestRedeemAndSwap is a paid mutator transaction binding the contract method 0xdbb176f4.
//
// Solidity: function testRedeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestRedeemAndSwap(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testRedeemAndSwap", to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// TestRedeemAndSwap is a paid mutator transaction binding the contract method 0xdbb176f4.
//
// Solidity: function testRedeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestRedeemAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestRedeemAndSwap(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// TestRedeemAndSwap is a paid mutator transaction binding the contract method 0xdbb176f4.
//
// Solidity: function testRedeemAndSwap(address to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestRedeemAndSwap(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, tokenIndexFrom uint8, tokenIndexTo uint8, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestRedeemAndSwap(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount, tokenIndexFrom, tokenIndexTo, minDy, deadline)
}

// TestRedeemV2 is a paid mutator transaction binding the contract method 0x98c3e142.
//
// Solidity: function testRedeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestRedeemV2(opts *bind.TransactOpts, to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testRedeemV2", to, chainId, token, amount)
}

// TestRedeemV2 is a paid mutator transaction binding the contract method 0x98c3e142.
//
// Solidity: function testRedeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestRedeemV2(to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestRedeemV2(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// TestRedeemV2 is a paid mutator transaction binding the contract method 0x98c3e142.
//
// Solidity: function testRedeemV2(bytes32 to, uint256 chainId, address token, uint256 amount) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestRedeemV2(to [32]byte, chainId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestRedeemV2(&_TestSynapseBridge.TransactOpts, to, chainId, token, amount)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xffafc49b.
//
// Solidity: function testWithdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestWithdraw(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testWithdraw", to, token, amount, fee, kappa)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xffafc49b.
//
// Solidity: function testWithdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestWithdraw(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestWithdraw(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xffafc49b.
//
// Solidity: function testWithdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestWithdraw(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestWithdraw(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// TestWithdrawAndRemove is a paid mutator transaction binding the contract method 0xe2176fbe.
//
// Solidity: function testWithdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) TestWithdrawAndRemove(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "testWithdrawAndRemove", to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// TestWithdrawAndRemove is a paid mutator transaction binding the contract method 0xe2176fbe.
//
// Solidity: function testWithdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) TestWithdrawAndRemove(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestWithdrawAndRemove(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// TestWithdrawAndRemove is a paid mutator transaction binding the contract method 0xe2176fbe.
//
// Solidity: function testWithdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) TestWithdrawAndRemove(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.TestWithdrawAndRemove(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) Unpause() (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Unpause(&_TestSynapseBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Unpause(&_TestSynapseBridge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1cf5f07f.
//
// Solidity: function withdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "withdraw", to, token, amount, fee, kappa)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1cf5f07f.
//
// Solidity: function withdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) Withdraw(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Withdraw(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1cf5f07f.
//
// Solidity: function withdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) Withdraw(to common.Address, token common.Address, amount *big.Int, fee *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Withdraw(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, kappa)
}

// WithdrawAndRemove is a paid mutator transaction binding the contract method 0xd57eafac.
//
// Solidity: function withdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) WithdrawAndRemove(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "withdrawAndRemove", to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// WithdrawAndRemove is a paid mutator transaction binding the contract method 0xd57eafac.
//
// Solidity: function withdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) WithdrawAndRemove(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.WithdrawAndRemove(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// WithdrawAndRemove is a paid mutator transaction binding the contract method 0xd57eafac.
//
// Solidity: function withdrawAndRemove(address to, address token, uint256 amount, uint256 fee, address pool, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bytes32 kappa) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) WithdrawAndRemove(to common.Address, token common.Address, amount *big.Int, fee *big.Int, pool common.Address, swapTokenIndex uint8, swapMinAmount *big.Int, swapDeadline *big.Int, kappa [32]byte) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.WithdrawAndRemove(&_TestSynapseBridge.TransactOpts, to, token, amount, fee, pool, swapTokenIndex, swapMinAmount, swapDeadline, kappa)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address token, address to) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) WithdrawFees(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.Transact(opts, "withdrawFees", token, to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address token, address to) returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) WithdrawFees(token common.Address, to common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.WithdrawFees(&_TestSynapseBridge.TransactOpts, token, to)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address token, address to) returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) WithdrawFees(token common.Address, to common.Address) (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.WithdrawFees(&_TestSynapseBridge.TransactOpts, token, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSynapseBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestSynapseBridge *TestSynapseBridgeSession) Receive() (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Receive(&_TestSynapseBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestSynapseBridge *TestSynapseBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _TestSynapseBridge.Contract.Receive(&_TestSynapseBridge.TransactOpts)
}

// TestSynapseBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TestSynapseBridge contract.
type TestSynapseBridgePausedIterator struct {
	Event *TestSynapseBridgePaused // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgePaused)
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
		it.Event = new(TestSynapseBridgePaused)
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
func (it *TestSynapseBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgePaused represents a Paused event raised by the TestSynapseBridge contract.
type TestSynapseBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*TestSynapseBridgePausedIterator, error) {

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgePausedIterator{contract: _TestSynapseBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgePaused) (event.Subscription, error) {

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgePaused)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParsePaused(log types.Log) (*TestSynapseBridgePaused, error) {
	event := new(TestSynapseBridgePaused)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TestSynapseBridge contract.
type TestSynapseBridgeRoleAdminChangedIterator struct {
	Event *TestSynapseBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeRoleAdminChanged)
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
		it.Event = new(TestSynapseBridgeRoleAdminChanged)
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
func (it *TestSynapseBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the TestSynapseBridge contract.
type TestSynapseBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TestSynapseBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeRoleAdminChangedIterator{contract: _TestSynapseBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeRoleAdminChanged)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*TestSynapseBridgeRoleAdminChanged, error) {
	event := new(TestSynapseBridgeRoleAdminChanged)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TestSynapseBridge contract.
type TestSynapseBridgeRoleGrantedIterator struct {
	Event *TestSynapseBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeRoleGranted)
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
		it.Event = new(TestSynapseBridgeRoleGranted)
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
func (it *TestSynapseBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeRoleGranted represents a RoleGranted event raised by the TestSynapseBridge contract.
type TestSynapseBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TestSynapseBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeRoleGrantedIterator{contract: _TestSynapseBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeRoleGranted)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseRoleGranted(log types.Log) (*TestSynapseBridgeRoleGranted, error) {
	event := new(TestSynapseBridgeRoleGranted)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TestSynapseBridge contract.
type TestSynapseBridgeRoleRevokedIterator struct {
	Event *TestSynapseBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeRoleRevoked)
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
		it.Event = new(TestSynapseBridgeRoleRevoked)
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
func (it *TestSynapseBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeRoleRevoked represents a RoleRevoked event raised by the TestSynapseBridge contract.
type TestSynapseBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TestSynapseBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeRoleRevokedIterator{contract: _TestSynapseBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeRoleRevoked)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseRoleRevoked(log types.Log) (*TestSynapseBridgeRoleRevoked, error) {
	event := new(TestSynapseBridgeRoleRevoked)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenDepositIterator is returned from FilterTokenDeposit and is used to iterate over the raw logs and unpacked data for TokenDeposit events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenDepositIterator struct {
	Event *TestSynapseBridgeTokenDeposit // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenDeposit)
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
		it.Event = new(TestSynapseBridgeTokenDeposit)
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
func (it *TestSynapseBridgeTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenDeposit represents a TokenDeposit event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenDeposit struct {
	To      common.Address
	ChainId *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposit is a free log retrieval operation binding the contract event 0xda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b.
//
// Solidity: event TokenDeposit(address indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenDeposit(opts *bind.FilterOpts, to []common.Address) (*TestSynapseBridgeTokenDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenDeposit", toRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenDepositIterator{contract: _TestSynapseBridge.contract, event: "TokenDeposit", logs: logs, sub: sub}, nil
}

// WatchTokenDeposit is a free log subscription operation binding the contract event 0xda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b.
//
// Solidity: event TokenDeposit(address indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenDeposit(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenDeposit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenDeposit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenDeposit)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
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

// ParseTokenDeposit is a log parse operation binding the contract event 0xda5273705dbef4bf1b902a131c2eac086b7e1476a8ab0cb4da08af1fe1bd8e3b.
//
// Solidity: event TokenDeposit(address indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenDeposit(log types.Log) (*TestSynapseBridgeTokenDeposit, error) {
	event := new(TestSynapseBridgeTokenDeposit)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenDepositAndSwapIterator is returned from FilterTokenDepositAndSwap and is used to iterate over the raw logs and unpacked data for TokenDepositAndSwap events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenDepositAndSwapIterator struct {
	Event *TestSynapseBridgeTokenDepositAndSwap // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenDepositAndSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenDepositAndSwap)
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
		it.Event = new(TestSynapseBridgeTokenDepositAndSwap)
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
func (it *TestSynapseBridgeTokenDepositAndSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenDepositAndSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenDepositAndSwap represents a TokenDepositAndSwap event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenDepositAndSwap struct {
	To             common.Address
	ChainId        *big.Int
	Token          common.Address
	Amount         *big.Int
	TokenIndexFrom uint8
	TokenIndexTo   uint8
	MinDy          *big.Int
	Deadline       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenDepositAndSwap is a free log retrieval operation binding the contract event 0x79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f.
//
// Solidity: event TokenDepositAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenDepositAndSwap(opts *bind.FilterOpts, to []common.Address) (*TestSynapseBridgeTokenDepositAndSwapIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenDepositAndSwap", toRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenDepositAndSwapIterator{contract: _TestSynapseBridge.contract, event: "TokenDepositAndSwap", logs: logs, sub: sub}, nil
}

// WatchTokenDepositAndSwap is a free log subscription operation binding the contract event 0x79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f.
//
// Solidity: event TokenDepositAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenDepositAndSwap(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenDepositAndSwap, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenDepositAndSwap", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenDepositAndSwap)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenDepositAndSwap", log); err != nil {
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

// ParseTokenDepositAndSwap is a log parse operation binding the contract event 0x79c15604b92ef54d3f61f0c40caab8857927ca3d5092367163b4562c1699eb5f.
//
// Solidity: event TokenDepositAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenDepositAndSwap(log types.Log) (*TestSynapseBridgeTokenDepositAndSwap, error) {
	event := new(TestSynapseBridgeTokenDepositAndSwap)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenDepositAndSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenMintIterator is returned from FilterTokenMint and is used to iterate over the raw logs and unpacked data for TokenMint events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenMintIterator struct {
	Event *TestSynapseBridgeTokenMint // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenMint)
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
		it.Event = new(TestSynapseBridgeTokenMint)
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
func (it *TestSynapseBridgeTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenMint represents a TokenMint event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenMint struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Fee    *big.Int
	Kappa  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenMint is a free log retrieval operation binding the contract event 0xbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd38.
//
// Solidity: event TokenMint(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenMint(opts *bind.FilterOpts, to []common.Address, kappa [][32]byte) (*TestSynapseBridgeTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenMint", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenMintIterator{contract: _TestSynapseBridge.contract, event: "TokenMint", logs: logs, sub: sub}, nil
}

// WatchTokenMint is a free log subscription operation binding the contract event 0xbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd38.
//
// Solidity: event TokenMint(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenMint(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenMint, to []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenMint", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenMint)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenMint", log); err != nil {
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

// ParseTokenMint is a log parse operation binding the contract event 0xbf14b9fde87f6e1c29a7e0787ad1d0d64b4648d8ae63da21524d9fd0f283dd38.
//
// Solidity: event TokenMint(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenMint(log types.Log) (*TestSynapseBridgeTokenMint, error) {
	event := new(TestSynapseBridgeTokenMint)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenMintAndSwapIterator is returned from FilterTokenMintAndSwap and is used to iterate over the raw logs and unpacked data for TokenMintAndSwap events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenMintAndSwapIterator struct {
	Event *TestSynapseBridgeTokenMintAndSwap // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenMintAndSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenMintAndSwap)
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
		it.Event = new(TestSynapseBridgeTokenMintAndSwap)
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
func (it *TestSynapseBridgeTokenMintAndSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenMintAndSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenMintAndSwap represents a TokenMintAndSwap event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenMintAndSwap struct {
	To             common.Address
	Token          common.Address
	Amount         *big.Int
	Fee            *big.Int
	TokenIndexFrom uint8
	TokenIndexTo   uint8
	MinDy          *big.Int
	Deadline       *big.Int
	SwapSuccess    bool
	Kappa          [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenMintAndSwap is a free log retrieval operation binding the contract event 0x4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c65.
//
// Solidity: event TokenMintAndSwap(address indexed to, address token, uint256 amount, uint256 fee, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bool swapSuccess, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenMintAndSwap(opts *bind.FilterOpts, to []common.Address, kappa [][32]byte) (*TestSynapseBridgeTokenMintAndSwapIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenMintAndSwap", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenMintAndSwapIterator{contract: _TestSynapseBridge.contract, event: "TokenMintAndSwap", logs: logs, sub: sub}, nil
}

// WatchTokenMintAndSwap is a free log subscription operation binding the contract event 0x4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c65.
//
// Solidity: event TokenMintAndSwap(address indexed to, address token, uint256 amount, uint256 fee, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bool swapSuccess, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenMintAndSwap(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenMintAndSwap, to []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenMintAndSwap", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenMintAndSwap)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenMintAndSwap", log); err != nil {
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

// ParseTokenMintAndSwap is a log parse operation binding the contract event 0x4f56ec39e98539920503fd54ee56ae0cbebe9eb15aa778f18de67701eeae7c65.
//
// Solidity: event TokenMintAndSwap(address indexed to, address token, uint256 amount, uint256 fee, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline, bool swapSuccess, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenMintAndSwap(log types.Log) (*TestSynapseBridgeTokenMintAndSwap, error) {
	event := new(TestSynapseBridgeTokenMintAndSwap)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenMintAndSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenRedeemIterator is returned from FilterTokenRedeem and is used to iterate over the raw logs and unpacked data for TokenRedeem events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenRedeemIterator struct {
	Event *TestSynapseBridgeTokenRedeem // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenRedeem)
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
		it.Event = new(TestSynapseBridgeTokenRedeem)
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
func (it *TestSynapseBridgeTokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenRedeem represents a TokenRedeem event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenRedeem struct {
	To      common.Address
	ChainId *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeem is a free log retrieval operation binding the contract event 0xdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c.
//
// Solidity: event TokenRedeem(address indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenRedeem(opts *bind.FilterOpts, to []common.Address) (*TestSynapseBridgeTokenRedeemIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenRedeem", toRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenRedeemIterator{contract: _TestSynapseBridge.contract, event: "TokenRedeem", logs: logs, sub: sub}, nil
}

// WatchTokenRedeem is a free log subscription operation binding the contract event 0xdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c.
//
// Solidity: event TokenRedeem(address indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenRedeem(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenRedeem, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenRedeem", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenRedeem)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenRedeem", log); err != nil {
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

// ParseTokenRedeem is a log parse operation binding the contract event 0xdc5bad4651c5fbe9977a696aadc65996c468cde1448dd468ec0d83bf61c4b57c.
//
// Solidity: event TokenRedeem(address indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenRedeem(log types.Log) (*TestSynapseBridgeTokenRedeem, error) {
	event := new(TestSynapseBridgeTokenRedeem)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenRedeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenRedeemAndRemoveIterator is returned from FilterTokenRedeemAndRemove and is used to iterate over the raw logs and unpacked data for TokenRedeemAndRemove events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenRedeemAndRemoveIterator struct {
	Event *TestSynapseBridgeTokenRedeemAndRemove // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenRedeemAndRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenRedeemAndRemove)
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
		it.Event = new(TestSynapseBridgeTokenRedeemAndRemove)
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
func (it *TestSynapseBridgeTokenRedeemAndRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenRedeemAndRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenRedeemAndRemove represents a TokenRedeemAndRemove event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenRedeemAndRemove struct {
	To             common.Address
	ChainId        *big.Int
	Token          common.Address
	Amount         *big.Int
	SwapTokenIndex uint8
	SwapMinAmount  *big.Int
	SwapDeadline   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeemAndRemove is a free log retrieval operation binding the contract event 0x9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c.
//
// Solidity: event TokenRedeemAndRemove(address indexed to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenRedeemAndRemove(opts *bind.FilterOpts, to []common.Address) (*TestSynapseBridgeTokenRedeemAndRemoveIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenRedeemAndRemove", toRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenRedeemAndRemoveIterator{contract: _TestSynapseBridge.contract, event: "TokenRedeemAndRemove", logs: logs, sub: sub}, nil
}

// WatchTokenRedeemAndRemove is a free log subscription operation binding the contract event 0x9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c.
//
// Solidity: event TokenRedeemAndRemove(address indexed to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenRedeemAndRemove(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenRedeemAndRemove, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenRedeemAndRemove", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenRedeemAndRemove)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenRedeemAndRemove", log); err != nil {
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

// ParseTokenRedeemAndRemove is a log parse operation binding the contract event 0x9a7024cde1920aa50cdde09ca396229e8c4d530d5cfdc6233590def70a94408c.
//
// Solidity: event TokenRedeemAndRemove(address indexed to, uint256 chainId, address token, uint256 amount, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenRedeemAndRemove(log types.Log) (*TestSynapseBridgeTokenRedeemAndRemove, error) {
	event := new(TestSynapseBridgeTokenRedeemAndRemove)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenRedeemAndRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenRedeemAndSwapIterator is returned from FilterTokenRedeemAndSwap and is used to iterate over the raw logs and unpacked data for TokenRedeemAndSwap events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenRedeemAndSwapIterator struct {
	Event *TestSynapseBridgeTokenRedeemAndSwap // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenRedeemAndSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenRedeemAndSwap)
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
		it.Event = new(TestSynapseBridgeTokenRedeemAndSwap)
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
func (it *TestSynapseBridgeTokenRedeemAndSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenRedeemAndSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenRedeemAndSwap represents a TokenRedeemAndSwap event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenRedeemAndSwap struct {
	To             common.Address
	ChainId        *big.Int
	Token          common.Address
	Amount         *big.Int
	TokenIndexFrom uint8
	TokenIndexTo   uint8
	MinDy          *big.Int
	Deadline       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeemAndSwap is a free log retrieval operation binding the contract event 0x91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc319425.
//
// Solidity: event TokenRedeemAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenRedeemAndSwap(opts *bind.FilterOpts, to []common.Address) (*TestSynapseBridgeTokenRedeemAndSwapIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenRedeemAndSwap", toRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenRedeemAndSwapIterator{contract: _TestSynapseBridge.contract, event: "TokenRedeemAndSwap", logs: logs, sub: sub}, nil
}

// WatchTokenRedeemAndSwap is a free log subscription operation binding the contract event 0x91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc319425.
//
// Solidity: event TokenRedeemAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenRedeemAndSwap(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenRedeemAndSwap, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenRedeemAndSwap", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenRedeemAndSwap)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenRedeemAndSwap", log); err != nil {
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

// ParseTokenRedeemAndSwap is a log parse operation binding the contract event 0x91f25e9be0134ec851830e0e76dc71e06f9dade75a9b84e9524071dbbc319425.
//
// Solidity: event TokenRedeemAndSwap(address indexed to, uint256 chainId, address token, uint256 amount, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 minDy, uint256 deadline)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenRedeemAndSwap(log types.Log) (*TestSynapseBridgeTokenRedeemAndSwap, error) {
	event := new(TestSynapseBridgeTokenRedeemAndSwap)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenRedeemAndSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenRedeemV2Iterator is returned from FilterTokenRedeemV2 and is used to iterate over the raw logs and unpacked data for TokenRedeemV2 events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenRedeemV2Iterator struct {
	Event *TestSynapseBridgeTokenRedeemV2 // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenRedeemV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenRedeemV2)
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
		it.Event = new(TestSynapseBridgeTokenRedeemV2)
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
func (it *TestSynapseBridgeTokenRedeemV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenRedeemV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenRedeemV2 represents a TokenRedeemV2 event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenRedeemV2 struct {
	To      [32]byte
	ChainId *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeemV2 is a free log retrieval operation binding the contract event 0x8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499.
//
// Solidity: event TokenRedeemV2(bytes32 indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenRedeemV2(opts *bind.FilterOpts, to [][32]byte) (*TestSynapseBridgeTokenRedeemV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenRedeemV2", toRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenRedeemV2Iterator{contract: _TestSynapseBridge.contract, event: "TokenRedeemV2", logs: logs, sub: sub}, nil
}

// WatchTokenRedeemV2 is a free log subscription operation binding the contract event 0x8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499.
//
// Solidity: event TokenRedeemV2(bytes32 indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenRedeemV2(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenRedeemV2, to [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenRedeemV2", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenRedeemV2)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenRedeemV2", log); err != nil {
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

// ParseTokenRedeemV2 is a log parse operation binding the contract event 0x8e57e8c5fea426159af69d47eda6c5052c7605c9f70967cf749d4aa55b70b499.
//
// Solidity: event TokenRedeemV2(bytes32 indexed to, uint256 chainId, address token, uint256 amount)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenRedeemV2(log types.Log) (*TestSynapseBridgeTokenRedeemV2, error) {
	event := new(TestSynapseBridgeTokenRedeemV2)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenRedeemV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenWithdrawIterator struct {
	Event *TestSynapseBridgeTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenWithdraw)
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
		it.Event = new(TestSynapseBridgeTokenWithdraw)
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
func (it *TestSynapseBridgeTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenWithdraw represents a TokenWithdraw event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenWithdraw struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Fee    *big.Int
	Kappa  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26.
//
// Solidity: event TokenWithdraw(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenWithdraw(opts *bind.FilterOpts, to []common.Address, kappa [][32]byte) (*TestSynapseBridgeTokenWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenWithdraw", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenWithdrawIterator{contract: _TestSynapseBridge.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26.
//
// Solidity: event TokenWithdraw(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenWithdraw, to []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenWithdraw", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenWithdraw)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x8b0afdc777af6946e53045a4a75212769075d30455a212ac51c9b16f9c5c9b26.
//
// Solidity: event TokenWithdraw(address indexed to, address token, uint256 amount, uint256 fee, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenWithdraw(log types.Log) (*TestSynapseBridgeTokenWithdraw, error) {
	event := new(TestSynapseBridgeTokenWithdraw)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeTokenWithdrawAndRemoveIterator is returned from FilterTokenWithdrawAndRemove and is used to iterate over the raw logs and unpacked data for TokenWithdrawAndRemove events raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenWithdrawAndRemoveIterator struct {
	Event *TestSynapseBridgeTokenWithdrawAndRemove // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeTokenWithdrawAndRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeTokenWithdrawAndRemove)
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
		it.Event = new(TestSynapseBridgeTokenWithdrawAndRemove)
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
func (it *TestSynapseBridgeTokenWithdrawAndRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeTokenWithdrawAndRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeTokenWithdrawAndRemove represents a TokenWithdrawAndRemove event raised by the TestSynapseBridge contract.
type TestSynapseBridgeTokenWithdrawAndRemove struct {
	To             common.Address
	Token          common.Address
	Amount         *big.Int
	Fee            *big.Int
	SwapTokenIndex uint8
	SwapMinAmount  *big.Int
	SwapDeadline   *big.Int
	SwapSuccess    bool
	Kappa          [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawAndRemove is a free log retrieval operation binding the contract event 0xc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d.
//
// Solidity: event TokenWithdrawAndRemove(address indexed to, address token, uint256 amount, uint256 fee, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bool swapSuccess, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterTokenWithdrawAndRemove(opts *bind.FilterOpts, to []common.Address, kappa [][32]byte) (*TestSynapseBridgeTokenWithdrawAndRemoveIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "TokenWithdrawAndRemove", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeTokenWithdrawAndRemoveIterator{contract: _TestSynapseBridge.contract, event: "TokenWithdrawAndRemove", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawAndRemove is a free log subscription operation binding the contract event 0xc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d.
//
// Solidity: event TokenWithdrawAndRemove(address indexed to, address token, uint256 amount, uint256 fee, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bool swapSuccess, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchTokenWithdrawAndRemove(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeTokenWithdrawAndRemove, to []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "TokenWithdrawAndRemove", toRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeTokenWithdrawAndRemove)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenWithdrawAndRemove", log); err != nil {
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

// ParseTokenWithdrawAndRemove is a log parse operation binding the contract event 0xc1a608d0f8122d014d03cc915a91d98cef4ebaf31ea3552320430cba05211b6d.
//
// Solidity: event TokenWithdrawAndRemove(address indexed to, address token, uint256 amount, uint256 fee, uint8 swapTokenIndex, uint256 swapMinAmount, uint256 swapDeadline, bool swapSuccess, bytes32 indexed kappa)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseTokenWithdrawAndRemove(log types.Log) (*TestSynapseBridgeTokenWithdrawAndRemove, error) {
	event := new(TestSynapseBridgeTokenWithdrawAndRemove)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "TokenWithdrawAndRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSynapseBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TestSynapseBridge contract.
type TestSynapseBridgeUnpausedIterator struct {
	Event *TestSynapseBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *TestSynapseBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSynapseBridgeUnpaused)
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
		it.Event = new(TestSynapseBridgeUnpaused)
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
func (it *TestSynapseBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSynapseBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSynapseBridgeUnpaused represents a Unpaused event raised by the TestSynapseBridge contract.
type TestSynapseBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TestSynapseBridgeUnpausedIterator, error) {

	logs, sub, err := _TestSynapseBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TestSynapseBridgeUnpausedIterator{contract: _TestSynapseBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TestSynapseBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _TestSynapseBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSynapseBridgeUnpaused)
				if err := _TestSynapseBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TestSynapseBridge *TestSynapseBridgeFilterer) ParseUnpaused(log types.Log) (*TestSynapseBridgeUnpaused, error) {
	event := new(TestSynapseBridgeUnpaused)
	if err := _TestSynapseBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
