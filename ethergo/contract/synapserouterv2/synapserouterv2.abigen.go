// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package synapserouterv2

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

// BridgeToken is an auto generated low-level Go binding around an user-defined struct.
type BridgeToken struct {
	Symbol string
	Token  common.Address
}

// DestRequest is an auto generated low-level Go binding around an user-defined struct.
type DestRequest struct {
	Symbol   string
	AmountIn *big.Int
}

// LimitedToken is an auto generated low-level Go binding around an user-defined struct.
type LimitedToken struct {
	ActionMask *big.Int
	Token      common.Address
}

// Pool is an auto generated low-level Go binding around an user-defined struct.
type Pool struct {
	Pool    common.Address
	LpToken common.Address
	Tokens  []PoolToken
}

// PoolToken is an auto generated low-level Go binding around an user-defined struct.
type PoolToken struct {
	IsWeth bool
	Token  common.Address
}

// SwapQuery is an auto generated low-level Go binding around an user-defined struct.
type SwapQuery struct {
	RouterAdapter common.Address
	TokenOut      common.Address
	MinAmountOut  *big.Int
	Deadline      *big.Int
	RawParams     []byte
}

// ActionLibMetaData contains all meta data concerning the ActionLib contract.
var ActionLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f2f0c2e703038d9bcd0b65eac503583c9ca7b6f45cee488f9fd58559c4e3634264736f6c63430008110033",
}

// ActionLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ActionLibMetaData.ABI instead.
var ActionLibABI = ActionLibMetaData.ABI

// ActionLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ActionLibMetaData.Bin instead.
var ActionLibBin = ActionLibMetaData.Bin

// DeployActionLib deploys a new Ethereum contract, binding an instance of ActionLib to it.
func DeployActionLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ActionLib, error) {
	parsed, err := ActionLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ActionLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ActionLib{ActionLibCaller: ActionLibCaller{contract: contract}, ActionLibTransactor: ActionLibTransactor{contract: contract}, ActionLibFilterer: ActionLibFilterer{contract: contract}}, nil
}

// ActionLib is an auto generated Go binding around an Ethereum contract.
type ActionLib struct {
	ActionLibCaller     // Read-only binding to the contract
	ActionLibTransactor // Write-only binding to the contract
	ActionLibFilterer   // Log filterer for contract events
}

// ActionLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActionLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActionLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActionLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActionLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActionLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActionLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActionLibSession struct {
	Contract     *ActionLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActionLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActionLibCallerSession struct {
	Contract *ActionLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ActionLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActionLibTransactorSession struct {
	Contract     *ActionLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ActionLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActionLibRaw struct {
	Contract *ActionLib // Generic contract binding to access the raw methods on
}

// ActionLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActionLibCallerRaw struct {
	Contract *ActionLibCaller // Generic read-only contract binding to access the raw methods on
}

// ActionLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActionLibTransactorRaw struct {
	Contract *ActionLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActionLib creates a new instance of ActionLib, bound to a specific deployed contract.
func NewActionLib(address common.Address, backend bind.ContractBackend) (*ActionLib, error) {
	contract, err := bindActionLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActionLib{ActionLibCaller: ActionLibCaller{contract: contract}, ActionLibTransactor: ActionLibTransactor{contract: contract}, ActionLibFilterer: ActionLibFilterer{contract: contract}}, nil
}

// NewActionLibCaller creates a new read-only instance of ActionLib, bound to a specific deployed contract.
func NewActionLibCaller(address common.Address, caller bind.ContractCaller) (*ActionLibCaller, error) {
	contract, err := bindActionLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActionLibCaller{contract: contract}, nil
}

// NewActionLibTransactor creates a new write-only instance of ActionLib, bound to a specific deployed contract.
func NewActionLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ActionLibTransactor, error) {
	contract, err := bindActionLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActionLibTransactor{contract: contract}, nil
}

// NewActionLibFilterer creates a new log filterer instance of ActionLib, bound to a specific deployed contract.
func NewActionLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ActionLibFilterer, error) {
	contract, err := bindActionLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActionLibFilterer{contract: contract}, nil
}

// bindActionLib binds a generic wrapper to an already deployed contract.
func bindActionLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ActionLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActionLib *ActionLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ActionLib.Contract.ActionLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActionLib *ActionLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActionLib.Contract.ActionLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActionLib *ActionLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActionLib.Contract.ActionLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActionLib *ActionLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ActionLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActionLib *ActionLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActionLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActionLib *ActionLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActionLib.Contract.contract.Transact(opts, method, params...)
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122066a6d6aeb1acd9ca5c808c53e53abbefe76d936e1f108dd2373ffefa64a0166b64736f6c63430008110033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ArraysMetaData contains all meta data concerning the Arrays contract.
var ArraysMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"ArrayLengthInvalid\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a26b66c076f23e7401008ec92d9a72ab6fcaddf8d33e1f0eb46e0b80b6cf100864736f6c63430008110033",
}

// ArraysABI is the input ABI used to generate the binding from.
// Deprecated: Use ArraysMetaData.ABI instead.
var ArraysABI = ArraysMetaData.ABI

// ArraysBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArraysMetaData.Bin instead.
var ArraysBin = ArraysMetaData.Bin

// DeployArrays deploys a new Ethereum contract, binding an instance of Arrays to it.
func DeployArrays(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Arrays, error) {
	parsed, err := ArraysMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArraysBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Arrays{ArraysCaller: ArraysCaller{contract: contract}, ArraysTransactor: ArraysTransactor{contract: contract}, ArraysFilterer: ArraysFilterer{contract: contract}}, nil
}

// Arrays is an auto generated Go binding around an Ethereum contract.
type Arrays struct {
	ArraysCaller     // Read-only binding to the contract
	ArraysTransactor // Write-only binding to the contract
	ArraysFilterer   // Log filterer for contract events
}

// ArraysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArraysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArraysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArraysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArraysSession struct {
	Contract     *Arrays           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArraysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArraysCallerSession struct {
	Contract *ArraysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArraysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArraysTransactorSession struct {
	Contract     *ArraysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArraysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArraysRaw struct {
	Contract *Arrays // Generic contract binding to access the raw methods on
}

// ArraysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArraysCallerRaw struct {
	Contract *ArraysCaller // Generic read-only contract binding to access the raw methods on
}

// ArraysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArraysTransactorRaw struct {
	Contract *ArraysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArrays creates a new instance of Arrays, bound to a specific deployed contract.
func NewArrays(address common.Address, backend bind.ContractBackend) (*Arrays, error) {
	contract, err := bindArrays(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arrays{ArraysCaller: ArraysCaller{contract: contract}, ArraysTransactor: ArraysTransactor{contract: contract}, ArraysFilterer: ArraysFilterer{contract: contract}}, nil
}

// NewArraysCaller creates a new read-only instance of Arrays, bound to a specific deployed contract.
func NewArraysCaller(address common.Address, caller bind.ContractCaller) (*ArraysCaller, error) {
	contract, err := bindArrays(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArraysCaller{contract: contract}, nil
}

// NewArraysTransactor creates a new write-only instance of Arrays, bound to a specific deployed contract.
func NewArraysTransactor(address common.Address, transactor bind.ContractTransactor) (*ArraysTransactor, error) {
	contract, err := bindArrays(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArraysTransactor{contract: contract}, nil
}

// NewArraysFilterer creates a new log filterer instance of Arrays, bound to a specific deployed contract.
func NewArraysFilterer(address common.Address, filterer bind.ContractFilterer) (*ArraysFilterer, error) {
	contract, err := bindArrays(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArraysFilterer{contract: contract}, nil
}

// bindArrays binds a generic wrapper to an already deployed contract.
func bindArrays(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArraysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arrays *ArraysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arrays.Contract.ArraysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arrays *ArraysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arrays.Contract.ArraysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arrays *ArraysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arrays.Contract.ArraysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arrays *ArraysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arrays.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arrays *ArraysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arrays.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arrays *ArraysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arrays.Contract.contract.Transact(opts, method, params...)
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

// DefaultAdapterMetaData contains all meta data concerning the DefaultAdapter contract.
var DefaultAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MsgValueIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAddressMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensIdentical\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"name\":\"adapterSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"24a98f11": "adapterSwap(address,address,uint256,address,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b506117b6806100206000396000f3fe6080604052600436106100225760003560e01c806324a98f111461002e57600080fd5b3661002957005b600080fd5b61004161003c366004611382565b610053565b60405190815260200160405180910390f35b6000610062868686868661006c565b9695505050505050565b60008061007a86858561011a565b905061008886868684610233565b955060006100978786846102d1565b90506100a58787838561032b565b92507fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff8616016100ee576100ee8184610512565b61010f73ffffffffffffffffffffffffffffffffffffffff86168985610596565b505095945050505050565b6040805160808101825260008082526020820181905291810182905260608101919091528273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036101a3576040517f0b839a1f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b818060200190518101906101b7919061148d565b602081015190915073ffffffffffffffffffffffffffffffffffffffff161580156101f557506003815160038111156101f2576101f2611512565b14155b1561022c576040517f76ecffc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9392505050565b60007fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff86160161028f5761027e838360016106f0565b905061028a81856107c2565b6102c9565b508334156102c9576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b949350505050565b60007fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff8416016103235761031c848360006106f0565b905061022c565b509092915050565b600060038251600381111561034257610342611512565b0361034e5750826102c9565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa1580156103b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103dc9190611541565b60208301519091506104069073ffffffffffffffffffffffffffffffffffffffff87169086610860565b60008251600381111561041b5761041b611512565b036104355761043082602001518386866109dc565b61046f565b60018251600381111561044a5761044a611512565b0361045f576104308260200151838686610ba6565b61046f8260200151838686610d55565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152819073ffffffffffffffffffffffffffffffffffffffff8516906370a0823190602401602060405180830381865afa1580156104db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ff9190611541565b6105099190611589565b95945050505050565b6040517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526004810182905273ffffffffffffffffffffffffffffffffffffffff831690632e1a7d4d90602401600060405180830381600087803b15801561057a57600080fd5b505af115801561058e573d6000803e3d6000fd5b505050505050565b3073ffffffffffffffffffffffffffffffffffffffff8316036105b857505050565b7fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff8416016106ca5760008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d806000811461064f576040519150601f19603f3d011682016040523d82523d6000602084013e610654565b606091505b50509050806106c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c65640000000000000000000000000060448201526064015b60405180910390fd5b50505050565b6106eb73ffffffffffffffffffffffffffffffffffffffff84168383610ecd565b505050565b600060038351600381111561070757610707611512565b0361071357508261022c565b826020015173ffffffffffffffffffffffffffffffffffffffff166382b8660083610742578460600151610748565b84604001515b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815260ff9091166004820152602401602060405180830381865afa15801561079e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c991906115a2565b3481146107fb576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561084357600080fd5b505af1158015610857573d6000803e3d6000fd5b50505050505050565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156108d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fb9190611541565b61090591906115bf565b60405173ffffffffffffffffffffffffffffffffffffffff85166024820152604481018290529091506106c49085907f095ea7b300000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610f23565b60608301516040517f82b8660000000000000000000000000000000000000000000000000000000000815260ff909116600482015273ffffffffffffffffffffffffffffffffffffffff82811691908616906382b8660090602401602060405180830381865afa158015610a54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7891906115a2565b73ffffffffffffffffffffffffffffffffffffffff1614610ac5576040517f28716b9200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080840151606085015191517f9169558600000000000000000000000000000000000000000000000000000000815260ff91821660048201529116602482015260448101839052600060648201527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608482015273ffffffffffffffffffffffffffffffffffffffff85169063916955869060a4015b6020604051808303816000875af1158015610b7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9f9190611541565b5050505050565b6000610bb18561102f565b90506000610bbe86611100565b90508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c25576040517f28716b9200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008267ffffffffffffffff811115610c4057610c40611304565b604051908082528060200260200182016040528015610c69578160200160208202803683370190505b5090508481876040015160ff1681518110610c8657610c866115d2565b60209081029190910101526040517f4d49e87d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff881690634d49e87d90610d089084906000907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90600401611601565b6020604051808303816000875af1158015610d27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4b9190611541565b5050505050505050565b60608301516040517f82b8660000000000000000000000000000000000000000000000000000000000815260ff909116600482015273ffffffffffffffffffffffffffffffffffffffff82811691908616906382b8660090602401602060405180830381865afa158015610dcd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df191906115a2565b73ffffffffffffffffffffffffffffffffffffffff1614610e3e576040517f28716b9200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608301516040517f3e3a15600000000000000000000000000000000000000000000000000000000081526004810184905260ff9091166024820152600060448201527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff606482015273ffffffffffffffffffffffffffffffffffffffff851690633e3a156090608401610b5c565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526106eb9084907fa9059cbb000000000000000000000000000000000000000000000000000000009060640161095a565b6000610f85826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661117d9092919063ffffffff16565b8051909150156106eb5780806020019051810190610fa3919061164e565b6106eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016106bb565b6000805b6040517f82b8660000000000000000000000000000000000000000000000000000000000815260ff8216600482015273ffffffffffffffffffffffffffffffffffffffff8416906382b8660090602401602060405180830381865afa9250505080156110da575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526110d7918101906115a2565b60015b6110e9578060ff1691506110fa565b506110f381611670565b9050611033565b50919050565b60008173ffffffffffffffffffffffffffffffffffffffff16635fd65f0f6040518163ffffffff1660e01b815260040160e060405180830381865afa15801561114d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611171919061168f565b98975050505050505050565b60606102c984846000858573ffffffffffffffffffffffffffffffffffffffff85163b611206576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016106bb565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161122f9190611713565b60006040518083038185875af1925050503d806000811461126c576040519150601f19603f3d011682016040523d82523d6000602084013e611271565b606091505b509150915061128182828661128c565b979650505050505050565b6060831561129b57508161022c565b8251156112ab5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bb919061172f565b73ffffffffffffffffffffffffffffffffffffffff8116811461130157600080fd5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561137a5761137a611304565b604052919050565b600080600080600060a0868803121561139a57600080fd5b85356113a5816112df565b94506020868101356113b6816112df565b94506040870135935060608701356113cd816112df565b9250608087013567ffffffffffffffff808211156113ea57600080fd5b818901915089601f8301126113fe57600080fd5b81358181111561141057611410611304565b611440847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611333565b91508082528a8482850101111561145657600080fd5b80848401858401376000848284010152508093505050509295509295909350565b805160ff8116811461148857600080fd5b919050565b60006080828403121561149f57600080fd5b6040516080810181811067ffffffffffffffff821117156114c2576114c2611304565b6040528251600481106114d457600080fd5b815260208301516114e4816112df565b60208201526114f560408401611477565b604082015261150660608401611477565b60608201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60006020828403121561155357600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561159c5761159c61155a565b92915050565b6000602082840312156115b457600080fd5b815161022c816112df565b8082018082111561159c5761159c61155a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b606080825284519082018190526000906020906080840190828801845b8281101561163a5781518452928401929084019060010161161e565b505050908301949094525060400152919050565b60006020828403121561166057600080fd5b8151801515811461022c57600080fd5b600060ff821660ff81036116865761168661155a565b60010192915050565b600080600080600080600060e0888a0312156116aa57600080fd5b875196506020880151955060408801519450606088015193506080880151925060a0880151915060c08801516116df816112df565b8091505092959891949750929550565b60005b8381101561170a5781810151838201526020016116f2565b50506000910152565b600082516117258184602087016116ef565b9190910192915050565b602081526000825180602084015261174e8160408501602087016116ef565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea2646970667358221220c9dcd0c0465024964615cf3eb09a30cd865bd2daa7b4331fea5836144a34108a64736f6c63430008110033",
}

// DefaultAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use DefaultAdapterMetaData.ABI instead.
var DefaultAdapterABI = DefaultAdapterMetaData.ABI

// Deprecated: Use DefaultAdapterMetaData.Sigs instead.
// DefaultAdapterFuncSigs maps the 4-byte function signature to its string representation.
var DefaultAdapterFuncSigs = DefaultAdapterMetaData.Sigs

// DefaultAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DefaultAdapterMetaData.Bin instead.
var DefaultAdapterBin = DefaultAdapterMetaData.Bin

// DeployDefaultAdapter deploys a new Ethereum contract, binding an instance of DefaultAdapter to it.
func DeployDefaultAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DefaultAdapter, error) {
	parsed, err := DefaultAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DefaultAdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DefaultAdapter{DefaultAdapterCaller: DefaultAdapterCaller{contract: contract}, DefaultAdapterTransactor: DefaultAdapterTransactor{contract: contract}, DefaultAdapterFilterer: DefaultAdapterFilterer{contract: contract}}, nil
}

// DefaultAdapter is an auto generated Go binding around an Ethereum contract.
type DefaultAdapter struct {
	DefaultAdapterCaller     // Read-only binding to the contract
	DefaultAdapterTransactor // Write-only binding to the contract
	DefaultAdapterFilterer   // Log filterer for contract events
}

// DefaultAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefaultAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefaultAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefaultAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefaultAdapterSession struct {
	Contract     *DefaultAdapter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefaultAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefaultAdapterCallerSession struct {
	Contract *DefaultAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DefaultAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefaultAdapterTransactorSession struct {
	Contract     *DefaultAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DefaultAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefaultAdapterRaw struct {
	Contract *DefaultAdapter // Generic contract binding to access the raw methods on
}

// DefaultAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefaultAdapterCallerRaw struct {
	Contract *DefaultAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// DefaultAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefaultAdapterTransactorRaw struct {
	Contract *DefaultAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefaultAdapter creates a new instance of DefaultAdapter, bound to a specific deployed contract.
func NewDefaultAdapter(address common.Address, backend bind.ContractBackend) (*DefaultAdapter, error) {
	contract, err := bindDefaultAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefaultAdapter{DefaultAdapterCaller: DefaultAdapterCaller{contract: contract}, DefaultAdapterTransactor: DefaultAdapterTransactor{contract: contract}, DefaultAdapterFilterer: DefaultAdapterFilterer{contract: contract}}, nil
}

// NewDefaultAdapterCaller creates a new read-only instance of DefaultAdapter, bound to a specific deployed contract.
func NewDefaultAdapterCaller(address common.Address, caller bind.ContractCaller) (*DefaultAdapterCaller, error) {
	contract, err := bindDefaultAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultAdapterCaller{contract: contract}, nil
}

// NewDefaultAdapterTransactor creates a new write-only instance of DefaultAdapter, bound to a specific deployed contract.
func NewDefaultAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*DefaultAdapterTransactor, error) {
	contract, err := bindDefaultAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultAdapterTransactor{contract: contract}, nil
}

// NewDefaultAdapterFilterer creates a new log filterer instance of DefaultAdapter, bound to a specific deployed contract.
func NewDefaultAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*DefaultAdapterFilterer, error) {
	contract, err := bindDefaultAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefaultAdapterFilterer{contract: contract}, nil
}

// bindDefaultAdapter binds a generic wrapper to an already deployed contract.
func bindDefaultAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefaultAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultAdapter *DefaultAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultAdapter.Contract.DefaultAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultAdapter *DefaultAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAdapter.Contract.DefaultAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultAdapter *DefaultAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultAdapter.Contract.DefaultAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultAdapter *DefaultAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultAdapter *DefaultAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultAdapter *DefaultAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultAdapter.Contract.contract.Transact(opts, method, params...)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_DefaultAdapter *DefaultAdapterTransactor) AdapterSwap(opts *bind.TransactOpts, recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _DefaultAdapter.contract.Transact(opts, "adapterSwap", recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_DefaultAdapter *DefaultAdapterSession) AdapterSwap(recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _DefaultAdapter.Contract.AdapterSwap(&_DefaultAdapter.TransactOpts, recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_DefaultAdapter *DefaultAdapterTransactorSession) AdapterSwap(recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _DefaultAdapter.Contract.AdapterSwap(&_DefaultAdapter.TransactOpts, recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAdapter *DefaultAdapterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAdapter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAdapter *DefaultAdapterSession) Receive() (*types.Transaction, error) {
	return _DefaultAdapter.Contract.Receive(&_DefaultAdapter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAdapter *DefaultAdapterTransactorSession) Receive() (*types.Transaction, error) {
	return _DefaultAdapter.Contract.Receive(&_DefaultAdapter.TransactOpts)
}

// DefaultRouterMetaData contains all meta data concerning the DefaultRouter contract.
var DefaultRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MsgValueIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAddressMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensIdentical\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"name\":\"adapterSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"24a98f11": "adapterSwap(address,address,uint256,address,bytes)",
	},
}

// DefaultRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use DefaultRouterMetaData.ABI instead.
var DefaultRouterABI = DefaultRouterMetaData.ABI

// Deprecated: Use DefaultRouterMetaData.Sigs instead.
// DefaultRouterFuncSigs maps the 4-byte function signature to its string representation.
var DefaultRouterFuncSigs = DefaultRouterMetaData.Sigs

// DefaultRouter is an auto generated Go binding around an Ethereum contract.
type DefaultRouter struct {
	DefaultRouterCaller     // Read-only binding to the contract
	DefaultRouterTransactor // Write-only binding to the contract
	DefaultRouterFilterer   // Log filterer for contract events
}

// DefaultRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefaultRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefaultRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefaultRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefaultRouterSession struct {
	Contract     *DefaultRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefaultRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefaultRouterCallerSession struct {
	Contract *DefaultRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DefaultRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefaultRouterTransactorSession struct {
	Contract     *DefaultRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DefaultRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefaultRouterRaw struct {
	Contract *DefaultRouter // Generic contract binding to access the raw methods on
}

// DefaultRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefaultRouterCallerRaw struct {
	Contract *DefaultRouterCaller // Generic read-only contract binding to access the raw methods on
}

// DefaultRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefaultRouterTransactorRaw struct {
	Contract *DefaultRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefaultRouter creates a new instance of DefaultRouter, bound to a specific deployed contract.
func NewDefaultRouter(address common.Address, backend bind.ContractBackend) (*DefaultRouter, error) {
	contract, err := bindDefaultRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefaultRouter{DefaultRouterCaller: DefaultRouterCaller{contract: contract}, DefaultRouterTransactor: DefaultRouterTransactor{contract: contract}, DefaultRouterFilterer: DefaultRouterFilterer{contract: contract}}, nil
}

// NewDefaultRouterCaller creates a new read-only instance of DefaultRouter, bound to a specific deployed contract.
func NewDefaultRouterCaller(address common.Address, caller bind.ContractCaller) (*DefaultRouterCaller, error) {
	contract, err := bindDefaultRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultRouterCaller{contract: contract}, nil
}

// NewDefaultRouterTransactor creates a new write-only instance of DefaultRouter, bound to a specific deployed contract.
func NewDefaultRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*DefaultRouterTransactor, error) {
	contract, err := bindDefaultRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultRouterTransactor{contract: contract}, nil
}

// NewDefaultRouterFilterer creates a new log filterer instance of DefaultRouter, bound to a specific deployed contract.
func NewDefaultRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*DefaultRouterFilterer, error) {
	contract, err := bindDefaultRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefaultRouterFilterer{contract: contract}, nil
}

// bindDefaultRouter binds a generic wrapper to an already deployed contract.
func bindDefaultRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefaultRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultRouter *DefaultRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultRouter.Contract.DefaultRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultRouter *DefaultRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultRouter.Contract.DefaultRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultRouter *DefaultRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultRouter.Contract.DefaultRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultRouter *DefaultRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultRouter *DefaultRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultRouter *DefaultRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultRouter.Contract.contract.Transact(opts, method, params...)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_DefaultRouter *DefaultRouterTransactor) AdapterSwap(opts *bind.TransactOpts, recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _DefaultRouter.contract.Transact(opts, "adapterSwap", recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_DefaultRouter *DefaultRouterSession) AdapterSwap(recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _DefaultRouter.Contract.AdapterSwap(&_DefaultRouter.TransactOpts, recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_DefaultRouter *DefaultRouterTransactorSession) AdapterSwap(recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _DefaultRouter.Contract.AdapterSwap(&_DefaultRouter.TransactOpts, recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultRouter *DefaultRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultRouter *DefaultRouterSession) Receive() (*types.Transaction, error) {
	return _DefaultRouter.Contract.Receive(&_DefaultRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultRouter *DefaultRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _DefaultRouter.Contract.Receive(&_DefaultRouter.TransactOpts)
}

// EnumerableMapMetaData contains all meta data concerning the EnumerableMap contract.
var EnumerableMapMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122076dd1db60671830f4f33275bcc3394b8dfa502ad306c7cfb35ad9b5b9d661de464736f6c63430008110033",
}

// EnumerableMapABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableMapMetaData.ABI instead.
var EnumerableMapABI = EnumerableMapMetaData.ABI

// EnumerableMapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableMapMetaData.Bin instead.
var EnumerableMapBin = EnumerableMapMetaData.Bin

// DeployEnumerableMap deploys a new Ethereum contract, binding an instance of EnumerableMap to it.
func DeployEnumerableMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableMap, error) {
	parsed, err := EnumerableMapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableMapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// EnumerableMap is an auto generated Go binding around an Ethereum contract.
type EnumerableMap struct {
	EnumerableMapCaller     // Read-only binding to the contract
	EnumerableMapTransactor // Write-only binding to the contract
	EnumerableMapFilterer   // Log filterer for contract events
}

// EnumerableMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableMapSession struct {
	Contract     *EnumerableMap    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableMapCallerSession struct {
	Contract *EnumerableMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableMapTransactorSession struct {
	Contract     *EnumerableMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableMapRaw struct {
	Contract *EnumerableMap // Generic contract binding to access the raw methods on
}

// EnumerableMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableMapCallerRaw struct {
	Contract *EnumerableMapCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableMapTransactorRaw struct {
	Contract *EnumerableMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableMap creates a new instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMap(address common.Address, backend bind.ContractBackend) (*EnumerableMap, error) {
	contract, err := bindEnumerableMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// NewEnumerableMapCaller creates a new read-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapCaller(address common.Address, caller bind.ContractCaller) (*EnumerableMapCaller, error) {
	contract, err := bindEnumerableMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapCaller{contract: contract}, nil
}

// NewEnumerableMapTransactor creates a new write-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableMapTransactor, error) {
	contract, err := bindEnumerableMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapTransactor{contract: contract}, nil
}

// NewEnumerableMapFilterer creates a new log filterer instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableMapFilterer, error) {
	contract, err := bindEnumerableMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapFilterer{contract: contract}, nil
}

// bindEnumerableMap binds a generic wrapper to an already deployed contract.
func bindEnumerableMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableMapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.EnumerableMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200a141b9612aa319129b811f2d73db44e94c615da02ccc8ab036293080db57a4464736f6c63430008110033",
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

// IBridgeModuleMetaData contains all meta data concerning the IBridgeModule contract.
var IBridgeModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSwap\",\"type\":\"bool\"}],\"name\":\"calculateFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"destQuery\",\"type\":\"tuple\"}],\"name\":\"delegateBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"bridgeTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getMaxBridgedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenToActionMask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actionMask\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenToSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0d25aafe": "calculateFeeAmount(address,uint256,bool)",
		"436f3aa5": "delegateBridge(address,uint256,address,uint256,(address,address,uint256,uint256,bytes))",
		"9c1d060e": "getBridgeTokens()",
		"04b1ac29": "getMaxBridgedAmount(address)",
		"a5bc29c2": "symbolToToken(string)",
		"98b57505": "tokenToActionMask(address)",
		"0ba36121": "tokenToSymbol(address)",
	},
}

// IBridgeModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use IBridgeModuleMetaData.ABI instead.
var IBridgeModuleABI = IBridgeModuleMetaData.ABI

// Deprecated: Use IBridgeModuleMetaData.Sigs instead.
// IBridgeModuleFuncSigs maps the 4-byte function signature to its string representation.
var IBridgeModuleFuncSigs = IBridgeModuleMetaData.Sigs

// IBridgeModule is an auto generated Go binding around an Ethereum contract.
type IBridgeModule struct {
	IBridgeModuleCaller     // Read-only binding to the contract
	IBridgeModuleTransactor // Write-only binding to the contract
	IBridgeModuleFilterer   // Log filterer for contract events
}

// IBridgeModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeModuleSession struct {
	Contract     *IBridgeModule    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeModuleCallerSession struct {
	Contract *IBridgeModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IBridgeModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeModuleTransactorSession struct {
	Contract     *IBridgeModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBridgeModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeModuleRaw struct {
	Contract *IBridgeModule // Generic contract binding to access the raw methods on
}

// IBridgeModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeModuleCallerRaw struct {
	Contract *IBridgeModuleCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeModuleTransactorRaw struct {
	Contract *IBridgeModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridgeModule creates a new instance of IBridgeModule, bound to a specific deployed contract.
func NewIBridgeModule(address common.Address, backend bind.ContractBackend) (*IBridgeModule, error) {
	contract, err := bindIBridgeModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridgeModule{IBridgeModuleCaller: IBridgeModuleCaller{contract: contract}, IBridgeModuleTransactor: IBridgeModuleTransactor{contract: contract}, IBridgeModuleFilterer: IBridgeModuleFilterer{contract: contract}}, nil
}

// NewIBridgeModuleCaller creates a new read-only instance of IBridgeModule, bound to a specific deployed contract.
func NewIBridgeModuleCaller(address common.Address, caller bind.ContractCaller) (*IBridgeModuleCaller, error) {
	contract, err := bindIBridgeModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeModuleCaller{contract: contract}, nil
}

// NewIBridgeModuleTransactor creates a new write-only instance of IBridgeModule, bound to a specific deployed contract.
func NewIBridgeModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeModuleTransactor, error) {
	contract, err := bindIBridgeModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeModuleTransactor{contract: contract}, nil
}

// NewIBridgeModuleFilterer creates a new log filterer instance of IBridgeModule, bound to a specific deployed contract.
func NewIBridgeModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeModuleFilterer, error) {
	contract, err := bindIBridgeModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeModuleFilterer{contract: contract}, nil
}

// bindIBridgeModule binds a generic wrapper to an already deployed contract.
func bindIBridgeModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBridgeModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeModule *IBridgeModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeModule.Contract.IBridgeModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeModule *IBridgeModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeModule.Contract.IBridgeModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeModule *IBridgeModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeModule.Contract.IBridgeModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeModule *IBridgeModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeModule *IBridgeModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeModule *IBridgeModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeModule.Contract.contract.Transact(opts, method, params...)
}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_IBridgeModule *IBridgeModuleCaller) CalculateFeeAmount(opts *bind.CallOpts, token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	var out []interface{}
	err := _IBridgeModule.contract.Call(opts, &out, "calculateFeeAmount", token, amount, isSwap)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_IBridgeModule *IBridgeModuleSession) CalculateFeeAmount(token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	return _IBridgeModule.Contract.CalculateFeeAmount(&_IBridgeModule.CallOpts, token, amount, isSwap)
}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_IBridgeModule *IBridgeModuleCallerSession) CalculateFeeAmount(token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	return _IBridgeModule.Contract.CalculateFeeAmount(&_IBridgeModule.CallOpts, token, amount, isSwap)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_IBridgeModule *IBridgeModuleCaller) GetBridgeTokens(opts *bind.CallOpts) ([]BridgeToken, error) {
	var out []interface{}
	err := _IBridgeModule.contract.Call(opts, &out, "getBridgeTokens")

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_IBridgeModule *IBridgeModuleSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _IBridgeModule.Contract.GetBridgeTokens(&_IBridgeModule.CallOpts)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_IBridgeModule *IBridgeModuleCallerSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _IBridgeModule.Contract.GetBridgeTokens(&_IBridgeModule.CallOpts)
}

// GetMaxBridgedAmount is a free data retrieval call binding the contract method 0x04b1ac29.
//
// Solidity: function getMaxBridgedAmount(address token) view returns(uint256 amount)
func (_IBridgeModule *IBridgeModuleCaller) GetMaxBridgedAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBridgeModule.contract.Call(opts, &out, "getMaxBridgedAmount", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxBridgedAmount is a free data retrieval call binding the contract method 0x04b1ac29.
//
// Solidity: function getMaxBridgedAmount(address token) view returns(uint256 amount)
func (_IBridgeModule *IBridgeModuleSession) GetMaxBridgedAmount(token common.Address) (*big.Int, error) {
	return _IBridgeModule.Contract.GetMaxBridgedAmount(&_IBridgeModule.CallOpts, token)
}

// GetMaxBridgedAmount is a free data retrieval call binding the contract method 0x04b1ac29.
//
// Solidity: function getMaxBridgedAmount(address token) view returns(uint256 amount)
func (_IBridgeModule *IBridgeModuleCallerSession) GetMaxBridgedAmount(token common.Address) (*big.Int, error) {
	return _IBridgeModule.Contract.GetMaxBridgedAmount(&_IBridgeModule.CallOpts, token)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string symbol) view returns(address token)
func (_IBridgeModule *IBridgeModuleCaller) SymbolToToken(opts *bind.CallOpts, symbol string) (common.Address, error) {
	var out []interface{}
	err := _IBridgeModule.contract.Call(opts, &out, "symbolToToken", symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string symbol) view returns(address token)
func (_IBridgeModule *IBridgeModuleSession) SymbolToToken(symbol string) (common.Address, error) {
	return _IBridgeModule.Contract.SymbolToToken(&_IBridgeModule.CallOpts, symbol)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string symbol) view returns(address token)
func (_IBridgeModule *IBridgeModuleCallerSession) SymbolToToken(symbol string) (common.Address, error) {
	return _IBridgeModule.Contract.SymbolToToken(&_IBridgeModule.CallOpts, symbol)
}

// TokenToActionMask is a free data retrieval call binding the contract method 0x98b57505.
//
// Solidity: function tokenToActionMask(address token) view returns(uint256 actionMask)
func (_IBridgeModule *IBridgeModuleCaller) TokenToActionMask(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBridgeModule.contract.Call(opts, &out, "tokenToActionMask", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToActionMask is a free data retrieval call binding the contract method 0x98b57505.
//
// Solidity: function tokenToActionMask(address token) view returns(uint256 actionMask)
func (_IBridgeModule *IBridgeModuleSession) TokenToActionMask(token common.Address) (*big.Int, error) {
	return _IBridgeModule.Contract.TokenToActionMask(&_IBridgeModule.CallOpts, token)
}

// TokenToActionMask is a free data retrieval call binding the contract method 0x98b57505.
//
// Solidity: function tokenToActionMask(address token) view returns(uint256 actionMask)
func (_IBridgeModule *IBridgeModuleCallerSession) TokenToActionMask(token common.Address) (*big.Int, error) {
	return _IBridgeModule.Contract.TokenToActionMask(&_IBridgeModule.CallOpts, token)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address token) view returns(string symbol)
func (_IBridgeModule *IBridgeModuleCaller) TokenToSymbol(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _IBridgeModule.contract.Call(opts, &out, "tokenToSymbol", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address token) view returns(string symbol)
func (_IBridgeModule *IBridgeModuleSession) TokenToSymbol(token common.Address) (string, error) {
	return _IBridgeModule.Contract.TokenToSymbol(&_IBridgeModule.CallOpts, token)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address token) view returns(string symbol)
func (_IBridgeModule *IBridgeModuleCallerSession) TokenToSymbol(token common.Address) (string, error) {
	return _IBridgeModule.Contract.TokenToSymbol(&_IBridgeModule.CallOpts, token)
}

// DelegateBridge is a paid mutator transaction binding the contract method 0x436f3aa5.
//
// Solidity: function delegateBridge(address to, uint256 chainId, address token, uint256 amount, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_IBridgeModule *IBridgeModuleTransactor) DelegateBridge(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, destQuery SwapQuery) (*types.Transaction, error) {
	return _IBridgeModule.contract.Transact(opts, "delegateBridge", to, chainId, token, amount, destQuery)
}

// DelegateBridge is a paid mutator transaction binding the contract method 0x436f3aa5.
//
// Solidity: function delegateBridge(address to, uint256 chainId, address token, uint256 amount, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_IBridgeModule *IBridgeModuleSession) DelegateBridge(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, destQuery SwapQuery) (*types.Transaction, error) {
	return _IBridgeModule.Contract.DelegateBridge(&_IBridgeModule.TransactOpts, to, chainId, token, amount, destQuery)
}

// DelegateBridge is a paid mutator transaction binding the contract method 0x436f3aa5.
//
// Solidity: function delegateBridge(address to, uint256 chainId, address token, uint256 amount, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_IBridgeModule *IBridgeModuleTransactorSession) DelegateBridge(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, destQuery SwapQuery) (*types.Transaction, error) {
	return _IBridgeModule.Contract.DelegateBridge(&_IBridgeModule.TransactOpts, to, chainId, token, amount, destQuery)
}

// IDefaultExtendedPoolMetaData contains all meta data concerning the IDefaultExtendedPool contract.
var IDefaultExtendedPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateRemoveLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAPrecise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialATime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureATime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adminFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4d49e87d": "addLiquidity(uint256[],uint256,uint256)",
		"f2fad2b6": "calculateRemoveLiquidity(uint256)",
		"342a87a1": "calculateRemoveLiquidityOneToken(uint256,uint8)",
		"a95b089f": "calculateSwap(uint8,uint8,uint256)",
		"0ba81959": "getAPrecise()",
		"82b86600": "getToken(uint8)",
		"91ceb3eb": "getTokenBalance(uint8)",
		"3e3a1560": "removeLiquidityOneToken(uint256,uint8,uint256,uint256)",
		"91695586": "swap(uint8,uint8,uint256,uint256,uint256)",
		"5fd65f0f": "swapStorage()",
	},
}

// IDefaultExtendedPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IDefaultExtendedPoolMetaData.ABI instead.
var IDefaultExtendedPoolABI = IDefaultExtendedPoolMetaData.ABI

// Deprecated: Use IDefaultExtendedPoolMetaData.Sigs instead.
// IDefaultExtendedPoolFuncSigs maps the 4-byte function signature to its string representation.
var IDefaultExtendedPoolFuncSigs = IDefaultExtendedPoolMetaData.Sigs

// IDefaultExtendedPool is an auto generated Go binding around an Ethereum contract.
type IDefaultExtendedPool struct {
	IDefaultExtendedPoolCaller     // Read-only binding to the contract
	IDefaultExtendedPoolTransactor // Write-only binding to the contract
	IDefaultExtendedPoolFilterer   // Log filterer for contract events
}

// IDefaultExtendedPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDefaultExtendedPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDefaultExtendedPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDefaultExtendedPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDefaultExtendedPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDefaultExtendedPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDefaultExtendedPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDefaultExtendedPoolSession struct {
	Contract     *IDefaultExtendedPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IDefaultExtendedPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDefaultExtendedPoolCallerSession struct {
	Contract *IDefaultExtendedPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IDefaultExtendedPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDefaultExtendedPoolTransactorSession struct {
	Contract     *IDefaultExtendedPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IDefaultExtendedPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDefaultExtendedPoolRaw struct {
	Contract *IDefaultExtendedPool // Generic contract binding to access the raw methods on
}

// IDefaultExtendedPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDefaultExtendedPoolCallerRaw struct {
	Contract *IDefaultExtendedPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IDefaultExtendedPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDefaultExtendedPoolTransactorRaw struct {
	Contract *IDefaultExtendedPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDefaultExtendedPool creates a new instance of IDefaultExtendedPool, bound to a specific deployed contract.
func NewIDefaultExtendedPool(address common.Address, backend bind.ContractBackend) (*IDefaultExtendedPool, error) {
	contract, err := bindIDefaultExtendedPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDefaultExtendedPool{IDefaultExtendedPoolCaller: IDefaultExtendedPoolCaller{contract: contract}, IDefaultExtendedPoolTransactor: IDefaultExtendedPoolTransactor{contract: contract}, IDefaultExtendedPoolFilterer: IDefaultExtendedPoolFilterer{contract: contract}}, nil
}

// NewIDefaultExtendedPoolCaller creates a new read-only instance of IDefaultExtendedPool, bound to a specific deployed contract.
func NewIDefaultExtendedPoolCaller(address common.Address, caller bind.ContractCaller) (*IDefaultExtendedPoolCaller, error) {
	contract, err := bindIDefaultExtendedPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDefaultExtendedPoolCaller{contract: contract}, nil
}

// NewIDefaultExtendedPoolTransactor creates a new write-only instance of IDefaultExtendedPool, bound to a specific deployed contract.
func NewIDefaultExtendedPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IDefaultExtendedPoolTransactor, error) {
	contract, err := bindIDefaultExtendedPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDefaultExtendedPoolTransactor{contract: contract}, nil
}

// NewIDefaultExtendedPoolFilterer creates a new log filterer instance of IDefaultExtendedPool, bound to a specific deployed contract.
func NewIDefaultExtendedPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IDefaultExtendedPoolFilterer, error) {
	contract, err := bindIDefaultExtendedPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDefaultExtendedPoolFilterer{contract: contract}, nil
}

// bindIDefaultExtendedPool binds a generic wrapper to an already deployed contract.
func bindIDefaultExtendedPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDefaultExtendedPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDefaultExtendedPool *IDefaultExtendedPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDefaultExtendedPool.Contract.IDefaultExtendedPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDefaultExtendedPool *IDefaultExtendedPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.IDefaultExtendedPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDefaultExtendedPool *IDefaultExtendedPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.IDefaultExtendedPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDefaultExtendedPool *IDefaultExtendedPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDefaultExtendedPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDefaultExtendedPool *IDefaultExtendedPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDefaultExtendedPool *IDefaultExtendedPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.contract.Transact(opts, method, params...)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_IDefaultExtendedPool *IDefaultExtendedPoolCaller) CalculateRemoveLiquidity(opts *bind.CallOpts, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IDefaultExtendedPool.contract.Call(opts, &out, "calculateRemoveLiquidity", amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _IDefaultExtendedPool.Contract.CalculateRemoveLiquidity(&_IDefaultExtendedPool.CallOpts, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_IDefaultExtendedPool *IDefaultExtendedPoolCallerSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _IDefaultExtendedPool.Contract.CalculateRemoveLiquidity(&_IDefaultExtendedPool.CallOpts, amount)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCaller) CalculateRemoveLiquidityOneToken(opts *bind.CallOpts, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _IDefaultExtendedPool.contract.Call(opts, &out, "calculateRemoveLiquidityOneToken", tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _IDefaultExtendedPool.Contract.CalculateRemoveLiquidityOneToken(&_IDefaultExtendedPool.CallOpts, tokenAmount, tokenIndex)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCallerSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _IDefaultExtendedPool.Contract.CalculateRemoveLiquidityOneToken(&_IDefaultExtendedPool.CallOpts, tokenAmount, tokenIndex)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCaller) CalculateSwap(opts *bind.CallOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDefaultExtendedPool.contract.Call(opts, &out, "calculateSwap", tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _IDefaultExtendedPool.Contract.CalculateSwap(&_IDefaultExtendedPool.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCallerSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _IDefaultExtendedPool.Contract.CalculateSwap(&_IDefaultExtendedPool.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCaller) GetAPrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDefaultExtendedPool.contract.Call(opts, &out, "getAPrecise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) GetAPrecise() (*big.Int, error) {
	return _IDefaultExtendedPool.Contract.GetAPrecise(&_IDefaultExtendedPool.CallOpts)
}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCallerSession) GetAPrecise() (*big.Int, error) {
	return _IDefaultExtendedPool.Contract.GetAPrecise(&_IDefaultExtendedPool.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address token)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCaller) GetToken(opts *bind.CallOpts, index uint8) (common.Address, error) {
	var out []interface{}
	err := _IDefaultExtendedPool.contract.Call(opts, &out, "getToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address token)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) GetToken(index uint8) (common.Address, error) {
	return _IDefaultExtendedPool.Contract.GetToken(&_IDefaultExtendedPool.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address token)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCallerSession) GetToken(index uint8) (common.Address, error) {
	return _IDefaultExtendedPool.Contract.GetToken(&_IDefaultExtendedPool.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCaller) GetTokenBalance(opts *bind.CallOpts, index uint8) (*big.Int, error) {
	var out []interface{}
	err := _IDefaultExtendedPool.contract.Call(opts, &out, "getTokenBalance", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _IDefaultExtendedPool.Contract.GetTokenBalance(&_IDefaultExtendedPool.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCallerSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _IDefaultExtendedPool.Contract.GetTokenBalance(&_IDefaultExtendedPool.CallOpts, index)
}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCaller) SwapStorage(opts *bind.CallOpts) (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	var out []interface{}
	err := _IDefaultExtendedPool.contract.Call(opts, &out, "swapStorage")

	outstruct := new(struct {
		InitialA     *big.Int
		FutureA      *big.Int
		InitialATime *big.Int
		FutureATime  *big.Int
		SwapFee      *big.Int
		AdminFee     *big.Int
		LpToken      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FutureA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitialATime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FutureATime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SwapFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AdminFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LpToken = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) SwapStorage() (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	return _IDefaultExtendedPool.Contract.SwapStorage(&_IDefaultExtendedPool.CallOpts)
}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_IDefaultExtendedPool *IDefaultExtendedPoolCallerSession) SwapStorage() (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	return _IDefaultExtendedPool.Contract.SwapStorage(&_IDefaultExtendedPool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolTransactor) AddLiquidity(opts *bind.TransactOpts, amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.contract.Transact(opts, "addLiquidity", amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.AddLiquidity(&_IDefaultExtendedPool.TransactOpts, amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolTransactorSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.AddLiquidity(&_IDefaultExtendedPool.TransactOpts, amounts, minToMint, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolTransactor) RemoveLiquidityOneToken(opts *bind.TransactOpts, tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.contract.Transact(opts, "removeLiquidityOneToken", tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.RemoveLiquidityOneToken(&_IDefaultExtendedPool.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_IDefaultExtendedPool *IDefaultExtendedPoolTransactorSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.RemoveLiquidityOneToken(&_IDefaultExtendedPool.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256 amountOut)
func (_IDefaultExtendedPool *IDefaultExtendedPoolTransactor) Swap(opts *bind.TransactOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.contract.Transact(opts, "swap", tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256 amountOut)
func (_IDefaultExtendedPool *IDefaultExtendedPoolSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.Swap(&_IDefaultExtendedPool.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256 amountOut)
func (_IDefaultExtendedPool *IDefaultExtendedPoolTransactorSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultExtendedPool.Contract.Swap(&_IDefaultExtendedPool.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// IDefaultPoolMetaData contains all meta data concerning the IDefaultPool contract.
var IDefaultPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a95b089f": "calculateSwap(uint8,uint8,uint256)",
		"82b86600": "getToken(uint8)",
		"91695586": "swap(uint8,uint8,uint256,uint256,uint256)",
	},
}

// IDefaultPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IDefaultPoolMetaData.ABI instead.
var IDefaultPoolABI = IDefaultPoolMetaData.ABI

// Deprecated: Use IDefaultPoolMetaData.Sigs instead.
// IDefaultPoolFuncSigs maps the 4-byte function signature to its string representation.
var IDefaultPoolFuncSigs = IDefaultPoolMetaData.Sigs

// IDefaultPool is an auto generated Go binding around an Ethereum contract.
type IDefaultPool struct {
	IDefaultPoolCaller     // Read-only binding to the contract
	IDefaultPoolTransactor // Write-only binding to the contract
	IDefaultPoolFilterer   // Log filterer for contract events
}

// IDefaultPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDefaultPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDefaultPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDefaultPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDefaultPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDefaultPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDefaultPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDefaultPoolSession struct {
	Contract     *IDefaultPool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDefaultPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDefaultPoolCallerSession struct {
	Contract *IDefaultPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IDefaultPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDefaultPoolTransactorSession struct {
	Contract     *IDefaultPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IDefaultPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDefaultPoolRaw struct {
	Contract *IDefaultPool // Generic contract binding to access the raw methods on
}

// IDefaultPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDefaultPoolCallerRaw struct {
	Contract *IDefaultPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IDefaultPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDefaultPoolTransactorRaw struct {
	Contract *IDefaultPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDefaultPool creates a new instance of IDefaultPool, bound to a specific deployed contract.
func NewIDefaultPool(address common.Address, backend bind.ContractBackend) (*IDefaultPool, error) {
	contract, err := bindIDefaultPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDefaultPool{IDefaultPoolCaller: IDefaultPoolCaller{contract: contract}, IDefaultPoolTransactor: IDefaultPoolTransactor{contract: contract}, IDefaultPoolFilterer: IDefaultPoolFilterer{contract: contract}}, nil
}

// NewIDefaultPoolCaller creates a new read-only instance of IDefaultPool, bound to a specific deployed contract.
func NewIDefaultPoolCaller(address common.Address, caller bind.ContractCaller) (*IDefaultPoolCaller, error) {
	contract, err := bindIDefaultPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDefaultPoolCaller{contract: contract}, nil
}

// NewIDefaultPoolTransactor creates a new write-only instance of IDefaultPool, bound to a specific deployed contract.
func NewIDefaultPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IDefaultPoolTransactor, error) {
	contract, err := bindIDefaultPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDefaultPoolTransactor{contract: contract}, nil
}

// NewIDefaultPoolFilterer creates a new log filterer instance of IDefaultPool, bound to a specific deployed contract.
func NewIDefaultPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IDefaultPoolFilterer, error) {
	contract, err := bindIDefaultPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDefaultPoolFilterer{contract: contract}, nil
}

// bindIDefaultPool binds a generic wrapper to an already deployed contract.
func bindIDefaultPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDefaultPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDefaultPool *IDefaultPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDefaultPool.Contract.IDefaultPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDefaultPool *IDefaultPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDefaultPool.Contract.IDefaultPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDefaultPool *IDefaultPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDefaultPool.Contract.IDefaultPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDefaultPool *IDefaultPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDefaultPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDefaultPool *IDefaultPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDefaultPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDefaultPool *IDefaultPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDefaultPool.Contract.contract.Transact(opts, method, params...)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_IDefaultPool *IDefaultPoolCaller) CalculateSwap(opts *bind.CallOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDefaultPool.contract.Call(opts, &out, "calculateSwap", tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_IDefaultPool *IDefaultPoolSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _IDefaultPool.Contract.CalculateSwap(&_IDefaultPool.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_IDefaultPool *IDefaultPoolCallerSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _IDefaultPool.Contract.CalculateSwap(&_IDefaultPool.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address token)
func (_IDefaultPool *IDefaultPoolCaller) GetToken(opts *bind.CallOpts, index uint8) (common.Address, error) {
	var out []interface{}
	err := _IDefaultPool.contract.Call(opts, &out, "getToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address token)
func (_IDefaultPool *IDefaultPoolSession) GetToken(index uint8) (common.Address, error) {
	return _IDefaultPool.Contract.GetToken(&_IDefaultPool.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address token)
func (_IDefaultPool *IDefaultPoolCallerSession) GetToken(index uint8) (common.Address, error) {
	return _IDefaultPool.Contract.GetToken(&_IDefaultPool.CallOpts, index)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256 amountOut)
func (_IDefaultPool *IDefaultPoolTransactor) Swap(opts *bind.TransactOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultPool.contract.Transact(opts, "swap", tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256 amountOut)
func (_IDefaultPool *IDefaultPoolSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultPool.Contract.Swap(&_IDefaultPool.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256 amountOut)
func (_IDefaultPool *IDefaultPoolTransactorSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IDefaultPool.Contract.Swap(&_IDefaultPool.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRouterAdapterMetaData contains all meta data concerning the IRouterAdapter contract.
var IRouterAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"name\":\"adapterSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"24a98f11": "adapterSwap(address,address,uint256,address,bytes)",
	},
}

// IRouterAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use IRouterAdapterMetaData.ABI instead.
var IRouterAdapterABI = IRouterAdapterMetaData.ABI

// Deprecated: Use IRouterAdapterMetaData.Sigs instead.
// IRouterAdapterFuncSigs maps the 4-byte function signature to its string representation.
var IRouterAdapterFuncSigs = IRouterAdapterMetaData.Sigs

// IRouterAdapter is an auto generated Go binding around an Ethereum contract.
type IRouterAdapter struct {
	IRouterAdapterCaller     // Read-only binding to the contract
	IRouterAdapterTransactor // Write-only binding to the contract
	IRouterAdapterFilterer   // Log filterer for contract events
}

// IRouterAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRouterAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRouterAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRouterAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRouterAdapterSession struct {
	Contract     *IRouterAdapter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRouterAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRouterAdapterCallerSession struct {
	Contract *IRouterAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IRouterAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRouterAdapterTransactorSession struct {
	Contract     *IRouterAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IRouterAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRouterAdapterRaw struct {
	Contract *IRouterAdapter // Generic contract binding to access the raw methods on
}

// IRouterAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRouterAdapterCallerRaw struct {
	Contract *IRouterAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// IRouterAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRouterAdapterTransactorRaw struct {
	Contract *IRouterAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRouterAdapter creates a new instance of IRouterAdapter, bound to a specific deployed contract.
func NewIRouterAdapter(address common.Address, backend bind.ContractBackend) (*IRouterAdapter, error) {
	contract, err := bindIRouterAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRouterAdapter{IRouterAdapterCaller: IRouterAdapterCaller{contract: contract}, IRouterAdapterTransactor: IRouterAdapterTransactor{contract: contract}, IRouterAdapterFilterer: IRouterAdapterFilterer{contract: contract}}, nil
}

// NewIRouterAdapterCaller creates a new read-only instance of IRouterAdapter, bound to a specific deployed contract.
func NewIRouterAdapterCaller(address common.Address, caller bind.ContractCaller) (*IRouterAdapterCaller, error) {
	contract, err := bindIRouterAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRouterAdapterCaller{contract: contract}, nil
}

// NewIRouterAdapterTransactor creates a new write-only instance of IRouterAdapter, bound to a specific deployed contract.
func NewIRouterAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*IRouterAdapterTransactor, error) {
	contract, err := bindIRouterAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRouterAdapterTransactor{contract: contract}, nil
}

// NewIRouterAdapterFilterer creates a new log filterer instance of IRouterAdapter, bound to a specific deployed contract.
func NewIRouterAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*IRouterAdapterFilterer, error) {
	contract, err := bindIRouterAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRouterAdapterFilterer{contract: contract}, nil
}

// bindIRouterAdapter binds a generic wrapper to an already deployed contract.
func bindIRouterAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRouterAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRouterAdapter *IRouterAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRouterAdapter.Contract.IRouterAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRouterAdapter *IRouterAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRouterAdapter.Contract.IRouterAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRouterAdapter *IRouterAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRouterAdapter.Contract.IRouterAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRouterAdapter *IRouterAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRouterAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRouterAdapter *IRouterAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRouterAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRouterAdapter *IRouterAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRouterAdapter.Contract.contract.Transact(opts, method, params...)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_IRouterAdapter *IRouterAdapterTransactor) AdapterSwap(opts *bind.TransactOpts, recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _IRouterAdapter.contract.Transact(opts, "adapterSwap", recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_IRouterAdapter *IRouterAdapterSession) AdapterSwap(recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _IRouterAdapter.Contract.AdapterSwap(&_IRouterAdapter.TransactOpts, recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_IRouterAdapter *IRouterAdapterTransactorSession) AdapterSwap(recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _IRouterAdapter.Contract.AdapterSwap(&_IRouterAdapter.TransactOpts, recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// IRouterV2MetaData contains all meta data concerning the IRouterV2 contract.
var IRouterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"originQuery\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"destQuery\",\"type\":\"tuple\"}],\"name\":\"bridgeViaSynapse\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"name\":\"connectBridgeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"}],\"name\":\"disconnectBridgeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"bridgeTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structDestRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getDestinationAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"destQuery\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getDestinationBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"destTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getOriginAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"originQuery\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"}],\"name\":\"getOriginBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"originTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"supportedTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"}],\"name\":\"idToModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"name\":\"moduleToId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISwapQuoterV2\",\"name\":\"_swapQuoter\",\"type\":\"address\"}],\"name\":\"setSwapQuoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"query\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"name\":\"updateBridgeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c95fafd2": "bridgeViaSynapse(address,uint256,bytes32,address,uint256,(address,address,uint256,uint256,bytes),(address,address,uint256,uint256,bytes))",
		"b3bce952": "connectBridgeModule(bytes32,address)",
		"b68e4302": "disconnectBridgeModule(bytes32)",
		"9c1d060e": "getBridgeTokens()",
		"7de31c74": "getDestinationAmountOut((string,uint256),address)",
		"1d04879b": "getDestinationBridgeTokens(address)",
		"f533941d": "getOriginAmountOut(address,string,uint256)",
		"3e811a5c": "getOriginBridgeTokens(address)",
		"d3c7c2c7": "getSupportedTokens()",
		"53e2e8e7": "idToModule(bytes32)",
		"9f2671fa": "moduleToId(address)",
		"da46098c": "setAllowance(address,address,uint256)",
		"804b3dff": "setSwapQuoter(address)",
		"b5d1cdd4": "swap(address,address,uint256,(address,address,uint256,uint256,bytes))",
		"70a1cdc9": "updateBridgeModule(bytes32,address)",
	},
}

// IRouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use IRouterV2MetaData.ABI instead.
var IRouterV2ABI = IRouterV2MetaData.ABI

// Deprecated: Use IRouterV2MetaData.Sigs instead.
// IRouterV2FuncSigs maps the 4-byte function signature to its string representation.
var IRouterV2FuncSigs = IRouterV2MetaData.Sigs

// IRouterV2 is an auto generated Go binding around an Ethereum contract.
type IRouterV2 struct {
	IRouterV2Caller     // Read-only binding to the contract
	IRouterV2Transactor // Write-only binding to the contract
	IRouterV2Filterer   // Log filterer for contract events
}

// IRouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IRouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IRouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRouterV2Session struct {
	Contract     *IRouterV2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRouterV2CallerSession struct {
	Contract *IRouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IRouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRouterV2TransactorSession struct {
	Contract     *IRouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IRouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IRouterV2Raw struct {
	Contract *IRouterV2 // Generic contract binding to access the raw methods on
}

// IRouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRouterV2CallerRaw struct {
	Contract *IRouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// IRouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRouterV2TransactorRaw struct {
	Contract *IRouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIRouterV2 creates a new instance of IRouterV2, bound to a specific deployed contract.
func NewIRouterV2(address common.Address, backend bind.ContractBackend) (*IRouterV2, error) {
	contract, err := bindIRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRouterV2{IRouterV2Caller: IRouterV2Caller{contract: contract}, IRouterV2Transactor: IRouterV2Transactor{contract: contract}, IRouterV2Filterer: IRouterV2Filterer{contract: contract}}, nil
}

// NewIRouterV2Caller creates a new read-only instance of IRouterV2, bound to a specific deployed contract.
func NewIRouterV2Caller(address common.Address, caller bind.ContractCaller) (*IRouterV2Caller, error) {
	contract, err := bindIRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRouterV2Caller{contract: contract}, nil
}

// NewIRouterV2Transactor creates a new write-only instance of IRouterV2, bound to a specific deployed contract.
func NewIRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*IRouterV2Transactor, error) {
	contract, err := bindIRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRouterV2Transactor{contract: contract}, nil
}

// NewIRouterV2Filterer creates a new log filterer instance of IRouterV2, bound to a specific deployed contract.
func NewIRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*IRouterV2Filterer, error) {
	contract, err := bindIRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRouterV2Filterer{contract: contract}, nil
}

// bindIRouterV2 binds a generic wrapper to an already deployed contract.
func bindIRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRouterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRouterV2 *IRouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRouterV2.Contract.IRouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRouterV2 *IRouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRouterV2.Contract.IRouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRouterV2 *IRouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRouterV2.Contract.IRouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRouterV2 *IRouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRouterV2 *IRouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRouterV2 *IRouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRouterV2.Contract.contract.Transact(opts, method, params...)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_IRouterV2 *IRouterV2Caller) GetBridgeTokens(opts *bind.CallOpts) ([]BridgeToken, error) {
	var out []interface{}
	err := _IRouterV2.contract.Call(opts, &out, "getBridgeTokens")

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_IRouterV2 *IRouterV2Session) GetBridgeTokens() ([]BridgeToken, error) {
	return _IRouterV2.Contract.GetBridgeTokens(&_IRouterV2.CallOpts)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_IRouterV2 *IRouterV2CallerSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _IRouterV2.Contract.GetBridgeTokens(&_IRouterV2.CallOpts)
}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x7de31c74.
//
// Solidity: function getDestinationAmountOut((string,uint256) request, address tokenOut) view returns((address,address,uint256,uint256,bytes) destQuery)
func (_IRouterV2 *IRouterV2Caller) GetDestinationAmountOut(opts *bind.CallOpts, request DestRequest, tokenOut common.Address) (SwapQuery, error) {
	var out []interface{}
	err := _IRouterV2.contract.Call(opts, &out, "getDestinationAmountOut", request, tokenOut)

	if err != nil {
		return *new(SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new(SwapQuery)).(*SwapQuery)

	return out0, err

}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x7de31c74.
//
// Solidity: function getDestinationAmountOut((string,uint256) request, address tokenOut) view returns((address,address,uint256,uint256,bytes) destQuery)
func (_IRouterV2 *IRouterV2Session) GetDestinationAmountOut(request DestRequest, tokenOut common.Address) (SwapQuery, error) {
	return _IRouterV2.Contract.GetDestinationAmountOut(&_IRouterV2.CallOpts, request, tokenOut)
}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x7de31c74.
//
// Solidity: function getDestinationAmountOut((string,uint256) request, address tokenOut) view returns((address,address,uint256,uint256,bytes) destQuery)
func (_IRouterV2 *IRouterV2CallerSession) GetDestinationAmountOut(request DestRequest, tokenOut common.Address) (SwapQuery, error) {
	return _IRouterV2.Contract.GetDestinationAmountOut(&_IRouterV2.CallOpts, request, tokenOut)
}

// GetDestinationBridgeTokens is a free data retrieval call binding the contract method 0x1d04879b.
//
// Solidity: function getDestinationBridgeTokens(address tokenOut) view returns((string,address)[] destTokens)
func (_IRouterV2 *IRouterV2Caller) GetDestinationBridgeTokens(opts *bind.CallOpts, tokenOut common.Address) ([]BridgeToken, error) {
	var out []interface{}
	err := _IRouterV2.contract.Call(opts, &out, "getDestinationBridgeTokens", tokenOut)

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetDestinationBridgeTokens is a free data retrieval call binding the contract method 0x1d04879b.
//
// Solidity: function getDestinationBridgeTokens(address tokenOut) view returns((string,address)[] destTokens)
func (_IRouterV2 *IRouterV2Session) GetDestinationBridgeTokens(tokenOut common.Address) ([]BridgeToken, error) {
	return _IRouterV2.Contract.GetDestinationBridgeTokens(&_IRouterV2.CallOpts, tokenOut)
}

// GetDestinationBridgeTokens is a free data retrieval call binding the contract method 0x1d04879b.
//
// Solidity: function getDestinationBridgeTokens(address tokenOut) view returns((string,address)[] destTokens)
func (_IRouterV2 *IRouterV2CallerSession) GetDestinationBridgeTokens(tokenOut common.Address) ([]BridgeToken, error) {
	return _IRouterV2.Contract.GetDestinationBridgeTokens(&_IRouterV2.CallOpts, tokenOut)
}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0xf533941d.
//
// Solidity: function getOriginAmountOut(address tokenIn, string tokenSymbol, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) originQuery)
func (_IRouterV2 *IRouterV2Caller) GetOriginAmountOut(opts *bind.CallOpts, tokenIn common.Address, tokenSymbol string, amountIn *big.Int) (SwapQuery, error) {
	var out []interface{}
	err := _IRouterV2.contract.Call(opts, &out, "getOriginAmountOut", tokenIn, tokenSymbol, amountIn)

	if err != nil {
		return *new(SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new(SwapQuery)).(*SwapQuery)

	return out0, err

}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0xf533941d.
//
// Solidity: function getOriginAmountOut(address tokenIn, string tokenSymbol, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) originQuery)
func (_IRouterV2 *IRouterV2Session) GetOriginAmountOut(tokenIn common.Address, tokenSymbol string, amountIn *big.Int) (SwapQuery, error) {
	return _IRouterV2.Contract.GetOriginAmountOut(&_IRouterV2.CallOpts, tokenIn, tokenSymbol, amountIn)
}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0xf533941d.
//
// Solidity: function getOriginAmountOut(address tokenIn, string tokenSymbol, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) originQuery)
func (_IRouterV2 *IRouterV2CallerSession) GetOriginAmountOut(tokenIn common.Address, tokenSymbol string, amountIn *big.Int) (SwapQuery, error) {
	return _IRouterV2.Contract.GetOriginAmountOut(&_IRouterV2.CallOpts, tokenIn, tokenSymbol, amountIn)
}

// GetOriginBridgeTokens is a free data retrieval call binding the contract method 0x3e811a5c.
//
// Solidity: function getOriginBridgeTokens(address tokenIn) view returns((string,address)[] originTokens)
func (_IRouterV2 *IRouterV2Caller) GetOriginBridgeTokens(opts *bind.CallOpts, tokenIn common.Address) ([]BridgeToken, error) {
	var out []interface{}
	err := _IRouterV2.contract.Call(opts, &out, "getOriginBridgeTokens", tokenIn)

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetOriginBridgeTokens is a free data retrieval call binding the contract method 0x3e811a5c.
//
// Solidity: function getOriginBridgeTokens(address tokenIn) view returns((string,address)[] originTokens)
func (_IRouterV2 *IRouterV2Session) GetOriginBridgeTokens(tokenIn common.Address) ([]BridgeToken, error) {
	return _IRouterV2.Contract.GetOriginBridgeTokens(&_IRouterV2.CallOpts, tokenIn)
}

// GetOriginBridgeTokens is a free data retrieval call binding the contract method 0x3e811a5c.
//
// Solidity: function getOriginBridgeTokens(address tokenIn) view returns((string,address)[] originTokens)
func (_IRouterV2 *IRouterV2CallerSession) GetOriginBridgeTokens(tokenIn common.Address) ([]BridgeToken, error) {
	return _IRouterV2.Contract.GetOriginBridgeTokens(&_IRouterV2.CallOpts, tokenIn)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] supportedTokens)
func (_IRouterV2 *IRouterV2Caller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IRouterV2.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] supportedTokens)
func (_IRouterV2 *IRouterV2Session) GetSupportedTokens() ([]common.Address, error) {
	return _IRouterV2.Contract.GetSupportedTokens(&_IRouterV2.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] supportedTokens)
func (_IRouterV2 *IRouterV2CallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _IRouterV2.Contract.GetSupportedTokens(&_IRouterV2.CallOpts)
}

// IdToModule is a free data retrieval call binding the contract method 0x53e2e8e7.
//
// Solidity: function idToModule(bytes32 moduleId) view returns(address bridgeModule)
func (_IRouterV2 *IRouterV2Caller) IdToModule(opts *bind.CallOpts, moduleId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IRouterV2.contract.Call(opts, &out, "idToModule", moduleId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdToModule is a free data retrieval call binding the contract method 0x53e2e8e7.
//
// Solidity: function idToModule(bytes32 moduleId) view returns(address bridgeModule)
func (_IRouterV2 *IRouterV2Session) IdToModule(moduleId [32]byte) (common.Address, error) {
	return _IRouterV2.Contract.IdToModule(&_IRouterV2.CallOpts, moduleId)
}

// IdToModule is a free data retrieval call binding the contract method 0x53e2e8e7.
//
// Solidity: function idToModule(bytes32 moduleId) view returns(address bridgeModule)
func (_IRouterV2 *IRouterV2CallerSession) IdToModule(moduleId [32]byte) (common.Address, error) {
	return _IRouterV2.Contract.IdToModule(&_IRouterV2.CallOpts, moduleId)
}

// ModuleToId is a free data retrieval call binding the contract method 0x9f2671fa.
//
// Solidity: function moduleToId(address bridgeModule) view returns(bytes32 moduleId)
func (_IRouterV2 *IRouterV2Caller) ModuleToId(opts *bind.CallOpts, bridgeModule common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IRouterV2.contract.Call(opts, &out, "moduleToId", bridgeModule)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ModuleToId is a free data retrieval call binding the contract method 0x9f2671fa.
//
// Solidity: function moduleToId(address bridgeModule) view returns(bytes32 moduleId)
func (_IRouterV2 *IRouterV2Session) ModuleToId(bridgeModule common.Address) ([32]byte, error) {
	return _IRouterV2.Contract.ModuleToId(&_IRouterV2.CallOpts, bridgeModule)
}

// ModuleToId is a free data retrieval call binding the contract method 0x9f2671fa.
//
// Solidity: function moduleToId(address bridgeModule) view returns(bytes32 moduleId)
func (_IRouterV2 *IRouterV2CallerSession) ModuleToId(bridgeModule common.Address) ([32]byte, error) {
	return _IRouterV2.Contract.ModuleToId(&_IRouterV2.CallOpts, bridgeModule)
}

// BridgeViaSynapse is a paid mutator transaction binding the contract method 0xc95fafd2.
//
// Solidity: function bridgeViaSynapse(address to, uint256 chainId, bytes32 moduleId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_IRouterV2 *IRouterV2Transactor) BridgeViaSynapse(opts *bind.TransactOpts, to common.Address, chainId *big.Int, moduleId [32]byte, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _IRouterV2.contract.Transact(opts, "bridgeViaSynapse", to, chainId, moduleId, token, amount, originQuery, destQuery)
}

// BridgeViaSynapse is a paid mutator transaction binding the contract method 0xc95fafd2.
//
// Solidity: function bridgeViaSynapse(address to, uint256 chainId, bytes32 moduleId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_IRouterV2 *IRouterV2Session) BridgeViaSynapse(to common.Address, chainId *big.Int, moduleId [32]byte, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _IRouterV2.Contract.BridgeViaSynapse(&_IRouterV2.TransactOpts, to, chainId, moduleId, token, amount, originQuery, destQuery)
}

// BridgeViaSynapse is a paid mutator transaction binding the contract method 0xc95fafd2.
//
// Solidity: function bridgeViaSynapse(address to, uint256 chainId, bytes32 moduleId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_IRouterV2 *IRouterV2TransactorSession) BridgeViaSynapse(to common.Address, chainId *big.Int, moduleId [32]byte, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _IRouterV2.Contract.BridgeViaSynapse(&_IRouterV2.TransactOpts, to, chainId, moduleId, token, amount, originQuery, destQuery)
}

// ConnectBridgeModule is a paid mutator transaction binding the contract method 0xb3bce952.
//
// Solidity: function connectBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_IRouterV2 *IRouterV2Transactor) ConnectBridgeModule(opts *bind.TransactOpts, moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _IRouterV2.contract.Transact(opts, "connectBridgeModule", moduleId, bridgeModule)
}

// ConnectBridgeModule is a paid mutator transaction binding the contract method 0xb3bce952.
//
// Solidity: function connectBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_IRouterV2 *IRouterV2Session) ConnectBridgeModule(moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _IRouterV2.Contract.ConnectBridgeModule(&_IRouterV2.TransactOpts, moduleId, bridgeModule)
}

// ConnectBridgeModule is a paid mutator transaction binding the contract method 0xb3bce952.
//
// Solidity: function connectBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_IRouterV2 *IRouterV2TransactorSession) ConnectBridgeModule(moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _IRouterV2.Contract.ConnectBridgeModule(&_IRouterV2.TransactOpts, moduleId, bridgeModule)
}

// DisconnectBridgeModule is a paid mutator transaction binding the contract method 0xb68e4302.
//
// Solidity: function disconnectBridgeModule(bytes32 moduleId) returns()
func (_IRouterV2 *IRouterV2Transactor) DisconnectBridgeModule(opts *bind.TransactOpts, moduleId [32]byte) (*types.Transaction, error) {
	return _IRouterV2.contract.Transact(opts, "disconnectBridgeModule", moduleId)
}

// DisconnectBridgeModule is a paid mutator transaction binding the contract method 0xb68e4302.
//
// Solidity: function disconnectBridgeModule(bytes32 moduleId) returns()
func (_IRouterV2 *IRouterV2Session) DisconnectBridgeModule(moduleId [32]byte) (*types.Transaction, error) {
	return _IRouterV2.Contract.DisconnectBridgeModule(&_IRouterV2.TransactOpts, moduleId)
}

// DisconnectBridgeModule is a paid mutator transaction binding the contract method 0xb68e4302.
//
// Solidity: function disconnectBridgeModule(bytes32 moduleId) returns()
func (_IRouterV2 *IRouterV2TransactorSession) DisconnectBridgeModule(moduleId [32]byte) (*types.Transaction, error) {
	return _IRouterV2.Contract.DisconnectBridgeModule(&_IRouterV2.TransactOpts, moduleId)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_IRouterV2 *IRouterV2Transactor) SetAllowance(opts *bind.TransactOpts, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IRouterV2.contract.Transact(opts, "setAllowance", token, spender, amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_IRouterV2 *IRouterV2Session) SetAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IRouterV2.Contract.SetAllowance(&_IRouterV2.TransactOpts, token, spender, amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_IRouterV2 *IRouterV2TransactorSession) SetAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IRouterV2.Contract.SetAllowance(&_IRouterV2.TransactOpts, token, spender, amount)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_IRouterV2 *IRouterV2Transactor) SetSwapQuoter(opts *bind.TransactOpts, _swapQuoter common.Address) (*types.Transaction, error) {
	return _IRouterV2.contract.Transact(opts, "setSwapQuoter", _swapQuoter)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_IRouterV2 *IRouterV2Session) SetSwapQuoter(_swapQuoter common.Address) (*types.Transaction, error) {
	return _IRouterV2.Contract.SetSwapQuoter(&_IRouterV2.TransactOpts, _swapQuoter)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_IRouterV2 *IRouterV2TransactorSession) SetSwapQuoter(_swapQuoter common.Address) (*types.Transaction, error) {
	return _IRouterV2.Contract.SetSwapQuoter(&_IRouterV2.TransactOpts, _swapQuoter)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_IRouterV2 *IRouterV2Transactor) Swap(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _IRouterV2.contract.Transact(opts, "swap", to, token, amount, query)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_IRouterV2 *IRouterV2Session) Swap(to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _IRouterV2.Contract.Swap(&_IRouterV2.TransactOpts, to, token, amount, query)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_IRouterV2 *IRouterV2TransactorSession) Swap(to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _IRouterV2.Contract.Swap(&_IRouterV2.TransactOpts, to, token, amount, query)
}

// UpdateBridgeModule is a paid mutator transaction binding the contract method 0x70a1cdc9.
//
// Solidity: function updateBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_IRouterV2 *IRouterV2Transactor) UpdateBridgeModule(opts *bind.TransactOpts, moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _IRouterV2.contract.Transact(opts, "updateBridgeModule", moduleId, bridgeModule)
}

// UpdateBridgeModule is a paid mutator transaction binding the contract method 0x70a1cdc9.
//
// Solidity: function updateBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_IRouterV2 *IRouterV2Session) UpdateBridgeModule(moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _IRouterV2.Contract.UpdateBridgeModule(&_IRouterV2.TransactOpts, moduleId, bridgeModule)
}

// UpdateBridgeModule is a paid mutator transaction binding the contract method 0x70a1cdc9.
//
// Solidity: function updateBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_IRouterV2 *IRouterV2TransactorSession) UpdateBridgeModule(moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _IRouterV2.Contract.UpdateBridgeModule(&_IRouterV2.TransactOpts, moduleId, bridgeModule)
}

// ISwapQuoterV1MetaData contains all meta data concerning the ISwapQuoterV1 contract.
var ISwapQuoterV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"allPools\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWeth\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPoolToken[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"internalType\":\"structPool[]\",\"name\":\"pools\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateWithdrawOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultPoolCalc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"actionMask\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structLimitedToken[]\",\"name\":\"bridgeTokensIn\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"findConnectedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountFound\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"isConnected\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"actionMask\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structLimitedToken\",\"name\":\"tokenIn\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"query\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"poolTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWeth\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPoolToken[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amtPools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c5c63e65": "allPools()",
		"4d864496": "calculateAddLiquidity(address,uint256[])",
		"7c61e561": "calculateRemoveLiquidity(address,uint256)",
		"798af720": "calculateSwap(address,uint8,uint8,uint256)",
		"ccc1bbc1": "calculateWithdrawOneToken(address,uint256,uint8)",
		"65170fa2": "defaultPoolCalc()",
		"a08129ce": "findConnectedTokens((uint256,address)[],address)",
		"e6b00009": "getAmountOut((uint256,address),address,uint256)",
		"9a7b5f11": "poolInfo(address)",
		"a9126169": "poolTokens(address)",
		"ba7d536e": "poolsAmount()",
		"3fc8cef3": "weth()",
	},
}

// ISwapQuoterV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapQuoterV1MetaData.ABI instead.
var ISwapQuoterV1ABI = ISwapQuoterV1MetaData.ABI

// Deprecated: Use ISwapQuoterV1MetaData.Sigs instead.
// ISwapQuoterV1FuncSigs maps the 4-byte function signature to its string representation.
var ISwapQuoterV1FuncSigs = ISwapQuoterV1MetaData.Sigs

// ISwapQuoterV1 is an auto generated Go binding around an Ethereum contract.
type ISwapQuoterV1 struct {
	ISwapQuoterV1Caller     // Read-only binding to the contract
	ISwapQuoterV1Transactor // Write-only binding to the contract
	ISwapQuoterV1Filterer   // Log filterer for contract events
}

// ISwapQuoterV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapQuoterV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapQuoterV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapQuoterV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapQuoterV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapQuoterV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapQuoterV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapQuoterV1Session struct {
	Contract     *ISwapQuoterV1    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapQuoterV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapQuoterV1CallerSession struct {
	Contract *ISwapQuoterV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISwapQuoterV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapQuoterV1TransactorSession struct {
	Contract     *ISwapQuoterV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISwapQuoterV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapQuoterV1Raw struct {
	Contract *ISwapQuoterV1 // Generic contract binding to access the raw methods on
}

// ISwapQuoterV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapQuoterV1CallerRaw struct {
	Contract *ISwapQuoterV1Caller // Generic read-only contract binding to access the raw methods on
}

// ISwapQuoterV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapQuoterV1TransactorRaw struct {
	Contract *ISwapQuoterV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapQuoterV1 creates a new instance of ISwapQuoterV1, bound to a specific deployed contract.
func NewISwapQuoterV1(address common.Address, backend bind.ContractBackend) (*ISwapQuoterV1, error) {
	contract, err := bindISwapQuoterV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapQuoterV1{ISwapQuoterV1Caller: ISwapQuoterV1Caller{contract: contract}, ISwapQuoterV1Transactor: ISwapQuoterV1Transactor{contract: contract}, ISwapQuoterV1Filterer: ISwapQuoterV1Filterer{contract: contract}}, nil
}

// NewISwapQuoterV1Caller creates a new read-only instance of ISwapQuoterV1, bound to a specific deployed contract.
func NewISwapQuoterV1Caller(address common.Address, caller bind.ContractCaller) (*ISwapQuoterV1Caller, error) {
	contract, err := bindISwapQuoterV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapQuoterV1Caller{contract: contract}, nil
}

// NewISwapQuoterV1Transactor creates a new write-only instance of ISwapQuoterV1, bound to a specific deployed contract.
func NewISwapQuoterV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ISwapQuoterV1Transactor, error) {
	contract, err := bindISwapQuoterV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapQuoterV1Transactor{contract: contract}, nil
}

// NewISwapQuoterV1Filterer creates a new log filterer instance of ISwapQuoterV1, bound to a specific deployed contract.
func NewISwapQuoterV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ISwapQuoterV1Filterer, error) {
	contract, err := bindISwapQuoterV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapQuoterV1Filterer{contract: contract}, nil
}

// bindISwapQuoterV1 binds a generic wrapper to an already deployed contract.
func bindISwapQuoterV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISwapQuoterV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapQuoterV1 *ISwapQuoterV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapQuoterV1.Contract.ISwapQuoterV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapQuoterV1 *ISwapQuoterV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapQuoterV1.Contract.ISwapQuoterV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapQuoterV1 *ISwapQuoterV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapQuoterV1.Contract.ISwapQuoterV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapQuoterV1 *ISwapQuoterV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapQuoterV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapQuoterV1 *ISwapQuoterV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapQuoterV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapQuoterV1 *ISwapQuoterV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapQuoterV1.Contract.contract.Transact(opts, method, params...)
}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) AllPools(opts *bind.CallOpts) ([]Pool, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "allPools")

	if err != nil {
		return *new([]Pool), err
	}

	out0 := *abi.ConvertType(out[0], new([]Pool)).(*[]Pool)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) AllPools() ([]Pool, error) {
	return _ISwapQuoterV1.Contract.AllPools(&_ISwapQuoterV1.CallOpts)
}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) AllPools() ([]Pool, error) {
	return _ISwapQuoterV1.Contract.AllPools(&_ISwapQuoterV1.CallOpts)
}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) CalculateAddLiquidity(opts *bind.CallOpts, pool common.Address, amounts []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "calculateAddLiquidity", pool, amounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) CalculateAddLiquidity(pool common.Address, amounts []*big.Int) (*big.Int, error) {
	return _ISwapQuoterV1.Contract.CalculateAddLiquidity(&_ISwapQuoterV1.CallOpts, pool, amounts)
}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) CalculateAddLiquidity(pool common.Address, amounts []*big.Int) (*big.Int, error) {
	return _ISwapQuoterV1.Contract.CalculateAddLiquidity(&_ISwapQuoterV1.CallOpts, pool, amounts)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) CalculateRemoveLiquidity(opts *bind.CallOpts, pool common.Address, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "calculateRemoveLiquidity", pool, amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) CalculateRemoveLiquidity(pool common.Address, amount *big.Int) ([]*big.Int, error) {
	return _ISwapQuoterV1.Contract.CalculateRemoveLiquidity(&_ISwapQuoterV1.CallOpts, pool, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) CalculateRemoveLiquidity(pool common.Address, amount *big.Int) ([]*big.Int, error) {
	return _ISwapQuoterV1.Contract.CalculateRemoveLiquidity(&_ISwapQuoterV1.CallOpts, pool, amount)
}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) CalculateSwap(opts *bind.CallOpts, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "calculateSwap", pool, tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) CalculateSwap(pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _ISwapQuoterV1.Contract.CalculateSwap(&_ISwapQuoterV1.CallOpts, pool, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) CalculateSwap(pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _ISwapQuoterV1.Contract.CalculateSwap(&_ISwapQuoterV1.CallOpts, pool, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) CalculateWithdrawOneToken(opts *bind.CallOpts, pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "calculateWithdrawOneToken", pool, tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) CalculateWithdrawOneToken(pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _ISwapQuoterV1.Contract.CalculateWithdrawOneToken(&_ISwapQuoterV1.CallOpts, pool, tokenAmount, tokenIndex)
}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) CalculateWithdrawOneToken(pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _ISwapQuoterV1.Contract.CalculateWithdrawOneToken(&_ISwapQuoterV1.CallOpts, pool, tokenAmount, tokenIndex)
}

// DefaultPoolCalc is a free data retrieval call binding the contract method 0x65170fa2.
//
// Solidity: function defaultPoolCalc() view returns(address)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) DefaultPoolCalc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "defaultPoolCalc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultPoolCalc is a free data retrieval call binding the contract method 0x65170fa2.
//
// Solidity: function defaultPoolCalc() view returns(address)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) DefaultPoolCalc() (common.Address, error) {
	return _ISwapQuoterV1.Contract.DefaultPoolCalc(&_ISwapQuoterV1.CallOpts)
}

// DefaultPoolCalc is a free data retrieval call binding the contract method 0x65170fa2.
//
// Solidity: function defaultPoolCalc() view returns(address)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) DefaultPoolCalc() (common.Address, error) {
	return _ISwapQuoterV1.Contract.DefaultPoolCalc(&_ISwapQuoterV1.CallOpts)
}

// FindConnectedTokens is a free data retrieval call binding the contract method 0xa08129ce.
//
// Solidity: function findConnectedTokens((uint256,address)[] bridgeTokensIn, address tokenOut) view returns(uint256 amountFound, bool[] isConnected)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) FindConnectedTokens(opts *bind.CallOpts, bridgeTokensIn []LimitedToken, tokenOut common.Address) (struct {
	AmountFound *big.Int
	IsConnected []bool
}, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "findConnectedTokens", bridgeTokensIn, tokenOut)

	outstruct := new(struct {
		AmountFound *big.Int
		IsConnected []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountFound = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsConnected = *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return *outstruct, err

}

// FindConnectedTokens is a free data retrieval call binding the contract method 0xa08129ce.
//
// Solidity: function findConnectedTokens((uint256,address)[] bridgeTokensIn, address tokenOut) view returns(uint256 amountFound, bool[] isConnected)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) FindConnectedTokens(bridgeTokensIn []LimitedToken, tokenOut common.Address) (struct {
	AmountFound *big.Int
	IsConnected []bool
}, error) {
	return _ISwapQuoterV1.Contract.FindConnectedTokens(&_ISwapQuoterV1.CallOpts, bridgeTokensIn, tokenOut)
}

// FindConnectedTokens is a free data retrieval call binding the contract method 0xa08129ce.
//
// Solidity: function findConnectedTokens((uint256,address)[] bridgeTokensIn, address tokenOut) view returns(uint256 amountFound, bool[] isConnected)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) FindConnectedTokens(bridgeTokensIn []LimitedToken, tokenOut common.Address) (struct {
	AmountFound *big.Int
	IsConnected []bool
}, error) {
	return _ISwapQuoterV1.Contract.FindConnectedTokens(&_ISwapQuoterV1.CallOpts, bridgeTokensIn, tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xe6b00009.
//
// Solidity: function getAmountOut((uint256,address) tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) query)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) GetAmountOut(opts *bind.CallOpts, tokenIn LimitedToken, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "getAmountOut", tokenIn, tokenOut, amountIn)

	if err != nil {
		return *new(SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new(SwapQuery)).(*SwapQuery)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xe6b00009.
//
// Solidity: function getAmountOut((uint256,address) tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) query)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) GetAmountOut(tokenIn LimitedToken, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	return _ISwapQuoterV1.Contract.GetAmountOut(&_ISwapQuoterV1.CallOpts, tokenIn, tokenOut, amountIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xe6b00009.
//
// Solidity: function getAmountOut((uint256,address) tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) query)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) GetAmountOut(tokenIn LimitedToken, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	return _ISwapQuoterV1.Contract.GetAmountOut(&_ISwapQuoterV1.CallOpts, tokenIn, tokenOut, amountIn)
}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256 numTokens, address lpToken)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) PoolInfo(opts *bind.CallOpts, pool common.Address) (struct {
	NumTokens *big.Int
	LpToken   common.Address
}, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "poolInfo", pool)

	outstruct := new(struct {
		NumTokens *big.Int
		LpToken   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumTokens = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LpToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256 numTokens, address lpToken)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) PoolInfo(pool common.Address) (struct {
	NumTokens *big.Int
	LpToken   common.Address
}, error) {
	return _ISwapQuoterV1.Contract.PoolInfo(&_ISwapQuoterV1.CallOpts, pool)
}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256 numTokens, address lpToken)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) PoolInfo(pool common.Address) (struct {
	NumTokens *big.Int
	LpToken   common.Address
}, error) {
	return _ISwapQuoterV1.Contract.PoolInfo(&_ISwapQuoterV1.CallOpts, pool)
}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) PoolTokens(opts *bind.CallOpts, pool common.Address) ([]PoolToken, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "poolTokens", pool)

	if err != nil {
		return *new([]PoolToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolToken)).(*[]PoolToken)

	return out0, err

}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) PoolTokens(pool common.Address) ([]PoolToken, error) {
	return _ISwapQuoterV1.Contract.PoolTokens(&_ISwapQuoterV1.CallOpts, pool)
}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) PoolTokens(pool common.Address) ([]PoolToken, error) {
	return _ISwapQuoterV1.Contract.PoolTokens(&_ISwapQuoterV1.CallOpts, pool)
}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amtPools)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) PoolsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "poolsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amtPools)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) PoolsAmount() (*big.Int, error) {
	return _ISwapQuoterV1.Contract.PoolsAmount(&_ISwapQuoterV1.CallOpts)
}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amtPools)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) PoolsAmount() (*big.Int, error) {
	return _ISwapQuoterV1.Contract.PoolsAmount(&_ISwapQuoterV1.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ISwapQuoterV1 *ISwapQuoterV1Caller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISwapQuoterV1.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ISwapQuoterV1 *ISwapQuoterV1Session) Weth() (common.Address, error) {
	return _ISwapQuoterV1.Contract.Weth(&_ISwapQuoterV1.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ISwapQuoterV1 *ISwapQuoterV1CallerSession) Weth() (common.Address, error) {
	return _ISwapQuoterV1.Contract.Weth(&_ISwapQuoterV1.CallOpts)
}

// ISwapQuoterV2MetaData contains all meta data concerning the ISwapQuoterV2 contract.
var ISwapQuoterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"allPools\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWeth\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPoolToken[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"internalType\":\"structPool[]\",\"name\":\"pools\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"actionMask\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structLimitedToken\",\"name\":\"tokenIn\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"areConnectedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateWithdrawOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultPoolCalc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"actionMask\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structLimitedToken[]\",\"name\":\"bridgeTokensIn\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"findConnectedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountFound\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"isConnected\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"actionMask\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structLimitedToken\",\"name\":\"tokenIn\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"query\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"poolTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWeth\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPoolToken[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amtPools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"synapseRouter_\",\"type\":\"address\"}],\"name\":\"setSynapseRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c5c63e65": "allPools()",
		"285078fc": "areConnectedTokens((uint256,address),address)",
		"4d864496": "calculateAddLiquidity(address,uint256[])",
		"7c61e561": "calculateRemoveLiquidity(address,uint256)",
		"798af720": "calculateSwap(address,uint8,uint8,uint256)",
		"ccc1bbc1": "calculateWithdrawOneToken(address,uint256,uint8)",
		"65170fa2": "defaultPoolCalc()",
		"a08129ce": "findConnectedTokens((uint256,address)[],address)",
		"e6b00009": "getAmountOut((uint256,address),address,uint256)",
		"9a7b5f11": "poolInfo(address)",
		"a9126169": "poolTokens(address)",
		"ba7d536e": "poolsAmount()",
		"446bac69": "setSynapseRouter(address)",
		"3fc8cef3": "weth()",
	},
}

// ISwapQuoterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapQuoterV2MetaData.ABI instead.
var ISwapQuoterV2ABI = ISwapQuoterV2MetaData.ABI

// Deprecated: Use ISwapQuoterV2MetaData.Sigs instead.
// ISwapQuoterV2FuncSigs maps the 4-byte function signature to its string representation.
var ISwapQuoterV2FuncSigs = ISwapQuoterV2MetaData.Sigs

// ISwapQuoterV2 is an auto generated Go binding around an Ethereum contract.
type ISwapQuoterV2 struct {
	ISwapQuoterV2Caller     // Read-only binding to the contract
	ISwapQuoterV2Transactor // Write-only binding to the contract
	ISwapQuoterV2Filterer   // Log filterer for contract events
}

// ISwapQuoterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapQuoterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapQuoterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapQuoterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapQuoterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapQuoterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapQuoterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapQuoterV2Session struct {
	Contract     *ISwapQuoterV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapQuoterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapQuoterV2CallerSession struct {
	Contract *ISwapQuoterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISwapQuoterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapQuoterV2TransactorSession struct {
	Contract     *ISwapQuoterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISwapQuoterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapQuoterV2Raw struct {
	Contract *ISwapQuoterV2 // Generic contract binding to access the raw methods on
}

// ISwapQuoterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapQuoterV2CallerRaw struct {
	Contract *ISwapQuoterV2Caller // Generic read-only contract binding to access the raw methods on
}

// ISwapQuoterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapQuoterV2TransactorRaw struct {
	Contract *ISwapQuoterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapQuoterV2 creates a new instance of ISwapQuoterV2, bound to a specific deployed contract.
func NewISwapQuoterV2(address common.Address, backend bind.ContractBackend) (*ISwapQuoterV2, error) {
	contract, err := bindISwapQuoterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapQuoterV2{ISwapQuoterV2Caller: ISwapQuoterV2Caller{contract: contract}, ISwapQuoterV2Transactor: ISwapQuoterV2Transactor{contract: contract}, ISwapQuoterV2Filterer: ISwapQuoterV2Filterer{contract: contract}}, nil
}

// NewISwapQuoterV2Caller creates a new read-only instance of ISwapQuoterV2, bound to a specific deployed contract.
func NewISwapQuoterV2Caller(address common.Address, caller bind.ContractCaller) (*ISwapQuoterV2Caller, error) {
	contract, err := bindISwapQuoterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapQuoterV2Caller{contract: contract}, nil
}

// NewISwapQuoterV2Transactor creates a new write-only instance of ISwapQuoterV2, bound to a specific deployed contract.
func NewISwapQuoterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ISwapQuoterV2Transactor, error) {
	contract, err := bindISwapQuoterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapQuoterV2Transactor{contract: contract}, nil
}

