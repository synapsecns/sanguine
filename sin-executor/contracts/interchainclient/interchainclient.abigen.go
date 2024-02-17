// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interchainclient

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

// InterchainEntry is an auto generated low-level Go binding around an user-defined struct.
type InterchainEntry struct {
	SrcChainId  *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	DataHash    [32]byte
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

// IInterchainAppMetaData contains all meta data concerning the IInterchainApp contract.
var IInterchainAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"getLinkedIApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptimisticTimePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequiredResponses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"chainIDs\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"linkedIApps\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sendingModules\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"receivingModules\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"optimisticTimePeriod\",\"type\":\"uint64\"}],\"name\":\"setAppConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7bcad630": "appReceive()",
		"bfc849ee": "getLinkedIApp(uint64)",
		"7c9abd3e": "getOptimisticTimePeriod()",
		"a45e107a": "getReceivingModules()",
		"f31b19a9": "getRequiredResponses()",
		"ea13398f": "getSendingModules()",
		"e1ef3b3f": "send(bytes32,uint256,bytes)",
		"dd34f56a": "setAppConfig(uint64[],address[],address[],address[],uint256,uint64)",
	},
}

// IInterchainAppABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainAppMetaData.ABI instead.
var IInterchainAppABI = IInterchainAppMetaData.ABI

// Deprecated: Use IInterchainAppMetaData.Sigs instead.
// IInterchainAppFuncSigs maps the 4-byte function signature to its string representation.
var IInterchainAppFuncSigs = IInterchainAppMetaData.Sigs

// IInterchainApp is an auto generated Go binding around an Ethereum contract.
type IInterchainApp struct {
	IInterchainAppCaller     // Read-only binding to the contract
	IInterchainAppTransactor // Write-only binding to the contract
	IInterchainAppFilterer   // Log filterer for contract events
}

// IInterchainAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainAppSession struct {
	Contract     *IInterchainApp   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInterchainAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainAppCallerSession struct {
	Contract *IInterchainAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IInterchainAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainAppTransactorSession struct {
	Contract     *IInterchainAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IInterchainAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainAppRaw struct {
	Contract *IInterchainApp // Generic contract binding to access the raw methods on
}

// IInterchainAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainAppCallerRaw struct {
	Contract *IInterchainAppCaller // Generic read-only contract binding to access the raw methods on
}

// IInterchainAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainAppTransactorRaw struct {
	Contract *IInterchainAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainApp creates a new instance of IInterchainApp, bound to a specific deployed contract.
func NewIInterchainApp(address common.Address, backend bind.ContractBackend) (*IInterchainApp, error) {
	contract, err := bindIInterchainApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainApp{IInterchainAppCaller: IInterchainAppCaller{contract: contract}, IInterchainAppTransactor: IInterchainAppTransactor{contract: contract}, IInterchainAppFilterer: IInterchainAppFilterer{contract: contract}}, nil
}

// NewIInterchainAppCaller creates a new read-only instance of IInterchainApp, bound to a specific deployed contract.
func NewIInterchainAppCaller(address common.Address, caller bind.ContractCaller) (*IInterchainAppCaller, error) {
	contract, err := bindIInterchainApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppCaller{contract: contract}, nil
}

// NewIInterchainAppTransactor creates a new write-only instance of IInterchainApp, bound to a specific deployed contract.
func NewIInterchainAppTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainAppTransactor, error) {
	contract, err := bindIInterchainApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppTransactor{contract: contract}, nil
}

// NewIInterchainAppFilterer creates a new log filterer instance of IInterchainApp, bound to a specific deployed contract.
func NewIInterchainAppFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainAppFilterer, error) {
	contract, err := bindIInterchainApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppFilterer{contract: contract}, nil
}

