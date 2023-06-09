// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cctp

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220dd1790f715b9503832d7cbd531a84f9ba1d192821e322c245940712f589b2ec564736f6c63430008110033",
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

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
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

// IMessageTransmitterMetaData contains all meta data concerning the IMessageTransmitter contract.
var IMessageTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvailableNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"8371744e": "nextAvailableNonce()",
		"57ecfd28": "receiveMessage(bytes,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
}

// IMessageTransmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageTransmitterMetaData.ABI instead.
var IMessageTransmitterABI = IMessageTransmitterMetaData.ABI

// Deprecated: Use IMessageTransmitterMetaData.Sigs instead.
// IMessageTransmitterFuncSigs maps the 4-byte function signature to its string representation.
var IMessageTransmitterFuncSigs = IMessageTransmitterMetaData.Sigs

// IMessageTransmitter is an auto generated Go binding around an Ethereum contract.
type IMessageTransmitter struct {
	IMessageTransmitterCaller     // Read-only binding to the contract
	IMessageTransmitterTransactor // Write-only binding to the contract
	IMessageTransmitterFilterer   // Log filterer for contract events
}

// IMessageTransmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageTransmitterSession struct {
	Contract     *IMessageTransmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMessageTransmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageTransmitterCallerSession struct {
	Contract *IMessageTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMessageTransmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageTransmitterTransactorSession struct {
	Contract     *IMessageTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMessageTransmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageTransmitterRaw struct {
	Contract *IMessageTransmitter // Generic contract binding to access the raw methods on
}

// IMessageTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageTransmitterCallerRaw struct {
	Contract *IMessageTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactorRaw struct {
	Contract *IMessageTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageTransmitter creates a new instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitter(address common.Address, backend bind.ContractBackend) (*IMessageTransmitter, error) {
	contract, err := bindIMessageTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitter{IMessageTransmitterCaller: IMessageTransmitterCaller{contract: contract}, IMessageTransmitterTransactor: IMessageTransmitterTransactor{contract: contract}, IMessageTransmitterFilterer: IMessageTransmitterFilterer{contract: contract}}, nil
}

// NewIMessageTransmitterCaller creates a new read-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterCaller(address common.Address, caller bind.ContractCaller) (*IMessageTransmitterCaller, error) {
	contract, err := bindIMessageTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterCaller{contract: contract}, nil
}

// NewIMessageTransmitterTransactor creates a new write-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageTransmitterTransactor, error) {
	contract, err := bindIMessageTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterTransactor{contract: contract}, nil
}

// NewIMessageTransmitterFilterer creates a new log filterer instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageTransmitterFilterer, error) {
	contract, err := bindIMessageTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterFilterer{contract: contract}, nil
}

// bindIMessageTransmitter binds a generic wrapper to an already deployed contract.
func bindIMessageTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageTransmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.IMessageTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IMessageTransmitter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterSession) LocalDomain() (uint32, error) {
	return _IMessageTransmitter.Contract.LocalDomain(&_IMessageTransmitter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterCallerSession) LocalDomain() (uint32, error) {
	return _IMessageTransmitter.Contract.LocalDomain(&_IMessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterCaller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IMessageTransmitter.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) NextAvailableNonce() (uint64, error) {
	return _IMessageTransmitter.Contract.NextAvailableNonce(&_IMessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterCallerSession) NextAvailableNonce() (uint64, error) {
	return _IMessageTransmitter.Contract.NextAvailableNonce(&_IMessageTransmitter.CallOpts)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// ISynapseCCTPMetaData contains all meta data concerning the ISynapseCCTP contract.
var ISynapseCCTPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"}],\"name\":\"receiveCircleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"swapParams\",\"type\":\"bytes\"}],\"name\":\"sendCircleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4a5ae51d": "receiveCircleToken(bytes,bytes,uint32,bytes)",
		"304ddb4c": "sendCircleToken(address,uint256,address,uint256,uint32,bytes)",
	},
}

// ISynapseCCTPABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseCCTPMetaData.ABI instead.
var ISynapseCCTPABI = ISynapseCCTPMetaData.ABI

// Deprecated: Use ISynapseCCTPMetaData.Sigs instead.
// ISynapseCCTPFuncSigs maps the 4-byte function signature to its string representation.
var ISynapseCCTPFuncSigs = ISynapseCCTPMetaData.Sigs

// ISynapseCCTP is an auto generated Go binding around an Ethereum contract.
type ISynapseCCTP struct {
	ISynapseCCTPCaller     // Read-only binding to the contract
	ISynapseCCTPTransactor // Write-only binding to the contract
	ISynapseCCTPFilterer   // Log filterer for contract events
}

// ISynapseCCTPCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseCCTPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseCCTPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseCCTPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseCCTPSession struct {
	Contract     *ISynapseCCTP     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynapseCCTPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseCCTPCallerSession struct {
	Contract *ISynapseCCTPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ISynapseCCTPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseCCTPTransactorSession struct {
	Contract     *ISynapseCCTPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ISynapseCCTPRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseCCTPRaw struct {
	Contract *ISynapseCCTP // Generic contract binding to access the raw methods on
}

// ISynapseCCTPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseCCTPCallerRaw struct {
	Contract *ISynapseCCTPCaller // Generic read-only contract binding to access the raw methods on
}

// ISynapseCCTPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseCCTPTransactorRaw struct {
	Contract *ISynapseCCTPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseCCTP creates a new instance of ISynapseCCTP, bound to a specific deployed contract.
func NewISynapseCCTP(address common.Address, backend bind.ContractBackend) (*ISynapseCCTP, error) {
	contract, err := bindISynapseCCTP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTP{ISynapseCCTPCaller: ISynapseCCTPCaller{contract: contract}, ISynapseCCTPTransactor: ISynapseCCTPTransactor{contract: contract}, ISynapseCCTPFilterer: ISynapseCCTPFilterer{contract: contract}}, nil
}

// NewISynapseCCTPCaller creates a new read-only instance of ISynapseCCTP, bound to a specific deployed contract.
func NewISynapseCCTPCaller(address common.Address, caller bind.ContractCaller) (*ISynapseCCTPCaller, error) {
	contract, err := bindISynapseCCTP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPCaller{contract: contract}, nil
}

// NewISynapseCCTPTransactor creates a new write-only instance of ISynapseCCTP, bound to a specific deployed contract.
func NewISynapseCCTPTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseCCTPTransactor, error) {
	contract, err := bindISynapseCCTP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPTransactor{contract: contract}, nil
}

// NewISynapseCCTPFilterer creates a new log filterer instance of ISynapseCCTP, bound to a specific deployed contract.
func NewISynapseCCTPFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseCCTPFilterer, error) {
	contract, err := bindISynapseCCTP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPFilterer{contract: contract}, nil
}

// bindISynapseCCTP binds a generic wrapper to an already deployed contract.
func bindISynapseCCTP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISynapseCCTPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseCCTP *ISynapseCCTPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseCCTP.Contract.ISynapseCCTPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseCCTP *ISynapseCCTPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.ISynapseCCTPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseCCTP *ISynapseCCTPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.ISynapseCCTPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseCCTP *ISynapseCCTPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseCCTP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseCCTP *ISynapseCCTPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseCCTP *ISynapseCCTPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.contract.Transact(opts, method, params...)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) returns()
func (_ISynapseCCTP *ISynapseCCTPTransactor) ReceiveCircleToken(opts *bind.TransactOpts, message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.contract.Transact(opts, "receiveCircleToken", message, signature, requestVersion, formattedRequest)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) returns()
func (_ISynapseCCTP *ISynapseCCTPSession) ReceiveCircleToken(message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.ReceiveCircleToken(&_ISynapseCCTP.TransactOpts, message, signature, requestVersion, formattedRequest)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) returns()
func (_ISynapseCCTP *ISynapseCCTPTransactorSession) ReceiveCircleToken(message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.ReceiveCircleToken(&_ISynapseCCTP.TransactOpts, message, signature, requestVersion, formattedRequest)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_ISynapseCCTP *ISynapseCCTPTransactor) SendCircleToken(opts *bind.TransactOpts, recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.contract.Transact(opts, "sendCircleToken", recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_ISynapseCCTP *ISynapseCCTPSession) SendCircleToken(recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.SendCircleToken(&_ISynapseCCTP.TransactOpts, recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_ISynapseCCTP *ISynapseCCTPTransactorSession) SendCircleToken(recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.SendCircleToken(&_ISynapseCCTP.TransactOpts, recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// ITokenMessengerMetaData contains all meta data concerning the ITokenMessenger contract.
var ITokenMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"depositForBurnWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMessageTransmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f856ddb6": "depositForBurnWithCaller(uint256,uint32,bytes32,address,bytes32)",
		"96abeb70": "handleReceiveMessage(uint32,bytes32,bytes)",
		"2c121921": "localMessageTransmitter()",
		"cb75c11c": "localMinter()",
	},
}

// ITokenMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMessengerMetaData.ABI instead.
var ITokenMessengerABI = ITokenMessengerMetaData.ABI

// Deprecated: Use ITokenMessengerMetaData.Sigs instead.
// ITokenMessengerFuncSigs maps the 4-byte function signature to its string representation.
var ITokenMessengerFuncSigs = ITokenMessengerMetaData.Sigs

// ITokenMessenger is an auto generated Go binding around an Ethereum contract.
type ITokenMessenger struct {
	ITokenMessengerCaller     // Read-only binding to the contract
	ITokenMessengerTransactor // Write-only binding to the contract
	ITokenMessengerFilterer   // Log filterer for contract events
}

// ITokenMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenMessengerSession struct {
	Contract     *ITokenMessenger  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenMessengerCallerSession struct {
	Contract *ITokenMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ITokenMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenMessengerTransactorSession struct {
	Contract     *ITokenMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ITokenMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenMessengerRaw struct {
	Contract *ITokenMessenger // Generic contract binding to access the raw methods on
}

// ITokenMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenMessengerCallerRaw struct {
	Contract *ITokenMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenMessengerTransactorRaw struct {
	Contract *ITokenMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenMessenger creates a new instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessenger(address common.Address, backend bind.ContractBackend) (*ITokenMessenger, error) {
	contract, err := bindITokenMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenMessenger{ITokenMessengerCaller: ITokenMessengerCaller{contract: contract}, ITokenMessengerTransactor: ITokenMessengerTransactor{contract: contract}, ITokenMessengerFilterer: ITokenMessengerFilterer{contract: contract}}, nil
}

// NewITokenMessengerCaller creates a new read-only instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerCaller(address common.Address, caller bind.ContractCaller) (*ITokenMessengerCaller, error) {
	contract, err := bindITokenMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerCaller{contract: contract}, nil
}

// NewITokenMessengerTransactor creates a new write-only instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenMessengerTransactor, error) {
	contract, err := bindITokenMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerTransactor{contract: contract}, nil
}

// NewITokenMessengerFilterer creates a new log filterer instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenMessengerFilterer, error) {
	contract, err := bindITokenMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerFilterer{contract: contract}, nil
}

// bindITokenMessenger binds a generic wrapper to an already deployed contract.
func bindITokenMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMessenger *ITokenMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMessenger.Contract.ITokenMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMessenger *ITokenMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ITokenMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMessenger *ITokenMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ITokenMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMessenger *ITokenMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMessenger *ITokenMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMessenger *ITokenMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.contract.Transact(opts, method, params...)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCaller) LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITokenMessenger.contract.Call(opts, &out, "localMessageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerSession) LocalMessageTransmitter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMessageTransmitter(&_ITokenMessenger.CallOpts)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCallerSession) LocalMessageTransmitter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMessageTransmitter(&_ITokenMessenger.CallOpts)
}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCaller) LocalMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITokenMessenger.contract.Call(opts, &out, "localMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerSession) LocalMinter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMinter(&_ITokenMessenger.CallOpts)
}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCallerSession) LocalMinter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMinter(&_ITokenMessenger.CallOpts)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerTransactor) DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "depositForBurnWithCaller", amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurnWithCaller(&_ITokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerTransactorSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurnWithCaller(&_ITokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerTransactor) HandleReceiveMessage(opts *bind.TransactOpts, remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "handleReceiveMessage", remoteDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerSession) HandleReceiveMessage(remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.HandleReceiveMessage(&_ITokenMessenger.TransactOpts, remoteDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerTransactorSession) HandleReceiveMessage(remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.HandleReceiveMessage(&_ITokenMessenger.TransactOpts, remoteDomain, sender, messageBody)
}

// ITokenMinterMetaData contains all meta data concerning the ITokenMinter contract.
var ITokenMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"remoteToken\",\"type\":\"bytes32\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"burnToken\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9dc29fac": "burn(address,uint256)",
		"78a0565e": "getLocalToken(uint32,bytes32)",
		"d54de06f": "mint(uint32,bytes32,address,uint256)",
	},
}

// ITokenMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMinterMetaData.ABI instead.
var ITokenMinterABI = ITokenMinterMetaData.ABI

// Deprecated: Use ITokenMinterMetaData.Sigs instead.
// ITokenMinterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenMinterFuncSigs = ITokenMinterMetaData.Sigs

// ITokenMinter is an auto generated Go binding around an Ethereum contract.
type ITokenMinter struct {
	ITokenMinterCaller     // Read-only binding to the contract
	ITokenMinterTransactor // Write-only binding to the contract
	ITokenMinterFilterer   // Log filterer for contract events
}

// ITokenMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenMinterSession struct {
	Contract     *ITokenMinter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenMinterCallerSession struct {
	Contract *ITokenMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenMinterTransactorSession struct {
	Contract     *ITokenMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenMinterRaw struct {
	Contract *ITokenMinter // Generic contract binding to access the raw methods on
}

// ITokenMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenMinterCallerRaw struct {
	Contract *ITokenMinterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenMinterTransactorRaw struct {
	Contract *ITokenMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenMinter creates a new instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinter(address common.Address, backend bind.ContractBackend) (*ITokenMinter, error) {
	contract, err := bindITokenMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenMinter{ITokenMinterCaller: ITokenMinterCaller{contract: contract}, ITokenMinterTransactor: ITokenMinterTransactor{contract: contract}, ITokenMinterFilterer: ITokenMinterFilterer{contract: contract}}, nil
}

// NewITokenMinterCaller creates a new read-only instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterCaller(address common.Address, caller bind.ContractCaller) (*ITokenMinterCaller, error) {
	contract, err := bindITokenMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterCaller{contract: contract}, nil
}

// NewITokenMinterTransactor creates a new write-only instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenMinterTransactor, error) {
	contract, err := bindITokenMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterTransactor{contract: contract}, nil
}

// NewITokenMinterFilterer creates a new log filterer instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenMinterFilterer, error) {
	contract, err := bindITokenMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterFilterer{contract: contract}, nil
}

