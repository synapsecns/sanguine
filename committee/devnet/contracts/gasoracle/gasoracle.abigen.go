// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gasoracle

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

// ISynapseGasOracleV1RemoteGasData is an auto generated low-level Go binding around an user-defined struct.
type ISynapseGasOracleV1RemoteGasData struct {
	CalldataPrice *big.Int
	GasPrice      *big.Int
	NativePrice   *big.Int
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

// ISynapseGasOracleV1MetaData contains all meta data concerning the ISynapseGasOracleV1 contract.
var ISynapseGasOracleV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"SynapseGasOracleV1__ChainIdNotRemote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"SynapseGasOracleV1__NativePriceNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseGasOracleV1__NativePriceZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLocalGasData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLocalNativePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getRemoteGasData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"calldataPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"internalType\":\"structISynapseGasOracleV1.RemoteGasData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveRemoteGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"name\":\"setLocalNativePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"calldataPrice\",\"type\":\"uint256\"}],\"name\":\"setRemoteCallDataPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"calldataPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"internalType\":\"structISynapseGasOracleV1.RemoteGasData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"setRemoteGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"setRemoteGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"name\":\"setRemoteNativePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40658a74": "convertRemoteValueToLocalUnits(uint64,uint256)",
		"bf495c88": "estimateTxCostInLocalUnits(uint64,uint256,uint256)",
		"b376a688": "estimateTxCostInRemoteUnits(uint64,uint256,uint256)",
		"6f928aa7": "getLocalGasData()",
		"50fc83a6": "getLocalNativePrice()",
		"d0634296": "getRemoteGasData(uint64)",
		"83389de7": "receiveRemoteGasData(uint64,bytes)",
		"d1bdedaa": "setLocalNativePrice(uint256)",
		"5819a378": "setRemoteCallDataPrice(uint64,uint256)",
		"afe6c6b5": "setRemoteGasData(uint64,(uint256,uint256,uint256))",
		"fc4721ad": "setRemoteGasPrice(uint64,uint256)",
		"c91c6724": "setRemoteNativePrice(uint64,uint256)",
	},
}

// ISynapseGasOracleV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseGasOracleV1MetaData.ABI instead.
var ISynapseGasOracleV1ABI = ISynapseGasOracleV1MetaData.ABI

// Deprecated: Use ISynapseGasOracleV1MetaData.Sigs instead.
// ISynapseGasOracleV1FuncSigs maps the 4-byte function signature to its string representation.
var ISynapseGasOracleV1FuncSigs = ISynapseGasOracleV1MetaData.Sigs

// ISynapseGasOracleV1 is an auto generated Go binding around an Ethereum contract.
type ISynapseGasOracleV1 struct {
	ISynapseGasOracleV1Caller     // Read-only binding to the contract
	ISynapseGasOracleV1Transactor // Write-only binding to the contract
	ISynapseGasOracleV1Filterer   // Log filterer for contract events
}

// ISynapseGasOracleV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseGasOracleV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseGasOracleV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseGasOracleV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseGasOracleV1Session struct {
	Contract     *ISynapseGasOracleV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ISynapseGasOracleV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseGasOracleV1CallerSession struct {
	Contract *ISynapseGasOracleV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ISynapseGasOracleV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseGasOracleV1TransactorSession struct {
	Contract     *ISynapseGasOracleV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ISynapseGasOracleV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseGasOracleV1Raw struct {
	Contract *ISynapseGasOracleV1 // Generic contract binding to access the raw methods on
}

// ISynapseGasOracleV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseGasOracleV1CallerRaw struct {
	Contract *ISynapseGasOracleV1Caller // Generic read-only contract binding to access the raw methods on
}

// ISynapseGasOracleV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseGasOracleV1TransactorRaw struct {
	Contract *ISynapseGasOracleV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseGasOracleV1 creates a new instance of ISynapseGasOracleV1, bound to a specific deployed contract.
func NewISynapseGasOracleV1(address common.Address, backend bind.ContractBackend) (*ISynapseGasOracleV1, error) {
	contract, err := bindISynapseGasOracleV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleV1{ISynapseGasOracleV1Caller: ISynapseGasOracleV1Caller{contract: contract}, ISynapseGasOracleV1Transactor: ISynapseGasOracleV1Transactor{contract: contract}, ISynapseGasOracleV1Filterer: ISynapseGasOracleV1Filterer{contract: contract}}, nil
}

// NewISynapseGasOracleV1Caller creates a new read-only instance of ISynapseGasOracleV1, bound to a specific deployed contract.
func NewISynapseGasOracleV1Caller(address common.Address, caller bind.ContractCaller) (*ISynapseGasOracleV1Caller, error) {
	contract, err := bindISynapseGasOracleV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleV1Caller{contract: contract}, nil
}

// NewISynapseGasOracleV1Transactor creates a new write-only instance of ISynapseGasOracleV1, bound to a specific deployed contract.
func NewISynapseGasOracleV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseGasOracleV1Transactor, error) {
	contract, err := bindISynapseGasOracleV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleV1Transactor{contract: contract}, nil
}

// NewISynapseGasOracleV1Filterer creates a new log filterer instance of ISynapseGasOracleV1, bound to a specific deployed contract.
func NewISynapseGasOracleV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseGasOracleV1Filterer, error) {
	contract, err := bindISynapseGasOracleV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleV1Filterer{contract: contract}, nil
}

