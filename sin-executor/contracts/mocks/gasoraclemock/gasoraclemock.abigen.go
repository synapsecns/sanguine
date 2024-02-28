// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gasoraclemock

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

// GasOracleMockMetaData contains all meta data concerning the GasOracleMock contract.
var GasOracleMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1e7b9287": "convertRemoteValueToLocalUnits(uint256,uint256)",
		"5cbd3c48": "estimateTxCostInLocalUnits(uint256,uint256,uint256)",
		"fd6a7167": "estimateTxCostInRemoteUnits(uint256,uint256,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060fc8061001f6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80631e7b92871460415780635cbd3c48146066578063fd6a7167146066575b600080fd5b6054604c366004607a565b600092915050565b60405190815260200160405180910390f35b60546071366004609b565b60009392505050565b60008060408385031215608c57600080fd5b50508035926020909101359150565b60008060006060848603121560af57600080fd5b50508135936020830135935060409092013591905056fea2646970667358221220d8e802cf67e61b25c6022989c4dcd8d2d60a5a11a7319acb8f31a0840e26db5864736f6c63430008140033",
}

// GasOracleMockABI is the input ABI used to generate the binding from.
// Deprecated: Use GasOracleMockMetaData.ABI instead.
var GasOracleMockABI = GasOracleMockMetaData.ABI

// Deprecated: Use GasOracleMockMetaData.Sigs instead.
// GasOracleMockFuncSigs maps the 4-byte function signature to its string representation.
var GasOracleMockFuncSigs = GasOracleMockMetaData.Sigs

// GasOracleMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasOracleMockMetaData.Bin instead.
var GasOracleMockBin = GasOracleMockMetaData.Bin

// DeployGasOracleMock deploys a new Ethereum contract, binding an instance of GasOracleMock to it.
func DeployGasOracleMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasOracleMock, error) {
	parsed, err := GasOracleMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasOracleMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasOracleMock{GasOracleMockCaller: GasOracleMockCaller{contract: contract}, GasOracleMockTransactor: GasOracleMockTransactor{contract: contract}, GasOracleMockFilterer: GasOracleMockFilterer{contract: contract}}, nil
}

// GasOracleMock is an auto generated Go binding around an Ethereum contract.
type GasOracleMock struct {
	GasOracleMockCaller     // Read-only binding to the contract
	GasOracleMockTransactor // Write-only binding to the contract
	GasOracleMockFilterer   // Log filterer for contract events
}

// GasOracleMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasOracleMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasOracleMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasOracleMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasOracleMockSession struct {
	Contract     *GasOracleMock    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasOracleMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasOracleMockCallerSession struct {
	Contract *GasOracleMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GasOracleMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasOracleMockTransactorSession struct {
	Contract     *GasOracleMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GasOracleMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasOracleMockRaw struct {
	Contract *GasOracleMock // Generic contract binding to access the raw methods on
}

// GasOracleMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasOracleMockCallerRaw struct {
	Contract *GasOracleMockCaller // Generic read-only contract binding to access the raw methods on
}

// GasOracleMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasOracleMockTransactorRaw struct {
	Contract *GasOracleMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasOracleMock creates a new instance of GasOracleMock, bound to a specific deployed contract.
func NewGasOracleMock(address common.Address, backend bind.ContractBackend) (*GasOracleMock, error) {
	contract, err := bindGasOracleMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasOracleMock{GasOracleMockCaller: GasOracleMockCaller{contract: contract}, GasOracleMockTransactor: GasOracleMockTransactor{contract: contract}, GasOracleMockFilterer: GasOracleMockFilterer{contract: contract}}, nil
}

// NewGasOracleMockCaller creates a new read-only instance of GasOracleMock, bound to a specific deployed contract.
func NewGasOracleMockCaller(address common.Address, caller bind.ContractCaller) (*GasOracleMockCaller, error) {
	contract, err := bindGasOracleMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasOracleMockCaller{contract: contract}, nil
}

// NewGasOracleMockTransactor creates a new write-only instance of GasOracleMock, bound to a specific deployed contract.
func NewGasOracleMockTransactor(address common.Address, transactor bind.ContractTransactor) (*GasOracleMockTransactor, error) {
	contract, err := bindGasOracleMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasOracleMockTransactor{contract: contract}, nil
}

// NewGasOracleMockFilterer creates a new log filterer instance of GasOracleMock, bound to a specific deployed contract.
func NewGasOracleMockFilterer(address common.Address, filterer bind.ContractFilterer) (*GasOracleMockFilterer, error) {
	contract, err := bindGasOracleMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasOracleMockFilterer{contract: contract}, nil
}

// bindGasOracleMock binds a generic wrapper to an already deployed contract.
func bindGasOracleMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasOracleMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasOracleMock *GasOracleMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasOracleMock.Contract.GasOracleMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasOracleMock *GasOracleMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasOracleMock.Contract.GasOracleMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasOracleMock *GasOracleMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasOracleMock.Contract.GasOracleMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasOracleMock *GasOracleMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasOracleMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasOracleMock *GasOracleMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasOracleMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasOracleMock *GasOracleMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasOracleMock.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_GasOracleMock *GasOracleMockCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasOracleMock.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_GasOracleMock *GasOracleMockSession) ConvertRemoteValueToLocalUnits(remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	return _GasOracleMock.Contract.ConvertRemoteValueToLocalUnits(&_GasOracleMock.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_GasOracleMock *GasOracleMockCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	return _GasOracleMock.Contract.ConvertRemoteValueToLocalUnits(&_GasOracleMock.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_GasOracleMock *GasOracleMockCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasOracleMock.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_GasOracleMock *GasOracleMockSession) EstimateTxCostInLocalUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _GasOracleMock.Contract.EstimateTxCostInLocalUnits(&_GasOracleMock.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_GasOracleMock *GasOracleMockCallerSession) EstimateTxCostInLocalUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _GasOracleMock.Contract.EstimateTxCostInLocalUnits(&_GasOracleMock.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_GasOracleMock *GasOracleMockCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasOracleMock.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_GasOracleMock *GasOracleMockSession) EstimateTxCostInRemoteUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _GasOracleMock.Contract.EstimateTxCostInRemoteUnits(&_GasOracleMock.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_GasOracleMock *GasOracleMockCallerSession) EstimateTxCostInRemoteUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _GasOracleMock.Contract.EstimateTxCostInRemoteUnits(&_GasOracleMock.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// IGasOracleMetaData contains all meta data concerning the IGasOracle contract.
var IGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1e7b9287": "convertRemoteValueToLocalUnits(uint256,uint256)",
		"5cbd3c48": "estimateTxCostInLocalUnits(uint256,uint256,uint256)",
		"fd6a7167": "estimateTxCostInRemoteUnits(uint256,uint256,uint256)",
	},
}

// IGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IGasOracleMetaData.ABI instead.
var IGasOracleABI = IGasOracleMetaData.ABI

// Deprecated: Use IGasOracleMetaData.Sigs instead.
// IGasOracleFuncSigs maps the 4-byte function signature to its string representation.
var IGasOracleFuncSigs = IGasOracleMetaData.Sigs

// IGasOracle is an auto generated Go binding around an Ethereum contract.
type IGasOracle struct {
	IGasOracleCaller     // Read-only binding to the contract
	IGasOracleTransactor // Write-only binding to the contract
	IGasOracleFilterer   // Log filterer for contract events
}

// IGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGasOracleSession struct {
	Contract     *IGasOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGasOracleCallerSession struct {
	Contract *IGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGasOracleTransactorSession struct {
	Contract     *IGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGasOracleRaw struct {
	Contract *IGasOracle // Generic contract binding to access the raw methods on
}

// IGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGasOracleCallerRaw struct {
	Contract *IGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGasOracleTransactorRaw struct {
	Contract *IGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGasOracle creates a new instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracle(address common.Address, backend bind.ContractBackend) (*IGasOracle, error) {
	contract, err := bindIGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGasOracle{IGasOracleCaller: IGasOracleCaller{contract: contract}, IGasOracleTransactor: IGasOracleTransactor{contract: contract}, IGasOracleFilterer: IGasOracleFilterer{contract: contract}}, nil
}

// NewIGasOracleCaller creates a new read-only instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleCaller(address common.Address, caller bind.ContractCaller) (*IGasOracleCaller, error) {
	contract, err := bindIGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGasOracleCaller{contract: contract}, nil
}

// NewIGasOracleTransactor creates a new write-only instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IGasOracleTransactor, error) {
	contract, err := bindIGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGasOracleTransactor{contract: contract}, nil
}

// NewIGasOracleFilterer creates a new log filterer instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IGasOracleFilterer, error) {
	contract, err := bindIGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGasOracleFilterer{contract: contract}, nil
}

// bindIGasOracle binds a generic wrapper to an already deployed contract.
func bindIGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGasOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasOracle *IGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasOracle.Contract.IGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasOracle *IGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasOracle.Contract.IGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasOracle *IGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasOracle.Contract.IGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasOracle *IGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasOracle *IGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasOracle *IGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasOracle.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleSession) ConvertRemoteValueToLocalUnits(remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInLocalUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInLocalUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInRemoteUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInRemoteUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}