// bindIInterchainApp binds a generic wrapper to an already deployed contract.
func bindIInterchainApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainApp *IInterchainAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainApp.Contract.IInterchainAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainApp *IInterchainAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainApp.Contract.IInterchainAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainApp *IInterchainAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainApp.Contract.IInterchainAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainApp *IInterchainAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainApp *IInterchainAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainApp *IInterchainAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainApp.Contract.contract.Transact(opts, method, params...)
}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_IInterchainApp *IInterchainAppCaller) GetLinkedIApp(opts *bind.CallOpts, chainID uint64) (common.Address, error) {
	var out []interface{}
	err := _IInterchainApp.contract.Call(opts, &out, "getLinkedIApp", chainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_IInterchainApp *IInterchainAppSession) GetLinkedIApp(chainID uint64) (common.Address, error) {
	return _IInterchainApp.Contract.GetLinkedIApp(&_IInterchainApp.CallOpts, chainID)
}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_IInterchainApp *IInterchainAppCallerSession) GetLinkedIApp(chainID uint64) (common.Address, error) {
	return _IInterchainApp.Contract.GetLinkedIApp(&_IInterchainApp.CallOpts, chainID)
}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() view returns(uint64)
func (_IInterchainApp *IInterchainAppCaller) GetOptimisticTimePeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IInterchainApp.contract.Call(opts, &out, "getOptimisticTimePeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() view returns(uint64)
func (_IInterchainApp *IInterchainAppSession) GetOptimisticTimePeriod() (uint64, error) {
	return _IInterchainApp.Contract.GetOptimisticTimePeriod(&_IInterchainApp.CallOpts)
}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() view returns(uint64)
func (_IInterchainApp *IInterchainAppCallerSession) GetOptimisticTimePeriod() (uint64, error) {
	return _IInterchainApp.Contract.GetOptimisticTimePeriod(&_IInterchainApp.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_IInterchainApp *IInterchainAppCaller) GetReceivingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IInterchainApp.contract.Call(opts, &out, "getReceivingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_IInterchainApp *IInterchainAppSession) GetReceivingModules() ([]common.Address, error) {
	return _IInterchainApp.Contract.GetReceivingModules(&_IInterchainApp.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_IInterchainApp *IInterchainAppCallerSession) GetReceivingModules() ([]common.Address, error) {
	return _IInterchainApp.Contract.GetReceivingModules(&_IInterchainApp.CallOpts)
}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() view returns(uint256)
func (_IInterchainApp *IInterchainAppCaller) GetRequiredResponses(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainApp.contract.Call(opts, &out, "getRequiredResponses")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() view returns(uint256)
func (_IInterchainApp *IInterchainAppSession) GetRequiredResponses() (*big.Int, error) {
	return _IInterchainApp.Contract.GetRequiredResponses(&_IInterchainApp.CallOpts)
}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() view returns(uint256)
func (_IInterchainApp *IInterchainAppCallerSession) GetRequiredResponses() (*big.Int, error) {
	return _IInterchainApp.Contract.GetRequiredResponses(&_IInterchainApp.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_IInterchainApp *IInterchainAppCaller) GetSendingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IInterchainApp.contract.Call(opts, &out, "getSendingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_IInterchainApp *IInterchainAppSession) GetSendingModules() ([]common.Address, error) {
	return _IInterchainApp.Contract.GetSendingModules(&_IInterchainApp.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_IInterchainApp *IInterchainAppCallerSession) GetSendingModules() ([]common.Address, error) {
	return _IInterchainApp.Contract.GetSendingModules(&_IInterchainApp.CallOpts)
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_IInterchainApp *IInterchainAppTransactor) AppReceive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainApp.contract.Transact(opts, "appReceive")
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_IInterchainApp *IInterchainAppSession) AppReceive() (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts)
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_IInterchainApp *IInterchainAppTransactorSession) AppReceive() (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactor) Send(opts *bind.TransactOpts, receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.contract.Transact(opts, "send", receiver, dstChainId, message)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppSession) Send(receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.Send(&_IInterchainApp.TransactOpts, receiver, dstChainId, message)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactorSession) Send(receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.Send(&_IInterchainApp.TransactOpts, receiver, dstChainId, message)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] sendingModules, address[] receivingModules, uint256 requiredResponses, uint64 optimisticTimePeriod) returns()
func (_IInterchainApp *IInterchainAppTransactor) SetAppConfig(opts *bind.TransactOpts, chainIDs []uint64, linkedIApps []common.Address, sendingModules []common.Address, receivingModules []common.Address, requiredResponses *big.Int, optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _IInterchainApp.contract.Transact(opts, "setAppConfig", chainIDs, linkedIApps, sendingModules, receivingModules, requiredResponses, optimisticTimePeriod)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] sendingModules, address[] receivingModules, uint256 requiredResponses, uint64 optimisticTimePeriod) returns()
func (_IInterchainApp *IInterchainAppSession) SetAppConfig(chainIDs []uint64, linkedIApps []common.Address, sendingModules []common.Address, receivingModules []common.Address, requiredResponses *big.Int, optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _IInterchainApp.Contract.SetAppConfig(&_IInterchainApp.TransactOpts, chainIDs, linkedIApps, sendingModules, receivingModules, requiredResponses, optimisticTimePeriod)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] sendingModules, address[] receivingModules, uint256 requiredResponses, uint64 optimisticTimePeriod) returns()
func (_IInterchainApp *IInterchainAppTransactorSession) SetAppConfig(chainIDs []uint64, linkedIApps []common.Address, sendingModules []common.Address, receivingModules []common.Address, requiredResponses *big.Int, optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _IInterchainApp.Contract.SetAppConfig(&_IInterchainApp.TransactOpts, chainIDs, linkedIApps, sendingModules, receivingModules, requiredResponses, optimisticTimePeriod)
}

// IInterchainClientV1MetaData contains all meta data concerning the IInterchainClientV1 contract.
var IInterchainClientV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"convertAddressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"convertBytes32ToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"interchainSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchainDB\",\"type\":\"address\"}],\"name\":\"setInterchainDB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5893740e": "convertAddressToBytes32(address)",
		"1efa2220": "convertBytes32ToAddress(bytes32)",
		"d3388b80": "interchainExecute(bytes32,bytes)",
		"8366a109": "interchainSend(bytes32,uint256,bytes,address[])",
		"b7ce2078": "setInterchainDB(address)",
		"f34234c8": "setLinkedClient(uint256,bytes32)",
	},
}

// IInterchainClientV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainClientV1MetaData.ABI instead.
var IInterchainClientV1ABI = IInterchainClientV1MetaData.ABI

// Deprecated: Use IInterchainClientV1MetaData.Sigs instead.
// IInterchainClientV1FuncSigs maps the 4-byte function signature to its string representation.
var IInterchainClientV1FuncSigs = IInterchainClientV1MetaData.Sigs

// IInterchainClientV1 is an auto generated Go binding around an Ethereum contract.
type IInterchainClientV1 struct {
	IInterchainClientV1Caller     // Read-only binding to the contract
	IInterchainClientV1Transactor // Write-only binding to the contract
	IInterchainClientV1Filterer   // Log filterer for contract events
}