// NewISwapQuoterV2Filterer creates a new log filterer instance of ISwapQuoterV2, bound to a specific deployed contract.
func NewISwapQuoterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ISwapQuoterV2Filterer, error) {
	contract, err := bindISwapQuoterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapQuoterV2Filterer{contract: contract}, nil
}

// bindISwapQuoterV2 binds a generic wrapper to an already deployed contract.
func bindISwapQuoterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISwapQuoterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapQuoterV2 *ISwapQuoterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapQuoterV2.Contract.ISwapQuoterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapQuoterV2 *ISwapQuoterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapQuoterV2.Contract.ISwapQuoterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapQuoterV2 *ISwapQuoterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapQuoterV2.Contract.ISwapQuoterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapQuoterV2 *ISwapQuoterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapQuoterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapQuoterV2 *ISwapQuoterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapQuoterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapQuoterV2 *ISwapQuoterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapQuoterV2.Contract.contract.Transact(opts, method, params...)
}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) AllPools(opts *bind.CallOpts) ([]Pool, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "allPools")

	if err != nil {
		return *new([]Pool), err
	}

	out0 := *abi.ConvertType(out[0], new([]Pool)).(*[]Pool)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) AllPools() ([]Pool, error) {
	return _ISwapQuoterV2.Contract.AllPools(&_ISwapQuoterV2.CallOpts)
}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) AllPools() ([]Pool, error) {
	return _ISwapQuoterV2.Contract.AllPools(&_ISwapQuoterV2.CallOpts)
}