// bindITokenMinter binds a generic wrapper to an already deployed contract.
func bindITokenMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMinter *ITokenMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMinter.Contract.ITokenMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMinter *ITokenMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMinter.Contract.ITokenMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMinter *ITokenMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMinter.Contract.ITokenMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMinter *ITokenMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMinter *ITokenMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMinter *ITokenMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMinter.Contract.contract.Transact(opts, method, params...)
}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterCaller) GetLocalToken(opts *bind.CallOpts, remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITokenMinter.contract.Call(opts, &out, "getLocalToken", remoteDomain, remoteToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterSession) GetLocalToken(remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	return _ITokenMinter.Contract.GetLocalToken(&_ITokenMinter.CallOpts, remoteDomain, remoteToken)
}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterCallerSession) GetLocalToken(remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	return _ITokenMinter.Contract.GetLocalToken(&_ITokenMinter.CallOpts, remoteDomain, remoteToken)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterTransactor) Burn(opts *bind.TransactOpts, burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "burn", burnToken, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterSession) Burn(burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Burn(&_ITokenMinter.TransactOpts, burnToken, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterTransactorSession) Burn(burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Burn(&_ITokenMinter.TransactOpts, burnToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterTransactor) Mint(opts *bind.TransactOpts, sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "mint", sourceDomain, burnToken, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterSession) Mint(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Mint(&_ITokenMinter.TransactOpts, sourceDomain, burnToken, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterTransactorSession) Mint(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Mint(&_ITokenMinter.TransactOpts, sourceDomain, burnToken, to, amount)
}

// MinimalForwarderLibMetaData contains all meta data concerning the MinimalForwarderLib contract.
var MinimalForwarderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122062b83228cb4cad1893c1d418e8705515a544b4a769800ee731c552ce3f9cf81c64736f6c63430008110033",
}

// MinimalForwarderLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MinimalForwarderLibMetaData.ABI instead.
var MinimalForwarderLibABI = MinimalForwarderLibMetaData.ABI

// MinimalForwarderLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MinimalForwarderLibMetaData.Bin instead.
var MinimalForwarderLibBin = MinimalForwarderLibMetaData.Bin

// DeployMinimalForwarderLib deploys a new Ethereum contract, binding an instance of MinimalForwarderLib to it.
func DeployMinimalForwarderLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MinimalForwarderLib, error) {
	parsed, err := MinimalForwarderLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MinimalForwarderLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MinimalForwarderLib{MinimalForwarderLibCaller: MinimalForwarderLibCaller{contract: contract}, MinimalForwarderLibTransactor: MinimalForwarderLibTransactor{contract: contract}, MinimalForwarderLibFilterer: MinimalForwarderLibFilterer{contract: contract}}, nil
}

// MinimalForwarderLib is an auto generated Go binding around an Ethereum contract.
type MinimalForwarderLib struct {
	MinimalForwarderLibCaller     // Read-only binding to the contract
	MinimalForwarderLibTransactor // Write-only binding to the contract
	MinimalForwarderLibFilterer   // Log filterer for contract events
}

// MinimalForwarderLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinimalForwarderLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinimalForwarderLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinimalForwarderLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinimalForwarderLibSession struct {
	Contract     *MinimalForwarderLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MinimalForwarderLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinimalForwarderLibCallerSession struct {
	Contract *MinimalForwarderLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MinimalForwarderLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinimalForwarderLibTransactorSession struct {
	Contract     *MinimalForwarderLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MinimalForwarderLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinimalForwarderLibRaw struct {
	Contract *MinimalForwarderLib // Generic contract binding to access the raw methods on
}

// MinimalForwarderLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinimalForwarderLibCallerRaw struct {
	Contract *MinimalForwarderLibCaller // Generic read-only contract binding to access the raw methods on
}

// MinimalForwarderLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinimalForwarderLibTransactorRaw struct {
	Contract *MinimalForwarderLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinimalForwarderLib creates a new instance of MinimalForwarderLib, bound to a specific deployed contract.
func NewMinimalForwarderLib(address common.Address, backend bind.ContractBackend) (*MinimalForwarderLib, error) {
	contract, err := bindMinimalForwarderLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderLib{MinimalForwarderLibCaller: MinimalForwarderLibCaller{contract: contract}, MinimalForwarderLibTransactor: MinimalForwarderLibTransactor{contract: contract}, MinimalForwarderLibFilterer: MinimalForwarderLibFilterer{contract: contract}}, nil
}

// NewMinimalForwarderLibCaller creates a new read-only instance of MinimalForwarderLib, bound to a specific deployed contract.
func NewMinimalForwarderLibCaller(address common.Address, caller bind.ContractCaller) (*MinimalForwarderLibCaller, error) {
	contract, err := bindMinimalForwarderLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderLibCaller{contract: contract}, nil
}

// NewMinimalForwarderLibTransactor creates a new write-only instance of MinimalForwarderLib, bound to a specific deployed contract.
func NewMinimalForwarderLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MinimalForwarderLibTransactor, error) {
	contract, err := bindMinimalForwarderLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderLibTransactor{contract: contract}, nil
}

// NewMinimalForwarderLibFilterer creates a new log filterer instance of MinimalForwarderLib, bound to a specific deployed contract.
func NewMinimalForwarderLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MinimalForwarderLibFilterer, error) {
	contract, err := bindMinimalForwarderLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderLibFilterer{contract: contract}, nil
}

// bindMinimalForwarderLib binds a generic wrapper to an already deployed contract.
func bindMinimalForwarderLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinimalForwarderLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinimalForwarderLib *MinimalForwarderLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinimalForwarderLib.Contract.MinimalForwarderLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinimalForwarderLib *MinimalForwarderLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinimalForwarderLib.Contract.MinimalForwarderLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinimalForwarderLib *MinimalForwarderLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinimalForwarderLib.Contract.MinimalForwarderLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinimalForwarderLib *MinimalForwarderLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinimalForwarderLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinimalForwarderLib *MinimalForwarderLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinimalForwarderLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinimalForwarderLib *MinimalForwarderLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinimalForwarderLib.Contract.contract.Transact(opts, method, params...)
}

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220dce9c2238eef7e699856efa83369c374cb4d872745cd3deefd0571ff5fe22fbb64736f6c63430008110033",
}

// RequestLibABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestLibMetaData.ABI instead.
var RequestLibABI = RequestLibMetaData.ABI

// RequestLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequestLibMetaData.Bin instead.
var RequestLibBin = RequestLibMetaData.Bin

// DeployRequestLib deploys a new Ethereum contract, binding an instance of RequestLib to it.
func DeployRequestLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestLib, error) {
	parsed, err := RequestLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequestLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestLib{RequestLibCaller: RequestLibCaller{contract: contract}, RequestLibTransactor: RequestLibTransactor{contract: contract}, RequestLibFilterer: RequestLibFilterer{contract: contract}}, nil
}

// RequestLib is an auto generated Go binding around an Ethereum contract.
type RequestLib struct {
	RequestLibCaller     // Read-only binding to the contract
	RequestLibTransactor // Write-only binding to the contract
	RequestLibFilterer   // Log filterer for contract events
}

// RequestLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestLibSession struct {
	Contract     *RequestLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestLibCallerSession struct {
	Contract *RequestLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RequestLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestLibTransactorSession struct {
	Contract     *RequestLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RequestLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestLibRaw struct {
	Contract *RequestLib // Generic contract binding to access the raw methods on
}

// RequestLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestLibCallerRaw struct {
	Contract *RequestLibCaller // Generic read-only contract binding to access the raw methods on
}

// RequestLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestLibTransactorRaw struct {
	Contract *RequestLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestLib creates a new instance of RequestLib, bound to a specific deployed contract.
func NewRequestLib(address common.Address, backend bind.ContractBackend) (*RequestLib, error) {
	contract, err := bindRequestLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestLib{RequestLibCaller: RequestLibCaller{contract: contract}, RequestLibTransactor: RequestLibTransactor{contract: contract}, RequestLibFilterer: RequestLibFilterer{contract: contract}}, nil
}

// NewRequestLibCaller creates a new read-only instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibCaller(address common.Address, caller bind.ContractCaller) (*RequestLibCaller, error) {
	contract, err := bindRequestLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestLibCaller{contract: contract}, nil
}

// NewRequestLibTransactor creates a new write-only instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestLibTransactor, error) {
	contract, err := bindRequestLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestLibTransactor{contract: contract}, nil
}

// NewRequestLibFilterer creates a new log filterer instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestLibFilterer, error) {
	contract, err := bindRequestLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestLibFilterer{contract: contract}, nil
}

// bindRequestLib binds a generic wrapper to an already deployed contract.
func bindRequestLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestLib *RequestLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestLib.Contract.RequestLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestLib *RequestLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestLib.Contract.RequestLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestLib *RequestLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestLib.Contract.RequestLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestLib *RequestLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestLib *RequestLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestLib *RequestLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestLib.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220db481b01e27340107f8773a209c67cc43e84d8758d685f09c04b63dab17b8b4164736f6c63430008110033",
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

// SlicerLibMetaData contains all meta data concerning the SlicerLib contract.
var SlicerLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206bc0ea76823bc81c8dc727db42575df1d55cb6d977435627f0cf6b152e41d33b64736f6c63430008110033",
}

// SlicerLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SlicerLibMetaData.ABI instead.
var SlicerLibABI = SlicerLibMetaData.ABI

// SlicerLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlicerLibMetaData.Bin instead.
var SlicerLibBin = SlicerLibMetaData.Bin

// DeploySlicerLib deploys a new Ethereum contract, binding an instance of SlicerLib to it.
func DeploySlicerLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlicerLib, error) {
	parsed, err := SlicerLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlicerLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlicerLib{SlicerLibCaller: SlicerLibCaller{contract: contract}, SlicerLibTransactor: SlicerLibTransactor{contract: contract}, SlicerLibFilterer: SlicerLibFilterer{contract: contract}}, nil
}

// SlicerLib is an auto generated Go binding around an Ethereum contract.
type SlicerLib struct {
	SlicerLibCaller     // Read-only binding to the contract
	SlicerLibTransactor // Write-only binding to the contract
	SlicerLibFilterer   // Log filterer for contract events
}

// SlicerLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlicerLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlicerLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlicerLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlicerLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlicerLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlicerLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlicerLibSession struct {
	Contract     *SlicerLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlicerLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlicerLibCallerSession struct {
	Contract *SlicerLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SlicerLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlicerLibTransactorSession struct {
	Contract     *SlicerLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SlicerLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlicerLibRaw struct {
	Contract *SlicerLib // Generic contract binding to access the raw methods on
}

// SlicerLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlicerLibCallerRaw struct {
	Contract *SlicerLibCaller // Generic read-only contract binding to access the raw methods on
}

// SlicerLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlicerLibTransactorRaw struct {
	Contract *SlicerLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlicerLib creates a new instance of SlicerLib, bound to a specific deployed contract.
func NewSlicerLib(address common.Address, backend bind.ContractBackend) (*SlicerLib, error) {
	contract, err := bindSlicerLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlicerLib{SlicerLibCaller: SlicerLibCaller{contract: contract}, SlicerLibTransactor: SlicerLibTransactor{contract: contract}, SlicerLibFilterer: SlicerLibFilterer{contract: contract}}, nil
}

// NewSlicerLibCaller creates a new read-only instance of SlicerLib, bound to a specific deployed contract.
func NewSlicerLibCaller(address common.Address, caller bind.ContractCaller) (*SlicerLibCaller, error) {
	contract, err := bindSlicerLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlicerLibCaller{contract: contract}, nil
}

// NewSlicerLibTransactor creates a new write-only instance of SlicerLib, bound to a specific deployed contract.
func NewSlicerLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SlicerLibTransactor, error) {
	contract, err := bindSlicerLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlicerLibTransactor{contract: contract}, nil
}

// NewSlicerLibFilterer creates a new log filterer instance of SlicerLib, bound to a specific deployed contract.
func NewSlicerLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SlicerLibFilterer, error) {
	contract, err := bindSlicerLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlicerLibFilterer{contract: contract}, nil
}

// bindSlicerLib binds a generic wrapper to an already deployed contract.
func bindSlicerLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SlicerLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlicerLib *SlicerLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlicerLib.Contract.SlicerLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlicerLib *SlicerLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlicerLib.Contract.SlicerLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlicerLib *SlicerLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlicerLib.Contract.SlicerLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlicerLib *SlicerLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlicerLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlicerLib *SlicerLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlicerLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlicerLib *SlicerLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlicerLib.Contract.contract.Transact(opts, method, params...)
}