// IInterchainClientV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainClientV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainClientV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainClientV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainClientV1Session struct {
	Contract     *IInterchainClientV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IInterchainClientV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainClientV1CallerSession struct {
	Contract *IInterchainClientV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IInterchainClientV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainClientV1TransactorSession struct {
	Contract     *IInterchainClientV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IInterchainClientV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainClientV1Raw struct {
	Contract *IInterchainClientV1 // Generic contract binding to access the raw methods on
}

// IInterchainClientV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainClientV1CallerRaw struct {
	Contract *IInterchainClientV1Caller // Generic read-only contract binding to access the raw methods on
}

// IInterchainClientV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainClientV1TransactorRaw struct {
	Contract *IInterchainClientV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainClientV1 creates a new instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1(address common.Address, backend bind.ContractBackend) (*IInterchainClientV1, error) {
	contract, err := bindIInterchainClientV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1{IInterchainClientV1Caller: IInterchainClientV1Caller{contract: contract}, IInterchainClientV1Transactor: IInterchainClientV1Transactor{contract: contract}, IInterchainClientV1Filterer: IInterchainClientV1Filterer{contract: contract}}, nil
}

// NewIInterchainClientV1Caller creates a new read-only instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Caller(address common.Address, caller bind.ContractCaller) (*IInterchainClientV1Caller, error) {
	contract, err := bindIInterchainClientV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Caller{contract: contract}, nil
}

// NewIInterchainClientV1Transactor creates a new write-only instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Transactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainClientV1Transactor, error) {
	contract, err := bindIInterchainClientV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Transactor{contract: contract}, nil
}

// NewIInterchainClientV1Filterer creates a new log filterer instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Filterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainClientV1Filterer, error) {
	contract, err := bindIInterchainClientV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Filterer{contract: contract}, nil
}

// bindIInterchainClientV1 binds a generic wrapper to an already deployed contract.
func bindIInterchainClientV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainClientV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainClientV1.Contract.IInterchainClientV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.IInterchainClientV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.IInterchainClientV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainClientV1 *IInterchainClientV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainClientV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainClientV1 *IInterchainClientV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainClientV1 *IInterchainClientV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.contract.Transact(opts, method, params...)
}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1Caller) ConvertAddressToBytes32(opts *bind.CallOpts, _address common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "convertAddressToBytes32", _address)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1Session) ConvertAddressToBytes32(_address common.Address) ([32]byte, error) {
	return _IInterchainClientV1.Contract.ConvertAddressToBytes32(&_IInterchainClientV1.CallOpts, _address)
}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) ConvertAddressToBytes32(_address common.Address) ([32]byte, error) {
	return _IInterchainClientV1.Contract.ConvertAddressToBytes32(&_IInterchainClientV1.CallOpts, _address)
}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Caller) ConvertBytes32ToAddress(opts *bind.CallOpts, _bytes32 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "convertBytes32ToAddress", _bytes32)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Session) ConvertBytes32ToAddress(_bytes32 [32]byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.ConvertBytes32ToAddress(&_IInterchainClientV1.CallOpts, _bytes32)
}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) ConvertBytes32ToAddress(_bytes32 [32]byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.ConvertBytes32ToAddress(&_IInterchainClientV1.CallOpts, _bytes32)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0xd3388b80.
//
// Solidity: function interchainExecute(bytes32 transactionID, bytes transaction) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainExecute(opts *bind.TransactOpts, transactionID [32]byte, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainExecute", transactionID, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0xd3388b80.
//
// Solidity: function interchainExecute(bytes32 transactionID, bytes transaction) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainExecute(transactionID [32]byte, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, transactionID, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0xd3388b80.
//
// Solidity: function interchainExecute(bytes32 transactionID, bytes transaction) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainExecute(transactionID [32]byte, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, transactionID, transaction)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] srcModules) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainSend(opts *bind.TransactOpts, receiver [32]byte, dstChainId *big.Int, message []byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainSend", receiver, dstChainId, message, srcModules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] srcModules) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, receiver, dstChainId, message, srcModules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] srcModules) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, receiver, dstChainId, message, srcModules)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) SetInterchainDB(opts *bind.TransactOpts, _interchainDB common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "setInterchainDB", _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetInterchainDB(&_IInterchainClientV1.TransactOpts, _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetInterchainDB(&_IInterchainClientV1.TransactOpts, _interchainDB)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) SetLinkedClient(opts *bind.TransactOpts, chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "setLinkedClient", chainId, client)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) SetLinkedClient(chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetLinkedClient(&_IInterchainClientV1.TransactOpts, chainId, client)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) SetLinkedClient(chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetLinkedClient(&_IInterchainClientV1.TransactOpts, chainId, client)
}

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"existingDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"newEntry\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingEntries\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__EntryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"getWriterNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"readEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b8a740e0": "getEntry(address,uint256)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"4a30a686": "getWriterNonce(address)",
		"d48588e0": "readEntry(address,(uint256,bytes32,uint256,bytes32))",
		"b4f16bae": "requestVerification(uint256,address,uint256,address[])",
		"9cbc6dd5": "verifyEntry((uint256,bytes32,uint256,bytes32))",
		"2ad8c706": "writeEntry(bytes32)",
		"67c769af": "writeEntryWithVerification(uint256,bytes32,address[])",
	},
}

// IInterchainDBABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainDBMetaData.ABI instead.
var IInterchainDBABI = IInterchainDBMetaData.ABI

// Deprecated: Use IInterchainDBMetaData.Sigs instead.
// IInterchainDBFuncSigs maps the 4-byte function signature to its string representation.
var IInterchainDBFuncSigs = IInterchainDBMetaData.Sigs

// IInterchainDB is an auto generated Go binding around an Ethereum contract.
type IInterchainDB struct {
	IInterchainDBCaller     // Read-only binding to the contract
	IInterchainDBTransactor // Write-only binding to the contract
	IInterchainDBFilterer   // Log filterer for contract events
}

// IInterchainDBCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainDBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainDBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainDBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainDBSession struct {
	Contract     *IInterchainDB    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInterchainDBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainDBCallerSession struct {
	Contract *IInterchainDBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IInterchainDBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainDBTransactorSession struct {
	Contract     *IInterchainDBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IInterchainDBRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainDBRaw struct {
	Contract *IInterchainDB // Generic contract binding to access the raw methods on
}

// IInterchainDBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainDBCallerRaw struct {
	Contract *IInterchainDBCaller // Generic read-only contract binding to access the raw methods on
}

// IInterchainDBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainDBTransactorRaw struct {
	Contract *IInterchainDBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainDB creates a new instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDB(address common.Address, backend bind.ContractBackend) (*IInterchainDB, error) {
	contract, err := bindIInterchainDB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainDB{IInterchainDBCaller: IInterchainDBCaller{contract: contract}, IInterchainDBTransactor: IInterchainDBTransactor{contract: contract}, IInterchainDBFilterer: IInterchainDBFilterer{contract: contract}}, nil
}

// NewIInterchainDBCaller creates a new read-only instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBCaller(address common.Address, caller bind.ContractCaller) (*IInterchainDBCaller, error) {
	contract, err := bindIInterchainDB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBCaller{contract: contract}, nil
}

// NewIInterchainDBTransactor creates a new write-only instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainDBTransactor, error) {
	contract, err := bindIInterchainDB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBTransactor{contract: contract}, nil
}

// NewIInterchainDBFilterer creates a new log filterer instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainDBFilterer, error) {
	contract, err := bindIInterchainDB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBFilterer{contract: contract}, nil
}

