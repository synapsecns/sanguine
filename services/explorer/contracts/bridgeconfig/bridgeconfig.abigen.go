// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeconfig

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

// BridgeConfigV3Pool is an auto generated low-level Go binding around an user-defined struct.
type BridgeConfigV3Pool struct {
	TokenAddress common.Address
	ChainId      *big.Int
	PoolAddress  common.Address
	Metaswap     bool
}

// BridgeConfigV3Token is an auto generated low-level Go binding around an user-defined struct.
type BridgeConfigV3Token struct {
	ChainId       *big.Int
	TokenAddress  string
	TokenDecimals uint8
	MaxSwap       *big.Int
	MinSwap       *big.Int
	SwapFee       *big.Int
	MaxSwapFee    *big.Int
	MinSwapFee    *big.Int
	HasUnderlying bool
	IsUnderlying  bool
}

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
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

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

// Deprecated: Use AccessControlMetaData.Sigs instead.
// AccessControlFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlFuncSigs = AccessControlMetaData.Sigs

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControl.Contract.GetRoleMember(&_AccessControl.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControl *AccessControlCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControl.Contract.GetRoleMember(&_AccessControl.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.GetRoleMemberCount(&_AccessControl.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.GetRoleMemberCount(&_AccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// AccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControl contract.
type AccessControlRoleAdminChangedIterator struct {
	Event *AccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleAdminChanged)
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
		it.Event = new(AccessControlRoleAdminChanged)
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
func (it *AccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControl contract.
type AccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleAdminChangedIterator{contract: _AccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleAdminChanged)
				if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlRoleAdminChanged, error) {
	event := new(AccessControlRoleAdminChanged)
	if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleGrantedIterator, error) {

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

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleRevokedIterator, error) {

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

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207bff0e2bd493579e1cca2e287bbaf08915fed4ee1fd5bdda0a91d5b8feca389164736f6c634300060c0033",
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

// BridgeConfigV3MetaData contains all meta data concerning the BridgeConfigV3 contract.
var BridgeConfigV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGEMANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeConfigVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTokenIDs\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"result\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getPoolConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"metaswap\",\"type\":\"bool\"}],\"internalType\":\"structBridgeConfigV3.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasUnderlying\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUnderlying\",\"type\":\"bool\"}],\"internalType\":\"structBridgeConfigV3.Token\",\"name\":\"token\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getTokenByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasUnderlying\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUnderlying\",\"type\":\"bool\"}],\"internalType\":\"structBridgeConfigV3.Token\",\"name\":\"token\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getTokenByEVMAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasUnderlying\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUnderlying\",\"type\":\"bool\"}],\"internalType\":\"structBridgeConfigV3.Token\",\"name\":\"token\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getTokenByID\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasUnderlying\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUnderlying\",\"type\":\"bool\"}],\"internalType\":\"structBridgeConfigV3.Token\",\"name\":\"token\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getTokenID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getTokenID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"getUnderlyingToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasUnderlying\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUnderlying\",\"type\":\"bool\"}],\"internalType\":\"structBridgeConfigV3.Token\",\"name\":\"token\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"hasUnderlyingToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"}],\"name\":\"isTokenIDExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"setMaxGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"metaswap\",\"type\":\"bool\"}],\"name\":\"setPoolConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"metaswap\",\"type\":\"bool\"}],\"internalType\":\"structBridgeConfigV3.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasUnderlying\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUnderlying\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenAddress\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasUnderlying\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isUnderlying\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ff9106c7": "BRIDGEMANAGER_ROLE()",
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"2c02799e": "bridgeConfigVersion()",
		"fc7cc4cb": "calculateSwapFee(address,uint256,uint256)",
		"0a62a9cb": "calculateSwapFee(string,uint256,uint256)",
		"684a10b3": "getAllTokenIDs()",
		"fd534b33": "getMaxGasPrice(uint256)",
		"72fb43d9": "getPoolConfig(address,uint256)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"324980b5": "getToken(string,uint256)",
		"e814157d": "getTokenByAddress(string,uint256)",
		"558dae3a": "getTokenByEVMAddress(address,uint256)",
		"77b8cbf7": "getTokenByID(string,uint256)",
		"3cc1c7e0": "getTokenID(address,uint256)",
		"efd7516e": "getTokenID(string,uint256)",
		"58dfe6f1": "getUnderlyingToken(string)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"074b7e97": "hasUnderlyingToken(string)",
		"af611ca0": "isTokenIDExist(string)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"abaac008": "setMaxGasPrice(uint256,uint256)",
		"7e355e5e": "setPoolConfig(address,uint256,address,bool)",
		"59053bfe": "setTokenConfig(string,uint256,address,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool)",
		"ddb54399": "setTokenConfig(string,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506200001f60003362000025565b62000139565b62000031828262000035565b5050565b6000828152602081815260409091206200005a91839062001098620000ae821b17901c565b1562000031576200006a620000ce565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000620000c5836001600160a01b038416620000d2565b90505b92915050565b3390565b6000620000e0838362000121565b6200011857508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620000c8565b506000620000c8565b60009081526001919091016020526040902054151590565b612ae180620001496000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80637e355e5e116100f9578063d547741f11610097578063efd7516e11610071578063efd7516e146103c3578063fc7cc4cb146103d6578063fd534b33146103e9578063ff9106c7146103fc576101c4565b8063d547741f1461038a578063ddb543991461039d578063e814157d146103b0576101c4565b8063a217fddf116100d3578063a217fddf14610349578063abaac00814610351578063af611ca014610364578063ca15c87314610377576101c4565b80637e355e5e146103035780639010d07c1461031657806391d1485414610336576101c4565b80633cc1c7e01161016657806359053bfe1161014057806359053bfe146102bb578063684a10b3146102ce57806372fb43d9146102e357806377b8cbf714610242576101c4565b80633cc1c7e014610275578063558dae3a1461029557806358dfe6f1146102a8576101c4565b80632c02799e116101a25780632c02799e146102255780632f2ff15d1461022d578063324980b51461024257806336568abe14610262576101c4565b8063074b7e97146101c95780630a62a9cb146101f2578063248a9ca314610212575b600080fd5b6101dc6101d7366004612551565b610404565b6040516101e99190612725565b60405180910390f35b6102056102003660046125c7565b6105d3565b6040516101e99190612730565b610205610220366004612397565b6105f2565b610205610607565b61024061023b3660046123af565b61060c565b005b610255610250366004612584565b610677565b6040516101e991906129d7565b6102406102703660046123af565b6107cb565b6102886102833660046122e6565b610841565b6040516101e99190612739565b6102556102a33660046122e6565b61085d565b6102556102b6366004612551565b610942565b6101dc6102c9366004612417565b610b25565b6102d6610b87565b6040516101e991906126a7565b6102f66102f13660046122e6565b610c2a565b6040516101e9919061298b565b6102f6610311366004612310565b610cb4565b6103296103243660046123f6565b610e0f565b6040516101e99190612686565b6101dc6103443660046123af565b610e27565b610205610e3f565b61024061035f3660046123f6565b610e44565b6101dc610372366004612551565b610e89565b610205610385366004612397565b610e9c565b6102406103983660046123af565b610eb3565b6101dc6103ab3660046124d9565b610f07565b6102556103be366004612584565b611020565b6102886103d1366004612584565b611043565b6102056103e4366004612364565b611051565b6102056103f7366004612397565b611062565b610205611074565b600080610410836110ba565b600081815260026020908152604080832080548251818502810185019093528083529495506060949193909284015b8282101561058257838290600052602060002090600902016040518061014001604052908160008201548152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105065780601f106104db57610100808354040283529160200191610506565b820191906000526020600020905b8154815290600101906020018083116104e957829003601f168201915b5050509183525050600282015460ff908116602080840191909152600384015460408401526004840154606084015260058401546080840152600684015460a0840152600784015460c0840152600890930154808216151560e0840152610100908190049091161515910152908252600192909201910161043f565b50505050905060005b81518110156105c6578181815181106105a057fe5b60200260200101516101000151156105be57600193505050506105ce565b60010161058b565b506000925050505b919050565b60006105e86105e1856110d3565b848461122d565b90505b9392505050565b60009081526020819052604090206002015490565b600381565b60008281526020819052604090206002015461062a90610344611410565b610669576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610660906127a9565b60405180910390fd5b6106738282611414565b5050565b61067f6120aa565b6004600061068c856110ba565b815260200190815260200160002060008381526020019081526020016000206040518061014001604052908160008201548152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561075a5780601f1061072f5761010080835404028352916020019161075a565b820191906000526020600020905b81548152906001019060200180831161073d57829003601f168201915b5050509183525050600282015460ff9081166020830152600383015460408301526004830154606083015260058301546080830152600683015460a0830152600783015460c0830152600890920154808316151560e083015261010090819004909216151591015290505b92915050565b6107d3611410565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610837576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106609061292e565b6106738282611497565b60606105eb6108576108528561151a565b6110d3565b8361168b565b6108656120aa565b6000828152600360205260408120600491906108836108528761151a565b604051610890919061266a565b908152602001604051809103902054815260200190815260200160002060008381526020019081526020016000206040518061014001604052908160008201548152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561075a5780601f1061072f5761010080835404028352916020019161075a565b61094a6120aa565b6000610955836110ba565b600081815260026020908152604080832080548251818502810185019093528083529495506060949193909284015b82821015610ac757838290600052602060002090600902016040518061014001604052908160008201548152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a4b5780601f10610a2057610100808354040283529160200191610a4b565b820191906000526020600020905b815481529060010190602001808311610a2e57829003601f168201915b5050509183525050600282015460ff908116602080840191909152600384015460408401526004840154606084015260058401546080840152600684015460a0840152600784015460c0840152600890930154808216151560e08401526101009081900490911615159101529082526001929092019101610984565b50505050905060005b8151811015610b1d57818181518110610ae557fe5b6020026020010151610120015115610b1557818181518110610b0357fe5b602002602001015193505050506105ce565b600101610ad0565b505050919050565b6000610b517f4370dcf3e42e4d5b773a451bb8390ee8e7308f47681d1414cff87c2ad0512c8533610e27565b610b5a57600080fd5b610b768d8d8d610b698e61151a565b8d8d8d8d8d8d8d8d610f07565b9d9c50505050505050505050505050565b6001546060908067ffffffffffffffff81118015610ba457600080fd5b50604051908082528060200260200182016040528015610bd857816020015b6060815260200190600190039081610bc35790505b50915060005b81811015610c2557610c0660018281548110610bf657fe5b90600052602060002001546116c1565b838281518110610c1257fe5b6020908102919091010152600101610bde565b505090565b610c32612104565b5073ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093835292815290829020825160808101845281548516815260018201549281019290925260020154928316918101919091527401000000000000000000000000000000000000000090910460ff161515606082015290565b610cbc612104565b610ce67f4370dcf3e42e4d5b773a451bb8390ee8e7308f47681d1414cff87c2ad0512c8533610e27565b610d1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610660906128f7565b610d24612104565b50506040805160808101825273ffffffffffffffffffffffffffffffffffffffff8087168083526020808401888152878416858701908152871515606087019081526000948552600584528785208b865290935295909220845181549085167fffffffffffffffffffffffff0000000000000000000000000000000000000000918216178255925160018201559451600290950180549151151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9690941691909216179390931617909155949350505050565b60008281526020819052604081206105eb90836117cd565b60008281526020819052604081206105eb90836117d9565b600081565b610e6e7f4370dcf3e42e4d5b773a451bb8390ee8e7308f47681d1414cff87c2ad0512c8533610e27565b610e7757600080fd5b60009182526006602052604090912055565b60006107c5610e97836110ba565b6117fb565b60008181526020819052604081206107c590611843565b600082815260208190526040902060020154610ed190610344611410565b610837576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106609061283d565b6000610f337f4370dcf3e42e4d5b773a451bb8390ee8e7308f47681d1414cff87c2ad0512c8533610e27565b610f3c57600080fd5b610f446120aa565b610f4d8b6110d3565b816020018190525089816040019060ff16908160ff16815250508881606001818152505087816080018181525050868160a0018181525050858160c0018181525050848160e0018181525050838161010001901515908115158152505082816101200190151590811515815250508b81600001818152505061100e6110078f8f8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506110ba92505050565b8d8361184e565b9e9d5050505050505050505050505050565b6110286120aa565b600082815260036020526040812060049190610883866110d3565b60606105eb610857846110d3565b60006105e86105e16108528661151a565b60009081526006602052604090205490565b7f4370dcf3e42e4d5b773a451bb8390ee8e7308f47681d1414cff87c2ad0512c8581565b60006105eb8373ffffffffffffffffffffffffffffffffffffffff8416611c77565b60006020825111156110cb57600080fd5b506020015190565b6060808290506060815167ffffffffffffffff811180156110f357600080fd5b506040519080825280601f01601f19166020018201604052801561111e576020820181803683370190505b50905060005b825181101561122557604183828151811061113b57fe5b016020015160f81c108015906111655750605a83828151811061115a57fe5b016020015160f81c11155b156111ca5782818151811061117657fe5b602001015160f81c60f81b60f81c60200160f81b82828151811061119657fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061121d565b8281815181106111d657fe5b602001015160f81c60f81b8282815181106111ed57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505b600101611124565b509392505050565b60006112376120aa565b6000848152600360205260408082209051600492919061125890899061266a565b908152602001604051809103902054815260200190815260200160002060008581526020019081526020016000206040518061014001604052908160008201548152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113355780601f1061130a57610100808354040283529160200191611335565b820191906000526020600020905b81548152906001019060200180831161131857829003601f168201915b5050509183525050600282015460ff9081166020830152600383015460408301526004830154606083015260058301546080830152600683015460a080840191909152600784015460c0840152600890930154808216151560e08401526101009081900490911615159101528101519091506000906113c6906402540be400906113c0908790611cc1565b90611d15565b90508160e00151811180156113de57508160c0015181105b156113ec5791506105eb9050565b8160c00151811115611404575060c0015190506105eb565b5060e0015190506105eb565b3390565b600082815260208190526040902061142c9082611098565b1561067357611439611410565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526020819052604090206114af9082611d61565b15610673576114bc611410565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b604080516028808252606082810190935282919060208201818036833701905050905060005b60148110156116445760008160130360080260020a8573ffffffffffffffffffffffffffffffffffffffff168161157357fe5b0460f81b9050600060108260f81c60ff168161158b57fe5b0460f81b905060008160f81c6010028360f81c0360f81b90506115ad82611d83565b8585600202815181106115bc57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506115f481611d83565b85856002026001018151811061160657fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535050600190920191506115409050565b5060408051808201909152600281527f307800000000000000000000000000000000000000000000000000000000000060208201526116838183611db1565b949350505050565b60606105eb60036000848152602001908152602001600020846040516116b1919061266a565b9081526020016040518091039020545b606060005b60208160ff1610801561170c5750828160ff16602081106116e357fe5b1a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615155b15611719576001016116c6565b60608160ff1667ffffffffffffffff8111801561173557600080fd5b506040519080825280601f01601f191660200182016040528015611760576020820181803683370190505b50905060005b8260ff168160ff16101561122557848160ff166020811061178357fe5b1a60f81b828260ff168151811061179657fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101611766565b60006105eb8383611ef2565b60006105eb8373ffffffffffffffffffffffffffffffffffffffff8416611f51565b6000805b60015481101561183a57826001828154811061181757fe5b906000526020600020015414156118325760019150506105ce565b6001016117ff565b50600092915050565b60006107c582611f69565b60008381526004602090815260408083208584528252822083518155818401518051859361188392600185019291019061212b565b50604082015160028201805460ff9092167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00928316179055606083015160038301556080830151600483015560a0830151600583015560c0830151600683015560e0830151600783015561010080840151600890930180546101209095015115159091027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff931515949092169390931791909116179055611943846117fb565b61197c576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018490555b6000848152600260205260408120905b8154811015611b4857848282815481106119a257fe5b9060005260206000209060090201600001541415611b405760608282815481106119c857fe5b90600052602060002090600902016001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611a6d5780601f10611a4257610100808354040283529160200191611a6d565b820191906000526020600020905b815481529060010190602001808311611a5057829003601f168201915b50505050509050611a82856020015182611f6d565b611b3e578460200151838381548110611a9757fe5b90600052602060002090600902016001019080519060200190611abb92919061212b565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4706003600088815260200190815260200160002082604051611afe919061266a565b90815260408051602092819003830181209390935560008981526003835220908701518992611b2d919061266a565b908152604051908190036020019020555b505b60010161198c565b5080546001818101835560008381526020908190208651600990940201928355808601518051879493611b809390850192019061212b565b5060408281015160028301805460ff9092167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0092831617905560608401516003808501919091556080850151600485015560a0850151600585015560c0850151600685015560e0850151600785015561010080860151600890950180546101209097015115159091027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff95151596909316959095179390931617909255600086815260209182528290209085015191518792611c5b9161266a565b9081526040519081900360200190205550600190509392505050565b6000611c838383611f51565b611cb9575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556107c5565b5060006107c5565b600082611cd0575060006107c5565b82820282848281611cdd57fe5b04146105eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106609061289a565b6000808211611d50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066090612806565b818381611d5957fe5b049392505050565b60006105eb8373ffffffffffffffffffffffffffffffffffffffff8416611fc6565b6000600a60f883901c1015611da3578160f81c60300160f81b90506105ce565b50605760f891821c01901b90565b805182516060918491849184910167ffffffffffffffff81118015611dd557600080fd5b506040519080825280601f01601f191660200182016040528015611e00576020820181803683370190505b509050806000805b8551821015611e7457858281518110611e1d57fe5b602001015160f81c60f81b838280600101935081518110611e3a57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600190910190611e08565b600091505b8451821015611ee557848281518110611e8e57fe5b602001015160f81c60f81b838280600101935081518110611eab57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600190910190611e79565b5090979650505050505050565b81546000908210611f2f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106609061274c565b826000018281548110611f3e57fe5b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b5490565b600081604051602001611f80919061266a565b6040516020818303038152906040528051906020012083604051602001611fa7919061266a565b6040516020818303038152906040528051906020012014905092915050565b600081815260018301602052604081205480156120a05783547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808301919081019060009087908390811061201757fe5b906000526020600020015490508087600001848154811061203457fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061206457fe5b600190038181906000526020600020016000905590558660010160008781526020019081526020016000206000905560019450505050506107c5565b60009150506107c5565b6040518061014001604052806000815260200160608152602001600060ff16815260200160008152602001600081526020016000815260200160008152602001600081526020016000151581526020016000151581525090565b60408051608081018252600080825260208201819052918101829052606081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061216c57805160ff1916838001178555612199565b82800160010185558215612199579182015b8281111561219957825182559160200191906001019061217e565b506121a59291506121a9565b5090565b5b808211156121a557600081556001016121aa565b803573ffffffffffffffffffffffffffffffffffffffff811681146107c557600080fd5b803580151581146107c557600080fd5b60008083601f840112612203578182fd5b50813567ffffffffffffffff81111561221a578182fd5b60208301915083602082850101111561223257600080fd5b9250929050565b600082601f830112612249578081fd5b813567ffffffffffffffff80821115612260578283fd5b60405160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f850116820101818110838211171561229e578485fd5b6040528281529250828483016020018610156122b957600080fd5b8260208601602083013760006020848301015250505092915050565b803560ff811681146107c557600080fd5b600080604083850312156122f8578182fd5b61230284846121be565b946020939093013593505050565b60008060008060808587031215612325578182fd5b61232f86866121be565b93506020850135925061234586604087016121be565b915060608501358015158114612359578182fd5b939692955090935050565b600080600060608486031215612378578283fd5b61238285856121be565b95602085013595506040909401359392505050565b6000602082840312156123a8578081fd5b5035919050565b600080604083850312156123c1578182fd5b82359150602083013573ffffffffffffffffffffffffffffffffffffffff811681146123eb578182fd5b809150509250929050565b60008060408385031215612408578182fd5b50508035926020909101359150565b6000806000806000806000806000806000806101608d8f031215612439578788fd5b67ffffffffffffffff8d35111561244e578788fd5b61245b8e8e358f016121f2565b909c509a5060208d013599506124748e60408f016121be565b98506124838e60608f016122d5565b975060808d0135965060a08d0135955060c08d0135945060e08d013593506101008d013592506124b78e6101208f016121e2565b91506124c78e6101408f016121e2565b90509295989b509295989b509295989b565b6000806000806000806000806000806000806101608d8f0312156124fb578081fd5b67ffffffffffffffff8d351115612510578081fd5b61251d8e8e358f016121f2565b909c509a5060208d0135995067ffffffffffffffff60408e01351115612541578081fd5b6124748e60408f01358f01612239565b600060208284031215612562578081fd5b813567ffffffffffffffff811115612578578182fd5b61168384828501612239565b60008060408385031215612596578182fd5b823567ffffffffffffffff8111156125ac578283fd5b6125b885828601612239565b95602094909401359450505050565b6000806000606084860312156125db578081fd5b833567ffffffffffffffff8111156125f1578182fd5b6125fd86828701612239565b9660208601359650604090950135949350505050565b15159052565b60008151808452612631816020860160208601612a7b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60ff169052565b6000825161267c818460208701612a7b565b9190910192915050565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015612718577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0888603018452612706858351612619565b945092850192908501906001016126cc565b5092979650505050505050565b901515815260200190565b90815260200190565b6000602082526105eb6020830184612619565b60208082526022908201527f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e60408201527f6473000000000000000000000000000000000000000000000000000000000000606082015260800190565b6020808252602f908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201527f2061646d696e20746f206772616e740000000000000000000000000000000000606082015260800190565b6020808252601a908201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604082015260600190565b60208082526030908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201527f2061646d696e20746f207265766f6b6500000000000000000000000000000000606082015260800190565b60208082526021908201527f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60408201527f7700000000000000000000000000000000000000000000000000000000000000606082015260800190565b6020808252601c908201527f43616c6c6572206973206e6f7420427269646765204d616e6167657200000000604082015260600190565b6020808252602f908201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560408201527f20726f6c657320666f722073656c660000000000000000000000000000000000606082015260800190565b600060808201905073ffffffffffffffffffffffffffffffffffffffff808451168352602084015160208401528060408501511660408401525060608301511515606083015292915050565b600060208252825160208301526020830151610140806040850152612a00610160850183612619565b91506040850151612a146060860182612663565b5060608501516080850152608085015160a085015260a085015160c085015260c085015160e085015260e0850151610100818187015280870151915050610120612a6081870183612613565b8601519050612a7185830182612613565b5090949350505050565b60005b83811015612a96578181015183820152602001612a7e565b83811115612aa5576000848401525b5050505056fea2646970667358221220183585375986449f7549fd6c95f8d4314e8678e4a8d4e150abebb12f59796dca64736f6c634300060c0033",
}

// BridgeConfigV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeConfigV3MetaData.ABI instead.
var BridgeConfigV3ABI = BridgeConfigV3MetaData.ABI

// Deprecated: Use BridgeConfigV3MetaData.Sigs instead.
// BridgeConfigV3FuncSigs maps the 4-byte function signature to its string representation.
var BridgeConfigV3FuncSigs = BridgeConfigV3MetaData.Sigs

// BridgeConfigV3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeConfigV3MetaData.Bin instead.
var BridgeConfigV3Bin = BridgeConfigV3MetaData.Bin

// DeployBridgeConfigV3 deploys a new Ethereum contract, binding an instance of BridgeConfigV3 to it.
func DeployBridgeConfigV3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeConfigV3, error) {
	parsed, err := BridgeConfigV3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeConfigV3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeConfigV3{BridgeConfigV3Caller: BridgeConfigV3Caller{contract: contract}, BridgeConfigV3Transactor: BridgeConfigV3Transactor{contract: contract}, BridgeConfigV3Filterer: BridgeConfigV3Filterer{contract: contract}}, nil
}

// BridgeConfigV3 is an auto generated Go binding around an Ethereum contract.
type BridgeConfigV3 struct {
	BridgeConfigV3Caller     // Read-only binding to the contract
	BridgeConfigV3Transactor // Write-only binding to the contract
	BridgeConfigV3Filterer   // Log filterer for contract events
}

// BridgeConfigV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeConfigV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeConfigV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeConfigV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeConfigV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeConfigV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeConfigV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeConfigV3Session struct {
	Contract     *BridgeConfigV3   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeConfigV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeConfigV3CallerSession struct {
	Contract *BridgeConfigV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeConfigV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeConfigV3TransactorSession struct {
	Contract     *BridgeConfigV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeConfigV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeConfigV3Raw struct {
	Contract *BridgeConfigV3 // Generic contract binding to access the raw methods on
}

// BridgeConfigV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeConfigV3CallerRaw struct {
	Contract *BridgeConfigV3Caller // Generic read-only contract binding to access the raw methods on
}

// BridgeConfigV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeConfigV3TransactorRaw struct {
	Contract *BridgeConfigV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeConfigV3 creates a new instance of BridgeConfigV3, bound to a specific deployed contract.
func NewBridgeConfigV3(address common.Address, backend bind.ContractBackend) (*BridgeConfigV3, error) {
	contract, err := bindBridgeConfigV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigV3{BridgeConfigV3Caller: BridgeConfigV3Caller{contract: contract}, BridgeConfigV3Transactor: BridgeConfigV3Transactor{contract: contract}, BridgeConfigV3Filterer: BridgeConfigV3Filterer{contract: contract}}, nil
}

// NewBridgeConfigV3Caller creates a new read-only instance of BridgeConfigV3, bound to a specific deployed contract.
func NewBridgeConfigV3Caller(address common.Address, caller bind.ContractCaller) (*BridgeConfigV3Caller, error) {
	contract, err := bindBridgeConfigV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigV3Caller{contract: contract}, nil
}

// NewBridgeConfigV3Transactor creates a new write-only instance of BridgeConfigV3, bound to a specific deployed contract.
func NewBridgeConfigV3Transactor(address common.Address, transactor bind.ContractTransactor) (*BridgeConfigV3Transactor, error) {
	contract, err := bindBridgeConfigV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigV3Transactor{contract: contract}, nil
}

// NewBridgeConfigV3Filterer creates a new log filterer instance of BridgeConfigV3, bound to a specific deployed contract.
func NewBridgeConfigV3Filterer(address common.Address, filterer bind.ContractFilterer) (*BridgeConfigV3Filterer, error) {
	contract, err := bindBridgeConfigV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigV3Filterer{contract: contract}, nil
}

// bindBridgeConfigV3 binds a generic wrapper to an already deployed contract.
func bindBridgeConfigV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeConfigV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeConfigV3 *BridgeConfigV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeConfigV3.Contract.BridgeConfigV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeConfigV3 *BridgeConfigV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.BridgeConfigV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeConfigV3 *BridgeConfigV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.BridgeConfigV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeConfigV3 *BridgeConfigV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeConfigV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeConfigV3 *BridgeConfigV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeConfigV3 *BridgeConfigV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEMANAGERROLE is a free data retrieval call binding the contract method 0xff9106c7.
//
// Solidity: function BRIDGEMANAGER_ROLE() view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3Caller) BRIDGEMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "BRIDGEMANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BRIDGEMANAGERROLE is a free data retrieval call binding the contract method 0xff9106c7.
//
// Solidity: function BRIDGEMANAGER_ROLE() view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3Session) BRIDGEMANAGERROLE() ([32]byte, error) {
	return _BridgeConfigV3.Contract.BRIDGEMANAGERROLE(&_BridgeConfigV3.CallOpts)
}

// BRIDGEMANAGERROLE is a free data retrieval call binding the contract method 0xff9106c7.
//
// Solidity: function BRIDGEMANAGER_ROLE() view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) BRIDGEMANAGERROLE() ([32]byte, error) {
	return _BridgeConfigV3.Contract.BRIDGEMANAGERROLE(&_BridgeConfigV3.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeConfigV3.Contract.DEFAULTADMINROLE(&_BridgeConfigV3.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeConfigV3.Contract.DEFAULTADMINROLE(&_BridgeConfigV3.CallOpts)
}

// BridgeConfigVersion is a free data retrieval call binding the contract method 0x2c02799e.
//
// Solidity: function bridgeConfigVersion() view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Caller) BridgeConfigVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "bridgeConfigVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeConfigVersion is a free data retrieval call binding the contract method 0x2c02799e.
//
// Solidity: function bridgeConfigVersion() view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Session) BridgeConfigVersion() (*big.Int, error) {
	return _BridgeConfigV3.Contract.BridgeConfigVersion(&_BridgeConfigV3.CallOpts)
}

// BridgeConfigVersion is a free data retrieval call binding the contract method 0x2c02799e.
//
// Solidity: function bridgeConfigVersion() view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) BridgeConfigVersion() (*big.Int, error) {
	return _BridgeConfigV3.Contract.BridgeConfigVersion(&_BridgeConfigV3.CallOpts)
}

// CalculateSwapFee is a free data retrieval call binding the contract method 0x0a62a9cb.
//
// Solidity: function calculateSwapFee(string tokenAddress, uint256 chainID, uint256 amount) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Caller) CalculateSwapFee(opts *bind.CallOpts, tokenAddress string, chainID *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "calculateSwapFee", tokenAddress, chainID, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwapFee is a free data retrieval call binding the contract method 0x0a62a9cb.
//
// Solidity: function calculateSwapFee(string tokenAddress, uint256 chainID, uint256 amount) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Session) CalculateSwapFee(tokenAddress string, chainID *big.Int, amount *big.Int) (*big.Int, error) {
	return _BridgeConfigV3.Contract.CalculateSwapFee(&_BridgeConfigV3.CallOpts, tokenAddress, chainID, amount)
}

// CalculateSwapFee is a free data retrieval call binding the contract method 0x0a62a9cb.
//
// Solidity: function calculateSwapFee(string tokenAddress, uint256 chainID, uint256 amount) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) CalculateSwapFee(tokenAddress string, chainID *big.Int, amount *big.Int) (*big.Int, error) {
	return _BridgeConfigV3.Contract.CalculateSwapFee(&_BridgeConfigV3.CallOpts, tokenAddress, chainID, amount)
}

