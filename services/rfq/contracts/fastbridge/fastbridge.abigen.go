// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fastbridge

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

// IFastBridgeBridgeParams is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeBridgeParams struct {
	DstChainId   uint32
	To           common.Address
	OriginToken  common.Address
	DestToken    common.Address
	OriginAmount *big.Int
	DestAmount   *big.Int
	SendChainGas bool
	Deadline     *big.Int
}

// IFastBridgeBridgeTransaction is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeBridgeTransaction struct {
	OriginChainId   uint32
	DestChainId     uint32
	OriginSender    common.Address
	DestRecipient   common.Address
	OriginToken     common.Address
	DestToken       common.Address
	OriginAmount    *big.Int
	DestAmount      *big.Int
	OriginFeeAmount *big.Int
	SendChainGas    bool
	Deadline        *big.Int
	Nonce           *big.Int
}

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
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
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, callerConfirmation)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202d6c62037e144f9ef289f9f324d4afaead691fba08bce487bb17ddeab9d9b53864736f6c63430008140033",
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
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// AdminMetaData contains all meta data concerning the Admin contract.
var AdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldChainGasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_RATE_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"addGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"addRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"removeGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"removeRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"bf333f2c": "FEE_BPS()",
		"0f5f6ed7": "FEE_RATE_MAX()",
		"ccc57490": "GOVERNOR_ROLE()",
		"03ed0ee5": "GUARD_ROLE()",
		"926d7d7f": "RELAYER_ROLE()",
		"3c4a25d0": "addGovernor(address)",
		"6913a63c": "addGuard(address)",
		"dd39f00d": "addRelayer(address)",
		"e00a83e0": "chainGasAmount()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"58f85880": "protocolFeeRate()",
		"dcf844a7": "protocolFees(address)",
		"eecdac88": "removeGovernor(address)",
		"b6235016": "removeGuard(address)",
		"60f0a5ac": "removeRelayer(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"b250fe6b": "setChainGasAmount(uint256)",
		"b13aa2d6": "setProtocolFeeRate(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"06f333f2": "sweepProtocolFees(address,address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051620015c2380380620015c2833981016040819052610031916100ef565b61003c600082610043565b505061011f565b6000828152602081815260408083206001600160a01b038516845290915281205460ff166100e5576000838152602081815260408083206001600160a01b03861684529091529020805460ff1916600117905561009d3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016100e9565b5060005b92915050565b60006020828403121561010157600080fd5b81516001600160a01b038116811461011857600080fd5b9392505050565b611493806200012f6000396000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c8063926d7d7f116100e3578063ccc574901161008c578063dd39f00d11610066578063dd39f00d146103b3578063e00a83e0146103c6578063eecdac88146103cf57600080fd5b8063ccc5749014610359578063d547741f14610380578063dcf844a71461039357600080fd5b8063b250fe6b116100bd578063b250fe6b14610329578063b62350161461033c578063bf333f2c1461034f57600080fd5b8063926d7d7f146102e7578063a217fddf1461030e578063b13aa2d61461031657600080fd5b806336568abe1161014557806360f0a5ac1161011f57806360f0a5ac1461027d5780636913a63c1461029057806391d14854146102a357600080fd5b806336568abe1461024e5780633c4a25d01461026157806358f858801461027457600080fd5b80630f5f6ed7116101765780630f5f6ed71461020f578063248a9ca3146102185780632f2ff15d1461023b57600080fd5b806301ffc9a71461019d57806303ed0ee5146101c557806306f333f2146101fa575b600080fd5b6101b06101ab366004611317565b6103e2565b60405190151581526020015b60405180910390f35b6101ec7f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d81565b6040519081526020016101bc565b61020d610208366004611382565b61047b565b005b6101ec61271081565b6101ec6102263660046113b5565b60009081526020819052604090206001015490565b61020d6102493660046113ce565b6105da565b61020d61025c3660046113ce565b610605565b61020d61026f3660046113f1565b610663565b6101ec60015481565b61020d61028b3660046113f1565b610716565b61020d61029e3660046113f1565b6107c2565b6101b06102b13660046113ce565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101ec7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc481565b6101ec600081565b61020d6103243660046113b5565b61086e565b61020d6103373660046113b5565b6109b8565b61020d61034a3660046113f1565b610a8e565b6101ec620f424081565b6101ec7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b61020d61038e3660046113ce565b610b3a565b6101ec6103a13660046113f1565b60026020526000908152604090205481565b61020d6103c13660046113f1565b610b5f565b6101ec60035481565b61020d6103dd3660046113f1565b610c0b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061047557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b3360009081527f6f3b488caa0dc03e5f3d9fd6a25fa8755afcfd62ccffb6e436c696bfa6b866de602052604090205460ff16610518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f43616c6c6572206973206e6f74206120676f7665726e6f72000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600260205260408120549081900361054b57505050565b73ffffffffffffffffffffffffffffffffffffffff831660008181526002602052604081205561057c908383610cb7565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd9060600160405180910390a1505b5050565b6000828152602081905260409020600101546105f581610e01565b6105ff8383610e0e565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610654576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61065e8282610f0a565b505050565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1661069e57600080fd5b6106c87f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5582610e0e565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5906020015b60405180910390a150565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1661075157600080fd5b61077b7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc482610f0a565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b9060200161070b565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff166107fd57600080fd5b6108277f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d82610e0e565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527f93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f599060200161070b565b3360009081527f6f3b488caa0dc03e5f3d9fd6a25fa8755afcfd62ccffb6e436c696bfa6b866de602052604090205460ff16610906576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f43616c6c6572206973206e6f74206120676f7665726e6f720000000000000000604482015260640161050f565b612710811115610972576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f6e657746656552617465203e206d617800000000000000000000000000000000604482015260640161050f565b600180549082905560408051828152602081018490527f14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb95791015b60405180910390a15050565b3360009081527f6f3b488caa0dc03e5f3d9fd6a25fa8755afcfd62ccffb6e436c696bfa6b866de602052604090205460ff16610a50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f43616c6c6572206973206e6f74206120676f7665726e6f720000000000000000604482015260640161050f565b600380549082905560408051828152602081018490527f5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa91016109ac565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16610ac957600080fd5b610af37f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d82610f0a565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527f59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d0489060200161070b565b600082815260208190526040902060010154610b5581610e01565b6105ff8383610f0a565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16610b9a57600080fd5b610bc47fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc482610e0e565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c59060200161070b565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16610c4657600080fd5b610c707f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5582610f0a565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b9060200161070b565b3073ffffffffffffffffffffffffffffffffffffffff831603610cd957505050565b7fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff841601610de05760008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114610d70576040519150601f19603f3d011682016040523d82523d6000602084013e610d75565b606091505b50509050806105ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c656400000000000000000000000000604482015260640161050f565b61065e73ffffffffffffffffffffffffffffffffffffffff84168383610fc5565b610e0b8133611052565b50565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610f025760008381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610ea03390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610475565b506000610475565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610f025760008381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610475565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261065e9084906110d8565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166105d6576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161050f565b60006110fa73ffffffffffffffffffffffffffffffffffffffff84168361116e565b9050805160001415801561111f57508080602001905181019061111d919061140c565b155b1561065e576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161050f565b606061117c83836000611183565b9392505050565b6060814710156111c1576040517fcd78605900000000000000000000000000000000000000000000000000000000815230600482015260240161050f565b6000808573ffffffffffffffffffffffffffffffffffffffff1684866040516111ea919061142e565b60006040518083038185875af1925050503d8060008114611227576040519150601f19603f3d011682016040523d82523d6000602084013e61122c565b606091505b509150915061123c868383611246565b9695505050505050565b60608261125b57611256826112d5565b61117c565b815115801561127f575073ffffffffffffffffffffffffffffffffffffffff84163b155b156112ce576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161050f565b508061117c565b8051156112e55780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561132957600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461117c57600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461137d57600080fd5b919050565b6000806040838503121561139557600080fd5b61139e83611359565b91506113ac60208401611359565b90509250929050565b6000602082840312156113c757600080fd5b5035919050565b600080604083850312156113e157600080fd5b823591506113ac60208401611359565b60006020828403121561140357600080fd5b61117c82611359565b60006020828403121561141e57600080fd5b8151801515811461117c57600080fd5b6000825160005b8181101561144f5760208186018101518583015201611435565b50600092019182525091905056fea264697066735822122078d3d5ae4bb368717cef758d16ba8fed6cfff0a2da35671340e16ff844d0591b64736f6c63430008140033",
}

// AdminABI is the input ABI used to generate the binding from.
// Deprecated: Use AdminMetaData.ABI instead.
var AdminABI = AdminMetaData.ABI

// Deprecated: Use AdminMetaData.Sigs instead.
// AdminFuncSigs maps the 4-byte function signature to its string representation.
var AdminFuncSigs = AdminMetaData.Sigs

// AdminBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdminMetaData.Bin instead.
var AdminBin = AdminMetaData.Bin