// bindISynapseGasOracleV1 binds a generic wrapper to an already deployed contract.
func bindISynapseGasOracleV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISynapseGasOracleV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseGasOracleV1.Contract.ISynapseGasOracleV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.ISynapseGasOracleV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.ISynapseGasOracleV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseGasOracleV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Caller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId uint64, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracleV1.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _ISynapseGasOracleV1.Contract.ConvertRemoteValueToLocalUnits(&_ISynapseGasOracleV1.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1CallerSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _ISynapseGasOracleV1.Contract.ConvertRemoteValueToLocalUnits(&_ISynapseGasOracleV1.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Caller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracleV1.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracleV1.Contract.EstimateTxCostInLocalUnits(&_ISynapseGasOracleV1.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1CallerSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracleV1.Contract.EstimateTxCostInLocalUnits(&_ISynapseGasOracleV1.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Caller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracleV1.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracleV1.Contract.EstimateTxCostInRemoteUnits(&_ISynapseGasOracleV1.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1CallerSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracleV1.Contract.EstimateTxCostInRemoteUnits(&_ISynapseGasOracleV1.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Caller) GetLocalGasData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ISynapseGasOracleV1.contract.Call(opts, &out, "getLocalGasData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) GetLocalGasData() ([]byte, error) {
	return _ISynapseGasOracleV1.Contract.GetLocalGasData(&_ISynapseGasOracleV1.CallOpts)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1CallerSession) GetLocalGasData() ([]byte, error) {
	return _ISynapseGasOracleV1.Contract.GetLocalGasData(&_ISynapseGasOracleV1.CallOpts)
}

// GetLocalNativePrice is a free data retrieval call binding the contract method 0x50fc83a6.
//
// Solidity: function getLocalNativePrice() view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Caller) GetLocalNativePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracleV1.contract.Call(opts, &out, "getLocalNativePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLocalNativePrice is a free data retrieval call binding the contract method 0x50fc83a6.
//
// Solidity: function getLocalNativePrice() view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) GetLocalNativePrice() (*big.Int, error) {
	return _ISynapseGasOracleV1.Contract.GetLocalNativePrice(&_ISynapseGasOracleV1.CallOpts)
}

// GetLocalNativePrice is a free data retrieval call binding the contract method 0x50fc83a6.
//
// Solidity: function getLocalNativePrice() view returns(uint256)
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1CallerSession) GetLocalNativePrice() (*big.Int, error) {
	return _ISynapseGasOracleV1.Contract.GetLocalNativePrice(&_ISynapseGasOracleV1.CallOpts)
}

// GetRemoteGasData is a free data retrieval call binding the contract method 0xd0634296.
//
// Solidity: function getRemoteGasData(uint64 chainId) view returns((uint256,uint256,uint256))
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Caller) GetRemoteGasData(opts *bind.CallOpts, chainId uint64) (ISynapseGasOracleV1RemoteGasData, error) {
	var out []interface{}
	err := _ISynapseGasOracleV1.contract.Call(opts, &out, "getRemoteGasData", chainId)

	if err != nil {
		return *new(ISynapseGasOracleV1RemoteGasData), err
	}

	out0 := *abi.ConvertType(out[0], new(ISynapseGasOracleV1RemoteGasData)).(*ISynapseGasOracleV1RemoteGasData)

	return out0, err

}

// GetRemoteGasData is a free data retrieval call binding the contract method 0xd0634296.
//
// Solidity: function getRemoteGasData(uint64 chainId) view returns((uint256,uint256,uint256))
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) GetRemoteGasData(chainId uint64) (ISynapseGasOracleV1RemoteGasData, error) {
	return _ISynapseGasOracleV1.Contract.GetRemoteGasData(&_ISynapseGasOracleV1.CallOpts, chainId)
}

// GetRemoteGasData is a free data retrieval call binding the contract method 0xd0634296.
//
// Solidity: function getRemoteGasData(uint64 chainId) view returns((uint256,uint256,uint256))
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1CallerSession) GetRemoteGasData(chainId uint64) (ISynapseGasOracleV1RemoteGasData, error) {
	return _ISynapseGasOracleV1.Contract.GetRemoteGasData(&_ISynapseGasOracleV1.CallOpts, chainId)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Transactor) ReceiveRemoteGasData(opts *bind.TransactOpts, srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.contract.Transact(opts, "receiveRemoteGasData", srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.ReceiveRemoteGasData(&_ISynapseGasOracleV1.TransactOpts, srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1TransactorSession) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.ReceiveRemoteGasData(&_ISynapseGasOracleV1.TransactOpts, srcChainId, data)
}

// SetLocalNativePrice is a paid mutator transaction binding the contract method 0xd1bdedaa.
//
// Solidity: function setLocalNativePrice(uint256 nativePrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Transactor) SetLocalNativePrice(opts *bind.TransactOpts, nativePrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.contract.Transact(opts, "setLocalNativePrice", nativePrice)
}

// SetLocalNativePrice is a paid mutator transaction binding the contract method 0xd1bdedaa.
//
// Solidity: function setLocalNativePrice(uint256 nativePrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) SetLocalNativePrice(nativePrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetLocalNativePrice(&_ISynapseGasOracleV1.TransactOpts, nativePrice)
}

// SetLocalNativePrice is a paid mutator transaction binding the contract method 0xd1bdedaa.
//
// Solidity: function setLocalNativePrice(uint256 nativePrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1TransactorSession) SetLocalNativePrice(nativePrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetLocalNativePrice(&_ISynapseGasOracleV1.TransactOpts, nativePrice)
}

// SetRemoteCallDataPrice is a paid mutator transaction binding the contract method 0x5819a378.
//
// Solidity: function setRemoteCallDataPrice(uint64 chainId, uint256 calldataPrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Transactor) SetRemoteCallDataPrice(opts *bind.TransactOpts, chainId uint64, calldataPrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.contract.Transact(opts, "setRemoteCallDataPrice", chainId, calldataPrice)
}

// SetRemoteCallDataPrice is a paid mutator transaction binding the contract method 0x5819a378.
//
// Solidity: function setRemoteCallDataPrice(uint64 chainId, uint256 calldataPrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) SetRemoteCallDataPrice(chainId uint64, calldataPrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetRemoteCallDataPrice(&_ISynapseGasOracleV1.TransactOpts, chainId, calldataPrice)
}

// SetRemoteCallDataPrice is a paid mutator transaction binding the contract method 0x5819a378.
//
// Solidity: function setRemoteCallDataPrice(uint64 chainId, uint256 calldataPrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1TransactorSession) SetRemoteCallDataPrice(chainId uint64, calldataPrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetRemoteCallDataPrice(&_ISynapseGasOracleV1.TransactOpts, chainId, calldataPrice)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xafe6c6b5.
//
// Solidity: function setRemoteGasData(uint64 chainId, (uint256,uint256,uint256) data) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Transactor) SetRemoteGasData(opts *bind.TransactOpts, chainId uint64, data ISynapseGasOracleV1RemoteGasData) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.contract.Transact(opts, "setRemoteGasData", chainId, data)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xafe6c6b5.
//
// Solidity: function setRemoteGasData(uint64 chainId, (uint256,uint256,uint256) data) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) SetRemoteGasData(chainId uint64, data ISynapseGasOracleV1RemoteGasData) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetRemoteGasData(&_ISynapseGasOracleV1.TransactOpts, chainId, data)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xafe6c6b5.
//
// Solidity: function setRemoteGasData(uint64 chainId, (uint256,uint256,uint256) data) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1TransactorSession) SetRemoteGasData(chainId uint64, data ISynapseGasOracleV1RemoteGasData) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetRemoteGasData(&_ISynapseGasOracleV1.TransactOpts, chainId, data)
}

// SetRemoteGasPrice is a paid mutator transaction binding the contract method 0xfc4721ad.
//
// Solidity: function setRemoteGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Transactor) SetRemoteGasPrice(opts *bind.TransactOpts, chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.contract.Transact(opts, "setRemoteGasPrice", chainId, gasPrice)
}

// SetRemoteGasPrice is a paid mutator transaction binding the contract method 0xfc4721ad.
//
// Solidity: function setRemoteGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) SetRemoteGasPrice(chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetRemoteGasPrice(&_ISynapseGasOracleV1.TransactOpts, chainId, gasPrice)
}

// SetRemoteGasPrice is a paid mutator transaction binding the contract method 0xfc4721ad.
//
// Solidity: function setRemoteGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1TransactorSession) SetRemoteGasPrice(chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetRemoteGasPrice(&_ISynapseGasOracleV1.TransactOpts, chainId, gasPrice)
}

// SetRemoteNativePrice is a paid mutator transaction binding the contract method 0xc91c6724.
//
// Solidity: function setRemoteNativePrice(uint64 chainId, uint256 nativePrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Transactor) SetRemoteNativePrice(opts *bind.TransactOpts, chainId uint64, nativePrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.contract.Transact(opts, "setRemoteNativePrice", chainId, nativePrice)
}

// SetRemoteNativePrice is a paid mutator transaction binding the contract method 0xc91c6724.
//
// Solidity: function setRemoteNativePrice(uint64 chainId, uint256 nativePrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1Session) SetRemoteNativePrice(chainId uint64, nativePrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetRemoteNativePrice(&_ISynapseGasOracleV1.TransactOpts, chainId, nativePrice)
}

