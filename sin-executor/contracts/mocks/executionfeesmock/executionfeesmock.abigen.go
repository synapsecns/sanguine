// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executionfeesmock

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

// ExecutionFeesMockMetaData contains all meta data concerning the ExecutionFeesMock contract.
var ExecutionFeesMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionFees__AlreadyRecorded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFees__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFees__ZeroAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"accumulatedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"addExecutionFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"claimExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"executionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"recordExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"recordedExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"unclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unclaimed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"73f273fc": "accumulatedRewards(address)",
		"7b18c25c": "addExecutionFee(uint64,bytes32)",
		"4e497dac": "claimExecutionFees(address)",
		"656a96d9": "executionFee(uint64,bytes32)",
		"fd411b43": "recordExecutor(uint64,bytes32,address)",
		"c2bc3357": "recordedExecutor(uint64,bytes32)",
		"949813b8": "unclaimedRewards(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061025d806100206000396000f3fe6080604052600436106100705760003560e01c80637b18c25c1161004e5780637b18c25c146100ec578063949813b8146100cb578063c2bc3357146100fe578063fd411b431461013e57600080fd5b80634e497dac14610075578063656a96d91461009557806373f273fc146100cb575b600080fd5b34801561008157600080fd5b50610093610090366004610187565b50565b005b3480156100a157600080fd5b506100b86100b03660046101c1565b600092915050565b6040519081526020015b60405180910390f35b3480156100d757600080fd5b506100b86100e6366004610187565b50600090565b6100936100fa3660046101c1565b5050565b34801561010a57600080fd5b506101196100b03660046101c1565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100c2565b34801561014a57600080fd5b506100936101593660046101eb565b505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461018257600080fd5b919050565b60006020828403121561019957600080fd5b6101a28261015e565b9392505050565b803567ffffffffffffffff8116811461018257600080fd5b600080604083850312156101d457600080fd5b6101dd836101a9565b946020939093013593505050565b60008060006060848603121561020057600080fd5b610209846101a9565b92506020840135915061021e6040850161015e565b9050925092509256fea2646970667358221220eed4b061a2e1f9d5d4ec4069a18a2b9b43e9c4813a7e6f2f7639a225fb8120d064736f6c63430008140033",
}

// ExecutionFeesMockABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionFeesMockMetaData.ABI instead.
var ExecutionFeesMockABI = ExecutionFeesMockMetaData.ABI

// Deprecated: Use ExecutionFeesMockMetaData.Sigs instead.
// ExecutionFeesMockFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionFeesMockFuncSigs = ExecutionFeesMockMetaData.Sigs

// ExecutionFeesMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutionFeesMockMetaData.Bin instead.
var ExecutionFeesMockBin = ExecutionFeesMockMetaData.Bin

// DeployExecutionFeesMock deploys a new Ethereum contract, binding an instance of ExecutionFeesMock to it.
func DeployExecutionFeesMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExecutionFeesMock, error) {
	parsed, err := ExecutionFeesMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutionFeesMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutionFeesMock{ExecutionFeesMockCaller: ExecutionFeesMockCaller{contract: contract}, ExecutionFeesMockTransactor: ExecutionFeesMockTransactor{contract: contract}, ExecutionFeesMockFilterer: ExecutionFeesMockFilterer{contract: contract}}, nil
}

// ExecutionFeesMock is an auto generated Go binding around an Ethereum contract.
type ExecutionFeesMock struct {
	ExecutionFeesMockCaller     // Read-only binding to the contract
	ExecutionFeesMockTransactor // Write-only binding to the contract
	ExecutionFeesMockFilterer   // Log filterer for contract events
}

// ExecutionFeesMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionFeesMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionFeesMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionFeesMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionFeesMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionFeesMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionFeesMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionFeesMockSession struct {
	Contract     *ExecutionFeesMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExecutionFeesMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionFeesMockCallerSession struct {
	Contract *ExecutionFeesMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ExecutionFeesMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionFeesMockTransactorSession struct {
	Contract     *ExecutionFeesMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ExecutionFeesMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionFeesMockRaw struct {
	Contract *ExecutionFeesMock // Generic contract binding to access the raw methods on
}

// ExecutionFeesMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionFeesMockCallerRaw struct {
	Contract *ExecutionFeesMockCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionFeesMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionFeesMockTransactorRaw struct {
	Contract *ExecutionFeesMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionFeesMock creates a new instance of ExecutionFeesMock, bound to a specific deployed contract.
func NewExecutionFeesMock(address common.Address, backend bind.ContractBackend) (*ExecutionFeesMock, error) {
	contract, err := bindExecutionFeesMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionFeesMock{ExecutionFeesMockCaller: ExecutionFeesMockCaller{contract: contract}, ExecutionFeesMockTransactor: ExecutionFeesMockTransactor{contract: contract}, ExecutionFeesMockFilterer: ExecutionFeesMockFilterer{contract: contract}}, nil
}

// NewExecutionFeesMockCaller creates a new read-only instance of ExecutionFeesMock, bound to a specific deployed contract.
func NewExecutionFeesMockCaller(address common.Address, caller bind.ContractCaller) (*ExecutionFeesMockCaller, error) {
	contract, err := bindExecutionFeesMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionFeesMockCaller{contract: contract}, nil
}

// NewExecutionFeesMockTransactor creates a new write-only instance of ExecutionFeesMock, bound to a specific deployed contract.
func NewExecutionFeesMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionFeesMockTransactor, error) {
	contract, err := bindExecutionFeesMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionFeesMockTransactor{contract: contract}, nil
}

// NewExecutionFeesMockFilterer creates a new log filterer instance of ExecutionFeesMock, bound to a specific deployed contract.
func NewExecutionFeesMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionFeesMockFilterer, error) {
	contract, err := bindExecutionFeesMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionFeesMockFilterer{contract: contract}, nil
}

// bindExecutionFeesMock binds a generic wrapper to an already deployed contract.
func bindExecutionFeesMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutionFeesMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionFeesMock *ExecutionFeesMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionFeesMock.Contract.ExecutionFeesMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionFeesMock *ExecutionFeesMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.ExecutionFeesMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionFeesMock *ExecutionFeesMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.ExecutionFeesMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionFeesMock *ExecutionFeesMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionFeesMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionFeesMock *ExecutionFeesMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionFeesMock *ExecutionFeesMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address executor) view returns(uint256 accumulated)
func (_ExecutionFeesMock *ExecutionFeesMockCaller) AccumulatedRewards(opts *bind.CallOpts, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExecutionFeesMock.contract.Call(opts, &out, "accumulatedRewards", executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address executor) view returns(uint256 accumulated)
func (_ExecutionFeesMock *ExecutionFeesMockSession) AccumulatedRewards(executor common.Address) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.AccumulatedRewards(&_ExecutionFeesMock.CallOpts, executor)
}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address executor) view returns(uint256 accumulated)
func (_ExecutionFeesMock *ExecutionFeesMockCallerSession) AccumulatedRewards(executor common.Address) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.AccumulatedRewards(&_ExecutionFeesMock.CallOpts, executor)
}

// ExecutionFee is a free data retrieval call binding the contract method 0x656a96d9.
//
// Solidity: function executionFee(uint64 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_ExecutionFeesMock *ExecutionFeesMockCaller) ExecutionFee(opts *bind.CallOpts, dstChainId uint64, transactionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ExecutionFeesMock.contract.Call(opts, &out, "executionFee", dstChainId, transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutionFee is a free data retrieval call binding the contract method 0x656a96d9.
//
// Solidity: function executionFee(uint64 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_ExecutionFeesMock *ExecutionFeesMockSession) ExecutionFee(dstChainId uint64, transactionId [32]byte) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.ExecutionFee(&_ExecutionFeesMock.CallOpts, dstChainId, transactionId)
}

