// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interchainapp

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

// AppConfigV1 is an auto generated low-level Go binding around an user-defined struct.
type AppConfigV1 struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
}

// InterchainTxDescriptor is an auto generated low-level Go binding around an user-defined struct.
type InterchainTxDescriptor struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
}

// AppConfigLibMetaData contains all meta data concerning the AppConfigLib contract.
var AppConfigLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"AppConfigLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122077de10f8314ffff7df1f15f98d9e3e48b3f00f8f9482c9f5c10d397f5491854864736f6c63430008140033",
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

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122000db264a03c73894bca4a5c7dfb8c07932924528de5e8b59b532569ad817582664736f6c63430008140033",
}

// EnumerableSetABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetMetaData.ABI instead.
var EnumerableSetABI = EnumerableSetMetaData.ABI

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetMetaData.Bin instead.
var EnumerableSetBin = EnumerableSetMetaData.Bin

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// IInterchainAppMetaData contains all meta data concerning the IInterchainApp contract.
var IInterchainAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"287bc057": "getReceivingConfig()",
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

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// IInterchainClientV1MetaData contains all meta data concerning the IInterchainClientV1 contract.
var IInterchainClientV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectDstChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NoLinkedClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__NotEVMClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotEnoughResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotRemoteChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxNotExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"getExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"getExecutorById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedClient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedClientEVM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSend\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"internalType\":\"structInterchainTxDescriptor\",\"name\":\"desc\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSendEVM\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"internalType\":\"structInterchainTxDescriptor\",\"name\":\"desc\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"isExecutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionFees_\",\"type\":\"address\"}],\"name\":\"setExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"writeExecutionProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f92a79ff": "getExecutor(bytes)",
		"f1a61fac": "getExecutorById(bytes32)",
		"3c383e7b": "getInterchainFee(uint256,address,address[],bytes,bytes)",
		"aa102ec4": "getLinkedClient(uint256)",
		"02172a35": "getLinkedClientEVM(uint256)",
		"53b67d74": "interchainExecute(uint256,bytes,bytes32[])",
		"98939d28": "interchainSend(uint256,bytes32,address,address[],bytes,bytes)",
		"827f940d": "interchainSendEVM(uint256,address,address,address[],bytes,bytes)",
		"1450c281": "isExecutable(bytes,bytes32[])",
		"3dc68b87": "setExecutionFees(address)",
		"f34234c8": "setLinkedClient(uint256,bytes32)",
		"90e81077": "writeExecutionProof(bytes32)",
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

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes transaction) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetExecutor(opts *bind.CallOpts, transaction []byte) (common.Address, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getExecutor", transaction)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes transaction) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetExecutor(transaction []byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetExecutor(&_IInterchainClientV1.CallOpts, transaction)
}

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes transaction) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetExecutor(transaction []byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetExecutor(&_IInterchainClientV1.CallOpts, transaction)
}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetExecutorById(opts *bind.CallOpts, transactionId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getExecutorById", transactionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetExecutorById(transactionId [32]byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetExecutorById(&_IInterchainClientV1.CallOpts, transactionId)
}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetExecutorById(transactionId [32]byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetExecutorById(&_IInterchainClientV1.CallOpts, transactionId)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0x3c383e7b.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, bytes message) view returns(uint256)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetInterchainFee(opts *bind.CallOpts, dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcExecutionService, srcModules, options, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0x3c383e7b.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, bytes message) view returns(uint256)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetInterchainFee(dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*big.Int, error) {
	return _IInterchainClientV1.Contract.GetInterchainFee(&_IInterchainClientV1.CallOpts, dstChainId, srcExecutionService, srcModules, options, message)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0x3c383e7b.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, bytes message) view returns(uint256)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetInterchainFee(dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*big.Int, error) {
	return _IInterchainClientV1.Contract.GetInterchainFee(&_IInterchainClientV1.CallOpts, dstChainId, srcExecutionService, srcModules, options, message)
}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetLinkedClient(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getLinkedClient", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetLinkedClient(chainId *big.Int) ([32]byte, error) {
	return _IInterchainClientV1.Contract.GetLinkedClient(&_IInterchainClientV1.CallOpts, chainId)
}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetLinkedClient(chainId *big.Int) ([32]byte, error) {
	return _IInterchainClientV1.Contract.GetLinkedClient(&_IInterchainClientV1.CallOpts, chainId)
}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetLinkedClientEVM(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getLinkedClientEVM", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetLinkedClientEVM(chainId *big.Int) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetLinkedClientEVM(&_IInterchainClientV1.CallOpts, chainId)
}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetLinkedClientEVM(chainId *big.Int) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetLinkedClientEVM(&_IInterchainClientV1.CallOpts, chainId)
}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes transaction, bytes32[] proof) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1Caller) IsExecutable(opts *bind.CallOpts, transaction []byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "isExecutable", transaction, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes transaction, bytes32[] proof) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1Session) IsExecutable(transaction []byte, proof [][32]byte) (bool, error) {
	return _IInterchainClientV1.Contract.IsExecutable(&_IInterchainClientV1.CallOpts, transaction, proof)
}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes transaction, bytes32[] proof) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) IsExecutable(transaction []byte, proof [][32]byte) (bool, error) {
	return _IInterchainClientV1.Contract.IsExecutable(&_IInterchainClientV1.CallOpts, transaction, proof)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainExecute(opts *bind.TransactOpts, gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainExecute", gasLimit, transaction, proof)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainExecute(gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, gasLimit, transaction, proof)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainExecute(gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, gasLimit, transaction, proof)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainSend(opts *bind.TransactOpts, dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainSend", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainSendEVM(opts *bind.TransactOpts, dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainSendEVM", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainSendEVM(dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSendEVM(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainSendEVM(dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSendEVM(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) SetExecutionFees(opts *bind.TransactOpts, executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "setExecutionFees", executionFees_)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) SetExecutionFees(executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetExecutionFees(&_IInterchainClientV1.TransactOpts, executionFees_)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) SetExecutionFees(executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetExecutionFees(&_IInterchainClientV1.TransactOpts, executionFees_)
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

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) WriteExecutionProof(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "writeExecutionProof", transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainClientV1 *IInterchainClientV1Session) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.WriteExecutionProof(&_IInterchainClientV1.TransactOpts, transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.WriteExecutionProof(&_IInterchainClientV1.TransactOpts, transactionId)
}

// InterchainAppExampleMetaData contains all meta data concerning the InterchainAppExample contract.
var InterchainAppExampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainApp__CallerNotInterchainClient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__InterchainClientNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleNotAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__ReceiverNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"AppConfigV1Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"AppLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"ExecutionServiceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"interchainClient\",\"type\":\"address\"}],\"name\":\"InterchainClientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"addTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppConfigV1\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedApp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"isAllowedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"linkRemoteApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteApp\",\"type\":\"address\"}],\"name\":\"linkRemoteAppEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"removeTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"appConfig\",\"type\":\"tuple\"}],\"name\":\"setAppConfigV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"setExecutionService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"interchain_\",\"type\":\"address\"}],\"name\":\"setInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cb5038fb": "addTrustedModule(address)",
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"7717a647": "getAppConfigV1()",
		"a20ce510": "getAppVersion()",
		"c313c807": "getExecutionService()",
		"b9b74b18": "getLinkedApp(uint256)",
		"287bc057": "getReceivingConfig()",
		"a45e107a": "getReceivingModules()",
		"ea13398f": "getSendingModules()",
		"70838975": "interchain()",
		"dc2b9075": "isAllowedSender(uint256,bytes32)",
		"51a30802": "linkRemoteApp(uint256,bytes32)",
		"af8fcc8e": "linkRemoteAppEVM(uint256,address)",
		"8da5cb5b": "owner()",
		"b70c40b3": "removeTrustedModule(address)",
		"715018a6": "renounceOwnership()",
		"0ca709ee": "sendMessage(uint256,uint256,bytes)",
		"0d32b505": "setAppConfigV1((uint256,uint256))",
		"496774b1": "setExecutionService(address)",
		"27efcbb7": "setInterchainClient(address)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161158738038061158783398101604081905261002f916100c2565b80806001600160a01b03811661005f57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61006881610070565b5050506100f2565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000602082840312156100d457600080fd5b81516001600160a01b03811681146100eb57600080fd5b9392505050565b611486806101016000396000f3fe60806040526004361061016a5760003560e01c80638da5cb5b116100cb578063b9b74b181161007f578063dc2b907511610059578063dc2b90751461040f578063ea13398f14610335578063f2fde38b1461043f57600080fd5b8063b9b74b1814610397578063c313c807146103c4578063cb5038fb146103ef57600080fd5b8063a45e107a116100b0578063a45e107a14610335578063af8fcc8e14610357578063b70c40b31461037757600080fd5b80638da5cb5b146102ec578063a20ce5101461031757600080fd5b806351a308021161012257806370838975116101075780637083897514610243578063715018a6146102955780637717a647146102aa57600080fd5b806351a308021461021057806368a698471461023057600080fd5b806327efcbb71161015357806327efcbb7146101a4578063287bc057146101c4578063496774b1146101f057600080fd5b80630ca709ee1461016f5780630d32b50514610184575b600080fd5b61018261017d366004610f29565b61045f565b005b34801561019057600080fd5b5061018261019f366004610f7c565b61051c565b3480156101b057600080fd5b506101826101bf36600461101b565b610530565b3480156101d057600080fd5b506101d9610541565b6040516101e79291906110eb565b60405180910390f35b3480156101fc57600080fd5b5061018261020b36600461101b565b61058a565b34801561021c57600080fd5b5061018261022b366004611110565b61059b565b61018261023e366004611148565b6105b1565b34801561024f57600080fd5b506000546102709073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e7565b3480156102a157600080fd5b506101826106a5565b3480156102b657600080fd5b50604080518082018252600080825260209182015281518083019092526001548252600254908201526040516101e791906111bb565b3480156102f857600080fd5b5060075473ffffffffffffffffffffffffffffffffffffffff16610270565b34801561032357600080fd5b5060015b6040519081526020016101e7565b34801561034157600080fd5b5061034a6106b9565b6040516101e791906111d2565b34801561036357600080fd5b506101826103723660046111e5565b6106ca565b34801561038357600080fd5b5061018261039236600461101b565b6106dc565b3480156103a357600080fd5b506103276103b2366004611211565b60009081526003602052604090205490565b3480156103d057600080fd5b5060065473ffffffffffffffffffffffffffffffffffffffff16610270565b3480156103fb57600080fd5b5061018261040a36600461101b565b6106ed565b34801561041b57600080fd5b5061042f61042a366004611110565b6106fe565b60405190151581526020016101e7565b34801561044b57600080fd5b5061018261045a36600461101b565b610716565b60006104b885346040518060400160405280888152602001600081525086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061077792505050565b602080820151604080840151845182518b81529485019390935267ffffffffffffffff169083015260608201529091507f2ef16db2691a32543ce5591798c4992f4cfbbcd446874f1437d99da53d600e7c9060800160405180910390a15050505050565b6105246107bf565b61052d81610812565b50565b6105386107bf565b61052d8161085e565b60608061057a6105756040805180820182526000808252602091820152815180830190925260015482526002549082015290565b6108d1565b91506105846106b9565b90509091565b6105926107bf565b61052d816108fd565b6105a36107bf565b6105ad8282610970565b5050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610609576040517f3e336bbb0000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b468603610645576040517fbfae2eb700000000000000000000000000000000000000000000000000000000815260048101879052602401610600565b61064f86866106fe565b61068f576040517f327f41230000000000000000000000000000000000000000000000000000000081526004810187905260248101869052604401610600565b61069d8686868686866109af565b505050505050565b6106ad6107bf565b6106b760006109f8565b565b60606106c56004610a6f565b905090565b6106d26107bf565b6105ad8282610a83565b6106e46107bf565b61052d81610aa3565b6106f56107bf565b61052d81610b4f565b60008281526003602052604090205481145b92915050565b61071e6107bf565b73ffffffffffffffffffffffffffffffffffffffff811661076e576040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260006004820152602401610600565b61052d816109f8565b60408051606081018252600080825260208201819052918101919091526107b6856107ae8760009081526003602052604090205490565b868686610bf3565b95945050505050565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b7576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610600565b80516001819055602080830151600281905560408051938452918301527f156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf91015b60405180910390a150565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527feec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d790602001610853565b60606107106001836040516020016108e991906111bb565b604051602081830303815290604052610cd2565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e90602001610853565b60008281526003602052604080822083905551829184917f622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a319190a35050565b7fc5b07cf0d424748241636d3a67366b44ac88118c42bef024abcb7fd138df79a68686868686866040516109e89695949392919061122a565b60405180910390a1505050505050565b6007805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000610a7c83610cfe565b9392505050565b6105ad8273ffffffffffffffffffffffffffffffffffffffff8316610970565b6000610ab0600483610d5a565b905080610b01576040517fb12a48e600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610600565b60405173ffffffffffffffffffffffffffffffffffffffff831681527f91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385906020015b60405180910390a15050565b6000610b5c600483610d7c565b905080610bad576040517f856e38ac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610600565b60405173ffffffffffffffffffffffffffffffffffffffff831681527f0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a5011750990602001610b43565b604080516060810182526000808252602082018190529181019190915260005473ffffffffffffffffffffffffffffffffffffffff16806398939d28868989610c5160065473ffffffffffffffffffffffffffffffffffffffff1690565b610c596106b9565b610c628b6108d1565b8a6040518863ffffffff1660e01b8152600401610c849695949392919061129d565b60606040518083038185885af1158015610ca2573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610cc7919061130b565b979650505050505050565b60608282604051602001610ce7929190611394565b604051602081830303815290604052905092915050565b606081600001805480602002602001604051908101604052809291908181526020018280548015610d4e57602002820191906000526020600020905b815481526020019060010190808311610d3a575b50505050509050919050565b6000610a7c8373ffffffffffffffffffffffffffffffffffffffff8416610d9e565b6000610a7c8373ffffffffffffffffffffffffffffffffffffffff8416610e91565b60008181526001830160205260408120548015610e87576000610dc26001836113b8565b8554909150600090610dd6906001906113b8565b9050808214610e3b576000866000018281548110610df657610df66113f2565b9060005260206000200154905080876000018481548110610e1957610e196113f2565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610e4c57610e4c611421565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610710565b6000915050610710565b6000818152600183016020526040812054610ed857508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610710565b506000610710565b60008083601f840112610ef257600080fd5b50813567ffffffffffffffff811115610f0a57600080fd5b602083019150836020828501011115610f2257600080fd5b9250929050565b60008060008060608587031215610f3f57600080fd5b8435935060208501359250604085013567ffffffffffffffff811115610f6457600080fd5b610f7087828801610ee0565b95989497509550505050565b600060408284031215610f8e57600080fd5b6040516040810181811067ffffffffffffffff82111715610fd8577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052823581526020928301359281019290925250919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461101657600080fd5b919050565b60006020828403121561102d57600080fd5b610a7c82610ff2565b6000815180845260005b8181101561105c57602081850181015186830182015201611040565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b600081518084526020808501945080840160005b838110156110e057815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016110ae565b509495945050505050565b6040815260006110fe6040830185611036565b82810360208401526107b6818561109a565b6000806040838503121561112357600080fd5b50508035926020909101359150565b67ffffffffffffffff8116811461052d57600080fd5b60008060008060008060a0878903121561116157600080fd5b863595506020870135945060408701359350606087013561118181611132565b9250608087013567ffffffffffffffff81111561119d57600080fd5b6111a989828a01610ee0565b979a9699509497509295939492505050565b815181526020808301519082015260408101610710565b602081526000610a7c602083018461109a565b600080604083850312156111f857600080fd5b8235915061120860208401610ff2565b90509250929050565b60006020828403121561122357600080fd5b5035919050565b86815285602082015284604082015267ffffffffffffffff8416606082015260a060808201528160a0820152818360c0830137600081830160c090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010195945050505050565b86815285602082015273ffffffffffffffffffffffffffffffffffffffff8516604082015260c0606082015260006112d860c083018661109a565b82810360808401526112ea8186611036565b905082810360a08401526112fe8185611036565b9998505050505050505050565b60006060828403121561131d57600080fd5b6040516060810181811067ffffffffffffffff82111715611367577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b80604052508251815260208301516020820152604083015161138881611132565b60408201529392505050565b60ff831681526040602082015260006113b06040830184611036565b949350505050565b81810381811115610710577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212205d4a883d1ff7dbc74a92ec381cf7755b51ad1b3f6e3ae11207e77e1ac2f711f964736f6c63430008140033",
}

// InterchainAppExampleABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainAppExampleMetaData.ABI instead.
var InterchainAppExampleABI = InterchainAppExampleMetaData.ABI

// Deprecated: Use InterchainAppExampleMetaData.Sigs instead.
// InterchainAppExampleFuncSigs maps the 4-byte function signature to its string representation.
var InterchainAppExampleFuncSigs = InterchainAppExampleMetaData.Sigs

// InterchainAppExampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainAppExampleMetaData.Bin instead.
var InterchainAppExampleBin = InterchainAppExampleMetaData.Bin

// DeployInterchainAppExample deploys a new Ethereum contract, binding an instance of InterchainAppExample to it.
func DeployInterchainAppExample(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address) (common.Address, *types.Transaction, *InterchainAppExample, error) {
	parsed, err := InterchainAppExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainAppExampleBin), backend, owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainAppExample{InterchainAppExampleCaller: InterchainAppExampleCaller{contract: contract}, InterchainAppExampleTransactor: InterchainAppExampleTransactor{contract: contract}, InterchainAppExampleFilterer: InterchainAppExampleFilterer{contract: contract}}, nil
}