// CalculateSwapFee0 is a free data retrieval call binding the contract method 0xfc7cc4cb.
//
// Solidity: function calculateSwapFee(address tokenAddress, uint256 chainID, uint256 amount) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Caller) CalculateSwapFee0(opts *bind.CallOpts, tokenAddress common.Address, chainID *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "calculateSwapFee0", tokenAddress, chainID, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwapFee0 is a free data retrieval call binding the contract method 0xfc7cc4cb.
//
// Solidity: function calculateSwapFee(address tokenAddress, uint256 chainID, uint256 amount) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Session) CalculateSwapFee0(tokenAddress common.Address, chainID *big.Int, amount *big.Int) (*big.Int, error) {
	return _BridgeConfigV3.Contract.CalculateSwapFee0(&_BridgeConfigV3.CallOpts, tokenAddress, chainID, amount)
}

// CalculateSwapFee0 is a free data retrieval call binding the contract method 0xfc7cc4cb.
//
// Solidity: function calculateSwapFee(address tokenAddress, uint256 chainID, uint256 amount) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) CalculateSwapFee0(tokenAddress common.Address, chainID *big.Int, amount *big.Int) (*big.Int, error) {
	return _BridgeConfigV3.Contract.CalculateSwapFee0(&_BridgeConfigV3.CallOpts, tokenAddress, chainID, amount)
}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetAllTokenIDs(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getAllTokenIDs")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetAllTokenIDs() ([]string, error) {
	return _BridgeConfigV3.Contract.GetAllTokenIDs(&_BridgeConfigV3.CallOpts)
}

