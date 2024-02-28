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
	SrcChainId *big.Int
	DbNonce    *big.Int
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
	Options     []byte
	Message     []byte
}

// OptionsV1 is an auto generated low-level Go binding around an user-defined struct.
type OptionsV1 struct {
	GasLimit   *big.Int
	GasAirdrop *big.Int
}

// AppConfigLibMetaData contains all meta data concerning the AppConfigLib contract.
var AppConfigLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"AppConfigLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122091c24ac5f930697e7f1a115a22d165a12584d9783f66383540206abd7a95ed4b64736f6c63430008140033",
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
	ABI: "[{\"inputs\":[],\"name\":\"ExecutionFees__AlreadyRecorded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFees__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFees__ZeroAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"accumulatedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accumulated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"addExecutionFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"claimExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"executionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"recordExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"recordedExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"unclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unclaimed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b399470d": "appReceive(uint256,bytes32,uint256,bytes)",
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

// AppReceive is a paid mutator transaction binding the contract method 0xb399470d.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0xb399470d.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts, srcChainId, sender, dbNonce, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0xb399470d.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts, srcChainId, sender, dbNonce, message)
}

// IInterchainClientV1MetaData contains all meta data concerning the IInterchainClientV1 contract.
var IInterchainClientV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectDstChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectSrcChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotEnoughResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxNotExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"getExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"getExecutorById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSendEVM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"isExecutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionFees_\",\"type\":\"address\"}],\"name\":\"setExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"writeExecutionProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f92a79ff": "getExecutor(bytes)",
		"f1a61fac": "getExecutorById(bytes32)",
		"3c383e7b": "getInterchainFee(uint256,address,address[],bytes,bytes)",
		"80efe777": "interchainExecute(uint256,bytes)",
		"98939d28": "interchainSend(uint256,bytes32,address,address[],bytes,bytes)",
		"827f940d": "interchainSendEVM(uint256,address,address,address[],bytes,bytes)",
		"31afa7de": "isExecutable(bytes)",
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

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes transaction) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1Caller) IsExecutable(opts *bind.CallOpts, transaction []byte) (bool, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "isExecutable", transaction)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes transaction) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1Session) IsExecutable(transaction []byte) (bool, error) {
	return _IInterchainClientV1.Contract.IsExecutable(&_IInterchainClientV1.CallOpts, transaction)
}

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes transaction) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) IsExecutable(transaction []byte) (bool, error) {
	return _IInterchainClientV1.Contract.IsExecutable(&_IInterchainClientV1.CallOpts, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainExecute(opts *bind.TransactOpts, gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainExecute", gasLimit, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainExecute(gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, gasLimit, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainExecute(gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, gasLimit, transaction)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainSend(opts *bind.TransactOpts, dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainSend", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainSendEVM(opts *bind.TransactOpts, dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainSendEVM", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainSendEVM(dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSendEVM(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
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
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) WriteExecutionProof(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "writeExecutionProof", transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce)
func (_IInterchainClientV1 *IInterchainClientV1Session) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.WriteExecutionProof(&_IInterchainClientV1.TransactOpts, transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce)
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.WriteExecutionProof(&_IInterchainClientV1.TransactOpts, transactionId)
}

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"existingEntryValue\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"newEntry\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingEntries\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__EntryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"readEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f338140e": "getDBNonce()",
		"bae78d7b": "getEntry(uint256)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"a9c9cff1": "readEntry(address,(uint256,uint256,bytes32,bytes32))",
		"81ab5b5a": "requestVerification(uint256,uint256,address[])",
		"54941dfa": "verifyEntry((uint256,uint256,bytes32,bytes32))",
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

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetEntry(opts *bind.CallOpts, dbNonce *big.Int) (InterchainEntry, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntry", dbNonce)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetEntry(dbNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, dbNonce)
}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetEntry(dbNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, dbNonce)
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

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) ReadEntry(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "readEntry", dstModule, entry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestVerification", destChainId, dbNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestVerification(destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, dbNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestVerification(destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, dbNonce, srcModules)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "verifyEntry", entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactorSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// InterchainClientV1MetaData contains all meta data concerning the InterchainClientV1 contract.
var InterchainClientV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"interchainDB\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"AppConfigLib__IncorrectVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectDstChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectSrcChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotEnoughResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxNotExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionProofWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"}],\"name\":\"InterchainTransactionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"InterchainTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INTERCHAIN_DB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedOptions\",\"type\":\"bytes\"}],\"name\":\"decodeOptions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAirdrop\",\"type\":\"uint256\"}],\"internalType\":\"structOptionsV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structInterchainTransaction\",\"name\":\"icTx\",\"type\":\"tuple\"}],\"name\":\"encodeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionFees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"getExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"getExecutorById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSend\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSendEVM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedTx\",\"type\":\"bytes\"}],\"name\":\"isExecutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"linkedClients\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteClient\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionFees_\",\"type\":\"address\"}],\"name\":\"setExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"writeExecutionProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e4c61247": "INTERCHAIN_DB()",
		"d5e788a0": "decodeOptions(bytes)",
		"2743f804": "encodeTransaction((uint256,bytes32,uint256,bytes32,uint256,bytes,bytes))",
		"7341eaf9": "executionFees()",
		"f92a79ff": "getExecutor(bytes)",
		"f1a61fac": "getExecutorById(bytes32)",
		"3c383e7b": "getInterchainFee(uint256,address,address[],bytes,bytes)",
		"80efe777": "interchainExecute(uint256,bytes)",
		"98939d28": "interchainSend(uint256,bytes32,address,address[],bytes,bytes)",
		"827f940d": "interchainSendEVM(uint256,address,address,address[],bytes,bytes)",
		"31afa7de": "isExecutable(bytes)",
		"7268b08f": "linkedClients(uint256)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"3dc68b87": "setExecutionFees(address)",
		"f34234c8": "setLinkedClient(uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
		"90e81077": "writeExecutionProof(bytes32)",
	},
	Bin: "0x60a06040523480156200001157600080fd5b50604051620023e9380380620023e98339810160408190526200003491620000c9565b33806200005b57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b620000668162000079565b506001600160a01b0316608052620000fb565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215620000dc57600080fd5b81516001600160a01b0381168114620000f457600080fd5b9392505050565b6080516122a8620001416000396000818161035e015281816104cc015281816108ae01528181610f460152818161100e01528181611136015261153c01526122a86000f3fe6080604052600436106101295760003560e01c80638da5cb5b116100a5578063e4c6124711610074578063f2fde38b11610059578063f2fde38b146103c3578063f34234c8146103e3578063f92a79ff1461040357600080fd5b8063e4c612471461034c578063f1a61fac1461038057600080fd5b80638da5cb5b146102b357806390e81077146102de57806398939d28146102fe578063d5e788a01461031157600080fd5b8063715018a6116100fc5780637341eaf9116100e15780637341eaf91461022657806380efe77714610278578063827f940d1461028b57600080fd5b8063715018a6146101e45780637268b08f146101f957600080fd5b80632743f8041461012e57806331afa7de146101645780633c383e7b146101945780633dc68b87146101c2575b600080fd5b34801561013a57600080fd5b5061014e610149366004611853565b610423565b60405161015b9190611966565b60405180910390f35b34801561017057600080fd5b5061018461017f3660046119c9565b610434565b604051901515815260200161015b565b3480156101a057600080fd5b506101b46101af366004611a72565b61048c565b60405190815260200161015b565b3480156101ce57600080fd5b506101e26101dd366004611b28565b610553565b005b3480156101f057600080fd5b506101e26105a2565b34801561020557600080fd5b506101b4610214366004611b45565b60026020526000908152604090205481565b34801561023257600080fd5b506001546102539073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015b565b6101e2610286366004611b5e565b6105b6565b61029e610299366004611baa565b610799565b6040805192835260208301919091520161015b565b3480156102bf57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610253565b3480156102ea57600080fd5b506101b46102f9366004611b45565b6107d7565b61029e61030c366004611c75565b610987565b34801561031d57600080fd5b5061033161032c366004611cac565b6109ad565b6040805182518152602092830151928101929092520161015b565b34801561035857600080fd5b506102537f000000000000000000000000000000000000000000000000000000000000000081565b34801561038c57600080fd5b5061025361039b366004611b45565b60009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b3480156103cf57600080fd5b506101e26103de366004611b28565b6109ca565b3480156103ef57600080fd5b506101e26103fe366004611ce1565b610a2e565b34801561040f57600080fd5b5061025361041e3660046119c9565b610a48565b606061042e82610ac8565b92915050565b60008061047684848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610af192505050565b905061048181610b4b565b506001949350505050565b6040517ffc7686ec00000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063fc7686ec90610505908c908b908b90600401611d59565b602060405180830381865afa158015610522573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105469190611d7c565b9998505050505050505050565b61055b610d7b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6105aa610d7b565b6105b46000610dce565b565b60006105f783838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610af192505050565b9050600061060482610b4b565b600081815260036020526040812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000163317905560a08401519192509061064c90610e43565b90508060200151341461069f5760208101516040517f2b36102500000000000000000000000000000000000000000000000000000000815234600482015260248101919091526044015b60405180910390fd5b80518610156106ad57805195505b606083015173ffffffffffffffffffffffffffffffffffffffff1663b399470d87348660000151876020015188608001518960c001516040518763ffffffff1660e01b81526004016107029493929190611d95565b6000604051808303818589803b15801561071b57600080fd5b5088f115801561072f573d6000803e3d6000fd5b50505060808601518651602088015160608901516040519396508895507fdb272c7b2f65161c0573b8b7c21bc3e25cf53bcd8c1524fb3fb78bac8c7df94e9450610789939283526020830191909152604082015260600190565b60405180910390a3505050505050565b60008073ffffffffffffffffffffffffffffffffffffffff8a166107c48c828c8c8c8c8c8c8c610ec7565b9250925050995099975050505050505050565b60008181526003602052604081205473ffffffffffffffffffffffffffffffffffffffff1680610836576040517fe99eb48d00000000000000000000000000000000000000000000000000000000815260048101849052602401610696565b6000838260405160200161086a92919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60408051808303601f1901815290829052805160208201207f2ad8c706000000000000000000000000000000000000000000000000000000008352600483015291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632ad8c706906024016020604051808303816000875af115801561090c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109309190611d7c565b60405173ffffffffffffffffffffffffffffffffffffffff84168152909350839085907ff9e8d86f4b1412ca26874e2b00b21627542627093b367c3dc95e612c4de927c19060200160405180910390a35050919050565b60008061099b8b8b8b8b8b8b8b8b8b610ec7565b91509150995099975050505050505050565b604080518082019091526000808252602082015261042e82610e43565b6109d2610d7b565b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260006004820152602401610696565b610a2b81610dce565b50565b610a36610d7b565b60009182526002602052604090912055565b600080610a8a84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610af192505050565b905060036000610a9983611493565b815260208101919091526040016000205473ffffffffffffffffffffffffffffffffffffffff16949350505050565b606081604051602001610adb9190611dc4565b6040516020818303038152906040529050919050565b610b376040518060e00160405280600081526020016000801916815260200160008152602001600080191681526020016000815260200160608152602001606081525090565b8180602001905181019061042e9190611e75565b600046826000015103610b905781516040517f9720ce050000000000000000000000000000000000000000000000000000000081526004810191909152602401610696565b46826040015114610bd55781604001516040517f9732538200000000000000000000000000000000000000000000000000000000815260040161069691815260200190565b610bde82611493565b60008181526003602052604090205490915073ffffffffffffffffffffffffffffffffffffffff1615610c40576040517fd80aeb9100000000000000000000000000000000000000000000000000000000815260048101829052602401610696565b6000604051806080016040528084600001518152602001846080015181526020016002600086600001518152602001908152602001600020548152602001838152509050600080610c92856060015190565b73ffffffffffffffffffffffffffffffffffffffff1663287bc0576040518163ffffffff1660e01b8152600401600060405180830381865afa158015610cdc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d049190810190611f29565b915091506000610d13836114c3565b90506000610d268386846020015161152b565b8251909150811015610d715781516040517f0bce4e85000000000000000000000000000000000000000000000000000000008152610696918391600401918252602082015260400190565b5050505050919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105b4576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610696565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040805180820190915260008082526020820152600080610e6384611689565b9092509050600160ff83161015610eab576040517fbd91a21500000000000000000000000000000000000000000000000000000000815260ff83166004820152602401610696565b80806020019051810190610ebf919061204e565b949350505050565b600080468b03610f06576040517f97325382000000000000000000000000000000000000000000000000000000008152600481018c9052602401610696565b6040517ffc7686ec00000000000000000000000000000000000000000000000000000000815260009073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063fc7686ec90610f7f908f908d908d90600401611d59565b602060405180830381865afa158015610f9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc09190611d7c565b905080341015611005576040517f9557ee4900000000000000000000000000000000000000000000000000000000815234600482015260248101829052604401610696565b600081340390507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f338140e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611077573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109b9190611d7c565b92506000611127338f8f878d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508c8c8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116ab92505050565b905061113281611493565b94507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166367c769af848360400151888f8f6040518663ffffffff1660e01b8152600401611198949392919061206a565b60206040518083038185885af11580156111b6573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906111db9190611d7c565b84146111e9576111e961208a565b60015460408281015190517fffecec7e00000000000000000000000000000000000000000000000000000000815260048101919091526024810187905273ffffffffffffffffffffffffffffffffffffffff9091169063ffecec7e9084906044016000604051808303818588803b15801561126357600080fd5b505af1158015611277573d6000803e3d6000fd5b5050505073ffffffffffffffffffffffffffffffffffffffff8d16159050611425578b73ffffffffffffffffffffffffffffffffffffffff1663e4e065228f6112bf84610ac8565b5188868e8e6040518763ffffffff1660e01b81526004016112e5969594939291906120b9565b600060405180830381600087803b1580156112ff57600080fd5b505af1158015611313573d6000803e3d6000fd5b5050505060008c73ffffffffffffffffffffffffffffffffffffffff166362014bad6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611364573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113889190612104565b60015460408481015190517f0676b70600000000000000000000000000000000000000000000000000000000815260048101919091526024810189905273ffffffffffffffffffffffffffffffffffffffff8084166044830152929350911690630676b70690606401600060405180830381600087803b15801561140b57600080fd5b505af115801561141f573d6000803e3d6000fd5b50505050505b8060800151857f1e88c89e51d85a8c4b383cced79b9e9f26ef41f4719f5bee6e758f76e3c0450983604001518460200151856060015188888860a001518960c0015160405161147a9796959493929190612121565b60405180910390a3505050995099975050505050505050565b6000816040516020016114a69190611dc4565b604051602081830303815290604052805190602001209050919050565b60408051808201909152600080825260208201526000806114e384611689565b9092509050600160ff83161015610eab576040517fc3e3b66600000000000000000000000000000000000000000000000000000000815260ff83166004820152602401610696565b6000805b84518110156116815760007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9c9cff187848151811061158857611588612172565b602090810291909101810151604080517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff9092166004830152895160248301529189015160448201529088015160648201526060880151608482015260a401602060405180830381865afa158015611621573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116459190611d7c565b9050801580159061165f57504261165c85836121d0565b11155b156116705761166d836121e3565b92505b5061167a816121e3565b905061152f565b509392505050565b60006060828060200190518101906116a1919061221b565b9094909350915050565b6116f16040518060e00160405280600081526020016000801916815260200160008152602001600080191681526020016000815260200160608152602001606081525090565b6040518060e0016040528046815260200161171f8973ffffffffffffffffffffffffffffffffffffffff1690565b81526020018781526020018681526020018581526020018481526020018381525090509695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561179e5761179e61174c565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156117cd576117cd61174c565b604052919050565b600067ffffffffffffffff8211156117ef576117ef61174c565b50601f01601f191660200190565b600082601f83011261180e57600080fd5b813561182161181c826117d5565b6117a4565b81815284602083860101111561183657600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561186557600080fd5b813567ffffffffffffffff8082111561187d57600080fd5b9083019060e0828603121561189157600080fd5b61189961177b565b823581526020830135602082015260408301356040820152606083013560608201526080830135608082015260a0830135828111156118d757600080fd5b6118e3878286016117fd565b60a08301525060c0830135828111156118fb57600080fd5b611907878286016117fd565b60c08301525095945050505050565b60005b83811015611931578181015183820152602001611919565b50506000910152565b60008151808452611952816020860160208601611916565b601f01601f19169290920160200192915050565b602081526000611979602083018461193a565b9392505050565b60008083601f84011261199257600080fd5b50813567ffffffffffffffff8111156119aa57600080fd5b6020830191508360208285010111156119c257600080fd5b9250929050565b600080602083850312156119dc57600080fd5b823567ffffffffffffffff8111156119f357600080fd5b6119ff85828601611980565b90969095509350505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610a2b57600080fd5b60008083601f840112611a3f57600080fd5b50813567ffffffffffffffff811115611a5757600080fd5b6020830191508360208260051b85010111156119c257600080fd5b60008060008060008060008060a0898b031215611a8e57600080fd5b883597506020890135611aa081611a0b565b9650604089013567ffffffffffffffff80821115611abd57600080fd5b611ac98c838d01611a2d565b909850965060608b0135915080821115611ae257600080fd5b611aee8c838d01611980565b909650945060808b0135915080821115611b0757600080fd5b50611b148b828c01611980565b999c989b5096995094979396929594505050565b600060208284031215611b3a57600080fd5b813561197981611a0b565b600060208284031215611b5757600080fd5b5035919050565b600080600060408486031215611b7357600080fd5b83359250602084013567ffffffffffffffff811115611b9157600080fd5b611b9d86828701611980565b9497909650939450505050565b600080600080600080600080600060c08a8c031215611bc857600080fd5b8935985060208a0135611bda81611a0b565b975060408a0135611bea81611a0b565b965060608a013567ffffffffffffffff80821115611c0757600080fd5b611c138d838e01611a2d565b909850965060808c0135915080821115611c2c57600080fd5b611c388d838e01611980565b909650945060a08c0135915080821115611c5157600080fd5b50611c5e8c828d01611980565b915080935050809150509295985092959850929598565b600080600080600080600080600060c08a8c031215611c9357600080fd5b8935985060208a0135975060408a0135611bea81611a0b565b600060208284031215611cbe57600080fd5b813567ffffffffffffffff811115611cd557600080fd5b610ebf848285016117fd565b60008060408385031215611cf457600080fd5b50508035926020909101359150565b8183526000602080850194508260005b85811015611d4e578135611d2681611a0b565b73ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611d13565b509495945050505050565b838152604060208201526000611d73604083018486611d03565b95945050505050565b600060208284031215611d8e57600080fd5b5051919050565b848152836020820152826040820152608060608201526000611dba608083018461193a565b9695505050505050565b6020815281516020820152602082015160408201526040820151606082015260608201516080820152608082015160a0820152600060a083015160e060c0840152611e1361010084018261193a565b905060c0840151601f198483030160e0850152611d73828261193a565b600082601f830112611e4157600080fd5b8151611e4f61181c826117d5565b818152846020838601011115611e6457600080fd5b610ebf826020830160208701611916565b600060208284031215611e8757600080fd5b815167ffffffffffffffff80821115611e9f57600080fd5b9083019060e08286031215611eb357600080fd5b611ebb61177b565b825181526020830151602082015260408301516040820152606083015160608201526080830151608082015260a083015182811115611ef957600080fd5b611f0587828601611e30565b60a08301525060c083015182811115611f1d57600080fd5b61190787828601611e30565b60008060408385031215611f3c57600080fd5b825167ffffffffffffffff80821115611f5457600080fd5b611f6086838701611e30565b9350602091508185015181811115611f7757600080fd5b8501601f81018713611f8857600080fd5b805182811115611f9a57611f9a61174c565b8060051b9250611fab8484016117a4565b8181529282018401928481019089851115611fc557600080fd5b928501925b84841015611fef5783519250611fdf83611a0b565b8282529285019290850190611fca565b8096505050505050509250929050565b60006040828403121561201157600080fd5b6040516040810181811067ffffffffffffffff821117156120345761203461174c565b604052825181526020928301519281019290925250919050565b60006040828403121561206057600080fd5b6119798383611fff565b848152836020820152606060408201526000611dba606083018486611d03565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b86815285602082015284604082015283606082015260a060808201528160a0820152818360c0830137600081830160c090810191909152601f909201601f1916010195945050505050565b60006020828403121561211657600080fd5b815161197981611a0b565b87815286602082015285604082015284606082015283608082015260e060a0820152600061215260e083018561193a565b82810360c0840152612164818561193a565b9a9950505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561042e5761042e6121a1565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612214576122146121a1565b5060010190565b6000806040838503121561222e57600080fd5b825160ff8116811461223f57600080fd5b602084015190925067ffffffffffffffff81111561225c57600080fd5b61226885828601611e30565b915050925092905056fea2646970667358221220f8c79bac41fec192301dcc517e52d750e64c060a03eebcea90968b159c0029b564736f6c63430008140033",
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
func DeployInterchainClientV1(auth *bind.TransactOpts, backend bind.ContractBackend, interchainDB common.Address) (common.Address, *types.Transaction, *InterchainClientV1, error) {
	parsed, err := InterchainClientV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainClientV1Bin), backend, interchainDB)
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

