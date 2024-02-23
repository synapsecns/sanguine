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

// AppConfigLibMetaData contains all meta data concerning the AppConfigLib contract.
var AppConfigLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"AppConfigLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200142d42fccb109a3d975169ccdb6375a4d1fb7c05b92cf0bd7e217f1903fd2df64736f6c63430008140033",
}

// AppConfigLibABI is the input ABI used to generate the binding from.
// Deprecated: Use AppConfigLibMetaData.ABI instead.
var AppConfigLibABI = AppConfigLibMetaData.ABI

// AppConfigLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AppConfigLibMetaData.Bin instead.
var AppConfigLibBin = AppConfigLibMetaData.Bin

// DeployAppConfigLib deploys a new Ethereum contract, binding an instance of AppConfigLib to it.
func DeployAppConfigLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AppConfigLib, error) {
	parsed, err := AppConfigLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AppConfigLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AppConfigLib{AppConfigLibCaller: AppConfigLibCaller{contract: contract}, AppConfigLibTransactor: AppConfigLibTransactor{contract: contract}, AppConfigLibFilterer: AppConfigLibFilterer{contract: contract}}, nil
}

// AppConfigLib is an auto generated Go binding around an Ethereum contract.
type AppConfigLib struct {
	AppConfigLibCaller     // Read-only binding to the contract
	AppConfigLibTransactor // Write-only binding to the contract
	AppConfigLibFilterer   // Log filterer for contract events
}

// AppConfigLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppConfigLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppConfigLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppConfigLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppConfigLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppConfigLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppConfigLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppConfigLibSession struct {
	Contract     *AppConfigLib     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppConfigLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppConfigLibCallerSession struct {
	Contract *AppConfigLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AppConfigLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppConfigLibTransactorSession struct {
	Contract     *AppConfigLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AppConfigLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppConfigLibRaw struct {
	Contract *AppConfigLib // Generic contract binding to access the raw methods on
}

// AppConfigLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppConfigLibCallerRaw struct {
	Contract *AppConfigLibCaller // Generic read-only contract binding to access the raw methods on
}

// AppConfigLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppConfigLibTransactorRaw struct {
	Contract *AppConfigLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppConfigLib creates a new instance of AppConfigLib, bound to a specific deployed contract.
func NewAppConfigLib(address common.Address, backend bind.ContractBackend) (*AppConfigLib, error) {
	contract, err := bindAppConfigLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppConfigLib{AppConfigLibCaller: AppConfigLibCaller{contract: contract}, AppConfigLibTransactor: AppConfigLibTransactor{contract: contract}, AppConfigLibFilterer: AppConfigLibFilterer{contract: contract}}, nil
}

// NewAppConfigLibCaller creates a new read-only instance of AppConfigLib, bound to a specific deployed contract.
func NewAppConfigLibCaller(address common.Address, caller bind.ContractCaller) (*AppConfigLibCaller, error) {
	contract, err := bindAppConfigLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppConfigLibCaller{contract: contract}, nil
}

// NewAppConfigLibTransactor creates a new write-only instance of AppConfigLib, bound to a specific deployed contract.
func NewAppConfigLibTransactor(address common.Address, transactor bind.ContractTransactor) (*AppConfigLibTransactor, error) {
	contract, err := bindAppConfigLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppConfigLibTransactor{contract: contract}, nil
}

// NewAppConfigLibFilterer creates a new log filterer instance of AppConfigLib, bound to a specific deployed contract.
func NewAppConfigLibFilterer(address common.Address, filterer bind.ContractFilterer) (*AppConfigLibFilterer, error) {
	contract, err := bindAppConfigLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppConfigLibFilterer{contract: contract}, nil
}

// bindAppConfigLib binds a generic wrapper to an already deployed contract.
func bindAppConfigLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppConfigLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppConfigLib *AppConfigLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppConfigLib.Contract.AppConfigLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppConfigLib *AppConfigLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppConfigLib.Contract.AppConfigLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppConfigLib *AppConfigLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppConfigLib.Contract.AppConfigLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppConfigLib *AppConfigLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppConfigLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppConfigLib *AppConfigLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppConfigLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppConfigLib *AppConfigLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppConfigLib.Contract.contract.Transact(opts, method, params...)
}

