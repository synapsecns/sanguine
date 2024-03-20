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

// InterchainBatch is an auto generated low-level Go binding around an user-defined struct.
type InterchainBatch struct {
	SrcChainId *big.Int
	DbNonce    *big.Int
	BatchRoot  [32]byte
}

// InterchainEntry is an auto generated low-level Go binding around an user-defined struct.
type InterchainEntry struct {
	SrcChainId *big.Int
	DbNonce    *big.Int
	EntryIndex uint64
	SrcWriter  [32]byte
	DataHash   [32]byte
}

// InterchainTransaction is an auto generated low-level Go binding around an user-defined struct.
type InterchainTransaction struct {
	SrcChainId  *big.Int
	SrcSender   [32]byte
	DstChainId  *big.Int
	DstReceiver [32]byte
	DbNonce     *big.Int
	EntryIndex  uint64
	Options     []byte
	Message     []byte
}

// InterchainTxDescriptor is an auto generated low-level Go binding around an user-defined struct.
type InterchainTxDescriptor struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
}

// OptionsV1 is an auto generated low-level Go binding around an user-defined struct.
type OptionsV1 struct {
	GasLimit   *big.Int
	GasAirdrop *big.Int
}

// AppConfigLibMetaData contains all meta data concerning the AppConfigLib contract.
var AppConfigLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"AppConfigLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ae2dd045e155a35e99c1c7a6d3ebc0bae66e3ae75ad5df373f775d9b68b07efa64736f6c63430008140033",
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

// IExecutionFeesMetaData contains all meta data concerning the IExecutionFees contract.
var IExecutionFeesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionFees__AlreadyRecorded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFees__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFees__ZeroAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"accumulatedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"addExecutionFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"claimExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"executionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"recordExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"recordedExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"unclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unclaimed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"73f273fc": "accumulatedRewards(address)",
		"ffecec7e": "addExecutionFee(uint256,bytes32)",
		"4e497dac": "claimExecutionFees(address)",
		"936fd4db": "executionFee(uint256,bytes32)",
		"0676b706": "recordExecutor(uint256,bytes32,address)",
		"d01e09a6": "recordedExecutor(uint256,bytes32)",
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

// ExecutionFee is a free data retrieval call binding the contract method 0x936fd4db.
//
// Solidity: function executionFee(uint256 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_IExecutionFees *IExecutionFeesCaller) ExecutionFee(opts *bind.CallOpts, dstChainId *big.Int, transactionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionFees.contract.Call(opts, &out, "executionFee", dstChainId, transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutionFee is a free data retrieval call binding the contract method 0x936fd4db.
//
// Solidity: function executionFee(uint256 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_IExecutionFees *IExecutionFeesSession) ExecutionFee(dstChainId *big.Int, transactionId [32]byte) (*big.Int, error) {
	return _IExecutionFees.Contract.ExecutionFee(&_IExecutionFees.CallOpts, dstChainId, transactionId)
}

// ExecutionFee is a free data retrieval call binding the contract method 0x936fd4db.
//
// Solidity: function executionFee(uint256 dstChainId, bytes32 transactionId) view returns(uint256 fee)
func (_IExecutionFees *IExecutionFeesCallerSession) ExecutionFee(dstChainId *big.Int, transactionId [32]byte) (*big.Int, error) {
	return _IExecutionFees.Contract.ExecutionFee(&_IExecutionFees.CallOpts, dstChainId, transactionId)
}