// EncodeTransaction is a free data retrieval call binding the contract method 0x2743f804.
//
// Solidity: function encodeTransaction((uint256,bytes32,uint256,bytes32,uint256,bytes,bytes) icTx) pure returns(bytes)
func (_InterchainClientV1 *InterchainClientV1Caller) EncodeTransaction(opts *bind.CallOpts, icTx InterchainTransaction) ([]byte, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "encodeTransaction", icTx)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransaction is a free data retrieval call binding the contract method 0x2743f804.
//
// Solidity: function encodeTransaction((uint256,bytes32,uint256,bytes32,uint256,bytes,bytes) icTx) pure returns(bytes)
func (_InterchainClientV1 *InterchainClientV1Session) EncodeTransaction(icTx InterchainTransaction) ([]byte, error) {
	return _InterchainClientV1.Contract.EncodeTransaction(&_InterchainClientV1.CallOpts, icTx)
}

// EncodeTransaction is a free data retrieval call binding the contract method 0x2743f804.
//
// Solidity: function encodeTransaction((uint256,bytes32,uint256,bytes32,uint256,bytes,bytes) icTx) pure returns(bytes)
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

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes encodedTx) view returns(bool)
func (_InterchainClientV1 *InterchainClientV1Caller) IsExecutable(opts *bind.CallOpts, encodedTx []byte) (bool, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "isExecutable", encodedTx)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes encodedTx) view returns(bool)
func (_InterchainClientV1 *InterchainClientV1Session) IsExecutable(encodedTx []byte) (bool, error) {
	return _InterchainClientV1.Contract.IsExecutable(&_InterchainClientV1.CallOpts, encodedTx)
}

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes encodedTx) view returns(bool)
func (_InterchainClientV1 *InterchainClientV1CallerSession) IsExecutable(encodedTx []byte) (bool, error) {
	return _InterchainClientV1.Contract.IsExecutable(&_InterchainClientV1.CallOpts, encodedTx)
}

