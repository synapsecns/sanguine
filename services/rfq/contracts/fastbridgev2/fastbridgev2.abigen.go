// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fastbridgev2

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

// IFastBridgeBridgeParams is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeBridgeParams struct {
	DstChainId   uint32
	Sender       common.Address
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

// IFastBridgeV2BridgeParamsV2 is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeV2BridgeParamsV2 struct {
	QuoteRelayer            common.Address
	QuoteExclusivitySeconds *big.Int
	QuoteId                 []byte
	ZapNative               *big.Int
	ZapData                 []byte
}

// IFastBridgeV2BridgeTransactionV2 is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeV2BridgeTransactionV2 struct {
	OriginChainId      uint32
	DestChainId        uint32
	OriginSender       common.Address
	DestRecipient      common.Address
	OriginToken        common.Address
	DestToken          common.Address
	OriginAmount       *big.Int
	DestAmount         *big.Int
	OriginFeeAmount    *big.Int
	Deadline           *big.Int
	Nonce              *big.Int
	ExclusivityRelayer common.Address
	ExclusivityEndTime *big.Int
	ZapNative          *big.Int
	ZapData            []byte
}

// IMulticallTargetResult is an auto generated low-level Go binding around an user-defined struct.
type IMulticallTargetResult struct {
	Success    bool
	ReturnData []byte
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

// AccessControlEnumerableMetaData contains all meta data concerning the AccessControlEnumerable contract.
var AccessControlEnumerableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// AccessControlEnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlEnumerableMetaData.ABI instead.
var AccessControlEnumerableABI = AccessControlEnumerableMetaData.ABI

// Deprecated: Use AccessControlEnumerableMetaData.Sigs instead.
// AccessControlEnumerableFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlEnumerableFuncSigs = AccessControlEnumerableMetaData.Sigs

// AccessControlEnumerable is an auto generated Go binding around an Ethereum contract.
type AccessControlEnumerable struct {
	AccessControlEnumerableCaller     // Read-only binding to the contract
	AccessControlEnumerableTransactor // Write-only binding to the contract
	AccessControlEnumerableFilterer   // Log filterer for contract events
}

// AccessControlEnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlEnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlEnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlEnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlEnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlEnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlEnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlEnumerableSession struct {
	Contract     *AccessControlEnumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlEnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlEnumerableCallerSession struct {
	Contract *AccessControlEnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// AccessControlEnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlEnumerableTransactorSession struct {
	Contract     *AccessControlEnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AccessControlEnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlEnumerableRaw struct {
	Contract *AccessControlEnumerable // Generic contract binding to access the raw methods on
}

// AccessControlEnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlEnumerableCallerRaw struct {
	Contract *AccessControlEnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlEnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlEnumerableTransactorRaw struct {
	Contract *AccessControlEnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlEnumerable creates a new instance of AccessControlEnumerable, bound to a specific deployed contract.
func NewAccessControlEnumerable(address common.Address, backend bind.ContractBackend) (*AccessControlEnumerable, error) {
	contract, err := bindAccessControlEnumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerable{AccessControlEnumerableCaller: AccessControlEnumerableCaller{contract: contract}, AccessControlEnumerableTransactor: AccessControlEnumerableTransactor{contract: contract}, AccessControlEnumerableFilterer: AccessControlEnumerableFilterer{contract: contract}}, nil
}

// NewAccessControlEnumerableCaller creates a new read-only instance of AccessControlEnumerable, bound to a specific deployed contract.
func NewAccessControlEnumerableCaller(address common.Address, caller bind.ContractCaller) (*AccessControlEnumerableCaller, error) {
	contract, err := bindAccessControlEnumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableCaller{contract: contract}, nil
}

// NewAccessControlEnumerableTransactor creates a new write-only instance of AccessControlEnumerable, bound to a specific deployed contract.
func NewAccessControlEnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlEnumerableTransactor, error) {
	contract, err := bindAccessControlEnumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableTransactor{contract: contract}, nil
}

// NewAccessControlEnumerableFilterer creates a new log filterer instance of AccessControlEnumerable, bound to a specific deployed contract.
func NewAccessControlEnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlEnumerableFilterer, error) {
	contract, err := bindAccessControlEnumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableFilterer{contract: contract}, nil
}

// bindAccessControlEnumerable binds a generic wrapper to an already deployed contract.
func bindAccessControlEnumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlEnumerableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlEnumerable *AccessControlEnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlEnumerable.Contract.AccessControlEnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlEnumerable *AccessControlEnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.AccessControlEnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlEnumerable *AccessControlEnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.AccessControlEnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlEnumerable *AccessControlEnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlEnumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlEnumerable *AccessControlEnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlEnumerable *AccessControlEnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlEnumerable.Contract.DEFAULTADMINROLE(&_AccessControlEnumerable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlEnumerable.Contract.DEFAULTADMINROLE(&_AccessControlEnumerable.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlEnumerable.Contract.GetRoleAdmin(&_AccessControlEnumerable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlEnumerable.Contract.GetRoleAdmin(&_AccessControlEnumerable.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlEnumerable *AccessControlEnumerableSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControlEnumerable.Contract.GetRoleMember(&_AccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControlEnumerable.Contract.GetRoleMember(&_AccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlEnumerable *AccessControlEnumerableSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControlEnumerable.Contract.GetRoleMemberCount(&_AccessControlEnumerable.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControlEnumerable.Contract.GetRoleMemberCount(&_AccessControlEnumerable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlEnumerable.Contract.HasRole(&_AccessControlEnumerable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlEnumerable.Contract.HasRole(&_AccessControlEnumerable.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControlEnumerable.Contract.SupportsInterface(&_AccessControlEnumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControlEnumerable.Contract.SupportsInterface(&_AccessControlEnumerable.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.GrantRole(&_AccessControlEnumerable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.GrantRole(&_AccessControlEnumerable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlEnumerable *AccessControlEnumerableSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.RenounceRole(&_AccessControlEnumerable.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.RenounceRole(&_AccessControlEnumerable.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.RevokeRole(&_AccessControlEnumerable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.RevokeRole(&_AccessControlEnumerable.TransactOpts, role, account)
}

// AccessControlEnumerableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleAdminChangedIterator struct {
	Event *AccessControlEnumerableRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlEnumerableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlEnumerableRoleAdminChanged)
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
		it.Event = new(AccessControlEnumerableRoleAdminChanged)
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
func (it *AccessControlEnumerableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlEnumerableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlEnumerableRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlEnumerableRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AccessControlEnumerable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableRoleAdminChangedIterator{contract: _AccessControlEnumerable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlEnumerableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlEnumerable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlEnumerableRoleAdminChanged)
				if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlEnumerableRoleAdminChanged, error) {
	event := new(AccessControlEnumerableRoleAdminChanged)
	if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlEnumerableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleGrantedIterator struct {
	Event *AccessControlEnumerableRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlEnumerableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlEnumerableRoleGranted)
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
		it.Event = new(AccessControlEnumerableRoleGranted)
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
func (it *AccessControlEnumerableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlEnumerableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlEnumerableRoleGranted represents a RoleGranted event raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlEnumerableRoleGrantedIterator, error) {

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

	logs, sub, err := _AccessControlEnumerable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableRoleGrantedIterator{contract: _AccessControlEnumerable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlEnumerableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlEnumerable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlEnumerableRoleGranted)
				if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) ParseRoleGranted(log types.Log) (*AccessControlEnumerableRoleGranted, error) {
	event := new(AccessControlEnumerableRoleGranted)
	if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlEnumerableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleRevokedIterator struct {
	Event *AccessControlEnumerableRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlEnumerableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlEnumerableRoleRevoked)
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
		it.Event = new(AccessControlEnumerableRoleRevoked)
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
func (it *AccessControlEnumerableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlEnumerableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlEnumerableRoleRevoked represents a RoleRevoked event raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlEnumerableRoleRevokedIterator, error) {

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

	logs, sub, err := _AccessControlEnumerable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableRoleRevokedIterator{contract: _AccessControlEnumerable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlEnumerableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlEnumerable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlEnumerableRoleRevoked)
				if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) ParseRoleRevoked(log types.Log) (*AccessControlEnumerableRoleRevoked, error) {
	event := new(AccessControlEnumerableRoleRevoked)
	if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bf1b9d7d74d58f093eea1b088fdbcad7f67374661fae1ab5815d7a31130f2fb864736f6c63430008180033",
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

// AdminMetaData contains all meta data concerning the Admin contract.
var AdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldChainGasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_RATE_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUNDER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"bf333f2c": "FEE_BPS()",
		"0f5f6ed7": "FEE_RATE_MAX()",
		"ccc57490": "GOVERNOR_ROLE()",
		"03ed0ee5": "GUARD_ROLE()",
		"5960ccf2": "REFUNDER_ROLE()",
		"926d7d7f": "RELAYER_ROLE()",
		"e00a83e0": "chainGasAmount()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"58f85880": "protocolFeeRate()",
		"dcf844a7": "protocolFees(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"b250fe6b": "setChainGasAmount(uint256)",
		"b13aa2d6": "setProtocolFeeRate(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"06f333f2": "sweepProtocolFees(address,address)",
	},
	Bin: "0x60806040523480156200001157600080fd5b50604051620014123803806200141283398101604081905262000034916200018e565b6200004160008262000049565b5050620001b9565b60008062000058848462000086565b905080156200007d5760008481526001602052604090206200007b908462000134565b505b90505b92915050565b6000828152602081815260408083206001600160a01b038516845290915281205460ff166200012b576000838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055620000e23390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600162000080565b50600062000080565b60006200007d836001600160a01b03841660008181526001830160205260408120546200012b5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000080565b600060208284031215620001a157600080fd5b81516001600160a01b03811681146200007d57600080fd5b61124980620001c96000396000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806391d14854116100d8578063bf333f2c1161008c578063d547741f11610066578063d547741f14610385578063dcf844a714610398578063e00a83e0146103b857600080fd5b8063bf333f2c14610341578063ca15c8731461034b578063ccc574901461035e57600080fd5b8063a217fddf116100bd578063a217fddf14610313578063b13aa2d61461031b578063b250fe6b1461032e57600080fd5b806391d14854146102a8578063926d7d7f146102ec57600080fd5b80632f2ff15d1161012f57806358f858801161011457806358f85880146102405780635960ccf2146102495780639010d07c1461027057600080fd5b80632f2ff15d1461021a57806336568abe1461022d57600080fd5b806306f333f21161016057806306f333f2146101d95780630f5f6ed7146101ee578063248a9ca3146101f757600080fd5b806301ffc9a71461017c57806303ed0ee5146101a4575b600080fd5b61018f61018a366004611013565b6103c1565b60405190151581526020015b60405180910390f35b6101cb7f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d81565b60405190815260200161019b565b6101ec6101e736600461107e565b61041d565b005b6101cb61271081565b6101cb6102053660046110b1565b60009081526020819052604090206001015490565b6101ec6102283660046110ca565b61050b565b6101ec61023b3660046110ca565b610536565b6101cb60025481565b6101cb7fdb9556138406326f00296e13ea2ad7db24ba82381212d816b1a40c23b466b32781565b61028361027e3660046110ed565b61058f565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b61018f6102b63660046110ca565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101cb7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc481565b6101cb600081565b6101ec6103293660046110b1565b6105ae565b6101ec61033c3660046110b1565b610690565b6101cb620f424081565b6101cb6103593660046110b1565b6106f8565b6101cb7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b6101ec6103933660046110ca565b61070f565b6101cb6103a636600461110f565b60036020526000908152604090205481565b6101cb60045481565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f000000000000000000000000000000000000000000000000000000001480610417575061041782610734565b92915050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610447816107cb565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260408120549081900361047b5750505050565b73ffffffffffffffffffffffffffffffffffffffff84166000818152600360205260408120556104ac9084836107d8565b6040805173ffffffffffffffffffffffffffffffffffffffff8087168252851660208201529081018290527f244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd9060600160405180910390a1505b505050565b600082815260208190526040902060010154610526816107cb565b610530838361092f565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610585576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105068282610964565b60008281526001602052604081206105a79083610991565b9392505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f556105d8816107cb565b612710821115610649576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f6e657746656552617465203e206d61780000000000000000000000000000000060448201526064015b60405180910390fd5b600280549083905560408051828152602081018590527f14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb95791015b60405180910390a1505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f556106ba816107cb565b600480549083905560408051828152602081018590527f5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa9101610683565b60008181526001602052604081206104179061099d565b60008281526020819052604090206001015461072a816107cb565b6105308383610964565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061041757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610417565b6107d581336109a7565b50565b3073ffffffffffffffffffffffffffffffffffffffff8316036107fa57505050565b8060000361080757505050565b7fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff84160161090e5760008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d806000811461089e576040519150601f19603f3d011682016040523d82523d6000602084013e6108a3565b606091505b5050905080610530576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401610640565b61050673ffffffffffffffffffffffffffffffffffffffff84168383610a31565b60008061093c8484610abe565b905080156105a757600084815260016020526040902061095c9084610bba565b509392505050565b6000806109718484610bdc565b905080156105a757600084815260016020526040902061095c9084610c97565b60006105a78383610cb9565b6000610417825490565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610a2d576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610640565b5050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610506908490610ce3565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610bb25760008381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610b503390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610417565b506000610417565b60006105a78373ffffffffffffffffffffffffffffffffffffffff8416610d79565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610bb25760008381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610417565b60006105a78373ffffffffffffffffffffffffffffffffffffffff8416610dc0565b6000826000018281548110610cd057610cd061112a565b9060005260206000200154905092915050565b6000610d0573ffffffffffffffffffffffffffffffffffffffff841683610eb3565b90508051600014158015610d2a575080806020019051810190610d289190611159565b155b15610506576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610640565b6000818152600183016020526040812054610bb257508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610417565b60008181526001830160205260408120548015610ea9576000610de460018361117b565b8554909150600090610df89060019061117b565b9050808214610e5d576000866000018281548110610e1857610e1861112a565b9060005260206000200154905080876000018481548110610e3b57610e3b61112a565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610e6e57610e6e6111b5565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610417565b6000915050610417565b60606105a783836000846000808573ffffffffffffffffffffffffffffffffffffffff168486604051610ee691906111e4565b60006040518083038185875af1925050503d8060008114610f23576040519150601f19603f3d011682016040523d82523d6000602084013e610f28565b606091505b5091509150610f38868383610f42565b9695505050505050565b606082610f5757610f5282610fd1565b6105a7565b8151158015610f7b575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610fca576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610640565b50806105a7565b805115610fe15780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561102557600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146105a757600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461107957600080fd5b919050565b6000806040838503121561109157600080fd5b61109a83611055565b91506110a860208401611055565b90509250929050565b6000602082840312156110c357600080fd5b5035919050565b600080604083850312156110dd57600080fd5b823591506110a860208401611055565b6000806040838503121561110057600080fd5b50508035926020909101359150565b60006020828403121561112157600080fd5b6105a782611055565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561116b57600080fd5b815180151581146105a757600080fd5b81810381811115610417577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000825160005b8181101561120557602081860181015185830152016111eb565b50600092019182525091905056fea2646970667358221220cf33c3d4da69e4a6aeddfcb46cd394bfc44899f75312d31fd8c8e38dff96ab9764736f6c63430008180033",
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
	parsed, err := AdminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// REFUNDERROLE is a free data retrieval call binding the contract method 0x5960ccf2.
//
// Solidity: function REFUNDER_ROLE() view returns(bytes32)
func (_Admin *AdminCaller) REFUNDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "REFUNDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REFUNDERROLE is a free data retrieval call binding the contract method 0x5960ccf2.
//
// Solidity: function REFUNDER_ROLE() view returns(bytes32)
func (_Admin *AdminSession) REFUNDERROLE() ([32]byte, error) {
	return _Admin.Contract.REFUNDERROLE(&_Admin.CallOpts)
}

// REFUNDERROLE is a free data retrieval call binding the contract method 0x5960ccf2.
//
// Solidity: function REFUNDER_ROLE() view returns(bytes32)
func (_Admin *AdminCallerSession) REFUNDERROLE() ([32]byte, error) {
	return _Admin.Contract.REFUNDERROLE(&_Admin.CallOpts)
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

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Admin *AdminCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Admin *AdminSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Admin.Contract.GetRoleMember(&_Admin.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Admin *AdminCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Admin.Contract.GetRoleMember(&_Admin.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Admin *AdminCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Admin *AdminSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Admin.Contract.GetRoleMemberCount(&_Admin.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Admin *AdminCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Admin.Contract.GetRoleMemberCount(&_Admin.CallOpts, role)
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

// BridgeTransactionV2LibMetaData contains all meta data concerning the BridgeTransactionV2Lib contract.
var BridgeTransactionV2LibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BridgeTransactionV2__InvalidEncodedTx\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"BridgeTransactionV2__UnsupportedVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c0d1e2256fb16b1e19ad4848dc80aafe3c3f70b204e139f256b5d6095b52dd5564736f6c63430008180033",
}

// BridgeTransactionV2LibABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTransactionV2LibMetaData.ABI instead.
var BridgeTransactionV2LibABI = BridgeTransactionV2LibMetaData.ABI

// BridgeTransactionV2LibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeTransactionV2LibMetaData.Bin instead.
var BridgeTransactionV2LibBin = BridgeTransactionV2LibMetaData.Bin

// DeployBridgeTransactionV2Lib deploys a new Ethereum contract, binding an instance of BridgeTransactionV2Lib to it.
func DeployBridgeTransactionV2Lib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeTransactionV2Lib, error) {
	parsed, err := BridgeTransactionV2LibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTransactionV2LibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeTransactionV2Lib{BridgeTransactionV2LibCaller: BridgeTransactionV2LibCaller{contract: contract}, BridgeTransactionV2LibTransactor: BridgeTransactionV2LibTransactor{contract: contract}, BridgeTransactionV2LibFilterer: BridgeTransactionV2LibFilterer{contract: contract}}, nil
}

// BridgeTransactionV2Lib is an auto generated Go binding around an Ethereum contract.
type BridgeTransactionV2Lib struct {
	BridgeTransactionV2LibCaller     // Read-only binding to the contract
	BridgeTransactionV2LibTransactor // Write-only binding to the contract
	BridgeTransactionV2LibFilterer   // Log filterer for contract events
}

// BridgeTransactionV2LibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTransactionV2LibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactionV2LibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactionV2LibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactionV2LibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTransactionV2LibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactionV2LibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTransactionV2LibSession struct {
	Contract     *BridgeTransactionV2Lib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeTransactionV2LibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTransactionV2LibCallerSession struct {
	Contract *BridgeTransactionV2LibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BridgeTransactionV2LibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactionV2LibTransactorSession struct {
	Contract     *BridgeTransactionV2LibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BridgeTransactionV2LibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTransactionV2LibRaw struct {
	Contract *BridgeTransactionV2Lib // Generic contract binding to access the raw methods on
}

// BridgeTransactionV2LibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTransactionV2LibCallerRaw struct {
	Contract *BridgeTransactionV2LibCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactionV2LibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactionV2LibTransactorRaw struct {
	Contract *BridgeTransactionV2LibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTransactionV2Lib creates a new instance of BridgeTransactionV2Lib, bound to a specific deployed contract.
func NewBridgeTransactionV2Lib(address common.Address, backend bind.ContractBackend) (*BridgeTransactionV2Lib, error) {
	contract, err := bindBridgeTransactionV2Lib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactionV2Lib{BridgeTransactionV2LibCaller: BridgeTransactionV2LibCaller{contract: contract}, BridgeTransactionV2LibTransactor: BridgeTransactionV2LibTransactor{contract: contract}, BridgeTransactionV2LibFilterer: BridgeTransactionV2LibFilterer{contract: contract}}, nil
}

// NewBridgeTransactionV2LibCaller creates a new read-only instance of BridgeTransactionV2Lib, bound to a specific deployed contract.
func NewBridgeTransactionV2LibCaller(address common.Address, caller bind.ContractCaller) (*BridgeTransactionV2LibCaller, error) {
	contract, err := bindBridgeTransactionV2Lib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactionV2LibCaller{contract: contract}, nil
}

// NewBridgeTransactionV2LibTransactor creates a new write-only instance of BridgeTransactionV2Lib, bound to a specific deployed contract.
func NewBridgeTransactionV2LibTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactionV2LibTransactor, error) {
	contract, err := bindBridgeTransactionV2Lib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactionV2LibTransactor{contract: contract}, nil
}

// NewBridgeTransactionV2LibFilterer creates a new log filterer instance of BridgeTransactionV2Lib, bound to a specific deployed contract.
func NewBridgeTransactionV2LibFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTransactionV2LibFilterer, error) {
	contract, err := bindBridgeTransactionV2Lib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactionV2LibFilterer{contract: contract}, nil
}

// bindBridgeTransactionV2Lib binds a generic wrapper to an already deployed contract.
func bindBridgeTransactionV2Lib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeTransactionV2LibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTransactionV2Lib.Contract.BridgeTransactionV2LibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransactionV2Lib.Contract.BridgeTransactionV2LibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransactionV2Lib.Contract.BridgeTransactionV2LibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTransactionV2Lib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransactionV2Lib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransactionV2Lib.Contract.contract.Transact(opts, method, params...)
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
	parsed, err := ERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202a0689afaebf6ffa0f4d0a6fe04b9b6e642f8ef83c8db74e665e9fc77d71999964736f6c63430008180033",
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

// FastBridgeV2MetaData contains all meta data concerning the FastBridgeV2 contract.
var FastBridgeV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeTransactionV2__InvalidEncodedTx\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"BridgeTransactionV2__UnsupportedVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineNotExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExclusivityParamsIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExclusivityPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MulticallTarget__UndeterminedRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecipientIncorrectReturnValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecipientNoReturnValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StatusIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZapDataLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZapNativeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"}],\"name\":\"BridgeQuoteDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldChainGasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_RATE_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ZAP_DATA_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEADLINE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_GAS_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUNDER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUND_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"quoteRelayer\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"quoteExclusivitySeconds\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeParamsV2\",\"name\":\"paramsV2\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeProofs\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"timestamp\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelayDetails\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"blockNumber\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"blockTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelays\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeStatuses\",\"outputs\":[{\"internalType\":\"enumIFastBridgeV2.BridgeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeTxDetails\",\"outputs\":[{\"internalType\":\"enumIFastBridgeV2.BridgeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint56\",\"name\":\"proofBlockTimestamp\",\"type\":\"uint56\"},{\"internalType\":\"address\",\"name\":\"proofRelayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransactionV2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"exclusivityRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exclusivityEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeTransactionV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"ignoreReverts\",\"type\":\"bool\"}],\"name\":\"multicallNoResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"ignoreReverts\",\"type\":\"bool\"}],\"name\":\"multicallWithResults\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticallTarget.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"senderNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"a5bbe22b": "DISPUTE_PERIOD()",
		"bf333f2c": "FEE_BPS()",
		"0f5f6ed7": "FEE_RATE_MAX()",
		"ccc57490": "GOVERNOR_ROLE()",
		"03ed0ee5": "GUARD_ROLE()",
		"54eff068": "MAX_ZAP_DATA_LENGTH()",
		"820688d5": "MIN_DEADLINE_PERIOD()",
		"0f862f1e": "NATIVE_GAS_TOKEN()",
		"5960ccf2": "REFUNDER_ROLE()",
		"190da595": "REFUND_DELAY()",
		"926d7d7f": "RELAYER_ROLE()",
		"45851694": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256))",
		"bfc7c607": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256),(address,int256,bytes,uint256,bytes))",
		"91ad5039": "bridgeProofs(bytes32)",
		"c79371b1": "bridgeRelayDetails(bytes32)",
		"8379a24f": "bridgeRelays(bytes32)",
		"051287bc": "bridgeStatuses(bytes32)",
		"63787e52": "bridgeTxDetails(bytes32)",
		"aa9641ab": "canClaim(bytes32,address)",
		"e00a83e0": "chainGasAmount()",
		"c63ff8dd": "claim(bytes)",
		"41fcb612": "claim(bytes,address)",
		"a3ec191a": "deployBlock()",
		"add98c70": "dispute(bytes32)",
		"ac11fb1a": "getBridgeTransaction(bytes)",
		"5aa6ccba": "getBridgeTransactionV2(bytes)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"3f61331d": "multicallNoResults(bytes[],bool)",
		"385c1d2f": "multicallWithResults(bytes[],bool)",
		"affed0e0": "nonce()",
		"58f85880": "protocolFeeRate()",
		"dcf844a7": "protocolFees(address)",
		"886d36ff": "prove(bytes,bytes32)",
		"18e4357d": "prove(bytes32,bytes32,address)",
		"5eb7d946": "refund(bytes)",
		"8f0d6f17": "relay(bytes)",
		"9c9545f0": "relay(bytes,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"295710ff": "senderNonces(address)",
		"b250fe6b": "setChainGasAmount(uint256)",
		"b13aa2d6": "setProtocolFeeRate(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"06f333f2": "sweepProtocolFees(address,address)",
	},
	Bin: "0x60c060405260006080523480156200001657600080fd5b506040516200475c3803806200475c833981016040819052620000399162000199565b806200004760008262000054565b50504360a05250620001c4565b60008062000063848462000091565b90508015620000885760008481526001602052604090206200008690846200013f565b505b90505b92915050565b6000828152602081815260408083206001600160a01b038516845290915281205460ff1662000136576000838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055620000ed3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016200008b565b5060006200008b565b600062000088836001600160a01b038416600081815260018301602052604081205462000136575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200008b565b600060208284031215620001ac57600080fd5b81516001600160a01b03811681146200008857600080fd5b60805160a051614572620001ea60003960006108b30152600061095401526145726000f3fe60806040526004361061031e5760003560e01c8063886d36ff116101a5578063add98c70116100ec578063c63ff8dd11610095578063ccc574901161006f578063ccc5749014610aad578063d547741f14610ae1578063dcf844a714610b01578063e00a83e014610b2e57600080fd5b8063c63ff8dd146109e0578063c79371b114610a00578063ca15c87314610a8d57600080fd5b8063b250fe6b116100c6578063b250fe6b14610996578063bf333f2c146109b6578063bfc7c607146109cd57600080fd5b8063add98c7014610922578063affed0e014610942578063b13aa2d61461097657600080fd5b80639c9545f01161014e578063a5bbe22b11610128578063a5bbe22b146106ca578063aa9641ab146108d5578063ac11fb1a146108f557600080fd5b80639c9545f014610879578063a217fddf1461088c578063a3ec191a146108a157600080fd5b806391ad50391161017f57806391ad50391461077b57806391d1485414610801578063926d7d7f1461084557600080fd5b8063886d36ff146107285780638f0d6f17146107485780639010d07c1461075b57600080fd5b8063385c1d2f116102695780635960ccf21161021257806363787e52116101ec57806363787e5214610650578063820688d5146106ca5780638379a24f146106e057600080fd5b80635960ccf2146105cf5780635aa6ccba146106035780635eb7d9461461063057600080fd5b80634585169411610243578063458516941461059057806354eff068146105a357806358f85880146105b957600080fd5b8063385c1d2f146105235780633f61331d1461055057806341fcb6121461057057600080fd5b806318e4357d116102cb578063295710ff116102a5578063295710ff146104b65780632f2ff15d146104e357806336568abe1461050357600080fd5b806318e4357d1461044f578063190da5951461046f578063248a9ca31461048657600080fd5b806306f333f2116102fc57806306f333f2146103d75780630f5f6ed7146103f95780630f862f1e1461040f57600080fd5b806301ffc9a71461032357806303ed0ee514610358578063051287bc1461039a575b600080fd5b34801561032f57600080fd5b5061034361033e36600461351e565b610b44565b60405190151581526020015b60405180910390f35b34801561036457600080fd5b5061038c7f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d81565b60405190815260200161034f565b3480156103a657600080fd5b506103ca6103b5366004613560565b60009081526005602052604090205460ff1690565b60405161034f91906135e3565b3480156103e357600080fd5b506103f76103f2366004613616565b610ba0565b005b34801561040557600080fd5b5061038c61271081565b34801561041b57600080fd5b5061043773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b6040516001600160a01b03909116815260200161034f565b34801561045b57600080fd5b506103f761046a36600461364f565b610c67565b34801561047b57600080fd5b5061038c62093a8081565b34801561049257600080fd5b5061038c6104a1366004613560565b60009081526020819052604090206001015490565b3480156104c257600080fd5b5061038c6104d1366004613688565b60076020526000908152604090205481565b3480156104ef57600080fd5b506103f76104fe3660046136a5565b610d7c565b34801561050f57600080fd5b506103f761051e3660046136a5565b610da7565b34801561052f57600080fd5b5061054361053e3660046136e3565b610df3565b60405161034f91906137ae565b34801561055c57600080fd5b506103f761056b3660046136e3565b610f81565b34801561057c57600080fd5b506103f761058b36600461388d565b61102b565b6103f761059e366004613a8e565b61129d565b3480156105af57600080fd5b5061038c61ffff81565b3480156105c557600080fd5b5061038c60025481565b3480156105db57600080fd5b5061038c7fdb9556138406326f00296e13ea2ad7db24ba82381212d816b1a40c23b466b32781565b34801561060f57600080fd5b5061062361061e366004613aab565b6112fa565b60405161034f9190613aed565b34801561063c57600080fd5b506103f761064b366004613aab565b6113c7565b34801561065c57600080fd5b506106ba61066b366004613560565b60056020526000908152604090205460ff811690610100810463ffffffff169065010000000000810466ffffffffffffff16906c0100000000000000000000000090046001600160a01b031684565b60405161034f9493929190613c0a565b3480156106d657600080fd5b5061038c61070881565b3480156106ec57600080fd5b506103436106fb366004613560565b6000908152600660205260409020546c0100000000000000000000000090046001600160a01b0316151590565b34801561073457600080fd5b506103f7610743366004613c4b565b6115c9565b6103f7610756366004613aab565b6115f5565b34801561076757600080fd5b50610437610776366004613c97565b611604565b34801561078757600080fd5b506107d5610796366004613560565b60009081526005602052604090205465010000000000810466ffffffffffffff16916c010000000000000000000000009091046001600160a01b031690565b604080516bffffffffffffffffffffffff90931683526001600160a01b0390911660208301520161034f565b34801561080d57600080fd5b5061034361081c3660046136a5565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561085157600080fd5b5061038c7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc481565b6103f761088736600461388d565b61161c565b34801561089857600080fd5b5061038c600081565b3480156108ad57600080fd5b5061038c7f000000000000000000000000000000000000000000000000000000000000000081565b3480156108e157600080fd5b506103436108f03660046136a5565b6118d6565b34801561090157600080fd5b50610915610910366004613aab565b6119ad565b60405161034f9190613cb9565b34801561092e57600080fd5b506103f761093d366004613560565b611b61565b34801561094e57600080fd5b5061038c7f000000000000000000000000000000000000000000000000000000000000000081565b34801561098257600080fd5b506103f7610991366004613560565b611cab565b3480156109a257600080fd5b506103f76109b1366004613560565b611d8d565b3480156109c257600080fd5b5061038c620f424081565b6103f76109db366004613e1d565b611df5565b3480156109ec57600080fd5b506103f76109fb366004613aab565b612084565b348015610a0c57600080fd5b50610a5f610a1b366004613560565b60066020526000908152604090205465ffffffffffff8082169166010000000000008104909116906c0100000000000000000000000090046001600160a01b031683565b6040805165ffffffffffff94851681529390921660208401526001600160a01b03169082015260600161034f565b348015610a9957600080fd5b5061038c610aa8366004613560565b612090565b348015610ab957600080fd5b5061038c7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b348015610aed57600080fd5b506103f7610afc3660046136a5565b6120a7565b348015610b0d57600080fd5b5061038c610b1c366004613688565b60036020526000908152604090205481565b348015610b3a57600080fd5b5061038c60045481565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f000000000000000000000000000000000000000000000000000000001480610b9a5750610b9a826120cc565b92915050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610bca81612163565b6001600160a01b03831660009081526003602052604081205490819003610bf15750505050565b6001600160a01b038416600081815260036020526040812055610c1590848361216d565b604080516001600160a01b038087168252851660208201529081018290527f244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd9060600160405180910390a1505b505050565b7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4610c9181612163565b60008481526005602052604090206001815460ff166004811115610cb757610cb7613579565b14610cee576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80546001600160a01b0384166c0100000000000000000000000081026bffffffffffffffffffffffff66ffffffffffffff421665010000000000021664ffffffff009093169290921791909117600217825560405185815286907f4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e49060200160405180910390a35050505050565b600082815260208190526040902060010154610d9781612163565b610da18383612290565b50505050565b6001600160a01b0381163314610de9576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c6282826122bd565b60608267ffffffffffffffff811115610e0e57610e0e6138d9565b604051908082528060200260200182016040528015610e5457816020015b604080518082019091526000815260606020820152815260200190600190039081610e2c5790505b50905060005b83811015610f795730858583818110610e7557610e75613eeb565b9050602002810190610e879190613f1a565b604051610e95929190613f7f565b600060405180830381855af49150503d8060008114610ed0576040519150601f19603f3d011682016040523d82523d6000602084013e610ed5565b606091505b50838381518110610ee857610ee8613eeb565b6020026020010151600001848481518110610f0557610f05613eeb565b602002602001015160200182905282151515158152505050818181518110610f2f57610f2f613eeb565b602002602001015160000151158015610f46575082155b15610f7157610f71828281518110610f6057610f60613eeb565b6020026020010151602001516122ea565b600101610e5a565b509392505050565b60005b82811015610da15760008030868685818110610fa257610fa2613eeb565b9050602002810190610fb49190613f1a565b604051610fc2929190613f7f565b600060405180830381855af49150503d8060008114610ffd576040519150601f19603f3d011682016040523d82523d6000602084013e611002565b606091505b509150915081158015611013575083155b1561102157611021816122ea565b5050600101610f84565b611035838361232c565b60008383604051611047929190613f7f565b604080519182900390912060008181526005602052919091208054919250906c0100000000000000000000000081046001600160a01b03169060ff81169065010000000000900466ffffffffffffff1660028260048111156110ab576110ab613579565b146110e2576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107084282900366ffffffffffffff1611611129576040517f1992d0bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03861661113f57829550611181565b6001600160a01b0383163314611181576040517f4af43a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166003178455603288013560601c605a890135609a8a013580156111f0576001600160a01b038316600090815260036020526040812080548392906111ea908490613fbe565b90915550505b604080516001600160a01b03858116825260208201859052808c1692908916918b917f582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678910160405180910390a47fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b0384160161127c5761127789836123ad565b611290565b6112906001600160a01b0384168a84612476565b5050505050505050505050565b6112f7816040518060a0016040528060006001600160a01b03168152602001600081526020016040518060200160405280600081525081526020016000815260200160405180602001604052806000815250815250611df5565b50565b6113ac604051806101e00160405280600063ffffffff168152602001600063ffffffff16815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016000815260200160008152602001606081525090565b6113b6838361232c565b6113c083836124ea565b9392505050565b6113d1828261232c565b600082826040516113e3929190613f7f565b604080519182900390912060008181526005602052919091209091506001815460ff16600481111561141757611417613579565b1461144e576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081527fd2043bf65931af3dbecf60d0db8f40e4160406d7beb00522f4200cf4944a1eb8602052604090205460ba8501359060ff1661149a5761149762093a8082613fbe565b90505b8042116114d3576040517fe15ff9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166004178255600a850135606090811c906032870135901c6000611522609a890135605a8a0135613fbe565b604080516001600160a01b03858116825260208201849052825193945086169289927fb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958928290030190a37fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b038316016115ab576115a683826123ad565b6115bf565b6115bf6001600160a01b0383168483612476565b5050505050505050565b6115d3838361232c565b610c6283836040516115e6929190613f7f565b60405180910390208233610c67565b61160082823361161c565b5050565b60008281526001602052604081206113c0908361268c565b611626838361232c565b60008383604051611638929190613f7f565b6040518091039020905061164e84848385612698565b6040805160608101825265ffffffffffff438116825242811660208084019182526001600160a01b038088168587019081526000888152600690935295822094518554935196519091166c01000000000000000000000000026bffffffffffffffffffffffff9685166601000000000000027fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009094169190941617919091179390931617905561170285601e013560601c90565b9050604685013560601c607a86013561012e8701356001600160a01b03808516908716867ff8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c60028c013560e01c6040805163ffffffff92909216825260328e0135606090811c60208401526001600160a01b038a1683830152605a8f0135908301526080820188905260a08201879052519081900360c00190a47fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03841601611842578015611804576040517f846dedc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81341461183d576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611890565b80341461187b576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118906001600160a01b03841633868561282d565b36600061189d8a8a612866565b909250905080156118ba576118b58686868585612882565b6118ca565b34156118ca576118ca86346123ad565b50505050505050505050565b60008281526005602052604081206002815460ff1660048111156118fc576118fc613579565b14611933576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80546001600160a01b038481166c010000000000000000000000009092041614611989576040517f4af43a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b546107086501000000000090910466ffffffffffffff908116420316119392505050565b6040805161018081018252600080825260208201819052818301819052606082018190526080820181905260a0820181905260c0820181905260e0820181905261010082018190526101208201819052610140820181905261016082015290517f5aa6ccba0000000000000000000000000000000000000000000000000000000081523090635aa6ccba90611a489086908690600401613ffc565b600060405180830381865afa925050508015611a8657506040513d6000823e601f3d908101601f19168201604052611a83919081019061406b565b60015b611a9d57611a96828401846141a0565b9050610b9a565b604051806101800160405280826000015163ffffffff168152602001826020015163ffffffff16815260200182604001516001600160a01b0316815260200182606001516001600160a01b0316815260200182608001516001600160a01b031681526020018260a001516001600160a01b031681526020018260c0015181526020018260e0015181526020018261010001518152602001826101a0015160001415151581526020018261012001518152602001826101400151815250915050610b9a565b7f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d611b8b81612163565b600082815260056020526040902080546c0100000000000000000000000081046001600160a01b03169060ff81169065010000000000900466ffffffffffffff166002826004811115611be057611be0613579565b14611c17576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107084282900366ffffffffffffff161115611c5f576040517f3e908aac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835464ffffffff001660011784556040516001600160a01b0384169087907f0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad490600090a3505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55611cd581612163565b612710821115611d46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f6e657746656552617465203e206d61780000000000000000000000000000000060448201526064015b60405180910390fd5b600280549083905560408051828152602081018590527f14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb95791015b60405180910390a1505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55611db781612163565b600480549083905560408051828152602081018590527f5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa9101611d80565b80516000906001600160a01b031615611e1a576020820151611e17904261426c565b90505b611e258383836129de565b6000611e3984606001518560a00151612c62565b90506000806002541115611e7257620f424060025483611e599190614294565b611e6391906142ab565b9050611e6f81836142e6565b91505b6000611f7b604051806101e001604052804663ffffffff168152602001886000015163ffffffff16815260200188602001516001600160a01b0316815260200188604001516001600160a01b0316815260200188606001516001600160a01b0316815260200188608001516001600160a01b031681526020018581526020018860c0015181526020018481526020018861010001518152602001600760008a602001516001600160a01b03166001600160a01b031681526020019081526020016000206000815480929190611f46906142f9565b90915550815287516001600160a01b031660208201526040810187905260608089015190820152608080890151910152612e75565b805160208083019190912060008181526005835260409081902080548b5163ffffffff8116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000090921691909117600117909155928a01516060808c015160808d015160c08e0151928d0151945197985094966001600160a01b039093169587957f120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a95612037958b959394938e92909190151590614331565b60405180910390a3807f3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e87604001516040516120739190614387565b60405180910390a250505050505050565b6116008282600061102b565b6000818152600160205260408120610b9a90612fba565b6000828152602081905260409020600101546120c281612163565b610da183836122bd565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610b9a57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610b9a565b6112f78133612fc4565b306001600160a01b0383160361218257505050565b8060000361218f57505050565b7fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b0384160161227c576000826001600160a01b03168260405160006040518083038185875af1925050503d806000811461220c576040519150601f19603f3d011682016040523d82523d6000602084013e612211565b606091505b5050905080610da1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c6564000000000000000000000000006044820152606401611d3d565b610c626001600160a01b0384168383612476565b60008061229d8484613030565b905080156113c0576000848152600160205260409020610f7990846130f8565b6000806122ca848461310d565b905080156113c0576000848152600160205260409020610f7990846131ae565b8051156122fa5780518082602001fd5b6040517f5ead5a9d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61014e811015612368576040517f6ee0e17100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b813560f01c60028114610c62576040517f8bd977e700000000000000000000000000000000000000000000000000000000815261ffff82166004820152602401611d3d565b804710156123e9576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611d3d565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114612436576040519150601f19603f3d011682016040523d82523d6000602084013e61243b565b606091505b5050905080610c62576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516001600160a01b03838116602483015260448201839052610c6291859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506131c3565b61259c604051806101e00160405280600063ffffffff168152602001600063ffffffff16815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016000815260200160008152602001606081525090565b600283013560e090811c82526006840135811c6020830152600a840135606090811c6040840152601e850135811c818401526032850135811c60808401526046850135811c60a0840152605a85013560c0840152607a85013591830191909152609a84013561010083015260ba84013561012083015260da84013561014083015260fa840135901c61016082015261010e83013561018082015261012e8301356101a082015261264c8383612866565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506101c082015292915050565b60006113c0838361323f565b6001600160a01b0381166126d8576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152600660205260409020546c0100000000000000000000000090046001600160a01b031615612737576040517fbef7bb7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600684013560e01c4614612777576040517f7029fdf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60ba8401354211156127b5576040517f559895a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60fa84013560601c80158015906127de5750816001600160a01b0316816001600160a01b031614155b80156127ef575061010e8501354211155b15612826576040517f14be123200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6040516001600160a01b038481166024830152838116604483015260648201839052610da19186918216906323b872dd906084016124a3565b3660006128778361014e818761439a565b915091509250929050565b60006128fd868686868660405160240161289f94939291906143c4565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe85e13dd0000000000000000000000000000000000000000000000000000000017905234613269565b9050805160000361293a576040517f1a95ad2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051602014612975576040517ff3725cca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe85e13dd0000000000000000000000000000000000000000000000000000000061299f826143ed565b146129d6576040517ff3725cca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050565b46836000015163ffffffff1603612a21576040517f7029fdf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08301511580612a34575060c0830151155b15612a6b576040517fe38820c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208301516001600160a01b03161580612a90575060408301516001600160a01b0316155b15612ac7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608301516001600160a01b03161580612aec575060808301516001600160a01b0316155b15612b23576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612b2f61070842613fbe565b8361010001511015612b6d576040517f04b7fcc800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff8260800151511115612bae576040517f177a70e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082015115801590612be1575060808301516001600160a01b031673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee145b15612c18576040517f846dedc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000811280612c2b575082610100015181135b15610c62576040517f352807b900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03841601612cd457348214612ccd576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5034610b9a565b3415612d0c576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b03163b600003612d50576040517f7f523fe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038416906370a0823190602401602060405180830381865afa158015612dad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dd19190614432565b9050612de86001600160a01b03841633308561282d565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281906001600160a01b038516906370a0823190602401602060405180830381865afa158015612e47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e6b9190614432565b6113c091906142e6565b8051602080830151604080850151606086810151608088015160a089015160c08a015195517e02000000000000000000000000000000000000000000000000000000000000988101989098527fffffffff0000000000000000000000000000000000000000000000000000000060e0998a1b811660228a01529690981b90951660268701527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000092821b8316602a870152811b8216603e86015292831b8116605285015293821b9093166066830152607a820192909252600090609a0160408051601f198184030181529082905260e08501516101008601516101208701516101408801516101608901516101808a01516101a08b01516101c08c0151979950612fa3988a989060200161444b565b604051602081830303815290604052915050919050565b6000610b9a825490565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16611600576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611d3d565b6000828152602081815260408083206001600160a01b038516845290915281205460ff166130f0576000838152602081815260408083206001600160a01b0386168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556130a83390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610b9a565b506000610b9a565b60006113c0836001600160a01b03841661331f565b6000828152602081815260408083206001600160a01b038516845290915281205460ff16156130f0576000838152602081815260408083206001600160a01b038616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610b9a565b60006113c0836001600160a01b038416613366565b60006131d86001600160a01b03841683613459565b905080516000141580156131fd5750808060200190518101906131fb91906144d4565b155b15610c62576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611d3d565b600082600001828154811061325657613256613eeb565b9060005260206000200154905092915050565b6060814710156132a7576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611d3d565b600080856001600160a01b031684866040516132c391906144f1565b60006040518083038185875af1925050503d8060008114613300576040519150601f19603f3d011682016040523d82523d6000602084013e613305565b606091505b5091509150613315868383613467565b9695505050505050565b60008181526001830160205260408120546130f057508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610b9a565b6000818152600183016020526040812054801561344f57600061338a6001836142e6565b855490915060009061339e906001906142e6565b90508082146134035760008660000182815481106133be576133be613eeb565b90600052602060002001549050808760000184815481106133e1576133e1613eeb565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806134145761341461450d565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610b9a565b6000915050610b9a565b60606113c083836000613269565b60608261347c57613477826134dc565b6113c0565b815115801561349357506001600160a01b0384163b155b156134d5576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611d3d565b50806113c0565b8051156134ec5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561353057600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113c057600080fd5b60006020828403121561357257600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600581106135df577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60208101610b9a82846135a8565b6001600160a01b03811681146112f757600080fd5b8035613611816135f1565b919050565b6000806040838503121561362957600080fd5b8235613634816135f1565b91506020830135613644816135f1565b809150509250929050565b60008060006060848603121561366457600080fd5b8335925060208401359150604084013561367d816135f1565b809150509250925092565b60006020828403121561369a57600080fd5b81356113c0816135f1565b600080604083850312156136b857600080fd5b823591506020830135613644816135f1565b80151581146112f757600080fd5b8035613611816136ca565b6000806000604084860312156136f857600080fd5b833567ffffffffffffffff8082111561371057600080fd5b818601915086601f83011261372457600080fd5b81358181111561373357600080fd5b8760208260051b850101111561374857600080fd5b6020928301955093505084013561367d816136ca565b60005b83811015613779578181015183820152602001613761565b50506000910152565b6000815180845261379a81602086016020860161375e565b601f01601f19169290920160200192915050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b83811015613836578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180511515845287015187840187905261382387850182613782565b95880195935050908601906001016137d7565b509098975050505050505050565b60008083601f84011261385657600080fd5b50813567ffffffffffffffff81111561386e57600080fd5b60208301915083602082850101111561388657600080fd5b9250929050565b6000806000604084860312156138a257600080fd5b833567ffffffffffffffff8111156138b957600080fd5b6138c586828701613844565b909450925050602084013561367d816135f1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff8111828210171561392c5761392c6138d9565b60405290565b60405160a0810167ffffffffffffffff8111828210171561392c5761392c6138d9565b6040516101e0810167ffffffffffffffff8111828210171561392c5761392c6138d9565b604051610180810167ffffffffffffffff8111828210171561392c5761392c6138d9565b604051601f8201601f1916810167ffffffffffffffff811182821017156139c6576139c66138d9565b604052919050565b63ffffffff811681146112f757600080fd5b8035613611816139ce565b600061012082840312156139fe57600080fd5b613a06613908565b9050613a11826139e0565b8152613a1f60208301613606565b6020820152613a3060408301613606565b6040820152613a4160608301613606565b6060820152613a5260808301613606565b608082015260a082013560a082015260c082013560c0820152613a7760e083016136d8565b60e082015261010080830135818301525092915050565b60006101208284031215613aa157600080fd5b6113c083836139eb565b60008060208385031215613abe57600080fd5b823567ffffffffffffffff811115613ad557600080fd5b613ae185828601613844565b90969095509350505050565b60208152613b0460208201835163ffffffff169052565b60006020830151613b1d604084018263ffffffff169052565b5060408301516001600160a01b03811660608401525060608301516001600160a01b03811660808401525060808301516001600160a01b03811660a08401525060a08301516001600160a01b03811660c08401525060c083015160e08381019190915283015161010080840191909152830151610120808401919091528301516101408084019190915283015161016080840191909152830151610180613bce818501836001600160a01b03169052565b8401516101a0848101919091528401516101c0808501919091528401516101e0808501529050613c02610200840182613782565b949350505050565b60808101613c1882876135a8565b63ffffffff8516602083015266ffffffffffffff841660408301526001600160a01b038316606083015295945050505050565b600080600060408486031215613c6057600080fd5b833567ffffffffffffffff811115613c7757600080fd5b613c8386828701613844565b909790965060209590950135949350505050565b60008060408385031215613caa57600080fd5b50508035926020909101359150565b815163ffffffff16815261018081016020830151613cdf602084018263ffffffff169052565b506040830151613cfa60408401826001600160a01b03169052565b506060830151613d1560608401826001600160a01b03169052565b506080830151613d3060808401826001600160a01b03169052565b5060a0830151613d4b60a08401826001600160a01b03169052565b5060c083015160c083015260e083015160e083015261010080840151818401525061012080840151613d808285018215159052565b5050610140838101519083015261016092830151929091019190915290565b600067ffffffffffffffff821115613db957613db96138d9565b50601f01601f191660200190565b600082601f830112613dd857600080fd5b8135613deb613de682613d9f565b61399d565b818152846020838601011115613e0057600080fd5b816020850160208301376000918101602001919091529392505050565b6000806101408385031215613e3157600080fd5b613e3b84846139eb565b915061012083013567ffffffffffffffff80821115613e5957600080fd5b9084019060a08287031215613e6d57600080fd5b613e75613932565b8235613e80816135f1565b815260208381013590820152604083013582811115613e9e57600080fd5b613eaa88828601613dc7565b60408301525060608301356060820152608083013582811115613ecc57600080fd5b613ed888828601613dc7565b6080830152508093505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613f4f57600080fd5b83018035915067ffffffffffffffff821115613f6a57600080fd5b60200191503681900382131561388657600080fd5b8183823760009101908152919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610b9a57610b9a613f8f565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b602081526000613c02602083018486613fd1565b8051613611816139ce565b8051613611816135f1565b600082601f83011261403757600080fd5b8151614045613de682613d9f565b81815284602083860101111561405a57600080fd5b613c0282602083016020870161375e565b60006020828403121561407d57600080fd5b815167ffffffffffffffff8082111561409557600080fd5b908301906101e082860312156140aa57600080fd5b6140b2613955565b6140bb83614010565b81526140c960208401614010565b60208201526140da6040840161401b565b60408201526140eb6060840161401b565b60608201526140fc6080840161401b565b608082015261410d60a0840161401b565b60a082015260c0838101519082015260e0808401519082015261010080840151908201526101208084015190820152610140808401519082015261016061415581850161401b565b9082015261018083810151908201526101a080840151908201526101c0808401518381111561418357600080fd5b61418f88828701614026565b918301919091525095945050505050565b600061018082840312156141b357600080fd5b6141bb613979565b6141c4836139e0565b81526141d2602084016139e0565b60208201526141e360408401613606565b60408201526141f460608401613606565b606082015261420560808401613606565b608082015261421660a08401613606565b60a082015260c083013560c082015260e083013560e08201526101008084013581830152506101206142498185016136d8565b908201526101408381013590820152610160928301359281019290925250919050565b808201828112600083128015821682158216171561428c5761428c613f8f565b505092915050565b8082028115828204841417610b9a57610b9a613f8f565b6000826142e1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b81810381811115610b9a57610b9a613f8f565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361432a5761432a613f8f565b5060010190565b60e08152600061434460e083018a613782565b63ffffffff989098166020830152506001600160a01b039586166040820152939094166060840152608083019190915260a082015290151560c090910152919050565b6020815260006113c06020830184613782565b600080858511156143aa57600080fd5b838611156143b757600080fd5b5050820193919092039150565b6001600160a01b0385168152836020820152606060408201526000613315606083018486613fd1565b8051602080830151919081101561442c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8160200360031b1b821691505b50919050565b60006020828403121561444457600080fd5b5051919050565b60008a5161445d818460208f0161375e565b80830190508a81528960208201528860408201528760608201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008760601b1660808201528560948201528460b482015283516144c18160d484016020880161375e565b0160d4019b9a5050505050505050505050565b6000602082840312156144e657600080fd5b81516113c0816136ca565b6000825161450381846020870161375e565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212209a48ddc6013fa0ddcb3b8a930e8c7e75c40dab4d351f72ddd07978c3a1539f2764736f6c63430008180033",
}

// FastBridgeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FastBridgeV2MetaData.ABI instead.
var FastBridgeV2ABI = FastBridgeV2MetaData.ABI

// Deprecated: Use FastBridgeV2MetaData.Sigs instead.
// FastBridgeV2FuncSigs maps the 4-byte function signature to its string representation.
var FastBridgeV2FuncSigs = FastBridgeV2MetaData.Sigs

// FastBridgeV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FastBridgeV2MetaData.Bin instead.
var FastBridgeV2Bin = FastBridgeV2MetaData.Bin

// DeployFastBridgeV2 deploys a new Ethereum contract, binding an instance of FastBridgeV2 to it.
func DeployFastBridgeV2(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *FastBridgeV2, error) {
	parsed, err := FastBridgeV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FastBridgeV2Bin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FastBridgeV2{FastBridgeV2Caller: FastBridgeV2Caller{contract: contract}, FastBridgeV2Transactor: FastBridgeV2Transactor{contract: contract}, FastBridgeV2Filterer: FastBridgeV2Filterer{contract: contract}}, nil
}

// FastBridgeV2 is an auto generated Go binding around an Ethereum contract.
type FastBridgeV2 struct {
	FastBridgeV2Caller     // Read-only binding to the contract
	FastBridgeV2Transactor // Write-only binding to the contract
	FastBridgeV2Filterer   // Log filterer for contract events
}

// FastBridgeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FastBridgeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FastBridgeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastBridgeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastBridgeV2Session struct {
	Contract     *FastBridgeV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FastBridgeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastBridgeV2CallerSession struct {
	Contract *FastBridgeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FastBridgeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastBridgeV2TransactorSession struct {
	Contract     *FastBridgeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FastBridgeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FastBridgeV2Raw struct {
	Contract *FastBridgeV2 // Generic contract binding to access the raw methods on
}

// FastBridgeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastBridgeV2CallerRaw struct {
	Contract *FastBridgeV2Caller // Generic read-only contract binding to access the raw methods on
}

// FastBridgeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastBridgeV2TransactorRaw struct {
	Contract *FastBridgeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFastBridgeV2 creates a new instance of FastBridgeV2, bound to a specific deployed contract.
func NewFastBridgeV2(address common.Address, backend bind.ContractBackend) (*FastBridgeV2, error) {
	contract, err := bindFastBridgeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2{FastBridgeV2Caller: FastBridgeV2Caller{contract: contract}, FastBridgeV2Transactor: FastBridgeV2Transactor{contract: contract}, FastBridgeV2Filterer: FastBridgeV2Filterer{contract: contract}}, nil
}

// NewFastBridgeV2Caller creates a new read-only instance of FastBridgeV2, bound to a specific deployed contract.
func NewFastBridgeV2Caller(address common.Address, caller bind.ContractCaller) (*FastBridgeV2Caller, error) {
	contract, err := bindFastBridgeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2Caller{contract: contract}, nil
}

// NewFastBridgeV2Transactor creates a new write-only instance of FastBridgeV2, bound to a specific deployed contract.
func NewFastBridgeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FastBridgeV2Transactor, error) {
	contract, err := bindFastBridgeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2Transactor{contract: contract}, nil
}

// NewFastBridgeV2Filterer creates a new log filterer instance of FastBridgeV2, bound to a specific deployed contract.
func NewFastBridgeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FastBridgeV2Filterer, error) {
	contract, err := bindFastBridgeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2Filterer{contract: contract}, nil
}

// bindFastBridgeV2 binds a generic wrapper to an already deployed contract.
func bindFastBridgeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastBridgeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastBridgeV2 *FastBridgeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastBridgeV2.Contract.FastBridgeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastBridgeV2 *FastBridgeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.FastBridgeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastBridgeV2 *FastBridgeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.FastBridgeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastBridgeV2 *FastBridgeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastBridgeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastBridgeV2 *FastBridgeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastBridgeV2 *FastBridgeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.DEFAULTADMINROLE(&_FastBridgeV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.DEFAULTADMINROLE(&_FastBridgeV2.CallOpts)
}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) DISPUTEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "DISPUTE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) DISPUTEPERIOD() (*big.Int, error) {
	return _FastBridgeV2.Contract.DISPUTEPERIOD(&_FastBridgeV2.CallOpts)
}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) DISPUTEPERIOD() (*big.Int, error) {
	return _FastBridgeV2.Contract.DISPUTEPERIOD(&_FastBridgeV2.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) FEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) FEEBPS() (*big.Int, error) {
	return _FastBridgeV2.Contract.FEEBPS(&_FastBridgeV2.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) FEEBPS() (*big.Int, error) {
	return _FastBridgeV2.Contract.FEEBPS(&_FastBridgeV2.CallOpts)
}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) FEERATEMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "FEE_RATE_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) FEERATEMAX() (*big.Int, error) {
	return _FastBridgeV2.Contract.FEERATEMAX(&_FastBridgeV2.CallOpts)
}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) FEERATEMAX() (*big.Int, error) {
	return _FastBridgeV2.Contract.FEERATEMAX(&_FastBridgeV2.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) GOVERNORROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.GOVERNORROLE(&_FastBridgeV2.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) GOVERNORROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.GOVERNORROLE(&_FastBridgeV2.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) GUARDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "GUARD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) GUARDROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.GUARDROLE(&_FastBridgeV2.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) GUARDROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.GUARDROLE(&_FastBridgeV2.CallOpts)
}

// MAXZAPDATALENGTH is a free data retrieval call binding the contract method 0x54eff068.
//
// Solidity: function MAX_ZAP_DATA_LENGTH() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) MAXZAPDATALENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "MAX_ZAP_DATA_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXZAPDATALENGTH is a free data retrieval call binding the contract method 0x54eff068.
//
// Solidity: function MAX_ZAP_DATA_LENGTH() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) MAXZAPDATALENGTH() (*big.Int, error) {
	return _FastBridgeV2.Contract.MAXZAPDATALENGTH(&_FastBridgeV2.CallOpts)
}

// MAXZAPDATALENGTH is a free data retrieval call binding the contract method 0x54eff068.
//
// Solidity: function MAX_ZAP_DATA_LENGTH() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) MAXZAPDATALENGTH() (*big.Int, error) {
	return _FastBridgeV2.Contract.MAXZAPDATALENGTH(&_FastBridgeV2.CallOpts)
}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) MINDEADLINEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "MIN_DEADLINE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) MINDEADLINEPERIOD() (*big.Int, error) {
	return _FastBridgeV2.Contract.MINDEADLINEPERIOD(&_FastBridgeV2.CallOpts)
}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) MINDEADLINEPERIOD() (*big.Int, error) {
	return _FastBridgeV2.Contract.MINDEADLINEPERIOD(&_FastBridgeV2.CallOpts)
}

// NATIVEGASTOKEN is a free data retrieval call binding the contract method 0x0f862f1e.
//
// Solidity: function NATIVE_GAS_TOKEN() view returns(address)
func (_FastBridgeV2 *FastBridgeV2Caller) NATIVEGASTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "NATIVE_GAS_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEGASTOKEN is a free data retrieval call binding the contract method 0x0f862f1e.
//
// Solidity: function NATIVE_GAS_TOKEN() view returns(address)
func (_FastBridgeV2 *FastBridgeV2Session) NATIVEGASTOKEN() (common.Address, error) {
	return _FastBridgeV2.Contract.NATIVEGASTOKEN(&_FastBridgeV2.CallOpts)
}

// NATIVEGASTOKEN is a free data retrieval call binding the contract method 0x0f862f1e.
//
// Solidity: function NATIVE_GAS_TOKEN() view returns(address)
func (_FastBridgeV2 *FastBridgeV2CallerSession) NATIVEGASTOKEN() (common.Address, error) {
	return _FastBridgeV2.Contract.NATIVEGASTOKEN(&_FastBridgeV2.CallOpts)
}

// REFUNDERROLE is a free data retrieval call binding the contract method 0x5960ccf2.
//
// Solidity: function REFUNDER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) REFUNDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "REFUNDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REFUNDERROLE is a free data retrieval call binding the contract method 0x5960ccf2.
//
// Solidity: function REFUNDER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) REFUNDERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.REFUNDERROLE(&_FastBridgeV2.CallOpts)
}

// REFUNDERROLE is a free data retrieval call binding the contract method 0x5960ccf2.
//
// Solidity: function REFUNDER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) REFUNDERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.REFUNDERROLE(&_FastBridgeV2.CallOpts)
}