// RecordedExecutor is a free data retrieval call binding the contract method 0xd01e09a6.
//
// Solidity: function recordedExecutor(uint256 dstChainId, bytes32 transactionId) view returns(address executor)
func (_IExecutionFees *IExecutionFeesCaller) RecordedExecutor(opts *bind.CallOpts, dstChainId *big.Int, transactionId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IExecutionFees.contract.Call(opts, &out, "recordedExecutor", dstChainId, transactionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecordedExecutor is a free data retrieval call binding the contract method 0xd01e09a6.
//
// Solidity: function recordedExecutor(uint256 dstChainId, bytes32 transactionId) view returns(address executor)
func (_IExecutionFees *IExecutionFeesSession) RecordedExecutor(dstChainId *big.Int, transactionId [32]byte) (common.Address, error) {
	return _IExecutionFees.Contract.RecordedExecutor(&_IExecutionFees.CallOpts, dstChainId, transactionId)
}

// RecordedExecutor is a free data retrieval call binding the contract method 0xd01e09a6.
//
// Solidity: function recordedExecutor(uint256 dstChainId, bytes32 transactionId) view returns(address executor)
func (_IExecutionFees *IExecutionFeesCallerSession) RecordedExecutor(dstChainId *big.Int, transactionId [32]byte) (common.Address, error) {
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectDstChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NoLinkedClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__NotEVMClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotEnoughResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotRemoteChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxNotExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainClientV1__ZeroReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainClientV1__ZeroRequiredResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"getExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"getExecutorById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedClient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedClientEVM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSend\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"internalType\":\"structInterchainTxDescriptor\",\"name\":\"desc\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSendEVM\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"internalType\":\"structInterchainTxDescriptor\",\"name\":\"desc\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"isExecutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionFees_\",\"type\":\"address\"}],\"name\":\"setExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"writeExecutionProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchNotFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"existingBatchRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"newBatch\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingBatches\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__InvalidEntryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafsPaginated\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEntryIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"67b1f42e": "checkVerification(address,(uint256,uint256,uint64,bytes32,bytes32),bytes32[])",
		"5ac44282": "getBatch(uint256)",
		"d63020bb": "getBatchLeafs(uint256)",
		"25a1641d": "getBatchLeafsPaginated(uint256,uint64,uint64)",
		"b955e9b9": "getBatchSize(uint256)",
		"f338140e": "getDBNonce()",
		"1725fd30": "getEntry(uint256,uint64)",
		"4f84d040": "getEntryProof(uint256,uint64)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"aa2f06ae": "getNextEntryIndex()",
		"84b1c8b8": "requestBatchVerification(uint256,uint256,address[])",
		"05d0728c": "verifyRemoteBatch((uint256,uint256,bytes32))",
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

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) CheckVerification(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "checkVerification", dstModule, entry, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) CheckVerification(dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	return _IInterchainDB.Contract.CheckVerification(&_IInterchainDB.CallOpts, dstModule, entry, proof)
}

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) CheckVerification(dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	return _IInterchainDB.Contract.CheckVerification(&_IInterchainDB.CallOpts, dstModule, entry, proof)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetBatch(opts *bind.CallOpts, dbNonce *big.Int) (InterchainBatch, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatch", dbNonce)

	if err != nil {
		return *new(InterchainBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainBatch)).(*InterchainBatch)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetBatch(dbNonce *big.Int) (InterchainBatch, error) {
	return _IInterchainDB.Contract.GetBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetBatch(dbNonce *big.Int) (InterchainBatch, error) {
	return _IInterchainDB.Contract.GetBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCaller) GetBatchLeafs(opts *bind.CallOpts, dbNonce *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchLeafs", dbNonce)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBSession) GetBatchLeafs(dbNonce *big.Int) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafs(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchLeafs(dbNonce *big.Int) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafs(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCaller) GetBatchLeafsPaginated(opts *bind.CallOpts, dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchLeafsPaginated", dbNonce, start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBSession) GetBatchLeafsPaginated(dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafsPaginated(&_IInterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchLeafsPaginated(dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafsPaginated(&_IInterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBCaller) GetBatchSize(opts *bind.CallOpts, dbNonce *big.Int) (uint64, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchSize", dbNonce)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBSession) GetBatchSize(dbNonce *big.Int) (uint64, error) {
	return _IInterchainDB.Contract.GetBatchSize(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchSize(dbNonce *big.Int) (uint64, error) {
	return _IInterchainDB.Contract.GetBatchSize(&_IInterchainDB.CallOpts, dbNonce)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetDBNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getDBNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetDBNonce() (*big.Int, error) {
	return _IInterchainDB.Contract.GetDBNonce(&_IInterchainDB.CallOpts)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetDBNonce() (*big.Int, error) {
	return _IInterchainDB.Contract.GetDBNonce(&_IInterchainDB.CallOpts)
}

// GetEntry is a free data retrieval call binding the contract method 0x1725fd30.
//
// Solidity: function getEntry(uint256 dbNonce, uint64 entryIndex) view returns((uint256,uint256,uint64,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetEntry(opts *bind.CallOpts, dbNonce *big.Int, entryIndex uint64) (InterchainEntry, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntry", dbNonce, entryIndex)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0x1725fd30.
//
// Solidity: function getEntry(uint256 dbNonce, uint64 entryIndex) view returns((uint256,uint256,uint64,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetEntry(dbNonce *big.Int, entryIndex uint64) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntry is a free data retrieval call binding the contract method 0x1725fd30.
//
// Solidity: function getEntry(uint256 dbNonce, uint64 entryIndex) view returns((uint256,uint256,uint64,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetEntry(dbNonce *big.Int, entryIndex uint64) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBCaller) GetEntryProof(opts *bind.CallOpts, dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntryProof", dbNonce, entryIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBSession) GetEntryProof(dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetEntryProof(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBCallerSession) GetEntryProof(dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetEntryProof(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetInterchainFee(dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, dstChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetInterchainFee(dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, dstChainId, srcModules)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBCaller) GetNextEntryIndex(opts *bind.CallOpts) (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getNextEntryIndex")

	outstruct := new(struct {
		DbNonce    *big.Int
		EntryIndex uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DbNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EntryIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) GetNextEntryIndex() (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	return _IInterchainDB.Contract.GetNextEntryIndex(&_IInterchainDB.CallOpts)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBCallerSession) GetNextEntryIndex() (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	return _IInterchainDB.Contract.GetNextEntryIndex(&_IInterchainDB.CallOpts)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestBatchVerification", dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestBatchVerification(dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestBatchVerification(&_IInterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestBatchVerification(dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestBatchVerification(&_IInterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_IInterchainDB *IInterchainDBTransactor) VerifyRemoteBatch(opts *bind.TransactOpts, batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "verifyRemoteBatch", batch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_IInterchainDB *IInterchainDBSession) VerifyRemoteBatch(batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyRemoteBatch(&_IInterchainDB.TransactOpts, batch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_IInterchainDB *IInterchainDBTransactorSession) VerifyRemoteBatch(batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyRemoteBatch(&_IInterchainDB.TransactOpts, batch)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// InterchainBatchLibMetaData contains all meta data concerning the InterchainBatchLib contract.
var InterchainBatchLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208099eef78f304061842ba5fc306dd72b140378638234593b4604163e85e9286964736f6c63430008140033",
}

// InterchainBatchLibABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainBatchLibMetaData.ABI instead.
var InterchainBatchLibABI = InterchainBatchLibMetaData.ABI

// InterchainBatchLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainBatchLibMetaData.Bin instead.
var InterchainBatchLibBin = InterchainBatchLibMetaData.Bin

// DeployInterchainBatchLib deploys a new Ethereum contract, binding an instance of InterchainBatchLib to it.
func DeployInterchainBatchLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainBatchLib, error) {
	parsed, err := InterchainBatchLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainBatchLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainBatchLib{InterchainBatchLibCaller: InterchainBatchLibCaller{contract: contract}, InterchainBatchLibTransactor: InterchainBatchLibTransactor{contract: contract}, InterchainBatchLibFilterer: InterchainBatchLibFilterer{contract: contract}}, nil
}

// InterchainBatchLib is an auto generated Go binding around an Ethereum contract.
type InterchainBatchLib struct {
	InterchainBatchLibCaller     // Read-only binding to the contract
	InterchainBatchLibTransactor // Write-only binding to the contract
	InterchainBatchLibFilterer   // Log filterer for contract events
}

// InterchainBatchLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainBatchLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainBatchLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainBatchLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainBatchLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainBatchLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainBatchLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainBatchLibSession struct {
	Contract     *InterchainBatchLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterchainBatchLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainBatchLibCallerSession struct {
	Contract *InterchainBatchLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterchainBatchLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainBatchLibTransactorSession struct {
	Contract     *InterchainBatchLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterchainBatchLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainBatchLibRaw struct {
	Contract *InterchainBatchLib // Generic contract binding to access the raw methods on
}

// InterchainBatchLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainBatchLibCallerRaw struct {
	Contract *InterchainBatchLibCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainBatchLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainBatchLibTransactorRaw struct {
	Contract *InterchainBatchLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainBatchLib creates a new instance of InterchainBatchLib, bound to a specific deployed contract.
func NewInterchainBatchLib(address common.Address, backend bind.ContractBackend) (*InterchainBatchLib, error) {
	contract, err := bindInterchainBatchLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainBatchLib{InterchainBatchLibCaller: InterchainBatchLibCaller{contract: contract}, InterchainBatchLibTransactor: InterchainBatchLibTransactor{contract: contract}, InterchainBatchLibFilterer: InterchainBatchLibFilterer{contract: contract}}, nil
}

// NewInterchainBatchLibCaller creates a new read-only instance of InterchainBatchLib, bound to a specific deployed contract.
func NewInterchainBatchLibCaller(address common.Address, caller bind.ContractCaller) (*InterchainBatchLibCaller, error) {
	contract, err := bindInterchainBatchLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainBatchLibCaller{contract: contract}, nil
}

// NewInterchainBatchLibTransactor creates a new write-only instance of InterchainBatchLib, bound to a specific deployed contract.
func NewInterchainBatchLibTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainBatchLibTransactor, error) {
	contract, err := bindInterchainBatchLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainBatchLibTransactor{contract: contract}, nil
}

// NewInterchainBatchLibFilterer creates a new log filterer instance of InterchainBatchLib, bound to a specific deployed contract.
func NewInterchainBatchLibFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainBatchLibFilterer, error) {
	contract, err := bindInterchainBatchLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainBatchLibFilterer{contract: contract}, nil
}

// bindInterchainBatchLib binds a generic wrapper to an already deployed contract.
func bindInterchainBatchLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainBatchLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainBatchLib *InterchainBatchLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainBatchLib.Contract.InterchainBatchLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainBatchLib *InterchainBatchLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainBatchLib.Contract.InterchainBatchLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainBatchLib *InterchainBatchLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainBatchLib.Contract.InterchainBatchLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainBatchLib *InterchainBatchLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainBatchLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainBatchLib *InterchainBatchLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainBatchLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainBatchLib *InterchainBatchLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainBatchLib.Contract.contract.Transact(opts, method, params...)
}

// InterchainClientV1MetaData contains all meta data concerning the InterchainClientV1 contract.
var InterchainClientV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"interchainDB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"AppConfigLib__IncorrectVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectDstChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NoLinkedClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__NotEVMClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotEnoughResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotRemoteChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxNotExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainClientV1__ZeroReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainClientV1__ZeroRequiredResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionFees\",\"type\":\"address\"}],\"name\":\"ExecutionFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionProofWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"}],\"name\":\"InterchainTransactionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"InterchainTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"LinkedClientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INTERCHAIN_DB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedOptions\",\"type\":\"bytes\"}],\"name\":\"decodeOptions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAirdrop\",\"type\":\"uint256\"}],\"internalType\":\"structOptionsV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structInterchainTransaction\",\"name\":\"icTx\",\"type\":\"tuple\"}],\"name\":\"encodeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionFees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"getExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"getExecutorById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedClient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedClientEVM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"linkedClientEVM\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSend\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"internalType\":\"structInterchainTxDescriptor\",\"name\":\"desc\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSendEVM\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"internalType\":\"structInterchainTxDescriptor\",\"name\":\"desc\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"isExecutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionFees_\",\"type\":\"address\"}],\"name\":\"setExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"writeExecutionProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e4c61247": "INTERCHAIN_DB()",
		"d5e788a0": "decodeOptions(bytes)",
		"7c80a90f": "encodeTransaction((uint256,bytes32,uint256,bytes32,uint256,uint64,bytes,bytes))",
		"7341eaf9": "executionFees()",
		"f92a79ff": "getExecutor(bytes)",
		"f1a61fac": "getExecutorById(bytes32)",
		"3c383e7b": "getInterchainFee(uint256,address,address[],bytes,bytes)",
		"aa102ec4": "getLinkedClient(uint256)",
		"02172a35": "getLinkedClientEVM(uint256)",
		"53b67d74": "interchainExecute(uint256,bytes,bytes32[])",
		"98939d28": "interchainSend(uint256,bytes32,address,address[],bytes,bytes)",
		"827f940d": "interchainSendEVM(uint256,address,address,address[],bytes,bytes)",
		"1450c281": "isExecutable(bytes,bytes32[])",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"3dc68b87": "setExecutionFees(address)",
		"f34234c8": "setLinkedClient(uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
		"90e81077": "writeExecutionProof(bytes32)",
	},
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002b7638038062002b768339810160408190526200003491620000f0565b806001600160a01b0381166200006457604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6200006f8162000083565b50506001600160a01b031660805262000128565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114620000eb57600080fd5b919050565b600080604083850312156200010457600080fd5b6200010f83620000d3565b91506200011f60208401620000d3565b90509250929050565b608051612a086200016e600039600081816103bd0152818161062101528181610bc0015281816114de0152818161159f015281816116fd0152611b9c0152612a086000f3fe6080604052600436106101445760003560e01c80638da5cb5b116100c0578063e4c6124711610074578063f2fde38b11610059578063f2fde38b14610422578063f34234c814610442578063f92a79ff1461046257600080fd5b8063e4c61247146103ab578063f1a61fac146103df57600080fd5b806398939d28116100a557806398939d281461033d578063aa102ec414610350578063d5e788a01461037057600080fd5b80638da5cb5b146102d457806390e81077146102ff57600080fd5b806353b67d74116101175780637341eaf9116100fc5780637341eaf91461023b5780637c80a90f14610268578063827f940d1461029557600080fd5b806353b67d7414610213578063715018a61461022657600080fd5b806302172a35146101495780631450c281146101935780633c383e7b146101c35780633dc68b87146101f1575b600080fd5b34801561015557600080fd5b50610169610164366004611c98565b610482565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561019f57600080fd5b506101b36101ae366004611d3f565b61052e565b604051901515815260200161018a565b3480156101cf57600080fd5b506101e36101de366004611dcd565b610598565b60405190815260200161018a565b3480156101fd57600080fd5b5061021161020c366004611e83565b6107f4565b005b610211610221366004611ea7565b610875565b34801561023257600080fd5b50610211610a6e565b34801561024757600080fd5b506001546101699073ffffffffffffffffffffffffffffffffffffffff1681565b34801561027457600080fd5b5061028861028336600461204a565b610a82565b60405161018a919061216f565b6102a86102a3366004612182565b610a93565b6040805182518152602080840151908201529181015167ffffffffffffffff169082015260600161018a565b3480156102e057600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610169565b34801561030b57600080fd5b5061031f61031a366004611c98565b610ae7565b6040805192835267ffffffffffffffff90911660208301520161018a565b6102a861034b36600461224d565b610ca6565b34801561035c57600080fd5b506101e361036b366004611c98565b610ce2565b34801561037c57600080fd5b5061039061038b366004612284565b610d33565b6040805182518152602092830151928101929092520161018a565b3480156103b757600080fd5b506101697f000000000000000000000000000000000000000000000000000000000000000081565b3480156103eb57600080fd5b506101696103fa366004611c98565b60009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b34801561042e57600080fd5b5061021161043d366004611e83565b610d50565b34801561044e57600080fd5b5061021161045d3660046122b9565b610db4565b34801561046e57600080fd5b5061016961047d3660046122db565b610e0b565b60004682036104c5576040517f57516f69000000000000000000000000000000000000000000000000000000008152600481018390526024015b60405180910390fd5b506000818152600260205260409020548073ffffffffffffffffffffffffffffffffffffffff81168114610528576040517f0a55a4eb000000000000000000000000000000000000000000000000000000008152600481018290526024016104bc565b50919050565b60008061057086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8b92505050565b905061057f8160c00151610ef7565b5061058b818585610f7b565b5060019695505050505050565b60006105a3896111b6565b506105e385858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ef792505050565b506040517ffc7686ec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063fc7686ec9061065a908c908b908b90600401612373565b602060405180830381865afa158015610677573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069b9190612396565b905073ffffffffffffffffffffffffffffffffffffffff8816156107e857600061073860008b6000801b6000808b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8f018190048102820181019092528d815292508d91508c908190840183828082843760009201919091525061124692505050565b90508873ffffffffffffffffffffffffffffffffffffffff1663c473e7e88b8360405160200161076891906123af565b6040516020818303038152906040525189896040518563ffffffff1660e01b81526004016107999493929190612465565b602060405180830381865afa1580156107b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107da9190612396565b6107e490836124b4565b9150505b98975050505050505050565b6107fc61130b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f829060200160405180910390a150565b60006108b685858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8b92505050565b905060006108c5828585610f7b565b600081815260036020526040812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000163317905560c08401519192509061090d90610ef7565b90508060200151341461095b5760208101516040517f2b36102500000000000000000000000000000000000000000000000000000000815234600482015260248101919091526044016104bc565b805188101561096957805197505b606083015173ffffffffffffffffffffffffffffffffffffffff166368a6984789348660000151876020015188608001518960a001518a60e001516040518863ffffffff1660e01b81526004016109c49594939291906124c7565b6000604051808303818589803b1580156109dd57600080fd5b5088f11580156109f1573d6000803e3d6000fd5b5050505050508260a0015167ffffffffffffffff168360800151837f9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c866000015187602001518860600151604051610a5c939291909283526020830191909152604082015260600190565b60405180910390a45050505050505050565b610a7661130b565b610a80600061135e565b565b6060610a8d826113d3565b92915050565b604080516060810182526000808252602082018190529181019190915273ffffffffffffffffffffffffffffffffffffffff8916610ad88b828b8b8b8b8b8b8b6113fc565b9b9a5050505050505050505050565b600081815260036020526040812054819073ffffffffffffffffffffffffffffffffffffffff1680610b48576040517fe99eb48d000000000000000000000000000000000000000000000000000000008152600481018590526024016104bc565b60008482604051602001610b7c92919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60408051808303601f1901815290829052805160208201207f2ad8c706000000000000000000000000000000000000000000000000000000008352600483015291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632ad8c7069060240160408051808303816000875af1158015610c1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c419190612512565b60405173ffffffffffffffffffffffffffffffffffffffff85168152919550935067ffffffffffffffff841690859087907f810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c9060200160405180910390a45050915091565b6040805160608101825260008082526020820181905291810191909152610cd48a8a8a8a8a8a8a8a8a6113fc565b9a9950505050505050505050565b6000468203610d20576040517f57516f69000000000000000000000000000000000000000000000000000000008152600481018390526024016104bc565b5060009081526002602052604090205490565b6040805180820190915260008082526020820152610a8d82610ef7565b610d5861130b565b73ffffffffffffffffffffffffffffffffffffffff8116610da8576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024016104bc565b610db18161135e565b50565b610dbc61130b565b60008281526002602090815260409182902083905581518481529081018390527fb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a910160405180910390a15050565b600080610e4d84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8b92505050565b905060036000610e5c83611ad1565b815260208101919091526040016000205473ffffffffffffffffffffffffffffffffffffffff16949350505050565b610ee36040518061010001604052806000815260200160008019168152602001600081526020016000801916815260200160008152602001600067ffffffffffffffff16815260200160608152602001606081525090565b81806020019051810190610a8d9190612587565b6040805180820190915260008082526020820152600080610f1784611b01565b9092509050600160ff83161015610f5f576040517fbd91a21500000000000000000000000000000000000000000000000000000000815260ff831660048201526024016104bc565b80806020019051810190610f73919061269c565b949350505050565b600080610f8b85600001516111b6565b905046856040015114610fd25784604001516040517f973253820000000000000000000000000000000000000000000000000000000081526004016104bc91815260200190565b610fdb85611ad1565b60008181526003602052604090205490925073ffffffffffffffffffffffffffffffffffffffff161561103d576040517fd80aeb91000000000000000000000000000000000000000000000000000000008152600481018390526024016104bc565b60006040518060a0016040528087600001518152602001876080015181526020018760a0015167ffffffffffffffff16815260200183815260200184815250905060008061108c886060015190565b73ffffffffffffffffffffffffffffffffffffffff1663287bc0576040518163ffffffff1660e01b8152600401600060405180830381865afa1580156110d6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110fe91908101906126b8565b91509150600061110d83611b23565b805190915060000361114b576040517fa09e214300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061115e83868b8b8660200151611b8b565b82519091508110156111a95781516040517f0bce4e850000000000000000000000000000000000000000000000000000000081526104bc918391600401918252602082015260400190565b5050505050509392505050565b60004682036111f4576040517f57516f69000000000000000000000000000000000000000000000000000000008152600481018390526024016104bc565b5060008181526002602052604081205490819003611241576040517f9a45110f000000000000000000000000000000000000000000000000000000008152600481018390526024016104bc565b919050565b61129e6040518061010001604052806000815260200160008019168152602001600081526020016000801916815260200160008152602001600067ffffffffffffffff16815260200160608152602001606081525090565b6040518061010001604052804681526020016112cd8a73ffffffffffffffffffffffffffffffffffffffff1690565b81526020018881526020018781526020018681526020018567ffffffffffffffff168152602001848152602001838152509050979650505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a80576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016104bc565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6060816040516020016113e691906123af565b6040516020818303038152906040529050919050565b60408051606081018252600080825260208201819052918101919091526114228a6111b6565b50600089900361145e576040517f0b3e41a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61149d85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ef792505050565b506040517ffc7686ec00000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063fc7686ec90611517908e908c908c90600401612373565b602060405180830381865afa158015611534573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115589190612396565b90508034101561159d576040517f9557ee49000000000000000000000000000000000000000000000000000000008152346004820152602481018290526044016104bc565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663aa2f06ae6040518163ffffffff1660e01b81526004016040805180830381865afa158015611607573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162b9190612512565b83602001846040018267ffffffffffffffff1667ffffffffffffffff1681525082815250505060006116e4338d8d866020015187604001518c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061124692505050565b90506116ef81611ad1565b8360000181815250506000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166367c769af85856040015188600001518f8f6040518663ffffffff1660e01b8152600401611763949392919061278e565b604080518083038185885af1158015611780573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906117a59190612512565b915091508460200151821480156117d35750846040015167ffffffffffffffff168167ffffffffffffffff16145b6117df576117df6127ae565b50503482810390831461188857600154604080840151865191517fffecec7e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9093169263ffecec7e92859261185592909190600401918252602082015260400190565b6000604051808303818588803b15801561186e57600080fd5b505af1158015611882573d6000803e3d6000fd5b50505050505b73ffffffffffffffffffffffffffffffffffffffff8b1615611a51578a73ffffffffffffffffffffffffffffffffffffffff1663e4e065228e6118ca856113d3565b5187516040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815261190e9392919087908f908f906004016127dd565b600060405180830381600087803b15801561192857600080fd5b505af115801561193c573d6000803e3d6000fd5b5050505060008b73ffffffffffffffffffffffffffffffffffffffff166362014bad6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561198d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119b19190612809565b600154604085810151885191517f0676b7060000000000000000000000000000000000000000000000000000000081526004810191909152602481019190915273ffffffffffffffffffffffffffffffffffffffff8084166044830152929350911690630676b70690606401600060405180830381600087803b158015611a3757600080fd5b505af1158015611a4b573d6000803e3d6000fd5b50505050505b8160a0015167ffffffffffffffff16826080015185600001517f1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b85604001518660200151876060015189888a60c001518b60e00151604051611ab99796959493929190612826565b60405180910390a45050509998505050505050505050565b600081604051602001611ae491906123af565b604051602081830303815290604052805190602001209050919050565b6000606082806020019051810190611b199190612869565b9094909350915050565b6040805180820190915260008082526020820152600080611b4384611b01565b9092509050600160ff83161015610f5f576040517fc3e3b66600000000000000000000000000000000000000000000000000000000815260ff831660048201526024016104bc565b6000805b8651811015611c8e5760007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166367b1f42e898481518110611be857611be86128c0565b60200260200101518989896040518563ffffffff1660e01b8152600401611c1294939291906128ef565b602060405180830381865afa158015611c2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c539190612396565b90508015801590611c6c575042611c6a85836124b4565b105b15611c7d57611c7a8361299a565b92505b50611c878161299a565b9050611b8f565b5095945050505050565b600060208284031215611caa57600080fd5b5035919050565b60008083601f840112611cc357600080fd5b50813567ffffffffffffffff811115611cdb57600080fd5b602083019150836020828501011115611cf357600080fd5b9250929050565b60008083601f840112611d0c57600080fd5b50813567ffffffffffffffff811115611d2457600080fd5b6020830191508360208260051b8501011115611cf357600080fd5b60008060008060408587031215611d5557600080fd5b843567ffffffffffffffff80821115611d6d57600080fd5b611d7988838901611cb1565b90965094506020870135915080821115611d9257600080fd5b50611d9f87828801611cfa565b95989497509550505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610db157600080fd5b60008060008060008060008060a0898b031215611de957600080fd5b883597506020890135611dfb81611dab565b9650604089013567ffffffffffffffff80821115611e1857600080fd5b611e248c838d01611cfa565b909850965060608b0135915080821115611e3d57600080fd5b611e498c838d01611cb1565b909650945060808b0135915080821115611e6257600080fd5b50611e6f8b828c01611cb1565b999c989b5096995094979396929594505050565b600060208284031215611e9557600080fd5b8135611ea081611dab565b9392505050565b600080600080600060608688031215611ebf57600080fd5b85359450602086013567ffffffffffffffff80821115611ede57600080fd5b611eea89838a01611cb1565b90965094506040880135915080821115611f0357600080fd5b50611f1088828901611cfa565b969995985093965092949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715611f7457611f74611f21565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611fa357611fa3611f21565b604052919050565b67ffffffffffffffff81168114610db157600080fd5b803561124181611fab565b600067ffffffffffffffff821115611fe657611fe6611f21565b50601f01601f191660200190565b600082601f83011261200557600080fd5b813561201861201382611fcc565b611f7a565b81815284602083860101111561202d57600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561205c57600080fd5b813567ffffffffffffffff8082111561207457600080fd5b90830190610100828603121561208957600080fd5b612091611f50565b82358152602083013560208201526040830135604082015260608301356060820152608083013560808201526120c960a08401611fc1565b60a082015260c0830135828111156120e057600080fd5b6120ec87828601611ff4565b60c08301525060e08301358281111561210457600080fd5b61211087828601611ff4565b60e08301525095945050505050565b60005b8381101561213a578181015183820152602001612122565b50506000910152565b6000815180845261215b81602086016020860161211f565b601f01601f19169290920160200192915050565b602081526000611ea06020830184612143565b600080600080600080600080600060c08a8c0312156121a057600080fd5b8935985060208a01356121b281611dab565b975060408a01356121c281611dab565b965060608a013567ffffffffffffffff808211156121df57600080fd5b6121eb8d838e01611cfa565b909850965060808c013591508082111561220457600080fd5b6122108d838e01611cb1565b909650945060a08c013591508082111561222957600080fd5b506122368c828d01611cb1565b915080935050809150509295985092959850929598565b600080600080600080600080600060c08a8c03121561226b57600080fd5b8935985060208a0135975060408a01356121c281611dab565b60006020828403121561229657600080fd5b813567ffffffffffffffff8111156122ad57600080fd5b610f7384828501611ff4565b600080604083850312156122cc57600080fd5b50508035926020909101359150565b600080602083850312156122ee57600080fd5b823567ffffffffffffffff81111561230557600080fd5b61231185828601611cb1565b90969095509350505050565b8183526000602080850194508260005b8581101561236857813561234081611dab565b73ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010161232d565b509495945050505050565b83815260406020820152600061238d60408301848661231d565b95945050505050565b6000602082840312156123a857600080fd5b5051919050565b6020815281516020820152602082015160408201526040820151606082015260608201516080820152608082015160a082015267ffffffffffffffff60a08301511660c0820152600060c08301516101008060e0850152612414610120850183612143565b915060e0850151601f1985840301828601526124308382612143565b9695505050505050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b84815283602082015260606040820152600061243060608301848661243a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610a8d57610a8d612485565b85815284602082015283604082015267ffffffffffffffff8316606082015260a0608082015260006124fc60a0830184612143565b979650505050505050565b805161124181611fab565b6000806040838503121561252557600080fd5b82519150602083015161253781611fab565b809150509250929050565b600082601f83011261255357600080fd5b815161256161201382611fcc565b81815284602083860101111561257657600080fd5b610f7382602083016020870161211f565b60006020828403121561259957600080fd5b815167ffffffffffffffff808211156125b157600080fd5b9083019061010082860312156125c657600080fd5b6125ce611f50565b825181526020830151602082015260408301516040820152606083015160608201526080830151608082015261260660a08401612507565b60a082015260c08301518281111561261d57600080fd5b61262987828601612542565b60c08301525060e08301518281111561264157600080fd5b61211087828601612542565b60006040828403121561265f57600080fd5b6040516040810181811067ffffffffffffffff8211171561268257612682611f21565b604052825181526020928301519281019290925250919050565b6000604082840312156126ae57600080fd5b611ea0838361264d565b600080604083850312156126cb57600080fd5b825167ffffffffffffffff808211156126e357600080fd5b6126ef86838701612542565b935060209150818501518181111561270657600080fd5b8501601f8101871361271757600080fd5b80518281111561272957612729611f21565b8060051b925061273a848401611f7a565b818152928201840192848101908985111561275457600080fd5b928501925b8484101561277e578351925061276e83611dab565b8282529285019290850190612759565b8096505050505050509250929050565b84815283602082015260606040820152600061243060608301848661231d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b86815285602082015284604082015283606082015260a0608082015260006107e860a08301848661243a565b60006020828403121561281b57600080fd5b8151611ea081611dab565b87815286602082015285604082015284606082015283608082015260e060a0820152600061285760e0830185612143565b82810360c08401526107e48185612143565b6000806040838503121561287c57600080fd5b825160ff8116811461288d57600080fd5b602084015190925067ffffffffffffffff8111156128aa57600080fd5b6128b685828601612542565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff85168152835160208201526020840151604082015267ffffffffffffffff604085015116606082015260608401516080820152608084015160a082015260e060c08201528160e082015260006101007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84111561298057600080fd5b8360051b8086838601379290920190910195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036129cb576129cb612485565b506001019056fea2646970667358221220932249f487c52b866750d955e59505543c8a7c08da55e4397f2122de7aa0104264736f6c63430008140033",
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
func DeployInterchainClientV1(auth *bind.TransactOpts, backend bind.ContractBackend, interchainDB common.Address, owner_ common.Address) (common.Address, *types.Transaction, *InterchainClientV1, error) {
	parsed, err := InterchainClientV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainClientV1Bin), backend, interchainDB, owner_)
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

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_InterchainClientV1 *InterchainClientV1Caller) INTERCHAINDB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "INTERCHAIN_DB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_InterchainClientV1 *InterchainClientV1Session) INTERCHAINDB() (common.Address, error) {
	return _InterchainClientV1.Contract.INTERCHAINDB(&_InterchainClientV1.CallOpts)
}

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_InterchainClientV1 *InterchainClientV1CallerSession) INTERCHAINDB() (common.Address, error) {
	return _InterchainClientV1.Contract.INTERCHAINDB(&_InterchainClientV1.CallOpts)
}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes encodedOptions) pure returns((uint256,uint256))
func (_InterchainClientV1 *InterchainClientV1Caller) DecodeOptions(opts *bind.CallOpts, encodedOptions []byte) (OptionsV1, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "decodeOptions", encodedOptions)

	if err != nil {
		return *new(OptionsV1), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionsV1)).(*OptionsV1)

	return out0, err

}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes encodedOptions) pure returns((uint256,uint256))
func (_InterchainClientV1 *InterchainClientV1Session) DecodeOptions(encodedOptions []byte) (OptionsV1, error) {
	return _InterchainClientV1.Contract.DecodeOptions(&_InterchainClientV1.CallOpts, encodedOptions)
}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes encodedOptions) pure returns((uint256,uint256))
func (_InterchainClientV1 *InterchainClientV1CallerSession) DecodeOptions(encodedOptions []byte) (OptionsV1, error) {
	return _InterchainClientV1.Contract.DecodeOptions(&_InterchainClientV1.CallOpts, encodedOptions)
}