// SynapseCCTPMetaData contains all meta data concerning the SynapseCCTP contract.
var SynapseCCTPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"tokenMessenger_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CCTPMessageNotReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwarderDeploymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectRequestLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOrRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LocalCCTPTokenNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoteCCTPDeploymentNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoteCCTPTokenNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SliceOverrun\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownRequestVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageTransmitter\",\"outputs\":[{\"internalType\":\"contractIMessageTransmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"}],\"name\":\"receiveCircleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remoteDomainConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"synapseCCTP\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"swapParams\",\"type\":\"bytes\"}],\"name\":\"sendCircleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"setLocalToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"remoteSynapseCCTP\",\"type\":\"address\"}],\"name\":\"setRemoteDomainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f879a41a": "getLocalToken(uint32,address)",
		"8d3638f4": "localDomain()",
		"7b04c181": "messageTransmitter()",
		"4a5ae51d": "receiveCircleToken(bytes,bytes,uint32,bytes)",
		"e9259ab9": "remoteDomainConfig(uint256)",
		"304ddb4c": "sendCircleToken(address,uint256,address,uint256,uint32,bytes)",
		"393e5b60": "setLocalToken(uint32,address)",
		"e9bbb36d": "setRemoteDomainConfig(uint256,uint32,address)",
		"46117830": "tokenMessenger()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b506040516200239e3803806200239e83398101604081905262000034916200013e565b6001600160a01b03811660c081905260408051632c12192160e01b81529051632c121921916004808201926020929091908290030181865afa1580156200007f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000a591906200013e565b6001600160a01b031660a08190526040805163234d8e3d60e21b81529051638d3638f4916004808201926020929091908290030181865afa158015620000ef573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000115919062000165565b63ffffffff16608052506200018d565b6001600160a01b03811681146200013b57600080fd5b50565b6000602082840312156200015157600080fd5b81516200015e8162000125565b9392505050565b6000602082840312156200017857600080fd5b815163ffffffff811681146200015e57600080fd5b60805160a05160c0516121a3620001fb6000396000818160d5015281816104f40152818161062c01528181610d2001528181610dc30152610e05015260008181610139015281816102bc01526110850152600081816101600152818161034f015261087c01526121a36000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80637b04c18111610076578063e9259ab91161005b578063e9259ab914610197578063e9bbb36d1461020e578063f879a41a1461029957600080fd5b80637b04c181146101345780638d3638f41461015b57600080fd5b8063304ddb4c146100a8578063393e5b60146100bd57806346117830146100d05780634a5ae51d14610121575b600080fd5b6100bb6100b6366004611c88565b6102ac565b005b6100bb6100cb366004611d0d565b610628565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100bb61012f366004611d86565b61085a565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610118565b6101dd6101a5366004611e1f565b60006020819052908152604090205463ffffffff811690640100000000900473ffffffffffffffffffffffffffffffffffffffff1682565b6040805163ffffffff909316835273ffffffffffffffffffffffffffffffffffffffff909116602083015201610118565b6100bb61021c366004611e38565b60408051808201825263ffffffff938416815273ffffffffffffffffffffffffffffffffffffffff9283166020808301918252600096875286905291909420935184549151909216640100000000027fffffffffffffffff0000000000000000000000000000000000000000000000009091169190921617179055565b6100f76102a7366004611d0d565b610970565b6102b684846109bf565b925060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16638371744e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610325573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103499190611e78565b604080517f000000000000000000000000000000000000000000000000000000000000000060e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260c083901b7fffffffffffffffff000000000000000000000000000000000000000000000000166024820152606088811b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000908116602c8401528284018990528b821b169082015281516054818303018152607490910190915290915060009061042490859085610b18565b60008881526020818152604080832081518083019092525463ffffffff81168252640100000000900473ffffffffffffffffffffffffffffffffffffffff16918101829052929350908190036104a6576040517fa86a3b0e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151835160208086019190912067ffffffff0000000083831b1663ffffffff8a161760009081529152604090206104dd8a8a610cdd565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663f856ddb68a84868e6105278288610e4f565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b168152600481019590955263ffffffff939093166024850152604484019190915273ffffffffffffffffffffffffffffffffffffffff166064830152608482015260a4016020604051808303816000875af11580156105b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d79190611e78565b50807f4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a8c888d8d8d8b60405161061296959493929190611ef2565b60405180910390a2505050505050505050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663cb75c11c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610695573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b99190611f55565b905060008173ffffffffffffffffffffffffffffffffffffffff166378a0565e8561070d8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1690565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815263ffffffff9290921660048301526024820152604401602060405180830381865afa15801561076c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107909190611f55565b905073ffffffffffffffffffffffffffffffffffffffff81166107df576040517f74e3d32e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b77ffffffff000000000000000000000000000000000000000060a09490941b9390931673ffffffffffffffffffffffffffffffffffffffff92831617600090815260016020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169390921692909217905550565b60006108668383610e94565b825160208085019190912067ffffffff000000007f0000000000000000000000000000000000000000000000000000000000000000831b1663ffffffff87161760009081529152604090209091506108c18888888885610fd8565b6000806108cd84611100565b9150915060006108dd83836111a1565b9092509050600080806108f186868a6111ac565b6040805173ffffffffffffffffffffffffffffffffffffffff8b81168252602082018a9052848116828401526060820184905291519497509295509093508992908616917feaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f919081900360800190a35050505050505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81811660a084901b77ffffffff00000000000000000000000000000000000000001617600090815260016020526040902054165b92915050565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff8516906370a0823190602401602060405180830381865afa158015610a2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a529190611f72565b9050610a7673ffffffffffffffffffffffffffffffffffffffff85163330866111ec565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152819073ffffffffffffffffffffffffffffffffffffffff8616906370a0823190602401602060405180830381865afa158015610ae2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b069190611f72565b610b109190611fba565b949350505050565b6060600163ffffffff85161115610b5b576040517f523fa8d500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b6760006020611fcd565b610b72906020611fcd565b610b7d906014611fcd565b835114610bb6576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8416158015610bca5750815115155b15610c01576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff84166001148015610c7b5750610c1e60006020611fcd565b610c29906020611fcd565b610c34906014611fcd565b610c4060006020611fcd565b610c4b906020611fcd565b610c56906014611fcd565b610c61906020611fcd565b610c6c906020611fcd565b610c769190611fba565b825114155b15610cb2576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282604051602001610cc5929190611fe0565b60405160208183030381529060405290509392505050565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301526000919084169063dd62ed3e90604401602060405180830381865afa158015610d73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d979190611f72565b905081811015610e4a578015610de957610de973ffffffffffffffffffffffffffffffffffffffff84167f000000000000000000000000000000000000000000000000000000000000000060006112b0565b610e4a73ffffffffffffffffffffffffffffffffffffffff84167f00000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6112b0565b505050565b6000610e8d610e7473ffffffffffffffffffffffffffffffffffffffff851684611437565b73ffffffffffffffffffffffffffffffffffffffff1690565b9392505050565b6000600163ffffffff84161115610ed7576040517f523fa8d500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8316158015610f0d5750610ef260006020611fcd565b610efd906020611fcd565b610f08906014611fcd565b825114155b15610f44576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff83166001148015610f925750610f6160006020611fcd565b610f6c906020611fcd565b610f77906014611fcd565b610f82906020611fcd565b610f8d906020611fcd565b825114155b15610fc9576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81516020830160801b17610e8d565b6000610fe382611523565b905060006357ecfd2860e01b87878787604051602401611006949392919061203a565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152905060006110aa73ffffffffffffffffffffffffffffffffffffffff84167f0000000000000000000000000000000000000000000000000000000000000000846115dc565b9050808060200190518101906110c09190612061565b6110f6576040517f182f34eb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050565b60008060008061110f856115eb565b73ffffffffffffffffffffffffffffffffffffffff80831660a086901b77ffffffff00000000000000000000000000000000000000001617600090815260016020526040902054169750955091935090915084905061119a576040517f5ec62f3300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050915091565b8060005b9250929050565b60008060006111ba84611652565b92508591508490506111e373ffffffffffffffffffffffffffffffffffffffff83168483611674565b93509350939050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526112aa9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526116ca565b50505050565b80158061135057506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015284169063dd62ed3e90604401602060405180830381865afa15801561132a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134e9190611f72565b155b6113e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e63650000000000000000000060648201526084015b60405180910390fd5b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610e4a9084907f095ea7b30000000000000000000000000000000000000000000000000000000090606401611246565b6000610e8d83836040518060400160405280602081526020017f602036038060203d373d3d3d923d343d355af13d82803e903d91601e57fd5bf38152506040516020016114849190612083565b6040516020818303038152906040528051906020012060405160200161150a939291907fff00000000000000000000000000000000000000000000000000000000000000815260609390931b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660018401526015830191909152603582015260550190565b6040516020818303038152906040528051906020012090565b6000806040518060400160405280602081526020017f602036038060203d373d3d3d923d343d355af13d82803e903d91601e57fd5bf381525060405160200161156c9190612083565b6040516020818303038152906040529050828151602083016000f5915073ffffffffffffffffffffffffffffffffffffffff82166115d6576040517f27afa9fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50919050565b6060610b1084848460006117d6565b6000808080806115fd81875b9061183b565b90508060e01c945067ffffffffffffffff8160a01c16935073ffffffffffffffffffffffffffffffffffffffff811692506116456000602061163f9190611fcd565b876115f7565b60001c9150509193509193565b60006109b9611662826020611fcd565b61166d906020611fcd565b83906118d6565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610e4a9084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401611246565b600061172c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166119749092919063ffffffff16565b805190915015610e4a578080602001905181019061174a9190612061565b610e4a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016113d8565b606061183273ffffffffffffffffffffffffffffffffffffffff8516846040516020016118049291906120ef565b60408051601f1981840301815291905273ffffffffffffffffffffffffffffffffffffffff87169084611983565b95945050505050565b6000608083901c6fffffffffffffffffffffffffffffffff841680841061188e576040517fdfc52d7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b808460200111156118cb576040517f4b72f29d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b509091015192915050565b6000608083901c6fffffffffffffffffffffffffffffffff8416808410611929576040517fdfc52d7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80846014011115611966576040517f4b72f29d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b509091015160601c92915050565b6060610b1084846000856119a5565b6060610b10848484604051806060016040528060298152602001612145602991395b606082471015611a37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016113d8565b73ffffffffffffffffffffffffffffffffffffffff85163b611ab5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016113d8565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611ade9190612115565b60006040518083038185875af1925050503d8060008114611b1b576040519150601f19603f3d011682016040523d82523d6000602084013e611b20565b606091505b5091509150611b30828286611b3b565b979650505050505050565b60608315611b4a575081610e8d565b825115611b5a5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113d89190612131565b73ffffffffffffffffffffffffffffffffffffffff81168114611bb057600080fd5b50565b803563ffffffff81168114611bc757600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112611c0c57600080fd5b813567ffffffffffffffff80821115611c2757611c27611bcc565b604051601f8301601f19908116603f01168101908282118183101715611c4f57611c4f611bcc565b81604052838152866020858801011115611c6857600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060008060008060c08789031215611ca157600080fd5b8635611cac81611b8e565b9550602087013594506040870135611cc381611b8e565b935060608701359250611cd860808801611bb3565b915060a087013567ffffffffffffffff811115611cf457600080fd5b611d0089828a01611bfb565b9150509295509295509295565b60008060408385031215611d2057600080fd5b611d2983611bb3565b91506020830135611d3981611b8e565b809150509250929050565b60008083601f840112611d5657600080fd5b50813567ffffffffffffffff811115611d6e57600080fd5b6020830191508360208285010111156111a557600080fd5b60008060008060008060808789031215611d9f57600080fd5b863567ffffffffffffffff80821115611db757600080fd5b611dc38a838b01611d44565b90985096506020890135915080821115611ddc57600080fd5b611de88a838b01611d44565b9096509450849150611dfc60408a01611bb3565b93506060890135915080821115611e1257600080fd5b50611d0089828a01611bfb565b600060208284031215611e3157600080fd5b5035919050565b600080600060608486031215611e4d57600080fd5b83359250611e5d60208501611bb3565b91506040840135611e6d81611b8e565b809150509250925092565b600060208284031215611e8a57600080fd5b815167ffffffffffffffff81168114610e8d57600080fd5b60005b83811015611ebd578181015183820152602001611ea5565b50506000910152565b60008151808452611ede816020860160208601611ea2565b601f01601f19169290920160200192915050565b86815267ffffffffffffffff8616602082015273ffffffffffffffffffffffffffffffffffffffff8516604082015283606082015263ffffffff8316608082015260c060a08201526000611f4960c0830184611ec6565b98975050505050505050565b600060208284031215611f6757600080fd5b8151610e8d81611b8e565b600060208284031215611f8457600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156109b9576109b9611f8b565b808201808211156109b9576109b9611f8b565b60008351611ff2818460208801611ea2565b835190830190612006818360208801611ea2565b01949350505050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b60408152600061204e60408301868861200f565b8281036020840152611b3081858761200f565b60006020828403121561207357600080fd5b81518015158114610e8d57600080fd5b7f7f000000000000000000000000000000000000000000000000000000000000008152600082516120bb816001850160208701611ea2565b7f3d5260203df300000000000000000000000000000000000000000000000000006001939091019283015250600701919050565b82815260008251612107816020850160208701611ea2565b919091016020019392505050565b60008251612127818460208701611ea2565b9190910192915050565b602081526000610e8d6020830184611ec656fe416464726573733a206c6f772d6c6576656c2063616c6c20776974682076616c7565206661696c6564a2646970667358221220b20f91da13ced161bd28ee9ec04443facf85562f061afbf72510e53e231d859164736f6c63430008110033",
}

// SynapseCCTPABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseCCTPMetaData.ABI instead.
var SynapseCCTPABI = SynapseCCTPMetaData.ABI

// Deprecated: Use SynapseCCTPMetaData.Sigs instead.
// SynapseCCTPFuncSigs maps the 4-byte function signature to its string representation.
var SynapseCCTPFuncSigs = SynapseCCTPMetaData.Sigs

// SynapseCCTPBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseCCTPMetaData.Bin instead.
var SynapseCCTPBin = SynapseCCTPMetaData.Bin

// DeploySynapseCCTP deploys a new Ethereum contract, binding an instance of SynapseCCTP to it.
func DeploySynapseCCTP(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger_ common.Address) (common.Address, *types.Transaction, *SynapseCCTP, error) {
	parsed, err := SynapseCCTPMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseCCTPBin), backend, tokenMessenger_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseCCTP{SynapseCCTPCaller: SynapseCCTPCaller{contract: contract}, SynapseCCTPTransactor: SynapseCCTPTransactor{contract: contract}, SynapseCCTPFilterer: SynapseCCTPFilterer{contract: contract}}, nil
}

