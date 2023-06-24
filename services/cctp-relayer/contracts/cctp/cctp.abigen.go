// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cctp

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
)

// BridgeToken is an auto generated low-level Go binding around an user-defined struct.
type BridgeToken struct {
	Symbol string
	Token  common.Address
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ad8c086bc175b1cee1eee8b8d1c0ffd82d2e53d3e2f173762003b2e0fbe2ffb664736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204b082ea0884db8043388c27893ebae2e2b7af80913695d64dfacf7fb7262a05664736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	parsed, err := abi.JSON(strings.NewReader(IDefaultPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// IMessageTransmitterMetaData contains all meta data concerning the IMessageTransmitter contract.
var IMessageTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvailableNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"8371744e": "nextAvailableNonce()",
		"57ecfd28": "receiveMessage(bytes,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
}

// IMessageTransmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageTransmitterMetaData.ABI instead.
var IMessageTransmitterABI = IMessageTransmitterMetaData.ABI

// Deprecated: Use IMessageTransmitterMetaData.Sigs instead.
// IMessageTransmitterFuncSigs maps the 4-byte function signature to its string representation.
var IMessageTransmitterFuncSigs = IMessageTransmitterMetaData.Sigs

// IMessageTransmitter is an auto generated Go binding around an Ethereum contract.
type IMessageTransmitter struct {
	IMessageTransmitterCaller     // Read-only binding to the contract
	IMessageTransmitterTransactor // Write-only binding to the contract
	IMessageTransmitterFilterer   // Log filterer for contract events
}

// IMessageTransmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageTransmitterSession struct {
	Contract     *IMessageTransmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMessageTransmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageTransmitterCallerSession struct {
	Contract *IMessageTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMessageTransmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageTransmitterTransactorSession struct {
	Contract     *IMessageTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMessageTransmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageTransmitterRaw struct {
	Contract *IMessageTransmitter // Generic contract binding to access the raw methods on
}

// IMessageTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageTransmitterCallerRaw struct {
	Contract *IMessageTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactorRaw struct {
	Contract *IMessageTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageTransmitter creates a new instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitter(address common.Address, backend bind.ContractBackend) (*IMessageTransmitter, error) {
	contract, err := bindIMessageTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitter{IMessageTransmitterCaller: IMessageTransmitterCaller{contract: contract}, IMessageTransmitterTransactor: IMessageTransmitterTransactor{contract: contract}, IMessageTransmitterFilterer: IMessageTransmitterFilterer{contract: contract}}, nil
}

// NewIMessageTransmitterCaller creates a new read-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterCaller(address common.Address, caller bind.ContractCaller) (*IMessageTransmitterCaller, error) {
	contract, err := bindIMessageTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterCaller{contract: contract}, nil
}

// NewIMessageTransmitterTransactor creates a new write-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageTransmitterTransactor, error) {
	contract, err := bindIMessageTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterTransactor{contract: contract}, nil
}

// NewIMessageTransmitterFilterer creates a new log filterer instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageTransmitterFilterer, error) {
	contract, err := bindIMessageTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterFilterer{contract: contract}, nil
}

// bindIMessageTransmitter binds a generic wrapper to an already deployed contract.
func bindIMessageTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageTransmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.IMessageTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IMessageTransmitter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterSession) LocalDomain() (uint32, error) {
	return _IMessageTransmitter.Contract.LocalDomain(&_IMessageTransmitter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterCallerSession) LocalDomain() (uint32, error) {
	return _IMessageTransmitter.Contract.LocalDomain(&_IMessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterCaller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IMessageTransmitter.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) NextAvailableNonce() (uint64, error) {
	return _IMessageTransmitter.Contract.NextAvailableNonce(&_IMessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterCallerSession) NextAvailableNonce() (uint64, error) {
	return _IMessageTransmitter.Contract.NextAvailableNonce(&_IMessageTransmitter.CallOpts)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// ISynapseCCTPMetaData contains all meta data concerning the ISynapseCCTP contract.
var ISynapseCCTPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"}],\"name\":\"receiveCircleToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"swapParams\",\"type\":\"bytes\"}],\"name\":\"sendCircleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4a5ae51d": "receiveCircleToken(bytes,bytes,uint32,bytes)",
		"304ddb4c": "sendCircleToken(address,uint256,address,uint256,uint32,bytes)",
	},
}

// ISynapseCCTPABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseCCTPMetaData.ABI instead.
var ISynapseCCTPABI = ISynapseCCTPMetaData.ABI

// Deprecated: Use ISynapseCCTPMetaData.Sigs instead.
// ISynapseCCTPFuncSigs maps the 4-byte function signature to its string representation.
var ISynapseCCTPFuncSigs = ISynapseCCTPMetaData.Sigs

// ISynapseCCTP is an auto generated Go binding around an Ethereum contract.
type ISynapseCCTP struct {
	ISynapseCCTPCaller     // Read-only binding to the contract
	ISynapseCCTPTransactor // Write-only binding to the contract
	ISynapseCCTPFilterer   // Log filterer for contract events
}

// ISynapseCCTPCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseCCTPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseCCTPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseCCTPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseCCTPSession struct {
	Contract     *ISynapseCCTP     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynapseCCTPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseCCTPCallerSession struct {
	Contract *ISynapseCCTPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ISynapseCCTPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseCCTPTransactorSession struct {
	Contract     *ISynapseCCTPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ISynapseCCTPRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseCCTPRaw struct {
	Contract *ISynapseCCTP // Generic contract binding to access the raw methods on
}

// ISynapseCCTPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseCCTPCallerRaw struct {
	Contract *ISynapseCCTPCaller // Generic read-only contract binding to access the raw methods on
}

// ISynapseCCTPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseCCTPTransactorRaw struct {
	Contract *ISynapseCCTPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseCCTP creates a new instance of ISynapseCCTP, bound to a specific deployed contract.
func NewISynapseCCTP(address common.Address, backend bind.ContractBackend) (*ISynapseCCTP, error) {
	contract, err := bindISynapseCCTP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTP{ISynapseCCTPCaller: ISynapseCCTPCaller{contract: contract}, ISynapseCCTPTransactor: ISynapseCCTPTransactor{contract: contract}, ISynapseCCTPFilterer: ISynapseCCTPFilterer{contract: contract}}, nil
}

// NewISynapseCCTPCaller creates a new read-only instance of ISynapseCCTP, bound to a specific deployed contract.
func NewISynapseCCTPCaller(address common.Address, caller bind.ContractCaller) (*ISynapseCCTPCaller, error) {
	contract, err := bindISynapseCCTP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPCaller{contract: contract}, nil
}

// NewISynapseCCTPTransactor creates a new write-only instance of ISynapseCCTP, bound to a specific deployed contract.
func NewISynapseCCTPTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseCCTPTransactor, error) {
	contract, err := bindISynapseCCTP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPTransactor{contract: contract}, nil
}

// NewISynapseCCTPFilterer creates a new log filterer instance of ISynapseCCTP, bound to a specific deployed contract.
func NewISynapseCCTPFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseCCTPFilterer, error) {
	contract, err := bindISynapseCCTP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPFilterer{contract: contract}, nil
}

// bindISynapseCCTP binds a generic wrapper to an already deployed contract.
func bindISynapseCCTP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISynapseCCTPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseCCTP *ISynapseCCTPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseCCTP.Contract.ISynapseCCTPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseCCTP *ISynapseCCTPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.ISynapseCCTPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseCCTP *ISynapseCCTPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.ISynapseCCTPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseCCTP *ISynapseCCTPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseCCTP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseCCTP *ISynapseCCTPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseCCTP *ISynapseCCTPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.contract.Transact(opts, method, params...)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) payable returns()
func (_ISynapseCCTP *ISynapseCCTPTransactor) ReceiveCircleToken(opts *bind.TransactOpts, message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.contract.Transact(opts, "receiveCircleToken", message, signature, requestVersion, formattedRequest)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) payable returns()
func (_ISynapseCCTP *ISynapseCCTPSession) ReceiveCircleToken(message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.ReceiveCircleToken(&_ISynapseCCTP.TransactOpts, message, signature, requestVersion, formattedRequest)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) payable returns()
func (_ISynapseCCTP *ISynapseCCTPTransactorSession) ReceiveCircleToken(message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.ReceiveCircleToken(&_ISynapseCCTP.TransactOpts, message, signature, requestVersion, formattedRequest)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_ISynapseCCTP *ISynapseCCTPTransactor) SendCircleToken(opts *bind.TransactOpts, recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.contract.Transact(opts, "sendCircleToken", recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_ISynapseCCTP *ISynapseCCTPSession) SendCircleToken(recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.SendCircleToken(&_ISynapseCCTP.TransactOpts, recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_ISynapseCCTP *ISynapseCCTPTransactorSession) SendCircleToken(recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _ISynapseCCTP.Contract.SendCircleToken(&_ISynapseCCTP.TransactOpts, recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// ITokenMessengerMetaData contains all meta data concerning the ITokenMessenger contract.
var ITokenMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"depositForBurnWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMessageTransmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f856ddb6": "depositForBurnWithCaller(uint256,uint32,bytes32,address,bytes32)",
		"96abeb70": "handleReceiveMessage(uint32,bytes32,bytes)",
		"2c121921": "localMessageTransmitter()",
		"cb75c11c": "localMinter()",
	},
}

// ITokenMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMessengerMetaData.ABI instead.
var ITokenMessengerABI = ITokenMessengerMetaData.ABI

// Deprecated: Use ITokenMessengerMetaData.Sigs instead.
// ITokenMessengerFuncSigs maps the 4-byte function signature to its string representation.
var ITokenMessengerFuncSigs = ITokenMessengerMetaData.Sigs

// ITokenMessenger is an auto generated Go binding around an Ethereum contract.
type ITokenMessenger struct {
	ITokenMessengerCaller     // Read-only binding to the contract
	ITokenMessengerTransactor // Write-only binding to the contract
	ITokenMessengerFilterer   // Log filterer for contract events
}

// ITokenMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenMessengerSession struct {
	Contract     *ITokenMessenger  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenMessengerCallerSession struct {
	Contract *ITokenMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ITokenMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenMessengerTransactorSession struct {
	Contract     *ITokenMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ITokenMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenMessengerRaw struct {
	Contract *ITokenMessenger // Generic contract binding to access the raw methods on
}

// ITokenMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenMessengerCallerRaw struct {
	Contract *ITokenMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenMessengerTransactorRaw struct {
	Contract *ITokenMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenMessenger creates a new instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessenger(address common.Address, backend bind.ContractBackend) (*ITokenMessenger, error) {
	contract, err := bindITokenMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenMessenger{ITokenMessengerCaller: ITokenMessengerCaller{contract: contract}, ITokenMessengerTransactor: ITokenMessengerTransactor{contract: contract}, ITokenMessengerFilterer: ITokenMessengerFilterer{contract: contract}}, nil
}

// NewITokenMessengerCaller creates a new read-only instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerCaller(address common.Address, caller bind.ContractCaller) (*ITokenMessengerCaller, error) {
	contract, err := bindITokenMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerCaller{contract: contract}, nil
}

// NewITokenMessengerTransactor creates a new write-only instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenMessengerTransactor, error) {
	contract, err := bindITokenMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerTransactor{contract: contract}, nil
}

// NewITokenMessengerFilterer creates a new log filterer instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenMessengerFilterer, error) {
	contract, err := bindITokenMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerFilterer{contract: contract}, nil
}

// bindITokenMessenger binds a generic wrapper to an already deployed contract.
func bindITokenMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMessenger *ITokenMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMessenger.Contract.ITokenMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMessenger *ITokenMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ITokenMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMessenger *ITokenMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ITokenMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMessenger *ITokenMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMessenger *ITokenMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMessenger *ITokenMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.contract.Transact(opts, method, params...)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCaller) LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITokenMessenger.contract.Call(opts, &out, "localMessageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerSession) LocalMessageTransmitter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMessageTransmitter(&_ITokenMessenger.CallOpts)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCallerSession) LocalMessageTransmitter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMessageTransmitter(&_ITokenMessenger.CallOpts)
}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCaller) LocalMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITokenMessenger.contract.Call(opts, &out, "localMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerSession) LocalMinter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMinter(&_ITokenMessenger.CallOpts)
}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCallerSession) LocalMinter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMinter(&_ITokenMessenger.CallOpts)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerTransactor) DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "depositForBurnWithCaller", amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurnWithCaller(&_ITokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerTransactorSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurnWithCaller(&_ITokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerTransactor) HandleReceiveMessage(opts *bind.TransactOpts, remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "handleReceiveMessage", remoteDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerSession) HandleReceiveMessage(remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.HandleReceiveMessage(&_ITokenMessenger.TransactOpts, remoteDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerTransactorSession) HandleReceiveMessage(remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.HandleReceiveMessage(&_ITokenMessenger.TransactOpts, remoteDomain, sender, messageBody)
}

// ITokenMinterMetaData contains all meta data concerning the ITokenMinter contract.
var ITokenMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"remoteToken\",\"type\":\"bytes32\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"burnToken\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9dc29fac": "burn(address,uint256)",
		"78a0565e": "getLocalToken(uint32,bytes32)",
		"d54de06f": "mint(uint32,bytes32,address,uint256)",
	},
}

// ITokenMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMinterMetaData.ABI instead.
var ITokenMinterABI = ITokenMinterMetaData.ABI

// Deprecated: Use ITokenMinterMetaData.Sigs instead.
// ITokenMinterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenMinterFuncSigs = ITokenMinterMetaData.Sigs

// ITokenMinter is an auto generated Go binding around an Ethereum contract.
type ITokenMinter struct {
	ITokenMinterCaller     // Read-only binding to the contract
	ITokenMinterTransactor // Write-only binding to the contract
	ITokenMinterFilterer   // Log filterer for contract events
}

// ITokenMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenMinterSession struct {
	Contract     *ITokenMinter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenMinterCallerSession struct {
	Contract *ITokenMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenMinterTransactorSession struct {
	Contract     *ITokenMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenMinterRaw struct {
	Contract *ITokenMinter // Generic contract binding to access the raw methods on
}

// ITokenMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenMinterCallerRaw struct {
	Contract *ITokenMinterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenMinterTransactorRaw struct {
	Contract *ITokenMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenMinter creates a new instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinter(address common.Address, backend bind.ContractBackend) (*ITokenMinter, error) {
	contract, err := bindITokenMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenMinter{ITokenMinterCaller: ITokenMinterCaller{contract: contract}, ITokenMinterTransactor: ITokenMinterTransactor{contract: contract}, ITokenMinterFilterer: ITokenMinterFilterer{contract: contract}}, nil
}

// NewITokenMinterCaller creates a new read-only instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterCaller(address common.Address, caller bind.ContractCaller) (*ITokenMinterCaller, error) {
	contract, err := bindITokenMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterCaller{contract: contract}, nil
}

// NewITokenMinterTransactor creates a new write-only instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenMinterTransactor, error) {
	contract, err := bindITokenMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterTransactor{contract: contract}, nil
}

// NewITokenMinterFilterer creates a new log filterer instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenMinterFilterer, error) {
	contract, err := bindITokenMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterFilterer{contract: contract}, nil
}

