// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executionservicemock

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

// ExecutionServiceMockMetaData contains all meta data concerning the ExecutionServiceMock contract.
var ExecutionServiceMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62014bad": "executorEOA()",
		"c473e7e8": "getExecutionFee(uint256,uint256,bytes)",
		"e4e06522": "requestExecution(uint256,uint256,bytes32,uint256,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610290806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806362014bad14610046578063c473e7e814610090578063e4e06522146100b5575b600080fd5b6000546100669073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100a761009e3660046101a6565b60009392505050565b604051908152602001610087565b6100ca6100c33660046101f6565b5050505050565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261010c57600080fd5b813567ffffffffffffffff80821115610127576101276100cc565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561016d5761016d6100cc565b8160405283815286602085880101111561018657600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156101bb57600080fd5b8335925060208401359150604084013567ffffffffffffffff8111156101e057600080fd5b6101ec868287016100fb565b9150509250925092565b600080600080600060a0868803121561020e57600080fd5b85359450602086013593506040860135925060608601359150608086013567ffffffffffffffff81111561024157600080fd5b61024d888289016100fb565b915050929550929590935056fea2646970667358221220c15c20205132b64aecd1865fd1e7602a4f626d5e1e0fd8064e46f96b970b8ecf64736f6c63430008140033",
}

// ExecutionServiceMockABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionServiceMockMetaData.ABI instead.
var ExecutionServiceMockABI = ExecutionServiceMockMetaData.ABI

// Deprecated: Use ExecutionServiceMockMetaData.Sigs instead.
// ExecutionServiceMockFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionServiceMockFuncSigs = ExecutionServiceMockMetaData.Sigs

// ExecutionServiceMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutionServiceMockMetaData.Bin instead.
var ExecutionServiceMockBin = ExecutionServiceMockMetaData.Bin

// DeployExecutionServiceMock deploys a new Ethereum contract, binding an instance of ExecutionServiceMock to it.
func DeployExecutionServiceMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExecutionServiceMock, error) {
	parsed, err := ExecutionServiceMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutionServiceMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutionServiceMock{ExecutionServiceMockCaller: ExecutionServiceMockCaller{contract: contract}, ExecutionServiceMockTransactor: ExecutionServiceMockTransactor{contract: contract}, ExecutionServiceMockFilterer: ExecutionServiceMockFilterer{contract: contract}}, nil
}

// ExecutionServiceMock is an auto generated Go binding around an Ethereum contract.
type ExecutionServiceMock struct {
	ExecutionServiceMockCaller     // Read-only binding to the contract
	ExecutionServiceMockTransactor // Write-only binding to the contract
	ExecutionServiceMockFilterer   // Log filterer for contract events
}

// ExecutionServiceMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionServiceMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionServiceMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionServiceMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionServiceMockSession struct {
	Contract     *ExecutionServiceMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExecutionServiceMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionServiceMockCallerSession struct {
	Contract *ExecutionServiceMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ExecutionServiceMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionServiceMockTransactorSession struct {
	Contract     *ExecutionServiceMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ExecutionServiceMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionServiceMockRaw struct {
	Contract *ExecutionServiceMock // Generic contract binding to access the raw methods on
}

// ExecutionServiceMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionServiceMockCallerRaw struct {
	Contract *ExecutionServiceMockCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionServiceMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionServiceMockTransactorRaw struct {
	Contract *ExecutionServiceMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionServiceMock creates a new instance of ExecutionServiceMock, bound to a specific deployed contract.
func NewExecutionServiceMock(address common.Address, backend bind.ContractBackend) (*ExecutionServiceMock, error) {
	contract, err := bindExecutionServiceMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceMock{ExecutionServiceMockCaller: ExecutionServiceMockCaller{contract: contract}, ExecutionServiceMockTransactor: ExecutionServiceMockTransactor{contract: contract}, ExecutionServiceMockFilterer: ExecutionServiceMockFilterer{contract: contract}}, nil
}

// NewExecutionServiceMockCaller creates a new read-only instance of ExecutionServiceMock, bound to a specific deployed contract.
func NewExecutionServiceMockCaller(address common.Address, caller bind.ContractCaller) (*ExecutionServiceMockCaller, error) {
	contract, err := bindExecutionServiceMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceMockCaller{contract: contract}, nil
}

// NewExecutionServiceMockTransactor creates a new write-only instance of ExecutionServiceMock, bound to a specific deployed contract.
func NewExecutionServiceMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionServiceMockTransactor, error) {
	contract, err := bindExecutionServiceMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceMockTransactor{contract: contract}, nil
}

// NewExecutionServiceMockFilterer creates a new log filterer instance of ExecutionServiceMock, bound to a specific deployed contract.
func NewExecutionServiceMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionServiceMockFilterer, error) {
	contract, err := bindExecutionServiceMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceMockFilterer{contract: contract}, nil
}

// bindExecutionServiceMock binds a generic wrapper to an already deployed contract.
func bindExecutionServiceMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutionServiceMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionServiceMock *ExecutionServiceMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionServiceMock.Contract.ExecutionServiceMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionServiceMock *ExecutionServiceMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionServiceMock.Contract.ExecutionServiceMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionServiceMock *ExecutionServiceMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionServiceMock.Contract.ExecutionServiceMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionServiceMock *ExecutionServiceMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionServiceMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionServiceMock *ExecutionServiceMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionServiceMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionServiceMock *ExecutionServiceMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionServiceMock.Contract.contract.Transact(opts, method, params...)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ExecutionServiceMock *ExecutionServiceMockCaller) ExecutorEOA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionServiceMock.contract.Call(opts, &out, "executorEOA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ExecutionServiceMock *ExecutionServiceMockSession) ExecutorEOA() (common.Address, error) {
	return _ExecutionServiceMock.Contract.ExecutorEOA(&_ExecutionServiceMock.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ExecutionServiceMock *ExecutionServiceMockCallerSession) ExecutorEOA() (common.Address, error) {
	return _ExecutionServiceMock.Contract.ExecutorEOA(&_ExecutionServiceMock.CallOpts)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ExecutionServiceMock *ExecutionServiceMockCaller) GetExecutionFee(opts *bind.CallOpts, dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	var out []interface{}
	err := _ExecutionServiceMock.contract.Call(opts, &out, "getExecutionFee", dstChainId, txPayloadSize, options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ExecutionServiceMock *ExecutionServiceMockSession) GetExecutionFee(dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _ExecutionServiceMock.Contract.GetExecutionFee(&_ExecutionServiceMock.CallOpts, dstChainId, txPayloadSize, options)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ExecutionServiceMock *ExecutionServiceMockCallerSession) GetExecutionFee(dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _ExecutionServiceMock.Contract.GetExecutionFee(&_ExecutionServiceMock.CallOpts, dstChainId, txPayloadSize, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ExecutionServiceMock *ExecutionServiceMockTransactor) RequestExecution(opts *bind.TransactOpts, dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ExecutionServiceMock.contract.Transact(opts, "requestExecution", dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ExecutionServiceMock *ExecutionServiceMockSession) RequestExecution(dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ExecutionServiceMock.Contract.RequestExecution(&_ExecutionServiceMock.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ExecutionServiceMock *ExecutionServiceMockTransactorSession) RequestExecution(dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ExecutionServiceMock.Contract.RequestExecution(&_ExecutionServiceMock.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// IExecutionServiceMetaData contains all meta data concerning the IExecutionService contract.
var IExecutionServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62014bad": "executorEOA()",
		"c473e7e8": "getExecutionFee(uint256,uint256,bytes)",
		"e4e06522": "requestExecution(uint256,uint256,bytes32,uint256,bytes)",
	},
}

// IExecutionServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionServiceMetaData.ABI instead.
var IExecutionServiceABI = IExecutionServiceMetaData.ABI

// Deprecated: Use IExecutionServiceMetaData.Sigs instead.
// IExecutionServiceFuncSigs maps the 4-byte function signature to its string representation.
var IExecutionServiceFuncSigs = IExecutionServiceMetaData.Sigs

// IExecutionService is an auto generated Go binding around an Ethereum contract.
type IExecutionService struct {
	IExecutionServiceCaller     // Read-only binding to the contract
	IExecutionServiceTransactor // Write-only binding to the contract
	IExecutionServiceFilterer   // Log filterer for contract events
}

// IExecutionServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionServiceSession struct {
	Contract     *IExecutionService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IExecutionServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionServiceCallerSession struct {
	Contract *IExecutionServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IExecutionServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionServiceTransactorSession struct {
	Contract     *IExecutionServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IExecutionServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionServiceRaw struct {
	Contract *IExecutionService // Generic contract binding to access the raw methods on
}

// IExecutionServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionServiceCallerRaw struct {
	Contract *IExecutionServiceCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionServiceTransactorRaw struct {
	Contract *IExecutionServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecutionService creates a new instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionService(address common.Address, backend bind.ContractBackend) (*IExecutionService, error) {
	contract, err := bindIExecutionService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecutionService{IExecutionServiceCaller: IExecutionServiceCaller{contract: contract}, IExecutionServiceTransactor: IExecutionServiceTransactor{contract: contract}, IExecutionServiceFilterer: IExecutionServiceFilterer{contract: contract}}, nil
}

// NewIExecutionServiceCaller creates a new read-only instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceCaller(address common.Address, caller bind.ContractCaller) (*IExecutionServiceCaller, error) {
	contract, err := bindIExecutionService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceCaller{contract: contract}, nil
}

// NewIExecutionServiceTransactor creates a new write-only instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionServiceTransactor, error) {
	contract, err := bindIExecutionService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceTransactor{contract: contract}, nil
}

// NewIExecutionServiceFilterer creates a new log filterer instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionServiceFilterer, error) {
	contract, err := bindIExecutionService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceFilterer{contract: contract}, nil
}

// bindIExecutionService binds a generic wrapper to an already deployed contract.
func bindIExecutionService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExecutionServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionService *IExecutionServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionService.Contract.IExecutionServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionService *IExecutionServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionService.Contract.IExecutionServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionService *IExecutionServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionService.Contract.IExecutionServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionService *IExecutionServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionService *IExecutionServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionService *IExecutionServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionService.Contract.contract.Transact(opts, method, params...)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceCaller) ExecutorEOA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IExecutionService.contract.Call(opts, &out, "executorEOA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceSession) ExecutorEOA() (common.Address, error) {
	return _IExecutionService.Contract.ExecutorEOA(&_IExecutionService.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceCallerSession) ExecutorEOA() (common.Address, error) {
	return _IExecutionService.Contract.ExecutorEOA(&_IExecutionService.CallOpts)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceCaller) GetExecutionFee(opts *bind.CallOpts, dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionService.contract.Call(opts, &out, "getExecutionFee", dstChainId, txPayloadSize, options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceSession) GetExecutionFee(dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _IExecutionService.Contract.GetExecutionFee(&_IExecutionService.CallOpts, dstChainId, txPayloadSize, options)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceCallerSession) GetExecutionFee(dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _IExecutionService.Contract.GetExecutionFee(&_IExecutionService.CallOpts, dstChainId, txPayloadSize, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceTransactor) RequestExecution(opts *bind.TransactOpts, dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.contract.Transact(opts, "requestExecution", dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceSession) RequestExecution(dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.Contract.RequestExecution(&_IExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceTransactorSession) RequestExecution(dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.Contract.RequestExecution(&_IExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}