// SetRemoteNativePrice is a paid mutator transaction binding the contract method 0xc91c6724.
//
// Solidity: function setRemoteNativePrice(uint64 chainId, uint256 nativePrice) returns()
func (_ISynapseGasOracleV1 *ISynapseGasOracleV1TransactorSession) SetRemoteNativePrice(chainId uint64, nativePrice *big.Int) (*types.Transaction, error) {
	return _ISynapseGasOracleV1.Contract.SetRemoteNativePrice(&_ISynapseGasOracleV1.TransactOpts, chainId, nativePrice)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntToUint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintToInt\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220611545909c1073bb2bcd4fbf787621a7fe6c780fb6b2dde246ad59dfe69b721f64736f6c63430008140033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// SynapseGasOracleV1MetaData contains all meta data concerning the SynapseGasOracleV1 contract.
var SynapseGasOracleV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"SynapseGasOracleV1__ChainIdNotRemote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"SynapseGasOracleV1__NativePriceNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseGasOracleV1__NativePriceZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"calldataPrice\",\"type\":\"uint256\"}],\"name\":\"CalldataPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"GasPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"name\":\"NativePriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLocalGasData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLocalNativePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getRemoteGasData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"calldataPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"internalType\":\"structISynapseGasOracleV1.RemoteGasData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveRemoteGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"name\":\"setLocalNativePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"calldataPrice\",\"type\":\"uint256\"}],\"name\":\"setRemoteCallDataPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"calldataPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"internalType\":\"structISynapseGasOracleV1.RemoteGasData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"setRemoteGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"setRemoteGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"name\":\"setRemoteNativePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40658a74": "convertRemoteValueToLocalUnits(uint64,uint256)",
		"bf495c88": "estimateTxCostInLocalUnits(uint64,uint256,uint256)",
		"b376a688": "estimateTxCostInRemoteUnits(uint64,uint256,uint256)",
		"6f928aa7": "getLocalGasData()",
		"50fc83a6": "getLocalNativePrice()",
		"d0634296": "getRemoteGasData(uint64)",
		"8da5cb5b": "owner()",
		"83389de7": "receiveRemoteGasData(uint64,bytes)",
		"715018a6": "renounceOwnership()",
		"d1bdedaa": "setLocalNativePrice(uint256)",
		"5819a378": "setRemoteCallDataPrice(uint64,uint256)",
		"afe6c6b5": "setRemoteGasData(uint64,(uint256,uint256,uint256))",
		"fc4721ad": "setRemoteGasPrice(uint64,uint256)",
		"c91c6724": "setRemoteNativePrice(uint64,uint256)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161113c38038061113c83398101604081905261002f916100be565b806001600160a01b03811661005e57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6100678161006e565b50506100ee565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100d057600080fd5b81516001600160a01b03811681146100e757600080fd5b9392505050565b61103f806100fd6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063afe6c6b511610097578063d063429611610066578063d0634296146101db578063d1bdedaa14610210578063f2fde38b14610223578063fc4721ad1461023657600080fd5b8063afe6c6b51461018f578063b376a688146101a2578063bf495c88146101b5578063c91c6724146101c857600080fd5b80636f928aa7116100d35780636f928aa71461013d578063715018a61461014c57806383389de7146101545780638da5cb5b1461016757600080fd5b806340658a74146100fa57806350fc83a6146101205780635819a37814610128575b600080fd5b61010d610108366004610cf3565b610249565b6040519081526020015b60405180910390f35b60015461010d565b61013b610136366004610cf3565b610315565b005b60606040516101179190610d1d565b61013b6103df565b61013b610162366004610d89565b505050565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610117565b61013b61019d366004610e0c565b6103f3565b61010d6101b0366004610ed1565b6104b4565b61010d6101c3366004610ed1565b61057d565b61013b6101d6366004610cf3565b610655565b6101ee6101e9366004610f04565b6106f2565b6040805182518152602080840151908201529181015190820152606001610117565b61013b61021e366004610f1f565b6107aa565b61013b610231366004610f38565b61084c565b61013b610244366004610cf3565b6108b0565b6000828067ffffffffffffffff1646036102a0576040517fbcaca17400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024015b60405180910390fd5b67ffffffffffffffff8416600090815260026020819052604082200154859103610302576040517f856e7cf000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b61030c8585610974565b95945050505050565b61031d610a06565b818067ffffffffffffffff16460361036d576040517fbcaca17400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b67ffffffffffffffff83166000908152600260208190526040822001548491036103cf576040517f856e7cf000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b6103d98484610a59565b50505050565b6103e7610a06565b6103f16000610ac7565b565b6103fb610a06565b818067ffffffffffffffff16460361044b576040517fbcaca17400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b81604001518060000361048a576040517f8e2d4f4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610498848460000151610a59565b6104a6848460200151610b3c565b6103d9848460400151610bb0565b6000838067ffffffffffffffff164603610506576040517fbcaca17400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b67ffffffffffffffff8516600090815260026020819052604082200154869103610568576040517f856e7cf000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b610573868686610c25565b9695505050505050565b6000838067ffffffffffffffff1646036105cf576040517fbcaca17400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b67ffffffffffffffff8516600090815260026020819052604082200154869103610631576040517f856e7cf000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b600061063e878787610c25565b905061064a8782610974565b979650505050505050565b61065d610a06565b818067ffffffffffffffff1646036106ad576040517fbcaca17400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b81806000036106e8576040517f8e2d4f4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103d98484610bb0565b61071660405180606001604052806000815260200160008152602001600081525090565b818067ffffffffffffffff164603610766576040517fbcaca17400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b505067ffffffffffffffff16600090815260026020818152604092839020835160608101855281548152600182015492810192909252909101549181019190915290565b6107b2610a06565b80806000036107ed576040517f8e2d4f4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816001541461084857600182905561080446610c82565b67ffffffffffffffff167f01e22a2f804ee98d53c44fad5dd2f7bd9dafc2737327f816f5330d66103f6dfd8360405161083f91815260200190565b60405180910390a25b5050565b610854610a06565b73ffffffffffffffffffffffffffffffffffffffff81166108a4576040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260006004820152602401610297565b6108ad81610ac7565b50565b6108b8610a06565b818067ffffffffffffffff164603610908576040517fbcaca17400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b67ffffffffffffffff831660009081526002602081905260408220015484910361096a576040517f856e7cf000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610297565b6103d98484610b3c565b60006001546000036109c85761098946610c82565b6040517f856e7cf000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610297565b60015467ffffffffffffffff8416600090815260026020819052604090912001546109f39084610fa4565b6109fd9190610fbb565b90505b92915050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103f1576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610297565b67ffffffffffffffff821660009081526002602052604090205481146108485767ffffffffffffffff821660008181526002602052604090819020839055517f3f67e101936b3cb094183a2ed4bf880810a36909c55c27a7d145ee4bbf3dd5159061083f9084815260200190565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b67ffffffffffffffff821660009081526002602052604090206001015481146108485767ffffffffffffffff821660008181526002602052604090819020600101839055517f3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff9061083f9084815260200190565b67ffffffffffffffff82166000908152600260208190526040909120015481146108485767ffffffffffffffff821660008181526002602081905260409182902001839055517f01e22a2f804ee98d53c44fad5dd2f7bd9dafc2737327f816f5330d66103f6dfd9061083f9084815260200190565b67ffffffffffffffff8316600090815260026020526040812054610c499083610fa4565b67ffffffffffffffff8516600090815260026020526040902060010154610c709085610fa4565b610c7a9190610ff6565b949350505050565b600067ffffffffffffffff821115610cd257604080517f6dfcc650000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610297565b5090565b803567ffffffffffffffff81168114610cee57600080fd5b919050565b60008060408385031215610d0657600080fd5b610d0f83610cd6565b946020939093013593505050565b600060208083528351808285015260005b81811015610d4a57858101830151858201604001528201610d2e565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b600080600060408486031215610d9e57600080fd5b610da784610cd6565b9250602084013567ffffffffffffffff80821115610dc457600080fd5b818601915086601f830112610dd857600080fd5b813581811115610de757600080fd5b876020828501011115610df957600080fd5b6020830194508093505050509250925092565b6000808284036080811215610e2057600080fd5b610e2984610cd6565b925060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082011215610e5b57600080fd5b506040516060810181811067ffffffffffffffff82111715610ea6577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8060405250602084013581526040840135602082015260608401356040820152809150509250929050565b600080600060608486031215610ee657600080fd5b610eef84610cd6565b95602085013595506040909401359392505050565b600060208284031215610f1657600080fd5b6109fd82610cd6565b600060208284031215610f3157600080fd5b5035919050565b600060208284031215610f4a57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610f6e57600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417610a0057610a00610f75565b600082610ff1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b80820180821115610a0057610a00610f7556fea26469706673582212203b20a6344ca372570cdde1ec02c77c4383a0d8967330f9b39acd857ec105009664736f6c63430008140033",
}