// bindIInterchainDB binds a generic wrapper to an already deployed contract.
func bindIInterchainDB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainDBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainDB *IInterchainDBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainDB.Contract.IInterchainDBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainDB *IInterchainDBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainDB.Contract.IInterchainDBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainDB *IInterchainDBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainDB.Contract.IInterchainDBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainDB *IInterchainDBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainDB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainDB *IInterchainDBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainDB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainDB *IInterchainDBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainDB.Contract.contract.Transact(opts, method, params...)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetEntry(opts *bind.CallOpts, writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntry", writer, writerNonce)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, writer, writerNonce)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, writer, writerNonce)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getInterchainFee", destChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, destChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, destChainId, srcModules)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetWriterNonce(opts *bind.CallOpts, writer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getWriterNonce", writer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetWriterNonce(&_IInterchainDB.CallOpts, writer)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetWriterNonce(&_IInterchainDB.CallOpts, writer)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) ReadEntry(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "readEntry", dstModule, entry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestVerification", destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "verifyEntry", entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactorSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// InterchainClientV1MetaData contains all meta data concerning the InterchainClientV1 contract.
var InterchainClientV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbWriterNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clientNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"convertAddressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"convertBytes32ToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchainDB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"interchainSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchainDB\",\"type\":\"address\"}],\"name\":\"setInterchainDB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0d898416": "clientNonce()",
		"5893740e": "convertAddressToBytes32(address)",
		"1efa2220": "convertBytes32ToAddress(bytes32)",
		"0e785ce0": "interchainDB()",
		"d3388b80": "interchainExecute(bytes32,bytes)",
		"8366a109": "interchainSend(bytes32,uint256,bytes,address[])",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"b7ce2078": "setInterchainDB(address)",
		"f34234c8": "setLinkedClient(uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50338061003757604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61004081610046565b50610096565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611298806100a56000396000f3fe6080604052600436106100bc5760003560e01c80638366a10911610074578063d3388b801161004e578063d3388b801461023f578063f2fde38b1461025f578063f34234c81461027f57600080fd5b80638366a109146101e15780638da5cb5b146101f4578063b7ce20781461021f57600080fd5b80631efa2220116100a55780631efa22201461016a5780635893740e14610188578063715018a6146101ca57600080fd5b80630d898416146100c15780630e785ce014610118575b600080fd5b3480156100cd57600080fd5b506000546100fa9074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b34801561012457600080fd5b506001546101459073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010f565b34801561017657600080fd5b50610145610185366004610b5c565b90565b34801561019457600080fd5b506101bc6101a3366004610b97565b73ffffffffffffffffffffffffffffffffffffffff1690565b60405190815260200161010f565b3480156101d657600080fd5b506101df6102ac565b005b6101df6101ef366004610c04565b6102c0565b34801561020057600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610145565b34801561022b57600080fd5b506101df61023a366004610b97565b610486565b34801561024b57600080fd5b506101df61025a366004610cb6565b6104d5565b34801561026b57600080fd5b506101df61027a366004610b97565b610a30565b34801561028b57600080fd5b506101df61029a366004610d02565b60009182526002602052604090912055565b6102b4610a94565b6102be6000610ae7565b565b6000805460405134923392909161030891849146918d918d918d918d9174010000000000000000000000000000000000000000900467ffffffffffffffff1690602001610d4f565b60408051601f198184030181529082905280516020909101206001547f67c769af00000000000000000000000000000000000000000000000000000000835290925060009173ffffffffffffffffffffffffffffffffffffffff909116906367c769af908690610382908d9087908c908c90600401610d99565b60206040518083038185885af11580156103a0573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906103c59190610e05565b905081898b7ff679da022706ad89d2486ce3bf6cba510d619af7c1b2f6799637ed085f9287d186468d8d600060149054906101000a900467ffffffffffffffff168960405161041996959493929190610e1e565b60405180910390a46000805474010000000000000000000000000000000000000000900467ffffffffffffffff1690601461045383610e8c565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050505050505050505050565b61048e610a94565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006104e382840184610fd3565b9050600060405180608001604052808360200151815260200160026000856020015181526020019081526020016000205481526020018360e0015181526020018360c0015181525090506000826000015183602001518460400151856060015186608001518760a0015160405160200161056296959493929190611091565b604051602081830303815290604052905080805190602001208260600151146105ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e76616c6964207472616e73616374696f6e2049440000000000000000000060448201526064015b60405180910390fd5b60006105f9846040015190565b905060008173ffffffffffffffffffffffffffffffffffffffff1663a45e107a6040518163ffffffff1660e01b8152600401600060405180830381865afa158015610648573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106709190810190611113565b905060008273ffffffffffffffffffffffffffffffffffffffff1663f31b19a96040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e39190610e05565b905060008373ffffffffffffffffffffffffffffffffffffffff16637c9abd3e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610732573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061075691906111c5565b67ffffffffffffffff1690506000835167ffffffffffffffff81111561077e5761077e610eb3565b6040519080825280602002602001820160405280156107a7578160200160208202803683370190505b50905060005b84518110156108d857600154855173ffffffffffffffffffffffffffffffffffffffff9091169063d48588e0908790849081106107ec576107ec6111e2565b602090810291909101810151604080517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301528c516024830152918c01516044820152908b0151606482015260608b0151608482015260a401602060405180830381865afa158015610885573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a99190610e05565b8282815181106108bb576108bb6111e2565b6020908102919091010152806108d081611211565b9150506107ad565b506000805b82518110156109315742848483815181106108fa576108fa6111e2565b602002602001015161090c9190611249565b1061091f578161091b81611211565b9250505b8061092981611211565b9150506108dd565b50838110156109c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4e6f7420656e6f7567682076616c696420726573706f6e73657320746f206d6560448201527f657420746865207468726573686f6c640000000000000000000000000000000060648201526084016105e3565b8573ffffffffffffffffffffffffffffffffffffffff16637bcad6306040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610a0a57600080fd5b505af1158015610a1e573d6000803e3d6000fd5b50505050505050505050505050505050565b610a38610a94565b73ffffffffffffffffffffffffffffffffffffffff8116610a88576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024016105e3565b610a9181610ae7565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146102be576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016105e3565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215610b6e57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610a9157600080fd5b600060208284031215610ba957600080fd5b8135610bb481610b75565b9392505050565b60008083601f840112610bcd57600080fd5b50813567ffffffffffffffff811115610be557600080fd5b602083019150836020828501011115610bfd57600080fd5b9250929050565b60008060008060008060808789031215610c1d57600080fd5b8635955060208701359450604087013567ffffffffffffffff80821115610c4357600080fd5b610c4f8a838b01610bbb565b90965094506060890135915080821115610c6857600080fd5b818901915089601f830112610c7c57600080fd5b813581811115610c8b57600080fd5b8a60208260051b8501011115610ca057600080fd5b6020830194508093505050509295509295509295565b600080600060408486031215610ccb57600080fd5b83359250602084013567ffffffffffffffff811115610ce957600080fd5b610cf586828701610bbb565b9497909650939450505050565b60008060408385031215610d1557600080fd5b50508035926020909101359150565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b87815286602082015285604082015284606082015260c060808201526000610d7b60c083018587610d24565b905067ffffffffffffffff831660a083015298975050505050505050565b84815260208082018590526060604083018190528201839052600090849060808401835b86811015610df8578335610dd081610b75565b73ffffffffffffffffffffffffffffffffffffffff1682529282019290820190600101610dbd565b5098975050505050505050565b600060208284031215610e1757600080fd5b5051919050565b86815285602082015260a060408201526000610e3e60a083018688610d24565b67ffffffffffffffff9490941660608301525060800152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818103610ea957610ea9610e5d565b6001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715610f0657610f06610eb3565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610f3557610f35610eb3565b604052919050565b600082601f830112610f4e57600080fd5b813567ffffffffffffffff811115610f6857610f68610eb3565b610f7b6020601f19601f84011601610f0c565b818152846020838601011115610f9057600080fd5b816020850160208301376000918101602001919091529392505050565b67ffffffffffffffff81168114610a9157600080fd5b8035610fce81610fad565b919050565b600060208284031215610fe557600080fd5b813567ffffffffffffffff80821115610ffd57600080fd5b90830190610100828603121561101257600080fd5b61101a610ee2565b8235815260208301356020820152604083013560408201526060830135606082015260808301358281111561104e57600080fd5b61105a87828601610f3d565b60808301525061106c60a08401610fc3565b60a082015260c083013560c082015260e083013560e082015280935050505092915050565b86815260006020878184015286604084015285606084015260c0608084015284518060c085015260005b818110156110d75786810183015185820160e0015282016110bb565b50600060e0828601015260e0601f19601f8301168501019250505061110860a083018467ffffffffffffffff169052565b979650505050505050565b6000602080838503121561112657600080fd5b825167ffffffffffffffff8082111561113e57600080fd5b818501915085601f83011261115257600080fd5b81518181111561116457611164610eb3565b8060051b9150611175848301610f0c565b818152918301840191848101908884111561118f57600080fd5b938501935b838510156111b957845192506111a983610b75565b8282529385019390850190611194565b98975050505050505050565b6000602082840312156111d757600080fd5b8151610bb481610fad565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361124257611242610e5d565b5060010190565b8082018082111561125c5761125c610e5d565b9291505056fea2646970667358221220deb17baf20f7e8ce8635f4a7b487f8fccc999d99fa0ac55fadf98d9ba40105ae64736f6c63430008140033",
}

// InterchainClientV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainClientV1MetaData.ABI instead.
var InterchainClientV1ABI = InterchainClientV1MetaData.ABI

// Deprecated: Use InterchainClientV1MetaData.Sigs instead.
// InterchainClientV1FuncSigs maps the 4-byte function signature to its string representation.
var InterchainClientV1FuncSigs = InterchainClientV1MetaData.Sigs

// InterchainClientV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainClientV1MetaData.Bin instead.
var InterchainClientV1Bin = InterchainClientV1MetaData.Bin