// bindITokenMinter binds a generic wrapper to an already deployed contract.
func bindITokenMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMinter *ITokenMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMinter.Contract.ITokenMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMinter *ITokenMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMinter.Contract.ITokenMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMinter *ITokenMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMinter.Contract.ITokenMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMinter *ITokenMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMinter *ITokenMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMinter *ITokenMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMinter.Contract.contract.Transact(opts, method, params...)
}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterCaller) GetLocalToken(opts *bind.CallOpts, remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITokenMinter.contract.Call(opts, &out, "getLocalToken", remoteDomain, remoteToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterSession) GetLocalToken(remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	return _ITokenMinter.Contract.GetLocalToken(&_ITokenMinter.CallOpts, remoteDomain, remoteToken)
}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterCallerSession) GetLocalToken(remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	return _ITokenMinter.Contract.GetLocalToken(&_ITokenMinter.CallOpts, remoteDomain, remoteToken)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterTransactor) Burn(opts *bind.TransactOpts, burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "burn", burnToken, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterSession) Burn(burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Burn(&_ITokenMinter.TransactOpts, burnToken, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterTransactorSession) Burn(burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Burn(&_ITokenMinter.TransactOpts, burnToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterTransactor) Mint(opts *bind.TransactOpts, sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "mint", sourceDomain, burnToken, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterSession) Mint(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Mint(&_ITokenMinter.TransactOpts, sourceDomain, burnToken, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterTransactorSession) Mint(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Mint(&_ITokenMinter.TransactOpts, sourceDomain, burnToken, to, amount)
}

// MinimalForwarderLibMetaData contains all meta data concerning the MinimalForwarderLib contract.
var MinimalForwarderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206bf9d67355d8175d675ec9d1517c5d7607cbc418783ece2e0065781e85d5190a64736f6c63430008110033",
}

// MinimalForwarderLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MinimalForwarderLibMetaData.ABI instead.
var MinimalForwarderLibABI = MinimalForwarderLibMetaData.ABI

// MinimalForwarderLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MinimalForwarderLibMetaData.Bin instead.
var MinimalForwarderLibBin = MinimalForwarderLibMetaData.Bin

// DeployMinimalForwarderLib deploys a new Ethereum contract, binding an instance of MinimalForwarderLib to it.
func DeployMinimalForwarderLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MinimalForwarderLib, error) {
	parsed, err := MinimalForwarderLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MinimalForwarderLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MinimalForwarderLib{MinimalForwarderLibCaller: MinimalForwarderLibCaller{contract: contract}, MinimalForwarderLibTransactor: MinimalForwarderLibTransactor{contract: contract}, MinimalForwarderLibFilterer: MinimalForwarderLibFilterer{contract: contract}}, nil
}

// MinimalForwarderLib is an auto generated Go binding around an Ethereum contract.
type MinimalForwarderLib struct {
	MinimalForwarderLibCaller     // Read-only binding to the contract
	MinimalForwarderLibTransactor // Write-only binding to the contract
	MinimalForwarderLibFilterer   // Log filterer for contract events
}

// MinimalForwarderLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinimalForwarderLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinimalForwarderLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinimalForwarderLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinimalForwarderLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinimalForwarderLibSession struct {
	Contract     *MinimalForwarderLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MinimalForwarderLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinimalForwarderLibCallerSession struct {
	Contract *MinimalForwarderLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MinimalForwarderLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinimalForwarderLibTransactorSession struct {
	Contract     *MinimalForwarderLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MinimalForwarderLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinimalForwarderLibRaw struct {
	Contract *MinimalForwarderLib // Generic contract binding to access the raw methods on
}

// MinimalForwarderLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinimalForwarderLibCallerRaw struct {
	Contract *MinimalForwarderLibCaller // Generic read-only contract binding to access the raw methods on
}

// MinimalForwarderLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinimalForwarderLibTransactorRaw struct {
	Contract *MinimalForwarderLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinimalForwarderLib creates a new instance of MinimalForwarderLib, bound to a specific deployed contract.
func NewMinimalForwarderLib(address common.Address, backend bind.ContractBackend) (*MinimalForwarderLib, error) {
	contract, err := bindMinimalForwarderLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderLib{MinimalForwarderLibCaller: MinimalForwarderLibCaller{contract: contract}, MinimalForwarderLibTransactor: MinimalForwarderLibTransactor{contract: contract}, MinimalForwarderLibFilterer: MinimalForwarderLibFilterer{contract: contract}}, nil
}

// NewMinimalForwarderLibCaller creates a new read-only instance of MinimalForwarderLib, bound to a specific deployed contract.
func NewMinimalForwarderLibCaller(address common.Address, caller bind.ContractCaller) (*MinimalForwarderLibCaller, error) {
	contract, err := bindMinimalForwarderLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderLibCaller{contract: contract}, nil
}

// NewMinimalForwarderLibTransactor creates a new write-only instance of MinimalForwarderLib, bound to a specific deployed contract.
func NewMinimalForwarderLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MinimalForwarderLibTransactor, error) {
	contract, err := bindMinimalForwarderLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderLibTransactor{contract: contract}, nil
}

// NewMinimalForwarderLibFilterer creates a new log filterer instance of MinimalForwarderLib, bound to a specific deployed contract.
func NewMinimalForwarderLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MinimalForwarderLibFilterer, error) {
	contract, err := bindMinimalForwarderLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinimalForwarderLibFilterer{contract: contract}, nil
}

// bindMinimalForwarderLib binds a generic wrapper to an already deployed contract.
func bindMinimalForwarderLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinimalForwarderLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinimalForwarderLib *MinimalForwarderLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinimalForwarderLib.Contract.MinimalForwarderLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinimalForwarderLib *MinimalForwarderLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinimalForwarderLib.Contract.MinimalForwarderLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinimalForwarderLib *MinimalForwarderLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinimalForwarderLib.Contract.MinimalForwarderLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinimalForwarderLib *MinimalForwarderLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinimalForwarderLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinimalForwarderLib *MinimalForwarderLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinimalForwarderLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinimalForwarderLib *MinimalForwarderLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinimalForwarderLib.Contract.contract.Transact(opts, method, params...)
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
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204b8d4b235bdaedbfb3ecf4d3b800f5a0edaa2bdf26c32b60316a4b0dbf4334aa64736f6c63430008110033",
}

// RequestLibABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestLibMetaData.ABI instead.
var RequestLibABI = RequestLibMetaData.ABI

// RequestLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequestLibMetaData.Bin instead.
var RequestLibBin = RequestLibMetaData.Bin

// DeployRequestLib deploys a new Ethereum contract, binding an instance of RequestLib to it.
func DeployRequestLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestLib, error) {
	parsed, err := RequestLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequestLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestLib{RequestLibCaller: RequestLibCaller{contract: contract}, RequestLibTransactor: RequestLibTransactor{contract: contract}, RequestLibFilterer: RequestLibFilterer{contract: contract}}, nil
}

// RequestLib is an auto generated Go binding around an Ethereum contract.
type RequestLib struct {
	RequestLibCaller     // Read-only binding to the contract
	RequestLibTransactor // Write-only binding to the contract
	RequestLibFilterer   // Log filterer for contract events
}

// RequestLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestLibSession struct {
	Contract     *RequestLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestLibCallerSession struct {
	Contract *RequestLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RequestLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestLibTransactorSession struct {
	Contract     *RequestLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RequestLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestLibRaw struct {
	Contract *RequestLib // Generic contract binding to access the raw methods on
}

// RequestLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestLibCallerRaw struct {
	Contract *RequestLibCaller // Generic read-only contract binding to access the raw methods on
}

// RequestLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestLibTransactorRaw struct {
	Contract *RequestLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestLib creates a new instance of RequestLib, bound to a specific deployed contract.
func NewRequestLib(address common.Address, backend bind.ContractBackend) (*RequestLib, error) {
	contract, err := bindRequestLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestLib{RequestLibCaller: RequestLibCaller{contract: contract}, RequestLibTransactor: RequestLibTransactor{contract: contract}, RequestLibFilterer: RequestLibFilterer{contract: contract}}, nil
}

// NewRequestLibCaller creates a new read-only instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibCaller(address common.Address, caller bind.ContractCaller) (*RequestLibCaller, error) {
	contract, err := bindRequestLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestLibCaller{contract: contract}, nil
}

// NewRequestLibTransactor creates a new write-only instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestLibTransactor, error) {
	contract, err := bindRequestLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestLibTransactor{contract: contract}, nil
}

// NewRequestLibFilterer creates a new log filterer instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestLibFilterer, error) {
	contract, err := bindRequestLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestLibFilterer{contract: contract}, nil
}