// SynapseGasOracleV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseGasOracleV1MetaData.ABI instead.
var SynapseGasOracleV1ABI = SynapseGasOracleV1MetaData.ABI

// Deprecated: Use SynapseGasOracleV1MetaData.Sigs instead.
// SynapseGasOracleV1FuncSigs maps the 4-byte function signature to its string representation.
var SynapseGasOracleV1FuncSigs = SynapseGasOracleV1MetaData.Sigs

// SynapseGasOracleV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseGasOracleV1MetaData.Bin instead.
var SynapseGasOracleV1Bin = SynapseGasOracleV1MetaData.Bin

// DeploySynapseGasOracleV1 deploys a new Ethereum contract, binding an instance of SynapseGasOracleV1 to it.
func DeploySynapseGasOracleV1(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address) (common.Address, *types.Transaction, *SynapseGasOracleV1, error) {
	parsed, err := SynapseGasOracleV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseGasOracleV1Bin), backend, owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseGasOracleV1{SynapseGasOracleV1Caller: SynapseGasOracleV1Caller{contract: contract}, SynapseGasOracleV1Transactor: SynapseGasOracleV1Transactor{contract: contract}, SynapseGasOracleV1Filterer: SynapseGasOracleV1Filterer{contract: contract}}, nil
}

// SynapseGasOracleV1 is an auto generated Go binding around an Ethereum contract.
type SynapseGasOracleV1 struct {
	SynapseGasOracleV1Caller     // Read-only binding to the contract
	SynapseGasOracleV1Transactor // Write-only binding to the contract
	SynapseGasOracleV1Filterer   // Log filterer for contract events
}

// SynapseGasOracleV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseGasOracleV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseGasOracleV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseGasOracleV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseGasOracleV1Session struct {
	Contract     *SynapseGasOracleV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SynapseGasOracleV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseGasOracleV1CallerSession struct {
	Contract *SynapseGasOracleV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SynapseGasOracleV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseGasOracleV1TransactorSession struct {
	Contract     *SynapseGasOracleV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SynapseGasOracleV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseGasOracleV1Raw struct {
	Contract *SynapseGasOracleV1 // Generic contract binding to access the raw methods on
}

// SynapseGasOracleV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseGasOracleV1CallerRaw struct {
	Contract *SynapseGasOracleV1Caller // Generic read-only contract binding to access the raw methods on
}

// SynapseGasOracleV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseGasOracleV1TransactorRaw struct {
	Contract *SynapseGasOracleV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseGasOracleV1 creates a new instance of SynapseGasOracleV1, bound to a specific deployed contract.
func NewSynapseGasOracleV1(address common.Address, backend bind.ContractBackend) (*SynapseGasOracleV1, error) {
	contract, err := bindSynapseGasOracleV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1{SynapseGasOracleV1Caller: SynapseGasOracleV1Caller{contract: contract}, SynapseGasOracleV1Transactor: SynapseGasOracleV1Transactor{contract: contract}, SynapseGasOracleV1Filterer: SynapseGasOracleV1Filterer{contract: contract}}, nil
}

// NewSynapseGasOracleV1Caller creates a new read-only instance of SynapseGasOracleV1, bound to a specific deployed contract.
func NewSynapseGasOracleV1Caller(address common.Address, caller bind.ContractCaller) (*SynapseGasOracleV1Caller, error) {
	contract, err := bindSynapseGasOracleV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1Caller{contract: contract}, nil
}

// NewSynapseGasOracleV1Transactor creates a new write-only instance of SynapseGasOracleV1, bound to a specific deployed contract.
func NewSynapseGasOracleV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SynapseGasOracleV1Transactor, error) {
	contract, err := bindSynapseGasOracleV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1Transactor{contract: contract}, nil
}

// NewSynapseGasOracleV1Filterer creates a new log filterer instance of SynapseGasOracleV1, bound to a specific deployed contract.
func NewSynapseGasOracleV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SynapseGasOracleV1Filterer, error) {
	contract, err := bindSynapseGasOracleV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1Filterer{contract: contract}, nil
}