// EncodeTransaction is a free data retrieval call binding the contract method 0x7c80a90f.
//
// Solidity: function encodeTransaction((uint256,bytes32,uint256,bytes32,uint256,uint64,bytes,bytes) icTx) pure returns(bytes)
func (_InterchainClientV1 *InterchainClientV1Caller) EncodeTransaction(opts *bind.CallOpts, icTx InterchainTransaction) ([]byte, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "encodeTransaction", icTx)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransaction is a free data retrieval call binding the contract method 0x7c80a90f.
//
// Solidity: function encodeTransaction((uint256,bytes32,uint256,bytes32,uint256,uint64,bytes,bytes) icTx) pure returns(bytes)
func (_InterchainClientV1 *InterchainClientV1Session) EncodeTransaction(icTx InterchainTransaction) ([]byte, error) {
	return _InterchainClientV1.Contract.EncodeTransaction(&_InterchainClientV1.CallOpts, icTx)
}

// EncodeTransaction is a free data retrieval call binding the contract method 0x7c80a90f.
//
// Solidity: function encodeTransaction((uint256,bytes32,uint256,bytes32,uint256,uint64,bytes,bytes) icTx) pure returns(bytes)
func (_InterchainClientV1 *InterchainClientV1CallerSession) EncodeTransaction(icTx InterchainTransaction) ([]byte, error) {
	return _InterchainClientV1.Contract.EncodeTransaction(&_InterchainClientV1.CallOpts, icTx)
}