// GetAllTokenIDs is a free data retrieval call binding the contract method 0x684a10b3.
//
// Solidity: function getAllTokenIDs() view returns(string[] result)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetAllTokenIDs() ([]string, error) {
	return _BridgeConfigV3.Contract.GetAllTokenIDs(&_BridgeConfigV3.CallOpts)
}

// GetMaxGasPrice is a free data retrieval call binding the contract method 0xfd534b33.
//
// Solidity: function getMaxGasPrice(uint256 chainID) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetMaxGasPrice(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getMaxGasPrice", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxGasPrice is a free data retrieval call binding the contract method 0xfd534b33.
//
// Solidity: function getMaxGasPrice(uint256 chainID) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetMaxGasPrice(chainID *big.Int) (*big.Int, error) {
	return _BridgeConfigV3.Contract.GetMaxGasPrice(&_BridgeConfigV3.CallOpts, chainID)
}

// GetMaxGasPrice is a free data retrieval call binding the contract method 0xfd534b33.
//
// Solidity: function getMaxGasPrice(uint256 chainID) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetMaxGasPrice(chainID *big.Int) (*big.Int, error) {
	return _BridgeConfigV3.Contract.GetMaxGasPrice(&_BridgeConfigV3.CallOpts, chainID)
}

// GetPoolConfig is a free data retrieval call binding the contract method 0x72fb43d9.
//
// Solidity: function getPoolConfig(address tokenAddress, uint256 chainID) view returns((address,uint256,address,bool))
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetPoolConfig(opts *bind.CallOpts, tokenAddress common.Address, chainID *big.Int) (BridgeConfigV3Pool, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getPoolConfig", tokenAddress, chainID)

	if err != nil {
		return *new(BridgeConfigV3Pool), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeConfigV3Pool)).(*BridgeConfigV3Pool)

	return out0, err

}