// DeployAdmin deploys a new Ethereum contract, binding an instance of Admin to it.
func DeployAdmin(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Admin, error) {
	parsed, err := AdminMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdminBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// Admin is an auto generated Go binding around an Ethereum contract.
type Admin struct {
	AdminCaller     // Read-only binding to the contract
	AdminTransactor // Write-only binding to the contract
	AdminFilterer   // Log filterer for contract events
}

// AdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminSession struct {
	Contract     *Admin            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminCallerSession struct {
	Contract *AdminCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminTransactorSession struct {
	Contract     *AdminTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminRaw struct {
	Contract *Admin // Generic contract binding to access the raw methods on
}

// AdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminCallerRaw struct {
	Contract *AdminCaller // Generic read-only contract binding to access the raw methods on
}

// AdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminTransactorRaw struct {
	Contract *AdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdmin creates a new instance of Admin, bound to a specific deployed contract.
func NewAdmin(address common.Address, backend bind.ContractBackend) (*Admin, error) {
	contract, err := bindAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// NewAdminCaller creates a new read-only instance of Admin, bound to a specific deployed contract.
func NewAdminCaller(address common.Address, caller bind.ContractCaller) (*AdminCaller, error) {
	contract, err := bindAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminCaller{contract: contract}, nil
}

// NewAdminTransactor creates a new write-only instance of Admin, bound to a specific deployed contract.
func NewAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminTransactor, error) {
	contract, err := bindAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminTransactor{contract: contract}, nil
}

// NewAdminFilterer creates a new log filterer instance of Admin, bound to a specific deployed contract.
func NewAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminFilterer, error) {
	contract, err := bindAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminFilterer{contract: contract}, nil
}

// bindAdmin binds a generic wrapper to an already deployed contract.
func bindAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.AdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Admin.Contract.DEFAULTADMINROLE(&_Admin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Admin.Contract.DEFAULTADMINROLE(&_Admin.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_Admin *AdminCaller) FEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_Admin *AdminSession) FEEBPS() (*big.Int, error) {
	return _Admin.Contract.FEEBPS(&_Admin.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_Admin *AdminCallerSession) FEEBPS() (*big.Int, error) {
	return _Admin.Contract.FEEBPS(&_Admin.CallOpts)
}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_Admin *AdminCaller) FEERATEMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "FEE_RATE_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_Admin *AdminSession) FEERATEMAX() (*big.Int, error) {
	return _Admin.Contract.FEERATEMAX(&_Admin.CallOpts)
}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_Admin *AdminCallerSession) FEERATEMAX() (*big.Int, error) {
	return _Admin.Contract.FEERATEMAX(&_Admin.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Admin *AdminCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Admin *AdminSession) GOVERNORROLE() ([32]byte, error) {
	return _Admin.Contract.GOVERNORROLE(&_Admin.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_Admin *AdminCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _Admin.Contract.GOVERNORROLE(&_Admin.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_Admin *AdminCaller) GUARDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "GUARD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_Admin *AdminSession) GUARDROLE() ([32]byte, error) {
	return _Admin.Contract.GUARDROLE(&_Admin.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_Admin *AdminCallerSession) GUARDROLE() ([32]byte, error) {
	return _Admin.Contract.GUARDROLE(&_Admin.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Admin *AdminCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Admin *AdminSession) RELAYERROLE() ([32]byte, error) {
	return _Admin.Contract.RELAYERROLE(&_Admin.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Admin *AdminCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Admin.Contract.RELAYERROLE(&_Admin.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_Admin *AdminCaller) ChainGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "chainGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_Admin *AdminSession) ChainGasAmount() (*big.Int, error) {
	return _Admin.Contract.ChainGasAmount(&_Admin.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_Admin *AdminCallerSession) ChainGasAmount() (*big.Int, error) {
	return _Admin.Contract.ChainGasAmount(&_Admin.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Admin *AdminCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Admin *AdminSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Admin.Contract.GetRoleAdmin(&_Admin.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Admin *AdminCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Admin.Contract.GetRoleAdmin(&_Admin.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Admin *AdminCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Admin *AdminSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Admin.Contract.HasRole(&_Admin.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Admin *AdminCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Admin.Contract.HasRole(&_Admin.CallOpts, role, account)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_Admin *AdminCaller) ProtocolFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "protocolFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_Admin *AdminSession) ProtocolFeeRate() (*big.Int, error) {
	return _Admin.Contract.ProtocolFeeRate(&_Admin.CallOpts)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_Admin *AdminCallerSession) ProtocolFeeRate() (*big.Int, error) {
	return _Admin.Contract.ProtocolFeeRate(&_Admin.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_Admin *AdminCaller) ProtocolFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "protocolFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_Admin *AdminSession) ProtocolFees(arg0 common.Address) (*big.Int, error) {
	return _Admin.Contract.ProtocolFees(&_Admin.CallOpts, arg0)
}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_Admin *AdminCallerSession) ProtocolFees(arg0 common.Address) (*big.Int, error) {
	return _Admin.Contract.ProtocolFees(&_Admin.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Admin *AdminCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Admin *AdminSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Admin.Contract.SupportsInterface(&_Admin.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Admin *AdminCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Admin.Contract.SupportsInterface(&_Admin.CallOpts, interfaceId)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_Admin *AdminTransactor) AddGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "addGovernor", _governor)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_Admin *AdminSession) AddGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddGovernor(&_Admin.TransactOpts, _governor)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_Admin *AdminTransactorSession) AddGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddGovernor(&_Admin.TransactOpts, _governor)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_Admin *AdminTransactor) AddGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "addGuard", _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_Admin *AdminSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddGuard(&_Admin.TransactOpts, _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_Admin *AdminTransactorSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddGuard(&_Admin.TransactOpts, _guard)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_Admin *AdminTransactor) AddRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "addRelayer", _relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_Admin *AdminSession) AddRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddRelayer(&_Admin.TransactOpts, _relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_Admin *AdminTransactorSession) AddRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddRelayer(&_Admin.TransactOpts, _relayer)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Admin *AdminTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Admin *AdminSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.Contract.GrantRole(&_Admin.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Admin *AdminTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.Contract.GrantRole(&_Admin.TransactOpts, role, account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_Admin *AdminTransactor) RemoveGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "removeGovernor", _governor)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_Admin *AdminSession) RemoveGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RemoveGovernor(&_Admin.TransactOpts, _governor)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_Admin *AdminTransactorSession) RemoveGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RemoveGovernor(&_Admin.TransactOpts, _governor)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_Admin *AdminTransactor) RemoveGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "removeGuard", _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_Admin *AdminSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RemoveGuard(&_Admin.TransactOpts, _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_Admin *AdminTransactorSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RemoveGuard(&_Admin.TransactOpts, _guard)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_Admin *AdminTransactor) RemoveRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "removeRelayer", _relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_Admin *AdminSession) RemoveRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RemoveRelayer(&_Admin.TransactOpts, _relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_Admin *AdminTransactorSession) RemoveRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RemoveRelayer(&_Admin.TransactOpts, _relayer)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Admin *AdminTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Admin *AdminSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RenounceRole(&_Admin.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Admin *AdminTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RenounceRole(&_Admin.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Admin *AdminTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Admin *AdminSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RevokeRole(&_Admin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Admin *AdminTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RevokeRole(&_Admin.TransactOpts, role, account)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_Admin *AdminTransactor) SetChainGasAmount(opts *bind.TransactOpts, newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "setChainGasAmount", newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_Admin *AdminSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.SetChainGasAmount(&_Admin.TransactOpts, newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_Admin *AdminTransactorSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.SetChainGasAmount(&_Admin.TransactOpts, newChainGasAmount)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_Admin *AdminTransactor) SetProtocolFeeRate(opts *bind.TransactOpts, newFeeRate *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "setProtocolFeeRate", newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_Admin *AdminSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.SetProtocolFeeRate(&_Admin.TransactOpts, newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_Admin *AdminTransactorSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.SetProtocolFeeRate(&_Admin.TransactOpts, newFeeRate)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_Admin *AdminTransactor) SweepProtocolFees(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "sweepProtocolFees", token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_Admin *AdminSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Admin.Contract.SweepProtocolFees(&_Admin.TransactOpts, token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_Admin *AdminTransactorSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Admin.Contract.SweepProtocolFees(&_Admin.TransactOpts, token, recipient)
}

// AdminChainGasAmountUpdatedIterator is returned from FilterChainGasAmountUpdated and is used to iterate over the raw logs and unpacked data for ChainGasAmountUpdated events raised by the Admin contract.
type AdminChainGasAmountUpdatedIterator struct {
	Event *AdminChainGasAmountUpdated // Event containing the contract specifics and raw log

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
func (it *AdminChainGasAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminChainGasAmountUpdated)
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
		it.Event = new(AdminChainGasAmountUpdated)
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
func (it *AdminChainGasAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminChainGasAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminChainGasAmountUpdated represents a ChainGasAmountUpdated event raised by the Admin contract.
type AdminChainGasAmountUpdated struct {
	OldChainGasAmount *big.Int
	NewChainGasAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChainGasAmountUpdated is a free log retrieval operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_Admin *AdminFilterer) FilterChainGasAmountUpdated(opts *bind.FilterOpts) (*AdminChainGasAmountUpdatedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &AdminChainGasAmountUpdatedIterator{contract: _Admin.contract, event: "ChainGasAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchChainGasAmountUpdated is a free log subscription operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_Admin *AdminFilterer) WatchChainGasAmountUpdated(opts *bind.WatchOpts, sink chan<- *AdminChainGasAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminChainGasAmountUpdated)
				if err := _Admin.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
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

// ParseChainGasAmountUpdated is a log parse operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_Admin *AdminFilterer) ParseChainGasAmountUpdated(log types.Log) (*AdminChainGasAmountUpdated, error) {
	event := new(AdminChainGasAmountUpdated)
	if err := _Admin.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminFeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the Admin contract.
type AdminFeeRateUpdatedIterator struct {
	Event *AdminFeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *AdminFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminFeeRateUpdated)
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
		it.Event = new(AdminFeeRateUpdated)
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
func (it *AdminFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminFeeRateUpdated represents a FeeRateUpdated event raised by the Admin contract.
type AdminFeeRateUpdated struct {
	OldFeeRate *big.Int
	NewFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_Admin *AdminFilterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*AdminFeeRateUpdatedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &AdminFeeRateUpdatedIterator{contract: _Admin.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_Admin *AdminFilterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *AdminFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminFeeRateUpdated)
				if err := _Admin.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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

// ParseFeeRateUpdated is a log parse operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_Admin *AdminFilterer) ParseFeeRateUpdated(log types.Log) (*AdminFeeRateUpdated, error) {
	event := new(AdminFeeRateUpdated)
	if err := _Admin.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminFeesSweptIterator is returned from FilterFeesSwept and is used to iterate over the raw logs and unpacked data for FeesSwept events raised by the Admin contract.
type AdminFeesSweptIterator struct {
	Event *AdminFeesSwept // Event containing the contract specifics and raw log

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
func (it *AdminFeesSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminFeesSwept)
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
		it.Event = new(AdminFeesSwept)
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
func (it *AdminFeesSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminFeesSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminFeesSwept represents a FeesSwept event raised by the Admin contract.
type AdminFeesSwept struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesSwept is a free log retrieval operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_Admin *AdminFilterer) FilterFeesSwept(opts *bind.FilterOpts) (*AdminFeesSweptIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return &AdminFeesSweptIterator{contract: _Admin.contract, event: "FeesSwept", logs: logs, sub: sub}, nil
}

// WatchFeesSwept is a free log subscription operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_Admin *AdminFilterer) WatchFeesSwept(opts *bind.WatchOpts, sink chan<- *AdminFeesSwept) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminFeesSwept)
				if err := _Admin.contract.UnpackLog(event, "FeesSwept", log); err != nil {
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

// ParseFeesSwept is a log parse operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_Admin *AdminFilterer) ParseFeesSwept(log types.Log) (*AdminFeesSwept, error) {
	event := new(AdminFeesSwept)
	if err := _Admin.contract.UnpackLog(event, "FeesSwept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the Admin contract.
type AdminGovernorAddedIterator struct {
	Event *AdminGovernorAdded // Event containing the contract specifics and raw log

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
func (it *AdminGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminGovernorAdded)
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
		it.Event = new(AdminGovernorAdded)
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
func (it *AdminGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminGovernorAdded represents a GovernorAdded event raised by the Admin contract.
type AdminGovernorAdded struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_Admin *AdminFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*AdminGovernorAddedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &AdminGovernorAddedIterator{contract: _Admin.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_Admin *AdminFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *AdminGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminGovernorAdded)
				if err := _Admin.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_Admin *AdminFilterer) ParseGovernorAdded(log types.Log) (*AdminGovernorAdded, error) {
	event := new(AdminGovernorAdded)
	if err := _Admin.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the Admin contract.
type AdminGovernorRemovedIterator struct {
	Event *AdminGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *AdminGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminGovernorRemoved)
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
		it.Event = new(AdminGovernorRemoved)
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
func (it *AdminGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminGovernorRemoved represents a GovernorRemoved event raised by the Admin contract.
type AdminGovernorRemoved struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_Admin *AdminFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*AdminGovernorRemovedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &AdminGovernorRemovedIterator{contract: _Admin.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_Admin *AdminFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *AdminGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminGovernorRemoved)
				if err := _Admin.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_Admin *AdminFilterer) ParseGovernorRemoved(log types.Log) (*AdminGovernorRemoved, error) {
	event := new(AdminGovernorRemoved)
	if err := _Admin.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the Admin contract.
type AdminGuardAddedIterator struct {
	Event *AdminGuardAdded // Event containing the contract specifics and raw log

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
func (it *AdminGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminGuardAdded)
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
		it.Event = new(AdminGuardAdded)
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
func (it *AdminGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminGuardAdded represents a GuardAdded event raised by the Admin contract.
type AdminGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_Admin *AdminFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*AdminGuardAddedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &AdminGuardAddedIterator{contract: _Admin.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_Admin *AdminFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *AdminGuardAdded) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminGuardAdded)
				if err := _Admin.contract.UnpackLog(event, "GuardAdded", log); err != nil {
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

// ParseGuardAdded is a log parse operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_Admin *AdminFilterer) ParseGuardAdded(log types.Log) (*AdminGuardAdded, error) {
	event := new(AdminGuardAdded)
	if err := _Admin.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the Admin contract.
type AdminGuardRemovedIterator struct {
	Event *AdminGuardRemoved // Event containing the contract specifics and raw log

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
func (it *AdminGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminGuardRemoved)
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
		it.Event = new(AdminGuardRemoved)
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
func (it *AdminGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminGuardRemoved represents a GuardRemoved event raised by the Admin contract.
type AdminGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_Admin *AdminFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*AdminGuardRemovedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &AdminGuardRemovedIterator{contract: _Admin.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_Admin *AdminFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *AdminGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminGuardRemoved)
				if err := _Admin.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
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

// ParseGuardRemoved is a log parse operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_Admin *AdminFilterer) ParseGuardRemoved(log types.Log) (*AdminGuardRemoved, error) {
	event := new(AdminGuardRemoved)
	if err := _Admin.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the Admin contract.
type AdminRelayerAddedIterator struct {
	Event *AdminRelayerAdded // Event containing the contract specifics and raw log

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
func (it *AdminRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminRelayerAdded)
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
		it.Event = new(AdminRelayerAdded)
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
func (it *AdminRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminRelayerAdded represents a RelayerAdded event raised by the Admin contract.
type AdminRelayerAdded struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_Admin *AdminFilterer) FilterRelayerAdded(opts *bind.FilterOpts) (*AdminRelayerAddedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "RelayerAdded")
	if err != nil {
		return nil, err
	}
	return &AdminRelayerAddedIterator{contract: _Admin.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_Admin *AdminFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *AdminRelayerAdded) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "RelayerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminRelayerAdded)
				if err := _Admin.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
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

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_Admin *AdminFilterer) ParseRelayerAdded(log types.Log) (*AdminRelayerAdded, error) {
	event := new(AdminRelayerAdded)
	if err := _Admin.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the Admin contract.
type AdminRelayerRemovedIterator struct {
	Event *AdminRelayerRemoved // Event containing the contract specifics and raw log

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
func (it *AdminRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminRelayerRemoved)
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
		it.Event = new(AdminRelayerRemoved)
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
func (it *AdminRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminRelayerRemoved represents a RelayerRemoved event raised by the Admin contract.
type AdminRelayerRemoved struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_Admin *AdminFilterer) FilterRelayerRemoved(opts *bind.FilterOpts) (*AdminRelayerRemovedIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "RelayerRemoved")
	if err != nil {
		return nil, err
	}
	return &AdminRelayerRemovedIterator{contract: _Admin.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_Admin *AdminFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *AdminRelayerRemoved) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "RelayerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminRelayerRemoved)
				if err := _Admin.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
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

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_Admin *AdminFilterer) ParseRelayerRemoved(log types.Log) (*AdminRelayerRemoved, error) {
	event := new(AdminRelayerRemoved)
	if err := _Admin.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Admin contract.
type AdminRoleAdminChangedIterator struct {
	Event *AdminRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AdminRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminRoleAdminChanged)
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
		it.Event = new(AdminRoleAdminChanged)
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
func (it *AdminRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminRoleAdminChanged represents a RoleAdminChanged event raised by the Admin contract.
type AdminRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Admin *AdminFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AdminRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Admin.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AdminRoleAdminChangedIterator{contract: _Admin.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Admin *AdminFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AdminRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Admin.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminRoleAdminChanged)
				if err := _Admin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Admin *AdminFilterer) ParseRoleAdminChanged(log types.Log) (*AdminRoleAdminChanged, error) {
	event := new(AdminRoleAdminChanged)
	if err := _Admin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Admin contract.
type AdminRoleGrantedIterator struct {
	Event *AdminRoleGranted // Event containing the contract specifics and raw log

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
func (it *AdminRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminRoleGranted)
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
		it.Event = new(AdminRoleGranted)
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
func (it *AdminRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminRoleGranted represents a RoleGranted event raised by the Admin contract.
type AdminRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdminRoleGrantedIterator, error) {

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

	logs, sub, err := _Admin.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdminRoleGrantedIterator{contract: _Admin.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AdminRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Admin.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminRoleGranted)
				if err := _Admin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Admin *AdminFilterer) ParseRoleGranted(log types.Log) (*AdminRoleGranted, error) {
	event := new(AdminRoleGranted)
	if err := _Admin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Admin contract.
type AdminRoleRevokedIterator struct {
	Event *AdminRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AdminRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminRoleRevoked)
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
		it.Event = new(AdminRoleRevoked)
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
func (it *AdminRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminRoleRevoked represents a RoleRevoked event raised by the Admin contract.
type AdminRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdminRoleRevokedIterator, error) {

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

	logs, sub, err := _Admin.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdminRoleRevokedIterator{contract: _Admin.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AdminRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Admin.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminRoleRevoked)
				if err := _Admin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Admin *AdminFilterer) ParseRoleRevoked(log types.Log) (*AdminRoleRevoked, error) {
	event := new(AdminRoleRevoked)
	if err := _Admin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// ERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165MetaData.ABI instead.
var ERC165ABI = ERC165MetaData.ABI

// Deprecated: Use ERC165MetaData.Sigs instead.
// ERC165FuncSigs maps the 4-byte function signature to its string representation.
var ERC165FuncSigs = ERC165MetaData.Sigs

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// FastBridgeMetaData contains all meta data concerning the FastBridge contract.
var FastBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineNotExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueIncorrect\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StatusIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldChainGasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_RATE_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEADLINE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"addGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"addRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeProofs\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"timestamp\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelays\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeStatuses\",\"outputs\":[{\"internalType\":\"enumFastBridge.BridgeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"removeGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"removeRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"a5bbe22b": "DISPUTE_PERIOD()",
		"bf333f2c": "FEE_BPS()",
		"0f5f6ed7": "FEE_RATE_MAX()",
		"ccc57490": "GOVERNOR_ROLE()",
		"03ed0ee5": "GUARD_ROLE()",
		"820688d5": "MIN_DEADLINE_PERIOD()",
		"926d7d7f": "RELAYER_ROLE()",
		"3c4a25d0": "addGovernor(address)",
		"6913a63c": "addGuard(address)",
		"dd39f00d": "addRelayer(address)",
		"778d31ff": "bridge((uint32,address,address,address,uint256,uint256,bool,uint256))",
		"91ad5039": "bridgeProofs(bytes32)",
		"8379a24f": "bridgeRelays(bytes32)",
		"051287bc": "bridgeStatuses(bytes32)",
		"aa9641ab": "canClaim(bytes32,address)",
		"e00a83e0": "chainGasAmount()",
		"41fcb612": "claim(bytes,address)",
		"a3ec191a": "deployBlock()",
		"add98c70": "dispute(bytes32)",
		"ac11fb1a": "getBridgeTransaction(bytes)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"affed0e0": "nonce()",
		"58f85880": "protocolFeeRate()",
		"dcf844a7": "protocolFees(address)",
		"886d36ff": "prove(bytes,bytes32)",
		"ef059bd3": "refund(bytes,address)",
		"8f0d6f17": "relay(bytes)",
		"eecdac88": "removeGovernor(address)",
		"b6235016": "removeGuard(address)",
		"60f0a5ac": "removeRelayer(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"b250fe6b": "setChainGasAmount(uint256)",
		"b13aa2d6": "setProtocolFeeRate(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"06f333f2": "sweepProtocolFees(address,address)",
	},
	Bin: "0x60a06040523480156200001157600080fd5b5060405162003047380380620030478339810160408190526200003491620000fe565b80620000426000826200004f565b5050436080525062000130565b6000828152602081815260408083206001600160a01b038516845290915281205460ff16620000f4576000838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055620000ab3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001620000f8565b5060005b92915050565b6000602082840312156200011157600080fd5b81516001600160a01b03811681146200012957600080fd5b9392505050565b608051612efb6200014c60003960006106440152612efb6000f3fe6080604052600436106102a05760003560e01c806391d148541161016e578063b250fe6b116100cb578063dcf844a71161007f578063e00a83e011610064578063e00a83e014610801578063eecdac8814610817578063ef059bd31461083757600080fd5b8063dcf844a7146107b4578063dd39f00d146107e157600080fd5b8063bf333f2c116100b0578063bf333f2c14610749578063ccc5749014610760578063d547741f1461079457600080fd5b8063b250fe6b14610709578063b62350161461072957600080fd5b8063aa9641ab11610122578063add98c7011610107578063add98c70146106b3578063affed0e0146106d3578063b13aa2d6146106e957600080fd5b8063aa9641ab14610666578063ac11fb1a1461068657600080fd5b8063a217fddf11610153578063a217fddf1461061d578063a3ec191a14610632578063a5bbe22b146104aa57600080fd5b806391d14854146105a5578063926d7d7f146105e957600080fd5b806341fcb6121161021c578063820688d5116101d0578063886d36ff116101b5578063886d36ff146104f05780638f0d6f171461051057806391ad50391461052357600080fd5b8063820688d5146104aa5780638379a24f146104c057600080fd5b806360f0a5ac1161020157806360f0a5ac146104575780636913a63c14610477578063778d31ff1461049757600080fd5b806341fcb6121461042157806358f858801461044157600080fd5b80630f5f6ed7116102735780632f2ff15d116102585780632f2ff15d146103c157806336568abe146103e15780633c4a25d01461040157600080fd5b80630f5f6ed71461037b578063248a9ca31461039157600080fd5b806301ffc9a7146102a557806303ed0ee5146102da578063051287bc1461031c57806306f333f214610359575b600080fd5b3480156102b157600080fd5b506102c56102c036600461271a565b610857565b60405190151581526020015b60405180910390f35b3480156102e657600080fd5b5061030e7f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d81565b6040519081526020016102d1565b34801561032857600080fd5b5061034c61033736600461275c565b60046020526000908152604090205460ff1681565b6040516102d191906127a4565b34801561036557600080fd5b5061037961037436600461280a565b6108f0565b005b34801561038757600080fd5b5061030e61271081565b34801561039d57600080fd5b5061030e6103ac36600461275c565b60009081526020819052604090206001015490565b3480156103cd57600080fd5b506103796103dc366004612843565b610a28565b3480156103ed57600080fd5b506103796103fc366004612843565b610a53565b34801561040d57600080fd5b5061037961041c366004612868565b610aa4565b34801561042d57600080fd5b5061037961043c366004612989565b610b4a565b34801561044d57600080fd5b5061030e60015481565b34801561046357600080fd5b50610379610472366004612868565b610e03565b34801561048357600080fd5b50610379610492366004612868565b610ea2565b6103796104a53660046129fb565b610f41565b3480156104b657600080fd5b5061030e61070881565b3480156104cc57600080fd5b506102c56104db36600461275c565b60066020526000908152604090205460ff1681565b3480156104fc57600080fd5b5061037961050b366004612ab3565b61121c565b61037961051e366004612af8565b611417565b34801561052f57600080fd5b5061057961053e36600461275c565b6005602052600090815260409020546bffffffffffffffffffffffff8116906c0100000000000000000000000090046001600160a01b031682565b604080516bffffffffffffffffffffffff90931683526001600160a01b039091166020830152016102d1565b3480156105b157600080fd5b506102c56105c0366004612843565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156105f557600080fd5b5061030e7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc481565b34801561062957600080fd5b5061030e600081565b34801561063e57600080fd5b5061030e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561067257600080fd5b506102c5610681366004612843565b61166e565b34801561069257600080fd5b506106a66106a1366004612af8565b611773565b6040516102d19190612b35565b3480156106bf57600080fd5b506103796106ce36600461275c565b6117e6565b3480156106df57600080fd5b5061030e60075481565b3480156106f557600080fd5b5061037961070436600461275c565b6119c7565b34801561071557600080fd5b5061037961072436600461275c565b611b11565b34801561073557600080fd5b50610379610744366004612868565b611be7565b34801561075557600080fd5b5061030e620f424081565b34801561076c57600080fd5b5061030e7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b3480156107a057600080fd5b506103796107af366004612843565b611c86565b3480156107c057600080fd5b5061030e6107cf366004612868565b60026020526000908152604090205481565b3480156107ed57600080fd5b506103796107fc366004612868565b611cab565b34801561080d57600080fd5b5061030e60035481565b34801561082357600080fd5b50610379610832366004612868565b611d4a565b34801561084357600080fd5b50610379610852366004612989565b611de9565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108ea57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b3360009081527f6f3b488caa0dc03e5f3d9fd6a25fa8755afcfd62ccffb6e436c696bfa6b866de602052604090205460ff1661098d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f43616c6c6572206973206e6f74206120676f7665726e6f72000000000000000060448201526064015b60405180910390fd5b6001600160a01b038216600090815260026020526040812054908190036109b357505050565b6001600160a01b0383166000818152600260205260408120556109d7908383611f86565b604080516001600160a01b038086168252841660208201529081018290527f244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd9060600160405180910390a1505b5050565b600082815260208190526040902060010154610a438161209c565b610a4d83836120a9565b50505050565b6001600160a01b0381163314610a95576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a9f8282612153565b505050565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16610adf57600080fd5b610b097f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55826120a9565b506040516001600160a01b03821681527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5906020015b60405180910390a150565b3360009081527ffaf93c3d007e112089dc8351e013e6685ef67703975d0224b26fc45941d4f1f5602052604090205460ff16610be2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f7420612072656c617965720000000000000000006044820152606401610984565b815160208301206000610bf484611773565b9050600260008381526004602081905260409091205460ff1690811115610c1d57610c1d612775565b14610c54576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152600560209081526040918290208251808401909352546bffffffffffffffffffffffff811683526c0100000000000000000000000090046001600160a01b03169082018190523314610cd7576040517f4af43a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516107089042036bffffffffffffffffffffffff1611610d24576040517f1992d0bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838152600460205260409020805460ff1916600317905561010082015115610d805761010082015160808301516001600160a01b031660009081526002602052604081208054909190610d7a908490612c4a565b90915550505b608082015160c0830151610d9e6001600160a01b0383168783611f86565b604080518681523360208201526001600160a01b03808916928201929092529083166060820152608081018290527f582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac6789060a0015b60405180910390a150505050505050565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16610e3e57600080fd5b610e687fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc482612153565b506040516001600160a01b03821681527f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b90602001610b3f565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16610edd57600080fd5b610f077f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d826120a9565b506040516001600160a01b03821681527f93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f5990602001610b3f565b46816000015163ffffffff1603610f84576040517f7029fdf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60808101511580610f97575060a0810151155b15610fce576040517fe38820c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408101516001600160a01b03161580610ff3575060608101516001600160a01b0316155b1561102a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61103661070842612c4a565b8160e001511015611073576040517f04b7fcc800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061108830836040015184608001516121d6565b905060008060015411156110b557620f4240600154836110a89190612c5d565b6110b29190612c74565b90505b6110bf8183612caf565b915060006040518061018001604052804663ffffffff168152602001856000015163ffffffff168152602001336001600160a01b0316815260200185602001516001600160a01b0316815260200185604001516001600160a01b0316815260200185606001516001600160a01b031681526020018481526020018560a0015181526020018381526020018560c00151151581526020018560e0015181526020016007600081548092919061117290612cc2565b9091555090526040516111889190602001612b35565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152828252805160208083019190912060008181526004909252929020805460ff191660011790559250907f2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf9061120d90839033908690612d1e565b60405180910390a15050505050565b3360009081527ffaf93c3d007e112089dc8351e013e6685ef67703975d0224b26fc45941d4f1f5602052604090205460ff166112b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f7420612072656c617965720000000000000000006044820152606401610984565b8151602083012060006112c684611773565b9050806101400151421115611307576040517f559895a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160008381526004602081905260409091205460ff169081111561132e5761132e612775565b14611365576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152600460209081526040808320805460ff19166002179055805180820182526bffffffffffffffffffffffff4281168252338285018181528887526005865295849020925195516001600160a01b03166c0100000000000000000000000002959091169490941790558051858152918201929092529081018490527f4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e49060600160405180910390a150505050565b3360009081527ffaf93c3d007e112089dc8351e013e6685ef67703975d0224b26fc45941d4f1f5602052604090205460ff166114af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f7420612072656c617965720000000000000000006044820152606401610984565b8051602082012060006114c183611773565b90504663ffffffff16816020015163ffffffff161461150c576040517f7029fdf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526006602052604090205460ff1615611555576040517fbef7bb7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152600660205260409020805460ff19166001179055606081015160a082015160e083015160035461012085015161159e575060006115988484846121d6565b5061160f565b7fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b038416016115e25761159884846115dd8486612c4a565b6121d6565b6115ed8484846121d6565b5061160d8473eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee836121d6565b505b604080518781523360208201526001600160a01b038087169282019290925290841660608201526080810183905260a081018290527ff5d0e166883753ed3b66d60c893dda5f6760044b7bba1317907a9385c6620f9b9060c001610df2565b6000600260008481526004602081905260409091205460ff169081111561169757611697612775565b146116ce576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838152600560209081526040918290208251808401909352546bffffffffffffffffffffffff811683526001600160a01b036c010000000000000000000000009091048116918301829052841614611754576040517f4af43a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516107089042036bffffffffffffffffffffffff1611949350505050565b604080516101808101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c0820181905260e08201819052610100820181905261012082018190526101408201819052610160820152825190916108ea9184018101908401612da7565b3360009081527f92bf2f06c618f6add65571bf52eb90d22c9b52fd65028daa87260b82998b30bd602052604090205460ff1661187e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f43616c6c6572206973206e6f74206120677561726400000000000000000000006044820152606401610984565b600260008281526004602081905260409091205460ff16908111156118a5576118a5612775565b146118dc576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000818152600560209081526040918290208251808401909352546bffffffffffffffffffffffff8082168085526c010000000000000000000000009092046001600160a01b0316939092019290925261070891420316111561196b576040517f3e908aac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000818152600460209081526040808320805460ff191660011790556005825280832092909255815183815233918101919091527f0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad49101610b3f565b3360009081527f6f3b488caa0dc03e5f3d9fd6a25fa8755afcfd62ccffb6e436c696bfa6b866de602052604090205460ff16611a5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f43616c6c6572206973206e6f74206120676f7665726e6f7200000000000000006044820152606401610984565b612710811115611acb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f6e657746656552617465203e206d6178000000000000000000000000000000006044820152606401610984565b600180549082905560408051828152602081018490527f14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb95791015b60405180910390a15050565b3360009081527f6f3b488caa0dc03e5f3d9fd6a25fa8755afcfd62ccffb6e436c696bfa6b866de602052604090205460ff16611ba9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f43616c6c6572206973206e6f74206120676f7665726e6f7200000000000000006044820152606401610984565b600380549082905560408051828152602081018490527f5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa9101611b05565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16611c2257600080fd5b611c4c7f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d82612153565b506040516001600160a01b03821681527f59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d04890602001610b3f565b600082815260208190526040902060010154611ca18161209c565b610a4d8383612153565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16611ce657600080fd5b611d107fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4826120a9565b506040516001600160a01b03821681527f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c590602001610b3f565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16611d8557600080fd5b611daf7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5582612153565b506040516001600160a01b03821681527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b90602001610b3f565b815160208301206000611dfb84611773565b60408101519091506001600160a01b03163314611e44576040517f4af43a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8061014001514211611e82576040517fe15ff9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160008381526004602081905260409091205460ff1690811115611ea957611ea9612775565b14611ee0576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526004602081905260408220805460ff19169091179055608082015161010083015160c0840151919291611f189190612c4a565b9050611f2e6001600160a01b0383168683611f86565b604080518581526001600160a01b0387811660208301528416818301526060810183905290517fb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc9589181900360800190a1505050505050565b306001600160a01b03831603611f9b57505050565b7fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03841601612088576000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114612018576040519150601f19603f3d011682016040523d82523d6000602084013e61201d565b606091505b5050905080610a4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610984565b610a9f6001600160a01b03841683836123a6565b6120a6813361241a565b50565b6000828152602081815260408083206001600160a01b038516845290915281205460ff1661214b576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556121033390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016108ea565b5060006108ea565b6000828152602081815260408083206001600160a01b038516845290915281205460ff161561214b576000838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016108ea565b60006001600160a01b03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1461233f5761220e836001600160a01b0316612486565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301528416906370a0823190602401602060405180830381865afa15801561226d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122919190612e73565b90506122a86001600160a01b03841633868561252c565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301528291908516906370a0823190602401602060405180830381865afa15801561230a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061232e9190612e73565b6123389190612caf565b905061239f565b348214612378576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038416301461239c5761239c6001600160a01b0384168584611f86565b50345b9392505050565b6040516001600160a01b03838116602483015260448201839052610a9f91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612565565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16610a24576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401610984565b7fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b038216016124e8576040517f7f523fe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b03163b6000036120a6576040517f7f523fe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516001600160a01b038481166024830152838116604483015260648201839052610a4d9186918216906323b872dd906084016123d3565b600061257a6001600160a01b038416836125e1565b9050805160001415801561259f57508080602001905181019061259d9190612e8c565b155b15610a9f576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610984565b606061239f8383600084600080856001600160a01b031684866040516126079190612ea9565b60006040518083038185875af1925050503d8060008114612644576040519150601f19603f3d011682016040523d82523d6000602084013e612649565b606091505b5091509150612659868383612663565b9695505050505050565b60608261267857612673826126d8565b61239f565b815115801561268f57506001600160a01b0384163b155b156126d1576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401610984565b508061239f565b8051156126e85780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561272c57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461239f57600080fd5b60006020828403121561276e57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208101600583106127df577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6001600160a01b03811681146120a657600080fd5b8035612805816127e5565b919050565b6000806040838503121561281d57600080fd5b8235612828816127e5565b91506020830135612838816127e5565b809150509250929050565b6000806040838503121561285657600080fd5b823591506020830135612838816127e5565b60006020828403121561287a57600080fd5b813561239f816127e5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610180810167ffffffffffffffff811182821017156128d8576128d8612885565b60405290565b600082601f8301126128ef57600080fd5b813567ffffffffffffffff8082111561290a5761290a612885565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561295057612950612885565b8160405283815286602085880101111561296957600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561299c57600080fd5b823567ffffffffffffffff8111156129b357600080fd5b6129bf858286016128de565b9250506020830135612838816127e5565b63ffffffff811681146120a657600080fd5b80151581146120a657600080fd5b8035612805816129e2565b6000610100808385031215612a0f57600080fd5b6040519081019067ffffffffffffffff82118183101715612a3257612a32612885565b8160405283359150612a43826129d0565b818152612a52602085016127fa565b6020820152612a63604085016127fa565b6040820152612a74606085016127fa565b60608201526080840135608082015260a084013560a0820152612a9960c085016129f0565b60c082015260e084013560e0820152809250505092915050565b60008060408385031215612ac657600080fd5b823567ffffffffffffffff811115612add57600080fd5b612ae9858286016128de565b95602094909401359450505050565b600060208284031215612b0a57600080fd5b813567ffffffffffffffff811115612b2157600080fd5b612b2d848285016128de565b949350505050565b815163ffffffff16815261018081016020830151612b5b602084018263ffffffff169052565b506040830151612b7660408401826001600160a01b03169052565b506060830151612b9160608401826001600160a01b03169052565b506080830151612bac60808401826001600160a01b03169052565b5060a0830151612bc760a08401826001600160a01b03169052565b5060c083015160c083015260e083015160e083015261010080840151818401525061012080840151612bfc8285018215159052565b5050610140838101519083015261016092830151929091019190915290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156108ea576108ea612c1b565b80820281158282048414176108ea576108ea612c1b565b600082612caa577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b818103818111156108ea576108ea612c1b565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612cf357612cf3612c1b565b5060010190565b60005b83811015612d15578181015183820152602001612cfd565b50506000910152565b8381526001600160a01b03831660208201526060604082015260008251806060840152612d52816080850160208701612cfa565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01691909101608001949350505050565b8051612805816129d0565b8051612805816127e5565b8051612805816129e2565b60006101808284031215612dba57600080fd5b612dc26128b4565b612dcb83612d86565b8152612dd960208401612d86565b6020820152612dea60408401612d91565b6040820152612dfb60608401612d91565b6060820152612e0c60808401612d91565b6080820152612e1d60a08401612d91565b60a082015260c083015160c082015260e083015160e0820152610100808401518183015250610120612e50818501612d9c565b908201526101408381015190820152610160928301519281019290925250919050565b600060208284031215612e8557600080fd5b5051919050565b600060208284031215612e9e57600080fd5b815161239f816129e2565b60008251612ebb818460208701612cfa565b919091019291505056fea2646970667358221220774806a328ca2212f7448aa1fc3c9c91c3c6fc680fd1ee8476e14b4e62b6d8fe64736f6c63430008140033",
}

// FastBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use FastBridgeMetaData.ABI instead.
var FastBridgeABI = FastBridgeMetaData.ABI

// Deprecated: Use FastBridgeMetaData.Sigs instead.
// FastBridgeFuncSigs maps the 4-byte function signature to its string representation.
var FastBridgeFuncSigs = FastBridgeMetaData.Sigs

// FastBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FastBridgeMetaData.Bin instead.
var FastBridgeBin = FastBridgeMetaData.Bin

// DeployFastBridge deploys a new Ethereum contract, binding an instance of FastBridge to it.
func DeployFastBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *FastBridge, error) {
	parsed, err := FastBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FastBridgeBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FastBridge{FastBridgeCaller: FastBridgeCaller{contract: contract}, FastBridgeTransactor: FastBridgeTransactor{contract: contract}, FastBridgeFilterer: FastBridgeFilterer{contract: contract}}, nil
}

// FastBridge is an auto generated Go binding around an Ethereum contract.
type FastBridge struct {
	FastBridgeCaller     // Read-only binding to the contract
	FastBridgeTransactor // Write-only binding to the contract
	FastBridgeFilterer   // Log filterer for contract events
}

// FastBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FastBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FastBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastBridgeSession struct {
	Contract     *FastBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FastBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastBridgeCallerSession struct {
	Contract *FastBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FastBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastBridgeTransactorSession struct {
	Contract     *FastBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FastBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FastBridgeRaw struct {
	Contract *FastBridge // Generic contract binding to access the raw methods on
}

// FastBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastBridgeCallerRaw struct {
	Contract *FastBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// FastBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastBridgeTransactorRaw struct {
	Contract *FastBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFastBridge creates a new instance of FastBridge, bound to a specific deployed contract.
func NewFastBridge(address common.Address, backend bind.ContractBackend) (*FastBridge, error) {
	contract, err := bindFastBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastBridge{FastBridgeCaller: FastBridgeCaller{contract: contract}, FastBridgeTransactor: FastBridgeTransactor{contract: contract}, FastBridgeFilterer: FastBridgeFilterer{contract: contract}}, nil
}

// NewFastBridgeCaller creates a new read-only instance of FastBridge, bound to a specific deployed contract.
func NewFastBridgeCaller(address common.Address, caller bind.ContractCaller) (*FastBridgeCaller, error) {
	contract, err := bindFastBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastBridgeCaller{contract: contract}, nil
}

// NewFastBridgeTransactor creates a new write-only instance of FastBridge, bound to a specific deployed contract.
func NewFastBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*FastBridgeTransactor, error) {
	contract, err := bindFastBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastBridgeTransactor{contract: contract}, nil
}

// NewFastBridgeFilterer creates a new log filterer instance of FastBridge, bound to a specific deployed contract.
func NewFastBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*FastBridgeFilterer, error) {
	contract, err := bindFastBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastBridgeFilterer{contract: contract}, nil
}

// bindFastBridge binds a generic wrapper to an already deployed contract.
func bindFastBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FastBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastBridge *FastBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastBridge.Contract.FastBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastBridge *FastBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastBridge.Contract.FastBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastBridge *FastBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastBridge.Contract.FastBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastBridge *FastBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastBridge *FastBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastBridge *FastBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FastBridge.Contract.DEFAULTADMINROLE(&_FastBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FastBridge.Contract.DEFAULTADMINROLE(&_FastBridge.CallOpts)
}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeCaller) DISPUTEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "DISPUTE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeSession) DISPUTEPERIOD() (*big.Int, error) {
	return _FastBridge.Contract.DISPUTEPERIOD(&_FastBridge.CallOpts)
}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) DISPUTEPERIOD() (*big.Int, error) {
	return _FastBridge.Contract.DISPUTEPERIOD(&_FastBridge.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_FastBridge *FastBridgeCaller) FEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_FastBridge *FastBridgeSession) FEEBPS() (*big.Int, error) {
	return _FastBridge.Contract.FEEBPS(&_FastBridge.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) FEEBPS() (*big.Int, error) {
	return _FastBridge.Contract.FEEBPS(&_FastBridge.CallOpts)
}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_FastBridge *FastBridgeCaller) FEERATEMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "FEE_RATE_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_FastBridge *FastBridgeSession) FEERATEMAX() (*big.Int, error) {
	return _FastBridge.Contract.FEERATEMAX(&_FastBridge.CallOpts)
}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) FEERATEMAX() (*big.Int, error) {
	return _FastBridge.Contract.FEERATEMAX(&_FastBridge.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCaller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeSession) GOVERNORROLE() ([32]byte, error) {
	return _FastBridge.Contract.GOVERNORROLE(&_FastBridge.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) GOVERNORROLE() ([32]byte, error) {
	return _FastBridge.Contract.GOVERNORROLE(&_FastBridge.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCaller) GUARDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "GUARD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeSession) GUARDROLE() ([32]byte, error) {
	return _FastBridge.Contract.GUARDROLE(&_FastBridge.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) GUARDROLE() ([32]byte, error) {
	return _FastBridge.Contract.GUARDROLE(&_FastBridge.CallOpts)
}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeCaller) MINDEADLINEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "MIN_DEADLINE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeSession) MINDEADLINEPERIOD() (*big.Int, error) {
	return _FastBridge.Contract.MINDEADLINEPERIOD(&_FastBridge.CallOpts)
}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) MINDEADLINEPERIOD() (*big.Int, error) {
	return _FastBridge.Contract.MINDEADLINEPERIOD(&_FastBridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeSession) RELAYERROLE() ([32]byte, error) {
	return _FastBridge.Contract.RELAYERROLE(&_FastBridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) RELAYERROLE() ([32]byte, error) {
	return _FastBridge.Contract.RELAYERROLE(&_FastBridge.CallOpts)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 ) view returns(uint96 timestamp, address relayer)
func (_FastBridge *FastBridgeCaller) BridgeProofs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "bridgeProofs", arg0)

	outstruct := new(struct {
		Timestamp *big.Int
		Relayer   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Relayer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 ) view returns(uint96 timestamp, address relayer)
func (_FastBridge *FastBridgeSession) BridgeProofs(arg0 [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _FastBridge.Contract.BridgeProofs(&_FastBridge.CallOpts, arg0)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 ) view returns(uint96 timestamp, address relayer)
func (_FastBridge *FastBridgeCallerSession) BridgeProofs(arg0 [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _FastBridge.Contract.BridgeProofs(&_FastBridge.CallOpts, arg0)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 ) view returns(bool)
func (_FastBridge *FastBridgeCaller) BridgeRelays(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "bridgeRelays", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 ) view returns(bool)
func (_FastBridge *FastBridgeSession) BridgeRelays(arg0 [32]byte) (bool, error) {
	return _FastBridge.Contract.BridgeRelays(&_FastBridge.CallOpts, arg0)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 ) view returns(bool)
func (_FastBridge *FastBridgeCallerSession) BridgeRelays(arg0 [32]byte) (bool, error) {
	return _FastBridge.Contract.BridgeRelays(&_FastBridge.CallOpts, arg0)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 ) view returns(uint8)
func (_FastBridge *FastBridgeCaller) BridgeStatuses(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "bridgeStatuses", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 ) view returns(uint8)
func (_FastBridge *FastBridgeSession) BridgeStatuses(arg0 [32]byte) (uint8, error) {
	return _FastBridge.Contract.BridgeStatuses(&_FastBridge.CallOpts, arg0)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 ) view returns(uint8)
func (_FastBridge *FastBridgeCallerSession) BridgeStatuses(arg0 [32]byte) (uint8, error) {
	return _FastBridge.Contract.BridgeStatuses(&_FastBridge.CallOpts, arg0)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_FastBridge *FastBridgeCaller) CanClaim(opts *bind.CallOpts, transactionId [32]byte, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "canClaim", transactionId, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_FastBridge *FastBridgeSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _FastBridge.Contract.CanClaim(&_FastBridge.CallOpts, transactionId, relayer)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_FastBridge *FastBridgeCallerSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _FastBridge.Contract.CanClaim(&_FastBridge.CallOpts, transactionId, relayer)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_FastBridge *FastBridgeCaller) ChainGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "chainGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_FastBridge *FastBridgeSession) ChainGasAmount() (*big.Int, error) {
	return _FastBridge.Contract.ChainGasAmount(&_FastBridge.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) ChainGasAmount() (*big.Int, error) {
	return _FastBridge.Contract.ChainGasAmount(&_FastBridge.CallOpts)
}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_FastBridge *FastBridgeCaller) DeployBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "deployBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_FastBridge *FastBridgeSession) DeployBlock() (*big.Int, error) {
	return _FastBridge.Contract.DeployBlock(&_FastBridge.CallOpts)
}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) DeployBlock() (*big.Int, error) {
	return _FastBridge.Contract.DeployBlock(&_FastBridge.CallOpts)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_FastBridge *FastBridgeCaller) GetBridgeTransaction(opts *bind.CallOpts, request []byte) (IFastBridgeBridgeTransaction, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "getBridgeTransaction", request)

	if err != nil {
		return *new(IFastBridgeBridgeTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeBridgeTransaction)).(*IFastBridgeBridgeTransaction)

	return out0, err

}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_FastBridge *FastBridgeSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _FastBridge.Contract.GetBridgeTransaction(&_FastBridge.CallOpts, request)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_FastBridge *FastBridgeCallerSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _FastBridge.Contract.GetBridgeTransaction(&_FastBridge.CallOpts, request)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridge *FastBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridge *FastBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FastBridge.Contract.GetRoleAdmin(&_FastBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FastBridge.Contract.GetRoleAdmin(&_FastBridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridge *FastBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridge *FastBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FastBridge.Contract.HasRole(&_FastBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridge *FastBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FastBridge.Contract.HasRole(&_FastBridge.CallOpts, role, account)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridge *FastBridgeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridge *FastBridgeSession) Nonce() (*big.Int, error) {
	return _FastBridge.Contract.Nonce(&_FastBridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) Nonce() (*big.Int, error) {
	return _FastBridge.Contract.Nonce(&_FastBridge.CallOpts)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_FastBridge *FastBridgeCaller) ProtocolFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "protocolFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_FastBridge *FastBridgeSession) ProtocolFeeRate() (*big.Int, error) {
	return _FastBridge.Contract.ProtocolFeeRate(&_FastBridge.CallOpts)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) ProtocolFeeRate() (*big.Int, error) {
	return _FastBridge.Contract.ProtocolFeeRate(&_FastBridge.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_FastBridge *FastBridgeCaller) ProtocolFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "protocolFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_FastBridge *FastBridgeSession) ProtocolFees(arg0 common.Address) (*big.Int, error) {
	return _FastBridge.Contract.ProtocolFees(&_FastBridge.CallOpts, arg0)
}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) ProtocolFees(arg0 common.Address) (*big.Int, error) {
	return _FastBridge.Contract.ProtocolFees(&_FastBridge.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridge *FastBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridge *FastBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FastBridge.Contract.SupportsInterface(&_FastBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridge *FastBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FastBridge.Contract.SupportsInterface(&_FastBridge.CallOpts, interfaceId)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_FastBridge *FastBridgeTransactor) AddGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "addGovernor", _governor)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_FastBridge *FastBridgeSession) AddGovernor(_governor common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddGovernor(&_FastBridge.TransactOpts, _governor)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_FastBridge *FastBridgeTransactorSession) AddGovernor(_governor common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddGovernor(&_FastBridge.TransactOpts, _governor)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_FastBridge *FastBridgeTransactor) AddGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "addGuard", _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_FastBridge *FastBridgeSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddGuard(&_FastBridge.TransactOpts, _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_FastBridge *FastBridgeTransactorSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddGuard(&_FastBridge.TransactOpts, _guard)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeTransactor) AddRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "addRelayer", _relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeSession) AddRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddRelayer(&_FastBridge.TransactOpts, _relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeTransactorSession) AddRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddRelayer(&_FastBridge.TransactOpts, _relayer)
}

// Bridge is a paid mutator transaction binding the contract method 0x778d31ff.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_FastBridge *FastBridgeTransactor) Bridge(opts *bind.TransactOpts, params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "bridge", params)
}

// Bridge is a paid mutator transaction binding the contract method 0x778d31ff.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_FastBridge *FastBridgeSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridge.Contract.Bridge(&_FastBridge.TransactOpts, params)
}

// Bridge is a paid mutator transaction binding the contract method 0x778d31ff.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_FastBridge *FastBridgeTransactorSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridge.Contract.Bridge(&_FastBridge.TransactOpts, params)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridge *FastBridgeTransactor) Claim(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "claim", request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridge *FastBridgeSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.Claim(&_FastBridge.TransactOpts, request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridge *FastBridgeTransactorSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.Claim(&_FastBridge.TransactOpts, request, to)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridge *FastBridgeTransactor) Dispute(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "dispute", transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridge *FastBridgeSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Dispute(&_FastBridge.TransactOpts, transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridge *FastBridgeTransactorSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Dispute(&_FastBridge.TransactOpts, transactionId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.GrantRole(&_FastBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.GrantRole(&_FastBridge.TransactOpts, role, account)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridge *FastBridgeTransactor) Prove(opts *bind.TransactOpts, request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "prove", request, destTxHash)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridge *FastBridgeSession) Prove(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Prove(&_FastBridge.TransactOpts, request, destTxHash)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridge *FastBridgeTransactorSession) Prove(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Prove(&_FastBridge.TransactOpts, request, destTxHash)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_FastBridge *FastBridgeTransactor) Refund(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "refund", request, to)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_FastBridge *FastBridgeSession) Refund(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.Refund(&_FastBridge.TransactOpts, request, to)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_FastBridge *FastBridgeTransactorSession) Refund(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.Refund(&_FastBridge.TransactOpts, request, to)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridge *FastBridgeTransactor) Relay(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "relay", request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridge *FastBridgeSession) Relay(request []byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Relay(&_FastBridge.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridge *FastBridgeTransactorSession) Relay(request []byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Relay(&_FastBridge.TransactOpts, request)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_FastBridge *FastBridgeTransactor) RemoveGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "removeGovernor", _governor)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_FastBridge *FastBridgeSession) RemoveGovernor(_governor common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveGovernor(&_FastBridge.TransactOpts, _governor)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_FastBridge *FastBridgeTransactorSession) RemoveGovernor(_governor common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveGovernor(&_FastBridge.TransactOpts, _governor)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_FastBridge *FastBridgeTransactor) RemoveGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "removeGuard", _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_FastBridge *FastBridgeSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveGuard(&_FastBridge.TransactOpts, _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_FastBridge *FastBridgeTransactorSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveGuard(&_FastBridge.TransactOpts, _guard)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeTransactor) RemoveRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "removeRelayer", _relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeSession) RemoveRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveRelayer(&_FastBridge.TransactOpts, _relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeTransactorSession) RemoveRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveRelayer(&_FastBridge.TransactOpts, _relayer)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridge *FastBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridge *FastBridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RenounceRole(&_FastBridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridge *FastBridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RenounceRole(&_FastBridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RevokeRole(&_FastBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RevokeRole(&_FastBridge.TransactOpts, role, account)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_FastBridge *FastBridgeTransactor) SetChainGasAmount(opts *bind.TransactOpts, newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "setChainGasAmount", newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_FastBridge *FastBridgeSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _FastBridge.Contract.SetChainGasAmount(&_FastBridge.TransactOpts, newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_FastBridge *FastBridgeTransactorSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _FastBridge.Contract.SetChainGasAmount(&_FastBridge.TransactOpts, newChainGasAmount)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_FastBridge *FastBridgeTransactor) SetProtocolFeeRate(opts *bind.TransactOpts, newFeeRate *big.Int) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "setProtocolFeeRate", newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_FastBridge *FastBridgeSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _FastBridge.Contract.SetProtocolFeeRate(&_FastBridge.TransactOpts, newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_FastBridge *FastBridgeTransactorSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _FastBridge.Contract.SetProtocolFeeRate(&_FastBridge.TransactOpts, newFeeRate)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_FastBridge *FastBridgeTransactor) SweepProtocolFees(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "sweepProtocolFees", token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_FastBridge *FastBridgeSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.SweepProtocolFees(&_FastBridge.TransactOpts, token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_FastBridge *FastBridgeTransactorSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.SweepProtocolFees(&_FastBridge.TransactOpts, token, recipient)
}

// FastBridgeBridgeDepositClaimedIterator is returned from FilterBridgeDepositClaimed and is used to iterate over the raw logs and unpacked data for BridgeDepositClaimed events raised by the FastBridge contract.
type FastBridgeBridgeDepositClaimedIterator struct {
	Event *FastBridgeBridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeDepositClaimed)
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
		it.Event = new(FastBridgeBridgeDepositClaimed)
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
func (it *FastBridgeBridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeDepositClaimed represents a BridgeDepositClaimed event raised by the FastBridge contract.
type FastBridgeBridgeDepositClaimed struct {
	TransactionId [32]byte
	Relayer       common.Address
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositClaimed is a free log retrieval operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) FilterBridgeDepositClaimed(opts *bind.FilterOpts) (*FastBridgeBridgeDepositClaimedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeDepositClaimed")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeDepositClaimedIterator{contract: _FastBridge.contract, event: "BridgeDepositClaimed", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositClaimed is a free log subscription operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) WatchBridgeDepositClaimed(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeDepositClaimed) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeDepositClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeDepositClaimed)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
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

// ParseBridgeDepositClaimed is a log parse operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) ParseBridgeDepositClaimed(log types.Log) (*FastBridgeBridgeDepositClaimed, error) {
	event := new(FastBridgeBridgeDepositClaimed)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeDepositRefundedIterator is returned from FilterBridgeDepositRefunded and is used to iterate over the raw logs and unpacked data for BridgeDepositRefunded events raised by the FastBridge contract.
type FastBridgeBridgeDepositRefundedIterator struct {
	Event *FastBridgeBridgeDepositRefunded // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeDepositRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeDepositRefunded)
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
		it.Event = new(FastBridgeBridgeDepositRefunded)
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
func (it *FastBridgeBridgeDepositRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeDepositRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeDepositRefunded represents a BridgeDepositRefunded event raised by the FastBridge contract.
type FastBridgeBridgeDepositRefunded struct {
	TransactionId [32]byte
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositRefunded is a free log retrieval operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) FilterBridgeDepositRefunded(opts *bind.FilterOpts) (*FastBridgeBridgeDepositRefundedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeDepositRefunded")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeDepositRefundedIterator{contract: _FastBridge.contract, event: "BridgeDepositRefunded", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositRefunded is a free log subscription operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) WatchBridgeDepositRefunded(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeDepositRefunded) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeDepositRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeDepositRefunded)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
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

// ParseBridgeDepositRefunded is a log parse operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) ParseBridgeDepositRefunded(log types.Log) (*FastBridgeBridgeDepositRefunded, error) {
	event := new(FastBridgeBridgeDepositRefunded)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeProofDisputedIterator is returned from FilterBridgeProofDisputed and is used to iterate over the raw logs and unpacked data for BridgeProofDisputed events raised by the FastBridge contract.
type FastBridgeBridgeProofDisputedIterator struct {
	Event *FastBridgeBridgeProofDisputed // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeProofDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeProofDisputed)
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
		it.Event = new(FastBridgeBridgeProofDisputed)
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
func (it *FastBridgeBridgeProofDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeProofDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeProofDisputed represents a BridgeProofDisputed event raised by the FastBridge contract.
type FastBridgeBridgeProofDisputed struct {
	TransactionId [32]byte
	Relayer       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofDisputed is a free log retrieval operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_FastBridge *FastBridgeFilterer) FilterBridgeProofDisputed(opts *bind.FilterOpts) (*FastBridgeBridgeProofDisputedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeProofDisputed")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeProofDisputedIterator{contract: _FastBridge.contract, event: "BridgeProofDisputed", logs: logs, sub: sub}, nil
}

// WatchBridgeProofDisputed is a free log subscription operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_FastBridge *FastBridgeFilterer) WatchBridgeProofDisputed(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeProofDisputed) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeProofDisputed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeProofDisputed)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
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

// ParseBridgeProofDisputed is a log parse operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_FastBridge *FastBridgeFilterer) ParseBridgeProofDisputed(log types.Log) (*FastBridgeBridgeProofDisputed, error) {
	event := new(FastBridgeBridgeProofDisputed)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeProofProvidedIterator is returned from FilterBridgeProofProvided and is used to iterate over the raw logs and unpacked data for BridgeProofProvided events raised by the FastBridge contract.
type FastBridgeBridgeProofProvidedIterator struct {
	Event *FastBridgeBridgeProofProvided // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeProofProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeProofProvided)
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
		it.Event = new(FastBridgeBridgeProofProvided)
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
func (it *FastBridgeBridgeProofProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeProofProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeProofProvided represents a BridgeProofProvided event raised by the FastBridge contract.
type FastBridgeBridgeProofProvided struct {
	TransactionId   [32]byte
	Relayer         common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofProvided is a free log retrieval operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_FastBridge *FastBridgeFilterer) FilterBridgeProofProvided(opts *bind.FilterOpts) (*FastBridgeBridgeProofProvidedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeProofProvided")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeProofProvidedIterator{contract: _FastBridge.contract, event: "BridgeProofProvided", logs: logs, sub: sub}, nil
}

// WatchBridgeProofProvided is a free log subscription operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_FastBridge *FastBridgeFilterer) WatchBridgeProofProvided(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeProofProvided) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeProofProvided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeProofProvided)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
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

// ParseBridgeProofProvided is a log parse operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_FastBridge *FastBridgeFilterer) ParseBridgeProofProvided(log types.Log) (*FastBridgeBridgeProofProvided, error) {
	event := new(FastBridgeBridgeProofProvided)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeRelayedIterator is returned from FilterBridgeRelayed and is used to iterate over the raw logs and unpacked data for BridgeRelayed events raised by the FastBridge contract.
type FastBridgeBridgeRelayedIterator struct {
	Event *FastBridgeBridgeRelayed // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeRelayed)
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
		it.Event = new(FastBridgeBridgeRelayed)
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
func (it *FastBridgeBridgeRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeRelayed represents a BridgeRelayed event raised by the FastBridge contract.
type FastBridgeBridgeRelayed struct {
	TransactionId  [32]byte
	Relayer        common.Address
	To             common.Address
	Token          common.Address
	Amount         *big.Int
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeRelayed is a free log retrieval operation binding the contract event 0xf5d0e166883753ed3b66d60c893dda5f6760044b7bba1317907a9385c6620f9b.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount, uint256 chainGasAmount)
func (_FastBridge *FastBridgeFilterer) FilterBridgeRelayed(opts *bind.FilterOpts) (*FastBridgeBridgeRelayedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeRelayed")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeRelayedIterator{contract: _FastBridge.contract, event: "BridgeRelayed", logs: logs, sub: sub}, nil
}

// WatchBridgeRelayed is a free log subscription operation binding the contract event 0xf5d0e166883753ed3b66d60c893dda5f6760044b7bba1317907a9385c6620f9b.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount, uint256 chainGasAmount)
func (_FastBridge *FastBridgeFilterer) WatchBridgeRelayed(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeRelayed) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeRelayed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeRelayed)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
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

// ParseBridgeRelayed is a log parse operation binding the contract event 0xf5d0e166883753ed3b66d60c893dda5f6760044b7bba1317907a9385c6620f9b.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount, uint256 chainGasAmount)
func (_FastBridge *FastBridgeFilterer) ParseBridgeRelayed(log types.Log) (*FastBridgeBridgeRelayed, error) {
	event := new(FastBridgeBridgeRelayed)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeRequestedIterator is returned from FilterBridgeRequested and is used to iterate over the raw logs and unpacked data for BridgeRequested events raised by the FastBridge contract.
type FastBridgeBridgeRequestedIterator struct {
	Event *FastBridgeBridgeRequested // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeRequested)
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
		it.Event = new(FastBridgeBridgeRequested)
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
func (it *FastBridgeBridgeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeRequested represents a BridgeRequested event raised by the FastBridge contract.
type FastBridgeBridgeRequested struct {
	TransactionId [32]byte
	Sender        common.Address
	Request       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRequested is a free log retrieval operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_FastBridge *FastBridgeFilterer) FilterBridgeRequested(opts *bind.FilterOpts) (*FastBridgeBridgeRequestedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeRequested")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeRequestedIterator{contract: _FastBridge.contract, event: "BridgeRequested", logs: logs, sub: sub}, nil
}

// WatchBridgeRequested is a free log subscription operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_FastBridge *FastBridgeFilterer) WatchBridgeRequested(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeRequested) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeRequested)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
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

// ParseBridgeRequested is a log parse operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_FastBridge *FastBridgeFilterer) ParseBridgeRequested(log types.Log) (*FastBridgeBridgeRequested, error) {
	event := new(FastBridgeBridgeRequested)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeChainGasAmountUpdatedIterator is returned from FilterChainGasAmountUpdated and is used to iterate over the raw logs and unpacked data for ChainGasAmountUpdated events raised by the FastBridge contract.
type FastBridgeChainGasAmountUpdatedIterator struct {
	Event *FastBridgeChainGasAmountUpdated // Event containing the contract specifics and raw log

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
func (it *FastBridgeChainGasAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeChainGasAmountUpdated)
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
		it.Event = new(FastBridgeChainGasAmountUpdated)
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
func (it *FastBridgeChainGasAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeChainGasAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeChainGasAmountUpdated represents a ChainGasAmountUpdated event raised by the FastBridge contract.
type FastBridgeChainGasAmountUpdated struct {
	OldChainGasAmount *big.Int
	NewChainGasAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChainGasAmountUpdated is a free log retrieval operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_FastBridge *FastBridgeFilterer) FilterChainGasAmountUpdated(opts *bind.FilterOpts) (*FastBridgeChainGasAmountUpdatedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &FastBridgeChainGasAmountUpdatedIterator{contract: _FastBridge.contract, event: "ChainGasAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchChainGasAmountUpdated is a free log subscription operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_FastBridge *FastBridgeFilterer) WatchChainGasAmountUpdated(opts *bind.WatchOpts, sink chan<- *FastBridgeChainGasAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeChainGasAmountUpdated)
				if err := _FastBridge.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
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

// ParseChainGasAmountUpdated is a log parse operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_FastBridge *FastBridgeFilterer) ParseChainGasAmountUpdated(log types.Log) (*FastBridgeChainGasAmountUpdated, error) {
	event := new(FastBridgeChainGasAmountUpdated)
	if err := _FastBridge.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeFeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the FastBridge contract.
type FastBridgeFeeRateUpdatedIterator struct {
	Event *FastBridgeFeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *FastBridgeFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeFeeRateUpdated)
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
		it.Event = new(FastBridgeFeeRateUpdated)
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
func (it *FastBridgeFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeFeeRateUpdated represents a FeeRateUpdated event raised by the FastBridge contract.
type FastBridgeFeeRateUpdated struct {
	OldFeeRate *big.Int
	NewFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_FastBridge *FastBridgeFilterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*FastBridgeFeeRateUpdatedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &FastBridgeFeeRateUpdatedIterator{contract: _FastBridge.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_FastBridge *FastBridgeFilterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *FastBridgeFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeFeeRateUpdated)
				if err := _FastBridge.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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

// ParseFeeRateUpdated is a log parse operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_FastBridge *FastBridgeFilterer) ParseFeeRateUpdated(log types.Log) (*FastBridgeFeeRateUpdated, error) {
	event := new(FastBridgeFeeRateUpdated)
	if err := _FastBridge.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeFeesSweptIterator is returned from FilterFeesSwept and is used to iterate over the raw logs and unpacked data for FeesSwept events raised by the FastBridge contract.
type FastBridgeFeesSweptIterator struct {
	Event *FastBridgeFeesSwept // Event containing the contract specifics and raw log

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
func (it *FastBridgeFeesSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeFeesSwept)
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
		it.Event = new(FastBridgeFeesSwept)
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
func (it *FastBridgeFeesSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeFeesSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeFeesSwept represents a FeesSwept event raised by the FastBridge contract.
type FastBridgeFeesSwept struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesSwept is a free log retrieval operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_FastBridge *FastBridgeFilterer) FilterFeesSwept(opts *bind.FilterOpts) (*FastBridgeFeesSweptIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return &FastBridgeFeesSweptIterator{contract: _FastBridge.contract, event: "FeesSwept", logs: logs, sub: sub}, nil
}

// WatchFeesSwept is a free log subscription operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_FastBridge *FastBridgeFilterer) WatchFeesSwept(opts *bind.WatchOpts, sink chan<- *FastBridgeFeesSwept) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeFeesSwept)
				if err := _FastBridge.contract.UnpackLog(event, "FeesSwept", log); err != nil {
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

// ParseFeesSwept is a log parse operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_FastBridge *FastBridgeFilterer) ParseFeesSwept(log types.Log) (*FastBridgeFeesSwept, error) {
	event := new(FastBridgeFeesSwept)
	if err := _FastBridge.contract.UnpackLog(event, "FeesSwept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the FastBridge contract.
type FastBridgeGovernorAddedIterator struct {
	Event *FastBridgeGovernorAdded // Event containing the contract specifics and raw log

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
func (it *FastBridgeGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeGovernorAdded)
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
		it.Event = new(FastBridgeGovernorAdded)
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
func (it *FastBridgeGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeGovernorAdded represents a GovernorAdded event raised by the FastBridge contract.
type FastBridgeGovernorAdded struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_FastBridge *FastBridgeFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*FastBridgeGovernorAddedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &FastBridgeGovernorAddedIterator{contract: _FastBridge.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_FastBridge *FastBridgeFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *FastBridgeGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeGovernorAdded)
				if err := _FastBridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_FastBridge *FastBridgeFilterer) ParseGovernorAdded(log types.Log) (*FastBridgeGovernorAdded, error) {
	event := new(FastBridgeGovernorAdded)
	if err := _FastBridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the FastBridge contract.
type FastBridgeGovernorRemovedIterator struct {
	Event *FastBridgeGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *FastBridgeGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeGovernorRemoved)
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
		it.Event = new(FastBridgeGovernorRemoved)
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
func (it *FastBridgeGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeGovernorRemoved represents a GovernorRemoved event raised by the FastBridge contract.
type FastBridgeGovernorRemoved struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_FastBridge *FastBridgeFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*FastBridgeGovernorRemovedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &FastBridgeGovernorRemovedIterator{contract: _FastBridge.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_FastBridge *FastBridgeFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *FastBridgeGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeGovernorRemoved)
				if err := _FastBridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_FastBridge *FastBridgeFilterer) ParseGovernorRemoved(log types.Log) (*FastBridgeGovernorRemoved, error) {
	event := new(FastBridgeGovernorRemoved)
	if err := _FastBridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the FastBridge contract.
type FastBridgeGuardAddedIterator struct {
	Event *FastBridgeGuardAdded // Event containing the contract specifics and raw log

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
func (it *FastBridgeGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeGuardAdded)
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
		it.Event = new(FastBridgeGuardAdded)
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
func (it *FastBridgeGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeGuardAdded represents a GuardAdded event raised by the FastBridge contract.
type FastBridgeGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_FastBridge *FastBridgeFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*FastBridgeGuardAddedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &FastBridgeGuardAddedIterator{contract: _FastBridge.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_FastBridge *FastBridgeFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *FastBridgeGuardAdded) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeGuardAdded)
				if err := _FastBridge.contract.UnpackLog(event, "GuardAdded", log); err != nil {
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

// ParseGuardAdded is a log parse operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_FastBridge *FastBridgeFilterer) ParseGuardAdded(log types.Log) (*FastBridgeGuardAdded, error) {
	event := new(FastBridgeGuardAdded)
	if err := _FastBridge.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the FastBridge contract.
type FastBridgeGuardRemovedIterator struct {
	Event *FastBridgeGuardRemoved // Event containing the contract specifics and raw log

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
func (it *FastBridgeGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeGuardRemoved)
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
		it.Event = new(FastBridgeGuardRemoved)
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
func (it *FastBridgeGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeGuardRemoved represents a GuardRemoved event raised by the FastBridge contract.
type FastBridgeGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_FastBridge *FastBridgeFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*FastBridgeGuardRemovedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &FastBridgeGuardRemovedIterator{contract: _FastBridge.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_FastBridge *FastBridgeFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *FastBridgeGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeGuardRemoved)
				if err := _FastBridge.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
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

// ParseGuardRemoved is a log parse operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_FastBridge *FastBridgeFilterer) ParseGuardRemoved(log types.Log) (*FastBridgeGuardRemoved, error) {
	event := new(FastBridgeGuardRemoved)
	if err := _FastBridge.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the FastBridge contract.
type FastBridgeRelayerAddedIterator struct {
	Event *FastBridgeRelayerAdded // Event containing the contract specifics and raw log

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
func (it *FastBridgeRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRelayerAdded)
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
		it.Event = new(FastBridgeRelayerAdded)
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
func (it *FastBridgeRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRelayerAdded represents a RelayerAdded event raised by the FastBridge contract.
type FastBridgeRelayerAdded struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_FastBridge *FastBridgeFilterer) FilterRelayerAdded(opts *bind.FilterOpts) (*FastBridgeRelayerAddedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RelayerAdded")
	if err != nil {
		return nil, err
	}
	return &FastBridgeRelayerAddedIterator{contract: _FastBridge.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_FastBridge *FastBridgeFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *FastBridgeRelayerAdded) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RelayerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRelayerAdded)
				if err := _FastBridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
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

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_FastBridge *FastBridgeFilterer) ParseRelayerAdded(log types.Log) (*FastBridgeRelayerAdded, error) {
	event := new(FastBridgeRelayerAdded)
	if err := _FastBridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the FastBridge contract.
type FastBridgeRelayerRemovedIterator struct {
	Event *FastBridgeRelayerRemoved // Event containing the contract specifics and raw log

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
func (it *FastBridgeRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRelayerRemoved)
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
		it.Event = new(FastBridgeRelayerRemoved)
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
func (it *FastBridgeRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRelayerRemoved represents a RelayerRemoved event raised by the FastBridge contract.
type FastBridgeRelayerRemoved struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_FastBridge *FastBridgeFilterer) FilterRelayerRemoved(opts *bind.FilterOpts) (*FastBridgeRelayerRemovedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RelayerRemoved")
	if err != nil {
		return nil, err
	}
	return &FastBridgeRelayerRemovedIterator{contract: _FastBridge.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_FastBridge *FastBridgeFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *FastBridgeRelayerRemoved) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RelayerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRelayerRemoved)
				if err := _FastBridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
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

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_FastBridge *FastBridgeFilterer) ParseRelayerRemoved(log types.Log) (*FastBridgeRelayerRemoved, error) {
	event := new(FastBridgeRelayerRemoved)
	if err := _FastBridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FastBridge contract.
type FastBridgeRoleAdminChangedIterator struct {
	Event *FastBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FastBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRoleAdminChanged)
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
		it.Event = new(FastBridgeRoleAdminChanged)
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
func (it *FastBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the FastBridge contract.
type FastBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FastBridge *FastBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FastBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeRoleAdminChangedIterator{contract: _FastBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FastBridge *FastBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FastBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRoleAdminChanged)
				if err := _FastBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_FastBridge *FastBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*FastBridgeRoleAdminChanged, error) {
	event := new(FastBridgeRoleAdminChanged)
	if err := _FastBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FastBridge contract.
type FastBridgeRoleGrantedIterator struct {
	Event *FastBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *FastBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRoleGranted)
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
		it.Event = new(FastBridgeRoleGranted)
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
func (it *FastBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRoleGranted represents a RoleGranted event raised by the FastBridge contract.
type FastBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FastBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeRoleGrantedIterator{contract: _FastBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FastBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRoleGranted)
				if err := _FastBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_FastBridge *FastBridgeFilterer) ParseRoleGranted(log types.Log) (*FastBridgeRoleGranted, error) {
	event := new(FastBridgeRoleGranted)
	if err := _FastBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FastBridge contract.
type FastBridgeRoleRevokedIterator struct {
	Event *FastBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FastBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRoleRevoked)
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
		it.Event = new(FastBridgeRoleRevoked)
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
func (it *FastBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRoleRevoked represents a RoleRevoked event raised by the FastBridge contract.
type FastBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FastBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeRoleRevokedIterator{contract: _FastBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FastBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRoleRevoked)
				if err := _FastBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_FastBridge *FastBridgeFilterer) ParseRoleRevoked(log types.Log) (*FastBridgeRoleRevoked, error) {
	event := new(FastBridgeRoleRevoked)
	if err := _FastBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
	parsed, err := abi.JSON(strings.NewReader(IAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// IAdminMetaData contains all meta data concerning the IAdmin contract.
var IAdminMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldChainGasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"addGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"addRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"removeGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"removeRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3c4a25d0": "addGovernor(address)",
		"6913a63c": "addGuard(address)",
		"dd39f00d": "addRelayer(address)",
		"eecdac88": "removeGovernor(address)",
		"b6235016": "removeGuard(address)",
		"60f0a5ac": "removeRelayer(address)",
		"b250fe6b": "setChainGasAmount(uint256)",
		"b13aa2d6": "setProtocolFeeRate(uint256)",
		"06f333f2": "sweepProtocolFees(address,address)",
	},
}

// IAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use IAdminMetaData.ABI instead.
var IAdminABI = IAdminMetaData.ABI

// Deprecated: Use IAdminMetaData.Sigs instead.
// IAdminFuncSigs maps the 4-byte function signature to its string representation.
var IAdminFuncSigs = IAdminMetaData.Sigs

// IAdmin is an auto generated Go binding around an Ethereum contract.
type IAdmin struct {
	IAdminCaller     // Read-only binding to the contract
	IAdminTransactor // Write-only binding to the contract
	IAdminFilterer   // Log filterer for contract events
}

// IAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAdminSession struct {
	Contract     *IAdmin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAdminCallerSession struct {
	Contract *IAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAdminTransactorSession struct {
	Contract     *IAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAdminRaw struct {
	Contract *IAdmin // Generic contract binding to access the raw methods on
}

// IAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAdminCallerRaw struct {
	Contract *IAdminCaller // Generic read-only contract binding to access the raw methods on
}

// IAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAdminTransactorRaw struct {
	Contract *IAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAdmin creates a new instance of IAdmin, bound to a specific deployed contract.
func NewIAdmin(address common.Address, backend bind.ContractBackend) (*IAdmin, error) {
	contract, err := bindIAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAdmin{IAdminCaller: IAdminCaller{contract: contract}, IAdminTransactor: IAdminTransactor{contract: contract}, IAdminFilterer: IAdminFilterer{contract: contract}}, nil
}

// NewIAdminCaller creates a new read-only instance of IAdmin, bound to a specific deployed contract.
func NewIAdminCaller(address common.Address, caller bind.ContractCaller) (*IAdminCaller, error) {
	contract, err := bindIAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminCaller{contract: contract}, nil
}

// NewIAdminTransactor creates a new write-only instance of IAdmin, bound to a specific deployed contract.
func NewIAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*IAdminTransactor, error) {
	contract, err := bindIAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminTransactor{contract: contract}, nil
}

// NewIAdminFilterer creates a new log filterer instance of IAdmin, bound to a specific deployed contract.
func NewIAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*IAdminFilterer, error) {
	contract, err := bindIAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAdminFilterer{contract: contract}, nil
}

// bindIAdmin binds a generic wrapper to an already deployed contract.
func bindIAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAdmin *IAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAdmin.Contract.IAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAdmin *IAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAdmin.Contract.IAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAdmin *IAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAdmin.Contract.IAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAdmin *IAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAdmin *IAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAdmin *IAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAdmin.Contract.contract.Transact(opts, method, params...)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_IAdmin *IAdminTransactor) AddGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "addGovernor", _governor)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_IAdmin *IAdminSession) AddGovernor(_governor common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.AddGovernor(&_IAdmin.TransactOpts, _governor)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _governor) returns()
func (_IAdmin *IAdminTransactorSession) AddGovernor(_governor common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.AddGovernor(&_IAdmin.TransactOpts, _governor)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_IAdmin *IAdminTransactor) AddGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "addGuard", _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_IAdmin *IAdminSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.AddGuard(&_IAdmin.TransactOpts, _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_IAdmin *IAdminTransactorSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.AddGuard(&_IAdmin.TransactOpts, _guard)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_IAdmin *IAdminTransactor) AddRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "addRelayer", _relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_IAdmin *IAdminSession) AddRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.AddRelayer(&_IAdmin.TransactOpts, _relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_IAdmin *IAdminTransactorSession) AddRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.AddRelayer(&_IAdmin.TransactOpts, _relayer)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_IAdmin *IAdminTransactor) RemoveGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "removeGovernor", _governor)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_IAdmin *IAdminSession) RemoveGovernor(_governor common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.RemoveGovernor(&_IAdmin.TransactOpts, _governor)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _governor) returns()
func (_IAdmin *IAdminTransactorSession) RemoveGovernor(_governor common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.RemoveGovernor(&_IAdmin.TransactOpts, _governor)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_IAdmin *IAdminTransactor) RemoveGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "removeGuard", _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_IAdmin *IAdminSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.RemoveGuard(&_IAdmin.TransactOpts, _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_IAdmin *IAdminTransactorSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.RemoveGuard(&_IAdmin.TransactOpts, _guard)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_IAdmin *IAdminTransactor) RemoveRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "removeRelayer", _relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_IAdmin *IAdminSession) RemoveRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.RemoveRelayer(&_IAdmin.TransactOpts, _relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_IAdmin *IAdminTransactorSession) RemoveRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.RemoveRelayer(&_IAdmin.TransactOpts, _relayer)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_IAdmin *IAdminTransactor) SetChainGasAmount(opts *bind.TransactOpts, newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "setChainGasAmount", newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_IAdmin *IAdminSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _IAdmin.Contract.SetChainGasAmount(&_IAdmin.TransactOpts, newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_IAdmin *IAdminTransactorSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _IAdmin.Contract.SetChainGasAmount(&_IAdmin.TransactOpts, newChainGasAmount)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_IAdmin *IAdminTransactor) SetProtocolFeeRate(opts *bind.TransactOpts, newFeeRate *big.Int) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "setProtocolFeeRate", newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_IAdmin *IAdminSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _IAdmin.Contract.SetProtocolFeeRate(&_IAdmin.TransactOpts, newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_IAdmin *IAdminTransactorSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _IAdmin.Contract.SetProtocolFeeRate(&_IAdmin.TransactOpts, newFeeRate)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_IAdmin *IAdminTransactor) SweepProtocolFees(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "sweepProtocolFees", token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_IAdmin *IAdminSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.SweepProtocolFees(&_IAdmin.TransactOpts, token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_IAdmin *IAdminTransactorSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.SweepProtocolFees(&_IAdmin.TransactOpts, token, recipient)
}

// IAdminChainGasAmountUpdatedIterator is returned from FilterChainGasAmountUpdated and is used to iterate over the raw logs and unpacked data for ChainGasAmountUpdated events raised by the IAdmin contract.
type IAdminChainGasAmountUpdatedIterator struct {
	Event *IAdminChainGasAmountUpdated // Event containing the contract specifics and raw log

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
func (it *IAdminChainGasAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminChainGasAmountUpdated)
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
		it.Event = new(IAdminChainGasAmountUpdated)
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
func (it *IAdminChainGasAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminChainGasAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminChainGasAmountUpdated represents a ChainGasAmountUpdated event raised by the IAdmin contract.
type IAdminChainGasAmountUpdated struct {
	OldChainGasAmount *big.Int
	NewChainGasAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChainGasAmountUpdated is a free log retrieval operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_IAdmin *IAdminFilterer) FilterChainGasAmountUpdated(opts *bind.FilterOpts) (*IAdminChainGasAmountUpdatedIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &IAdminChainGasAmountUpdatedIterator{contract: _IAdmin.contract, event: "ChainGasAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchChainGasAmountUpdated is a free log subscription operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_IAdmin *IAdminFilterer) WatchChainGasAmountUpdated(opts *bind.WatchOpts, sink chan<- *IAdminChainGasAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminChainGasAmountUpdated)
				if err := _IAdmin.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
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

// ParseChainGasAmountUpdated is a log parse operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_IAdmin *IAdminFilterer) ParseChainGasAmountUpdated(log types.Log) (*IAdminChainGasAmountUpdated, error) {
	event := new(IAdminChainGasAmountUpdated)
	if err := _IAdmin.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminFeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the IAdmin contract.
type IAdminFeeRateUpdatedIterator struct {
	Event *IAdminFeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *IAdminFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminFeeRateUpdated)
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
		it.Event = new(IAdminFeeRateUpdated)
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
func (it *IAdminFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminFeeRateUpdated represents a FeeRateUpdated event raised by the IAdmin contract.
type IAdminFeeRateUpdated struct {
	OldFeeRate *big.Int
	NewFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_IAdmin *IAdminFilterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*IAdminFeeRateUpdatedIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &IAdminFeeRateUpdatedIterator{contract: _IAdmin.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_IAdmin *IAdminFilterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *IAdminFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminFeeRateUpdated)
				if err := _IAdmin.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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

// ParseFeeRateUpdated is a log parse operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_IAdmin *IAdminFilterer) ParseFeeRateUpdated(log types.Log) (*IAdminFeeRateUpdated, error) {
	event := new(IAdminFeeRateUpdated)
	if err := _IAdmin.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminFeesSweptIterator is returned from FilterFeesSwept and is used to iterate over the raw logs and unpacked data for FeesSwept events raised by the IAdmin contract.
type IAdminFeesSweptIterator struct {
	Event *IAdminFeesSwept // Event containing the contract specifics and raw log

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
func (it *IAdminFeesSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminFeesSwept)
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
		it.Event = new(IAdminFeesSwept)
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
func (it *IAdminFeesSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminFeesSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminFeesSwept represents a FeesSwept event raised by the IAdmin contract.
type IAdminFeesSwept struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesSwept is a free log retrieval operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_IAdmin *IAdminFilterer) FilterFeesSwept(opts *bind.FilterOpts) (*IAdminFeesSweptIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return &IAdminFeesSweptIterator{contract: _IAdmin.contract, event: "FeesSwept", logs: logs, sub: sub}, nil
}

// WatchFeesSwept is a free log subscription operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_IAdmin *IAdminFilterer) WatchFeesSwept(opts *bind.WatchOpts, sink chan<- *IAdminFeesSwept) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminFeesSwept)
				if err := _IAdmin.contract.UnpackLog(event, "FeesSwept", log); err != nil {
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

// ParseFeesSwept is a log parse operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_IAdmin *IAdminFilterer) ParseFeesSwept(log types.Log) (*IAdminFeesSwept, error) {
	event := new(IAdminFeesSwept)
	if err := _IAdmin.contract.UnpackLog(event, "FeesSwept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the IAdmin contract.
type IAdminGovernorAddedIterator struct {
	Event *IAdminGovernorAdded // Event containing the contract specifics and raw log

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
func (it *IAdminGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminGovernorAdded)
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
		it.Event = new(IAdminGovernorAdded)
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
func (it *IAdminGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminGovernorAdded represents a GovernorAdded event raised by the IAdmin contract.
type IAdminGovernorAdded struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_IAdmin *IAdminFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*IAdminGovernorAddedIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &IAdminGovernorAddedIterator{contract: _IAdmin.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_IAdmin *IAdminFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *IAdminGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminGovernorAdded)
				if err := _IAdmin.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address governor)
func (_IAdmin *IAdminFilterer) ParseGovernorAdded(log types.Log) (*IAdminGovernorAdded, error) {
	event := new(IAdminGovernorAdded)
	if err := _IAdmin.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the IAdmin contract.
type IAdminGovernorRemovedIterator struct {
	Event *IAdminGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *IAdminGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminGovernorRemoved)
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
		it.Event = new(IAdminGovernorRemoved)
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
func (it *IAdminGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminGovernorRemoved represents a GovernorRemoved event raised by the IAdmin contract.
type IAdminGovernorRemoved struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_IAdmin *IAdminFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*IAdminGovernorRemovedIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &IAdminGovernorRemovedIterator{contract: _IAdmin.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_IAdmin *IAdminFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *IAdminGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminGovernorRemoved)
				if err := _IAdmin.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address governor)
func (_IAdmin *IAdminFilterer) ParseGovernorRemoved(log types.Log) (*IAdminGovernorRemoved, error) {
	event := new(IAdminGovernorRemoved)
	if err := _IAdmin.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the IAdmin contract.
type IAdminGuardAddedIterator struct {
	Event *IAdminGuardAdded // Event containing the contract specifics and raw log

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
func (it *IAdminGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminGuardAdded)
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
		it.Event = new(IAdminGuardAdded)
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
func (it *IAdminGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminGuardAdded represents a GuardAdded event raised by the IAdmin contract.
type IAdminGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_IAdmin *IAdminFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*IAdminGuardAddedIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &IAdminGuardAddedIterator{contract: _IAdmin.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_IAdmin *IAdminFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *IAdminGuardAdded) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminGuardAdded)
				if err := _IAdmin.contract.UnpackLog(event, "GuardAdded", log); err != nil {
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

// ParseGuardAdded is a log parse operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_IAdmin *IAdminFilterer) ParseGuardAdded(log types.Log) (*IAdminGuardAdded, error) {
	event := new(IAdminGuardAdded)
	if err := _IAdmin.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the IAdmin contract.
type IAdminGuardRemovedIterator struct {
	Event *IAdminGuardRemoved // Event containing the contract specifics and raw log

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
func (it *IAdminGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminGuardRemoved)
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
		it.Event = new(IAdminGuardRemoved)
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
func (it *IAdminGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminGuardRemoved represents a GuardRemoved event raised by the IAdmin contract.
type IAdminGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_IAdmin *IAdminFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*IAdminGuardRemovedIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &IAdminGuardRemovedIterator{contract: _IAdmin.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_IAdmin *IAdminFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *IAdminGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminGuardRemoved)
				if err := _IAdmin.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
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

// ParseGuardRemoved is a log parse operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_IAdmin *IAdminFilterer) ParseGuardRemoved(log types.Log) (*IAdminGuardRemoved, error) {
	event := new(IAdminGuardRemoved)
	if err := _IAdmin.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the IAdmin contract.
type IAdminRelayerAddedIterator struct {
	Event *IAdminRelayerAdded // Event containing the contract specifics and raw log

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
func (it *IAdminRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminRelayerAdded)
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
		it.Event = new(IAdminRelayerAdded)
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
func (it *IAdminRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminRelayerAdded represents a RelayerAdded event raised by the IAdmin contract.
type IAdminRelayerAdded struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_IAdmin *IAdminFilterer) FilterRelayerAdded(opts *bind.FilterOpts) (*IAdminRelayerAddedIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "RelayerAdded")
	if err != nil {
		return nil, err
	}
	return &IAdminRelayerAddedIterator{contract: _IAdmin.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_IAdmin *IAdminFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *IAdminRelayerAdded) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "RelayerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminRelayerAdded)
				if err := _IAdmin.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
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

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_IAdmin *IAdminFilterer) ParseRelayerAdded(log types.Log) (*IAdminRelayerAdded, error) {
	event := new(IAdminRelayerAdded)
	if err := _IAdmin.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the IAdmin contract.
type IAdminRelayerRemovedIterator struct {
	Event *IAdminRelayerRemoved // Event containing the contract specifics and raw log

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
func (it *IAdminRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminRelayerRemoved)
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
		it.Event = new(IAdminRelayerRemoved)
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
func (it *IAdminRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminRelayerRemoved represents a RelayerRemoved event raised by the IAdmin contract.
type IAdminRelayerRemoved struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_IAdmin *IAdminFilterer) FilterRelayerRemoved(opts *bind.FilterOpts) (*IAdminRelayerRemovedIterator, error) {

	logs, sub, err := _IAdmin.contract.FilterLogs(opts, "RelayerRemoved")
	if err != nil {
		return nil, err
	}
	return &IAdminRelayerRemovedIterator{contract: _IAdmin.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_IAdmin *IAdminFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *IAdminRelayerRemoved) (event.Subscription, error) {

	logs, sub, err := _IAdmin.contract.WatchLogs(opts, "RelayerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminRelayerRemoved)
				if err := _IAdmin.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
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

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_IAdmin *IAdminFilterer) ParseRelayerRemoved(log types.Log) (*IAdminRelayerRemoved, error) {
	event := new(IAdminRelayerRemoved)
	if err := _IAdmin.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
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
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
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

// IERC20PermitMetaData contains all meta data concerning the IERC20Permit contract.
var IERC20PermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"7ecebe00": "nonces(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
	},
}

// IERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitMetaData.ABI instead.
var IERC20PermitABI = IERC20PermitMetaData.ABI

// Deprecated: Use IERC20PermitMetaData.Sigs instead.
// IERC20PermitFuncSigs maps the 4-byte function signature to its string representation.
var IERC20PermitFuncSigs = IERC20PermitMetaData.Sigs

// IERC20Permit is an auto generated Go binding around an Ethereum contract.
type IERC20Permit struct {
	IERC20PermitCaller     // Read-only binding to the contract
	IERC20PermitTransactor // Write-only binding to the contract
	IERC20PermitFilterer   // Log filterer for contract events
}

// IERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitSession struct {
	Contract     *IERC20Permit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitCallerSession struct {
	Contract *IERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitTransactorSession struct {
	Contract     *IERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitRaw struct {
	Contract *IERC20Permit // Generic contract binding to access the raw methods on
}

// IERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitCallerRaw struct {
	Contract *IERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitTransactorRaw struct {
	Contract *IERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Permit creates a new instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20Permit(address common.Address, backend bind.ContractBackend) (*IERC20Permit, error) {
	contract, err := bindIERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Permit{IERC20PermitCaller: IERC20PermitCaller{contract: contract}, IERC20PermitTransactor: IERC20PermitTransactor{contract: contract}, IERC20PermitFilterer: IERC20PermitFilterer{contract: contract}}, nil
}

// NewIERC20PermitCaller creates a new read-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitCaller, error) {
	contract, err := bindIERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitCaller{contract: contract}, nil
}

// NewIERC20PermitTransactor creates a new write-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitTransactor, error) {
	contract, err := bindIERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitTransactor{contract: contract}, nil
}

// NewIERC20PermitFilterer creates a new log filterer instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitFilterer, error) {
	contract, err := bindIERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitFilterer{contract: contract}, nil
}

// bindIERC20Permit binds a generic wrapper to an already deployed contract.
func bindIERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20PermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.IERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// IFastBridgeMetaData contains all meta data concerning the IFastBridge contract.
var IFastBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"778d31ff": "bridge((uint32,address,address,address,uint256,uint256,bool,uint256))",
		"aa9641ab": "canClaim(bytes32,address)",
		"41fcb612": "claim(bytes,address)",
		"add98c70": "dispute(bytes32)",
		"ac11fb1a": "getBridgeTransaction(bytes)",
		"886d36ff": "prove(bytes,bytes32)",
		"ef059bd3": "refund(bytes,address)",
		"8f0d6f17": "relay(bytes)",
	},
}

// IFastBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IFastBridgeMetaData.ABI instead.
var IFastBridgeABI = IFastBridgeMetaData.ABI

// Deprecated: Use IFastBridgeMetaData.Sigs instead.
// IFastBridgeFuncSigs maps the 4-byte function signature to its string representation.
var IFastBridgeFuncSigs = IFastBridgeMetaData.Sigs

// IFastBridge is an auto generated Go binding around an Ethereum contract.
type IFastBridge struct {
	IFastBridgeCaller     // Read-only binding to the contract
	IFastBridgeTransactor // Write-only binding to the contract
	IFastBridgeFilterer   // Log filterer for contract events
}

// IFastBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFastBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFastBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFastBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFastBridgeSession struct {
	Contract     *IFastBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFastBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFastBridgeCallerSession struct {
	Contract *IFastBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IFastBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFastBridgeTransactorSession struct {
	Contract     *IFastBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IFastBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFastBridgeRaw struct {
	Contract *IFastBridge // Generic contract binding to access the raw methods on
}

// IFastBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFastBridgeCallerRaw struct {
	Contract *IFastBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IFastBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFastBridgeTransactorRaw struct {
	Contract *IFastBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFastBridge creates a new instance of IFastBridge, bound to a specific deployed contract.
func NewIFastBridge(address common.Address, backend bind.ContractBackend) (*IFastBridge, error) {
	contract, err := bindIFastBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFastBridge{IFastBridgeCaller: IFastBridgeCaller{contract: contract}, IFastBridgeTransactor: IFastBridgeTransactor{contract: contract}, IFastBridgeFilterer: IFastBridgeFilterer{contract: contract}}, nil
}

// NewIFastBridgeCaller creates a new read-only instance of IFastBridge, bound to a specific deployed contract.
func NewIFastBridgeCaller(address common.Address, caller bind.ContractCaller) (*IFastBridgeCaller, error) {
	contract, err := bindIFastBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeCaller{contract: contract}, nil
}

// NewIFastBridgeTransactor creates a new write-only instance of IFastBridge, bound to a specific deployed contract.
func NewIFastBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IFastBridgeTransactor, error) {
	contract, err := bindIFastBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeTransactor{contract: contract}, nil
}

// NewIFastBridgeFilterer creates a new log filterer instance of IFastBridge, bound to a specific deployed contract.
func NewIFastBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IFastBridgeFilterer, error) {
	contract, err := bindIFastBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeFilterer{contract: contract}, nil
}

// bindIFastBridge binds a generic wrapper to an already deployed contract.
func bindIFastBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFastBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridge *IFastBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridge.Contract.IFastBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridge *IFastBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridge.Contract.IFastBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridge *IFastBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridge.Contract.IFastBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridge *IFastBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridge *IFastBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridge *IFastBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridge.Contract.contract.Transact(opts, method, params...)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridge *IFastBridgeCaller) CanClaim(opts *bind.CallOpts, transactionId [32]byte, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _IFastBridge.contract.Call(opts, &out, "canClaim", transactionId, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridge *IFastBridgeSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _IFastBridge.Contract.CanClaim(&_IFastBridge.CallOpts, transactionId, relayer)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridge *IFastBridgeCallerSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _IFastBridge.Contract.CanClaim(&_IFastBridge.CallOpts, transactionId, relayer)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridge *IFastBridgeCaller) GetBridgeTransaction(opts *bind.CallOpts, request []byte) (IFastBridgeBridgeTransaction, error) {
	var out []interface{}
	err := _IFastBridge.contract.Call(opts, &out, "getBridgeTransaction", request)

	if err != nil {
		return *new(IFastBridgeBridgeTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeBridgeTransaction)).(*IFastBridgeBridgeTransaction)

	return out0, err

}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridge *IFastBridgeSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridge.Contract.GetBridgeTransaction(&_IFastBridge.CallOpts, request)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridge *IFastBridgeCallerSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridge.Contract.GetBridgeTransaction(&_IFastBridge.CallOpts, request)
}

// Bridge is a paid mutator transaction binding the contract method 0x778d31ff.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridge *IFastBridgeTransactor) Bridge(opts *bind.TransactOpts, params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "bridge", params)
}

// Bridge is a paid mutator transaction binding the contract method 0x778d31ff.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridge *IFastBridgeSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridge.Contract.Bridge(&_IFastBridge.TransactOpts, params)
}

// Bridge is a paid mutator transaction binding the contract method 0x778d31ff.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridge *IFastBridgeTransactorSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridge.Contract.Bridge(&_IFastBridge.TransactOpts, params)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeTransactor) Claim(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "claim", request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.Contract.Claim(&_IFastBridge.TransactOpts, request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.Contract.Claim(&_IFastBridge.TransactOpts, request, to)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridge *IFastBridgeTransactor) Dispute(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "dispute", transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridge *IFastBridgeSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Dispute(&_IFastBridge.TransactOpts, transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Dispute(&_IFastBridge.TransactOpts, transactionId)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridge *IFastBridgeTransactor) Prove(opts *bind.TransactOpts, request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "prove", request, destTxHash)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridge *IFastBridgeSession) Prove(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Prove(&_IFastBridge.TransactOpts, request, destTxHash)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Prove(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Prove(&_IFastBridge.TransactOpts, request, destTxHash)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeTransactor) Refund(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "refund", request, to)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeSession) Refund(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.Contract.Refund(&_IFastBridge.TransactOpts, request, to)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Refund(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.Contract.Refund(&_IFastBridge.TransactOpts, request, to)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridge *IFastBridgeTransactor) Relay(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "relay", request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridge *IFastBridgeSession) Relay(request []byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Relay(&_IFastBridge.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridge *IFastBridgeTransactorSession) Relay(request []byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Relay(&_IFastBridge.TransactOpts, request)
}

// IFastBridgeBridgeDepositClaimedIterator is returned from FilterBridgeDepositClaimed and is used to iterate over the raw logs and unpacked data for BridgeDepositClaimed events raised by the IFastBridge contract.
type IFastBridgeBridgeDepositClaimedIterator struct {
	Event *IFastBridgeBridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeDepositClaimed)
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
		it.Event = new(IFastBridgeBridgeDepositClaimed)
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
func (it *IFastBridgeBridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeDepositClaimed represents a BridgeDepositClaimed event raised by the IFastBridge contract.
type IFastBridgeBridgeDepositClaimed struct {
	TransactionId [32]byte
	Relayer       common.Address
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositClaimed is a free log retrieval operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeDepositClaimed(opts *bind.FilterOpts) (*IFastBridgeBridgeDepositClaimedIterator, error) {

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeDepositClaimed")
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeDepositClaimedIterator{contract: _IFastBridge.contract, event: "BridgeDepositClaimed", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositClaimed is a free log subscription operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeDepositClaimed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeDepositClaimed) (event.Subscription, error) {

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeDepositClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeDepositClaimed)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
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

// ParseBridgeDepositClaimed is a log parse operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeDepositClaimed(log types.Log) (*IFastBridgeBridgeDepositClaimed, error) {
	event := new(IFastBridgeBridgeDepositClaimed)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeDepositRefundedIterator is returned from FilterBridgeDepositRefunded and is used to iterate over the raw logs and unpacked data for BridgeDepositRefunded events raised by the IFastBridge contract.
type IFastBridgeBridgeDepositRefundedIterator struct {
	Event *IFastBridgeBridgeDepositRefunded // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeDepositRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeDepositRefunded)
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
		it.Event = new(IFastBridgeBridgeDepositRefunded)
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
func (it *IFastBridgeBridgeDepositRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeDepositRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeDepositRefunded represents a BridgeDepositRefunded event raised by the IFastBridge contract.
type IFastBridgeBridgeDepositRefunded struct {
	TransactionId [32]byte
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositRefunded is a free log retrieval operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeDepositRefunded(opts *bind.FilterOpts) (*IFastBridgeBridgeDepositRefundedIterator, error) {

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeDepositRefunded")
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeDepositRefundedIterator{contract: _IFastBridge.contract, event: "BridgeDepositRefunded", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositRefunded is a free log subscription operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeDepositRefunded(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeDepositRefunded) (event.Subscription, error) {

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeDepositRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeDepositRefunded)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
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

// ParseBridgeDepositRefunded is a log parse operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeDepositRefunded(log types.Log) (*IFastBridgeBridgeDepositRefunded, error) {
	event := new(IFastBridgeBridgeDepositRefunded)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeProofDisputedIterator is returned from FilterBridgeProofDisputed and is used to iterate over the raw logs and unpacked data for BridgeProofDisputed events raised by the IFastBridge contract.
type IFastBridgeBridgeProofDisputedIterator struct {
	Event *IFastBridgeBridgeProofDisputed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeProofDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeProofDisputed)
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
		it.Event = new(IFastBridgeBridgeProofDisputed)
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
func (it *IFastBridgeBridgeProofDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeProofDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeProofDisputed represents a BridgeProofDisputed event raised by the IFastBridge contract.
type IFastBridgeBridgeProofDisputed struct {
	TransactionId [32]byte
	Relayer       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofDisputed is a free log retrieval operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeProofDisputed(opts *bind.FilterOpts) (*IFastBridgeBridgeProofDisputedIterator, error) {

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeProofDisputed")
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeProofDisputedIterator{contract: _IFastBridge.contract, event: "BridgeProofDisputed", logs: logs, sub: sub}, nil
}

// WatchBridgeProofDisputed is a free log subscription operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeProofDisputed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeProofDisputed) (event.Subscription, error) {

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeProofDisputed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeProofDisputed)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
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

// ParseBridgeProofDisputed is a log parse operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeProofDisputed(log types.Log) (*IFastBridgeBridgeProofDisputed, error) {
	event := new(IFastBridgeBridgeProofDisputed)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeProofProvidedIterator is returned from FilterBridgeProofProvided and is used to iterate over the raw logs and unpacked data for BridgeProofProvided events raised by the IFastBridge contract.
type IFastBridgeBridgeProofProvidedIterator struct {
	Event *IFastBridgeBridgeProofProvided // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeProofProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeProofProvided)
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
		it.Event = new(IFastBridgeBridgeProofProvided)
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
func (it *IFastBridgeBridgeProofProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeProofProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeProofProvided represents a BridgeProofProvided event raised by the IFastBridge contract.
type IFastBridgeBridgeProofProvided struct {
	TransactionId   [32]byte
	Relayer         common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofProvided is a free log retrieval operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeProofProvided(opts *bind.FilterOpts) (*IFastBridgeBridgeProofProvidedIterator, error) {

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeProofProvided")
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeProofProvidedIterator{contract: _IFastBridge.contract, event: "BridgeProofProvided", logs: logs, sub: sub}, nil
}

// WatchBridgeProofProvided is a free log subscription operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeProofProvided(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeProofProvided) (event.Subscription, error) {

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeProofProvided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeProofProvided)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
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

// ParseBridgeProofProvided is a log parse operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeProofProvided(log types.Log) (*IFastBridgeBridgeProofProvided, error) {
	event := new(IFastBridgeBridgeProofProvided)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeRelayedIterator is returned from FilterBridgeRelayed and is used to iterate over the raw logs and unpacked data for BridgeRelayed events raised by the IFastBridge contract.
type IFastBridgeBridgeRelayedIterator struct {
	Event *IFastBridgeBridgeRelayed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeRelayed)
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
		it.Event = new(IFastBridgeBridgeRelayed)
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
func (it *IFastBridgeBridgeRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeRelayed represents a BridgeRelayed event raised by the IFastBridge contract.
type IFastBridgeBridgeRelayed struct {
	TransactionId  [32]byte
	Relayer        common.Address
	To             common.Address
	Token          common.Address
	Amount         *big.Int
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeRelayed is a free log retrieval operation binding the contract event 0xf5d0e166883753ed3b66d60c893dda5f6760044b7bba1317907a9385c6620f9b.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount, uint256 chainGasAmount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeRelayed(opts *bind.FilterOpts) (*IFastBridgeBridgeRelayedIterator, error) {

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeRelayed")
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeRelayedIterator{contract: _IFastBridge.contract, event: "BridgeRelayed", logs: logs, sub: sub}, nil
}

// WatchBridgeRelayed is a free log subscription operation binding the contract event 0xf5d0e166883753ed3b66d60c893dda5f6760044b7bba1317907a9385c6620f9b.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount, uint256 chainGasAmount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeRelayed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeRelayed) (event.Subscription, error) {

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeRelayed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeRelayed)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
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

// ParseBridgeRelayed is a log parse operation binding the contract event 0xf5d0e166883753ed3b66d60c893dda5f6760044b7bba1317907a9385c6620f9b.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount, uint256 chainGasAmount)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeRelayed(log types.Log) (*IFastBridgeBridgeRelayed, error) {
	event := new(IFastBridgeBridgeRelayed)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeRequestedIterator is returned from FilterBridgeRequested and is used to iterate over the raw logs and unpacked data for BridgeRequested events raised by the IFastBridge contract.
type IFastBridgeBridgeRequestedIterator struct {
	Event *IFastBridgeBridgeRequested // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeRequested)
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
		it.Event = new(IFastBridgeBridgeRequested)
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
func (it *IFastBridgeBridgeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeRequested represents a BridgeRequested event raised by the IFastBridge contract.
type IFastBridgeBridgeRequested struct {
	TransactionId [32]byte
	Sender        common.Address
	Request       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRequested is a free log retrieval operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeRequested(opts *bind.FilterOpts) (*IFastBridgeBridgeRequestedIterator, error) {

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeRequested")
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeRequestedIterator{contract: _IFastBridge.contract, event: "BridgeRequested", logs: logs, sub: sub}, nil
}

// WatchBridgeRequested is a free log subscription operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeRequested(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeRequested) (event.Subscription, error) {

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeRequested)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
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

// ParseBridgeRequested is a log parse operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeRequested(log types.Log) (*IFastBridgeBridgeRequested, error) {
	event := new(IFastBridgeBridgeRequested)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedDecrease\",\"type\":\"uint256\"}],\"name\":\"SafeERC20FailedDecreaseAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208ec5d76181f59dd8c7e42c1d6e60399b6554c91ac82cfd016d7890cb3c12d3cd64736f6c63430008140033",
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
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// UniversalTokenLibMetaData contains all meta data concerning the UniversalTokenLib contract.
var UniversalTokenLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e1f79dedb4224dbf083f29b363d21bbc917393c8bde42513b43d8be8b3053a9164736f6c63430008140033",
}

// UniversalTokenLibABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalTokenLibMetaData.ABI instead.
var UniversalTokenLibABI = UniversalTokenLibMetaData.ABI

// UniversalTokenLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniversalTokenLibMetaData.Bin instead.
var UniversalTokenLibBin = UniversalTokenLibMetaData.Bin

// DeployUniversalTokenLib deploys a new Ethereum contract, binding an instance of UniversalTokenLib to it.
func DeployUniversalTokenLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniversalTokenLib, error) {
	parsed, err := UniversalTokenLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniversalTokenLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniversalTokenLib{UniversalTokenLibCaller: UniversalTokenLibCaller{contract: contract}, UniversalTokenLibTransactor: UniversalTokenLibTransactor{contract: contract}, UniversalTokenLibFilterer: UniversalTokenLibFilterer{contract: contract}}, nil
}

// UniversalTokenLib is an auto generated Go binding around an Ethereum contract.
type UniversalTokenLib struct {
	UniversalTokenLibCaller     // Read-only binding to the contract
	UniversalTokenLibTransactor // Write-only binding to the contract
	UniversalTokenLibFilterer   // Log filterer for contract events
}

// UniversalTokenLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalTokenLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalTokenLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalTokenLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalTokenLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalTokenLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalTokenLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalTokenLibSession struct {
	Contract     *UniversalTokenLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniversalTokenLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalTokenLibCallerSession struct {
	Contract *UniversalTokenLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniversalTokenLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalTokenLibTransactorSession struct {
	Contract     *UniversalTokenLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniversalTokenLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalTokenLibRaw struct {
	Contract *UniversalTokenLib // Generic contract binding to access the raw methods on
}

// UniversalTokenLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalTokenLibCallerRaw struct {
	Contract *UniversalTokenLibCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalTokenLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalTokenLibTransactorRaw struct {
	Contract *UniversalTokenLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalTokenLib creates a new instance of UniversalTokenLib, bound to a specific deployed contract.
func NewUniversalTokenLib(address common.Address, backend bind.ContractBackend) (*UniversalTokenLib, error) {
	contract, err := bindUniversalTokenLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalTokenLib{UniversalTokenLibCaller: UniversalTokenLibCaller{contract: contract}, UniversalTokenLibTransactor: UniversalTokenLibTransactor{contract: contract}, UniversalTokenLibFilterer: UniversalTokenLibFilterer{contract: contract}}, nil
}

// NewUniversalTokenLibCaller creates a new read-only instance of UniversalTokenLib, bound to a specific deployed contract.
func NewUniversalTokenLibCaller(address common.Address, caller bind.ContractCaller) (*UniversalTokenLibCaller, error) {
	contract, err := bindUniversalTokenLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalTokenLibCaller{contract: contract}, nil
}

// NewUniversalTokenLibTransactor creates a new write-only instance of UniversalTokenLib, bound to a specific deployed contract.
func NewUniversalTokenLibTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalTokenLibTransactor, error) {
	contract, err := bindUniversalTokenLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalTokenLibTransactor{contract: contract}, nil
}

// NewUniversalTokenLibFilterer creates a new log filterer instance of UniversalTokenLib, bound to a specific deployed contract.
func NewUniversalTokenLibFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalTokenLibFilterer, error) {
	contract, err := bindUniversalTokenLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalTokenLibFilterer{contract: contract}, nil
}

// bindUniversalTokenLib binds a generic wrapper to an already deployed contract.
func bindUniversalTokenLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniversalTokenLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalTokenLib *UniversalTokenLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalTokenLib.Contract.UniversalTokenLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalTokenLib *UniversalTokenLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalTokenLib.Contract.UniversalTokenLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalTokenLib *UniversalTokenLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalTokenLib.Contract.UniversalTokenLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalTokenLib *UniversalTokenLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalTokenLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalTokenLib *UniversalTokenLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalTokenLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalTokenLib *UniversalTokenLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalTokenLib.Contract.contract.Transact(opts, method, params...)
}