// InterchainAppExample is an auto generated Go binding around an Ethereum contract.
type InterchainAppExample struct {
	InterchainAppExampleCaller     // Read-only binding to the contract
	InterchainAppExampleTransactor // Write-only binding to the contract
	InterchainAppExampleFilterer   // Log filterer for contract events
}

// InterchainAppExampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainAppExampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppExampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainAppExampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppExampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainAppExampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppExampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainAppExampleSession struct {
	Contract     *InterchainAppExample // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InterchainAppExampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainAppExampleCallerSession struct {
	Contract *InterchainAppExampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// InterchainAppExampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainAppExampleTransactorSession struct {
	Contract     *InterchainAppExampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// InterchainAppExampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainAppExampleRaw struct {
	Contract *InterchainAppExample // Generic contract binding to access the raw methods on
}

// InterchainAppExampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainAppExampleCallerRaw struct {
	Contract *InterchainAppExampleCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainAppExampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainAppExampleTransactorRaw struct {
	Contract *InterchainAppExampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainAppExample creates a new instance of InterchainAppExample, bound to a specific deployed contract.
func NewInterchainAppExample(address common.Address, backend bind.ContractBackend) (*InterchainAppExample, error) {
	contract, err := bindInterchainAppExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainAppExample{InterchainAppExampleCaller: InterchainAppExampleCaller{contract: contract}, InterchainAppExampleTransactor: InterchainAppExampleTransactor{contract: contract}, InterchainAppExampleFilterer: InterchainAppExampleFilterer{contract: contract}}, nil
}

// NewInterchainAppExampleCaller creates a new read-only instance of InterchainAppExample, bound to a specific deployed contract.
func NewInterchainAppExampleCaller(address common.Address, caller bind.ContractCaller) (*InterchainAppExampleCaller, error) {
	contract, err := bindInterchainAppExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleCaller{contract: contract}, nil
}

// NewInterchainAppExampleTransactor creates a new write-only instance of InterchainAppExample, bound to a specific deployed contract.
func NewInterchainAppExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainAppExampleTransactor, error) {
	contract, err := bindInterchainAppExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleTransactor{contract: contract}, nil
}

// NewInterchainAppExampleFilterer creates a new log filterer instance of InterchainAppExample, bound to a specific deployed contract.
func NewInterchainAppExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainAppExampleFilterer, error) {
	contract, err := bindInterchainAppExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleFilterer{contract: contract}, nil
}