// bindSynapseGasOracleV1 binds a generic wrapper to an already deployed contract.
func bindSynapseGasOracleV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseGasOracleV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseGasOracleV1 *SynapseGasOracleV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseGasOracleV1.Contract.SynapseGasOracleV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseGasOracleV1 *SynapseGasOracleV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SynapseGasOracleV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseGasOracleV1 *SynapseGasOracleV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SynapseGasOracleV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseGasOracleV1 *SynapseGasOracleV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseGasOracleV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Caller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId uint64, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseGasOracleV1.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _SynapseGasOracleV1.Contract.ConvertRemoteValueToLocalUnits(&_SynapseGasOracleV1.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1CallerSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _SynapseGasOracleV1.Contract.ConvertRemoteValueToLocalUnits(&_SynapseGasOracleV1.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Caller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseGasOracleV1.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _SynapseGasOracleV1.Contract.EstimateTxCostInLocalUnits(&_SynapseGasOracleV1.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1CallerSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _SynapseGasOracleV1.Contract.EstimateTxCostInLocalUnits(&_SynapseGasOracleV1.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Caller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseGasOracleV1.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _SynapseGasOracleV1.Contract.EstimateTxCostInRemoteUnits(&_SynapseGasOracleV1.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1CallerSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _SynapseGasOracleV1.Contract.EstimateTxCostInRemoteUnits(&_SynapseGasOracleV1.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Caller) GetLocalGasData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SynapseGasOracleV1.contract.Call(opts, &out, "getLocalGasData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) GetLocalGasData() ([]byte, error) {
	return _SynapseGasOracleV1.Contract.GetLocalGasData(&_SynapseGasOracleV1.CallOpts)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_SynapseGasOracleV1 *SynapseGasOracleV1CallerSession) GetLocalGasData() ([]byte, error) {
	return _SynapseGasOracleV1.Contract.GetLocalGasData(&_SynapseGasOracleV1.CallOpts)
}

// GetLocalNativePrice is a free data retrieval call binding the contract method 0x50fc83a6.
//
// Solidity: function getLocalNativePrice() view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Caller) GetLocalNativePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseGasOracleV1.contract.Call(opts, &out, "getLocalNativePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLocalNativePrice is a free data retrieval call binding the contract method 0x50fc83a6.
//
// Solidity: function getLocalNativePrice() view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) GetLocalNativePrice() (*big.Int, error) {
	return _SynapseGasOracleV1.Contract.GetLocalNativePrice(&_SynapseGasOracleV1.CallOpts)
}

// GetLocalNativePrice is a free data retrieval call binding the contract method 0x50fc83a6.
//
// Solidity: function getLocalNativePrice() view returns(uint256)
func (_SynapseGasOracleV1 *SynapseGasOracleV1CallerSession) GetLocalNativePrice() (*big.Int, error) {
	return _SynapseGasOracleV1.Contract.GetLocalNativePrice(&_SynapseGasOracleV1.CallOpts)
}

// GetRemoteGasData is a free data retrieval call binding the contract method 0xd0634296.
//
// Solidity: function getRemoteGasData(uint64 chainId) view returns((uint256,uint256,uint256))
func (_SynapseGasOracleV1 *SynapseGasOracleV1Caller) GetRemoteGasData(opts *bind.CallOpts, chainId uint64) (ISynapseGasOracleV1RemoteGasData, error) {
	var out []interface{}
	err := _SynapseGasOracleV1.contract.Call(opts, &out, "getRemoteGasData", chainId)

	if err != nil {
		return *new(ISynapseGasOracleV1RemoteGasData), err
	}

	out0 := *abi.ConvertType(out[0], new(ISynapseGasOracleV1RemoteGasData)).(*ISynapseGasOracleV1RemoteGasData)

	return out0, err

}

// GetRemoteGasData is a free data retrieval call binding the contract method 0xd0634296.
//
// Solidity: function getRemoteGasData(uint64 chainId) view returns((uint256,uint256,uint256))
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) GetRemoteGasData(chainId uint64) (ISynapseGasOracleV1RemoteGasData, error) {
	return _SynapseGasOracleV1.Contract.GetRemoteGasData(&_SynapseGasOracleV1.CallOpts, chainId)
}

// GetRemoteGasData is a free data retrieval call binding the contract method 0xd0634296.
//
// Solidity: function getRemoteGasData(uint64 chainId) view returns((uint256,uint256,uint256))
func (_SynapseGasOracleV1 *SynapseGasOracleV1CallerSession) GetRemoteGasData(chainId uint64) (ISynapseGasOracleV1RemoteGasData, error) {
	return _SynapseGasOracleV1.Contract.GetRemoteGasData(&_SynapseGasOracleV1.CallOpts, chainId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseGasOracleV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) Owner() (common.Address, error) {
	return _SynapseGasOracleV1.Contract.Owner(&_SynapseGasOracleV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseGasOracleV1 *SynapseGasOracleV1CallerSession) Owner() (common.Address, error) {
	return _SynapseGasOracleV1.Contract.Owner(&_SynapseGasOracleV1.CallOpts)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Transactor) ReceiveRemoteGasData(opts *bind.TransactOpts, srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _SynapseGasOracleV1.contract.Transact(opts, "receiveRemoteGasData", srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.ReceiveRemoteGasData(&_SynapseGasOracleV1.TransactOpts, srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorSession) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.ReceiveRemoteGasData(&_SynapseGasOracleV1.TransactOpts, srcChainId, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasOracleV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.RenounceOwnership(&_SynapseGasOracleV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.RenounceOwnership(&_SynapseGasOracleV1.TransactOpts)
}

// SetLocalNativePrice is a paid mutator transaction binding the contract method 0xd1bdedaa.
//
// Solidity: function setLocalNativePrice(uint256 nativePrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Transactor) SetLocalNativePrice(opts *bind.TransactOpts, nativePrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.contract.Transact(opts, "setLocalNativePrice", nativePrice)
}

// SetLocalNativePrice is a paid mutator transaction binding the contract method 0xd1bdedaa.
//
// Solidity: function setLocalNativePrice(uint256 nativePrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) SetLocalNativePrice(nativePrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetLocalNativePrice(&_SynapseGasOracleV1.TransactOpts, nativePrice)
}

// SetLocalNativePrice is a paid mutator transaction binding the contract method 0xd1bdedaa.
//
// Solidity: function setLocalNativePrice(uint256 nativePrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorSession) SetLocalNativePrice(nativePrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetLocalNativePrice(&_SynapseGasOracleV1.TransactOpts, nativePrice)
}

// SetRemoteCallDataPrice is a paid mutator transaction binding the contract method 0x5819a378.
//
// Solidity: function setRemoteCallDataPrice(uint64 chainId, uint256 calldataPrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Transactor) SetRemoteCallDataPrice(opts *bind.TransactOpts, chainId uint64, calldataPrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.contract.Transact(opts, "setRemoteCallDataPrice", chainId, calldataPrice)
}

// SetRemoteCallDataPrice is a paid mutator transaction binding the contract method 0x5819a378.
//
// Solidity: function setRemoteCallDataPrice(uint64 chainId, uint256 calldataPrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) SetRemoteCallDataPrice(chainId uint64, calldataPrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetRemoteCallDataPrice(&_SynapseGasOracleV1.TransactOpts, chainId, calldataPrice)
}

// SetRemoteCallDataPrice is a paid mutator transaction binding the contract method 0x5819a378.
//
// Solidity: function setRemoteCallDataPrice(uint64 chainId, uint256 calldataPrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorSession) SetRemoteCallDataPrice(chainId uint64, calldataPrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetRemoteCallDataPrice(&_SynapseGasOracleV1.TransactOpts, chainId, calldataPrice)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xafe6c6b5.
//
// Solidity: function setRemoteGasData(uint64 chainId, (uint256,uint256,uint256) data) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Transactor) SetRemoteGasData(opts *bind.TransactOpts, chainId uint64, data ISynapseGasOracleV1RemoteGasData) (*types.Transaction, error) {
	return _SynapseGasOracleV1.contract.Transact(opts, "setRemoteGasData", chainId, data)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xafe6c6b5.
//
// Solidity: function setRemoteGasData(uint64 chainId, (uint256,uint256,uint256) data) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) SetRemoteGasData(chainId uint64, data ISynapseGasOracleV1RemoteGasData) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetRemoteGasData(&_SynapseGasOracleV1.TransactOpts, chainId, data)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xafe6c6b5.
//
// Solidity: function setRemoteGasData(uint64 chainId, (uint256,uint256,uint256) data) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorSession) SetRemoteGasData(chainId uint64, data ISynapseGasOracleV1RemoteGasData) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetRemoteGasData(&_SynapseGasOracleV1.TransactOpts, chainId, data)
}

// SetRemoteGasPrice is a paid mutator transaction binding the contract method 0xfc4721ad.
//
// Solidity: function setRemoteGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Transactor) SetRemoteGasPrice(opts *bind.TransactOpts, chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.contract.Transact(opts, "setRemoteGasPrice", chainId, gasPrice)
}

// SetRemoteGasPrice is a paid mutator transaction binding the contract method 0xfc4721ad.
//
// Solidity: function setRemoteGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) SetRemoteGasPrice(chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetRemoteGasPrice(&_SynapseGasOracleV1.TransactOpts, chainId, gasPrice)
}