// GetPoolConfig is a free data retrieval call binding the contract method 0x72fb43d9.
//
// Solidity: function getPoolConfig(address tokenAddress, uint256 chainID) view returns((address,uint256,address,bool))
func (_BridgeConfigV3 *BridgeConfigV3Session) GetPoolConfig(tokenAddress common.Address, chainID *big.Int) (BridgeConfigV3Pool, error) {
	return _BridgeConfigV3.Contract.GetPoolConfig(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetPoolConfig is a free data retrieval call binding the contract method 0x72fb43d9.
//
// Solidity: function getPoolConfig(address tokenAddress, uint256 chainID) view returns((address,uint256,address,bool))
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetPoolConfig(tokenAddress common.Address, chainID *big.Int) (BridgeConfigV3Pool, error) {
	return _BridgeConfigV3.Contract.GetPoolConfig(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeConfigV3.Contract.GetRoleAdmin(&_BridgeConfigV3.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeConfigV3.Contract.GetRoleAdmin(&_BridgeConfigV3.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _BridgeConfigV3.Contract.GetRoleMember(&_BridgeConfigV3.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _BridgeConfigV3.Contract.GetRoleMember(&_BridgeConfigV3.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BridgeConfigV3.Contract.GetRoleMemberCount(&_BridgeConfigV3.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BridgeConfigV3.Contract.GetRoleMemberCount(&_BridgeConfigV3.CallOpts, role)
}

// GetToken is a free data retrieval call binding the contract method 0x324980b5.
//
// Solidity: function getToken(string tokenID, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetToken(opts *bind.CallOpts, tokenID string, chainID *big.Int) (BridgeConfigV3Token, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getToken", tokenID, chainID)

	if err != nil {
		return *new(BridgeConfigV3Token), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeConfigV3Token)).(*BridgeConfigV3Token)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x324980b5.
//
// Solidity: function getToken(string tokenID, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetToken(tokenID string, chainID *big.Int) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetToken(&_BridgeConfigV3.CallOpts, tokenID, chainID)
}

// GetToken is a free data retrieval call binding the contract method 0x324980b5.
//
// Solidity: function getToken(string tokenID, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetToken(tokenID string, chainID *big.Int) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetToken(&_BridgeConfigV3.CallOpts, tokenID, chainID)
}

// GetTokenByAddress is a free data retrieval call binding the contract method 0xe814157d.
//
// Solidity: function getTokenByAddress(string tokenAddress, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetTokenByAddress(opts *bind.CallOpts, tokenAddress string, chainID *big.Int) (BridgeConfigV3Token, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getTokenByAddress", tokenAddress, chainID)

	if err != nil {
		return *new(BridgeConfigV3Token), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeConfigV3Token)).(*BridgeConfigV3Token)

	return out0, err

}

// GetTokenByAddress is a free data retrieval call binding the contract method 0xe814157d.
//
// Solidity: function getTokenByAddress(string tokenAddress, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetTokenByAddress(tokenAddress string, chainID *big.Int) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetTokenByAddress(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetTokenByAddress is a free data retrieval call binding the contract method 0xe814157d.
//
// Solidity: function getTokenByAddress(string tokenAddress, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetTokenByAddress(tokenAddress string, chainID *big.Int) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetTokenByAddress(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetTokenByEVMAddress is a free data retrieval call binding the contract method 0x558dae3a.
//
// Solidity: function getTokenByEVMAddress(address tokenAddress, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetTokenByEVMAddress(opts *bind.CallOpts, tokenAddress common.Address, chainID *big.Int) (BridgeConfigV3Token, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getTokenByEVMAddress", tokenAddress, chainID)

	if err != nil {
		return *new(BridgeConfigV3Token), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeConfigV3Token)).(*BridgeConfigV3Token)

	return out0, err

}

// GetTokenByEVMAddress is a free data retrieval call binding the contract method 0x558dae3a.
//
// Solidity: function getTokenByEVMAddress(address tokenAddress, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetTokenByEVMAddress(tokenAddress common.Address, chainID *big.Int) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetTokenByEVMAddress(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetTokenByEVMAddress is a free data retrieval call binding the contract method 0x558dae3a.
//
// Solidity: function getTokenByEVMAddress(address tokenAddress, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetTokenByEVMAddress(tokenAddress common.Address, chainID *big.Int) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetTokenByEVMAddress(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetTokenByID is a free data retrieval call binding the contract method 0x77b8cbf7.
//
// Solidity: function getTokenByID(string tokenID, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetTokenByID(opts *bind.CallOpts, tokenID string, chainID *big.Int) (BridgeConfigV3Token, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getTokenByID", tokenID, chainID)

	if err != nil {
		return *new(BridgeConfigV3Token), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeConfigV3Token)).(*BridgeConfigV3Token)

	return out0, err

}