// bindInterchainAppExample binds a generic wrapper to an already deployed contract.
func bindInterchainAppExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainAppExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppExample *InterchainAppExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppExample.Contract.InterchainAppExampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppExample *InterchainAppExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.InterchainAppExampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppExample *InterchainAppExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.InterchainAppExampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppExample *InterchainAppExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppExample.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppExample *InterchainAppExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppExample *InterchainAppExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.contract.Transact(opts, method, params...)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_InterchainAppExample *InterchainAppExampleCaller) GetAppConfigV1(opts *bind.CallOpts) (AppConfigV1, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "getAppConfigV1")

	if err != nil {
		return *new(AppConfigV1), err
	}

	out0 := *abi.ConvertType(out[0], new(AppConfigV1)).(*AppConfigV1)

	return out0, err

}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_InterchainAppExample *InterchainAppExampleSession) GetAppConfigV1() (AppConfigV1, error) {
	return _InterchainAppExample.Contract.GetAppConfigV1(&_InterchainAppExample.CallOpts)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_InterchainAppExample *InterchainAppExampleCallerSession) GetAppConfigV1() (AppConfigV1, error) {
	return _InterchainAppExample.Contract.GetAppConfigV1(&_InterchainAppExample.CallOpts)
}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_InterchainAppExample *InterchainAppExampleCaller) GetAppVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "getAppVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_InterchainAppExample *InterchainAppExampleSession) GetAppVersion() (*big.Int, error) {
	return _InterchainAppExample.Contract.GetAppVersion(&_InterchainAppExample.CallOpts)
}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_InterchainAppExample *InterchainAppExampleCallerSession) GetAppVersion() (*big.Int, error) {
	return _InterchainAppExample.Contract.GetAppVersion(&_InterchainAppExample.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_InterchainAppExample *InterchainAppExampleCaller) GetExecutionService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "getExecutionService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_InterchainAppExample *InterchainAppExampleSession) GetExecutionService() (common.Address, error) {
	return _InterchainAppExample.Contract.GetExecutionService(&_InterchainAppExample.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_InterchainAppExample *InterchainAppExampleCallerSession) GetExecutionService() (common.Address, error) {
	return _InterchainAppExample.Contract.GetExecutionService(&_InterchainAppExample.CallOpts)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_InterchainAppExample *InterchainAppExampleCaller) GetLinkedApp(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "getLinkedApp", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_InterchainAppExample *InterchainAppExampleSession) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _InterchainAppExample.Contract.GetLinkedApp(&_InterchainAppExample.CallOpts, chainId)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_InterchainAppExample *InterchainAppExampleCallerSession) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _InterchainAppExample.Contract.GetLinkedApp(&_InterchainAppExample.CallOpts, chainId)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_InterchainAppExample *InterchainAppExampleCaller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "getReceivingConfig")

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
func (_InterchainAppExample *InterchainAppExampleSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _InterchainAppExample.Contract.GetReceivingConfig(&_InterchainAppExample.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_InterchainAppExample *InterchainAppExampleCallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _InterchainAppExample.Contract.GetReceivingConfig(&_InterchainAppExample.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppExample *InterchainAppExampleCaller) GetReceivingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "getReceivingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppExample *InterchainAppExampleSession) GetReceivingModules() ([]common.Address, error) {
	return _InterchainAppExample.Contract.GetReceivingModules(&_InterchainAppExample.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppExample *InterchainAppExampleCallerSession) GetReceivingModules() ([]common.Address, error) {
	return _InterchainAppExample.Contract.GetReceivingModules(&_InterchainAppExample.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppExample *InterchainAppExampleCaller) GetSendingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "getSendingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppExample *InterchainAppExampleSession) GetSendingModules() ([]common.Address, error) {
	return _InterchainAppExample.Contract.GetSendingModules(&_InterchainAppExample.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppExample *InterchainAppExampleCallerSession) GetSendingModules() ([]common.Address, error) {
	return _InterchainAppExample.Contract.GetSendingModules(&_InterchainAppExample.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainAppExample *InterchainAppExampleCaller) Interchain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "interchain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainAppExample *InterchainAppExampleSession) Interchain() (common.Address, error) {
	return _InterchainAppExample.Contract.Interchain(&_InterchainAppExample.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainAppExample *InterchainAppExampleCallerSession) Interchain() (common.Address, error) {
	return _InterchainAppExample.Contract.Interchain(&_InterchainAppExample.CallOpts)
}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_InterchainAppExample *InterchainAppExampleCaller) IsAllowedSender(opts *bind.CallOpts, srcChainId *big.Int, sender [32]byte) (bool, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "isAllowedSender", srcChainId, sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_InterchainAppExample *InterchainAppExampleSession) IsAllowedSender(srcChainId *big.Int, sender [32]byte) (bool, error) {
	return _InterchainAppExample.Contract.IsAllowedSender(&_InterchainAppExample.CallOpts, srcChainId, sender)
}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_InterchainAppExample *InterchainAppExampleCallerSession) IsAllowedSender(srcChainId *big.Int, sender [32]byte) (bool, error) {
	return _InterchainAppExample.Contract.IsAllowedSender(&_InterchainAppExample.CallOpts, srcChainId, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainAppExample *InterchainAppExampleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainAppExample.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainAppExample *InterchainAppExampleSession) Owner() (common.Address, error) {
	return _InterchainAppExample.Contract.Owner(&_InterchainAppExample.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainAppExample *InterchainAppExampleCallerSession) Owner() (common.Address, error) {
	return _InterchainAppExample.Contract.Owner(&_InterchainAppExample.CallOpts)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) AddTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "addTrustedModule", module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_InterchainAppExample *InterchainAppExampleSession) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.AddTrustedModule(&_InterchainAppExample.TransactOpts, module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.AddTrustedModule(&_InterchainAppExample.TransactOpts, module)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_InterchainAppExample *InterchainAppExampleSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.AppReceive(&_InterchainAppExample.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.AppReceive(&_InterchainAppExample.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) LinkRemoteApp(opts *bind.TransactOpts, chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "linkRemoteApp", chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_InterchainAppExample *InterchainAppExampleSession) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.LinkRemoteApp(&_InterchainAppExample.TransactOpts, chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.LinkRemoteApp(&_InterchainAppExample.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) LinkRemoteAppEVM(opts *bind.TransactOpts, chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "linkRemoteAppEVM", chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_InterchainAppExample *InterchainAppExampleSession) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.LinkRemoteAppEVM(&_InterchainAppExample.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.LinkRemoteAppEVM(&_InterchainAppExample.TransactOpts, chainId, remoteApp)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) RemoveTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "removeTrustedModule", module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_InterchainAppExample *InterchainAppExampleSession) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.RemoveTrustedModule(&_InterchainAppExample.TransactOpts, module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.RemoveTrustedModule(&_InterchainAppExample.TransactOpts, module)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainAppExample *InterchainAppExampleSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterchainAppExample.Contract.RenounceOwnership(&_InterchainAppExample.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterchainAppExample.Contract.RenounceOwnership(&_InterchainAppExample.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ca709ee.
//
// Solidity: function sendMessage(uint256 dstChainId, uint256 gasLimit, bytes message) payable returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) SendMessage(opts *bind.TransactOpts, dstChainId *big.Int, gasLimit *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "sendMessage", dstChainId, gasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ca709ee.
//
// Solidity: function sendMessage(uint256 dstChainId, uint256 gasLimit, bytes message) payable returns()
func (_InterchainAppExample *InterchainAppExampleSession) SendMessage(dstChainId *big.Int, gasLimit *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.SendMessage(&_InterchainAppExample.TransactOpts, dstChainId, gasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ca709ee.
//
// Solidity: function sendMessage(uint256 dstChainId, uint256 gasLimit, bytes message) payable returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) SendMessage(dstChainId *big.Int, gasLimit *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.SendMessage(&_InterchainAppExample.TransactOpts, dstChainId, gasLimit, message)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) SetAppConfigV1(opts *bind.TransactOpts, appConfig AppConfigV1) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "setAppConfigV1", appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_InterchainAppExample *InterchainAppExampleSession) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.SetAppConfigV1(&_InterchainAppExample.TransactOpts, appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.SetAppConfigV1(&_InterchainAppExample.TransactOpts, appConfig)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) SetExecutionService(opts *bind.TransactOpts, executionService common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "setExecutionService", executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_InterchainAppExample *InterchainAppExampleSession) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.SetExecutionService(&_InterchainAppExample.TransactOpts, executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.SetExecutionService(&_InterchainAppExample.TransactOpts, executionService)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address interchain_) returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) SetInterchainClient(opts *bind.TransactOpts, interchain_ common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "setInterchainClient", interchain_)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address interchain_) returns()
func (_InterchainAppExample *InterchainAppExampleSession) SetInterchainClient(interchain_ common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.SetInterchainClient(&_InterchainAppExample.TransactOpts, interchain_)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address interchain_) returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) SetInterchainClient(interchain_ common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.SetInterchainClient(&_InterchainAppExample.TransactOpts, interchain_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainAppExample *InterchainAppExampleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainAppExample *InterchainAppExampleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.TransferOwnership(&_InterchainAppExample.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainAppExample *InterchainAppExampleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterchainAppExample.Contract.TransferOwnership(&_InterchainAppExample.TransactOpts, newOwner)
}

// InterchainAppExampleAppConfigV1SetIterator is returned from FilterAppConfigV1Set and is used to iterate over the raw logs and unpacked data for AppConfigV1Set events raised by the InterchainAppExample contract.
type InterchainAppExampleAppConfigV1SetIterator struct {
	Event *InterchainAppExampleAppConfigV1Set // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleAppConfigV1SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleAppConfigV1Set)
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
		it.Event = new(InterchainAppExampleAppConfigV1Set)
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
func (it *InterchainAppExampleAppConfigV1SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleAppConfigV1SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleAppConfigV1Set represents a AppConfigV1Set event raised by the InterchainAppExample contract.
type InterchainAppExampleAppConfigV1Set struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAppConfigV1Set is a free log retrieval operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterAppConfigV1Set(opts *bind.FilterOpts) (*InterchainAppExampleAppConfigV1SetIterator, error) {

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleAppConfigV1SetIterator{contract: _InterchainAppExample.contract, event: "AppConfigV1Set", logs: logs, sub: sub}, nil
}

// WatchAppConfigV1Set is a free log subscription operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchAppConfigV1Set(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleAppConfigV1Set) (event.Subscription, error) {

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleAppConfigV1Set)
				if err := _InterchainAppExample.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
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

// ParseAppConfigV1Set is a log parse operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseAppConfigV1Set(log types.Log) (*InterchainAppExampleAppConfigV1Set, error) {
	event := new(InterchainAppExampleAppConfigV1Set)
	if err := _InterchainAppExample.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppExampleAppLinkedIterator is returned from FilterAppLinked and is used to iterate over the raw logs and unpacked data for AppLinked events raised by the InterchainAppExample contract.
type InterchainAppExampleAppLinkedIterator struct {
	Event *InterchainAppExampleAppLinked // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleAppLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleAppLinked)
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
		it.Event = new(InterchainAppExampleAppLinked)
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
func (it *InterchainAppExampleAppLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleAppLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleAppLinked represents a AppLinked event raised by the InterchainAppExample contract.
type InterchainAppExampleAppLinked struct {
	ChainId   *big.Int
	RemoteApp [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppLinked is a free log retrieval operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterAppLinked(opts *bind.FilterOpts, chainId []*big.Int, remoteApp [][32]byte) (*InterchainAppExampleAppLinkedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleAppLinkedIterator{contract: _InterchainAppExample.contract, event: "AppLinked", logs: logs, sub: sub}, nil
}

// WatchAppLinked is a free log subscription operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchAppLinked(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleAppLinked, chainId []*big.Int, remoteApp [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleAppLinked)
				if err := _InterchainAppExample.contract.UnpackLog(event, "AppLinked", log); err != nil {
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

// ParseAppLinked is a log parse operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseAppLinked(log types.Log) (*InterchainAppExampleAppLinked, error) {
	event := new(InterchainAppExampleAppLinked)
	if err := _InterchainAppExample.contract.UnpackLog(event, "AppLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppExampleExecutionServiceSetIterator is returned from FilterExecutionServiceSet and is used to iterate over the raw logs and unpacked data for ExecutionServiceSet events raised by the InterchainAppExample contract.
type InterchainAppExampleExecutionServiceSetIterator struct {
	Event *InterchainAppExampleExecutionServiceSet // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleExecutionServiceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleExecutionServiceSet)
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
		it.Event = new(InterchainAppExampleExecutionServiceSet)
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
func (it *InterchainAppExampleExecutionServiceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleExecutionServiceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleExecutionServiceSet represents a ExecutionServiceSet event raised by the InterchainAppExample contract.
type InterchainAppExampleExecutionServiceSet struct {
	ExecutionService common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutionServiceSet is a free log retrieval operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterExecutionServiceSet(opts *bind.FilterOpts) (*InterchainAppExampleExecutionServiceSetIterator, error) {

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleExecutionServiceSetIterator{contract: _InterchainAppExample.contract, event: "ExecutionServiceSet", logs: logs, sub: sub}, nil
}

// WatchExecutionServiceSet is a free log subscription operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchExecutionServiceSet(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleExecutionServiceSet) (event.Subscription, error) {

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleExecutionServiceSet)
				if err := _InterchainAppExample.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
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

// ParseExecutionServiceSet is a log parse operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseExecutionServiceSet(log types.Log) (*InterchainAppExampleExecutionServiceSet, error) {
	event := new(InterchainAppExampleExecutionServiceSet)
	if err := _InterchainAppExample.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppExampleInterchainClientSetIterator is returned from FilterInterchainClientSet and is used to iterate over the raw logs and unpacked data for InterchainClientSet events raised by the InterchainAppExample contract.
type InterchainAppExampleInterchainClientSetIterator struct {
	Event *InterchainAppExampleInterchainClientSet // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleInterchainClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleInterchainClientSet)
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
		it.Event = new(InterchainAppExampleInterchainClientSet)
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
func (it *InterchainAppExampleInterchainClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleInterchainClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleInterchainClientSet represents a InterchainClientSet event raised by the InterchainAppExample contract.
type InterchainAppExampleInterchainClientSet struct {
	InterchainClient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientSet is a free log retrieval operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterInterchainClientSet(opts *bind.FilterOpts) (*InterchainAppExampleInterchainClientSetIterator, error) {

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "InterchainClientSet")
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleInterchainClientSetIterator{contract: _InterchainAppExample.contract, event: "InterchainClientSet", logs: logs, sub: sub}, nil
}

// WatchInterchainClientSet is a free log subscription operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchInterchainClientSet(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleInterchainClientSet) (event.Subscription, error) {

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "InterchainClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleInterchainClientSet)
				if err := _InterchainAppExample.contract.UnpackLog(event, "InterchainClientSet", log); err != nil {
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

// ParseInterchainClientSet is a log parse operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseInterchainClientSet(log types.Log) (*InterchainAppExampleInterchainClientSet, error) {
	event := new(InterchainAppExampleInterchainClientSet)
	if err := _InterchainAppExample.contract.UnpackLog(event, "InterchainClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppExampleMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the InterchainAppExample contract.
type InterchainAppExampleMessageReceivedIterator struct {
	Event *InterchainAppExampleMessageReceived // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleMessageReceived)
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
		it.Event = new(InterchainAppExampleMessageReceived)
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
func (it *InterchainAppExampleMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleMessageReceived represents a MessageReceived event raised by the InterchainAppExample contract.
type InterchainAppExampleMessageReceived struct {
	SrcChainId *big.Int
	Sender     [32]byte
	DbNonce    *big.Int
	EntryIndex uint64
	Message    []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0xc5b07cf0d424748241636d3a67366b44ac88118c42bef024abcb7fd138df79a6.
//
// Solidity: event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*InterchainAppExampleMessageReceivedIterator, error) {

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleMessageReceivedIterator{contract: _InterchainAppExample.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0xc5b07cf0d424748241636d3a67366b44ac88118c42bef024abcb7fd138df79a6.
//
// Solidity: event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleMessageReceived) (event.Subscription, error) {

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleMessageReceived)
				if err := _InterchainAppExample.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0xc5b07cf0d424748241636d3a67366b44ac88118c42bef024abcb7fd138df79a6.
//
// Solidity: event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message)
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseMessageReceived(log types.Log) (*InterchainAppExampleMessageReceived, error) {
	event := new(InterchainAppExampleMessageReceived)
	if err := _InterchainAppExample.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppExampleMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the InterchainAppExample contract.
type InterchainAppExampleMessageSentIterator struct {
	Event *InterchainAppExampleMessageSent // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleMessageSent)
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
		it.Event = new(InterchainAppExampleMessageSent)
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
func (it *InterchainAppExampleMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleMessageSent represents a MessageSent event raised by the InterchainAppExample contract.
type InterchainAppExampleMessageSent struct {
	DstChainId    *big.Int
	DbNonce       *big.Int
	EntryIndex    uint64
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x2ef16db2691a32543ce5591798c4992f4cfbbcd446874f1437d99da53d600e7c.
//
// Solidity: event MessageSent(uint256 dstChainId, uint256 dbNonce, uint64 entryIndex, bytes32 transactionId)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterMessageSent(opts *bind.FilterOpts) (*InterchainAppExampleMessageSentIterator, error) {

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleMessageSentIterator{contract: _InterchainAppExample.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x2ef16db2691a32543ce5591798c4992f4cfbbcd446874f1437d99da53d600e7c.
//
// Solidity: event MessageSent(uint256 dstChainId, uint256 dbNonce, uint64 entryIndex, bytes32 transactionId)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleMessageSent) (event.Subscription, error) {

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleMessageSent)
				if err := _InterchainAppExample.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x2ef16db2691a32543ce5591798c4992f4cfbbcd446874f1437d99da53d600e7c.
//
// Solidity: event MessageSent(uint256 dstChainId, uint256 dbNonce, uint64 entryIndex, bytes32 transactionId)
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseMessageSent(log types.Log) (*InterchainAppExampleMessageSent, error) {
	event := new(InterchainAppExampleMessageSent)
	if err := _InterchainAppExample.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppExampleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InterchainAppExample contract.
type InterchainAppExampleOwnershipTransferredIterator struct {
	Event *InterchainAppExampleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleOwnershipTransferred)
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
		it.Event = new(InterchainAppExampleOwnershipTransferred)
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
func (it *InterchainAppExampleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleOwnershipTransferred represents a OwnershipTransferred event raised by the InterchainAppExample contract.
type InterchainAppExampleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InterchainAppExampleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleOwnershipTransferredIterator{contract: _InterchainAppExample.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleOwnershipTransferred)
				if err := _InterchainAppExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseOwnershipTransferred(log types.Log) (*InterchainAppExampleOwnershipTransferred, error) {
	event := new(InterchainAppExampleOwnershipTransferred)
	if err := _InterchainAppExample.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppExampleTrustedModuleAddedIterator is returned from FilterTrustedModuleAdded and is used to iterate over the raw logs and unpacked data for TrustedModuleAdded events raised by the InterchainAppExample contract.
type InterchainAppExampleTrustedModuleAddedIterator struct {
	Event *InterchainAppExampleTrustedModuleAdded // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleTrustedModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleTrustedModuleAdded)
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
		it.Event = new(InterchainAppExampleTrustedModuleAdded)
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
func (it *InterchainAppExampleTrustedModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleTrustedModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleTrustedModuleAdded represents a TrustedModuleAdded event raised by the InterchainAppExample contract.
type InterchainAppExampleTrustedModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleAdded is a free log retrieval operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterTrustedModuleAdded(opts *bind.FilterOpts) (*InterchainAppExampleTrustedModuleAddedIterator, error) {

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleTrustedModuleAddedIterator{contract: _InterchainAppExample.contract, event: "TrustedModuleAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleAdded is a free log subscription operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchTrustedModuleAdded(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleTrustedModuleAdded) (event.Subscription, error) {

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleTrustedModuleAdded)
				if err := _InterchainAppExample.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
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

// ParseTrustedModuleAdded is a log parse operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseTrustedModuleAdded(log types.Log) (*InterchainAppExampleTrustedModuleAdded, error) {
	event := new(InterchainAppExampleTrustedModuleAdded)
	if err := _InterchainAppExample.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppExampleTrustedModuleRemovedIterator is returned from FilterTrustedModuleRemoved and is used to iterate over the raw logs and unpacked data for TrustedModuleRemoved events raised by the InterchainAppExample contract.
type InterchainAppExampleTrustedModuleRemovedIterator struct {
	Event *InterchainAppExampleTrustedModuleRemoved // Event containing the contract specifics and raw log

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
func (it *InterchainAppExampleTrustedModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppExampleTrustedModuleRemoved)
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
		it.Event = new(InterchainAppExampleTrustedModuleRemoved)
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
func (it *InterchainAppExampleTrustedModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppExampleTrustedModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppExampleTrustedModuleRemoved represents a TrustedModuleRemoved event raised by the InterchainAppExample contract.
type InterchainAppExampleTrustedModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleRemoved is a free log retrieval operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppExample *InterchainAppExampleFilterer) FilterTrustedModuleRemoved(opts *bind.FilterOpts) (*InterchainAppExampleTrustedModuleRemovedIterator, error) {

	logs, sub, err := _InterchainAppExample.contract.FilterLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &InterchainAppExampleTrustedModuleRemovedIterator{contract: _InterchainAppExample.contract, event: "TrustedModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleRemoved is a free log subscription operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppExample *InterchainAppExampleFilterer) WatchTrustedModuleRemoved(opts *bind.WatchOpts, sink chan<- *InterchainAppExampleTrustedModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _InterchainAppExample.contract.WatchLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppExampleTrustedModuleRemoved)
				if err := _InterchainAppExample.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
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

// ParseTrustedModuleRemoved is a log parse operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppExample *InterchainAppExampleFilterer) ParseTrustedModuleRemoved(log types.Log) (*InterchainAppExampleTrustedModuleRemoved, error) {
	event := new(InterchainAppExampleTrustedModuleRemoved)
	if err := _InterchainAppExample.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1MetaData contains all meta data concerning the InterchainAppV1 contract.
var InterchainAppV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainApp__CallerNotInterchainClient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__InterchainClientNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleNotAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__ReceiverNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__SenderNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"AppConfigV1Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"AppLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"ExecutionServiceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"interchainClient\",\"type\":\"address\"}],\"name\":\"InterchainClientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppConfigV1\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedApp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"isAllowedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"7717a647": "getAppConfigV1()",
		"a20ce510": "getAppVersion()",
		"c313c807": "getExecutionService()",
		"b9b74b18": "getLinkedApp(uint256)",
		"287bc057": "getReceivingConfig()",
		"a45e107a": "getReceivingModules()",
		"ea13398f": "getSendingModules()",
		"70838975": "interchain()",
		"dc2b9075": "isAllowedSender(uint256,bytes32)",
	},
}

// InterchainAppV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainAppV1MetaData.ABI instead.
var InterchainAppV1ABI = InterchainAppV1MetaData.ABI

// Deprecated: Use InterchainAppV1MetaData.Sigs instead.
// InterchainAppV1FuncSigs maps the 4-byte function signature to its string representation.
var InterchainAppV1FuncSigs = InterchainAppV1MetaData.Sigs

// InterchainAppV1 is an auto generated Go binding around an Ethereum contract.
type InterchainAppV1 struct {
	InterchainAppV1Caller     // Read-only binding to the contract
	InterchainAppV1Transactor // Write-only binding to the contract
	InterchainAppV1Filterer   // Log filterer for contract events
}

// InterchainAppV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainAppV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainAppV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainAppV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainAppV1Session struct {
	Contract     *InterchainAppV1  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterchainAppV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainAppV1CallerSession struct {
	Contract *InterchainAppV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// InterchainAppV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainAppV1TransactorSession struct {
	Contract     *InterchainAppV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InterchainAppV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainAppV1Raw struct {
	Contract *InterchainAppV1 // Generic contract binding to access the raw methods on
}

// InterchainAppV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainAppV1CallerRaw struct {
	Contract *InterchainAppV1Caller // Generic read-only contract binding to access the raw methods on
}

// InterchainAppV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainAppV1TransactorRaw struct {
	Contract *InterchainAppV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainAppV1 creates a new instance of InterchainAppV1, bound to a specific deployed contract.
func NewInterchainAppV1(address common.Address, backend bind.ContractBackend) (*InterchainAppV1, error) {
	contract, err := bindInterchainAppV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1{InterchainAppV1Caller: InterchainAppV1Caller{contract: contract}, InterchainAppV1Transactor: InterchainAppV1Transactor{contract: contract}, InterchainAppV1Filterer: InterchainAppV1Filterer{contract: contract}}, nil
}

// NewInterchainAppV1Caller creates a new read-only instance of InterchainAppV1, bound to a specific deployed contract.
func NewInterchainAppV1Caller(address common.Address, caller bind.ContractCaller) (*InterchainAppV1Caller, error) {
	contract, err := bindInterchainAppV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1Caller{contract: contract}, nil
}

// NewInterchainAppV1Transactor creates a new write-only instance of InterchainAppV1, bound to a specific deployed contract.
func NewInterchainAppV1Transactor(address common.Address, transactor bind.ContractTransactor) (*InterchainAppV1Transactor, error) {
	contract, err := bindInterchainAppV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1Transactor{contract: contract}, nil
}

// NewInterchainAppV1Filterer creates a new log filterer instance of InterchainAppV1, bound to a specific deployed contract.
func NewInterchainAppV1Filterer(address common.Address, filterer bind.ContractFilterer) (*InterchainAppV1Filterer, error) {
	contract, err := bindInterchainAppV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1Filterer{contract: contract}, nil
}

// bindInterchainAppV1 binds a generic wrapper to an already deployed contract.
func bindInterchainAppV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainAppV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppV1 *InterchainAppV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppV1.Contract.InterchainAppV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppV1 *InterchainAppV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppV1.Contract.InterchainAppV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppV1 *InterchainAppV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppV1.Contract.InterchainAppV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppV1 *InterchainAppV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppV1 *InterchainAppV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppV1 *InterchainAppV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppV1.Contract.contract.Transact(opts, method, params...)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_InterchainAppV1 *InterchainAppV1Caller) GetAppConfigV1(opts *bind.CallOpts) (AppConfigV1, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "getAppConfigV1")

	if err != nil {
		return *new(AppConfigV1), err
	}

	out0 := *abi.ConvertType(out[0], new(AppConfigV1)).(*AppConfigV1)

	return out0, err

}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_InterchainAppV1 *InterchainAppV1Session) GetAppConfigV1() (AppConfigV1, error) {
	return _InterchainAppV1.Contract.GetAppConfigV1(&_InterchainAppV1.CallOpts)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_InterchainAppV1 *InterchainAppV1CallerSession) GetAppConfigV1() (AppConfigV1, error) {
	return _InterchainAppV1.Contract.GetAppConfigV1(&_InterchainAppV1.CallOpts)
}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_InterchainAppV1 *InterchainAppV1Caller) GetAppVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "getAppVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_InterchainAppV1 *InterchainAppV1Session) GetAppVersion() (*big.Int, error) {
	return _InterchainAppV1.Contract.GetAppVersion(&_InterchainAppV1.CallOpts)
}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_InterchainAppV1 *InterchainAppV1CallerSession) GetAppVersion() (*big.Int, error) {
	return _InterchainAppV1.Contract.GetAppVersion(&_InterchainAppV1.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_InterchainAppV1 *InterchainAppV1Caller) GetExecutionService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "getExecutionService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_InterchainAppV1 *InterchainAppV1Session) GetExecutionService() (common.Address, error) {
	return _InterchainAppV1.Contract.GetExecutionService(&_InterchainAppV1.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_InterchainAppV1 *InterchainAppV1CallerSession) GetExecutionService() (common.Address, error) {
	return _InterchainAppV1.Contract.GetExecutionService(&_InterchainAppV1.CallOpts)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_InterchainAppV1 *InterchainAppV1Caller) GetLinkedApp(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "getLinkedApp", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_InterchainAppV1 *InterchainAppV1Session) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _InterchainAppV1.Contract.GetLinkedApp(&_InterchainAppV1.CallOpts, chainId)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_InterchainAppV1 *InterchainAppV1CallerSession) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _InterchainAppV1.Contract.GetLinkedApp(&_InterchainAppV1.CallOpts, chainId)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_InterchainAppV1 *InterchainAppV1Caller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "getReceivingConfig")

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
func (_InterchainAppV1 *InterchainAppV1Session) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _InterchainAppV1.Contract.GetReceivingConfig(&_InterchainAppV1.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_InterchainAppV1 *InterchainAppV1CallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _InterchainAppV1.Contract.GetReceivingConfig(&_InterchainAppV1.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppV1 *InterchainAppV1Caller) GetReceivingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "getReceivingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppV1 *InterchainAppV1Session) GetReceivingModules() ([]common.Address, error) {
	return _InterchainAppV1.Contract.GetReceivingModules(&_InterchainAppV1.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainAppV1 *InterchainAppV1CallerSession) GetReceivingModules() ([]common.Address, error) {
	return _InterchainAppV1.Contract.GetReceivingModules(&_InterchainAppV1.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppV1 *InterchainAppV1Caller) GetSendingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "getSendingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppV1 *InterchainAppV1Session) GetSendingModules() ([]common.Address, error) {
	return _InterchainAppV1.Contract.GetSendingModules(&_InterchainAppV1.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainAppV1 *InterchainAppV1CallerSession) GetSendingModules() ([]common.Address, error) {
	return _InterchainAppV1.Contract.GetSendingModules(&_InterchainAppV1.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainAppV1 *InterchainAppV1Caller) Interchain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "interchain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainAppV1 *InterchainAppV1Session) Interchain() (common.Address, error) {
	return _InterchainAppV1.Contract.Interchain(&_InterchainAppV1.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainAppV1 *InterchainAppV1CallerSession) Interchain() (common.Address, error) {
	return _InterchainAppV1.Contract.Interchain(&_InterchainAppV1.CallOpts)
}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_InterchainAppV1 *InterchainAppV1Caller) IsAllowedSender(opts *bind.CallOpts, srcChainId *big.Int, sender [32]byte) (bool, error) {
	var out []interface{}
	err := _InterchainAppV1.contract.Call(opts, &out, "isAllowedSender", srcChainId, sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_InterchainAppV1 *InterchainAppV1Session) IsAllowedSender(srcChainId *big.Int, sender [32]byte) (bool, error) {
	return _InterchainAppV1.Contract.IsAllowedSender(&_InterchainAppV1.CallOpts, srcChainId, sender)
}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_InterchainAppV1 *InterchainAppV1CallerSession) IsAllowedSender(srcChainId *big.Int, sender [32]byte) (bool, error) {
	return _InterchainAppV1.Contract.IsAllowedSender(&_InterchainAppV1.CallOpts, srcChainId, sender)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_InterchainAppV1 *InterchainAppV1Transactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppV1.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_InterchainAppV1 *InterchainAppV1Session) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppV1.Contract.AppReceive(&_InterchainAppV1.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_InterchainAppV1 *InterchainAppV1TransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _InterchainAppV1.Contract.AppReceive(&_InterchainAppV1.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// InterchainAppV1AppConfigV1SetIterator is returned from FilterAppConfigV1Set and is used to iterate over the raw logs and unpacked data for AppConfigV1Set events raised by the InterchainAppV1 contract.
type InterchainAppV1AppConfigV1SetIterator struct {
	Event *InterchainAppV1AppConfigV1Set // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1AppConfigV1SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1AppConfigV1Set)
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
		it.Event = new(InterchainAppV1AppConfigV1Set)
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
func (it *InterchainAppV1AppConfigV1SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1AppConfigV1SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1AppConfigV1Set represents a AppConfigV1Set event raised by the InterchainAppV1 contract.
type InterchainAppV1AppConfigV1Set struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAppConfigV1Set is a free log retrieval operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1 *InterchainAppV1Filterer) FilterAppConfigV1Set(opts *bind.FilterOpts) (*InterchainAppV1AppConfigV1SetIterator, error) {

	logs, sub, err := _InterchainAppV1.contract.FilterLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1AppConfigV1SetIterator{contract: _InterchainAppV1.contract, event: "AppConfigV1Set", logs: logs, sub: sub}, nil
}

// WatchAppConfigV1Set is a free log subscription operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1 *InterchainAppV1Filterer) WatchAppConfigV1Set(opts *bind.WatchOpts, sink chan<- *InterchainAppV1AppConfigV1Set) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1.contract.WatchLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1AppConfigV1Set)
				if err := _InterchainAppV1.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
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

// ParseAppConfigV1Set is a log parse operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1 *InterchainAppV1Filterer) ParseAppConfigV1Set(log types.Log) (*InterchainAppV1AppConfigV1Set, error) {
	event := new(InterchainAppV1AppConfigV1Set)
	if err := _InterchainAppV1.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1AppLinkedIterator is returned from FilterAppLinked and is used to iterate over the raw logs and unpacked data for AppLinked events raised by the InterchainAppV1 contract.
type InterchainAppV1AppLinkedIterator struct {
	Event *InterchainAppV1AppLinked // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1AppLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1AppLinked)
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
		it.Event = new(InterchainAppV1AppLinked)
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
func (it *InterchainAppV1AppLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1AppLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1AppLinked represents a AppLinked event raised by the InterchainAppV1 contract.
type InterchainAppV1AppLinked struct {
	ChainId   *big.Int
	RemoteApp [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppLinked is a free log retrieval operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1 *InterchainAppV1Filterer) FilterAppLinked(opts *bind.FilterOpts, chainId []*big.Int, remoteApp [][32]byte) (*InterchainAppV1AppLinkedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _InterchainAppV1.contract.FilterLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1AppLinkedIterator{contract: _InterchainAppV1.contract, event: "AppLinked", logs: logs, sub: sub}, nil
}

// WatchAppLinked is a free log subscription operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1 *InterchainAppV1Filterer) WatchAppLinked(opts *bind.WatchOpts, sink chan<- *InterchainAppV1AppLinked, chainId []*big.Int, remoteApp [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _InterchainAppV1.contract.WatchLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1AppLinked)
				if err := _InterchainAppV1.contract.UnpackLog(event, "AppLinked", log); err != nil {
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

// ParseAppLinked is a log parse operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1 *InterchainAppV1Filterer) ParseAppLinked(log types.Log) (*InterchainAppV1AppLinked, error) {
	event := new(InterchainAppV1AppLinked)
	if err := _InterchainAppV1.contract.UnpackLog(event, "AppLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1ExecutionServiceSetIterator is returned from FilterExecutionServiceSet and is used to iterate over the raw logs and unpacked data for ExecutionServiceSet events raised by the InterchainAppV1 contract.
type InterchainAppV1ExecutionServiceSetIterator struct {
	Event *InterchainAppV1ExecutionServiceSet // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1ExecutionServiceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1ExecutionServiceSet)
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
		it.Event = new(InterchainAppV1ExecutionServiceSet)
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
func (it *InterchainAppV1ExecutionServiceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1ExecutionServiceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1ExecutionServiceSet represents a ExecutionServiceSet event raised by the InterchainAppV1 contract.
type InterchainAppV1ExecutionServiceSet struct {
	ExecutionService common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutionServiceSet is a free log retrieval operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1 *InterchainAppV1Filterer) FilterExecutionServiceSet(opts *bind.FilterOpts) (*InterchainAppV1ExecutionServiceSetIterator, error) {

	logs, sub, err := _InterchainAppV1.contract.FilterLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1ExecutionServiceSetIterator{contract: _InterchainAppV1.contract, event: "ExecutionServiceSet", logs: logs, sub: sub}, nil
}

// WatchExecutionServiceSet is a free log subscription operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1 *InterchainAppV1Filterer) WatchExecutionServiceSet(opts *bind.WatchOpts, sink chan<- *InterchainAppV1ExecutionServiceSet) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1.contract.WatchLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1ExecutionServiceSet)
				if err := _InterchainAppV1.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
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

// ParseExecutionServiceSet is a log parse operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1 *InterchainAppV1Filterer) ParseExecutionServiceSet(log types.Log) (*InterchainAppV1ExecutionServiceSet, error) {
	event := new(InterchainAppV1ExecutionServiceSet)
	if err := _InterchainAppV1.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1InterchainClientSetIterator is returned from FilterInterchainClientSet and is used to iterate over the raw logs and unpacked data for InterchainClientSet events raised by the InterchainAppV1 contract.
type InterchainAppV1InterchainClientSetIterator struct {
	Event *InterchainAppV1InterchainClientSet // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1InterchainClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1InterchainClientSet)
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
		it.Event = new(InterchainAppV1InterchainClientSet)
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
func (it *InterchainAppV1InterchainClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1InterchainClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1InterchainClientSet represents a InterchainClientSet event raised by the InterchainAppV1 contract.
type InterchainAppV1InterchainClientSet struct {
	InterchainClient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientSet is a free log retrieval operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppV1 *InterchainAppV1Filterer) FilterInterchainClientSet(opts *bind.FilterOpts) (*InterchainAppV1InterchainClientSetIterator, error) {

	logs, sub, err := _InterchainAppV1.contract.FilterLogs(opts, "InterchainClientSet")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1InterchainClientSetIterator{contract: _InterchainAppV1.contract, event: "InterchainClientSet", logs: logs, sub: sub}, nil
}

// WatchInterchainClientSet is a free log subscription operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppV1 *InterchainAppV1Filterer) WatchInterchainClientSet(opts *bind.WatchOpts, sink chan<- *InterchainAppV1InterchainClientSet) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1.contract.WatchLogs(opts, "InterchainClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1InterchainClientSet)
				if err := _InterchainAppV1.contract.UnpackLog(event, "InterchainClientSet", log); err != nil {
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

// ParseInterchainClientSet is a log parse operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppV1 *InterchainAppV1Filterer) ParseInterchainClientSet(log types.Log) (*InterchainAppV1InterchainClientSet, error) {
	event := new(InterchainAppV1InterchainClientSet)
	if err := _InterchainAppV1.contract.UnpackLog(event, "InterchainClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1TrustedModuleAddedIterator is returned from FilterTrustedModuleAdded and is used to iterate over the raw logs and unpacked data for TrustedModuleAdded events raised by the InterchainAppV1 contract.
type InterchainAppV1TrustedModuleAddedIterator struct {
	Event *InterchainAppV1TrustedModuleAdded // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1TrustedModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1TrustedModuleAdded)
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
		it.Event = new(InterchainAppV1TrustedModuleAdded)
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
func (it *InterchainAppV1TrustedModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1TrustedModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1TrustedModuleAdded represents a TrustedModuleAdded event raised by the InterchainAppV1 contract.
type InterchainAppV1TrustedModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleAdded is a free log retrieval operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1 *InterchainAppV1Filterer) FilterTrustedModuleAdded(opts *bind.FilterOpts) (*InterchainAppV1TrustedModuleAddedIterator, error) {

	logs, sub, err := _InterchainAppV1.contract.FilterLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1TrustedModuleAddedIterator{contract: _InterchainAppV1.contract, event: "TrustedModuleAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleAdded is a free log subscription operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1 *InterchainAppV1Filterer) WatchTrustedModuleAdded(opts *bind.WatchOpts, sink chan<- *InterchainAppV1TrustedModuleAdded) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1.contract.WatchLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1TrustedModuleAdded)
				if err := _InterchainAppV1.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
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

// ParseTrustedModuleAdded is a log parse operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1 *InterchainAppV1Filterer) ParseTrustedModuleAdded(log types.Log) (*InterchainAppV1TrustedModuleAdded, error) {
	event := new(InterchainAppV1TrustedModuleAdded)
	if err := _InterchainAppV1.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1TrustedModuleRemovedIterator is returned from FilterTrustedModuleRemoved and is used to iterate over the raw logs and unpacked data for TrustedModuleRemoved events raised by the InterchainAppV1 contract.
type InterchainAppV1TrustedModuleRemovedIterator struct {
	Event *InterchainAppV1TrustedModuleRemoved // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1TrustedModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1TrustedModuleRemoved)
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
		it.Event = new(InterchainAppV1TrustedModuleRemoved)
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
func (it *InterchainAppV1TrustedModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1TrustedModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1TrustedModuleRemoved represents a TrustedModuleRemoved event raised by the InterchainAppV1 contract.
type InterchainAppV1TrustedModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleRemoved is a free log retrieval operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1 *InterchainAppV1Filterer) FilterTrustedModuleRemoved(opts *bind.FilterOpts) (*InterchainAppV1TrustedModuleRemovedIterator, error) {

	logs, sub, err := _InterchainAppV1.contract.FilterLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1TrustedModuleRemovedIterator{contract: _InterchainAppV1.contract, event: "TrustedModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleRemoved is a free log subscription operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1 *InterchainAppV1Filterer) WatchTrustedModuleRemoved(opts *bind.WatchOpts, sink chan<- *InterchainAppV1TrustedModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1.contract.WatchLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1TrustedModuleRemoved)
				if err := _InterchainAppV1.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
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

// ParseTrustedModuleRemoved is a log parse operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1 *InterchainAppV1Filterer) ParseTrustedModuleRemoved(log types.Log) (*InterchainAppV1TrustedModuleRemoved, error) {
	event := new(InterchainAppV1TrustedModuleRemoved)
	if err := _InterchainAppV1.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsMetaData contains all meta data concerning the InterchainAppV1Events contract.
var InterchainAppV1EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"AppConfigV1Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"AppLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"ExecutionServiceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"interchainClient\",\"type\":\"address\"}],\"name\":\"InterchainClientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleRemoved\",\"type\":\"event\"}]",
}

// InterchainAppV1EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainAppV1EventsMetaData.ABI instead.
var InterchainAppV1EventsABI = InterchainAppV1EventsMetaData.ABI

// InterchainAppV1Events is an auto generated Go binding around an Ethereum contract.
type InterchainAppV1Events struct {
	InterchainAppV1EventsCaller     // Read-only binding to the contract
	InterchainAppV1EventsTransactor // Write-only binding to the contract
	InterchainAppV1EventsFilterer   // Log filterer for contract events
}

// InterchainAppV1EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainAppV1EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainAppV1EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainAppV1EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainAppV1EventsSession struct {
	Contract     *InterchainAppV1Events // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InterchainAppV1EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainAppV1EventsCallerSession struct {
	Contract *InterchainAppV1EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// InterchainAppV1EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainAppV1EventsTransactorSession struct {
	Contract     *InterchainAppV1EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// InterchainAppV1EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainAppV1EventsRaw struct {
	Contract *InterchainAppV1Events // Generic contract binding to access the raw methods on
}

// InterchainAppV1EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainAppV1EventsCallerRaw struct {
	Contract *InterchainAppV1EventsCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainAppV1EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainAppV1EventsTransactorRaw struct {
	Contract *InterchainAppV1EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainAppV1Events creates a new instance of InterchainAppV1Events, bound to a specific deployed contract.
func NewInterchainAppV1Events(address common.Address, backend bind.ContractBackend) (*InterchainAppV1Events, error) {
	contract, err := bindInterchainAppV1Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1Events{InterchainAppV1EventsCaller: InterchainAppV1EventsCaller{contract: contract}, InterchainAppV1EventsTransactor: InterchainAppV1EventsTransactor{contract: contract}, InterchainAppV1EventsFilterer: InterchainAppV1EventsFilterer{contract: contract}}, nil
}

// NewInterchainAppV1EventsCaller creates a new read-only instance of InterchainAppV1Events, bound to a specific deployed contract.
func NewInterchainAppV1EventsCaller(address common.Address, caller bind.ContractCaller) (*InterchainAppV1EventsCaller, error) {
	contract, err := bindInterchainAppV1Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsCaller{contract: contract}, nil
}

// NewInterchainAppV1EventsTransactor creates a new write-only instance of InterchainAppV1Events, bound to a specific deployed contract.
func NewInterchainAppV1EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainAppV1EventsTransactor, error) {
	contract, err := bindInterchainAppV1Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsTransactor{contract: contract}, nil
}

// NewInterchainAppV1EventsFilterer creates a new log filterer instance of InterchainAppV1Events, bound to a specific deployed contract.
func NewInterchainAppV1EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainAppV1EventsFilterer, error) {
	contract, err := bindInterchainAppV1Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsFilterer{contract: contract}, nil
}

// bindInterchainAppV1Events binds a generic wrapper to an already deployed contract.
func bindInterchainAppV1Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainAppV1EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppV1Events *InterchainAppV1EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppV1Events.Contract.InterchainAppV1EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppV1Events *InterchainAppV1EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppV1Events.Contract.InterchainAppV1EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppV1Events *InterchainAppV1EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppV1Events.Contract.InterchainAppV1EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppV1Events *InterchainAppV1EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppV1Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppV1Events *InterchainAppV1EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppV1Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppV1Events *InterchainAppV1EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppV1Events.Contract.contract.Transact(opts, method, params...)
}

// InterchainAppV1EventsAppConfigV1SetIterator is returned from FilterAppConfigV1Set and is used to iterate over the raw logs and unpacked data for AppConfigV1Set events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsAppConfigV1SetIterator struct {
	Event *InterchainAppV1EventsAppConfigV1Set // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1EventsAppConfigV1SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsAppConfigV1Set)
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
		it.Event = new(InterchainAppV1EventsAppConfigV1Set)
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
func (it *InterchainAppV1EventsAppConfigV1SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsAppConfigV1SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsAppConfigV1Set represents a AppConfigV1Set event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsAppConfigV1Set struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAppConfigV1Set is a free log retrieval operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterAppConfigV1Set(opts *bind.FilterOpts) (*InterchainAppV1EventsAppConfigV1SetIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsAppConfigV1SetIterator{contract: _InterchainAppV1Events.contract, event: "AppConfigV1Set", logs: logs, sub: sub}, nil
}

// WatchAppConfigV1Set is a free log subscription operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchAppConfigV1Set(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsAppConfigV1Set) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsAppConfigV1Set)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
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

// ParseAppConfigV1Set is a log parse operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseAppConfigV1Set(log types.Log) (*InterchainAppV1EventsAppConfigV1Set, error) {
	event := new(InterchainAppV1EventsAppConfigV1Set)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsAppLinkedIterator is returned from FilterAppLinked and is used to iterate over the raw logs and unpacked data for AppLinked events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsAppLinkedIterator struct {
	Event *InterchainAppV1EventsAppLinked // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1EventsAppLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsAppLinked)
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
		it.Event = new(InterchainAppV1EventsAppLinked)
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
func (it *InterchainAppV1EventsAppLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsAppLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsAppLinked represents a AppLinked event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsAppLinked struct {
	ChainId   *big.Int
	RemoteApp [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppLinked is a free log retrieval operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterAppLinked(opts *bind.FilterOpts, chainId []*big.Int, remoteApp [][32]byte) (*InterchainAppV1EventsAppLinkedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsAppLinkedIterator{contract: _InterchainAppV1Events.contract, event: "AppLinked", logs: logs, sub: sub}, nil
}

// WatchAppLinked is a free log subscription operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchAppLinked(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsAppLinked, chainId []*big.Int, remoteApp [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsAppLinked)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "AppLinked", log); err != nil {
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

// ParseAppLinked is a log parse operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseAppLinked(log types.Log) (*InterchainAppV1EventsAppLinked, error) {
	event := new(InterchainAppV1EventsAppLinked)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "AppLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsExecutionServiceSetIterator is returned from FilterExecutionServiceSet and is used to iterate over the raw logs and unpacked data for ExecutionServiceSet events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsExecutionServiceSetIterator struct {
	Event *InterchainAppV1EventsExecutionServiceSet // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1EventsExecutionServiceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsExecutionServiceSet)
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
		it.Event = new(InterchainAppV1EventsExecutionServiceSet)
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
func (it *InterchainAppV1EventsExecutionServiceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsExecutionServiceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsExecutionServiceSet represents a ExecutionServiceSet event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsExecutionServiceSet struct {
	ExecutionService common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutionServiceSet is a free log retrieval operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterExecutionServiceSet(opts *bind.FilterOpts) (*InterchainAppV1EventsExecutionServiceSetIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsExecutionServiceSetIterator{contract: _InterchainAppV1Events.contract, event: "ExecutionServiceSet", logs: logs, sub: sub}, nil
}

// WatchExecutionServiceSet is a free log subscription operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchExecutionServiceSet(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsExecutionServiceSet) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsExecutionServiceSet)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
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

// ParseExecutionServiceSet is a log parse operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseExecutionServiceSet(log types.Log) (*InterchainAppV1EventsExecutionServiceSet, error) {
	event := new(InterchainAppV1EventsExecutionServiceSet)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsInterchainClientSetIterator is returned from FilterInterchainClientSet and is used to iterate over the raw logs and unpacked data for InterchainClientSet events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsInterchainClientSetIterator struct {
	Event *InterchainAppV1EventsInterchainClientSet // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1EventsInterchainClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsInterchainClientSet)
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
		it.Event = new(InterchainAppV1EventsInterchainClientSet)
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
func (it *InterchainAppV1EventsInterchainClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsInterchainClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsInterchainClientSet represents a InterchainClientSet event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsInterchainClientSet struct {
	InterchainClient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientSet is a free log retrieval operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterInterchainClientSet(opts *bind.FilterOpts) (*InterchainAppV1EventsInterchainClientSetIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "InterchainClientSet")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsInterchainClientSetIterator{contract: _InterchainAppV1Events.contract, event: "InterchainClientSet", logs: logs, sub: sub}, nil
}

// WatchInterchainClientSet is a free log subscription operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchInterchainClientSet(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsInterchainClientSet) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "InterchainClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsInterchainClientSet)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "InterchainClientSet", log); err != nil {
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

// ParseInterchainClientSet is a log parse operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseInterchainClientSet(log types.Log) (*InterchainAppV1EventsInterchainClientSet, error) {
	event := new(InterchainAppV1EventsInterchainClientSet)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "InterchainClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsTrustedModuleAddedIterator is returned from FilterTrustedModuleAdded and is used to iterate over the raw logs and unpacked data for TrustedModuleAdded events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsTrustedModuleAddedIterator struct {
	Event *InterchainAppV1EventsTrustedModuleAdded // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1EventsTrustedModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsTrustedModuleAdded)
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
		it.Event = new(InterchainAppV1EventsTrustedModuleAdded)
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
func (it *InterchainAppV1EventsTrustedModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsTrustedModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsTrustedModuleAdded represents a TrustedModuleAdded event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsTrustedModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleAdded is a free log retrieval operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterTrustedModuleAdded(opts *bind.FilterOpts) (*InterchainAppV1EventsTrustedModuleAddedIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsTrustedModuleAddedIterator{contract: _InterchainAppV1Events.contract, event: "TrustedModuleAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleAdded is a free log subscription operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchTrustedModuleAdded(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsTrustedModuleAdded) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsTrustedModuleAdded)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
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

// ParseTrustedModuleAdded is a log parse operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseTrustedModuleAdded(log types.Log) (*InterchainAppV1EventsTrustedModuleAdded, error) {
	event := new(InterchainAppV1EventsTrustedModuleAdded)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsTrustedModuleRemovedIterator is returned from FilterTrustedModuleRemoved and is used to iterate over the raw logs and unpacked data for TrustedModuleRemoved events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsTrustedModuleRemovedIterator struct {
	Event *InterchainAppV1EventsTrustedModuleRemoved // Event containing the contract specifics and raw log

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
func (it *InterchainAppV1EventsTrustedModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsTrustedModuleRemoved)
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
		it.Event = new(InterchainAppV1EventsTrustedModuleRemoved)
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
func (it *InterchainAppV1EventsTrustedModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsTrustedModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsTrustedModuleRemoved represents a TrustedModuleRemoved event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsTrustedModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleRemoved is a free log retrieval operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterTrustedModuleRemoved(opts *bind.FilterOpts) (*InterchainAppV1EventsTrustedModuleRemovedIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsTrustedModuleRemovedIterator{contract: _InterchainAppV1Events.contract, event: "TrustedModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleRemoved is a free log subscription operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchTrustedModuleRemoved(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsTrustedModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsTrustedModuleRemoved)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
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

// ParseTrustedModuleRemoved is a log parse operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseTrustedModuleRemoved(log types.Log) (*InterchainAppV1EventsTrustedModuleRemoved, error) {
	event := new(InterchainAppV1EventsTrustedModuleRemoved)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainTransactionLibMetaData contains all meta data concerning the InterchainTransactionLib contract.
var InterchainTransactionLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220615335ab93fa354ba3e6905e54aa39796f40054f50849e6888a95320e9db9ffd64736f6c63430008140033",
}

// InterchainTransactionLibABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainTransactionLibMetaData.ABI instead.
var InterchainTransactionLibABI = InterchainTransactionLibMetaData.ABI

// InterchainTransactionLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainTransactionLibMetaData.Bin instead.
var InterchainTransactionLibBin = InterchainTransactionLibMetaData.Bin

// DeployInterchainTransactionLib deploys a new Ethereum contract, binding an instance of InterchainTransactionLib to it.
func DeployInterchainTransactionLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainTransactionLib, error) {
	parsed, err := InterchainTransactionLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainTransactionLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainTransactionLib{InterchainTransactionLibCaller: InterchainTransactionLibCaller{contract: contract}, InterchainTransactionLibTransactor: InterchainTransactionLibTransactor{contract: contract}, InterchainTransactionLibFilterer: InterchainTransactionLibFilterer{contract: contract}}, nil
}

// InterchainTransactionLib is an auto generated Go binding around an Ethereum contract.
type InterchainTransactionLib struct {
	InterchainTransactionLibCaller     // Read-only binding to the contract
	InterchainTransactionLibTransactor // Write-only binding to the contract
	InterchainTransactionLibFilterer   // Log filterer for contract events
}

// InterchainTransactionLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainTransactionLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainTransactionLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainTransactionLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainTransactionLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainTransactionLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainTransactionLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainTransactionLibSession struct {
	Contract     *InterchainTransactionLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InterchainTransactionLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainTransactionLibCallerSession struct {
	Contract *InterchainTransactionLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// InterchainTransactionLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainTransactionLibTransactorSession struct {
	Contract     *InterchainTransactionLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// InterchainTransactionLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainTransactionLibRaw struct {
	Contract *InterchainTransactionLib // Generic contract binding to access the raw methods on
}

// InterchainTransactionLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainTransactionLibCallerRaw struct {
	Contract *InterchainTransactionLibCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainTransactionLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainTransactionLibTransactorRaw struct {
	Contract *InterchainTransactionLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainTransactionLib creates a new instance of InterchainTransactionLib, bound to a specific deployed contract.
func NewInterchainTransactionLib(address common.Address, backend bind.ContractBackend) (*InterchainTransactionLib, error) {
	contract, err := bindInterchainTransactionLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionLib{InterchainTransactionLibCaller: InterchainTransactionLibCaller{contract: contract}, InterchainTransactionLibTransactor: InterchainTransactionLibTransactor{contract: contract}, InterchainTransactionLibFilterer: InterchainTransactionLibFilterer{contract: contract}}, nil
}

// NewInterchainTransactionLibCaller creates a new read-only instance of InterchainTransactionLib, bound to a specific deployed contract.
func NewInterchainTransactionLibCaller(address common.Address, caller bind.ContractCaller) (*InterchainTransactionLibCaller, error) {
	contract, err := bindInterchainTransactionLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionLibCaller{contract: contract}, nil
}

// NewInterchainTransactionLibTransactor creates a new write-only instance of InterchainTransactionLib, bound to a specific deployed contract.
func NewInterchainTransactionLibTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainTransactionLibTransactor, error) {
	contract, err := bindInterchainTransactionLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionLibTransactor{contract: contract}, nil
}

// NewInterchainTransactionLibFilterer creates a new log filterer instance of InterchainTransactionLib, bound to a specific deployed contract.
func NewInterchainTransactionLibFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainTransactionLibFilterer, error) {
	contract, err := bindInterchainTransactionLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionLibFilterer{contract: contract}, nil
}

// bindInterchainTransactionLib binds a generic wrapper to an already deployed contract.
func bindInterchainTransactionLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainTransactionLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainTransactionLib *InterchainTransactionLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainTransactionLib.Contract.InterchainTransactionLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainTransactionLib *InterchainTransactionLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainTransactionLib.Contract.InterchainTransactionLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainTransactionLib *InterchainTransactionLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainTransactionLib.Contract.InterchainTransactionLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainTransactionLib *InterchainTransactionLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainTransactionLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainTransactionLib *InterchainTransactionLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainTransactionLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainTransactionLib *InterchainTransactionLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainTransactionLib.Contract.contract.Transact(opts, method, params...)
}

// OptionsLibMetaData contains all meta data concerning the OptionsLib contract.
var OptionsLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209be4b74281be2da67b16c7389df5391b0537bf6017253468393d1b53da752ac664736f6c63430008140033",
}

// OptionsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OptionsLibMetaData.ABI instead.
var OptionsLibABI = OptionsLibMetaData.ABI

// OptionsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptionsLibMetaData.Bin instead.
var OptionsLibBin = OptionsLibMetaData.Bin

// DeployOptionsLib deploys a new Ethereum contract, binding an instance of OptionsLib to it.
func DeployOptionsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptionsLib, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptionsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// OptionsLib is an auto generated Go binding around an Ethereum contract.
type OptionsLib struct {
	OptionsLibCaller     // Read-only binding to the contract
	OptionsLibTransactor // Write-only binding to the contract
	OptionsLibFilterer   // Log filterer for contract events
}

// OptionsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptionsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptionsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptionsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptionsLibSession struct {
	Contract     *OptionsLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptionsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptionsLibCallerSession struct {
	Contract *OptionsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OptionsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptionsLibTransactorSession struct {
	Contract     *OptionsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OptionsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptionsLibRaw struct {
	Contract *OptionsLib // Generic contract binding to access the raw methods on
}

// OptionsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptionsLibCallerRaw struct {
	Contract *OptionsLibCaller // Generic read-only contract binding to access the raw methods on
}

// OptionsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptionsLibTransactorRaw struct {
	Contract *OptionsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptionsLib creates a new instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLib(address common.Address, backend bind.ContractBackend) (*OptionsLib, error) {
	contract, err := bindOptionsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// NewOptionsLibCaller creates a new read-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibCaller(address common.Address, caller bind.ContractCaller) (*OptionsLibCaller, error) {
	contract, err := bindOptionsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibCaller{contract: contract}, nil
}

// NewOptionsLibTransactor creates a new write-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OptionsLibTransactor, error) {
	contract, err := bindOptionsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibTransactor{contract: contract}, nil
}

// NewOptionsLibFilterer creates a new log filterer instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OptionsLibFilterer, error) {
	contract, err := bindOptionsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptionsLibFilterer{contract: contract}, nil
}

// bindOptionsLib binds a generic wrapper to an already deployed contract.
func bindOptionsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.OptionsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transact(opts, method, params...)
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

// OwnableAppMetaData contains all meta data concerning the OwnableApp contract.
var OwnableAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainApp__CallerNotInterchainClient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__InterchainClientNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleNotAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__ReceiverNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"AppConfigV1Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"AppLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"ExecutionServiceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"interchainClient\",\"type\":\"address\"}],\"name\":\"InterchainClientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"addTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppConfigV1\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedApp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"isAllowedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"linkRemoteApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteApp\",\"type\":\"address\"}],\"name\":\"linkRemoteAppEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"removeTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"appConfig\",\"type\":\"tuple\"}],\"name\":\"setAppConfigV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"setExecutionService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"interchain_\",\"type\":\"address\"}],\"name\":\"setInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cb5038fb": "addTrustedModule(address)",
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"7717a647": "getAppConfigV1()",
		"a20ce510": "getAppVersion()",
		"c313c807": "getExecutionService()",
		"b9b74b18": "getLinkedApp(uint256)",
		"287bc057": "getReceivingConfig()",
		"a45e107a": "getReceivingModules()",
		"ea13398f": "getSendingModules()",
		"70838975": "interchain()",
		"dc2b9075": "isAllowedSender(uint256,bytes32)",
		"51a30802": "linkRemoteApp(uint256,bytes32)",
		"af8fcc8e": "linkRemoteAppEVM(uint256,address)",
		"8da5cb5b": "owner()",
		"b70c40b3": "removeTrustedModule(address)",
		"715018a6": "renounceOwnership()",
		"0d32b505": "setAppConfigV1((uint256,uint256))",
		"496774b1": "setExecutionService(address)",
		"27efcbb7": "setInterchainClient(address)",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableAppABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableAppMetaData.ABI instead.
var OwnableAppABI = OwnableAppMetaData.ABI

// Deprecated: Use OwnableAppMetaData.Sigs instead.
// OwnableAppFuncSigs maps the 4-byte function signature to its string representation.
var OwnableAppFuncSigs = OwnableAppMetaData.Sigs

// OwnableApp is an auto generated Go binding around an Ethereum contract.
type OwnableApp struct {
	OwnableAppCaller     // Read-only binding to the contract
	OwnableAppTransactor // Write-only binding to the contract
	OwnableAppFilterer   // Log filterer for contract events
}

// OwnableAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableAppSession struct {
	Contract     *OwnableApp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableAppCallerSession struct {
	Contract *OwnableAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OwnableAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableAppTransactorSession struct {
	Contract     *OwnableAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OwnableAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableAppRaw struct {
	Contract *OwnableApp // Generic contract binding to access the raw methods on
}

// OwnableAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableAppCallerRaw struct {
	Contract *OwnableAppCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableAppTransactorRaw struct {
	Contract *OwnableAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableApp creates a new instance of OwnableApp, bound to a specific deployed contract.
func NewOwnableApp(address common.Address, backend bind.ContractBackend) (*OwnableApp, error) {
	contract, err := bindOwnableApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableApp{OwnableAppCaller: OwnableAppCaller{contract: contract}, OwnableAppTransactor: OwnableAppTransactor{contract: contract}, OwnableAppFilterer: OwnableAppFilterer{contract: contract}}, nil
}

// NewOwnableAppCaller creates a new read-only instance of OwnableApp, bound to a specific deployed contract.
func NewOwnableAppCaller(address common.Address, caller bind.ContractCaller) (*OwnableAppCaller, error) {
	contract, err := bindOwnableApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableAppCaller{contract: contract}, nil
}

// NewOwnableAppTransactor creates a new write-only instance of OwnableApp, bound to a specific deployed contract.
func NewOwnableAppTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableAppTransactor, error) {
	contract, err := bindOwnableApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableAppTransactor{contract: contract}, nil
}

// NewOwnableAppFilterer creates a new log filterer instance of OwnableApp, bound to a specific deployed contract.
func NewOwnableAppFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableAppFilterer, error) {
	contract, err := bindOwnableApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableAppFilterer{contract: contract}, nil
}

// bindOwnableApp binds a generic wrapper to an already deployed contract.
func bindOwnableApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableApp *OwnableAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableApp.Contract.OwnableAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableApp *OwnableAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableApp.Contract.OwnableAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableApp *OwnableAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableApp.Contract.OwnableAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableApp *OwnableAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableApp *OwnableAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableApp *OwnableAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableApp.Contract.contract.Transact(opts, method, params...)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_OwnableApp *OwnableAppCaller) GetAppConfigV1(opts *bind.CallOpts) (AppConfigV1, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "getAppConfigV1")

	if err != nil {
		return *new(AppConfigV1), err
	}

	out0 := *abi.ConvertType(out[0], new(AppConfigV1)).(*AppConfigV1)

	return out0, err

}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_OwnableApp *OwnableAppSession) GetAppConfigV1() (AppConfigV1, error) {
	return _OwnableApp.Contract.GetAppConfigV1(&_OwnableApp.CallOpts)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_OwnableApp *OwnableAppCallerSession) GetAppConfigV1() (AppConfigV1, error) {
	return _OwnableApp.Contract.GetAppConfigV1(&_OwnableApp.CallOpts)
}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_OwnableApp *OwnableAppCaller) GetAppVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "getAppVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_OwnableApp *OwnableAppSession) GetAppVersion() (*big.Int, error) {
	return _OwnableApp.Contract.GetAppVersion(&_OwnableApp.CallOpts)
}

// GetAppVersion is a free data retrieval call binding the contract method 0xa20ce510.
//
// Solidity: function getAppVersion() pure returns(uint256)
func (_OwnableApp *OwnableAppCallerSession) GetAppVersion() (*big.Int, error) {
	return _OwnableApp.Contract.GetAppVersion(&_OwnableApp.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_OwnableApp *OwnableAppCaller) GetExecutionService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "getExecutionService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_OwnableApp *OwnableAppSession) GetExecutionService() (common.Address, error) {
	return _OwnableApp.Contract.GetExecutionService(&_OwnableApp.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_OwnableApp *OwnableAppCallerSession) GetExecutionService() (common.Address, error) {
	return _OwnableApp.Contract.GetExecutionService(&_OwnableApp.CallOpts)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_OwnableApp *OwnableAppCaller) GetLinkedApp(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "getLinkedApp", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_OwnableApp *OwnableAppSession) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _OwnableApp.Contract.GetLinkedApp(&_OwnableApp.CallOpts, chainId)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_OwnableApp *OwnableAppCallerSession) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _OwnableApp.Contract.GetLinkedApp(&_OwnableApp.CallOpts, chainId)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_OwnableApp *OwnableAppCaller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "getReceivingConfig")

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
func (_OwnableApp *OwnableAppSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _OwnableApp.Contract.GetReceivingConfig(&_OwnableApp.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_OwnableApp *OwnableAppCallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _OwnableApp.Contract.GetReceivingConfig(&_OwnableApp.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_OwnableApp *OwnableAppCaller) GetReceivingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "getReceivingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_OwnableApp *OwnableAppSession) GetReceivingModules() ([]common.Address, error) {
	return _OwnableApp.Contract.GetReceivingModules(&_OwnableApp.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_OwnableApp *OwnableAppCallerSession) GetReceivingModules() ([]common.Address, error) {
	return _OwnableApp.Contract.GetReceivingModules(&_OwnableApp.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_OwnableApp *OwnableAppCaller) GetSendingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "getSendingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_OwnableApp *OwnableAppSession) GetSendingModules() ([]common.Address, error) {
	return _OwnableApp.Contract.GetSendingModules(&_OwnableApp.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_OwnableApp *OwnableAppCallerSession) GetSendingModules() ([]common.Address, error) {
	return _OwnableApp.Contract.GetSendingModules(&_OwnableApp.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_OwnableApp *OwnableAppCaller) Interchain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "interchain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_OwnableApp *OwnableAppSession) Interchain() (common.Address, error) {
	return _OwnableApp.Contract.Interchain(&_OwnableApp.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_OwnableApp *OwnableAppCallerSession) Interchain() (common.Address, error) {
	return _OwnableApp.Contract.Interchain(&_OwnableApp.CallOpts)
}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_OwnableApp *OwnableAppCaller) IsAllowedSender(opts *bind.CallOpts, srcChainId *big.Int, sender [32]byte) (bool, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "isAllowedSender", srcChainId, sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_OwnableApp *OwnableAppSession) IsAllowedSender(srcChainId *big.Int, sender [32]byte) (bool, error) {
	return _OwnableApp.Contract.IsAllowedSender(&_OwnableApp.CallOpts, srcChainId, sender)
}

// IsAllowedSender is a free data retrieval call binding the contract method 0xdc2b9075.
//
// Solidity: function isAllowedSender(uint256 srcChainId, bytes32 sender) view returns(bool)
func (_OwnableApp *OwnableAppCallerSession) IsAllowedSender(srcChainId *big.Int, sender [32]byte) (bool, error) {
	return _OwnableApp.Contract.IsAllowedSender(&_OwnableApp.CallOpts, srcChainId, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableApp *OwnableAppCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableApp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableApp *OwnableAppSession) Owner() (common.Address, error) {
	return _OwnableApp.Contract.Owner(&_OwnableApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableApp *OwnableAppCallerSession) Owner() (common.Address, error) {
	return _OwnableApp.Contract.Owner(&_OwnableApp.CallOpts)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_OwnableApp *OwnableAppTransactor) AddTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "addTrustedModule", module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_OwnableApp *OwnableAppSession) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.AddTrustedModule(&_OwnableApp.TransactOpts, module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_OwnableApp *OwnableAppTransactorSession) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.AddTrustedModule(&_OwnableApp.TransactOpts, module)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_OwnableApp *OwnableAppTransactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_OwnableApp *OwnableAppSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _OwnableApp.Contract.AppReceive(&_OwnableApp.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_OwnableApp *OwnableAppTransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _OwnableApp.Contract.AppReceive(&_OwnableApp.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_OwnableApp *OwnableAppTransactor) LinkRemoteApp(opts *bind.TransactOpts, chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "linkRemoteApp", chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_OwnableApp *OwnableAppSession) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _OwnableApp.Contract.LinkRemoteApp(&_OwnableApp.TransactOpts, chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_OwnableApp *OwnableAppTransactorSession) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _OwnableApp.Contract.LinkRemoteApp(&_OwnableApp.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_OwnableApp *OwnableAppTransactor) LinkRemoteAppEVM(opts *bind.TransactOpts, chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "linkRemoteAppEVM", chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_OwnableApp *OwnableAppSession) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.LinkRemoteAppEVM(&_OwnableApp.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_OwnableApp *OwnableAppTransactorSession) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.LinkRemoteAppEVM(&_OwnableApp.TransactOpts, chainId, remoteApp)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_OwnableApp *OwnableAppTransactor) RemoveTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "removeTrustedModule", module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_OwnableApp *OwnableAppSession) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.RemoveTrustedModule(&_OwnableApp.TransactOpts, module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_OwnableApp *OwnableAppTransactorSession) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.RemoveTrustedModule(&_OwnableApp.TransactOpts, module)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableApp *OwnableAppTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableApp *OwnableAppSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableApp.Contract.RenounceOwnership(&_OwnableApp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableApp *OwnableAppTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableApp.Contract.RenounceOwnership(&_OwnableApp.TransactOpts)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_OwnableApp *OwnableAppTransactor) SetAppConfigV1(opts *bind.TransactOpts, appConfig AppConfigV1) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "setAppConfigV1", appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_OwnableApp *OwnableAppSession) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _OwnableApp.Contract.SetAppConfigV1(&_OwnableApp.TransactOpts, appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_OwnableApp *OwnableAppTransactorSession) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _OwnableApp.Contract.SetAppConfigV1(&_OwnableApp.TransactOpts, appConfig)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_OwnableApp *OwnableAppTransactor) SetExecutionService(opts *bind.TransactOpts, executionService common.Address) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "setExecutionService", executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_OwnableApp *OwnableAppSession) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.SetExecutionService(&_OwnableApp.TransactOpts, executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_OwnableApp *OwnableAppTransactorSession) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.SetExecutionService(&_OwnableApp.TransactOpts, executionService)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address interchain_) returns()
func (_OwnableApp *OwnableAppTransactor) SetInterchainClient(opts *bind.TransactOpts, interchain_ common.Address) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "setInterchainClient", interchain_)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address interchain_) returns()
func (_OwnableApp *OwnableAppSession) SetInterchainClient(interchain_ common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.SetInterchainClient(&_OwnableApp.TransactOpts, interchain_)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address interchain_) returns()
func (_OwnableApp *OwnableAppTransactorSession) SetInterchainClient(interchain_ common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.SetInterchainClient(&_OwnableApp.TransactOpts, interchain_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableApp *OwnableAppTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableApp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableApp *OwnableAppSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.TransferOwnership(&_OwnableApp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableApp *OwnableAppTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableApp.Contract.TransferOwnership(&_OwnableApp.TransactOpts, newOwner)
}

// OwnableAppAppConfigV1SetIterator is returned from FilterAppConfigV1Set and is used to iterate over the raw logs and unpacked data for AppConfigV1Set events raised by the OwnableApp contract.
type OwnableAppAppConfigV1SetIterator struct {
	Event *OwnableAppAppConfigV1Set // Event containing the contract specifics and raw log

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
func (it *OwnableAppAppConfigV1SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableAppAppConfigV1Set)
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
		it.Event = new(OwnableAppAppConfigV1Set)
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
func (it *OwnableAppAppConfigV1SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableAppAppConfigV1SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableAppAppConfigV1Set represents a AppConfigV1Set event raised by the OwnableApp contract.
type OwnableAppAppConfigV1Set struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAppConfigV1Set is a free log retrieval operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_OwnableApp *OwnableAppFilterer) FilterAppConfigV1Set(opts *bind.FilterOpts) (*OwnableAppAppConfigV1SetIterator, error) {

	logs, sub, err := _OwnableApp.contract.FilterLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return &OwnableAppAppConfigV1SetIterator{contract: _OwnableApp.contract, event: "AppConfigV1Set", logs: logs, sub: sub}, nil
}

// WatchAppConfigV1Set is a free log subscription operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_OwnableApp *OwnableAppFilterer) WatchAppConfigV1Set(opts *bind.WatchOpts, sink chan<- *OwnableAppAppConfigV1Set) (event.Subscription, error) {

	logs, sub, err := _OwnableApp.contract.WatchLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableAppAppConfigV1Set)
				if err := _OwnableApp.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
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

// ParseAppConfigV1Set is a log parse operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_OwnableApp *OwnableAppFilterer) ParseAppConfigV1Set(log types.Log) (*OwnableAppAppConfigV1Set, error) {
	event := new(OwnableAppAppConfigV1Set)
	if err := _OwnableApp.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableAppAppLinkedIterator is returned from FilterAppLinked and is used to iterate over the raw logs and unpacked data for AppLinked events raised by the OwnableApp contract.
type OwnableAppAppLinkedIterator struct {
	Event *OwnableAppAppLinked // Event containing the contract specifics and raw log

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
func (it *OwnableAppAppLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableAppAppLinked)
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
		it.Event = new(OwnableAppAppLinked)
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
func (it *OwnableAppAppLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableAppAppLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableAppAppLinked represents a AppLinked event raised by the OwnableApp contract.
type OwnableAppAppLinked struct {
	ChainId   *big.Int
	RemoteApp [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppLinked is a free log retrieval operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_OwnableApp *OwnableAppFilterer) FilterAppLinked(opts *bind.FilterOpts, chainId []*big.Int, remoteApp [][32]byte) (*OwnableAppAppLinkedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _OwnableApp.contract.FilterLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return &OwnableAppAppLinkedIterator{contract: _OwnableApp.contract, event: "AppLinked", logs: logs, sub: sub}, nil
}

// WatchAppLinked is a free log subscription operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_OwnableApp *OwnableAppFilterer) WatchAppLinked(opts *bind.WatchOpts, sink chan<- *OwnableAppAppLinked, chainId []*big.Int, remoteApp [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _OwnableApp.contract.WatchLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableAppAppLinked)
				if err := _OwnableApp.contract.UnpackLog(event, "AppLinked", log); err != nil {
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

// ParseAppLinked is a log parse operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_OwnableApp *OwnableAppFilterer) ParseAppLinked(log types.Log) (*OwnableAppAppLinked, error) {
	event := new(OwnableAppAppLinked)
	if err := _OwnableApp.contract.UnpackLog(event, "AppLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableAppExecutionServiceSetIterator is returned from FilterExecutionServiceSet and is used to iterate over the raw logs and unpacked data for ExecutionServiceSet events raised by the OwnableApp contract.
type OwnableAppExecutionServiceSetIterator struct {
	Event *OwnableAppExecutionServiceSet // Event containing the contract specifics and raw log

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
func (it *OwnableAppExecutionServiceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableAppExecutionServiceSet)
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
		it.Event = new(OwnableAppExecutionServiceSet)
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
func (it *OwnableAppExecutionServiceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableAppExecutionServiceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableAppExecutionServiceSet represents a ExecutionServiceSet event raised by the OwnableApp contract.
type OwnableAppExecutionServiceSet struct {
	ExecutionService common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutionServiceSet is a free log retrieval operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_OwnableApp *OwnableAppFilterer) FilterExecutionServiceSet(opts *bind.FilterOpts) (*OwnableAppExecutionServiceSetIterator, error) {

	logs, sub, err := _OwnableApp.contract.FilterLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return &OwnableAppExecutionServiceSetIterator{contract: _OwnableApp.contract, event: "ExecutionServiceSet", logs: logs, sub: sub}, nil
}

// WatchExecutionServiceSet is a free log subscription operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_OwnableApp *OwnableAppFilterer) WatchExecutionServiceSet(opts *bind.WatchOpts, sink chan<- *OwnableAppExecutionServiceSet) (event.Subscription, error) {

	logs, sub, err := _OwnableApp.contract.WatchLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableAppExecutionServiceSet)
				if err := _OwnableApp.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
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

// ParseExecutionServiceSet is a log parse operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_OwnableApp *OwnableAppFilterer) ParseExecutionServiceSet(log types.Log) (*OwnableAppExecutionServiceSet, error) {
	event := new(OwnableAppExecutionServiceSet)
	if err := _OwnableApp.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableAppInterchainClientSetIterator is returned from FilterInterchainClientSet and is used to iterate over the raw logs and unpacked data for InterchainClientSet events raised by the OwnableApp contract.
type OwnableAppInterchainClientSetIterator struct {
	Event *OwnableAppInterchainClientSet // Event containing the contract specifics and raw log

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
func (it *OwnableAppInterchainClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableAppInterchainClientSet)
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
		it.Event = new(OwnableAppInterchainClientSet)
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
func (it *OwnableAppInterchainClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableAppInterchainClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableAppInterchainClientSet represents a InterchainClientSet event raised by the OwnableApp contract.
type OwnableAppInterchainClientSet struct {
	InterchainClient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientSet is a free log retrieval operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_OwnableApp *OwnableAppFilterer) FilterInterchainClientSet(opts *bind.FilterOpts) (*OwnableAppInterchainClientSetIterator, error) {

	logs, sub, err := _OwnableApp.contract.FilterLogs(opts, "InterchainClientSet")
	if err != nil {
		return nil, err
	}
	return &OwnableAppInterchainClientSetIterator{contract: _OwnableApp.contract, event: "InterchainClientSet", logs: logs, sub: sub}, nil
}

// WatchInterchainClientSet is a free log subscription operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_OwnableApp *OwnableAppFilterer) WatchInterchainClientSet(opts *bind.WatchOpts, sink chan<- *OwnableAppInterchainClientSet) (event.Subscription, error) {

	logs, sub, err := _OwnableApp.contract.WatchLogs(opts, "InterchainClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableAppInterchainClientSet)
				if err := _OwnableApp.contract.UnpackLog(event, "InterchainClientSet", log); err != nil {
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

// ParseInterchainClientSet is a log parse operation binding the contract event 0xeec21067aa320b611516f448454be9fae691403167636e737345cab1f262d5d7.
//
// Solidity: event InterchainClientSet(address interchainClient)
func (_OwnableApp *OwnableAppFilterer) ParseInterchainClientSet(log types.Log) (*OwnableAppInterchainClientSet, error) {
	event := new(OwnableAppInterchainClientSet)
	if err := _OwnableApp.contract.UnpackLog(event, "InterchainClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableAppOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableApp contract.
type OwnableAppOwnershipTransferredIterator struct {
	Event *OwnableAppOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableAppOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableAppOwnershipTransferred)
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
		it.Event = new(OwnableAppOwnershipTransferred)
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
func (it *OwnableAppOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableAppOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableAppOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableApp contract.
type OwnableAppOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableApp *OwnableAppFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableAppOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableApp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableAppOwnershipTransferredIterator{contract: _OwnableApp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableApp *OwnableAppFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableAppOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableApp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableAppOwnershipTransferred)
				if err := _OwnableApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OwnableApp *OwnableAppFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableAppOwnershipTransferred, error) {
	event := new(OwnableAppOwnershipTransferred)
	if err := _OwnableApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableAppTrustedModuleAddedIterator is returned from FilterTrustedModuleAdded and is used to iterate over the raw logs and unpacked data for TrustedModuleAdded events raised by the OwnableApp contract.
type OwnableAppTrustedModuleAddedIterator struct {
	Event *OwnableAppTrustedModuleAdded // Event containing the contract specifics and raw log

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
func (it *OwnableAppTrustedModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableAppTrustedModuleAdded)
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
		it.Event = new(OwnableAppTrustedModuleAdded)
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
func (it *OwnableAppTrustedModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableAppTrustedModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableAppTrustedModuleAdded represents a TrustedModuleAdded event raised by the OwnableApp contract.
type OwnableAppTrustedModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleAdded is a free log retrieval operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_OwnableApp *OwnableAppFilterer) FilterTrustedModuleAdded(opts *bind.FilterOpts) (*OwnableAppTrustedModuleAddedIterator, error) {

	logs, sub, err := _OwnableApp.contract.FilterLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return &OwnableAppTrustedModuleAddedIterator{contract: _OwnableApp.contract, event: "TrustedModuleAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleAdded is a free log subscription operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_OwnableApp *OwnableAppFilterer) WatchTrustedModuleAdded(opts *bind.WatchOpts, sink chan<- *OwnableAppTrustedModuleAdded) (event.Subscription, error) {

	logs, sub, err := _OwnableApp.contract.WatchLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableAppTrustedModuleAdded)
				if err := _OwnableApp.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
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

// ParseTrustedModuleAdded is a log parse operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_OwnableApp *OwnableAppFilterer) ParseTrustedModuleAdded(log types.Log) (*OwnableAppTrustedModuleAdded, error) {
	event := new(OwnableAppTrustedModuleAdded)
	if err := _OwnableApp.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableAppTrustedModuleRemovedIterator is returned from FilterTrustedModuleRemoved and is used to iterate over the raw logs and unpacked data for TrustedModuleRemoved events raised by the OwnableApp contract.
type OwnableAppTrustedModuleRemovedIterator struct {
	Event *OwnableAppTrustedModuleRemoved // Event containing the contract specifics and raw log

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
func (it *OwnableAppTrustedModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableAppTrustedModuleRemoved)
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
		it.Event = new(OwnableAppTrustedModuleRemoved)
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
func (it *OwnableAppTrustedModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableAppTrustedModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableAppTrustedModuleRemoved represents a TrustedModuleRemoved event raised by the OwnableApp contract.
type OwnableAppTrustedModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleRemoved is a free log retrieval operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_OwnableApp *OwnableAppFilterer) FilterTrustedModuleRemoved(opts *bind.FilterOpts) (*OwnableAppTrustedModuleRemovedIterator, error) {

	logs, sub, err := _OwnableApp.contract.FilterLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &OwnableAppTrustedModuleRemovedIterator{contract: _OwnableApp.contract, event: "TrustedModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleRemoved is a free log subscription operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_OwnableApp *OwnableAppFilterer) WatchTrustedModuleRemoved(opts *bind.WatchOpts, sink chan<- *OwnableAppTrustedModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _OwnableApp.contract.WatchLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableAppTrustedModuleRemoved)
				if err := _OwnableApp.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
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

// ParseTrustedModuleRemoved is a log parse operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_OwnableApp *OwnableAppFilterer) ParseTrustedModuleRemoved(log types.Log) (*OwnableAppTrustedModuleRemoved, error) {
	event := new(OwnableAppTrustedModuleRemoved)
	if err := _OwnableApp.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204666bb590d693f4485a5c61cc399e8450e0be7b19c7bd1bd2a78b5d9ca2520bd64736f6c63430008140033",
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