// AreConnectedTokens is a free data retrieval call binding the contract method 0x285078fc.
//
// Solidity: function areConnectedTokens((uint256,address) tokenIn, address tokenOut) view returns(bool)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) AreConnectedTokens(opts *bind.CallOpts, tokenIn LimitedToken, tokenOut common.Address) (bool, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "areConnectedTokens", tokenIn, tokenOut)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreConnectedTokens is a free data retrieval call binding the contract method 0x285078fc.
//
// Solidity: function areConnectedTokens((uint256,address) tokenIn, address tokenOut) view returns(bool)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) AreConnectedTokens(tokenIn LimitedToken, tokenOut common.Address) (bool, error) {
	return _ISwapQuoterV2.Contract.AreConnectedTokens(&_ISwapQuoterV2.CallOpts, tokenIn, tokenOut)
}

// AreConnectedTokens is a free data retrieval call binding the contract method 0x285078fc.
//
// Solidity: function areConnectedTokens((uint256,address) tokenIn, address tokenOut) view returns(bool)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) AreConnectedTokens(tokenIn LimitedToken, tokenOut common.Address) (bool, error) {
	return _ISwapQuoterV2.Contract.AreConnectedTokens(&_ISwapQuoterV2.CallOpts, tokenIn, tokenOut)
}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) CalculateAddLiquidity(opts *bind.CallOpts, pool common.Address, amounts []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "calculateAddLiquidity", pool, amounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) CalculateAddLiquidity(pool common.Address, amounts []*big.Int) (*big.Int, error) {
	return _ISwapQuoterV2.Contract.CalculateAddLiquidity(&_ISwapQuoterV2.CallOpts, pool, amounts)
}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) CalculateAddLiquidity(pool common.Address, amounts []*big.Int) (*big.Int, error) {
	return _ISwapQuoterV2.Contract.CalculateAddLiquidity(&_ISwapQuoterV2.CallOpts, pool, amounts)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) CalculateRemoveLiquidity(opts *bind.CallOpts, pool common.Address, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "calculateRemoveLiquidity", pool, amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) CalculateRemoveLiquidity(pool common.Address, amount *big.Int) ([]*big.Int, error) {
	return _ISwapQuoterV2.Contract.CalculateRemoveLiquidity(&_ISwapQuoterV2.CallOpts, pool, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) CalculateRemoveLiquidity(pool common.Address, amount *big.Int) ([]*big.Int, error) {
	return _ISwapQuoterV2.Contract.CalculateRemoveLiquidity(&_ISwapQuoterV2.CallOpts, pool, amount)
}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) CalculateSwap(opts *bind.CallOpts, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "calculateSwap", pool, tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) CalculateSwap(pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _ISwapQuoterV2.Contract.CalculateSwap(&_ISwapQuoterV2.CallOpts, pool, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) CalculateSwap(pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _ISwapQuoterV2.Contract.CalculateSwap(&_ISwapQuoterV2.CallOpts, pool, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) CalculateWithdrawOneToken(opts *bind.CallOpts, pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "calculateWithdrawOneToken", pool, tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) CalculateWithdrawOneToken(pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _ISwapQuoterV2.Contract.CalculateWithdrawOneToken(&_ISwapQuoterV2.CallOpts, pool, tokenAmount, tokenIndex)
}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) CalculateWithdrawOneToken(pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _ISwapQuoterV2.Contract.CalculateWithdrawOneToken(&_ISwapQuoterV2.CallOpts, pool, tokenAmount, tokenIndex)
}

// DefaultPoolCalc is a free data retrieval call binding the contract method 0x65170fa2.
//
// Solidity: function defaultPoolCalc() view returns(address)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) DefaultPoolCalc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "defaultPoolCalc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultPoolCalc is a free data retrieval call binding the contract method 0x65170fa2.
//
// Solidity: function defaultPoolCalc() view returns(address)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) DefaultPoolCalc() (common.Address, error) {
	return _ISwapQuoterV2.Contract.DefaultPoolCalc(&_ISwapQuoterV2.CallOpts)
}

// DefaultPoolCalc is a free data retrieval call binding the contract method 0x65170fa2.
//
// Solidity: function defaultPoolCalc() view returns(address)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) DefaultPoolCalc() (common.Address, error) {
	return _ISwapQuoterV2.Contract.DefaultPoolCalc(&_ISwapQuoterV2.CallOpts)
}

// FindConnectedTokens is a free data retrieval call binding the contract method 0xa08129ce.
//
// Solidity: function findConnectedTokens((uint256,address)[] bridgeTokensIn, address tokenOut) view returns(uint256 amountFound, bool[] isConnected)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) FindConnectedTokens(opts *bind.CallOpts, bridgeTokensIn []LimitedToken, tokenOut common.Address) (struct {
	AmountFound *big.Int
	IsConnected []bool
}, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "findConnectedTokens", bridgeTokensIn, tokenOut)

	outstruct := new(struct {
		AmountFound *big.Int
		IsConnected []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountFound = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsConnected = *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return *outstruct, err

}

// FindConnectedTokens is a free data retrieval call binding the contract method 0xa08129ce.
//
// Solidity: function findConnectedTokens((uint256,address)[] bridgeTokensIn, address tokenOut) view returns(uint256 amountFound, bool[] isConnected)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) FindConnectedTokens(bridgeTokensIn []LimitedToken, tokenOut common.Address) (struct {
	AmountFound *big.Int
	IsConnected []bool
}, error) {
	return _ISwapQuoterV2.Contract.FindConnectedTokens(&_ISwapQuoterV2.CallOpts, bridgeTokensIn, tokenOut)
}

// FindConnectedTokens is a free data retrieval call binding the contract method 0xa08129ce.
//
// Solidity: function findConnectedTokens((uint256,address)[] bridgeTokensIn, address tokenOut) view returns(uint256 amountFound, bool[] isConnected)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) FindConnectedTokens(bridgeTokensIn []LimitedToken, tokenOut common.Address) (struct {
	AmountFound *big.Int
	IsConnected []bool
}, error) {
	return _ISwapQuoterV2.Contract.FindConnectedTokens(&_ISwapQuoterV2.CallOpts, bridgeTokensIn, tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xe6b00009.
//
// Solidity: function getAmountOut((uint256,address) tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) query)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) GetAmountOut(opts *bind.CallOpts, tokenIn LimitedToken, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "getAmountOut", tokenIn, tokenOut, amountIn)

	if err != nil {
		return *new(SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new(SwapQuery)).(*SwapQuery)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xe6b00009.
//
// Solidity: function getAmountOut((uint256,address) tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) query)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) GetAmountOut(tokenIn LimitedToken, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	return _ISwapQuoterV2.Contract.GetAmountOut(&_ISwapQuoterV2.CallOpts, tokenIn, tokenOut, amountIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xe6b00009.
//
// Solidity: function getAmountOut((uint256,address) tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) query)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) GetAmountOut(tokenIn LimitedToken, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	return _ISwapQuoterV2.Contract.GetAmountOut(&_ISwapQuoterV2.CallOpts, tokenIn, tokenOut, amountIn)
}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256 numTokens, address lpToken)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) PoolInfo(opts *bind.CallOpts, pool common.Address) (struct {
	NumTokens *big.Int
	LpToken   common.Address
}, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "poolInfo", pool)

	outstruct := new(struct {
		NumTokens *big.Int
		LpToken   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumTokens = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LpToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256 numTokens, address lpToken)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) PoolInfo(pool common.Address) (struct {
	NumTokens *big.Int
	LpToken   common.Address
}, error) {
	return _ISwapQuoterV2.Contract.PoolInfo(&_ISwapQuoterV2.CallOpts, pool)
}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256 numTokens, address lpToken)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) PoolInfo(pool common.Address) (struct {
	NumTokens *big.Int
	LpToken   common.Address
}, error) {
	return _ISwapQuoterV2.Contract.PoolInfo(&_ISwapQuoterV2.CallOpts, pool)
}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) PoolTokens(opts *bind.CallOpts, pool common.Address) ([]PoolToken, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "poolTokens", pool)

	if err != nil {
		return *new([]PoolToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolToken)).(*[]PoolToken)

	return out0, err

}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) PoolTokens(pool common.Address) ([]PoolToken, error) {
	return _ISwapQuoterV2.Contract.PoolTokens(&_ISwapQuoterV2.CallOpts, pool)
}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) PoolTokens(pool common.Address) ([]PoolToken, error) {
	return _ISwapQuoterV2.Contract.PoolTokens(&_ISwapQuoterV2.CallOpts, pool)
}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amtPools)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) PoolsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "poolsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amtPools)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) PoolsAmount() (*big.Int, error) {
	return _ISwapQuoterV2.Contract.PoolsAmount(&_ISwapQuoterV2.CallOpts)
}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amtPools)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) PoolsAmount() (*big.Int, error) {
	return _ISwapQuoterV2.Contract.PoolsAmount(&_ISwapQuoterV2.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ISwapQuoterV2 *ISwapQuoterV2Caller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISwapQuoterV2.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ISwapQuoterV2 *ISwapQuoterV2Session) Weth() (common.Address, error) {
	return _ISwapQuoterV2.Contract.Weth(&_ISwapQuoterV2.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_ISwapQuoterV2 *ISwapQuoterV2CallerSession) Weth() (common.Address, error) {
	return _ISwapQuoterV2.Contract.Weth(&_ISwapQuoterV2.CallOpts)
}

// SetSynapseRouter is a paid mutator transaction binding the contract method 0x446bac69.
//
// Solidity: function setSynapseRouter(address synapseRouter_) returns()
func (_ISwapQuoterV2 *ISwapQuoterV2Transactor) SetSynapseRouter(opts *bind.TransactOpts, synapseRouter_ common.Address) (*types.Transaction, error) {
	return _ISwapQuoterV2.contract.Transact(opts, "setSynapseRouter", synapseRouter_)
}

// SetSynapseRouter is a paid mutator transaction binding the contract method 0x446bac69.
//
// Solidity: function setSynapseRouter(address synapseRouter_) returns()
func (_ISwapQuoterV2 *ISwapQuoterV2Session) SetSynapseRouter(synapseRouter_ common.Address) (*types.Transaction, error) {
	return _ISwapQuoterV2.Contract.SetSynapseRouter(&_ISwapQuoterV2.TransactOpts, synapseRouter_)
}

// SetSynapseRouter is a paid mutator transaction binding the contract method 0x446bac69.
//
// Solidity: function setSynapseRouter(address synapseRouter_) returns()
func (_ISwapQuoterV2 *ISwapQuoterV2TransactorSession) SetSynapseRouter(synapseRouter_ common.Address) (*types.Transaction, error) {
	return _ISwapQuoterV2.Contract.SetSynapseRouter(&_ISwapQuoterV2.TransactOpts, synapseRouter_)
}

// IWETH9MetaData contains all meta data concerning the IWETH9 contract.
var IWETH9MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d0e30db0": "deposit()",
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// IWETH9ABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETH9MetaData.ABI instead.
var IWETH9ABI = IWETH9MetaData.ABI

// Deprecated: Use IWETH9MetaData.Sigs instead.
// IWETH9FuncSigs maps the 4-byte function signature to its string representation.
var IWETH9FuncSigs = IWETH9MetaData.Sigs

// IWETH9 is an auto generated Go binding around an Ethereum contract.
type IWETH9 struct {
	IWETH9Caller     // Read-only binding to the contract
	IWETH9Transactor // Write-only binding to the contract
	IWETH9Filterer   // Log filterer for contract events
}

// IWETH9Caller is an auto generated read-only Go binding around an Ethereum contract.
type IWETH9Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETH9Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETH9Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETH9Session struct {
	Contract     *IWETH9           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETH9CallerSession struct {
	Contract *IWETH9Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETH9TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETH9TransactorSession struct {
	Contract     *IWETH9Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9Raw is an auto generated low-level Go binding around an Ethereum contract.
type IWETH9Raw struct {
	Contract *IWETH9 // Generic contract binding to access the raw methods on
}

// IWETH9CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETH9CallerRaw struct {
	Contract *IWETH9Caller // Generic read-only contract binding to access the raw methods on
}

// IWETH9TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETH9TransactorRaw struct {
	Contract *IWETH9Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH9 creates a new instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9(address common.Address, backend bind.ContractBackend) (*IWETH9, error) {
	contract, err := bindIWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH9{IWETH9Caller: IWETH9Caller{contract: contract}, IWETH9Transactor: IWETH9Transactor{contract: contract}, IWETH9Filterer: IWETH9Filterer{contract: contract}}, nil
}

// NewIWETH9Caller creates a new read-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Caller(address common.Address, caller bind.ContractCaller) (*IWETH9Caller, error) {
	contract, err := bindIWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Caller{contract: contract}, nil
}

// NewIWETH9Transactor creates a new write-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*IWETH9Transactor, error) {
	contract, err := bindIWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Transactor{contract: contract}, nil
}

// NewIWETH9Filterer creates a new log filterer instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*IWETH9Filterer, error) {
	contract, err := bindIWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETH9Filterer{contract: contract}, nil
}

// bindIWETH9 binds a generic wrapper to an already deployed contract.
func bindIWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWETH9MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.IWETH9Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Session) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9TransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWETH9 *IWETH9Transactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWETH9 *IWETH9Session) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_IWETH9 *IWETH9TransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, wad)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206b970769f18a9b7f66a34ebe5506d065a5867642184e7f5e456173c975e5bc8b64736f6c63430008110033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SwapQueryLibMetaData contains all meta data concerning the SwapQueryLib contract.
var SwapQueryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122025ce66cc54bdbf742df79150788eacb4aff8f6818b1eff240e12aea5718410ef64736f6c63430008110033",
}

// SwapQueryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapQueryLibMetaData.ABI instead.
var SwapQueryLibABI = SwapQueryLibMetaData.ABI

// SwapQueryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapQueryLibMetaData.Bin instead.
var SwapQueryLibBin = SwapQueryLibMetaData.Bin

// DeploySwapQueryLib deploys a new Ethereum contract, binding an instance of SwapQueryLib to it.
func DeploySwapQueryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwapQueryLib, error) {
	parsed, err := SwapQueryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapQueryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapQueryLib{SwapQueryLibCaller: SwapQueryLibCaller{contract: contract}, SwapQueryLibTransactor: SwapQueryLibTransactor{contract: contract}, SwapQueryLibFilterer: SwapQueryLibFilterer{contract: contract}}, nil
}

// SwapQueryLib is an auto generated Go binding around an Ethereum contract.
type SwapQueryLib struct {
	SwapQueryLibCaller     // Read-only binding to the contract
	SwapQueryLibTransactor // Write-only binding to the contract
	SwapQueryLibFilterer   // Log filterer for contract events
}

// SwapQueryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapQueryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapQueryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapQueryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapQueryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapQueryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapQueryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapQueryLibSession struct {
	Contract     *SwapQueryLib     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapQueryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapQueryLibCallerSession struct {
	Contract *SwapQueryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SwapQueryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapQueryLibTransactorSession struct {
	Contract     *SwapQueryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SwapQueryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapQueryLibRaw struct {
	Contract *SwapQueryLib // Generic contract binding to access the raw methods on
}

// SwapQueryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapQueryLibCallerRaw struct {
	Contract *SwapQueryLibCaller // Generic read-only contract binding to access the raw methods on
}

// SwapQueryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapQueryLibTransactorRaw struct {
	Contract *SwapQueryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapQueryLib creates a new instance of SwapQueryLib, bound to a specific deployed contract.
func NewSwapQueryLib(address common.Address, backend bind.ContractBackend) (*SwapQueryLib, error) {
	contract, err := bindSwapQueryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapQueryLib{SwapQueryLibCaller: SwapQueryLibCaller{contract: contract}, SwapQueryLibTransactor: SwapQueryLibTransactor{contract: contract}, SwapQueryLibFilterer: SwapQueryLibFilterer{contract: contract}}, nil
}

// NewSwapQueryLibCaller creates a new read-only instance of SwapQueryLib, bound to a specific deployed contract.
func NewSwapQueryLibCaller(address common.Address, caller bind.ContractCaller) (*SwapQueryLibCaller, error) {
	contract, err := bindSwapQueryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapQueryLibCaller{contract: contract}, nil
}

// NewSwapQueryLibTransactor creates a new write-only instance of SwapQueryLib, bound to a specific deployed contract.
func NewSwapQueryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapQueryLibTransactor, error) {
	contract, err := bindSwapQueryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapQueryLibTransactor{contract: contract}, nil
}

// NewSwapQueryLibFilterer creates a new log filterer instance of SwapQueryLib, bound to a specific deployed contract.
func NewSwapQueryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapQueryLibFilterer, error) {
	contract, err := bindSwapQueryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapQueryLibFilterer{contract: contract}, nil
}

// bindSwapQueryLib binds a generic wrapper to an already deployed contract.
func bindSwapQueryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapQueryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapQueryLib *SwapQueryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapQueryLib.Contract.SwapQueryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapQueryLib *SwapQueryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapQueryLib.Contract.SwapQueryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapQueryLib *SwapQueryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapQueryLib.Contract.SwapQueryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapQueryLib *SwapQueryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapQueryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapQueryLib *SwapQueryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapQueryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapQueryLib *SwapQueryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapQueryLib.Contract.contract.Transact(opts, method, params...)
}