// GetTokenByID is a free data retrieval call binding the contract method 0x77b8cbf7.
//
// Solidity: function getTokenByID(string tokenID, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetTokenByID(tokenID string, chainID *big.Int) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetTokenByID(&_BridgeConfigV3.CallOpts, tokenID, chainID)
}

// GetTokenByID is a free data retrieval call binding the contract method 0x77b8cbf7.
//
// Solidity: function getTokenByID(string tokenID, uint256 chainID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetTokenByID(tokenID string, chainID *big.Int) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetTokenByID(&_BridgeConfigV3.CallOpts, tokenID, chainID)
}

// GetTokenID is a free data retrieval call binding the contract method 0x3cc1c7e0.
//
// Solidity: function getTokenID(address tokenAddress, uint256 chainID) view returns(string)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetTokenID(opts *bind.CallOpts, tokenAddress common.Address, chainID *big.Int) (string, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getTokenID", tokenAddress, chainID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenID is a free data retrieval call binding the contract method 0x3cc1c7e0.
//
// Solidity: function getTokenID(address tokenAddress, uint256 chainID) view returns(string)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetTokenID(tokenAddress common.Address, chainID *big.Int) (string, error) {
	return _BridgeConfigV3.Contract.GetTokenID(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetTokenID is a free data retrieval call binding the contract method 0x3cc1c7e0.
//
// Solidity: function getTokenID(address tokenAddress, uint256 chainID) view returns(string)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetTokenID(tokenAddress common.Address, chainID *big.Int) (string, error) {
	return _BridgeConfigV3.Contract.GetTokenID(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetTokenID0 is a free data retrieval call binding the contract method 0xefd7516e.
//
// Solidity: function getTokenID(string tokenAddress, uint256 chainID) view returns(string)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetTokenID0(opts *bind.CallOpts, tokenAddress string, chainID *big.Int) (string, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getTokenID0", tokenAddress, chainID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenID0 is a free data retrieval call binding the contract method 0xefd7516e.
//
// Solidity: function getTokenID(string tokenAddress, uint256 chainID) view returns(string)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetTokenID0(tokenAddress string, chainID *big.Int) (string, error) {
	return _BridgeConfigV3.Contract.GetTokenID0(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetTokenID0 is a free data retrieval call binding the contract method 0xefd7516e.
//
// Solidity: function getTokenID(string tokenAddress, uint256 chainID) view returns(string)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetTokenID0(tokenAddress string, chainID *big.Int) (string, error) {
	return _BridgeConfigV3.Contract.GetTokenID0(&_BridgeConfigV3.CallOpts, tokenAddress, chainID)
}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0x58dfe6f1.
//
// Solidity: function getUnderlyingToken(string tokenID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Caller) GetUnderlyingToken(opts *bind.CallOpts, tokenID string) (BridgeConfigV3Token, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "getUnderlyingToken", tokenID)

	if err != nil {
		return *new(BridgeConfigV3Token), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeConfigV3Token)).(*BridgeConfigV3Token)

	return out0, err

}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0x58dfe6f1.
//
// Solidity: function getUnderlyingToken(string tokenID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3Session) GetUnderlyingToken(tokenID string) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetUnderlyingToken(&_BridgeConfigV3.CallOpts, tokenID)
}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0x58dfe6f1.
//
// Solidity: function getUnderlyingToken(string tokenID) view returns((uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool) token)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) GetUnderlyingToken(tokenID string) (BridgeConfigV3Token, error) {
	return _BridgeConfigV3.Contract.GetUnderlyingToken(&_BridgeConfigV3.CallOpts, tokenID)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeConfigV3.Contract.HasRole(&_BridgeConfigV3.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeConfigV3.Contract.HasRole(&_BridgeConfigV3.CallOpts, role, account)
}

