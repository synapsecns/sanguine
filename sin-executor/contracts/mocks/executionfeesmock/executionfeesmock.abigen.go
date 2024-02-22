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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"addExecutionFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"getAccumulatedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"getUnclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unclaimed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"recordExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ffecec7e": "addExecutionFee(uint256,bytes32)",
		"10886ac4": "claimExecutionFees()",
		"5ee09669": "getAccumulatedRewards(address)",
		"69a69e29": "getUnclaimedRewards(address)",
		"0676b706": "recordExecutor(uint256,bytes32,address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506101aa806100206000396000f3fe60806040526004361061005a5760003560e01c80635ee09669116100435780635ee096691461008d57806369a69e291461008d578063ffecec7e146100c057600080fd5b80630676b7061461005f57806310886ac414610081575b600080fd5b34801561006b57600080fd5b5061007f61007a3660046100fb565b505050565b005b34801561007f57600080fd5b34801561009957600080fd5b506100ae6100a8366004610130565b50600090565b60405190815260200160405180910390f35b61007f6100ce366004610152565b5050565b803573ffffffffffffffffffffffffffffffffffffffff811681146100f657600080fd5b919050565b60008060006060848603121561011057600080fd5b8335925060208401359150610127604085016100d2565b90509250925092565b60006020828403121561014257600080fd5b61014b826100d2565b9392505050565b6000806040838503121561016557600080fd5b5050803592602090910135915056fea2646970667358221220a86f7bc8f1b15dbf063f62b92f027b8f75b3126747b9f94226d39da3394ec07d64736f6c63430008140033",
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