// DeployInterchainClientV1 deploys a new Ethereum contract, binding an instance of InterchainClientV1 to it.
func DeployInterchainClientV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainClientV1, error) {
	parsed, err := InterchainClientV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainClientV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainClientV1{InterchainClientV1Caller: InterchainClientV1Caller{contract: contract}, InterchainClientV1Transactor: InterchainClientV1Transactor{contract: contract}, InterchainClientV1Filterer: InterchainClientV1Filterer{contract: contract}}, nil
}

// InterchainClientV1 is an auto generated Go binding around an Ethereum contract.
type InterchainClientV1 struct {
	InterchainClientV1Caller     // Read-only binding to the contract
	InterchainClientV1Transactor // Write-only binding to the contract
	InterchainClientV1Filterer   // Log filterer for contract events
}

// InterchainClientV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainClientV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainClientV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainClientV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainClientV1Session struct {
	Contract     *InterchainClientV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterchainClientV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainClientV1CallerSession struct {
	Contract *InterchainClientV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterchainClientV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainClientV1TransactorSession struct {
	Contract     *InterchainClientV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterchainClientV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainClientV1Raw struct {
	Contract *InterchainClientV1 // Generic contract binding to access the raw methods on
}

// InterchainClientV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainClientV1CallerRaw struct {
	Contract *InterchainClientV1Caller // Generic read-only contract binding to access the raw methods on
}

// InterchainClientV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainClientV1TransactorRaw struct {
	Contract *InterchainClientV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainClientV1 creates a new instance of InterchainClientV1, bound to a specific deployed contract.
func NewInterchainClientV1(address common.Address, backend bind.ContractBackend) (*InterchainClientV1, error) {
	contract, err := bindInterchainClientV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1{InterchainClientV1Caller: InterchainClientV1Caller{contract: contract}, InterchainClientV1Transactor: InterchainClientV1Transactor{contract: contract}, InterchainClientV1Filterer: InterchainClientV1Filterer{contract: contract}}, nil
}

// NewInterchainClientV1Caller creates a new read-only instance of InterchainClientV1, bound to a specific deployed contract.
func NewInterchainClientV1Caller(address common.Address, caller bind.ContractCaller) (*InterchainClientV1Caller, error) {
	contract, err := bindInterchainClientV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1Caller{contract: contract}, nil
}

// NewInterchainClientV1Transactor creates a new write-only instance of InterchainClientV1, bound to a specific deployed contract.
func NewInterchainClientV1Transactor(address common.Address, transactor bind.ContractTransactor) (*InterchainClientV1Transactor, error) {
	contract, err := bindInterchainClientV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1Transactor{contract: contract}, nil
}

// NewInterchainClientV1Filterer creates a new log filterer instance of InterchainClientV1, bound to a specific deployed contract.
func NewInterchainClientV1Filterer(address common.Address, filterer bind.ContractFilterer) (*InterchainClientV1Filterer, error) {
	contract, err := bindInterchainClientV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1Filterer{contract: contract}, nil
}

// bindInterchainClientV1 binds a generic wrapper to an already deployed contract.
func bindInterchainClientV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainClientV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainClientV1 *InterchainClientV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainClientV1.Contract.InterchainClientV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainClientV1 *InterchainClientV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainClientV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainClientV1 *InterchainClientV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainClientV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainClientV1 *InterchainClientV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainClientV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainClientV1 *InterchainClientV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainClientV1 *InterchainClientV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.contract.Transact(opts, method, params...)
}

// ClientNonce is a free data retrieval call binding the contract method 0x0d898416.
//
// Solidity: function clientNonce() view returns(uint64)
func (_InterchainClientV1 *InterchainClientV1Caller) ClientNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "clientNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ClientNonce is a free data retrieval call binding the contract method 0x0d898416.
//
// Solidity: function clientNonce() view returns(uint64)
func (_InterchainClientV1 *InterchainClientV1Session) ClientNonce() (uint64, error) {
	return _InterchainClientV1.Contract.ClientNonce(&_InterchainClientV1.CallOpts)
}

// ClientNonce is a free data retrieval call binding the contract method 0x0d898416.
//
// Solidity: function clientNonce() view returns(uint64)
func (_InterchainClientV1 *InterchainClientV1CallerSession) ClientNonce() (uint64, error) {
	return _InterchainClientV1.Contract.ClientNonce(&_InterchainClientV1.CallOpts)
}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_InterchainClientV1 *InterchainClientV1Caller) ConvertAddressToBytes32(opts *bind.CallOpts, _address common.Address) ([32]byte, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "convertAddressToBytes32", _address)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_InterchainClientV1 *InterchainClientV1Session) ConvertAddressToBytes32(_address common.Address) ([32]byte, error) {
	return _InterchainClientV1.Contract.ConvertAddressToBytes32(&_InterchainClientV1.CallOpts, _address)
}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_InterchainClientV1 *InterchainClientV1CallerSession) ConvertAddressToBytes32(_address common.Address) ([32]byte, error) {
	return _InterchainClientV1.Contract.ConvertAddressToBytes32(&_InterchainClientV1.CallOpts, _address)
}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_InterchainClientV1 *InterchainClientV1Caller) ConvertBytes32ToAddress(opts *bind.CallOpts, _bytes32 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "convertBytes32ToAddress", _bytes32)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_InterchainClientV1 *InterchainClientV1Session) ConvertBytes32ToAddress(_bytes32 [32]byte) (common.Address, error) {
	return _InterchainClientV1.Contract.ConvertBytes32ToAddress(&_InterchainClientV1.CallOpts, _bytes32)
}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_InterchainClientV1 *InterchainClientV1CallerSession) ConvertBytes32ToAddress(_bytes32 [32]byte) (common.Address, error) {
	return _InterchainClientV1.Contract.ConvertBytes32ToAddress(&_InterchainClientV1.CallOpts, _bytes32)
}