// HasUnderlyingToken is a free data retrieval call binding the contract method 0x074b7e97.
//
// Solidity: function hasUnderlyingToken(string tokenID) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Caller) HasUnderlyingToken(opts *bind.CallOpts, tokenID string) (bool, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "hasUnderlyingToken", tokenID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUnderlyingToken is a free data retrieval call binding the contract method 0x074b7e97.
//
// Solidity: function hasUnderlyingToken(string tokenID) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Session) HasUnderlyingToken(tokenID string) (bool, error) {
	return _BridgeConfigV3.Contract.HasUnderlyingToken(&_BridgeConfigV3.CallOpts, tokenID)
}

// HasUnderlyingToken is a free data retrieval call binding the contract method 0x074b7e97.
//
// Solidity: function hasUnderlyingToken(string tokenID) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) HasUnderlyingToken(tokenID string) (bool, error) {
	return _BridgeConfigV3.Contract.HasUnderlyingToken(&_BridgeConfigV3.CallOpts, tokenID)
}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Caller) IsTokenIDExist(opts *bind.CallOpts, tokenID string) (bool, error) {
	var out []interface{}
	err := _BridgeConfigV3.contract.Call(opts, &out, "isTokenIDExist", tokenID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Session) IsTokenIDExist(tokenID string) (bool, error) {
	return _BridgeConfigV3.Contract.IsTokenIDExist(&_BridgeConfigV3.CallOpts, tokenID)
}

// IsTokenIDExist is a free data retrieval call binding the contract method 0xaf611ca0.
//
// Solidity: function isTokenIDExist(string tokenID) view returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3CallerSession) IsTokenIDExist(tokenID string) (bool, error) {
	return _BridgeConfigV3.Contract.IsTokenIDExist(&_BridgeConfigV3.CallOpts, tokenID)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.GrantRole(&_BridgeConfigV3.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.GrantRole(&_BridgeConfigV3.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.RenounceRole(&_BridgeConfigV3.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.RenounceRole(&_BridgeConfigV3.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.RevokeRole(&_BridgeConfigV3.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeConfigV3 *BridgeConfigV3TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.RevokeRole(&_BridgeConfigV3.TransactOpts, role, account)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xabaac008.
//
// Solidity: function setMaxGasPrice(uint256 chainID, uint256 maxPrice) returns()
func (_BridgeConfigV3 *BridgeConfigV3Transactor) SetMaxGasPrice(opts *bind.TransactOpts, chainID *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BridgeConfigV3.contract.Transact(opts, "setMaxGasPrice", chainID, maxPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xabaac008.
//
// Solidity: function setMaxGasPrice(uint256 chainID, uint256 maxPrice) returns()
func (_BridgeConfigV3 *BridgeConfigV3Session) SetMaxGasPrice(chainID *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.SetMaxGasPrice(&_BridgeConfigV3.TransactOpts, chainID, maxPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xabaac008.
//
// Solidity: function setMaxGasPrice(uint256 chainID, uint256 maxPrice) returns()
func (_BridgeConfigV3 *BridgeConfigV3TransactorSession) SetMaxGasPrice(chainID *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.SetMaxGasPrice(&_BridgeConfigV3.TransactOpts, chainID, maxPrice)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0x7e355e5e.
//
// Solidity: function setPoolConfig(address tokenAddress, uint256 chainID, address poolAddress, bool metaswap) returns((address,uint256,address,bool))
func (_BridgeConfigV3 *BridgeConfigV3Transactor) SetPoolConfig(opts *bind.TransactOpts, tokenAddress common.Address, chainID *big.Int, poolAddress common.Address, metaswap bool) (*types.Transaction, error) {
	return _BridgeConfigV3.contract.Transact(opts, "setPoolConfig", tokenAddress, chainID, poolAddress, metaswap)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0x7e355e5e.
//
// Solidity: function setPoolConfig(address tokenAddress, uint256 chainID, address poolAddress, bool metaswap) returns((address,uint256,address,bool))
func (_BridgeConfigV3 *BridgeConfigV3Session) SetPoolConfig(tokenAddress common.Address, chainID *big.Int, poolAddress common.Address, metaswap bool) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.SetPoolConfig(&_BridgeConfigV3.TransactOpts, tokenAddress, chainID, poolAddress, metaswap)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0x7e355e5e.
//
// Solidity: function setPoolConfig(address tokenAddress, uint256 chainID, address poolAddress, bool metaswap) returns((address,uint256,address,bool))
func (_BridgeConfigV3 *BridgeConfigV3TransactorSession) SetPoolConfig(tokenAddress common.Address, chainID *big.Int, poolAddress common.Address, metaswap bool) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.SetPoolConfig(&_BridgeConfigV3.TransactOpts, tokenAddress, chainID, poolAddress, metaswap)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x59053bfe.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, address tokenAddress, uint8 tokenDecimals, uint256 maxSwap, uint256 minSwap, uint256 swapFee, uint256 maxSwapFee, uint256 minSwapFee, bool hasUnderlying, bool isUnderlying) returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Transactor) SetTokenConfig(opts *bind.TransactOpts, tokenID string, chainID *big.Int, tokenAddress common.Address, tokenDecimals uint8, maxSwap *big.Int, minSwap *big.Int, swapFee *big.Int, maxSwapFee *big.Int, minSwapFee *big.Int, hasUnderlying bool, isUnderlying bool) (*types.Transaction, error) {
	return _BridgeConfigV3.contract.Transact(opts, "setTokenConfig", tokenID, chainID, tokenAddress, tokenDecimals, maxSwap, minSwap, swapFee, maxSwapFee, minSwapFee, hasUnderlying, isUnderlying)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x59053bfe.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, address tokenAddress, uint8 tokenDecimals, uint256 maxSwap, uint256 minSwap, uint256 swapFee, uint256 maxSwapFee, uint256 minSwapFee, bool hasUnderlying, bool isUnderlying) returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Session) SetTokenConfig(tokenID string, chainID *big.Int, tokenAddress common.Address, tokenDecimals uint8, maxSwap *big.Int, minSwap *big.Int, swapFee *big.Int, maxSwapFee *big.Int, minSwapFee *big.Int, hasUnderlying bool, isUnderlying bool) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.SetTokenConfig(&_BridgeConfigV3.TransactOpts, tokenID, chainID, tokenAddress, tokenDecimals, maxSwap, minSwap, swapFee, maxSwapFee, minSwapFee, hasUnderlying, isUnderlying)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x59053bfe.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, address tokenAddress, uint8 tokenDecimals, uint256 maxSwap, uint256 minSwap, uint256 swapFee, uint256 maxSwapFee, uint256 minSwapFee, bool hasUnderlying, bool isUnderlying) returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3TransactorSession) SetTokenConfig(tokenID string, chainID *big.Int, tokenAddress common.Address, tokenDecimals uint8, maxSwap *big.Int, minSwap *big.Int, swapFee *big.Int, maxSwapFee *big.Int, minSwapFee *big.Int, hasUnderlying bool, isUnderlying bool) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.SetTokenConfig(&_BridgeConfigV3.TransactOpts, tokenID, chainID, tokenAddress, tokenDecimals, maxSwap, minSwap, swapFee, maxSwapFee, minSwapFee, hasUnderlying, isUnderlying)
}

// SetTokenConfig0 is a paid mutator transaction binding the contract method 0xddb54399.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddress, uint8 tokenDecimals, uint256 maxSwap, uint256 minSwap, uint256 swapFee, uint256 maxSwapFee, uint256 minSwapFee, bool hasUnderlying, bool isUnderlying) returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Transactor) SetTokenConfig0(opts *bind.TransactOpts, tokenID string, chainID *big.Int, tokenAddress string, tokenDecimals uint8, maxSwap *big.Int, minSwap *big.Int, swapFee *big.Int, maxSwapFee *big.Int, minSwapFee *big.Int, hasUnderlying bool, isUnderlying bool) (*types.Transaction, error) {
	return _BridgeConfigV3.contract.Transact(opts, "setTokenConfig0", tokenID, chainID, tokenAddress, tokenDecimals, maxSwap, minSwap, swapFee, maxSwapFee, minSwapFee, hasUnderlying, isUnderlying)
}

// SetTokenConfig0 is a paid mutator transaction binding the contract method 0xddb54399.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddress, uint8 tokenDecimals, uint256 maxSwap, uint256 minSwap, uint256 swapFee, uint256 maxSwapFee, uint256 minSwapFee, bool hasUnderlying, bool isUnderlying) returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3Session) SetTokenConfig0(tokenID string, chainID *big.Int, tokenAddress string, tokenDecimals uint8, maxSwap *big.Int, minSwap *big.Int, swapFee *big.Int, maxSwapFee *big.Int, minSwapFee *big.Int, hasUnderlying bool, isUnderlying bool) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.SetTokenConfig0(&_BridgeConfigV3.TransactOpts, tokenID, chainID, tokenAddress, tokenDecimals, maxSwap, minSwap, swapFee, maxSwapFee, minSwapFee, hasUnderlying, isUnderlying)
}

// SetTokenConfig0 is a paid mutator transaction binding the contract method 0xddb54399.
//
// Solidity: function setTokenConfig(string tokenID, uint256 chainID, string tokenAddress, uint8 tokenDecimals, uint256 maxSwap, uint256 minSwap, uint256 swapFee, uint256 maxSwapFee, uint256 minSwapFee, bool hasUnderlying, bool isUnderlying) returns(bool)
func (_BridgeConfigV3 *BridgeConfigV3TransactorSession) SetTokenConfig0(tokenID string, chainID *big.Int, tokenAddress string, tokenDecimals uint8, maxSwap *big.Int, minSwap *big.Int, swapFee *big.Int, maxSwapFee *big.Int, minSwapFee *big.Int, hasUnderlying bool, isUnderlying bool) (*types.Transaction, error) {
	return _BridgeConfigV3.Contract.SetTokenConfig0(&_BridgeConfigV3.TransactOpts, tokenID, chainID, tokenAddress, tokenDecimals, maxSwap, minSwap, swapFee, maxSwapFee, minSwapFee, hasUnderlying, isUnderlying)
}

// BridgeConfigV3RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeConfigV3 contract.
type BridgeConfigV3RoleAdminChangedIterator struct {
	Event *BridgeConfigV3RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeConfigV3RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeConfigV3RoleAdminChanged)
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
		it.Event = new(BridgeConfigV3RoleAdminChanged)
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
func (it *BridgeConfigV3RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeConfigV3RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeConfigV3RoleAdminChanged represents a RoleAdminChanged event raised by the BridgeConfigV3 contract.
type BridgeConfigV3RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeConfigV3 *BridgeConfigV3Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeConfigV3RoleAdminChangedIterator, error) {

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

	logs, sub, err := _BridgeConfigV3.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigV3RoleAdminChangedIterator{contract: _BridgeConfigV3.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeConfigV3 *BridgeConfigV3Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeConfigV3RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BridgeConfigV3.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeConfigV3RoleAdminChanged)
				if err := _BridgeConfigV3.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BridgeConfigV3 *BridgeConfigV3Filterer) ParseRoleAdminChanged(log types.Log) (*BridgeConfigV3RoleAdminChanged, error) {
	event := new(BridgeConfigV3RoleAdminChanged)
	if err := _BridgeConfigV3.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeConfigV3RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeConfigV3 contract.
type BridgeConfigV3RoleGrantedIterator struct {
	Event *BridgeConfigV3RoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeConfigV3RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeConfigV3RoleGranted)
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
		it.Event = new(BridgeConfigV3RoleGranted)
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
func (it *BridgeConfigV3RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeConfigV3RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeConfigV3RoleGranted represents a RoleGranted event raised by the BridgeConfigV3 contract.
type BridgeConfigV3RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeConfigV3 *BridgeConfigV3Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeConfigV3RoleGrantedIterator, error) {

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

	logs, sub, err := _BridgeConfigV3.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigV3RoleGrantedIterator{contract: _BridgeConfigV3.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeConfigV3 *BridgeConfigV3Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeConfigV3RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeConfigV3.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeConfigV3RoleGranted)
				if err := _BridgeConfigV3.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BridgeConfigV3 *BridgeConfigV3Filterer) ParseRoleGranted(log types.Log) (*BridgeConfigV3RoleGranted, error) {
	event := new(BridgeConfigV3RoleGranted)
	if err := _BridgeConfigV3.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeConfigV3RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeConfigV3 contract.
type BridgeConfigV3RoleRevokedIterator struct {
	Event *BridgeConfigV3RoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeConfigV3RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeConfigV3RoleRevoked)
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
		it.Event = new(BridgeConfigV3RoleRevoked)
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
func (it *BridgeConfigV3RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeConfigV3RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeConfigV3RoleRevoked represents a RoleRevoked event raised by the BridgeConfigV3 contract.
type BridgeConfigV3RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeConfigV3 *BridgeConfigV3Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeConfigV3RoleRevokedIterator, error) {

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

	logs, sub, err := _BridgeConfigV3.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigV3RoleRevokedIterator{contract: _BridgeConfigV3.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeConfigV3 *BridgeConfigV3Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeConfigV3RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeConfigV3.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeConfigV3RoleRevoked)
				if err := _BridgeConfigV3.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BridgeConfigV3 *BridgeConfigV3Filterer) ParseRoleRevoked(log types.Log) (*BridgeConfigV3RoleRevoked, error) {
	event := new(BridgeConfigV3RoleRevoked)
	if err := _BridgeConfigV3.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220405e2324b796616fb7f971f7b3275378e9fe83959e1a656abdd6c0338eca2e6f64736f6c634300060c0033",
}

// EnumerableSetABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetMetaData.ABI instead.
var EnumerableSetABI = EnumerableSetMetaData.ABI

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetMetaData.Bin instead.
var EnumerableSetBin = EnumerableSetMetaData.Bin

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d4eddff4055697d351f5a3da789f263dfbf77e5fafc13033c5bee8f2fc527b9964736f6c634300060c0033",
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
