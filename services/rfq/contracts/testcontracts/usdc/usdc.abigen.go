// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc

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

// USDCMetaData contains all meta data concerning the USDC contract.
var USDCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"40c10f19": "mint(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610985806100206000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610276578063a9059cbb1461027e578063dd62ed3e146102b7576100be565b806340c10f191461020857806370a0823114610243576100be565b806318160ddd116100a757806318160ddd1461018d57806323b872dd146101a7578063313ce567146101ea576100be565b806306fdde03146100c3578063095ea7b314610140575b600080fd5b6100cb6102f2565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101055781810151838201526020016100ed565b50505050905090810190601f1680156101325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101796004803603604081101561015657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020013561032b565b604080519115158252519081900360200190f35b610195610343565b60408051918252519081900360200190f35b610179600480360360608110156101bd57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101359091169060400135610349565b6101f261036d565b6040805160ff9092168252519081900360200190f35b6102416004803603604081101561021e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610372565b005b6101956004803603602081101561025957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610380565b6100cb6103a8565b6101796004803603604081101561029457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356103e1565b610195600480360360408110156102cd57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166103ef565b6040518060400160405280600881526020017f55534420436f696e00000000000000000000000000000000000000000000000081525081565b600033610339818585610427565b5060019392505050565b60005490565b60003361035785828561059a565b61036285858561064f565b506001949350505050565b600681565b61037c8282610860565b5050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b6040518060400160405280600481526020017f555344430000000000000000000000000000000000000000000000000000000081525081565b60003361033981858561064f565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b73ffffffffffffffffffffffffffffffffffffffff83166104a957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f417070726f76652066726f6d207a65726f000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff821661052b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f417070726f766520746f207a65726f0000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600260209081526040808320938616835292905220548181101561063c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e73756666696369656e7420616c6c6f77616e636500000000000000000000604482015290519081900360640190fd5b6106498484848403610427565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166106d157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f5472616e736665722066726f6d207a65726f0000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff821661075357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5472616e7366657220746f207a65726f00000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff83166000908152600160205260409020548111156107e757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e6365000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff808416600081815260016020908152604080832080548790039055938616808352918490208054860190558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a3505050565b73ffffffffffffffffffffffffffffffffffffffff82166108e257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4d696e7420746f207a65726f0000000000000000000000000000000000000000604482015290519081900360640190fd5b600080548201815573ffffffffffffffffffffffffffffffffffffffff8316808252600160209081526040808420805486019055805185815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3505056fea264697066735822122035f3015cd88cbb1fe0103279c17496a3b49acf5d38f6844bbc2d7730b1d2946164736f6c63430006020033",
}

// USDCABI is the input ABI used to generate the binding from.
// Deprecated: Use USDCMetaData.ABI instead.
var USDCABI = USDCMetaData.ABI

// Deprecated: Use USDCMetaData.Sigs instead.
// USDCFuncSigs maps the 4-byte function signature to its string representation.
var USDCFuncSigs = USDCMetaData.Sigs

// USDCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use USDCMetaData.Bin instead.
var USDCBin = USDCMetaData.Bin

// DeployUSDC deploys a new Ethereum contract, binding an instance of USDC to it.
func DeployUSDC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *USDC, error) {
	parsed, err := USDCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDC{USDCCaller: USDCCaller{contract: contract}, USDCTransactor: USDCTransactor{contract: contract}, USDCFilterer: USDCFilterer{contract: contract}}, nil
}

// USDC is an auto generated Go binding around an Ethereum contract.
type USDC struct {
	USDCCaller     // Read-only binding to the contract
	USDCTransactor // Write-only binding to the contract
	USDCFilterer   // Log filterer for contract events
}

// USDCCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDCSession struct {
	Contract     *USDC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDCCallerSession struct {
	Contract *USDCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// USDCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDCTransactorSession struct {
	Contract     *USDCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDCRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDCRaw struct {
	Contract *USDC // Generic contract binding to access the raw methods on
}

// USDCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDCCallerRaw struct {
	Contract *USDCCaller // Generic read-only contract binding to access the raw methods on
}

// USDCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDCTransactorRaw struct {
	Contract *USDCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDC creates a new instance of USDC, bound to a specific deployed contract.