// IInterchainAppMetaData contains all meta data concerning the IInterchainApp contract.
var IInterchainAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"getLinkedIApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"chainIDs\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"linkedIApps\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sendingModules\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"receivingModules\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"optimisticTimePeriod\",\"type\":\"uint64\"}],\"name\":\"setAppConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d1eb8bde": "appReceive(uint256,bytes32,uint64,bytes)",
		"bfc849ee": "getLinkedIApp(uint64)",
		"287bc057": "getReceivingConfig()",
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

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainApp *IInterchainAppCaller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _IInterchainApp.contract.Call(opts, &out, "getReceivingConfig")

	outstruct := new(struct {
		AppConfig []byte
		Modules   []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppConfig = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Modules = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainApp *IInterchainAppSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _IInterchainApp.Contract.GetReceivingConfig(&_IInterchainApp.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainApp *IInterchainAppCallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _IInterchainApp.Contract.GetReceivingConfig(&_IInterchainApp.CallOpts)
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

// AppReceive is a paid mutator transaction binding the contract method 0xd1eb8bde.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, nonce uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.contract.Transact(opts, "appReceive", srcChainId, sender, nonce, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0xd1eb8bde.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppSession) AppReceive(srcChainId *big.Int, sender [32]byte, nonce uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts, srcChainId, sender, nonce, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0xd1eb8bde.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, nonce uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts, srcChainId, sender, nonce, message)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"getLinkedIApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receivingModules\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"chainIDs\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"linkedIApps\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sendingModules\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_receivingModules\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"optimisticTimePeriod\",\"type\":\"uint64\"}],\"name\":\"setAppConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receivingModule\",\"type\":\"address\"}],\"name\":\"setReceivingModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d1eb8bde": "appReceive(uint256,bytes32,uint64,bytes)",
		"bfc849ee": "getLinkedIApp(uint64)",
		"287bc057": "getReceivingConfig()",
		"ea13398f": "getSendingModules()",
		"e079da63": "receivingModules(uint256)",
		"e1ef3b3f": "send(bytes32,uint256,bytes)",
		"dd34f56a": "setAppConfig(uint64[],address[],address[],address[],uint256,uint64)",
		"92c2f0c3": "setReceivingModule(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610731806100206000396000f3fe60806040526004361061007b5760003560e01c8063dd34f56a1161004e578063dd34f56a14610193578063e079da63146101b6578063e1ef3b3f146101d6578063ea13398f146101ea57600080fd5b8063287bc0571461008057806392c2f0c3146100a7578063bfc849ee14610138578063d1eb8bde1461017e575b600080fd5b34801561008c57600080fd5b5060608060405161009e92919061028e565b60405180910390f35b3480156100b357600080fd5b506101366100c236600461033a565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b005b34801561014457600080fd5b50610159610153366004610374565b50600090565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161009e565b61013661018c3660046103d8565b5050505050565b34801561019f57600080fd5b506101366101ae3660046105b6565b505050505050565b3480156101c257600080fd5b506101596101d136600461067c565b610206565b6101366101e4366004610695565b50505050565b3480156101f657600080fd5b50606060405161009e91906106e8565b6000818154811061021657600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b600081518084526020808501945080840160005b8381101561028357815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101610251565b509495945050505050565b604081526000835180604084015260005b818110156102bc576020818701810151606086840101520161029f565b506000606082850101527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f820116830190506060838203016020840152610308606082018561023d565b95945050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461033557600080fd5b919050565b60006020828403121561034c57600080fd5b61035582610311565b9392505050565b803567ffffffffffffffff8116811461033557600080fd5b60006020828403121561038657600080fd5b6103558261035c565b60008083601f8401126103a157600080fd5b50813567ffffffffffffffff8111156103b957600080fd5b6020830191508360208285010111156103d157600080fd5b9250929050565b6000806000806000608086880312156103f057600080fd5b85359450602086013593506104076040870161035c565b9250606086013567ffffffffffffffff81111561042357600080fd5b61042f8882890161038f565b969995985093965092949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156104b6576104b6610440565b604052919050565b600067ffffffffffffffff8211156104d8576104d8610440565b5060051b60200190565b600082601f8301126104f357600080fd5b81356020610508610503836104be565b61046f565b82815260059290921b8401810191818101908684111561052757600080fd5b8286015b848110156105495761053c8161035c565b835291830191830161052b565b509695505050505050565b600082601f83011261056557600080fd5b81356020610575610503836104be565b82815260059290921b8401810191818101908684111561059457600080fd5b8286015b84811015610549576105a981610311565b8352918301918301610598565b60008060008060008060c087890312156105cf57600080fd5b863567ffffffffffffffff808211156105e757600080fd5b6105f38a838b016104e2565b9750602089013591508082111561060957600080fd5b6106158a838b01610554565b9650604089013591508082111561062b57600080fd5b6106378a838b01610554565b9550606089013591508082111561064d57600080fd5b5061065a89828a01610554565b9350506080870135915061067060a0880161035c565b90509295509295509295565b60006020828403121561068e57600080fd5b5035919050565b600080600080606085870312156106ab57600080fd5b8435935060208501359250604085013567ffffffffffffffff8111156106d057600080fd5b6106dc8782880161038f565b95989497509550505050565b602081526000610355602083018461023d56fea2646970667358221220c2dde84297266b9295de690effb1f6971f3ecd8430f5e1ac23c0b8e8d1b577f264736f6c63430008140033",
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

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_InterchainAppMock *InterchainAppMockCaller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _InterchainAppMock.contract.Call(opts, &out, "getReceivingConfig")

	outstruct := new(struct {
		AppConfig []byte
		Modules   []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppConfig = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Modules = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_InterchainAppMock *InterchainAppMockSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _InterchainAppMock.Contract.GetReceivingConfig(&_InterchainAppMock.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_InterchainAppMock *InterchainAppMockCallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _InterchainAppMock.Contract.GetReceivingConfig(&_InterchainAppMock.CallOpts)
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

// AppReceive is a paid mutator transaction binding the contract method 0xd1eb8bde.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes message) payable returns()
func (_InterchainAppMock *InterchainAppMockTransactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, nonce uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppMock.contract.Transact(opts, "appReceive", srcChainId, sender, nonce, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0xd1eb8bde.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes message) payable returns()
func (_InterchainAppMock *InterchainAppMockSession) AppReceive(srcChainId *big.Int, sender [32]byte, nonce uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.AppReceive(&_InterchainAppMock.TransactOpts, srcChainId, sender, nonce, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0xd1eb8bde.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes message) payable returns()
func (_InterchainAppMock *InterchainAppMockTransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, nonce uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppMock.Contract.AppReceive(&_InterchainAppMock.TransactOpts, srcChainId, sender, nonce, message)
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