// SynapseCCTP is an auto generated Go binding around an Ethereum contract.
type SynapseCCTP struct {
	SynapseCCTPCaller     // Read-only binding to the contract
	SynapseCCTPTransactor // Write-only binding to the contract
	SynapseCCTPFilterer   // Log filterer for contract events
}

// SynapseCCTPCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseCCTPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseCCTPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseCCTPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseCCTPSession struct {
	Contract     *SynapseCCTP      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseCCTPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseCCTPCallerSession struct {
	Contract *SynapseCCTPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SynapseCCTPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseCCTPTransactorSession struct {
	Contract     *SynapseCCTPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SynapseCCTPRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseCCTPRaw struct {
	Contract *SynapseCCTP // Generic contract binding to access the raw methods on
}

// SynapseCCTPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseCCTPCallerRaw struct {
	Contract *SynapseCCTPCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseCCTPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseCCTPTransactorRaw struct {
	Contract *SynapseCCTPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseCCTP creates a new instance of SynapseCCTP, bound to a specific deployed contract.
func NewSynapseCCTP(address common.Address, backend bind.ContractBackend) (*SynapseCCTP, error) {
	contract, err := bindSynapseCCTP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTP{SynapseCCTPCaller: SynapseCCTPCaller{contract: contract}, SynapseCCTPTransactor: SynapseCCTPTransactor{contract: contract}, SynapseCCTPFilterer: SynapseCCTPFilterer{contract: contract}}, nil
}

// NewSynapseCCTPCaller creates a new read-only instance of SynapseCCTP, bound to a specific deployed contract.
func NewSynapseCCTPCaller(address common.Address, caller bind.ContractCaller) (*SynapseCCTPCaller, error) {
	contract, err := bindSynapseCCTP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPCaller{contract: contract}, nil
}

// NewSynapseCCTPTransactor creates a new write-only instance of SynapseCCTP, bound to a specific deployed contract.
func NewSynapseCCTPTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseCCTPTransactor, error) {
	contract, err := bindSynapseCCTP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPTransactor{contract: contract}, nil
}

// NewSynapseCCTPFilterer creates a new log filterer instance of SynapseCCTP, bound to a specific deployed contract.
func NewSynapseCCTPFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseCCTPFilterer, error) {
	contract, err := bindSynapseCCTP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFilterer{contract: contract}, nil
}

// bindSynapseCCTP binds a generic wrapper to an already deployed contract.
func bindSynapseCCTP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseCCTPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTP *SynapseCCTPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTP.Contract.SynapseCCTPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTP *SynapseCCTPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SynapseCCTPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTP *SynapseCCTPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SynapseCCTPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTP *SynapseCCTPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTP *SynapseCCTPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTP *SynapseCCTPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.contract.Transact(opts, method, params...)
}

// GetLocalToken is a free data retrieval call binding the contract method 0xf879a41a.
//
// Solidity: function getLocalToken(uint32 remoteDomain, address remoteToken) view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) GetLocalToken(opts *bind.CallOpts, remoteDomain uint32, remoteToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "getLocalToken", remoteDomain, remoteToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocalToken is a free data retrieval call binding the contract method 0xf879a41a.
//
// Solidity: function getLocalToken(uint32 remoteDomain, address remoteToken) view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) GetLocalToken(remoteDomain uint32, remoteToken common.Address) (common.Address, error) {
	return _SynapseCCTP.Contract.GetLocalToken(&_SynapseCCTP.CallOpts, remoteDomain, remoteToken)
}

// GetLocalToken is a free data retrieval call binding the contract method 0xf879a41a.
//
// Solidity: function getLocalToken(uint32 remoteDomain, address remoteToken) view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) GetLocalToken(remoteDomain uint32, remoteToken common.Address) (common.Address, error) {
	return _SynapseCCTP.Contract.GetLocalToken(&_SynapseCCTP.CallOpts, remoteDomain, remoteToken)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseCCTP *SynapseCCTPCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseCCTP *SynapseCCTPSession) LocalDomain() (uint32, error) {
	return _SynapseCCTP.Contract.LocalDomain(&_SynapseCCTP.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseCCTP *SynapseCCTPCallerSession) LocalDomain() (uint32, error) {
	return _SynapseCCTP.Contract.LocalDomain(&_SynapseCCTP.CallOpts)
}

// MessageTransmitter is a free data retrieval call binding the contract method 0x7b04c181.
//
// Solidity: function messageTransmitter() view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) MessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "messageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageTransmitter is a free data retrieval call binding the contract method 0x7b04c181.
//
// Solidity: function messageTransmitter() view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) MessageTransmitter() (common.Address, error) {
	return _SynapseCCTP.Contract.MessageTransmitter(&_SynapseCCTP.CallOpts)
}

// MessageTransmitter is a free data retrieval call binding the contract method 0x7b04c181.
//
// Solidity: function messageTransmitter() view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) MessageTransmitter() (common.Address, error) {
	return _SynapseCCTP.Contract.MessageTransmitter(&_SynapseCCTP.CallOpts)
}

// RemoteDomainConfig is a free data retrieval call binding the contract method 0xe9259ab9.
//
// Solidity: function remoteDomainConfig(uint256 ) view returns(uint32 domain, address synapseCCTP)
func (_SynapseCCTP *SynapseCCTPCaller) RemoteDomainConfig(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Domain      uint32
	SynapseCCTP common.Address
}, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "remoteDomainConfig", arg0)

	outstruct := new(struct {
		Domain      uint32
		SynapseCCTP common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Domain = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.SynapseCCTP = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RemoteDomainConfig is a free data retrieval call binding the contract method 0xe9259ab9.
//
// Solidity: function remoteDomainConfig(uint256 ) view returns(uint32 domain, address synapseCCTP)
func (_SynapseCCTP *SynapseCCTPSession) RemoteDomainConfig(arg0 *big.Int) (struct {
	Domain      uint32
	SynapseCCTP common.Address
}, error) {
	return _SynapseCCTP.Contract.RemoteDomainConfig(&_SynapseCCTP.CallOpts, arg0)
}

// RemoteDomainConfig is a free data retrieval call binding the contract method 0xe9259ab9.
//
// Solidity: function remoteDomainConfig(uint256 ) view returns(uint32 domain, address synapseCCTP)
func (_SynapseCCTP *SynapseCCTPCallerSession) RemoteDomainConfig(arg0 *big.Int) (struct {
	Domain      uint32
	SynapseCCTP common.Address
}, error) {
	return _SynapseCCTP.Contract.RemoteDomainConfig(&_SynapseCCTP.CallOpts, arg0)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) TokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) TokenMessenger() (common.Address, error) {
	return _SynapseCCTP.Contract.TokenMessenger(&_SynapseCCTP.CallOpts)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) TokenMessenger() (common.Address, error) {
	return _SynapseCCTP.Contract.TokenMessenger(&_SynapseCCTP.CallOpts)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) ReceiveCircleToken(opts *bind.TransactOpts, message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "receiveCircleToken", message, signature, requestVersion, formattedRequest)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) returns()
func (_SynapseCCTP *SynapseCCTPSession) ReceiveCircleToken(message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.ReceiveCircleToken(&_SynapseCCTP.TransactOpts, message, signature, requestVersion, formattedRequest)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) ReceiveCircleToken(message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.ReceiveCircleToken(&_SynapseCCTP.TransactOpts, message, signature, requestVersion, formattedRequest)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SendCircleToken(opts *bind.TransactOpts, recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "sendCircleToken", recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_SynapseCCTP *SynapseCCTPSession) SendCircleToken(recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SendCircleToken(&_SynapseCCTP.TransactOpts, recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SendCircleToken(recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SendCircleToken(&_SynapseCCTP.TransactOpts, recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SetLocalToken is a paid mutator transaction binding the contract method 0x393e5b60.
//
// Solidity: function setLocalToken(uint32 remoteDomain, address remoteToken) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetLocalToken(opts *bind.TransactOpts, remoteDomain uint32, remoteToken common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setLocalToken", remoteDomain, remoteToken)
}

// SetLocalToken is a paid mutator transaction binding the contract method 0x393e5b60.
//
// Solidity: function setLocalToken(uint32 remoteDomain, address remoteToken) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetLocalToken(remoteDomain uint32, remoteToken common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetLocalToken(&_SynapseCCTP.TransactOpts, remoteDomain, remoteToken)
}

// SetLocalToken is a paid mutator transaction binding the contract method 0x393e5b60.
//
// Solidity: function setLocalToken(uint32 remoteDomain, address remoteToken) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SetLocalToken(remoteDomain uint32, remoteToken common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetLocalToken(&_SynapseCCTP.TransactOpts, remoteDomain, remoteToken)
}

// SetRemoteDomainConfig is a paid mutator transaction binding the contract method 0xe9bbb36d.
//
// Solidity: function setRemoteDomainConfig(uint256 remoteChainId, uint32 remoteDomain, address remoteSynapseCCTP) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetRemoteDomainConfig(opts *bind.TransactOpts, remoteChainId *big.Int, remoteDomain uint32, remoteSynapseCCTP common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setRemoteDomainConfig", remoteChainId, remoteDomain, remoteSynapseCCTP)
}

// SetRemoteDomainConfig is a paid mutator transaction binding the contract method 0xe9bbb36d.
//
// Solidity: function setRemoteDomainConfig(uint256 remoteChainId, uint32 remoteDomain, address remoteSynapseCCTP) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetRemoteDomainConfig(remoteChainId *big.Int, remoteDomain uint32, remoteSynapseCCTP common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetRemoteDomainConfig(&_SynapseCCTP.TransactOpts, remoteChainId, remoteDomain, remoteSynapseCCTP)
}

