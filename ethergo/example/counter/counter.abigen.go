// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decrementCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVitalikCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vitalikIncrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f5c5ad83": "decrementCounter()",
		"a87d942c": "getCount()",
		"9f6f1ec1": "getVitalikCount()",
		"5b34b966": "incrementCounter()",
		"6c573535": "vitalikIncrement()",
	},
	Bin: "0x608060405260008055600060015534801561001957600080fd5b506102b0806100296000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80639f6f1ec1116100505780639f6f1ec11461007e578063a87d942c14610094578063f5c5ad831461009c57600080fd5b80635b34b9661461006c5780636c57353514610076575b600080fd5b6100746100a4565b005b6100746100bd565b6001545b60405190815260200160405180910390f35b600054610082565b610074610151565b60016000808282546100b69190610163565b9091555050565b3373d8da6bf26964af9d7eed9e03e53415d37aa960451461013e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4f6e6c7920566974616c696b2063616e20636f756e7420627920313000000000604482015260640160405180910390fd5b600a600160008282546100b69190610163565b60016000808282546100b691906101d7565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561019d5761019d61024b565b827f80000000000000000000000000000000000000000000000000000000000000000384128116156101d1576101d161024b565b50500190565b6000808312837f8000000000000000000000000000000000000000000000000000000000000000018312811516156102115761021161024b565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0183138116156102455761024561024b565b50500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea26469706673582212201e03c8b68dcbcef6344fae810afa5d485b33bc238c0e8cb1508114b9c0ca702964736f6c63430008040033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// Deprecated: Use CounterMetaData.Sigs instead.
// CounterFuncSigs maps the 4-byte function signature to its string representation.
var CounterFuncSigs = CounterMetaData.Sigs

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterCallerSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetVitalikCount is a free data retrieval call binding the contract method 0x9f6f1ec1.
//
// Solidity: function getVitalikCount() view returns(int256)
func (_Counter *CounterCaller) GetVitalikCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getVitalikCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVitalikCount is a free data retrieval call binding the contract method 0x9f6f1ec1.
//
// Solidity: function getVitalikCount() view returns(int256)
func (_Counter *CounterSession) GetVitalikCount() (*big.Int, error) {
	return _Counter.Contract.GetVitalikCount(&_Counter.CallOpts)
}

// GetVitalikCount is a free data retrieval call binding the contract method 0x9f6f1ec1.
//
// Solidity: function getVitalikCount() view returns(int256)
func (_Counter *CounterCallerSession) GetVitalikCount() (*big.Int, error) {
	return _Counter.Contract.GetVitalikCount(&_Counter.CallOpts)
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_Counter *CounterTransactor) DecrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "decrementCounter")
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_Counter *CounterSession) DecrementCounter() (*types.Transaction, error) {
	return _Counter.Contract.DecrementCounter(&_Counter.TransactOpts)
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_Counter *CounterTransactorSession) DecrementCounter() (*types.Transaction, error) {
	return _Counter.Contract.DecrementCounter(&_Counter.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Counter *CounterTransactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Counter *CounterSession) IncrementCounter() (*types.Transaction, error) {
	return _Counter.Contract.IncrementCounter(&_Counter.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Counter *CounterTransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _Counter.Contract.IncrementCounter(&_Counter.TransactOpts)
}

// VitalikIncrement is a paid mutator transaction binding the contract method 0x6c573535.
//
// Solidity: function vitalikIncrement() returns()
func (_Counter *CounterTransactor) VitalikIncrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "vitalikIncrement")
}

// VitalikIncrement is a paid mutator transaction binding the contract method 0x6c573535.
//
// Solidity: function vitalikIncrement() returns()
func (_Counter *CounterSession) VitalikIncrement() (*types.Transaction, error) {
	return _Counter.Contract.VitalikIncrement(&_Counter.TransactOpts)
}

// VitalikIncrement is a paid mutator transaction binding the contract method 0x6c573535.
//
// Solidity: function vitalikIncrement() returns()
func (_Counter *CounterTransactorSession) VitalikIncrement() (*types.Transaction, error) {
	return _Counter.Contract.VitalikIncrement(&_Counter.TransactOpts)
}