// ExecutionFee is a free data retrieval call binding the contract method 0x656a96d9.
//
// Solidity: function executionFee(uint64 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_ExecutionFeesMock *ExecutionFeesMockCallerSession) ExecutionFee(dstChainId uint64, transactionId [32]byte) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.ExecutionFee(&_ExecutionFeesMock.CallOpts, dstChainId, transactionId)
}

// RecordedExecutor is a free data retrieval call binding the contract method 0xc2bc3357.
//
// Solidity: function recordedExecutor(uint64 dstChainId, bytes32 transactionId) view returns(address executor)
func (_ExecutionFeesMock *ExecutionFeesMockCaller) RecordedExecutor(opts *bind.CallOpts, dstChainId uint64, transactionId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ExecutionFeesMock.contract.Call(opts, &out, "recordedExecutor", dstChainId, transactionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecordedExecutor is a free data retrieval call binding the contract method 0xc2bc3357.
//
// Solidity: function recordedExecutor(uint64 dstChainId, bytes32 transactionId) view returns(address executor)
func (_ExecutionFeesMock *ExecutionFeesMockSession) RecordedExecutor(dstChainId uint64, transactionId [32]byte) (common.Address, error) {
	return _ExecutionFeesMock.Contract.RecordedExecutor(&_ExecutionFeesMock.CallOpts, dstChainId, transactionId)
}

// RecordedExecutor is a free data retrieval call binding the contract method 0xc2bc3357.
//
// Solidity: function recordedExecutor(uint64 dstChainId, bytes32 transactionId) view returns(address executor)
func (_ExecutionFeesMock *ExecutionFeesMockCallerSession) RecordedExecutor(dstChainId uint64, transactionId [32]byte) (common.Address, error) {
	return _ExecutionFeesMock.Contract.RecordedExecutor(&_ExecutionFeesMock.CallOpts, dstChainId, transactionId)
}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_ExecutionFeesMock *ExecutionFeesMockCaller) UnclaimedRewards(opts *bind.CallOpts, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExecutionFeesMock.contract.Call(opts, &out, "unclaimedRewards", executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_ExecutionFeesMock *ExecutionFeesMockSession) UnclaimedRewards(executor common.Address) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.UnclaimedRewards(&_ExecutionFeesMock.CallOpts, executor)
}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_ExecutionFeesMock *ExecutionFeesMockCallerSession) UnclaimedRewards(executor common.Address) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.UnclaimedRewards(&_ExecutionFeesMock.CallOpts, executor)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0x7b18c25c.
//
// Solidity: function addExecutionFee(uint64 dstChainId, bytes32 transactionId) payable returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactor) AddExecutionFee(opts *bind.TransactOpts, dstChainId uint64, transactionId [32]byte) (*types.Transaction, error) {
	return _ExecutionFeesMock.contract.Transact(opts, "addExecutionFee", dstChainId, transactionId)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0x7b18c25c.
//
// Solidity: function addExecutionFee(uint64 dstChainId, bytes32 transactionId) payable returns()
func (_ExecutionFeesMock *ExecutionFeesMockSession) AddExecutionFee(dstChainId uint64, transactionId [32]byte) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.AddExecutionFee(&_ExecutionFeesMock.TransactOpts, dstChainId, transactionId)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0x7b18c25c.
//
// Solidity: function addExecutionFee(uint64 dstChainId, bytes32 transactionId) payable returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactorSession) AddExecutionFee(dstChainId uint64, transactionId [32]byte) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.AddExecutionFee(&_ExecutionFeesMock.TransactOpts, dstChainId, transactionId)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x4e497dac.
//
// Solidity: function claimExecutionFees(address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactor) ClaimExecutionFees(opts *bind.TransactOpts, executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.contract.Transact(opts, "claimExecutionFees", executor)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x4e497dac.
//
// Solidity: function claimExecutionFees(address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockSession) ClaimExecutionFees(executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.ClaimExecutionFees(&_ExecutionFeesMock.TransactOpts, executor)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x4e497dac.
//
// Solidity: function claimExecutionFees(address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactorSession) ClaimExecutionFees(executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.ClaimExecutionFees(&_ExecutionFeesMock.TransactOpts, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0xfd411b43.
//
// Solidity: function recordExecutor(uint64 dstChainId, bytes32 transactionId, address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactor) RecordExecutor(opts *bind.TransactOpts, dstChainId uint64, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.contract.Transact(opts, "recordExecutor", dstChainId, transactionId, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0xfd411b43.
//
// Solidity: function recordExecutor(uint64 dstChainId, bytes32 transactionId, address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockSession) RecordExecutor(dstChainId uint64, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.RecordExecutor(&_ExecutionFeesMock.TransactOpts, dstChainId, transactionId, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0xfd411b43.
//
// Solidity: function recordExecutor(uint64 dstChainId, bytes32 transactionId, address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactorSession) RecordExecutor(dstChainId uint64, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.RecordExecutor(&_ExecutionFeesMock.TransactOpts, dstChainId, transactionId, executor)
}