// REFUNDDELAY is a free data retrieval call binding the contract method 0x190da595.
//
// Solidity: function REFUND_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) REFUNDDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "REFUND_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REFUNDDELAY is a free data retrieval call binding the contract method 0x190da595.
//
// Solidity: function REFUND_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) REFUNDDELAY() (*big.Int, error) {
	return _FastBridgeV2.Contract.REFUNDDELAY(&_FastBridgeV2.CallOpts)
}

// REFUNDDELAY is a free data retrieval call binding the contract method 0x190da595.
//
// Solidity: function REFUND_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) REFUNDDELAY() (*big.Int, error) {
	return _FastBridgeV2.Contract.REFUNDDELAY(&_FastBridgeV2.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) RELAYERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.RELAYERROLE(&_FastBridgeV2.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) RELAYERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.RELAYERROLE(&_FastBridgeV2.CallOpts)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_FastBridgeV2 *FastBridgeV2Caller) BridgeProofs(opts *bind.CallOpts, transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "bridgeProofs", transactionId)

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
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_FastBridgeV2 *FastBridgeV2Session) BridgeProofs(transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _FastBridgeV2.Contract.BridgeProofs(&_FastBridgeV2.CallOpts, transactionId)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_FastBridgeV2 *FastBridgeV2CallerSession) BridgeProofs(transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _FastBridgeV2.Contract.BridgeProofs(&_FastBridgeV2.CallOpts, transactionId)
}