// bindRequestLib binds a generic wrapper to an already deployed contract.
func bindRequestLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestLib *RequestLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestLib.Contract.RequestLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestLib *RequestLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestLib.Contract.RequestLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestLib *RequestLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestLib.Contract.RequestLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestLib *RequestLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestLib *RequestLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestLib *RequestLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestLib.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ac00af76b85ffcb4463bd221b7eb9501e240801b098bc859bddc3713a7dc6d3664736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// SynapseCCTPMetaData contains all meta data concerning the SynapseCCTP contract.
var SynapseCCTPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"tokenMessenger_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CCTPGasRescueFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectGasAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectProtocolFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPInsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPMessageNotReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPSymbolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPSymbolIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPTokenAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPTokenNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CastOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwarderDeploymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectRequestLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoteCCTPDeploymentNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownRequestVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAirdropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayerFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accumulatedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSwap\",\"type\":\"bool\"}],\"name\":\"calculateFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"circleTokenPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeStructures\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"relayerFee\",\"type\":\"uint40\"},{\"internalType\":\"uint72\",\"name\":\"minBaseFee\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"minSwapFee\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"maxFee\",\"type\":\"uint72\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"bridgeTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"isRequestFulfilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageTransmitter\",\"outputs\":[{\"internalType\":\"contractIMessageTransmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"}],\"name\":\"receiveCircleToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayerFeeCollectors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remoteDomainConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"synapseCCTP\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescueGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"swapParams\",\"type\":\"bytes\"}],\"name\":\"sendCircleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"circleToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setCircleTokenPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"remoteSynapseCCTP\",\"type\":\"address\"}],\"name\":\"setRemoteDomainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"setTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawRelayerFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d4a67c6d": "accumulatedFees(address,address)",
		"4a85178d": "addToken(string,address,uint256,uint256,uint256,uint256)",
		"0d25aafe": "calculateFeeAmount(address,uint256,bool)",
		"e00a83e0": "chainGasAmount()",
		"a4b1d034": "circleTokenPool(address)",
		"dc72495b": "feeStructures(address)",
		"9c1d060e": "getBridgeTokens()",
		"f879a41a": "getLocalToken(uint32,address)",
		"92a442ea": "isRequestFulfilled(bytes32)",
		"8d3638f4": "localDomain()",
		"7b04c181": "messageTransmitter()",
		"8da5cb5b": "owner()",
		"b0e21e8a": "protocolFee()",
		"4a5ae51d": "receiveCircleToken(bytes,bytes,uint32,bytes)",
		"41f355ee": "relayerFeeCollectors(address)",
		"e9259ab9": "remoteDomainConfig(uint256)",
		"5fa7b584": "removeToken(address)",
		"715018a6": "renounceOwnership()",
		"40432d51": "rescueGas()",
		"304ddb4c": "sendCircleToken(address,uint256,address,uint256,uint32,bytes)",
		"b250fe6b": "setChainGasAmount(uint256)",
		"2cc9e7e5": "setCircleTokenPool(address,address)",
		"a42dce80": "setFeeCollector(address)",
		"787dce3d": "setProtocolFee(uint256)",
		"e9bbb36d": "setRemoteDomainConfig(uint256,uint32,address)",
		"4bdb4eed": "setTokenFee(address,uint256,uint256,uint256,uint256)",
		"a5bc29c2": "symbolToToken(string)",
		"46117830": "tokenMessenger()",
		"0ba36121": "tokenToSymbol(address)",
		"f2fde38b": "transferOwnership(address)",
		"2d80caa5": "withdrawProtocolFees(address)",
		"f7265b3a": "withdrawRelayerFees(address)",
	},
	Bin: "0x60e06040523480156200001157600080fd5b506040516200458638038062004586833981016040819052620000349162000199565b6200003f3362000130565b6001600160a01b03811660c081905260408051632c12192160e01b81529051632c121921916004808201926020929091908290030181865afa1580156200008a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b0919062000199565b6001600160a01b031660a08190526040805163234d8e3d60e21b81529051638d3638f4916004808201926020929091908290030181865afa158015620000fa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001209190620001c0565b63ffffffff1660805250620001e8565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146200019657600080fd5b50565b600060208284031215620001ac57600080fd5b8151620001b98162000180565b9392505050565b600060208284031215620001d357600080fd5b815163ffffffff81168114620001b957600080fd5b60805160a05160c051614336620002506000396000818161032301528181610c8001528181610caf015261232e0152600081816103ff01528181610ad101526122ae01526000818161043301528181610b5c01528181610f2401526118c001526143366000f3fe6080604052600436106101e35760003560e01c80638da5cb5b11610102578063d4a67c6d11610095578063e9bbb36d11610064578063e9bbb36d14610716578063f2fde38b14610736578063f7265b3a14610756578063f879a41a1461077657600080fd5b8063d4a67c6d146105a7578063dc72495b146105df578063e00a83e014610696578063e9259ab9146106ac57600080fd5b8063a4b1d034116100d1578063a4b1d034146104fa578063a5bc29c214610530578063b0e21e8a14610571578063b250fe6b1461058757600080fd5b80638da5cb5b1461046a57806392a442ea146104885780639c1d060e146104b8578063a42dce80146104da57600080fd5b80634a5ae51d1161017a578063715018a611610149578063715018a6146103b8578063787dce3d146103cd5780637b04c181146103ed5780638d3638f41461042157600080fd5b80634a5ae51d146103455780634a85178d146103585780634bdb4eed146103785780635fa7b5841461039857600080fd5b8063304ddb4c116101b6578063304ddb4c1461028e57806340432d51146102ae57806341f355ee146102c3578063461178301461031157600080fd5b80630ba36121146101e85780630d25aafe1461021e5780632cc9e7e51461024c5780632d80caa51461026e575b600080fd5b3480156101f457600080fd5b50610208610203366004613785565b610796565b60405161021591906137f2565b60405180910390f35b34801561022a57600080fd5b5061023e610239366004613813565b610830565b604051908152602001610215565b34801561025857600080fd5b5061026c610267366004613855565b610847565b005b34801561027a57600080fd5b5061026c610289366004613785565b61096d565b34801561029a57600080fd5b5061026c6102a936600461397e565b610a80565b3480156102ba57600080fd5b5061026c610dd6565b3480156102cf57600080fd5b506102f96102de366004613785565b6005602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610215565b34801561031d57600080fd5b506102f97f000000000000000000000000000000000000000000000000000000000000000081565b61026c610353366004613a47565b610eb5565b34801561036457600080fd5b5061026c610373366004613ae1565b61102f565b34801561038457600080fd5b5061026c610393366004613b53565b6111a3565b3480156103a457600080fd5b5061026c6103b3366004613785565b611252565b3480156103c457600080fd5b5061026c61141a565b3480156103d957600080fd5b5061026c6103e8366004613b97565b611480565b3480156103f957600080fd5b506102f97f000000000000000000000000000000000000000000000000000000000000000081565b34801561042d57600080fd5b506104557f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610215565b34801561047657600080fd5b506000546001600160a01b03166102f9565b34801561049457600080fd5b506104a86104a3366004613b97565b61155f565b6040519015158152602001610215565b3480156104c457600080fd5b506104cd61157d565b6040516102159190613bb0565b3480156104e657600080fd5b5061026c6104f5366004613785565b61170a565b34801561050657600080fd5b506102f9610515366004613785565b600b602052600090815260409020546001600160a01b031681565b34801561053c57600080fd5b506102f961054b366004613c4d565b80516020818301810180516002825292820191909301209152546001600160a01b031681565b34801561057d57600080fd5b5061023e60065481565b34801561059357600080fd5b5061026c6105a2366004613b97565b611792565b3480156105b357600080fd5b5061023e6105c2366004613855565b600460209081526000928352604080842090915290825290205481565b3480156105eb57600080fd5b5061065d6105fa366004613785565b60036020526000908152604090205464ffffffffff81169068ffffffffffffffffff6501000000000082048116916e0100000000000000000000000000008104821691770100000000000000000000000000000000000000000000009091041684565b6040805164ffffffffff909516855268ffffffffffffffffff938416602086015291831691840191909152166060820152608001610215565b3480156106a257600080fd5b5061023e60075481565b3480156106b857600080fd5b506106f26106c7366004613b97565b600a6020526000908152604090205463ffffffff81169064010000000090046001600160a01b031682565b6040805163ffffffff90931683526001600160a01b03909116602083015201610215565b34801561072257600080fd5b5061026c610731366004613c82565b611821565b34801561074257600080fd5b5061026c610751366004613785565b611a17565b34801561076257600080fd5b5061026c610771366004613785565b611af6565b34801561078257600080fd5b506102f9610791366004613cb9565b611b85565b600160205260009081526040902080546107af90613cd7565b80601f01602080910402602001604051908101604052809291908181526020018280546107db90613cd7565b80156108285780601f106107fd57610100808354040283529160200191610828565b820191906000526020600020905b81548152906001019060200180831161080b57829003601f168201915b505050505081565b600061083d848484611b9a565b90505b9392505050565b6000546001600160a01b031633146108a65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b0382166108e6576040517f24305eca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108f1600883611ca4565b610927576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b039182166000908152600b6020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001691909216179055565b6000546001600160a01b031633146109c75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b6001600160a01b03811660009081527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec602052604081205490819003610a39576040517f30b93f1d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821660008181527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec6020526040812055610a7c903383611cc6565b5050565b610a8b600885611ca4565b610ac1576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610acb8484611d74565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638371744e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b519190613d41565b6040805163ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015267ffffffffffffffff8316818301526001600160a01b038089166060830152608082018890528a1660a0808301919091528251808303909101815260c0909101909152909150600090610bd690859085611ea6565b6000888152600a6020908152604080832081518083019092525463ffffffff8116825264010000000090046001600160a01b031691810182905292935090819003610c4d576040517fa86a3b0e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151835160208086019190912067ffffffff0000000083831b1663ffffffff8a16176000908152915260409020610ca58a7f00000000000000000000000000000000000000000000000000000000000000008b611fd3565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663f856ddb68a84868e610ce282886120a1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b168152600481019590955263ffffffff93909316602485015260448401919091526001600160a01b03166064830152608482015260a4016020604051808303816000875af1158015610d61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d859190613d41565b50807f4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a8c888d8d8d8b604051610dc096959493929190613d5c565b60405180910390a2505050505050505050505050565b6000546001600160a01b03163314610e305760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b604051600090339047908381818185875af1925050503d8060008114610e72576040519150601f19603f3d011682016040523d82523d6000602084013e610e77565b606091505b5050905080610eb2576040517f4e5610fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b6007543414610ef0576040517fc561806500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080610efd84846120c5565b91509150600080600080610f10866121a5565b8b516020808e019190912063ffffffff8f167f0000000000000000000000000000000000000000000000000000000000000000831b67ffffffff000000001617600090815291526040812095995092975090955093509190509050610f788d8d8d8d8561220e565b6000610f848686612329565b90506000610f9c828663ffffffff8e166001146124a5565b9095509050600080610fb08685898d6126f5565b90925090503415610fc457610fc48661280e565b604080516001600160a01b038681168252602082018690528481168284015260608201849052915187928916917feaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f919081900360800190a35050505050505050505050505050505050565b6000546001600160a01b031633146110895760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b6001600160a01b0385166110c9576040517f76998feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110d46008866128a8565b61110a576040517f1191732500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611113866128bd565b6001600160a01b03851660009081526001602052604090206111358782613df8565b50846002876040516111479190613eb8565b90815260405190819003602001902080546001600160a01b03929092167fffffffffffffffffffffffff000000000000000000000000000000000000000090921691909117905561119b8585858585612a5d565b505050505050565b6000546001600160a01b031633146111fd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b611208600886611ca4565b61123e576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61124b8585858585612a5d565b5050505050565b6000546001600160a01b031633146112ac5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b6112b7600882612c53565b6112ed576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166000908152600160205260408120805461131090613cd7565b80601f016020809104026020016040519081016040528092919081815260200182805461133c90613cd7565b80156113895780601f1061135e57610100808354040283529160200191611389565b820191906000526020600020905b81548152906001019060200180831161136c57829003601f168201915b505050506001600160a01b03841660009081526001602052604081209293506113b3929150613726565b6002816040516113c39190613eb8565b908152604080516020928190038301902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556001600160a01b0393909316600090815260039091529182209190915550565b6000546001600160a01b031633146114745760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b61147e6000612c68565b565b6000546001600160a01b031633146114da5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b6114ea60026402540be400613f03565b811115611523576040517f28562c4700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60068190556040518181527fd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40906020015b60405180910390a150565b60008061156c3084612cd0565b6001600160a01b03163b1192915050565b6060600061158b6008612dbc565b90508067ffffffffffffffff8111156115a6576115a66138a0565b6040519080825280602002602001820160405280156115ec57816020015b6040805180820190915260608152600060208201528152602001906001900390816115c45790505b50915060005b81811015611705576000611607600883612dc6565b9050604051806040016040528060016000846001600160a01b03166001600160a01b03168152602001908152602001600020805461164490613cd7565b80601f016020809104026020016040519081016040528092919081815260200182805461167090613cd7565b80156116bd5780601f10611692576101008083540402835291602001916116bd565b820191906000526020600020905b8154815290600101906020018083116116a057829003601f168201915b50505050508152602001826001600160a01b03168152508483815181106116e6576116e6613f3e565b60200260200101819052505080806116fd90613f6d565b9150506115f2565b505090565b3360008181526005602090815260409182902080547fffffffffffffffffffffffff000000000000000000000000000000000000000081166001600160a01b03878116918217909355845192909116808352928201529092917f9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e910160405180910390a25050565b6000546001600160a01b031633146117ec5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b60078190556040518181527f5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b490602001611554565b6000546001600160a01b0316331461187b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b82158061188757504683145b156118be576040517f3f8f40a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168263ffffffff1603611923576040517f93c970c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8216156001841414611966576040517f93c970c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166119a6576040517f24305eca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051808201825263ffffffff93841681526001600160a01b0392831660208083019182526000968752600a905291909420935184549151909216640100000000027fffffffffffffffff0000000000000000000000000000000000000000000000009091169190921617179055565b6000546001600160a01b03163314611a715760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161089d565b6001600160a01b038116611aed5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161089d565b610eb281612c68565b3360009081526004602090815260408083206001600160a01b038516845290915281205490819003611b54576040517f30b93f1d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360008181526004602090815260408083206001600160a01b0387168085529252822091909155610a7c9183611cc6565b6000611b918383612329565b90505b92915050565b6001600160a01b03831660009081526003602090815260408083208151608081018352905464ffffffffff811680835268ffffffffffffffffff6501000000000083048116958401959095526e010000000000000000000000000000820485169383019390935277010000000000000000000000000000000000000000000000900490921660608301526402540be40090611c359086613f87565b611c3f9190613f03565b9150600083611c52578160200151611c58565b81604001515b68ffffffffffffffffff16905080831015611c71578092505b816060015168ffffffffffffffffff16831115611c9b57816060015168ffffffffffffffffff1692505b50509392505050565b6001600160a01b03811660009081526001830160205260408120541515611b91565b6040516001600160a01b038316602482015260448101829052611d6f9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152612dd2565b505050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015611dd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dfa9190613f9e565b9050611e116001600160a01b038516333086612eb7565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281906001600160a01b038616906370a0823190602401602060405180830381865afa158015611e70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e949190613f9e565b611e9e9190613fb7565b949350505050565b606060a0835114611ee3576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8416611f2e57815115611f27576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5081610840565b60001963ffffffff851601611fa1576080825114611f78576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8282604051602001611f8b929190613fca565b6040516020818303038152906040529050610840565b6040517f523fa8d500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0383811660248301526000919085169063dd62ed3e90604401602060405180830381865afa15801561203c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120609190613f9e565b90508181101561209b578015612085576120856001600160a01b038516846000612f08565b61209b6001600160a01b03851684600019612f08565b50505050565b6000611b916120b96001600160a01b03851684612cd0565b6001600160a01b031690565b60608063ffffffff84166121265760a083511461210e576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050604080516020810190915260008152819061219e565b60001963ffffffff851601611fa157608061214260a082613fef565b61214c9190613fef565b835114612185576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b828060200190518101906121999190614047565b915091505b9250929050565b600080600080600060a08651146121e8576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858060200190518101906121fc91906140ab565b939a9299509097509550909350915050565b600061221982613056565b905060006357ecfd2860e01b8787878760405160240161223c949392919061413d565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152905060006122d36001600160a01b0384167f000000000000000000000000000000000000000000000000000000000000000084613102565b9050808060200190518101906122e99190614164565b61231f576040517f182f34eb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cb75c11c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561238a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123ae9190614181565b9050806001600160a01b03166378a0565e856123d9866001600160a01b03166001600160a01b031690565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815263ffffffff9290921660048301526024820152604401602060405180830381865afa158015612438573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061245c9190614181565b91506001600160a01b03821661249e576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5092915050565b6000806124b3600886611ca4565b6124e9576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124f4858585611b9a565b905083811061252f576040517f3eae42e400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526005602052604090205481850392506001600160a01b0316806125e4576001600160a01b03861660009081527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec602052604081208054849290612599908490613fef565b909155505060408051600080825260208201529081018390527f108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca5619060600160405180910390a16126ec565b60006402540be400600654846125fa9190613f87565b6126049190613f03565b905060006126128285613fb7565b6001600160a01b03891660009081527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec602052604081208054929350849290919061265e908490613fef565b90915550506001600160a01b038084166000908152600460209081526040808320938c168352929052908120805483929061269a908490613fef565b9091555050604080516001600160a01b0385168152602081018390529081018390527f108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca5619060600160405180910390a150505b50935093915050565b600080825160000361271f576127156001600160a01b0386168786611cc6565b5083905082612805565b6001600160a01b038086166000908152600b6020526040902054168061275f576127536001600160a01b0387168887611cc6565b85859250925050612805565b60008060008061276e88613111565b93509350935093506127808584613175565b96506001600160a01b0387166127b4576127a46001600160a01b038b168c8b611cc6565b8989965096505050505050612805565b6127bf8a868b611fd3565b6127cd8585858c8686613278565b9550856000036127eb576127a46001600160a01b038b168c8b611cc6565b6127ff6001600160a01b0388168c88611cc6565b50505050505b94509492505050565b6000816001600160a01b03163460405160006040518083038185875af1925050503d806000811461285b576040519150601f19603f3d011682016040523d82523d6000602084013e612860565b606091505b505090507ff9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd68381612891576000612893565b345b60405190815260200160405180910390a15050565b6000611b91836001600160a01b03841661332f565b60006001600160a01b03166002826040516128d89190613eb8565b908152604051908190036020019020546001600160a01b031614612928576040517f82ca3adf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80518190600510612965576040517f3f8fe5a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b6005811015611d6f576040518060400160405280600581526020017f434354502e00000000000000000000000000000000000000000000000000000081525081815181106129b8576129b8613f3e565b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168282815181106129f7576129f7613f3e565b01602001517fff000000000000000000000000000000000000000000000000000000000000001614612a55576040517f3f8fe5a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600101612968565b62989680841115612a9a576040517f76998feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81831115612ad4576040517f76998feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80821115612b0e576040517f76998feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060800160405280612b228661337e565b64ffffffffff168152602001612b37856133c3565b68ffffffffffffffffff168152602001612b50846133c3565b68ffffffffffffffffff168152602001612b69836133c3565b68ffffffffffffffffff9081169091526001600160a01b039096166000908152600360209081526040918290208351815492850151938501516060909501518a16770100000000000000000000000000000000000000000000000276ffffffffffffffffffffffffffffffffffffffffffffff958b166e01000000000000000000000000000002959095166dffffffffffffffffffffffffffff94909a1665010000000000027fffffffffffffffffffffffffffffffffffff000000000000000000000000000090931664ffffffffff909116179190911791909116969096171790945550505050565b6000611b91836001600160a01b038416613408565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000611b9183836040518060400160405280602081526020017f602036038060203d373d3d3d923d343d355af13d82803e903d91601e57fd5bf3815250604051602001612d1d919061419e565b60405160208183030381529060405280519060200120604051602001612da3939291907fff00000000000000000000000000000000000000000000000000000000000000815260609390931b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660018401526015830191909152603582015260550190565b6040516020818303038152906040528051906020012090565b6000611b94825490565b6000611b9183836134fb565b6000612e27826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166135259092919063ffffffff16565b805190915015611d6f5780806020019051810190612e459190614164565b611d6f5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161089d565b6040516001600160a01b038085166024830152831660448201526064810182905261209b9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611d0b565b801580612f9b57506040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03838116602483015284169063dd62ed3e90604401602060405180830381865afa158015612f75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f999190613f9e565b155b61300d5760405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000606482015260840161089d565b6040516001600160a01b038316602482015260448101829052611d6f9084907f095ea7b30000000000000000000000000000000000000000000000000000000090606401611d0b565b6000806040518060400160405280602081526020017f602036038060203d373d3d3d923d343d355af13d82803e903d91601e57fd5bf381525060405160200161309f919061419e565b6040516020818303038152906040529050828151602083016000f591506001600160a01b0382166130fc576040517f27afa9fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50919050565b606061083d8484846000613534565b6000806000806080855114613152576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84806020019051810190613166919061421b565b93509350935093509193509193565b6040805160ff831660248083019190915282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f82b86600000000000000000000000000000000000000000000000000000000001790529051600091829182916001600160a01b038716916131fd9190613eb8565b600060405180830381855afa9150503d8060008114613238576040519150601f19603f3d011682016040523d82523d6000602084013e61323d565b606091505b5091509150818015613250575080516020145b1561326b576132646132618261425e565b90565b9250613270565b600092505b505092915050565b6040517f9169558600000000000000000000000000000000000000000000000000000000815260ff8087166004830152851660248201526044810184905260648101829052608481018390526000906001600160a01b0388169063916955869060a4016020604051808303816000875af1925050508015613316575060408051601f3d908101601f1916820190925261331391810190613f9e565b60015b61332257506000613325565b90505b9695505050505050565b600081815260018301602052604081205461337657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611b94565b506000611b94565b600064ffffffffff8211156133bf576040517fe58d471800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b600068ffffffffffffffffff8211156133bf576040517fe58d471800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260018301602052604081205480156134f157600061342c600183613fb7565b855490915060009061344090600190613fb7565b90508181146134a557600086600001828154811061346057613460613f3e565b906000526020600020015490508087600001848154811061348357613483613f3e565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806134b6576134b6614282565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050611b94565b6000915050611b94565b600082600001828154811061351257613512613f3e565b9060005260206000200154905092915050565b606061083d848460008561357f565b60606135766001600160a01b038516846040516020016135559291906142b1565b60408051601f198184030181529190526001600160a01b03871690846136c7565b95945050505050565b6060824710156135f75760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161089d565b6001600160a01b0385163b61364e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161089d565b600080866001600160a01b0316858760405161366a9190613eb8565b60006040518083038185875af1925050503d80600081146136a7576040519150601f19603f3d011682016040523d82523d6000602084013e6136ac565b606091505b50915091506136bc8282866136ed565b979650505050505050565b606061083d8484846040518060600160405280602981526020016142d86029913961357f565b606083156136fc575081610840565b82511561370c5782518084602001fd5b8160405162461bcd60e51b815260040161089d91906137f2565b50805461373290613cd7565b6000825580601f10613742575050565b601f016020900490600052602060002090810190610eb291905b808211156133bf576000815560010161375c565b6001600160a01b0381168114610eb257600080fd5b60006020828403121561379757600080fd5b813561084081613770565b60005b838110156137bd5781810151838201526020016137a5565b50506000910152565b600081518084526137de8160208601602086016137a2565b601f01601f19169290920160200192915050565b602081526000611b9160208301846137c6565b8015158114610eb257600080fd5b60008060006060848603121561382857600080fd5b833561383381613770565b925060208401359150604084013561384a81613805565b809150509250925092565b6000806040838503121561386857600080fd5b823561387381613770565b9150602083013561388381613770565b809150509250929050565b63ffffffff81168114610eb257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156138f8576138f86138a0565b604052919050565b600067ffffffffffffffff82111561391a5761391a6138a0565b50601f01601f191660200190565b600082601f83011261393957600080fd5b813561394c61394782613900565b6138cf565b81815284602083860101111561396157600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561399757600080fd5b86356139a281613770565b95506020870135945060408701356139b981613770565b93506060870135925060808701356139d08161388e565b915060a087013567ffffffffffffffff8111156139ec57600080fd5b6139f889828a01613928565b9150509295509295509295565b60008083601f840112613a1757600080fd5b50813567ffffffffffffffff811115613a2f57600080fd5b60208301915083602082850101111561219e57600080fd5b60008060008060008060808789031215613a6057600080fd5b863567ffffffffffffffff80821115613a7857600080fd5b613a848a838b01613a05565b90985096506020890135915080821115613a9d57600080fd5b613aa98a838b01613a05565b909650945060408901359150613abe8261388e565b90925060608801359080821115613ad457600080fd5b506139f889828a01613928565b60008060008060008060c08789031215613afa57600080fd5b863567ffffffffffffffff811115613b1157600080fd5b613b1d89828a01613928565b9650506020870135613b2e81613770565b95989597505050506040840135936060810135936080820135935060a0909101359150565b600080600080600060a08688031215613b6b57600080fd5b8535613b7681613770565b97602087013597506040870135966060810135965060800135945092505050565b600060208284031215613ba957600080fd5b5035919050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015613c3f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc089840301855281518051878552613c19888601826137c6565b918901516001600160a01b03169489019490945294870194925090860190600101613bd7565b509098975050505050505050565b600060208284031215613c5f57600080fd5b813567ffffffffffffffff811115613c7657600080fd5b611e9e84828501613928565b600080600060608486031215613c9757600080fd5b833592506020840135613ca98161388e565b9150604084013561384a81613770565b60008060408385031215613ccc57600080fd5b82356138738161388e565b600181811c90821680613ceb57607f821691505b6020821081036130fc577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b805167ffffffffffffffff81168114613d3c57600080fd5b919050565b600060208284031215613d5357600080fd5b611b9182613d24565b86815267ffffffffffffffff861660208201526001600160a01b038516604082015283606082015263ffffffff8316608082015260c060a08201526000613da660c08301846137c6565b98975050505050505050565b601f821115611d6f57600081815260208120601f850160051c81016020861015613dd95750805b601f850160051c820191505b8181101561119b57828155600101613de5565b815167ffffffffffffffff811115613e1257613e126138a0565b613e2681613e208454613cd7565b84613db2565b602080601f831160018114613e5b5760008415613e435750858301515b600019600386901b1c1916600185901b17855561119b565b600085815260208120601f198616915b82811015613e8a57888601518255948401946001909101908401613e6b565b5085821015613ea85787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008251613eca8184602087016137a2565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082613f39577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006000198203613f8057613f80613ed4565b5060010190565b8082028115828204841417611b9457611b94613ed4565b600060208284031215613fb057600080fd5b5051919050565b81810381811115611b9457611b94613ed4565b604081526000613fdd60408301856137c6565b828103602084015261357681856137c6565b80820180821115611b9457611b94613ed4565b600082601f83011261401357600080fd5b815161402161394782613900565b81815284602083860101111561403657600080fd5b611e9e8260208301602087016137a2565b6000806040838503121561405a57600080fd5b825167ffffffffffffffff8082111561407257600080fd5b61407e86838701614002565b9350602085015191508082111561409457600080fd5b506140a185828601614002565b9150509250929050565b600080600080600060a086880312156140c357600080fd5b85516140ce8161388e565b94506140dc60208701613d24565b935060408601516140ec81613770565b60608701516080880151919450925061410481613770565b809150509295509295909350565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b604081526000614151604083018688614112565b82810360208401526136bc818587614112565b60006020828403121561417657600080fd5b815161084081613805565b60006020828403121561419357600080fd5b815161084081613770565b7f7f000000000000000000000000000000000000000000000000000000000000008152600082516141d68160018501602087016137a2565b7f3d5260203df300000000000000000000000000000000000000000000000000006001939091019283015250600701919050565b805160ff81168114613d3c57600080fd5b6000806000806080858703121561423157600080fd5b61423a8561420a565b93506142486020860161420a565b6040860151606090960151949790965092505050565b805160208083015191908110156130fc5760001960209190910360031b1b16919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b828152600082516142c98160208501602087016137a2565b91909101602001939250505056fe416464726573733a206c6f772d6c6576656c2063616c6c20776974682076616c7565206661696c6564a26469706673582212205affd02a1e59facad6cac5fd476d11455f305706f06f4f81215cc01d07e1f76f64736f6c63430008110033",
}

// SynapseCCTPABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseCCTPMetaData.ABI instead.
var SynapseCCTPABI = SynapseCCTPMetaData.ABI

// Deprecated: Use SynapseCCTPMetaData.Sigs instead.
// SynapseCCTPFuncSigs maps the 4-byte function signature to its string representation.
var SynapseCCTPFuncSigs = SynapseCCTPMetaData.Sigs

// SynapseCCTPBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseCCTPMetaData.Bin instead.
var SynapseCCTPBin = SynapseCCTPMetaData.Bin

// DeploySynapseCCTP deploys a new Ethereum contract, binding an instance of SynapseCCTP to it.
func DeploySynapseCCTP(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger_ common.Address) (common.Address, *types.Transaction, *SynapseCCTP, error) {
	parsed, err := SynapseCCTPMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseCCTPBin), backend, tokenMessenger_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseCCTP{SynapseCCTPCaller: SynapseCCTPCaller{contract: contract}, SynapseCCTPTransactor: SynapseCCTPTransactor{contract: contract}, SynapseCCTPFilterer: SynapseCCTPFilterer{contract: contract}}, nil
}

// SynapseCCTP is an auto generated Go binding around an Ethereum contract.
type SynapseCCTP struct {
	SynapseCCTPCaller     // Read-only binding to the contract
	SynapseCCTPTransactor // Write-only binding to the contract
	SynapseCCTPFilterer   // Log filterer for contract events
}

// SynapseCCTPCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseCCTPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseCCTPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseCCTPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseCCTPSession struct {
	Contract     *SynapseCCTP      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseCCTPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseCCTPCallerSession struct {
	Contract *SynapseCCTPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SynapseCCTPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseCCTPTransactorSession struct {
	Contract     *SynapseCCTPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SynapseCCTPRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseCCTPRaw struct {
	Contract *SynapseCCTP // Generic contract binding to access the raw methods on
}

// SynapseCCTPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseCCTPCallerRaw struct {
	Contract *SynapseCCTPCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseCCTPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseCCTPTransactorRaw struct {
	Contract *SynapseCCTPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseCCTP creates a new instance of SynapseCCTP, bound to a specific deployed contract.
func NewSynapseCCTP(address common.Address, backend bind.ContractBackend) (*SynapseCCTP, error) {
	contract, err := bindSynapseCCTP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTP{SynapseCCTPCaller: SynapseCCTPCaller{contract: contract}, SynapseCCTPTransactor: SynapseCCTPTransactor{contract: contract}, SynapseCCTPFilterer: SynapseCCTPFilterer{contract: contract}}, nil
}

// NewSynapseCCTPCaller creates a new read-only instance of SynapseCCTP, bound to a specific deployed contract.
func NewSynapseCCTPCaller(address common.Address, caller bind.ContractCaller) (*SynapseCCTPCaller, error) {
	contract, err := bindSynapseCCTP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPCaller{contract: contract}, nil
}

// NewSynapseCCTPTransactor creates a new write-only instance of SynapseCCTP, bound to a specific deployed contract.
func NewSynapseCCTPTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseCCTPTransactor, error) {
	contract, err := bindSynapseCCTP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPTransactor{contract: contract}, nil
}

// NewSynapseCCTPFilterer creates a new log filterer instance of SynapseCCTP, bound to a specific deployed contract.
func NewSynapseCCTPFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseCCTPFilterer, error) {
	contract, err := bindSynapseCCTP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFilterer{contract: contract}, nil
}

// bindSynapseCCTP binds a generic wrapper to an already deployed contract.
func bindSynapseCCTP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseCCTPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTP *SynapseCCTPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTP.Contract.SynapseCCTPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTP *SynapseCCTPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SynapseCCTPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTP *SynapseCCTPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SynapseCCTPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTP *SynapseCCTPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTP *SynapseCCTPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTP *SynapseCCTPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0xd4a67c6d.
//
// Solidity: function accumulatedFees(address , address ) view returns(uint256)
func (_SynapseCCTP *SynapseCCTPCaller) AccumulatedFees(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "accumulatedFees", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFees is a free data retrieval call binding the contract method 0xd4a67c6d.
//
// Solidity: function accumulatedFees(address , address ) view returns(uint256)
func (_SynapseCCTP *SynapseCCTPSession) AccumulatedFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SynapseCCTP.Contract.AccumulatedFees(&_SynapseCCTP.CallOpts, arg0, arg1)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0xd4a67c6d.
//
// Solidity: function accumulatedFees(address , address ) view returns(uint256)
func (_SynapseCCTP *SynapseCCTPCallerSession) AccumulatedFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SynapseCCTP.Contract.AccumulatedFees(&_SynapseCCTP.CallOpts, arg0, arg1)
}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_SynapseCCTP *SynapseCCTPCaller) CalculateFeeAmount(opts *bind.CallOpts, token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "calculateFeeAmount", token, amount, isSwap)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_SynapseCCTP *SynapseCCTPSession) CalculateFeeAmount(token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	return _SynapseCCTP.Contract.CalculateFeeAmount(&_SynapseCCTP.CallOpts, token, amount, isSwap)
}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_SynapseCCTP *SynapseCCTPCallerSession) CalculateFeeAmount(token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	return _SynapseCCTP.Contract.CalculateFeeAmount(&_SynapseCCTP.CallOpts, token, amount, isSwap)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseCCTP *SynapseCCTPCaller) ChainGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "chainGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseCCTP *SynapseCCTPSession) ChainGasAmount() (*big.Int, error) {
	return _SynapseCCTP.Contract.ChainGasAmount(&_SynapseCCTP.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseCCTP *SynapseCCTPCallerSession) ChainGasAmount() (*big.Int, error) {
	return _SynapseCCTP.Contract.ChainGasAmount(&_SynapseCCTP.CallOpts)
}

// CircleTokenPool is a free data retrieval call binding the contract method 0xa4b1d034.
//
// Solidity: function circleTokenPool(address ) view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) CircleTokenPool(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "circleTokenPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CircleTokenPool is a free data retrieval call binding the contract method 0xa4b1d034.
//
// Solidity: function circleTokenPool(address ) view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) CircleTokenPool(arg0 common.Address) (common.Address, error) {
	return _SynapseCCTP.Contract.CircleTokenPool(&_SynapseCCTP.CallOpts, arg0)
}

// CircleTokenPool is a free data retrieval call binding the contract method 0xa4b1d034.
//
// Solidity: function circleTokenPool(address ) view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) CircleTokenPool(arg0 common.Address) (common.Address, error) {
	return _SynapseCCTP.Contract.CircleTokenPool(&_SynapseCCTP.CallOpts, arg0)
}

// FeeStructures is a free data retrieval call binding the contract method 0xdc72495b.
//
// Solidity: function feeStructures(address ) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_SynapseCCTP *SynapseCCTPCaller) FeeStructures(opts *bind.CallOpts, arg0 common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "feeStructures", arg0)

	outstruct := new(struct {
		RelayerFee *big.Int
		MinBaseFee *big.Int
		MinSwapFee *big.Int
		MaxFee     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RelayerFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinBaseFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinSwapFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FeeStructures is a free data retrieval call binding the contract method 0xdc72495b.
//
// Solidity: function feeStructures(address ) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_SynapseCCTP *SynapseCCTPSession) FeeStructures(arg0 common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	return _SynapseCCTP.Contract.FeeStructures(&_SynapseCCTP.CallOpts, arg0)
}

// FeeStructures is a free data retrieval call binding the contract method 0xdc72495b.
//
// Solidity: function feeStructures(address ) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_SynapseCCTP *SynapseCCTPCallerSession) FeeStructures(arg0 common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	return _SynapseCCTP.Contract.FeeStructures(&_SynapseCCTP.CallOpts, arg0)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_SynapseCCTP *SynapseCCTPCaller) GetBridgeTokens(opts *bind.CallOpts) ([]BridgeToken, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "getBridgeTokens")

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_SynapseCCTP *SynapseCCTPSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _SynapseCCTP.Contract.GetBridgeTokens(&_SynapseCCTP.CallOpts)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_SynapseCCTP *SynapseCCTPCallerSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _SynapseCCTP.Contract.GetBridgeTokens(&_SynapseCCTP.CallOpts)
}

// GetLocalToken is a free data retrieval call binding the contract method 0xf879a41a.
//
// Solidity: function getLocalToken(uint32 remoteDomain, address remoteToken) view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) GetLocalToken(opts *bind.CallOpts, remoteDomain uint32, remoteToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "getLocalToken", remoteDomain, remoteToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocalToken is a free data retrieval call binding the contract method 0xf879a41a.
//
// Solidity: function getLocalToken(uint32 remoteDomain, address remoteToken) view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) GetLocalToken(remoteDomain uint32, remoteToken common.Address) (common.Address, error) {
	return _SynapseCCTP.Contract.GetLocalToken(&_SynapseCCTP.CallOpts, remoteDomain, remoteToken)
}

// GetLocalToken is a free data retrieval call binding the contract method 0xf879a41a.
//
// Solidity: function getLocalToken(uint32 remoteDomain, address remoteToken) view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) GetLocalToken(remoteDomain uint32, remoteToken common.Address) (common.Address, error) {
	return _SynapseCCTP.Contract.GetLocalToken(&_SynapseCCTP.CallOpts, remoteDomain, remoteToken)
}