// SetRemoteGasPrice is a paid mutator transaction binding the contract method 0xfc4721ad.
//
// Solidity: function setRemoteGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorSession) SetRemoteGasPrice(chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetRemoteGasPrice(&_SynapseGasOracleV1.TransactOpts, chainId, gasPrice)
}

// SetRemoteNativePrice is a paid mutator transaction binding the contract method 0xc91c6724.
//
// Solidity: function setRemoteNativePrice(uint64 chainId, uint256 nativePrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Transactor) SetRemoteNativePrice(opts *bind.TransactOpts, chainId uint64, nativePrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.contract.Transact(opts, "setRemoteNativePrice", chainId, nativePrice)
}

// SetRemoteNativePrice is a paid mutator transaction binding the contract method 0xc91c6724.
//
// Solidity: function setRemoteNativePrice(uint64 chainId, uint256 nativePrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) SetRemoteNativePrice(chainId uint64, nativePrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetRemoteNativePrice(&_SynapseGasOracleV1.TransactOpts, chainId, nativePrice)
}

// SetRemoteNativePrice is a paid mutator transaction binding the contract method 0xc91c6724.
//
// Solidity: function setRemoteNativePrice(uint64 chainId, uint256 nativePrice) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorSession) SetRemoteNativePrice(chainId uint64, nativePrice *big.Int) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.SetRemoteNativePrice(&_SynapseGasOracleV1.TransactOpts, chainId, nativePrice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynapseGasOracleV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.TransferOwnership(&_SynapseGasOracleV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseGasOracleV1 *SynapseGasOracleV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseGasOracleV1.Contract.TransferOwnership(&_SynapseGasOracleV1.TransactOpts, newOwner)
}

