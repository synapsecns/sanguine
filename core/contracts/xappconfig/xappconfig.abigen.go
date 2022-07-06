// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xappconfig

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a5c6fca0fd672236e4b5dba0312948c5d3249fd3a47379cc28707ddb3665001964736f6c634300080d0033",
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

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220399610f74191be940c1541cc14acc31ebf0045ee753c764f7f145c35f4324e6e64736f6c634300080d0033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
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

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContextUpgradeable contract.
type ContextUpgradeableInitializedIterator struct {
	Event *ContextUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ContextUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContextUpgradeableInitialized)
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
		it.Event = new(ContextUpgradeableInitialized)
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
func (it *ContextUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContextUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContextUpgradeableInitialized represents a Initialized event raised by the ContextUpgradeable contract.
type ContextUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContextUpgradeableInitializedIterator, error) {

	logs, sub, err := _ContextUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableInitializedIterator{contract: _ContextUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContextUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ContextUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContextUpgradeableInitialized)
				if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) ParseInitialized(log types.Log) (*ContextUpgradeableInitialized, error) {
	event := new(ContextUpgradeableInitialized)
	if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a37cd07f5ec5a836f7376ac96b489d5bfb20cbc8e61c8980a78e57c384b2e4ee64736f6c634300080d0033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// HomeMetaData contains all meta data concerning the Home contract.
var HomeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destinationAndNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"committedRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"newRoot\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature2\",\"type\":\"bytes\"}],\"name\":\"DoubleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"ImproperUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updaterManager\",\"type\":\"address\"}],\"name\":\"NewUpdaterManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"UpdaterSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_BODY_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committedRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_oldRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"_newRoot\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature2\",\"type\":\"bytes\"}],\"name\":\"doubleUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_oldRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"improperUpdate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUpdaterManager\",\"name\":\"_updaterManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_item\",\"type\":\"bytes32\"}],\"name\":\"queueContains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueEnd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updaterManager\",\"type\":\"address\"}],\"name\":\"setUpdaterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumHome.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_committedRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_new\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_committedRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updaterManager\",\"outputs\":[{\"internalType\":\"contractIUpdaterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"522ae002": "MAX_MESSAGE_BODY_BYTES()",
		"ffa1ad74": "VERSION()",
		"67a6771d": "committedRoot()",
		"06661abd": "count()",
		"59e62328": "dispatch(uint32,bytes32,uint32,bytes)",
		"19d9d21a": "doubleUpdate(bytes32,bytes32[2],bytes,bytes)",
		"45630b1a": "homeDomainHash()",
		"8e4e30e0": "improperUpdate(bytes32,bytes32,bytes)",
		"c4d66de8": "initialize(address)",
		"8d3638f4": "localDomain()",
		"b95a2001": "nonces(uint32)",
		"8da5cb5b": "owner()",
		"2bef2892": "queueContains(bytes32)",
		"f6d16102": "queueEnd()",
		"ab91c7b0": "queueLength()",
		"715018a6": "renounceOwnership()",
		"ebf0c717": "root()",
		"9d54f419": "setUpdater(address)",
		"9776120e": "setUpdaterManager(address)",
		"c19d93fb": "state()",
		"36e104de": "suggestUpdate()",
		"f2fde38b": "transferOwnership(address)",
		"fd54b228": "tree()",
		"b31c01fb": "update(bytes32,bytes32,bytes)",
		"df034cd0": "updater()",
		"9df6c8e1": "updaterManager()",
	},
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002c7b38038062002c7b833981016040819052620000349162000043565b63ffffffff1660805262000072565b6000602082840312156200005657600080fd5b815163ffffffff811681146200006b57600080fd5b9392505050565b608051612bd1620000aa6000396000818161025a01528181610669015281816107d201528181610cf00152610ff30152612bd16000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c80639d54f419116100ee578063c4d66de811610097578063f2fde38b11610071578063f2fde38b146103da578063f6d16102146103ed578063fd54b2281461042a578063ffa1ad741461043457600080fd5b8063c4d66de81461039f578063df034cd0146103b2578063ebf0c717146103d257600080fd5b8063b31c01fb116100c8578063b31c01fb14610332578063b95a200114610345578063c19d93fb1461036c57600080fd5b80639d54f419146102f65780639df6c8e114610309578063ab91c7b01461032a57600080fd5b806359e623281161015b5780638d3638f4116101355780638d3638f4146102555780638da5cb5b146102915780638e4e30e0146102d05780639776120e146102e357600080fd5b806359e623281461023057806367a6771d14610243578063715018a61461024d57600080fd5b806336e104de1161018c57806336e104de1461020257806345630b1a1461021f578063522ae0021461022757600080fd5b806306661abd146101b357806319d9d21a146101ca5780632bef2892146101df575b600080fd5b6054545b6040519081526020015b60405180910390f35b6101dd6101d836600461247e565b61044e565b005b6101f26101ed36600461250d565b6105cb565b60405190151581526020016101c1565b61020a6105de565b604080519283526020830191909152016101c1565b6101b7610662565b6101b761080081565b6101dd61023e366004612614565b610692565b6101b761011e5481565b6101dd6108da565b61027c7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101c1565b60b85473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c1565b6101f26102de36600461267c565b610943565b6101dd6102f13660046126ee565b610ad2565b6101dd6103043660046126ee565b610b45565b61011d546102ab9073ffffffffffffffffffffffffffffffffffffffff1681565b6101b7610bf9565b6101dd61034036600461267c565b610c34565b61027c61035336600461270b565b61011c6020526000908152604090205463ffffffff1681565b61011d546103929074010000000000000000000000000000000000000000900460ff1681565b6040516101c19190612755565b6101dd6103ad3660046126ee565b610d52565b60ea546102ab9073ffffffffffffffffffffffffffffffffffffffff1681565b6101b7610ee7565b6101dd6103e83660046126ee565b610ef3565b60015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff166000908152600260205260409020546101b7565b6054546101b79081565b61043c600081565b60405160ff90911681526020016101c1565b600261011d5474010000000000000000000000000000000000000000900460ff16600281111561048057610480612726565b036104d25760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064015b60405180910390fd5b604080516020601f86018190048102820181019092528481526105149188918835918890889081908401838280828437600092019190915250610fec92505050565b8015610563575061056386866001602002013584848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fec92505050565b801561057457508435602086013514155b156105c357610581611022565b7f2c3f60bab4170347826231b75a920b5053941ddebc6eed6fd2c25721648b186f8686868686866040516105ba9695949392919061280e565b60405180910390a15b505050505050565b60006105d860018361112d565b92915050565b60008061061a6001546fffffffffffffffffffffffffffffffff8082167001000000000000000000000000000000009092048116919091031690565b1561065e57505061011e5460015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff166000908152600260205260409020545b9091565b600061068d7f00000000000000000000000000000000000000000000000000000000000000006111f3565b905090565b600261011d5474010000000000000000000000000000000000000000900460ff1660028111156106c4576106c4612726565b036107115760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064016104c9565b610800815111156107645760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e67000000000000000000000000000000000000000060448201526064016104c9565b63ffffffff808516600090815261011c60205260409020541661078881600161287f565b63ffffffff868116600090815261011c6020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000016939092169290921790556107fc7f000000000000000000000000000000000000000000000000000000000000000033848989898961126c565b805160208201209091506108116034826112a7565b61086d61081c610ee7565b60018054700100000000000000000000000000000000908190046fffffffffffffffffffffffffffffffff908116830181166000818152600260205260409020949094558254169202919091179055565b60545467ffffffff00000000602089901b1663ffffffff85161790610894906001906128a7565b827f9d4c83d2e57d7d381feb264b44a5015e7f9ef26340f4fc46b558a6dc16dd811a61011e54866040516108c9929190612938565b60405180910390a450505050505050565b60b85473ffffffffffffffffffffffffffffffffffffffff1633146109415760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104c9565b565b6000600261011d5474010000000000000000000000000000000000000000900460ff16600281111561097757610977612726565b036109c45760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064016104c9565b6109cf848484610fec565b610a1b5760405162461bcd60e51b815260206004820152600c60248201527f217570646174657220736967000000000000000000000000000000000000000060448201526064016104c9565b61011e548414610a6d5760405162461bcd60e51b815260206004820152601460248201527f6e6f7420612063757272656e742075706461746500000000000000000000000060448201526064016104c9565b610a7860018461112d565b610ac757610a84611022565b7f6844fd5e21c932b5197b78ac11bf96e2eaa4e882dd0c88087060cf2065c04ab2848484604051610ab793929190612951565b60405180910390a1506001610acb565b5060005b9392505050565b60b85473ffffffffffffffffffffffffffffffffffffffff163314610b395760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104c9565b610b42816113c5565b50565b61011d5473ffffffffffffffffffffffffffffffffffffffff163314610bad5760405162461bcd60e51b815260206004820152600f60248201527f21757064617465724d616e61676572000000000000000000000000000000000060448201526064016104c9565b610bb6816114a3565b5061011d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000179055565b600061068d6001546fffffffffffffffffffffffffffffffff8082167001000000000000000000000000000000009092048116919091031690565b600261011d5474010000000000000000000000000000000000000000900460ff166002811115610c6657610c66612726565b03610cb35760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064016104c9565b610cbe838383610943565b610d4d575b6000610ccf6001611522565b9050828103610cde5750610ce4565b50610cc3565b8161011e8190555081837f000000000000000000000000000000000000000000000000000000000000000063ffffffff167f608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b284604051610d449190612979565b60405180910390a45b505050565b6000610d5e60016115ec565b90508015610d9357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610d9b611743565b610da4826113c5565b61011d54604080517fdf034cd00000000000000000000000000000000000000000000000000000000081529051610e3f9273ffffffffffffffffffffffffffffffffffffffff169163df034cd09160048083019260209291908290030181865afa158015610e16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3a919061298c565b6117c0565b61011d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610ee357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b600061068d603461184e565b60b85473ffffffffffffffffffffffffffffffffffffffff163314610f5a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104c9565b73ffffffffffffffffffffffffffffffffffffffff8116610fe35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104c9565b610b4281611861565b600061101a7f00000000000000000000000000000000000000000000000000000000000000008585856118d8565b949350505050565b61011d8054740200000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff8216179091556040517f5b3c2cbf00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911690635b3c2cbf90602401600060405180830381600087803b1580156110cc57600080fd5b505af11580156110e0573d6000803e3d6000fd5b505060ea5460405133935073ffffffffffffffffffffffffffffffffffffffff90911691507f98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce90600090a3565b81546000906fffffffffffffffffffffffffffffffff7001000000000000000000000000000000008204811691165b806fffffffffffffffffffffffffffffffff16826fffffffffffffffffffffffffffffffff1611156111e8576fffffffffffffffffffffffffffffffff821660009081526001860160205260409020548490036111be576001925050506105d8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9091019061115c565b506000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e083901b1660208201527f53594e000000000000000000000000000000000000000000000000000000000060248201526000906027015b604051602081830303815290604052805190602001209050919050565b60608787878787878760405160200161128b97969594939291906129a9565b6040516020818303038152906040529050979650505050505050565b602080830154906001906112bc906002612b48565b6112c691906128a7565b81106113145760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c0000000000000000000000000000000060448201526064016104c9565b6001016020830181905560005b60208110156113bc5781600116600103611350578284826020811061134857611348612796565b015550505050565b83816020811061136257611362612796565b01546040805160208101929092528101849052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209250600191821c9101611321565b50610d4d612b54565b73ffffffffffffffffffffffffffffffffffffffff81163b6114295760405162461bcd60e51b815260206004820152601860248201527f21636f6e747261637420757064617465724d616e61676572000000000000000060448201526064016104c9565b61011d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf9060200160405180910390a150565b60ea805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a9101610eda565b600061155882546fffffffffffffffffffffffffffffffff80821670010000000000000000000000000000000090920416111590565b1561158f576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50805460016fffffffffffffffffffffffffffffffff9182168101909116600081815291830160205260408220805492905582547fffffffffffffffffffffffffffffffff00000000000000000000000000000000161790915590565b60008054610100900460ff1615611689578160ff16600114801561160f5750303b155b6116815760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104c9565b506000919050565b60005460ff8084169116106117065760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104c9565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b600054610100900460ff166109415760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c9565b600054610100900460ff1661183d5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c9565b611845611967565b610b42816114a3565b60006105d88261185c6119ec565b611ead565b60b8805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000806118e4866111f3565b60408051602081019290925281018690526060810185905260800160405160208183030381529060405280519060200120905061192081611f74565b60ea5490915073ffffffffffffffffffffffffffffffffffffffff166119468285611faf565b73ffffffffffffffffffffffffffffffffffffffff16149695505050505050565b600054610100900460ff166119e45760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c9565b610941611fd3565b6119f461241d565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b6020811015611f6c57600182821c166000868360208110611eda57611eda612796565b0154905081600103611f17576040805160208101839052908101869052606001604051602081830303815290604052805190602001209450611f62565b84868460208110611f2a57611f2a612796565b6020020151604051602001611f49929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b5050600101611eb7565b505092915050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c0161124f565b6000806000611fbe8585612059565b91509150611fcb816120c7565b509392505050565b600054610100900460ff166120505760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c9565b61094133611861565b600080825160410361208f5760208301516040840151606085015160001a612083878285856122b3565b945094505050506120c0565b82516040036120b857602083015160408401516120ad8683836123cb565b9350935050506120c0565b506000905060025b9250929050565b60008160048111156120db576120db612726565b036120e35750565b60018160048111156120f7576120f7612726565b036121445760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104c9565b600281600481111561215857612158612726565b036121a55760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104c9565b60038160048111156121b9576121b9612726565b0361222c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104c9565b600481600481111561224057612240612726565b03610b425760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104c9565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156122ea57506000905060036123c2565b8460ff16601b1415801561230257508460ff16601c14155b1561231357506000905060046123c2565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612367573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166123bb576000600192509250506123c2565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83168161240160ff86901c601b612b83565b905061240f878288856122b3565b935093505050935093915050565b6040518061040001604052806020906020820280368337509192915050565b60008083601f84011261244e57600080fd5b50813567ffffffffffffffff81111561246657600080fd5b6020830191508360208285010111156120c057600080fd5b60008060008060008060a0878903121561249757600080fd5b8635955060608701888111156124ac57600080fd5b6020880195503567ffffffffffffffff808211156124c957600080fd5b6124d58a838b0161243c565b909650945060808901359150808211156124ee57600080fd5b506124fb89828a0161243c565b979a9699509497509295939492505050565b60006020828403121561251f57600080fd5b5035919050565b803563ffffffff8116811461173e57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261257a57600080fd5b813567ffffffffffffffff808211156125955761259561253a565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156125db576125db61253a565b816040528381528660208588010111156125f457600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561262a57600080fd5b61263385612526565b93506020850135925061264860408601612526565b9150606085013567ffffffffffffffff81111561266457600080fd5b61267087828801612569565b91505092959194509250565b60008060006060848603121561269157600080fd5b8335925060208401359150604084013567ffffffffffffffff8111156126b657600080fd5b6126c286828701612569565b9150509250925092565b73ffffffffffffffffffffffffffffffffffffffff81168114610b4257600080fd5b60006020828403121561270057600080fd5b8135610acb816126cc565b60006020828403121561271d57600080fd5b610acb82612526565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160038310612790577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b868152604086602083013760a06060820152600061283060a0830186886127c5565b82810360808401526128438185876127c5565b9998505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff80831681851680830382111561289e5761289e612850565b01949350505050565b6000828210156128b9576128b9612850565b500390565b60005b838110156128d95781810151838201526020016128c1565b838111156128e8576000848401525b50505050565b600081518084526129068160208601602086016128be565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b82815260406020820152600061101a60408301846128ee565b83815282602082015260606040820152600061297060608301846128ee565b95945050505050565b602081526000610acb60208301846128ee565b60006020828403121561299e57600080fd5b8151610acb816126cc565b60007fffffffff00000000000000000000000000000000000000000000000000000000808a60e01b168352886004840152808860e01b166024840152808760e01b16602884015285602c840152808560e01b16604c840152508251612a158160508501602087016128be565b9190910160500198975050505050505050565b600181815b80851115612a8157817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612a6757612a67612850565b80851615612a7457918102915b93841c9390800290612a2d565b509250929050565b600082612a98575060016105d8565b81612aa5575060006105d8565b8160018114612abb5760028114612ac557612ae1565b60019150506105d8565b60ff841115612ad657612ad6612850565b50506001821b6105d8565b5060208310610133831016604e8410600b8410161715612b04575081810a6105d8565b612b0e8383612a28565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612b4057612b40612850565b029392505050565b6000610acb8383612a89565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60008219821115612b9657612b96612850565b50019056fea264697066735822122063677658764ad613238e646022d9594e4df00325ef07be1057c032da5aa582a764736f6c634300080d0033",
}

// HomeABI is the input ABI used to generate the binding from.
// Deprecated: Use HomeMetaData.ABI instead.
var HomeABI = HomeMetaData.ABI

// Deprecated: Use HomeMetaData.Sigs instead.
// HomeFuncSigs maps the 4-byte function signature to its string representation.
var HomeFuncSigs = HomeMetaData.Sigs

// HomeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HomeMetaData.Bin instead.
var HomeBin = HomeMetaData.Bin

// DeployHome deploys a new Ethereum contract, binding an instance of Home to it.
func DeployHome(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32) (common.Address, *types.Transaction, *Home, error) {
	parsed, err := HomeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HomeBin), backend, _localDomain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Home{HomeCaller: HomeCaller{contract: contract}, HomeTransactor: HomeTransactor{contract: contract}, HomeFilterer: HomeFilterer{contract: contract}}, nil
}

// Home is an auto generated Go binding around an Ethereum contract.
type Home struct {
	HomeCaller     // Read-only binding to the contract
	HomeTransactor // Write-only binding to the contract
	HomeFilterer   // Log filterer for contract events
}

// HomeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HomeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HomeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HomeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HomeSession struct {
	Contract     *Home             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HomeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HomeCallerSession struct {
	Contract *HomeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HomeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HomeTransactorSession struct {
	Contract     *HomeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HomeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HomeRaw struct {
	Contract *Home // Generic contract binding to access the raw methods on
}

// HomeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HomeCallerRaw struct {
	Contract *HomeCaller // Generic read-only contract binding to access the raw methods on
}

// HomeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HomeTransactorRaw struct {
	Contract *HomeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHome creates a new instance of Home, bound to a specific deployed contract.
func NewHome(address common.Address, backend bind.ContractBackend) (*Home, error) {
	contract, err := bindHome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Home{HomeCaller: HomeCaller{contract: contract}, HomeTransactor: HomeTransactor{contract: contract}, HomeFilterer: HomeFilterer{contract: contract}}, nil
}

// NewHomeCaller creates a new read-only instance of Home, bound to a specific deployed contract.
func NewHomeCaller(address common.Address, caller bind.ContractCaller) (*HomeCaller, error) {
	contract, err := bindHome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HomeCaller{contract: contract}, nil
}

// NewHomeTransactor creates a new write-only instance of Home, bound to a specific deployed contract.
func NewHomeTransactor(address common.Address, transactor bind.ContractTransactor) (*HomeTransactor, error) {
	contract, err := bindHome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HomeTransactor{contract: contract}, nil
}

// NewHomeFilterer creates a new log filterer instance of Home, bound to a specific deployed contract.
func NewHomeFilterer(address common.Address, filterer bind.ContractFilterer) (*HomeFilterer, error) {
	contract, err := bindHome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HomeFilterer{contract: contract}, nil
}

// bindHome binds a generic wrapper to an already deployed contract.
func bindHome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HomeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Home *HomeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Home.Contract.HomeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Home *HomeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Home.Contract.HomeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Home *HomeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Home.Contract.HomeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Home *HomeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Home.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Home *HomeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Home.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Home *HomeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Home.Contract.contract.Transact(opts, method, params...)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Home *HomeCaller) MAXMESSAGEBODYBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "MAX_MESSAGE_BODY_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Home *HomeSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _Home.Contract.MAXMESSAGEBODYBYTES(&_Home.CallOpts)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Home *HomeCallerSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _Home.Contract.MAXMESSAGEBODYBYTES(&_Home.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Home *HomeCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Home *HomeSession) VERSION() (uint8, error) {
	return _Home.Contract.VERSION(&_Home.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Home *HomeCallerSession) VERSION() (uint8, error) {
	return _Home.Contract.VERSION(&_Home.CallOpts)
}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_Home *HomeCaller) CommittedRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "committedRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_Home *HomeSession) CommittedRoot() ([32]byte, error) {
	return _Home.Contract.CommittedRoot(&_Home.CallOpts)
}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_Home *HomeCallerSession) CommittedRoot() ([32]byte, error) {
	return _Home.Contract.CommittedRoot(&_Home.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Home *HomeCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Home *HomeSession) Count() (*big.Int, error) {
	return _Home.Contract.Count(&_Home.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Home *HomeCallerSession) Count() (*big.Int, error) {
	return _Home.Contract.Count(&_Home.CallOpts)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_Home *HomeCaller) HomeDomainHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "homeDomainHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_Home *HomeSession) HomeDomainHash() ([32]byte, error) {
	return _Home.Contract.HomeDomainHash(&_Home.CallOpts)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_Home *HomeCallerSession) HomeDomainHash() ([32]byte, error) {
	return _Home.Contract.HomeDomainHash(&_Home.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Home *HomeCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Home *HomeSession) LocalDomain() (uint32, error) {
	return _Home.Contract.LocalDomain(&_Home.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Home *HomeCallerSession) LocalDomain() (uint32, error) {
	return _Home.Contract.LocalDomain(&_Home.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0xb95a2001.
//
// Solidity: function nonces(uint32 ) view returns(uint32)
func (_Home *HomeCaller) Nonces(opts *bind.CallOpts, arg0 uint32) (uint32, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xb95a2001.
//
// Solidity: function nonces(uint32 ) view returns(uint32)
func (_Home *HomeSession) Nonces(arg0 uint32) (uint32, error) {
	return _Home.Contract.Nonces(&_Home.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0xb95a2001.
//
// Solidity: function nonces(uint32 ) view returns(uint32)
func (_Home *HomeCallerSession) Nonces(arg0 uint32) (uint32, error) {
	return _Home.Contract.Nonces(&_Home.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Home *HomeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Home *HomeSession) Owner() (common.Address, error) {
	return _Home.Contract.Owner(&_Home.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Home *HomeCallerSession) Owner() (common.Address, error) {
	return _Home.Contract.Owner(&_Home.CallOpts)
}

// QueueContains is a free data retrieval call binding the contract method 0x2bef2892.
//
// Solidity: function queueContains(bytes32 _item) view returns(bool)
func (_Home *HomeCaller) QueueContains(opts *bind.CallOpts, _item [32]byte) (bool, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "queueContains", _item)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// QueueContains is a free data retrieval call binding the contract method 0x2bef2892.
//
// Solidity: function queueContains(bytes32 _item) view returns(bool)
func (_Home *HomeSession) QueueContains(_item [32]byte) (bool, error) {
	return _Home.Contract.QueueContains(&_Home.CallOpts, _item)
}

// QueueContains is a free data retrieval call binding the contract method 0x2bef2892.
//
// Solidity: function queueContains(bytes32 _item) view returns(bool)
func (_Home *HomeCallerSession) QueueContains(_item [32]byte) (bool, error) {
	return _Home.Contract.QueueContains(&_Home.CallOpts, _item)
}

// QueueEnd is a free data retrieval call binding the contract method 0xf6d16102.
//
// Solidity: function queueEnd() view returns(bytes32)
func (_Home *HomeCaller) QueueEnd(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "queueEnd")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// QueueEnd is a free data retrieval call binding the contract method 0xf6d16102.
//
// Solidity: function queueEnd() view returns(bytes32)
func (_Home *HomeSession) QueueEnd() ([32]byte, error) {
	return _Home.Contract.QueueEnd(&_Home.CallOpts)
}

// QueueEnd is a free data retrieval call binding the contract method 0xf6d16102.
//
// Solidity: function queueEnd() view returns(bytes32)
func (_Home *HomeCallerSession) QueueEnd() ([32]byte, error) {
	return _Home.Contract.QueueEnd(&_Home.CallOpts)
}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_Home *HomeCaller) QueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "queueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_Home *HomeSession) QueueLength() (*big.Int, error) {
	return _Home.Contract.QueueLength(&_Home.CallOpts)
}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_Home *HomeCallerSession) QueueLength() (*big.Int, error) {
	return _Home.Contract.QueueLength(&_Home.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Home *HomeCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Home *HomeSession) Root() ([32]byte, error) {
	return _Home.Contract.Root(&_Home.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Home *HomeCallerSession) Root() ([32]byte, error) {
	return _Home.Contract.Root(&_Home.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Home *HomeCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Home *HomeSession) State() (uint8, error) {
	return _Home.Contract.State(&_Home.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Home *HomeCallerSession) State() (uint8, error) {
	return _Home.Contract.State(&_Home.CallOpts)
}

// SuggestUpdate is a free data retrieval call binding the contract method 0x36e104de.
//
// Solidity: function suggestUpdate() view returns(bytes32 _committedRoot, bytes32 _new)
func (_Home *HomeCaller) SuggestUpdate(opts *bind.CallOpts) (struct {
	CommittedRoot [32]byte
	New           [32]byte
}, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "suggestUpdate")

	outstruct := new(struct {
		CommittedRoot [32]byte
		New           [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CommittedRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.New = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SuggestUpdate is a free data retrieval call binding the contract method 0x36e104de.
//
// Solidity: function suggestUpdate() view returns(bytes32 _committedRoot, bytes32 _new)
func (_Home *HomeSession) SuggestUpdate() (struct {
	CommittedRoot [32]byte
	New           [32]byte
}, error) {
	return _Home.Contract.SuggestUpdate(&_Home.CallOpts)
}

// SuggestUpdate is a free data retrieval call binding the contract method 0x36e104de.
//
// Solidity: function suggestUpdate() view returns(bytes32 _committedRoot, bytes32 _new)
func (_Home *HomeCallerSession) SuggestUpdate() (struct {
	CommittedRoot [32]byte
	New           [32]byte
}, error) {
	return _Home.Contract.SuggestUpdate(&_Home.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Home *HomeCaller) Tree(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "tree")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Home *HomeSession) Tree() (*big.Int, error) {
	return _Home.Contract.Tree(&_Home.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Home *HomeCallerSession) Tree() (*big.Int, error) {
	return _Home.Contract.Tree(&_Home.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_Home *HomeCaller) Updater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "updater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_Home *HomeSession) Updater() (common.Address, error) {
	return _Home.Contract.Updater(&_Home.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_Home *HomeCallerSession) Updater() (common.Address, error) {
	return _Home.Contract.Updater(&_Home.CallOpts)
}

// UpdaterManager is a free data retrieval call binding the contract method 0x9df6c8e1.
//
// Solidity: function updaterManager() view returns(address)
func (_Home *HomeCaller) UpdaterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "updaterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpdaterManager is a free data retrieval call binding the contract method 0x9df6c8e1.
//
// Solidity: function updaterManager() view returns(address)
func (_Home *HomeSession) UpdaterManager() (common.Address, error) {
	return _Home.Contract.UpdaterManager(&_Home.CallOpts)
}

// UpdaterManager is a free data retrieval call binding the contract method 0x9df6c8e1.
//
// Solidity: function updaterManager() view returns(address)
func (_Home *HomeCallerSession) UpdaterManager() (common.Address, error) {
	return _Home.Contract.UpdaterManager(&_Home.CallOpts)
}

// Dispatch is a paid mutator transaction binding the contract method 0x59e62328.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _messageBody) returns()
func (_Home *HomeTransactor) Dispatch(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _messageBody []byte) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "dispatch", _destinationDomain, _recipientAddress, _optimisticSeconds, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0x59e62328.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _messageBody) returns()
func (_Home *HomeSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _messageBody []byte) (*types.Transaction, error) {
	return _Home.Contract.Dispatch(&_Home.TransactOpts, _destinationDomain, _recipientAddress, _optimisticSeconds, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0x59e62328.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _messageBody) returns()
func (_Home *HomeTransactorSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _messageBody []byte) (*types.Transaction, error) {
	return _Home.Contract.Dispatch(&_Home.TransactOpts, _destinationDomain, _recipientAddress, _optimisticSeconds, _messageBody)
}

// DoubleUpdate is a paid mutator transaction binding the contract method 0x19d9d21a.
//
// Solidity: function doubleUpdate(bytes32 _oldRoot, bytes32[2] _newRoot, bytes _signature, bytes _signature2) returns()
func (_Home *HomeTransactor) DoubleUpdate(opts *bind.TransactOpts, _oldRoot [32]byte, _newRoot [2][32]byte, _signature []byte, _signature2 []byte) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "doubleUpdate", _oldRoot, _newRoot, _signature, _signature2)
}

// DoubleUpdate is a paid mutator transaction binding the contract method 0x19d9d21a.
//
// Solidity: function doubleUpdate(bytes32 _oldRoot, bytes32[2] _newRoot, bytes _signature, bytes _signature2) returns()
func (_Home *HomeSession) DoubleUpdate(_oldRoot [32]byte, _newRoot [2][32]byte, _signature []byte, _signature2 []byte) (*types.Transaction, error) {
	return _Home.Contract.DoubleUpdate(&_Home.TransactOpts, _oldRoot, _newRoot, _signature, _signature2)
}

// DoubleUpdate is a paid mutator transaction binding the contract method 0x19d9d21a.
//
// Solidity: function doubleUpdate(bytes32 _oldRoot, bytes32[2] _newRoot, bytes _signature, bytes _signature2) returns()
func (_Home *HomeTransactorSession) DoubleUpdate(_oldRoot [32]byte, _newRoot [2][32]byte, _signature []byte, _signature2 []byte) (*types.Transaction, error) {
	return _Home.Contract.DoubleUpdate(&_Home.TransactOpts, _oldRoot, _newRoot, _signature, _signature2)
}

// ImproperUpdate is a paid mutator transaction binding the contract method 0x8e4e30e0.
//
// Solidity: function improperUpdate(bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns(bool)
func (_Home *HomeTransactor) ImproperUpdate(opts *bind.TransactOpts, _oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "improperUpdate", _oldRoot, _newRoot, _signature)
}

// ImproperUpdate is a paid mutator transaction binding the contract method 0x8e4e30e0.
//
// Solidity: function improperUpdate(bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns(bool)
func (_Home *HomeSession) ImproperUpdate(_oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Home.Contract.ImproperUpdate(&_Home.TransactOpts, _oldRoot, _newRoot, _signature)
}

// ImproperUpdate is a paid mutator transaction binding the contract method 0x8e4e30e0.
//
// Solidity: function improperUpdate(bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns(bool)
func (_Home *HomeTransactorSession) ImproperUpdate(_oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Home.Contract.ImproperUpdate(&_Home.TransactOpts, _oldRoot, _newRoot, _signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _updaterManager) returns()
func (_Home *HomeTransactor) Initialize(opts *bind.TransactOpts, _updaterManager common.Address) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "initialize", _updaterManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _updaterManager) returns()
func (_Home *HomeSession) Initialize(_updaterManager common.Address) (*types.Transaction, error) {
	return _Home.Contract.Initialize(&_Home.TransactOpts, _updaterManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _updaterManager) returns()
func (_Home *HomeTransactorSession) Initialize(_updaterManager common.Address) (*types.Transaction, error) {
	return _Home.Contract.Initialize(&_Home.TransactOpts, _updaterManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Home *HomeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Home *HomeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Home.Contract.RenounceOwnership(&_Home.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Home *HomeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Home.Contract.RenounceOwnership(&_Home.TransactOpts)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_Home *HomeTransactor) SetUpdater(opts *bind.TransactOpts, _updater common.Address) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "setUpdater", _updater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_Home *HomeSession) SetUpdater(_updater common.Address) (*types.Transaction, error) {
	return _Home.Contract.SetUpdater(&_Home.TransactOpts, _updater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_Home *HomeTransactorSession) SetUpdater(_updater common.Address) (*types.Transaction, error) {
	return _Home.Contract.SetUpdater(&_Home.TransactOpts, _updater)
}

// SetUpdaterManager is a paid mutator transaction binding the contract method 0x9776120e.
//
// Solidity: function setUpdaterManager(address _updaterManager) returns()
func (_Home *HomeTransactor) SetUpdaterManager(opts *bind.TransactOpts, _updaterManager common.Address) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "setUpdaterManager", _updaterManager)
}

// SetUpdaterManager is a paid mutator transaction binding the contract method 0x9776120e.
//
// Solidity: function setUpdaterManager(address _updaterManager) returns()
func (_Home *HomeSession) SetUpdaterManager(_updaterManager common.Address) (*types.Transaction, error) {
	return _Home.Contract.SetUpdaterManager(&_Home.TransactOpts, _updaterManager)
}

// SetUpdaterManager is a paid mutator transaction binding the contract method 0x9776120e.
//
// Solidity: function setUpdaterManager(address _updaterManager) returns()
func (_Home *HomeTransactorSession) SetUpdaterManager(_updaterManager common.Address) (*types.Transaction, error) {
	return _Home.Contract.SetUpdaterManager(&_Home.TransactOpts, _updaterManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Home *HomeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Home *HomeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Home.Contract.TransferOwnership(&_Home.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Home *HomeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Home.Contract.TransferOwnership(&_Home.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0xb31c01fb.
//
// Solidity: function update(bytes32 _committedRoot, bytes32 _newRoot, bytes _signature) returns()
func (_Home *HomeTransactor) Update(opts *bind.TransactOpts, _committedRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "update", _committedRoot, _newRoot, _signature)
}

// Update is a paid mutator transaction binding the contract method 0xb31c01fb.
//
// Solidity: function update(bytes32 _committedRoot, bytes32 _newRoot, bytes _signature) returns()
func (_Home *HomeSession) Update(_committedRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Home.Contract.Update(&_Home.TransactOpts, _committedRoot, _newRoot, _signature)
}

// Update is a paid mutator transaction binding the contract method 0xb31c01fb.
//
// Solidity: function update(bytes32 _committedRoot, bytes32 _newRoot, bytes _signature) returns()
func (_Home *HomeTransactorSession) Update(_committedRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Home.Contract.Update(&_Home.TransactOpts, _committedRoot, _newRoot, _signature)
}

// HomeDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the Home contract.
type HomeDispatchIterator struct {
	Event *HomeDispatch // Event containing the contract specifics and raw log

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
func (it *HomeDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeDispatch)
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
		it.Event = new(HomeDispatch)
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
func (it *HomeDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeDispatch represents a Dispatch event raised by the Home contract.
type HomeDispatch struct {
	MessageHash         [32]byte
	LeafIndex           *big.Int
	DestinationAndNonce uint64
	CommittedRoot       [32]byte
	Message             []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0x9d4c83d2e57d7d381feb264b44a5015e7f9ef26340f4fc46b558a6dc16dd811a.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes32 committedRoot, bytes message)
func (_Home *HomeFilterer) FilterDispatch(opts *bind.FilterOpts, messageHash [][32]byte, leafIndex []*big.Int, destinationAndNonce []uint64) (*HomeDispatchIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var leafIndexRule []interface{}
	for _, leafIndexItem := range leafIndex {
		leafIndexRule = append(leafIndexRule, leafIndexItem)
	}
	var destinationAndNonceRule []interface{}
	for _, destinationAndNonceItem := range destinationAndNonce {
		destinationAndNonceRule = append(destinationAndNonceRule, destinationAndNonceItem)
	}

	logs, sub, err := _Home.contract.FilterLogs(opts, "Dispatch", messageHashRule, leafIndexRule, destinationAndNonceRule)
	if err != nil {
		return nil, err
	}
	return &HomeDispatchIterator{contract: _Home.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0x9d4c83d2e57d7d381feb264b44a5015e7f9ef26340f4fc46b558a6dc16dd811a.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes32 committedRoot, bytes message)
func (_Home *HomeFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *HomeDispatch, messageHash [][32]byte, leafIndex []*big.Int, destinationAndNonce []uint64) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var leafIndexRule []interface{}
	for _, leafIndexItem := range leafIndex {
		leafIndexRule = append(leafIndexRule, leafIndexItem)
	}
	var destinationAndNonceRule []interface{}
	for _, destinationAndNonceItem := range destinationAndNonce {
		destinationAndNonceRule = append(destinationAndNonceRule, destinationAndNonceItem)
	}

	logs, sub, err := _Home.contract.WatchLogs(opts, "Dispatch", messageHashRule, leafIndexRule, destinationAndNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeDispatch)
				if err := _Home.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0x9d4c83d2e57d7d381feb264b44a5015e7f9ef26340f4fc46b558a6dc16dd811a.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes32 committedRoot, bytes message)
func (_Home *HomeFilterer) ParseDispatch(log types.Log) (*HomeDispatch, error) {
	event := new(HomeDispatch)
	if err := _Home.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeDoubleUpdateIterator is returned from FilterDoubleUpdate and is used to iterate over the raw logs and unpacked data for DoubleUpdate events raised by the Home contract.
type HomeDoubleUpdateIterator struct {
	Event *HomeDoubleUpdate // Event containing the contract specifics and raw log

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
func (it *HomeDoubleUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeDoubleUpdate)
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
		it.Event = new(HomeDoubleUpdate)
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
func (it *HomeDoubleUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeDoubleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeDoubleUpdate represents a DoubleUpdate event raised by the Home contract.
type HomeDoubleUpdate struct {
	OldRoot    [32]byte
	NewRoot    [2][32]byte
	Signature  []byte
	Signature2 []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDoubleUpdate is a free log retrieval operation binding the contract event 0x2c3f60bab4170347826231b75a920b5053941ddebc6eed6fd2c25721648b186f.
//
// Solidity: event DoubleUpdate(bytes32 oldRoot, bytes32[2] newRoot, bytes signature, bytes signature2)
func (_Home *HomeFilterer) FilterDoubleUpdate(opts *bind.FilterOpts) (*HomeDoubleUpdateIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "DoubleUpdate")
	if err != nil {
		return nil, err
	}
	return &HomeDoubleUpdateIterator{contract: _Home.contract, event: "DoubleUpdate", logs: logs, sub: sub}, nil
}

// WatchDoubleUpdate is a free log subscription operation binding the contract event 0x2c3f60bab4170347826231b75a920b5053941ddebc6eed6fd2c25721648b186f.
//
// Solidity: event DoubleUpdate(bytes32 oldRoot, bytes32[2] newRoot, bytes signature, bytes signature2)
func (_Home *HomeFilterer) WatchDoubleUpdate(opts *bind.WatchOpts, sink chan<- *HomeDoubleUpdate) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "DoubleUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeDoubleUpdate)
				if err := _Home.contract.UnpackLog(event, "DoubleUpdate", log); err != nil {
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

// ParseDoubleUpdate is a log parse operation binding the contract event 0x2c3f60bab4170347826231b75a920b5053941ddebc6eed6fd2c25721648b186f.
//
// Solidity: event DoubleUpdate(bytes32 oldRoot, bytes32[2] newRoot, bytes signature, bytes signature2)
func (_Home *HomeFilterer) ParseDoubleUpdate(log types.Log) (*HomeDoubleUpdate, error) {
	event := new(HomeDoubleUpdate)
	if err := _Home.contract.UnpackLog(event, "DoubleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeImproperUpdateIterator is returned from FilterImproperUpdate and is used to iterate over the raw logs and unpacked data for ImproperUpdate events raised by the Home contract.
type HomeImproperUpdateIterator struct {
	Event *HomeImproperUpdate // Event containing the contract specifics and raw log

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
func (it *HomeImproperUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeImproperUpdate)
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
		it.Event = new(HomeImproperUpdate)
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
func (it *HomeImproperUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeImproperUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeImproperUpdate represents a ImproperUpdate event raised by the Home contract.
type HomeImproperUpdate struct {
	OldRoot   [32]byte
	NewRoot   [32]byte
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterImproperUpdate is a free log retrieval operation binding the contract event 0x6844fd5e21c932b5197b78ac11bf96e2eaa4e882dd0c88087060cf2065c04ab2.
//
// Solidity: event ImproperUpdate(bytes32 oldRoot, bytes32 newRoot, bytes signature)
func (_Home *HomeFilterer) FilterImproperUpdate(opts *bind.FilterOpts) (*HomeImproperUpdateIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "ImproperUpdate")
	if err != nil {
		return nil, err
	}
	return &HomeImproperUpdateIterator{contract: _Home.contract, event: "ImproperUpdate", logs: logs, sub: sub}, nil
}

// WatchImproperUpdate is a free log subscription operation binding the contract event 0x6844fd5e21c932b5197b78ac11bf96e2eaa4e882dd0c88087060cf2065c04ab2.
//
// Solidity: event ImproperUpdate(bytes32 oldRoot, bytes32 newRoot, bytes signature)
func (_Home *HomeFilterer) WatchImproperUpdate(opts *bind.WatchOpts, sink chan<- *HomeImproperUpdate) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "ImproperUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeImproperUpdate)
				if err := _Home.contract.UnpackLog(event, "ImproperUpdate", log); err != nil {
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

// ParseImproperUpdate is a log parse operation binding the contract event 0x6844fd5e21c932b5197b78ac11bf96e2eaa4e882dd0c88087060cf2065c04ab2.
//
// Solidity: event ImproperUpdate(bytes32 oldRoot, bytes32 newRoot, bytes signature)
func (_Home *HomeFilterer) ParseImproperUpdate(log types.Log) (*HomeImproperUpdate, error) {
	event := new(HomeImproperUpdate)
	if err := _Home.contract.UnpackLog(event, "ImproperUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Home contract.
type HomeInitializedIterator struct {
	Event *HomeInitialized // Event containing the contract specifics and raw log

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
func (it *HomeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeInitialized)
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
		it.Event = new(HomeInitialized)
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
func (it *HomeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeInitialized represents a Initialized event raised by the Home contract.
type HomeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Home *HomeFilterer) FilterInitialized(opts *bind.FilterOpts) (*HomeInitializedIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HomeInitializedIterator{contract: _Home.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Home *HomeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HomeInitialized) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeInitialized)
				if err := _Home.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Home *HomeFilterer) ParseInitialized(log types.Log) (*HomeInitialized, error) {
	event := new(HomeInitialized)
	if err := _Home.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeNewUpdaterIterator is returned from FilterNewUpdater and is used to iterate over the raw logs and unpacked data for NewUpdater events raised by the Home contract.
type HomeNewUpdaterIterator struct {
	Event *HomeNewUpdater // Event containing the contract specifics and raw log

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
func (it *HomeNewUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeNewUpdater)
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
		it.Event = new(HomeNewUpdater)
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
func (it *HomeNewUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeNewUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeNewUpdater represents a NewUpdater event raised by the Home contract.
type HomeNewUpdater struct {
	OldUpdater common.Address
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUpdater is a free log retrieval operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_Home *HomeFilterer) FilterNewUpdater(opts *bind.FilterOpts) (*HomeNewUpdaterIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return &HomeNewUpdaterIterator{contract: _Home.contract, event: "NewUpdater", logs: logs, sub: sub}, nil
}

// WatchNewUpdater is a free log subscription operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_Home *HomeFilterer) WatchNewUpdater(opts *bind.WatchOpts, sink chan<- *HomeNewUpdater) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeNewUpdater)
				if err := _Home.contract.UnpackLog(event, "NewUpdater", log); err != nil {
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

// ParseNewUpdater is a log parse operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_Home *HomeFilterer) ParseNewUpdater(log types.Log) (*HomeNewUpdater, error) {
	event := new(HomeNewUpdater)
	if err := _Home.contract.UnpackLog(event, "NewUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeNewUpdaterManagerIterator is returned from FilterNewUpdaterManager and is used to iterate over the raw logs and unpacked data for NewUpdaterManager events raised by the Home contract.
type HomeNewUpdaterManagerIterator struct {
	Event *HomeNewUpdaterManager // Event containing the contract specifics and raw log

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
func (it *HomeNewUpdaterManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeNewUpdaterManager)
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
		it.Event = new(HomeNewUpdaterManager)
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
func (it *HomeNewUpdaterManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeNewUpdaterManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeNewUpdaterManager represents a NewUpdaterManager event raised by the Home contract.
type HomeNewUpdaterManager struct {
	UpdaterManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewUpdaterManager is a free log retrieval operation binding the contract event 0x958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf.
//
// Solidity: event NewUpdaterManager(address updaterManager)
func (_Home *HomeFilterer) FilterNewUpdaterManager(opts *bind.FilterOpts) (*HomeNewUpdaterManagerIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "NewUpdaterManager")
	if err != nil {
		return nil, err
	}
	return &HomeNewUpdaterManagerIterator{contract: _Home.contract, event: "NewUpdaterManager", logs: logs, sub: sub}, nil
}

// WatchNewUpdaterManager is a free log subscription operation binding the contract event 0x958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf.
//
// Solidity: event NewUpdaterManager(address updaterManager)
func (_Home *HomeFilterer) WatchNewUpdaterManager(opts *bind.WatchOpts, sink chan<- *HomeNewUpdaterManager) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "NewUpdaterManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeNewUpdaterManager)
				if err := _Home.contract.UnpackLog(event, "NewUpdaterManager", log); err != nil {
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

// ParseNewUpdaterManager is a log parse operation binding the contract event 0x958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf.
//
// Solidity: event NewUpdaterManager(address updaterManager)
func (_Home *HomeFilterer) ParseNewUpdaterManager(log types.Log) (*HomeNewUpdaterManager, error) {
	event := new(HomeNewUpdaterManager)
	if err := _Home.contract.UnpackLog(event, "NewUpdaterManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Home contract.
type HomeOwnershipTransferredIterator struct {
	Event *HomeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HomeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeOwnershipTransferred)
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
		it.Event = new(HomeOwnershipTransferred)
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
func (it *HomeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeOwnershipTransferred represents a OwnershipTransferred event raised by the Home contract.
type HomeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Home *HomeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HomeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Home.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HomeOwnershipTransferredIterator{contract: _Home.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Home *HomeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HomeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Home.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeOwnershipTransferred)
				if err := _Home.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Home *HomeFilterer) ParseOwnershipTransferred(log types.Log) (*HomeOwnershipTransferred, error) {
	event := new(HomeOwnershipTransferred)
	if err := _Home.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the Home contract.
type HomeUpdateIterator struct {
	Event *HomeUpdate // Event containing the contract specifics and raw log

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
func (it *HomeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeUpdate)
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
		it.Event = new(HomeUpdate)
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
func (it *HomeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeUpdate represents a Update event raised by the Home contract.
type HomeUpdate struct {
	HomeDomain uint32
	OldRoot    [32]byte
	NewRoot    [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_Home *HomeFilterer) FilterUpdate(opts *bind.FilterOpts, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (*HomeUpdateIterator, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var oldRootRule []interface{}
	for _, oldRootItem := range oldRoot {
		oldRootRule = append(oldRootRule, oldRootItem)
	}
	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}

	logs, sub, err := _Home.contract.FilterLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return &HomeUpdateIterator{contract: _Home.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_Home *HomeFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *HomeUpdate, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (event.Subscription, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var oldRootRule []interface{}
	for _, oldRootItem := range oldRoot {
		oldRootRule = append(oldRootRule, oldRootItem)
	}
	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}

	logs, sub, err := _Home.contract.WatchLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeUpdate)
				if err := _Home.contract.UnpackLog(event, "Update", log); err != nil {
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

// ParseUpdate is a log parse operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_Home *HomeFilterer) ParseUpdate(log types.Log) (*HomeUpdate, error) {
	event := new(HomeUpdate)
	if err := _Home.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeUpdaterSlashedIterator is returned from FilterUpdaterSlashed and is used to iterate over the raw logs and unpacked data for UpdaterSlashed events raised by the Home contract.
type HomeUpdaterSlashedIterator struct {
	Event *HomeUpdaterSlashed // Event containing the contract specifics and raw log

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
func (it *HomeUpdaterSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeUpdaterSlashed)
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
		it.Event = new(HomeUpdaterSlashed)
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
func (it *HomeUpdaterSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeUpdaterSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeUpdaterSlashed represents a UpdaterSlashed event raised by the Home contract.
type HomeUpdaterSlashed struct {
	Updater  common.Address
	Reporter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdaterSlashed is a free log retrieval operation binding the contract event 0x98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce.
//
// Solidity: event UpdaterSlashed(address indexed updater, address indexed reporter)
func (_Home *HomeFilterer) FilterUpdaterSlashed(opts *bind.FilterOpts, updater []common.Address, reporter []common.Address) (*HomeUpdaterSlashedIterator, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Home.contract.FilterLogs(opts, "UpdaterSlashed", updaterRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return &HomeUpdaterSlashedIterator{contract: _Home.contract, event: "UpdaterSlashed", logs: logs, sub: sub}, nil
}

// WatchUpdaterSlashed is a free log subscription operation binding the contract event 0x98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce.
//
// Solidity: event UpdaterSlashed(address indexed updater, address indexed reporter)
func (_Home *HomeFilterer) WatchUpdaterSlashed(opts *bind.WatchOpts, sink chan<- *HomeUpdaterSlashed, updater []common.Address, reporter []common.Address) (event.Subscription, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Home.contract.WatchLogs(opts, "UpdaterSlashed", updaterRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeUpdaterSlashed)
				if err := _Home.contract.UnpackLog(event, "UpdaterSlashed", log); err != nil {
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

// ParseUpdaterSlashed is a log parse operation binding the contract event 0x98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce.
//
// Solidity: event UpdaterSlashed(address indexed updater, address indexed reporter)
func (_Home *HomeFilterer) ParseUpdaterSlashed(log types.Log) (*HomeUpdaterSlashed, error) {
	event := new(HomeUpdaterSlashed)
	if err := _Home.contract.UnpackLog(event, "UpdaterSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMessageRecipientMetaData contains all meta data concerning the IMessageRecipient contract.
var IMessageRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_rootTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e4d16d62": "handle(uint32,uint32,bytes32,uint256,bytes)",
	},
}

// IMessageRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageRecipientMetaData.ABI instead.
var IMessageRecipientABI = IMessageRecipientMetaData.ABI

// Deprecated: Use IMessageRecipientMetaData.Sigs instead.
// IMessageRecipientFuncSigs maps the 4-byte function signature to its string representation.
var IMessageRecipientFuncSigs = IMessageRecipientMetaData.Sigs

// IMessageRecipient is an auto generated Go binding around an Ethereum contract.
type IMessageRecipient struct {
	IMessageRecipientCaller     // Read-only binding to the contract
	IMessageRecipientTransactor // Write-only binding to the contract
	IMessageRecipientFilterer   // Log filterer for contract events
}

// IMessageRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageRecipientSession struct {
	Contract     *IMessageRecipient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IMessageRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageRecipientCallerSession struct {
	Contract *IMessageRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IMessageRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageRecipientTransactorSession struct {
	Contract     *IMessageRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IMessageRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageRecipientRaw struct {
	Contract *IMessageRecipient // Generic contract binding to access the raw methods on
}

// IMessageRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageRecipientCallerRaw struct {
	Contract *IMessageRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageRecipientTransactorRaw struct {
	Contract *IMessageRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageRecipient creates a new instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipient(address common.Address, backend bind.ContractBackend) (*IMessageRecipient, error) {
	contract, err := bindIMessageRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipient{IMessageRecipientCaller: IMessageRecipientCaller{contract: contract}, IMessageRecipientTransactor: IMessageRecipientTransactor{contract: contract}, IMessageRecipientFilterer: IMessageRecipientFilterer{contract: contract}}, nil
}

// NewIMessageRecipientCaller creates a new read-only instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientCaller(address common.Address, caller bind.ContractCaller) (*IMessageRecipientCaller, error) {
	contract, err := bindIMessageRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientCaller{contract: contract}, nil
}

// NewIMessageRecipientTransactor creates a new write-only instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageRecipientTransactor, error) {
	contract, err := bindIMessageRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientTransactor{contract: contract}, nil
}

// NewIMessageRecipientFilterer creates a new log filterer instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageRecipientFilterer, error) {
	contract, err := bindIMessageRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientFilterer{contract: contract}, nil
}

// bindIMessageRecipient binds a generic wrapper to an already deployed contract.
func bindIMessageRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageRecipient *IMessageRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageRecipient.Contract.IMessageRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageRecipient *IMessageRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.IMessageRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageRecipient *IMessageRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.IMessageRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageRecipient *IMessageRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageRecipient *IMessageRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageRecipient *IMessageRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.contract.Transact(opts, method, params...)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, uint256 _rootTimestamp, bytes _message) returns()
func (_IMessageRecipient *IMessageRecipientTransactor) Handle(opts *bind.TransactOpts, _origin uint32, _nonce uint32, _sender [32]byte, _rootTimestamp *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageRecipient.contract.Transact(opts, "handle", _origin, _nonce, _sender, _rootTimestamp, _message)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, uint256 _rootTimestamp, bytes _message) returns()
func (_IMessageRecipient *IMessageRecipientSession) Handle(_origin uint32, _nonce uint32, _sender [32]byte, _rootTimestamp *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.Handle(&_IMessageRecipient.TransactOpts, _origin, _nonce, _sender, _rootTimestamp, _message)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, uint256 _rootTimestamp, bytes _message) returns()
func (_IMessageRecipient *IMessageRecipientTransactorSession) Handle(_origin uint32, _nonce uint32, _sender [32]byte, _rootTimestamp *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.Handle(&_IMessageRecipient.TransactOpts, _origin, _nonce, _sender, _rootTimestamp, _message)
}

// IUpdaterManagerMetaData contains all meta data concerning the IUpdaterManager contract.
var IUpdaterManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_reporter\",\"type\":\"address\"}],\"name\":\"slashUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5b3c2cbf": "slashUpdater(address)",
		"df034cd0": "updater()",
	},
}

// IUpdaterManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IUpdaterManagerMetaData.ABI instead.
var IUpdaterManagerABI = IUpdaterManagerMetaData.ABI

// Deprecated: Use IUpdaterManagerMetaData.Sigs instead.
// IUpdaterManagerFuncSigs maps the 4-byte function signature to its string representation.
var IUpdaterManagerFuncSigs = IUpdaterManagerMetaData.Sigs

// IUpdaterManager is an auto generated Go binding around an Ethereum contract.
type IUpdaterManager struct {
	IUpdaterManagerCaller     // Read-only binding to the contract
	IUpdaterManagerTransactor // Write-only binding to the contract
	IUpdaterManagerFilterer   // Log filterer for contract events
}

// IUpdaterManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUpdaterManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpdaterManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUpdaterManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpdaterManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUpdaterManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpdaterManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUpdaterManagerSession struct {
	Contract     *IUpdaterManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUpdaterManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUpdaterManagerCallerSession struct {
	Contract *IUpdaterManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IUpdaterManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUpdaterManagerTransactorSession struct {
	Contract     *IUpdaterManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IUpdaterManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUpdaterManagerRaw struct {
	Contract *IUpdaterManager // Generic contract binding to access the raw methods on
}

// IUpdaterManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUpdaterManagerCallerRaw struct {
	Contract *IUpdaterManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IUpdaterManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUpdaterManagerTransactorRaw struct {
	Contract *IUpdaterManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUpdaterManager creates a new instance of IUpdaterManager, bound to a specific deployed contract.
func NewIUpdaterManager(address common.Address, backend bind.ContractBackend) (*IUpdaterManager, error) {
	contract, err := bindIUpdaterManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUpdaterManager{IUpdaterManagerCaller: IUpdaterManagerCaller{contract: contract}, IUpdaterManagerTransactor: IUpdaterManagerTransactor{contract: contract}, IUpdaterManagerFilterer: IUpdaterManagerFilterer{contract: contract}}, nil
}

// NewIUpdaterManagerCaller creates a new read-only instance of IUpdaterManager, bound to a specific deployed contract.
func NewIUpdaterManagerCaller(address common.Address, caller bind.ContractCaller) (*IUpdaterManagerCaller, error) {
	contract, err := bindIUpdaterManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUpdaterManagerCaller{contract: contract}, nil
}

// NewIUpdaterManagerTransactor creates a new write-only instance of IUpdaterManager, bound to a specific deployed contract.
func NewIUpdaterManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IUpdaterManagerTransactor, error) {
	contract, err := bindIUpdaterManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUpdaterManagerTransactor{contract: contract}, nil
}

// NewIUpdaterManagerFilterer creates a new log filterer instance of IUpdaterManager, bound to a specific deployed contract.
func NewIUpdaterManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IUpdaterManagerFilterer, error) {
	contract, err := bindIUpdaterManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUpdaterManagerFilterer{contract: contract}, nil
}

// bindIUpdaterManager binds a generic wrapper to an already deployed contract.
func bindIUpdaterManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUpdaterManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpdaterManager *IUpdaterManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUpdaterManager.Contract.IUpdaterManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpdaterManager *IUpdaterManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpdaterManager.Contract.IUpdaterManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpdaterManager *IUpdaterManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpdaterManager.Contract.IUpdaterManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpdaterManager *IUpdaterManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUpdaterManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpdaterManager *IUpdaterManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpdaterManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpdaterManager *IUpdaterManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpdaterManager.Contract.contract.Transact(opts, method, params...)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_IUpdaterManager *IUpdaterManagerCaller) Updater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IUpdaterManager.contract.Call(opts, &out, "updater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_IUpdaterManager *IUpdaterManagerSession) Updater() (common.Address, error) {
	return _IUpdaterManager.Contract.Updater(&_IUpdaterManager.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_IUpdaterManager *IUpdaterManagerCallerSession) Updater() (common.Address, error) {
	return _IUpdaterManager.Contract.Updater(&_IUpdaterManager.CallOpts)
}

// SlashUpdater is a paid mutator transaction binding the contract method 0x5b3c2cbf.
//
// Solidity: function slashUpdater(address _reporter) returns()
func (_IUpdaterManager *IUpdaterManagerTransactor) SlashUpdater(opts *bind.TransactOpts, _reporter common.Address) (*types.Transaction, error) {
	return _IUpdaterManager.contract.Transact(opts, "slashUpdater", _reporter)
}

// SlashUpdater is a paid mutator transaction binding the contract method 0x5b3c2cbf.
//
// Solidity: function slashUpdater(address _reporter) returns()
func (_IUpdaterManager *IUpdaterManagerSession) SlashUpdater(_reporter common.Address) (*types.Transaction, error) {
	return _IUpdaterManager.Contract.SlashUpdater(&_IUpdaterManager.TransactOpts, _reporter)
}

// SlashUpdater is a paid mutator transaction binding the contract method 0x5b3c2cbf.
//
// Solidity: function slashUpdater(address _reporter) returns()
func (_IUpdaterManager *IUpdaterManagerTransactorSession) SlashUpdater(_reporter common.Address) (*types.Transaction, error) {
	return _IUpdaterManager.Contract.SlashUpdater(&_IUpdaterManager.TransactOpts, _reporter)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

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
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
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
		it.Event = new(InitializableInitialized)
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
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {

	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {

	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleLibMetaData contains all meta data concerning the MerkleLib contract.
var MerkleLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122022ec91b921e4f19ba968cb989907b59d34b76cf2e46174209f28d1e89307a10464736f6c634300080d0033",
}

// MerkleLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleLibMetaData.ABI instead.
var MerkleLibABI = MerkleLibMetaData.ABI

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleLibMetaData.Bin instead.
var MerkleLibBin = MerkleLibMetaData.Bin

// DeployMerkleLib deploys a new Ethereum contract, binding an instance of MerkleLib to it.
func DeployMerkleLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleLib, error) {
	parsed, err := MerkleLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// MerkleLib is an auto generated Go binding around an Ethereum contract.
type MerkleLib struct {
	MerkleLibCaller     // Read-only binding to the contract
	MerkleLibTransactor // Write-only binding to the contract
	MerkleLibFilterer   // Log filterer for contract events
}

// MerkleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleLibSession struct {
	Contract     *MerkleLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleLibCallerSession struct {
	Contract *MerkleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MerkleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleLibTransactorSession struct {
	Contract     *MerkleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MerkleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleLibRaw struct {
	Contract *MerkleLib // Generic contract binding to access the raw methods on
}

// MerkleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleLibCallerRaw struct {
	Contract *MerkleLibCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleLibTransactorRaw struct {
	Contract *MerkleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleLib creates a new instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLib(address common.Address, backend bind.ContractBackend) (*MerkleLib, error) {
	contract, err := bindMerkleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// NewMerkleLibCaller creates a new read-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibCaller(address common.Address, caller bind.ContractCaller) (*MerkleLibCaller, error) {
	contract, err := bindMerkleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibCaller{contract: contract}, nil
}

// NewMerkleLibTransactor creates a new write-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleLibTransactor, error) {
	contract, err := bindMerkleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibTransactor{contract: contract}, nil
}

// NewMerkleLibFilterer creates a new log filterer instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleLibFilterer, error) {
	contract, err := bindMerkleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleLibFilterer{contract: contract}, nil
}

// bindMerkleLib binds a generic wrapper to an already deployed contract.
func bindMerkleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.MerkleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transact(opts, method, params...)
}

// MerkleTreeManagerMetaData contains all meta data concerning the MerkleTreeManager contract.
var MerkleTreeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"06661abd": "count()",
		"ebf0c717": "root()",
		"fd54b228": "tree()",
	},
	Bin: "0x608060405234801561001057600080fd5b506106a4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806306661abd14610046578063ebf0c7171461005c578063fd54b22814610064575b600080fd5b6020545b60405190815260200160405180910390f35b61004a61006e565b60205461004a9081565b600061007a600061007f565b905090565b60006100928261008d610098565b610559565b92915050565b6100a0610620565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b602081101561061857600182821c1660008683602081106105865761058661063f565b01549050816001036105c357604080516020810183905290810186905260600160405160208183030381529060405280519060200120945061060e565b848684602081106105d6576105d661063f565b60200201516040516020016105f5929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b5050600101610563565b505092915050565b6040518061040001604052806020906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220c4381f9d3b4754b34ce1dbe68a3b00e518d704e6c3f02f08d611fc409b532ecb64736f6c634300080d0033",
}

// MerkleTreeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTreeManagerMetaData.ABI instead.
var MerkleTreeManagerABI = MerkleTreeManagerMetaData.ABI

// Deprecated: Use MerkleTreeManagerMetaData.Sigs instead.
// MerkleTreeManagerFuncSigs maps the 4-byte function signature to its string representation.
var MerkleTreeManagerFuncSigs = MerkleTreeManagerMetaData.Sigs

// MerkleTreeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleTreeManagerMetaData.Bin instead.
var MerkleTreeManagerBin = MerkleTreeManagerMetaData.Bin

// DeployMerkleTreeManager deploys a new Ethereum contract, binding an instance of MerkleTreeManager to it.
func DeployMerkleTreeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleTreeManager, error) {
	parsed, err := MerkleTreeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleTreeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleTreeManager{MerkleTreeManagerCaller: MerkleTreeManagerCaller{contract: contract}, MerkleTreeManagerTransactor: MerkleTreeManagerTransactor{contract: contract}, MerkleTreeManagerFilterer: MerkleTreeManagerFilterer{contract: contract}}, nil
}

// MerkleTreeManager is an auto generated Go binding around an Ethereum contract.
type MerkleTreeManager struct {
	MerkleTreeManagerCaller     // Read-only binding to the contract
	MerkleTreeManagerTransactor // Write-only binding to the contract
	MerkleTreeManagerFilterer   // Log filterer for contract events
}

// MerkleTreeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTreeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTreeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTreeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTreeManagerSession struct {
	Contract     *MerkleTreeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MerkleTreeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTreeManagerCallerSession struct {
	Contract *MerkleTreeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MerkleTreeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTreeManagerTransactorSession struct {
	Contract     *MerkleTreeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MerkleTreeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTreeManagerRaw struct {
	Contract *MerkleTreeManager // Generic contract binding to access the raw methods on
}

// MerkleTreeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTreeManagerCallerRaw struct {
	Contract *MerkleTreeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTreeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTreeManagerTransactorRaw struct {
	Contract *MerkleTreeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTreeManager creates a new instance of MerkleTreeManager, bound to a specific deployed contract.
func NewMerkleTreeManager(address common.Address, backend bind.ContractBackend) (*MerkleTreeManager, error) {
	contract, err := bindMerkleTreeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeManager{MerkleTreeManagerCaller: MerkleTreeManagerCaller{contract: contract}, MerkleTreeManagerTransactor: MerkleTreeManagerTransactor{contract: contract}, MerkleTreeManagerFilterer: MerkleTreeManagerFilterer{contract: contract}}, nil
}

// NewMerkleTreeManagerCaller creates a new read-only instance of MerkleTreeManager, bound to a specific deployed contract.
func NewMerkleTreeManagerCaller(address common.Address, caller bind.ContractCaller) (*MerkleTreeManagerCaller, error) {
	contract, err := bindMerkleTreeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeManagerCaller{contract: contract}, nil
}

// NewMerkleTreeManagerTransactor creates a new write-only instance of MerkleTreeManager, bound to a specific deployed contract.
func NewMerkleTreeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTreeManagerTransactor, error) {
	contract, err := bindMerkleTreeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeManagerTransactor{contract: contract}, nil
}

// NewMerkleTreeManagerFilterer creates a new log filterer instance of MerkleTreeManager, bound to a specific deployed contract.
func NewMerkleTreeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTreeManagerFilterer, error) {
	contract, err := bindMerkleTreeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeManagerFilterer{contract: contract}, nil
}

// bindMerkleTreeManager binds a generic wrapper to an already deployed contract.
func bindMerkleTreeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleTreeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTreeManager *MerkleTreeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTreeManager.Contract.MerkleTreeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTreeManager *MerkleTreeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeManager.Contract.MerkleTreeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTreeManager *MerkleTreeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTreeManager.Contract.MerkleTreeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTreeManager *MerkleTreeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTreeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTreeManager *MerkleTreeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTreeManager *MerkleTreeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTreeManager.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MerkleTreeManager *MerkleTreeManagerCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeManager.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MerkleTreeManager *MerkleTreeManagerSession) Count() (*big.Int, error) {
	return _MerkleTreeManager.Contract.Count(&_MerkleTreeManager.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MerkleTreeManager *MerkleTreeManagerCallerSession) Count() (*big.Int, error) {
	return _MerkleTreeManager.Contract.Count(&_MerkleTreeManager.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_MerkleTreeManager *MerkleTreeManagerCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MerkleTreeManager.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_MerkleTreeManager *MerkleTreeManagerSession) Root() ([32]byte, error) {
	return _MerkleTreeManager.Contract.Root(&_MerkleTreeManager.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_MerkleTreeManager *MerkleTreeManagerCallerSession) Root() ([32]byte, error) {
	return _MerkleTreeManager.Contract.Root(&_MerkleTreeManager.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_MerkleTreeManager *MerkleTreeManagerCaller) Tree(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeManager.contract.Call(opts, &out, "tree")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_MerkleTreeManager *MerkleTreeManagerSession) Tree() (*big.Int, error) {
	return _MerkleTreeManager.Contract.Tree(&_MerkleTreeManager.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_MerkleTreeManager *MerkleTreeManagerCallerSession) Tree() (*big.Int, error) {
	return _MerkleTreeManager.Contract.Tree(&_MerkleTreeManager.CallOpts)
}

// MessageMetaData contains all meta data concerning the Message contract.
var MessageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203c21cedd68aa6bad8396ccf24b5f69a34efb4c9244b73ef8a01cae6e8a52dfa064736f6c634300080d0033",
}

// MessageABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageMetaData.ABI instead.
var MessageABI = MessageMetaData.ABI

// MessageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageMetaData.Bin instead.
var MessageBin = MessageMetaData.Bin

// DeployMessage deploys a new Ethereum contract, binding an instance of Message to it.
func DeployMessage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Message, error) {
	parsed, err := MessageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Message{MessageCaller: MessageCaller{contract: contract}, MessageTransactor: MessageTransactor{contract: contract}, MessageFilterer: MessageFilterer{contract: contract}}, nil
}

// Message is an auto generated Go binding around an Ethereum contract.
type Message struct {
	MessageCaller     // Read-only binding to the contract
	MessageTransactor // Write-only binding to the contract
	MessageFilterer   // Log filterer for contract events
}

// MessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageSession struct {
	Contract     *Message          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageCallerSession struct {
	Contract *MessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTransactorSession struct {
	Contract     *MessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageRaw struct {
	Contract *Message // Generic contract binding to access the raw methods on
}

// MessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageCallerRaw struct {
	Contract *MessageCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTransactorRaw struct {
	Contract *MessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessage creates a new instance of Message, bound to a specific deployed contract.
func NewMessage(address common.Address, backend bind.ContractBackend) (*Message, error) {
	contract, err := bindMessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Message{MessageCaller: MessageCaller{contract: contract}, MessageTransactor: MessageTransactor{contract: contract}, MessageFilterer: MessageFilterer{contract: contract}}, nil
}

// NewMessageCaller creates a new read-only instance of Message, bound to a specific deployed contract.
func NewMessageCaller(address common.Address, caller bind.ContractCaller) (*MessageCaller, error) {
	contract, err := bindMessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageCaller{contract: contract}, nil
}

// NewMessageTransactor creates a new write-only instance of Message, bound to a specific deployed contract.
func NewMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTransactor, error) {
	contract, err := bindMessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransactor{contract: contract}, nil
}

// NewMessageFilterer creates a new log filterer instance of Message, bound to a specific deployed contract.
func NewMessageFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageFilterer, error) {
	contract, err := bindMessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageFilterer{contract: contract}, nil
}

// bindMessage binds a generic wrapper to an already deployed contract.
func bindMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Message *MessageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Message.Contract.MessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Message *MessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Message.Contract.MessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Message *MessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Message.Contract.MessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Message *MessageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Message.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Message *MessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Message.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Message *MessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Message.Contract.contract.Transact(opts, method, params...)
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

// OwnableUpgradeableMetaData contains all meta data concerning the OwnableUpgradeable contract.
var OwnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableUpgradeableMetaData.ABI instead.
var OwnableUpgradeableABI = OwnableUpgradeableMetaData.ABI

// Deprecated: Use OwnableUpgradeableMetaData.Sigs instead.
// OwnableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableUpgradeableFuncSigs = OwnableUpgradeableMetaData.Sigs

// OwnableUpgradeable is an auto generated Go binding around an Ethereum contract.
type OwnableUpgradeable struct {
	OwnableUpgradeableCaller     // Read-only binding to the contract
	OwnableUpgradeableTransactor // Write-only binding to the contract
	OwnableUpgradeableFilterer   // Log filterer for contract events
}

// OwnableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableUpgradeableSession struct {
	Contract     *OwnableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OwnableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableUpgradeableCallerSession struct {
	Contract *OwnableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OwnableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableUpgradeableTransactorSession struct {
	Contract     *OwnableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OwnableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableUpgradeableRaw struct {
	Contract *OwnableUpgradeable // Generic contract binding to access the raw methods on
}

// OwnableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCallerRaw struct {
	Contract *OwnableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactorRaw struct {
	Contract *OwnableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableUpgradeable creates a new instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeable(address common.Address, backend bind.ContractBackend) (*OwnableUpgradeable, error) {
	contract, err := bindOwnableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeable{OwnableUpgradeableCaller: OwnableUpgradeableCaller{contract: contract}, OwnableUpgradeableTransactor: OwnableUpgradeableTransactor{contract: contract}, OwnableUpgradeableFilterer: OwnableUpgradeableFilterer{contract: contract}}, nil
}

// NewOwnableUpgradeableCaller creates a new read-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*OwnableUpgradeableCaller, error) {
	contract, err := bindOwnableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableCaller{contract: contract}, nil
}

// NewOwnableUpgradeableTransactor creates a new write-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableUpgradeableTransactor, error) {
	contract, err := bindOwnableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableTransactor{contract: contract}, nil
}

// NewOwnableUpgradeableFilterer creates a new log filterer instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableUpgradeableFilterer, error) {
	contract, err := bindOwnableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableFilterer{contract: contract}, nil
}

// bindOwnableUpgradeable binds a generic wrapper to an already deployed contract.
func bindOwnableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCallerSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// OwnableUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitializedIterator struct {
	Event *OwnableUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableInitialized)
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
		it.Event = new(OwnableUpgradeableInitialized)
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
func (it *OwnableUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableInitialized represents a Initialized event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*OwnableUpgradeableInitializedIterator, error) {

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableInitializedIterator{contract: _OwnableUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableInitialized)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseInitialized(log types.Log) (*OwnableUpgradeableInitialized, error) {
	event := new(OwnableUpgradeableInitialized)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferredIterator struct {
	Event *OwnableUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableOwnershipTransferred)
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
		it.Event = new(OwnableUpgradeableOwnershipTransferred)
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
func (it *OwnableUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableOwnershipTransferredIterator{contract: _OwnableUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableOwnershipTransferred)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableUpgradeableOwnershipTransferred, error) {
	event := new(OwnableUpgradeableOwnershipTransferred)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QueueLibMetaData contains all meta data concerning the QueueLib contract.
var QueueLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200b871e825acd525db93c660467d34727329c256dfaa6b5ab72ac9e9ddb30422664736f6c634300080d0033",
}

// QueueLibABI is the input ABI used to generate the binding from.
// Deprecated: Use QueueLibMetaData.ABI instead.
var QueueLibABI = QueueLibMetaData.ABI

// QueueLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QueueLibMetaData.Bin instead.
var QueueLibBin = QueueLibMetaData.Bin

// DeployQueueLib deploys a new Ethereum contract, binding an instance of QueueLib to it.
func DeployQueueLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QueueLib, error) {
	parsed, err := QueueLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QueueLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// QueueLib is an auto generated Go binding around an Ethereum contract.
type QueueLib struct {
	QueueLibCaller     // Read-only binding to the contract
	QueueLibTransactor // Write-only binding to the contract
	QueueLibFilterer   // Log filterer for contract events
}

// QueueLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type QueueLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueueLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueueLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QueueLibSession struct {
	Contract     *QueueLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueueLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueueLibCallerSession struct {
	Contract *QueueLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QueueLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueueLibTransactorSession struct {
	Contract     *QueueLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QueueLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueueLibRaw struct {
	Contract *QueueLib // Generic contract binding to access the raw methods on
}

// QueueLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueueLibCallerRaw struct {
	Contract *QueueLibCaller // Generic read-only contract binding to access the raw methods on
}

// QueueLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueueLibTransactorRaw struct {
	Contract *QueueLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQueueLib creates a new instance of QueueLib, bound to a specific deployed contract.
func NewQueueLib(address common.Address, backend bind.ContractBackend) (*QueueLib, error) {
	contract, err := bindQueueLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// NewQueueLibCaller creates a new read-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibCaller(address common.Address, caller bind.ContractCaller) (*QueueLibCaller, error) {
	contract, err := bindQueueLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibCaller{contract: contract}, nil
}

// NewQueueLibTransactor creates a new write-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibTransactor(address common.Address, transactor bind.ContractTransactor) (*QueueLibTransactor, error) {
	contract, err := bindQueueLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibTransactor{contract: contract}, nil
}

// NewQueueLibFilterer creates a new log filterer instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibFilterer(address common.Address, filterer bind.ContractFilterer) (*QueueLibFilterer, error) {
	contract, err := bindQueueLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueueLibFilterer{contract: contract}, nil
}

// bindQueueLib binds a generic wrapper to an already deployed contract.
func bindQueueLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.QueueLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transact(opts, method, params...)
}

// QueueManagerMetaData contains all meta data concerning the QueueManager contract.
var QueueManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_item\",\"type\":\"bytes32\"}],\"name\":\"queueContains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueEnd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2bef2892": "queueContains(bytes32)",
		"f6d16102": "queueEnd()",
		"ab91c7b0": "queueLength()",
	},
	Bin: "0x608060405234801561001057600080fd5b50610229806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632bef289214610046578063ab91c7b01461006e578063f6d1610214610084575b600080fd5b6100596100543660046101da565b6100c1565b60405190151581526020015b60405180910390f35b6100766100d4565b604051908152602001610065565b60015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16600090815260026020526040902054610076565b60006100ce600183610114565b92915050565b600061010f6001546fffffffffffffffffffffffffffffffff8082167001000000000000000000000000000000009092048116919091031690565b905090565b81546000906fffffffffffffffffffffffffffffffff7001000000000000000000000000000000008204811691165b806fffffffffffffffffffffffffffffffff16826fffffffffffffffffffffffffffffffff1611156101cf576fffffffffffffffffffffffffffffffff821660009081526001860160205260409020548490036101a5576001925050506100ce565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90910190610143565b506000949350505050565b6000602082840312156101ec57600080fd5b503591905056fea2646970667358221220a2a79d1e000077710d8a9d53295eab3dfeb86cabcb03c8e84e73f224d23f7f3364736f6c634300080d0033",
}

// QueueManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use QueueManagerMetaData.ABI instead.
var QueueManagerABI = QueueManagerMetaData.ABI

// Deprecated: Use QueueManagerMetaData.Sigs instead.
// QueueManagerFuncSigs maps the 4-byte function signature to its string representation.
var QueueManagerFuncSigs = QueueManagerMetaData.Sigs

// QueueManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QueueManagerMetaData.Bin instead.
var QueueManagerBin = QueueManagerMetaData.Bin

// DeployQueueManager deploys a new Ethereum contract, binding an instance of QueueManager to it.
func DeployQueueManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QueueManager, error) {
	parsed, err := QueueManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QueueManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QueueManager{QueueManagerCaller: QueueManagerCaller{contract: contract}, QueueManagerTransactor: QueueManagerTransactor{contract: contract}, QueueManagerFilterer: QueueManagerFilterer{contract: contract}}, nil
}

// QueueManager is an auto generated Go binding around an Ethereum contract.
type QueueManager struct {
	QueueManagerCaller     // Read-only binding to the contract
	QueueManagerTransactor // Write-only binding to the contract
	QueueManagerFilterer   // Log filterer for contract events
}

// QueueManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type QueueManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueueManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueueManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QueueManagerSession struct {
	Contract     *QueueManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueueManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueueManagerCallerSession struct {
	Contract *QueueManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// QueueManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueueManagerTransactorSession struct {
	Contract     *QueueManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// QueueManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueueManagerRaw struct {
	Contract *QueueManager // Generic contract binding to access the raw methods on
}

// QueueManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueueManagerCallerRaw struct {
	Contract *QueueManagerCaller // Generic read-only contract binding to access the raw methods on
}

// QueueManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueueManagerTransactorRaw struct {
	Contract *QueueManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQueueManager creates a new instance of QueueManager, bound to a specific deployed contract.
func NewQueueManager(address common.Address, backend bind.ContractBackend) (*QueueManager, error) {
	contract, err := bindQueueManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QueueManager{QueueManagerCaller: QueueManagerCaller{contract: contract}, QueueManagerTransactor: QueueManagerTransactor{contract: contract}, QueueManagerFilterer: QueueManagerFilterer{contract: contract}}, nil
}

// NewQueueManagerCaller creates a new read-only instance of QueueManager, bound to a specific deployed contract.
func NewQueueManagerCaller(address common.Address, caller bind.ContractCaller) (*QueueManagerCaller, error) {
	contract, err := bindQueueManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueueManagerCaller{contract: contract}, nil
}

// NewQueueManagerTransactor creates a new write-only instance of QueueManager, bound to a specific deployed contract.
func NewQueueManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*QueueManagerTransactor, error) {
	contract, err := bindQueueManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueueManagerTransactor{contract: contract}, nil
}

// NewQueueManagerFilterer creates a new log filterer instance of QueueManager, bound to a specific deployed contract.
func NewQueueManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*QueueManagerFilterer, error) {
	contract, err := bindQueueManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueueManagerFilterer{contract: contract}, nil
}

// bindQueueManager binds a generic wrapper to an already deployed contract.
func bindQueueManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueManager *QueueManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueueManager.Contract.QueueManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueManager *QueueManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueManager.Contract.QueueManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueManager *QueueManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueManager.Contract.QueueManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueManager *QueueManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueueManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueManager *QueueManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueManager *QueueManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueManager.Contract.contract.Transact(opts, method, params...)
}

// QueueContains is a free data retrieval call binding the contract method 0x2bef2892.
//
// Solidity: function queueContains(bytes32 _item) view returns(bool)
func (_QueueManager *QueueManagerCaller) QueueContains(opts *bind.CallOpts, _item [32]byte) (bool, error) {
	var out []interface{}
	err := _QueueManager.contract.Call(opts, &out, "queueContains", _item)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// QueueContains is a free data retrieval call binding the contract method 0x2bef2892.
//
// Solidity: function queueContains(bytes32 _item) view returns(bool)
func (_QueueManager *QueueManagerSession) QueueContains(_item [32]byte) (bool, error) {
	return _QueueManager.Contract.QueueContains(&_QueueManager.CallOpts, _item)
}

// QueueContains is a free data retrieval call binding the contract method 0x2bef2892.
//
// Solidity: function queueContains(bytes32 _item) view returns(bool)
func (_QueueManager *QueueManagerCallerSession) QueueContains(_item [32]byte) (bool, error) {
	return _QueueManager.Contract.QueueContains(&_QueueManager.CallOpts, _item)
}

// QueueEnd is a free data retrieval call binding the contract method 0xf6d16102.
//
// Solidity: function queueEnd() view returns(bytes32)
func (_QueueManager *QueueManagerCaller) QueueEnd(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QueueManager.contract.Call(opts, &out, "queueEnd")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// QueueEnd is a free data retrieval call binding the contract method 0xf6d16102.
//
// Solidity: function queueEnd() view returns(bytes32)
func (_QueueManager *QueueManagerSession) QueueEnd() ([32]byte, error) {
	return _QueueManager.Contract.QueueEnd(&_QueueManager.CallOpts)
}

// QueueEnd is a free data retrieval call binding the contract method 0xf6d16102.
//
// Solidity: function queueEnd() view returns(bytes32)
func (_QueueManager *QueueManagerCallerSession) QueueEnd() ([32]byte, error) {
	return _QueueManager.Contract.QueueEnd(&_QueueManager.CallOpts)
}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_QueueManager *QueueManagerCaller) QueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QueueManager.contract.Call(opts, &out, "queueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_QueueManager *QueueManagerSession) QueueLength() (*big.Int, error) {
	return _QueueManager.Contract.QueueLength(&_QueueManager.CallOpts)
}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_QueueManager *QueueManagerCallerSession) QueueLength() (*big.Int, error) {
	return _QueueManager.Contract.QueueLength(&_QueueManager.CallOpts)
}

// QueueManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the QueueManager contract.
type QueueManagerInitializedIterator struct {
	Event *QueueManagerInitialized // Event containing the contract specifics and raw log

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
func (it *QueueManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QueueManagerInitialized)
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
		it.Event = new(QueueManagerInitialized)
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
func (it *QueueManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QueueManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QueueManagerInitialized represents a Initialized event raised by the QueueManager contract.
type QueueManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QueueManager *QueueManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*QueueManagerInitializedIterator, error) {

	logs, sub, err := _QueueManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &QueueManagerInitializedIterator{contract: _QueueManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QueueManager *QueueManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *QueueManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _QueueManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QueueManagerInitialized)
				if err := _QueueManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QueueManager *QueueManagerFilterer) ParseInitialized(log types.Log) (*QueueManagerInitialized, error) {
	event := new(QueueManagerInitialized)
	if err := _QueueManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaLibMetaData contains all meta data concerning the ReplicaLib contract.
var ReplicaLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MESSAGE_STATUS_NONE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSAGE_STATUS_PROCESSED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b0075818": "MESSAGE_STATUS_NONE()",
		"643d8086": "MESSAGE_STATUS_PROCESSED()",
	},
	Bin: "0x6098610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c8063643d8086146042578063b007581814605b575b600080fd5b6049600181565b60405190815260200160405180910390f35b604960008156fea2646970667358221220ea0145a5e7762d9fd4aa75bec9244866bf7080ec22ab00e061bb639a8af4753764736f6c634300080d0033",
}

// ReplicaLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ReplicaLibMetaData.ABI instead.
var ReplicaLibABI = ReplicaLibMetaData.ABI

// Deprecated: Use ReplicaLibMetaData.Sigs instead.
// ReplicaLibFuncSigs maps the 4-byte function signature to its string representation.
var ReplicaLibFuncSigs = ReplicaLibMetaData.Sigs

// ReplicaLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReplicaLibMetaData.Bin instead.
var ReplicaLibBin = ReplicaLibMetaData.Bin

// DeployReplicaLib deploys a new Ethereum contract, binding an instance of ReplicaLib to it.
func DeployReplicaLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReplicaLib, error) {
	parsed, err := ReplicaLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReplicaLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReplicaLib{ReplicaLibCaller: ReplicaLibCaller{contract: contract}, ReplicaLibTransactor: ReplicaLibTransactor{contract: contract}, ReplicaLibFilterer: ReplicaLibFilterer{contract: contract}}, nil
}

// ReplicaLib is an auto generated Go binding around an Ethereum contract.
type ReplicaLib struct {
	ReplicaLibCaller     // Read-only binding to the contract
	ReplicaLibTransactor // Write-only binding to the contract
	ReplicaLibFilterer   // Log filterer for contract events
}

// ReplicaLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReplicaLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReplicaLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReplicaLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReplicaLibSession struct {
	Contract     *ReplicaLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReplicaLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReplicaLibCallerSession struct {
	Contract *ReplicaLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReplicaLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReplicaLibTransactorSession struct {
	Contract     *ReplicaLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReplicaLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReplicaLibRaw struct {
	Contract *ReplicaLib // Generic contract binding to access the raw methods on
}

// ReplicaLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReplicaLibCallerRaw struct {
	Contract *ReplicaLibCaller // Generic read-only contract binding to access the raw methods on
}

// ReplicaLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReplicaLibTransactorRaw struct {
	Contract *ReplicaLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReplicaLib creates a new instance of ReplicaLib, bound to a specific deployed contract.
func NewReplicaLib(address common.Address, backend bind.ContractBackend) (*ReplicaLib, error) {
	contract, err := bindReplicaLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReplicaLib{ReplicaLibCaller: ReplicaLibCaller{contract: contract}, ReplicaLibTransactor: ReplicaLibTransactor{contract: contract}, ReplicaLibFilterer: ReplicaLibFilterer{contract: contract}}, nil
}

// NewReplicaLibCaller creates a new read-only instance of ReplicaLib, bound to a specific deployed contract.
func NewReplicaLibCaller(address common.Address, caller bind.ContractCaller) (*ReplicaLibCaller, error) {
	contract, err := bindReplicaLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicaLibCaller{contract: contract}, nil
}

// NewReplicaLibTransactor creates a new write-only instance of ReplicaLib, bound to a specific deployed contract.
func NewReplicaLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ReplicaLibTransactor, error) {
	contract, err := bindReplicaLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicaLibTransactor{contract: contract}, nil
}

// NewReplicaLibFilterer creates a new log filterer instance of ReplicaLib, bound to a specific deployed contract.
func NewReplicaLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ReplicaLibFilterer, error) {
	contract, err := bindReplicaLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReplicaLibFilterer{contract: contract}, nil
}

// bindReplicaLib binds a generic wrapper to an already deployed contract.
func bindReplicaLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReplicaLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReplicaLib *ReplicaLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReplicaLib.Contract.ReplicaLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReplicaLib *ReplicaLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReplicaLib.Contract.ReplicaLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReplicaLib *ReplicaLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReplicaLib.Contract.ReplicaLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReplicaLib *ReplicaLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReplicaLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReplicaLib *ReplicaLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReplicaLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReplicaLib *ReplicaLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReplicaLib.Contract.contract.Transact(opts, method, params...)
}

// MESSAGESTATUSNONE is a free data retrieval call binding the contract method 0xb0075818.
//
// Solidity: function MESSAGE_STATUS_NONE() view returns(bytes32)
func (_ReplicaLib *ReplicaLibCaller) MESSAGESTATUSNONE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ReplicaLib.contract.Call(opts, &out, "MESSAGE_STATUS_NONE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSAGESTATUSNONE is a free data retrieval call binding the contract method 0xb0075818.
//
// Solidity: function MESSAGE_STATUS_NONE() view returns(bytes32)
func (_ReplicaLib *ReplicaLibSession) MESSAGESTATUSNONE() ([32]byte, error) {
	return _ReplicaLib.Contract.MESSAGESTATUSNONE(&_ReplicaLib.CallOpts)
}

// MESSAGESTATUSNONE is a free data retrieval call binding the contract method 0xb0075818.
//
// Solidity: function MESSAGE_STATUS_NONE() view returns(bytes32)
func (_ReplicaLib *ReplicaLibCallerSession) MESSAGESTATUSNONE() ([32]byte, error) {
	return _ReplicaLib.Contract.MESSAGESTATUSNONE(&_ReplicaLib.CallOpts)
}

// MESSAGESTATUSPROCESSED is a free data retrieval call binding the contract method 0x643d8086.
//
// Solidity: function MESSAGE_STATUS_PROCESSED() view returns(bytes32)
func (_ReplicaLib *ReplicaLibCaller) MESSAGESTATUSPROCESSED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ReplicaLib.contract.Call(opts, &out, "MESSAGE_STATUS_PROCESSED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSAGESTATUSPROCESSED is a free data retrieval call binding the contract method 0x643d8086.
//
// Solidity: function MESSAGE_STATUS_PROCESSED() view returns(bytes32)
func (_ReplicaLib *ReplicaLibSession) MESSAGESTATUSPROCESSED() ([32]byte, error) {
	return _ReplicaLib.Contract.MESSAGESTATUSPROCESSED(&_ReplicaLib.CallOpts)
}

// MESSAGESTATUSPROCESSED is a free data retrieval call binding the contract method 0x643d8086.
//
// Solidity: function MESSAGE_STATUS_PROCESSED() view returns(bytes32)
func (_ReplicaLib *ReplicaLibCallerSession) MESSAGESTATUSPROCESSED() ([32]byte, error) {
	return _ReplicaLib.Contract.MESSAGESTATUSPROCESSED(&_ReplicaLib.CallOpts)
}

// ReplicaManagerMetaData contains all meta data concerning the ReplicaManager contract.
var ReplicaManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_processGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveGas\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"Process\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousConfirmAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newConfirmAt\",\"type\":\"uint256\"}],\"name\":\"SetConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PROCESS_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"acceptableRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"}],\"name\":\"activeReplicaCommittedRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"activeReplicaConfirmedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"activeReplicaMessageStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_homeDomain\",\"type\":\"uint32\"}],\"name\":\"homeDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"process\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_proof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"prove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_proof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"proveAndProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmAt\",\"type\":\"uint256\"}],\"name\":\"setConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_oldRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d88beda2": "PROCESS_GAS()",
		"25e3beda": "RESERVE_GAS()",
		"ffa1ad74": "VERSION()",
		"15a046aa": "acceptableRoot(uint32,uint32,bytes32)",
		"f1e74e06": "activeReplicaCommittedRoot(uint32)",
		"7dfdba28": "activeReplicaConfirmedAt(uint32,bytes32)",
		"63415514": "activeReplicaMessageStatus(uint32,bytes32)",
		"fd749546": "homeDomainHash(uint32)",
		"8624c35c": "initialize(uint32,address)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"928bc4b2": "process(bytes)",
		"4f63be3f": "prove(uint32,bytes,bytes32[32],uint256)",
		"68705275": "proveAndProcess(uint32,bytes,bytes32[32],uint256)",
		"715018a6": "renounceOwnership()",
		"9df7d36d": "setConfirmation(uint32,bytes32,uint256)",
		"9d54f419": "setUpdater(address)",
		"f2fde38b": "transferOwnership(address)",
		"dec48b67": "update(uint32,bytes32,bytes32,bytes)",
		"df034cd0": "updater()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b5060405162002c0a38038062002c0a8339810160408190526200003491620000d9565b63ffffffff8316608052620cf850821015620000865760405162461bcd60e51b815260206004820152600c60248201526b2170726f636573732067617360a01b60448201526064015b60405180910390fd5b613a98811015620000c95760405162461bcd60e51b815260206004820152600c60248201526b21726573657276652067617360a01b60448201526064016200007d565b60a09190915260c052506200011b565b600080600060608486031215620000ef57600080fd5b835163ffffffff811681146200010457600080fd5b602085015160409095015190969495509392505050565b60805160a05160c051612aa3620001676000396000818161019e01526109d601526000818161034c015281816109f70152610b9a01526000818161029801526107df0152612aa36000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c8063928bc4b2116100cd578063df034cd011610081578063f2fde38b11610066578063f2fde38b146103d4578063fd749546146103e7578063ffa1ad74146103fa57600080fd5b8063df034cd014610381578063f1e74e06146103a157600080fd5b80639df7d36d116100b25780639df7d36d14610334578063d88beda214610347578063dec48b671461036e57600080fd5b8063928bc4b21461030e5780639d54f4191461032157600080fd5b8063715018a6116101245780638624c35c116101095780638624c35c146102805780638d3638f4146102935780638da5cb5b146102cf57600080fd5b8063715018a6146102375780637dfdba281461023f57600080fd5b80634f63be3f116101555780634f63be3f146101ce57806363415514146101e1578063687052751461022257600080fd5b806315a046aa1461017157806325e3beda14610199575b600080fd5b61018461017f3660046123a7565b610414565b60405190151581526020015b60405180910390f35b6101c07f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610190565b6101846101dc3660046124fd565b610471565b6101c06101ef36600461256d565b63ffffffff919091166000908152609a602090815260408083205483526099825280832093835260039093019052205490565b6102356102303660046124fd565b610578565b005b6102356105e0565b6101c061024d36600461256d565b63ffffffff919091166000908152609a602090815260408083205483526099825280832093835260029093019052205490565b61023561028e3660046125bb565b610649565b6102ba7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610190565b60335473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610190565b61018461031c3660046125ee565b610797565b61023561032f366004612623565b610c9b565b61023561034236600461263e565b610d0e565b6101c07f000000000000000000000000000000000000000000000000000000000000000081565b61023561037c366004612671565b610e02565b6065546102e99073ffffffffffffffffffffffffffffffffffffffff1681565b6101c06103af3660046126d2565b63ffffffff166000908152609a60209081526040808320548352609990915290205490565b6102356103e2366004612623565b610f29565b6101c06103f53660046126d2565b611022565b610402600081565b60405160ff9091168152602001610190565b63ffffffff83166000908152609a602090815260408083205483526099825280832084845260020190915281205480820361045357600091505061046a565b61046363ffffffff85168261271c565b4210159150505b9392505050565b825160208085019190912063ffffffff86166000908152609a835260408082205482526099845280822083835260038101909452812054909290156104fd5760405162461bcd60e51b815260206004820152601360248201527f214d6573736167655374617475732e4e6f6e650000000000000000000000000060448201526064015b60405180910390fd5b60006105338387602080602002604051908101604052809291908260208002808284376000920191909152508991506110339050565b600081815260028401602052604090205490915015610568576000928352600391909101602052604090912055506001610570565b600093505050505b949350505050565b61058484848484610471565b6105d05760405162461bcd60e51b815260206004820152600660248201527f2170726f7665000000000000000000000000000000000000000000000000000060448201526064016104f4565b6105d983610797565b5050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146106475760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104f4565b565b600061065560016110d9565b9050801561068a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61069382611230565b609780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610719836098546000818152609960205260409020600101805463ffffffff84167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009091161764010000000017905560018101609855919050565b63ffffffff84166000908152609a6020526040902055801561079257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6000806107a483826112b5565b905060006107b762ffffff1983166112d9565b63ffffffff8082166000908152609a60209081526040808320548352609990915290209192507f00000000000000000000000000000000000000000000000000000000000000001661080e62ffffff1985166112ed565b63ffffffff16146108615760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e000000000000000000000000000000000000000060448201526064016104f4565b600061087262ffffff198516611302565b600081815260038401602052604090205490915061088f8161135f565b6108db5760405162461bcd60e51b815260206004820152601460248201527f21657869737473207c7c2070726f63657373656400000000000000000000000060448201526064016104f4565b6108f4846108ee62ffffff198816611373565b83610414565b6109405760405162461bcd60e51b815260206004820152601260248201527f216f7074696d69737469635365636f6e6473000000000000000000000000000060448201526064016104f4565b60975460ff166001146109955760405162461bcd60e51b815260206004820152600a60248201527f217265656e7472616e740000000000000000000000000000000000000000000060448201526064016104f4565b609780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556000828152600384016020526040902060019055610a1b7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000061271c565b5a1015610a6c5760405162461bcd60e51b81526004016104f49060208082526004908201527f2167617300000000000000000000000000000000000000000000000000000000604082015260600190565b60007fe4d16d620000000000000000000000000000000000000000000000000000000085610a9f62ffffff198916611388565b610aae62ffffff198a1661139d565b6000868152600289016020526040902054610adc610ad162ffffff198d166113b2565b62ffffff19166113e5565b604051602401610af09594939291906127ae565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915290506000610b8062ffffff198816611438565b6040805161010080825261012082019092529192506000917f00000000000000000000000000000000000000000000000000000000000000009083908360208201818036833701905050905060008087516020890160008987f19b503d935082841115610beb578293505b838152836000602083013e8b610c1d57610c0481611449565b60405162461bcd60e51b81526004016104f491906127ee565b8b1515888b63ffffffff167f223de0966a99342a66dcd8e6b41362efb8e142d6ea63bca2fa73514df1d1f48f84604051610c5791906127ee565b60405180910390a45050609780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905550979998505050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610d025760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104f4565b610d0b816114a8565b50565b60335473ffffffffffffffffffffffffffffffffffffffff163314610d755760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104f4565b63ffffffff8084166000908152609a602090815260408083205483526099825280832086845260028101909252909120549091610db89083908690869061152e16565b6040805182815260208101859052859163ffffffff8816917f6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730910160405180910390a35050505050565b63ffffffff84166000908152609a602090815260408083205483526099909152902080548414610e745760405162461bcd60e51b815260206004820152601260248201527f6e6f742063757272656e7420757064617465000000000000000000000000000060448201526064016104f4565b610e8085858585611542565b610ecc5760405162461bcd60e51b815260206004820152600c60248201527f217570646174657220736967000000000000000000000000000000000000000060448201526064016104f4565b6000838152600282016020526040902042905582815582848663ffffffff167f608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b285604051610f1a91906127ee565b60405180910390a45050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610f905760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104f4565b73ffffffffffffffffffffffffffffffffffffffff81166110195760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104f4565b610d0b816115d1565b600061102d82611648565b92915050565b8260005b60208110156110d157600183821c16600085836020811061105a5761105a612801565b602002015190508160010361109a5760408051602081018390529081018590526060016040516020818303038152906040528051906020012093506110c7565b60408051602081018690529081018290526060016040516020818303038152906040528051906020012093505b5050600101611037565b509392505050565b60008054610100900460ff1615611176578160ff1660011480156110fc5750303b155b61116e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104f4565b506000919050565b60005460ff8084169116106111f35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104f4565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b600054610100900460ff166112ad5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f4565b610d026116c1565b8151600090602084016112d064ffffffffff85168284611746565b95945050505050565b600061102d62ffffff19831682600461178d565b600061102d62ffffff1983166028600461178d565b60008061131d8360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006113478460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6000811580159061102d5750506001141590565b600061102d62ffffff198316604c600461178d565b600061102d62ffffff1983166024600461178d565b600061102d62ffffff198316600460206117bd565b600061102d60506113d581601886901c6bffffffffffffffffffffffff16612830565b62ffffff19851691906000611962565b60606000806114028460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061142784836020016119e6565b508181016020016040529052919050565b600061102d61144683611b81565b90565b606060448251101561148e57505060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c79000000602082015290565b6004820191508180602001905181019061102d9190612847565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a910160405180910390a15050565b600091825260029092016020526040902055565b60008061154e86611648565b60408051602081019290925281018690526060810185905260800160405160208183030381529060405280519060200120905061158a81611b96565b60655490915073ffffffffffffffffffffffffffffffffffffffff166115b08285611bd1565b73ffffffffffffffffffffffffffffffffffffffff16149695505050505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e083901b1660208201527f53594e000000000000000000000000000000000000000000000000000000000060248201526000906027015b604051602081830303815290604052805190602001209050919050565b600054610100900460ff1661173e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f4565b610647611bed565b600080611753838561271c565b9050604051811115611763575060005b806000036117785762ffffff1991505061046a565b5050606092831b9190911790911b1760181b90565b600061179a8260206128e4565b6117a5906008612907565b60ff166117b38585856117bd565b901c949350505050565b60008160ff166000036117d25750600061046a565b6117ea8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1661180560ff84168561271c565b111561186457610c046118268560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1661184c8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16611c73565b60208260ff1611156118de5760405162461bcd60e51b815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e20333220627974657300000000000060648201526084016104f4565b6008820260006118fc8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b60008061197d8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905061199686611ce1565b846119a1878461271c565b6119ab919061271c565b11156119be5762ffffff19915050610570565b6119c8858261271c565b90506119dc8364ffffffffff168286611746565b9695505050505050565b600062ffffff1980841603611a635760405162461bcd60e51b815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f657220646572656600000000000000000000000000000000000000000000000060648201526084016104f4565b611a6c83611d29565b611ade5760405162461bcd60e51b815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e74657220646572656600000000000000000000000000000000000000000060648201526084016104f4565b6000611af88460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000611b228560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000604051905084811115611b475760206060fd5b8285848460045afa506119dc611b5d8760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b600061102d62ffffff198316602c60206117bd565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c016116a4565b6000806000611be08585611d66565b915091506110d181611dd4565b600054610100900460ff16611c6a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f4565b610647336115d1565b60606000611c8086611fc0565b9150506000611c8e86611fc0565b9150506000611c9c86611fc0565b9150506000611caa86611fc0565b91505083838383604051602001611cc49493929190612930565b604051602081830303815290604052945050505050949350505050565b6000611cfb8260181c6bffffffffffffffffffffffff1690565b611d138360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b6000611d358260d81c90565b64ffffffffff1664ffffffffff03611d4f57506000919050565b6000611d5a83611ce1565b60405110199392505050565b6000808251604103611d9c5760208301516040840151606085015160001a611d90878285856120aa565b94509450505050611dcd565b8251604003611dc55760208301516040840151611dba8683836121c2565b935093505050611dcd565b506000905060025b9250929050565b6000816004811115611de857611de86128b5565b03611df05750565b6001816004811115611e0457611e046128b5565b03611e515760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104f4565b6002816004811115611e6557611e656128b5565b03611eb25760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104f4565b6003816004811115611ec657611ec66128b5565b03611f395760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104f4565b6004816004811115611f4d57611f4d6128b5565b03610d0b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104f4565b600080601f5b600f8160ff161115612033576000611fdf826008612907565b60ff1685901c9050611ff081612214565b61ffff16841793508160ff1660101461200b57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611fc6565b50600f5b60ff8160ff1610156120a4576000612050826008612907565b60ff1685901c905061206181612214565b61ffff16831792508160ff1660001461207c57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612037565b50915091565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156120e157506000905060036121b9565b8460ff16601b141580156120f957508460ff16601c14155b1561210a57506000905060046121b9565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561215e573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166121b2576000600192509250506121b9565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316816121f860ff86901c601b61271c565b9050612206878288856120aa565b935093505050935093915050565b600061222660048360ff16901c612246565b60ff1661ffff919091161760081b61223d82612246565b60ff1617919050565b600060f08083179060ff821690036122615750603092915050565b8060ff1660f1036122755750603192915050565b8060ff1660f2036122895750603292915050565b8060ff1660f30361229d5750603392915050565b8060ff1660f4036122b15750603492915050565b8060ff1660f5036122c55750603592915050565b8060ff1660f6036122d95750603692915050565b8060ff1660f7036122ed5750603792915050565b8060ff1660f8036123015750603892915050565b8060ff1660f9036123155750603992915050565b8060ff1660fa036123295750606192915050565b8060ff1660fb0361233d5750606292915050565b8060ff1660fc036123515750606392915050565b8060ff1660fd036123655750606492915050565b8060ff1660fe036123795750606592915050565b8060ff1660ff0361238d5750606692915050565b50919050565b803563ffffffff8116811461122b57600080fd5b6000806000606084860312156123bc57600080fd5b6123c584612393565b92506123d360208501612393565b9150604084013590509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612459576124596123e3565b604052919050565b600067ffffffffffffffff82111561247b5761247b6123e3565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f8301126124b857600080fd5b81356124cb6124c682612461565b612412565b8181528460208386010111156124e057600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080610460858703121561251457600080fd5b61251d85612393565b9350602085013567ffffffffffffffff81111561253957600080fd5b612545878288016124a7565b93505061044085018681111561255a57600080fd5b9396929550505060409290920191903590565b6000806040838503121561258057600080fd5b61258983612393565b946020939093013593505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461122b57600080fd5b600080604083850312156125ce57600080fd5b6125d783612393565b91506125e560208401612597565b90509250929050565b60006020828403121561260057600080fd5b813567ffffffffffffffff81111561261757600080fd5b610570848285016124a7565b60006020828403121561263557600080fd5b61046a82612597565b60008060006060848603121561265357600080fd5b61265c84612393565b95602085013595506040909401359392505050565b6000806000806080858703121561268757600080fd5b61269085612393565b93506020850135925060408501359150606085013567ffffffffffffffff8111156126ba57600080fd5b6126c6878288016124a7565b91505092959194509250565b6000602082840312156126e457600080fd5b61046a82612393565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561272f5761272f6126ed565b500190565b60005b8381101561274f578181015183820152602001612737565b8381111561275e576000848401525b50505050565b6000815180845261277c816020860160208601612734565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600063ffffffff808816835280871660208401525084604083015283606083015260a060808301526127e360a0830184612764565b979650505050505050565b60208152600061046a6020830184612764565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082821015612842576128426126ed565b500390565b60006020828403121561285957600080fd5b815167ffffffffffffffff81111561287057600080fd5b8201601f8101841361288157600080fd5b805161288f6124c682612461565b8181528560208385010111156128a457600080fd5b6112d0826020830160208601612734565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600060ff821660ff8416808210156128fe576128fe6126ed565b90039392505050565b600060ff821660ff84168160ff0481118215151615612928576129286126ed565b029392505050565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d82016119dc56fea26469706673582212202d276f5cbcfef8a48e8380b97ac5c0ad80d90de81cc90e0385e10a2f840928b664736f6c634300080d0033",
}

// ReplicaManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ReplicaManagerMetaData.ABI instead.
var ReplicaManagerABI = ReplicaManagerMetaData.ABI

// Deprecated: Use ReplicaManagerMetaData.Sigs instead.
// ReplicaManagerFuncSigs maps the 4-byte function signature to its string representation.
var ReplicaManagerFuncSigs = ReplicaManagerMetaData.Sigs

// ReplicaManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReplicaManagerMetaData.Bin instead.
var ReplicaManagerBin = ReplicaManagerMetaData.Bin

// DeployReplicaManager deploys a new Ethereum contract, binding an instance of ReplicaManager to it.
func DeployReplicaManager(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32, _processGas *big.Int, _reserveGas *big.Int) (common.Address, *types.Transaction, *ReplicaManager, error) {
	parsed, err := ReplicaManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReplicaManagerBin), backend, _localDomain, _processGas, _reserveGas)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReplicaManager{ReplicaManagerCaller: ReplicaManagerCaller{contract: contract}, ReplicaManagerTransactor: ReplicaManagerTransactor{contract: contract}, ReplicaManagerFilterer: ReplicaManagerFilterer{contract: contract}}, nil
}

// ReplicaManager is an auto generated Go binding around an Ethereum contract.
type ReplicaManager struct {
	ReplicaManagerCaller     // Read-only binding to the contract
	ReplicaManagerTransactor // Write-only binding to the contract
	ReplicaManagerFilterer   // Log filterer for contract events
}

// ReplicaManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReplicaManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReplicaManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReplicaManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReplicaManagerSession struct {
	Contract     *ReplicaManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReplicaManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReplicaManagerCallerSession struct {
	Contract *ReplicaManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ReplicaManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReplicaManagerTransactorSession struct {
	Contract     *ReplicaManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReplicaManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReplicaManagerRaw struct {
	Contract *ReplicaManager // Generic contract binding to access the raw methods on
}

// ReplicaManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReplicaManagerCallerRaw struct {
	Contract *ReplicaManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ReplicaManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReplicaManagerTransactorRaw struct {
	Contract *ReplicaManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReplicaManager creates a new instance of ReplicaManager, bound to a specific deployed contract.
func NewReplicaManager(address common.Address, backend bind.ContractBackend) (*ReplicaManager, error) {
	contract, err := bindReplicaManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReplicaManager{ReplicaManagerCaller: ReplicaManagerCaller{contract: contract}, ReplicaManagerTransactor: ReplicaManagerTransactor{contract: contract}, ReplicaManagerFilterer: ReplicaManagerFilterer{contract: contract}}, nil
}

// NewReplicaManagerCaller creates a new read-only instance of ReplicaManager, bound to a specific deployed contract.
func NewReplicaManagerCaller(address common.Address, caller bind.ContractCaller) (*ReplicaManagerCaller, error) {
	contract, err := bindReplicaManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerCaller{contract: contract}, nil
}

// NewReplicaManagerTransactor creates a new write-only instance of ReplicaManager, bound to a specific deployed contract.
func NewReplicaManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ReplicaManagerTransactor, error) {
	contract, err := bindReplicaManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerTransactor{contract: contract}, nil
}

// NewReplicaManagerFilterer creates a new log filterer instance of ReplicaManager, bound to a specific deployed contract.
func NewReplicaManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ReplicaManagerFilterer, error) {
	contract, err := bindReplicaManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerFilterer{contract: contract}, nil
}

// bindReplicaManager binds a generic wrapper to an already deployed contract.
func bindReplicaManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReplicaManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReplicaManager *ReplicaManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReplicaManager.Contract.ReplicaManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReplicaManager *ReplicaManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReplicaManager.Contract.ReplicaManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReplicaManager *ReplicaManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReplicaManager.Contract.ReplicaManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReplicaManager *ReplicaManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReplicaManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReplicaManager *ReplicaManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReplicaManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReplicaManager *ReplicaManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReplicaManager.Contract.contract.Transact(opts, method, params...)
}

// PROCESSGAS is a free data retrieval call binding the contract method 0xd88beda2.
//
// Solidity: function PROCESS_GAS() view returns(uint256)
func (_ReplicaManager *ReplicaManagerCaller) PROCESSGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "PROCESS_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROCESSGAS is a free data retrieval call binding the contract method 0xd88beda2.
//
// Solidity: function PROCESS_GAS() view returns(uint256)
func (_ReplicaManager *ReplicaManagerSession) PROCESSGAS() (*big.Int, error) {
	return _ReplicaManager.Contract.PROCESSGAS(&_ReplicaManager.CallOpts)
}

// PROCESSGAS is a free data retrieval call binding the contract method 0xd88beda2.
//
// Solidity: function PROCESS_GAS() view returns(uint256)
func (_ReplicaManager *ReplicaManagerCallerSession) PROCESSGAS() (*big.Int, error) {
	return _ReplicaManager.Contract.PROCESSGAS(&_ReplicaManager.CallOpts)
}

// RESERVEGAS is a free data retrieval call binding the contract method 0x25e3beda.
//
// Solidity: function RESERVE_GAS() view returns(uint256)
func (_ReplicaManager *ReplicaManagerCaller) RESERVEGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "RESERVE_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RESERVEGAS is a free data retrieval call binding the contract method 0x25e3beda.
//
// Solidity: function RESERVE_GAS() view returns(uint256)
func (_ReplicaManager *ReplicaManagerSession) RESERVEGAS() (*big.Int, error) {
	return _ReplicaManager.Contract.RESERVEGAS(&_ReplicaManager.CallOpts)
}

// RESERVEGAS is a free data retrieval call binding the contract method 0x25e3beda.
//
// Solidity: function RESERVE_GAS() view returns(uint256)
func (_ReplicaManager *ReplicaManagerCallerSession) RESERVEGAS() (*big.Int, error) {
	return _ReplicaManager.Contract.RESERVEGAS(&_ReplicaManager.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_ReplicaManager *ReplicaManagerCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_ReplicaManager *ReplicaManagerSession) VERSION() (uint8, error) {
	return _ReplicaManager.Contract.VERSION(&_ReplicaManager.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_ReplicaManager *ReplicaManagerCallerSession) VERSION() (uint8, error) {
	return _ReplicaManager.Contract.VERSION(&_ReplicaManager.CallOpts)
}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_ReplicaManager *ReplicaManagerCaller) AcceptableRoot(opts *bind.CallOpts, _remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "acceptableRoot", _remoteDomain, _optimisticSeconds, _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_ReplicaManager *ReplicaManagerSession) AcceptableRoot(_remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	return _ReplicaManager.Contract.AcceptableRoot(&_ReplicaManager.CallOpts, _remoteDomain, _optimisticSeconds, _root)
}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_ReplicaManager *ReplicaManagerCallerSession) AcceptableRoot(_remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	return _ReplicaManager.Contract.AcceptableRoot(&_ReplicaManager.CallOpts, _remoteDomain, _optimisticSeconds, _root)
}

// ActiveReplicaCommittedRoot is a free data retrieval call binding the contract method 0xf1e74e06.
//
// Solidity: function activeReplicaCommittedRoot(uint32 _remoteDomain) view returns(bytes32)
func (_ReplicaManager *ReplicaManagerCaller) ActiveReplicaCommittedRoot(opts *bind.CallOpts, _remoteDomain uint32) ([32]byte, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "activeReplicaCommittedRoot", _remoteDomain)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActiveReplicaCommittedRoot is a free data retrieval call binding the contract method 0xf1e74e06.
//
// Solidity: function activeReplicaCommittedRoot(uint32 _remoteDomain) view returns(bytes32)
func (_ReplicaManager *ReplicaManagerSession) ActiveReplicaCommittedRoot(_remoteDomain uint32) ([32]byte, error) {
	return _ReplicaManager.Contract.ActiveReplicaCommittedRoot(&_ReplicaManager.CallOpts, _remoteDomain)
}

// ActiveReplicaCommittedRoot is a free data retrieval call binding the contract method 0xf1e74e06.
//
// Solidity: function activeReplicaCommittedRoot(uint32 _remoteDomain) view returns(bytes32)
func (_ReplicaManager *ReplicaManagerCallerSession) ActiveReplicaCommittedRoot(_remoteDomain uint32) ([32]byte, error) {
	return _ReplicaManager.Contract.ActiveReplicaCommittedRoot(&_ReplicaManager.CallOpts, _remoteDomain)
}

// ActiveReplicaConfirmedAt is a free data retrieval call binding the contract method 0x7dfdba28.
//
// Solidity: function activeReplicaConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_ReplicaManager *ReplicaManagerCaller) ActiveReplicaConfirmedAt(opts *bind.CallOpts, _remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "activeReplicaConfirmedAt", _remoteDomain, _root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveReplicaConfirmedAt is a free data retrieval call binding the contract method 0x7dfdba28.
//
// Solidity: function activeReplicaConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_ReplicaManager *ReplicaManagerSession) ActiveReplicaConfirmedAt(_remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	return _ReplicaManager.Contract.ActiveReplicaConfirmedAt(&_ReplicaManager.CallOpts, _remoteDomain, _root)
}

// ActiveReplicaConfirmedAt is a free data retrieval call binding the contract method 0x7dfdba28.
//
// Solidity: function activeReplicaConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_ReplicaManager *ReplicaManagerCallerSession) ActiveReplicaConfirmedAt(_remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	return _ReplicaManager.Contract.ActiveReplicaConfirmedAt(&_ReplicaManager.CallOpts, _remoteDomain, _root)
}

// ActiveReplicaMessageStatus is a free data retrieval call binding the contract method 0x63415514.
//
// Solidity: function activeReplicaMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_ReplicaManager *ReplicaManagerCaller) ActiveReplicaMessageStatus(opts *bind.CallOpts, _remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "activeReplicaMessageStatus", _remoteDomain, _messageId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActiveReplicaMessageStatus is a free data retrieval call binding the contract method 0x63415514.
//
// Solidity: function activeReplicaMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_ReplicaManager *ReplicaManagerSession) ActiveReplicaMessageStatus(_remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	return _ReplicaManager.Contract.ActiveReplicaMessageStatus(&_ReplicaManager.CallOpts, _remoteDomain, _messageId)
}

// ActiveReplicaMessageStatus is a free data retrieval call binding the contract method 0x63415514.
//
// Solidity: function activeReplicaMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_ReplicaManager *ReplicaManagerCallerSession) ActiveReplicaMessageStatus(_remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	return _ReplicaManager.Contract.ActiveReplicaMessageStatus(&_ReplicaManager.CallOpts, _remoteDomain, _messageId)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0xfd749546.
//
// Solidity: function homeDomainHash(uint32 _homeDomain) pure returns(bytes32)
func (_ReplicaManager *ReplicaManagerCaller) HomeDomainHash(opts *bind.CallOpts, _homeDomain uint32) ([32]byte, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "homeDomainHash", _homeDomain)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HomeDomainHash is a free data retrieval call binding the contract method 0xfd749546.
//
// Solidity: function homeDomainHash(uint32 _homeDomain) pure returns(bytes32)
func (_ReplicaManager *ReplicaManagerSession) HomeDomainHash(_homeDomain uint32) ([32]byte, error) {
	return _ReplicaManager.Contract.HomeDomainHash(&_ReplicaManager.CallOpts, _homeDomain)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0xfd749546.
//
// Solidity: function homeDomainHash(uint32 _homeDomain) pure returns(bytes32)
func (_ReplicaManager *ReplicaManagerCallerSession) HomeDomainHash(_homeDomain uint32) ([32]byte, error) {
	return _ReplicaManager.Contract.HomeDomainHash(&_ReplicaManager.CallOpts, _homeDomain)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ReplicaManager *ReplicaManagerCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ReplicaManager *ReplicaManagerSession) LocalDomain() (uint32, error) {
	return _ReplicaManager.Contract.LocalDomain(&_ReplicaManager.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ReplicaManager *ReplicaManagerCallerSession) LocalDomain() (uint32, error) {
	return _ReplicaManager.Contract.LocalDomain(&_ReplicaManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReplicaManager *ReplicaManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReplicaManager *ReplicaManagerSession) Owner() (common.Address, error) {
	return _ReplicaManager.Contract.Owner(&_ReplicaManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReplicaManager *ReplicaManagerCallerSession) Owner() (common.Address, error) {
	return _ReplicaManager.Contract.Owner(&_ReplicaManager.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_ReplicaManager *ReplicaManagerCaller) Updater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReplicaManager.contract.Call(opts, &out, "updater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_ReplicaManager *ReplicaManagerSession) Updater() (common.Address, error) {
	return _ReplicaManager.Contract.Updater(&_ReplicaManager.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_ReplicaManager *ReplicaManagerCallerSession) Updater() (common.Address, error) {
	return _ReplicaManager.Contract.Updater(&_ReplicaManager.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _updater) returns()
func (_ReplicaManager *ReplicaManagerTransactor) Initialize(opts *bind.TransactOpts, _remoteDomain uint32, _updater common.Address) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "initialize", _remoteDomain, _updater)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _updater) returns()
func (_ReplicaManager *ReplicaManagerSession) Initialize(_remoteDomain uint32, _updater common.Address) (*types.Transaction, error) {
	return _ReplicaManager.Contract.Initialize(&_ReplicaManager.TransactOpts, _remoteDomain, _updater)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _updater) returns()
func (_ReplicaManager *ReplicaManagerTransactorSession) Initialize(_remoteDomain uint32, _updater common.Address) (*types.Transaction, error) {
	return _ReplicaManager.Contract.Initialize(&_ReplicaManager.TransactOpts, _remoteDomain, _updater)
}

// Process is a paid mutator transaction binding the contract method 0x928bc4b2.
//
// Solidity: function process(bytes _message) returns(bool _success)
func (_ReplicaManager *ReplicaManagerTransactor) Process(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "process", _message)
}

// Process is a paid mutator transaction binding the contract method 0x928bc4b2.
//
// Solidity: function process(bytes _message) returns(bool _success)
func (_ReplicaManager *ReplicaManagerSession) Process(_message []byte) (*types.Transaction, error) {
	return _ReplicaManager.Contract.Process(&_ReplicaManager.TransactOpts, _message)
}

// Process is a paid mutator transaction binding the contract method 0x928bc4b2.
//
// Solidity: function process(bytes _message) returns(bool _success)
func (_ReplicaManager *ReplicaManagerTransactorSession) Process(_message []byte) (*types.Transaction, error) {
	return _ReplicaManager.Contract.Process(&_ReplicaManager.TransactOpts, _message)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_ReplicaManager *ReplicaManagerTransactor) Prove(opts *bind.TransactOpts, _remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "prove", _remoteDomain, _message, _proof, _index)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_ReplicaManager *ReplicaManagerSession) Prove(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.Contract.Prove(&_ReplicaManager.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_ReplicaManager *ReplicaManagerTransactorSession) Prove(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.Contract.Prove(&_ReplicaManager.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// ProveAndProcess is a paid mutator transaction binding the contract method 0x68705275.
//
// Solidity: function proveAndProcess(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_ReplicaManager *ReplicaManagerTransactor) ProveAndProcess(opts *bind.TransactOpts, _remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "proveAndProcess", _remoteDomain, _message, _proof, _index)
}

// ProveAndProcess is a paid mutator transaction binding the contract method 0x68705275.
//
// Solidity: function proveAndProcess(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_ReplicaManager *ReplicaManagerSession) ProveAndProcess(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.Contract.ProveAndProcess(&_ReplicaManager.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// ProveAndProcess is a paid mutator transaction binding the contract method 0x68705275.
//
// Solidity: function proveAndProcess(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_ReplicaManager *ReplicaManagerTransactorSession) ProveAndProcess(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.Contract.ProveAndProcess(&_ReplicaManager.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReplicaManager *ReplicaManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReplicaManager *ReplicaManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReplicaManager.Contract.RenounceOwnership(&_ReplicaManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReplicaManager *ReplicaManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReplicaManager.Contract.RenounceOwnership(&_ReplicaManager.TransactOpts)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_ReplicaManager *ReplicaManagerTransactor) SetConfirmation(opts *bind.TransactOpts, _remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "setConfirmation", _remoteDomain, _root, _confirmAt)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_ReplicaManager *ReplicaManagerSession) SetConfirmation(_remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.Contract.SetConfirmation(&_ReplicaManager.TransactOpts, _remoteDomain, _root, _confirmAt)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_ReplicaManager *ReplicaManagerTransactorSession) SetConfirmation(_remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _ReplicaManager.Contract.SetConfirmation(&_ReplicaManager.TransactOpts, _remoteDomain, _root, _confirmAt)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_ReplicaManager *ReplicaManagerTransactor) SetUpdater(opts *bind.TransactOpts, _updater common.Address) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "setUpdater", _updater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_ReplicaManager *ReplicaManagerSession) SetUpdater(_updater common.Address) (*types.Transaction, error) {
	return _ReplicaManager.Contract.SetUpdater(&_ReplicaManager.TransactOpts, _updater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_ReplicaManager *ReplicaManagerTransactorSession) SetUpdater(_updater common.Address) (*types.Transaction, error) {
	return _ReplicaManager.Contract.SetUpdater(&_ReplicaManager.TransactOpts, _updater)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReplicaManager *ReplicaManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReplicaManager *ReplicaManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReplicaManager.Contract.TransferOwnership(&_ReplicaManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReplicaManager *ReplicaManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReplicaManager.Contract.TransferOwnership(&_ReplicaManager.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0xdec48b67.
//
// Solidity: function update(uint32 _remoteDomain, bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns()
func (_ReplicaManager *ReplicaManagerTransactor) Update(opts *bind.TransactOpts, _remoteDomain uint32, _oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ReplicaManager.contract.Transact(opts, "update", _remoteDomain, _oldRoot, _newRoot, _signature)
}

// Update is a paid mutator transaction binding the contract method 0xdec48b67.
//
// Solidity: function update(uint32 _remoteDomain, bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns()
func (_ReplicaManager *ReplicaManagerSession) Update(_remoteDomain uint32, _oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ReplicaManager.Contract.Update(&_ReplicaManager.TransactOpts, _remoteDomain, _oldRoot, _newRoot, _signature)
}

// Update is a paid mutator transaction binding the contract method 0xdec48b67.
//
// Solidity: function update(uint32 _remoteDomain, bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns()
func (_ReplicaManager *ReplicaManagerTransactorSession) Update(_remoteDomain uint32, _oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _ReplicaManager.Contract.Update(&_ReplicaManager.TransactOpts, _remoteDomain, _oldRoot, _newRoot, _signature)
}

// ReplicaManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ReplicaManager contract.
type ReplicaManagerInitializedIterator struct {
	Event *ReplicaManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ReplicaManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaManagerInitialized)
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
		it.Event = new(ReplicaManagerInitialized)
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
func (it *ReplicaManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaManagerInitialized represents a Initialized event raised by the ReplicaManager contract.
type ReplicaManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReplicaManager *ReplicaManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReplicaManagerInitializedIterator, error) {

	logs, sub, err := _ReplicaManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerInitializedIterator{contract: _ReplicaManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReplicaManager *ReplicaManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReplicaManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ReplicaManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaManagerInitialized)
				if err := _ReplicaManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReplicaManager *ReplicaManagerFilterer) ParseInitialized(log types.Log) (*ReplicaManagerInitialized, error) {
	event := new(ReplicaManagerInitialized)
	if err := _ReplicaManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaManagerNewUpdaterIterator is returned from FilterNewUpdater and is used to iterate over the raw logs and unpacked data for NewUpdater events raised by the ReplicaManager contract.
type ReplicaManagerNewUpdaterIterator struct {
	Event *ReplicaManagerNewUpdater // Event containing the contract specifics and raw log

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
func (it *ReplicaManagerNewUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaManagerNewUpdater)
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
		it.Event = new(ReplicaManagerNewUpdater)
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
func (it *ReplicaManagerNewUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaManagerNewUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaManagerNewUpdater represents a NewUpdater event raised by the ReplicaManager contract.
type ReplicaManagerNewUpdater struct {
	OldUpdater common.Address
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUpdater is a free log retrieval operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_ReplicaManager *ReplicaManagerFilterer) FilterNewUpdater(opts *bind.FilterOpts) (*ReplicaManagerNewUpdaterIterator, error) {

	logs, sub, err := _ReplicaManager.contract.FilterLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerNewUpdaterIterator{contract: _ReplicaManager.contract, event: "NewUpdater", logs: logs, sub: sub}, nil
}

// WatchNewUpdater is a free log subscription operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_ReplicaManager *ReplicaManagerFilterer) WatchNewUpdater(opts *bind.WatchOpts, sink chan<- *ReplicaManagerNewUpdater) (event.Subscription, error) {

	logs, sub, err := _ReplicaManager.contract.WatchLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaManagerNewUpdater)
				if err := _ReplicaManager.contract.UnpackLog(event, "NewUpdater", log); err != nil {
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

// ParseNewUpdater is a log parse operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_ReplicaManager *ReplicaManagerFilterer) ParseNewUpdater(log types.Log) (*ReplicaManagerNewUpdater, error) {
	event := new(ReplicaManagerNewUpdater)
	if err := _ReplicaManager.contract.UnpackLog(event, "NewUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ReplicaManager contract.
type ReplicaManagerOwnershipTransferredIterator struct {
	Event *ReplicaManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReplicaManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaManagerOwnershipTransferred)
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
		it.Event = new(ReplicaManagerOwnershipTransferred)
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
func (it *ReplicaManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ReplicaManager contract.
type ReplicaManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReplicaManager *ReplicaManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReplicaManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReplicaManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerOwnershipTransferredIterator{contract: _ReplicaManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReplicaManager *ReplicaManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReplicaManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReplicaManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaManagerOwnershipTransferred)
				if err := _ReplicaManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ReplicaManager *ReplicaManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ReplicaManagerOwnershipTransferred, error) {
	event := new(ReplicaManagerOwnershipTransferred)
	if err := _ReplicaManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaManagerProcessIterator is returned from FilterProcess and is used to iterate over the raw logs and unpacked data for Process events raised by the ReplicaManager contract.
type ReplicaManagerProcessIterator struct {
	Event *ReplicaManagerProcess // Event containing the contract specifics and raw log

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
func (it *ReplicaManagerProcessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaManagerProcess)
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
		it.Event = new(ReplicaManagerProcess)
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
func (it *ReplicaManagerProcessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaManagerProcessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaManagerProcess represents a Process event raised by the ReplicaManager contract.
type ReplicaManagerProcess struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Success      bool
	ReturnData   []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProcess is a free log retrieval operation binding the contract event 0x223de0966a99342a66dcd8e6b41362efb8e142d6ea63bca2fa73514df1d1f48f.
//
// Solidity: event Process(uint32 indexed remoteDomain, bytes32 indexed messageHash, bool indexed success, bytes returnData)
func (_ReplicaManager *ReplicaManagerFilterer) FilterProcess(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte, success []bool) (*ReplicaManagerProcessIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ReplicaManager.contract.FilterLogs(opts, "Process", remoteDomainRule, messageHashRule, successRule)
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerProcessIterator{contract: _ReplicaManager.contract, event: "Process", logs: logs, sub: sub}, nil
}

// WatchProcess is a free log subscription operation binding the contract event 0x223de0966a99342a66dcd8e6b41362efb8e142d6ea63bca2fa73514df1d1f48f.
//
// Solidity: event Process(uint32 indexed remoteDomain, bytes32 indexed messageHash, bool indexed success, bytes returnData)
func (_ReplicaManager *ReplicaManagerFilterer) WatchProcess(opts *bind.WatchOpts, sink chan<- *ReplicaManagerProcess, remoteDomain []uint32, messageHash [][32]byte, success []bool) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ReplicaManager.contract.WatchLogs(opts, "Process", remoteDomainRule, messageHashRule, successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaManagerProcess)
				if err := _ReplicaManager.contract.UnpackLog(event, "Process", log); err != nil {
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

// ParseProcess is a log parse operation binding the contract event 0x223de0966a99342a66dcd8e6b41362efb8e142d6ea63bca2fa73514df1d1f48f.
//
// Solidity: event Process(uint32 indexed remoteDomain, bytes32 indexed messageHash, bool indexed success, bytes returnData)
func (_ReplicaManager *ReplicaManagerFilterer) ParseProcess(log types.Log) (*ReplicaManagerProcess, error) {
	event := new(ReplicaManagerProcess)
	if err := _ReplicaManager.contract.UnpackLog(event, "Process", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaManagerSetConfirmationIterator is returned from FilterSetConfirmation and is used to iterate over the raw logs and unpacked data for SetConfirmation events raised by the ReplicaManager contract.
type ReplicaManagerSetConfirmationIterator struct {
	Event *ReplicaManagerSetConfirmation // Event containing the contract specifics and raw log

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
func (it *ReplicaManagerSetConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaManagerSetConfirmation)
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
		it.Event = new(ReplicaManagerSetConfirmation)
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
func (it *ReplicaManagerSetConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaManagerSetConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaManagerSetConfirmation represents a SetConfirmation event raised by the ReplicaManager contract.
type ReplicaManagerSetConfirmation struct {
	RemoteDomain      uint32
	Root              [32]byte
	PreviousConfirmAt *big.Int
	NewConfirmAt      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetConfirmation is a free log retrieval operation binding the contract event 0x6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730.
//
// Solidity: event SetConfirmation(uint32 indexed remoteDomain, bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_ReplicaManager *ReplicaManagerFilterer) FilterSetConfirmation(opts *bind.FilterOpts, remoteDomain []uint32, root [][32]byte) (*ReplicaManagerSetConfirmationIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ReplicaManager.contract.FilterLogs(opts, "SetConfirmation", remoteDomainRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerSetConfirmationIterator{contract: _ReplicaManager.contract, event: "SetConfirmation", logs: logs, sub: sub}, nil
}

// WatchSetConfirmation is a free log subscription operation binding the contract event 0x6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730.
//
// Solidity: event SetConfirmation(uint32 indexed remoteDomain, bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_ReplicaManager *ReplicaManagerFilterer) WatchSetConfirmation(opts *bind.WatchOpts, sink chan<- *ReplicaManagerSetConfirmation, remoteDomain []uint32, root [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ReplicaManager.contract.WatchLogs(opts, "SetConfirmation", remoteDomainRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaManagerSetConfirmation)
				if err := _ReplicaManager.contract.UnpackLog(event, "SetConfirmation", log); err != nil {
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

// ParseSetConfirmation is a log parse operation binding the contract event 0x6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730.
//
// Solidity: event SetConfirmation(uint32 indexed remoteDomain, bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_ReplicaManager *ReplicaManagerFilterer) ParseSetConfirmation(log types.Log) (*ReplicaManagerSetConfirmation, error) {
	event := new(ReplicaManagerSetConfirmation)
	if err := _ReplicaManager.contract.UnpackLog(event, "SetConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaManagerUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the ReplicaManager contract.
type ReplicaManagerUpdateIterator struct {
	Event *ReplicaManagerUpdate // Event containing the contract specifics and raw log

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
func (it *ReplicaManagerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaManagerUpdate)
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
		it.Event = new(ReplicaManagerUpdate)
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
func (it *ReplicaManagerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaManagerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaManagerUpdate represents a Update event raised by the ReplicaManager contract.
type ReplicaManagerUpdate struct {
	HomeDomain uint32
	OldRoot    [32]byte
	NewRoot    [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_ReplicaManager *ReplicaManagerFilterer) FilterUpdate(opts *bind.FilterOpts, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (*ReplicaManagerUpdateIterator, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var oldRootRule []interface{}
	for _, oldRootItem := range oldRoot {
		oldRootRule = append(oldRootRule, oldRootItem)
	}
	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}

	logs, sub, err := _ReplicaManager.contract.FilterLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return &ReplicaManagerUpdateIterator{contract: _ReplicaManager.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_ReplicaManager *ReplicaManagerFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *ReplicaManagerUpdate, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (event.Subscription, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var oldRootRule []interface{}
	for _, oldRootItem := range oldRoot {
		oldRootRule = append(oldRootRule, oldRootItem)
	}
	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}

	logs, sub, err := _ReplicaManager.contract.WatchLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaManagerUpdate)
				if err := _ReplicaManager.contract.UnpackLog(event, "Update", log); err != nil {
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

// ParseUpdate is a log parse operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_ReplicaManager *ReplicaManagerFilterer) ParseUpdate(log types.Log) (*ReplicaManagerUpdate, error) {
	event := new(ReplicaManagerUpdate)
	if err := _ReplicaManager.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207d6f4e0f6399708032ce1aa35108087d0fd847b9f9d81bdfe66f6696c75b750164736f6c634300080d0033",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// StringsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsMetaData.Bin instead.
var StringsBin = StringsMetaData.Bin

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220290bf685525c5d2b0b6e71963d9d52df53ebe472be87e00183a41933a8c7fc1664736f6c634300080d0033",
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

// TypedMemViewMetaData contains all meta data concerning the TypedMemView contract.
var TypedMemViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f26be3fc": "NULL()",
	},
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea2646970667358221220dbda491f76eadddf6980306838ba13d10fecceecbf12d8efad192d32faf34e4264736f6c634300080d0033",
}

// TypedMemViewABI is the input ABI used to generate the binding from.
// Deprecated: Use TypedMemViewMetaData.ABI instead.
var TypedMemViewABI = TypedMemViewMetaData.ABI

// Deprecated: Use TypedMemViewMetaData.Sigs instead.
// TypedMemViewFuncSigs maps the 4-byte function signature to its string representation.
var TypedMemViewFuncSigs = TypedMemViewMetaData.Sigs

// TypedMemViewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TypedMemViewMetaData.Bin instead.
var TypedMemViewBin = TypedMemViewMetaData.Bin

// DeployTypedMemView deploys a new Ethereum contract, binding an instance of TypedMemView to it.
func DeployTypedMemView(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TypedMemView, error) {
	parsed, err := TypedMemViewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TypedMemViewBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TypedMemView{TypedMemViewCaller: TypedMemViewCaller{contract: contract}, TypedMemViewTransactor: TypedMemViewTransactor{contract: contract}, TypedMemViewFilterer: TypedMemViewFilterer{contract: contract}}, nil
}

// TypedMemView is an auto generated Go binding around an Ethereum contract.
type TypedMemView struct {
	TypedMemViewCaller     // Read-only binding to the contract
	TypedMemViewTransactor // Write-only binding to the contract
	TypedMemViewFilterer   // Log filterer for contract events
}

// TypedMemViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypedMemViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypedMemViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypedMemViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypedMemViewSession struct {
	Contract     *TypedMemView     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypedMemViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypedMemViewCallerSession struct {
	Contract *TypedMemViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TypedMemViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypedMemViewTransactorSession struct {
	Contract     *TypedMemViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TypedMemViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypedMemViewRaw struct {
	Contract *TypedMemView // Generic contract binding to access the raw methods on
}

// TypedMemViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypedMemViewCallerRaw struct {
	Contract *TypedMemViewCaller // Generic read-only contract binding to access the raw methods on
}

// TypedMemViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypedMemViewTransactorRaw struct {
	Contract *TypedMemViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTypedMemView creates a new instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemView(address common.Address, backend bind.ContractBackend) (*TypedMemView, error) {
	contract, err := bindTypedMemView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TypedMemView{TypedMemViewCaller: TypedMemViewCaller{contract: contract}, TypedMemViewTransactor: TypedMemViewTransactor{contract: contract}, TypedMemViewFilterer: TypedMemViewFilterer{contract: contract}}, nil
}

// NewTypedMemViewCaller creates a new read-only instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewCaller(address common.Address, caller bind.ContractCaller) (*TypedMemViewCaller, error) {
	contract, err := bindTypedMemView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewCaller{contract: contract}, nil
}

// NewTypedMemViewTransactor creates a new write-only instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewTransactor(address common.Address, transactor bind.ContractTransactor) (*TypedMemViewTransactor, error) {
	contract, err := bindTypedMemView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewTransactor{contract: contract}, nil
}

// NewTypedMemViewFilterer creates a new log filterer instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewFilterer(address common.Address, filterer bind.ContractFilterer) (*TypedMemViewFilterer, error) {
	contract, err := bindTypedMemView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewFilterer{contract: contract}, nil
}

// bindTypedMemView binds a generic wrapper to an already deployed contract.
func bindTypedMemView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypedMemViewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypedMemView *TypedMemViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypedMemView.Contract.TypedMemViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypedMemView *TypedMemViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypedMemView.Contract.TypedMemViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypedMemView *TypedMemViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypedMemView.Contract.TypedMemViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypedMemView *TypedMemViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypedMemView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypedMemView *TypedMemViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypedMemView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypedMemView *TypedMemViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypedMemView.Contract.contract.Transact(opts, method, params...)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewCaller) NULL(opts *bind.CallOpts) ([29]byte, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "NULL")

	if err != nil {
		return *new([29]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([29]byte)).(*[29]byte)

	return out0, err

}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewSession) NULL() ([29]byte, error) {
	return _TypedMemView.Contract.NULL(&_TypedMemView.CallOpts)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewCallerSession) NULL() ([29]byte, error) {
	return _TypedMemView.Contract.NULL(&_TypedMemView.CallOpts)
}

// UpdaterStorageMetaData contains all meta data concerning the UpdaterStorage contract.
var UpdaterStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"df034cd0": "updater()",
	},
}

// UpdaterStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use UpdaterStorageMetaData.ABI instead.
var UpdaterStorageABI = UpdaterStorageMetaData.ABI

// Deprecated: Use UpdaterStorageMetaData.Sigs instead.
// UpdaterStorageFuncSigs maps the 4-byte function signature to its string representation.
var UpdaterStorageFuncSigs = UpdaterStorageMetaData.Sigs

// UpdaterStorage is an auto generated Go binding around an Ethereum contract.
type UpdaterStorage struct {
	UpdaterStorageCaller     // Read-only binding to the contract
	UpdaterStorageTransactor // Write-only binding to the contract
	UpdaterStorageFilterer   // Log filterer for contract events
}

// UpdaterStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpdaterStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdaterStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpdaterStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdaterStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpdaterStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdaterStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpdaterStorageSession struct {
	Contract     *UpdaterStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpdaterStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpdaterStorageCallerSession struct {
	Contract *UpdaterStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UpdaterStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpdaterStorageTransactorSession struct {
	Contract     *UpdaterStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UpdaterStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpdaterStorageRaw struct {
	Contract *UpdaterStorage // Generic contract binding to access the raw methods on
}

// UpdaterStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpdaterStorageCallerRaw struct {
	Contract *UpdaterStorageCaller // Generic read-only contract binding to access the raw methods on
}

// UpdaterStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpdaterStorageTransactorRaw struct {
	Contract *UpdaterStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpdaterStorage creates a new instance of UpdaterStorage, bound to a specific deployed contract.
func NewUpdaterStorage(address common.Address, backend bind.ContractBackend) (*UpdaterStorage, error) {
	contract, err := bindUpdaterStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpdaterStorage{UpdaterStorageCaller: UpdaterStorageCaller{contract: contract}, UpdaterStorageTransactor: UpdaterStorageTransactor{contract: contract}, UpdaterStorageFilterer: UpdaterStorageFilterer{contract: contract}}, nil
}

// NewUpdaterStorageCaller creates a new read-only instance of UpdaterStorage, bound to a specific deployed contract.
func NewUpdaterStorageCaller(address common.Address, caller bind.ContractCaller) (*UpdaterStorageCaller, error) {
	contract, err := bindUpdaterStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpdaterStorageCaller{contract: contract}, nil
}

// NewUpdaterStorageTransactor creates a new write-only instance of UpdaterStorage, bound to a specific deployed contract.
func NewUpdaterStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*UpdaterStorageTransactor, error) {
	contract, err := bindUpdaterStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpdaterStorageTransactor{contract: contract}, nil
}

// NewUpdaterStorageFilterer creates a new log filterer instance of UpdaterStorage, bound to a specific deployed contract.
func NewUpdaterStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*UpdaterStorageFilterer, error) {
	contract, err := bindUpdaterStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpdaterStorageFilterer{contract: contract}, nil
}

// bindUpdaterStorage binds a generic wrapper to an already deployed contract.
func bindUpdaterStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpdaterStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpdaterStorage *UpdaterStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpdaterStorage.Contract.UpdaterStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpdaterStorage *UpdaterStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpdaterStorage.Contract.UpdaterStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpdaterStorage *UpdaterStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpdaterStorage.Contract.UpdaterStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpdaterStorage *UpdaterStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpdaterStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpdaterStorage *UpdaterStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpdaterStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpdaterStorage *UpdaterStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpdaterStorage.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_UpdaterStorage *UpdaterStorageCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _UpdaterStorage.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_UpdaterStorage *UpdaterStorageSession) LocalDomain() (uint32, error) {
	return _UpdaterStorage.Contract.LocalDomain(&_UpdaterStorage.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_UpdaterStorage *UpdaterStorageCallerSession) LocalDomain() (uint32, error) {
	return _UpdaterStorage.Contract.LocalDomain(&_UpdaterStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpdaterStorage *UpdaterStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpdaterStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpdaterStorage *UpdaterStorageSession) Owner() (common.Address, error) {
	return _UpdaterStorage.Contract.Owner(&_UpdaterStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpdaterStorage *UpdaterStorageCallerSession) Owner() (common.Address, error) {
	return _UpdaterStorage.Contract.Owner(&_UpdaterStorage.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_UpdaterStorage *UpdaterStorageCaller) Updater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpdaterStorage.contract.Call(opts, &out, "updater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_UpdaterStorage *UpdaterStorageSession) Updater() (common.Address, error) {
	return _UpdaterStorage.Contract.Updater(&_UpdaterStorage.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_UpdaterStorage *UpdaterStorageCallerSession) Updater() (common.Address, error) {
	return _UpdaterStorage.Contract.Updater(&_UpdaterStorage.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpdaterStorage *UpdaterStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpdaterStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpdaterStorage *UpdaterStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpdaterStorage.Contract.RenounceOwnership(&_UpdaterStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpdaterStorage *UpdaterStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpdaterStorage.Contract.RenounceOwnership(&_UpdaterStorage.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpdaterStorage *UpdaterStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UpdaterStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpdaterStorage *UpdaterStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpdaterStorage.Contract.TransferOwnership(&_UpdaterStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpdaterStorage *UpdaterStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpdaterStorage.Contract.TransferOwnership(&_UpdaterStorage.TransactOpts, newOwner)
}

// UpdaterStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UpdaterStorage contract.
type UpdaterStorageInitializedIterator struct {
	Event *UpdaterStorageInitialized // Event containing the contract specifics and raw log

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
func (it *UpdaterStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdaterStorageInitialized)
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
		it.Event = new(UpdaterStorageInitialized)
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
func (it *UpdaterStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdaterStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdaterStorageInitialized represents a Initialized event raised by the UpdaterStorage contract.
type UpdaterStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpdaterStorage *UpdaterStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*UpdaterStorageInitializedIterator, error) {

	logs, sub, err := _UpdaterStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UpdaterStorageInitializedIterator{contract: _UpdaterStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpdaterStorage *UpdaterStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UpdaterStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _UpdaterStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdaterStorageInitialized)
				if err := _UpdaterStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpdaterStorage *UpdaterStorageFilterer) ParseInitialized(log types.Log) (*UpdaterStorageInitialized, error) {
	event := new(UpdaterStorageInitialized)
	if err := _UpdaterStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpdaterStorageNewUpdaterIterator is returned from FilterNewUpdater and is used to iterate over the raw logs and unpacked data for NewUpdater events raised by the UpdaterStorage contract.
type UpdaterStorageNewUpdaterIterator struct {
	Event *UpdaterStorageNewUpdater // Event containing the contract specifics and raw log

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
func (it *UpdaterStorageNewUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdaterStorageNewUpdater)
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
		it.Event = new(UpdaterStorageNewUpdater)
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
func (it *UpdaterStorageNewUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdaterStorageNewUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdaterStorageNewUpdater represents a NewUpdater event raised by the UpdaterStorage contract.
type UpdaterStorageNewUpdater struct {
	OldUpdater common.Address
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUpdater is a free log retrieval operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_UpdaterStorage *UpdaterStorageFilterer) FilterNewUpdater(opts *bind.FilterOpts) (*UpdaterStorageNewUpdaterIterator, error) {

	logs, sub, err := _UpdaterStorage.contract.FilterLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return &UpdaterStorageNewUpdaterIterator{contract: _UpdaterStorage.contract, event: "NewUpdater", logs: logs, sub: sub}, nil
}

// WatchNewUpdater is a free log subscription operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_UpdaterStorage *UpdaterStorageFilterer) WatchNewUpdater(opts *bind.WatchOpts, sink chan<- *UpdaterStorageNewUpdater) (event.Subscription, error) {

	logs, sub, err := _UpdaterStorage.contract.WatchLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdaterStorageNewUpdater)
				if err := _UpdaterStorage.contract.UnpackLog(event, "NewUpdater", log); err != nil {
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

// ParseNewUpdater is a log parse operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_UpdaterStorage *UpdaterStorageFilterer) ParseNewUpdater(log types.Log) (*UpdaterStorageNewUpdater, error) {
	event := new(UpdaterStorageNewUpdater)
	if err := _UpdaterStorage.contract.UnpackLog(event, "NewUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpdaterStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UpdaterStorage contract.
type UpdaterStorageOwnershipTransferredIterator struct {
	Event *UpdaterStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UpdaterStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdaterStorageOwnershipTransferred)
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
		it.Event = new(UpdaterStorageOwnershipTransferred)
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
func (it *UpdaterStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdaterStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdaterStorageOwnershipTransferred represents a OwnershipTransferred event raised by the UpdaterStorage contract.
type UpdaterStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpdaterStorage *UpdaterStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UpdaterStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpdaterStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UpdaterStorageOwnershipTransferredIterator{contract: _UpdaterStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpdaterStorage *UpdaterStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UpdaterStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpdaterStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdaterStorageOwnershipTransferred)
				if err := _UpdaterStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UpdaterStorage *UpdaterStorageFilterer) ParseOwnershipTransferred(log types.Log) (*UpdaterStorageOwnershipTransferred, error) {
	event := new(UpdaterStorageOwnershipTransferred)
	if err := _UpdaterStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpdaterStorageUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the UpdaterStorage contract.
type UpdaterStorageUpdateIterator struct {
	Event *UpdaterStorageUpdate // Event containing the contract specifics and raw log

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
func (it *UpdaterStorageUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdaterStorageUpdate)
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
		it.Event = new(UpdaterStorageUpdate)
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
func (it *UpdaterStorageUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdaterStorageUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdaterStorageUpdate represents a Update event raised by the UpdaterStorage contract.
type UpdaterStorageUpdate struct {
	HomeDomain uint32
	OldRoot    [32]byte
	NewRoot    [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_UpdaterStorage *UpdaterStorageFilterer) FilterUpdate(opts *bind.FilterOpts, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (*UpdaterStorageUpdateIterator, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var oldRootRule []interface{}
	for _, oldRootItem := range oldRoot {
		oldRootRule = append(oldRootRule, oldRootItem)
	}
	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}

	logs, sub, err := _UpdaterStorage.contract.FilterLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return &UpdaterStorageUpdateIterator{contract: _UpdaterStorage.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_UpdaterStorage *UpdaterStorageFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *UpdaterStorageUpdate, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (event.Subscription, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var oldRootRule []interface{}
	for _, oldRootItem := range oldRoot {
		oldRootRule = append(oldRootRule, oldRootItem)
	}
	var newRootRule []interface{}
	for _, newRootItem := range newRoot {
		newRootRule = append(newRootRule, newRootItem)
	}

	logs, sub, err := _UpdaterStorage.contract.WatchLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdaterStorageUpdate)
				if err := _UpdaterStorage.contract.UnpackLog(event, "Update", log); err != nil {
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

// ParseUpdate is a log parse operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_UpdaterStorage *UpdaterStorageFilterer) ParseUpdate(log types.Log) (*UpdaterStorageUpdate, error) {
	event := new(UpdaterStorageUpdate)
	if err := _UpdaterStorage.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Version0MetaData contains all meta data concerning the Version0 contract.
var Version0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ffa1ad74": "VERSION()",
	},
	Bin: "0x6080604052348015600f57600080fd5b5060808061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063ffa1ad7414602d575b600080fd5b6034600081565b60405160ff909116815260200160405180910390f3fea2646970667358221220b1d36cd12342d470d8f41375f6875d36e453c8767efb7a480bb349ad0ab3ddd164736f6c634300080d0033",
}

// Version0ABI is the input ABI used to generate the binding from.
// Deprecated: Use Version0MetaData.ABI instead.
var Version0ABI = Version0MetaData.ABI

// Deprecated: Use Version0MetaData.Sigs instead.
// Version0FuncSigs maps the 4-byte function signature to its string representation.
var Version0FuncSigs = Version0MetaData.Sigs

// Version0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Version0MetaData.Bin instead.
var Version0Bin = Version0MetaData.Bin

// DeployVersion0 deploys a new Ethereum contract, binding an instance of Version0 to it.
func DeployVersion0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Version0, error) {
	parsed, err := Version0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Version0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Version0{Version0Caller: Version0Caller{contract: contract}, Version0Transactor: Version0Transactor{contract: contract}, Version0Filterer: Version0Filterer{contract: contract}}, nil
}

// Version0 is an auto generated Go binding around an Ethereum contract.
type Version0 struct {
	Version0Caller     // Read-only binding to the contract
	Version0Transactor // Write-only binding to the contract
	Version0Filterer   // Log filterer for contract events
}

// Version0Caller is an auto generated read-only Go binding around an Ethereum contract.
type Version0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Version0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Version0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Version0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Version0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Version0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Version0Session struct {
	Contract     *Version0         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Version0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Version0CallerSession struct {
	Contract *Version0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Version0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Version0TransactorSession struct {
	Contract     *Version0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Version0Raw is an auto generated low-level Go binding around an Ethereum contract.
type Version0Raw struct {
	Contract *Version0 // Generic contract binding to access the raw methods on
}

// Version0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Version0CallerRaw struct {
	Contract *Version0Caller // Generic read-only contract binding to access the raw methods on
}

// Version0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Version0TransactorRaw struct {
	Contract *Version0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVersion0 creates a new instance of Version0, bound to a specific deployed contract.
func NewVersion0(address common.Address, backend bind.ContractBackend) (*Version0, error) {
	contract, err := bindVersion0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Version0{Version0Caller: Version0Caller{contract: contract}, Version0Transactor: Version0Transactor{contract: contract}, Version0Filterer: Version0Filterer{contract: contract}}, nil
}

// NewVersion0Caller creates a new read-only instance of Version0, bound to a specific deployed contract.
func NewVersion0Caller(address common.Address, caller bind.ContractCaller) (*Version0Caller, error) {
	contract, err := bindVersion0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Version0Caller{contract: contract}, nil
}

// NewVersion0Transactor creates a new write-only instance of Version0, bound to a specific deployed contract.
func NewVersion0Transactor(address common.Address, transactor bind.ContractTransactor) (*Version0Transactor, error) {
	contract, err := bindVersion0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Version0Transactor{contract: contract}, nil
}

// NewVersion0Filterer creates a new log filterer instance of Version0, bound to a specific deployed contract.
func NewVersion0Filterer(address common.Address, filterer bind.ContractFilterer) (*Version0Filterer, error) {
	contract, err := bindVersion0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Version0Filterer{contract: contract}, nil
}

// bindVersion0 binds a generic wrapper to an already deployed contract.
func bindVersion0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Version0ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Version0 *Version0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Version0.Contract.Version0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Version0 *Version0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Version0.Contract.Version0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Version0 *Version0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Version0.Contract.Version0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Version0 *Version0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Version0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Version0 *Version0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Version0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Version0 *Version0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Version0.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Version0 *Version0Caller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Version0.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Version0 *Version0Session) VERSION() (uint8, error) {
	return _Version0.Contract.VERSION(&_Version0.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Version0 *Version0CallerSession) VERSION() (uint8, error) {
	return _Version0.Contract.VERSION(&_Version0.CallOpts)
}

// XAppConfigMetaData contains all meta data concerning the XAppConfig contract.
var XAppConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"replica\",\"type\":\"address\"}],\"name\":\"ReplicaEnrolled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"replica\",\"type\":\"address\"}],\"name\":\"ReplicaUnenrolled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"watcher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"access\",\"type\":\"bool\"}],\"name\":\"WatcherPermissionSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"domainToReplica\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"home\",\"outputs\":[{\"internalType\":\"contractHome\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_replica\",\"type\":\"address\"}],\"name\":\"isReplica\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_replica\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"ownerEnrollReplica\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_replica\",\"type\":\"address\"}],\"name\":\"ownerUnenrollReplica\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"replicaToDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_home\",\"type\":\"address\"}],\"name\":\"setHome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_watcher\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"_access\",\"type\":\"bool\"}],\"name\":\"setWatcherPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_updater\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"unenrollReplica\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_watcher\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"watcherPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b9cff162": "domainToReplica(uint32)",
		"9fa92f9d": "home()",
		"5190bc53": "isReplica(address)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"f31faefb": "ownerEnrollReplica(address,uint32)",
		"8f5d90e0": "ownerUnenrollReplica(address)",
		"715018a6": "renounceOwnership()",
		"5f8b1dba": "replicaToDomain(address)",
		"6ef0f37f": "setHome(address)",
		"916c3470": "setWatcherPermission(address,uint32,bool)",
		"f2fde38b": "transferOwnership(address)",
		"e0e7a913": "unenrollReplica(uint32,bytes32,bytes)",
		"427ebef5": "watcherPermission(address,uint32)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6114078061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638f5d90e01161008c578063b9cff16211610066578063b9cff16214610275578063e0e7a913146102ab578063f2fde38b146102be578063f31faefb146102d157600080fd5b80638f5d90e01461022f578063916c3470146102425780639fa92f9d1461025557600080fd5b80636ef0f37f116100c85780636ef0f37f146101cb578063715018a6146101e05780638d3638f4146101e85780638da5cb5b146101f057600080fd5b8063427ebef5146100ef5780635190bc53146101525780635f8b1dba14610190575b600080fd5b61013d6100fd36600461112c565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260046020908152604080832063ffffffff8516845290915290205460ff1692915050565b60405190151581526020015b60405180910390f35b61013d610160366004611165565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205463ffffffff16151590565b6101b661019e366004611165565b60026020526000908152604090205463ffffffff1681565b60405163ffffffff9091168152602001610149565b6101de6101d9366004611165565b6102e4565b005b6101de610397565b6101b6610400565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610149565b6101de61023d366004611165565b610499565b6101de610250366004611189565b61050c565b60015461020a9073ffffffffffffffffffffffffffffffffffffffff1681565b61020a6102833660046111d9565b60036020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6101de6102b9366004611225565b610617565b6101de6102cc366004611165565b610836565b6101de6102df36600461112c565b61092f565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103505760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103fe5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b565b600154604080517f8d3638f4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691638d3638f49160048083019260209291908290030181865afa158015610470573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104949190611310565b905090565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105005760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b61050981610a61565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b73ffffffffffffffffffffffffffffffffffffffff8316600081815260046020908152604080832063ffffffff87168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915581519485529184019190915290917f517de16b526853f481451c5151e87484e1b251ec7d0302efa1019c2ece179c2c910160405180910390a2505050565b63ffffffff831660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16806106905760405162461bcd60e51b815260206004820152600f60248201527f217265706c6963612065786973747300000000000000000000000000000000006044820152606401610347565b8273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663df034cd06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610716919061132d565b73ffffffffffffffffffffffffffffffffffffffff16146107795760405162461bcd60e51b815260206004820152601060248201527f2163757272656e742075706461746572000000000000000000000000000000006044820152606401610347565b600061079d8573ffffffffffffffffffffffffffffffffffffffff84168686610b20565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260046020908152604080832063ffffffff8a16845290915290205490915060ff166108265760405162461bcd60e51b815260206004820152600e60248201527f2176616c696420776174636865720000000000000000000000000000000000006044820152606401610347565b61082f82610a61565b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461089d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b73ffffffffffffffffffffffffffffffffffffffff81166109265760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610347565b61050981610c9b565b60005473ffffffffffffffffffffffffffffffffffffffff1633146109965760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b61099f82610a61565b73ffffffffffffffffffffffffffffffffffffffff8216600081815260026020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8716908117909155808452600383529281902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016851790555192835290917f8440df9bf8a8542634a9eb196da1735b786ed9aa2fc12b080ac34c5fa81a923491015b60405180910390a25050565b73ffffffffffffffffffffffffffffffffffffffff81166000818152600260208181526040808420805463ffffffff168086526003845282862080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690559486905292825282547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000169092559051928352909182917fce1533133fb359ace801d3176bbad25ace030d714aed35e38a6293c8a60b115b9101610a55565b600080846040517ffd74954600000000000000000000000000000000000000000000000000000000815263ffffffff8816600482015273ffffffffffffffffffffffffffffffffffffffff919091169063fd74954690602401602060405180830381865afa158015610b96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bba919061134a565b90506000818786604051602001610c099392919092835260e09190911b7fffffffff00000000000000000000000000000000000000000000000000000000166020830152602482015260440190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815282825280516020918201207f19457468657265756d205369676e6564204d6573736167653a0a33320000000082850152603c8085019190915282518085039091018152605c90930190915281519101209050610c908185610d10565b979650505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806000610d1f8585610d34565b91509150610d2c81610da2565b509392505050565b6000808251604103610d6a5760208301516040840151606085015160001a610d5e87828585610f8e565b94509450505050610d9b565b8251604003610d935760208301516040840151610d888683836110a6565b935093505050610d9b565b506000905060025b9250929050565b6000816004811115610db657610db6611363565b03610dbe5750565b6001816004811115610dd257610dd2611363565b03610e1f5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610347565b6002816004811115610e3357610e33611363565b03610e805760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610347565b6003816004811115610e9457610e94611363565b03610f075760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610347565b6004816004811115610f1b57610f1b611363565b036105095760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610347565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610fc5575060009050600361109d565b8460ff16601b14158015610fdd57508460ff16601c14155b15610fee575060009050600461109d565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611042573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166110965760006001925092505061109d565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316816110dc60ff86901c601b611392565b90506110ea87828885610f8e565b935093505050935093915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461050957600080fd5b63ffffffff8116811461050957600080fd5b6000806040838503121561113f57600080fd5b823561114a816110f8565b9150602083013561115a8161111a565b809150509250929050565b60006020828403121561117757600080fd5b8135611182816110f8565b9392505050565b60008060006060848603121561119e57600080fd5b83356111a9816110f8565b925060208401356111b98161111a565b9150604084013580151581146111ce57600080fd5b809150509250925092565b6000602082840312156111eb57600080fd5b81356111828161111a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561123a57600080fd5b83356112458161111a565b925060208401359150604084013567ffffffffffffffff8082111561126957600080fd5b818601915086601f83011261127d57600080fd5b81358181111561128f5761128f6111f6565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156112d5576112d56111f6565b816040528281528960208487010111156112ee57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60006020828403121561132257600080fd5b81516111828161111a565b60006020828403121561133f57600080fd5b8151611182816110f8565b60006020828403121561135c57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600082198211156113cc577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50019056fea2646970667358221220863add28eb90748be404231844e184279078b47012f6910655f4e8253a0a9ff464736f6c634300080d0033",
}

// XAppConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use XAppConfigMetaData.ABI instead.
var XAppConfigABI = XAppConfigMetaData.ABI

// Deprecated: Use XAppConfigMetaData.Sigs instead.
// XAppConfigFuncSigs maps the 4-byte function signature to its string representation.
var XAppConfigFuncSigs = XAppConfigMetaData.Sigs

// XAppConfigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use XAppConfigMetaData.Bin instead.
var XAppConfigBin = XAppConfigMetaData.Bin

// DeployXAppConfig deploys a new Ethereum contract, binding an instance of XAppConfig to it.
func DeployXAppConfig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XAppConfig, error) {
	parsed, err := XAppConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(XAppConfigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XAppConfig{XAppConfigCaller: XAppConfigCaller{contract: contract}, XAppConfigTransactor: XAppConfigTransactor{contract: contract}, XAppConfigFilterer: XAppConfigFilterer{contract: contract}}, nil
}

// XAppConfig is an auto generated Go binding around an Ethereum contract.
type XAppConfig struct {
	XAppConfigCaller     // Read-only binding to the contract
	XAppConfigTransactor // Write-only binding to the contract
	XAppConfigFilterer   // Log filterer for contract events
}

// XAppConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type XAppConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XAppConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XAppConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XAppConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XAppConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XAppConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XAppConfigSession struct {
	Contract     *XAppConfig       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XAppConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XAppConfigCallerSession struct {
	Contract *XAppConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// XAppConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XAppConfigTransactorSession struct {
	Contract     *XAppConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// XAppConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type XAppConfigRaw struct {
	Contract *XAppConfig // Generic contract binding to access the raw methods on
}

// XAppConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XAppConfigCallerRaw struct {
	Contract *XAppConfigCaller // Generic read-only contract binding to access the raw methods on
}

// XAppConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XAppConfigTransactorRaw struct {
	Contract *XAppConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXAppConfig creates a new instance of XAppConfig, bound to a specific deployed contract.
func NewXAppConfig(address common.Address, backend bind.ContractBackend) (*XAppConfig, error) {
	contract, err := bindXAppConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XAppConfig{XAppConfigCaller: XAppConfigCaller{contract: contract}, XAppConfigTransactor: XAppConfigTransactor{contract: contract}, XAppConfigFilterer: XAppConfigFilterer{contract: contract}}, nil
}

// NewXAppConfigCaller creates a new read-only instance of XAppConfig, bound to a specific deployed contract.
func NewXAppConfigCaller(address common.Address, caller bind.ContractCaller) (*XAppConfigCaller, error) {
	contract, err := bindXAppConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XAppConfigCaller{contract: contract}, nil
}

// NewXAppConfigTransactor creates a new write-only instance of XAppConfig, bound to a specific deployed contract.
func NewXAppConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*XAppConfigTransactor, error) {
	contract, err := bindXAppConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XAppConfigTransactor{contract: contract}, nil
}

// NewXAppConfigFilterer creates a new log filterer instance of XAppConfig, bound to a specific deployed contract.
func NewXAppConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*XAppConfigFilterer, error) {
	contract, err := bindXAppConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XAppConfigFilterer{contract: contract}, nil
}

// bindXAppConfig binds a generic wrapper to an already deployed contract.
func bindXAppConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XAppConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XAppConfig *XAppConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XAppConfig.Contract.XAppConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XAppConfig *XAppConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XAppConfig.Contract.XAppConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XAppConfig *XAppConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XAppConfig.Contract.XAppConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XAppConfig *XAppConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XAppConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XAppConfig *XAppConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XAppConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XAppConfig *XAppConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XAppConfig.Contract.contract.Transact(opts, method, params...)
}

// DomainToReplica is a free data retrieval call binding the contract method 0xb9cff162.
//
// Solidity: function domainToReplica(uint32 ) view returns(address)
func (_XAppConfig *XAppConfigCaller) DomainToReplica(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _XAppConfig.contract.Call(opts, &out, "domainToReplica", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DomainToReplica is a free data retrieval call binding the contract method 0xb9cff162.
//
// Solidity: function domainToReplica(uint32 ) view returns(address)
func (_XAppConfig *XAppConfigSession) DomainToReplica(arg0 uint32) (common.Address, error) {
	return _XAppConfig.Contract.DomainToReplica(&_XAppConfig.CallOpts, arg0)
}

// DomainToReplica is a free data retrieval call binding the contract method 0xb9cff162.
//
// Solidity: function domainToReplica(uint32 ) view returns(address)
func (_XAppConfig *XAppConfigCallerSession) DomainToReplica(arg0 uint32) (common.Address, error) {
	return _XAppConfig.Contract.DomainToReplica(&_XAppConfig.CallOpts, arg0)
}

// Home is a free data retrieval call binding the contract method 0x9fa92f9d.
//
// Solidity: function home() view returns(address)
func (_XAppConfig *XAppConfigCaller) Home(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XAppConfig.contract.Call(opts, &out, "home")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Home is a free data retrieval call binding the contract method 0x9fa92f9d.
//
// Solidity: function home() view returns(address)
func (_XAppConfig *XAppConfigSession) Home() (common.Address, error) {
	return _XAppConfig.Contract.Home(&_XAppConfig.CallOpts)
}

// Home is a free data retrieval call binding the contract method 0x9fa92f9d.
//
// Solidity: function home() view returns(address)
func (_XAppConfig *XAppConfigCallerSession) Home() (common.Address, error) {
	return _XAppConfig.Contract.Home(&_XAppConfig.CallOpts)
}

// IsReplica is a free data retrieval call binding the contract method 0x5190bc53.
//
// Solidity: function isReplica(address _replica) view returns(bool)
func (_XAppConfig *XAppConfigCaller) IsReplica(opts *bind.CallOpts, _replica common.Address) (bool, error) {
	var out []interface{}
	err := _XAppConfig.contract.Call(opts, &out, "isReplica", _replica)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReplica is a free data retrieval call binding the contract method 0x5190bc53.
//
// Solidity: function isReplica(address _replica) view returns(bool)
func (_XAppConfig *XAppConfigSession) IsReplica(_replica common.Address) (bool, error) {
	return _XAppConfig.Contract.IsReplica(&_XAppConfig.CallOpts, _replica)
}

// IsReplica is a free data retrieval call binding the contract method 0x5190bc53.
//
// Solidity: function isReplica(address _replica) view returns(bool)
func (_XAppConfig *XAppConfigCallerSession) IsReplica(_replica common.Address) (bool, error) {
	return _XAppConfig.Contract.IsReplica(&_XAppConfig.CallOpts, _replica)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_XAppConfig *XAppConfigCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _XAppConfig.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_XAppConfig *XAppConfigSession) LocalDomain() (uint32, error) {
	return _XAppConfig.Contract.LocalDomain(&_XAppConfig.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_XAppConfig *XAppConfigCallerSession) LocalDomain() (uint32, error) {
	return _XAppConfig.Contract.LocalDomain(&_XAppConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XAppConfig *XAppConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XAppConfig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XAppConfig *XAppConfigSession) Owner() (common.Address, error) {
	return _XAppConfig.Contract.Owner(&_XAppConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XAppConfig *XAppConfigCallerSession) Owner() (common.Address, error) {
	return _XAppConfig.Contract.Owner(&_XAppConfig.CallOpts)
}

// ReplicaToDomain is a free data retrieval call binding the contract method 0x5f8b1dba.
//
// Solidity: function replicaToDomain(address ) view returns(uint32)
func (_XAppConfig *XAppConfigCaller) ReplicaToDomain(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _XAppConfig.contract.Call(opts, &out, "replicaToDomain", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ReplicaToDomain is a free data retrieval call binding the contract method 0x5f8b1dba.
//
// Solidity: function replicaToDomain(address ) view returns(uint32)
func (_XAppConfig *XAppConfigSession) ReplicaToDomain(arg0 common.Address) (uint32, error) {
	return _XAppConfig.Contract.ReplicaToDomain(&_XAppConfig.CallOpts, arg0)
}

// ReplicaToDomain is a free data retrieval call binding the contract method 0x5f8b1dba.
//
// Solidity: function replicaToDomain(address ) view returns(uint32)
func (_XAppConfig *XAppConfigCallerSession) ReplicaToDomain(arg0 common.Address) (uint32, error) {
	return _XAppConfig.Contract.ReplicaToDomain(&_XAppConfig.CallOpts, arg0)
}

// WatcherPermission is a free data retrieval call binding the contract method 0x427ebef5.
//
// Solidity: function watcherPermission(address _watcher, uint32 _domain) view returns(bool)
func (_XAppConfig *XAppConfigCaller) WatcherPermission(opts *bind.CallOpts, _watcher common.Address, _domain uint32) (bool, error) {
	var out []interface{}
	err := _XAppConfig.contract.Call(opts, &out, "watcherPermission", _watcher, _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WatcherPermission is a free data retrieval call binding the contract method 0x427ebef5.
//
// Solidity: function watcherPermission(address _watcher, uint32 _domain) view returns(bool)
func (_XAppConfig *XAppConfigSession) WatcherPermission(_watcher common.Address, _domain uint32) (bool, error) {
	return _XAppConfig.Contract.WatcherPermission(&_XAppConfig.CallOpts, _watcher, _domain)
}

// WatcherPermission is a free data retrieval call binding the contract method 0x427ebef5.
//
// Solidity: function watcherPermission(address _watcher, uint32 _domain) view returns(bool)
func (_XAppConfig *XAppConfigCallerSession) WatcherPermission(_watcher common.Address, _domain uint32) (bool, error) {
	return _XAppConfig.Contract.WatcherPermission(&_XAppConfig.CallOpts, _watcher, _domain)
}

// OwnerEnrollReplica is a paid mutator transaction binding the contract method 0xf31faefb.
//
// Solidity: function ownerEnrollReplica(address _replica, uint32 _domain) returns()
func (_XAppConfig *XAppConfigTransactor) OwnerEnrollReplica(opts *bind.TransactOpts, _replica common.Address, _domain uint32) (*types.Transaction, error) {
	return _XAppConfig.contract.Transact(opts, "ownerEnrollReplica", _replica, _domain)
}

// OwnerEnrollReplica is a paid mutator transaction binding the contract method 0xf31faefb.
//
// Solidity: function ownerEnrollReplica(address _replica, uint32 _domain) returns()
func (_XAppConfig *XAppConfigSession) OwnerEnrollReplica(_replica common.Address, _domain uint32) (*types.Transaction, error) {
	return _XAppConfig.Contract.OwnerEnrollReplica(&_XAppConfig.TransactOpts, _replica, _domain)
}

// OwnerEnrollReplica is a paid mutator transaction binding the contract method 0xf31faefb.
//
// Solidity: function ownerEnrollReplica(address _replica, uint32 _domain) returns()
func (_XAppConfig *XAppConfigTransactorSession) OwnerEnrollReplica(_replica common.Address, _domain uint32) (*types.Transaction, error) {
	return _XAppConfig.Contract.OwnerEnrollReplica(&_XAppConfig.TransactOpts, _replica, _domain)
}

// OwnerUnenrollReplica is a paid mutator transaction binding the contract method 0x8f5d90e0.
//
// Solidity: function ownerUnenrollReplica(address _replica) returns()
func (_XAppConfig *XAppConfigTransactor) OwnerUnenrollReplica(opts *bind.TransactOpts, _replica common.Address) (*types.Transaction, error) {
	return _XAppConfig.contract.Transact(opts, "ownerUnenrollReplica", _replica)
}

// OwnerUnenrollReplica is a paid mutator transaction binding the contract method 0x8f5d90e0.
//
// Solidity: function ownerUnenrollReplica(address _replica) returns()
func (_XAppConfig *XAppConfigSession) OwnerUnenrollReplica(_replica common.Address) (*types.Transaction, error) {
	return _XAppConfig.Contract.OwnerUnenrollReplica(&_XAppConfig.TransactOpts, _replica)
}

// OwnerUnenrollReplica is a paid mutator transaction binding the contract method 0x8f5d90e0.
//
// Solidity: function ownerUnenrollReplica(address _replica) returns()
func (_XAppConfig *XAppConfigTransactorSession) OwnerUnenrollReplica(_replica common.Address) (*types.Transaction, error) {
	return _XAppConfig.Contract.OwnerUnenrollReplica(&_XAppConfig.TransactOpts, _replica)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XAppConfig *XAppConfigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XAppConfig.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XAppConfig *XAppConfigSession) RenounceOwnership() (*types.Transaction, error) {
	return _XAppConfig.Contract.RenounceOwnership(&_XAppConfig.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XAppConfig *XAppConfigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _XAppConfig.Contract.RenounceOwnership(&_XAppConfig.TransactOpts)
}

// SetHome is a paid mutator transaction binding the contract method 0x6ef0f37f.
//
// Solidity: function setHome(address _home) returns()
func (_XAppConfig *XAppConfigTransactor) SetHome(opts *bind.TransactOpts, _home common.Address) (*types.Transaction, error) {
	return _XAppConfig.contract.Transact(opts, "setHome", _home)
}

// SetHome is a paid mutator transaction binding the contract method 0x6ef0f37f.
//
// Solidity: function setHome(address _home) returns()
func (_XAppConfig *XAppConfigSession) SetHome(_home common.Address) (*types.Transaction, error) {
	return _XAppConfig.Contract.SetHome(&_XAppConfig.TransactOpts, _home)
}

// SetHome is a paid mutator transaction binding the contract method 0x6ef0f37f.
//
// Solidity: function setHome(address _home) returns()
func (_XAppConfig *XAppConfigTransactorSession) SetHome(_home common.Address) (*types.Transaction, error) {
	return _XAppConfig.Contract.SetHome(&_XAppConfig.TransactOpts, _home)
}

// SetWatcherPermission is a paid mutator transaction binding the contract method 0x916c3470.
//
// Solidity: function setWatcherPermission(address _watcher, uint32 _domain, bool _access) returns()
func (_XAppConfig *XAppConfigTransactor) SetWatcherPermission(opts *bind.TransactOpts, _watcher common.Address, _domain uint32, _access bool) (*types.Transaction, error) {
	return _XAppConfig.contract.Transact(opts, "setWatcherPermission", _watcher, _domain, _access)
}

// SetWatcherPermission is a paid mutator transaction binding the contract method 0x916c3470.
//
// Solidity: function setWatcherPermission(address _watcher, uint32 _domain, bool _access) returns()
func (_XAppConfig *XAppConfigSession) SetWatcherPermission(_watcher common.Address, _domain uint32, _access bool) (*types.Transaction, error) {
	return _XAppConfig.Contract.SetWatcherPermission(&_XAppConfig.TransactOpts, _watcher, _domain, _access)
}

// SetWatcherPermission is a paid mutator transaction binding the contract method 0x916c3470.
//
// Solidity: function setWatcherPermission(address _watcher, uint32 _domain, bool _access) returns()
func (_XAppConfig *XAppConfigTransactorSession) SetWatcherPermission(_watcher common.Address, _domain uint32, _access bool) (*types.Transaction, error) {
	return _XAppConfig.Contract.SetWatcherPermission(&_XAppConfig.TransactOpts, _watcher, _domain, _access)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XAppConfig *XAppConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _XAppConfig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XAppConfig *XAppConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _XAppConfig.Contract.TransferOwnership(&_XAppConfig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XAppConfig *XAppConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _XAppConfig.Contract.TransferOwnership(&_XAppConfig.TransactOpts, newOwner)
}

// UnenrollReplica is a paid mutator transaction binding the contract method 0xe0e7a913.
//
// Solidity: function unenrollReplica(uint32 _domain, bytes32 _updater, bytes _signature) returns()
func (_XAppConfig *XAppConfigTransactor) UnenrollReplica(opts *bind.TransactOpts, _domain uint32, _updater [32]byte, _signature []byte) (*types.Transaction, error) {
	return _XAppConfig.contract.Transact(opts, "unenrollReplica", _domain, _updater, _signature)
}

// UnenrollReplica is a paid mutator transaction binding the contract method 0xe0e7a913.
//
// Solidity: function unenrollReplica(uint32 _domain, bytes32 _updater, bytes _signature) returns()
func (_XAppConfig *XAppConfigSession) UnenrollReplica(_domain uint32, _updater [32]byte, _signature []byte) (*types.Transaction, error) {
	return _XAppConfig.Contract.UnenrollReplica(&_XAppConfig.TransactOpts, _domain, _updater, _signature)
}

// UnenrollReplica is a paid mutator transaction binding the contract method 0xe0e7a913.
//
// Solidity: function unenrollReplica(uint32 _domain, bytes32 _updater, bytes _signature) returns()
func (_XAppConfig *XAppConfigTransactorSession) UnenrollReplica(_domain uint32, _updater [32]byte, _signature []byte) (*types.Transaction, error) {
	return _XAppConfig.Contract.UnenrollReplica(&_XAppConfig.TransactOpts, _domain, _updater, _signature)
}

// XAppConfigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the XAppConfig contract.
type XAppConfigOwnershipTransferredIterator struct {
	Event *XAppConfigOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *XAppConfigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XAppConfigOwnershipTransferred)
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
		it.Event = new(XAppConfigOwnershipTransferred)
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
func (it *XAppConfigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XAppConfigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XAppConfigOwnershipTransferred represents a OwnershipTransferred event raised by the XAppConfig contract.
type XAppConfigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XAppConfig *XAppConfigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*XAppConfigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _XAppConfig.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &XAppConfigOwnershipTransferredIterator{contract: _XAppConfig.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XAppConfig *XAppConfigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *XAppConfigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _XAppConfig.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XAppConfigOwnershipTransferred)
				if err := _XAppConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_XAppConfig *XAppConfigFilterer) ParseOwnershipTransferred(log types.Log) (*XAppConfigOwnershipTransferred, error) {
	event := new(XAppConfigOwnershipTransferred)
	if err := _XAppConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XAppConfigReplicaEnrolledIterator is returned from FilterReplicaEnrolled and is used to iterate over the raw logs and unpacked data for ReplicaEnrolled events raised by the XAppConfig contract.
type XAppConfigReplicaEnrolledIterator struct {
	Event *XAppConfigReplicaEnrolled // Event containing the contract specifics and raw log

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
func (it *XAppConfigReplicaEnrolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XAppConfigReplicaEnrolled)
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
		it.Event = new(XAppConfigReplicaEnrolled)
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
func (it *XAppConfigReplicaEnrolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XAppConfigReplicaEnrolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XAppConfigReplicaEnrolled represents a ReplicaEnrolled event raised by the XAppConfig contract.
type XAppConfigReplicaEnrolled struct {
	Domain  uint32
	Replica common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReplicaEnrolled is a free log retrieval operation binding the contract event 0x8440df9bf8a8542634a9eb196da1735b786ed9aa2fc12b080ac34c5fa81a9234.
//
// Solidity: event ReplicaEnrolled(uint32 indexed domain, address replica)
func (_XAppConfig *XAppConfigFilterer) FilterReplicaEnrolled(opts *bind.FilterOpts, domain []uint32) (*XAppConfigReplicaEnrolledIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _XAppConfig.contract.FilterLogs(opts, "ReplicaEnrolled", domainRule)
	if err != nil {
		return nil, err
	}
	return &XAppConfigReplicaEnrolledIterator{contract: _XAppConfig.contract, event: "ReplicaEnrolled", logs: logs, sub: sub}, nil
}

// WatchReplicaEnrolled is a free log subscription operation binding the contract event 0x8440df9bf8a8542634a9eb196da1735b786ed9aa2fc12b080ac34c5fa81a9234.
//
// Solidity: event ReplicaEnrolled(uint32 indexed domain, address replica)
func (_XAppConfig *XAppConfigFilterer) WatchReplicaEnrolled(opts *bind.WatchOpts, sink chan<- *XAppConfigReplicaEnrolled, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _XAppConfig.contract.WatchLogs(opts, "ReplicaEnrolled", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XAppConfigReplicaEnrolled)
				if err := _XAppConfig.contract.UnpackLog(event, "ReplicaEnrolled", log); err != nil {
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

// ParseReplicaEnrolled is a log parse operation binding the contract event 0x8440df9bf8a8542634a9eb196da1735b786ed9aa2fc12b080ac34c5fa81a9234.
//
// Solidity: event ReplicaEnrolled(uint32 indexed domain, address replica)
func (_XAppConfig *XAppConfigFilterer) ParseReplicaEnrolled(log types.Log) (*XAppConfigReplicaEnrolled, error) {
	event := new(XAppConfigReplicaEnrolled)
	if err := _XAppConfig.contract.UnpackLog(event, "ReplicaEnrolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XAppConfigReplicaUnenrolledIterator is returned from FilterReplicaUnenrolled and is used to iterate over the raw logs and unpacked data for ReplicaUnenrolled events raised by the XAppConfig contract.
type XAppConfigReplicaUnenrolledIterator struct {
	Event *XAppConfigReplicaUnenrolled // Event containing the contract specifics and raw log

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
func (it *XAppConfigReplicaUnenrolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XAppConfigReplicaUnenrolled)
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
		it.Event = new(XAppConfigReplicaUnenrolled)
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
func (it *XAppConfigReplicaUnenrolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XAppConfigReplicaUnenrolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XAppConfigReplicaUnenrolled represents a ReplicaUnenrolled event raised by the XAppConfig contract.
type XAppConfigReplicaUnenrolled struct {
	Domain  uint32
	Replica common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReplicaUnenrolled is a free log retrieval operation binding the contract event 0xce1533133fb359ace801d3176bbad25ace030d714aed35e38a6293c8a60b115b.
//
// Solidity: event ReplicaUnenrolled(uint32 indexed domain, address replica)
func (_XAppConfig *XAppConfigFilterer) FilterReplicaUnenrolled(opts *bind.FilterOpts, domain []uint32) (*XAppConfigReplicaUnenrolledIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _XAppConfig.contract.FilterLogs(opts, "ReplicaUnenrolled", domainRule)
	if err != nil {
		return nil, err
	}
	return &XAppConfigReplicaUnenrolledIterator{contract: _XAppConfig.contract, event: "ReplicaUnenrolled", logs: logs, sub: sub}, nil
}

// WatchReplicaUnenrolled is a free log subscription operation binding the contract event 0xce1533133fb359ace801d3176bbad25ace030d714aed35e38a6293c8a60b115b.
//
// Solidity: event ReplicaUnenrolled(uint32 indexed domain, address replica)
func (_XAppConfig *XAppConfigFilterer) WatchReplicaUnenrolled(opts *bind.WatchOpts, sink chan<- *XAppConfigReplicaUnenrolled, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _XAppConfig.contract.WatchLogs(opts, "ReplicaUnenrolled", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XAppConfigReplicaUnenrolled)
				if err := _XAppConfig.contract.UnpackLog(event, "ReplicaUnenrolled", log); err != nil {
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

// ParseReplicaUnenrolled is a log parse operation binding the contract event 0xce1533133fb359ace801d3176bbad25ace030d714aed35e38a6293c8a60b115b.
//
// Solidity: event ReplicaUnenrolled(uint32 indexed domain, address replica)
func (_XAppConfig *XAppConfigFilterer) ParseReplicaUnenrolled(log types.Log) (*XAppConfigReplicaUnenrolled, error) {
	event := new(XAppConfigReplicaUnenrolled)
	if err := _XAppConfig.contract.UnpackLog(event, "ReplicaUnenrolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XAppConfigWatcherPermissionSetIterator is returned from FilterWatcherPermissionSet and is used to iterate over the raw logs and unpacked data for WatcherPermissionSet events raised by the XAppConfig contract.
type XAppConfigWatcherPermissionSetIterator struct {
	Event *XAppConfigWatcherPermissionSet // Event containing the contract specifics and raw log

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
func (it *XAppConfigWatcherPermissionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XAppConfigWatcherPermissionSet)
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
		it.Event = new(XAppConfigWatcherPermissionSet)
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
func (it *XAppConfigWatcherPermissionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XAppConfigWatcherPermissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XAppConfigWatcherPermissionSet represents a WatcherPermissionSet event raised by the XAppConfig contract.
type XAppConfigWatcherPermissionSet struct {
	Domain  uint32
	Watcher common.Address
	Access  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWatcherPermissionSet is a free log retrieval operation binding the contract event 0x517de16b526853f481451c5151e87484e1b251ec7d0302efa1019c2ece179c2c.
//
// Solidity: event WatcherPermissionSet(uint32 indexed domain, address watcher, bool access)
func (_XAppConfig *XAppConfigFilterer) FilterWatcherPermissionSet(opts *bind.FilterOpts, domain []uint32) (*XAppConfigWatcherPermissionSetIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _XAppConfig.contract.FilterLogs(opts, "WatcherPermissionSet", domainRule)
	if err != nil {
		return nil, err
	}
	return &XAppConfigWatcherPermissionSetIterator{contract: _XAppConfig.contract, event: "WatcherPermissionSet", logs: logs, sub: sub}, nil
}

// WatchWatcherPermissionSet is a free log subscription operation binding the contract event 0x517de16b526853f481451c5151e87484e1b251ec7d0302efa1019c2ece179c2c.
//
// Solidity: event WatcherPermissionSet(uint32 indexed domain, address watcher, bool access)
func (_XAppConfig *XAppConfigFilterer) WatchWatcherPermissionSet(opts *bind.WatchOpts, sink chan<- *XAppConfigWatcherPermissionSet, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _XAppConfig.contract.WatchLogs(opts, "WatcherPermissionSet", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XAppConfigWatcherPermissionSet)
				if err := _XAppConfig.contract.UnpackLog(event, "WatcherPermissionSet", log); err != nil {
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

// ParseWatcherPermissionSet is a log parse operation binding the contract event 0x517de16b526853f481451c5151e87484e1b251ec7d0302efa1019c2ece179c2c.
//
// Solidity: event WatcherPermissionSet(uint32 indexed domain, address watcher, bool access)
func (_XAppConfig *XAppConfigFilterer) ParseWatcherPermissionSet(log types.Log) (*XAppConfigWatcherPermissionSet, error) {
	event := new(XAppConfigWatcherPermissionSet)
	if err := _XAppConfig.contract.UnpackLog(event, "WatcherPermissionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
