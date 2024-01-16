// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dfktear

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

// TearBridgeMessageFormat is an auto generated low-level Go binding around an user-defined struct.
type TearBridgeMessageFormat struct {
	DstUser       common.Address
	DstTearAmount *big.Int
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

// IInventoryItemMetaData contains all meta data concerning the IInventoryItem contract.
var IInventoryItemMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"79cc6790": "burnFrom(address,uint256)",
		"40c10f19": "mint(address,uint256)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IInventoryItemABI is the input ABI used to generate the binding from.
// Deprecated: Use IInventoryItemMetaData.ABI instead.
var IInventoryItemABI = IInventoryItemMetaData.ABI

// Deprecated: Use IInventoryItemMetaData.Sigs instead.
// IInventoryItemFuncSigs maps the 4-byte function signature to its string representation.
var IInventoryItemFuncSigs = IInventoryItemMetaData.Sigs

// IInventoryItem is an auto generated Go binding around an Ethereum contract.
type IInventoryItem struct {
	IInventoryItemCaller     // Read-only binding to the contract
	IInventoryItemTransactor // Write-only binding to the contract
	IInventoryItemFilterer   // Log filterer for contract events
}

// IInventoryItemCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInventoryItemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInventoryItemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInventoryItemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInventoryItemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInventoryItemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInventoryItemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInventoryItemSession struct {
	Contract     *IInventoryItem   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInventoryItemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInventoryItemCallerSession struct {
	Contract *IInventoryItemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IInventoryItemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInventoryItemTransactorSession struct {
	Contract     *IInventoryItemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IInventoryItemRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInventoryItemRaw struct {
	Contract *IInventoryItem // Generic contract binding to access the raw methods on
}

// IInventoryItemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInventoryItemCallerRaw struct {
	Contract *IInventoryItemCaller // Generic read-only contract binding to access the raw methods on
}

// IInventoryItemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInventoryItemTransactorRaw struct {
	Contract *IInventoryItemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInventoryItem creates a new instance of IInventoryItem, bound to a specific deployed contract.
func NewIInventoryItem(address common.Address, backend bind.ContractBackend) (*IInventoryItem, error) {
	contract, err := bindIInventoryItem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInventoryItem{IInventoryItemCaller: IInventoryItemCaller{contract: contract}, IInventoryItemTransactor: IInventoryItemTransactor{contract: contract}, IInventoryItemFilterer: IInventoryItemFilterer{contract: contract}}, nil
}

// NewIInventoryItemCaller creates a new read-only instance of IInventoryItem, bound to a specific deployed contract.
func NewIInventoryItemCaller(address common.Address, caller bind.ContractCaller) (*IInventoryItemCaller, error) {
	contract, err := bindIInventoryItem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInventoryItemCaller{contract: contract}, nil
}

// NewIInventoryItemTransactor creates a new write-only instance of IInventoryItem, bound to a specific deployed contract.
func NewIInventoryItemTransactor(address common.Address, transactor bind.ContractTransactor) (*IInventoryItemTransactor, error) {
	contract, err := bindIInventoryItem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInventoryItemTransactor{contract: contract}, nil
}

// NewIInventoryItemFilterer creates a new log filterer instance of IInventoryItem, bound to a specific deployed contract.
func NewIInventoryItemFilterer(address common.Address, filterer bind.ContractFilterer) (*IInventoryItemFilterer, error) {
	contract, err := bindIInventoryItem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInventoryItemFilterer{contract: contract}, nil
}

// bindIInventoryItem binds a generic wrapper to an already deployed contract.
func bindIInventoryItem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInventoryItemMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInventoryItem *IInventoryItemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInventoryItem.Contract.IInventoryItemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInventoryItem *IInventoryItemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInventoryItem.Contract.IInventoryItemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInventoryItem *IInventoryItemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInventoryItem.Contract.IInventoryItemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInventoryItem *IInventoryItemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInventoryItem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInventoryItem *IInventoryItemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInventoryItem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInventoryItem *IInventoryItemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInventoryItem.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IInventoryItem *IInventoryItemCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInventoryItem.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IInventoryItem *IInventoryItemSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IInventoryItem.Contract.Allowance(&_IInventoryItem.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IInventoryItem *IInventoryItemCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IInventoryItem.Contract.Allowance(&_IInventoryItem.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IInventoryItem *IInventoryItemCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInventoryItem.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IInventoryItem *IInventoryItemSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IInventoryItem.Contract.BalanceOf(&_IInventoryItem.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IInventoryItem *IInventoryItemCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IInventoryItem.Contract.BalanceOf(&_IInventoryItem.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IInventoryItem *IInventoryItemCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IInventoryItem.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IInventoryItem *IInventoryItemSession) TotalSupply() (*big.Int, error) {
	return _IInventoryItem.Contract.TotalSupply(&_IInventoryItem.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IInventoryItem *IInventoryItemCallerSession) TotalSupply() (*big.Int, error) {
	return _IInventoryItem.Contract.TotalSupply(&_IInventoryItem.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.Approve(&_IInventoryItem.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.Approve(&_IInventoryItem.TransactOpts, spender, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_IInventoryItem *IInventoryItemTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.contract.Transact(opts, "burnFrom", from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_IInventoryItem *IInventoryItemSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.BurnFrom(&_IInventoryItem.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (_IInventoryItem *IInventoryItemTransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.BurnFrom(&_IInventoryItem.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IInventoryItem *IInventoryItemTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IInventoryItem *IInventoryItemSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.Mint(&_IInventoryItem.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IInventoryItem *IInventoryItemTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.Mint(&_IInventoryItem.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.Transfer(&_IInventoryItem.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.Transfer(&_IInventoryItem.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.TransferFrom(&_IInventoryItem.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IInventoryItem *IInventoryItemTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IInventoryItem.Contract.TransferFrom(&_IInventoryItem.TransactOpts, from, to, amount)
}

// IInventoryItemApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IInventoryItem contract.
type IInventoryItemApprovalIterator struct {
	Event *IInventoryItemApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IInventoryItemApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInventoryItemApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IInventoryItemApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IInventoryItemApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInventoryItemApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInventoryItemApproval represents a Approval event raised by the IInventoryItem contract.
type IInventoryItemApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IInventoryItem *IInventoryItemFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IInventoryItemApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IInventoryItem.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IInventoryItemApprovalIterator{contract: _IInventoryItem.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IInventoryItem *IInventoryItemFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IInventoryItemApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IInventoryItem.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInventoryItemApproval)
				if err := _IInventoryItem.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_IInventoryItem *IInventoryItemFilterer) ParseApproval(log types.Log) (*IInventoryItemApproval, error) {
	event := new(IInventoryItemApproval)
	if err := _IInventoryItem.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInventoryItemTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IInventoryItem contract.
type IInventoryItemTransferIterator struct {
	Event *IInventoryItemTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IInventoryItemTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInventoryItemTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IInventoryItemTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IInventoryItemTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInventoryItemTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInventoryItemTransfer represents a Transfer event raised by the IInventoryItem contract.
type IInventoryItemTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IInventoryItem *IInventoryItemFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IInventoryItemTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IInventoryItem.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IInventoryItemTransferIterator{contract: _IInventoryItem.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IInventoryItem *IInventoryItemFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IInventoryItemTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IInventoryItem.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInventoryItemTransfer)
				if err := _IInventoryItem.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_IInventoryItem *IInventoryItemFilterer) ParseTransfer(log types.Log) (*IInventoryItemTransfer, error) {
	event := new(IInventoryItemTransfer)
	if err := _IInventoryItem.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMessageBusMetaData contains all meta data concerning the IMessageBus contract.
var IMessageBusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5da6d2c4": "estimateFee(uint256,bytes)",
		"21730efc": "executeMessage(uint256,bytes,address,uint256,uint256,bytes,bytes32)",
		"ac8a4c1b": "sendMessage(bytes32,uint256,bytes,bytes)",
		"1ac3ddeb": "withdrawFee(address)",
	},
}

// IMessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageBusMetaData.ABI instead.
var IMessageBusABI = IMessageBusMetaData.ABI

// Deprecated: Use IMessageBusMetaData.Sigs instead.
// IMessageBusFuncSigs maps the 4-byte function signature to its string representation.
var IMessageBusFuncSigs = IMessageBusMetaData.Sigs

// IMessageBus is an auto generated Go binding around an Ethereum contract.
type IMessageBus struct {
	IMessageBusCaller     // Read-only binding to the contract
	IMessageBusTransactor // Write-only binding to the contract
	IMessageBusFilterer   // Log filterer for contract events
}

// IMessageBusCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageBusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageBusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageBusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageBusSession struct {
	Contract     *IMessageBus      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageBusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageBusCallerSession struct {
	Contract *IMessageBusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IMessageBusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageBusTransactorSession struct {
	Contract     *IMessageBusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMessageBusRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageBusRaw struct {
	Contract *IMessageBus // Generic contract binding to access the raw methods on
}

// IMessageBusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageBusCallerRaw struct {
	Contract *IMessageBusCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageBusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageBusTransactorRaw struct {
	Contract *IMessageBusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageBus creates a new instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBus(address common.Address, backend bind.ContractBackend) (*IMessageBus, error) {
	contract, err := bindIMessageBus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageBus{IMessageBusCaller: IMessageBusCaller{contract: contract}, IMessageBusTransactor: IMessageBusTransactor{contract: contract}, IMessageBusFilterer: IMessageBusFilterer{contract: contract}}, nil
}

// NewIMessageBusCaller creates a new read-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusCaller(address common.Address, caller bind.ContractCaller) (*IMessageBusCaller, error) {
	contract, err := bindIMessageBus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusCaller{contract: contract}, nil
}

// NewIMessageBusTransactor creates a new write-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageBusTransactor, error) {
	contract, err := bindIMessageBus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusTransactor{contract: contract}, nil
}

// NewIMessageBusFilterer creates a new log filterer instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageBusFilterer, error) {
	contract, err := bindIMessageBus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageBusFilterer{contract: contract}, nil
}

// bindIMessageBus binds a generic wrapper to an already deployed contract.
func bindIMessageBus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageBusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.IMessageBusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transact(opts, method, params...)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IMessageBus *IMessageBusTransactor) EstimateFee(opts *bind.TransactOpts, _dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "estimateFee", _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IMessageBus *IMessageBusSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.EstimateFee(&_IMessageBus.TransactOpts, _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IMessageBus *IMessageBusTransactorSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.EstimateFee(&_IMessageBus.TransactOpts, _dstChainId, _options)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x21730efc.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress []byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessage", _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x21730efc.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress []byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x21730efc.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress []byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// SendMessage is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_IMessageBus *IMessageBusTransactor) SendMessage(opts *bind.TransactOpts, _receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message, _options)
}

// SendMessage is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_IMessageBus *IMessageBusSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// SendMessage is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1ac3ddeb.
//
// Solidity: function withdrawFee(address _account) returns()
func (_IMessageBus *IMessageBusTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "withdrawFee", _account)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1ac3ddeb.
//
// Solidity: function withdrawFee(address _account) returns()
func (_IMessageBus *IMessageBusSession) WithdrawFee(_account common.Address) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1ac3ddeb.
//
// Solidity: function withdrawFee(address _account) returns()
func (_IMessageBus *IMessageBusTransactorSession) WithdrawFee(_account common.Address) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account)
}

// ISynMessagingReceiverMetaData contains all meta data concerning the ISynMessagingReceiver contract.
var ISynMessagingReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a6060871": "executeMessage(bytes32,uint256,bytes,address)",
	},
}

// ISynMessagingReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynMessagingReceiverMetaData.ABI instead.
var ISynMessagingReceiverABI = ISynMessagingReceiverMetaData.ABI

// Deprecated: Use ISynMessagingReceiverMetaData.Sigs instead.
// ISynMessagingReceiverFuncSigs maps the 4-byte function signature to its string representation.
var ISynMessagingReceiverFuncSigs = ISynMessagingReceiverMetaData.Sigs

// ISynMessagingReceiver is an auto generated Go binding around an Ethereum contract.
type ISynMessagingReceiver struct {
	ISynMessagingReceiverCaller     // Read-only binding to the contract
	ISynMessagingReceiverTransactor // Write-only binding to the contract
	ISynMessagingReceiverFilterer   // Log filterer for contract events
}

// ISynMessagingReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynMessagingReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynMessagingReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynMessagingReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynMessagingReceiverSession struct {
	Contract     *ISynMessagingReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ISynMessagingReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynMessagingReceiverCallerSession struct {
	Contract *ISynMessagingReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ISynMessagingReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynMessagingReceiverTransactorSession struct {
	Contract     *ISynMessagingReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ISynMessagingReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynMessagingReceiverRaw struct {
	Contract *ISynMessagingReceiver // Generic contract binding to access the raw methods on
}

// ISynMessagingReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynMessagingReceiverCallerRaw struct {
	Contract *ISynMessagingReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ISynMessagingReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynMessagingReceiverTransactorRaw struct {
	Contract *ISynMessagingReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynMessagingReceiver creates a new instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiver(address common.Address, backend bind.ContractBackend) (*ISynMessagingReceiver, error) {
	contract, err := bindISynMessagingReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiver{ISynMessagingReceiverCaller: ISynMessagingReceiverCaller{contract: contract}, ISynMessagingReceiverTransactor: ISynMessagingReceiverTransactor{contract: contract}, ISynMessagingReceiverFilterer: ISynMessagingReceiverFilterer{contract: contract}}, nil
}

// NewISynMessagingReceiverCaller creates a new read-only instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverCaller(address common.Address, caller bind.ContractCaller) (*ISynMessagingReceiverCaller, error) {
	contract, err := bindISynMessagingReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverCaller{contract: contract}, nil
}

// NewISynMessagingReceiverTransactor creates a new write-only instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynMessagingReceiverTransactor, error) {
	contract, err := bindISynMessagingReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverTransactor{contract: contract}, nil
}

// NewISynMessagingReceiverFilterer creates a new log filterer instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynMessagingReceiverFilterer, error) {
	contract, err := bindISynMessagingReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverFilterer{contract: contract}, nil
}

// bindISynMessagingReceiver binds a generic wrapper to an already deployed contract.
func bindISynMessagingReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISynMessagingReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynMessagingReceiver *ISynMessagingReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynMessagingReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.contract.Transact(opts, "executeMessage", _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ExecuteMessage(&_ISynMessagingReceiver.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ExecuteMessage(&_ISynMessagingReceiver.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynMessagingReceiverMetaData contains all meta data concerning the SynMessagingReceiver contract.
var SynMessagingReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getTrustedRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"trustedRemote\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a6060871": "executeMessage(bytes32,uint256,bytes,address)",
		"84a12b0f": "getTrustedRemote(uint256)",
		"a1a227fa": "messageBus()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"547cad12": "setMessageBus(address)",
		"bd3583ae": "setTrustedRemote(uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
	},
}

// SynMessagingReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use SynMessagingReceiverMetaData.ABI instead.
var SynMessagingReceiverABI = SynMessagingReceiverMetaData.ABI

// Deprecated: Use SynMessagingReceiverMetaData.Sigs instead.
// SynMessagingReceiverFuncSigs maps the 4-byte function signature to its string representation.
var SynMessagingReceiverFuncSigs = SynMessagingReceiverMetaData.Sigs

// SynMessagingReceiver is an auto generated Go binding around an Ethereum contract.
type SynMessagingReceiver struct {
	SynMessagingReceiverCaller     // Read-only binding to the contract
	SynMessagingReceiverTransactor // Write-only binding to the contract
	SynMessagingReceiverFilterer   // Log filterer for contract events
}

// SynMessagingReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynMessagingReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynMessagingReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynMessagingReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynMessagingReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynMessagingReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynMessagingReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynMessagingReceiverSession struct {
	Contract     *SynMessagingReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SynMessagingReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynMessagingReceiverCallerSession struct {
	Contract *SynMessagingReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SynMessagingReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynMessagingReceiverTransactorSession struct {
	Contract     *SynMessagingReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SynMessagingReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynMessagingReceiverRaw struct {
	Contract *SynMessagingReceiver // Generic contract binding to access the raw methods on
}

// SynMessagingReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynMessagingReceiverCallerRaw struct {
	Contract *SynMessagingReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// SynMessagingReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynMessagingReceiverTransactorRaw struct {
	Contract *SynMessagingReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynMessagingReceiver creates a new instance of SynMessagingReceiver, bound to a specific deployed contract.
func NewSynMessagingReceiver(address common.Address, backend bind.ContractBackend) (*SynMessagingReceiver, error) {
	contract, err := bindSynMessagingReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiver{SynMessagingReceiverCaller: SynMessagingReceiverCaller{contract: contract}, SynMessagingReceiverTransactor: SynMessagingReceiverTransactor{contract: contract}, SynMessagingReceiverFilterer: SynMessagingReceiverFilterer{contract: contract}}, nil
}

// NewSynMessagingReceiverCaller creates a new read-only instance of SynMessagingReceiver, bound to a specific deployed contract.
func NewSynMessagingReceiverCaller(address common.Address, caller bind.ContractCaller) (*SynMessagingReceiverCaller, error) {
	contract, err := bindSynMessagingReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverCaller{contract: contract}, nil
}

// NewSynMessagingReceiverTransactor creates a new write-only instance of SynMessagingReceiver, bound to a specific deployed contract.
func NewSynMessagingReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*SynMessagingReceiverTransactor, error) {
	contract, err := bindSynMessagingReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverTransactor{contract: contract}, nil
}

// NewSynMessagingReceiverFilterer creates a new log filterer instance of SynMessagingReceiver, bound to a specific deployed contract.
func NewSynMessagingReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*SynMessagingReceiverFilterer, error) {
	contract, err := bindSynMessagingReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverFilterer{contract: contract}, nil
}

// bindSynMessagingReceiver binds a generic wrapper to an already deployed contract.
func bindSynMessagingReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynMessagingReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynMessagingReceiver *SynMessagingReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynMessagingReceiver.Contract.SynMessagingReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynMessagingReceiver *SynMessagingReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.SynMessagingReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynMessagingReceiver *SynMessagingReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.SynMessagingReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynMessagingReceiver *SynMessagingReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynMessagingReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynMessagingReceiver *SynMessagingReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynMessagingReceiver *SynMessagingReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.contract.Transact(opts, method, params...)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_SynMessagingReceiver *SynMessagingReceiverCaller) GetTrustedRemote(opts *bind.CallOpts, _chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SynMessagingReceiver.contract.Call(opts, &out, "getTrustedRemote", _chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_SynMessagingReceiver *SynMessagingReceiverSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _SynMessagingReceiver.Contract.GetTrustedRemote(&_SynMessagingReceiver.CallOpts, _chainId)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_SynMessagingReceiver *SynMessagingReceiverCallerSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _SynMessagingReceiver.Contract.GetTrustedRemote(&_SynMessagingReceiver.CallOpts, _chainId)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_SynMessagingReceiver *SynMessagingReceiverCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynMessagingReceiver.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_SynMessagingReceiver *SynMessagingReceiverSession) MessageBus() (common.Address, error) {
	return _SynMessagingReceiver.Contract.MessageBus(&_SynMessagingReceiver.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_SynMessagingReceiver *SynMessagingReceiverCallerSession) MessageBus() (common.Address, error) {
	return _SynMessagingReceiver.Contract.MessageBus(&_SynMessagingReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynMessagingReceiver *SynMessagingReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynMessagingReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynMessagingReceiver *SynMessagingReceiverSession) Owner() (common.Address, error) {
	return _SynMessagingReceiver.Contract.Owner(&_SynMessagingReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynMessagingReceiver *SynMessagingReceiverCallerSession) Owner() (common.Address, error) {
	return _SynMessagingReceiver.Contract.Owner(&_SynMessagingReceiver.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.contract.Transact(opts, "executeMessage", _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_SynMessagingReceiver *SynMessagingReceiverSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.ExecuteMessage(&_SynMessagingReceiver.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactorSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.ExecuteMessage(&_SynMessagingReceiver.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynMessagingReceiver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynMessagingReceiver *SynMessagingReceiverSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.RenounceOwnership(&_SynMessagingReceiver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.RenounceOwnership(&_SynMessagingReceiver.TransactOpts)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_SynMessagingReceiver *SynMessagingReceiverSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.SetMessageBus(&_SynMessagingReceiver.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.SetMessageBus(&_SynMessagingReceiver.TransactOpts, _messageBus)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _SynMessagingReceiver.contract.Transact(opts, "setTrustedRemote", _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_SynMessagingReceiver *SynMessagingReceiverSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.SetTrustedRemote(&_SynMessagingReceiver.TransactOpts, _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactorSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.SetTrustedRemote(&_SynMessagingReceiver.TransactOpts, _srcChainId, _srcAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynMessagingReceiver *SynMessagingReceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.TransferOwnership(&_SynMessagingReceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynMessagingReceiver *SynMessagingReceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiver.Contract.TransferOwnership(&_SynMessagingReceiver.TransactOpts, newOwner)
}

// SynMessagingReceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynMessagingReceiver contract.
type SynMessagingReceiverOwnershipTransferredIterator struct {
	Event *SynMessagingReceiverOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SynMessagingReceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynMessagingReceiverOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SynMessagingReceiverOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SynMessagingReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynMessagingReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynMessagingReceiverOwnershipTransferred represents a OwnershipTransferred event raised by the SynMessagingReceiver contract.
type SynMessagingReceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynMessagingReceiver *SynMessagingReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynMessagingReceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynMessagingReceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverOwnershipTransferredIterator{contract: _SynMessagingReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynMessagingReceiver *SynMessagingReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynMessagingReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynMessagingReceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynMessagingReceiverOwnershipTransferred)
				if err := _SynMessagingReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynMessagingReceiver *SynMessagingReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*SynMessagingReceiverOwnershipTransferred, error) {
	event := new(SynMessagingReceiverOwnershipTransferred)
	if err := _SynMessagingReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynMessagingReceiverSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the SynMessagingReceiver contract.
type SynMessagingReceiverSetTrustedRemoteIterator struct {
	Event *SynMessagingReceiverSetTrustedRemote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SynMessagingReceiverSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynMessagingReceiverSetTrustedRemote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SynMessagingReceiverSetTrustedRemote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SynMessagingReceiverSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynMessagingReceiverSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynMessagingReceiverSetTrustedRemote represents a SetTrustedRemote event raised by the SynMessagingReceiver contract.
type SynMessagingReceiverSetTrustedRemote struct {
	SrcChainId *big.Int
	SrcAddress [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_SynMessagingReceiver *SynMessagingReceiverFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*SynMessagingReceiverSetTrustedRemoteIterator, error) {

	logs, sub, err := _SynMessagingReceiver.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverSetTrustedRemoteIterator{contract: _SynMessagingReceiver.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_SynMessagingReceiver *SynMessagingReceiverFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *SynMessagingReceiverSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _SynMessagingReceiver.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynMessagingReceiverSetTrustedRemote)
				if err := _SynMessagingReceiver.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemote is a log parse operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_SynMessagingReceiver *SynMessagingReceiverFilterer) ParseSetTrustedRemote(log types.Log) (*SynMessagingReceiverSetTrustedRemote, error) {
	event := new(SynMessagingReceiverSetTrustedRemote)
	if err := _SynMessagingReceiver.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TearBridgeMetaData contains all meta data concerning the TearBridge contract.
var TearBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gaiaTear\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arrivalChainId\",\"type\":\"uint256\"}],\"name\":\"GaiaArrived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arrivalChainId\",\"type\":\"uint256\"}],\"name\":\"GaiaSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"decodeMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"dstUser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstTearAmount\",\"type\":\"uint256\"}],\"internalType\":\"structTearBridge.MessageFormat\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gaiaTears\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getTrustedRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"trustedRemote\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tearsAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"}],\"name\":\"sendTear\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_msgGasLimit\",\"type\":\"uint256\"}],\"name\":\"setMsgGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"634d45b2": "decodeMessage(bytes)",
		"a6060871": "executeMessage(bytes32,uint256,bytes,address)",
		"acac4bdd": "gaiaTears()",
		"84a12b0f": "getTrustedRemote(uint256)",
		"a1a227fa": "messageBus()",
		"c0e07f28": "msgGasLimit()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"82731903": "sendTear(uint256,uint256)",
		"547cad12": "setMessageBus(address)",
		"f9ecc6f5": "setMsgGasLimit(uint256)",
		"bd3583ae": "setTrustedRemote(uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x60a060405234801561001057600080fd5b5060405161116e38038061116e83398101604081905261002f916100c9565b6100383361005d565b600180546001600160a01b0319166001600160a01b03938416179055166080526100fc565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146100c457600080fd5b919050565b600080604083850312156100dc57600080fd5b6100e5836100ad565b91506100f3602084016100ad565b90509250929050565b6080516110496101256000396000818161025d015281816104ab0152610c5b01526110496000f3fe6080604052600436106100d25760003560e01c8063a1a227fa1161007f578063bd3583ae11610059578063bd3583ae1461027f578063c0e07f281461029f578063f2fde38b146102b5578063f9ecc6f5146102d557600080fd5b8063a1a227fa146101fe578063a60608711461022b578063acac4bdd1461024b57600080fd5b806382731903116100b0578063827319031461016457806384a12b0f146101775780638da5cb5b146101b257600080fd5b8063547cad12146100d7578063634d45b2146100f9578063715018a61461014f575b600080fd5b3480156100e357600080fd5b506100f76100f2366004610d23565b6102f5565b005b34801561010557600080fd5b50610119610114366004610d6f565b6103c2565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260209283015192810192909252015b60405180910390f35b34801561015b57600080fd5b506100f76103e5565b6100f7610172366004610e3e565b610472565b34801561018357600080fd5b506101a4610192366004610e60565b60009081526002602052604090205490565b604051908152602001610146565b3480156101be57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610146565b34801561020a57600080fd5b506001546101d99073ffffffffffffffffffffffffffffffffffffffff1681565b34801561023757600080fd5b506100f7610246366004610e79565b6105ec565b34801561025757600080fd5b506101d97f000000000000000000000000000000000000000000000000000000000000000081565b34801561028b57600080fd5b506100f761029a366004610e3e565b61072e565b3480156102ab57600080fd5b506101a460035481565b3480156102c157600080fd5b506100f76102d0366004610d23565b6107fe565b3480156102e157600080fd5b506100f76102f0366004610e60565b61092e565b60005473ffffffffffffffffffffffffffffffffffffffff16331461037b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60408051808201909152600080825260208201526103df826109b4565b92915050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610372565b61047060006109e5565b565b6040517f79cc679000000000000000000000000000000000000000000000000000000000815233600482015260248101839052829082907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b15801561050457600080fd5b505af1158015610518573d6000803e3d6000fd5b50505060008281526002602090815260409182902054825180840184523380825290830187815284519384019190915251828401528251808303840181526060830184526003547e0100000000000000000000000000000000000000000000000000000000000060808501526082808501919091528451808503909101815260a290930190935292506105ad83858484610a5a565b60405185815233907fe82273e05845454dcf88823968e5c722028bc4cb17ed03bdc06eaa32cc58ee66906020015b60405180910390a250505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461066d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610372565b60008481526002602052604090205485146106e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f496e76616c696420736f757263652073656e64696e67206170700000000000006044820152606401610372565b610727858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250610bf6915050565b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146107af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610372565b60008281526002602090815260409182902083905581518481529081018390527f642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03910160405180910390a15050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461087f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610372565b73ffffffffffffffffffffffffffffffffffffffff8116610922576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610372565b61092b816109e5565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146109af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610372565b600355565b60408051808201909152600080825260208201526000828060200190518101906109de9190610f14565b9392505050565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008381526002602052604090205480610ad0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4e6f2072656d6f7465206170702073657420666f722064737420636861696e006044820152606401610372565b848114610b5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f5265636569766572206973206e6f7420696e20747275737465642072656d6f7460448201527f65206170707300000000000000000000000000000000000000000000000000006064820152608401610372565b6001546040517fac8a4c1b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063ac8a4c1b903490610bbd908990899089908990600401610fd7565b6000604051808303818588803b158015610bd657600080fd5b505af1158015610bea573d6000803e3d6000fd5b50505050505050505050565b6000610c01836109b4565b805160208201516040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808416600483015260248201839052939450919290917f0000000000000000000000000000000000000000000000000000000000000000909116906340c10f1990604401600060405180830381600087803b158015610ca157600080fd5b505af1158015610cb5573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff167f07b815cd29685803e7213231371fa19ce2e23221109bf847d949305e6b7464a4826040516105db91815260200190565b73ffffffffffffffffffffffffffffffffffffffff8116811461092b57600080fd5b600060208284031215610d3557600080fd5b81356109de81610d01565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060208284031215610d8157600080fd5b813567ffffffffffffffff80821115610d9957600080fd5b818401915084601f830112610dad57600080fd5b813581811115610dbf57610dbf610d40565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610e0557610e05610d40565b81604052828152876020848701011115610e1e57600080fd5b826020860160208301376000928101602001929092525095945050505050565b60008060408385031215610e5157600080fd5b50508035926020909101359150565b600060208284031215610e7257600080fd5b5035919050565b600080600080600060808688031215610e9157600080fd5b8535945060208601359350604086013567ffffffffffffffff80821115610eb757600080fd5b818801915088601f830112610ecb57600080fd5b813581811115610eda57600080fd5b896020828501011115610eec57600080fd5b6020830195508094505050506060860135610f0681610d01565b809150509295509295909350565b600060408284031215610f2657600080fd5b6040516040810181811067ffffffffffffffff82111715610f4957610f49610d40565b6040528251610f5781610d01565b81526020928301519281019290925250919050565b6000815180845260005b81811015610f9257602081850181015186830182015201610f76565b81811115610fa4576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b848152836020820152608060408201526000610ff66080830185610f6c565b82810360608401526110088185610f6c565b97965050505050505056fea2646970667358221220cea6277f9f212d1506db098e9253d5cf9a52097aa7864aca3d8e1cacf1659e3964736f6c634300080d0033",
}

// TearBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use TearBridgeMetaData.ABI instead.
var TearBridgeABI = TearBridgeMetaData.ABI

// Deprecated: Use TearBridgeMetaData.Sigs instead.
// TearBridgeFuncSigs maps the 4-byte function signature to its string representation.
var TearBridgeFuncSigs = TearBridgeMetaData.Sigs

// TearBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TearBridgeMetaData.Bin instead.
var TearBridgeBin = TearBridgeMetaData.Bin

// DeployTearBridge deploys a new Ethereum contract, binding an instance of TearBridge to it.
func DeployTearBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _messageBus common.Address, _gaiaTear common.Address) (common.Address, *types.Transaction, *TearBridge, error) {
	parsed, err := TearBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TearBridgeBin), backend, _messageBus, _gaiaTear)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TearBridge{TearBridgeCaller: TearBridgeCaller{contract: contract}, TearBridgeTransactor: TearBridgeTransactor{contract: contract}, TearBridgeFilterer: TearBridgeFilterer{contract: contract}}, nil
}

// TearBridge is an auto generated Go binding around an Ethereum contract.
type TearBridge struct {
	TearBridgeCaller     // Read-only binding to the contract
	TearBridgeTransactor // Write-only binding to the contract
	TearBridgeFilterer   // Log filterer for contract events
}

// TearBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TearBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TearBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TearBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TearBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TearBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TearBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TearBridgeSession struct {
	Contract     *TearBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TearBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TearBridgeCallerSession struct {
	Contract *TearBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TearBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TearBridgeTransactorSession struct {
	Contract     *TearBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TearBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TearBridgeRaw struct {
	Contract *TearBridge // Generic contract binding to access the raw methods on
}

// TearBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TearBridgeCallerRaw struct {
	Contract *TearBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// TearBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TearBridgeTransactorRaw struct {
	Contract *TearBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTearBridge creates a new instance of TearBridge, bound to a specific deployed contract.
func NewTearBridge(address common.Address, backend bind.ContractBackend) (*TearBridge, error) {
	contract, err := bindTearBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TearBridge{TearBridgeCaller: TearBridgeCaller{contract: contract}, TearBridgeTransactor: TearBridgeTransactor{contract: contract}, TearBridgeFilterer: TearBridgeFilterer{contract: contract}}, nil
}

// NewTearBridgeCaller creates a new read-only instance of TearBridge, bound to a specific deployed contract.
func NewTearBridgeCaller(address common.Address, caller bind.ContractCaller) (*TearBridgeCaller, error) {
	contract, err := bindTearBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TearBridgeCaller{contract: contract}, nil
}

// NewTearBridgeTransactor creates a new write-only instance of TearBridge, bound to a specific deployed contract.
func NewTearBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*TearBridgeTransactor, error) {
	contract, err := bindTearBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TearBridgeTransactor{contract: contract}, nil
}

// NewTearBridgeFilterer creates a new log filterer instance of TearBridge, bound to a specific deployed contract.
func NewTearBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*TearBridgeFilterer, error) {
	contract, err := bindTearBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TearBridgeFilterer{contract: contract}, nil
}

// bindTearBridge binds a generic wrapper to an already deployed contract.
func bindTearBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TearBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TearBridge *TearBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TearBridge.Contract.TearBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TearBridge *TearBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TearBridge.Contract.TearBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TearBridge *TearBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TearBridge.Contract.TearBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TearBridge *TearBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TearBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TearBridge *TearBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TearBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TearBridge *TearBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TearBridge.Contract.contract.Transact(opts, method, params...)
}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns((address,uint256))
func (_TearBridge *TearBridgeCaller) DecodeMessage(opts *bind.CallOpts, _message []byte) (TearBridgeMessageFormat, error) {
	var out []interface{}
	err := _TearBridge.contract.Call(opts, &out, "decodeMessage", _message)

	if err != nil {
		return *new(TearBridgeMessageFormat), err
	}

	out0 := *abi.ConvertType(out[0], new(TearBridgeMessageFormat)).(*TearBridgeMessageFormat)

	return out0, err

}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns((address,uint256))
func (_TearBridge *TearBridgeSession) DecodeMessage(_message []byte) (TearBridgeMessageFormat, error) {
	return _TearBridge.Contract.DecodeMessage(&_TearBridge.CallOpts, _message)
}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns((address,uint256))
func (_TearBridge *TearBridgeCallerSession) DecodeMessage(_message []byte) (TearBridgeMessageFormat, error) {
	return _TearBridge.Contract.DecodeMessage(&_TearBridge.CallOpts, _message)
}

// GaiaTears is a free data retrieval call binding the contract method 0xacac4bdd.
//
// Solidity: function gaiaTears() view returns(address)
func (_TearBridge *TearBridgeCaller) GaiaTears(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TearBridge.contract.Call(opts, &out, "gaiaTears")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaiaTears is a free data retrieval call binding the contract method 0xacac4bdd.
//
// Solidity: function gaiaTears() view returns(address)
func (_TearBridge *TearBridgeSession) GaiaTears() (common.Address, error) {
	return _TearBridge.Contract.GaiaTears(&_TearBridge.CallOpts)
}

// GaiaTears is a free data retrieval call binding the contract method 0xacac4bdd.
//
// Solidity: function gaiaTears() view returns(address)
func (_TearBridge *TearBridgeCallerSession) GaiaTears() (common.Address, error) {
	return _TearBridge.Contract.GaiaTears(&_TearBridge.CallOpts)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_TearBridge *TearBridgeCaller) GetTrustedRemote(opts *bind.CallOpts, _chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TearBridge.contract.Call(opts, &out, "getTrustedRemote", _chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_TearBridge *TearBridgeSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _TearBridge.Contract.GetTrustedRemote(&_TearBridge.CallOpts, _chainId)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_TearBridge *TearBridgeCallerSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _TearBridge.Contract.GetTrustedRemote(&_TearBridge.CallOpts, _chainId)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TearBridge *TearBridgeCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TearBridge.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TearBridge *TearBridgeSession) MessageBus() (common.Address, error) {
	return _TearBridge.Contract.MessageBus(&_TearBridge.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TearBridge *TearBridgeCallerSession) MessageBus() (common.Address, error) {
	return _TearBridge.Contract.MessageBus(&_TearBridge.CallOpts)
}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_TearBridge *TearBridgeCaller) MsgGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TearBridge.contract.Call(opts, &out, "msgGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_TearBridge *TearBridgeSession) MsgGasLimit() (*big.Int, error) {
	return _TearBridge.Contract.MsgGasLimit(&_TearBridge.CallOpts)
}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_TearBridge *TearBridgeCallerSession) MsgGasLimit() (*big.Int, error) {
	return _TearBridge.Contract.MsgGasLimit(&_TearBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TearBridge *TearBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TearBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TearBridge *TearBridgeSession) Owner() (common.Address, error) {
	return _TearBridge.Contract.Owner(&_TearBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TearBridge *TearBridgeCallerSession) Owner() (common.Address, error) {
	return _TearBridge.Contract.Owner(&_TearBridge.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_TearBridge *TearBridgeTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _TearBridge.contract.Transact(opts, "executeMessage", _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_TearBridge *TearBridgeSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _TearBridge.Contract.ExecuteMessage(&_TearBridge.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_TearBridge *TearBridgeTransactorSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _TearBridge.Contract.ExecuteMessage(&_TearBridge.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TearBridge *TearBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TearBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TearBridge *TearBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _TearBridge.Contract.RenounceOwnership(&_TearBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TearBridge *TearBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TearBridge.Contract.RenounceOwnership(&_TearBridge.TransactOpts)
}

// SendTear is a paid mutator transaction binding the contract method 0x82731903.
//
// Solidity: function sendTear(uint256 _tearsAmount, uint256 _dstChainId) payable returns()
func (_TearBridge *TearBridgeTransactor) SendTear(opts *bind.TransactOpts, _tearsAmount *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _TearBridge.contract.Transact(opts, "sendTear", _tearsAmount, _dstChainId)
}

// SendTear is a paid mutator transaction binding the contract method 0x82731903.
//
// Solidity: function sendTear(uint256 _tearsAmount, uint256 _dstChainId) payable returns()
func (_TearBridge *TearBridgeSession) SendTear(_tearsAmount *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _TearBridge.Contract.SendTear(&_TearBridge.TransactOpts, _tearsAmount, _dstChainId)
}

// SendTear is a paid mutator transaction binding the contract method 0x82731903.
//
// Solidity: function sendTear(uint256 _tearsAmount, uint256 _dstChainId) payable returns()
func (_TearBridge *TearBridgeTransactorSession) SendTear(_tearsAmount *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _TearBridge.Contract.SendTear(&_TearBridge.TransactOpts, _tearsAmount, _dstChainId)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TearBridge *TearBridgeTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _TearBridge.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TearBridge *TearBridgeSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _TearBridge.Contract.SetMessageBus(&_TearBridge.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TearBridge *TearBridgeTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _TearBridge.Contract.SetMessageBus(&_TearBridge.TransactOpts, _messageBus)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_TearBridge *TearBridgeTransactor) SetMsgGasLimit(opts *bind.TransactOpts, _msgGasLimit *big.Int) (*types.Transaction, error) {
	return _TearBridge.contract.Transact(opts, "setMsgGasLimit", _msgGasLimit)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_TearBridge *TearBridgeSession) SetMsgGasLimit(_msgGasLimit *big.Int) (*types.Transaction, error) {
	return _TearBridge.Contract.SetMsgGasLimit(&_TearBridge.TransactOpts, _msgGasLimit)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_TearBridge *TearBridgeTransactorSession) SetMsgGasLimit(_msgGasLimit *big.Int) (*types.Transaction, error) {
	return _TearBridge.Contract.SetMsgGasLimit(&_TearBridge.TransactOpts, _msgGasLimit)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_TearBridge *TearBridgeTransactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _TearBridge.contract.Transact(opts, "setTrustedRemote", _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_TearBridge *TearBridgeSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _TearBridge.Contract.SetTrustedRemote(&_TearBridge.TransactOpts, _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_TearBridge *TearBridgeTransactorSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _TearBridge.Contract.SetTrustedRemote(&_TearBridge.TransactOpts, _srcChainId, _srcAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TearBridge *TearBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TearBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TearBridge *TearBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TearBridge.Contract.TransferOwnership(&_TearBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TearBridge *TearBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TearBridge.Contract.TransferOwnership(&_TearBridge.TransactOpts, newOwner)
}

// TearBridgeGaiaArrivedIterator is returned from FilterGaiaArrived and is used to iterate over the raw logs and unpacked data for GaiaArrived events raised by the TearBridge contract.
type TearBridgeGaiaArrivedIterator struct {
	Event *TearBridgeGaiaArrived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TearBridgeGaiaArrivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TearBridgeGaiaArrived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TearBridgeGaiaArrived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TearBridgeGaiaArrivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TearBridgeGaiaArrivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TearBridgeGaiaArrived represents a GaiaArrived event raised by the TearBridge contract.
type TearBridgeGaiaArrived struct {
	DstUser        common.Address
	ArrivalChainId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGaiaArrived is a free log retrieval operation binding the contract event 0x07b815cd29685803e7213231371fa19ce2e23221109bf847d949305e6b7464a4.
//
// Solidity: event GaiaArrived(address indexed dstUser, uint256 arrivalChainId)
func (_TearBridge *TearBridgeFilterer) FilterGaiaArrived(opts *bind.FilterOpts, dstUser []common.Address) (*TearBridgeGaiaArrivedIterator, error) {

	var dstUserRule []interface{}
	for _, dstUserItem := range dstUser {
		dstUserRule = append(dstUserRule, dstUserItem)
	}

	logs, sub, err := _TearBridge.contract.FilterLogs(opts, "GaiaArrived", dstUserRule)
	if err != nil {
		return nil, err
	}
	return &TearBridgeGaiaArrivedIterator{contract: _TearBridge.contract, event: "GaiaArrived", logs: logs, sub: sub}, nil
}

// WatchGaiaArrived is a free log subscription operation binding the contract event 0x07b815cd29685803e7213231371fa19ce2e23221109bf847d949305e6b7464a4.
//
// Solidity: event GaiaArrived(address indexed dstUser, uint256 arrivalChainId)
func (_TearBridge *TearBridgeFilterer) WatchGaiaArrived(opts *bind.WatchOpts, sink chan<- *TearBridgeGaiaArrived, dstUser []common.Address) (event.Subscription, error) {

	var dstUserRule []interface{}
	for _, dstUserItem := range dstUser {
		dstUserRule = append(dstUserRule, dstUserItem)
	}

	logs, sub, err := _TearBridge.contract.WatchLogs(opts, "GaiaArrived", dstUserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TearBridgeGaiaArrived)
				if err := _TearBridge.contract.UnpackLog(event, "GaiaArrived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGaiaArrived is a log parse operation binding the contract event 0x07b815cd29685803e7213231371fa19ce2e23221109bf847d949305e6b7464a4.
//
// Solidity: event GaiaArrived(address indexed dstUser, uint256 arrivalChainId)
func (_TearBridge *TearBridgeFilterer) ParseGaiaArrived(log types.Log) (*TearBridgeGaiaArrived, error) {
	event := new(TearBridgeGaiaArrived)
	if err := _TearBridge.contract.UnpackLog(event, "GaiaArrived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TearBridgeGaiaSentIterator is returned from FilterGaiaSent and is used to iterate over the raw logs and unpacked data for GaiaSent events raised by the TearBridge contract.
type TearBridgeGaiaSentIterator struct {
	Event *TearBridgeGaiaSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TearBridgeGaiaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TearBridgeGaiaSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TearBridgeGaiaSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TearBridgeGaiaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TearBridgeGaiaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TearBridgeGaiaSent represents a GaiaSent event raised by the TearBridge contract.
type TearBridgeGaiaSent struct {
	DstUser        common.Address
	ArrivalChainId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGaiaSent is a free log retrieval operation binding the contract event 0xe82273e05845454dcf88823968e5c722028bc4cb17ed03bdc06eaa32cc58ee66.
//
// Solidity: event GaiaSent(address indexed dstUser, uint256 arrivalChainId)
func (_TearBridge *TearBridgeFilterer) FilterGaiaSent(opts *bind.FilterOpts, dstUser []common.Address) (*TearBridgeGaiaSentIterator, error) {

	var dstUserRule []interface{}
	for _, dstUserItem := range dstUser {
		dstUserRule = append(dstUserRule, dstUserItem)
	}

	logs, sub, err := _TearBridge.contract.FilterLogs(opts, "GaiaSent", dstUserRule)
	if err != nil {
		return nil, err
	}
	return &TearBridgeGaiaSentIterator{contract: _TearBridge.contract, event: "GaiaSent", logs: logs, sub: sub}, nil
}

// WatchGaiaSent is a free log subscription operation binding the contract event 0xe82273e05845454dcf88823968e5c722028bc4cb17ed03bdc06eaa32cc58ee66.
//
// Solidity: event GaiaSent(address indexed dstUser, uint256 arrivalChainId)
func (_TearBridge *TearBridgeFilterer) WatchGaiaSent(opts *bind.WatchOpts, sink chan<- *TearBridgeGaiaSent, dstUser []common.Address) (event.Subscription, error) {

	var dstUserRule []interface{}
	for _, dstUserItem := range dstUser {
		dstUserRule = append(dstUserRule, dstUserItem)
	}

	logs, sub, err := _TearBridge.contract.WatchLogs(opts, "GaiaSent", dstUserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TearBridgeGaiaSent)
				if err := _TearBridge.contract.UnpackLog(event, "GaiaSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGaiaSent is a log parse operation binding the contract event 0xe82273e05845454dcf88823968e5c722028bc4cb17ed03bdc06eaa32cc58ee66.
//
// Solidity: event GaiaSent(address indexed dstUser, uint256 arrivalChainId)
func (_TearBridge *TearBridgeFilterer) ParseGaiaSent(log types.Log) (*TearBridgeGaiaSent, error) {
	event := new(TearBridgeGaiaSent)
	if err := _TearBridge.contract.UnpackLog(event, "GaiaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TearBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TearBridge contract.
type TearBridgeOwnershipTransferredIterator struct {
	Event *TearBridgeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TearBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TearBridgeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TearBridgeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TearBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TearBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TearBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the TearBridge contract.
type TearBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TearBridge *TearBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TearBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TearBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TearBridgeOwnershipTransferredIterator{contract: _TearBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TearBridge *TearBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TearBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TearBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TearBridgeOwnershipTransferred)
				if err := _TearBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TearBridge *TearBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*TearBridgeOwnershipTransferred, error) {
	event := new(TearBridgeOwnershipTransferred)
	if err := _TearBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TearBridgeSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the TearBridge contract.
type TearBridgeSetTrustedRemoteIterator struct {
	Event *TearBridgeSetTrustedRemote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TearBridgeSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TearBridgeSetTrustedRemote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TearBridgeSetTrustedRemote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TearBridgeSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TearBridgeSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TearBridgeSetTrustedRemote represents a SetTrustedRemote event raised by the TearBridge contract.
type TearBridgeSetTrustedRemote struct {
	SrcChainId *big.Int
	SrcAddress [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_TearBridge *TearBridgeFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*TearBridgeSetTrustedRemoteIterator, error) {

	logs, sub, err := _TearBridge.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &TearBridgeSetTrustedRemoteIterator{contract: _TearBridge.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_TearBridge *TearBridgeFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *TearBridgeSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _TearBridge.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TearBridgeSetTrustedRemote)
				if err := _TearBridge.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemote is a log parse operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_TearBridge *TearBridgeFilterer) ParseSetTrustedRemote(log types.Log) (*TearBridgeSetTrustedRemote, error) {
	event := new(TearBridgeSetTrustedRemote)
	if err := _TearBridge.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
