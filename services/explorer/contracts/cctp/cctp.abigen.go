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

// ActionLibMetaData contains all meta data concerning the ActionLib contract.
var ActionLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c1b14cecff239984d0910096ae2c879d0b6aa72288ea9d4d1cbfffa7a2d2e6ef64736f6c634300080d0033",
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
	parsed, err := abi.JSON(strings.NewReader(ActionLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220655c6d0fdf779c53ad8bbc9d91838f1d883314cce982482dabe34d66730ec6af64736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f1e69d77df4221ee9f4a6d53f3246fc1c7af269baad02b7bcdcafb65f8b5382b64736f6c634300080d0033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"circleTokenPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"swap\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"}],\"name\":\"receiveCircleToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"swapParams\",\"type\":\"bytes\"}],\"name\":\"sendCircleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a4b1d034": "circleTokenPool(address)",
		"4a5ae51d": "receiveCircleToken(bytes,bytes,uint32,bytes)",
		"304ddb4c": "sendCircleToken(address,uint256,address,uint256,uint32,bytes)",
		"46117830": "tokenMessenger()",
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

// CircleTokenPool is a free data retrieval call binding the contract method 0xa4b1d034.
//
// Solidity: function circleTokenPool(address token) view returns(address swap)
func (_ISynapseCCTP *ISynapseCCTPCaller) CircleTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _ISynapseCCTP.contract.Call(opts, &out, "circleTokenPool", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CircleTokenPool is a free data retrieval call binding the contract method 0xa4b1d034.
//
// Solidity: function circleTokenPool(address token) view returns(address swap)
func (_ISynapseCCTP *ISynapseCCTPSession) CircleTokenPool(token common.Address) (common.Address, error) {
	return _ISynapseCCTP.Contract.CircleTokenPool(&_ISynapseCCTP.CallOpts, token)
}

// CircleTokenPool is a free data retrieval call binding the contract method 0xa4b1d034.
//
// Solidity: function circleTokenPool(address token) view returns(address swap)
func (_ISynapseCCTP *ISynapseCCTPCallerSession) CircleTokenPool(token common.Address) (common.Address, error) {
	return _ISynapseCCTP.Contract.CircleTokenPool(&_ISynapseCCTP.CallOpts, token)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_ISynapseCCTP *ISynapseCCTPCaller) TokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISynapseCCTP.contract.Call(opts, &out, "tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_ISynapseCCTP *ISynapseCCTPSession) TokenMessenger() (common.Address, error) {
	return _ISynapseCCTP.Contract.TokenMessenger(&_ISynapseCCTP.CallOpts)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_ISynapseCCTP *ISynapseCCTPCallerSession) TokenMessenger() (common.Address, error) {
	return _ISynapseCCTP.Contract.TokenMessenger(&_ISynapseCCTP.CallOpts)
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

// ISynapseCCTPFeesMetaData contains all meta data concerning the ISynapseCCTPFees contract.
var ISynapseCCTPFeesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSwap\",\"type\":\"bool\"}],\"name\":\"calculateFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"feeStructures\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"relayerFee\",\"type\":\"uint40\"},{\"internalType\":\"uint72\",\"name\":\"minBaseFee\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"minSwapFee\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"maxFee\",\"type\":\"uint72\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"bridgeTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenToSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0d25aafe": "calculateFeeAmount(address,uint256,bool)",
		"dc72495b": "feeStructures(address)",
		"9c1d060e": "getBridgeTokens()",
		"a5bc29c2": "symbolToToken(string)",
		"0ba36121": "tokenToSymbol(address)",
	},
}

// ISynapseCCTPFeesABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseCCTPFeesMetaData.ABI instead.
var ISynapseCCTPFeesABI = ISynapseCCTPFeesMetaData.ABI

// Deprecated: Use ISynapseCCTPFeesMetaData.Sigs instead.
// ISynapseCCTPFeesFuncSigs maps the 4-byte function signature to its string representation.
var ISynapseCCTPFeesFuncSigs = ISynapseCCTPFeesMetaData.Sigs

// ISynapseCCTPFees is an auto generated Go binding around an Ethereum contract.
type ISynapseCCTPFees struct {
	ISynapseCCTPFeesCaller     // Read-only binding to the contract
	ISynapseCCTPFeesTransactor // Write-only binding to the contract
	ISynapseCCTPFeesFilterer   // Log filterer for contract events
}

// ISynapseCCTPFeesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseCCTPFeesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPFeesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseCCTPFeesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPFeesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseCCTPFeesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseCCTPFeesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseCCTPFeesSession struct {
	Contract     *ISynapseCCTPFees // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynapseCCTPFeesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseCCTPFeesCallerSession struct {
	Contract *ISynapseCCTPFeesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ISynapseCCTPFeesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseCCTPFeesTransactorSession struct {
	Contract     *ISynapseCCTPFeesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ISynapseCCTPFeesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseCCTPFeesRaw struct {
	Contract *ISynapseCCTPFees // Generic contract binding to access the raw methods on
}

// ISynapseCCTPFeesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseCCTPFeesCallerRaw struct {
	Contract *ISynapseCCTPFeesCaller // Generic read-only contract binding to access the raw methods on
}

// ISynapseCCTPFeesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseCCTPFeesTransactorRaw struct {
	Contract *ISynapseCCTPFeesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseCCTPFees creates a new instance of ISynapseCCTPFees, bound to a specific deployed contract.
func NewISynapseCCTPFees(address common.Address, backend bind.ContractBackend) (*ISynapseCCTPFees, error) {
	contract, err := bindISynapseCCTPFees(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPFees{ISynapseCCTPFeesCaller: ISynapseCCTPFeesCaller{contract: contract}, ISynapseCCTPFeesTransactor: ISynapseCCTPFeesTransactor{contract: contract}, ISynapseCCTPFeesFilterer: ISynapseCCTPFeesFilterer{contract: contract}}, nil
}

// NewISynapseCCTPFeesCaller creates a new read-only instance of ISynapseCCTPFees, bound to a specific deployed contract.
func NewISynapseCCTPFeesCaller(address common.Address, caller bind.ContractCaller) (*ISynapseCCTPFeesCaller, error) {
	contract, err := bindISynapseCCTPFees(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPFeesCaller{contract: contract}, nil
}

// NewISynapseCCTPFeesTransactor creates a new write-only instance of ISynapseCCTPFees, bound to a specific deployed contract.
func NewISynapseCCTPFeesTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseCCTPFeesTransactor, error) {
	contract, err := bindISynapseCCTPFees(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPFeesTransactor{contract: contract}, nil
}

// NewISynapseCCTPFeesFilterer creates a new log filterer instance of ISynapseCCTPFees, bound to a specific deployed contract.
func NewISynapseCCTPFeesFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseCCTPFeesFilterer, error) {
	contract, err := bindISynapseCCTPFees(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseCCTPFeesFilterer{contract: contract}, nil
}

// bindISynapseCCTPFees binds a generic wrapper to an already deployed contract.
func bindISynapseCCTPFees(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISynapseCCTPFeesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseCCTPFees *ISynapseCCTPFeesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseCCTPFees.Contract.ISynapseCCTPFeesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseCCTPFees *ISynapseCCTPFeesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseCCTPFees.Contract.ISynapseCCTPFeesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseCCTPFees *ISynapseCCTPFeesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseCCTPFees.Contract.ISynapseCCTPFeesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseCCTPFees *ISynapseCCTPFeesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseCCTPFees.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseCCTPFees *ISynapseCCTPFeesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseCCTPFees.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseCCTPFees *ISynapseCCTPFeesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseCCTPFees.Contract.contract.Transact(opts, method, params...)
}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCaller) CalculateFeeAmount(opts *bind.CallOpts, token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseCCTPFees.contract.Call(opts, &out, "calculateFeeAmount", token, amount, isSwap)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_ISynapseCCTPFees *ISynapseCCTPFeesSession) CalculateFeeAmount(token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	return _ISynapseCCTPFees.Contract.CalculateFeeAmount(&_ISynapseCCTPFees.CallOpts, token, amount, isSwap)
}

// CalculateFeeAmount is a free data retrieval call binding the contract method 0x0d25aafe.
//
// Solidity: function calculateFeeAmount(address token, uint256 amount, bool isSwap) view returns(uint256 fee)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCallerSession) CalculateFeeAmount(token common.Address, amount *big.Int, isSwap bool) (*big.Int, error) {
	return _ISynapseCCTPFees.Contract.CalculateFeeAmount(&_ISynapseCCTPFees.CallOpts, token, amount, isSwap)
}

// FeeStructures is a free data retrieval call binding the contract method 0xdc72495b.
//
// Solidity: function feeStructures(address token) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCaller) FeeStructures(opts *bind.CallOpts, token common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	var out []interface{}
	err := _ISynapseCCTPFees.contract.Call(opts, &out, "feeStructures", token)

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
// Solidity: function feeStructures(address token) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_ISynapseCCTPFees *ISynapseCCTPFeesSession) FeeStructures(token common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	return _ISynapseCCTPFees.Contract.FeeStructures(&_ISynapseCCTPFees.CallOpts, token)
}

// FeeStructures is a free data retrieval call binding the contract method 0xdc72495b.
//
// Solidity: function feeStructures(address token) view returns(uint40 relayerFee, uint72 minBaseFee, uint72 minSwapFee, uint72 maxFee)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCallerSession) FeeStructures(token common.Address) (struct {
	RelayerFee *big.Int
	MinBaseFee *big.Int
	MinSwapFee *big.Int
	MaxFee     *big.Int
}, error) {
	return _ISynapseCCTPFees.Contract.FeeStructures(&_ISynapseCCTPFees.CallOpts, token)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCaller) GetBridgeTokens(opts *bind.CallOpts) ([]BridgeToken, error) {
	var out []interface{}
	err := _ISynapseCCTPFees.contract.Call(opts, &out, "getBridgeTokens")

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_ISynapseCCTPFees *ISynapseCCTPFeesSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _ISynapseCCTPFees.Contract.GetBridgeTokens(&_ISynapseCCTPFees.CallOpts)
}