// GetAccumulatedRewards is a free data retrieval call binding the contract method 0x5ee09669.
//
// Solidity: function getAccumulatedRewards(address executor) view returns(uint256 accumulated)
func (_ExecutionFeesMock *ExecutionFeesMockCaller) GetAccumulatedRewards(opts *bind.CallOpts, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExecutionFeesMock.contract.Call(opts, &out, "getAccumulatedRewards", executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccumulatedRewards is a free data retrieval call binding the contract method 0x5ee09669.
//
// Solidity: function getAccumulatedRewards(address executor) view returns(uint256 accumulated)
func (_ExecutionFeesMock *ExecutionFeesMockSession) GetAccumulatedRewards(executor common.Address) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.GetAccumulatedRewards(&_ExecutionFeesMock.CallOpts, executor)
}

// GetAccumulatedRewards is a free data retrieval call binding the contract method 0x5ee09669.
//
// Solidity: function getAccumulatedRewards(address executor) view returns(uint256 accumulated)
func (_ExecutionFeesMock *ExecutionFeesMockCallerSession) GetAccumulatedRewards(executor common.Address) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.GetAccumulatedRewards(&_ExecutionFeesMock.CallOpts, executor)
}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_ExecutionFeesMock *ExecutionFeesMockCaller) GetUnclaimedRewards(opts *bind.CallOpts, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExecutionFeesMock.contract.Call(opts, &out, "getUnclaimedRewards", executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_ExecutionFeesMock *ExecutionFeesMockSession) GetUnclaimedRewards(executor common.Address) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.GetUnclaimedRewards(&_ExecutionFeesMock.CallOpts, executor)
}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_ExecutionFeesMock *ExecutionFeesMockCallerSession) GetUnclaimedRewards(executor common.Address) (*big.Int, error) {
	return _ExecutionFeesMock.Contract.GetUnclaimedRewards(&_ExecutionFeesMock.CallOpts, executor)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0xffecec7e.
//
// Solidity: function addExecutionFee(uint256 dstChainId, bytes32 transactionId) payable returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactor) AddExecutionFee(opts *bind.TransactOpts, dstChainId *big.Int, transactionId [32]byte) (*types.Transaction, error) {
	return _ExecutionFeesMock.contract.Transact(opts, "addExecutionFee", dstChainId, transactionId)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0xffecec7e.
//
// Solidity: function addExecutionFee(uint256 dstChainId, bytes32 transactionId) payable returns()
func (_ExecutionFeesMock *ExecutionFeesMockSession) AddExecutionFee(dstChainId *big.Int, transactionId [32]byte) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.AddExecutionFee(&_ExecutionFeesMock.TransactOpts, dstChainId, transactionId)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0xffecec7e.
//
// Solidity: function addExecutionFee(uint256 dstChainId, bytes32 transactionId) payable returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactorSession) AddExecutionFee(dstChainId *big.Int, transactionId [32]byte) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.AddExecutionFee(&_ExecutionFeesMock.TransactOpts, dstChainId, transactionId)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x10886ac4.
//
// Solidity: function claimExecutionFees() returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactor) ClaimExecutionFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionFeesMock.contract.Transact(opts, "claimExecutionFees")
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x10886ac4.
//
// Solidity: function claimExecutionFees() returns()
func (_ExecutionFeesMock *ExecutionFeesMockSession) ClaimExecutionFees() (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.ClaimExecutionFees(&_ExecutionFeesMock.TransactOpts)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x10886ac4.
//
// Solidity: function claimExecutionFees() returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactorSession) ClaimExecutionFees() (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.ClaimExecutionFees(&_ExecutionFeesMock.TransactOpts)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0x0676b706.
//
// Solidity: function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactor) RecordExecutor(opts *bind.TransactOpts, dstChainId *big.Int, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.contract.Transact(opts, "recordExecutor", dstChainId, transactionId, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0x0676b706.
//
// Solidity: function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockSession) RecordExecutor(dstChainId *big.Int, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.RecordExecutor(&_ExecutionFeesMock.TransactOpts, dstChainId, transactionId, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0x0676b706.
//
// Solidity: function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) returns()
func (_ExecutionFeesMock *ExecutionFeesMockTransactorSession) RecordExecutor(dstChainId *big.Int, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _ExecutionFeesMock.Contract.RecordExecutor(&_ExecutionFeesMock.TransactOpts, dstChainId, transactionId, executor)
}

// IExecutionFeesMetaData contains all meta data concerning the IExecutionFees contract.
var IExecutionFeesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"addExecutionFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"getAccumulatedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"getUnclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unclaimed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"recordExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ffecec7e": "addExecutionFee(uint256,bytes32)",
		"10886ac4": "claimExecutionFees()",
		"5ee09669": "getAccumulatedRewards(address)",
		"69a69e29": "getUnclaimedRewards(address)",
		"0676b706": "recordExecutor(uint256,bytes32,address)",
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

// GetAccumulatedRewards is a free data retrieval call binding the contract method 0x5ee09669.
//
// Solidity: function getAccumulatedRewards(address executor) view returns(uint256 accumulated)
func (_IExecutionFees *IExecutionFeesCaller) GetAccumulatedRewards(opts *bind.CallOpts, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionFees.contract.Call(opts, &out, "getAccumulatedRewards", executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccumulatedRewards is a free data retrieval call binding the contract method 0x5ee09669.
//
// Solidity: function getAccumulatedRewards(address executor) view returns(uint256 accumulated)
func (_IExecutionFees *IExecutionFeesSession) GetAccumulatedRewards(executor common.Address) (*big.Int, error) {
	return _IExecutionFees.Contract.GetAccumulatedRewards(&_IExecutionFees.CallOpts, executor)
}

// GetAccumulatedRewards is a free data retrieval call binding the contract method 0x5ee09669.
//
// Solidity: function getAccumulatedRewards(address executor) view returns(uint256 accumulated)
func (_IExecutionFees *IExecutionFeesCallerSession) GetAccumulatedRewards(executor common.Address) (*big.Int, error) {
	return _IExecutionFees.Contract.GetAccumulatedRewards(&_IExecutionFees.CallOpts, executor)
}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_IExecutionFees *IExecutionFeesCaller) GetUnclaimedRewards(opts *bind.CallOpts, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionFees.contract.Call(opts, &out, "getUnclaimedRewards", executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_IExecutionFees *IExecutionFeesSession) GetUnclaimedRewards(executor common.Address) (*big.Int, error) {
	return _IExecutionFees.Contract.GetUnclaimedRewards(&_IExecutionFees.CallOpts, executor)
}

// GetUnclaimedRewards is a free data retrieval call binding the contract method 0x69a69e29.
//
// Solidity: function getUnclaimedRewards(address executor) view returns(uint256 unclaimed)
func (_IExecutionFees *IExecutionFeesCallerSession) GetUnclaimedRewards(executor common.Address) (*big.Int, error) {
	return _IExecutionFees.Contract.GetUnclaimedRewards(&_IExecutionFees.CallOpts, executor)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0xffecec7e.
//
// Solidity: function addExecutionFee(uint256 dstChainId, bytes32 transactionId) payable returns()
func (_IExecutionFees *IExecutionFeesTransactor) AddExecutionFee(opts *bind.TransactOpts, dstChainId *big.Int, transactionId [32]byte) (*types.Transaction, error) {
	return _IExecutionFees.contract.Transact(opts, "addExecutionFee", dstChainId, transactionId)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0xffecec7e.
//
// Solidity: function addExecutionFee(uint256 dstChainId, bytes32 transactionId) payable returns()
func (_IExecutionFees *IExecutionFeesSession) AddExecutionFee(dstChainId *big.Int, transactionId [32]byte) (*types.Transaction, error) {
	return _IExecutionFees.Contract.AddExecutionFee(&_IExecutionFees.TransactOpts, dstChainId, transactionId)
}

// AddExecutionFee is a paid mutator transaction binding the contract method 0xffecec7e.
//
// Solidity: function addExecutionFee(uint256 dstChainId, bytes32 transactionId) payable returns()
func (_IExecutionFees *IExecutionFeesTransactorSession) AddExecutionFee(dstChainId *big.Int, transactionId [32]byte) (*types.Transaction, error) {
	return _IExecutionFees.Contract.AddExecutionFee(&_IExecutionFees.TransactOpts, dstChainId, transactionId)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x10886ac4.
//
// Solidity: function claimExecutionFees() returns()
func (_IExecutionFees *IExecutionFeesTransactor) ClaimExecutionFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionFees.contract.Transact(opts, "claimExecutionFees")
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x10886ac4.
//
// Solidity: function claimExecutionFees() returns()
func (_IExecutionFees *IExecutionFeesSession) ClaimExecutionFees() (*types.Transaction, error) {
	return _IExecutionFees.Contract.ClaimExecutionFees(&_IExecutionFees.TransactOpts)
}

// ClaimExecutionFees is a paid mutator transaction binding the contract method 0x10886ac4.
//
// Solidity: function claimExecutionFees() returns()
func (_IExecutionFees *IExecutionFeesTransactorSession) ClaimExecutionFees() (*types.Transaction, error) {
	return _IExecutionFees.Contract.ClaimExecutionFees(&_IExecutionFees.TransactOpts)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0x0676b706.
//
// Solidity: function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) returns()
func (_IExecutionFees *IExecutionFeesTransactor) RecordExecutor(opts *bind.TransactOpts, dstChainId *big.Int, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.contract.Transact(opts, "recordExecutor", dstChainId, transactionId, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0x0676b706.
//
// Solidity: function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) returns()
func (_IExecutionFees *IExecutionFeesSession) RecordExecutor(dstChainId *big.Int, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.Contract.RecordExecutor(&_IExecutionFees.TransactOpts, dstChainId, transactionId, executor)
}

// RecordExecutor is a paid mutator transaction binding the contract method 0x0676b706.
//
// Solidity: function recordExecutor(uint256 dstChainId, bytes32 transactionId, address executor) returns()
func (_IExecutionFees *IExecutionFeesTransactorSession) RecordExecutor(dstChainId *big.Int, transactionId [32]byte, executor common.Address) (*types.Transaction, error) {
	return _IExecutionFees.Contract.RecordExecutor(&_IExecutionFees.TransactOpts, dstChainId, transactionId, executor)
}