// LinkedClients is a free data retrieval call binding the contract method 0x7268b08f.
//
// Solidity: function linkedClients(uint256 chainId) view returns(bytes32 remoteClient)
func (_InterchainClientV1 *InterchainClientV1Caller) LinkedClients(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _InterchainClientV1.contract.Call(opts, &out, "linkedClients", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LinkedClients is a free data retrieval call binding the contract method 0x7268b08f.
//
// Solidity: function linkedClients(uint256 chainId) view returns(bytes32 remoteClient)
func (_InterchainClientV1 *InterchainClientV1Session) LinkedClients(chainId *big.Int) ([32]byte, error) {
	return _InterchainClientV1.Contract.LinkedClients(&_InterchainClientV1.CallOpts, chainId)
}

// LinkedClients is a free data retrieval call binding the contract method 0x7268b08f.
//
// Solidity: function linkedClients(uint256 chainId) view returns(bytes32 remoteClient)
func (_InterchainClientV1 *InterchainClientV1CallerSession) LinkedClients(chainId *big.Int) ([32]byte, error) {
	return _InterchainClientV1.Contract.LinkedClients(&_InterchainClientV1.CallOpts, chainId)
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

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_InterchainClientV1 *InterchainClientV1Transactor) InterchainExecute(opts *bind.TransactOpts, gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "interchainExecute", gasLimit, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_InterchainClientV1 *InterchainClientV1Session) InterchainExecute(gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainExecute(&_InterchainClientV1.TransactOpts, gasLimit, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_InterchainClientV1 *InterchainClientV1TransactorSession) InterchainExecute(gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainExecute(&_InterchainClientV1.TransactOpts, gasLimit, transaction)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_InterchainClientV1 *InterchainClientV1Transactor) InterchainSend(opts *bind.TransactOpts, dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "interchainSend", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_InterchainClientV1 *InterchainClientV1Session) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSend(&_InterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_InterchainClientV1 *InterchainClientV1TransactorSession) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSend(&_InterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_InterchainClientV1 *InterchainClientV1Transactor) InterchainSendEVM(opts *bind.TransactOpts, dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "interchainSendEVM", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
func (_InterchainClientV1 *InterchainClientV1Session) InterchainSendEVM(dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.InterchainSendEVM(&_InterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns(bytes32 transactionId, uint256 dbNonce)
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
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce)
func (_InterchainClientV1 *InterchainClientV1Transactor) WriteExecutionProof(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.contract.Transact(opts, "writeExecutionProof", transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce)
func (_InterchainClientV1 *InterchainClientV1Session) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.WriteExecutionProof(&_InterchainClientV1.TransactOpts, transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce)
func (_InterchainClientV1 *InterchainClientV1TransactorSession) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _InterchainClientV1.Contract.WriteExecutionProof(&_InterchainClientV1.TransactOpts, transactionId)
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
	Executor      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionProofWritten is a free log retrieval operation binding the contract event 0xf9e8d86f4b1412ca26874e2b00b21627542627093b367c3dc95e612c4de927c1.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, address executor)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterExecutionProofWritten(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int) (*InterchainClientV1ExecutionProofWrittenIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1ExecutionProofWrittenIterator{contract: _InterchainClientV1.contract, event: "ExecutionProofWritten", logs: logs, sub: sub}, nil
}

// WatchExecutionProofWritten is a free log subscription operation binding the contract event 0xf9e8d86f4b1412ca26874e2b00b21627542627093b367c3dc95e612c4de927c1.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, address executor)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchExecutionProofWritten(opts *bind.WatchOpts, sink chan<- *InterchainClientV1ExecutionProofWritten, transactionId [][32]byte, dbNonce []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule)
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

// ParseExecutionProofWritten is a log parse operation binding the contract event 0xf9e8d86f4b1412ca26874e2b00b21627542627093b367c3dc95e612c4de927c1.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, address executor)
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
	SrcChainId    *big.Int
	SrcSender     [32]byte
	DstReceiver   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionReceived is a free log retrieval operation binding the contract event 0xdb272c7b2f65161c0573b8b7c21bc3e25cf53bcd8c1524fb3fb78bac8c7df94e.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterInterchainTransactionReceived(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int) (*InterchainClientV1InterchainTransactionReceivedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1InterchainTransactionReceivedIterator{contract: _InterchainClientV1.contract, event: "InterchainTransactionReceived", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionReceived is a free log subscription operation binding the contract event 0xdb272c7b2f65161c0573b8b7c21bc3e25cf53bcd8c1524fb3fb78bac8c7df94e.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchInterchainTransactionReceived(opts *bind.WatchOpts, sink chan<- *InterchainClientV1InterchainTransactionReceived, transactionId [][32]byte, dbNonce []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule)
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

// ParseInterchainTransactionReceived is a log parse operation binding the contract event 0xdb272c7b2f65161c0573b8b7c21bc3e25cf53bcd8c1524fb3fb78bac8c7df94e.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
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
	DstChainId      *big.Int
	SrcSender       [32]byte
	DstReceiver     [32]byte
	VerificationFee *big.Int
	ExecutionFee    *big.Int
	Options         []byte
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionSent is a free log retrieval operation binding the contract event 0x1e88c89e51d85a8c4b383cced79b9e9f26ef41f4719f5bee6e758f76e3c04509.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1 *InterchainClientV1Filterer) FilterInterchainTransactionSent(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int) (*InterchainClientV1InterchainTransactionSentIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1.contract.FilterLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1InterchainTransactionSentIterator{contract: _InterchainClientV1.contract, event: "InterchainTransactionSent", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionSent is a free log subscription operation binding the contract event 0x1e88c89e51d85a8c4b383cced79b9e9f26ef41f4719f5bee6e758f76e3c04509.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1 *InterchainClientV1Filterer) WatchInterchainTransactionSent(opts *bind.WatchOpts, sink chan<- *InterchainClientV1InterchainTransactionSent, transactionId [][32]byte, dbNonce []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1.contract.WatchLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule)
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

// ParseInterchainTransactionSent is a log parse operation binding the contract event 0x1e88c89e51d85a8c4b383cced79b9e9f26ef41f4719f5bee6e758f76e3c04509.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
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

// InterchainClientV1EventsMetaData contains all meta data concerning the InterchainClientV1Events contract.
var InterchainClientV1EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionProofWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"}],\"name\":\"InterchainTransactionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"InterchainTransactionSent\",\"type\":\"event\"}]",
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
	Executor      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionProofWritten is a free log retrieval operation binding the contract event 0xf9e8d86f4b1412ca26874e2b00b21627542627093b367c3dc95e612c4de927c1.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, address executor)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterExecutionProofWritten(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int) (*InterchainClientV1EventsExecutionProofWrittenIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsExecutionProofWrittenIterator{contract: _InterchainClientV1Events.contract, event: "ExecutionProofWritten", logs: logs, sub: sub}, nil
}

// WatchExecutionProofWritten is a free log subscription operation binding the contract event 0xf9e8d86f4b1412ca26874e2b00b21627542627093b367c3dc95e612c4de927c1.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, address executor)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchExecutionProofWritten(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsExecutionProofWritten, transactionId [][32]byte, dbNonce []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule)
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

// ParseExecutionProofWritten is a log parse operation binding the contract event 0xf9e8d86f4b1412ca26874e2b00b21627542627093b367c3dc95e612c4de927c1.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, address executor)
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
	SrcChainId    *big.Int
	SrcSender     [32]byte
	DstReceiver   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionReceived is a free log retrieval operation binding the contract event 0xdb272c7b2f65161c0573b8b7c21bc3e25cf53bcd8c1524fb3fb78bac8c7df94e.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterInterchainTransactionReceived(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int) (*InterchainClientV1EventsInterchainTransactionReceivedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsInterchainTransactionReceivedIterator{contract: _InterchainClientV1Events.contract, event: "InterchainTransactionReceived", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionReceived is a free log subscription operation binding the contract event 0xdb272c7b2f65161c0573b8b7c21bc3e25cf53bcd8c1524fb3fb78bac8c7df94e.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchInterchainTransactionReceived(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsInterchainTransactionReceived, transactionId [][32]byte, dbNonce []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule)
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

// ParseInterchainTransactionReceived is a log parse operation binding the contract event 0xdb272c7b2f65161c0573b8b7c21bc3e25cf53bcd8c1524fb3fb78bac8c7df94e.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
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
	DstChainId      *big.Int
	SrcSender       [32]byte
	DstReceiver     [32]byte
	VerificationFee *big.Int
	ExecutionFee    *big.Int
	Options         []byte
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionSent is a free log retrieval operation binding the contract event 0x1e88c89e51d85a8c4b383cced79b9e9f26ef41f4719f5bee6e758f76e3c04509.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterInterchainTransactionSent(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int) (*InterchainClientV1EventsInterchainTransactionSentIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsInterchainTransactionSentIterator{contract: _InterchainClientV1Events.contract, event: "InterchainTransactionSent", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionSent is a free log subscription operation binding the contract event 0x1e88c89e51d85a8c4b383cced79b9e9f26ef41f4719f5bee6e758f76e3c04509.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchInterchainTransactionSent(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsInterchainTransactionSent, transactionId [][32]byte, dbNonce []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule)
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

// ParseInterchainTransactionSent is a log parse operation binding the contract event 0x1e88c89e51d85a8c4b383cced79b9e9f26ef41f4719f5bee6e758f76e3c04509.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseInterchainTransactionSent(log types.Log) (*InterchainClientV1EventsInterchainTransactionSent, error) {
	event := new(InterchainClientV1EventsInterchainTransactionSent)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209e4fbeac41cf438389a79111c16f60b590515b19d7989dca3536a81a4331911164736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122009ff36f74d8c162b5c04f4e766b1a10c4863d48ecce18aceb3cda87c99483ae664736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d985fd5f009464a9c28d3e65e3299ff37a12b72e62acda3c73a3b705de5f83e264736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209d5cc7bea7435ceb84ccdcc4d7260e463e2fc7205d7c623b354a6b98b5797d5064736f6c63430008140033",
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