// BridgeRelayDetails is a free data retrieval call binding the contract method 0xc79371b1.
//
// Solidity: function bridgeRelayDetails(bytes32 ) view returns(uint48 blockNumber, uint48 blockTimestamp, address relayer)
func (_FastBridgeV2 *FastBridgeV2Caller) BridgeRelayDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	BlockNumber    *big.Int
	BlockTimestamp *big.Int
	Relayer        common.Address
}, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "bridgeRelayDetails", arg0)

	outstruct := new(struct {
		BlockNumber    *big.Int
		BlockTimestamp *big.Int
		Relayer        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Relayer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BridgeRelayDetails is a free data retrieval call binding the contract method 0xc79371b1.
//
// Solidity: function bridgeRelayDetails(bytes32 ) view returns(uint48 blockNumber, uint48 blockTimestamp, address relayer)
func (_FastBridgeV2 *FastBridgeV2Session) BridgeRelayDetails(arg0 [32]byte) (struct {
	BlockNumber    *big.Int
	BlockTimestamp *big.Int
	Relayer        common.Address
}, error) {
	return _FastBridgeV2.Contract.BridgeRelayDetails(&_FastBridgeV2.CallOpts, arg0)
}

// BridgeRelayDetails is a free data retrieval call binding the contract method 0xc79371b1.
//
// Solidity: function bridgeRelayDetails(bytes32 ) view returns(uint48 blockNumber, uint48 blockTimestamp, address relayer)
func (_FastBridgeV2 *FastBridgeV2CallerSession) BridgeRelayDetails(arg0 [32]byte) (struct {
	BlockNumber    *big.Int
	BlockTimestamp *big.Int
	Relayer        common.Address
}, error) {
	return _FastBridgeV2.Contract.BridgeRelayDetails(&_FastBridgeV2.CallOpts, arg0)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2Caller) BridgeRelays(opts *bind.CallOpts, transactionId [32]byte) (bool, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "bridgeRelays", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2Session) BridgeRelays(transactionId [32]byte) (bool, error) {
	return _FastBridgeV2.Contract.BridgeRelays(&_FastBridgeV2.CallOpts, transactionId)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2CallerSession) BridgeRelays(transactionId [32]byte) (bool, error) {
	return _FastBridgeV2.Contract.BridgeRelays(&_FastBridgeV2.CallOpts, transactionId)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8 status)
func (_FastBridgeV2 *FastBridgeV2Caller) BridgeStatuses(opts *bind.CallOpts, transactionId [32]byte) (uint8, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "bridgeStatuses", transactionId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8 status)
func (_FastBridgeV2 *FastBridgeV2Session) BridgeStatuses(transactionId [32]byte) (uint8, error) {
	return _FastBridgeV2.Contract.BridgeStatuses(&_FastBridgeV2.CallOpts, transactionId)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8 status)
func (_FastBridgeV2 *FastBridgeV2CallerSession) BridgeStatuses(transactionId [32]byte) (uint8, error) {
	return _FastBridgeV2.Contract.BridgeStatuses(&_FastBridgeV2.CallOpts, transactionId)
}

// BridgeTxDetails is a free data retrieval call binding the contract method 0x63787e52.
//
// Solidity: function bridgeTxDetails(bytes32 ) view returns(uint8 status, uint32 destChainId, uint56 proofBlockTimestamp, address proofRelayer)
func (_FastBridgeV2 *FastBridgeV2Caller) BridgeTxDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status              uint8
	DestChainId         uint32
	ProofBlockTimestamp *big.Int
	ProofRelayer        common.Address
}, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "bridgeTxDetails", arg0)

	outstruct := new(struct {
		Status              uint8
		DestChainId         uint32
		ProofBlockTimestamp *big.Int
		ProofRelayer        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.DestChainId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ProofBlockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ProofRelayer = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BridgeTxDetails is a free data retrieval call binding the contract method 0x63787e52.
//
// Solidity: function bridgeTxDetails(bytes32 ) view returns(uint8 status, uint32 destChainId, uint56 proofBlockTimestamp, address proofRelayer)
func (_FastBridgeV2 *FastBridgeV2Session) BridgeTxDetails(arg0 [32]byte) (struct {
	Status              uint8
	DestChainId         uint32
	ProofBlockTimestamp *big.Int
	ProofRelayer        common.Address
}, error) {
	return _FastBridgeV2.Contract.BridgeTxDetails(&_FastBridgeV2.CallOpts, arg0)
}

// BridgeTxDetails is a free data retrieval call binding the contract method 0x63787e52.
//
// Solidity: function bridgeTxDetails(bytes32 ) view returns(uint8 status, uint32 destChainId, uint56 proofBlockTimestamp, address proofRelayer)
func (_FastBridgeV2 *FastBridgeV2CallerSession) BridgeTxDetails(arg0 [32]byte) (struct {
	Status              uint8
	DestChainId         uint32
	ProofBlockTimestamp *big.Int
	ProofRelayer        common.Address
}, error) {
	return _FastBridgeV2.Contract.BridgeTxDetails(&_FastBridgeV2.CallOpts, arg0)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2Caller) CanClaim(opts *bind.CallOpts, transactionId [32]byte, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "canClaim", transactionId, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2Session) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _FastBridgeV2.Contract.CanClaim(&_FastBridgeV2.CallOpts, transactionId, relayer)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2CallerSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _FastBridgeV2.Contract.CanClaim(&_FastBridgeV2.CallOpts, transactionId, relayer)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) ChainGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "chainGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) ChainGasAmount() (*big.Int, error) {
	return _FastBridgeV2.Contract.ChainGasAmount(&_FastBridgeV2.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) ChainGasAmount() (*big.Int, error) {
	return _FastBridgeV2.Contract.ChainGasAmount(&_FastBridgeV2.CallOpts)
}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) DeployBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "deployBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) DeployBlock() (*big.Int, error) {
	return _FastBridgeV2.Contract.DeployBlock(&_FastBridgeV2.CallOpts)
}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) DeployBlock() (*big.Int, error) {
	return _FastBridgeV2.Contract.DeployBlock(&_FastBridgeV2.CallOpts)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_FastBridgeV2 *FastBridgeV2Caller) GetBridgeTransaction(opts *bind.CallOpts, request []byte) (IFastBridgeBridgeTransaction, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "getBridgeTransaction", request)

	if err != nil {
		return *new(IFastBridgeBridgeTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeBridgeTransaction)).(*IFastBridgeBridgeTransaction)

	return out0, err

}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_FastBridgeV2 *FastBridgeV2Session) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _FastBridgeV2.Contract.GetBridgeTransaction(&_FastBridgeV2.CallOpts, request)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_FastBridgeV2 *FastBridgeV2CallerSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _FastBridgeV2.Contract.GetBridgeTransaction(&_FastBridgeV2.CallOpts, request)
}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_FastBridgeV2 *FastBridgeV2Caller) GetBridgeTransactionV2(opts *bind.CallOpts, request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "getBridgeTransactionV2", request)

	if err != nil {
		return *new(IFastBridgeV2BridgeTransactionV2), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeV2BridgeTransactionV2)).(*IFastBridgeV2BridgeTransactionV2)

	return out0, err

}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_FastBridgeV2 *FastBridgeV2Session) GetBridgeTransactionV2(request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	return _FastBridgeV2.Contract.GetBridgeTransactionV2(&_FastBridgeV2.CallOpts, request)
}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_FastBridgeV2 *FastBridgeV2CallerSession) GetBridgeTransactionV2(request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	return _FastBridgeV2.Contract.GetBridgeTransactionV2(&_FastBridgeV2.CallOpts, request)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FastBridgeV2.Contract.GetRoleAdmin(&_FastBridgeV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FastBridgeV2.Contract.GetRoleAdmin(&_FastBridgeV2.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FastBridgeV2 *FastBridgeV2Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FastBridgeV2 *FastBridgeV2Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _FastBridgeV2.Contract.GetRoleMember(&_FastBridgeV2.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_FastBridgeV2 *FastBridgeV2CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _FastBridgeV2.Contract.GetRoleMember(&_FastBridgeV2.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _FastBridgeV2.Contract.GetRoleMemberCount(&_FastBridgeV2.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _FastBridgeV2.Contract.GetRoleMemberCount(&_FastBridgeV2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FastBridgeV2.Contract.HasRole(&_FastBridgeV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FastBridgeV2.Contract.HasRole(&_FastBridgeV2.CallOpts, role, account)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) Nonce() (*big.Int, error) {
	return _FastBridgeV2.Contract.Nonce(&_FastBridgeV2.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) Nonce() (*big.Int, error) {
	return _FastBridgeV2.Contract.Nonce(&_FastBridgeV2.CallOpts)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) ProtocolFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "protocolFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) ProtocolFeeRate() (*big.Int, error) {
	return _FastBridgeV2.Contract.ProtocolFeeRate(&_FastBridgeV2.CallOpts)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) ProtocolFeeRate() (*big.Int, error) {
	return _FastBridgeV2.Contract.ProtocolFeeRate(&_FastBridgeV2.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) ProtocolFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "protocolFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) ProtocolFees(arg0 common.Address) (*big.Int, error) {
	return _FastBridgeV2.Contract.ProtocolFees(&_FastBridgeV2.CallOpts, arg0)
}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) ProtocolFees(arg0 common.Address) (*big.Int, error) {
	return _FastBridgeV2.Contract.ProtocolFees(&_FastBridgeV2.CallOpts, arg0)
}

// SenderNonces is a free data retrieval call binding the contract method 0x295710ff.
//
// Solidity: function senderNonces(address ) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) SenderNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "senderNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SenderNonces is a free data retrieval call binding the contract method 0x295710ff.
//
// Solidity: function senderNonces(address ) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) SenderNonces(arg0 common.Address) (*big.Int, error) {
	return _FastBridgeV2.Contract.SenderNonces(&_FastBridgeV2.CallOpts, arg0)
}

// SenderNonces is a free data retrieval call binding the contract method 0x295710ff.
//
// Solidity: function senderNonces(address ) view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) SenderNonces(arg0 common.Address) (*big.Int, error) {
	return _FastBridgeV2.Contract.SenderNonces(&_FastBridgeV2.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FastBridgeV2.Contract.SupportsInterface(&_FastBridgeV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridgeV2 *FastBridgeV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FastBridgeV2.Contract.SupportsInterface(&_FastBridgeV2.CallOpts, interfaceId)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Bridge(opts *bind.TransactOpts, params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "bridge", params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_FastBridgeV2 *FastBridgeV2Session) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Bridge(&_FastBridgeV2.TransactOpts, params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Bridge(&_FastBridgeV2.TransactOpts, params)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Bridge0(opts *bind.TransactOpts, params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "bridge0", params, paramsV2)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_FastBridgeV2 *FastBridgeV2Session) Bridge0(params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Bridge0(&_FastBridgeV2.TransactOpts, params, paramsV2)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Bridge0(params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Bridge0(&_FastBridgeV2.TransactOpts, params, paramsV2)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Claim(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "claim", request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridgeV2 *FastBridgeV2Session) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Claim(&_FastBridgeV2.TransactOpts, request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Claim(&_FastBridgeV2.TransactOpts, request, to)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Claim0(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "claim0", request)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2Session) Claim0(request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Claim0(&_FastBridgeV2.TransactOpts, request)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Claim0(request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Claim0(&_FastBridgeV2.TransactOpts, request)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Dispute(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "dispute", transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridgeV2 *FastBridgeV2Session) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Dispute(&_FastBridgeV2.TransactOpts, transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Dispute(&_FastBridgeV2.TransactOpts, transactionId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridgeV2 *FastBridgeV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.GrantRole(&_FastBridgeV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.GrantRole(&_FastBridgeV2.TransactOpts, role, account)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) MulticallNoResults(opts *bind.TransactOpts, data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "multicallNoResults", data, ignoreReverts)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_FastBridgeV2 *FastBridgeV2Session) MulticallNoResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.MulticallNoResults(&_FastBridgeV2.TransactOpts, data, ignoreReverts)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) MulticallNoResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.MulticallNoResults(&_FastBridgeV2.TransactOpts, data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_FastBridgeV2 *FastBridgeV2Transactor) MulticallWithResults(opts *bind.TransactOpts, data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "multicallWithResults", data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_FastBridgeV2 *FastBridgeV2Session) MulticallWithResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.MulticallWithResults(&_FastBridgeV2.TransactOpts, data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_FastBridgeV2 *FastBridgeV2TransactorSession) MulticallWithResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.MulticallWithResults(&_FastBridgeV2.TransactOpts, data, ignoreReverts)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Prove(opts *bind.TransactOpts, transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "prove", transactionId, destTxHash, relayer)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_FastBridgeV2 *FastBridgeV2Session) Prove(transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Prove(&_FastBridgeV2.TransactOpts, transactionId, destTxHash, relayer)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Prove(transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Prove(&_FastBridgeV2.TransactOpts, transactionId, destTxHash, relayer)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Prove0(opts *bind.TransactOpts, request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "prove0", request, destTxHash)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridgeV2 *FastBridgeV2Session) Prove0(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Prove0(&_FastBridgeV2.TransactOpts, request, destTxHash)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Prove0(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Prove0(&_FastBridgeV2.TransactOpts, request, destTxHash)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Refund(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "refund", request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2Session) Refund(request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Refund(&_FastBridgeV2.TransactOpts, request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Refund(request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Refund(&_FastBridgeV2.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Relay(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "relay", request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridgeV2 *FastBridgeV2Session) Relay(request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Relay(&_FastBridgeV2.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Relay(request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Relay(&_FastBridgeV2.TransactOpts, request)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Relay0(opts *bind.TransactOpts, request []byte, relayer common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "relay0", request, relayer)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_FastBridgeV2 *FastBridgeV2Session) Relay0(request []byte, relayer common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Relay0(&_FastBridgeV2.TransactOpts, request, relayer)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Relay0(request []byte, relayer common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Relay0(&_FastBridgeV2.TransactOpts, request, relayer)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridgeV2 *FastBridgeV2Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.RenounceRole(&_FastBridgeV2.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.RenounceRole(&_FastBridgeV2.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridgeV2 *FastBridgeV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.RevokeRole(&_FastBridgeV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.RevokeRole(&_FastBridgeV2.TransactOpts, role, account)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) SetChainGasAmount(opts *bind.TransactOpts, newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "setChainGasAmount", newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_FastBridgeV2 *FastBridgeV2Session) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.SetChainGasAmount(&_FastBridgeV2.TransactOpts, newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.SetChainGasAmount(&_FastBridgeV2.TransactOpts, newChainGasAmount)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) SetProtocolFeeRate(opts *bind.TransactOpts, newFeeRate *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "setProtocolFeeRate", newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_FastBridgeV2 *FastBridgeV2Session) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.SetProtocolFeeRate(&_FastBridgeV2.TransactOpts, newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.SetProtocolFeeRate(&_FastBridgeV2.TransactOpts, newFeeRate)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) SweepProtocolFees(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "sweepProtocolFees", token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_FastBridgeV2 *FastBridgeV2Session) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.SweepProtocolFees(&_FastBridgeV2.TransactOpts, token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.SweepProtocolFees(&_FastBridgeV2.TransactOpts, token, recipient)
}

// FastBridgeV2BridgeDepositClaimedIterator is returned from FilterBridgeDepositClaimed and is used to iterate over the raw logs and unpacked data for BridgeDepositClaimed events raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeDepositClaimedIterator struct {
	Event *FastBridgeV2BridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2BridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2BridgeDepositClaimed)
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
		it.Event = new(FastBridgeV2BridgeDepositClaimed)
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
func (it *FastBridgeV2BridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2BridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2BridgeDepositClaimed represents a BridgeDepositClaimed event raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeDepositClaimed struct {
	TransactionId [32]byte
	Relayer       common.Address
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositClaimed is a free log retrieval operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterBridgeDepositClaimed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*FastBridgeV2BridgeDepositClaimedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2BridgeDepositClaimedIterator{contract: _FastBridgeV2.contract, event: "BridgeDepositClaimed", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositClaimed is a free log subscription operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchBridgeDepositClaimed(opts *bind.WatchOpts, sink chan<- *FastBridgeV2BridgeDepositClaimed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2BridgeDepositClaimed)
				if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
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
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseBridgeDepositClaimed(log types.Log) (*FastBridgeV2BridgeDepositClaimed, error) {
	event := new(FastBridgeV2BridgeDepositClaimed)
	if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2BridgeDepositRefundedIterator is returned from FilterBridgeDepositRefunded and is used to iterate over the raw logs and unpacked data for BridgeDepositRefunded events raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeDepositRefundedIterator struct {
	Event *FastBridgeV2BridgeDepositRefunded // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2BridgeDepositRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2BridgeDepositRefunded)
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
		it.Event = new(FastBridgeV2BridgeDepositRefunded)
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
func (it *FastBridgeV2BridgeDepositRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2BridgeDepositRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2BridgeDepositRefunded represents a BridgeDepositRefunded event raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeDepositRefunded struct {
	TransactionId [32]byte
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositRefunded is a free log retrieval operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterBridgeDepositRefunded(opts *bind.FilterOpts, transactionId [][32]byte, to []common.Address) (*FastBridgeV2BridgeDepositRefundedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2BridgeDepositRefundedIterator{contract: _FastBridgeV2.contract, event: "BridgeDepositRefunded", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositRefunded is a free log subscription operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchBridgeDepositRefunded(opts *bind.WatchOpts, sink chan<- *FastBridgeV2BridgeDepositRefunded, transactionId [][32]byte, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2BridgeDepositRefunded)
				if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
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
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseBridgeDepositRefunded(log types.Log) (*FastBridgeV2BridgeDepositRefunded, error) {
	event := new(FastBridgeV2BridgeDepositRefunded)
	if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2BridgeProofDisputedIterator is returned from FilterBridgeProofDisputed and is used to iterate over the raw logs and unpacked data for BridgeProofDisputed events raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeProofDisputedIterator struct {
	Event *FastBridgeV2BridgeProofDisputed // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2BridgeProofDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2BridgeProofDisputed)
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
		it.Event = new(FastBridgeV2BridgeProofDisputed)
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
func (it *FastBridgeV2BridgeProofDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2BridgeProofDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2BridgeProofDisputed represents a BridgeProofDisputed event raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeProofDisputed struct {
	TransactionId [32]byte
	Relayer       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofDisputed is a free log retrieval operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterBridgeProofDisputed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*FastBridgeV2BridgeProofDisputedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2BridgeProofDisputedIterator{contract: _FastBridgeV2.contract, event: "BridgeProofDisputed", logs: logs, sub: sub}, nil
}

// WatchBridgeProofDisputed is a free log subscription operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchBridgeProofDisputed(opts *bind.WatchOpts, sink chan<- *FastBridgeV2BridgeProofDisputed, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2BridgeProofDisputed)
				if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
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
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseBridgeProofDisputed(log types.Log) (*FastBridgeV2BridgeProofDisputed, error) {
	event := new(FastBridgeV2BridgeProofDisputed)
	if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2BridgeProofProvidedIterator is returned from FilterBridgeProofProvided and is used to iterate over the raw logs and unpacked data for BridgeProofProvided events raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeProofProvidedIterator struct {
	Event *FastBridgeV2BridgeProofProvided // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2BridgeProofProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2BridgeProofProvided)
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
		it.Event = new(FastBridgeV2BridgeProofProvided)
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
func (it *FastBridgeV2BridgeProofProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2BridgeProofProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2BridgeProofProvided represents a BridgeProofProvided event raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeProofProvided struct {
	TransactionId   [32]byte
	Relayer         common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofProvided is a free log retrieval operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterBridgeProofProvided(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*FastBridgeV2BridgeProofProvidedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2BridgeProofProvidedIterator{contract: _FastBridgeV2.contract, event: "BridgeProofProvided", logs: logs, sub: sub}, nil
}

// WatchBridgeProofProvided is a free log subscription operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchBridgeProofProvided(opts *bind.WatchOpts, sink chan<- *FastBridgeV2BridgeProofProvided, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2BridgeProofProvided)
				if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
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
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseBridgeProofProvided(log types.Log) (*FastBridgeV2BridgeProofProvided, error) {
	event := new(FastBridgeV2BridgeProofProvided)
	if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2BridgeQuoteDetailsIterator is returned from FilterBridgeQuoteDetails and is used to iterate over the raw logs and unpacked data for BridgeQuoteDetails events raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeQuoteDetailsIterator struct {
	Event *FastBridgeV2BridgeQuoteDetails // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2BridgeQuoteDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2BridgeQuoteDetails)
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
		it.Event = new(FastBridgeV2BridgeQuoteDetails)
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
func (it *FastBridgeV2BridgeQuoteDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2BridgeQuoteDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2BridgeQuoteDetails represents a BridgeQuoteDetails event raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeQuoteDetails struct {
	TransactionId [32]byte
	QuoteId       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeQuoteDetails is a free log retrieval operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterBridgeQuoteDetails(opts *bind.FilterOpts, transactionId [][32]byte) (*FastBridgeV2BridgeQuoteDetailsIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "BridgeQuoteDetails", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2BridgeQuoteDetailsIterator{contract: _FastBridgeV2.contract, event: "BridgeQuoteDetails", logs: logs, sub: sub}, nil
}

// WatchBridgeQuoteDetails is a free log subscription operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchBridgeQuoteDetails(opts *bind.WatchOpts, sink chan<- *FastBridgeV2BridgeQuoteDetails, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "BridgeQuoteDetails", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2BridgeQuoteDetails)
				if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeQuoteDetails", log); err != nil {
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

// ParseBridgeQuoteDetails is a log parse operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseBridgeQuoteDetails(log types.Log) (*FastBridgeV2BridgeQuoteDetails, error) {
	event := new(FastBridgeV2BridgeQuoteDetails)
	if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeQuoteDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2BridgeRelayedIterator is returned from FilterBridgeRelayed and is used to iterate over the raw logs and unpacked data for BridgeRelayed events raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeRelayedIterator struct {
	Event *FastBridgeV2BridgeRelayed // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2BridgeRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2BridgeRelayed)
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
		it.Event = new(FastBridgeV2BridgeRelayed)
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
func (it *FastBridgeV2BridgeRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2BridgeRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2BridgeRelayed represents a BridgeRelayed event raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeRelayed struct {
	TransactionId  [32]byte
	Relayer        common.Address
	To             common.Address
	OriginChainId  uint32
	OriginToken    common.Address
	DestToken      common.Address
	OriginAmount   *big.Int
	DestAmount     *big.Int
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeRelayed is a free log retrieval operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterBridgeRelayed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*FastBridgeV2BridgeRelayedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2BridgeRelayedIterator{contract: _FastBridgeV2.contract, event: "BridgeRelayed", logs: logs, sub: sub}, nil
}

// WatchBridgeRelayed is a free log subscription operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchBridgeRelayed(opts *bind.WatchOpts, sink chan<- *FastBridgeV2BridgeRelayed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2BridgeRelayed)
				if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
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

// ParseBridgeRelayed is a log parse operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseBridgeRelayed(log types.Log) (*FastBridgeV2BridgeRelayed, error) {
	event := new(FastBridgeV2BridgeRelayed)
	if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2BridgeRequestedIterator is returned from FilterBridgeRequested and is used to iterate over the raw logs and unpacked data for BridgeRequested events raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeRequestedIterator struct {
	Event *FastBridgeV2BridgeRequested // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2BridgeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2BridgeRequested)
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
		it.Event = new(FastBridgeV2BridgeRequested)
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
func (it *FastBridgeV2BridgeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2BridgeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2BridgeRequested represents a BridgeRequested event raised by the FastBridgeV2 contract.
type FastBridgeV2BridgeRequested struct {
	TransactionId [32]byte
	Sender        common.Address
	Request       []byte
	DestChainId   uint32
	OriginToken   common.Address
	DestToken     common.Address
	OriginAmount  *big.Int
	DestAmount    *big.Int
	SendChainGas  bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRequested is a free log retrieval operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterBridgeRequested(opts *bind.FilterOpts, transactionId [][32]byte, sender []common.Address) (*FastBridgeV2BridgeRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2BridgeRequestedIterator{contract: _FastBridgeV2.contract, event: "BridgeRequested", logs: logs, sub: sub}, nil
}

// WatchBridgeRequested is a free log subscription operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchBridgeRequested(opts *bind.WatchOpts, sink chan<- *FastBridgeV2BridgeRequested, transactionId [][32]byte, sender []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2BridgeRequested)
				if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
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

// ParseBridgeRequested is a log parse operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseBridgeRequested(log types.Log) (*FastBridgeV2BridgeRequested, error) {
	event := new(FastBridgeV2BridgeRequested)
	if err := _FastBridgeV2.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2ChainGasAmountUpdatedIterator is returned from FilterChainGasAmountUpdated and is used to iterate over the raw logs and unpacked data for ChainGasAmountUpdated events raised by the FastBridgeV2 contract.
type FastBridgeV2ChainGasAmountUpdatedIterator struct {
	Event *FastBridgeV2ChainGasAmountUpdated // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2ChainGasAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2ChainGasAmountUpdated)
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
		it.Event = new(FastBridgeV2ChainGasAmountUpdated)
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
func (it *FastBridgeV2ChainGasAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2ChainGasAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2ChainGasAmountUpdated represents a ChainGasAmountUpdated event raised by the FastBridgeV2 contract.
type FastBridgeV2ChainGasAmountUpdated struct {
	OldChainGasAmount *big.Int
	NewChainGasAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChainGasAmountUpdated is a free log retrieval operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterChainGasAmountUpdated(opts *bind.FilterOpts) (*FastBridgeV2ChainGasAmountUpdatedIterator, error) {

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2ChainGasAmountUpdatedIterator{contract: _FastBridgeV2.contract, event: "ChainGasAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchChainGasAmountUpdated is a free log subscription operation binding the contract event 0x5cf09b12f3f56b4c564d51b25b40360af6d795198adb61ae0806a36c294323fa.
//
// Solidity: event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchChainGasAmountUpdated(opts *bind.WatchOpts, sink chan<- *FastBridgeV2ChainGasAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2ChainGasAmountUpdated)
				if err := _FastBridgeV2.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
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
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseChainGasAmountUpdated(log types.Log) (*FastBridgeV2ChainGasAmountUpdated, error) {
	event := new(FastBridgeV2ChainGasAmountUpdated)
	if err := _FastBridgeV2.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2FeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the FastBridgeV2 contract.
type FastBridgeV2FeeRateUpdatedIterator struct {
	Event *FastBridgeV2FeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2FeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2FeeRateUpdated)
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
		it.Event = new(FastBridgeV2FeeRateUpdated)
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
func (it *FastBridgeV2FeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2FeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2FeeRateUpdated represents a FeeRateUpdated event raised by the FastBridgeV2 contract.
type FastBridgeV2FeeRateUpdated struct {
	OldFeeRate *big.Int
	NewFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*FastBridgeV2FeeRateUpdatedIterator, error) {

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2FeeRateUpdatedIterator{contract: _FastBridgeV2.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *FastBridgeV2FeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2FeeRateUpdated)
				if err := _FastBridgeV2.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseFeeRateUpdated(log types.Log) (*FastBridgeV2FeeRateUpdated, error) {
	event := new(FastBridgeV2FeeRateUpdated)
	if err := _FastBridgeV2.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2FeesSweptIterator is returned from FilterFeesSwept and is used to iterate over the raw logs and unpacked data for FeesSwept events raised by the FastBridgeV2 contract.
type FastBridgeV2FeesSweptIterator struct {
	Event *FastBridgeV2FeesSwept // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2FeesSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2FeesSwept)
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
		it.Event = new(FastBridgeV2FeesSwept)
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
func (it *FastBridgeV2FeesSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2FeesSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2FeesSwept represents a FeesSwept event raised by the FastBridgeV2 contract.
type FastBridgeV2FeesSwept struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesSwept is a free log retrieval operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterFeesSwept(opts *bind.FilterOpts) (*FastBridgeV2FeesSweptIterator, error) {

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2FeesSweptIterator{contract: _FastBridgeV2.contract, event: "FeesSwept", logs: logs, sub: sub}, nil
}

// WatchFeesSwept is a free log subscription operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchFeesSwept(opts *bind.WatchOpts, sink chan<- *FastBridgeV2FeesSwept) (event.Subscription, error) {

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2FeesSwept)
				if err := _FastBridgeV2.contract.UnpackLog(event, "FeesSwept", log); err != nil {
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
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseFeesSwept(log types.Log) (*FastBridgeV2FeesSwept, error) {
	event := new(FastBridgeV2FeesSwept)
	if err := _FastBridgeV2.contract.UnpackLog(event, "FeesSwept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FastBridgeV2 contract.
type FastBridgeV2RoleAdminChangedIterator struct {
	Event *FastBridgeV2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2RoleAdminChanged)
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
		it.Event = new(FastBridgeV2RoleAdminChanged)
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
func (it *FastBridgeV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2RoleAdminChanged represents a RoleAdminChanged event raised by the FastBridgeV2 contract.
type FastBridgeV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FastBridgeV2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2RoleAdminChangedIterator{contract: _FastBridgeV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FastBridgeV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2RoleAdminChanged)
				if err := _FastBridgeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseRoleAdminChanged(log types.Log) (*FastBridgeV2RoleAdminChanged, error) {
	event := new(FastBridgeV2RoleAdminChanged)
	if err := _FastBridgeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FastBridgeV2 contract.
type FastBridgeV2RoleGrantedIterator struct {
	Event *FastBridgeV2RoleGranted // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2RoleGranted)
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
		it.Event = new(FastBridgeV2RoleGranted)
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
func (it *FastBridgeV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2RoleGranted represents a RoleGranted event raised by the FastBridgeV2 contract.
type FastBridgeV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FastBridgeV2RoleGrantedIterator, error) {

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

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2RoleGrantedIterator{contract: _FastBridgeV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FastBridgeV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2RoleGranted)
				if err := _FastBridgeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseRoleGranted(log types.Log) (*FastBridgeV2RoleGranted, error) {
	event := new(FastBridgeV2RoleGranted)
	if err := _FastBridgeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FastBridgeV2 contract.
type FastBridgeV2RoleRevokedIterator struct {
	Event *FastBridgeV2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2RoleRevoked)
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
		it.Event = new(FastBridgeV2RoleRevoked)
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
func (it *FastBridgeV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2RoleRevoked represents a RoleRevoked event raised by the FastBridgeV2 contract.
type FastBridgeV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FastBridgeV2RoleRevokedIterator, error) {

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

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2RoleRevokedIterator{contract: _FastBridgeV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FastBridgeV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2RoleRevoked)
				if err := _FastBridgeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseRoleRevoked(log types.Log) (*FastBridgeV2RoleRevoked, error) {
	event := new(FastBridgeV2RoleRevoked)
	if err := _FastBridgeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// IAccessControlEnumerableMetaData contains all meta data concerning the IAccessControlEnumerable contract.
var IAccessControlEnumerableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
	},
}

// IAccessControlEnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlEnumerableMetaData.ABI instead.
var IAccessControlEnumerableABI = IAccessControlEnumerableMetaData.ABI

// Deprecated: Use IAccessControlEnumerableMetaData.Sigs instead.
// IAccessControlEnumerableFuncSigs maps the 4-byte function signature to its string representation.
var IAccessControlEnumerableFuncSigs = IAccessControlEnumerableMetaData.Sigs

// IAccessControlEnumerable is an auto generated Go binding around an Ethereum contract.
type IAccessControlEnumerable struct {
	IAccessControlEnumerableCaller     // Read-only binding to the contract
	IAccessControlEnumerableTransactor // Write-only binding to the contract
	IAccessControlEnumerableFilterer   // Log filterer for contract events
}

// IAccessControlEnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlEnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlEnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlEnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlEnumerableSession struct {
	Contract     *IAccessControlEnumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlEnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlEnumerableCallerSession struct {
	Contract *IAccessControlEnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IAccessControlEnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlEnumerableTransactorSession struct {
	Contract     *IAccessControlEnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IAccessControlEnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlEnumerableRaw struct {
	Contract *IAccessControlEnumerable // Generic contract binding to access the raw methods on
}

// IAccessControlEnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlEnumerableCallerRaw struct {
	Contract *IAccessControlEnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlEnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlEnumerableTransactorRaw struct {
	Contract *IAccessControlEnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControlEnumerable creates a new instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerable(address common.Address, backend bind.ContractBackend) (*IAccessControlEnumerable, error) {
	contract, err := bindIAccessControlEnumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerable{IAccessControlEnumerableCaller: IAccessControlEnumerableCaller{contract: contract}, IAccessControlEnumerableTransactor: IAccessControlEnumerableTransactor{contract: contract}, IAccessControlEnumerableFilterer: IAccessControlEnumerableFilterer{contract: contract}}, nil
}

// NewIAccessControlEnumerableCaller creates a new read-only instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlEnumerableCaller, error) {
	contract, err := bindIAccessControlEnumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableCaller{contract: contract}, nil
}

// NewIAccessControlEnumerableTransactor creates a new write-only instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlEnumerableTransactor, error) {
	contract, err := bindIAccessControlEnumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableTransactor{contract: contract}, nil
}

// NewIAccessControlEnumerableFilterer creates a new log filterer instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlEnumerableFilterer, error) {
	contract, err := bindIAccessControlEnumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableFilterer{contract: contract}, nil
}

// bindIAccessControlEnumerable binds a generic wrapper to an already deployed contract.
func bindIAccessControlEnumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccessControlEnumerableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControlEnumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControlEnumerable.Contract.GetRoleAdmin(&_IAccessControlEnumerable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControlEnumerable.Contract.GetRoleAdmin(&_IAccessControlEnumerable.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMember(&_IAccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMember(&_IAccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMemberCount(&_IAccessControlEnumerable.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMemberCount(&_IAccessControlEnumerable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControlEnumerable.Contract.HasRole(&_IAccessControlEnumerable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControlEnumerable.Contract.HasRole(&_IAccessControlEnumerable.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.GrantRole(&_IAccessControlEnumerable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.GrantRole(&_IAccessControlEnumerable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.RenounceRole(&_IAccessControlEnumerable.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.RenounceRole(&_IAccessControlEnumerable.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.RevokeRole(&_IAccessControlEnumerable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.RevokeRole(&_IAccessControlEnumerable.TransactOpts, role, account)
}

// IAccessControlEnumerableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleAdminChangedIterator struct {
	Event *IAccessControlEnumerableRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IAccessControlEnumerableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlEnumerableRoleAdminChanged)
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
		it.Event = new(IAccessControlEnumerableRoleAdminChanged)
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
func (it *IAccessControlEnumerableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlEnumerableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlEnumerableRoleAdminChanged represents a RoleAdminChanged event raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IAccessControlEnumerableRoleAdminChangedIterator, error) {

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

	logs, sub, err := _IAccessControlEnumerable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableRoleAdminChangedIterator{contract: _IAccessControlEnumerable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IAccessControlEnumerableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IAccessControlEnumerable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlEnumerableRoleAdminChanged)
				if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) ParseRoleAdminChanged(log types.Log) (*IAccessControlEnumerableRoleAdminChanged, error) {
	event := new(IAccessControlEnumerableRoleAdminChanged)
	if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlEnumerableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleGrantedIterator struct {
	Event *IAccessControlEnumerableRoleGranted // Event containing the contract specifics and raw log

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
func (it *IAccessControlEnumerableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlEnumerableRoleGranted)
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
		it.Event = new(IAccessControlEnumerableRoleGranted)
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
func (it *IAccessControlEnumerableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlEnumerableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlEnumerableRoleGranted represents a RoleGranted event raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlEnumerableRoleGrantedIterator, error) {

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

	logs, sub, err := _IAccessControlEnumerable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableRoleGrantedIterator{contract: _IAccessControlEnumerable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IAccessControlEnumerableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IAccessControlEnumerable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlEnumerableRoleGranted)
				if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) ParseRoleGranted(log types.Log) (*IAccessControlEnumerableRoleGranted, error) {
	event := new(IAccessControlEnumerableRoleGranted)
	if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlEnumerableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleRevokedIterator struct {
	Event *IAccessControlEnumerableRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IAccessControlEnumerableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlEnumerableRoleRevoked)
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
		it.Event = new(IAccessControlEnumerableRoleRevoked)
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
func (it *IAccessControlEnumerableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlEnumerableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlEnumerableRoleRevoked represents a RoleRevoked event raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlEnumerableRoleRevokedIterator, error) {

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

	logs, sub, err := _IAccessControlEnumerable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableRoleRevokedIterator{contract: _IAccessControlEnumerable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IAccessControlEnumerableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IAccessControlEnumerable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlEnumerableRoleRevoked)
				if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) ParseRoleRevoked(log types.Log) (*IAccessControlEnumerableRoleRevoked, error) {
	event := new(IAccessControlEnumerableRoleRevoked)
	if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminMetaData contains all meta data concerning the IAdmin contract.
var IAdminMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldChainGasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
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
	parsed, err := IAdminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := IERC20PermitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"45851694": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256))",
		"aa9641ab": "canClaim(bytes32,address)",
		"41fcb612": "claim(bytes,address)",
		"add98c70": "dispute(bytes32)",
		"ac11fb1a": "getBridgeTransaction(bytes)",
		"886d36ff": "prove(bytes,bytes32)",
		"5eb7d946": "refund(bytes)",
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
	parsed, err := IFastBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
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
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridge *IFastBridgeSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridge.Contract.GetBridgeTransaction(&_IFastBridge.CallOpts, request)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridge *IFastBridgeCallerSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridge.Contract.GetBridgeTransaction(&_IFastBridge.CallOpts, request)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridge *IFastBridgeTransactor) Bridge(opts *bind.TransactOpts, params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "bridge", params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridge *IFastBridgeSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridge.Contract.Bridge(&_IFastBridge.TransactOpts, params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
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

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridge *IFastBridgeTransactor) Refund(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "refund", request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridge *IFastBridgeSession) Refund(request []byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Refund(&_IFastBridge.TransactOpts, request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Refund(request []byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Refund(&_IFastBridge.TransactOpts, request)
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
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeDepositClaimed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*IFastBridgeBridgeDepositClaimedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeDepositClaimedIterator{contract: _IFastBridge.contract, event: "BridgeDepositClaimed", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositClaimed is a free log subscription operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeDepositClaimed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeDepositClaimed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
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
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
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
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeDepositRefunded(opts *bind.FilterOpts, transactionId [][32]byte, to []common.Address) (*IFastBridgeBridgeDepositRefundedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeDepositRefundedIterator{contract: _IFastBridge.contract, event: "BridgeDepositRefunded", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositRefunded is a free log subscription operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeDepositRefunded(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeDepositRefunded, transactionId [][32]byte, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
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
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
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
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeProofDisputed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*IFastBridgeBridgeProofDisputedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeProofDisputedIterator{contract: _IFastBridge.contract, event: "BridgeProofDisputed", logs: logs, sub: sub}, nil
}

// WatchBridgeProofDisputed is a free log subscription operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeProofDisputed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeProofDisputed, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
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
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
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
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeProofProvided(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*IFastBridgeBridgeProofProvidedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeProofProvidedIterator{contract: _IFastBridge.contract, event: "BridgeProofProvided", logs: logs, sub: sub}, nil
}

// WatchBridgeProofProvided is a free log subscription operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeProofProvided(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeProofProvided, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
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
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
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
	OriginChainId  uint32
	OriginToken    common.Address
	DestToken      common.Address
	OriginAmount   *big.Int
	DestAmount     *big.Int
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeRelayed is a free log retrieval operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeRelayed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*IFastBridgeBridgeRelayedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeRelayedIterator{contract: _IFastBridge.contract, event: "BridgeRelayed", logs: logs, sub: sub}, nil
}

// WatchBridgeRelayed is a free log subscription operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeRelayed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeRelayed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
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

// ParseBridgeRelayed is a log parse operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
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
	DestChainId   uint32
	OriginToken   common.Address
	DestToken     common.Address
	OriginAmount  *big.Int
	DestAmount    *big.Int
	SendChainGas  bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRequested is a free log retrieval operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeRequested(opts *bind.FilterOpts, transactionId [][32]byte, sender []common.Address) (*IFastBridgeBridgeRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeRequestedIterator{contract: _IFastBridge.contract, event: "BridgeRequested", logs: logs, sub: sub}, nil
}

// WatchBridgeRequested is a free log subscription operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeRequested(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeRequested, transactionId [][32]byte, sender []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
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

// ParseBridgeRequested is a log parse operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeRequested(log types.Log) (*IFastBridgeBridgeRequested, error) {
	event := new(IFastBridgeBridgeRequested)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2MetaData contains all meta data concerning the IFastBridgeV2 contract.
var IFastBridgeV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"}],\"name\":\"BridgeQuoteDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"quoteRelayer\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"quoteExclusivitySeconds\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeParamsV2\",\"name\":\"paramsV2\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeProofs\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"timestamp\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelays\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeStatuses\",\"outputs\":[{\"internalType\":\"enumIFastBridgeV2.BridgeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransactionV2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"exclusivityRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exclusivityEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeTransactionV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"45851694": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256))",
		"bfc7c607": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256),(address,int256,bytes,uint256,bytes))",
		"91ad5039": "bridgeProofs(bytes32)",
		"8379a24f": "bridgeRelays(bytes32)",
		"051287bc": "bridgeStatuses(bytes32)",
		"aa9641ab": "canClaim(bytes32,address)",
		"c63ff8dd": "claim(bytes)",
		"41fcb612": "claim(bytes,address)",
		"add98c70": "dispute(bytes32)",
		"ac11fb1a": "getBridgeTransaction(bytes)",
		"5aa6ccba": "getBridgeTransactionV2(bytes)",
		"886d36ff": "prove(bytes,bytes32)",
		"18e4357d": "prove(bytes32,bytes32,address)",
		"5eb7d946": "refund(bytes)",
		"8f0d6f17": "relay(bytes)",
		"9c9545f0": "relay(bytes,address)",
	},
}

// IFastBridgeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use IFastBridgeV2MetaData.ABI instead.
var IFastBridgeV2ABI = IFastBridgeV2MetaData.ABI

// Deprecated: Use IFastBridgeV2MetaData.Sigs instead.
// IFastBridgeV2FuncSigs maps the 4-byte function signature to its string representation.
var IFastBridgeV2FuncSigs = IFastBridgeV2MetaData.Sigs

// IFastBridgeV2 is an auto generated Go binding around an Ethereum contract.
type IFastBridgeV2 struct {
	IFastBridgeV2Caller     // Read-only binding to the contract
	IFastBridgeV2Transactor // Write-only binding to the contract
	IFastBridgeV2Filterer   // Log filterer for contract events
}

// IFastBridgeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IFastBridgeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IFastBridgeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFastBridgeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFastBridgeV2Session struct {
	Contract     *IFastBridgeV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFastBridgeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFastBridgeV2CallerSession struct {
	Contract *IFastBridgeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IFastBridgeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFastBridgeV2TransactorSession struct {
	Contract     *IFastBridgeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IFastBridgeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IFastBridgeV2Raw struct {
	Contract *IFastBridgeV2 // Generic contract binding to access the raw methods on
}

// IFastBridgeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFastBridgeV2CallerRaw struct {
	Contract *IFastBridgeV2Caller // Generic read-only contract binding to access the raw methods on
}

// IFastBridgeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFastBridgeV2TransactorRaw struct {
	Contract *IFastBridgeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIFastBridgeV2 creates a new instance of IFastBridgeV2, bound to a specific deployed contract.
func NewIFastBridgeV2(address common.Address, backend bind.ContractBackend) (*IFastBridgeV2, error) {
	contract, err := bindIFastBridgeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2{IFastBridgeV2Caller: IFastBridgeV2Caller{contract: contract}, IFastBridgeV2Transactor: IFastBridgeV2Transactor{contract: contract}, IFastBridgeV2Filterer: IFastBridgeV2Filterer{contract: contract}}, nil
}

// NewIFastBridgeV2Caller creates a new read-only instance of IFastBridgeV2, bound to a specific deployed contract.
func NewIFastBridgeV2Caller(address common.Address, caller bind.ContractCaller) (*IFastBridgeV2Caller, error) {
	contract, err := bindIFastBridgeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2Caller{contract: contract}, nil
}

// NewIFastBridgeV2Transactor creates a new write-only instance of IFastBridgeV2, bound to a specific deployed contract.
func NewIFastBridgeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*IFastBridgeV2Transactor, error) {
	contract, err := bindIFastBridgeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2Transactor{contract: contract}, nil
}

// NewIFastBridgeV2Filterer creates a new log filterer instance of IFastBridgeV2, bound to a specific deployed contract.
func NewIFastBridgeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*IFastBridgeV2Filterer, error) {
	contract, err := bindIFastBridgeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2Filterer{contract: contract}, nil
}

// bindIFastBridgeV2 binds a generic wrapper to an already deployed contract.
func bindIFastBridgeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFastBridgeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridgeV2 *IFastBridgeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridgeV2.Contract.IFastBridgeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridgeV2 *IFastBridgeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.IFastBridgeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridgeV2 *IFastBridgeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.IFastBridgeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridgeV2 *IFastBridgeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridgeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridgeV2 *IFastBridgeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridgeV2 *IFastBridgeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.contract.Transact(opts, method, params...)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_IFastBridgeV2 *IFastBridgeV2Caller) BridgeProofs(opts *bind.CallOpts, transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "bridgeProofs", transactionId)

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
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_IFastBridgeV2 *IFastBridgeV2Session) BridgeProofs(transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _IFastBridgeV2.Contract.BridgeProofs(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) BridgeProofs(transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _IFastBridgeV2.Contract.BridgeProofs(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2Caller) BridgeRelays(opts *bind.CallOpts, transactionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "bridgeRelays", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2Session) BridgeRelays(transactionId [32]byte) (bool, error) {
	return _IFastBridgeV2.Contract.BridgeRelays(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) BridgeRelays(transactionId [32]byte) (bool, error) {
	return _IFastBridgeV2.Contract.BridgeRelays(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8)
func (_IFastBridgeV2 *IFastBridgeV2Caller) BridgeStatuses(opts *bind.CallOpts, transactionId [32]byte) (uint8, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "bridgeStatuses", transactionId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8)
func (_IFastBridgeV2 *IFastBridgeV2Session) BridgeStatuses(transactionId [32]byte) (uint8, error) {
	return _IFastBridgeV2.Contract.BridgeStatuses(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8)
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) BridgeStatuses(transactionId [32]byte) (uint8, error) {
	return _IFastBridgeV2.Contract.BridgeStatuses(&_IFastBridgeV2.CallOpts, transactionId)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2Caller) CanClaim(opts *bind.CallOpts, transactionId [32]byte, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "canClaim", transactionId, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2Session) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _IFastBridgeV2.Contract.CanClaim(&_IFastBridgeV2.CallOpts, transactionId, relayer)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _IFastBridgeV2.Contract.CanClaim(&_IFastBridgeV2.CallOpts, transactionId, relayer)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridgeV2 *IFastBridgeV2Caller) GetBridgeTransaction(opts *bind.CallOpts, request []byte) (IFastBridgeBridgeTransaction, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "getBridgeTransaction", request)

	if err != nil {
		return *new(IFastBridgeBridgeTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeBridgeTransaction)).(*IFastBridgeBridgeTransaction)

	return out0, err

}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridgeV2 *IFastBridgeV2Session) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridgeV2.Contract.GetBridgeTransaction(&_IFastBridgeV2.CallOpts, request)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridgeV2.Contract.GetBridgeTransaction(&_IFastBridgeV2.CallOpts, request)
}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_IFastBridgeV2 *IFastBridgeV2Caller) GetBridgeTransactionV2(opts *bind.CallOpts, request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "getBridgeTransactionV2", request)

	if err != nil {
		return *new(IFastBridgeV2BridgeTransactionV2), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeV2BridgeTransactionV2)).(*IFastBridgeV2BridgeTransactionV2)

	return out0, err

}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_IFastBridgeV2 *IFastBridgeV2Session) GetBridgeTransactionV2(request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	return _IFastBridgeV2.Contract.GetBridgeTransactionV2(&_IFastBridgeV2.CallOpts, request)
}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) GetBridgeTransactionV2(request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	return _IFastBridgeV2.Contract.GetBridgeTransactionV2(&_IFastBridgeV2.CallOpts, request)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Bridge(opts *bind.TransactOpts, params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "bridge", params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Bridge(&_IFastBridgeV2.TransactOpts, params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Bridge(&_IFastBridgeV2.TransactOpts, params)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Bridge0(opts *bind.TransactOpts, params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "bridge0", params, paramsV2)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Bridge0(params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Bridge0(&_IFastBridgeV2.TransactOpts, params, paramsV2)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Bridge0(params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Bridge0(&_IFastBridgeV2.TransactOpts, params, paramsV2)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Claim(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "claim", request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Claim(&_IFastBridgeV2.TransactOpts, request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Claim(&_IFastBridgeV2.TransactOpts, request, to)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Claim0(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "claim0", request)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Claim0(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Claim0(&_IFastBridgeV2.TransactOpts, request)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Claim0(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Claim0(&_IFastBridgeV2.TransactOpts, request)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Dispute(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "dispute", transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Dispute(&_IFastBridgeV2.TransactOpts, transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Dispute(&_IFastBridgeV2.TransactOpts, transactionId)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Prove(opts *bind.TransactOpts, transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "prove", transactionId, destTxHash, relayer)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Prove(transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Prove(&_IFastBridgeV2.TransactOpts, transactionId, destTxHash, relayer)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Prove(transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Prove(&_IFastBridgeV2.TransactOpts, transactionId, destTxHash, relayer)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Prove0(opts *bind.TransactOpts, request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "prove0", request, destTxHash)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Prove0(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Prove0(&_IFastBridgeV2.TransactOpts, request, destTxHash)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Prove0(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Prove0(&_IFastBridgeV2.TransactOpts, request, destTxHash)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Refund(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "refund", request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Refund(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Refund(&_IFastBridgeV2.TransactOpts, request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Refund(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Refund(&_IFastBridgeV2.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Relay(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "relay", request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Relay(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Relay(&_IFastBridgeV2.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Relay(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Relay(&_IFastBridgeV2.TransactOpts, request)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Relay0(opts *bind.TransactOpts, request []byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "relay0", request, relayer)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Relay0(request []byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Relay0(&_IFastBridgeV2.TransactOpts, request, relayer)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Relay0(request []byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Relay0(&_IFastBridgeV2.TransactOpts, request, relayer)
}

// IFastBridgeV2BridgeDepositClaimedIterator is returned from FilterBridgeDepositClaimed and is used to iterate over the raw logs and unpacked data for BridgeDepositClaimed events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeDepositClaimedIterator struct {
	Event *IFastBridgeV2BridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeDepositClaimed)
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
		it.Event = new(IFastBridgeV2BridgeDepositClaimed)
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
func (it *IFastBridgeV2BridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeDepositClaimed represents a BridgeDepositClaimed event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeDepositClaimed struct {
	TransactionId [32]byte
	Relayer       common.Address
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositClaimed is a free log retrieval operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeDepositClaimed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*IFastBridgeV2BridgeDepositClaimedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeDepositClaimedIterator{contract: _IFastBridgeV2.contract, event: "BridgeDepositClaimed", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositClaimed is a free log subscription operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeDepositClaimed(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeDepositClaimed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeDepositClaimed)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
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
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeDepositClaimed(log types.Log) (*IFastBridgeV2BridgeDepositClaimed, error) {
	event := new(IFastBridgeV2BridgeDepositClaimed)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeDepositRefundedIterator is returned from FilterBridgeDepositRefunded and is used to iterate over the raw logs and unpacked data for BridgeDepositRefunded events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeDepositRefundedIterator struct {
	Event *IFastBridgeV2BridgeDepositRefunded // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeDepositRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeDepositRefunded)
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
		it.Event = new(IFastBridgeV2BridgeDepositRefunded)
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
func (it *IFastBridgeV2BridgeDepositRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeDepositRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeDepositRefunded represents a BridgeDepositRefunded event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeDepositRefunded struct {
	TransactionId [32]byte
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositRefunded is a free log retrieval operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeDepositRefunded(opts *bind.FilterOpts, transactionId [][32]byte, to []common.Address) (*IFastBridgeV2BridgeDepositRefundedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeDepositRefundedIterator{contract: _IFastBridgeV2.contract, event: "BridgeDepositRefunded", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositRefunded is a free log subscription operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeDepositRefunded(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeDepositRefunded, transactionId [][32]byte, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeDepositRefunded)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
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
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeDepositRefunded(log types.Log) (*IFastBridgeV2BridgeDepositRefunded, error) {
	event := new(IFastBridgeV2BridgeDepositRefunded)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeProofDisputedIterator is returned from FilterBridgeProofDisputed and is used to iterate over the raw logs and unpacked data for BridgeProofDisputed events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeProofDisputedIterator struct {
	Event *IFastBridgeV2BridgeProofDisputed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeProofDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeProofDisputed)
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
		it.Event = new(IFastBridgeV2BridgeProofDisputed)
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
func (it *IFastBridgeV2BridgeProofDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeProofDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeProofDisputed represents a BridgeProofDisputed event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeProofDisputed struct {
	TransactionId [32]byte
	Relayer       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofDisputed is a free log retrieval operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeProofDisputed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*IFastBridgeV2BridgeProofDisputedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeProofDisputedIterator{contract: _IFastBridgeV2.contract, event: "BridgeProofDisputed", logs: logs, sub: sub}, nil
}

// WatchBridgeProofDisputed is a free log subscription operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeProofDisputed(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeProofDisputed, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeProofDisputed)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
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
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeProofDisputed(log types.Log) (*IFastBridgeV2BridgeProofDisputed, error) {
	event := new(IFastBridgeV2BridgeProofDisputed)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeProofProvidedIterator is returned from FilterBridgeProofProvided and is used to iterate over the raw logs and unpacked data for BridgeProofProvided events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeProofProvidedIterator struct {
	Event *IFastBridgeV2BridgeProofProvided // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeProofProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeProofProvided)
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
		it.Event = new(IFastBridgeV2BridgeProofProvided)
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
func (it *IFastBridgeV2BridgeProofProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeProofProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeProofProvided represents a BridgeProofProvided event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeProofProvided struct {
	TransactionId   [32]byte
	Relayer         common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofProvided is a free log retrieval operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeProofProvided(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*IFastBridgeV2BridgeProofProvidedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeProofProvidedIterator{contract: _IFastBridgeV2.contract, event: "BridgeProofProvided", logs: logs, sub: sub}, nil
}

// WatchBridgeProofProvided is a free log subscription operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeProofProvided(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeProofProvided, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeProofProvided)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
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
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeProofProvided(log types.Log) (*IFastBridgeV2BridgeProofProvided, error) {
	event := new(IFastBridgeV2BridgeProofProvided)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeQuoteDetailsIterator is returned from FilterBridgeQuoteDetails and is used to iterate over the raw logs and unpacked data for BridgeQuoteDetails events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeQuoteDetailsIterator struct {
	Event *IFastBridgeV2BridgeQuoteDetails // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeQuoteDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeQuoteDetails)
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
		it.Event = new(IFastBridgeV2BridgeQuoteDetails)
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
func (it *IFastBridgeV2BridgeQuoteDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeQuoteDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeQuoteDetails represents a BridgeQuoteDetails event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeQuoteDetails struct {
	TransactionId [32]byte
	QuoteId       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeQuoteDetails is a free log retrieval operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeQuoteDetails(opts *bind.FilterOpts, transactionId [][32]byte) (*IFastBridgeV2BridgeQuoteDetailsIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeQuoteDetails", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeQuoteDetailsIterator{contract: _IFastBridgeV2.contract, event: "BridgeQuoteDetails", logs: logs, sub: sub}, nil
}

// WatchBridgeQuoteDetails is a free log subscription operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeQuoteDetails(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeQuoteDetails, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeQuoteDetails", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeQuoteDetails)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeQuoteDetails", log); err != nil {
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

// ParseBridgeQuoteDetails is a log parse operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeQuoteDetails(log types.Log) (*IFastBridgeV2BridgeQuoteDetails, error) {
	event := new(IFastBridgeV2BridgeQuoteDetails)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeQuoteDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeRelayedIterator is returned from FilterBridgeRelayed and is used to iterate over the raw logs and unpacked data for BridgeRelayed events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeRelayedIterator struct {
	Event *IFastBridgeV2BridgeRelayed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeRelayed)
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
		it.Event = new(IFastBridgeV2BridgeRelayed)
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
func (it *IFastBridgeV2BridgeRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeRelayed represents a BridgeRelayed event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeRelayed struct {
	TransactionId  [32]byte
	Relayer        common.Address
	To             common.Address
	OriginChainId  uint32
	OriginToken    common.Address
	DestToken      common.Address
	OriginAmount   *big.Int
	DestAmount     *big.Int
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeRelayed is a free log retrieval operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeRelayed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*IFastBridgeV2BridgeRelayedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeRelayedIterator{contract: _IFastBridgeV2.contract, event: "BridgeRelayed", logs: logs, sub: sub}, nil
}

// WatchBridgeRelayed is a free log subscription operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeRelayed(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeRelayed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeRelayed)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
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

// ParseBridgeRelayed is a log parse operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeRelayed(log types.Log) (*IFastBridgeV2BridgeRelayed, error) {
	event := new(IFastBridgeV2BridgeRelayed)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeRequestedIterator is returned from FilterBridgeRequested and is used to iterate over the raw logs and unpacked data for BridgeRequested events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeRequestedIterator struct {
	Event *IFastBridgeV2BridgeRequested // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeRequested)
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
		it.Event = new(IFastBridgeV2BridgeRequested)
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
func (it *IFastBridgeV2BridgeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeRequested represents a BridgeRequested event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeRequested struct {
	TransactionId [32]byte
	Sender        common.Address
	Request       []byte
	DestChainId   uint32
	OriginToken   common.Address
	DestToken     common.Address
	OriginAmount  *big.Int
	DestAmount    *big.Int
	SendChainGas  bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRequested is a free log retrieval operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeRequested(opts *bind.FilterOpts, transactionId [][32]byte, sender []common.Address) (*IFastBridgeV2BridgeRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeRequestedIterator{contract: _IFastBridgeV2.contract, event: "BridgeRequested", logs: logs, sub: sub}, nil
}

// WatchBridgeRequested is a free log subscription operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeRequested(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeRequested, transactionId [][32]byte, sender []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeRequested)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
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

// ParseBridgeRequested is a log parse operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeRequested(log types.Log) (*IFastBridgeV2BridgeRequested, error) {
	event := new(IFastBridgeV2BridgeRequested)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2ErrorsMetaData contains all meta data concerning the IFastBridgeV2Errors contract.
var IFastBridgeV2ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AmountIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineNotExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExclusivityParamsIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExclusivityPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecipientIncorrectReturnValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecipientNoReturnValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StatusIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZapDataLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZapNativeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"}]",
}

// IFastBridgeV2ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IFastBridgeV2ErrorsMetaData.ABI instead.
var IFastBridgeV2ErrorsABI = IFastBridgeV2ErrorsMetaData.ABI

// IFastBridgeV2Errors is an auto generated Go binding around an Ethereum contract.
type IFastBridgeV2Errors struct {
	IFastBridgeV2ErrorsCaller     // Read-only binding to the contract
	IFastBridgeV2ErrorsTransactor // Write-only binding to the contract
	IFastBridgeV2ErrorsFilterer   // Log filterer for contract events
}

// IFastBridgeV2ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFastBridgeV2ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFastBridgeV2ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFastBridgeV2ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFastBridgeV2ErrorsSession struct {
	Contract     *IFastBridgeV2Errors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IFastBridgeV2ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFastBridgeV2ErrorsCallerSession struct {
	Contract *IFastBridgeV2ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IFastBridgeV2ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFastBridgeV2ErrorsTransactorSession struct {
	Contract     *IFastBridgeV2ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IFastBridgeV2ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFastBridgeV2ErrorsRaw struct {
	Contract *IFastBridgeV2Errors // Generic contract binding to access the raw methods on
}

// IFastBridgeV2ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFastBridgeV2ErrorsCallerRaw struct {
	Contract *IFastBridgeV2ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IFastBridgeV2ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFastBridgeV2ErrorsTransactorRaw struct {
	Contract *IFastBridgeV2ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFastBridgeV2Errors creates a new instance of IFastBridgeV2Errors, bound to a specific deployed contract.
func NewIFastBridgeV2Errors(address common.Address, backend bind.ContractBackend) (*IFastBridgeV2Errors, error) {
	contract, err := bindIFastBridgeV2Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2Errors{IFastBridgeV2ErrorsCaller: IFastBridgeV2ErrorsCaller{contract: contract}, IFastBridgeV2ErrorsTransactor: IFastBridgeV2ErrorsTransactor{contract: contract}, IFastBridgeV2ErrorsFilterer: IFastBridgeV2ErrorsFilterer{contract: contract}}, nil
}

// NewIFastBridgeV2ErrorsCaller creates a new read-only instance of IFastBridgeV2Errors, bound to a specific deployed contract.
func NewIFastBridgeV2ErrorsCaller(address common.Address, caller bind.ContractCaller) (*IFastBridgeV2ErrorsCaller, error) {
	contract, err := bindIFastBridgeV2Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2ErrorsCaller{contract: contract}, nil
}

// NewIFastBridgeV2ErrorsTransactor creates a new write-only instance of IFastBridgeV2Errors, bound to a specific deployed contract.
func NewIFastBridgeV2ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IFastBridgeV2ErrorsTransactor, error) {
	contract, err := bindIFastBridgeV2Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2ErrorsTransactor{contract: contract}, nil
}

// NewIFastBridgeV2ErrorsFilterer creates a new log filterer instance of IFastBridgeV2Errors, bound to a specific deployed contract.
func NewIFastBridgeV2ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IFastBridgeV2ErrorsFilterer, error) {
	contract, err := bindIFastBridgeV2Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2ErrorsFilterer{contract: contract}, nil
}

// bindIFastBridgeV2Errors binds a generic wrapper to an already deployed contract.
func bindIFastBridgeV2Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFastBridgeV2ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridgeV2Errors *IFastBridgeV2ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridgeV2Errors.Contract.IFastBridgeV2ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridgeV2Errors *IFastBridgeV2ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridgeV2Errors.Contract.IFastBridgeV2ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridgeV2Errors *IFastBridgeV2ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridgeV2Errors.Contract.IFastBridgeV2ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridgeV2Errors *IFastBridgeV2ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridgeV2Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridgeV2Errors *IFastBridgeV2ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridgeV2Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridgeV2Errors *IFastBridgeV2ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridgeV2Errors.Contract.contract.Transact(opts, method, params...)
}

// IMulticallTargetMetaData contains all meta data concerning the IMulticallTarget contract.
var IMulticallTargetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"ignoreReverts\",\"type\":\"bool\"}],\"name\":\"multicallNoResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"ignoreReverts\",\"type\":\"bool\"}],\"name\":\"multicallWithResults\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticallTarget.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3f61331d": "multicallNoResults(bytes[],bool)",
		"385c1d2f": "multicallWithResults(bytes[],bool)",
	},
}

// IMulticallTargetABI is the input ABI used to generate the binding from.
// Deprecated: Use IMulticallTargetMetaData.ABI instead.
var IMulticallTargetABI = IMulticallTargetMetaData.ABI

// Deprecated: Use IMulticallTargetMetaData.Sigs instead.
// IMulticallTargetFuncSigs maps the 4-byte function signature to its string representation.
var IMulticallTargetFuncSigs = IMulticallTargetMetaData.Sigs

// IMulticallTarget is an auto generated Go binding around an Ethereum contract.
type IMulticallTarget struct {
	IMulticallTargetCaller     // Read-only binding to the contract
	IMulticallTargetTransactor // Write-only binding to the contract
	IMulticallTargetFilterer   // Log filterer for contract events
}

// IMulticallTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMulticallTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticallTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMulticallTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticallTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMulticallTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticallTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMulticallTargetSession struct {
	Contract     *IMulticallTarget // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMulticallTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMulticallTargetCallerSession struct {
	Contract *IMulticallTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IMulticallTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMulticallTargetTransactorSession struct {
	Contract     *IMulticallTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IMulticallTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMulticallTargetRaw struct {
	Contract *IMulticallTarget // Generic contract binding to access the raw methods on
}

// IMulticallTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMulticallTargetCallerRaw struct {
	Contract *IMulticallTargetCaller // Generic read-only contract binding to access the raw methods on
}

// IMulticallTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMulticallTargetTransactorRaw struct {
	Contract *IMulticallTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMulticallTarget creates a new instance of IMulticallTarget, bound to a specific deployed contract.
func NewIMulticallTarget(address common.Address, backend bind.ContractBackend) (*IMulticallTarget, error) {
	contract, err := bindIMulticallTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMulticallTarget{IMulticallTargetCaller: IMulticallTargetCaller{contract: contract}, IMulticallTargetTransactor: IMulticallTargetTransactor{contract: contract}, IMulticallTargetFilterer: IMulticallTargetFilterer{contract: contract}}, nil
}

// NewIMulticallTargetCaller creates a new read-only instance of IMulticallTarget, bound to a specific deployed contract.
func NewIMulticallTargetCaller(address common.Address, caller bind.ContractCaller) (*IMulticallTargetCaller, error) {
	contract, err := bindIMulticallTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMulticallTargetCaller{contract: contract}, nil
}

// NewIMulticallTargetTransactor creates a new write-only instance of IMulticallTarget, bound to a specific deployed contract.
func NewIMulticallTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*IMulticallTargetTransactor, error) {
	contract, err := bindIMulticallTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMulticallTargetTransactor{contract: contract}, nil
}

// NewIMulticallTargetFilterer creates a new log filterer instance of IMulticallTarget, bound to a specific deployed contract.
func NewIMulticallTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*IMulticallTargetFilterer, error) {
	contract, err := bindIMulticallTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMulticallTargetFilterer{contract: contract}, nil
}

// bindIMulticallTarget binds a generic wrapper to an already deployed contract.
func bindIMulticallTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMulticallTargetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMulticallTarget *IMulticallTargetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMulticallTarget.Contract.IMulticallTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMulticallTarget *IMulticallTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMulticallTarget.Contract.IMulticallTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMulticallTarget *IMulticallTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMulticallTarget.Contract.IMulticallTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMulticallTarget *IMulticallTargetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMulticallTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMulticallTarget *IMulticallTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMulticallTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMulticallTarget *IMulticallTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMulticallTarget.Contract.contract.Transact(opts, method, params...)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_IMulticallTarget *IMulticallTargetTransactor) MulticallNoResults(opts *bind.TransactOpts, data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _IMulticallTarget.contract.Transact(opts, "multicallNoResults", data, ignoreReverts)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_IMulticallTarget *IMulticallTargetSession) MulticallNoResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _IMulticallTarget.Contract.MulticallNoResults(&_IMulticallTarget.TransactOpts, data, ignoreReverts)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_IMulticallTarget *IMulticallTargetTransactorSession) MulticallNoResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _IMulticallTarget.Contract.MulticallNoResults(&_IMulticallTarget.TransactOpts, data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_IMulticallTarget *IMulticallTargetTransactor) MulticallWithResults(opts *bind.TransactOpts, data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _IMulticallTarget.contract.Transact(opts, "multicallWithResults", data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_IMulticallTarget *IMulticallTargetSession) MulticallWithResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _IMulticallTarget.Contract.MulticallWithResults(&_IMulticallTarget.TransactOpts, data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_IMulticallTarget *IMulticallTargetTransactorSession) MulticallWithResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _IMulticallTarget.Contract.MulticallWithResults(&_IMulticallTarget.TransactOpts, data, ignoreReverts)
}

// IZapRecipientMetaData contains all meta data concerning the IZapRecipient contract.
var IZapRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"name\":\"zap\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e85e13dd": "zap(address,uint256,bytes)",
	},
}

// IZapRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use IZapRecipientMetaData.ABI instead.
var IZapRecipientABI = IZapRecipientMetaData.ABI

// Deprecated: Use IZapRecipientMetaData.Sigs instead.
// IZapRecipientFuncSigs maps the 4-byte function signature to its string representation.
var IZapRecipientFuncSigs = IZapRecipientMetaData.Sigs

// IZapRecipient is an auto generated Go binding around an Ethereum contract.
type IZapRecipient struct {
	IZapRecipientCaller     // Read-only binding to the contract
	IZapRecipientTransactor // Write-only binding to the contract
	IZapRecipientFilterer   // Log filterer for contract events
}

// IZapRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZapRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZapRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZapRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZapRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZapRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZapRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZapRecipientSession struct {
	Contract     *IZapRecipient    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZapRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZapRecipientCallerSession struct {
	Contract *IZapRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IZapRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZapRecipientTransactorSession struct {
	Contract     *IZapRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IZapRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZapRecipientRaw struct {
	Contract *IZapRecipient // Generic contract binding to access the raw methods on
}

// IZapRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZapRecipientCallerRaw struct {
	Contract *IZapRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// IZapRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZapRecipientTransactorRaw struct {
	Contract *IZapRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZapRecipient creates a new instance of IZapRecipient, bound to a specific deployed contract.
func NewIZapRecipient(address common.Address, backend bind.ContractBackend) (*IZapRecipient, error) {
	contract, err := bindIZapRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZapRecipient{IZapRecipientCaller: IZapRecipientCaller{contract: contract}, IZapRecipientTransactor: IZapRecipientTransactor{contract: contract}, IZapRecipientFilterer: IZapRecipientFilterer{contract: contract}}, nil
}

// NewIZapRecipientCaller creates a new read-only instance of IZapRecipient, bound to a specific deployed contract.
func NewIZapRecipientCaller(address common.Address, caller bind.ContractCaller) (*IZapRecipientCaller, error) {
	contract, err := bindIZapRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZapRecipientCaller{contract: contract}, nil
}

// NewIZapRecipientTransactor creates a new write-only instance of IZapRecipient, bound to a specific deployed contract.
func NewIZapRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*IZapRecipientTransactor, error) {
	contract, err := bindIZapRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZapRecipientTransactor{contract: contract}, nil
}

// NewIZapRecipientFilterer creates a new log filterer instance of IZapRecipient, bound to a specific deployed contract.
func NewIZapRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*IZapRecipientFilterer, error) {
	contract, err := bindIZapRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZapRecipientFilterer{contract: contract}, nil
}

// bindIZapRecipient binds a generic wrapper to an already deployed contract.
func bindIZapRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZapRecipientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZapRecipient *IZapRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZapRecipient.Contract.IZapRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZapRecipient *IZapRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZapRecipient.Contract.IZapRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZapRecipient *IZapRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZapRecipient.Contract.IZapRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZapRecipient *IZapRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZapRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZapRecipient *IZapRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZapRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZapRecipient *IZapRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZapRecipient.Contract.contract.Transact(opts, method, params...)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address token, uint256 amount, bytes zapData) payable returns(bytes4)
func (_IZapRecipient *IZapRecipientTransactor) Zap(opts *bind.TransactOpts, token common.Address, amount *big.Int, zapData []byte) (*types.Transaction, error) {
	return _IZapRecipient.contract.Transact(opts, "zap", token, amount, zapData)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address token, uint256 amount, bytes zapData) payable returns(bytes4)
func (_IZapRecipient *IZapRecipientSession) Zap(token common.Address, amount *big.Int, zapData []byte) (*types.Transaction, error) {
	return _IZapRecipient.Contract.Zap(&_IZapRecipient.TransactOpts, token, amount, zapData)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address token, uint256 amount, bytes zapData) payable returns(bytes4)
func (_IZapRecipient *IZapRecipientTransactorSession) Zap(token common.Address, amount *big.Int, zapData []byte) (*types.Transaction, error) {
	return _IZapRecipient.Contract.Zap(&_IZapRecipient.TransactOpts, token, amount, zapData)
}

// MulticallTargetMetaData contains all meta data concerning the MulticallTarget contract.
var MulticallTargetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MulticallTarget__UndeterminedRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"ignoreReverts\",\"type\":\"bool\"}],\"name\":\"multicallNoResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"ignoreReverts\",\"type\":\"bool\"}],\"name\":\"multicallWithResults\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticallTarget.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3f61331d": "multicallNoResults(bytes[],bool)",
		"385c1d2f": "multicallWithResults(bytes[],bool)",
	},
}

// MulticallTargetABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallTargetMetaData.ABI instead.
var MulticallTargetABI = MulticallTargetMetaData.ABI

// Deprecated: Use MulticallTargetMetaData.Sigs instead.
// MulticallTargetFuncSigs maps the 4-byte function signature to its string representation.
var MulticallTargetFuncSigs = MulticallTargetMetaData.Sigs

// MulticallTarget is an auto generated Go binding around an Ethereum contract.
type MulticallTarget struct {
	MulticallTargetCaller     // Read-only binding to the contract
	MulticallTargetTransactor // Write-only binding to the contract
	MulticallTargetFilterer   // Log filterer for contract events
}

// MulticallTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallTargetSession struct {
	Contract     *MulticallTarget  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallTargetCallerSession struct {
	Contract *MulticallTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MulticallTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallTargetTransactorSession struct {
	Contract     *MulticallTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MulticallTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallTargetRaw struct {
	Contract *MulticallTarget // Generic contract binding to access the raw methods on
}

// MulticallTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallTargetCallerRaw struct {
	Contract *MulticallTargetCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallTargetTransactorRaw struct {
	Contract *MulticallTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticallTarget creates a new instance of MulticallTarget, bound to a specific deployed contract.
func NewMulticallTarget(address common.Address, backend bind.ContractBackend) (*MulticallTarget, error) {
	contract, err := bindMulticallTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MulticallTarget{MulticallTargetCaller: MulticallTargetCaller{contract: contract}, MulticallTargetTransactor: MulticallTargetTransactor{contract: contract}, MulticallTargetFilterer: MulticallTargetFilterer{contract: contract}}, nil
}

// NewMulticallTargetCaller creates a new read-only instance of MulticallTarget, bound to a specific deployed contract.
func NewMulticallTargetCaller(address common.Address, caller bind.ContractCaller) (*MulticallTargetCaller, error) {
	contract, err := bindMulticallTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallTargetCaller{contract: contract}, nil
}

// NewMulticallTargetTransactor creates a new write-only instance of MulticallTarget, bound to a specific deployed contract.
func NewMulticallTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallTargetTransactor, error) {
	contract, err := bindMulticallTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallTargetTransactor{contract: contract}, nil
}

// NewMulticallTargetFilterer creates a new log filterer instance of MulticallTarget, bound to a specific deployed contract.
func NewMulticallTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallTargetFilterer, error) {
	contract, err := bindMulticallTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallTargetFilterer{contract: contract}, nil
}

// bindMulticallTarget binds a generic wrapper to an already deployed contract.
func bindMulticallTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MulticallTargetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallTarget *MulticallTargetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallTarget.Contract.MulticallTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallTarget *MulticallTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallTarget.Contract.MulticallTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallTarget *MulticallTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallTarget.Contract.MulticallTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallTarget *MulticallTargetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallTarget *MulticallTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallTarget *MulticallTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallTarget.Contract.contract.Transact(opts, method, params...)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_MulticallTarget *MulticallTargetTransactor) MulticallNoResults(opts *bind.TransactOpts, data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _MulticallTarget.contract.Transact(opts, "multicallNoResults", data, ignoreReverts)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_MulticallTarget *MulticallTargetSession) MulticallNoResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _MulticallTarget.Contract.MulticallNoResults(&_MulticallTarget.TransactOpts, data, ignoreReverts)
}

// MulticallNoResults is a paid mutator transaction binding the contract method 0x3f61331d.
//
// Solidity: function multicallNoResults(bytes[] data, bool ignoreReverts) returns()
func (_MulticallTarget *MulticallTargetTransactorSession) MulticallNoResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _MulticallTarget.Contract.MulticallNoResults(&_MulticallTarget.TransactOpts, data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_MulticallTarget *MulticallTargetTransactor) MulticallWithResults(opts *bind.TransactOpts, data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _MulticallTarget.contract.Transact(opts, "multicallWithResults", data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_MulticallTarget *MulticallTargetSession) MulticallWithResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _MulticallTarget.Contract.MulticallWithResults(&_MulticallTarget.TransactOpts, data, ignoreReverts)
}

// MulticallWithResults is a paid mutator transaction binding the contract method 0x385c1d2f.
//
// Solidity: function multicallWithResults(bytes[] data, bool ignoreReverts) returns((bool,bytes)[] results)
func (_MulticallTarget *MulticallTargetTransactorSession) MulticallWithResults(data [][]byte, ignoreReverts bool) (*types.Transaction, error) {
	return _MulticallTarget.Contract.MulticallWithResults(&_MulticallTarget.TransactOpts, data, ignoreReverts)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedDecrease\",\"type\":\"uint256\"}],\"name\":\"SafeERC20FailedDecreaseAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209fa7c4bd491bc75459d04b74ec3df4cf5631b4ff0438eacdbdda406ccb78986b64736f6c63430008180033",
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

// UniversalTokenLibMetaData contains all meta data concerning the UniversalTokenLib contract.
var UniversalTokenLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208ee6f45538d68dae1182af6006877f66fc2bb57b5fe35283c7b3ccf352c1148364736f6c63430008180033",
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
	parsed, err := UniversalTokenLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
