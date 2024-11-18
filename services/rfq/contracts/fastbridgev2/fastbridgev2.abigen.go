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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122053e3727ad57bc90f5560dc5a255131033ea91c54c8a147cf6df38af4aac448d564736f6c63430008180033",
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

// AdminV2MetaData contains all meta data concerning the AdminV2 contract.
var AdminV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CancelDelayBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRateAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCancelDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCancelDelay\",\"type\":\"uint256\"}],\"name\":\"CancelDelayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCELER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_RATE_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_GAS_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCancelDelay\",\"type\":\"uint256\"}],\"name\":\"setCancelDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"02d2ff66": "CANCELER_ROLE()",
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"930ac180": "DEFAULT_CANCEL_DELAY()",
		"bf333f2c": "FEE_BPS()",
		"0f5f6ed7": "FEE_RATE_MAX()",
		"ccc57490": "GOVERNOR_ROLE()",
		"03ed0ee5": "GUARD_ROLE()",
		"922b7487": "MIN_CANCEL_DELAY()",
		"0f862f1e": "NATIVE_GAS_TOKEN()",
		"dc9a4ef6": "PROVER_ROLE()",
		"7ebe815c": "QUOTER_ROLE()",
		"638a0f09": "cancelDelay()",
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
		"1ea327c5": "setCancelDelay(uint256)",
		"b13aa2d6": "setProtocolFeeRate(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"06f333f2": "sweepProtocolFees(address,address)",
	},
	Bin: "0x60a060405260006080523480156200001657600080fd5b50604051620015763803806200157683398101604081905262000039916200020a565b620000466000826200005c565b50620000556201518062000099565b5062000235565b6000806200006b848462000102565b90508015620000905760008481526001602052604090206200008e9084620001b0565b505b90505b92915050565b610e10811015620000bd57604051630e0ea5c760e01b815260040160405180910390fd5b600480549082905560408051828152602081018490527f909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95910160405180910390a15050565b6000828152602081815260408083206001600160a01b038516845290915281205460ff16620001a7576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556200015e3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600162000093565b50600062000093565b600062000090836001600160a01b0384166000818152600183016020526040812054620001a75750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000093565b6000602082840312156200021d57600080fd5b81516001600160a01b03811681146200009057600080fd5b60805161132562000251600039600061045201526113256000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c80639010d07c116100ee578063bf333f2c11610097578063d547741f11610071578063d547741f146103f3578063dc9a4ef614610406578063dcf844a71461042d578063e00a83e01461044d57600080fd5b8063bf333f2c146103af578063ca15c873146103b9578063ccc57490146103cc57600080fd5b8063930ac180116100c8578063930ac1801461038a578063a217fddf14610394578063b13aa2d61461039c57600080fd5b80639010d07c1461032a57806391d148541461033d578063922b74871461038157600080fd5b80631ea327c51161015b57806336568abe1161013557806336568abe146102de57806358f85880146102f1578063638a0f09146102fa5780637ebe815c1461030357600080fd5b80631ea327c514610295578063248a9ca3146102a85780632f2ff15d146102cb57600080fd5b806306f333f21161018c57806306f333f2146102375780630f5f6ed71461024c5780630f862f1e1461025557600080fd5b806301ffc9a7146101b357806302d2ff66146101db57806303ed0ee514610210575b600080fd5b6101c66101c13660046110ef565b610474565b60405190151581526020015b60405180910390f35b6102027febfdca8e46c0b8dacf9989ee613e35727eadd20a1d5e5ad01a53968c7e5fe07a81565b6040519081526020016101d2565b6102027f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d81565b61024a61024536600461115a565b6104d0565b005b61020261271081565b61027073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d2565b61024a6102a336600461118d565b610611565b6102026102b636600461118d565b60009081526020819052604090206001015490565b61024a6102d93660046111a6565b610648565b61024a6102ec3660046111a6565b61066d565b61020260025481565b61020260045481565b6102027f9a04aea0a349253cc7277afafdf6ead6729a3972a47ffb40eaef2c93d4e1bfea81565b6102706103383660046111c9565b6106c6565b6101c661034b3660046111a6565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610202610e1081565b6102026201518081565b610202600081565b61024a6103aa36600461118d565b6106e5565b610202620f424081565b6102026103c736600461118d565b610791565b6102027f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b61024a6104013660046111a6565b6107a8565b6102027f60044782a422109ac08a415e44260ad5d3bad63d96ebe1fac0363399642ee3f281565b61020261043b3660046111eb565b60036020526000908152604090205481565b6102027f000000000000000000000000000000000000000000000000000000000000000081565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f0000000000000000000000000000000000000000000000000000000014806104ca57506104ca826107cd565b92915050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f556104fa81610864565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600360205260408120549081900361052e5750505050565b73ffffffffffffffffffffffffffffffffffffffff8481166000818152600360209081526040808320929092558151928352928616928201929092529081018290527f244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd9060600160405180910390a17fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff8516016105e9576105e48382610871565b61060a565b61060a73ffffffffffffffffffffffffffffffffffffffff8516848361094c565b505b505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5561063b81610864565b610644826109d9565b5050565b60008281526020819052604090206001015461066381610864565b61060a8383610a5a565b73ffffffffffffffffffffffffffffffffffffffff811633146106bc576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61060c8282610a8f565b60008281526001602052604081206106de9083610abc565b9392505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5561070f81610864565b61271082111561074b576040517f9b569b8100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280549083905560408051828152602081018590527f14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957910160405180910390a1505050565b60008181526001602052604081206104ca90610ac8565b6000828152602081905260409020600101546107c381610864565b61060a8383610a8f565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806104ca57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146104ca565b61086e8133610ad2565b50565b804710156108b2576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024015b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d806000811461090c576040519150601f19603f3d011682016040523d82523d6000602084013e610911565b606091505b505090508061060c576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261060c908490610b58565b610e10811015610a15576040517f0e0ea5c700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480549082905560408051828152602081018490527f909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95910160405180910390a15050565b600080610a678484610bee565b905080156106de576000848152600160205260409020610a879084610cea565b509392505050565b600080610a9c8484610d0c565b905080156106de576000848152600160205260409020610a879084610dc7565b60006106de8383610de9565b60006104ca825490565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610644576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044016108a9565b6000610b7a73ffffffffffffffffffffffffffffffffffffffff841683610e13565b90508051600014158015610b9f575080806020019051810190610b9d9190611206565b155b1561060c576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016108a9565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610ce25760008381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610c803390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104ca565b5060006104ca565b60006106de8373ffffffffffffffffffffffffffffffffffffffff8416610e21565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610ce25760008381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104ca565b60006106de8373ffffffffffffffffffffffffffffffffffffffff8416610e68565b6000826000018281548110610e0057610e00611228565b9060005260206000200154905092915050565b60606106de83836000610f5b565b6000818152600183016020526040812054610ce2575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556104ca565b60008181526001830160205260408120548015610f51576000610e8c600183611257565b8554909150600090610ea090600190611257565b9050808214610f05576000866000018281548110610ec057610ec0611228565b9060005260206000200154905080876000018481548110610ee357610ee3611228565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610f1657610f16611291565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506104ca565b60009150506104ca565b606081471015610f99576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016108a9565b6000808573ffffffffffffffffffffffffffffffffffffffff168486604051610fc291906112c0565b60006040518083038185875af1925050503d8060008114610fff576040519150601f19603f3d011682016040523d82523d6000602084013e611004565b606091505b509150915061101486838361101e565b9695505050505050565b6060826110335761102e826110ad565b6106de565b8151158015611057575073ffffffffffffffffffffffffffffffffffffffff84163b155b156110a6576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016108a9565b50806106de565b8051156110bd5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561110157600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146106de57600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461115557600080fd5b919050565b6000806040838503121561116d57600080fd5b61117683611131565b915061118460208401611131565b90509250929050565b60006020828403121561119f57600080fd5b5035919050565b600080604083850312156111b957600080fd5b8235915061118460208401611131565b600080604083850312156111dc57600080fd5b50508035926020909101359150565b6000602082840312156111fd57600080fd5b6106de82611131565b60006020828403121561121857600080fd5b815180151581146106de57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b818103818111156104ca577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000825160005b818110156112e157602081860181015185830152016112c7565b50600092019182525091905056fea264697066735822122079d04eeec7c857b32065468d4b56344ced14fdee358cfecc2cfd327dbced278e64736f6c63430008180033",
}

// AdminV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AdminV2MetaData.ABI instead.
var AdminV2ABI = AdminV2MetaData.ABI

// Deprecated: Use AdminV2MetaData.Sigs instead.
// AdminV2FuncSigs maps the 4-byte function signature to its string representation.
var AdminV2FuncSigs = AdminV2MetaData.Sigs

// AdminV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdminV2MetaData.Bin instead.
var AdminV2Bin = AdminV2MetaData.Bin

// DeployAdminV2 deploys a new Ethereum contract, binding an instance of AdminV2 to it.
func DeployAdminV2(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *AdminV2, error) {
	parsed, err := AdminV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdminV2Bin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminV2{AdminV2Caller: AdminV2Caller{contract: contract}, AdminV2Transactor: AdminV2Transactor{contract: contract}, AdminV2Filterer: AdminV2Filterer{contract: contract}}, nil
}

// AdminV2 is an auto generated Go binding around an Ethereum contract.
type AdminV2 struct {
	AdminV2Caller     // Read-only binding to the contract
	AdminV2Transactor // Write-only binding to the contract
	AdminV2Filterer   // Log filterer for contract events
}

// AdminV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AdminV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminV2Session struct {
	Contract     *AdminV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminV2CallerSession struct {
	Contract *AdminV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AdminV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminV2TransactorSession struct {
	Contract     *AdminV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AdminV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AdminV2Raw struct {
	Contract *AdminV2 // Generic contract binding to access the raw methods on
}

// AdminV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminV2CallerRaw struct {
	Contract *AdminV2Caller // Generic read-only contract binding to access the raw methods on
}

// AdminV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminV2TransactorRaw struct {
	Contract *AdminV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminV2 creates a new instance of AdminV2, bound to a specific deployed contract.
func NewAdminV2(address common.Address, backend bind.ContractBackend) (*AdminV2, error) {
	contract, err := bindAdminV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminV2{AdminV2Caller: AdminV2Caller{contract: contract}, AdminV2Transactor: AdminV2Transactor{contract: contract}, AdminV2Filterer: AdminV2Filterer{contract: contract}}, nil
}

// NewAdminV2Caller creates a new read-only instance of AdminV2, bound to a specific deployed contract.
func NewAdminV2Caller(address common.Address, caller bind.ContractCaller) (*AdminV2Caller, error) {
	contract, err := bindAdminV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminV2Caller{contract: contract}, nil
}

// NewAdminV2Transactor creates a new write-only instance of AdminV2, bound to a specific deployed contract.
func NewAdminV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AdminV2Transactor, error) {
	contract, err := bindAdminV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminV2Transactor{contract: contract}, nil
}