// SynapseGasOracleV1CalldataPriceSetIterator is returned from FilterCalldataPriceSet and is used to iterate over the raw logs and unpacked data for CalldataPriceSet events raised by the SynapseGasOracleV1 contract.
type SynapseGasOracleV1CalldataPriceSetIterator struct {
	Event *SynapseGasOracleV1CalldataPriceSet // Event containing the contract specifics and raw log

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
func (it *SynapseGasOracleV1CalldataPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseGasOracleV1CalldataPriceSet)
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
		it.Event = new(SynapseGasOracleV1CalldataPriceSet)
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
func (it *SynapseGasOracleV1CalldataPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseGasOracleV1CalldataPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseGasOracleV1CalldataPriceSet represents a CalldataPriceSet event raised by the SynapseGasOracleV1 contract.
type SynapseGasOracleV1CalldataPriceSet struct {
	ChainId       uint64
	CalldataPrice *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalldataPriceSet is a free log retrieval operation binding the contract event 0x3f67e101936b3cb094183a2ed4bf880810a36909c55c27a7d145ee4bbf3dd515.
//
// Solidity: event CalldataPriceSet(uint64 indexed chainId, uint256 calldataPrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) FilterCalldataPriceSet(opts *bind.FilterOpts, chainId []uint64) (*SynapseGasOracleV1CalldataPriceSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1.contract.FilterLogs(opts, "CalldataPriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1CalldataPriceSetIterator{contract: _SynapseGasOracleV1.contract, event: "CalldataPriceSet", logs: logs, sub: sub}, nil
}

// WatchCalldataPriceSet is a free log subscription operation binding the contract event 0x3f67e101936b3cb094183a2ed4bf880810a36909c55c27a7d145ee4bbf3dd515.
//
// Solidity: event CalldataPriceSet(uint64 indexed chainId, uint256 calldataPrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) WatchCalldataPriceSet(opts *bind.WatchOpts, sink chan<- *SynapseGasOracleV1CalldataPriceSet, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1.contract.WatchLogs(opts, "CalldataPriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseGasOracleV1CalldataPriceSet)
				if err := _SynapseGasOracleV1.contract.UnpackLog(event, "CalldataPriceSet", log); err != nil {
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

// ParseCalldataPriceSet is a log parse operation binding the contract event 0x3f67e101936b3cb094183a2ed4bf880810a36909c55c27a7d145ee4bbf3dd515.
//
// Solidity: event CalldataPriceSet(uint64 indexed chainId, uint256 calldataPrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) ParseCalldataPriceSet(log types.Log) (*SynapseGasOracleV1CalldataPriceSet, error) {
	event := new(SynapseGasOracleV1CalldataPriceSet)
	if err := _SynapseGasOracleV1.contract.UnpackLog(event, "CalldataPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseGasOracleV1GasPriceSetIterator is returned from FilterGasPriceSet and is used to iterate over the raw logs and unpacked data for GasPriceSet events raised by the SynapseGasOracleV1 contract.
type SynapseGasOracleV1GasPriceSetIterator struct {
	Event *SynapseGasOracleV1GasPriceSet // Event containing the contract specifics and raw log

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
func (it *SynapseGasOracleV1GasPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseGasOracleV1GasPriceSet)
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
		it.Event = new(SynapseGasOracleV1GasPriceSet)
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
func (it *SynapseGasOracleV1GasPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseGasOracleV1GasPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseGasOracleV1GasPriceSet represents a GasPriceSet event raised by the SynapseGasOracleV1 contract.
type SynapseGasOracleV1GasPriceSet struct {
	ChainId  uint64
	GasPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasPriceSet is a free log retrieval operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 indexed chainId, uint256 gasPrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) FilterGasPriceSet(opts *bind.FilterOpts, chainId []uint64) (*SynapseGasOracleV1GasPriceSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1.contract.FilterLogs(opts, "GasPriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1GasPriceSetIterator{contract: _SynapseGasOracleV1.contract, event: "GasPriceSet", logs: logs, sub: sub}, nil
}

// WatchGasPriceSet is a free log subscription operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 indexed chainId, uint256 gasPrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) WatchGasPriceSet(opts *bind.WatchOpts, sink chan<- *SynapseGasOracleV1GasPriceSet, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1.contract.WatchLogs(opts, "GasPriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseGasOracleV1GasPriceSet)
				if err := _SynapseGasOracleV1.contract.UnpackLog(event, "GasPriceSet", log); err != nil {
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

// ParseGasPriceSet is a log parse operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 indexed chainId, uint256 gasPrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) ParseGasPriceSet(log types.Log) (*SynapseGasOracleV1GasPriceSet, error) {
	event := new(SynapseGasOracleV1GasPriceSet)
	if err := _SynapseGasOracleV1.contract.UnpackLog(event, "GasPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseGasOracleV1NativePriceSetIterator is returned from FilterNativePriceSet and is used to iterate over the raw logs and unpacked data for NativePriceSet events raised by the SynapseGasOracleV1 contract.
type SynapseGasOracleV1NativePriceSetIterator struct {
	Event *SynapseGasOracleV1NativePriceSet // Event containing the contract specifics and raw log

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
func (it *SynapseGasOracleV1NativePriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseGasOracleV1NativePriceSet)
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
		it.Event = new(SynapseGasOracleV1NativePriceSet)
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
func (it *SynapseGasOracleV1NativePriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseGasOracleV1NativePriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseGasOracleV1NativePriceSet represents a NativePriceSet event raised by the SynapseGasOracleV1 contract.
type SynapseGasOracleV1NativePriceSet struct {
	ChainId     uint64
	NativePrice *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNativePriceSet is a free log retrieval operation binding the contract event 0x01e22a2f804ee98d53c44fad5dd2f7bd9dafc2737327f816f5330d66103f6dfd.
//
// Solidity: event NativePriceSet(uint64 indexed chainId, uint256 nativePrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) FilterNativePriceSet(opts *bind.FilterOpts, chainId []uint64) (*SynapseGasOracleV1NativePriceSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1.contract.FilterLogs(opts, "NativePriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1NativePriceSetIterator{contract: _SynapseGasOracleV1.contract, event: "NativePriceSet", logs: logs, sub: sub}, nil
}

// WatchNativePriceSet is a free log subscription operation binding the contract event 0x01e22a2f804ee98d53c44fad5dd2f7bd9dafc2737327f816f5330d66103f6dfd.
//
// Solidity: event NativePriceSet(uint64 indexed chainId, uint256 nativePrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) WatchNativePriceSet(opts *bind.WatchOpts, sink chan<- *SynapseGasOracleV1NativePriceSet, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1.contract.WatchLogs(opts, "NativePriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseGasOracleV1NativePriceSet)
				if err := _SynapseGasOracleV1.contract.UnpackLog(event, "NativePriceSet", log); err != nil {
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

// ParseNativePriceSet is a log parse operation binding the contract event 0x01e22a2f804ee98d53c44fad5dd2f7bd9dafc2737327f816f5330d66103f6dfd.
//
// Solidity: event NativePriceSet(uint64 indexed chainId, uint256 nativePrice)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) ParseNativePriceSet(log types.Log) (*SynapseGasOracleV1NativePriceSet, error) {
	event := new(SynapseGasOracleV1NativePriceSet)
	if err := _SynapseGasOracleV1.contract.UnpackLog(event, "NativePriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseGasOracleV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynapseGasOracleV1 contract.
type SynapseGasOracleV1OwnershipTransferredIterator struct {
	Event *SynapseGasOracleV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SynapseGasOracleV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseGasOracleV1OwnershipTransferred)
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
		it.Event = new(SynapseGasOracleV1OwnershipTransferred)
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
func (it *SynapseGasOracleV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseGasOracleV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseGasOracleV1OwnershipTransferred represents a OwnershipTransferred event raised by the SynapseGasOracleV1 contract.
type SynapseGasOracleV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynapseGasOracleV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseGasOracleV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1OwnershipTransferredIterator{contract: _SynapseGasOracleV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynapseGasOracleV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseGasOracleV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseGasOracleV1OwnershipTransferred)
				if err := _SynapseGasOracleV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SynapseGasOracleV1 *SynapseGasOracleV1Filterer) ParseOwnershipTransferred(log types.Log) (*SynapseGasOracleV1OwnershipTransferred, error) {
	event := new(SynapseGasOracleV1OwnershipTransferred)
	if err := _SynapseGasOracleV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseGasOracleV1EventsMetaData contains all meta data concerning the SynapseGasOracleV1Events contract.
var SynapseGasOracleV1EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"calldataPrice\",\"type\":\"uint256\"}],\"name\":\"CalldataPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"GasPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nativePrice\",\"type\":\"uint256\"}],\"name\":\"NativePriceSet\",\"type\":\"event\"}]",
}

// SynapseGasOracleV1EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseGasOracleV1EventsMetaData.ABI instead.
var SynapseGasOracleV1EventsABI = SynapseGasOracleV1EventsMetaData.ABI

// SynapseGasOracleV1Events is an auto generated Go binding around an Ethereum contract.
type SynapseGasOracleV1Events struct {
	SynapseGasOracleV1EventsCaller     // Read-only binding to the contract
	SynapseGasOracleV1EventsTransactor // Write-only binding to the contract
	SynapseGasOracleV1EventsFilterer   // Log filterer for contract events
}

// SynapseGasOracleV1EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseGasOracleV1EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleV1EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseGasOracleV1EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleV1EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseGasOracleV1EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasOracleV1EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseGasOracleV1EventsSession struct {
	Contract     *SynapseGasOracleV1Events // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SynapseGasOracleV1EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseGasOracleV1EventsCallerSession struct {
	Contract *SynapseGasOracleV1EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SynapseGasOracleV1EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseGasOracleV1EventsTransactorSession struct {
	Contract     *SynapseGasOracleV1EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SynapseGasOracleV1EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseGasOracleV1EventsRaw struct {
	Contract *SynapseGasOracleV1Events // Generic contract binding to access the raw methods on
}

// SynapseGasOracleV1EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseGasOracleV1EventsCallerRaw struct {
	Contract *SynapseGasOracleV1EventsCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseGasOracleV1EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseGasOracleV1EventsTransactorRaw struct {
	Contract *SynapseGasOracleV1EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseGasOracleV1Events creates a new instance of SynapseGasOracleV1Events, bound to a specific deployed contract.
func NewSynapseGasOracleV1Events(address common.Address, backend bind.ContractBackend) (*SynapseGasOracleV1Events, error) {
	contract, err := bindSynapseGasOracleV1Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1Events{SynapseGasOracleV1EventsCaller: SynapseGasOracleV1EventsCaller{contract: contract}, SynapseGasOracleV1EventsTransactor: SynapseGasOracleV1EventsTransactor{contract: contract}, SynapseGasOracleV1EventsFilterer: SynapseGasOracleV1EventsFilterer{contract: contract}}, nil
}

// NewSynapseGasOracleV1EventsCaller creates a new read-only instance of SynapseGasOracleV1Events, bound to a specific deployed contract.
func NewSynapseGasOracleV1EventsCaller(address common.Address, caller bind.ContractCaller) (*SynapseGasOracleV1EventsCaller, error) {
	contract, err := bindSynapseGasOracleV1Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1EventsCaller{contract: contract}, nil
}

// NewSynapseGasOracleV1EventsTransactor creates a new write-only instance of SynapseGasOracleV1Events, bound to a specific deployed contract.
func NewSynapseGasOracleV1EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseGasOracleV1EventsTransactor, error) {
	contract, err := bindSynapseGasOracleV1Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1EventsTransactor{contract: contract}, nil
}

// NewSynapseGasOracleV1EventsFilterer creates a new log filterer instance of SynapseGasOracleV1Events, bound to a specific deployed contract.
func NewSynapseGasOracleV1EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseGasOracleV1EventsFilterer, error) {
	contract, err := bindSynapseGasOracleV1Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1EventsFilterer{contract: contract}, nil
}

// bindSynapseGasOracleV1Events binds a generic wrapper to an already deployed contract.
func bindSynapseGasOracleV1Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseGasOracleV1EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseGasOracleV1Events.Contract.SynapseGasOracleV1EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasOracleV1Events.Contract.SynapseGasOracleV1EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseGasOracleV1Events.Contract.SynapseGasOracleV1EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseGasOracleV1Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasOracleV1Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseGasOracleV1Events.Contract.contract.Transact(opts, method, params...)
}

// SynapseGasOracleV1EventsCalldataPriceSetIterator is returned from FilterCalldataPriceSet and is used to iterate over the raw logs and unpacked data for CalldataPriceSet events raised by the SynapseGasOracleV1Events contract.
type SynapseGasOracleV1EventsCalldataPriceSetIterator struct {
	Event *SynapseGasOracleV1EventsCalldataPriceSet // Event containing the contract specifics and raw log

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
func (it *SynapseGasOracleV1EventsCalldataPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseGasOracleV1EventsCalldataPriceSet)
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
		it.Event = new(SynapseGasOracleV1EventsCalldataPriceSet)
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
func (it *SynapseGasOracleV1EventsCalldataPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseGasOracleV1EventsCalldataPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseGasOracleV1EventsCalldataPriceSet represents a CalldataPriceSet event raised by the SynapseGasOracleV1Events contract.
type SynapseGasOracleV1EventsCalldataPriceSet struct {
	ChainId       uint64
	CalldataPrice *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalldataPriceSet is a free log retrieval operation binding the contract event 0x3f67e101936b3cb094183a2ed4bf880810a36909c55c27a7d145ee4bbf3dd515.
//
// Solidity: event CalldataPriceSet(uint64 indexed chainId, uint256 calldataPrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) FilterCalldataPriceSet(opts *bind.FilterOpts, chainId []uint64) (*SynapseGasOracleV1EventsCalldataPriceSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1Events.contract.FilterLogs(opts, "CalldataPriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1EventsCalldataPriceSetIterator{contract: _SynapseGasOracleV1Events.contract, event: "CalldataPriceSet", logs: logs, sub: sub}, nil
}

// WatchCalldataPriceSet is a free log subscription operation binding the contract event 0x3f67e101936b3cb094183a2ed4bf880810a36909c55c27a7d145ee4bbf3dd515.
//
// Solidity: event CalldataPriceSet(uint64 indexed chainId, uint256 calldataPrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) WatchCalldataPriceSet(opts *bind.WatchOpts, sink chan<- *SynapseGasOracleV1EventsCalldataPriceSet, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1Events.contract.WatchLogs(opts, "CalldataPriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseGasOracleV1EventsCalldataPriceSet)
				if err := _SynapseGasOracleV1Events.contract.UnpackLog(event, "CalldataPriceSet", log); err != nil {
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

// ParseCalldataPriceSet is a log parse operation binding the contract event 0x3f67e101936b3cb094183a2ed4bf880810a36909c55c27a7d145ee4bbf3dd515.
//
// Solidity: event CalldataPriceSet(uint64 indexed chainId, uint256 calldataPrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) ParseCalldataPriceSet(log types.Log) (*SynapseGasOracleV1EventsCalldataPriceSet, error) {
	event := new(SynapseGasOracleV1EventsCalldataPriceSet)
	if err := _SynapseGasOracleV1Events.contract.UnpackLog(event, "CalldataPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseGasOracleV1EventsGasPriceSetIterator is returned from FilterGasPriceSet and is used to iterate over the raw logs and unpacked data for GasPriceSet events raised by the SynapseGasOracleV1Events contract.
type SynapseGasOracleV1EventsGasPriceSetIterator struct {
	Event *SynapseGasOracleV1EventsGasPriceSet // Event containing the contract specifics and raw log

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
func (it *SynapseGasOracleV1EventsGasPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseGasOracleV1EventsGasPriceSet)
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
		it.Event = new(SynapseGasOracleV1EventsGasPriceSet)
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
func (it *SynapseGasOracleV1EventsGasPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseGasOracleV1EventsGasPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseGasOracleV1EventsGasPriceSet represents a GasPriceSet event raised by the SynapseGasOracleV1Events contract.
type SynapseGasOracleV1EventsGasPriceSet struct {
	ChainId  uint64
	GasPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasPriceSet is a free log retrieval operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 indexed chainId, uint256 gasPrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) FilterGasPriceSet(opts *bind.FilterOpts, chainId []uint64) (*SynapseGasOracleV1EventsGasPriceSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1Events.contract.FilterLogs(opts, "GasPriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1EventsGasPriceSetIterator{contract: _SynapseGasOracleV1Events.contract, event: "GasPriceSet", logs: logs, sub: sub}, nil
}

// WatchGasPriceSet is a free log subscription operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 indexed chainId, uint256 gasPrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) WatchGasPriceSet(opts *bind.WatchOpts, sink chan<- *SynapseGasOracleV1EventsGasPriceSet, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1Events.contract.WatchLogs(opts, "GasPriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseGasOracleV1EventsGasPriceSet)
				if err := _SynapseGasOracleV1Events.contract.UnpackLog(event, "GasPriceSet", log); err != nil {
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

// ParseGasPriceSet is a log parse operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 indexed chainId, uint256 gasPrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) ParseGasPriceSet(log types.Log) (*SynapseGasOracleV1EventsGasPriceSet, error) {
	event := new(SynapseGasOracleV1EventsGasPriceSet)
	if err := _SynapseGasOracleV1Events.contract.UnpackLog(event, "GasPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseGasOracleV1EventsNativePriceSetIterator is returned from FilterNativePriceSet and is used to iterate over the raw logs and unpacked data for NativePriceSet events raised by the SynapseGasOracleV1Events contract.
type SynapseGasOracleV1EventsNativePriceSetIterator struct {
	Event *SynapseGasOracleV1EventsNativePriceSet // Event containing the contract specifics and raw log

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
func (it *SynapseGasOracleV1EventsNativePriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseGasOracleV1EventsNativePriceSet)
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
		it.Event = new(SynapseGasOracleV1EventsNativePriceSet)
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
func (it *SynapseGasOracleV1EventsNativePriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseGasOracleV1EventsNativePriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseGasOracleV1EventsNativePriceSet represents a NativePriceSet event raised by the SynapseGasOracleV1Events contract.
type SynapseGasOracleV1EventsNativePriceSet struct {
	ChainId     uint64
	NativePrice *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNativePriceSet is a free log retrieval operation binding the contract event 0x01e22a2f804ee98d53c44fad5dd2f7bd9dafc2737327f816f5330d66103f6dfd.
//
// Solidity: event NativePriceSet(uint64 indexed chainId, uint256 nativePrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) FilterNativePriceSet(opts *bind.FilterOpts, chainId []uint64) (*SynapseGasOracleV1EventsNativePriceSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1Events.contract.FilterLogs(opts, "NativePriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseGasOracleV1EventsNativePriceSetIterator{contract: _SynapseGasOracleV1Events.contract, event: "NativePriceSet", logs: logs, sub: sub}, nil
}

// WatchNativePriceSet is a free log subscription operation binding the contract event 0x01e22a2f804ee98d53c44fad5dd2f7bd9dafc2737327f816f5330d66103f6dfd.
//
// Solidity: event NativePriceSet(uint64 indexed chainId, uint256 nativePrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) WatchNativePriceSet(opts *bind.WatchOpts, sink chan<- *SynapseGasOracleV1EventsNativePriceSet, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _SynapseGasOracleV1Events.contract.WatchLogs(opts, "NativePriceSet", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseGasOracleV1EventsNativePriceSet)
				if err := _SynapseGasOracleV1Events.contract.UnpackLog(event, "NativePriceSet", log); err != nil {
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

// ParseNativePriceSet is a log parse operation binding the contract event 0x01e22a2f804ee98d53c44fad5dd2f7bd9dafc2737327f816f5330d66103f6dfd.
//
// Solidity: event NativePriceSet(uint64 indexed chainId, uint256 nativePrice)
func (_SynapseGasOracleV1Events *SynapseGasOracleV1EventsFilterer) ParseNativePriceSet(log types.Log) (*SynapseGasOracleV1EventsNativePriceSet, error) {
	event := new(SynapseGasOracleV1EventsNativePriceSet)
	if err := _SynapseGasOracleV1Events.contract.UnpackLog(event, "NativePriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
