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

// IGasOracleMetaData contains all meta data concerning the IGasOracle contract.
var IGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40658a74": "convertRemoteValueToLocalUnits(uint64,uint256)",
		"bf495c88": "estimateTxCostInLocalUnits(uint64,uint256,uint256)",
		"b376a688": "estimateTxCostInRemoteUnits(uint64,uint256,uint256)",
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

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId uint64, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// ISynapseGasOracleMetaData contains all meta data concerning the ISynapseGasOracle contract.
var ISynapseGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLocalGasData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveRemoteGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40658a74": "convertRemoteValueToLocalUnits(uint64,uint256)",
		"bf495c88": "estimateTxCostInLocalUnits(uint64,uint256,uint256)",
		"b376a688": "estimateTxCostInRemoteUnits(uint64,uint256,uint256)",
		"6f928aa7": "getLocalGasData()",
		"83389de7": "receiveRemoteGasData(uint64,bytes)",
	},
}

// ISynapseGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseGasOracleMetaData.ABI instead.
var ISynapseGasOracleABI = ISynapseGasOracleMetaData.ABI

// Deprecated: Use ISynapseGasOracleMetaData.Sigs instead.
// ISynapseGasOracleFuncSigs maps the 4-byte function signature to its string representation.
var ISynapseGasOracleFuncSigs = ISynapseGasOracleMetaData.Sigs

// ISynapseGasOracle is an auto generated Go binding around an Ethereum contract.
type ISynapseGasOracle struct {
	ISynapseGasOracleCaller     // Read-only binding to the contract
	ISynapseGasOracleTransactor // Write-only binding to the contract
	ISynapseGasOracleFilterer   // Log filterer for contract events
}

// ISynapseGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseGasOracleSession struct {
	Contract     *ISynapseGasOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ISynapseGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseGasOracleCallerSession struct {
	Contract *ISynapseGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ISynapseGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseGasOracleTransactorSession struct {
	Contract     *ISynapseGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ISynapseGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseGasOracleRaw struct {
	Contract *ISynapseGasOracle // Generic contract binding to access the raw methods on
}

// ISynapseGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseGasOracleCallerRaw struct {
	Contract *ISynapseGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ISynapseGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseGasOracleTransactorRaw struct {
	Contract *ISynapseGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseGasOracle creates a new instance of ISynapseGasOracle, bound to a specific deployed contract.
func NewISynapseGasOracle(address common.Address, backend bind.ContractBackend) (*ISynapseGasOracle, error) {
	contract, err := bindISynapseGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracle{ISynapseGasOracleCaller: ISynapseGasOracleCaller{contract: contract}, ISynapseGasOracleTransactor: ISynapseGasOracleTransactor{contract: contract}, ISynapseGasOracleFilterer: ISynapseGasOracleFilterer{contract: contract}}, nil
}

// NewISynapseGasOracleCaller creates a new read-only instance of ISynapseGasOracle, bound to a specific deployed contract.
func NewISynapseGasOracleCaller(address common.Address, caller bind.ContractCaller) (*ISynapseGasOracleCaller, error) {
	contract, err := bindISynapseGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleCaller{contract: contract}, nil
}

// NewISynapseGasOracleTransactor creates a new write-only instance of ISynapseGasOracle, bound to a specific deployed contract.
func NewISynapseGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseGasOracleTransactor, error) {
	contract, err := bindISynapseGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleTransactor{contract: contract}, nil
}

// NewISynapseGasOracleFilterer creates a new log filterer instance of ISynapseGasOracle, bound to a specific deployed contract.
func NewISynapseGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseGasOracleFilterer, error) {
	contract, err := bindISynapseGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleFilterer{contract: contract}, nil
}

// bindISynapseGasOracle binds a generic wrapper to an already deployed contract.
func bindISynapseGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISynapseGasOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseGasOracle *ISynapseGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseGasOracle.Contract.ISynapseGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseGasOracle *ISynapseGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.ISynapseGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseGasOracle *ISynapseGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.ISynapseGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseGasOracle *ISynapseGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseGasOracle *ISynapseGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseGasOracle *ISynapseGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId uint64, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracle.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracle.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.EstimateTxCostInLocalUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCallerSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.EstimateTxCostInLocalUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracle.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.EstimateTxCostInRemoteUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCallerSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.EstimateTxCostInRemoteUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracle *ISynapseGasOracleCaller) GetLocalGasData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ISynapseGasOracle.contract.Call(opts, &out, "getLocalGasData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracle *ISynapseGasOracleSession) GetLocalGasData() ([]byte, error) {
	return _ISynapseGasOracle.Contract.GetLocalGasData(&_ISynapseGasOracle.CallOpts)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracle *ISynapseGasOracleCallerSession) GetLocalGasData() ([]byte, error) {
	return _ISynapseGasOracle.Contract.GetLocalGasData(&_ISynapseGasOracle.CallOpts)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracle *ISynapseGasOracleTransactor) ReceiveRemoteGasData(opts *bind.TransactOpts, srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracle.contract.Transact(opts, "receiveRemoteGasData", srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracle *ISynapseGasOracleSession) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.ReceiveRemoteGasData(&_ISynapseGasOracle.TransactOpts, srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracle *ISynapseGasOracleTransactorSession) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.ReceiveRemoteGasData(&_ISynapseGasOracle.TransactOpts, srcChainId, data)
}