// IExecutionFeesMetaData contains all meta data concerning the IExecutionFees contract.
var IExecutionFeesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionFees__AlreadyRecorded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFees__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFees__ZeroAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"accumulatedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"addExecutionFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"claimExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"executionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"recordExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"recordedExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"unclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unclaimed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"73f273fc": "accumulatedRewards(address)",
		"7b18c25c": "addExecutionFee(uint64,bytes32)",
		"4e497dac": "claimExecutionFees(address)",
		"656a96d9": "executionFee(uint64,bytes32)",
		"fd411b43": "recordExecutor(uint64,bytes32,address)",
		"c2bc3357": "recordedExecutor(uint64,bytes32)",
		"949813b8": "unclaimedRewards(address)",
	},
}

// IExecutionFeesABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionFeesMetaData.ABI instead.
var IExecutionFeesABI = IExecutionFeesMetaData.ABI

// Deprecated: Use IExecutionFeesMetaData.Sigs instead.
// IExecutionFeesFuncSigs maps the 4-byte function signature to its string representation.
var IExecutionFeesFuncSigs = IExecutionFeesMetaData.Sigs

// IExecutionFees is an auto generated Go binding around an Ethereum contract.
type IExecutionFees struct {
	IExecutionFeesCaller     // Read-only binding to the contract
	IExecutionFeesTransactor // Write-only binding to the contract
	IExecutionFeesFilterer   // Log filterer for contract events
}

// IExecutionFeesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionFeesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionFeesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionFeesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionFeesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionFeesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionFeesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionFeesSession struct {
	Contract     *IExecutionFees   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IExecutionFeesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionFeesCallerSession struct {
	Contract *IExecutionFeesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IExecutionFeesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionFeesTransactorSession struct {
	Contract     *IExecutionFeesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IExecutionFeesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionFeesRaw struct {
	Contract *IExecutionFees // Generic contract binding to access the raw methods on
}

// IExecutionFeesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionFeesCallerRaw struct {
	Contract *IExecutionFeesCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionFeesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionFeesTransactorRaw struct {
	Contract *IExecutionFeesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecutionFees creates a new instance of IExecutionFees, bound to a specific deployed contract.
func NewIExecutionFees(address common.Address, backend bind.ContractBackend) (*IExecutionFees, error) {
	contract, err := bindIExecutionFees(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecutionFees{IExecutionFeesCaller: IExecutionFeesCaller{contract: contract}, IExecutionFeesTransactor: IExecutionFeesTransactor{contract: contract}, IExecutionFeesFilterer: IExecutionFeesFilterer{contract: contract}}, nil
}

// NewIExecutionFeesCaller creates a new read-only instance of IExecutionFees, bound to a specific deployed contract.
func NewIExecutionFeesCaller(address common.Address, caller bind.ContractCaller) (*IExecutionFeesCaller, error) {
	contract, err := bindIExecutionFees(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionFeesCaller{contract: contract}, nil
}

// NewIExecutionFeesTransactor creates a new write-only instance of IExecutionFees, bound to a specific deployed contract.
func NewIExecutionFeesTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionFeesTransactor, error) {
	contract, err := bindIExecutionFees(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionFeesTransactor{contract: contract}, nil
}

// NewIExecutionFeesFilterer creates a new log filterer instance of IExecutionFees, bound to a specific deployed contract.
func NewIExecutionFeesFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionFeesFilterer, error) {
	contract, err := bindIExecutionFees(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionFeesFilterer{contract: contract}, nil
}

// bindIExecutionFees binds a generic wrapper to an already deployed contract.
func bindIExecutionFees(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExecutionFeesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionFees *IExecutionFeesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionFees.Contract.IExecutionFeesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionFees *IExecutionFeesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionFees.Contract.IExecutionFeesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionFees *IExecutionFeesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionFees.Contract.IExecutionFeesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionFees *IExecutionFeesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionFees.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionFees *IExecutionFeesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionFees.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionFees *IExecutionFeesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionFees.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address executor) view returns(uint256 accumulated)
func (_IExecutionFees *IExecutionFeesCaller) AccumulatedRewards(opts *bind.CallOpts, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionFees.contract.Call(opts, &out, "accumulatedRewards", executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address executor) view returns(uint256 accumulated)
func (_IExecutionFees *IExecutionFeesSession) AccumulatedRewards(executor common.Address) (*big.Int, error) {
	return _IExecutionFees.Contract.AccumulatedRewards(&_IExecutionFees.CallOpts, executor)
}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address executor) view returns(uint256 accumulated)
func (_IExecutionFees *IExecutionFeesCallerSession) AccumulatedRewards(executor common.Address) (*big.Int, error) {
	return _IExecutionFees.Contract.AccumulatedRewards(&_IExecutionFees.CallOpts, executor)
}

// ExecutionFee is a free data retrieval call binding the contract method 0x656a96d9.
//
// Solidity: function executionFee(uint64 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_IExecutionFees *IExecutionFeesCaller) ExecutionFee(opts *bind.CallOpts, dstChainId uint64, transactionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionFees.contract.Call(opts, &out, "executionFee", dstChainId, transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutionFee is a free data retrieval call binding the contract method 0x656a96d9.
//
// Solidity: function executionFee(uint64 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_IExecutionFees *IExecutionFeesSession) ExecutionFee(dstChainId uint64, transactionId [32]byte) (*big.Int, error) {
	return _IExecutionFees.Contract.ExecutionFee(&_IExecutionFees.CallOpts, dstChainId, transactionId)
}

// ExecutionFee is a free data retrieval call binding the contract method 0x656a96d9.
//
// Solidity: function executionFee(uint64 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_IExecutionFees *IExecutionFeesCallerSession) ExecutionFee(dstChainId uint64, transactionId [32]byte) (*big.Int, error) {
	return _IExecutionFees.Contract.ExecutionFee(&_IExecutionFees.CallOpts, dstChainId, transactionId)
}

// RecordedExecutor is a free data retrieval call binding the contract method 0xc2bc3357.
//
// Solidity: function recordedExecutor(uint64 dstChainId, bytes32 transactionId) view returns(address executor)
func (_IExecutionFees *IExecutionFeesCaller) RecordedExecutor(opts *bind.CallOpts, dstChainId uint64, transactionId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IExecutionFees.contract.Call(opts, &out, "recordedExecutor", dstChainId, transactionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecordedExecutor is a free data retrieval call binding the contract method 0xc2bc3357.
//
// Solidity: function recordedExecutor(uint64 dstChainId, bytes32 transactionId) view returns(address executor)
func (_IExecutionFees *IExecutionFeesSession) RecordedExecutor(dstChainId uint64, transactionId [32]byte) (common.Address, error) {
	return _IExecutionFees.Contract.RecordedExecutor(&_IExecutionFees.CallOpts, dstChainId, transactionId)
}

// RecordedExecutor is a free data retrieval call binding the contract method 0xc2bc3357.
//
// Solidity: function recordedExecutor(uint64 dstChainId, bytes32 transactionId) view returns(address executor)
func (_IExecutionFees *IExecutionFeesCallerSession) RecordedExecutor(dstChainId uint64, transactionId [32]byte) (common.Address, error) {
	return _IExecutionFees.Contract.RecordedExecutor(&_IExecutionFees.CallOpts, dstChainId, transactionId)
}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_IExecutionFees *IExecutionFeesCaller) UnclaimedRewards(opts *bind.CallOpts, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionFees.contract.Call(opts, &out, "unclaimedRewards", executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_IExecutionFees *IExecutionFeesSession) UnclaimedRewards(executor common.Address) (*big.Int, error) {
	return _IExecutionFees.Contract.UnclaimedRewards(&_IExecutionFees.CallOpts, executor)
}

// UnclaimedRewards is a free data retrieval call binding the contract method 0x949813b8.
//
// Solidity: function unclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_IExecutionFees *IExecutionFeesCallerSession) UnclaimedRewards(executor common.Address) (*big.Int, error) {
	return _IExecutionFees.Contract.UnclaimedRewards(&_IExecutionFees.CallOpts, executor)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0x7b18c25c.
//
// Solidity: function addExecutionFee(uint64 dstChainId, bytes32 transactionId) payable returns()
func (_IExecutionFees *IExecutionFeesTransactor) AddExecutionFee(opts *bind.TransactOpts, dstChainId uint64, transactionId [32]byte) (*types.Transaction, error) {
	return _IExecutionFees.contract.Transact(opts, "addExecutionFee", dstChainId, transactionId)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0x7b18c25c.
//
// Solidity: function addExecutionFee(uint64 dstChainId, bytes32 transactionId) payable returns()
func (_IExecutionFees *IExecutionFeesSession) AddExecutionFee(dstChainId uint64, transactionId [32]byte) (*types.Transaction, error) {
	return _IExecutionFees.Contract.AddExecutionFee(&_IExecutionFees.TransactOpts, dstChainId, transactionId)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0x7b18c25c.
//
// Solidity: function addExecutionFee(uint64 dstChainId, bytes32 transactionId) payable returns()
func (_IExecutionFees *IExecutionFeesTransactorSession) AddExecutionFee(dstChainId uint64, transactionId [32]byte) (*types.Transaction, error) {
	return _IExecutionFees.Contract.AddExecutionFee(&_IExecutionFees.TransactOpts, dstChainId, transactionId)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x4e497dac.
//
// Solidity: function claimExecutionFees(address executor) returns()
func (_IExecutionFees *IExecutionFeesTransactor) ClaimExecutionFees(opts *bind.TransactOpts, executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.contract.Transact(opts, "claimExecutionFees", executor)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x4e497dac.
//
// Solidity: function claimExecutionFees(address executor) returns()
func (_IExecutionFees *IExecutionFeesSession) ClaimExecutionFees(executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.Contract.ClaimExecutionFees(&_IExecutionFees.TransactOpts, executor)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x4e497dac.
//
// Solidity: function claimExecutionFees(address executor) returns()
func (_IExecutionFees *IExecutionFeesTransactorSession) ClaimExecutionFees(executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.Contract.ClaimExecutionFees(&_IExecutionFees.TransactOpts, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0xfd411b43.
//
// Solidity: function recordExecutor(uint64 dstChainId, bytes32 transactionId, address executor) returns()
func (_IExecutionFees *IExecutionFeesTransactor) RecordExecutor(opts *bind.TransactOpts, dstChainId uint64, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.contract.Transact(opts, "recordExecutor", dstChainId, transactionId, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0xfd411b43.
//
// Solidity: function recordExecutor(uint64 dstChainId, bytes32 transactionId, address executor) returns()
func (_IExecutionFees *IExecutionFeesSession) RecordExecutor(dstChainId uint64, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.Contract.RecordExecutor(&_IExecutionFees.TransactOpts, dstChainId, transactionId, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0xfd411b43.
//
// Solidity: function recordExecutor(uint64 dstChainId, bytes32 transactionId, address executor) returns()
func (_IExecutionFees *IExecutionFeesTransactorSession) RecordExecutor(dstChainId uint64, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.Contract.RecordExecutor(&_IExecutionFees.TransactOpts, dstChainId, transactionId, executor)
}