// IsRequestFulfilled is a free data retrieval call binding the contract method 0x92a442ea.
//
// Solidity: function isRequestFulfilled(bytes32 requestID) view returns(bool)
func (_SynapseCCTP *SynapseCCTPCaller) IsRequestFulfilled(opts *bind.CallOpts, requestID [32]byte) (bool, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "isRequestFulfilled", requestID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRequestFulfilled is a free data retrieval call binding the contract method 0x92a442ea.
//
// Solidity: function isRequestFulfilled(bytes32 requestID) view returns(bool)
func (_SynapseCCTP *SynapseCCTPSession) IsRequestFulfilled(requestID [32]byte) (bool, error) {
	return _SynapseCCTP.Contract.IsRequestFulfilled(&_SynapseCCTP.CallOpts, requestID)
}

// IsRequestFulfilled is a free data retrieval call binding the contract method 0x92a442ea.
//
// Solidity: function isRequestFulfilled(bytes32 requestID) view returns(bool)
func (_SynapseCCTP *SynapseCCTPCallerSession) IsRequestFulfilled(requestID [32]byte) (bool, error) {
	return _SynapseCCTP.Contract.IsRequestFulfilled(&_SynapseCCTP.CallOpts, requestID)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseCCTP *SynapseCCTPCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseCCTP *SynapseCCTPSession) LocalDomain() (uint32, error) {
	return _SynapseCCTP.Contract.LocalDomain(&_SynapseCCTP.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseCCTP *SynapseCCTPCallerSession) LocalDomain() (uint32, error) {
	return _SynapseCCTP.Contract.LocalDomain(&_SynapseCCTP.CallOpts)
}

// MessageTransmitter is a free data retrieval call binding the contract method 0x7b04c181.
//
// Solidity: function messageTransmitter() view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) MessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "messageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageTransmitter is a free data retrieval call binding the contract method 0x7b04c181.
//
// Solidity: function messageTransmitter() view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) MessageTransmitter() (common.Address, error) {
	return _SynapseCCTP.Contract.MessageTransmitter(&_SynapseCCTP.CallOpts)
}

// MessageTransmitter is a free data retrieval call binding the contract method 0x7b04c181.
//
// Solidity: function messageTransmitter() view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) MessageTransmitter() (common.Address, error) {
	return _SynapseCCTP.Contract.MessageTransmitter(&_SynapseCCTP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) Owner() (common.Address, error) {
	return _SynapseCCTP.Contract.Owner(&_SynapseCCTP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) Owner() (common.Address, error) {
	return _SynapseCCTP.Contract.Owner(&_SynapseCCTP.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SynapseCCTP *SynapseCCTPCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SynapseCCTP *SynapseCCTPSession) ProtocolFee() (*big.Int, error) {
	return _SynapseCCTP.Contract.ProtocolFee(&_SynapseCCTP.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SynapseCCTP *SynapseCCTPCallerSession) ProtocolFee() (*big.Int, error) {
	return _SynapseCCTP.Contract.ProtocolFee(&_SynapseCCTP.CallOpts)
}

// RelayerFeeCollectors is a free data retrieval call binding the contract method 0x41f355ee.
//
// Solidity: function relayerFeeCollectors(address ) view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) RelayerFeeCollectors(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "relayerFeeCollectors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerFeeCollectors is a free data retrieval call binding the contract method 0x41f355ee.
//
// Solidity: function relayerFeeCollectors(address ) view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) RelayerFeeCollectors(arg0 common.Address) (common.Address, error) {
	return _SynapseCCTP.Contract.RelayerFeeCollectors(&_SynapseCCTP.CallOpts, arg0)
}

// RelayerFeeCollectors is a free data retrieval call binding the contract method 0x41f355ee.
//
// Solidity: function relayerFeeCollectors(address ) view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) RelayerFeeCollectors(arg0 common.Address) (common.Address, error) {
	return _SynapseCCTP.Contract.RelayerFeeCollectors(&_SynapseCCTP.CallOpts, arg0)
}

// RemoteDomainConfig is a free data retrieval call binding the contract method 0xe9259ab9.
//
// Solidity: function remoteDomainConfig(uint256 ) view returns(uint32 domain, address synapseCCTP)
func (_SynapseCCTP *SynapseCCTPCaller) RemoteDomainConfig(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Domain      uint32
	SynapseCCTP common.Address
}, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "remoteDomainConfig", arg0)

	outstruct := new(struct {
		Domain      uint32
		SynapseCCTP common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Domain = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.SynapseCCTP = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RemoteDomainConfig is a free data retrieval call binding the contract method 0xe9259ab9.
//
// Solidity: function remoteDomainConfig(uint256 ) view returns(uint32 domain, address synapseCCTP)
func (_SynapseCCTP *SynapseCCTPSession) RemoteDomainConfig(arg0 *big.Int) (struct {
	Domain      uint32
	SynapseCCTP common.Address
}, error) {
	return _SynapseCCTP.Contract.RemoteDomainConfig(&_SynapseCCTP.CallOpts, arg0)
}

// RemoteDomainConfig is a free data retrieval call binding the contract method 0xe9259ab9.
//
// Solidity: function remoteDomainConfig(uint256 ) view returns(uint32 domain, address synapseCCTP)
func (_SynapseCCTP *SynapseCCTPCallerSession) RemoteDomainConfig(arg0 *big.Int) (struct {
	Domain      uint32
	SynapseCCTP common.Address
}, error) {
	return _SynapseCCTP.Contract.RemoteDomainConfig(&_SynapseCCTP.CallOpts, arg0)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) SymbolToToken(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "symbolToToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) SymbolToToken(arg0 string) (common.Address, error) {
	return _SynapseCCTP.Contract.SymbolToToken(&_SynapseCCTP.CallOpts, arg0)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) SymbolToToken(arg0 string) (common.Address, error) {
	return _SynapseCCTP.Contract.SymbolToToken(&_SynapseCCTP.CallOpts, arg0)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_SynapseCCTP *SynapseCCTPCaller) TokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_SynapseCCTP *SynapseCCTPSession) TokenMessenger() (common.Address, error) {
	return _SynapseCCTP.Contract.TokenMessenger(&_SynapseCCTP.CallOpts)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_SynapseCCTP *SynapseCCTPCallerSession) TokenMessenger() (common.Address, error) {
	return _SynapseCCTP.Contract.TokenMessenger(&_SynapseCCTP.CallOpts)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_SynapseCCTP *SynapseCCTPCaller) TokenToSymbol(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "tokenToSymbol", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_SynapseCCTP *SynapseCCTPSession) TokenToSymbol(arg0 common.Address) (string, error) {
	return _SynapseCCTP.Contract.TokenToSymbol(&_SynapseCCTP.CallOpts, arg0)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_SynapseCCTP *SynapseCCTPCallerSession) TokenToSymbol(arg0 common.Address) (string, error) {
	return _SynapseCCTP.Contract.TokenToSymbol(&_SynapseCCTP.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0x4a85178d.
//
// Solidity: function addToken(string symbol, address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) AddToken(opts *bind.TransactOpts, symbol string, token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "addToken", symbol, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// AddToken is a paid mutator transaction binding the contract method 0x4a85178d.
//
// Solidity: function addToken(string symbol, address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTP *SynapseCCTPSession) AddToken(symbol string, token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.AddToken(&_SynapseCCTP.TransactOpts, symbol, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// AddToken is a paid mutator transaction binding the contract method 0x4a85178d.
//
// Solidity: function addToken(string symbol, address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) AddToken(symbol string, token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.AddToken(&_SynapseCCTP.TransactOpts, symbol, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) payable returns()
func (_SynapseCCTP *SynapseCCTPTransactor) ReceiveCircleToken(opts *bind.TransactOpts, message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "receiveCircleToken", message, signature, requestVersion, formattedRequest)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) payable returns()
func (_SynapseCCTP *SynapseCCTPSession) ReceiveCircleToken(message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.ReceiveCircleToken(&_SynapseCCTP.TransactOpts, message, signature, requestVersion, formattedRequest)
}

// ReceiveCircleToken is a paid mutator transaction binding the contract method 0x4a5ae51d.
//
// Solidity: function receiveCircleToken(bytes message, bytes signature, uint32 requestVersion, bytes formattedRequest) payable returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) ReceiveCircleToken(message []byte, signature []byte, requestVersion uint32, formattedRequest []byte) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.ReceiveCircleToken(&_SynapseCCTP.TransactOpts, message, signature, requestVersion, formattedRequest)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) RemoveToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "removeToken", token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_SynapseCCTP *SynapseCCTPSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.RemoveToken(&_SynapseCCTP.TransactOpts, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.RemoveToken(&_SynapseCCTP.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseCCTP *SynapseCCTPTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseCCTP *SynapseCCTPSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseCCTP.Contract.RenounceOwnership(&_SynapseCCTP.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseCCTP.Contract.RenounceOwnership(&_SynapseCCTP.TransactOpts)
}

// RescueGas is a paid mutator transaction binding the contract method 0x40432d51.
//
// Solidity: function rescueGas() returns()
func (_SynapseCCTP *SynapseCCTPTransactor) RescueGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "rescueGas")
}

// RescueGas is a paid mutator transaction binding the contract method 0x40432d51.
//
// Solidity: function rescueGas() returns()
func (_SynapseCCTP *SynapseCCTPSession) RescueGas() (*types.Transaction, error) {
	return _SynapseCCTP.Contract.RescueGas(&_SynapseCCTP.TransactOpts)
}

// RescueGas is a paid mutator transaction binding the contract method 0x40432d51.
//
// Solidity: function rescueGas() returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) RescueGas() (*types.Transaction, error) {
	return _SynapseCCTP.Contract.RescueGas(&_SynapseCCTP.TransactOpts)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SendCircleToken(opts *bind.TransactOpts, recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "sendCircleToken", recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_SynapseCCTP *SynapseCCTPSession) SendCircleToken(recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SendCircleToken(&_SynapseCCTP.TransactOpts, recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SendCircleToken is a paid mutator transaction binding the contract method 0x304ddb4c.
//
// Solidity: function sendCircleToken(address recipient, uint256 chainId, address burnToken, uint256 amount, uint32 requestVersion, bytes swapParams) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SendCircleToken(recipient common.Address, chainId *big.Int, burnToken common.Address, amount *big.Int, requestVersion uint32, swapParams []byte) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SendCircleToken(&_SynapseCCTP.TransactOpts, recipient, chainId, burnToken, amount, requestVersion, swapParams)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetChainGasAmount(opts *bind.TransactOpts, newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setChainGasAmount", newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetChainGasAmount(&_SynapseCCTP.TransactOpts, newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetChainGasAmount(&_SynapseCCTP.TransactOpts, newChainGasAmount)
}

// SetCircleTokenPool is a paid mutator transaction binding the contract method 0x2cc9e7e5.
//
// Solidity: function setCircleTokenPool(address circleToken, address pool) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetCircleTokenPool(opts *bind.TransactOpts, circleToken common.Address, pool common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setCircleTokenPool", circleToken, pool)
}

// SetCircleTokenPool is a paid mutator transaction binding the contract method 0x2cc9e7e5.
//
// Solidity: function setCircleTokenPool(address circleToken, address pool) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetCircleTokenPool(circleToken common.Address, pool common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetCircleTokenPool(&_SynapseCCTP.TransactOpts, circleToken, pool)
}

// SetCircleTokenPool is a paid mutator transaction binding the contract method 0x2cc9e7e5.
//
// Solidity: function setCircleTokenPool(address circleToken, address pool) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SetCircleTokenPool(circleToken common.Address, pool common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetCircleTokenPool(&_SynapseCCTP.TransactOpts, circleToken, pool)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetFeeCollector(opts *bind.TransactOpts, feeCollector common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setFeeCollector", feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetFeeCollector(&_SynapseCCTP.TransactOpts, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetFeeCollector(&_SynapseCCTP.TransactOpts, feeCollector)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newProtocolFee) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetProtocolFee(opts *bind.TransactOpts, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setProtocolFee", newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newProtocolFee) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetProtocolFee(newProtocolFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetProtocolFee(&_SynapseCCTP.TransactOpts, newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newProtocolFee) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SetProtocolFee(newProtocolFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetProtocolFee(&_SynapseCCTP.TransactOpts, newProtocolFee)
}

// SetRemoteDomainConfig is a paid mutator transaction binding the contract method 0xe9bbb36d.
//
// Solidity: function setRemoteDomainConfig(uint256 remoteChainId, uint32 remoteDomain, address remoteSynapseCCTP) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetRemoteDomainConfig(opts *bind.TransactOpts, remoteChainId *big.Int, remoteDomain uint32, remoteSynapseCCTP common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setRemoteDomainConfig", remoteChainId, remoteDomain, remoteSynapseCCTP)
}

// SetRemoteDomainConfig is a paid mutator transaction binding the contract method 0xe9bbb36d.
//
// Solidity: function setRemoteDomainConfig(uint256 remoteChainId, uint32 remoteDomain, address remoteSynapseCCTP) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetRemoteDomainConfig(remoteChainId *big.Int, remoteDomain uint32, remoteSynapseCCTP common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetRemoteDomainConfig(&_SynapseCCTP.TransactOpts, remoteChainId, remoteDomain, remoteSynapseCCTP)
}

// SetRemoteDomainConfig is a paid mutator transaction binding the contract method 0xe9bbb36d.
//
// Solidity: function setRemoteDomainConfig(uint256 remoteChainId, uint32 remoteDomain, address remoteSynapseCCTP) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SetRemoteDomainConfig(remoteChainId *big.Int, remoteDomain uint32, remoteSynapseCCTP common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetRemoteDomainConfig(&_SynapseCCTP.TransactOpts, remoteChainId, remoteDomain, remoteSynapseCCTP)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x4bdb4eed.
//
// Solidity: function setTokenFee(address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetTokenFee(opts *bind.TransactOpts, token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setTokenFee", token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x4bdb4eed.
//
// Solidity: function setTokenFee(address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetTokenFee(token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetTokenFee(&_SynapseCCTP.TransactOpts, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x4bdb4eed.
//
// Solidity: function setTokenFee(address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) SetTokenFee(token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetTokenFee(&_SynapseCCTP.TransactOpts, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseCCTP *SynapseCCTPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.TransferOwnership(&_SynapseCCTP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.TransferOwnership(&_SynapseCCTP.TransactOpts, newOwner)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) WithdrawProtocolFees(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "withdrawProtocolFees", token)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) returns()
func (_SynapseCCTP *SynapseCCTPSession) WithdrawProtocolFees(token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.WithdrawProtocolFees(&_SynapseCCTP.TransactOpts, token)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) WithdrawProtocolFees(token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.WithdrawProtocolFees(&_SynapseCCTP.TransactOpts, token)
}

// WithdrawRelayerFees is a paid mutator transaction binding the contract method 0xf7265b3a.
//
// Solidity: function withdrawRelayerFees(address token) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) WithdrawRelayerFees(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "withdrawRelayerFees", token)
}

// WithdrawRelayerFees is a paid mutator transaction binding the contract method 0xf7265b3a.
//
// Solidity: function withdrawRelayerFees(address token) returns()
func (_SynapseCCTP *SynapseCCTPSession) WithdrawRelayerFees(token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.WithdrawRelayerFees(&_SynapseCCTP.TransactOpts, token)
}

// WithdrawRelayerFees is a paid mutator transaction binding the contract method 0xf7265b3a.
//
// Solidity: function withdrawRelayerFees(address token) returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) WithdrawRelayerFees(token common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.WithdrawRelayerFees(&_SynapseCCTP.TransactOpts, token)
}