// SynapseGasOracleMockMetaData contains all meta data concerning the SynapseGasOracleMock contract.
var SynapseGasOracleMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLocalGasData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveRemoteGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40658a74": "convertRemoteValueToLocalUnits(uint64,uint256)",
		"bf495c88": "estimateTxCostInLocalUnits(uint64,uint256,uint256)",
		"b376a688": "estimateTxCostInRemoteUnits(uint64,uint256,uint256)",
		"6f928aa7": "getLocalGasData()",
		"83389de7": "receiveRemoteGasData(uint64,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061026f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806383389de71161005057806383389de7146100a4578063b376a688146100b9578063bf495c88146100b957600080fd5b806340658a741461006c5780636f928aa714610095575b600080fd5b61008261007a3660046100ed565b600092915050565b6040519081526020015b60405180910390f35b606060405161008c9190610117565b6100b76100b2366004610183565b505050565b005b6100826100c7366004610206565b60009392505050565b803567ffffffffffffffff811681146100e857600080fd5b919050565b6000806040838503121561010057600080fd5b610109836100d0565b946020939093013593505050565b600060208083528351808285015260005b8181101561014457858101830151858201604001528201610128565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b60008060006040848603121561019857600080fd5b6101a1846100d0565b9250602084013567ffffffffffffffff808211156101be57600080fd5b818601915086601f8301126101d257600080fd5b8135818111156101e157600080fd5b8760208285010111156101f357600080fd5b6020830194508093505050509250925092565b60008060006060848603121561021b57600080fd5b610224846100d0565b9560208501359550604090940135939250505056fea264697066735822122031dbc394294ddb7ff922e1d852851eb8e27e036421dc325a5b4816f7eb19972c64736f6c63430008140033",
}

// SynapseGasOracleMockABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseGasOracleMockMetaData.ABI instead.
var SynapseGasOracleMockABI = SynapseGasOracleMockMetaData.ABI

// Deprecated: Use SynapseGasOracleMockMetaData.Sigs instead.
// SynapseGasOracleMockFuncSigs maps the 4-byte function signature to its string representation.
var SynapseGasOracleMockFuncSigs = SynapseGasOracleMockMetaData.Sigs

// SynapseGasOracleMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseGasOracleMockMetaData.Bin instead.
var SynapseGasOracleMockBin = SynapseGasOracleMockMetaData.Bin

// DeploySynapseGasOracleMock deploys a new Ethereum contract, binding an instance of SynapseGasOracleMock to it.
func DeploySynapseGasOracleMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseGasOracleMock, error) {
	parsed, err := SynapseGasOracleMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseGasOracleMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseGasOracleMock{SynapseGasOracleMockCaller: SynapseGasOracleMockCaller{contract: contract}, SynapseGasOracleMockTransactor: SynapseGasOracleMockTransactor{contract: contract}, SynapseGasOracleMockFilterer: SynapseGasOracleMockFilterer{contract: contract}}, nil
}

// SynapseGasOracleMock is an auto generated Go binding around an Ethereum contract.
type SynapseGasOracleMock struct {
	SynapseGasOracleMockCaller     // Read-only binding to the contract
	SynapseGasOracleMockTransactor // Write-only binding to the contract
	SynapseGasOracleMockFilterer   // Log filterer for contract events
}

// SynapseGasOracleMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseGasOracleMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseGasOracleMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseGasOracleMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseGasOracleMockSession struct {
	Contract     *SynapseGasOracleMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SynapseGasOracleMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseGasOracleMockCallerSession struct {
	Contract *SynapseGasOracleMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SynapseGasOracleMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseGasOracleMockTransactorSession struct {
	Contract     *SynapseGasOracleMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SynapseGasOracleMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseGasOracleMockRaw struct {
	Contract *SynapseGasOracleMock // Generic contract binding to access the raw methods on
}

// SynapseGasOracleMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseGasOracleMockCallerRaw struct {
	Contract *SynapseGasOracleMockCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseGasOracleMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseGasOracleMockTransactorRaw struct {
	Contract *SynapseGasOracleMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseGasOracleMock creates a new instance of SynapseGasOracleMock, bound to a specific deployed contract.
func NewSynapseGasOracleMock(address common.Address, backend bind.ContractBackend) (*SynapseGasOracleMock, error) {
	contract, err := bindSynapseGasOracleMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleMock{SynapseGasOracleMockCaller: SynapseGasOracleMockCaller{contract: contract}, SynapseGasOracleMockTransactor: SynapseGasOracleMockTransactor{contract: contract}, SynapseGasOracleMockFilterer: SynapseGasOracleMockFilterer{contract: contract}}, nil
}

// NewSynapseGasOracleMockCaller creates a new read-only instance of SynapseGasOracleMock, bound to a specific deployed contract.
func NewSynapseGasOracleMockCaller(address common.Address, caller bind.ContractCaller) (*SynapseGasOracleMockCaller, error) {
	contract, err := bindSynapseGasOracleMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleMockCaller{contract: contract}, nil
}

// NewSynapseGasOracleMockTransactor creates a new write-only instance of SynapseGasOracleMock, bound to a specific deployed contract.
func NewSynapseGasOracleMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseGasOracleMockTransactor, error) {
	contract, err := bindSynapseGasOracleMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleMockTransactor{contract: contract}, nil
}

// NewSynapseGasOracleMockFilterer creates a new log filterer instance of SynapseGasOracleMock, bound to a specific deployed contract.
func NewSynapseGasOracleMockFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseGasOracleMockFilterer, error) {
	contract, err := bindSynapseGasOracleMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleMockFilterer{contract: contract}, nil
}

// bindSynapseGasOracleMock binds a generic wrapper to an already deployed contract.
func bindSynapseGasOracleMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseGasOracleMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseGasOracleMock *SynapseGasOracleMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseGasOracleMock.Contract.SynapseGasOracleMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseGasOracleMock *SynapseGasOracleMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasOracleMock.Contract.SynapseGasOracleMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseGasOracleMock *SynapseGasOracleMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseGasOracleMock.Contract.SynapseGasOracleMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseGasOracleMock *SynapseGasOracleMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseGasOracleMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseGasOracleMock *SynapseGasOracleMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasOracleMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseGasOracleMock *SynapseGasOracleMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseGasOracleMock.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId uint64, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseGasOracleMock.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _SynapseGasOracleMock.Contract.ConvertRemoteValueToLocalUnits(&_SynapseGasOracleMock.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _SynapseGasOracleMock.Contract.ConvertRemoteValueToLocalUnits(&_SynapseGasOracleMock.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseGasOracleMock.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _SynapseGasOracleMock.Contract.EstimateTxCostInLocalUnits(&_SynapseGasOracleMock.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockCallerSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _SynapseGasOracleMock.Contract.EstimateTxCostInLocalUnits(&_SynapseGasOracleMock.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseGasOracleMock.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _SynapseGasOracleMock.Contract.EstimateTxCostInRemoteUnits(&_SynapseGasOracleMock.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleMock *SynapseGasOracleMockCallerSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _SynapseGasOracleMock.Contract.EstimateTxCostInRemoteUnits(&_SynapseGasOracleMock.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_SynapseGasOracleMock *SynapseGasOracleMockCaller) GetLocalGasData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SynapseGasOracleMock.contract.Call(opts, &out, "getLocalGasData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_SynapseGasOracleMock *SynapseGasOracleMockSession) GetLocalGasData() ([]byte, error) {
	return _SynapseGasOracleMock.Contract.GetLocalGasData(&_SynapseGasOracleMock.CallOpts)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_SynapseGasOracleMock *SynapseGasOracleMockCallerSession) GetLocalGasData() ([]byte, error) {
	return _SynapseGasOracleMock.Contract.GetLocalGasData(&_SynapseGasOracleMock.CallOpts)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_SynapseGasOracleMock *SynapseGasOracleMockTransactor) ReceiveRemoteGasData(opts *bind.TransactOpts, srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _SynapseGasOracleMock.contract.Transact(opts, "receiveRemoteGasData", srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_SynapseGasOracleMock *SynapseGasOracleMockSession) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _SynapseGasOracleMock.Contract.ReceiveRemoteGasData(&_SynapseGasOracleMock.TransactOpts, srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_SynapseGasOracleMock *SynapseGasOracleMockTransactorSession) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _SynapseGasOracleMock.Contract.ReceiveRemoteGasData(&_SynapseGasOracleMock.TransactOpts, srcChainId, data)
}
