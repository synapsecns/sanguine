// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interchainappmock

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

// InterchainAppMockMetaData contains all meta data concerning the InterchainAppMock contract.
var InterchainAppMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"getLinkedIApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptimisticTimePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequiredResponses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receivingModules\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"chainIDs\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"linkedIApps\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sendingModules\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_receivingModules\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"optimisticTimePeriod\",\"type\":\"uint64\"}],\"name\":\"setAppConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receivingModule\",\"type\":\"address\"}],\"name\":\"setReceivingModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7bcad630": "appReceive()",
		"bfc849ee": "getLinkedIApp(uint64)",
		"7c9abd3e": "getOptimisticTimePeriod()",
		"a45e107a": "getReceivingModules()",
		"f31b19a9": "getRequiredResponses()",
		"ea13398f": "getSendingModules()",
		"e079da63": "receivingModules(uint256)",
		"e1ef3b3f": "send(bytes32,uint256,bytes)",
		"dd34f56a": "setAppConfig(uint64[],address[],address[],address[],uint256,uint64)",
		"92c2f0c3": "setReceivingModule(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506106ee806100206000396000f3fe6080604052600436106100b15760003560e01c8063dd34f56a11610069578063e1ef3b3f1161004e578063e1ef3b3f14610221578063ea13398f14610235578063f31b19a91461024957600080fd5b8063dd34f56a146101de578063e079da631461020157600080fd5b806392c2f0c31161009a57806392c2f0c3146100e7578063a45e107a14610176578063bfc849ee1461019857600080fd5b80637bcad630146100b65780637c9abd3e146100c6575b600080fd5b3480156100c257600080fd5b505b005b3480156100d257600080fd5b50604051600081526020015b60405180910390f35b3480156100f357600080fd5b506100c4610102366004610334565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b34801561018257600080fd5b5061018b610265565b6040516100de9190610356565b3480156101a457600080fd5b506101b96101b33660046103c8565b50600090565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100de565b3480156101ea57600080fd5b506100c46101f9366004610559565b505050505050565b34801561020d57600080fd5b506101b961021c36600461061f565b6102d4565b6100c461022f366004610638565b50505050565b34801561024157600080fd5b50606061018b565b34801561025557600080fd5b50604051600181526020016100de565b606060008054806020026020016040519081016040528092919081815260200182805480156102ca57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161029f575b5050505050905090565b600081815481106102e457600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b803573ffffffffffffffffffffffffffffffffffffffff8116811461032f57600080fd5b919050565b60006020828403121561034657600080fd5b61034f8261030b565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156103a457835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101610372565b50909695505050505050565b803567ffffffffffffffff8116811461032f57600080fd5b6000602082840312156103da57600080fd5b61034f826103b0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610459576104596103e3565b604052919050565b600067ffffffffffffffff82111561047b5761047b6103e3565b5060051b60200190565b600082601f83011261049657600080fd5b813560206104ab6104a683610461565b610412565b82815260059290921b840181019181810190868411156104ca57600080fd5b8286015b848110156104ec576104df816103b0565b83529183019183016104ce565b509695505050505050565b600082601f83011261050857600080fd5b813560206105186104a683610461565b82815260059290921b8401810191818101908684111561053757600080fd5b8286015b848110156104ec5761054c8161030b565b835291830191830161053b565b60008060008060008060c0878903121561057257600080fd5b863567ffffffffffffffff8082111561058a57600080fd5b6105968a838b01610485565b975060208901359150808211156105ac57600080fd5b6105b88a838b016104f7565b965060408901359150808211156105ce57600080fd5b6105da8a838b016104f7565b955060608901359150808211156105f057600080fd5b506105fd89828a016104f7565b9350506080870135915061061360a088016103b0565b90509295509295509295565b60006020828403121561063157600080fd5b5035919050565b6000806000806060858703121561064e57600080fd5b8435935060208501359250604085013567ffffffffffffffff8082111561067457600080fd5b818701915087601f83011261068857600080fd5b81358181111561069757600080fd5b8860208285010111156106a957600080fd5b9598949750506020019450505056fea2646970667358221220dac0e43b3ffe004f8036018fbf75781d864eb7897b2b9ea4ed094cd0b6a328d464736f6c63430008140033",
}

// InterchainAppMockABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainAppMockMetaData.ABI instead.
var InterchainAppMockABI = InterchainAppMockMetaData.ABI

// Deprecated: Use InterchainAppMockMetaData.Sigs instead.
// InterchainAppMockFuncSigs maps the 4-byte function signature to its string representation.
var InterchainAppMockFuncSigs = InterchainAppMockMetaData.Sigs

// InterchainAppMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainAppMockMetaData.Bin instead.
var InterchainAppMockBin = InterchainAppMockMetaData.Bin

// DeployInterchainAppMock deploys a new Ethereum contract, binding an instance of InterchainAppMock to it.
func DeployInterchainAppMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainAppMock, error) {
	parsed, err := InterchainAppMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainAppMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainAppMock{InterchainAppMockCaller: InterchainAppMockCaller{contract: contract}, InterchainAppMockTransactor: InterchainAppMockTransactor{contract: contract}, InterchainAppMockFilterer: InterchainAppMockFilterer{contract: contract}}, nil
}

// InterchainAppMock is an auto generated Go binding around an Ethereum contract.
type InterchainAppMock struct {
	InterchainAppMockCaller     // Read-only binding to the contract
	InterchainAppMockTransactor // Write-only binding to the contract
	InterchainAppMockFilterer   // Log filterer for contract events
}

// InterchainAppMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainAppMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainAppMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainAppMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainAppMockSession struct {
	Contract     *InterchainAppMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InterchainAppMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainAppMockCallerSession struct {
	Contract *InterchainAppMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InterchainAppMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainAppMockTransactorSession struct {
	Contract     *InterchainAppMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InterchainAppMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainAppMockRaw struct {
	Contract *InterchainAppMock // Generic contract binding to access the raw methods on
}

// InterchainAppMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainAppMockCallerRaw struct {
	Contract *InterchainAppMockCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainAppMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainAppMockTransactorRaw struct {
	Contract *InterchainAppMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainAppMock creates a new instance of InterchainAppMock, bound to a specific deployed contract.
func NewInterchainAppMock(address common.Address, backend bind.ContractBackend) (*InterchainAppMock, error) {
	contract, err := bindInterchainAppMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainAppMock{InterchainAppMockCaller: InterchainAppMockCaller{contract: contract}, InterchainAppMockTransactor: InterchainAppMockTransactor{contract: contract}, InterchainAppMockFilterer: InterchainAppMockFilterer{contract: contract}}, nil
}

// NewInterchainAppMockCaller creates a new read-only instance of InterchainAppMock, bound to a specific deployed contract.
func NewInterchainAppMockCaller(address common.Address, caller bind.ContractCaller) (*InterchainAppMockCaller, error) {
	contract, err := bindInterchainAppMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppMockCaller{contract: contract}, nil
}

// NewInterchainAppMockTransactor creates a new write-only instance of InterchainAppMock, bound to a specific deployed contract.
func NewInterchainAppMockTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainAppMockTransactor, error) {
	contract, err := bindInterchainAppMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppMockTransactor{contract: contract}, nil
}

// NewInterchainAppMockFilterer creates a new log filterer instance of InterchainAppMock, bound to a specific deployed contract.
func NewInterchainAppMockFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainAppMockFilterer, error) {
	contract, err := bindInterchainAppMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainAppMockFilterer{contract: contract}, nil
}

// bindInterchainAppMock binds a generic wrapper to an already deployed contract.
func bindInterchainAppMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainAppMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppMock *InterchainAppMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppMock.Contract.InterchainAppMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppMock *InterchainAppMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.InterchainAppMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppMock *InterchainAppMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.InterchainAppMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppMock *InterchainAppMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppMock *InterchainAppMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppMock *InterchainAppMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.contract.Transact(opts, method, params...)
}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_InterchainAppMock *InterchainAppMockCaller) GetLinkedIApp(opts *bind.CallOpts, chainID uint64) (common.Address, error) {
	var out []interface{}
	err := _InterchainAppMock.contract.Call(opts, &out, "getLinkedIApp", chainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_InterchainAppMock *InterchainAppMockSession) GetLinkedIApp(chainID uint64) (common.Address, error) {
	return _InterchainAppMock.Contract.GetLinkedIApp(&_InterchainAppMock.CallOpts, chainID)
}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_InterchainAppMock *InterchainAppMockCallerSession) GetLinkedIApp(chainID uint64) (common.Address, error) {
	return _InterchainAppMock.Contract.GetLinkedIApp(&_InterchainAppMock.CallOpts, chainID)
}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() pure returns(uint64)
func (_InterchainAppMock *InterchainAppMockCaller) GetOptimisticTimePeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _InterchainAppMock.contract.Call(opts, &out, "getOptimisticTimePeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() pure returns(uint64)
func (_InterchainAppMock *InterchainAppMockSession) GetOptimisticTimePeriod() (uint64, error) {
	return _InterchainAppMock.Contract.GetOptimisticTimePeriod(&_InterchainAppMock.CallOpts)
}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() pure returns(uint64)
func (_InterchainAppMock *InterchainAppMockCallerSession) GetOptimisticTimePeriod() (uint64, error) {
	return _InterchainAppMock.Contract.GetOptimisticTimePeriod(&_InterchainAppMock.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppMock *InterchainAppMockCaller) GetReceivingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainAppMock.contract.Call(opts, &out, "getReceivingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppMock *InterchainAppMockSession) GetReceivingModules() ([]common.Address, error) {
	return _InterchainAppMock.Contract.GetReceivingModules(&_InterchainAppMock.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppMock *InterchainAppMockCallerSession) GetReceivingModules() ([]common.Address, error) {
	return _InterchainAppMock.Contract.GetReceivingModules(&_InterchainAppMock.CallOpts)
}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() pure returns(uint256)
func (_InterchainAppMock *InterchainAppMockCaller) GetRequiredResponses(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterchainAppMock.contract.Call(opts, &out, "getRequiredResponses")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() pure returns(uint256)
func (_InterchainAppMock *InterchainAppMockSession) GetRequiredResponses() (*big.Int, error) {
	return _InterchainAppMock.Contract.GetRequiredResponses(&_InterchainAppMock.CallOpts)
}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() pure returns(uint256)
func (_InterchainAppMock *InterchainAppMockCallerSession) GetRequiredResponses() (*big.Int, error) {
	return _InterchainAppMock.Contract.GetRequiredResponses(&_InterchainAppMock.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppMock *InterchainAppMockCaller) GetSendingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainAppMock.contract.Call(opts, &out, "getSendingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppMock *InterchainAppMockSession) GetSendingModules() ([]common.Address, error) {
	return _InterchainAppMock.Contract.GetSendingModules(&_InterchainAppMock.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppMock *InterchainAppMockCallerSession) GetSendingModules() ([]common.Address, error) {
	return _InterchainAppMock.Contract.GetSendingModules(&_InterchainAppMock.CallOpts)
}

// ReceivingModules is a free data retrieval call binding the contract method 0xe079da63.
//
// Solidity: function receivingModules(uint256 ) view returns(address)
func (_InterchainAppMock *InterchainAppMockCaller) ReceivingModules(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _InterchainAppMock.contract.Call(opts, &out, "receivingModules", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceivingModules is a free data retrieval call binding the contract method 0xe079da63.
//
// Solidity: function receivingModules(uint256 ) view returns(address)
func (_InterchainAppMock *InterchainAppMockSession) ReceivingModules(arg0 *big.Int) (common.Address, error) {
	return _InterchainAppMock.Contract.ReceivingModules(&_InterchainAppMock.CallOpts, arg0)
}

// ReceivingModules is a free data retrieval call binding the contract method 0xe079da63.
//
// Solidity: function receivingModules(uint256 ) view returns(address)
func (_InterchainAppMock *InterchainAppMockCallerSession) ReceivingModules(arg0 *big.Int) (common.Address, error) {
	return _InterchainAppMock.Contract.ReceivingModules(&_InterchainAppMock.CallOpts, arg0)
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_InterchainAppMock *InterchainAppMockTransactor) AppReceive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppMock.contract.Transact(opts, "appReceive")
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_InterchainAppMock *InterchainAppMockSession) AppReceive() (*types.Transaction, error) {
	return _InterchainAppMock.Contract.AppReceive(&_InterchainAppMock.TransactOpts)
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_InterchainAppMock *InterchainAppMockTransactorSession) AppReceive() (*types.Transaction, error) {
	return _InterchainAppMock.Contract.AppReceive(&_InterchainAppMock.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_InterchainAppMock *InterchainAppMockTransactor) Send(opts *bind.TransactOpts, receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainAppMock.contract.Transact(opts, "send", receiver, dstChainId, message)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_InterchainAppMock *InterchainAppMockSession) Send(receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.Send(&_InterchainAppMock.TransactOpts, receiver, dstChainId, message)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_InterchainAppMock *InterchainAppMockTransactorSession) Send(receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.Send(&_InterchainAppMock.TransactOpts, receiver, dstChainId, message)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] sendingModules, address[] _receivingModules, uint256 requiredResponses, uint64 optimisticTimePeriod) returns()
func (_InterchainAppMock *InterchainAppMockTransactor) SetAppConfig(opts *bind.TransactOpts, chainIDs []uint64, linkedIApps []common.Address, sendingModules []common.Address, _receivingModules []common.Address, requiredResponses *big.Int, optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _InterchainAppMock.contract.Transact(opts, "setAppConfig", chainIDs, linkedIApps, sendingModules, _receivingModules, requiredResponses, optimisticTimePeriod)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] sendingModules, address[] _receivingModules, uint256 requiredResponses, uint64 optimisticTimePeriod) returns()
func (_InterchainAppMock *InterchainAppMockSession) SetAppConfig(chainIDs []uint64, linkedIApps []common.Address, sendingModules []common.Address, _receivingModules []common.Address, requiredResponses *big.Int, optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.SetAppConfig(&_InterchainAppMock.TransactOpts, chainIDs, linkedIApps, sendingModules, _receivingModules, requiredResponses, optimisticTimePeriod)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] sendingModules, address[] _receivingModules, uint256 requiredResponses, uint64 optimisticTimePeriod) returns()
func (_InterchainAppMock *InterchainAppMockTransactorSession) SetAppConfig(chainIDs []uint64, linkedIApps []common.Address, sendingModules []common.Address, _receivingModules []common.Address, requiredResponses *big.Int, optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.SetAppConfig(&_InterchainAppMock.TransactOpts, chainIDs, linkedIApps, sendingModules, _receivingModules, requiredResponses, optimisticTimePeriod)
}

// SetReceivingModule is a paid mutator transaction binding the contract method 0x92c2f0c3.
//
// Solidity: function setReceivingModule(address _receivingModule) returns()
func (_InterchainAppMock *InterchainAppMockTransactor) SetReceivingModule(opts *bind.TransactOpts, _receivingModule common.Address) (*types.Transaction, error) {
	return _InterchainAppMock.contract.Transact(opts, "setReceivingModule", _receivingModule)
}

// SetReceivingModule is a paid mutator transaction binding the contract method 0x92c2f0c3.
//
// Solidity: function setReceivingModule(address _receivingModule) returns()
func (_InterchainAppMock *InterchainAppMockSession) SetReceivingModule(_receivingModule common.Address) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.SetReceivingModule(&_InterchainAppMock.TransactOpts, _receivingModule)
}

// SetReceivingModule is a paid mutator transaction binding the contract method 0x92c2f0c3.
//
// Solidity: function setReceivingModule(address _receivingModule) returns()
func (_InterchainAppMock *InterchainAppMockTransactorSession) SetReceivingModule(_receivingModule common.Address) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.SetReceivingModule(&_InterchainAppMock.TransactOpts, _receivingModule)
}