// ExecutionFees is a free data retrieval call binding the contract method 0x7341eaf9.
//
// Solidity: function executionFees() view returns(address)
func (_InterchainClientV1 *InterchainClientV1Caller) ExecutionFees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "executionFees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionFees is a free data retrieval call binding the contract method 0x7341eaf9.
//
// Solidity: function executionFees() view returns(address)
func (_InterchainClientV1 *InterchainClientV1Session) ExecutionFees() (common.Address, error) {
	return _InterchainClientV1.Contract.ExecutionFees(&_InterchainClientV1.CallOpts)
}

// ExecutionFees is a free data retrieval call binding the contract method 0x7341eaf9.
//
// Solidity: function executionFees() view returns(address)
func (_InterchainClientV1 *InterchainClientV1CallerSession) ExecutionFees() (common.Address, error) {
	return _InterchainClientV1.Contract.ExecutionFees(&_InterchainClientV1.CallOpts)
}

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes encodedTx) view returns(address)
func (_InterchainClientV1 *InterchainClientV1Caller) GetExecutor(opts *bind.CallOpts, encodedTx []byte) (common.Address, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "getExecutor", encodedTx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes encodedTx) view returns(address)
func (_InterchainClientV1 *InterchainClientV1Session) GetExecutor(encodedTx []byte) (common.Address, error) {
	return _InterchainClientV1.Contract.GetExecutor(&_InterchainClientV1.CallOpts, encodedTx)
}

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes encodedTx) view returns(address)
func (_InterchainClientV1 *InterchainClientV1CallerSession) GetExecutor(encodedTx []byte) (common.Address, error) {
	return _InterchainClientV1.Contract.GetExecutor(&_InterchainClientV1.CallOpts, encodedTx)
}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_InterchainClientV1 *InterchainClientV1Caller) GetExecutorById(opts *bind.CallOpts, transactionId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "getExecutorById", transactionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_InterchainClientV1 *InterchainClientV1Session) GetExecutorById(transactionId [32]byte) (common.Address, error) {
	return _InterchainClientV1.Contract.GetExecutorById(&_InterchainClientV1.CallOpts, transactionId)
}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_InterchainClientV1 *InterchainClientV1CallerSession) GetExecutorById(transactionId [32]byte) (common.Address, error) {
	return _InterchainClientV1.Contract.GetExecutorById(&_InterchainClientV1.CallOpts, transactionId)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0x3c383e7b.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, bytes message) view returns(uint256 fee)