// SynapseRouterV2MetaData contains all meta data concerning the SynapseRouterV2 contract.
var SynapseRouterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"ArrayLengthInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientOutputAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseRouterV2__ModuleExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseRouterV2__ModuleInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseRouterV2__ModuleNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseRouterV2__QueryEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAddressMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensIdentical\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"name\":\"ModuleConnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"}],\"name\":\"ModuleDisconnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBridgeModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBridgeModule\",\"type\":\"address\"}],\"name\":\"ModuleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSwapQuoter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSwapQuoter\",\"type\":\"address\"}],\"name\":\"QuoterSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"name\":\"adapterSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"originQuery\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"destQuery\",\"type\":\"tuple\"}],\"name\":\"bridgeViaSynapse\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"name\":\"connectBridgeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"}],\"name\":\"disconnectBridgeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structDestRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getDestinationAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"destQuery\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getDestinationBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"destTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getOriginAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"originQuery\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"}],\"name\":\"getOriginBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"originTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"supportedTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"}],\"name\":\"idToModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"name\":\"moduleToId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISwapQuoterV2\",\"name\":\"_swapQuoter\",\"type\":\"address\"}],\"name\":\"setSwapQuoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"routerAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"query\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapQuoter\",\"outputs\":[{\"internalType\":\"contractISwapQuoterV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bridgeModule\",\"type\":\"address\"}],\"name\":\"updateBridgeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"24a98f11": "adapterSwap(address,address,uint256,address,bytes)",
		"c95fafd2": "bridgeViaSynapse(address,uint256,bytes32,address,uint256,(address,address,uint256,uint256,bytes),(address,address,uint256,uint256,bytes))",
		"b3bce952": "connectBridgeModule(bytes32,address)",
		"b68e4302": "disconnectBridgeModule(bytes32)",
		"9c1d060e": "getBridgeTokens()",
		"7de31c74": "getDestinationAmountOut((string,uint256),address)",
		"1d04879b": "getDestinationBridgeTokens(address)",
		"f533941d": "getOriginAmountOut(address,string,uint256)",
		"3e811a5c": "getOriginBridgeTokens(address)",
		"d3c7c2c7": "getSupportedTokens()",
		"53e2e8e7": "idToModule(bytes32)",
		"9f2671fa": "moduleToId(address)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"da46098c": "setAllowance(address,address,uint256)",
		"804b3dff": "setSwapQuoter(address)",
		"b5d1cdd4": "swap(address,address,uint256,(address,address,uint256,uint256,bytes))",
		"34474c8c": "swapQuoter()",
		"f2fde38b": "transferOwnership(address)",
		"70a1cdc9": "updateBridgeModule(bytes32,address)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506200001d3362000023565b62000073565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b614b4f80620000836000396000f3fe6080604052600436106101635760003560e01c80639c1d060e116100c0578063c95fafd211610074578063da46098c11610059578063da46098c1461039d578063f2fde38b146103bd578063f533941d146103dd57600080fd5b8063c95fafd214610368578063d3c7c2c71461037b57600080fd5b8063b3bce952116100a5578063b3bce95214610315578063b5d1cdd414610335578063b68e43021461034857600080fd5b80639c1d060e146102e05780639f2671fa146102f557600080fd5b806370a1cdc9116101175780637de31c74116100fc5780637de31c7414610275578063804b3dff146102a25780638da5cb5b146102c257600080fd5b806370a1cdc91461023e578063715018a61461026057600080fd5b806334474c8c1161014857806334474c8c146101c65780633e811a5c146101fe57806353e2e8e71461021e57600080fd5b80631d04879b1461016f57806324a98f11146101a557600080fd5b3661016a57005b600080fd5b34801561017b57600080fd5b5061018f61018a366004613d51565b6103fd565b60405161019c9190613dbe565b60405180910390f35b6101b86101b3366004613fa8565b610410565b60405190815260200161019c565b3480156101d257600080fd5b506001546101e6906001600160a01b031681565b6040516001600160a01b03909116815260200161019c565b34801561020a57600080fd5b5061018f610219366004613d51565b610429565b34801561022a57600080fd5b506101e6610239366004614027565b610436565b34801561024a57600080fd5b5061025e610259366004614040565b610482565b005b34801561026c57600080fd5b5061025e6105c6565b34801561028157600080fd5b50610295610290366004614070565b61062c565b60405161019c9190614138565b3480156102ae57600080fd5b5061025e6102bd366004613d51565b61085c565b3480156102ce57600080fd5b506000546001600160a01b03166101e6565b3480156102ec57600080fd5b5061018f610937565b34801561030157600080fd5b506101b8610310366004613d51565b610a84565b34801561032157600080fd5b5061025e610330366004614040565b610b22565b6101b86103433660046141cf565b610c56565b34801561035457600080fd5b5061025e610363366004614027565b610cad565b61025e61037636600461423b565b610d80565b34801561038757600080fd5b50610390610e6c565b60405161019c91906142e1565b3480156103a957600080fd5b5061025e6103b836600461432e565b611243565b3480156103c957600080fd5b5061025e6103d8366004613d51565b6112b6565b3480156103e957600080fd5b506102956103f836600461436f565b611398565b606061040a82600061157c565b92915050565b600061041f8686868686611a34565b9695505050505050565b606061040a82600161157c565b600061044182611ac8565b610477576040517f45a6d1a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61040a600283611ad5565b6000546001600160a01b031633146104e15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b038116610521576040517f5920be2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61052a82611ac8565b610560576040517f45a6d1a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061056d600284611ad5565b905061057b60028484611ae1565b50604080516001600160a01b0380841682528416602082015284917fbe59b1ad7b24549601b98854029a8be9cd632ee55e4472692aa2542e668496e0910160405180910390a2505050565b6000546001600160a01b031633146106205760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104d8565b61062a6000611aff565b565b6106706040518060a0016040528060006001600160a01b0316815260200160006001600160a01b031681526020016000815260200160008152602001606081525090565b60008060006106828660000151611b67565b919450925090506001600160a01b03831661069f5750505061040a565b6000856001600160a01b0316846001600160a01b0316148061076f57506001600160a01b03861673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14801561076f5750600160009054906101000a90046001600160a01b03166001600160a01b0316633fc8cef36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610736573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061075a91906143c8565b6001600160a01b0316846001600160a01b0316145b159050600061078483868a6020015185611cd2565b90508060000361079857505050505061040a565b6040805180820182528581526001600160a01b038781166020830190815260015493517fe6b00009000000000000000000000000000000000000000000000000000000008152835160048201529051821660248201528a82166044820152606481018590529192169063e6b0000990608401600060405180830381865afa158015610827573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261084f9190810190614415565b9998505050505050505050565b6000546001600160a01b031633146108b65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104d8565b600154604080516001600160a01b03928316815291831660208301527f2e3d7d02ba3c4bd8b1f8995cd3a23ef0193922ebc4ee23249ead4d0ca2e34c68910160405180910390a1600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606060006109456002611d89565b905060008167ffffffffffffffff81111561096257610962613e5b565b60405190808252806020026020018201604052801561099557816020015b60608152602001906001900390816109805790505b5090506000805b83811015610a715760006109b1600283611d94565b915050806001600160a01b0316639c1d060e6040518163ffffffff1660e01b8152600401600060405180830381865afa1580156109f2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a1a91908101906144f9565b848381518110610a2c57610a2c614608565b6020026020010181905250838281518110610a4957610a49614608565b60200260200101515183610a5d9190614666565b92505080610a6a90614679565b905061099c565b50610a7c8282611db0565b935050505090565b600080610a916002611d89565b905060005b81811015610ae457600080610aac600284611d94565b91509150856001600160a01b0316816001600160a01b031603610ad157509250610ae4565b505080610add90614679565b9050610a96565b5081610b1c576040517f45a6d1a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50919050565b6000546001600160a01b03163314610b7c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104d8565b811580610b9057506001600160a01b038116155b15610bc7576040517f5920be2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bd082611ac8565b15610c07576040517f4b42265600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c1360028383611ae1565b506040516001600160a01b038216815282907fc92d33ac2d951a5a8265420d37e1664f111f74de9207f89f4c772142ec3aa5e39060200160405180910390a25050565b6000610c6b82516001600160a01b0316151590565b610ca1576040517fb50001b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61041f85858585611f0a565b6000546001600160a01b03163314610d075760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104d8565b610d1081611ac8565b610d46576040517f45a6d1a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d51600282612087565b5060405181907f594a6e72239daa84da65d4bbeb00fa0ccb8c579b2fd6b5601c82baf0614894cc90600090a250565b6000610d8b86610436565b9050610da083516001600160a01b0316151590565b15610dbb57610db130868686611f0a565b9095509350610dc9565b610dc6308686612093565b93505b600063436f3aa560e01b8989888887604051602401610dec959493929190614693565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091529050610e606001600160a01b03831682612278565b50505050505050505050565b60606000600160009054906101000a90046001600160a01b03166001600160a01b031663c5c63e656040518163ffffffff1660e01b8152600401600060405180830381865afa158015610ec3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610eeb91908101906146e0565b9050600081516001610efd9190614666565b67ffffffffffffffff811115610f1557610f15613e5b565b604051908082528060200260200182016040528015610f4857816020015b6060815260200190600190039081610f335790505b5090506000610f5d610f58610937565b61229d565b82845181518110610f7057610f70614608565b602002602001018190525081835181518110610f8e57610f8e614608565b60200260200101515181610fa29190614666565b905060005b8351811015611184576000848281518110610fc457610fc4614608565b6020026020010151905080604001515167ffffffffffffffff811115610fec57610fec613e5b565b604051908082528060200260200182016040528015611015578160200160208202803683370190505b5084838151811061102857611028614608565b6020026020010181905250600061106682602001518688518151811061105057611050614608565b602002602001015161234990919063ffffffff16565b905060005b826040015151811015611135578260400151818151811061108e5761108e614608565b6020026020010151602001518685815181106110ac576110ac614608565b602002602001015182815181106110c5576110c5614608565b60200260200101906001600160a01b031690816001600160a01b03168152505081611125576111228360400151828151811061110357611103614608565b6020026020010151602001518789518151811061105057611050614608565b91505b61112e81614679565b905061106b565b508061115e5784838151811061114d5761114d614608565b602002602001016060815250611171565b60408201515161116e9085614666565b93505b50508061117d90614679565b9050610fa7565b50611197611192838361239b565b6124a0565b935061121a600160009054906101000a90046001600160a01b03166001600160a01b0316633fc8cef36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111ef573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121391906143c8565b8590612349565b1561123d57610a7c8473eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee61261c565b50505090565b6000546001600160a01b0316331461129d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104d8565b6112b16001600160a01b0384168383612719565b505050565b6000546001600160a01b031633146113105760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104d8565b6001600160a01b03811661138c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104d8565b61139581611aff565b50565b6113dc6040518060a0016040528060006001600160a01b0316815260200160006001600160a01b031681526020016000815260200160008152602001606081525090565b6000806113e885611b67565b919350909150506001600160a01b038216611404575050611575565b6000604051806040016040528061141a60001990565b81526001600160a01b038981166020928301526001546040517fe6b00009000000000000000000000000000000000000000000000000000000008152845160048201529284015182166024840152868216604484015260648301899052929350600092169063e6b0000990608401600060405180830381865afa1580156114a5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114cd9190810190614415565b6040517f04b1ac290000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301529192506000918516906304b1ac2990602401602060405180830381865afa158015611532573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115569190614882565b9050808260400151111561156e575050505050611575565b5093505050505b9392505050565b6060600061158a6002611d89565b905060008167ffffffffffffffff8111156115a7576115a7613e5b565b6040519080825280602002602001820160405280156115da57816020015b60608152602001906001900390816115c55790505b5090506000805b83811015611a295760006115f6600283611d94565b9150506000816001600160a01b0316639c1d060e6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611639573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261166191908101906144f9565b9050600080825167ffffffffffffffff81111561168057611680613e5b565b6040519080825280602002602001820160405280156116a9578160200160208202803683370190505b50905060005b83518110156118ed5760008b611797576040518060400160405280876001600160a01b03166398b575058886815181106116eb576116eb614608565b6020026020010151602001516040518263ffffffff1660e01b815260040161172291906001600160a01b0391909116815260200190565b602060405180830381865afa15801561173f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117639190614882565b815260200186848151811061177a5761177a614608565b6020026020010151602001516001600160a01b03168152506117be565b60405180604001604052806117ab60001990565b81526020018e6001600160a01b03168152505b905060008c6117cd578d6117ec565b8583815181106117df576117df614608565b6020026020010151602001515b6001546040517f285078fc0000000000000000000000000000000000000000000000000000000081528451600482015260208501516001600160a01b039081166024830152808416604483015292935091169063285078fc90606401602060405180830381865afa158015611865573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611889919061489b565b84848151811061189b5761189b614608565b6020026020010190151590811515815250508383815181106118bf576118bf614608565b6020026020010151156118da57846118d681614679565b9550505b5050806118e690614679565b90506116af565b508167ffffffffffffffff81111561190757611907613e5b565b60405190808252806020026020018201604052801561194d57816020015b6040805180820190915260608152600060208201528152602001906001900390816119255790505b5087868151811061196057611960614608565b60209081029190910101526119758287614666565b95506000805b8451811015611a125782818151811061199657611996614608565b602002602001015115611a02578481815181106119b5576119b5614608565b60200260200101518988815181106119cf576119cf614608565b602002602001015183815181106119e8576119e8614608565b602002602001018190525081806119fe90614679565b9250505b611a0b81614679565b905061197b565b50505050505080611a2290614679565b90506115e1565b5061041f8282611db0565b600080611a428685856128c7565b9050611a50868686846129b2565b95506000611a5f878684612a3b565b9050611a6d87878385612a81565b92507fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03861601611aa957611aa98184612c41565b611abd6001600160a01b0386168985612cb8565b505095945050505050565b600061040a600283612dba565b60006115758383612dc6565b6000611af784846001600160a01b038516612e36565b949350505050565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600080600080611b776002611d89565b905060005b81811015611cc9576000611b91600283611d94565b6040517fa5bc29c20000000000000000000000000000000000000000000000000000000081529092506001600160a01b038316915063a5bc29c290611bda908a906004016148b6565b602060405180830381865afa158015611bf7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c1b91906143c8565b95506001600160a01b03861615611cb8576040517f98b575050000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301528216906398b5750590602401602060405180830381865afa158015611c8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611caf9190614882565b94509250611cc9565b50611cc281614679565b9050611b7c565b50509193909250565b6040517f0d25aafe0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301526024820184905282151560448301526000918291871690630d25aafe90606401602060405180830381865afa158015611d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d699190614882565b905083811015611d8057611d7d81856148c9565b91505b50949350505050565b600061040a82612e53565b6000808080611da38686612e5e565b9097909650945050505050565b60608167ffffffffffffffff811115611dcb57611dcb613e5b565b604051908082528060200260200182016040528015611e1157816020015b604080518082019091526060815260006020820152815260200190600190039081611de95790505b5090506000805b8451811015611ec65760005b858281518110611e3657611e36614608565b602002602001015151811015611eb557858281518110611e5857611e58614608565b60200260200101518181518110611e7157611e71614608565b6020026020010151848481518110611e8b57611e8b614608565b60200260200101819052508280611ea190614679565b93505080611eae90614679565b9050611e24565b50611ebf81614679565b9050611e18565b50828114611f03576040517f3726b986000000000000000000000000000000000000000000000000000000008152600481018290526024016104d8565b5092915050565b6000808260600151421115611f4b576040517f559895a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8251611f58908686612093565b602084015184519195509250306001600160a01b03821603611f8c57611f85878787868860800151611a34565b915061203f565b611f9f6001600160a01b03841688612e89565b9150806001600160a01b03166324a98f1134898989888a608001516040518763ffffffff1660e01b8152600401611fda9594939291906148dc565b60206040518083038185885af1158015611ff8573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061201d9190614882565b50816120326001600160a01b03851689612e89565b61203c91906148c9565b91505b836040015182101561207d576040517f42301c2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5094509492505050565b60006115758383612f56565b6000346000036121e0576120af836001600160a01b0316612f73565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301528416906370a0823190602401602060405180830381865afa15801561210e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121329190614882565b90506121496001600160a01b038416338685613019565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301528291908516906370a0823190602401602060405180830381865afa1580156121ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121cf9190614882565b6121d991906148c9565b9050611575565b6001600160a01b03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14612236576040517f2eac7efb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b34821461226f576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50349392505050565b60606115758383604051806060016040528060278152602001614af36027913961306a565b6060815167ffffffffffffffff8111156122b9576122b9613e5b565b6040519080825280602002602001820160405280156122e2578160200160208202803683370190505b50905060005b8251811015610b1c5782818151811061230357612303614608565b60200260200101516020015182828151811061232157612321614608565b6001600160a01b039092166020928302919091019091015261234281614679565b90506122e8565b6000805b8351811015611f035783818151811061236857612368614608565b60200260200101516001600160a01b0316836001600160a01b031614915081611f035761239481614679565b905061234d565b60608167ffffffffffffffff8111156123b6576123b6613e5b565b6040519080825280602002602001820160405280156123df578160200160208202803683370190505b5090506000805b8451811015611ec65760005b85828151811061240457612404614608565b60200260200101515181101561248f5785828151811061242657612426614608565b6020026020010151818151811061243f5761243f614608565b602002602001015184848151811061245957612459614608565b6001600160a01b03909216602092830291909101909101528261247b81614679565b9350508061248890614679565b90506123f2565b5061249981614679565b90506123e6565b60606000825167ffffffffffffffff8111156124be576124be613e5b565b6040519080825280602002602001820160405280156124e7578160200160208202803683370190505b5090506000805b845181101561256d57600085828151811061250b5761250b614608565b6020026020010151905061251f8482612349565b61255c578084848151811061253657612536614608565b6001600160a01b03909216602092830291909101909101528261255881614679565b9350505b5061256681614679565b90506124ee565b508067ffffffffffffffff81111561258757612587613e5b565b6040519080825280602002602001820160405280156125b0578160200160208202803683370190505b50925060005b81811015612614578281815181106125d0576125d0614608565b60200260200101518482815181106125ea576125ea614608565b6001600160a01b03909216602092830291909101909101528061260c81614679565b9150506125b6565b505050919050565b60608251600161262c9190614666565b67ffffffffffffffff81111561264457612644613e5b565b60405190808252806020026020018201604052801561266d578160200160208202803683370190505b50905060005b83518110156126d25783818151811061268e5761268e614608565b60200260200101518282815181106126a8576126a8614608565b6001600160a01b0390921660209283029190910190910152806126ca81614679565b915050612673565b508181600183516126e391906148c9565b815181106126f3576126f3614608565b60200260200101906001600160a01b031690816001600160a01b03168152505092915050565b8015806127ac57506040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03838116602483015284169063dd62ed3e90604401602060405180830381865afa158015612786573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127aa9190614882565b155b61281e5760405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e63650000000000000000000060648201526084016104d8565b6040516001600160a01b0383166024820152604481018290526112b19084907f095ea7b300000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613154565b604080516080810182526000808252602082018190529181018290526060810191909152826001600160a01b0316846001600160a01b031603612936576040517f0b839a1f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8180602001905181019061294a9190614927565b60208101519091506001600160a01b031615801561297b5750600381516003811115612978576129786149ac565b14155b15611575576040517f76ecffc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03861601612a01576129f083836001613239565b90506129fc81856132fe565b611af7565b50833415611af7576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03841601612a79576121d984836000613239565b509092915050565b6000600382516003811115612a9857612a986149ac565b03612aa4575082611af7565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038416906370a0823190602401602060405180830381865afa158015612b01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b259190614882565b6020830151909150612b42906001600160a01b038716908661338f565b600082516003811115612b5757612b576149ac565b03612b7157612b6c8260200151838686613473565b612bab565b600182516003811115612b8657612b866149ac565b03612b9b57612b6c82602001518386866135f8565b612bab8260200151838686613762565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281906001600160a01b038516906370a0823190602401602060405180830381865afa158015612c0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c2e9190614882565b612c3891906148c9565b95945050505050565b6040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018290526001600160a01b03831690632e1a7d4d90602401600060405180830381600087803b158015612c9c57600080fd5b505af1158015612cb0573d6000803e3d6000fd5b505050505050565b306001600160a01b03831603612ccd57505050565b7fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03841601612da6576000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114612d4a576040519150601f19603f3d011682016040523d82523d6000602084013e612d4f565b606091505b5050905080612da05760405162461bcd60e51b815260206004820152601360248201527f455448207472616e73666572206661696c65640000000000000000000000000060448201526064016104d8565b50505050565b6112b16001600160a01b0384168383613895565b600061157583836138de565b600081815260028301602052604081205480151580612dea5750612dea84846138de565b6115755760405162461bcd60e51b815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b6579000060448201526064016104d8565b60008281526002840160205260408120829055611af784846138ea565b600061040a826138f6565b60008080612e6c8585613900565b600081815260029690960160205260409095205494959350505050565b60007fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03841601612ecc57506001600160a01b0381163161040a565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301528416906370a0823190602401602060405180830381865afa158015612f2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f4f9190614882565b905061040a565b60008181526002830160205260408120819055611575838361390c565b7fffffffffffffffffffffffff11111111111111111111111111111111111111126001600160a01b03821601612fd5576040517f7f523fe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b03163b600003611395576040517f7f523fe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516001600160a01b0380851660248301528316604482015260648101829052612da09085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401612863565b60606001600160a01b0384163b6130e95760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084016104d8565b600080856001600160a01b03168560405161310491906149db565b600060405180830381855af49150503d806000811461313f576040519150601f19603f3d011682016040523d82523d6000602084013e613144565b606091505b509150915061041f828286613918565b60006131a9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166139519092919063ffffffff16565b8051909150156112b157808060200190518101906131c7919061489b565b6112b15760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016104d8565b6000600383516003811115613250576132506149ac565b0361325c575082611575565b82602001516001600160a01b03166382b866008361327e578460600151613284565b84604001515b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815260ff9091166004820152602401602060405180830381865afa1580156132da573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611af791906143c8565b348114613337576040517f81de0bf300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561337257600080fd5b505af1158015613386573d6000803e3d6000fd5b50505050505050565b6040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156133f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061341d9190614882565b6134279190614666565b6040516001600160a01b038516602482015260448101829052909150612da09085907f095ea7b30000000000000000000000000000000000000000000000000000000090606401612863565b60608301516040517f82b8660000000000000000000000000000000000000000000000000000000000815260ff90911660048201526001600160a01b0382811691908616906382b8660090602401602060405180830381865afa1580156134de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061350291906143c8565b6001600160a01b031614613542576040517f28716b9200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080840151606085015191517f9169558600000000000000000000000000000000000000000000000000000000815260ff918216600482015291166024820152604481018390526000606482015260001960848201526001600160a01b0385169063916955869060a4015b6020604051808303816000875af11580156135cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135f19190614882565b5050505050565b600061360385613960565b9050600061361086613a00565b9050826001600160a01b0316816001600160a01b03161461365d576040517f28716b9200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008267ffffffffffffffff81111561367857613678613e5b565b6040519080825280602002602001820160405280156136a1578160200160208202803683370190505b5090508481876040015160ff16815181106136be576136be614608565b60209081029190910101526040517f4d49e87d0000000000000000000000000000000000000000000000000000000081526001600160a01b03881690634d49e87d90613715908490600090600019906004016149f7565b6020604051808303816000875af1158015613734573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137589190614882565b5050505050505050565b60608301516040517f82b8660000000000000000000000000000000000000000000000000000000000815260ff90911660048201526001600160a01b0382811691908616906382b8660090602401602060405180830381865afa1580156137cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137f191906143c8565b6001600160a01b031614613831576040517f28716b9200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608301516040517f3e3a15600000000000000000000000000000000000000000000000000000000081526004810184905260ff90911660248201526000604482015260001960648201526001600160a01b03851690633e3a1560906084016135ae565b6040516001600160a01b0383166024820152604481018290526112b19084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401612863565b60006115758383613a70565b60006115758383613a88565b600061040a825490565b60006115758383613ad7565b60006115758383613b01565b60608315613927575081611575565b8251156139375782518084602001fd5b8160405162461bcd60e51b81526004016104d891906148b6565b6060611af78484600085613bf4565b6000805b6040517f82b8660000000000000000000000000000000000000000000000000000000000815260ff821660048201526001600160a01b038416906382b8660090602401602060405180830381865afa9250505080156139e0575060408051601f3d908101601f191682019092526139dd918101906143c8565b60015b6139ef578060ff169150610b1c565b506139f981614a44565b9050613964565b6000816001600160a01b0316635fd65f0f6040518163ffffffff1660e01b815260040160e060405180830381865afa158015613a40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a649190614a63565b98975050505050505050565b60008181526001830160205260408120541515611575565b6000818152600183016020526040812054613acf5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561040a565b50600061040a565b6000826000018281548110613aee57613aee614608565b9060005260206000200154905092915050565b60008181526001830160205260408120548015613bea576000613b256001836148c9565b8554909150600090613b39906001906148c9565b9050818114613b9e576000866000018281548110613b5957613b59614608565b9060005260206000200154905080876000018481548110613b7c57613b7c614608565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080613baf57613baf614ac3565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061040a565b600091505061040a565b606082471015613c6c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016104d8565b6001600160a01b0385163b613cc35760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104d8565b600080866001600160a01b03168587604051613cdf91906149db565b60006040518083038185875af1925050503d8060008114613d1c576040519150601f19603f3d011682016040523d82523d6000602084013e613d21565b606091505b5091509150613d31828286613918565b979650505050505050565b6001600160a01b038116811461139557600080fd5b600060208284031215613d6357600080fd5b813561157581613d3c565b60005b83811015613d89578181015183820152602001613d71565b50506000910152565b60008151808452613daa816020860160208601613d6e565b601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015613e4d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc089840301855281518051878552613e2788860182613d92565b918901516001600160a01b03169489019490945294870194925090860190600101613de5565b509098975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613ead57613ead613e5b565b60405290565b60405160a0810167ffffffffffffffff81118282101715613ead57613ead613e5b565b6040516060810167ffffffffffffffff81118282101715613ead57613ead613e5b565b604051601f8201601f1916810167ffffffffffffffff81118282101715613f2257613f22613e5b565b604052919050565b600067ffffffffffffffff821115613f4457613f44613e5b565b50601f01601f191660200190565b600082601f830112613f6357600080fd5b8135613f76613f7182613f2a565b613ef9565b818152846020838601011115613f8b57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a08688031215613fc057600080fd5b8535613fcb81613d3c565b94506020860135613fdb81613d3c565b9350604086013592506060860135613ff281613d3c565b9150608086013567ffffffffffffffff81111561400e57600080fd5b61401a88828901613f52565b9150509295509295909350565b60006020828403121561403957600080fd5b5035919050565b6000806040838503121561405357600080fd5b82359150602083013561406581613d3c565b809150509250929050565b6000806040838503121561408357600080fd5b823567ffffffffffffffff8082111561409b57600080fd5b90840190604082870312156140af57600080fd5b6140b7613e8a565b8235828111156140c657600080fd5b6140d288828601613f52565b8252506020928301358382015293505083013561406581613d3c565b60006001600160a01b03808351168452806020840151166020850152506040820151604084015260608201516060840152608082015160a06080850152611af760a0850182613d92565b60208152600061157560208301846140ee565b600060a0828403121561415d57600080fd5b614165613eb3565b9050813561417281613d3c565b8152602082013561418281613d3c565b806020830152506040820135604082015260608201356060820152608082013567ffffffffffffffff8111156141b757600080fd5b6141c384828501613f52565b60808301525092915050565b600080600080608085870312156141e557600080fd5b84356141f081613d3c565b9350602085013561420081613d3c565b925060408501359150606085013567ffffffffffffffff81111561422357600080fd5b61422f8782880161414b565b91505092959194509250565b600080600080600080600060e0888a03121561425657600080fd5b873561426181613d3c565b96506020880135955060408801359450606088013561427f81613d3c565b93506080880135925060a088013567ffffffffffffffff808211156142a357600080fd5b6142af8b838c0161414b565b935060c08a01359150808211156142c557600080fd5b506142d28a828b0161414b565b91505092959891949750929550565b6020808252825182820181905260009190848201906040850190845b818110156143225783516001600160a01b0316835292840192918401916001016142fd565b50909695505050505050565b60008060006060848603121561434357600080fd5b833561434e81613d3c565b9250602084013561435e81613d3c565b929592945050506040919091013590565b60008060006060848603121561438457600080fd5b833561438f81613d3c565b9250602084013567ffffffffffffffff8111156143ab57600080fd5b6143b786828701613f52565b925050604084013590509250925092565b6000602082840312156143da57600080fd5b815161157581613d3c565b60006143f3613f7184613f2a565b905082815283838301111561440757600080fd5b611575836020830184613d6e565b60006020828403121561442757600080fd5b815167ffffffffffffffff8082111561443f57600080fd5b9083019060a0828603121561445357600080fd5b61445b613eb3565b825161446681613d3c565b8152602083015161447681613d3c565b8060208301525060408301516040820152606083015160608201526080830151828111156144a357600080fd5b80840193505085601f8401126144b857600080fd5b6144c7868451602086016143e5565b608082015295945050505050565b600067ffffffffffffffff8211156144ef576144ef613e5b565b5060051b60200190565b6000602080838503121561450c57600080fd5b825167ffffffffffffffff8082111561452457600080fd5b818501915085601f83011261453857600080fd5b8151614546613f71826144d5565b81815260059190911b8301840190848101908883111561456557600080fd5b8585015b838110156145fb578051858111156145815760008081fd5b86016040818c03601f19018113156145995760008081fd5b6145a1613e8a565b89830151888111156145b35760008081fd5b8301603f81018e136145c55760008081fd5b6145d58e8c8301518584016143e5565b82525091810151916145e683613d3c565b808a0192909252508352918601918601614569565b5098975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561040a5761040a614637565b6000600019820361468c5761468c614637565b5060010190565b60006001600160a01b03808816835286602084015280861660408401525083606083015260a06080830152613d3160a08301846140ee565b805180151581146146db57600080fd5b919050565b6000602082840312156146f257600080fd5b815167ffffffffffffffff8082111561470a57600080fd5b818401915084601f83011261471e57600080fd5b815161472c613f71826144d5565b8082825260208201915060208360051b86010192508783111561474e57600080fd5b602085015b838110156148765780518581111561476a57600080fd5b86016060818b03601f1901121561478057600080fd5b614788613ed6565b602082015161479681613d3c565b815260408201516147a681613d3c565b60208201526060820151878111156147bd57600080fd5b8083019250508a603f8301126147d257600080fd5b60208201516147e3613f71826144d5565b81815260069190911b83016040019060208101908d83111561480457600080fd5b6040850194505b82851015614860576040858f03121561482357600080fd5b61482b613e8a565b614834866146cb565b8152602086015161484481613d3c565b806020830152508083525060208201915060408501945061480b565b6040840152505084525060209283019201614753565b50979650505050505050565b60006020828403121561489457600080fd5b5051919050565b6000602082840312156148ad57600080fd5b611575826146cb565b6020815260006115756020830184613d92565b8181038181111561040a5761040a614637565b60006001600160a01b038088168352808716602084015285604084015280851660608401525060a06080830152613d3160a0830184613d92565b805160ff811681146146db57600080fd5b60006080828403121561493957600080fd5b6040516080810181811067ffffffffffffffff8211171561495c5761495c613e5b565b60405282516004811061496e57600080fd5b8152602083015161497e81613d3c565b602082015261498f60408401614916565b60408201526149a060608401614916565b60608201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600082516149ed818460208701613d6e565b9190910192915050565b606080825284519082018190526000906020906080840190828801845b82811015614a3057815184529284019290840190600101614a14565b505050908301949094525060400152919050565b600060ff821660ff8103614a5a57614a5a614637565b60010192915050565b600080600080600080600060e0888a031215614a7e57600080fd5b875196506020880151955060408801519450606088015193506080880151925060a0880151915060c0880151614ab381613d3c565b8091505092959891949750929550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220fa6a143f537a228bb78bd144a6be8297f48c82927bfb46f5fb3a9a4390edc38f64736f6c63430008110033",
}

// SynapseRouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseRouterV2MetaData.ABI instead.
var SynapseRouterV2ABI = SynapseRouterV2MetaData.ABI

// Deprecated: Use SynapseRouterV2MetaData.Sigs instead.
// SynapseRouterV2FuncSigs maps the 4-byte function signature to its string representation.
var SynapseRouterV2FuncSigs = SynapseRouterV2MetaData.Sigs

// SynapseRouterV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseRouterV2MetaData.Bin instead.
var SynapseRouterV2Bin = SynapseRouterV2MetaData.Bin

// DeploySynapseRouterV2 deploys a new Ethereum contract, binding an instance of SynapseRouterV2 to it.
func DeploySynapseRouterV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseRouterV2, error) {
	parsed, err := SynapseRouterV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseRouterV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseRouterV2{SynapseRouterV2Caller: SynapseRouterV2Caller{contract: contract}, SynapseRouterV2Transactor: SynapseRouterV2Transactor{contract: contract}, SynapseRouterV2Filterer: SynapseRouterV2Filterer{contract: contract}}, nil
}

// SynapseRouterV2 is an auto generated Go binding around an Ethereum contract.
type SynapseRouterV2 struct {
	SynapseRouterV2Caller     // Read-only binding to the contract
	SynapseRouterV2Transactor // Write-only binding to the contract
	SynapseRouterV2Filterer   // Log filterer for contract events
}

// SynapseRouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseRouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseRouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseRouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseRouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseRouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseRouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseRouterV2Session struct {
	Contract     *SynapseRouterV2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseRouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseRouterV2CallerSession struct {
	Contract *SynapseRouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SynapseRouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseRouterV2TransactorSession struct {
	Contract     *SynapseRouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SynapseRouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseRouterV2Raw struct {
	Contract *SynapseRouterV2 // Generic contract binding to access the raw methods on
}

// SynapseRouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseRouterV2CallerRaw struct {
	Contract *SynapseRouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// SynapseRouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseRouterV2TransactorRaw struct {
	Contract *SynapseRouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseRouterV2 creates a new instance of SynapseRouterV2, bound to a specific deployed contract.
func NewSynapseRouterV2(address common.Address, backend bind.ContractBackend) (*SynapseRouterV2, error) {
	contract, err := bindSynapseRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2{SynapseRouterV2Caller: SynapseRouterV2Caller{contract: contract}, SynapseRouterV2Transactor: SynapseRouterV2Transactor{contract: contract}, SynapseRouterV2Filterer: SynapseRouterV2Filterer{contract: contract}}, nil
}

// NewSynapseRouterV2Caller creates a new read-only instance of SynapseRouterV2, bound to a specific deployed contract.
func NewSynapseRouterV2Caller(address common.Address, caller bind.ContractCaller) (*SynapseRouterV2Caller, error) {
	contract, err := bindSynapseRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2Caller{contract: contract}, nil
}

// NewSynapseRouterV2Transactor creates a new write-only instance of SynapseRouterV2, bound to a specific deployed contract.
func NewSynapseRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*SynapseRouterV2Transactor, error) {
	contract, err := bindSynapseRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2Transactor{contract: contract}, nil
}

// NewSynapseRouterV2Filterer creates a new log filterer instance of SynapseRouterV2, bound to a specific deployed contract.
func NewSynapseRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*SynapseRouterV2Filterer, error) {
	contract, err := bindSynapseRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2Filterer{contract: contract}, nil
}

// bindSynapseRouterV2 binds a generic wrapper to an already deployed contract.
func bindSynapseRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseRouterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseRouterV2 *SynapseRouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseRouterV2.Contract.SynapseRouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseRouterV2 *SynapseRouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.SynapseRouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseRouterV2 *SynapseRouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.SynapseRouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseRouterV2 *SynapseRouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseRouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseRouterV2 *SynapseRouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseRouterV2 *SynapseRouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.contract.Transact(opts, method, params...)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[])
func (_SynapseRouterV2 *SynapseRouterV2Caller) GetBridgeTokens(opts *bind.CallOpts) ([]BridgeToken, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "getBridgeTokens")

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[])
func (_SynapseRouterV2 *SynapseRouterV2Session) GetBridgeTokens() ([]BridgeToken, error) {
	return _SynapseRouterV2.Contract.GetBridgeTokens(&_SynapseRouterV2.CallOpts)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[])
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _SynapseRouterV2.Contract.GetBridgeTokens(&_SynapseRouterV2.CallOpts)
}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x7de31c74.
//
// Solidity: function getDestinationAmountOut((string,uint256) request, address tokenOut) view returns((address,address,uint256,uint256,bytes) destQuery)
func (_SynapseRouterV2 *SynapseRouterV2Caller) GetDestinationAmountOut(opts *bind.CallOpts, request DestRequest, tokenOut common.Address) (SwapQuery, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "getDestinationAmountOut", request, tokenOut)

	if err != nil {
		return *new(SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new(SwapQuery)).(*SwapQuery)

	return out0, err

}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x7de31c74.
//
// Solidity: function getDestinationAmountOut((string,uint256) request, address tokenOut) view returns((address,address,uint256,uint256,bytes) destQuery)
func (_SynapseRouterV2 *SynapseRouterV2Session) GetDestinationAmountOut(request DestRequest, tokenOut common.Address) (SwapQuery, error) {
	return _SynapseRouterV2.Contract.GetDestinationAmountOut(&_SynapseRouterV2.CallOpts, request, tokenOut)
}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x7de31c74.
//
// Solidity: function getDestinationAmountOut((string,uint256) request, address tokenOut) view returns((address,address,uint256,uint256,bytes) destQuery)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) GetDestinationAmountOut(request DestRequest, tokenOut common.Address) (SwapQuery, error) {
	return _SynapseRouterV2.Contract.GetDestinationAmountOut(&_SynapseRouterV2.CallOpts, request, tokenOut)
}

// GetDestinationBridgeTokens is a free data retrieval call binding the contract method 0x1d04879b.
//
// Solidity: function getDestinationBridgeTokens(address tokenOut) view returns((string,address)[] destTokens)
func (_SynapseRouterV2 *SynapseRouterV2Caller) GetDestinationBridgeTokens(opts *bind.CallOpts, tokenOut common.Address) ([]BridgeToken, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "getDestinationBridgeTokens", tokenOut)

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetDestinationBridgeTokens is a free data retrieval call binding the contract method 0x1d04879b.
//
// Solidity: function getDestinationBridgeTokens(address tokenOut) view returns((string,address)[] destTokens)
func (_SynapseRouterV2 *SynapseRouterV2Session) GetDestinationBridgeTokens(tokenOut common.Address) ([]BridgeToken, error) {
	return _SynapseRouterV2.Contract.GetDestinationBridgeTokens(&_SynapseRouterV2.CallOpts, tokenOut)
}

// GetDestinationBridgeTokens is a free data retrieval call binding the contract method 0x1d04879b.
//
// Solidity: function getDestinationBridgeTokens(address tokenOut) view returns((string,address)[] destTokens)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) GetDestinationBridgeTokens(tokenOut common.Address) ([]BridgeToken, error) {
	return _SynapseRouterV2.Contract.GetDestinationBridgeTokens(&_SynapseRouterV2.CallOpts, tokenOut)
}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0xf533941d.
//
// Solidity: function getOriginAmountOut(address tokenIn, string tokenSymbol, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) originQuery)
func (_SynapseRouterV2 *SynapseRouterV2Caller) GetOriginAmountOut(opts *bind.CallOpts, tokenIn common.Address, tokenSymbol string, amountIn *big.Int) (SwapQuery, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "getOriginAmountOut", tokenIn, tokenSymbol, amountIn)

	if err != nil {
		return *new(SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new(SwapQuery)).(*SwapQuery)

	return out0, err

}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0xf533941d.
//
// Solidity: function getOriginAmountOut(address tokenIn, string tokenSymbol, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) originQuery)
func (_SynapseRouterV2 *SynapseRouterV2Session) GetOriginAmountOut(tokenIn common.Address, tokenSymbol string, amountIn *big.Int) (SwapQuery, error) {
	return _SynapseRouterV2.Contract.GetOriginAmountOut(&_SynapseRouterV2.CallOpts, tokenIn, tokenSymbol, amountIn)
}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0xf533941d.
//
// Solidity: function getOriginAmountOut(address tokenIn, string tokenSymbol, uint256 amountIn) view returns((address,address,uint256,uint256,bytes) originQuery)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) GetOriginAmountOut(tokenIn common.Address, tokenSymbol string, amountIn *big.Int) (SwapQuery, error) {
	return _SynapseRouterV2.Contract.GetOriginAmountOut(&_SynapseRouterV2.CallOpts, tokenIn, tokenSymbol, amountIn)
}

// GetOriginBridgeTokens is a free data retrieval call binding the contract method 0x3e811a5c.
//
// Solidity: function getOriginBridgeTokens(address tokenIn) view returns((string,address)[] originTokens)
func (_SynapseRouterV2 *SynapseRouterV2Caller) GetOriginBridgeTokens(opts *bind.CallOpts, tokenIn common.Address) ([]BridgeToken, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "getOriginBridgeTokens", tokenIn)

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetOriginBridgeTokens is a free data retrieval call binding the contract method 0x3e811a5c.
//
// Solidity: function getOriginBridgeTokens(address tokenIn) view returns((string,address)[] originTokens)
func (_SynapseRouterV2 *SynapseRouterV2Session) GetOriginBridgeTokens(tokenIn common.Address) ([]BridgeToken, error) {
	return _SynapseRouterV2.Contract.GetOriginBridgeTokens(&_SynapseRouterV2.CallOpts, tokenIn)
}

// GetOriginBridgeTokens is a free data retrieval call binding the contract method 0x3e811a5c.
//
// Solidity: function getOriginBridgeTokens(address tokenIn) view returns((string,address)[] originTokens)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) GetOriginBridgeTokens(tokenIn common.Address) ([]BridgeToken, error) {
	return _SynapseRouterV2.Contract.GetOriginBridgeTokens(&_SynapseRouterV2.CallOpts, tokenIn)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] supportedTokens)
func (_SynapseRouterV2 *SynapseRouterV2Caller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] supportedTokens)
func (_SynapseRouterV2 *SynapseRouterV2Session) GetSupportedTokens() ([]common.Address, error) {
	return _SynapseRouterV2.Contract.GetSupportedTokens(&_SynapseRouterV2.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] supportedTokens)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _SynapseRouterV2.Contract.GetSupportedTokens(&_SynapseRouterV2.CallOpts)
}

// IdToModule is a free data retrieval call binding the contract method 0x53e2e8e7.
//
// Solidity: function idToModule(bytes32 moduleId) view returns(address bridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2Caller) IdToModule(opts *bind.CallOpts, moduleId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "idToModule", moduleId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdToModule is a free data retrieval call binding the contract method 0x53e2e8e7.
//
// Solidity: function idToModule(bytes32 moduleId) view returns(address bridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2Session) IdToModule(moduleId [32]byte) (common.Address, error) {
	return _SynapseRouterV2.Contract.IdToModule(&_SynapseRouterV2.CallOpts, moduleId)
}

// IdToModule is a free data retrieval call binding the contract method 0x53e2e8e7.
//
// Solidity: function idToModule(bytes32 moduleId) view returns(address bridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) IdToModule(moduleId [32]byte) (common.Address, error) {
	return _SynapseRouterV2.Contract.IdToModule(&_SynapseRouterV2.CallOpts, moduleId)
}

// ModuleToId is a free data retrieval call binding the contract method 0x9f2671fa.
//
// Solidity: function moduleToId(address bridgeModule) view returns(bytes32 moduleId)
func (_SynapseRouterV2 *SynapseRouterV2Caller) ModuleToId(opts *bind.CallOpts, bridgeModule common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "moduleToId", bridgeModule)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ModuleToId is a free data retrieval call binding the contract method 0x9f2671fa.
//
// Solidity: function moduleToId(address bridgeModule) view returns(bytes32 moduleId)
func (_SynapseRouterV2 *SynapseRouterV2Session) ModuleToId(bridgeModule common.Address) ([32]byte, error) {
	return _SynapseRouterV2.Contract.ModuleToId(&_SynapseRouterV2.CallOpts, bridgeModule)
}