// NewAdminV2Filterer creates a new log filterer instance of AdminV2, bound to a specific deployed contract.
func NewAdminV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AdminV2Filterer, error) {
	contract, err := bindAdminV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminV2Filterer{contract: contract}, nil
}

// bindAdminV2 binds a generic wrapper to an already deployed contract.
func bindAdminV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdminV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminV2 *AdminV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdminV2.Contract.AdminV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminV2 *AdminV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminV2.Contract.AdminV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminV2 *AdminV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminV2.Contract.AdminV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminV2 *AdminV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdminV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminV2 *AdminV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminV2 *AdminV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminV2.Contract.contract.Transact(opts, method, params...)
}

// CANCELERROLE is a free data retrieval call binding the contract method 0x02d2ff66.
//
// Solidity: function CANCELER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Caller) CANCELERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "CANCELER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELERROLE is a free data retrieval call binding the contract method 0x02d2ff66.
//
// Solidity: function CANCELER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Session) CANCELERROLE() ([32]byte, error) {
	return _AdminV2.Contract.CANCELERROLE(&_AdminV2.CallOpts)
}

// CANCELERROLE is a free data retrieval call binding the contract method 0x02d2ff66.
//
// Solidity: function CANCELER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2CallerSession) CANCELERROLE() ([32]byte, error) {
	return _AdminV2.Contract.CANCELERROLE(&_AdminV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _AdminV2.Contract.DEFAULTADMINROLE(&_AdminV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AdminV2.Contract.DEFAULTADMINROLE(&_AdminV2.CallOpts)
}

// DEFAULTCANCELDELAY is a free data retrieval call binding the contract method 0x930ac180.
//
// Solidity: function DEFAULT_CANCEL_DELAY() view returns(uint256)
func (_AdminV2 *AdminV2Caller) DEFAULTCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "DEFAULT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTCANCELDELAY is a free data retrieval call binding the contract method 0x930ac180.
//
// Solidity: function DEFAULT_CANCEL_DELAY() view returns(uint256)
func (_AdminV2 *AdminV2Session) DEFAULTCANCELDELAY() (*big.Int, error) {
	return _AdminV2.Contract.DEFAULTCANCELDELAY(&_AdminV2.CallOpts)
}

// DEFAULTCANCELDELAY is a free data retrieval call binding the contract method 0x930ac180.
//
// Solidity: function DEFAULT_CANCEL_DELAY() view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) DEFAULTCANCELDELAY() (*big.Int, error) {
	return _AdminV2.Contract.DEFAULTCANCELDELAY(&_AdminV2.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_AdminV2 *AdminV2Caller) FEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_AdminV2 *AdminV2Session) FEEBPS() (*big.Int, error) {
	return _AdminV2.Contract.FEEBPS(&_AdminV2.CallOpts)
}

// FEEBPS is a free data retrieval call binding the contract method 0xbf333f2c.
//
// Solidity: function FEE_BPS() view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) FEEBPS() (*big.Int, error) {
	return _AdminV2.Contract.FEEBPS(&_AdminV2.CallOpts)
}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_AdminV2 *AdminV2Caller) FEERATEMAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "FEE_RATE_MAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_AdminV2 *AdminV2Session) FEERATEMAX() (*big.Int, error) {
	return _AdminV2.Contract.FEERATEMAX(&_AdminV2.CallOpts)
}

// FEERATEMAX is a free data retrieval call binding the contract method 0x0f5f6ed7.
//
// Solidity: function FEE_RATE_MAX() view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) FEERATEMAX() (*big.Int, error) {
	return _AdminV2.Contract.FEERATEMAX(&_AdminV2.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Caller) GOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Session) GOVERNORROLE() ([32]byte, error) {
	return _AdminV2.Contract.GOVERNORROLE(&_AdminV2.CallOpts)
}

// GOVERNORROLE is a free data retrieval call binding the contract method 0xccc57490.
//
// Solidity: function GOVERNOR_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2CallerSession) GOVERNORROLE() ([32]byte, error) {
	return _AdminV2.Contract.GOVERNORROLE(&_AdminV2.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Caller) GUARDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "GUARD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Session) GUARDROLE() ([32]byte, error) {
	return _AdminV2.Contract.GUARDROLE(&_AdminV2.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2CallerSession) GUARDROLE() ([32]byte, error) {
	return _AdminV2.Contract.GUARDROLE(&_AdminV2.CallOpts)
}

// MINCANCELDELAY is a free data retrieval call binding the contract method 0x922b7487.
//
// Solidity: function MIN_CANCEL_DELAY() view returns(uint256)
func (_AdminV2 *AdminV2Caller) MINCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "MIN_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCANCELDELAY is a free data retrieval call binding the contract method 0x922b7487.
//
// Solidity: function MIN_CANCEL_DELAY() view returns(uint256)
func (_AdminV2 *AdminV2Session) MINCANCELDELAY() (*big.Int, error) {
	return _AdminV2.Contract.MINCANCELDELAY(&_AdminV2.CallOpts)
}

// MINCANCELDELAY is a free data retrieval call binding the contract method 0x922b7487.
//
// Solidity: function MIN_CANCEL_DELAY() view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) MINCANCELDELAY() (*big.Int, error) {
	return _AdminV2.Contract.MINCANCELDELAY(&_AdminV2.CallOpts)
}

// NATIVEGASTOKEN is a free data retrieval call binding the contract method 0x0f862f1e.
//
// Solidity: function NATIVE_GAS_TOKEN() view returns(address)
func (_AdminV2 *AdminV2Caller) NATIVEGASTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "NATIVE_GAS_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEGASTOKEN is a free data retrieval call binding the contract method 0x0f862f1e.
//
// Solidity: function NATIVE_GAS_TOKEN() view returns(address)
func (_AdminV2 *AdminV2Session) NATIVEGASTOKEN() (common.Address, error) {
	return _AdminV2.Contract.NATIVEGASTOKEN(&_AdminV2.CallOpts)
}

// NATIVEGASTOKEN is a free data retrieval call binding the contract method 0x0f862f1e.
//
// Solidity: function NATIVE_GAS_TOKEN() view returns(address)
func (_AdminV2 *AdminV2CallerSession) NATIVEGASTOKEN() (common.Address, error) {
	return _AdminV2.Contract.NATIVEGASTOKEN(&_AdminV2.CallOpts)
}

// PROVERROLE is a free data retrieval call binding the contract method 0xdc9a4ef6.
//
// Solidity: function PROVER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Caller) PROVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "PROVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVERROLE is a free data retrieval call binding the contract method 0xdc9a4ef6.
//
// Solidity: function PROVER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Session) PROVERROLE() ([32]byte, error) {
	return _AdminV2.Contract.PROVERROLE(&_AdminV2.CallOpts)
}

// PROVERROLE is a free data retrieval call binding the contract method 0xdc9a4ef6.
//
// Solidity: function PROVER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2CallerSession) PROVERROLE() ([32]byte, error) {
	return _AdminV2.Contract.PROVERROLE(&_AdminV2.CallOpts)
}

// QUOTERROLE is a free data retrieval call binding the contract method 0x7ebe815c.
//
// Solidity: function QUOTER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Caller) QUOTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "QUOTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// QUOTERROLE is a free data retrieval call binding the contract method 0x7ebe815c.
//
// Solidity: function QUOTER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2Session) QUOTERROLE() ([32]byte, error) {
	return _AdminV2.Contract.QUOTERROLE(&_AdminV2.CallOpts)
}