// SynapseCCTPChainGasAirdroppedIterator is returned from FilterChainGasAirdropped and is used to iterate over the raw logs and unpacked data for ChainGasAirdropped events raised by the SynapseCCTP contract.
type SynapseCCTPChainGasAirdroppedIterator struct {
	Event *SynapseCCTPChainGasAirdropped // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPChainGasAirdroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPChainGasAirdropped)
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
		it.Event = new(SynapseCCTPChainGasAirdropped)
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
func (it *SynapseCCTPChainGasAirdroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPChainGasAirdroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPChainGasAirdropped represents a ChainGasAirdropped event raised by the SynapseCCTP contract.
type SynapseCCTPChainGasAirdropped struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChainGasAirdropped is a free log retrieval operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterChainGasAirdropped(opts *bind.FilterOpts) (*SynapseCCTPChainGasAirdroppedIterator, error) {

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "ChainGasAirdropped")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPChainGasAirdroppedIterator{contract: _SynapseCCTP.contract, event: "ChainGasAirdropped", logs: logs, sub: sub}, nil
}

// WatchChainGasAirdropped is a free log subscription operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchChainGasAirdropped(opts *bind.WatchOpts, sink chan<- *SynapseCCTPChainGasAirdropped) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "ChainGasAirdropped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPChainGasAirdropped)
				if err := _SynapseCCTP.contract.UnpackLog(event, "ChainGasAirdropped", log); err != nil {
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

// ParseChainGasAirdropped is a log parse operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseChainGasAirdropped(log types.Log) (*SynapseCCTPChainGasAirdropped, error) {
	event := new(SynapseCCTPChainGasAirdropped)
	if err := _SynapseCCTP.contract.UnpackLog(event, "ChainGasAirdropped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPChainGasAmountUpdatedIterator is returned from FilterChainGasAmountUpdated and is used to iterate over the raw logs and unpacked data for ChainGasAmountUpdated events raised by the SynapseCCTP contract.
type SynapseCCTPChainGasAmountUpdatedIterator struct {
	Event *SynapseCCTPChainGasAmountUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPChainGasAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPChainGasAmountUpdated)
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
		it.Event = new(SynapseCCTPChainGasAmountUpdated)
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
func (it *SynapseCCTPChainGasAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPChainGasAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPChainGasAmountUpdated represents a ChainGasAmountUpdated event raised by the SynapseCCTP contract.
type SynapseCCTPChainGasAmountUpdated struct {
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChainGasAmountUpdated is a free log retrieval operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterChainGasAmountUpdated(opts *bind.FilterOpts) (*SynapseCCTPChainGasAmountUpdatedIterator, error) {

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPChainGasAmountUpdatedIterator{contract: _SynapseCCTP.contract, event: "ChainGasAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchChainGasAmountUpdated is a free log subscription operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchChainGasAmountUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPChainGasAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPChainGasAmountUpdated)
				if err := _SynapseCCTP.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
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

// ParseChainGasAmountUpdated is a log parse operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseChainGasAmountUpdated(log types.Log) (*SynapseCCTPChainGasAmountUpdated, error) {
	event := new(SynapseCCTPChainGasAmountUpdated)
	if err := _SynapseCCTP.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPCircleRequestFulfilledIterator is returned from FilterCircleRequestFulfilled and is used to iterate over the raw logs and unpacked data for CircleRequestFulfilled events raised by the SynapseCCTP contract.
type SynapseCCTPCircleRequestFulfilledIterator struct {
	Event *SynapseCCTPCircleRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPCircleRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPCircleRequestFulfilled)
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
		it.Event = new(SynapseCCTPCircleRequestFulfilled)
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
func (it *SynapseCCTPCircleRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPCircleRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPCircleRequestFulfilled represents a CircleRequestFulfilled event raised by the SynapseCCTP contract.
type SynapseCCTPCircleRequestFulfilled struct {
	Recipient common.Address
	MintToken common.Address
	Fee       *big.Int
	Token     common.Address
	Amount    *big.Int
	RequestID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestFulfilled is a free log retrieval operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterCircleRequestFulfilled(opts *bind.FilterOpts, recipient []common.Address, requestID [][32]byte) (*SynapseCCTPCircleRequestFulfilledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "CircleRequestFulfilled", recipientRule, requestIDRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPCircleRequestFulfilledIterator{contract: _SynapseCCTP.contract, event: "CircleRequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchCircleRequestFulfilled is a free log subscription operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchCircleRequestFulfilled(opts *bind.WatchOpts, sink chan<- *SynapseCCTPCircleRequestFulfilled, recipient []common.Address, requestID [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "CircleRequestFulfilled", recipientRule, requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPCircleRequestFulfilled)
				if err := _SynapseCCTP.contract.UnpackLog(event, "CircleRequestFulfilled", log); err != nil {
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

// ParseCircleRequestFulfilled is a log parse operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseCircleRequestFulfilled(log types.Log) (*SynapseCCTPCircleRequestFulfilled, error) {
	event := new(SynapseCCTPCircleRequestFulfilled)
	if err := _SynapseCCTP.contract.UnpackLog(event, "CircleRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPCircleRequestSentIterator is returned from FilterCircleRequestSent and is used to iterate over the raw logs and unpacked data for CircleRequestSent events raised by the SynapseCCTP contract.
type SynapseCCTPCircleRequestSentIterator struct {
	Event *SynapseCCTPCircleRequestSent // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPCircleRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPCircleRequestSent)
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
		it.Event = new(SynapseCCTPCircleRequestSent)
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
func (it *SynapseCCTPCircleRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPCircleRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPCircleRequestSent represents a CircleRequestSent event raised by the SynapseCCTP contract.
type SynapseCCTPCircleRequestSent struct {
	ChainId          *big.Int
	Nonce            uint64
	Token            common.Address
	Amount           *big.Int
	RequestVersion   uint32
	FormattedRequest []byte
	RequestID        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestSent is a free log retrieval operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterCircleRequestSent(opts *bind.FilterOpts, requestID [][32]byte) (*SynapseCCTPCircleRequestSentIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "CircleRequestSent", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPCircleRequestSentIterator{contract: _SynapseCCTP.contract, event: "CircleRequestSent", logs: logs, sub: sub}, nil
}

// WatchCircleRequestSent is a free log subscription operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchCircleRequestSent(opts *bind.WatchOpts, sink chan<- *SynapseCCTPCircleRequestSent, requestID [][32]byte) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "CircleRequestSent", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPCircleRequestSent)
				if err := _SynapseCCTP.contract.UnpackLog(event, "CircleRequestSent", log); err != nil {
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

// ParseCircleRequestSent is a log parse operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseCircleRequestSent(log types.Log) (*SynapseCCTPCircleRequestSent, error) {
	event := new(SynapseCCTPCircleRequestSent)
	if err := _SynapseCCTP.contract.UnpackLog(event, "CircleRequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the SynapseCCTP contract.
type SynapseCCTPFeeCollectedIterator struct {
	Event *SynapseCCTPFeeCollected // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeeCollected)
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
		it.Event = new(SynapseCCTPFeeCollected)
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
func (it *SynapseCCTPFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeeCollected represents a FeeCollected event raised by the SynapseCCTP contract.
type SynapseCCTPFeeCollected struct {
	FeeCollector      common.Address
	RelayerFeeAmount  *big.Int
	ProtocolFeeAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFeeCollected is a free log retrieval operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*SynapseCCTPFeeCollectedIterator, error) {

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeeCollectedIterator{contract: _SynapseCCTP.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeeCollected) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeeCollected)
				if err := _SynapseCCTP.contract.UnpackLog(event, "FeeCollected", log); err != nil {
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

// ParseFeeCollected is a log parse operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseFeeCollected(log types.Log) (*SynapseCCTPFeeCollected, error) {
	event := new(SynapseCCTPFeeCollected)
	if err := _SynapseCCTP.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeeCollectorUpdatedIterator is returned from FilterFeeCollectorUpdated and is used to iterate over the raw logs and unpacked data for FeeCollectorUpdated events raised by the SynapseCCTP contract.
type SynapseCCTPFeeCollectorUpdatedIterator struct {
	Event *SynapseCCTPFeeCollectorUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeeCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeeCollectorUpdated)
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
		it.Event = new(SynapseCCTPFeeCollectorUpdated)
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
func (it *SynapseCCTPFeeCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeeCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeeCollectorUpdated represents a FeeCollectorUpdated event raised by the SynapseCCTP contract.
type SynapseCCTPFeeCollectorUpdated struct {
	Relayer         common.Address
	OldFeeCollector common.Address
	NewFeeCollector common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorUpdated is a free log retrieval operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterFeeCollectorUpdated(opts *bind.FilterOpts, relayer []common.Address) (*SynapseCCTPFeeCollectorUpdatedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "FeeCollectorUpdated", relayerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeeCollectorUpdatedIterator{contract: _SynapseCCTP.contract, event: "FeeCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorUpdated is a free log subscription operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchFeeCollectorUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeeCollectorUpdated, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "FeeCollectorUpdated", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeeCollectorUpdated)
				if err := _SynapseCCTP.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
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

// ParseFeeCollectorUpdated is a log parse operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseFeeCollectorUpdated(log types.Log) (*SynapseCCTPFeeCollectorUpdated, error) {
	event := new(SynapseCCTPFeeCollectorUpdated)
	if err := _SynapseCCTP.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynapseCCTP contract.
type SynapseCCTPOwnershipTransferredIterator struct {
	Event *SynapseCCTPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPOwnershipTransferred)
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
		it.Event = new(SynapseCCTPOwnershipTransferred)
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
func (it *SynapseCCTPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPOwnershipTransferred represents a OwnershipTransferred event raised by the SynapseCCTP contract.
type SynapseCCTPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynapseCCTPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPOwnershipTransferredIterator{contract: _SynapseCCTP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynapseCCTPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPOwnershipTransferred)
				if err := _SynapseCCTP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SynapseCCTP *SynapseCCTPFilterer) ParseOwnershipTransferred(log types.Log) (*SynapseCCTPOwnershipTransferred, error) {
	event := new(SynapseCCTPOwnershipTransferred)
	if err := _SynapseCCTP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPProtocolFeeUpdatedIterator is returned from FilterProtocolFeeUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeUpdated events raised by the SynapseCCTP contract.
type SynapseCCTPProtocolFeeUpdatedIterator struct {
	Event *SynapseCCTPProtocolFeeUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPProtocolFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPProtocolFeeUpdated)
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
		it.Event = new(SynapseCCTPProtocolFeeUpdated)
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
func (it *SynapseCCTPProtocolFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPProtocolFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPProtocolFeeUpdated represents a ProtocolFeeUpdated event raised by the SynapseCCTP contract.
type SynapseCCTPProtocolFeeUpdated struct {
	NewProtocolFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeUpdated is a free log retrieval operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterProtocolFeeUpdated(opts *bind.FilterOpts) (*SynapseCCTPProtocolFeeUpdatedIterator, error) {

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "ProtocolFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPProtocolFeeUpdatedIterator{contract: _SynapseCCTP.contract, event: "ProtocolFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeUpdated is a free log subscription operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchProtocolFeeUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPProtocolFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "ProtocolFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPProtocolFeeUpdated)
				if err := _SynapseCCTP.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
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

// ParseProtocolFeeUpdated is a log parse operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseProtocolFeeUpdated(log types.Log) (*SynapseCCTPProtocolFeeUpdated, error) {
	event := new(SynapseCCTPProtocolFeeUpdated)
	if err := _SynapseCCTP.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPEventsMetaData contains all meta data concerning the SynapseCCTPEvents contract.
var SynapseCCTPEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestSent\",\"type\":\"event\"}]",
}

// SynapseCCTPEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseCCTPEventsMetaData.ABI instead.
var SynapseCCTPEventsABI = SynapseCCTPEventsMetaData.ABI

// SynapseCCTPEvents is an auto generated Go binding around an Ethereum contract.
type SynapseCCTPEvents struct {
	SynapseCCTPEventsCaller     // Read-only binding to the contract
	SynapseCCTPEventsTransactor // Write-only binding to the contract
	SynapseCCTPEventsFilterer   // Log filterer for contract events
}

// SynapseCCTPEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseCCTPEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseCCTPEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseCCTPEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseCCTPEventsSession struct {
	Contract     *SynapseCCTPEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SynapseCCTPEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseCCTPEventsCallerSession struct {
	Contract *SynapseCCTPEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SynapseCCTPEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseCCTPEventsTransactorSession struct {
	Contract     *SynapseCCTPEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SynapseCCTPEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseCCTPEventsRaw struct {
	Contract *SynapseCCTPEvents // Generic contract binding to access the raw methods on
}

// SynapseCCTPEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseCCTPEventsCallerRaw struct {
	Contract *SynapseCCTPEventsCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseCCTPEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseCCTPEventsTransactorRaw struct {
	Contract *SynapseCCTPEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseCCTPEvents creates a new instance of SynapseCCTPEvents, bound to a specific deployed contract.
func NewSynapseCCTPEvents(address common.Address, backend bind.ContractBackend) (*SynapseCCTPEvents, error) {
	contract, err := bindSynapseCCTPEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEvents{SynapseCCTPEventsCaller: SynapseCCTPEventsCaller{contract: contract}, SynapseCCTPEventsTransactor: SynapseCCTPEventsTransactor{contract: contract}, SynapseCCTPEventsFilterer: SynapseCCTPEventsFilterer{contract: contract}}, nil
}

// NewSynapseCCTPEventsCaller creates a new read-only instance of SynapseCCTPEvents, bound to a specific deployed contract.
func NewSynapseCCTPEventsCaller(address common.Address, caller bind.ContractCaller) (*SynapseCCTPEventsCaller, error) {
	contract, err := bindSynapseCCTPEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsCaller{contract: contract}, nil
}

// NewSynapseCCTPEventsTransactor creates a new write-only instance of SynapseCCTPEvents, bound to a specific deployed contract.
func NewSynapseCCTPEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseCCTPEventsTransactor, error) {
	contract, err := bindSynapseCCTPEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsTransactor{contract: contract}, nil
}

// NewSynapseCCTPEventsFilterer creates a new log filterer instance of SynapseCCTPEvents, bound to a specific deployed contract.
func NewSynapseCCTPEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseCCTPEventsFilterer, error) {
	contract, err := bindSynapseCCTPEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsFilterer{contract: contract}, nil
}

// bindSynapseCCTPEvents binds a generic wrapper to an already deployed contract.
func bindSynapseCCTPEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseCCTPEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTPEvents *SynapseCCTPEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTPEvents.Contract.SynapseCCTPEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTPEvents *SynapseCCTPEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPEvents.Contract.SynapseCCTPEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTPEvents *SynapseCCTPEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTPEvents.Contract.SynapseCCTPEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTPEvents *SynapseCCTPEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTPEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTPEvents *SynapseCCTPEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTPEvents *SynapseCCTPEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTPEvents.Contract.contract.Transact(opts, method, params...)
}

// SynapseCCTPEventsCircleRequestFulfilledIterator is returned from FilterCircleRequestFulfilled and is used to iterate over the raw logs and unpacked data for CircleRequestFulfilled events raised by the SynapseCCTPEvents contract.
type SynapseCCTPEventsCircleRequestFulfilledIterator struct {
	Event *SynapseCCTPEventsCircleRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPEventsCircleRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPEventsCircleRequestFulfilled)
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
		it.Event = new(SynapseCCTPEventsCircleRequestFulfilled)
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
func (it *SynapseCCTPEventsCircleRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPEventsCircleRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPEventsCircleRequestFulfilled represents a CircleRequestFulfilled event raised by the SynapseCCTPEvents contract.
type SynapseCCTPEventsCircleRequestFulfilled struct {
	Recipient common.Address
	MintToken common.Address
	Fee       *big.Int
	Token     common.Address
	Amount    *big.Int
	RequestID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestFulfilled is a free log retrieval operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) FilterCircleRequestFulfilled(opts *bind.FilterOpts, recipient []common.Address, requestID [][32]byte) (*SynapseCCTPEventsCircleRequestFulfilledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.FilterLogs(opts, "CircleRequestFulfilled", recipientRule, requestIDRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsCircleRequestFulfilledIterator{contract: _SynapseCCTPEvents.contract, event: "CircleRequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchCircleRequestFulfilled is a free log subscription operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) WatchCircleRequestFulfilled(opts *bind.WatchOpts, sink chan<- *SynapseCCTPEventsCircleRequestFulfilled, recipient []common.Address, requestID [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.WatchLogs(opts, "CircleRequestFulfilled", recipientRule, requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPEventsCircleRequestFulfilled)
				if err := _SynapseCCTPEvents.contract.UnpackLog(event, "CircleRequestFulfilled", log); err != nil {
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

// ParseCircleRequestFulfilled is a log parse operation binding the contract event 0xeaf2537b3a5c10387b14e2c0e57b1e11b46ff39b0f4ead5dac98cb0f4fd2118f.
//
// Solidity: event CircleRequestFulfilled(address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 indexed requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) ParseCircleRequestFulfilled(log types.Log) (*SynapseCCTPEventsCircleRequestFulfilled, error) {
	event := new(SynapseCCTPEventsCircleRequestFulfilled)
	if err := _SynapseCCTPEvents.contract.UnpackLog(event, "CircleRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPEventsCircleRequestSentIterator is returned from FilterCircleRequestSent and is used to iterate over the raw logs and unpacked data for CircleRequestSent events raised by the SynapseCCTPEvents contract.
type SynapseCCTPEventsCircleRequestSentIterator struct {
	Event *SynapseCCTPEventsCircleRequestSent // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPEventsCircleRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPEventsCircleRequestSent)
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
		it.Event = new(SynapseCCTPEventsCircleRequestSent)
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
func (it *SynapseCCTPEventsCircleRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPEventsCircleRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPEventsCircleRequestSent represents a CircleRequestSent event raised by the SynapseCCTPEvents contract.
type SynapseCCTPEventsCircleRequestSent struct {
	ChainId          *big.Int
	Nonce            uint64
	Token            common.Address
	Amount           *big.Int
	RequestVersion   uint32
	FormattedRequest []byte
	RequestID        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestSent is a free log retrieval operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) FilterCircleRequestSent(opts *bind.FilterOpts, requestID [][32]byte) (*SynapseCCTPEventsCircleRequestSentIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.FilterLogs(opts, "CircleRequestSent", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsCircleRequestSentIterator{contract: _SynapseCCTPEvents.contract, event: "CircleRequestSent", logs: logs, sub: sub}, nil
}

// WatchCircleRequestSent is a free log subscription operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) WatchCircleRequestSent(opts *bind.WatchOpts, sink chan<- *SynapseCCTPEventsCircleRequestSent, requestID [][32]byte) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.WatchLogs(opts, "CircleRequestSent", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPEventsCircleRequestSent)
				if err := _SynapseCCTPEvents.contract.UnpackLog(event, "CircleRequestSent", log); err != nil {
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

// ParseCircleRequestSent is a log parse operation binding the contract event 0x4ce96273f442a9bc593fea917ea7e8c2a009befc78ba3e334008948c7addf22a.
//
// Solidity: event CircleRequestSent(uint256 chainId, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 indexed requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) ParseCircleRequestSent(log types.Log) (*SynapseCCTPEventsCircleRequestSent, error) {
	event := new(SynapseCCTPEventsCircleRequestSent)
	if err := _SynapseCCTPEvents.contract.UnpackLog(event, "CircleRequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesMetaData contains all meta data concerning the SynapseCCTPFees contract.
var SynapseCCTPFeesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CCTPGasRescueFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectProtocolFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPSymbolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPSymbolIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPTokenAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPTokenNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CastOverflow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAirdropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayerFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accumulatedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSwap\",\"type\":\"bool\"}],\"name\":\"calculateFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeStructures\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"relayerFee\",\"type\":\"uint40\"},{\"internalType\":\"uint72\",\"name\":\"minBaseFee\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"minSwapFee\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"maxFee\",\"type\":\"uint72\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"bridgeTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayerFeeCollectors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescueGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"setTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d4a67c6d": "accumulatedFees(address,address)",
		"4a85178d": "addToken(string,address,uint256,uint256,uint256,uint256)",
		"0d25aafe": "calculateFeeAmount(address,uint256,bool)",
		"e00a83e0": "chainGasAmount()",
		"dc72495b": "feeStructures(address)",
		"9c1d060e": "getBridgeTokens()",
		"8da5cb5b": "owner()",
		"b0e21e8a": "protocolFee()",
		"41f355ee": "relayerFeeCollectors(address)",
		"5fa7b584": "removeToken(address)",
		"715018a6": "renounceOwnership()",
		"40432d51": "rescueGas()",
		"b250fe6b": "setChainGasAmount(uint256)",
		"a42dce80": "setFeeCollector(address)",
		"787dce3d": "setProtocolFee(uint256)",
		"4bdb4eed": "setTokenFee(address,uint256,uint256,uint256,uint256)",
		"a5bc29c2": "symbolToToken(string)",
		"0ba36121": "tokenToSymbol(address)",
		"f2fde38b": "transferOwnership(address)",
	},
}

// SynapseCCTPFeesABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseCCTPFeesMetaData.ABI instead.
var SynapseCCTPFeesABI = SynapseCCTPFeesMetaData.ABI

// Deprecated: Use SynapseCCTPFeesMetaData.Sigs instead.
// SynapseCCTPFeesFuncSigs maps the 4-byte function signature to its string representation.
var SynapseCCTPFeesFuncSigs = SynapseCCTPFeesMetaData.Sigs

// SynapseCCTPFees is an auto generated Go binding around an Ethereum contract.
type SynapseCCTPFees struct {
	SynapseCCTPFeesCaller     // Read-only binding to the contract
	SynapseCCTPFeesTransactor // Write-only binding to the contract
	SynapseCCTPFeesFilterer   // Log filterer for contract events
}

// SynapseCCTPFeesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseCCTPFeesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPFeesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseCCTPFeesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPFeesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseCCTPFeesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPFeesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseCCTPFeesSession struct {
	Contract     *SynapseCCTPFees  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseCCTPFeesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseCCTPFeesCallerSession struct {
	Contract *SynapseCCTPFeesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SynapseCCTPFeesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseCCTPFeesTransactorSession struct {
	Contract     *SynapseCCTPFeesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SynapseCCTPFeesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseCCTPFeesRaw struct {
	Contract *SynapseCCTPFees // Generic contract binding to access the raw methods on
}

// SynapseCCTPFeesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseCCTPFeesCallerRaw struct {
	Contract *SynapseCCTPFeesCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseCCTPFeesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseCCTPFeesTransactorRaw struct {
	Contract *SynapseCCTPFeesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseCCTPFees creates a new instance of SynapseCCTPFees, bound to a specific deployed contract.
func NewSynapseCCTPFees(address common.Address, backend bind.ContractBackend) (*SynapseCCTPFees, error) {
	contract, err := bindSynapseCCTPFees(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFees{SynapseCCTPFeesCaller: SynapseCCTPFeesCaller{contract: contract}, SynapseCCTPFeesTransactor: SynapseCCTPFeesTransactor{contract: contract}, SynapseCCTPFeesFilterer: SynapseCCTPFeesFilterer{contract: contract}}, nil
}

// NewSynapseCCTPFeesCaller creates a new read-only instance of SynapseCCTPFees, bound to a specific deployed contract.
func NewSynapseCCTPFeesCaller(address common.Address, caller bind.ContractCaller) (*SynapseCCTPFeesCaller, error) {
	contract, err := bindSynapseCCTPFees(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesCaller{contract: contract}, nil
}

// NewSynapseCCTPFeesTransactor creates a new write-only instance of SynapseCCTPFees, bound to a specific deployed contract.
func NewSynapseCCTPFeesTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseCCTPFeesTransactor, error) {
	contract, err := bindSynapseCCTPFees(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesTransactor{contract: contract}, nil
}

// NewSynapseCCTPFeesFilterer creates a new log filterer instance of SynapseCCTPFees, bound to a specific deployed contract.
func NewSynapseCCTPFeesFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseCCTPFeesFilterer, error) {
	contract, err := bindSynapseCCTPFees(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesFilterer{contract: contract}, nil
}

// bindSynapseCCTPFees binds a generic wrapper to an already deployed contract.
func bindSynapseCCTPFees(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseCCTPFeesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTPFees *SynapseCCTPFeesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTPFees.Contract.SynapseCCTPFeesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTPFees *SynapseCCTPFeesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SynapseCCTPFeesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTPFees *SynapseCCTPFeesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SynapseCCTPFeesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTPFees *SynapseCCTPFeesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTPFees.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0xd4a67c6d.
//
// Solidity: function accumulatedFees(address , address ) view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) AccumulatedFees(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "accumulatedFees", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFees is a free data retrieval call binding the contract method 0xd4a67c6d.
//
// Solidity: function accumulatedFees(address , address ) view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) AccumulatedFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SynapseCCTPFees.Contract.AccumulatedFees(&_SynapseCCTPFees.CallOpts, arg0, arg1)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0xd4a67c6d.
//
// Solidity: function accumulatedFees(address , address ) view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) AccumulatedFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SynapseCCTPFees.Contract.AccumulatedFees(&_SynapseCCTPFees.CallOpts, arg0, arg1)
}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) CalculateFeeAmount(opts *bind.CallOpts, token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "calculateFeeAmount", token, amount, isSwap)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) CalculateFeeAmount(token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	return _SynapseCCTPFees.Contract.CalculateFeeAmount(&_SynapseCCTPFees.CallOpts, token, amount, isSwap)
}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) CalculateFeeAmount(token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	return _SynapseCCTPFees.Contract.CalculateFeeAmount(&_SynapseCCTPFees.CallOpts, token, amount, isSwap)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) ChainGasAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "chainGasAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) ChainGasAmount() (*big.Int, error) {
	return _SynapseCCTPFees.Contract.ChainGasAmount(&_SynapseCCTPFees.CallOpts)
}

// ChainGasAmount is a free data retrieval call binding the contract method 0xe00a83e0.
//
// Solidity: function chainGasAmount() view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) ChainGasAmount() (*big.Int, error) {
	return _SynapseCCTPFees.Contract.ChainGasAmount(&_SynapseCCTPFees.CallOpts)
}

// FeeStructures is a free data retrieval call binding the contract method 0xdc72495b.
//
// Solidity: function feeStructures(address ) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) FeeStructures(opts *bind.CallOpts, arg0 common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "feeStructures", arg0)

	outstruct := new(struct {
		RelayerFee *big.Int
		MinBaseFee *big.Int
		MinSwapFee *big.Int
		MaxFee     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RelayerFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinBaseFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinSwapFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FeeStructures is a free data retrieval call binding the contract method 0xdc72495b.
//
// Solidity: function feeStructures(address ) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) FeeStructures(arg0 common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	return _SynapseCCTPFees.Contract.FeeStructures(&_SynapseCCTPFees.CallOpts, arg0)
}

// FeeStructures is a free data retrieval call binding the contract method 0xdc72495b.
//
// Solidity: function feeStructures(address ) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) FeeStructures(arg0 common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	return _SynapseCCTPFees.Contract.FeeStructures(&_SynapseCCTPFees.CallOpts, arg0)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) GetBridgeTokens(opts *bind.CallOpts) ([]BridgeToken, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "getBridgeTokens")

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _SynapseCCTPFees.Contract.GetBridgeTokens(&_SynapseCCTPFees.CallOpts)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _SynapseCCTPFees.Contract.GetBridgeTokens(&_SynapseCCTPFees.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) Owner() (common.Address, error) {
	return _SynapseCCTPFees.Contract.Owner(&_SynapseCCTPFees.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) Owner() (common.Address, error) {
	return _SynapseCCTPFees.Contract.Owner(&_SynapseCCTPFees.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) ProtocolFee() (*big.Int, error) {
	return _SynapseCCTPFees.Contract.ProtocolFee(&_SynapseCCTPFees.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) ProtocolFee() (*big.Int, error) {
	return _SynapseCCTPFees.Contract.ProtocolFee(&_SynapseCCTPFees.CallOpts)
}

// RelayerFeeCollectors is a free data retrieval call binding the contract method 0x41f355ee.
//
// Solidity: function relayerFeeCollectors(address ) view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) RelayerFeeCollectors(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "relayerFeeCollectors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerFeeCollectors is a free data retrieval call binding the contract method 0x41f355ee.
//
// Solidity: function relayerFeeCollectors(address ) view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) RelayerFeeCollectors(arg0 common.Address) (common.Address, error) {
	return _SynapseCCTPFees.Contract.RelayerFeeCollectors(&_SynapseCCTPFees.CallOpts, arg0)
}

// RelayerFeeCollectors is a free data retrieval call binding the contract method 0x41f355ee.
//
// Solidity: function relayerFeeCollectors(address ) view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) RelayerFeeCollectors(arg0 common.Address) (common.Address, error) {
	return _SynapseCCTPFees.Contract.RelayerFeeCollectors(&_SynapseCCTPFees.CallOpts, arg0)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) SymbolToToken(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "symbolToToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) SymbolToToken(arg0 string) (common.Address, error) {
	return _SynapseCCTPFees.Contract.SymbolToToken(&_SynapseCCTPFees.CallOpts, arg0)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) SymbolToToken(arg0 string) (common.Address, error) {
	return _SynapseCCTPFees.Contract.SymbolToToken(&_SynapseCCTPFees.CallOpts, arg0)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_SynapseCCTPFees *SynapseCCTPFeesCaller) TokenToSymbol(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _SynapseCCTPFees.contract.Call(opts, &out, "tokenToSymbol", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_SynapseCCTPFees *SynapseCCTPFeesSession) TokenToSymbol(arg0 common.Address) (string, error) {
	return _SynapseCCTPFees.Contract.TokenToSymbol(&_SynapseCCTPFees.CallOpts, arg0)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_SynapseCCTPFees *SynapseCCTPFeesCallerSession) TokenToSymbol(arg0 common.Address) (string, error) {
	return _SynapseCCTPFees.Contract.TokenToSymbol(&_SynapseCCTPFees.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0x4a85178d.
//
// Solidity: function addToken(string symbol, address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) AddToken(opts *bind.TransactOpts, symbol string, token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "addToken", symbol, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// AddToken is a paid mutator transaction binding the contract method 0x4a85178d.
//
// Solidity: function addToken(string symbol, address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) AddToken(symbol string, token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.AddToken(&_SynapseCCTPFees.TransactOpts, symbol, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// AddToken is a paid mutator transaction binding the contract method 0x4a85178d.
//
// Solidity: function addToken(string symbol, address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) AddToken(symbol string, token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.AddToken(&_SynapseCCTPFees.TransactOpts, symbol, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) RemoveToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "removeToken", token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.RemoveToken(&_SynapseCCTPFees.TransactOpts, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.RemoveToken(&_SynapseCCTPFees.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.RenounceOwnership(&_SynapseCCTPFees.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.RenounceOwnership(&_SynapseCCTPFees.TransactOpts)
}

// RescueGas is a paid mutator transaction binding the contract method 0x40432d51.
//
// Solidity: function rescueGas() returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) RescueGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "rescueGas")
}

// RescueGas is a paid mutator transaction binding the contract method 0x40432d51.
//
// Solidity: function rescueGas() returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) RescueGas() (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.RescueGas(&_SynapseCCTPFees.TransactOpts)
}

// RescueGas is a paid mutator transaction binding the contract method 0x40432d51.
//
// Solidity: function rescueGas() returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) RescueGas() (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.RescueGas(&_SynapseCCTPFees.TransactOpts)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) SetChainGasAmount(opts *bind.TransactOpts, newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "setChainGasAmount", newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SetChainGasAmount(&_SynapseCCTPFees.TransactOpts, newChainGasAmount)
}

// SetChainGasAmount is a paid mutator transaction binding the contract method 0xb250fe6b.
//
// Solidity: function setChainGasAmount(uint256 newChainGasAmount) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) SetChainGasAmount(newChainGasAmount *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SetChainGasAmount(&_SynapseCCTPFees.TransactOpts, newChainGasAmount)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) SetFeeCollector(opts *bind.TransactOpts, feeCollector common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "setFeeCollector", feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SetFeeCollector(&_SynapseCCTPFees.TransactOpts, feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) SetFeeCollector(feeCollector common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SetFeeCollector(&_SynapseCCTPFees.TransactOpts, feeCollector)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newProtocolFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) SetProtocolFee(opts *bind.TransactOpts, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "setProtocolFee", newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newProtocolFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) SetProtocolFee(newProtocolFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SetProtocolFee(&_SynapseCCTPFees.TransactOpts, newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newProtocolFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) SetProtocolFee(newProtocolFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SetProtocolFee(&_SynapseCCTPFees.TransactOpts, newProtocolFee)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x4bdb4eed.
//
// Solidity: function setTokenFee(address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) SetTokenFee(opts *bind.TransactOpts, token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "setTokenFee", token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x4bdb4eed.
//
// Solidity: function setTokenFee(address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) SetTokenFee(token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SetTokenFee(&_SynapseCCTPFees.TransactOpts, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x4bdb4eed.
//
// Solidity: function setTokenFee(address token, uint256 relayerFee, uint256 minBaseFee, uint256 minSwapFee, uint256 maxFee) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) SetTokenFee(token common.Address, relayerFee *big.Int, minBaseFee *big.Int, minSwapFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.SetTokenFee(&_SynapseCCTPFees.TransactOpts, token, relayerFee, minBaseFee, minSwapFee, maxFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.TransferOwnership(&_SynapseCCTPFees.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseCCTPFees *SynapseCCTPFeesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseCCTPFees.Contract.TransferOwnership(&_SynapseCCTPFees.TransactOpts, newOwner)
}

// SynapseCCTPFeesChainGasAirdroppedIterator is returned from FilterChainGasAirdropped and is used to iterate over the raw logs and unpacked data for ChainGasAirdropped events raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesChainGasAirdroppedIterator struct {
	Event *SynapseCCTPFeesChainGasAirdropped // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesChainGasAirdroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesChainGasAirdropped)
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
		it.Event = new(SynapseCCTPFeesChainGasAirdropped)
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
func (it *SynapseCCTPFeesChainGasAirdroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesChainGasAirdroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesChainGasAirdropped represents a ChainGasAirdropped event raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesChainGasAirdropped struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChainGasAirdropped is a free log retrieval operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) FilterChainGasAirdropped(opts *bind.FilterOpts) (*SynapseCCTPFeesChainGasAirdroppedIterator, error) {

	logs, sub, err := _SynapseCCTPFees.contract.FilterLogs(opts, "ChainGasAirdropped")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesChainGasAirdroppedIterator{contract: _SynapseCCTPFees.contract, event: "ChainGasAirdropped", logs: logs, sub: sub}, nil
}

// WatchChainGasAirdropped is a free log subscription operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) WatchChainGasAirdropped(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesChainGasAirdropped) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTPFees.contract.WatchLogs(opts, "ChainGasAirdropped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesChainGasAirdropped)
				if err := _SynapseCCTPFees.contract.UnpackLog(event, "ChainGasAirdropped", log); err != nil {
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

// ParseChainGasAirdropped is a log parse operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) ParseChainGasAirdropped(log types.Log) (*SynapseCCTPFeesChainGasAirdropped, error) {
	event := new(SynapseCCTPFeesChainGasAirdropped)
	if err := _SynapseCCTPFees.contract.UnpackLog(event, "ChainGasAirdropped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesChainGasAmountUpdatedIterator is returned from FilterChainGasAmountUpdated and is used to iterate over the raw logs and unpacked data for ChainGasAmountUpdated events raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesChainGasAmountUpdatedIterator struct {
	Event *SynapseCCTPFeesChainGasAmountUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesChainGasAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesChainGasAmountUpdated)
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
		it.Event = new(SynapseCCTPFeesChainGasAmountUpdated)
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
func (it *SynapseCCTPFeesChainGasAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesChainGasAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesChainGasAmountUpdated represents a ChainGasAmountUpdated event raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesChainGasAmountUpdated struct {
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChainGasAmountUpdated is a free log retrieval operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) FilterChainGasAmountUpdated(opts *bind.FilterOpts) (*SynapseCCTPFeesChainGasAmountUpdatedIterator, error) {

	logs, sub, err := _SynapseCCTPFees.contract.FilterLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesChainGasAmountUpdatedIterator{contract: _SynapseCCTPFees.contract, event: "ChainGasAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchChainGasAmountUpdated is a free log subscription operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) WatchChainGasAmountUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesChainGasAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTPFees.contract.WatchLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesChainGasAmountUpdated)
				if err := _SynapseCCTPFees.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
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

// ParseChainGasAmountUpdated is a log parse operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) ParseChainGasAmountUpdated(log types.Log) (*SynapseCCTPFeesChainGasAmountUpdated, error) {
	event := new(SynapseCCTPFeesChainGasAmountUpdated)
	if err := _SynapseCCTPFees.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesFeeCollectedIterator struct {
	Event *SynapseCCTPFeesFeeCollected // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesFeeCollected)
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
		it.Event = new(SynapseCCTPFeesFeeCollected)
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
func (it *SynapseCCTPFeesFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesFeeCollected represents a FeeCollected event raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesFeeCollected struct {
	FeeCollector      common.Address
	RelayerFeeAmount  *big.Int
	ProtocolFeeAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFeeCollected is a free log retrieval operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*SynapseCCTPFeesFeeCollectedIterator, error) {

	logs, sub, err := _SynapseCCTPFees.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesFeeCollectedIterator{contract: _SynapseCCTPFees.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesFeeCollected) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTPFees.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesFeeCollected)
				if err := _SynapseCCTPFees.contract.UnpackLog(event, "FeeCollected", log); err != nil {
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

// ParseFeeCollected is a log parse operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) ParseFeeCollected(log types.Log) (*SynapseCCTPFeesFeeCollected, error) {
	event := new(SynapseCCTPFeesFeeCollected)
	if err := _SynapseCCTPFees.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesFeeCollectorUpdatedIterator is returned from FilterFeeCollectorUpdated and is used to iterate over the raw logs and unpacked data for FeeCollectorUpdated events raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesFeeCollectorUpdatedIterator struct {
	Event *SynapseCCTPFeesFeeCollectorUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesFeeCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesFeeCollectorUpdated)
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
		it.Event = new(SynapseCCTPFeesFeeCollectorUpdated)
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
func (it *SynapseCCTPFeesFeeCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesFeeCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesFeeCollectorUpdated represents a FeeCollectorUpdated event raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesFeeCollectorUpdated struct {
	Relayer         common.Address
	OldFeeCollector common.Address
	NewFeeCollector common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorUpdated is a free log retrieval operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) FilterFeeCollectorUpdated(opts *bind.FilterOpts, relayer []common.Address) (*SynapseCCTPFeesFeeCollectorUpdatedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _SynapseCCTPFees.contract.FilterLogs(opts, "FeeCollectorUpdated", relayerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesFeeCollectorUpdatedIterator{contract: _SynapseCCTPFees.contract, event: "FeeCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorUpdated is a free log subscription operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) WatchFeeCollectorUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesFeeCollectorUpdated, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _SynapseCCTPFees.contract.WatchLogs(opts, "FeeCollectorUpdated", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesFeeCollectorUpdated)
				if err := _SynapseCCTPFees.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
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

// ParseFeeCollectorUpdated is a log parse operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) ParseFeeCollectorUpdated(log types.Log) (*SynapseCCTPFeesFeeCollectorUpdated, error) {
	event := new(SynapseCCTPFeesFeeCollectorUpdated)
	if err := _SynapseCCTPFees.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesOwnershipTransferredIterator struct {
	Event *SynapseCCTPFeesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesOwnershipTransferred)
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
		it.Event = new(SynapseCCTPFeesOwnershipTransferred)
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
func (it *SynapseCCTPFeesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesOwnershipTransferred represents a OwnershipTransferred event raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynapseCCTPFeesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseCCTPFees.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesOwnershipTransferredIterator{contract: _SynapseCCTPFees.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseCCTPFees.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesOwnershipTransferred)
				if err := _SynapseCCTPFees.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) ParseOwnershipTransferred(log types.Log) (*SynapseCCTPFeesOwnershipTransferred, error) {
	event := new(SynapseCCTPFeesOwnershipTransferred)
	if err := _SynapseCCTPFees.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesProtocolFeeUpdatedIterator is returned from FilterProtocolFeeUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeUpdated events raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesProtocolFeeUpdatedIterator struct {
	Event *SynapseCCTPFeesProtocolFeeUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesProtocolFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesProtocolFeeUpdated)
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
		it.Event = new(SynapseCCTPFeesProtocolFeeUpdated)
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
func (it *SynapseCCTPFeesProtocolFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesProtocolFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesProtocolFeeUpdated represents a ProtocolFeeUpdated event raised by the SynapseCCTPFees contract.
type SynapseCCTPFeesProtocolFeeUpdated struct {
	NewProtocolFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeUpdated is a free log retrieval operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) FilterProtocolFeeUpdated(opts *bind.FilterOpts) (*SynapseCCTPFeesProtocolFeeUpdatedIterator, error) {

	logs, sub, err := _SynapseCCTPFees.contract.FilterLogs(opts, "ProtocolFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesProtocolFeeUpdatedIterator{contract: _SynapseCCTPFees.contract, event: "ProtocolFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeUpdated is a free log subscription operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) WatchProtocolFeeUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesProtocolFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTPFees.contract.WatchLogs(opts, "ProtocolFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesProtocolFeeUpdated)
				if err := _SynapseCCTPFees.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
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

// ParseProtocolFeeUpdated is a log parse operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTPFees *SynapseCCTPFeesFilterer) ParseProtocolFeeUpdated(log types.Log) (*SynapseCCTPFeesProtocolFeeUpdated, error) {
	event := new(SynapseCCTPFeesProtocolFeeUpdated)
	if err := _SynapseCCTPFees.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesEventsMetaData contains all meta data concerning the SynapseCCTPFeesEvents contract.
var SynapseCCTPFeesEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAirdropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayerFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeUpdated\",\"type\":\"event\"}]",
}

// SynapseCCTPFeesEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseCCTPFeesEventsMetaData.ABI instead.
var SynapseCCTPFeesEventsABI = SynapseCCTPFeesEventsMetaData.ABI

// SynapseCCTPFeesEvents is an auto generated Go binding around an Ethereum contract.
type SynapseCCTPFeesEvents struct {
	SynapseCCTPFeesEventsCaller     // Read-only binding to the contract
	SynapseCCTPFeesEventsTransactor // Write-only binding to the contract
	SynapseCCTPFeesEventsFilterer   // Log filterer for contract events
}

// SynapseCCTPFeesEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseCCTPFeesEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPFeesEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseCCTPFeesEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPFeesEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseCCTPFeesEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseCCTPFeesEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseCCTPFeesEventsSession struct {
	Contract     *SynapseCCTPFeesEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SynapseCCTPFeesEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseCCTPFeesEventsCallerSession struct {
	Contract *SynapseCCTPFeesEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SynapseCCTPFeesEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseCCTPFeesEventsTransactorSession struct {
	Contract     *SynapseCCTPFeesEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SynapseCCTPFeesEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseCCTPFeesEventsRaw struct {
	Contract *SynapseCCTPFeesEvents // Generic contract binding to access the raw methods on
}

// SynapseCCTPFeesEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseCCTPFeesEventsCallerRaw struct {
	Contract *SynapseCCTPFeesEventsCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseCCTPFeesEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseCCTPFeesEventsTransactorRaw struct {
	Contract *SynapseCCTPFeesEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseCCTPFeesEvents creates a new instance of SynapseCCTPFeesEvents, bound to a specific deployed contract.
func NewSynapseCCTPFeesEvents(address common.Address, backend bind.ContractBackend) (*SynapseCCTPFeesEvents, error) {
	contract, err := bindSynapseCCTPFeesEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEvents{SynapseCCTPFeesEventsCaller: SynapseCCTPFeesEventsCaller{contract: contract}, SynapseCCTPFeesEventsTransactor: SynapseCCTPFeesEventsTransactor{contract: contract}, SynapseCCTPFeesEventsFilterer: SynapseCCTPFeesEventsFilterer{contract: contract}}, nil
}

// NewSynapseCCTPFeesEventsCaller creates a new read-only instance of SynapseCCTPFeesEvents, bound to a specific deployed contract.
func NewSynapseCCTPFeesEventsCaller(address common.Address, caller bind.ContractCaller) (*SynapseCCTPFeesEventsCaller, error) {
	contract, err := bindSynapseCCTPFeesEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEventsCaller{contract: contract}, nil
}

// NewSynapseCCTPFeesEventsTransactor creates a new write-only instance of SynapseCCTPFeesEvents, bound to a specific deployed contract.
func NewSynapseCCTPFeesEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseCCTPFeesEventsTransactor, error) {
	contract, err := bindSynapseCCTPFeesEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEventsTransactor{contract: contract}, nil
}

// NewSynapseCCTPFeesEventsFilterer creates a new log filterer instance of SynapseCCTPFeesEvents, bound to a specific deployed contract.
func NewSynapseCCTPFeesEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseCCTPFeesEventsFilterer, error) {
	contract, err := bindSynapseCCTPFeesEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEventsFilterer{contract: contract}, nil
}

// bindSynapseCCTPFeesEvents binds a generic wrapper to an already deployed contract.
func bindSynapseCCTPFeesEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseCCTPFeesEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTPFeesEvents.Contract.SynapseCCTPFeesEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPFeesEvents.Contract.SynapseCCTPFeesEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTPFeesEvents.Contract.SynapseCCTPFeesEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseCCTPFeesEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTPFeesEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseCCTPFeesEvents.Contract.contract.Transact(opts, method, params...)
}

// SynapseCCTPFeesEventsChainGasAirdroppedIterator is returned from FilterChainGasAirdropped and is used to iterate over the raw logs and unpacked data for ChainGasAirdropped events raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsChainGasAirdroppedIterator struct {
	Event *SynapseCCTPFeesEventsChainGasAirdropped // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesEventsChainGasAirdroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesEventsChainGasAirdropped)
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
		it.Event = new(SynapseCCTPFeesEventsChainGasAirdropped)
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
func (it *SynapseCCTPFeesEventsChainGasAirdroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesEventsChainGasAirdroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesEventsChainGasAirdropped represents a ChainGasAirdropped event raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsChainGasAirdropped struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChainGasAirdropped is a free log retrieval operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) FilterChainGasAirdropped(opts *bind.FilterOpts) (*SynapseCCTPFeesEventsChainGasAirdroppedIterator, error) {

	logs, sub, err := _SynapseCCTPFeesEvents.contract.FilterLogs(opts, "ChainGasAirdropped")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEventsChainGasAirdroppedIterator{contract: _SynapseCCTPFeesEvents.contract, event: "ChainGasAirdropped", logs: logs, sub: sub}, nil
}

// WatchChainGasAirdropped is a free log subscription operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) WatchChainGasAirdropped(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesEventsChainGasAirdropped) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTPFeesEvents.contract.WatchLogs(opts, "ChainGasAirdropped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesEventsChainGasAirdropped)
				if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "ChainGasAirdropped", log); err != nil {
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

// ParseChainGasAirdropped is a log parse operation binding the contract event 0xf9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd683.
//
// Solidity: event ChainGasAirdropped(uint256 amount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) ParseChainGasAirdropped(log types.Log) (*SynapseCCTPFeesEventsChainGasAirdropped, error) {
	event := new(SynapseCCTPFeesEventsChainGasAirdropped)
	if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "ChainGasAirdropped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesEventsChainGasAmountUpdatedIterator is returned from FilterChainGasAmountUpdated and is used to iterate over the raw logs and unpacked data for ChainGasAmountUpdated events raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsChainGasAmountUpdatedIterator struct {
	Event *SynapseCCTPFeesEventsChainGasAmountUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesEventsChainGasAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesEventsChainGasAmountUpdated)
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
		it.Event = new(SynapseCCTPFeesEventsChainGasAmountUpdated)
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
func (it *SynapseCCTPFeesEventsChainGasAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesEventsChainGasAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesEventsChainGasAmountUpdated represents a ChainGasAmountUpdated event raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsChainGasAmountUpdated struct {
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChainGasAmountUpdated is a free log retrieval operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) FilterChainGasAmountUpdated(opts *bind.FilterOpts) (*SynapseCCTPFeesEventsChainGasAmountUpdatedIterator, error) {

	logs, sub, err := _SynapseCCTPFeesEvents.contract.FilterLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEventsChainGasAmountUpdatedIterator{contract: _SynapseCCTPFeesEvents.contract, event: "ChainGasAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchChainGasAmountUpdated is a free log subscription operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) WatchChainGasAmountUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesEventsChainGasAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTPFeesEvents.contract.WatchLogs(opts, "ChainGasAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesEventsChainGasAmountUpdated)
				if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
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

// ParseChainGasAmountUpdated is a log parse operation binding the contract event 0x5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b4.
//
// Solidity: event ChainGasAmountUpdated(uint256 chainGasAmount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) ParseChainGasAmountUpdated(log types.Log) (*SynapseCCTPFeesEventsChainGasAmountUpdated, error) {
	event := new(SynapseCCTPFeesEventsChainGasAmountUpdated)
	if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "ChainGasAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesEventsFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsFeeCollectedIterator struct {
	Event *SynapseCCTPFeesEventsFeeCollected // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesEventsFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesEventsFeeCollected)
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
		it.Event = new(SynapseCCTPFeesEventsFeeCollected)
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
func (it *SynapseCCTPFeesEventsFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesEventsFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesEventsFeeCollected represents a FeeCollected event raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsFeeCollected struct {
	FeeCollector      common.Address
	RelayerFeeAmount  *big.Int
	ProtocolFeeAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFeeCollected is a free log retrieval operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*SynapseCCTPFeesEventsFeeCollectedIterator, error) {

	logs, sub, err := _SynapseCCTPFeesEvents.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEventsFeeCollectedIterator{contract: _SynapseCCTPFeesEvents.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesEventsFeeCollected) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTPFeesEvents.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesEventsFeeCollected)
				if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "FeeCollected", log); err != nil {
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

// ParseFeeCollected is a log parse operation binding the contract event 0x108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca561.
//
// Solidity: event FeeCollected(address feeCollector, uint256 relayerFeeAmount, uint256 protocolFeeAmount)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) ParseFeeCollected(log types.Log) (*SynapseCCTPFeesEventsFeeCollected, error) {
	event := new(SynapseCCTPFeesEventsFeeCollected)
	if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesEventsFeeCollectorUpdatedIterator is returned from FilterFeeCollectorUpdated and is used to iterate over the raw logs and unpacked data for FeeCollectorUpdated events raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsFeeCollectorUpdatedIterator struct {
	Event *SynapseCCTPFeesEventsFeeCollectorUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesEventsFeeCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesEventsFeeCollectorUpdated)
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
		it.Event = new(SynapseCCTPFeesEventsFeeCollectorUpdated)
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
func (it *SynapseCCTPFeesEventsFeeCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesEventsFeeCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesEventsFeeCollectorUpdated represents a FeeCollectorUpdated event raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsFeeCollectorUpdated struct {
	Relayer         common.Address
	OldFeeCollector common.Address
	NewFeeCollector common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorUpdated is a free log retrieval operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) FilterFeeCollectorUpdated(opts *bind.FilterOpts, relayer []common.Address) (*SynapseCCTPFeesEventsFeeCollectorUpdatedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _SynapseCCTPFeesEvents.contract.FilterLogs(opts, "FeeCollectorUpdated", relayerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEventsFeeCollectorUpdatedIterator{contract: _SynapseCCTPFeesEvents.contract, event: "FeeCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorUpdated is a free log subscription operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) WatchFeeCollectorUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesEventsFeeCollectorUpdated, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _SynapseCCTPFeesEvents.contract.WatchLogs(opts, "FeeCollectorUpdated", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesEventsFeeCollectorUpdated)
				if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
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

// ParseFeeCollectorUpdated is a log parse operation binding the contract event 0x9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e.
//
// Solidity: event FeeCollectorUpdated(address indexed relayer, address oldFeeCollector, address newFeeCollector)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) ParseFeeCollectorUpdated(log types.Log) (*SynapseCCTPFeesEventsFeeCollectorUpdated, error) {
	event := new(SynapseCCTPFeesEventsFeeCollectorUpdated)
	if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPFeesEventsProtocolFeeUpdatedIterator is returned from FilterProtocolFeeUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeUpdated events raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsProtocolFeeUpdatedIterator struct {
	Event *SynapseCCTPFeesEventsProtocolFeeUpdated // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPFeesEventsProtocolFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPFeesEventsProtocolFeeUpdated)
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
		it.Event = new(SynapseCCTPFeesEventsProtocolFeeUpdated)
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
func (it *SynapseCCTPFeesEventsProtocolFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPFeesEventsProtocolFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPFeesEventsProtocolFeeUpdated represents a ProtocolFeeUpdated event raised by the SynapseCCTPFeesEvents contract.
type SynapseCCTPFeesEventsProtocolFeeUpdated struct {
	NewProtocolFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeUpdated is a free log retrieval operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) FilterProtocolFeeUpdated(opts *bind.FilterOpts) (*SynapseCCTPFeesEventsProtocolFeeUpdatedIterator, error) {

	logs, sub, err := _SynapseCCTPFeesEvents.contract.FilterLogs(opts, "ProtocolFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPFeesEventsProtocolFeeUpdatedIterator{contract: _SynapseCCTPFeesEvents.contract, event: "ProtocolFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeUpdated is a free log subscription operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) WatchProtocolFeeUpdated(opts *bind.WatchOpts, sink chan<- *SynapseCCTPFeesEventsProtocolFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTPFeesEvents.contract.WatchLogs(opts, "ProtocolFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPFeesEventsProtocolFeeUpdated)
				if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
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

// ParseProtocolFeeUpdated is a log parse operation binding the contract event 0xd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40.
//
// Solidity: event ProtocolFeeUpdated(uint256 newProtocolFee)
func (_SynapseCCTPFeesEvents *SynapseCCTPFeesEventsFilterer) ParseProtocolFeeUpdated(log types.Log) (*SynapseCCTPFeesEventsProtocolFeeUpdated, error) {
	event := new(SynapseCCTPFeesEventsProtocolFeeUpdated)
	if err := _SynapseCCTPFeesEvents.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122041df85d2e8cef8badb468c4c896bcfabaac8cd155498513e1d5e0d66cac4b1c764736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(TypeCastsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