// SetRemoteDomainConfig is a paid mutator transaction binding the contract method 0xe9bbb36d.
//
// Solidity: function setRemoteDomainConfig(uint256 remoteChainId, uint32 remoteDomain, address remoteSynapseCCTP) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SetRemoteDomainConfig(remoteChainId *big.Int, remoteDomain uint32, remoteSynapseCCTP common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetRemoteDomainConfig(&_SynapseCCTP.TransactOpts, remoteChainId, remoteDomain, remoteSynapseCCTP)
}

// SynapseCCTPCircleRequestFulfilledIterator is returned from FilterCircleRequestFulfilled and is used to iterate over the raw logs and unpacked data for CircleRequestFulfilled events raised by the SynapseCCTP contract.
type SynapseCCTPCircleRequestFulfilledIterator struct {
	Event *SynapseCCTPCircleRequestFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SynapseCCTPCircleRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPCircleRequestFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SynapseCCTPCircleRequestFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SynapseCCTPCircleRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPCircleRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPCircleRequestFulfilled represents a CircleRequestFulfilled event raised by the SynapseCCTP contract.
type SynapseCCTPCircleRequestFulfilled struct {
	Recipient common.Address
	MintToken common.Address
	Fee       *big.Int
	Token     common.Address
	Amount    *big.Int
	Kappa     [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestFulfilled is a free log retrieval operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed kappa)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterCircleRequestFulfilled(opts *bind.FilterOpts, recipient []common.Address, kappa [][32]byte) (*SynapseCCTPCircleRequestFulfilledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "CircleRequestFulfilled", recipientRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPCircleRequestFulfilledIterator{contract: _SynapseCCTP.contract, event: "CircleRequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchCircleRequestFulfilled is a free log subscription operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed kappa)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchCircleRequestFulfilled(opts *bind.WatchOpts, sink chan<- *SynapseCCTPCircleRequestFulfilled, recipient []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "CircleRequestFulfilled", recipientRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPCircleRequestFulfilled)
				if err := _SynapseCCTP.contract.UnpackLog(event, "CircleRequestFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCircleRequestFulfilled is a log parse operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed kappa)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseCircleRequestFulfilled(log types.Log) (*SynapseCCTPCircleRequestFulfilled, error) {
	event := new(SynapseCCTPCircleRequestFulfilled)
	if err := _SynapseCCTP.contract.UnpackLog(event, "CircleRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPCircleRequestSentIterator is returned from FilterCircleRequestSent and is used to iterate over the raw logs and unpacked data for CircleRequestSent events raised by the SynapseCCTP contract.
type SynapseCCTPCircleRequestSentIterator struct {
	Event *SynapseCCTPCircleRequestSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SynapseCCTPCircleRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPCircleRequestSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SynapseCCTPCircleRequestSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SynapseCCTPCircleRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPCircleRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPCircleRequestSent represents a CircleRequestSent event raised by the SynapseCCTP contract.
type SynapseCCTPCircleRequestSent struct {
	ChainId          *big.Int
	Nonce            uint64
	Token            common.Address
	Amount           *big.Int
	RequestVersion   uint32
	FormattedRequest []byte
	Kappa            [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestSent is a free log retrieval operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed kappa)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterCircleRequestSent(opts *bind.FilterOpts, kappa [][32]byte) (*SynapseCCTPCircleRequestSentIterator, error) {

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "CircleRequestSent", kappaRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPCircleRequestSentIterator{contract: _SynapseCCTP.contract, event: "CircleRequestSent", logs: logs, sub: sub}, nil
}

// WatchCircleRequestSent is a free log subscription operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed kappa)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchCircleRequestSent(opts *bind.WatchOpts, sink chan<- *SynapseCCTPCircleRequestSent, kappa [][32]byte) (event.Subscription, error) {

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "CircleRequestSent", kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPCircleRequestSent)
				if err := _SynapseCCTP.contract.UnpackLog(event, "CircleRequestSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCircleRequestSent is a log parse operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed kappa)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseCircleRequestSent(log types.Log) (*SynapseCCTPCircleRequestSent, error) {
	event := new(SynapseCCTPCircleRequestSent)
	if err := _SynapseCCTP.contract.UnpackLog(event, "CircleRequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPEventsMetaData contains all meta data concerning the SynapseCCTPEvents contract.
var SynapseCCTPEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"kappa\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestSent\",\"type\":\"event\"}]",
}

// SynapseCCTPEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseCCTPEventsMetaData.ABI instead.
var SynapseCCTPEventsABI = SynapseCCTPEventsMetaData.ABI

// SynapseCCTPEvents is an auto generated Go binding around an Ethereum contract.
type SynapseCCTPEvents struct {
	SynapseCCTPEventsCaller     // Read-only binding to the contract
	SynapseCCTPEventsTransactor // Write-only binding to the contract
	SynapseCCTPEventsFilterer   // Log filterer for contract events
}

// SynapseCCTPEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseCCTPEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseCCTPEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseCCTPEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseCCTPEventsSession struct {
	Contract     *SynapseCCTPEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SynapseCCTPEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseCCTPEventsCallerSession struct {
	Contract *SynapseCCTPEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SynapseCCTPEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseCCTPEventsTransactorSession struct {
	Contract     *SynapseCCTPEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SynapseCCTPEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseCCTPEventsRaw struct {
	Contract *SynapseCCTPEvents // Generic contract binding to access the raw methods on
}

// SynapseCCTPEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseCCTPEventsCallerRaw struct {
	Contract *SynapseCCTPEventsCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseCCTPEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseCCTPEventsTransactorRaw struct {
	Contract *SynapseCCTPEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseCCTPEvents creates a new instance of SynapseCCTPEvents, bound to a specific deployed contract.
func NewSynapseCCTPEvents(address common.Address, backend bind.ContractBackend) (*SynapseCCTPEvents, error) {
	contract, err := bindSynapseCCTPEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEvents{SynapseCCTPEventsCaller: SynapseCCTPEventsCaller{contract: contract}, SynapseCCTPEventsTransactor: SynapseCCTPEventsTransactor{contract: contract}, SynapseCCTPEventsFilterer: SynapseCCTPEventsFilterer{contract: contract}}, nil
}

// NewSynapseCCTPEventsCaller creates a new read-only instance of SynapseCCTPEvents, bound to a specific deployed contract.
func NewSynapseCCTPEventsCaller(address common.Address, caller bind.ContractCaller) (*SynapseCCTPEventsCaller, error) {
	contract, err := bindSynapseCCTPEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsCaller{contract: contract}, nil
}

// NewSynapseCCTPEventsTransactor creates a new write-only instance of SynapseCCTPEvents, bound to a specific deployed contract.
func NewSynapseCCTPEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseCCTPEventsTransactor, error) {
	contract, err := bindSynapseCCTPEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsTransactor{contract: contract}, nil
}

// NewSynapseCCTPEventsFilterer creates a new log filterer instance of SynapseCCTPEvents, bound to a specific deployed contract.
func NewSynapseCCTPEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseCCTPEventsFilterer, error) {
	contract, err := bindSynapseCCTPEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsFilterer{contract: contract}, nil
}

// bindSynapseCCTPEvents binds a generic wrapper to an already deployed contract.
func bindSynapseCCTPEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseCCTPEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTPEvents *SynapseCCTPEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTPEvents.Contract.SynapseCCTPEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTPEvents *SynapseCCTPEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPEvents.Contract.SynapseCCTPEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTPEvents *SynapseCCTPEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTPEvents.Contract.SynapseCCTPEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTPEvents *SynapseCCTPEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTPEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTPEvents *SynapseCCTPEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTPEvents *SynapseCCTPEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTPEvents.Contract.contract.Transact(opts, method, params...)
}

// SynapseCCTPEventsCircleRequestFulfilledIterator is returned from FilterCircleRequestFulfilled and is used to iterate over the raw logs and unpacked data for CircleRequestFulfilled events raised by the SynapseCCTPEvents contract.
type SynapseCCTPEventsCircleRequestFulfilledIterator struct {
	Event *SynapseCCTPEventsCircleRequestFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SynapseCCTPEventsCircleRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPEventsCircleRequestFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SynapseCCTPEventsCircleRequestFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SynapseCCTPEventsCircleRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPEventsCircleRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPEventsCircleRequestFulfilled represents a CircleRequestFulfilled event raised by the SynapseCCTPEvents contract.
type SynapseCCTPEventsCircleRequestFulfilled struct {
	Recipient common.Address
	MintToken common.Address
	Fee       *big.Int
	Token     common.Address
	Amount    *big.Int
	Kappa     [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestFulfilled is a free log retrieval operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed kappa)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) FilterCircleRequestFulfilled(opts *bind.FilterOpts, recipient []common.Address, kappa [][32]byte) (*SynapseCCTPEventsCircleRequestFulfilledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.FilterLogs(opts, "CircleRequestFulfilled", recipientRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsCircleRequestFulfilledIterator{contract: _SynapseCCTPEvents.contract, event: "CircleRequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchCircleRequestFulfilled is a free log subscription operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed kappa)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) WatchCircleRequestFulfilled(opts *bind.WatchOpts, sink chan<- *SynapseCCTPEventsCircleRequestFulfilled, recipient []common.Address, kappa [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.WatchLogs(opts, "CircleRequestFulfilled", recipientRule, kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPEventsCircleRequestFulfilled)
				if err := _SynapseCCTPEvents.contract.UnpackLog(event, "CircleRequestFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCircleRequestFulfilled is a log parse operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed kappa)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) ParseCircleRequestFulfilled(log types.Log) (*SynapseCCTPEventsCircleRequestFulfilled, error) {
	event := new(SynapseCCTPEventsCircleRequestFulfilled)
	if err := _SynapseCCTPEvents.contract.UnpackLog(event, "CircleRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPEventsCircleRequestSentIterator is returned from FilterCircleRequestSent and is used to iterate over the raw logs and unpacked data for CircleRequestSent events raised by the SynapseCCTPEvents contract.
type SynapseCCTPEventsCircleRequestSentIterator struct {
	Event *SynapseCCTPEventsCircleRequestSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SynapseCCTPEventsCircleRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPEventsCircleRequestSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SynapseCCTPEventsCircleRequestSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SynapseCCTPEventsCircleRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPEventsCircleRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPEventsCircleRequestSent represents a CircleRequestSent event raised by the SynapseCCTPEvents contract.
type SynapseCCTPEventsCircleRequestSent struct {
	ChainId          *big.Int
	Nonce            uint64
	Token            common.Address
	Amount           *big.Int
	RequestVersion   uint32
	FormattedRequest []byte
	Kappa            [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestSent is a free log retrieval operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed kappa)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) FilterCircleRequestSent(opts *bind.FilterOpts, kappa [][32]byte) (*SynapseCCTPEventsCircleRequestSentIterator, error) {

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.FilterLogs(opts, "CircleRequestSent", kappaRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsCircleRequestSentIterator{contract: _SynapseCCTPEvents.contract, event: "CircleRequestSent", logs: logs, sub: sub}, nil
}

// WatchCircleRequestSent is a free log subscription operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed kappa)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) WatchCircleRequestSent(opts *bind.WatchOpts, sink chan<- *SynapseCCTPEventsCircleRequestSent, kappa [][32]byte) (event.Subscription, error) {

	var kappaRule []interface{}
	for _, kappaItem := range kappa {
		kappaRule = append(kappaRule, kappaItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.WatchLogs(opts, "CircleRequestSent", kappaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPEventsCircleRequestSent)
				if err := _SynapseCCTPEvents.contract.UnpackLog(event, "CircleRequestSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCircleRequestSent is a log parse operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed kappa)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) ParseCircleRequestSent(log types.Log) (*SynapseCCTPEventsCircleRequestSent, error) {
	event := new(SynapseCCTPEventsCircleRequestSent)
	if err := _SynapseCCTPEvents.contract.UnpackLog(event, "CircleRequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203dcc62b08e0ccaeb776e315e0258adfdd7dbc3da007caeda048df5bbb8699a9864736f6c63430008110033",
}

// TypeCastsABI is the input ABI used to generate the binding from.
// Deprecated: Use TypeCastsMetaData.ABI instead.
var TypeCastsABI = TypeCastsMetaData.ABI

// TypeCastsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TypeCastsMetaData.Bin instead.
var TypeCastsBin = TypeCastsMetaData.Bin

// DeployTypeCasts deploys a new Ethereum contract, binding an instance of TypeCasts to it.
func DeployTypeCasts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TypeCasts, error) {
	parsed, err := TypeCastsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TypeCastsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TypeCasts{TypeCastsCaller: TypeCastsCaller{contract: contract}, TypeCastsTransactor: TypeCastsTransactor{contract: contract}, TypeCastsFilterer: TypeCastsFilterer{contract: contract}}, nil
}

// TypeCasts is an auto generated Go binding around an Ethereum contract.
type TypeCasts struct {
	TypeCastsCaller     // Read-only binding to the contract
	TypeCastsTransactor // Write-only binding to the contract
	TypeCastsFilterer   // Log filterer for contract events
}

// TypeCastsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypeCastsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeCastsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypeCastsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeCastsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypeCastsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeCastsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypeCastsSession struct {
	Contract     *TypeCasts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypeCastsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypeCastsCallerSession struct {
	Contract *TypeCastsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TypeCastsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypeCastsTransactorSession struct {
	Contract     *TypeCastsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TypeCastsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypeCastsRaw struct {
	Contract *TypeCasts // Generic contract binding to access the raw methods on
}

// TypeCastsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypeCastsCallerRaw struct {
	Contract *TypeCastsCaller // Generic read-only contract binding to access the raw methods on
}

// TypeCastsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypeCastsTransactorRaw struct {
	Contract *TypeCastsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTypeCasts creates a new instance of TypeCasts, bound to a specific deployed contract.
func NewTypeCasts(address common.Address, backend bind.ContractBackend) (*TypeCasts, error) {
	contract, err := bindTypeCasts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TypeCasts{TypeCastsCaller: TypeCastsCaller{contract: contract}, TypeCastsTransactor: TypeCastsTransactor{contract: contract}, TypeCastsFilterer: TypeCastsFilterer{contract: contract}}, nil
}

// NewTypeCastsCaller creates a new read-only instance of TypeCasts, bound to a specific deployed contract.
func NewTypeCastsCaller(address common.Address, caller bind.ContractCaller) (*TypeCastsCaller, error) {
	contract, err := bindTypeCasts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypeCastsCaller{contract: contract}, nil
}

// NewTypeCastsTransactor creates a new write-only instance of TypeCasts, bound to a specific deployed contract.
func NewTypeCastsTransactor(address common.Address, transactor bind.ContractTransactor) (*TypeCastsTransactor, error) {
	contract, err := bindTypeCasts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypeCastsTransactor{contract: contract}, nil
}

// NewTypeCastsFilterer creates a new log filterer instance of TypeCasts, bound to a specific deployed contract.
func NewTypeCastsFilterer(address common.Address, filterer bind.ContractFilterer) (*TypeCastsFilterer, error) {
	contract, err := bindTypeCasts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypeCastsFilterer{contract: contract}, nil
}

// bindTypeCasts binds a generic wrapper to an already deployed contract.
func bindTypeCasts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypeCastsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypeCasts *TypeCastsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypeCasts.Contract.TypeCastsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypeCasts *TypeCastsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypeCasts.Contract.TypeCastsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypeCasts *TypeCastsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypeCasts.Contract.TypeCastsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypeCasts *TypeCastsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypeCasts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypeCasts *TypeCastsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypeCasts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypeCasts *TypeCastsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypeCasts.Contract.contract.Transact(opts, method, params...)
}