func (_InterchainClientV1 *InterchainClientV1Caller) GetInterchainFee(opts *bind.CallOpts, dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*big.Int, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcExecutionService, srcModules, options, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0x3c383e7b.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, bytes message) view returns(uint256 fee)
func (_InterchainClientV1 *InterchainClientV1Session) GetInterchainFee(dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*big.Int, error) {
	return _InterchainClientV1.Contract.GetInterchainFee(&_InterchainClientV1.CallOpts, dstChainId, srcExecutionService, srcModules, options, message)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0x3c383e7b.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, bytes message) view returns(uint256 fee)
func (_InterchainClientV1 *InterchainClientV1CallerSession) GetInterchainFee(dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*big.Int, error) {
	return _InterchainClientV1.Contract.GetInterchainFee(&_InterchainClientV1.CallOpts, dstChainId, srcExecutionService, srcModules, options, message)
}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_InterchainClientV1 *InterchainClientV1Caller) GetLinkedClient(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "getLinkedClient", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_InterchainClientV1 *InterchainClientV1Session) GetLinkedClient(chainId *big.Int) ([32]byte, error) {
	return _InterchainClientV1.Contract.GetLinkedClient(&_InterchainClientV1.CallOpts, chainId)
}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_InterchainClientV1 *InterchainClientV1CallerSession) GetLinkedClient(chainId *big.Int) ([32]byte, error) {
	return _InterchainClientV1.Contract.GetLinkedClient(&_InterchainClientV1.CallOpts, chainId)
}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address linkedClientEVM)
func (_InterchainClientV1 *InterchainClientV1Caller) GetLinkedClientEVM(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "getLinkedClientEVM", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address linkedClientEVM)
func (_InterchainClientV1 *InterchainClientV1Session) GetLinkedClientEVM(chainId *big.Int) (common.Address, error) {
	return _InterchainClientV1.Contract.GetLinkedClientEVM(&_InterchainClientV1.CallOpts, chainId)
}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address linkedClientEVM)
func (_InterchainClientV1 *InterchainClientV1CallerSession) GetLinkedClientEVM(chainId *big.Int) (common.Address, error) {
	return _InterchainClientV1.Contract.GetLinkedClientEVM(&_InterchainClientV1.CallOpts, chainId)
}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes encodedTx, bytes32[] proof) view returns(bool)
func (_InterchainClientV1 *InterchainClientV1Caller) IsExecutable(opts *bind.CallOpts, encodedTx []byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "isExecutable", encodedTx, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes encodedTx, bytes32[] proof) view returns(bool)
func (_InterchainClientV1 *InterchainClientV1Session) IsExecutable(encodedTx []byte, proof [][32]byte) (bool, error) {
	return _InterchainClientV1.Contract.IsExecutable(&_InterchainClientV1.CallOpts, encodedTx, proof)
}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes encodedTx, bytes32[] proof) view returns(bool)
func (_InterchainClientV1 *InterchainClientV1CallerSession) IsExecutable(encodedTx []byte, proof [][32]byte) (bool, error) {
	return _InterchainClientV1.Contract.IsExecutable(&_InterchainClientV1.CallOpts, encodedTx, proof)
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

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) InterchainExecute(opts *bind.TransactOpts, gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "interchainExecute", gasLimit, transaction, proof)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_InterchainClientV1 *InterchainClientV1Session) InterchainExecute(gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainExecute(&_InterchainClientV1.TransactOpts, gasLimit, transaction, proof)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) InterchainExecute(gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainExecute(&_InterchainClientV1.TransactOpts, gasLimit, transaction, proof)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_InterchainClientV1 *InterchainClientV1Transactor) InterchainSend(opts *bind.TransactOpts, dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "interchainSend", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_InterchainClientV1 *InterchainClientV1Session) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSend(&_InterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_InterchainClientV1 *InterchainClientV1TransactorSession) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSend(&_InterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_InterchainClientV1 *InterchainClientV1Transactor) InterchainSendEVM(opts *bind.TransactOpts, dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "interchainSendEVM", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_InterchainClientV1 *InterchainClientV1Session) InterchainSendEVM(dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSendEVM(&_InterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_InterchainClientV1 *InterchainClientV1TransactorSession) InterchainSendEVM(dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSendEVM(&_InterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
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

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) SetExecutionFees(opts *bind.TransactOpts, executionFees_ common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "setExecutionFees", executionFees_)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_InterchainClientV1 *InterchainClientV1Session) SetExecutionFees(executionFees_ common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.SetExecutionFees(&_InterchainClientV1.TransactOpts, executionFees_)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) SetExecutionFees(executionFees_ common.Address) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.SetExecutionFees(&_InterchainClientV1.TransactOpts, executionFees_)
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

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainClientV1 *InterchainClientV1Transactor) WriteExecutionProof(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "writeExecutionProof", transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainClientV1 *InterchainClientV1Session) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.WriteExecutionProof(&_InterchainClientV1.TransactOpts, transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainClientV1 *InterchainClientV1TransactorSession) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.WriteExecutionProof(&_InterchainClientV1.TransactOpts, transactionId)
}