func NewUSDC(address common.Address, backend bind.ContractBackend) (*USDC, error) {
	contract, err := bindUSDC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDC{USDCCaller: USDCCaller{contract: contract}, USDCTransactor: USDCTransactor{contract: contract}, USDCFilterer: USDCFilterer{contract: contract}}, nil
}

// NewUSDCCaller creates a new read-only instance of USDC, bound to a specific deployed contract.
func NewUSDCCaller(address common.Address, caller bind.ContractCaller) (*USDCCaller, error) {
	contract, err := bindUSDC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCCaller{contract: contract}, nil
}

// NewUSDCTransactor creates a new write-only instance of USDC, bound to a specific deployed contract.
func NewUSDCTransactor(address common.Address, transactor bind.ContractTransactor) (*USDCTransactor, error) {
	contract, err := bindUSDC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTransactor{contract: contract}, nil
}

// NewUSDCFilterer creates a new log filterer instance of USDC, bound to a specific deployed contract.
func NewUSDCFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCFilterer, error) {
	contract, err := bindUSDC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCFilterer{contract: contract}, nil
}

// bindUSDC binds a generic wrapper to an already deployed contract.
func bindUSDC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDC *USDCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDC.Contract.USDCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDC *USDCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDC.Contract.USDCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDC *USDCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDC.Contract.USDCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDC *USDCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDC *USDCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDC *USDCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDC *USDCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDC *USDCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USDC.Contract.Allowance(&_USDC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDC *USDCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USDC.Contract.Allowance(&_USDC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDC *USDCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDC *USDCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USDC.Contract.BalanceOf(&_USDC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDC *USDCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USDC.Contract.BalanceOf(&_USDC.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDC *USDCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USDC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDC *USDCSession) Decimals() (uint8, error) {
	return _USDC.Contract.Decimals(&_USDC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDC *USDCCallerSession) Decimals() (uint8, error) {
	return _USDC.Contract.Decimals(&_USDC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDC *USDCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDC *USDCSession) Name() (string, error) {
	return _USDC.Contract.Name(&_USDC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDC *USDCCallerSession) Name() (string, error) {
	return _USDC.Contract.Name(&_USDC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDC *USDCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDC *USDCSession) Symbol() (string, error) {
	return _USDC.Contract.Symbol(&_USDC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDC *USDCCallerSession) Symbol() (string, error) {
	return _USDC.Contract.Symbol(&_USDC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDC *USDCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDC *USDCSession) TotalSupply() (*big.Int, error) {
	return _USDC.Contract.TotalSupply(&_USDC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDC *USDCCallerSession) TotalSupply() (*big.Int, error) {
	return _USDC.Contract.TotalSupply(&_USDC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDC *USDCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDC *USDCSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.Contract.Approve(&_USDC.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDC *USDCTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.Contract.Approve(&_USDC.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_USDC *USDCTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_USDC *USDCSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.Contract.Mint(&_USDC.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_USDC *USDCTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.Contract.Mint(&_USDC.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_USDC *USDCTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_USDC *USDCSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.Contract.Transfer(&_USDC.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_USDC *USDCTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.Contract.Transfer(&_USDC.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_USDC *USDCTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_USDC *USDCSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.Contract.TransferFrom(&_USDC.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_USDC *USDCTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDC.Contract.TransferFrom(&_USDC.TransactOpts, from, to, amount)
}

// USDCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the USDC contract.
type USDCApprovalIterator struct {
	Event *USDCApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDCApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCApproval represents a Approval event raised by the USDC contract.
type USDCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDC *USDCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*USDCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &USDCApprovalIterator{contract: _USDC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDC *USDCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *USDCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCApproval)
				if err := _USDC.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_USDC *USDCFilterer) ParseApproval(log types.Log) (*USDCApproval, error) {
	event := new(USDCApproval)
	if err := _USDC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the USDC contract.
type USDCTransferIterator struct {
	Event *USDCTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDCTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDCTransfer represents a Transfer event raised by the USDC contract.
type USDCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDC *USDCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTransferIterator{contract: _USDC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDC *USDCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *USDCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDCTransfer)
				if err := _USDC.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_USDC *USDCFilterer) ParseTransfer(log types.Log) (*USDCTransfer, error) {
	event := new(USDCTransfer)
	if err := _USDC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