// GetBridgeTokens is a free data retrieval call binding the contract method 0x9c1d060e.
//
// Solidity: function getBridgeTokens() view returns((string,address)[] bridgeTokens)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCallerSession) GetBridgeTokens() ([]BridgeToken, error) {
	return _ISynapseCCTPFees.Contract.GetBridgeTokens(&_ISynapseCCTPFees.CallOpts)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string symbol) view returns(address token)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCaller) SymbolToToken(opts *bind.CallOpts, symbol string) (common.Address, error) {
	var out []interface{}
	err := _ISynapseCCTPFees.contract.Call(opts, &out, "symbolToToken", symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string symbol) view returns(address token)
func (_ISynapseCCTPFees *ISynapseCCTPFeesSession) SymbolToToken(symbol string) (common.Address, error) {
	return _ISynapseCCTPFees.Contract.SymbolToToken(&_ISynapseCCTPFees.CallOpts, symbol)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string symbol) view returns(address token)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCallerSession) SymbolToToken(symbol string) (common.Address, error) {
	return _ISynapseCCTPFees.Contract.SymbolToToken(&_ISynapseCCTPFees.CallOpts, symbol)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address token) view returns(string symbol)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCaller) TokenToSymbol(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _ISynapseCCTPFees.contract.Call(opts, &out, "tokenToSymbol", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address token) view returns(string symbol)
func (_ISynapseCCTPFees *ISynapseCCTPFeesSession) TokenToSymbol(token common.Address) (string, error) {
	return _ISynapseCCTPFees.Contract.TokenToSymbol(&_ISynapseCCTPFees.CallOpts, token)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address token) view returns(string symbol)
func (_ISynapseCCTPFees *ISynapseCCTPFeesCallerSession) TokenToSymbol(token common.Address) (string, error) {
	return _ISynapseCCTPFees.Contract.TokenToSymbol(&_ISynapseCCTPFees.CallOpts, token)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"burnLimitsPerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"remoteToken\",\"type\":\"bytes32\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"burnToken\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9dc29fac": "burn(address,uint256)",
		"a56ec632": "burnLimitsPerMessage(address)",
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

// BurnLimitsPerMessage is a free data retrieval call binding the contract method 0xa56ec632.
//
// Solidity: function burnLimitsPerMessage(address token) view returns(uint256)
func (_ITokenMinter *ITokenMinterCaller) BurnLimitsPerMessage(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ITokenMinter.contract.Call(opts, &out, "burnLimitsPerMessage", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnLimitsPerMessage is a free data retrieval call binding the contract method 0xa56ec632.
//
// Solidity: function burnLimitsPerMessage(address token) view returns(uint256)
func (_ITokenMinter *ITokenMinterSession) BurnLimitsPerMessage(token common.Address) (*big.Int, error) {
	return _ITokenMinter.Contract.BurnLimitsPerMessage(&_ITokenMinter.CallOpts, token)
}

// BurnLimitsPerMessage is a free data retrieval call binding the contract method 0xa56ec632.
//
// Solidity: function burnLimitsPerMessage(address token) view returns(uint256)
func (_ITokenMinter *ITokenMinterCallerSession) BurnLimitsPerMessage(token common.Address) (*big.Int, error) {
	return _ITokenMinter.Contract.BurnLimitsPerMessage(&_ITokenMinter.CallOpts, token)
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

// MessageTransmitterMetaData contains all meta data concerning the MessageTransmitter contract.
var MessageTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMessageTransmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvailableNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"2c121921": "localMessageTransmitter()",
		"8371744e": "nextAvailableNonce()",
		"57ecfd28": "receiveMessage(bytes,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610249806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80638371744e116100505780638371744e146100a75780638d3638f4146100c3578063f7259a75146100d257600080fd5b80632c1219211461006c57806357ecfd281461007f575b600080fd5b6040513081526020015b60405180910390f35b61009761008d366004610134565b6001949350505050565b6040519015158152602001610076565b60015b60405167ffffffffffffffff9091168152602001610076565b60405160018152602001610076565b6100aa6100e03660046101a0565b600195945050505050565b60008083601f8401126100fd57600080fd5b50813567ffffffffffffffff81111561011557600080fd5b60208301915083602082850101111561012d57600080fd5b9250929050565b6000806000806040858703121561014a57600080fd5b843567ffffffffffffffff8082111561016257600080fd5b61016e888389016100eb565b9096509450602087013591508082111561018757600080fd5b50610194878288016100eb565b95989497509550505050565b6000806000806000608086880312156101b857600080fd5b853563ffffffff811681146101cc57600080fd5b94506020860135935060408601359250606086013567ffffffffffffffff8111156101f657600080fd5b610202888289016100eb565b96999598509396509294939250505056fea26469706673582212202d056d4f7e5d33a986c7ac7e4529e3e6fe06c21d2cc169af4d552ca7e81f6b0764736f6c634300080d0033",
}

// MessageTransmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageTransmitterMetaData.ABI instead.
var MessageTransmitterABI = MessageTransmitterMetaData.ABI

// Deprecated: Use MessageTransmitterMetaData.Sigs instead.
// MessageTransmitterFuncSigs maps the 4-byte function signature to its string representation.
var MessageTransmitterFuncSigs = MessageTransmitterMetaData.Sigs

// MessageTransmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageTransmitterMetaData.Bin instead.
var MessageTransmitterBin = MessageTransmitterMetaData.Bin

// DeployMessageTransmitter deploys a new Ethereum contract, binding an instance of MessageTransmitter to it.
func DeployMessageTransmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageTransmitter, error) {
	parsed, err := MessageTransmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageTransmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageTransmitter{MessageTransmitterCaller: MessageTransmitterCaller{contract: contract}, MessageTransmitterTransactor: MessageTransmitterTransactor{contract: contract}, MessageTransmitterFilterer: MessageTransmitterFilterer{contract: contract}}, nil
}

// MessageTransmitter is an auto generated Go binding around an Ethereum contract.
type MessageTransmitter struct {
	MessageTransmitterCaller     // Read-only binding to the contract
	MessageTransmitterTransactor // Write-only binding to the contract
	MessageTransmitterFilterer   // Log filterer for contract events
}

// MessageTransmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageTransmitterSession struct {
	Contract     *MessageTransmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessageTransmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageTransmitterCallerSession struct {
	Contract *MessageTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MessageTransmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTransmitterTransactorSession struct {
	Contract     *MessageTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MessageTransmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageTransmitterRaw struct {
	Contract *MessageTransmitter // Generic contract binding to access the raw methods on
}

// MessageTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageTransmitterCallerRaw struct {
	Contract *MessageTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTransmitterTransactorRaw struct {
	Contract *MessageTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageTransmitter creates a new instance of MessageTransmitter, bound to a specific deployed contract.
func NewMessageTransmitter(address common.Address, backend bind.ContractBackend) (*MessageTransmitter, error) {
	contract, err := bindMessageTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitter{MessageTransmitterCaller: MessageTransmitterCaller{contract: contract}, MessageTransmitterTransactor: MessageTransmitterTransactor{contract: contract}, MessageTransmitterFilterer: MessageTransmitterFilterer{contract: contract}}, nil
}

// NewMessageTransmitterCaller creates a new read-only instance of MessageTransmitter, bound to a specific deployed contract.
func NewMessageTransmitterCaller(address common.Address, caller bind.ContractCaller) (*MessageTransmitterCaller, error) {
	contract, err := bindMessageTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterCaller{contract: contract}, nil
}

// NewMessageTransmitterTransactor creates a new write-only instance of MessageTransmitter, bound to a specific deployed contract.
func NewMessageTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTransmitterTransactor, error) {
	contract, err := bindMessageTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterTransactor{contract: contract}, nil
}

// NewMessageTransmitterFilterer creates a new log filterer instance of MessageTransmitter, bound to a specific deployed contract.
func NewMessageTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageTransmitterFilterer, error) {
	contract, err := bindMessageTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterFilterer{contract: contract}, nil
}

// bindMessageTransmitter binds a generic wrapper to an already deployed contract.
func bindMessageTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageTransmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTransmitter *MessageTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTransmitter.Contract.MessageTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTransmitter *MessageTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.MessageTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTransmitter *MessageTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.MessageTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTransmitter *MessageTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTransmitter *MessageTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTransmitter *MessageTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterSession) LocalDomain() (uint32, error) {
	return _MessageTransmitter.Contract.LocalDomain(&_MessageTransmitter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterCallerSession) LocalDomain() (uint32, error) {
	return _MessageTransmitter.Contract.LocalDomain(&_MessageTransmitter.CallOpts)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_MessageTransmitter *MessageTransmitterCaller) LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "localMessageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_MessageTransmitter *MessageTransmitterSession) LocalMessageTransmitter() (common.Address, error) {
	return _MessageTransmitter.Contract.LocalMessageTransmitter(&_MessageTransmitter.CallOpts)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_MessageTransmitter *MessageTransmitterCallerSession) LocalMessageTransmitter() (common.Address, error) {
	return _MessageTransmitter.Contract.LocalMessageTransmitter(&_MessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MessageTransmitter *MessageTransmitterCaller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MessageTransmitter *MessageTransmitterSession) NextAvailableNonce() (uint64, error) {
	return _MessageTransmitter.Contract.NextAvailableNonce(&_MessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MessageTransmitter *MessageTransmitterCallerSession) NextAvailableNonce() (uint64, error) {
	return _MessageTransmitter.Contract.NextAvailableNonce(&_MessageTransmitter.CallOpts)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_MessageTransmitter *MessageTransmitterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_MessageTransmitter *MessageTransmitterSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.ReceiveMessage(&_MessageTransmitter.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_MessageTransmitter *MessageTransmitterTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.ReceiveMessage(&_MessageTransmitter.TransactOpts, message, signature)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SendMessageWithCaller(&_MessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SendMessageWithCaller(&_MessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// MinimalForwarderLibMetaData contains all meta data concerning the MinimalForwarderLib contract.
var MinimalForwarderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d22902b1c05d6b3025fbd4b72fceca0666fbecd48c422cc220629906f9a9afb164736f6c634300080d0033",
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

// PausableMetaData contains all meta data concerning the Pausable contract.
var PausableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c975abb": "paused()",
	},
}

// PausableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableMetaData.ABI instead.
var PausableABI = PausableMetaData.ABI

// Deprecated: Use PausableMetaData.Sigs instead.
// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = PausableMetaData.Sigs

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209285cee3e89e6a53b82254d8bc43097a5aa351ef18d8b329cb5ad4f03b72c79d64736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122033edcb017b82f4414ecbfed038a6cf6f212fb5854cfdc0f06f5dc5353f49ca6a64736f6c634300080d0033",
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

// SwapQueryLibMetaData contains all meta data concerning the SwapQueryLib contract.
var SwapQueryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bbbeda60024ccff66bcc1e0ae109d1a5915d4973caffea3532870c1c68bd97f564736f6c634300080d0033",
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
	parsed, err := abi.JSON(strings.NewReader(SwapQueryLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// SynapseCCTPMetaData contains all meta data concerning the SynapseCCTP contract.
var SynapseCCTPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"tokenMessenger_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CCTPGasRescueFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectGasAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPIncorrectProtocolFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPInsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPMessageNotReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPSymbolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPSymbolIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPTokenAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPTokenNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CCTPZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CastOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForwarderDeploymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectRequestLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoteCCTPDeploymentNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownRequestVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAirdropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"ChainGasAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayerFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accumulatedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSwap\",\"type\":\"bool\"}],\"name\":\"calculateFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"circleTokenPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeStructures\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"relayerFee\",\"type\":\"uint40\"},{\"internalType\":\"uint72\",\"name\":\"minBaseFee\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"minSwapFee\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"maxFee\",\"type\":\"uint72\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"bridgeTokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"isRequestFulfilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageTransmitter\",\"outputs\":[{\"internalType\":\"contractIMessageTransmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseSending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"}],\"name\":\"receiveCircleToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayerFeeCollectors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remoteDomainConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"synapseCCTP\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescueGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"swapParams\",\"type\":\"bytes\"}],\"name\":\"sendCircleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChainGasAmount\",\"type\":\"uint256\"}],\"name\":\"setChainGasAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"circleToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swap\",\"type\":\"address\"}],\"name\":\"setCircleTokenPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"remoteSynapseCCTP\",\"type\":\"address\"}],\"name\":\"setRemoteDomainConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSwapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"setTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseSending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawRelayerFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
		"d77938e4": "pauseSending()",
		"5c975abb": "paused()",
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
		"e7a64a80": "unpauseSending()",
		"2d80caa5": "withdrawProtocolFees(address)",
		"f7265b3a": "withdrawRelayerFees(address)",
	},
	Bin: "0x60e06040523480156200001157600080fd5b50604051620048a5380380620048a58339810160408190526200003491620001af565b6200003f3362000146565b600a805460ff191690556001600160a01b03821660c081905260408051632c12192160e01b81529051632c121921916004808201926020929091908290030181865afa15801562000094573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000ba9190620001ee565b6001600160a01b031660a08190526040805163234d8e3d60e21b81529051638d3638f4916004808201926020929091908290030181865afa15801562000104573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200012a919062000215565b63ffffffff166080526200013e8162000146565b50506200023d565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114620001ac57600080fd5b50565b60008060408385031215620001c357600080fd5b8251620001d08162000196565b6020840151909250620001e38162000196565b809150509250929050565b6000602082840312156200020157600080fd5b81516200020e8162000196565b9392505050565b6000602082840312156200022857600080fd5b815163ffffffff811681146200020e57600080fd5b60805160a05160c051614600620002a56000396000818161037401528181610d6201528181610d9101526124f801526000818161047401528181610bb301526124780152600081816104a801528181610c3e015281816110110152611a8a01526146006000f3fe6080604052600436106102345760003560e01c80638da5cb5b11610138578063d77938e4116100b0578063e9259ab91161007f578063f2fde38b11610064578063f2fde38b146107c5578063f7265b3a146107e5578063f879a41a1461080557600080fd5b8063e9259ab91461073b578063e9bbb36d146107a557600080fd5b8063d77938e414610644578063dc72495b14610659578063e00a83e014610710578063e7a64a801461072657600080fd5b8063a4b1d03411610107578063b0e21e8a116100ec578063b0e21e8a146105d6578063b250fe6b146105ec578063d4a67c6d1461060c57600080fd5b8063a4b1d0341461055f578063a5bc29c21461059557600080fd5b80638da5cb5b146104df57806392a442ea146104fd5780639c1d060e1461051d578063a42dce801461053f57600080fd5b80634a5ae51d116101cb5780635fa7b5841161019a578063787dce3d1161017f578063787dce3d146104425780637b04c181146104625780638d3638f41461049657600080fd5b80635fa7b5841461040d578063715018a61461042d57600080fd5b80634a5ae51d146103965780634a85178d146103a95780634bdb4eed146103c95780635c975abb146103e957600080fd5b8063304ddb4c11610207578063304ddb4c146102df57806340432d51146102ff57806341f355ee14610314578063461178301461036257600080fd5b80630ba36121146102395780630d25aafe1461026f5780632cc9e7e51461029d5780632d80caa5146102bf575b600080fd5b34801561024557600080fd5b50610259610254366004613b34565b610825565b6040516102669190613ba9565b60405180910390f35b34801561027b57600080fd5b5061028f61028a366004613bca565b6108bf565b604051908152602001610266565b3480156102a957600080fd5b506102bd6102b8366004613c0c565b6108d6565b005b3480156102cb57600080fd5b506102bd6102da366004613b34565b6109fc565b3480156102eb57600080fd5b506102bd6102fa366004613d35565b610b0f565b34801561030b57600080fd5b506102bd610ec3565b34801561032057600080fd5b5061034a61032f366004613b34565b6005602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610266565b34801561036e57600080fd5b5061034a7f000000000000000000000000000000000000000000000000000000000000000081565b6102bd6103a4366004613dfe565b610fa2565b3480156103b557600080fd5b506102bd6103c4366004613e98565b61112e565b3480156103d557600080fd5b506102bd6103e4366004613f0a565b6112a9565b3480156103f557600080fd5b50600a5460ff165b6040519015158152602001610266565b34801561041957600080fd5b506102bd610428366004613b34565b611358565b34801561043957600080fd5b506102bd611520565b34801561044e57600080fd5b506102bd61045d366004613f4e565b611586565b34801561046e57600080fd5b5061034a7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104a257600080fd5b506104ca7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610266565b3480156104eb57600080fd5b506000546001600160a01b031661034a565b34801561050957600080fd5b506103fd610518366004613f4e565b611665565b34801561052957600080fd5b50610532611683565b6040516102669190613f67565b34801561054b57600080fd5b506102bd61055a366004613b34565b611810565b34801561056b57600080fd5b5061034a61057a366004613b34565b600c602052600090815260409020546001600160a01b031681565b3480156105a157600080fd5b5061034a6105b0366004614004565b80516020818301810180516002825292820191909301209152546001600160a01b031681565b3480156105e257600080fd5b5061028f60065481565b3480156105f857600080fd5b506102bd610607366004613f4e565b611898565b34801561061857600080fd5b5061028f610627366004613c0c565b600460209081526000928352604080842090915290825290205481565b34801561065057600080fd5b506102bd611927565b34801561066557600080fd5b506106d7610674366004613b34565b60036020526000908152604090205464ffffffffff81169068ffffffffffffffffff6501000000000082048116916e0100000000000000000000000000008104821691770100000000000000000000000000000000000000000000009091041684565b6040805164ffffffffff909516855268ffffffffffffffffff938416602086015291831691840191909152166060820152608001610266565b34801561071c57600080fd5b5061028f60075481565b34801561073257600080fd5b506102bd611989565b34801561074757600080fd5b50610781610756366004613f4e565b600b6020526000908152604090205463ffffffff81169064010000000090046001600160a01b031682565b6040805163ffffffff90931683526001600160a01b03909116602083015201610266565b3480156107b157600080fd5b506102bd6107c0366004614039565b6119eb565b3480156107d157600080fd5b506102bd6107e0366004613b34565b611be1565b3480156107f157600080fd5b506102bd610800366004613b34565b611cc0565b34801561081157600080fd5b5061034a610820366004614070565b611d4f565b6001602052600090815260409020805461083e9061408e565b80601f016020809104026020016040519081016040528092919081815260200182805461086a9061408e565b80156108b75780601f1061088c576101008083540402835291602001916108b7565b820191906000526020600020905b81548152906001019060200180831161089a57829003601f168201915b505050505081565b60006108cc848484611d64565b90505b9392505050565b6000546001600160a01b031633146109355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b038216610975576040517f24305eca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610980600883611e6e565b6109b6576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b039182166000908152600c6020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001691909216179055565b6000546001600160a01b03163314610a565760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b6001600160a01b03811660009081527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec602052604081205490819003610ac8576040517f30b93f1d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821660008181527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec6020526040812055610b0b903383611e90565b5050565b600a5460ff1615610b625760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161092c565b610b6d600885611e6e565b610ba3576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bad8484611f3e565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638371744e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3391906140f8565b6040805163ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015267ffffffffffffffff8316818301526001600160a01b038089166060830152608082018890528a1660a0808301919091528251808303909101815260c0909101909152909150600090610cb890859085612070565b6000888152600b6020908152604080832081518083019092525463ffffffff8116825264010000000090046001600160a01b031691810182905292935090819003610d2f576040517fa86a3b0e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151835160208086019190912067ffffffff0000000083831b1663ffffffff8a16176000908152915260409020610d878a7f00000000000000000000000000000000000000000000000000000000000000008b61219d565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663f856ddb68a84868e610dc4828861226b565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b168152600481019590955263ffffffff93909316602485015260448401919091526001600160a01b03166064830152608482015260a4016020604051808303816000875af1158015610e43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6791906140f8565b50326001600160a01b03167fc4980459837e213aedb84d9046eab1db050fec66cb9e046c4fe3b5578b01b20c8c888d8d8d8b88604051610ead9796959493929190614113565b60405180910390a2505050505050505050505050565b6000546001600160a01b03163314610f1d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b604051600090339047908381818185875af1925050503d8060008114610f5f576040519150601f19603f3d011682016040523d82523d6000602084013e610f64565b606091505b5050905080610f9f576040517f4e5610fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b6007543414610fdd576040517fc561806500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080610fea848461228f565b91509150600080600080610ffd8661236f565b8b516020808e019190912063ffffffff8f167f0000000000000000000000000000000000000000000000000000000000000000831b67ffffffff0000000016176000908152915260408120959950929750909550935091905090506110658d8d8d8d856123d8565b600061107186866124f3565b90506000611089828663ffffffff8e1660011461266f565b909550905060008061109d8685898d6128bf565b909250905034156110b1576110b1866129d8565b6040805163ffffffff8b1681526001600160a01b03868116602083015291810185905283821660608201526080810183905260a08101879052908716907f7864397c00beabf21ab17a04795e450354505d879a634dd2632f4fdc4b5ba04e9060c00160405180910390a25050505050505050505050505050505050565b6000546001600160a01b031633146111885760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b6001600160a01b0385166111c8576040517f76998feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111d3600886612a72565b611209576040517f1191732500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61121286612a87565b6001600160a01b0385166000908152600160209081526040909120875161123b92890190613a54565b508460028760405161124d9190614171565b90815260405190819003602001902080546001600160a01b03929092167fffffffffffffffffffffffff00000000000000000000000000000000000000009092169190911790556112a18585858585612c27565b505050505050565b6000546001600160a01b031633146113035760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b61130e600886611e6e565b611344576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113518585858585612c27565b5050505050565b6000546001600160a01b031633146113b25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b6113bd600882612e1d565b6113f3576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116600090815260016020526040812080546114169061408e565b80601f01602080910402602001604051908101604052809291908181526020018280546114429061408e565b801561148f5780601f106114645761010080835404028352916020019161148f565b820191906000526020600020905b81548152906001019060200180831161147257829003601f168201915b505050506001600160a01b03841660009081526001602052604081209293506114b9929150613ad4565b6002816040516114c99190614171565b908152604080516020928190038301902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556001600160a01b0393909316600090815260039091529182209190915550565b6000546001600160a01b0316331461157a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b6115846000612e32565b565b6000546001600160a01b031633146115e05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b6115f060026402540be4006141bc565b811115611629576040517f28562c4700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60068190556040518181527fd10d75876659a287a59a6ccfa2e3fff42f84d94b542837acd30bc184d562de40906020015b60405180910390a150565b6000806116723084612e9a565b6001600160a01b03163b1192915050565b606060006116916008612f86565b90508067ffffffffffffffff8111156116ac576116ac613c57565b6040519080825280602002602001820160405280156116f257816020015b6040805180820190915260608152600060208201528152602001906001900390816116ca5790505b50915060005b8181101561180b57600061170d600883612f90565b9050604051806040016040528060016000846001600160a01b03166001600160a01b03168152602001908152602001600020805461174a9061408e565b80601f01602080910402602001604051908101604052809291908181526020018280546117769061408e565b80156117c35780601f10611798576101008083540402835291602001916117c3565b820191906000526020600020905b8154815290600101906020018083116117a657829003601f168201915b50505050508152602001826001600160a01b03168152508483815181106117ec576117ec6141f7565b602002602001018190525050808061180390614226565b9150506116f8565b505090565b3360008181526005602090815260409182902080547fffffffffffffffffffffffff000000000000000000000000000000000000000081166001600160a01b03878116918217909355845192909116808352928201529092917f9dfcadd14a1ddfb19c51e84b87452ca32a43c5559e9750d1575c77105cdeac1e910160405180910390a25050565b6000546001600160a01b031633146118f25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b60078190556040518181527f5e8bad84cb22c143a6757c7f1252a7d53493816880330977cc99bb7c15aaf6b49060200161165a565b6000546001600160a01b031633146119815760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b611584612f9c565b6000546001600160a01b031633146119e35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b61158461305f565b6000546001600160a01b03163314611a455760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b821580611a5157504683145b15611a88576040517f3f8f40a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168263ffffffff1603611aed576040517f93c970c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8216156001841414611b30576040517f93c970c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116611b70576040517f24305eca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051808201825263ffffffff93841681526001600160a01b0392831660208083019182526000968752600b905291909420935184549151909216640100000000027fffffffffffffffff0000000000000000000000000000000000000000000000009091169190921617179055565b6000546001600160a01b03163314611c3b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161092c565b6001600160a01b038116611cb75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161092c565b610f9f81612e32565b3360009081526004602090815260408083206001600160a01b038516845290915281205490819003611d1e576040517f30b93f1d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360008181526004602090815260408083206001600160a01b0387168085529252822091909155610b0b9183611e90565b6000611d5b83836124f3565b90505b92915050565b6001600160a01b03831660009081526003602090815260408083208151608081018352905464ffffffffff811680835268ffffffffffffffffff6501000000000083048116958401959095526e010000000000000000000000000000820485169383019390935277010000000000000000000000000000000000000000000000900490921660608301526402540be40090611dff9086614240565b611e0991906141bc565b9150600083611e1c578160200151611e22565b81604001515b68ffffffffffffffffff16905080831015611e3b578092505b816060015168ffffffffffffffffff16831115611e6557816060015168ffffffffffffffffff1692505b50509392505050565b6001600160a01b03811660009081526001830160205260408120541515611d5b565b6040516001600160a01b038316602482015260448101829052611f399084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613100565b505050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015611fa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fc4919061425f565b9050611fdb6001600160a01b0385163330866131e5565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281906001600160a01b038616906370a0823190602401602060405180830381865afa15801561203a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061205e919061425f565b6120689190614278565b949350505050565b606060a08351146120ad576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff84166120f8578151156120f1576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50816108cf565b60001963ffffffff85160161216b576080825114612142576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b828260405160200161215592919061428f565b60405160208183030381529060405290506108cf565b6040517f523fa8d500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0383811660248301526000919085169063dd62ed3e90604401602060405180830381865afa158015612206573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061222a919061425f565b90508181101561226557801561224f5761224f6001600160a01b038516846000613236565b6122656001600160a01b03851684600019613236565b50505050565b6000611d5b6122836001600160a01b03851684612e9a565b6001600160a01b031690565b60608063ffffffff84166122f05760a08351146122d8576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50506040805160208101909152600081528190612368565b60001963ffffffff85160161216b57608061230c60a0826142b4565b61231691906142b4565b83511461234f576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b828060200190518101906123639190614311565b915091505b9250929050565b600080600080600060a08651146123b2576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858060200190518101906123c69190614375565b939a9299509097509550909350915050565b60006123e382613384565b905060006357ecfd2860e01b878787876040516024016124069493929190614407565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091529050600061249d6001600160a01b0384167f000000000000000000000000000000000000000000000000000000000000000084613430565b9050808060200190518101906124b3919061442e565b6124e9576040517f182f34eb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cb75c11c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612554573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612578919061444b565b9050806001600160a01b03166378a0565e856125a3866001600160a01b03166001600160a01b031690565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815263ffffffff9290921660048301526024820152604401602060405180830381865afa158015612602573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612626919061444b565b91506001600160a01b038216612668576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5092915050565b60008061267d600886611e6e565b6126b3576040517f53b5a66c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6126be858585611d64565b90508381106126f9576040517f3eae42e400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526005602052604090205481850392506001600160a01b0316806127ae576001600160a01b03861660009081527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec6020526040812080548492906127639084906142b4565b909155505060408051600080825260208201529081018390527f108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca5619060600160405180910390a16128b6565b60006402540be400600654846127c49190614240565b6127ce91906141bc565b905060006127dc8285614278565b6001600160a01b03891660009081527f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec60205260408120805492935084929091906128289084906142b4565b90915550506001600160a01b038084166000908152600460209081526040808320938c16835292905290812080548392906128649084906142b4565b9091555050604080516001600160a01b0385168152602081018390529081018390527f108516ddcf5ba43cea6bb2cd5ff6d59ac196c1c86ccb9178332b9dd72d1ca5619060600160405180910390a150505b50935093915050565b60008082516000036128e9576128df6001600160a01b0386168786611e90565b50839050826129cf565b6001600160a01b038086166000908152600c602052604090205416806129295761291d6001600160a01b0387168887611e90565b858592509250506129cf565b6000806000806129388861343f565b935093509350935061294a85846134a3565b96506001600160a01b03871661297e5761296e6001600160a01b038b168c8b611e90565b89899650965050505050506129cf565b6129898a868b61219d565b6129978585858c86866135a6565b9550856000036129b55761296e6001600160a01b038b168c8b611e90565b6129c96001600160a01b0388168c88611e90565b50505050505b94509492505050565b6000816001600160a01b03163460405160006040518083038185875af1925050503d8060008114612a25576040519150601f19603f3d011682016040523d82523d6000602084013e612a2a565b606091505b505090507ff9b0951a3a6282341e1ba9414555d42d04e99076337702ee6dc484a706bfd68381612a5b576000612a5d565b345b60405190815260200160405180910390a15050565b6000611d5b836001600160a01b03841661365d565b60006001600160a01b0316600282604051612aa29190614171565b908152604051908190036020019020546001600160a01b031614612af2576040517f82ca3adf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80518190600510612b2f576040517f3f8fe5a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b6005811015611f39576040518060400160405280600581526020017f434354502e0000000000000000000000000000000000000000000000000000008152508181518110612b8257612b826141f7565b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916828281518110612bc157612bc16141f7565b01602001517fff000000000000000000000000000000000000000000000000000000000000001614612c1f576040517f3f8fe5a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600101612b32565b62989680841115612c64576040517f76998feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81831115612c9e576040517f76998feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80821115612cd8576040517f76998feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060800160405280612cec866136ac565b64ffffffffff168152602001612d01856136f1565b68ffffffffffffffffff168152602001612d1a846136f1565b68ffffffffffffffffff168152602001612d33836136f1565b68ffffffffffffffffff9081169091526001600160a01b039096166000908152600360209081526040918290208351815492850151938501516060909501518a16770100000000000000000000000000000000000000000000000276ffffffffffffffffffffffffffffffffffffffffffffff958b166e01000000000000000000000000000002959095166dffffffffffffffffffffffffffff94909a1665010000000000027fffffffffffffffffffffffffffffffffffff000000000000000000000000000090931664ffffffffff909116179190911791909116969096171790945550505050565b6000611d5b836001600160a01b038416613736565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000611d5b83836040518060400160405280602081526020017f602036038060203d373d3d3d923d343d355af13d82803e903d91601e57fd5bf3815250604051602001612ee79190614468565b60405160208183030381529060405280519060200120604051602001612f6d939291907fff00000000000000000000000000000000000000000000000000000000000000815260609390931b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660018401526015830191909152603582015260550190565b6040516020818303038152906040528051906020012090565b6000611d5e825490565b6000611d5b8383613829565b600a5460ff1615612fef5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161092c565b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586130423390565b6040516001600160a01b03909116815260200160405180910390a1565b600a5460ff166130b15760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161092c565b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613042565b6000613155826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166138539092919063ffffffff16565b805190915015611f395780806020019051810190613173919061442e565b611f395760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161092c565b6040516001600160a01b03808516602483015283166044820152606481018290526122659085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611ed5565b8015806132c957506040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03838116602483015284169063dd62ed3e90604401602060405180830381865afa1580156132a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132c7919061425f565b155b61333b5760405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000606482015260840161092c565b6040516001600160a01b038316602482015260448101829052611f399084907f095ea7b30000000000000000000000000000000000000000000000000000000090606401611ed5565b6000806040518060400160405280602081526020017f602036038060203d373d3d3d923d343d355af13d82803e903d91601e57fd5bf38152506040516020016133cd9190614468565b6040516020818303038152906040529050828151602083016000f591506001600160a01b03821661342a576040517f27afa9fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50919050565b60606108cc8484846000613862565b6000806000806080855114613480576040517f74593f8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8480602001905181019061349491906144e5565b93509350935093509193509193565b6040805160ff831660248083019190915282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f82b86600000000000000000000000000000000000000000000000000000000001790529051600091829182916001600160a01b0387169161352b9190614171565b600060405180830381855afa9150503d8060008114613566576040519150601f19603f3d011682016040523d82523d6000602084013e61356b565b606091505b509150915081801561357e575080516020145b156135995761359261358f82614528565b90565b925061359e565b600092505b505092915050565b6040517f9169558600000000000000000000000000000000000000000000000000000000815260ff8087166004830152851660248201526044810184905260648101829052608481018390526000906001600160a01b0388169063916955869060a4016020604051808303816000875af1925050508015613644575060408051601f3d908101601f191682019092526136419181019061425f565b60015b61365057506000613653565b90505b9695505050505050565b60008181526001830160205260408120546136a457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611d5e565b506000611d5e565b600064ffffffffff8211156136ed576040517fe58d471800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b600068ffffffffffffffffff8211156136ed576040517fe58d471800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000818152600183016020526040812054801561381f57600061375a600183614278565b855490915060009061376e90600190614278565b90508181146137d357600086600001828154811061378e5761378e6141f7565b90600052602060002001549050808760000184815481106137b1576137b16141f7565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806137e4576137e461454c565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050611d5e565b6000915050611d5e565b6000826000018281548110613840576138406141f7565b9060005260206000200154905092915050565b60606108cc84846000856138ad565b60606138a46001600160a01b0385168460405160200161388392919061457b565b60408051601f198184030181529190526001600160a01b03871690846139f5565b95945050505050565b6060824710156139255760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161092c565b6001600160a01b0385163b61397c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161092c565b600080866001600160a01b031685876040516139989190614171565b60006040518083038185875af1925050503d80600081146139d5576040519150601f19603f3d011682016040523d82523d6000602084013e6139da565b606091505b50915091506139ea828286613a1b565b979650505050505050565b60606108cc8484846040518060600160405280602981526020016145a2602991396138ad565b60608315613a2a5750816108cf565b825115613a3a5782518084602001fd5b8160405162461bcd60e51b815260040161092c9190613ba9565b828054613a609061408e565b90600052602060002090601f016020900481019282613a825760008555613ac8565b82601f10613a9b57805160ff1916838001178555613ac8565b82800160010185558215613ac8579182015b82811115613ac8578251825591602001919060010190613aad565b506136ed929150613b0a565b508054613ae09061408e565b6000825580601f10613af0575050565b601f016020900490600052602060002090810190610f9f91905b5b808211156136ed5760008155600101613b0b565b6001600160a01b0381168114610f9f57600080fd5b600060208284031215613b4657600080fd5b81356108cf81613b1f565b60005b83811015613b6c578181015183820152602001613b54565b838111156122655750506000910152565b60008151808452613b95816020860160208601613b51565b601f01601f19169290920160200192915050565b602081526000611d5b6020830184613b7d565b8015158114610f9f57600080fd5b600080600060608486031215613bdf57600080fd5b8335613bea81613b1f565b9250602084013591506040840135613c0181613bbc565b809150509250925092565b60008060408385031215613c1f57600080fd5b8235613c2a81613b1f565b91506020830135613c3a81613b1f565b809150509250929050565b63ffffffff81168114610f9f57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715613caf57613caf613c57565b604052919050565b600067ffffffffffffffff821115613cd157613cd1613c57565b50601f01601f191660200190565b600082601f830112613cf057600080fd5b8135613d03613cfe82613cb7565b613c86565b818152846020838601011115613d1857600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c08789031215613d4e57600080fd5b8635613d5981613b1f565b9550602087013594506040870135613d7081613b1f565b9350606087013592506080870135613d8781613c45565b915060a087013567ffffffffffffffff811115613da357600080fd5b613daf89828a01613cdf565b9150509295509295509295565b60008083601f840112613dce57600080fd5b50813567ffffffffffffffff811115613de657600080fd5b60208301915083602082850101111561236857600080fd5b60008060008060008060808789031215613e1757600080fd5b863567ffffffffffffffff80821115613e2f57600080fd5b613e3b8a838b01613dbc565b90985096506020890135915080821115613e5457600080fd5b613e608a838b01613dbc565b909650945060408901359150613e7582613c45565b90925060608801359080821115613e8b57600080fd5b50613daf89828a01613cdf565b60008060008060008060c08789031215613eb157600080fd5b863567ffffffffffffffff811115613ec857600080fd5b613ed489828a01613cdf565b9650506020870135613ee581613b1f565b95989597505050506040840135936060810135936080820135935060a0909101359150565b600080600080600060a08688031215613f2257600080fd5b8535613f2d81613b1f565b97602087013597506040870135966060810135965060800135945092505050565b600060208284031215613f6057600080fd5b5035919050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015613ff6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc089840301855281518051878552613fd088860182613b7d565b918901516001600160a01b03169489019490945294870194925090860190600101613f8e565b509098975050505050505050565b60006020828403121561401657600080fd5b813567ffffffffffffffff81111561402d57600080fd5b61206884828501613cdf565b60008060006060848603121561404e57600080fd5b83359250602084013561406081613c45565b91506040840135613c0181613b1f565b6000806040838503121561408357600080fd5b8235613c2a81613c45565b600181811c908216806140a257607f821691505b60208210810361342a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b805167ffffffffffffffff811681146140f357600080fd5b919050565b60006020828403121561410a57600080fd5b611d5b826140db565b87815267ffffffffffffffff871660208201526001600160a01b038616604082015284606082015263ffffffff8416608082015260e060a0820152600061415d60e0830185613b7d565b90508260c083015298975050505050505050565b60008251614183818460208701613b51565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000826141f2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060001982036142395761423961418d565b5060010190565b600081600019048311821515161561425a5761425a61418d565b500290565b60006020828403121561427157600080fd5b5051919050565b60008282101561428a5761428a61418d565b500390565b6040815260006142a26040830185613b7d565b82810360208401526138a48185613b7d565b600082198211156142c7576142c761418d565b500190565b600082601f8301126142dd57600080fd5b81516142eb613cfe82613cb7565b81815284602083860101111561430057600080fd5b612068826020830160208701613b51565b6000806040838503121561432457600080fd5b825167ffffffffffffffff8082111561433c57600080fd5b614348868387016142cc565b9350602085015191508082111561435e57600080fd5b5061436b858286016142cc565b9150509250929050565b600080600080600060a0868803121561438d57600080fd5b855161439881613c45565b94506143a6602087016140db565b935060408601516143b681613b1f565b6060870151608088015191945092506143ce81613b1f565b809150509295509295909350565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b60408152600061441b6040830186886143dc565b82810360208401526139ea8185876143dc565b60006020828403121561444057600080fd5b81516108cf81613bbc565b60006020828403121561445d57600080fd5b81516108cf81613b1f565b7f7f000000000000000000000000000000000000000000000000000000000000008152600082516144a0816001850160208701613b51565b7f3d5260203df300000000000000000000000000000000000000000000000000006001939091019283015250600701919050565b805160ff811681146140f357600080fd5b600080600080608085870312156144fb57600080fd5b614504856144d4565b9350614512602086016144d4565b6040860151606090960151949790965092505050565b8051602080830151919081101561342a5760001960209190910360031b1b16919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b82815260008251614593816020850160208701613b51565b91909101602001939250505056fe416464726573733a206c6f772d6c6576656c2063616c6c20776974682076616c7565206661696c6564a2646970667358221220599a9c446de9cb152ce66f97e0b3fbf4a2a0aebd13e35d04745bca99d587c69064736f6c634300080d0033",
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
func DeploySynapseCCTP(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger_ common.Address, owner_ common.Address) (common.Address, *types.Transaction, *SynapseCCTP, error) {
	parsed, err := SynapseCCTPMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseCCTPBin), backend, tokenMessenger_, owner_)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SynapseCCTP *SynapseCCTPCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynapseCCTP.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SynapseCCTP *SynapseCCTPSession) Paused() (bool, error) {
	return _SynapseCCTP.Contract.Paused(&_SynapseCCTP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SynapseCCTP *SynapseCCTPCallerSession) Paused() (bool, error) {
	return _SynapseCCTP.Contract.Paused(&_SynapseCCTP.CallOpts)
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

// PauseSending is a paid mutator transaction binding the contract method 0xd77938e4.
//
// Solidity: function pauseSending() returns()
func (_SynapseCCTP *SynapseCCTPTransactor) PauseSending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "pauseSending")
}

// PauseSending is a paid mutator transaction binding the contract method 0xd77938e4.
//
// Solidity: function pauseSending() returns()
func (_SynapseCCTP *SynapseCCTPSession) PauseSending() (*types.Transaction, error) {
	return _SynapseCCTP.Contract.PauseSending(&_SynapseCCTP.TransactOpts)
}

// PauseSending is a paid mutator transaction binding the contract method 0xd77938e4.
//
// Solidity: function pauseSending() returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) PauseSending() (*types.Transaction, error) {
	return _SynapseCCTP.Contract.PauseSending(&_SynapseCCTP.TransactOpts)
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
// Solidity: function setCircleTokenPool(address circleToken, address swap) returns()
func (_SynapseCCTP *SynapseCCTPTransactor) SetCircleTokenPool(opts *bind.TransactOpts, circleToken common.Address, pool common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "setCircleTokenPool", circleToken, pool)
}

// SetCircleTokenPool is a paid mutator transaction binding the contract method 0x2cc9e7e5.
//
// Solidity: function setCircleTokenPool(address circleToken, address swap) returns()
func (_SynapseCCTP *SynapseCCTPSession) SetCircleTokenPool(circleToken common.Address, pool common.Address) (*types.Transaction, error) {
	return _SynapseCCTP.Contract.SetCircleTokenPool(&_SynapseCCTP.TransactOpts, circleToken, pool)
}

// SetCircleTokenPool is a paid mutator transaction binding the contract method 0x2cc9e7e5.
//
// Solidity: function setCircleTokenPool(address circleToken, address swap) returns()
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

// UnpauseSending is a paid mutator transaction binding the contract method 0xe7a64a80.
//
// Solidity: function unpauseSending() returns()
func (_SynapseCCTP *SynapseCCTPTransactor) UnpauseSending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseCCTP.contract.Transact(opts, "unpauseSending")
}

// UnpauseSending is a paid mutator transaction binding the contract method 0xe7a64a80.
//
// Solidity: function unpauseSending() returns()
func (_SynapseCCTP *SynapseCCTPSession) UnpauseSending() (*types.Transaction, error) {
	return _SynapseCCTP.Contract.UnpauseSending(&_SynapseCCTP.TransactOpts)
}

// UnpauseSending is a paid mutator transaction binding the contract method 0xe7a64a80.
//
// Solidity: function unpauseSending() returns()
func (_SynapseCCTP *SynapseCCTPTransactorSession) UnpauseSending() (*types.Transaction, error) {
	return _SynapseCCTP.Contract.UnpauseSending(&_SynapseCCTP.TransactOpts)
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
	OriginDomain uint32
	Recipient    common.Address
	MintToken    common.Address
	Fee          *big.Int
	Token        common.Address
	Amount       *big.Int
	RequestID    [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestFulfilled is a free log retrieval operation binding the contract event 0x7864397c00beabf21ab17a04795e450354505d879a634dd2632f4fdc4b5ba04e.
//
// Solidity: event CircleRequestFulfilled(uint32 originDomain, address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterCircleRequestFulfilled(opts *bind.FilterOpts, recipient []common.Address) (*SynapseCCTPCircleRequestFulfilledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "CircleRequestFulfilled", recipientRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPCircleRequestFulfilledIterator{contract: _SynapseCCTP.contract, event: "CircleRequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchCircleRequestFulfilled is a free log subscription operation binding the contract event 0x7864397c00beabf21ab17a04795e450354505d879a634dd2632f4fdc4b5ba04e.
//
// Solidity: event CircleRequestFulfilled(uint32 originDomain, address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchCircleRequestFulfilled(opts *bind.WatchOpts, sink chan<- *SynapseCCTPCircleRequestFulfilled, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "CircleRequestFulfilled", recipientRule)
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

// ParseCircleRequestFulfilled is a log parse operation binding the contract event 0x7864397c00beabf21ab17a04795e450354505d879a634dd2632f4fdc4b5ba04e.
//
// Solidity: event CircleRequestFulfilled(uint32 originDomain, address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 requestID)
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
	Sender           common.Address
	Nonce            uint64
	Token            common.Address
	Amount           *big.Int
	RequestVersion   uint32
	FormattedRequest []byte
	RequestID        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestSent is a free log retrieval operation binding the contract event 0xc4980459837e213aedb84d9046eab1db050fec66cb9e046c4fe3b5578b01b20c.
//
// Solidity: event CircleRequestSent(uint256 chainId, address indexed sender, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterCircleRequestSent(opts *bind.FilterOpts, sender []common.Address) (*SynapseCCTPCircleRequestSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "CircleRequestSent", senderRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPCircleRequestSentIterator{contract: _SynapseCCTP.contract, event: "CircleRequestSent", logs: logs, sub: sub}, nil
}

// WatchCircleRequestSent is a free log subscription operation binding the contract event 0xc4980459837e213aedb84d9046eab1db050fec66cb9e046c4fe3b5578b01b20c.
//
// Solidity: event CircleRequestSent(uint256 chainId, address indexed sender, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 requestID)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchCircleRequestSent(opts *bind.WatchOpts, sink chan<- *SynapseCCTPCircleRequestSent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "CircleRequestSent", senderRule)
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

// ParseCircleRequestSent is a log parse operation binding the contract event 0xc4980459837e213aedb84d9046eab1db050fec66cb9e046c4fe3b5578b01b20c.
//
// Solidity: event CircleRequestSent(uint256 chainId, address indexed sender, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 requestID)
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

// SynapseCCTPPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SynapseCCTP contract.
type SynapseCCTPPausedIterator struct {
	Event *SynapseCCTPPaused // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPPaused)
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
		it.Event = new(SynapseCCTPPaused)
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
func (it *SynapseCCTPPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPPaused represents a Paused event raised by the SynapseCCTP contract.
type SynapseCCTPPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterPaused(opts *bind.FilterOpts) (*SynapseCCTPPausedIterator, error) {

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPPausedIterator{contract: _SynapseCCTP.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SynapseCCTPPaused) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPPaused)
				if err := _SynapseCCTP.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SynapseCCTP *SynapseCCTPFilterer) ParsePaused(log types.Log) (*SynapseCCTPPaused, error) {
	event := new(SynapseCCTPPaused)
	if err := _SynapseCCTP.contract.UnpackLog(event, "Paused", log); err != nil {
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

// SynapseCCTPUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SynapseCCTP contract.
type SynapseCCTPUnpausedIterator struct {
	Event *SynapseCCTPUnpaused // Event containing the contract specifics and raw log

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
func (it *SynapseCCTPUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseCCTPUnpaused)
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
		it.Event = new(SynapseCCTPUnpaused)
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
func (it *SynapseCCTPUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseCCTPUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseCCTPUnpaused represents a Unpaused event raised by the SynapseCCTP contract.
type SynapseCCTPUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SynapseCCTP *SynapseCCTPFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SynapseCCTPUnpausedIterator, error) {

	logs, sub, err := _SynapseCCTP.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPUnpausedIterator{contract: _SynapseCCTP.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SynapseCCTP *SynapseCCTPFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SynapseCCTPUnpaused) (event.Subscription, error) {

	logs, sub, err := _SynapseCCTP.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseCCTPUnpaused)
				if err := _SynapseCCTP.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SynapseCCTP *SynapseCCTPFilterer) ParseUnpaused(log types.Log) (*SynapseCCTPUnpaused, error) {
	event := new(SynapseCCTPUnpaused)
	if err := _SynapseCCTP.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseCCTPEventsMetaData contains all meta data concerning the SynapseCCTPEvents contract.
var SynapseCCTPEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"requestVersion\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"formattedRequest\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"CircleRequestSent\",\"type\":\"event\"}]",
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
	OriginDomain uint32
	Recipient    common.Address
	MintToken    common.Address
	Fee          *big.Int
	Token        common.Address
	Amount       *big.Int
	RequestID    [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestFulfilled is a free log retrieval operation binding the contract event 0x7864397c00beabf21ab17a04795e450354505d879a634dd2632f4fdc4b5ba04e.
//
// Solidity: event CircleRequestFulfilled(uint32 originDomain, address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) FilterCircleRequestFulfilled(opts *bind.FilterOpts, recipient []common.Address) (*SynapseCCTPEventsCircleRequestFulfilledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.FilterLogs(opts, "CircleRequestFulfilled", recipientRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsCircleRequestFulfilledIterator{contract: _SynapseCCTPEvents.contract, event: "CircleRequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchCircleRequestFulfilled is a free log subscription operation binding the contract event 0x7864397c00beabf21ab17a04795e450354505d879a634dd2632f4fdc4b5ba04e.
//
// Solidity: event CircleRequestFulfilled(uint32 originDomain, address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) WatchCircleRequestFulfilled(opts *bind.WatchOpts, sink chan<- *SynapseCCTPEventsCircleRequestFulfilled, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.WatchLogs(opts, "CircleRequestFulfilled", recipientRule)
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

// ParseCircleRequestFulfilled is a log parse operation binding the contract event 0x7864397c00beabf21ab17a04795e450354505d879a634dd2632f4fdc4b5ba04e.
//
// Solidity: event CircleRequestFulfilled(uint32 originDomain, address indexed recipient, address mintToken, uint256 fee, address token, uint256 amount, bytes32 requestID)
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
	Sender           common.Address
	Nonce            uint64
	Token            common.Address
	Amount           *big.Int
	RequestVersion   uint32
	FormattedRequest []byte
	RequestID        [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCircleRequestSent is a free log retrieval operation binding the contract event 0xc4980459837e213aedb84d9046eab1db050fec66cb9e046c4fe3b5578b01b20c.
//
// Solidity: event CircleRequestSent(uint256 chainId, address indexed sender, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) FilterCircleRequestSent(opts *bind.FilterOpts, sender []common.Address) (*SynapseCCTPEventsCircleRequestSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.FilterLogs(opts, "CircleRequestSent", senderRule)
	if err != nil {
		return nil, err
	}
	return &SynapseCCTPEventsCircleRequestSentIterator{contract: _SynapseCCTPEvents.contract, event: "CircleRequestSent", logs: logs, sub: sub}, nil
}

// WatchCircleRequestSent is a free log subscription operation binding the contract event 0xc4980459837e213aedb84d9046eab1db050fec66cb9e046c4fe3b5578b01b20c.
//
// Solidity: event CircleRequestSent(uint256 chainId, address indexed sender, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 requestID)
func (_SynapseCCTPEvents *SynapseCCTPEventsFilterer) WatchCircleRequestSent(opts *bind.WatchOpts, sink chan<- *SynapseCCTPEventsCircleRequestSent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SynapseCCTPEvents.contract.WatchLogs(opts, "CircleRequestSent", senderRule)
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

// ParseCircleRequestSent is a log parse operation binding the contract event 0xc4980459837e213aedb84d9046eab1db050fec66cb9e046c4fe3b5578b01b20c.
//
// Solidity: event CircleRequestSent(uint256 chainId, address indexed sender, uint64 nonce, address token, uint256 amount, uint32 requestVersion, bytes formattedRequest, bytes32 requestID)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122079566f2a471471f474a29e755c18b605c06a9cd08e01e2e737b04d311e49116164736f6c634300080d0033",
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