// InterchainClientV1ExecutionFeesSetIterator is returned from FilterExecutionFeesSet and is used to iterate over the raw logs and unpacked data for ExecutionFeesSet events raised by the InterchainClientV1 contract.
type InterchainClientV1ExecutionFeesSetIterator struct {
	Event *InterchainClientV1ExecutionFeesSet // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1ExecutionFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1ExecutionFeesSet)
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
		it.Event = new(InterchainClientV1ExecutionFeesSet)
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
func (it *InterchainClientV1ExecutionFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1ExecutionFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1ExecutionFeesSet represents a ExecutionFeesSet event raised by the InterchainClientV1 contract.
type InterchainClientV1ExecutionFeesSet struct {
	ExecutionFees common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFeesSet is a free log retrieval operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterExecutionFeesSet(opts *bind.FilterOpts) (*InterchainClientV1ExecutionFeesSetIterator, error) {

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "ExecutionFeesSet")
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1ExecutionFeesSetIterator{contract: _InterchainClientV1.contract, event: "ExecutionFeesSet", logs: logs, sub: sub}, nil
}

// WatchExecutionFeesSet is a free log subscription operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchExecutionFeesSet(opts *bind.WatchOpts, sink chan<- *InterchainClientV1ExecutionFeesSet) (event.Subscription, error) {

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "ExecutionFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1ExecutionFeesSet)
				if err := _InterchainClientV1.contract.UnpackLog(event, "ExecutionFeesSet", log); err != nil {
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

// ParseExecutionFeesSet is a log parse operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1 *InterchainClientV1Filterer) ParseExecutionFeesSet(log types.Log) (*InterchainClientV1ExecutionFeesSet, error) {
	event := new(InterchainClientV1ExecutionFeesSet)
	if err := _InterchainClientV1.contract.UnpackLog(event, "ExecutionFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1ExecutionProofWrittenIterator is returned from FilterExecutionProofWritten and is used to iterate over the raw logs and unpacked data for ExecutionProofWritten events raised by the InterchainClientV1 contract.
type InterchainClientV1ExecutionProofWrittenIterator struct {
	Event *InterchainClientV1ExecutionProofWritten // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1ExecutionProofWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1ExecutionProofWritten)
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
		it.Event = new(InterchainClientV1ExecutionProofWritten)
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
func (it *InterchainClientV1ExecutionProofWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1ExecutionProofWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1ExecutionProofWritten represents a ExecutionProofWritten event raised by the InterchainClientV1 contract.
type InterchainClientV1ExecutionProofWritten struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
	Executor      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionProofWritten is a free log retrieval operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterExecutionProofWritten(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1ExecutionProofWrittenIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1ExecutionProofWrittenIterator{contract: _InterchainClientV1.contract, event: "ExecutionProofWritten", logs: logs, sub: sub}, nil
}

// WatchExecutionProofWritten is a free log subscription operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchExecutionProofWritten(opts *bind.WatchOpts, sink chan<- *InterchainClientV1ExecutionProofWritten, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1ExecutionProofWritten)
				if err := _InterchainClientV1.contract.UnpackLog(event, "ExecutionProofWritten", log); err != nil {
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

// ParseExecutionProofWritten is a log parse operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1 *InterchainClientV1Filterer) ParseExecutionProofWritten(log types.Log) (*InterchainClientV1ExecutionProofWritten, error) {
	event := new(InterchainClientV1ExecutionProofWritten)
	if err := _InterchainClientV1.contract.UnpackLog(event, "ExecutionProofWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1InterchainTransactionReceivedIterator is returned from FilterInterchainTransactionReceived and is used to iterate over the raw logs and unpacked data for InterchainTransactionReceived events raised by the InterchainClientV1 contract.
type InterchainClientV1InterchainTransactionReceivedIterator struct {
	Event *InterchainClientV1InterchainTransactionReceived // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1InterchainTransactionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1InterchainTransactionReceived)
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
		it.Event = new(InterchainClientV1InterchainTransactionReceived)
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
func (it *InterchainClientV1InterchainTransactionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1InterchainTransactionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1InterchainTransactionReceived represents a InterchainTransactionReceived event raised by the InterchainClientV1 contract.
type InterchainClientV1InterchainTransactionReceived struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
	SrcChainId    *big.Int
	SrcSender     [32]byte
	DstReceiver   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionReceived is a free log retrieval operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterInterchainTransactionReceived(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1InterchainTransactionReceivedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1InterchainTransactionReceivedIterator{contract: _InterchainClientV1.contract, event: "InterchainTransactionReceived", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionReceived is a free log subscription operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchInterchainTransactionReceived(opts *bind.WatchOpts, sink chan<- *InterchainClientV1InterchainTransactionReceived, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1InterchainTransactionReceived)
				if err := _InterchainClientV1.contract.UnpackLog(event, "InterchainTransactionReceived", log); err != nil {
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

// ParseInterchainTransactionReceived is a log parse operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1 *InterchainClientV1Filterer) ParseInterchainTransactionReceived(log types.Log) (*InterchainClientV1InterchainTransactionReceived, error) {
	event := new(InterchainClientV1InterchainTransactionReceived)
	if err := _InterchainClientV1.contract.UnpackLog(event, "InterchainTransactionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	TransactionId   [32]byte
	DbNonce         *big.Int
	EntryIndex      uint64
	DstChainId      *big.Int
	SrcSender       [32]byte
	DstReceiver     [32]byte
	VerificationFee *big.Int
	ExecutionFee    *big.Int
	Options         []byte
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionSent is a free log retrieval operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterInterchainTransactionSent(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1InterchainTransactionSentIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1InterchainTransactionSentIterator{contract: _InterchainClientV1.contract, event: "InterchainTransactionSent", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionSent is a free log subscription operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchInterchainTransactionSent(opts *bind.WatchOpts, sink chan<- *InterchainClientV1InterchainTransactionSent, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule, entryIndexRule)
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

// ParseInterchainTransactionSent is a log parse operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1 *InterchainClientV1Filterer) ParseInterchainTransactionSent(log types.Log) (*InterchainClientV1InterchainTransactionSent, error) {
	event := new(InterchainClientV1InterchainTransactionSent)
	if err := _InterchainClientV1.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1LinkedClientSetIterator is returned from FilterLinkedClientSet and is used to iterate over the raw logs and unpacked data for LinkedClientSet events raised by the InterchainClientV1 contract.
type InterchainClientV1LinkedClientSetIterator struct {
	Event *InterchainClientV1LinkedClientSet // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1LinkedClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1LinkedClientSet)
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
		it.Event = new(InterchainClientV1LinkedClientSet)
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
func (it *InterchainClientV1LinkedClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1LinkedClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1LinkedClientSet represents a LinkedClientSet event raised by the InterchainClientV1 contract.
type InterchainClientV1LinkedClientSet struct {
	ChainId *big.Int
	Client  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLinkedClientSet is a free log retrieval operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterLinkedClientSet(opts *bind.FilterOpts) (*InterchainClientV1LinkedClientSetIterator, error) {

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "LinkedClientSet")
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1LinkedClientSetIterator{contract: _InterchainClientV1.contract, event: "LinkedClientSet", logs: logs, sub: sub}, nil
}

// WatchLinkedClientSet is a free log subscription operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchLinkedClientSet(opts *bind.WatchOpts, sink chan<- *InterchainClientV1LinkedClientSet) (event.Subscription, error) {

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "LinkedClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1LinkedClientSet)
				if err := _InterchainClientV1.contract.UnpackLog(event, "LinkedClientSet", log); err != nil {
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

// ParseLinkedClientSet is a log parse operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1 *InterchainClientV1Filterer) ParseLinkedClientSet(log types.Log) (*InterchainClientV1LinkedClientSet, error) {
	event := new(InterchainClientV1LinkedClientSet)
	if err := _InterchainClientV1.contract.UnpackLog(event, "LinkedClientSet", log); err != nil {
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

// InterchainClientV1EventsMetaData contains all meta data concerning the InterchainClientV1Events contract.
var InterchainClientV1EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionFees\",\"type\":\"address\"}],\"name\":\"ExecutionFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionProofWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"}],\"name\":\"InterchainTransactionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"InterchainTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"LinkedClientSet\",\"type\":\"event\"}]",
}

// InterchainClientV1EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainClientV1EventsMetaData.ABI instead.
var InterchainClientV1EventsABI = InterchainClientV1EventsMetaData.ABI

// InterchainClientV1Events is an auto generated Go binding around an Ethereum contract.
type InterchainClientV1Events struct {
	InterchainClientV1EventsCaller     // Read-only binding to the contract
	InterchainClientV1EventsTransactor // Write-only binding to the contract
	InterchainClientV1EventsFilterer   // Log filterer for contract events
}

// InterchainClientV1EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainClientV1EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainClientV1EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainClientV1EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainClientV1EventsSession struct {
	Contract     *InterchainClientV1Events // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InterchainClientV1EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainClientV1EventsCallerSession struct {
	Contract *InterchainClientV1EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// InterchainClientV1EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainClientV1EventsTransactorSession struct {
	Contract     *InterchainClientV1EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// InterchainClientV1EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainClientV1EventsRaw struct {
	Contract *InterchainClientV1Events // Generic contract binding to access the raw methods on
}

// InterchainClientV1EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainClientV1EventsCallerRaw struct {
	Contract *InterchainClientV1EventsCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainClientV1EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainClientV1EventsTransactorRaw struct {
	Contract *InterchainClientV1EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainClientV1Events creates a new instance of InterchainClientV1Events, bound to a specific deployed contract.
func NewInterchainClientV1Events(address common.Address, backend bind.ContractBackend) (*InterchainClientV1Events, error) {
	contract, err := bindInterchainClientV1Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1Events{InterchainClientV1EventsCaller: InterchainClientV1EventsCaller{contract: contract}, InterchainClientV1EventsTransactor: InterchainClientV1EventsTransactor{contract: contract}, InterchainClientV1EventsFilterer: InterchainClientV1EventsFilterer{contract: contract}}, nil
}

// NewInterchainClientV1EventsCaller creates a new read-only instance of InterchainClientV1Events, bound to a specific deployed contract.
func NewInterchainClientV1EventsCaller(address common.Address, caller bind.ContractCaller) (*InterchainClientV1EventsCaller, error) {
	contract, err := bindInterchainClientV1Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsCaller{contract: contract}, nil
}

// NewInterchainClientV1EventsTransactor creates a new write-only instance of InterchainClientV1Events, bound to a specific deployed contract.
func NewInterchainClientV1EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainClientV1EventsTransactor, error) {
	contract, err := bindInterchainClientV1Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsTransactor{contract: contract}, nil
}

// NewInterchainClientV1EventsFilterer creates a new log filterer instance of InterchainClientV1Events, bound to a specific deployed contract.
func NewInterchainClientV1EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainClientV1EventsFilterer, error) {
	contract, err := bindInterchainClientV1Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsFilterer{contract: contract}, nil
}

// bindInterchainClientV1Events binds a generic wrapper to an already deployed contract.
func bindInterchainClientV1Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainClientV1EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainClientV1Events *InterchainClientV1EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainClientV1Events.Contract.InterchainClientV1EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainClientV1Events *InterchainClientV1EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientV1Events.Contract.InterchainClientV1EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainClientV1Events *InterchainClientV1EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainClientV1Events.Contract.InterchainClientV1EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainClientV1Events *InterchainClientV1EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainClientV1Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainClientV1Events *InterchainClientV1EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientV1Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainClientV1Events *InterchainClientV1EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainClientV1Events.Contract.contract.Transact(opts, method, params...)
}

// InterchainClientV1EventsExecutionFeesSetIterator is returned from FilterExecutionFeesSet and is used to iterate over the raw logs and unpacked data for ExecutionFeesSet events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsExecutionFeesSetIterator struct {
	Event *InterchainClientV1EventsExecutionFeesSet // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsExecutionFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsExecutionFeesSet)
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
		it.Event = new(InterchainClientV1EventsExecutionFeesSet)
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
func (it *InterchainClientV1EventsExecutionFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsExecutionFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsExecutionFeesSet represents a ExecutionFeesSet event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsExecutionFeesSet struct {
	ExecutionFees common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFeesSet is a free log retrieval operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterExecutionFeesSet(opts *bind.FilterOpts) (*InterchainClientV1EventsExecutionFeesSetIterator, error) {

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "ExecutionFeesSet")
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsExecutionFeesSetIterator{contract: _InterchainClientV1Events.contract, event: "ExecutionFeesSet", logs: logs, sub: sub}, nil
}

// WatchExecutionFeesSet is a free log subscription operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchExecutionFeesSet(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsExecutionFeesSet) (event.Subscription, error) {

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "ExecutionFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsExecutionFeesSet)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "ExecutionFeesSet", log); err != nil {
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

// ParseExecutionFeesSet is a log parse operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseExecutionFeesSet(log types.Log) (*InterchainClientV1EventsExecutionFeesSet, error) {
	event := new(InterchainClientV1EventsExecutionFeesSet)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "ExecutionFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsExecutionProofWrittenIterator is returned from FilterExecutionProofWritten and is used to iterate over the raw logs and unpacked data for ExecutionProofWritten events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsExecutionProofWrittenIterator struct {
	Event *InterchainClientV1EventsExecutionProofWritten // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsExecutionProofWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsExecutionProofWritten)
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
		it.Event = new(InterchainClientV1EventsExecutionProofWritten)
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
func (it *InterchainClientV1EventsExecutionProofWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsExecutionProofWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsExecutionProofWritten represents a ExecutionProofWritten event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsExecutionProofWritten struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
	Executor      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionProofWritten is a free log retrieval operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterExecutionProofWritten(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1EventsExecutionProofWrittenIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsExecutionProofWrittenIterator{contract: _InterchainClientV1Events.contract, event: "ExecutionProofWritten", logs: logs, sub: sub}, nil
}

// WatchExecutionProofWritten is a free log subscription operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchExecutionProofWritten(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsExecutionProofWritten, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsExecutionProofWritten)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "ExecutionProofWritten", log); err != nil {
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

// ParseExecutionProofWritten is a log parse operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseExecutionProofWritten(log types.Log) (*InterchainClientV1EventsExecutionProofWritten, error) {
	event := new(InterchainClientV1EventsExecutionProofWritten)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "ExecutionProofWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsInterchainTransactionReceivedIterator is returned from FilterInterchainTransactionReceived and is used to iterate over the raw logs and unpacked data for InterchainTransactionReceived events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsInterchainTransactionReceivedIterator struct {
	Event *InterchainClientV1EventsInterchainTransactionReceived // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsInterchainTransactionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsInterchainTransactionReceived)
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
		it.Event = new(InterchainClientV1EventsInterchainTransactionReceived)
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
func (it *InterchainClientV1EventsInterchainTransactionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsInterchainTransactionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsInterchainTransactionReceived represents a InterchainTransactionReceived event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsInterchainTransactionReceived struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
	SrcChainId    *big.Int
	SrcSender     [32]byte
	DstReceiver   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionReceived is a free log retrieval operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterInterchainTransactionReceived(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1EventsInterchainTransactionReceivedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsInterchainTransactionReceivedIterator{contract: _InterchainClientV1Events.contract, event: "InterchainTransactionReceived", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionReceived is a free log subscription operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchInterchainTransactionReceived(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsInterchainTransactionReceived, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsInterchainTransactionReceived)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionReceived", log); err != nil {
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

// ParseInterchainTransactionReceived is a log parse operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseInterchainTransactionReceived(log types.Log) (*InterchainClientV1EventsInterchainTransactionReceived, error) {
	event := new(InterchainClientV1EventsInterchainTransactionReceived)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsInterchainTransactionSentIterator is returned from FilterInterchainTransactionSent and is used to iterate over the raw logs and unpacked data for InterchainTransactionSent events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsInterchainTransactionSentIterator struct {
	Event *InterchainClientV1EventsInterchainTransactionSent // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsInterchainTransactionSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsInterchainTransactionSent)
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
		it.Event = new(InterchainClientV1EventsInterchainTransactionSent)
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
func (it *InterchainClientV1EventsInterchainTransactionSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsInterchainTransactionSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsInterchainTransactionSent represents a InterchainTransactionSent event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsInterchainTransactionSent struct {
	TransactionId   [32]byte
	DbNonce         *big.Int
	EntryIndex      uint64
	DstChainId      *big.Int
	SrcSender       [32]byte
	DstReceiver     [32]byte
	VerificationFee *big.Int
	ExecutionFee    *big.Int
	Options         []byte
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionSent is a free log retrieval operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterInterchainTransactionSent(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1EventsInterchainTransactionSentIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsInterchainTransactionSentIterator{contract: _InterchainClientV1Events.contract, event: "InterchainTransactionSent", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionSent is a free log subscription operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchInterchainTransactionSent(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsInterchainTransactionSent, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsInterchainTransactionSent)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
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

// ParseInterchainTransactionSent is a log parse operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseInterchainTransactionSent(log types.Log) (*InterchainClientV1EventsInterchainTransactionSent, error) {
	event := new(InterchainClientV1EventsInterchainTransactionSent)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsLinkedClientSetIterator is returned from FilterLinkedClientSet and is used to iterate over the raw logs and unpacked data for LinkedClientSet events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsLinkedClientSetIterator struct {
	Event *InterchainClientV1EventsLinkedClientSet // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsLinkedClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsLinkedClientSet)
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
		it.Event = new(InterchainClientV1EventsLinkedClientSet)
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
func (it *InterchainClientV1EventsLinkedClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsLinkedClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsLinkedClientSet represents a LinkedClientSet event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsLinkedClientSet struct {
	ChainId *big.Int
	Client  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLinkedClientSet is a free log retrieval operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterLinkedClientSet(opts *bind.FilterOpts) (*InterchainClientV1EventsLinkedClientSetIterator, error) {

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "LinkedClientSet")
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsLinkedClientSetIterator{contract: _InterchainClientV1Events.contract, event: "LinkedClientSet", logs: logs, sub: sub}, nil
}

// WatchLinkedClientSet is a free log subscription operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchLinkedClientSet(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsLinkedClientSet) (event.Subscription, error) {

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "LinkedClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsLinkedClientSet)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "LinkedClientSet", log); err != nil {
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

// ParseLinkedClientSet is a log parse operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseLinkedClientSet(log types.Log) (*InterchainClientV1EventsLinkedClientSet, error) {
	event := new(InterchainClientV1EventsLinkedClientSet)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "LinkedClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122042e9d0d617dda08d2003e847d0ffa768d3dcb1b5e8d5609184d6f5f994e490d964736f6c63430008140033",
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

// InterchainTransactionLibMetaData contains all meta data concerning the InterchainTransactionLib contract.
var InterchainTransactionLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201c6e7a0fb4b59dcaead35723454bb449764d05c94fd03b6bcc0ec5c473da95f764736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209b12c00facb35eb14d968a1595ed4c83fc8c9d9d1d14843863dbd8a6e142e3d164736f6c63430008140033",
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c7bac564ffbdae83721b5316f5ae26a9e1a3c94aec0098ab27e9821e8df2427e64736f6c63430008140033",
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