// ModuleToId is a free data retrieval call binding the contract method 0x9f2671fa.
//
// Solidity: function moduleToId(address bridgeModule) view returns(bytes32 moduleId)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) ModuleToId(bridgeModule common.Address) ([32]byte, error) {
	return _SynapseRouterV2.Contract.ModuleToId(&_SynapseRouterV2.CallOpts, bridgeModule)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseRouterV2 *SynapseRouterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseRouterV2 *SynapseRouterV2Session) Owner() (common.Address, error) {
	return _SynapseRouterV2.Contract.Owner(&_SynapseRouterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) Owner() (common.Address, error) {
	return _SynapseRouterV2.Contract.Owner(&_SynapseRouterV2.CallOpts)
}

// SwapQuoter is a free data retrieval call binding the contract method 0x34474c8c.
//
// Solidity: function swapQuoter() view returns(address)
func (_SynapseRouterV2 *SynapseRouterV2Caller) SwapQuoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseRouterV2.contract.Call(opts, &out, "swapQuoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapQuoter is a free data retrieval call binding the contract method 0x34474c8c.
//
// Solidity: function swapQuoter() view returns(address)
func (_SynapseRouterV2 *SynapseRouterV2Session) SwapQuoter() (common.Address, error) {
	return _SynapseRouterV2.Contract.SwapQuoter(&_SynapseRouterV2.CallOpts)
}

// SwapQuoter is a free data retrieval call binding the contract method 0x34474c8c.
//
// Solidity: function swapQuoter() view returns(address)
func (_SynapseRouterV2 *SynapseRouterV2CallerSession) SwapQuoter() (common.Address, error) {
	return _SynapseRouterV2.Contract.SwapQuoter(&_SynapseRouterV2.CallOpts)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_SynapseRouterV2 *SynapseRouterV2Transactor) AdapterSwap(opts *bind.TransactOpts, recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "adapterSwap", recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_SynapseRouterV2 *SynapseRouterV2Session) AdapterSwap(recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.AdapterSwap(&_SynapseRouterV2.TransactOpts, recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address recipient, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) AdapterSwap(recipient common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.AdapterSwap(&_SynapseRouterV2.TransactOpts, recipient, tokenIn, amountIn, tokenOut, rawParams)
}

// BridgeViaSynapse is a paid mutator transaction binding the contract method 0xc95fafd2.
//
// Solidity: function bridgeViaSynapse(address to, uint256 chainId, bytes32 moduleId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) BridgeViaSynapse(opts *bind.TransactOpts, to common.Address, chainId *big.Int, moduleId [32]byte, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "bridgeViaSynapse", to, chainId, moduleId, token, amount, originQuery, destQuery)
}

// BridgeViaSynapse is a paid mutator transaction binding the contract method 0xc95fafd2.
//
// Solidity: function bridgeViaSynapse(address to, uint256 chainId, bytes32 moduleId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) BridgeViaSynapse(to common.Address, chainId *big.Int, moduleId [32]byte, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.BridgeViaSynapse(&_SynapseRouterV2.TransactOpts, to, chainId, moduleId, token, amount, originQuery, destQuery)
}

// BridgeViaSynapse is a paid mutator transaction binding the contract method 0xc95fafd2.
//
// Solidity: function bridgeViaSynapse(address to, uint256 chainId, bytes32 moduleId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) BridgeViaSynapse(to common.Address, chainId *big.Int, moduleId [32]byte, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.BridgeViaSynapse(&_SynapseRouterV2.TransactOpts, to, chainId, moduleId, token, amount, originQuery, destQuery)
}

// ConnectBridgeModule is a paid mutator transaction binding the contract method 0xb3bce952.
//
// Solidity: function connectBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) ConnectBridgeModule(opts *bind.TransactOpts, moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "connectBridgeModule", moduleId, bridgeModule)
}

// ConnectBridgeModule is a paid mutator transaction binding the contract method 0xb3bce952.
//
// Solidity: function connectBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) ConnectBridgeModule(moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.ConnectBridgeModule(&_SynapseRouterV2.TransactOpts, moduleId, bridgeModule)
}

// ConnectBridgeModule is a paid mutator transaction binding the contract method 0xb3bce952.
//
// Solidity: function connectBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) ConnectBridgeModule(moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.ConnectBridgeModule(&_SynapseRouterV2.TransactOpts, moduleId, bridgeModule)
}

// DisconnectBridgeModule is a paid mutator transaction binding the contract method 0xb68e4302.
//
// Solidity: function disconnectBridgeModule(bytes32 moduleId) returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) DisconnectBridgeModule(opts *bind.TransactOpts, moduleId [32]byte) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "disconnectBridgeModule", moduleId)
}

// DisconnectBridgeModule is a paid mutator transaction binding the contract method 0xb68e4302.
//
// Solidity: function disconnectBridgeModule(bytes32 moduleId) returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) DisconnectBridgeModule(moduleId [32]byte) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.DisconnectBridgeModule(&_SynapseRouterV2.TransactOpts, moduleId)
}

// DisconnectBridgeModule is a paid mutator transaction binding the contract method 0xb68e4302.
//
// Solidity: function disconnectBridgeModule(bytes32 moduleId) returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) DisconnectBridgeModule(moduleId [32]byte) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.DisconnectBridgeModule(&_SynapseRouterV2.TransactOpts, moduleId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.RenounceOwnership(&_SynapseRouterV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.RenounceOwnership(&_SynapseRouterV2.TransactOpts)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) SetAllowance(opts *bind.TransactOpts, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "setAllowance", token, spender, amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) SetAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.SetAllowance(&_SynapseRouterV2.TransactOpts, token, spender, amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) SetAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.SetAllowance(&_SynapseRouterV2.TransactOpts, token, spender, amount)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) SetSwapQuoter(opts *bind.TransactOpts, _swapQuoter common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "setSwapQuoter", _swapQuoter)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) SetSwapQuoter(_swapQuoter common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.SetSwapQuoter(&_SynapseRouterV2.TransactOpts, _swapQuoter)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) SetSwapQuoter(_swapQuoter common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.SetSwapQuoter(&_SynapseRouterV2.TransactOpts, _swapQuoter)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_SynapseRouterV2 *SynapseRouterV2Transactor) Swap(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "swap", to, token, amount, query)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_SynapseRouterV2 *SynapseRouterV2Session) Swap(to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.Swap(&_SynapseRouterV2.TransactOpts, to, token, amount, query)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) Swap(to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.Swap(&_SynapseRouterV2.TransactOpts, to, token, amount, query)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.TransferOwnership(&_SynapseRouterV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.TransferOwnership(&_SynapseRouterV2.TransactOpts, newOwner)
}

// UpdateBridgeModule is a paid mutator transaction binding the contract method 0x70a1cdc9.
//
// Solidity: function updateBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) UpdateBridgeModule(opts *bind.TransactOpts, moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.Transact(opts, "updateBridgeModule", moduleId, bridgeModule)
}

// UpdateBridgeModule is a paid mutator transaction binding the contract method 0x70a1cdc9.
//
// Solidity: function updateBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) UpdateBridgeModule(moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.UpdateBridgeModule(&_SynapseRouterV2.TransactOpts, moduleId, bridgeModule)
}

// UpdateBridgeModule is a paid mutator transaction binding the contract method 0x70a1cdc9.
//
// Solidity: function updateBridgeModule(bytes32 moduleId, address bridgeModule) returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) UpdateBridgeModule(moduleId [32]byte, bridgeModule common.Address) (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.UpdateBridgeModule(&_SynapseRouterV2.TransactOpts, moduleId, bridgeModule)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SynapseRouterV2 *SynapseRouterV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseRouterV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SynapseRouterV2 *SynapseRouterV2Session) Receive() (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.Receive(&_SynapseRouterV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SynapseRouterV2 *SynapseRouterV2TransactorSession) Receive() (*types.Transaction, error) {
	return _SynapseRouterV2.Contract.Receive(&_SynapseRouterV2.TransactOpts)
}

// SynapseRouterV2ModuleConnectedIterator is returned from FilterModuleConnected and is used to iterate over the raw logs and unpacked data for ModuleConnected events raised by the SynapseRouterV2 contract.
type SynapseRouterV2ModuleConnectedIterator struct {
	Event *SynapseRouterV2ModuleConnected // Event containing the contract specifics and raw log

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
func (it *SynapseRouterV2ModuleConnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseRouterV2ModuleConnected)
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
		it.Event = new(SynapseRouterV2ModuleConnected)
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
func (it *SynapseRouterV2ModuleConnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseRouterV2ModuleConnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseRouterV2ModuleConnected represents a ModuleConnected event raised by the SynapseRouterV2 contract.
type SynapseRouterV2ModuleConnected struct {
	ModuleId     [32]byte
	BridgeModule common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterModuleConnected is a free log retrieval operation binding the contract event 0xc92d33ac2d951a5a8265420d37e1664f111f74de9207f89f4c772142ec3aa5e3.
//
// Solidity: event ModuleConnected(bytes32 indexed moduleId, address bridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) FilterModuleConnected(opts *bind.FilterOpts, moduleId [][32]byte) (*SynapseRouterV2ModuleConnectedIterator, error) {

	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _SynapseRouterV2.contract.FilterLogs(opts, "ModuleConnected", moduleIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2ModuleConnectedIterator{contract: _SynapseRouterV2.contract, event: "ModuleConnected", logs: logs, sub: sub}, nil
}

// WatchModuleConnected is a free log subscription operation binding the contract event 0xc92d33ac2d951a5a8265420d37e1664f111f74de9207f89f4c772142ec3aa5e3.
//
// Solidity: event ModuleConnected(bytes32 indexed moduleId, address bridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) WatchModuleConnected(opts *bind.WatchOpts, sink chan<- *SynapseRouterV2ModuleConnected, moduleId [][32]byte) (event.Subscription, error) {

	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _SynapseRouterV2.contract.WatchLogs(opts, "ModuleConnected", moduleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseRouterV2ModuleConnected)
				if err := _SynapseRouterV2.contract.UnpackLog(event, "ModuleConnected", log); err != nil {
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

// ParseModuleConnected is a log parse operation binding the contract event 0xc92d33ac2d951a5a8265420d37e1664f111f74de9207f89f4c772142ec3aa5e3.
//
// Solidity: event ModuleConnected(bytes32 indexed moduleId, address bridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) ParseModuleConnected(log types.Log) (*SynapseRouterV2ModuleConnected, error) {
	event := new(SynapseRouterV2ModuleConnected)
	if err := _SynapseRouterV2.contract.UnpackLog(event, "ModuleConnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseRouterV2ModuleDisconnectedIterator is returned from FilterModuleDisconnected and is used to iterate over the raw logs and unpacked data for ModuleDisconnected events raised by the SynapseRouterV2 contract.
type SynapseRouterV2ModuleDisconnectedIterator struct {
	Event *SynapseRouterV2ModuleDisconnected // Event containing the contract specifics and raw log

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
func (it *SynapseRouterV2ModuleDisconnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseRouterV2ModuleDisconnected)
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
		it.Event = new(SynapseRouterV2ModuleDisconnected)
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
func (it *SynapseRouterV2ModuleDisconnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseRouterV2ModuleDisconnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseRouterV2ModuleDisconnected represents a ModuleDisconnected event raised by the SynapseRouterV2 contract.
type SynapseRouterV2ModuleDisconnected struct {
	ModuleId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModuleDisconnected is a free log retrieval operation binding the contract event 0x594a6e72239daa84da65d4bbeb00fa0ccb8c579b2fd6b5601c82baf0614894cc.
//
// Solidity: event ModuleDisconnected(bytes32 indexed moduleId)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) FilterModuleDisconnected(opts *bind.FilterOpts, moduleId [][32]byte) (*SynapseRouterV2ModuleDisconnectedIterator, error) {

	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _SynapseRouterV2.contract.FilterLogs(opts, "ModuleDisconnected", moduleIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2ModuleDisconnectedIterator{contract: _SynapseRouterV2.contract, event: "ModuleDisconnected", logs: logs, sub: sub}, nil
}

// WatchModuleDisconnected is a free log subscription operation binding the contract event 0x594a6e72239daa84da65d4bbeb00fa0ccb8c579b2fd6b5601c82baf0614894cc.
//
// Solidity: event ModuleDisconnected(bytes32 indexed moduleId)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) WatchModuleDisconnected(opts *bind.WatchOpts, sink chan<- *SynapseRouterV2ModuleDisconnected, moduleId [][32]byte) (event.Subscription, error) {

	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _SynapseRouterV2.contract.WatchLogs(opts, "ModuleDisconnected", moduleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseRouterV2ModuleDisconnected)
				if err := _SynapseRouterV2.contract.UnpackLog(event, "ModuleDisconnected", log); err != nil {
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

// ParseModuleDisconnected is a log parse operation binding the contract event 0x594a6e72239daa84da65d4bbeb00fa0ccb8c579b2fd6b5601c82baf0614894cc.
//
// Solidity: event ModuleDisconnected(bytes32 indexed moduleId)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) ParseModuleDisconnected(log types.Log) (*SynapseRouterV2ModuleDisconnected, error) {
	event := new(SynapseRouterV2ModuleDisconnected)
	if err := _SynapseRouterV2.contract.UnpackLog(event, "ModuleDisconnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseRouterV2ModuleUpdatedIterator is returned from FilterModuleUpdated and is used to iterate over the raw logs and unpacked data for ModuleUpdated events raised by the SynapseRouterV2 contract.
type SynapseRouterV2ModuleUpdatedIterator struct {
	Event *SynapseRouterV2ModuleUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseRouterV2ModuleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseRouterV2ModuleUpdated)
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
		it.Event = new(SynapseRouterV2ModuleUpdated)
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
func (it *SynapseRouterV2ModuleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseRouterV2ModuleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseRouterV2ModuleUpdated represents a ModuleUpdated event raised by the SynapseRouterV2 contract.
type SynapseRouterV2ModuleUpdated struct {
	ModuleId        [32]byte
	OldBridgeModule common.Address
	NewBridgeModule common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModuleUpdated is a free log retrieval operation binding the contract event 0xbe59b1ad7b24549601b98854029a8be9cd632ee55e4472692aa2542e668496e0.
//
// Solidity: event ModuleUpdated(bytes32 indexed moduleId, address oldBridgeModule, address newBridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) FilterModuleUpdated(opts *bind.FilterOpts, moduleId [][32]byte) (*SynapseRouterV2ModuleUpdatedIterator, error) {

	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _SynapseRouterV2.contract.FilterLogs(opts, "ModuleUpdated", moduleIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2ModuleUpdatedIterator{contract: _SynapseRouterV2.contract, event: "ModuleUpdated", logs: logs, sub: sub}, nil
}

// WatchModuleUpdated is a free log subscription operation binding the contract event 0xbe59b1ad7b24549601b98854029a8be9cd632ee55e4472692aa2542e668496e0.
//
// Solidity: event ModuleUpdated(bytes32 indexed moduleId, address oldBridgeModule, address newBridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) WatchModuleUpdated(opts *bind.WatchOpts, sink chan<- *SynapseRouterV2ModuleUpdated, moduleId [][32]byte) (event.Subscription, error) {

	var moduleIdRule []interface{}
	for _, moduleIdItem := range moduleId {
		moduleIdRule = append(moduleIdRule, moduleIdItem)
	}

	logs, sub, err := _SynapseRouterV2.contract.WatchLogs(opts, "ModuleUpdated", moduleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseRouterV2ModuleUpdated)
				if err := _SynapseRouterV2.contract.UnpackLog(event, "ModuleUpdated", log); err != nil {
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

// ParseModuleUpdated is a log parse operation binding the contract event 0xbe59b1ad7b24549601b98854029a8be9cd632ee55e4472692aa2542e668496e0.
//
// Solidity: event ModuleUpdated(bytes32 indexed moduleId, address oldBridgeModule, address newBridgeModule)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) ParseModuleUpdated(log types.Log) (*SynapseRouterV2ModuleUpdated, error) {
	event := new(SynapseRouterV2ModuleUpdated)
	if err := _SynapseRouterV2.contract.UnpackLog(event, "ModuleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseRouterV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynapseRouterV2 contract.
type SynapseRouterV2OwnershipTransferredIterator struct {
	Event *SynapseRouterV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SynapseRouterV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseRouterV2OwnershipTransferred)
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
		it.Event = new(SynapseRouterV2OwnershipTransferred)
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
func (it *SynapseRouterV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseRouterV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseRouterV2OwnershipTransferred represents a OwnershipTransferred event raised by the SynapseRouterV2 contract.
type SynapseRouterV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynapseRouterV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseRouterV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2OwnershipTransferredIterator{contract: _SynapseRouterV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynapseRouterV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseRouterV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseRouterV2OwnershipTransferred)
				if err := _SynapseRouterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SynapseRouterV2 *SynapseRouterV2Filterer) ParseOwnershipTransferred(log types.Log) (*SynapseRouterV2OwnershipTransferred, error) {
	event := new(SynapseRouterV2OwnershipTransferred)
	if err := _SynapseRouterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseRouterV2QuoterSetIterator is returned from FilterQuoterSet and is used to iterate over the raw logs and unpacked data for QuoterSet events raised by the SynapseRouterV2 contract.
type SynapseRouterV2QuoterSetIterator struct {
	Event *SynapseRouterV2QuoterSet // Event containing the contract specifics and raw log

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
func (it *SynapseRouterV2QuoterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseRouterV2QuoterSet)
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
		it.Event = new(SynapseRouterV2QuoterSet)
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
func (it *SynapseRouterV2QuoterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseRouterV2QuoterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseRouterV2QuoterSet represents a QuoterSet event raised by the SynapseRouterV2 contract.
type SynapseRouterV2QuoterSet struct {
	OldSwapQuoter common.Address
	NewSwapQuoter common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterQuoterSet is a free log retrieval operation binding the contract event 0x2e3d7d02ba3c4bd8b1f8995cd3a23ef0193922ebc4ee23249ead4d0ca2e34c68.
//
// Solidity: event QuoterSet(address oldSwapQuoter, address newSwapQuoter)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) FilterQuoterSet(opts *bind.FilterOpts) (*SynapseRouterV2QuoterSetIterator, error) {

	logs, sub, err := _SynapseRouterV2.contract.FilterLogs(opts, "QuoterSet")
	if err != nil {
		return nil, err
	}
	return &SynapseRouterV2QuoterSetIterator{contract: _SynapseRouterV2.contract, event: "QuoterSet", logs: logs, sub: sub}, nil
}

// WatchQuoterSet is a free log subscription operation binding the contract event 0x2e3d7d02ba3c4bd8b1f8995cd3a23ef0193922ebc4ee23249ead4d0ca2e34c68.
//
// Solidity: event QuoterSet(address oldSwapQuoter, address newSwapQuoter)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) WatchQuoterSet(opts *bind.WatchOpts, sink chan<- *SynapseRouterV2QuoterSet) (event.Subscription, error) {

	logs, sub, err := _SynapseRouterV2.contract.WatchLogs(opts, "QuoterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseRouterV2QuoterSet)
				if err := _SynapseRouterV2.contract.UnpackLog(event, "QuoterSet", log); err != nil {
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

// ParseQuoterSet is a log parse operation binding the contract event 0x2e3d7d02ba3c4bd8b1f8995cd3a23ef0193922ebc4ee23249ead4d0ca2e34c68.
//
// Solidity: event QuoterSet(address oldSwapQuoter, address newSwapQuoter)
func (_SynapseRouterV2 *SynapseRouterV2Filterer) ParseQuoterSet(log types.Log) (*SynapseRouterV2QuoterSet, error) {
	event := new(SynapseRouterV2QuoterSet)
	if err := _SynapseRouterV2.contract.UnpackLog(event, "QuoterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalTokenLibMetaData contains all meta data concerning the UniversalTokenLib contract.
var UniversalTokenLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206119c53f4972e8f550594d951a759848cc63d2318cea815ea66ab7477769324964736f6c63430008110033",
}

// UniversalTokenLibABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalTokenLibMetaData.ABI instead.
var UniversalTokenLibABI = UniversalTokenLibMetaData.ABI

// UniversalTokenLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniversalTokenLibMetaData.Bin instead.
var UniversalTokenLibBin = UniversalTokenLibMetaData.Bin

// DeployUniversalTokenLib deploys a new Ethereum contract, binding an instance of UniversalTokenLib to it.
func DeployUniversalTokenLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniversalTokenLib, error) {
	parsed, err := UniversalTokenLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniversalTokenLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniversalTokenLib{UniversalTokenLibCaller: UniversalTokenLibCaller{contract: contract}, UniversalTokenLibTransactor: UniversalTokenLibTransactor{contract: contract}, UniversalTokenLibFilterer: UniversalTokenLibFilterer{contract: contract}}, nil
}

// UniversalTokenLib is an auto generated Go binding around an Ethereum contract.
type UniversalTokenLib struct {
	UniversalTokenLibCaller     // Read-only binding to the contract
	UniversalTokenLibTransactor // Write-only binding to the contract
	UniversalTokenLibFilterer   // Log filterer for contract events
}

// UniversalTokenLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalTokenLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalTokenLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalTokenLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalTokenLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalTokenLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalTokenLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalTokenLibSession struct {
	Contract     *UniversalTokenLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniversalTokenLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalTokenLibCallerSession struct {
	Contract *UniversalTokenLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniversalTokenLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalTokenLibTransactorSession struct {
	Contract     *UniversalTokenLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniversalTokenLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalTokenLibRaw struct {
	Contract *UniversalTokenLib // Generic contract binding to access the raw methods on
}

// UniversalTokenLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalTokenLibCallerRaw struct {
	Contract *UniversalTokenLibCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalTokenLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalTokenLibTransactorRaw struct {
	Contract *UniversalTokenLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalTokenLib creates a new instance of UniversalTokenLib, bound to a specific deployed contract.
func NewUniversalTokenLib(address common.Address, backend bind.ContractBackend) (*UniversalTokenLib, error) {
	contract, err := bindUniversalTokenLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalTokenLib{UniversalTokenLibCaller: UniversalTokenLibCaller{contract: contract}, UniversalTokenLibTransactor: UniversalTokenLibTransactor{contract: contract}, UniversalTokenLibFilterer: UniversalTokenLibFilterer{contract: contract}}, nil
}

// NewUniversalTokenLibCaller creates a new read-only instance of UniversalTokenLib, bound to a specific deployed contract.
func NewUniversalTokenLibCaller(address common.Address, caller bind.ContractCaller) (*UniversalTokenLibCaller, error) {
	contract, err := bindUniversalTokenLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalTokenLibCaller{contract: contract}, nil
}

// NewUniversalTokenLibTransactor creates a new write-only instance of UniversalTokenLib, bound to a specific deployed contract.
func NewUniversalTokenLibTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalTokenLibTransactor, error) {
	contract, err := bindUniversalTokenLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalTokenLibTransactor{contract: contract}, nil
}

// NewUniversalTokenLibFilterer creates a new log filterer instance of UniversalTokenLib, bound to a specific deployed contract.
func NewUniversalTokenLibFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalTokenLibFilterer, error) {
	contract, err := bindUniversalTokenLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalTokenLibFilterer{contract: contract}, nil
}

// bindUniversalTokenLib binds a generic wrapper to an already deployed contract.
func bindUniversalTokenLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniversalTokenLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalTokenLib *UniversalTokenLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalTokenLib.Contract.UniversalTokenLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalTokenLib *UniversalTokenLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalTokenLib.Contract.UniversalTokenLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalTokenLib *UniversalTokenLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalTokenLib.Contract.UniversalTokenLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalTokenLib *UniversalTokenLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalTokenLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalTokenLib *UniversalTokenLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalTokenLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalTokenLib *UniversalTokenLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalTokenLib.Contract.contract.Transact(opts, method, params...)
}