// InterchainDB is a free data retrieval call binding the contract method 0x0e785ce0.
//
// Solidity: function interchainDB() view returns(address)
func (_InterchainClientV1 *InterchainClientV1Caller) InterchainDB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "interchainDB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainDB is a free data retrieval call binding the contract method 0x0e785ce0.
//
// Solidity: function interchainDB() view returns(address)
func (_InterchainClientV1 *InterchainClientV1Session) InterchainDB() (common.Address, error) {
	return _InterchainClientV1.Contract.InterchainDB(&_InterchainClientV1.CallOpts)
}

// InterchainDB is a free data retrieval call binding the contract method 0x0e785ce0.
//
// Solidity: function interchainDB() view returns(address)
func (_InterchainClientV1 *InterchainClientV1CallerSession) InterchainDB() (common.Address, error) {
	return _InterchainClientV1.Contract.InterchainDB(&_InterchainClientV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainClientV1 *InterchainClientV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainClientV1 *InterchainClientV1Session) Owner() (common.Address, error) {
	return _InterchainClientV1.Contract.Owner(&_InterchainClientV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainClientV1 *InterchainClientV1CallerSession) Owner() (common.Address, error) {
	return _InterchainClientV1.Contract.Owner(&_InterchainClientV1.CallOpts)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0xd3388b80.
//
// Solidity: function interchainExecute(bytes32 transactionID, bytes transaction) returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) InterchainExecute(opts *bind.TransactOpts, transactionID [32]byte, transaction []byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "interchainExecute", transactionID, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0xd3388b80.
//
// Solidity: function interchainExecute(bytes32 transactionID, bytes transaction) returns()
func (_InterchainClientV1 *InterchainClientV1Session) InterchainExecute(transactionID [32]byte, transaction []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainExecute(&_InterchainClientV1.TransactOpts, transactionID, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0xd3388b80.
//
// Solidity: function interchainExecute(bytes32 transactionID, bytes transaction) returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) InterchainExecute(transactionID [32]byte, transaction []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainExecute(&_InterchainClientV1.TransactOpts, transactionID, transaction)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] srcModules) payable returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) InterchainSend(opts *bind.TransactOpts, receiver [32]byte, dstChainId *big.Int, message []byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "interchainSend", receiver, dstChainId, message, srcModules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] srcModules) payable returns()
func (_InterchainClientV1 *InterchainClientV1Session) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSend(&_InterchainClientV1.TransactOpts, receiver, dstChainId, message, srcModules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] srcModules) payable returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSend(&_InterchainClientV1.TransactOpts, receiver, dstChainId, message, srcModules)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainClientV1 *InterchainClientV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _InterchainClientV1.Contract.RenounceOwnership(&_InterchainClientV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterchainClientV1.Contract.RenounceOwnership(&_InterchainClientV1.TransactOpts)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) SetInterchainDB(opts *bind.TransactOpts, _interchainDB common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "setInterchainDB", _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_InterchainClientV1 *InterchainClientV1Session) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.SetInterchainDB(&_InterchainClientV1.TransactOpts, _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.SetInterchainDB(&_InterchainClientV1.TransactOpts, _interchainDB)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) SetLinkedClient(opts *bind.TransactOpts, chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "setLinkedClient", chainId, client)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_InterchainClientV1 *InterchainClientV1Session) SetLinkedClient(chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.SetLinkedClient(&_InterchainClientV1.TransactOpts, chainId, client)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) SetLinkedClient(chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.SetLinkedClient(&_InterchainClientV1.TransactOpts, chainId, client)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainClientV1 *InterchainClientV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.TransferOwnership(&_InterchainClientV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.TransferOwnership(&_InterchainClientV1.TransactOpts, newOwner)
}

// InterchainClientV1InterchainTransactionSentIterator is returned from FilterInterchainTransactionSent and is used to iterate over the raw logs and unpacked data for InterchainTransactionSent events raised by the InterchainClientV1 contract.
type InterchainClientV1InterchainTransactionSentIterator struct {
	Event *InterchainClientV1InterchainTransactionSent // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1InterchainTransactionSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1InterchainTransactionSent)
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
		it.Event = new(InterchainClientV1InterchainTransactionSent)
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
func (it *InterchainClientV1InterchainTransactionSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1InterchainTransactionSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1InterchainTransactionSent represents a InterchainTransactionSent event raised by the InterchainClientV1 contract.
type InterchainClientV1InterchainTransactionSent struct {
	SrcSender     [32]byte
	SrcChainId    *big.Int
	DstReceiver   [32]byte
	DstChainId    *big.Int
	Message       []byte
	Nonce         uint64
	TransactionId [32]byte
	DbWriterNonce *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionSent is a free log retrieval operation binding the contract event 0xf679da022706ad89d2486ce3bf6cba510d619af7c1b2f6799637ed085f9287d1.
//
// Solidity: event InterchainTransactionSent(bytes32 srcSender, uint256 srcChainId, bytes32 indexed dstReceiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes32 indexed transactionId, uint256 dbWriterNonce)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterInterchainTransactionSent(opts *bind.FilterOpts, dstReceiver [][32]byte, dstChainId []*big.Int, transactionId [][32]byte) (*InterchainClientV1InterchainTransactionSentIterator, error) {

	var dstReceiverRule []interface{}
	for _, dstReceiverItem := range dstReceiver {
		dstReceiverRule = append(dstReceiverRule, dstReceiverItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "InterchainTransactionSent", dstReceiverRule, dstChainIdRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1InterchainTransactionSentIterator{contract: _InterchainClientV1.contract, event: "InterchainTransactionSent", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionSent is a free log subscription operation binding the contract event 0xf679da022706ad89d2486ce3bf6cba510d619af7c1b2f6799637ed085f9287d1.
//
// Solidity: event InterchainTransactionSent(bytes32 srcSender, uint256 srcChainId, bytes32 indexed dstReceiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes32 indexed transactionId, uint256 dbWriterNonce)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchInterchainTransactionSent(opts *bind.WatchOpts, sink chan<- *InterchainClientV1InterchainTransactionSent, dstReceiver [][32]byte, dstChainId []*big.Int, transactionId [][32]byte) (event.Subscription, error) {

	var dstReceiverRule []interface{}
	for _, dstReceiverItem := range dstReceiver {
		dstReceiverRule = append(dstReceiverRule, dstReceiverItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "InterchainTransactionSent", dstReceiverRule, dstChainIdRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1InterchainTransactionSent)
				if err := _InterchainClientV1.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
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

// ParseInterchainTransactionSent is a log parse operation binding the contract event 0xf679da022706ad89d2486ce3bf6cba510d619af7c1b2f6799637ed085f9287d1.
//
// Solidity: event InterchainTransactionSent(bytes32 srcSender, uint256 srcChainId, bytes32 indexed dstReceiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes32 indexed transactionId, uint256 dbWriterNonce)
func (_InterchainClientV1 *InterchainClientV1Filterer) ParseInterchainTransactionSent(log types.Log) (*InterchainClientV1InterchainTransactionSent, error) {
	event := new(InterchainClientV1InterchainTransactionSent)
	if err := _InterchainClientV1.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InterchainClientV1 contract.
type InterchainClientV1OwnershipTransferredIterator struct {
	Event *InterchainClientV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1OwnershipTransferred)
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
		it.Event = new(InterchainClientV1OwnershipTransferred)
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
func (it *InterchainClientV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1OwnershipTransferred represents a OwnershipTransferred event raised by the InterchainClientV1 contract.
type InterchainClientV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InterchainClientV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1OwnershipTransferredIterator{contract: _InterchainClientV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InterchainClientV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1OwnershipTransferred)
				if err := _InterchainClientV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InterchainClientV1 *InterchainClientV1Filterer) ParseOwnershipTransferred(log types.Log) (*InterchainClientV1OwnershipTransferred, error) {
	event := new(InterchainClientV1OwnershipTransferred)
	if err := _InterchainClientV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a95898eae832fabe134bfc065c874495ab528dcf6f2ea618092430b693351e4c64736f6c63430008140033",
}

// InterchainEntryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainEntryLibMetaData.ABI instead.
var InterchainEntryLibABI = InterchainEntryLibMetaData.ABI

// InterchainEntryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainEntryLibMetaData.Bin instead.
var InterchainEntryLibBin = InterchainEntryLibMetaData.Bin

// DeployInterchainEntryLib deploys a new Ethereum contract, binding an instance of InterchainEntryLib to it.
func DeployInterchainEntryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainEntryLib, error) {
	parsed, err := InterchainEntryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainEntryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainEntryLib{InterchainEntryLibCaller: InterchainEntryLibCaller{contract: contract}, InterchainEntryLibTransactor: InterchainEntryLibTransactor{contract: contract}, InterchainEntryLibFilterer: InterchainEntryLibFilterer{contract: contract}}, nil
}

// InterchainEntryLib is an auto generated Go binding around an Ethereum contract.
type InterchainEntryLib struct {
	InterchainEntryLibCaller     // Read-only binding to the contract
	InterchainEntryLibTransactor // Write-only binding to the contract
	InterchainEntryLibFilterer   // Log filterer for contract events
}

// InterchainEntryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainEntryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainEntryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainEntryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainEntryLibSession struct {
	Contract     *InterchainEntryLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterchainEntryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainEntryLibCallerSession struct {
	Contract *InterchainEntryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterchainEntryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainEntryLibTransactorSession struct {
	Contract     *InterchainEntryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterchainEntryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainEntryLibRaw struct {
	Contract *InterchainEntryLib // Generic contract binding to access the raw methods on
}

// InterchainEntryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainEntryLibCallerRaw struct {
	Contract *InterchainEntryLibCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainEntryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainEntryLibTransactorRaw struct {
	Contract *InterchainEntryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainEntryLib creates a new instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLib(address common.Address, backend bind.ContractBackend) (*InterchainEntryLib, error) {
	contract, err := bindInterchainEntryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLib{InterchainEntryLibCaller: InterchainEntryLibCaller{contract: contract}, InterchainEntryLibTransactor: InterchainEntryLibTransactor{contract: contract}, InterchainEntryLibFilterer: InterchainEntryLibFilterer{contract: contract}}, nil
}

// NewInterchainEntryLibCaller creates a new read-only instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibCaller(address common.Address, caller bind.ContractCaller) (*InterchainEntryLibCaller, error) {
	contract, err := bindInterchainEntryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibCaller{contract: contract}, nil
}

// NewInterchainEntryLibTransactor creates a new write-only instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainEntryLibTransactor, error) {
	contract, err := bindInterchainEntryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibTransactor{contract: contract}, nil
}

// NewInterchainEntryLibFilterer creates a new log filterer instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainEntryLibFilterer, error) {
	contract, err := bindInterchainEntryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibFilterer{contract: contract}, nil
}

// bindInterchainEntryLib binds a generic wrapper to an already deployed contract.
func bindInterchainEntryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainEntryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainEntryLib *InterchainEntryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainEntryLib.Contract.InterchainEntryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainEntryLib *InterchainEntryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.InterchainEntryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainEntryLib *InterchainEntryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.InterchainEntryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainEntryLib *InterchainEntryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainEntryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainEntryLib *InterchainEntryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainEntryLib *InterchainEntryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.contract.Transact(opts, method, params...)
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122051a177649b9982e67f24164c71b3c0e94b78eaf8ea59631a67835bf49d39d38264736f6c63430008140033",
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
	parsed, err := TypeCastsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