// QUOTERROLE is a free data retrieval call binding the contract method 0x7ebe815c.
//
// Solidity: function QUOTER_ROLE() view returns(bytes32)
func (_AdminV2 *AdminV2CallerSession) QUOTERROLE() ([32]byte, error) {
	return _AdminV2.Contract.QUOTERROLE(&_AdminV2.CallOpts)
}

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_AdminV2 *AdminV2Caller) CancelDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "cancelDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_AdminV2 *AdminV2Session) CancelDelay() (*big.Int, error) {
	return _AdminV2.Contract.CancelDelay(&_AdminV2.CallOpts)
}

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) CancelDelay() (*big.Int, error) {
	return _AdminV2.Contract.CancelDelay(&_AdminV2.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_AdminV2 *AdminV2Caller) ChainGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "chainGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_AdminV2 *AdminV2Session) ChainGasAmount() (*big.Int, error) {
	return _AdminV2.Contract.ChainGasAmount(&_AdminV2.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) ChainGasAmount() (*big.Int, error) {
	return _AdminV2.Contract.ChainGasAmount(&_AdminV2.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdminV2 *AdminV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdminV2 *AdminV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AdminV2.Contract.GetRoleAdmin(&_AdminV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdminV2 *AdminV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AdminV2.Contract.GetRoleAdmin(&_AdminV2.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AdminV2 *AdminV2Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AdminV2 *AdminV2Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AdminV2.Contract.GetRoleMember(&_AdminV2.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AdminV2 *AdminV2CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AdminV2.Contract.GetRoleMember(&_AdminV2.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AdminV2 *AdminV2Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AdminV2 *AdminV2Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AdminV2.Contract.GetRoleMemberCount(&_AdminV2.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AdminV2.Contract.GetRoleMemberCount(&_AdminV2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdminV2 *AdminV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdminV2 *AdminV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AdminV2.Contract.HasRole(&_AdminV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdminV2 *AdminV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AdminV2.Contract.HasRole(&_AdminV2.CallOpts, role, account)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_AdminV2 *AdminV2Caller) ProtocolFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "protocolFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_AdminV2 *AdminV2Session) ProtocolFeeRate() (*big.Int, error) {
	return _AdminV2.Contract.ProtocolFeeRate(&_AdminV2.CallOpts)
}

// ProtocolFeeRate is a free data retrieval call binding the contract method 0x58f85880.
//
// Solidity: function protocolFeeRate() view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) ProtocolFeeRate() (*big.Int, error) {
	return _AdminV2.Contract.ProtocolFeeRate(&_AdminV2.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_AdminV2 *AdminV2Caller) ProtocolFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "protocolFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_AdminV2 *AdminV2Session) ProtocolFees(arg0 common.Address) (*big.Int, error) {
	return _AdminV2.Contract.ProtocolFees(&_AdminV2.CallOpts, arg0)
}

// ProtocolFees is a free data retrieval call binding the contract method 0xdcf844a7.
//
// Solidity: function protocolFees(address ) view returns(uint256)
func (_AdminV2 *AdminV2CallerSession) ProtocolFees(arg0 common.Address) (*big.Int, error) {
	return _AdminV2.Contract.ProtocolFees(&_AdminV2.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdminV2 *AdminV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AdminV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdminV2 *AdminV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AdminV2.Contract.SupportsInterface(&_AdminV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdminV2 *AdminV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AdminV2.Contract.SupportsInterface(&_AdminV2.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdminV2 *AdminV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdminV2 *AdminV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminV2.Contract.GrantRole(&_AdminV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdminV2 *AdminV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminV2.Contract.GrantRole(&_AdminV2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdminV2 *AdminV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdminV2.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdminV2 *AdminV2Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdminV2.Contract.RenounceRole(&_AdminV2.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdminV2 *AdminV2TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdminV2.Contract.RenounceRole(&_AdminV2.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdminV2 *AdminV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdminV2 *AdminV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminV2.Contract.RevokeRole(&_AdminV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdminV2 *AdminV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdminV2.Contract.RevokeRole(&_AdminV2.TransactOpts, role, account)
}

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_AdminV2 *AdminV2Transactor) SetCancelDelay(opts *bind.TransactOpts, newCancelDelay *big.Int) (*types.Transaction, error) {
	return _AdminV2.contract.Transact(opts, "setCancelDelay", newCancelDelay)
}

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_AdminV2 *AdminV2Session) SetCancelDelay(newCancelDelay *big.Int) (*types.Transaction, error) {
	return _AdminV2.Contract.SetCancelDelay(&_AdminV2.TransactOpts, newCancelDelay)
}

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_AdminV2 *AdminV2TransactorSession) SetCancelDelay(newCancelDelay *big.Int) (*types.Transaction, error) {
	return _AdminV2.Contract.SetCancelDelay(&_AdminV2.TransactOpts, newCancelDelay)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_AdminV2 *AdminV2Transactor) SetProtocolFeeRate(opts *bind.TransactOpts, newFeeRate *big.Int) (*types.Transaction, error) {
	return _AdminV2.contract.Transact(opts, "setProtocolFeeRate", newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_AdminV2 *AdminV2Session) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _AdminV2.Contract.SetProtocolFeeRate(&_AdminV2.TransactOpts, newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_AdminV2 *AdminV2TransactorSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _AdminV2.Contract.SetProtocolFeeRate(&_AdminV2.TransactOpts, newFeeRate)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_AdminV2 *AdminV2Transactor) SweepProtocolFees(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _AdminV2.contract.Transact(opts, "sweepProtocolFees", token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_AdminV2 *AdminV2Session) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _AdminV2.Contract.SweepProtocolFees(&_AdminV2.TransactOpts, token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_AdminV2 *AdminV2TransactorSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _AdminV2.Contract.SweepProtocolFees(&_AdminV2.TransactOpts, token, recipient)
}

// AdminV2CancelDelayUpdatedIterator is returned from FilterCancelDelayUpdated and is used to iterate over the raw logs and unpacked data for CancelDelayUpdated events raised by the AdminV2 contract.
type AdminV2CancelDelayUpdatedIterator struct {
	Event *AdminV2CancelDelayUpdated // Event containing the contract specifics and raw log

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
func (it *AdminV2CancelDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminV2CancelDelayUpdated)
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
		it.Event = new(AdminV2CancelDelayUpdated)
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
func (it *AdminV2CancelDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminV2CancelDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminV2CancelDelayUpdated represents a CancelDelayUpdated event raised by the AdminV2 contract.
type AdminV2CancelDelayUpdated struct {
	OldCancelDelay *big.Int
	NewCancelDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCancelDelayUpdated is a free log retrieval operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_AdminV2 *AdminV2Filterer) FilterCancelDelayUpdated(opts *bind.FilterOpts) (*AdminV2CancelDelayUpdatedIterator, error) {

	logs, sub, err := _AdminV2.contract.FilterLogs(opts, "CancelDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &AdminV2CancelDelayUpdatedIterator{contract: _AdminV2.contract, event: "CancelDelayUpdated", logs: logs, sub: sub}, nil
}

// WatchCancelDelayUpdated is a free log subscription operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_AdminV2 *AdminV2Filterer) WatchCancelDelayUpdated(opts *bind.WatchOpts, sink chan<- *AdminV2CancelDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _AdminV2.contract.WatchLogs(opts, "CancelDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminV2CancelDelayUpdated)
				if err := _AdminV2.contract.UnpackLog(event, "CancelDelayUpdated", log); err != nil {
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

// ParseCancelDelayUpdated is a log parse operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_AdminV2 *AdminV2Filterer) ParseCancelDelayUpdated(log types.Log) (*AdminV2CancelDelayUpdated, error) {
	event := new(AdminV2CancelDelayUpdated)
	if err := _AdminV2.contract.UnpackLog(event, "CancelDelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminV2FeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the AdminV2 contract.
type AdminV2FeeRateUpdatedIterator struct {
	Event *AdminV2FeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *AdminV2FeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminV2FeeRateUpdated)
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
		it.Event = new(AdminV2FeeRateUpdated)
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
func (it *AdminV2FeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminV2FeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminV2FeeRateUpdated represents a FeeRateUpdated event raised by the AdminV2 contract.
type AdminV2FeeRateUpdated struct {
	OldFeeRate *big.Int
	NewFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_AdminV2 *AdminV2Filterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*AdminV2FeeRateUpdatedIterator, error) {

	logs, sub, err := _AdminV2.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &AdminV2FeeRateUpdatedIterator{contract: _AdminV2.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_AdminV2 *AdminV2Filterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *AdminV2FeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _AdminV2.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminV2FeeRateUpdated)
				if err := _AdminV2.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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
func (_AdminV2 *AdminV2Filterer) ParseFeeRateUpdated(log types.Log) (*AdminV2FeeRateUpdated, error) {
	event := new(AdminV2FeeRateUpdated)
	if err := _AdminV2.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminV2FeesSweptIterator is returned from FilterFeesSwept and is used to iterate over the raw logs and unpacked data for FeesSwept events raised by the AdminV2 contract.
type AdminV2FeesSweptIterator struct {
	Event *AdminV2FeesSwept // Event containing the contract specifics and raw log

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
func (it *AdminV2FeesSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminV2FeesSwept)
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
		it.Event = new(AdminV2FeesSwept)
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
func (it *AdminV2FeesSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminV2FeesSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminV2FeesSwept represents a FeesSwept event raised by the AdminV2 contract.
type AdminV2FeesSwept struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesSwept is a free log retrieval operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_AdminV2 *AdminV2Filterer) FilterFeesSwept(opts *bind.FilterOpts) (*AdminV2FeesSweptIterator, error) {

	logs, sub, err := _AdminV2.contract.FilterLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return &AdminV2FeesSweptIterator{contract: _AdminV2.contract, event: "FeesSwept", logs: logs, sub: sub}, nil
}

// WatchFeesSwept is a free log subscription operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_AdminV2 *AdminV2Filterer) WatchFeesSwept(opts *bind.WatchOpts, sink chan<- *AdminV2FeesSwept) (event.Subscription, error) {

	logs, sub, err := _AdminV2.contract.WatchLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminV2FeesSwept)
				if err := _AdminV2.contract.UnpackLog(event, "FeesSwept", log); err != nil {
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
func (_AdminV2 *AdminV2Filterer) ParseFeesSwept(log types.Log) (*AdminV2FeesSwept, error) {
	event := new(AdminV2FeesSwept)
	if err := _AdminV2.contract.UnpackLog(event, "FeesSwept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AdminV2 contract.
type AdminV2RoleAdminChangedIterator struct {
	Event *AdminV2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AdminV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminV2RoleAdminChanged)
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
		it.Event = new(AdminV2RoleAdminChanged)
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
func (it *AdminV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminV2RoleAdminChanged represents a RoleAdminChanged event raised by the AdminV2 contract.
type AdminV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AdminV2 *AdminV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AdminV2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _AdminV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AdminV2RoleAdminChangedIterator{contract: _AdminV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AdminV2 *AdminV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AdminV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AdminV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminV2RoleAdminChanged)
				if err := _AdminV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AdminV2 *AdminV2Filterer) ParseRoleAdminChanged(log types.Log) (*AdminV2RoleAdminChanged, error) {
	event := new(AdminV2RoleAdminChanged)
	if err := _AdminV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AdminV2 contract.
type AdminV2RoleGrantedIterator struct {
	Event *AdminV2RoleGranted // Event containing the contract specifics and raw log

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
func (it *AdminV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminV2RoleGranted)
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
		it.Event = new(AdminV2RoleGranted)
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
func (it *AdminV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminV2RoleGranted represents a RoleGranted event raised by the AdminV2 contract.
type AdminV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdminV2 *AdminV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdminV2RoleGrantedIterator, error) {

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

	logs, sub, err := _AdminV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdminV2RoleGrantedIterator{contract: _AdminV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdminV2 *AdminV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AdminV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AdminV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminV2RoleGranted)
				if err := _AdminV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AdminV2 *AdminV2Filterer) ParseRoleGranted(log types.Log) (*AdminV2RoleGranted, error) {
	event := new(AdminV2RoleGranted)
	if err := _AdminV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AdminV2 contract.
type AdminV2RoleRevokedIterator struct {
	Event *AdminV2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *AdminV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminV2RoleRevoked)
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
		it.Event = new(AdminV2RoleRevoked)
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
func (it *AdminV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminV2RoleRevoked represents a RoleRevoked event raised by the AdminV2 contract.
type AdminV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdminV2 *AdminV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdminV2RoleRevokedIterator, error) {

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

	logs, sub, err := _AdminV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdminV2RoleRevokedIterator{contract: _AdminV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdminV2 *AdminV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AdminV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AdminV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminV2RoleRevoked)
				if err := _AdminV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AdminV2 *AdminV2Filterer) ParseRoleRevoked(log types.Log) (*AdminV2RoleRevoked, error) {
	event := new(AdminV2RoleRevoked)
	if err := _AdminV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTransactionV2LibMetaData contains all meta data concerning the BridgeTransactionV2Lib contract.
var BridgeTransactionV2LibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BridgeTransactionV2__InvalidEncodedTx\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"BridgeTransactionV2__UnsupportedVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220736bb5e8e495fb3a776c3cad8b30577c46279a088f4f89b0d509badbef3da01f64736f6c63430008180033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cb930ba14bc69bdf61a9326099568898423db2bde70050effda1b23aa9e9dc9e64736f6c63430008180033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeTransactionV2__InvalidEncodedTx\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"BridgeTransactionV2__UnsupportedVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CancelDelayBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineNotExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExclusivityParamsIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExclusivityPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRateAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MulticallTarget__UndeterminedRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecipientIncorrectReturnValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecipientNoReturnValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StatusIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZapDataLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZapNativeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"}],\"name\":\"BridgeQuoteDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCancelDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCancelDelay\",\"type\":\"uint256\"}],\"name\":\"CancelDelayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCELER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_RATE_MAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ZAP_DATA_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEADLINE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_GAS_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"quoteRelayer\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"quoteExclusivitySeconds\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeParamsV2\",\"name\":\"paramsV2\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeProofs\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"timestamp\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelayDetails\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"blockNumber\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"blockTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelays\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeStatuses\",\"outputs\":[{\"internalType\":\"enumIFastBridgeV2.BridgeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeTxDetails\",\"outputs\":[{\"internalType\":\"enumIFastBridgeV2.BridgeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint56\",\"name\":\"proofBlockTimestamp\",\"type\":\"uint56\"},{\"internalType\":\"address\",\"name\":\"proofRelayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransactionV2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"exclusivityRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exclusivityEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeTransactionV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"ignoreReverts\",\"type\":\"bool\"}],\"name\":\"multicallNoResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"ignoreReverts\",\"type\":\"bool\"}],\"name\":\"multicallWithResults\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticallTarget.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"senderNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCancelDelay\",\"type\":\"uint256\"}],\"name\":\"setCancelDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"02d2ff66": "CANCELER_ROLE()",
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"930ac180": "DEFAULT_CANCEL_DELAY()",
		"a5bbe22b": "DISPUTE_PERIOD()",
		"bf333f2c": "FEE_BPS()",
		"0f5f6ed7": "FEE_RATE_MAX()",
		"ccc57490": "GOVERNOR_ROLE()",
		"03ed0ee5": "GUARD_ROLE()",
		"54eff068": "MAX_ZAP_DATA_LENGTH()",
		"922b7487": "MIN_CANCEL_DELAY()",
		"820688d5": "MIN_DEADLINE_PERIOD()",
		"0f862f1e": "NATIVE_GAS_TOKEN()",
		"dc9a4ef6": "PROVER_ROLE()",
		"7ebe815c": "QUOTER_ROLE()",
		"45851694": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256))",
		"bfc7c607": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256),(address,int256,bytes,uint256,bytes))",
		"91ad5039": "bridgeProofs(bytes32)",
		"c79371b1": "bridgeRelayDetails(bytes32)",
		"8379a24f": "bridgeRelays(bytes32)",
		"051287bc": "bridgeStatuses(bytes32)",
		"63787e52": "bridgeTxDetails(bytes32)",
		"aa9641ab": "canClaim(bytes32,address)",
		"3c5beeb4": "cancel(bytes)",
		"638a0f09": "cancelDelay()",
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
		"1ea327c5": "setCancelDelay(uint256)",
		"b13aa2d6": "setProtocolFeeRate(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"06f333f2": "sweepProtocolFees(address,address)",
	},
	Bin: "0x60e06040526000608081905260a0523480156200001b57600080fd5b50604051620047ee380380620047ee8339810160408190526200003e9162000215565b806200004c60008262000067565b506200005b62015180620000a4565b50504360c05262000240565b6000806200007684846200010d565b905080156200009b576000848152600160205260409020620000999084620001bb565b505b90505b92915050565b610e10811015620000c857604051630e0ea5c760e01b815260040160405180910390fd5b600480549082905560408051828152602081018490527f909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95910160405180910390a15050565b6000828152602081815260408083206001600160a01b038516845290915281205460ff16620001b2576000838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055620001693390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016200009e565b5060006200009e565b60006200009b836001600160a01b0384166000818152600183016020526040812054620001b2575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200009e565b6000602082840312156200022857600080fd5b81516001600160a01b03811681146200009b57600080fd5b60805160a05160c05161457e62000270600039600061094b015260006109ec01526000610bec015261457e6000f3fe60806040526004361061034a5760003560e01c80638379a24f116101bb578063ac11fb1a116100f7578063c79371b111610095578063d547741f1161006f578063d547741f14610b59578063dc9a4ef614610b79578063dcf844a714610bad578063e00a83e014610bda57600080fd5b8063c79371b114610a78578063ca15c87314610b05578063ccc5749014610b2557600080fd5b8063b13aa2d6116100d1578063b13aa2d614610a0e578063bf333f2c14610a2e578063bfc7c60714610a45578063c63ff8dd14610a5857600080fd5b8063ac11fb1a1461098d578063add98c70146109ba578063affed0e0146109da57600080fd5b8063922b748711610164578063a217fddf1161013e578063a217fddf14610924578063a3ec191a14610939578063a5bbe22b14610769578063aa9641ab1461096d57600080fd5b8063922b7487146108e4578063930ac180146108fa5780639c9545f01461091157600080fd5b80639010d07c116101955780639010d07c146107fa57806391ad50391461081a57806391d14854146108a057600080fd5b80638379a24f1461077f578063886d36ff146107c75780638f0d6f17146107e757600080fd5b8063385c1d2f1161028a57806358f858801161023357806363787e521161020d57806363787e52146106a5578063638a0f091461071f5780637ebe815c14610735578063820688d51461076957600080fd5b806358f85880146106425780635aa6ccba146106585780635eb7d9461461068557600080fd5b806341fcb6121161026457806341fcb612146105f9578063458516941461061957806354eff0681461062c57600080fd5b8063385c1d2f1461058c5780633c5beeb4146105b95780633f61331d146105d957600080fd5b80630f862f1e116102f7578063248a9ca3116102d1578063248a9ca3146104ef578063295710ff1461051f5780632f2ff15d1461054c57806336568abe1461056c57600080fd5b80630f862f1e1461046f57806318e4357d146104af5780631ea327c5146104cf57600080fd5b8063051287bc11610328578063051287bc146103fa57806306f333f2146104375780630f5f6ed71461045957600080fd5b806301ffc9a71461034f57806302d2ff661461038457806303ed0ee5146103c6575b600080fd5b34801561035b57600080fd5b5061036f61036a36600461352a565b610c0e565b60405190151581526020015b60405180910390f35b34801561039057600080fd5b506103b87febfdca8e46c0b8dacf9989ee613e35727eadd20a1d5e5ad01a53968c7e5fe07a81565b60405190815260200161037b565b3480156103d257600080fd5b506103b87f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d81565b34801561040657600080fd5b5061042a61041536600461356c565b60009081526005602052604090205460ff1690565b60405161037b91906135ef565b34801561044357600080fd5b50610457610452366004613622565b610c6a565b005b34801561046557600080fd5b506103b861271081565b34801561047b57600080fd5b5061049773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b6040516001600160a01b03909116815260200161037b565b3480156104bb57600080fd5b506104576104ca36600461365b565b610d77565b3480156104db57600080fd5b506104576104ea36600461356c565b610e8c565b3480156104fb57600080fd5b506103b861050a36600461356c565b60009081526020819052604090206001015490565b34801561052b57600080fd5b506103b861053a366004613694565b60076020526000908152604090205481565b34801561055857600080fd5b506104576105673660046136b1565b610ec3565b34801561057857600080fd5b506104576105873660046136b1565b610ee8565b34801561059857600080fd5b506105ac6105a73660046136ef565b610f34565b60405161037b91906137ba565b3480156105c557600080fd5b506104576105d4366004613899565b6110c2565b3480156105e557600080fd5b506104576105f43660046136ef565b6112c4565b34801561060557600080fd5b506104576106143660046138db565b61136e565b610457610627366004613adc565b6115e0565b34801561063857600080fd5b506103b861ffff81565b34801561064e57600080fd5b506103b860025481565b34801561066457600080fd5b50610678610673366004613899565b61163d565b60405161037b9190613af9565b34801561069157600080fd5b506104576106a0366004613899565b61170a565b3480156106b157600080fd5b5061070f6106c036600461356c565b60056020526000908152604090205460ff811690610100810463ffffffff169065010000000000810466ffffffffffffff16906c0100000000000000000000000090046001600160a01b031684565b60405161037b9493929190613c16565b34801561072b57600080fd5b506103b860045481565b34801561074157600080fd5b506103b87f9a04aea0a349253cc7277afafdf6ead6729a3972a47ffb40eaef2c93d4e1bfea81565b34801561077557600080fd5b506103b861070881565b34801561078b57600080fd5b5061036f61079a36600461356c565b6000908152600660205260409020546c0100000000000000000000000090046001600160a01b0316151590565b3480156107d357600080fd5b506104576107e2366004613c57565b611714565b6104576107f5366004613899565b611740565b34801561080657600080fd5b50610497610815366004613ca3565b61174b565b34801561082657600080fd5b5061087461083536600461356c565b60009081526005602052604090205465010000000000810466ffffffffffffff16916c010000000000000000000000009091046001600160a01b031690565b604080516bffffffffffffffffffffffff90931683526001600160a01b0390911660208301520161037b565b3480156108ac57600080fd5b5061036f6108bb3660046136b1565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156108f057600080fd5b506103b8610e1081565b34801561090657600080fd5b506103b86201518081565b61045761091f3660046138db565b611763565b34801561093057600080fd5b506103b8600081565b34801561094557600080fd5b506103b87f000000000000000000000000000000000000000000000000000000000000000081565b34801561097957600080fd5b5061036f6109883660046136b1565b611a1d565b34801561099957600080fd5b506109ad6109a8366004613899565b611af4565b60405161037b9190613cc5565b3480156109c657600080fd5b506104576109d536600461356c565b611ca8565b3480156109e657600080fd5b506103b87f000000000000000000000000000000000000000000000000000000000000000081565b348015610a1a57600080fd5b50610457610a2936600461356c565b611df2565b348015610a3a57600080fd5b506103b8620f424081565b610457610a53366004613e29565b611e9e565b348015610a6457600080fd5b50610457610a73366004613899565b61212d565b348015610a8457600080fd5b50610ad7610a9336600461356c565b60066020526000908152604090205465ffffffffffff8082169166010000000000008104909116906c0100000000000000000000000090046001600160a01b031683565b6040805165ffffffffffff94851681529390921660208401526001600160a01b03169082015260600161037b565b348015610b1157600080fd5b506103b8610b2036600461356c565b612139565b348015610b3157600080fd5b506103b87f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f5581565b348015610b6557600080fd5b50610457610b743660046136b1565b612150565b348015610b8557600080fd5b506103b87f60044782a422109ac08a415e44260ad5d3bad63d96ebe1fac0363399642ee3f281565b348015610bb957600080fd5b506103b8610bc8366004613694565b60036020526000908152604090205481565b348015610be657600080fd5b506103b87f000000000000000000000000000000000000000000000000000000000000000081565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f000000000000000000000000000000000000000000000000000000001480610c645750610c6482612175565b92915050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610c948161220c565b6001600160a01b03831660009081526003602052604081205490819003610cbb5750505050565b6001600160a01b038481166000818152600360209081526040808320929092558151928352928616928201929092529081018290527f244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd9060600160405180910390a17fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03851601610d5c57610d578382612216565b610d70565b610d706001600160a01b03851684836122e4565b505b505050565b7f60044782a422109ac08a415e44260ad5d3bad63d96ebe1fac0363399642ee3f2610da18161220c565b60008481526005602052604090206001815460ff166004811115610dc757610dc7613585565b14610dfe576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80546001600160a01b0384166c0100000000000000000000000081026bffffffffffffffffffffffff66ffffffffffffff421665010000000000021664ffffffff009093169290921791909117600217825560405185815286907f4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e49060200160405180910390a35050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55610eb68161220c565b610ebf82612358565b5050565b600082815260208190526040902060010154610ede8161220c565b610d7083836123d9565b6001600160a01b0381163314610f2a576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d728282612406565b60608267ffffffffffffffff811115610f4f57610f4f613927565b604051908082528060200260200182016040528015610f9557816020015b604080518082019091526000815260606020820152815260200190600190039081610f6d5790505b50905060005b838110156110ba5730858583818110610fb657610fb6613ef7565b9050602002810190610fc89190613f26565b604051610fd6929190613f8b565b600060405180830381855af49150503d8060008114611011576040519150601f19603f3d011682016040523d82523d6000602084013e611016565b606091505b5083838151811061102957611029613ef7565b602002602001015160000184848151811061104657611046613ef7565b60200260200101516020018290528215151515815250505081818151811061107057611070613ef7565b602002602001015160000151158015611087575082155b156110b2576110b28282815181106110a1576110a1613ef7565b602002602001015160200151612433565b600101610f9b565b509392505050565b6110cc8282612475565b600082826040516110de929190613f8b565b604080519182900390912060008181526005602052919091209091506001815460ff16600481111561111257611112613585565b14611149576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081527f4527d8b28c468984933d33ae052f26b21c15ca0e7fdec762bba08e540e237001602052604090205460ba8501359060ff16611195576004546111929082613fca565b90505b8042116111ce576040517fe15ff9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166004178255600a850135606090811c906032870135901c600061121d609a890135605a8a0135613fca565b604080516001600160a01b03858116825260208201849052825193945086169289927fb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958928290030190a37fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b038316016112a6576112a18382612216565b6112ba565b6112ba6001600160a01b03831684836122e4565b5050505050505050565b60005b82811015610d7057600080308686858181106112e5576112e5613ef7565b90506020028101906112f79190613f26565b604051611305929190613f8b565b600060405180830381855af49150503d8060008114611340576040519150601f19603f3d011682016040523d82523d6000602084013e611345565b606091505b509150915081158015611356575083155b156113645761136481612433565b50506001016112c7565b6113788383612475565b6000838360405161138a929190613f8b565b604080519182900390912060008181526005602052919091208054919250906c0100000000000000000000000081046001600160a01b03169060ff81169065010000000000900466ffffffffffffff1660028260048111156113ee576113ee613585565b14611425576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107084282900366ffffffffffffff161161146c576040517f1992d0bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038616611482578295506114c4565b6001600160a01b03831633146114c4576040517f4af43a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166003178455603288013560601c605a890135609a8a01358015611533576001600160a01b0383166000908152600360205260408120805483929061152d908490613fca565b90915550505b604080516001600160a01b03858116825260208201859052808c1692908916918b917f582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678910160405180910390a47fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b038416016115bf576115ba8983612216565b6115d3565b6115d36001600160a01b0384168a846122e4565b5050505050505050505050565b61163a816040518060a0016040528060006001600160a01b03168152602001600081526020016040518060200160405280600081525081526020016000815260200160405180602001604052806000815250815250611e9e565b50565b6116ef604051806101e00160405280600063ffffffff168152602001600063ffffffff16815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016000815260200160008152602001606081525090565b6116f98383612475565b61170383836124f6565b9392505050565b610ebf82826110c2565b61171e8383612475565b610d728383604051611731929190613f8b565b60405180910390208233610d77565b610ebf828233611763565b60008281526001602052604081206117039083612698565b61176d8383612475565b6000838360405161177f929190613f8b565b60405180910390209050611795848483856126a4565b6040805160608101825265ffffffffffff438116825242811660208084019182526001600160a01b038088168587019081526000888152600690935295822094518554935196519091166c01000000000000000000000000026bffffffffffffffffffffffff9685166601000000000000027fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009094169190941617919091179390931617905561184985601e013560601c90565b9050604685013560601c607a86013561012e8701356001600160a01b03808516908716867ff8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c60028c013560e01c6040805163ffffffff92909216825260328e0135606090811c60208401526001600160a01b038a1683830152605a8f0135908301526080820188905260a08201879052519081900360c00190a47fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b0384160161198957801561194b576040517f846dedc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b813414611984576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119d7565b8034146119c2576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119d76001600160a01b038416338685612839565b3660006119e48a8a612872565b90925090508015611a01576119fc868686858561288e565b611a11565b3415611a1157611a118634612216565b50505050505050505050565b60008281526005602052604081206002815460ff166004811115611a4357611a43613585565b14611a7a576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80546001600160a01b038481166c010000000000000000000000009092041614611ad0576040517f4af43a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b546107086501000000000090910466ffffffffffffff908116420316119392505050565b6040805161018081018252600080825260208201819052818301819052606082018190526080820181905260a0820181905260c0820181905260e0820181905261010082018190526101208201819052610140820181905261016082015290517f5aa6ccba0000000000000000000000000000000000000000000000000000000081523090635aa6ccba90611b8f9086908690600401614008565b600060405180830381865afa925050508015611bcd57506040513d6000823e601f3d908101601f19168201604052611bca9190810190614077565b60015b611be457611bdd828401846141ac565b9050610c64565b604051806101800160405280826000015163ffffffff168152602001826020015163ffffffff16815260200182604001516001600160a01b0316815260200182606001516001600160a01b0316815260200182608001516001600160a01b031681526020018260a001516001600160a01b031681526020018260c0015181526020018260e0015181526020018261010001518152602001826101a0015160001415151581526020018261012001518152602001826101400151815250915050610c64565b7f043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639d611cd28161220c565b600082815260056020526040902080546c0100000000000000000000000081046001600160a01b03169060ff81169065010000000000900466ffffffffffffff166002826004811115611d2757611d27613585565b14611d5e576040517f4145817200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107084282900366ffffffffffffff161115611da6576040517f3e908aac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835464ffffffff001660011784556040516001600160a01b0384169087907f0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad490600090a3505050505050565b7f7935bd0ae54bc31f548c14dba4d37c5c64b3f8ca900cb468fb8abd54d5894f55611e1c8161220c565b612710821115611e58576040517f9b569b8100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280549083905560408051828152602081018590527f14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957910160405180910390a1505050565b80516000906001600160a01b031615611ec3576020820151611ec09042614278565b90505b611ece8383836129ea565b6000611ee284606001518560a00151612c6e565b90506000806002541115611f1b57620f424060025483611f0291906142a0565b611f0c91906142b7565b9050611f1881836142f2565b91505b6000612024604051806101e001604052804663ffffffff168152602001886000015163ffffffff16815260200188602001516001600160a01b0316815260200188604001516001600160a01b0316815260200188606001516001600160a01b0316815260200188608001516001600160a01b031681526020018581526020018860c0015181526020018481526020018861010001518152602001600760008a602001516001600160a01b03166001600160a01b031681526020019081526020016000206000815480929190611fef90614305565b90915550815287516001600160a01b031660208201526040810187905260608089015190820152608080890151910152612e81565b805160208083019190912060008181526005835260409081902080548b5163ffffffff8116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000090921691909117600117909155928a01516060808c015160808d015160c08e0151928d0151945197985094966001600160a01b039093169587957f120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a956120e0958b959394938e9290919015159061433d565b60405180910390a3807f3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e876040015160405161211c9190614393565b60405180910390a250505050505050565b610ebf8282600061136e565b6000818152600160205260408120610c6490612fc6565b60008281526020819052604090206001015461216b8161220c565b610d708383612406565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610c6457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610c64565b61163a8133612fd0565b80471015612257576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024015b60405180910390fd5b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146122a4576040519150601f19603f3d011682016040523d82523d6000602084013e6122a9565b606091505b5050905080610d72576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516001600160a01b03838116602483015260448201839052610d7291859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061303c565b610e10811015612394576040517f0e0ea5c700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480549082905560408051828152602081018490527f909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95910160405180910390a15050565b6000806123e684846130b8565b905080156117035760008481526001602052604090206110ba9084613180565b6000806124138484613195565b905080156117035760008481526001602052604090206110ba9084613236565b8051156124435780518082602001fd5b6040517f5ead5a9d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61014e8110156124b1576040517f6ee0e17100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b813560f01c60028114610d72576040517f8bd977e700000000000000000000000000000000000000000000000000000000815261ffff8216600482015260240161224e565b6125a8604051806101e00160405280600063ffffffff168152602001600063ffffffff16815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016000815260200160008152602001606081525090565b600283013560e090811c82526006840135811c6020830152600a840135606090811c6040840152601e850135811c818401526032850135811c60808401526046850135811c60a0840152605a85013560c0840152607a85013591830191909152609a84013561010083015260ba84013561012083015260da84013561014083015260fa840135901c61016082015261010e83013561018082015261012e8301356101a08201526126588383612872565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506101c082015292915050565b6000611703838361324b565b6001600160a01b0381166126e4576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152600660205260409020546c0100000000000000000000000090046001600160a01b031615612743576040517fbef7bb7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600684013560e01c4614612783576040517f7029fdf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60ba8401354211156127c1576040517f559895a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60fa84013560601c80158015906127ea5750816001600160a01b0316816001600160a01b031614155b80156127fb575061010e8501354211155b15612832576040517f14be123200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6040516001600160a01b038481166024830152838116604483015260648201839052610d709186918216906323b872dd90608401612311565b3660006128838361014e81876143a6565b915091509250929050565b600061290986868686866040516024016128ab94939291906143d0565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe85e13dd0000000000000000000000000000000000000000000000000000000017905234613275565b90508051600003612946576040517f1a95ad2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051602014612981576040517ff3725cca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fe85e13dd000000000000000000000000000000000000000000000000000000006129ab826143f9565b146129e2576040517ff3725cca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050565b46836000015163ffffffff1603612a2d576040517f7029fdf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08301511580612a40575060c0830151155b15612a77576040517fe38820c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208301516001600160a01b03161580612a9c575060408301516001600160a01b0316155b15612ad3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608301516001600160a01b03161580612af8575060808301516001600160a01b0316155b15612b2f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612b3b61070842613fca565b8361010001511015612b79576040517f04b7fcc800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61ffff8260800151511115612bba576040517f177a70e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082015115801590612bed575060808301516001600160a01b031673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee145b15612c24576040517f846dedc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000811280612c37575082610100015181135b15610d72576040517f352807b900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03841601612ce057348214612cd9576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5034610c64565b3415612d18576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b03163b600003612d5c576040517f7f523fe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038416906370a0823190602401602060405180830381865afa158015612db9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ddd919061443e565b9050612df46001600160a01b038416333085612839565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281906001600160a01b038516906370a0823190602401602060405180830381865afa158015612e53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e77919061443e565b61170391906142f2565b8051602080830151604080850151606086810151608088015160a089015160c08a015195517e02000000000000000000000000000000000000000000000000000000000000988101989098527fffffffff0000000000000000000000000000000000000000000000000000000060e0998a1b811660228a01529690981b90951660268701527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000092821b8316602a870152811b8216603e86015292831b8116605285015293821b9093166066830152607a820192909252600090609a0160408051601f198184030181529082905260e08501516101008601516101208701516101408801516101608901516101808a01516101a08b01516101c08c0151979950612faf988a9890602001614457565b604051602081830303815290604052915050919050565b6000610c64825490565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16610ebf576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161224e565b60006130516001600160a01b0384168361332b565b9050805160001415801561307657508080602001905181019061307491906144e0565b155b15610d72576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b038416600482015260240161224e565b6000828152602081815260408083206001600160a01b038516845290915281205460ff16613178576000838152602081815260408083206001600160a01b0386168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556131303390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610c64565b506000610c64565b6000611703836001600160a01b038416613339565b6000828152602081815260408083206001600160a01b038516845290915281205460ff1615613178576000838152602081815260408083206001600160a01b038616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610c64565b6000611703836001600160a01b038416613380565b600082600001828154811061326257613262613ef7565b9060005260206000200154905092915050565b6060814710156132b3576040517fcd78605900000000000000000000000000000000000000000000000000000000815230600482015260240161224e565b600080856001600160a01b031684866040516132cf91906144fd565b60006040518083038185875af1925050503d806000811461330c576040519150601f19603f3d011682016040523d82523d6000602084013e613311565b606091505b5091509150613321868383613473565b9695505050505050565b606061170383836000613275565b600081815260018301602052604081205461317857508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610c64565b600081815260018301602052604081205480156134695760006133a46001836142f2565b85549091506000906133b8906001906142f2565b905080821461341d5760008660000182815481106133d8576133d8613ef7565b90600052602060002001549050808760000184815481106133fb576133fb613ef7565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061342e5761342e614519565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610c64565b6000915050610c64565b60608261348857613483826134e8565b611703565b815115801561349f57506001600160a01b0384163b155b156134e1576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161224e565b5080611703565b8051156134f85780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561353c57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461170357600080fd5b60006020828403121561357e57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600581106135eb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60208101610c6482846135b4565b6001600160a01b038116811461163a57600080fd5b803561361d816135fd565b919050565b6000806040838503121561363557600080fd5b8235613640816135fd565b91506020830135613650816135fd565b809150509250929050565b60008060006060848603121561367057600080fd5b83359250602084013591506040840135613689816135fd565b809150509250925092565b6000602082840312156136a657600080fd5b8135611703816135fd565b600080604083850312156136c457600080fd5b823591506020830135613650816135fd565b801515811461163a57600080fd5b803561361d816136d6565b60008060006040848603121561370457600080fd5b833567ffffffffffffffff8082111561371c57600080fd5b818601915086601f83011261373057600080fd5b81358181111561373f57600080fd5b8760208260051b850101111561375457600080fd5b60209283019550935050840135613689816136d6565b60005b8381101561378557818101518382015260200161376d565b50506000910152565b600081518084526137a681602086016020860161376a565b601f01601f19169290920160200192915050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b83811015613842578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180511515845287015187840187905261382f8785018261378e565b95880195935050908601906001016137e3565b509098975050505050505050565b60008083601f84011261386257600080fd5b50813567ffffffffffffffff81111561387a57600080fd5b60208301915083602082850101111561389257600080fd5b9250929050565b600080602083850312156138ac57600080fd5b823567ffffffffffffffff8111156138c357600080fd5b6138cf85828601613850565b90969095509350505050565b6000806000604084860312156138f057600080fd5b833567ffffffffffffffff81111561390757600080fd5b61391386828701613850565b9094509250506020840135613689816135fd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff8111828210171561397a5761397a613927565b60405290565b60405160a0810167ffffffffffffffff8111828210171561397a5761397a613927565b6040516101e0810167ffffffffffffffff8111828210171561397a5761397a613927565b604051610180810167ffffffffffffffff8111828210171561397a5761397a613927565b604051601f8201601f1916810167ffffffffffffffff81118282101715613a1457613a14613927565b604052919050565b63ffffffff8116811461163a57600080fd5b803561361d81613a1c565b60006101208284031215613a4c57600080fd5b613a54613956565b9050613a5f82613a2e565b8152613a6d60208301613612565b6020820152613a7e60408301613612565b6040820152613a8f60608301613612565b6060820152613aa060808301613612565b608082015260a082013560a082015260c082013560c0820152613ac560e083016136e4565b60e082015261010080830135818301525092915050565b60006101208284031215613aef57600080fd5b6117038383613a39565b60208152613b1060208201835163ffffffff169052565b60006020830151613b29604084018263ffffffff169052565b5060408301516001600160a01b03811660608401525060608301516001600160a01b03811660808401525060808301516001600160a01b03811660a08401525060a08301516001600160a01b03811660c08401525060c083015160e08381019190915283015161010080840191909152830151610120808401919091528301516101408084019190915283015161016080840191909152830151610180613bda818501836001600160a01b03169052565b8401516101a0848101919091528401516101c0808501919091528401516101e0808501529050613c0e61020084018261378e565b949350505050565b60808101613c2482876135b4565b63ffffffff8516602083015266ffffffffffffff841660408301526001600160a01b038316606083015295945050505050565b600080600060408486031215613c6c57600080fd5b833567ffffffffffffffff811115613c8357600080fd5b613c8f86828701613850565b909790965060209590950135949350505050565b60008060408385031215613cb657600080fd5b50508035926020909101359150565b815163ffffffff16815261018081016020830151613ceb602084018263ffffffff169052565b506040830151613d0660408401826001600160a01b03169052565b506060830151613d2160608401826001600160a01b03169052565b506080830151613d3c60808401826001600160a01b03169052565b5060a0830151613d5760a08401826001600160a01b03169052565b5060c083015160c083015260e083015160e083015261010080840151818401525061012080840151613d8c8285018215159052565b5050610140838101519083015261016092830151929091019190915290565b600067ffffffffffffffff821115613dc557613dc5613927565b50601f01601f191660200190565b600082601f830112613de457600080fd5b8135613df7613df282613dab565b6139eb565b818152846020838601011115613e0c57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806101408385031215613e3d57600080fd5b613e478484613a39565b915061012083013567ffffffffffffffff80821115613e6557600080fd5b9084019060a08287031215613e7957600080fd5b613e81613980565b8235613e8c816135fd565b815260208381013590820152604083013582811115613eaa57600080fd5b613eb688828601613dd3565b60408301525060608301356060820152608083013582811115613ed857600080fd5b613ee488828601613dd3565b6080830152508093505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613f5b57600080fd5b83018035915067ffffffffffffffff821115613f7657600080fd5b60200191503681900382131561389257600080fd5b8183823760009101908152919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610c6457610c64613f9b565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b602081526000613c0e602083018486613fdd565b805161361d81613a1c565b805161361d816135fd565b600082601f83011261404357600080fd5b8151614051613df282613dab565b81815284602083860101111561406657600080fd5b613c0e82602083016020870161376a565b60006020828403121561408957600080fd5b815167ffffffffffffffff808211156140a157600080fd5b908301906101e082860312156140b657600080fd5b6140be6139a3565b6140c78361401c565b81526140d56020840161401c565b60208201526140e660408401614027565b60408201526140f760608401614027565b606082015261410860808401614027565b608082015261411960a08401614027565b60a082015260c0838101519082015260e08084015190820152610100808401519082015261012080840151908201526101408084015190820152610160614161818501614027565b9082015261018083810151908201526101a080840151908201526101c0808401518381111561418f57600080fd5b61419b88828701614032565b918301919091525095945050505050565b600061018082840312156141bf57600080fd5b6141c76139c7565b6141d083613a2e565b81526141de60208401613a2e565b60208201526141ef60408401613612565b604082015261420060608401613612565b606082015261421160808401613612565b608082015261422260a08401613612565b60a082015260c083013560c082015260e083013560e08201526101008084013581830152506101206142558185016136e4565b908201526101408381013590820152610160928301359281019290925250919050565b808201828112600083128015821682158216171561429857614298613f9b565b505092915050565b8082028115828204841417610c6457610c64613f9b565b6000826142ed577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b81810381811115610c6457610c64613f9b565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361433657614336613f9b565b5060010190565b60e08152600061435060e083018a61378e565b63ffffffff989098166020830152506001600160a01b039586166040820152939094166060840152608083019190915260a082015290151560c090910152919050565b602081526000611703602083018461378e565b600080858511156143b657600080fd5b838611156143c357600080fd5b5050820193919092039150565b6001600160a01b0385168152836020820152606060408201526000613321606083018486613fdd565b80516020808301519190811015614438577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8160200360031b1b821691505b50919050565b60006020828403121561445057600080fd5b5051919050565b60008a51614469818460208f0161376a565b80830190508a81528960208201528860408201528760608201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008760601b1660808201528560948201528460b482015283516144cd8160d484016020880161376a565b0160d4019b9a5050505050505050505050565b6000602082840312156144f257600080fd5b8151611703816136d6565b6000825161450f81846020870161376a565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220ff0963a32deca1c739f82fff423c2a4ab1b1e197296c902578ae57133c488fd564736f6c63430008180033",
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

// CANCELERROLE is a free data retrieval call binding the contract method 0x02d2ff66.
//
// Solidity: function CANCELER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) CANCELERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "CANCELER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELERROLE is a free data retrieval call binding the contract method 0x02d2ff66.
//
// Solidity: function CANCELER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) CANCELERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.CANCELERROLE(&_FastBridgeV2.CallOpts)
}

// CANCELERROLE is a free data retrieval call binding the contract method 0x02d2ff66.
//
// Solidity: function CANCELER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) CANCELERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.CANCELERROLE(&_FastBridgeV2.CallOpts)
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

// DEFAULTCANCELDELAY is a free data retrieval call binding the contract method 0x930ac180.
//
// Solidity: function DEFAULT_CANCEL_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) DEFAULTCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "DEFAULT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTCANCELDELAY is a free data retrieval call binding the contract method 0x930ac180.
//
// Solidity: function DEFAULT_CANCEL_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) DEFAULTCANCELDELAY() (*big.Int, error) {
	return _FastBridgeV2.Contract.DEFAULTCANCELDELAY(&_FastBridgeV2.CallOpts)
}

// DEFAULTCANCELDELAY is a free data retrieval call binding the contract method 0x930ac180.
//
// Solidity: function DEFAULT_CANCEL_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) DEFAULTCANCELDELAY() (*big.Int, error) {
	return _FastBridgeV2.Contract.DEFAULTCANCELDELAY(&_FastBridgeV2.CallOpts)
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

// MINCANCELDELAY is a free data retrieval call binding the contract method 0x922b7487.
//
// Solidity: function MIN_CANCEL_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) MINCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "MIN_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCANCELDELAY is a free data retrieval call binding the contract method 0x922b7487.
//
// Solidity: function MIN_CANCEL_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) MINCANCELDELAY() (*big.Int, error) {
	return _FastBridgeV2.Contract.MINCANCELDELAY(&_FastBridgeV2.CallOpts)
}

// MINCANCELDELAY is a free data retrieval call binding the contract method 0x922b7487.
//
// Solidity: function MIN_CANCEL_DELAY() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) MINCANCELDELAY() (*big.Int, error) {
	return _FastBridgeV2.Contract.MINCANCELDELAY(&_FastBridgeV2.CallOpts)
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

// PROVERROLE is a free data retrieval call binding the contract method 0xdc9a4ef6.
//
// Solidity: function PROVER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) PROVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "PROVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVERROLE is a free data retrieval call binding the contract method 0xdc9a4ef6.
//
// Solidity: function PROVER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) PROVERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.PROVERROLE(&_FastBridgeV2.CallOpts)
}

// PROVERROLE is a free data retrieval call binding the contract method 0xdc9a4ef6.
//
// Solidity: function PROVER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) PROVERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.PROVERROLE(&_FastBridgeV2.CallOpts)
}

// QUOTERROLE is a free data retrieval call binding the contract method 0x7ebe815c.
//
// Solidity: function QUOTER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Caller) QUOTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "QUOTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// QUOTERROLE is a free data retrieval call binding the contract method 0x7ebe815c.
//
// Solidity: function QUOTER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2Session) QUOTERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.QUOTERROLE(&_FastBridgeV2.CallOpts)
}

// QUOTERROLE is a free data retrieval call binding the contract method 0x7ebe815c.
//
// Solidity: function QUOTER_ROLE() view returns(bytes32)
func (_FastBridgeV2 *FastBridgeV2CallerSession) QUOTERROLE() ([32]byte, error) {
	return _FastBridgeV2.Contract.QUOTERROLE(&_FastBridgeV2.CallOpts)
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

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Caller) CancelDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridgeV2.contract.Call(opts, &out, "cancelDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2Session) CancelDelay() (*big.Int, error) {
	return _FastBridgeV2.Contract.CancelDelay(&_FastBridgeV2.CallOpts)
}

// CancelDelay is a free data retrieval call binding the contract method 0x638a0f09.
//
// Solidity: function cancelDelay() view returns(uint256)
func (_FastBridgeV2 *FastBridgeV2CallerSession) CancelDelay() (*big.Int, error) {
	return _FastBridgeV2.Contract.CancelDelay(&_FastBridgeV2.CallOpts)
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

// Cancel is a paid mutator transaction binding the contract method 0x3c5beeb4.
//
// Solidity: function cancel(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) Cancel(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "cancel", request)
}

// Cancel is a paid mutator transaction binding the contract method 0x3c5beeb4.
//
// Solidity: function cancel(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2Session) Cancel(request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Cancel(&_FastBridgeV2.TransactOpts, request)
}

// Cancel is a paid mutator transaction binding the contract method 0x3c5beeb4.
//
// Solidity: function cancel(bytes request) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) Cancel(request []byte) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.Cancel(&_FastBridgeV2.TransactOpts, request)
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

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_FastBridgeV2 *FastBridgeV2Transactor) SetCancelDelay(opts *bind.TransactOpts, newCancelDelay *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.contract.Transact(opts, "setCancelDelay", newCancelDelay)
}

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_FastBridgeV2 *FastBridgeV2Session) SetCancelDelay(newCancelDelay *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.SetCancelDelay(&_FastBridgeV2.TransactOpts, newCancelDelay)
}

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_FastBridgeV2 *FastBridgeV2TransactorSession) SetCancelDelay(newCancelDelay *big.Int) (*types.Transaction, error) {
	return _FastBridgeV2.Contract.SetCancelDelay(&_FastBridgeV2.TransactOpts, newCancelDelay)
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

// FastBridgeV2CancelDelayUpdatedIterator is returned from FilterCancelDelayUpdated and is used to iterate over the raw logs and unpacked data for CancelDelayUpdated events raised by the FastBridgeV2 contract.
type FastBridgeV2CancelDelayUpdatedIterator struct {
	Event *FastBridgeV2CancelDelayUpdated // Event containing the contract specifics and raw log

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
func (it *FastBridgeV2CancelDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeV2CancelDelayUpdated)
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
		it.Event = new(FastBridgeV2CancelDelayUpdated)
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
func (it *FastBridgeV2CancelDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeV2CancelDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeV2CancelDelayUpdated represents a CancelDelayUpdated event raised by the FastBridgeV2 contract.
type FastBridgeV2CancelDelayUpdated struct {
	OldCancelDelay *big.Int
	NewCancelDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCancelDelayUpdated is a free log retrieval operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_FastBridgeV2 *FastBridgeV2Filterer) FilterCancelDelayUpdated(opts *bind.FilterOpts) (*FastBridgeV2CancelDelayUpdatedIterator, error) {

	logs, sub, err := _FastBridgeV2.contract.FilterLogs(opts, "CancelDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &FastBridgeV2CancelDelayUpdatedIterator{contract: _FastBridgeV2.contract, event: "CancelDelayUpdated", logs: logs, sub: sub}, nil
}

// WatchCancelDelayUpdated is a free log subscription operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_FastBridgeV2 *FastBridgeV2Filterer) WatchCancelDelayUpdated(opts *bind.WatchOpts, sink chan<- *FastBridgeV2CancelDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _FastBridgeV2.contract.WatchLogs(opts, "CancelDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeV2CancelDelayUpdated)
				if err := _FastBridgeV2.contract.UnpackLog(event, "CancelDelayUpdated", log); err != nil {
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

// ParseCancelDelayUpdated is a log parse operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_FastBridgeV2 *FastBridgeV2Filterer) ParseCancelDelayUpdated(log types.Log) (*FastBridgeV2CancelDelayUpdated, error) {
	event := new(FastBridgeV2CancelDelayUpdated)
	if err := _FastBridgeV2.contract.UnpackLog(event, "CancelDelayUpdated", log); err != nil {
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

// IAdminV2MetaData contains all meta data concerning the IAdminV2 contract.
var IAdminV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCancelDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCancelDelay\",\"type\":\"uint256\"}],\"name\":\"CancelDelayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesSwept\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCancelDelay\",\"type\":\"uint256\"}],\"name\":\"setCancelDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1ea327c5": "setCancelDelay(uint256)",
		"b13aa2d6": "setProtocolFeeRate(uint256)",
		"06f333f2": "sweepProtocolFees(address,address)",
	},
}

// IAdminV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use IAdminV2MetaData.ABI instead.
var IAdminV2ABI = IAdminV2MetaData.ABI

// Deprecated: Use IAdminV2MetaData.Sigs instead.
// IAdminV2FuncSigs maps the 4-byte function signature to its string representation.
var IAdminV2FuncSigs = IAdminV2MetaData.Sigs

// IAdminV2 is an auto generated Go binding around an Ethereum contract.
type IAdminV2 struct {
	IAdminV2Caller     // Read-only binding to the contract
	IAdminV2Transactor // Write-only binding to the contract
	IAdminV2Filterer   // Log filterer for contract events
}

// IAdminV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IAdminV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IAdminV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAdminV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAdminV2Session struct {
	Contract     *IAdminV2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAdminV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAdminV2CallerSession struct {
	Contract *IAdminV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IAdminV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAdminV2TransactorSession struct {
	Contract     *IAdminV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IAdminV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IAdminV2Raw struct {
	Contract *IAdminV2 // Generic contract binding to access the raw methods on
}

// IAdminV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAdminV2CallerRaw struct {
	Contract *IAdminV2Caller // Generic read-only contract binding to access the raw methods on
}

// IAdminV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAdminV2TransactorRaw struct {
	Contract *IAdminV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIAdminV2 creates a new instance of IAdminV2, bound to a specific deployed contract.
func NewIAdminV2(address common.Address, backend bind.ContractBackend) (*IAdminV2, error) {
	contract, err := bindIAdminV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAdminV2{IAdminV2Caller: IAdminV2Caller{contract: contract}, IAdminV2Transactor: IAdminV2Transactor{contract: contract}, IAdminV2Filterer: IAdminV2Filterer{contract: contract}}, nil
}

// NewIAdminV2Caller creates a new read-only instance of IAdminV2, bound to a specific deployed contract.
func NewIAdminV2Caller(address common.Address, caller bind.ContractCaller) (*IAdminV2Caller, error) {
	contract, err := bindIAdminV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminV2Caller{contract: contract}, nil
}

// NewIAdminV2Transactor creates a new write-only instance of IAdminV2, bound to a specific deployed contract.
func NewIAdminV2Transactor(address common.Address, transactor bind.ContractTransactor) (*IAdminV2Transactor, error) {
	contract, err := bindIAdminV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminV2Transactor{contract: contract}, nil
}

// NewIAdminV2Filterer creates a new log filterer instance of IAdminV2, bound to a specific deployed contract.
func NewIAdminV2Filterer(address common.Address, filterer bind.ContractFilterer) (*IAdminV2Filterer, error) {
	contract, err := bindIAdminV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAdminV2Filterer{contract: contract}, nil
}

// bindIAdminV2 binds a generic wrapper to an already deployed contract.
func bindIAdminV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAdminV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAdminV2 *IAdminV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAdminV2.Contract.IAdminV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAdminV2 *IAdminV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAdminV2.Contract.IAdminV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAdminV2 *IAdminV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAdminV2.Contract.IAdminV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAdminV2 *IAdminV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAdminV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAdminV2 *IAdminV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAdminV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAdminV2 *IAdminV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAdminV2.Contract.contract.Transact(opts, method, params...)
}

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_IAdminV2 *IAdminV2Transactor) SetCancelDelay(opts *bind.TransactOpts, newCancelDelay *big.Int) (*types.Transaction, error) {
	return _IAdminV2.contract.Transact(opts, "setCancelDelay", newCancelDelay)
}

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_IAdminV2 *IAdminV2Session) SetCancelDelay(newCancelDelay *big.Int) (*types.Transaction, error) {
	return _IAdminV2.Contract.SetCancelDelay(&_IAdminV2.TransactOpts, newCancelDelay)
}

// SetCancelDelay is a paid mutator transaction binding the contract method 0x1ea327c5.
//
// Solidity: function setCancelDelay(uint256 newCancelDelay) returns()
func (_IAdminV2 *IAdminV2TransactorSession) SetCancelDelay(newCancelDelay *big.Int) (*types.Transaction, error) {
	return _IAdminV2.Contract.SetCancelDelay(&_IAdminV2.TransactOpts, newCancelDelay)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_IAdminV2 *IAdminV2Transactor) SetProtocolFeeRate(opts *bind.TransactOpts, newFeeRate *big.Int) (*types.Transaction, error) {
	return _IAdminV2.contract.Transact(opts, "setProtocolFeeRate", newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_IAdminV2 *IAdminV2Session) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _IAdminV2.Contract.SetProtocolFeeRate(&_IAdminV2.TransactOpts, newFeeRate)
}

// SetProtocolFeeRate is a paid mutator transaction binding the contract method 0xb13aa2d6.
//
// Solidity: function setProtocolFeeRate(uint256 newFeeRate) returns()
func (_IAdminV2 *IAdminV2TransactorSession) SetProtocolFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _IAdminV2.Contract.SetProtocolFeeRate(&_IAdminV2.TransactOpts, newFeeRate)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_IAdminV2 *IAdminV2Transactor) SweepProtocolFees(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IAdminV2.contract.Transact(opts, "sweepProtocolFees", token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_IAdminV2 *IAdminV2Session) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IAdminV2.Contract.SweepProtocolFees(&_IAdminV2.TransactOpts, token, recipient)
}

// SweepProtocolFees is a paid mutator transaction binding the contract method 0x06f333f2.
//
// Solidity: function sweepProtocolFees(address token, address recipient) returns()
func (_IAdminV2 *IAdminV2TransactorSession) SweepProtocolFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _IAdminV2.Contract.SweepProtocolFees(&_IAdminV2.TransactOpts, token, recipient)
}

// IAdminV2CancelDelayUpdatedIterator is returned from FilterCancelDelayUpdated and is used to iterate over the raw logs and unpacked data for CancelDelayUpdated events raised by the IAdminV2 contract.
type IAdminV2CancelDelayUpdatedIterator struct {
	Event *IAdminV2CancelDelayUpdated // Event containing the contract specifics and raw log

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
func (it *IAdminV2CancelDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminV2CancelDelayUpdated)
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
		it.Event = new(IAdminV2CancelDelayUpdated)
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
func (it *IAdminV2CancelDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminV2CancelDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminV2CancelDelayUpdated represents a CancelDelayUpdated event raised by the IAdminV2 contract.
type IAdminV2CancelDelayUpdated struct {
	OldCancelDelay *big.Int
	NewCancelDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCancelDelayUpdated is a free log retrieval operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_IAdminV2 *IAdminV2Filterer) FilterCancelDelayUpdated(opts *bind.FilterOpts) (*IAdminV2CancelDelayUpdatedIterator, error) {

	logs, sub, err := _IAdminV2.contract.FilterLogs(opts, "CancelDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &IAdminV2CancelDelayUpdatedIterator{contract: _IAdminV2.contract, event: "CancelDelayUpdated", logs: logs, sub: sub}, nil
}

// WatchCancelDelayUpdated is a free log subscription operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_IAdminV2 *IAdminV2Filterer) WatchCancelDelayUpdated(opts *bind.WatchOpts, sink chan<- *IAdminV2CancelDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _IAdminV2.contract.WatchLogs(opts, "CancelDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminV2CancelDelayUpdated)
				if err := _IAdminV2.contract.UnpackLog(event, "CancelDelayUpdated", log); err != nil {
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

// ParseCancelDelayUpdated is a log parse operation binding the contract event 0x909f9078b2f6eac1d10f2aae793656c5a67511eb627ebd1d2cb1eb9c0ab94c95.
//
// Solidity: event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay)
func (_IAdminV2 *IAdminV2Filterer) ParseCancelDelayUpdated(log types.Log) (*IAdminV2CancelDelayUpdated, error) {
	event := new(IAdminV2CancelDelayUpdated)
	if err := _IAdminV2.contract.UnpackLog(event, "CancelDelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminV2FeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the IAdminV2 contract.
type IAdminV2FeeRateUpdatedIterator struct {
	Event *IAdminV2FeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *IAdminV2FeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminV2FeeRateUpdated)
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
		it.Event = new(IAdminV2FeeRateUpdated)
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
func (it *IAdminV2FeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminV2FeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminV2FeeRateUpdated represents a FeeRateUpdated event raised by the IAdminV2 contract.
type IAdminV2FeeRateUpdated struct {
	OldFeeRate *big.Int
	NewFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_IAdminV2 *IAdminV2Filterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*IAdminV2FeeRateUpdatedIterator, error) {

	logs, sub, err := _IAdminV2.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &IAdminV2FeeRateUpdatedIterator{contract: _IAdminV2.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate)
func (_IAdminV2 *IAdminV2Filterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *IAdminV2FeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _IAdminV2.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminV2FeeRateUpdated)
				if err := _IAdminV2.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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
func (_IAdminV2 *IAdminV2Filterer) ParseFeeRateUpdated(log types.Log) (*IAdminV2FeeRateUpdated, error) {
	event := new(IAdminV2FeeRateUpdated)
	if err := _IAdminV2.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminV2FeesSweptIterator is returned from FilterFeesSwept and is used to iterate over the raw logs and unpacked data for FeesSwept events raised by the IAdminV2 contract.
type IAdminV2FeesSweptIterator struct {
	Event *IAdminV2FeesSwept // Event containing the contract specifics and raw log

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
func (it *IAdminV2FeesSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAdminV2FeesSwept)
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
		it.Event = new(IAdminV2FeesSwept)
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
func (it *IAdminV2FeesSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAdminV2FeesSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAdminV2FeesSwept represents a FeesSwept event raised by the IAdminV2 contract.
type IAdminV2FeesSwept struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesSwept is a free log retrieval operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_IAdminV2 *IAdminV2Filterer) FilterFeesSwept(opts *bind.FilterOpts) (*IAdminV2FeesSweptIterator, error) {

	logs, sub, err := _IAdminV2.contract.FilterLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return &IAdminV2FeesSweptIterator{contract: _IAdminV2.contract, event: "FeesSwept", logs: logs, sub: sub}, nil
}

// WatchFeesSwept is a free log subscription operation binding the contract event 0x244e51bc38c1452fa8aaf487bcb4bca36c2baa3a5fbdb776b1eabd8dc6d277cd.
//
// Solidity: event FeesSwept(address token, address recipient, uint256 amount)
func (_IAdminV2 *IAdminV2Filterer) WatchFeesSwept(opts *bind.WatchOpts, sink chan<- *IAdminV2FeesSwept) (event.Subscription, error) {

	logs, sub, err := _IAdminV2.contract.WatchLogs(opts, "FeesSwept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAdminV2FeesSwept)
				if err := _IAdminV2.contract.UnpackLog(event, "FeesSwept", log); err != nil {
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
func (_IAdminV2 *IAdminV2Filterer) ParseFeesSwept(log types.Log) (*IAdminV2FeesSwept, error) {
	event := new(IAdminV2FeesSwept)
	if err := _IAdminV2.contract.UnpackLog(event, "FeesSwept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAdminV2ErrorsMetaData contains all meta data concerning the IAdminV2Errors contract.
var IAdminV2ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CancelDelayBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRateAboveMax\",\"type\":\"error\"}]",
}

// IAdminV2ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IAdminV2ErrorsMetaData.ABI instead.
var IAdminV2ErrorsABI = IAdminV2ErrorsMetaData.ABI

// IAdminV2Errors is an auto generated Go binding around an Ethereum contract.
type IAdminV2Errors struct {
	IAdminV2ErrorsCaller     // Read-only binding to the contract
	IAdminV2ErrorsTransactor // Write-only binding to the contract
	IAdminV2ErrorsFilterer   // Log filterer for contract events
}

// IAdminV2ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAdminV2ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminV2ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAdminV2ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminV2ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAdminV2ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminV2ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAdminV2ErrorsSession struct {
	Contract     *IAdminV2Errors   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAdminV2ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAdminV2ErrorsCallerSession struct {
	Contract *IAdminV2ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAdminV2ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAdminV2ErrorsTransactorSession struct {
	Contract     *IAdminV2ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAdminV2ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAdminV2ErrorsRaw struct {
	Contract *IAdminV2Errors // Generic contract binding to access the raw methods on
}

// IAdminV2ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAdminV2ErrorsCallerRaw struct {
	Contract *IAdminV2ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IAdminV2ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAdminV2ErrorsTransactorRaw struct {
	Contract *IAdminV2ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAdminV2Errors creates a new instance of IAdminV2Errors, bound to a specific deployed contract.
func NewIAdminV2Errors(address common.Address, backend bind.ContractBackend) (*IAdminV2Errors, error) {
	contract, err := bindIAdminV2Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAdminV2Errors{IAdminV2ErrorsCaller: IAdminV2ErrorsCaller{contract: contract}, IAdminV2ErrorsTransactor: IAdminV2ErrorsTransactor{contract: contract}, IAdminV2ErrorsFilterer: IAdminV2ErrorsFilterer{contract: contract}}, nil
}

// NewIAdminV2ErrorsCaller creates a new read-only instance of IAdminV2Errors, bound to a specific deployed contract.
func NewIAdminV2ErrorsCaller(address common.Address, caller bind.ContractCaller) (*IAdminV2ErrorsCaller, error) {
	contract, err := bindIAdminV2Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminV2ErrorsCaller{contract: contract}, nil
}

// NewIAdminV2ErrorsTransactor creates a new write-only instance of IAdminV2Errors, bound to a specific deployed contract.
func NewIAdminV2ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IAdminV2ErrorsTransactor, error) {
	contract, err := bindIAdminV2Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminV2ErrorsTransactor{contract: contract}, nil
}

// NewIAdminV2ErrorsFilterer creates a new log filterer instance of IAdminV2Errors, bound to a specific deployed contract.
func NewIAdminV2ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IAdminV2ErrorsFilterer, error) {
	contract, err := bindIAdminV2Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAdminV2ErrorsFilterer{contract: contract}, nil
}

// bindIAdminV2Errors binds a generic wrapper to an already deployed contract.
func bindIAdminV2Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAdminV2ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAdminV2Errors *IAdminV2ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAdminV2Errors.Contract.IAdminV2ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAdminV2Errors *IAdminV2ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAdminV2Errors.Contract.IAdminV2ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAdminV2Errors *IAdminV2ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAdminV2Errors.Contract.IAdminV2ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAdminV2Errors *IAdminV2ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAdminV2Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAdminV2Errors *IAdminV2ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAdminV2Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAdminV2Errors *IAdminV2ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAdminV2Errors.Contract.contract.Transact(opts, method, params...)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"}],\"name\":\"BridgeQuoteDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"quoteRelayer\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"quoteExclusivitySeconds\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeParamsV2\",\"name\":\"paramsV2\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeProofs\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"timestamp\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelays\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeStatuses\",\"outputs\":[{\"internalType\":\"enumIFastBridgeV2.BridgeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransactionV2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"exclusivityRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exclusivityEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeTransactionV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"45851694": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256))",
		"bfc7c607": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256),(address,int256,bytes,uint256,bytes))",
		"91ad5039": "bridgeProofs(bytes32)",
		"8379a24f": "bridgeRelays(bytes32)",
		"051287bc": "bridgeStatuses(bytes32)",
		"aa9641ab": "canClaim(bytes32,address)",
		"3c5beeb4": "cancel(bytes)",
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

// Cancel is a paid mutator transaction binding the contract method 0x3c5beeb4.
//
// Solidity: function cancel(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Cancel(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "cancel", request)
}

// Cancel is a paid mutator transaction binding the contract method 0x3c5beeb4.
//
// Solidity: function cancel(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Cancel(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Cancel(&_IFastBridgeV2.TransactOpts, request)
}

// Cancel is a paid mutator transaction binding the contract method 0x3c5beeb4.
//
// Solidity: function cancel(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Cancel(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Cancel(&_IFastBridgeV2.TransactOpts, request)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fc15551e7cd0cae1476ceeb26ea6d8e21a222c930b99686cddd49777e215b49764736f6c63430008180033",
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
