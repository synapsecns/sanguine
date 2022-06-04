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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200def8aa5d33768aa2b4fe8c57ac802c46b81b2d2d5d899bd9b53a25833066d6a64736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209bf1c339f0f16230aa287e0e6f32009043ffaca3af43a557481b8f987f96baed64736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220481cd6dc15f78c461d1984bd9c92db046e580cb3c2dc01aa2f66e69f15b3dfaa64736f6c634300080d0033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destinationAndNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"committedRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[2]\",\"name\":\"newRoot\",\"type\":\"bytes32[2]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature2\",\"type\":\"bytes\"}],\"name\":\"DoubleUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"ImproperUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updaterManager\",\"type\":\"address\"}],\"name\":\"NewUpdaterManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"UpdaterSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_BODY_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committedRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_oldRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"_newRoot\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature2\",\"type\":\"bytes\"}],\"name\":\"doubleUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_oldRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"improperUpdate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUpdaterManager\",\"name\":\"_updaterManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_item\",\"type\":\"bytes32\"}],\"name\":\"queueContains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueEnd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updaterManager\",\"type\":\"address\"}],\"name\":\"setUpdaterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumSynapseBase.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_committedRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_new\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_committedRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updaterManager\",\"outputs\":[{\"internalType\":\"contractIUpdaterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"522ae002": "MAX_MESSAGE_BODY_BYTES()",
		"ffa1ad74": "VERSION()",
		"67a6771d": "committedRoot()",
		"06661abd": "count()",
		"fa31de01": "dispatch(uint32,bytes32,bytes)",
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
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002d8b38038062002d8b833981016040819052620000349162000043565b63ffffffff1660805262000072565b6000602082840312156200005657600080fd5b815163ffffffff811681146200006b57600080fd5b9392505050565b608051612ce8620000a3600039600081816102460152818161063601528181610a410152610e3c0152612ce86000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c80639df6c8e1116100ee578063df034cd011610097578063f6d1610211610071578063f6d16102146103d8578063fa31de0114610415578063fd54b22814610428578063ffa1ad741461043257600080fd5b8063df034cd01461039d578063ebf0c717146103bd578063f2fde38b146103c557600080fd5b8063b95a2001116100c8578063b95a200114610331578063c19d93fb14610358578063c4d66de81461038a57600080fd5b80639df6c8e1146102f5578063ab91c7b014610316578063b31c01fb1461031e57600080fd5b806367a6771d1161015b5780638da5cb5b116101355780638da5cb5b1461027d5780638e4e30e0146102bc5780639776120e146102cf5780639d54f419146102e257600080fd5b806367a6771d14610230578063715018a6146102395780638d3638f41461024157600080fd5b806336e104de1161018c57806336e104de1461020257806345630b1a1461021f578063522ae0021461022757600080fd5b806306661abd146101b357806319d9d21a146101ca5780632bef2892146101df575b600080fd5b6054545b6040519081526020015b60405180910390f35b6101dd6101d83660046124c4565b61044c565b005b6101f26101ed366004612553565b6105c8565b60405190151581526020016101c1565b61020a6105db565b604080519283526020830191909152016101c1565b6101b761062f565b6101b761080081565b6101b760eb5481565b6101dd61065f565b6102687f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101c1565b60b85473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c1565b6101f26102ca366004612646565b6106c8565b6101dd6102dd3660046126b8565b610855565b6101dd6102f03660046126b8565b6108c8565b61011c546102979073ffffffffffffffffffffffffffffffffffffffff1681565b6101b761097b565b6101dd61032c366004612646565b610987565b61026861033f3660046126e9565b61011b6020526000908152604090205463ffffffff1681565b60ea5461037d9074010000000000000000000000000000000000000000900460ff1681565b6040516101c19190612733565b6101dd6103983660046126b8565b610aa3565b60ea546102979073ffffffffffffffffffffffffffffffffffffffff1681565b6101b7610bf8565b6101dd6103d33660046126b8565b610c04565b60015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff166000908152600260205260409020546101b7565b6101dd610423366004612774565b610cfd565b6054546101b79081565b61043a600081565b60405160ff90911681526020016101c1565b600260ea5474010000000000000000000000000000000000000000900460ff16600281111561047d5761047d612704565b036104cf5760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064015b60405180910390fd5b604080516020601f86018190048102820181019092528481526105119188918835918890889081908401838280828437600092019190915250610ef992505050565b8015610560575061056086866001602002013584848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ef992505050565b801561057157508435602086013514155b156105c05761057e610f86565b7f2c3f60bab4170347826231b75a920b5053941ddebc6eed6fd2c25721648b186f8686868686866040516105b79695949392919061282d565b60405180910390a15b505050505050565b60006105d5600183611094565b92915050565b6000806105e8600161110e565b1561062b57505060eb5460015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff166000908152600260205260409020545b9091565b600061065a7f000000000000000000000000000000000000000000000000000000000000000061114e565b905090565b60b85473ffffffffffffffffffffffffffffffffffffffff1633146106c65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104c6565b565b6000600260ea5474010000000000000000000000000000000000000000900460ff1660028111156106fb576106fb612704565b036107485760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064016104c6565b610753848484610ef9565b61079f5760405162461bcd60e51b815260206004820152600c60248201527f217570646174657220736967000000000000000000000000000000000000000060448201526064016104c6565b60eb5484146107f05760405162461bcd60e51b815260206004820152601460248201527f6e6f7420612063757272656e742075706461746500000000000000000000000060448201526064016104c6565b6107fb600184611094565b61084a57610807610f86565b7f6844fd5e21c932b5197b78ac11bf96e2eaa4e882dd0c88087060cf2065c04ab284848460405161083a939291906128e9565b60405180910390a150600161084e565b5060005b9392505050565b60b85473ffffffffffffffffffffffffffffffffffffffff1633146108bc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104c6565b6108c5816111c7565b50565b61011c5473ffffffffffffffffffffffffffffffffffffffff1633146109305760405162461bcd60e51b815260206004820152600f60248201527f21757064617465724d616e61676572000000000000000000000000000000000060448201526064016104c6565b610939816112a5565b5060ea80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000179055565b600061065a600161110e565b600260ea5474010000000000000000000000000000000000000000900460ff1660028111156109b8576109b8612704565b03610a055760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064016104c6565b610a108383836106c8565b610a9e575b6000610a216001611324565b9050828103610a305750610a36565b50610a15565b8160eb8190555081837f000000000000000000000000000000000000000000000000000000000000000063ffffffff167f608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b284604051610a959190612911565b60405180910390a45b505050565b6000610aaf600161144a565b90508015610ae457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610aec6115a1565b610af5826111c7565b61011c54604080517fdf034cd00000000000000000000000000000000000000000000000000000000081529051610b909273ffffffffffffffffffffffffffffffffffffffff169163df034cd09160048083019260209291908290030181865afa158015610b67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8b9190612924565b611628565b8015610bf457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b600061065a60346116ad565b60b85473ffffffffffffffffffffffffffffffffffffffff163314610c6b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104c6565b73ffffffffffffffffffffffffffffffffffffffff8116610cf45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104c6565b6108c5816116c0565b600260ea5474010000000000000000000000000000000000000000900460ff166002811115610d2e57610d2e612704565b03610d7b5760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064016104c6565b61080081511115610dce5760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e67000000000000000000000000000000000000000060448201526064016104c6565b63ffffffff808416600090815261011b602052604090205416610df2816001612970565b63ffffffff858116600090815261011b6020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001693909216929092179055610e657f00000000000000000000000000000000000000000000000000000000000000003384888888611737565b80516020820120909150610e7a60348261176f565b610e8d610e85610bf8565b600190611894565b5060545467ffffffff00000000602088901b1663ffffffff85161790610eb590600190612998565b827f9d4c83d2e57d7d381feb264b44a5015e7f9ef26340f4fc46b558a6dc16dd811a60eb5486604051610ee99291906129af565b60405180910390a4505050505050565b600080610f0461062f565b604080516020810192909252810186905260608101859052608001604051602081830303815290604052805190602001209050610f408161192a565b60ea5490915073ffffffffffffffffffffffffffffffffffffffff16610f668285611965565b73ffffffffffffffffffffffffffffffffffffffff161495945050505050565b60ea8054740200000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff90911617905561011c546040517f5b3c2cbf00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911690635b3c2cbf90602401600060405180830381600087803b15801561103357600080fd5b505af1158015611047573d6000803e3d6000fd5b505060ea5460405133935073ffffffffffffffffffffffffffffffffffffffff90911691507f98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce90600090a3565b815460009070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff165b83546fffffffffffffffffffffffffffffffff16811061084a5760008181526001850160205260409020548390036110fc5760019150506105d5565b80611106816129c8565b9150506110c0565b80546000906fffffffffffffffffffffffffffffffff7001000000000000000000000000000000008204811691166111468282611989565b949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e083901b1660208201527f53594e000000000000000000000000000000000000000000000000000000000060248201526000906027015b604051602081830303815290604052805190602001209050919050565b73ffffffffffffffffffffffffffffffffffffffff81163b61122b5760405162461bcd60e51b815260206004820152601860248201527f21636f6e747261637420757064617465724d616e61676572000000000000000060448201526064016104c6565b61011c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf9060200160405180910390a150565b60ea805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a9101610beb565b80546000906fffffffffffffffffffffffffffffffff70010000000000000000000000000000000082048116911661135c8282611989565b6000036113ab5760405162461bcd60e51b815260206004820152600560248201527f456d70747900000000000000000000000000000000000000000000000000000060448201526064016104c6565b6fffffffffffffffffffffffffffffffff81166000908152600185016020526040902054925082156113fc576fffffffffffffffffffffffffffffffff811660009081526001850160205260408120555b6114078160016129fd565b84547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff919091161790935550919050565b60008054610100900460ff16156114e7578160ff16600114801561146d5750303b155b6114df5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104c6565b506000919050565b60005460ff8084169116106115645760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104c6565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b600054610100900460ff1661161e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c6565b6106c660016119ba565b600054610100900460ff166116a55760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c6565b610930611a01565b60006105d5826116bb611a86565b611f47565b60b8805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606086868686868660405160200161175496959493929190612a28565b60405160208183030381529060405290509695505050505050565b600161177d60206002612bbb565b6117879190612998565b8260200154106117d95760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c0000000000000000000000000000000060448201526064016104c6565b60018260200160008282546117ee9190612bc7565b9091555050602082015460005b602081101561188b578160011660010361182a5782848260208110611822576118226127b5565b015550505050565b83816020811061183c5761183c6127b5565b015460408051602081019290925281018490526060016040516020818303038152906040528051906020012092506002826118779190612bdf565b91508061188381612c1a565b9150506117fb565b50610a9e612c52565b81546000906118ca9070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1660016129fd565b83546fffffffffffffffffffffffffffffffff808316700100000000000000000000000000000000029116178455905081156105d5576fffffffffffffffffffffffffffffffff8116600090815260019390930160205260409092205590565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c016111aa565b60008060006119748585612019565b9150915061198181612087565b509392505050565b6000816119978460016129fd565b6119a19190612c81565b6fffffffffffffffffffffffffffffffff169392505050565b80546fffffffffffffffffffffffffffffffff166000036108c55780547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016600117815550565b600054610100900460ff16611a7e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c6565b6106c6612273565b611a8e612463565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b602081101561201157600182821c166000868360208110611f7457611f746127b5565b0154905081600103611fb1576040805160208101839052908101869052606001604051602081830303815290604052805190602001209450611ffc565b84868460208110611fc457611fc46127b5565b6020020151604051602001611fe3929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b5050808061200990612c1a565b915050611f51565b505092915050565b600080825160410361204f5760208301516040840151606085015160001a612043878285856122f9565b94509450505050612080565b8251604003612078576020830151604084015161206d868383612411565b935093505050612080565b506000905060025b9250929050565b600081600481111561209b5761209b612704565b036120a35750565b60018160048111156120b7576120b7612704565b036121045760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104c6565b600281600481111561211857612118612704565b036121655760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104c6565b600381600481111561217957612179612704565b036121ec5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104c6565b600481600481111561220057612200612704565b036108c55760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104c6565b600054610100900460ff166122f05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104c6565b6106c6336116c0565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156123305750600090506003612408565b8460ff16601b1415801561234857508460ff16601c14155b156123595750600090506004612408565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156123ad573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661240157600060019250925050612408565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83168161244760ff86901c601b612bc7565b9050612455878288856122f9565b935093505050935093915050565b6040518061040001604052806020906020820280368337509192915050565b60008083601f84011261249457600080fd5b50813567ffffffffffffffff8111156124ac57600080fd5b60208301915083602082850101111561208057600080fd5b60008060008060008060a087890312156124dd57600080fd5b8635955060608701888111156124f257600080fd5b6020880195503567ffffffffffffffff8082111561250f57600080fd5b61251b8a838b01612482565b9096509450608089013591508082111561253457600080fd5b5061254189828a01612482565b979a9699509497509295939492505050565b60006020828403121561256557600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126125ac57600080fd5b813567ffffffffffffffff808211156125c7576125c761256c565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561260d5761260d61256c565b8160405283815286602085880101111561262657600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561265b57600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561268057600080fd5b61268c8682870161259b565b9150509250925092565b73ffffffffffffffffffffffffffffffffffffffff811681146108c557600080fd5b6000602082840312156126ca57600080fd5b813561084e81612696565b803563ffffffff8116811461159c57600080fd5b6000602082840312156126fb57600080fd5b61084e826126d5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b602081016003831061276e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60008060006060848603121561278957600080fd5b612792846126d5565b925060208401359150604084013567ffffffffffffffff81111561268057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b868152604086602083013760a06060820152600061284f60a0830186886127e4565b82810360808401526128628185876127e4565b9998505050505050505050565b60005b8381101561288a578181015183820152602001612872565b83811115612899576000848401525b50505050565b600081518084526128b781602086016020860161286f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b838152826020820152606060408201526000612908606083018461289f565b95945050505050565b60208152600061084e602083018461289f565b60006020828403121561293657600080fd5b815161084e81612696565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff80831681851680830382111561298f5761298f612941565b01949350505050565b6000828210156129aa576129aa612941565b500390565b828152604060208201526000611146604083018461289f565b6000816129d7576129d7612941565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60006fffffffffffffffffffffffffffffffff80831681851680830382111561298f5761298f612941565b60007fffffffff00000000000000000000000000000000000000000000000000000000808960e01b168352876004840152808760e01b166024840152808660e01b1660288401525083602c8301528251612a8981604c85016020870161286f565b91909101604c01979650505050505050565b600181815b80851115612af457817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612ada57612ada612941565b80851615612ae757918102915b93841c9390800290612aa0565b509250929050565b600082612b0b575060016105d5565b81612b18575060006105d5565b8160018114612b2e5760028114612b3857612b54565b60019150506105d5565b60ff841115612b4957612b49612941565b50506001821b6105d5565b5060208310610133831016604e8410600b8410161715612b77575081810a6105d5565b612b818383612a9b565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612bb357612bb3612941565b029392505050565b600061084e8383612afc565b60008219821115612bda57612bda612941565b500190565b600082612c15577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612c4b57612c4b612941565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006fffffffffffffffffffffffffffffffff83811690831681811015612caa57612caa612941565b03939250505056fea264697066735822122012823c98c7fcbdbc69ed9b09be364fead769a8a7b4f7443d82036f562815713764736f6c634300080d0033",
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

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns()
func (_Home *HomeTransactor) Dispatch(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "dispatch", _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns()
func (_Home *HomeSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Home.Contract.Dispatch(&_Home.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns()
func (_Home *HomeTransactorSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Home.Contract.Dispatch(&_Home.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122081a160ebfc37acf759f0ba154e86d3414f7ae06684359020d6d86f30484eecb564736f6c634300080d0033",
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
	Bin: "0x608060405234801561001057600080fd5b5061070e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806306661abd14610046578063ebf0c7171461005c578063fd54b22814610064575b600080fd5b6020545b60405190815260200160405180910390f35b61004a61006e565b60205461004a9081565b600061007a600061007f565b905090565b60006100928261008d610098565b610559565b92915050565b6100a061062b565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b602081101561062357600182821c1660008683602081106105865761058661064a565b01549050816001036105c357604080516020810183905290810186905260600160405160208183030381529060405280519060200120945061060e565b848684602081106105d6576105d661064a565b60200201516040516020016105f5929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b5050808061061b90610679565b915050610563565b505092915050565b6040518061040001604052806020906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036106d1577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea26469706673582212205f07116268606fa2463bae8f7be162b5f072e9f7f0de112526a6ed41be6d7dee64736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220aebab3096b47feabf86a7cf1fcc9212e2e550728711d634fdd10509930510de164736f6c634300080d0033",
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
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207d757c9ed1257bd8975da00299b6ae5a258da03ec0e3c7265f7a41cc0cda5e0c64736f6c634300080d0033",
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
	Bin: "0x608060405234801561001057600080fd5b506102f2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632bef289214610046578063ab91c7b01461006e578063f6d1610214610084575b600080fd5b6100596100543660046101da565b6100c1565b60405190151581526020015b60405180910390f35b6100766100d4565b604051908152602001610065565b60015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16600090815260026020526040902054610076565b60006100ce6001836100e5565b92915050565b60006100e06001610169565b905090565b815460009070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff165b83546fffffffffffffffffffffffffffffffff16811061015f57600081815260018501602052604090205483900361014d5760019150506100ce565b8061015781610222565b915050610111565b5060009392505050565b80546000906fffffffffffffffffffffffffffffffff7001000000000000000000000000000000008204811691166101a182826101a9565b949350505050565b6000816101b7846001610257565b6101c1919061028b565b6fffffffffffffffffffffffffffffffff169392505050565b6000602082840312156101ec57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600081610231576102316101f3565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60006fffffffffffffffffffffffffffffffff808316818516808303821115610282576102826101f3565b01949350505050565b60006fffffffffffffffffffffffffffffffff838116908316818110156102b4576102b46101f3565b03939250505056fea264697066735822122058027d1934143aeb00543cfd715ccfe29a171abe801d6bd61fac4775fbcb0eed64736f6c634300080d0033",
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

// ReplicaMetaData contains all meta data concerning the Replica contract.
var ReplicaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_processGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveGas\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"Process\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousConfirmAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newConfirmAt\",\"type\":\"uint256\"}],\"name\":\"SetConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"SetOptimisticTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PROCESS_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"acceptableRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committedRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"confirmAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_committedRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_optimisticSeconds\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"messages\",\"outputs\":[{\"internalType\":\"enumReplica.MessageStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"process\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_proof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"prove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_proof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"proveAndProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmAt\",\"type\":\"uint256\"}],\"name\":\"setConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_optimisticSeconds\",\"type\":\"uint256\"}],\"name\":\"setOptimisticTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumSynapseBase.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_oldRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d88beda2": "PROCESS_GAS()",
		"25e3beda": "RESERVE_GAS()",
		"ffa1ad74": "VERSION()",
		"a3f81d68": "acceptableRoot(bytes32)",
		"67a6771d": "committedRoot()",
		"71bfb7b8": "confirmAt(bytes32)",
		"45630b1a": "homeDomainHash()",
		"e7e7a7b7": "initialize(uint32,address,bytes32,uint256)",
		"8d3638f4": "localDomain()",
		"2bbd59ca": "messages(bytes32)",
		"39992668": "optimisticSeconds()",
		"8da5cb5b": "owner()",
		"928bc4b2": "process(bytes)",
		"371d3071": "prove(bytes32,bytes32[32],uint256)",
		"6188af0e": "proveAndProcess(bytes,bytes32[32],uint256)",
		"961681dc": "remoteDomain()",
		"715018a6": "renounceOwnership()",
		"088b5ed3": "setConfirmation(bytes32,uint256)",
		"885b5e2d": "setOptimisticTimeout(uint256)",
		"9d54f419": "setUpdater(address)",
		"c19d93fb": "state()",
		"f2fde38b": "transferOwnership(address)",
		"b31c01fb": "update(bytes32,bytes32,bytes)",
		"df034cd0": "updater()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b5060405162002b1538038062002b158339810160408190526200003491620000d9565b63ffffffff8316608052620cf850821015620000865760405162461bcd60e51b815260206004820152600c60248201526b2170726f636573732067617360a01b60448201526064015b60405180910390fd5b613a98811015620000c95760405162461bcd60e51b815260206004820152600c60248201526b21726573657276652067617360a01b60448201526064016200007d565b60a09190915260c052506200011b565b600080600060608486031215620000ef57600080fd5b835163ffffffff811681146200010457600080fd5b602085015160409095015190969495509392505050565b60805160a05160c0516129ae62000167600039600081816101b7015261093f0152600081816103a8015281816109600152610a000152600081816102ac015261079b01526129ae6000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c80638da5cb5b116100e3578063c19d93fb1161008c578063e7e7a7b711610066578063e7e7a7b7146103ea578063f2fde38b146103fd578063ffa1ad741461041057600080fd5b8063c19d93fb1461037e578063d88beda2146103a3578063df034cd0146103ca57600080fd5b80639d54f419116100bd5780639d54f41914610345578063a3f81d6814610358578063b31c01fb1461036b57600080fd5b80638da5cb5b146102e3578063928bc4b214610322578063961681dc1461033557600080fd5b80636188af0e1161014557806371bfb7b81161011f57806371bfb7b814610274578063885b5e2d146102945780638d3638f4146102a757600080fd5b80636188af0e1461025057806367a6771d14610263578063715018a61461026c57600080fd5b8063371d307111610176578063371d30711461021c578063399926681461023f57806345630b1a1461024857600080fd5b8063088b5ed31461019d57806325e3beda146101b25780632bbd59ca146101ec575b600080fd5b6101b06101ab3660046122f4565b61042a565b005b6101d97f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b61020f6101fa366004612316565b609a6020526000908152604090205460ff1681565b6040516101e39190612395565b61022f61022a3660046123ba565b6104ec565b60405190151581526020016101e3565b6101d960975481565b6101d96105f3565b6101b061025e3660046124cc565b61060d565b6101d960665481565b6101b061067a565b6101d9610282366004612316565b60996020526000908152604090205481565b6101b06102a2366004612316565b6106e3565b6102ce7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101e3565b60335473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e3565b61022f610330366004612515565b610785565b6096546102ce9063ffffffff1681565b6101b061035336600461256e565b610bae565b61022f610366366004612316565b610c21565b6101b0610379366004612589565b610c49565b60655461020f9074010000000000000000000000000000000000000000900460ff1681565b6101d97f000000000000000000000000000000000000000000000000000000000000000081565b6065546102fd9073ffffffffffffffffffffffffffffffffffffffff1681565b6101b06103f83660046125d9565b610d63565b6101b061040b36600461256e565b610ec5565b610418600081565b60405160ff90911681526020016101e3565b60335473ffffffffffffffffffffffffffffffffffffffff1633146104965760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b60008281526099602090815260409182902080549084905582518181529182018490529184917fe906d465281862f8f37a1bdd7d810e50347f06b0b493701cca6ccb8e6a7b2694910160405180910390a2505050565b6000806000858152609a602052604090205460ff1660028111156105125761051261232f565b1461055f5760405162461bcd60e51b815260206004820152601360248201527f214d6573736167655374617475732e4e6f6e6500000000000000000000000000604482015260640161048d565b6000610595858560208060200260405190810160405280929190826020800280828437600092019190915250879150610fbe9050565b90506105a081610c21565b156105e65750506000838152609a6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091556105ec565b60009150505b9392505050565b6096546000906106089063ffffffff1661106f565b905090565b61061f838051906020012083836104ec565b61066b5760405162461bcd60e51b815260206004820152600660248201527f2170726f76650000000000000000000000000000000000000000000000000000604482015260640161048d565b61067483610785565b50505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146106e15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161048d565b565b60335473ffffffffffffffffffffffffffffffffffffffff16331461074a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161048d565b60978190556040518181527fa8de34d9c26dbb20579bc8c8957813e3168af7c5092d7668390e34471dc4ffb49060200160405180910390a150565b60008061079283826110e8565b905063ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166107ca62ffffff19831661110e565b63ffffffff161461081d5760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e0000000000000000000000000000000000000000604482015260640161048d565b600061082e62ffffff198316611123565b905060016000828152609a602052604090205460ff1660028111156108555761085561232f565b146108a25760405162461bcd60e51b815260206004820152600760248201527f2170726f76656e00000000000000000000000000000000000000000000000000604482015260640161048d565b60985460ff166001146108f75760405162461bcd60e51b815260206004820152600a60248201527f217265656e7472616e7400000000000000000000000000000000000000000000604482015260640161048d565b609880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009081169091556000828152609a60205260409020805490911660021790556109847f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000612655565b5a10156109d55760405162461bcd60e51b815260040161048d9060208082526004908201527f2167617300000000000000000000000000000000000000000000000000000000604082015260600190565b60006109e662ffffff198416611180565b6040805161010080825261012082019092529192506000917f0000000000000000000000000000000000000000000000000000000000000000908390836020820181803683370190505090506000610a4362ffffff198916611191565b610a5262ffffff198a166111a5565b610a6162ffffff198b166111ba565b610a7e610a7362ffffff198d166111cf565b62ffffff1916611202565b604051602401610a9194939291906126e3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fab2dc3f50000000000000000000000000000000000000000000000000000000017815281519192506000918291828a88f198503d945083851115610b26578394505b848252846000602084013e81604051610b3f9190612712565b604051908190038120908a15159089907fd42de95a9b26f1be134c8ecce389dc4fcfa18753d01661b7b361233569e8fe4890600090a45050609880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905550949695505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610c155760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161048d565b610c1e8161125c565b50565b600081815260996020526040812054808203610c405750600092915050565b42101592915050565b6066548314610c9a5760405162461bcd60e51b815260206004820152601260248201527f6e6f742063757272656e74207570646174650000000000000000000000000000604482015260640161048d565b610ca58383836112e2565b610cf15760405162461bcd60e51b815260206004820152600c60248201527f2175706461746572207369670000000000000000000000000000000000000000604482015260640161048d565b609754610cfe9042612655565b6000838152609960205260409081902091909155606683905560965490518391859163ffffffff909116907f608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b290610d5690869061272e565b60405180910390a4505050565b6000610d6f600161136f565b90508015610da457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610dad846114c6565b6098805460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009091168117909155609680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff88161790556066849055600084815260996020908152604091829020929092556097849055518381527fa8de34d9c26dbb20579bc8c8957813e3168af7c5092d7668390e34471dc4ffb4910160405180910390a18015610ebe57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610f2c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161048d565b73ffffffffffffffffffffffffffffffffffffffff8116610fb55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161048d565b610c1e81611596565b8260005b602081101561106757600183821c166000858360208110610fe557610fe5612741565b6020020151905081600103611025576040805160208101839052908101859052606001604051602081830303815290604052805190602001209350611052565b60408051602081018690529081018290526060016040516020818303038152906040528051906020012093505b5050808061105f90612770565b915050610fc2565b509392505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e083901b1660208201527f53594e000000000000000000000000000000000000000000000000000000000060248201526000906027015b604051602081830303815290604052805190602001209050919050565b81516000906020840161110364ffffffffff8516828461160d565b925050505b92915050565b600061110862ffffff19831660286004611654565b60008061113e8360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006111688460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b600061110861118e83611684565b90565b600061110862ffffff198316826004611654565b600061110862ffffff19831660246004611654565b600061110862ffffff19831660046020611695565b6000611108604c6111f281601886901c6bffffffffffffffffffffffff166127a8565b62ffffff1985169190600061185e565b606060008061121f8460181c6bffffffffffffffffffffffff1690565b60405193508392506bffffffffffffffffffffffff16905061124b84611246846020612655565b6118e3565b508181016020016040529052919050565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a910160405180910390a15050565b6000806112ed6105f3565b60408051602081019290925281018690526060810185905260800160405160208183030381529060405280519060200120905061132981611a88565b60655490915073ffffffffffffffffffffffffffffffffffffffff1661134f8285611ac3565b73ffffffffffffffffffffffffffffffffffffffff161495945050505050565b60008054610100900460ff161561140c578160ff1660011480156113925750303b155b6114045760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161048d565b506000919050565b60005460ff8084169116106114895760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161048d565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b600054610100900460ff166115435760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161048d565b61154b611adf565b6115548161125c565b50606580547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000179055565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008061161a8484611b64565b905060405181111561162a575060005b8060000361163f5762ffffff199150506105ec565b5050606092831b9190911790911b1760181b90565b60006116618260206127bf565b61166c9060086127e2565b60ff1661167a858585611695565b901c949350505050565b600061110862ffffff198316602c60205b60008160ff166000036116aa575060006105ec565b6116c28460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166116dd8460ff8516611b64565b11156117555761173c6116fe8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166117248660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16611b70565b60405162461bcd60e51b815260040161048d919061272e565b60208260ff1611156117cf5760405162461bcd60e51b815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e203332206279746573000000000000606482015260840161048d565b60006117dc8360086127e2565b905060006117f88660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b6000806118798660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905061189286611bde565b6118a6856118a08489611b64565b90611b64565b11156118b95762ffffff199150506118db565b6118c38186611b64565b90506118d78364ffffffffff16828661160d565b9150505b949350505050565b600062ffffff19808416036119605760405162461bcd60e51b815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f6572206465726566000000000000000000000000000000000000000000000000606482015260840161048d565b61196983611c2e565b6119db5760405162461bcd60e51b815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e746572206465726566000000000000000000000000000000000000000000606482015260840161048d565b60006119f58460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000611a1f8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000604051905084811115611a445760206060fd5b8285848460045afa50611a7e611a5a8760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b9695505050505050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c016110cb565b6000806000611ad28585611c6b565b9150915061106781611cd9565b600054610100900460ff16611b5c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161048d565b6106e1611ec5565b60006105ec8284612655565b60606000611b7d86611f4b565b9150506000611b8b86611f4b565b9150506000611b9986611f4b565b9150506000611ba786611f4b565b91505083838383604051602001611bc1949392919061280b565b604051602081830303815290604052945050505050949350505050565b6000611bf88260181c6bffffffffffffffffffffffff1690565b611c108360781c6bffffffffffffffffffffffff1690565b611c1a9190612948565b6bffffffffffffffffffffffff1692915050565b6000611c3a8260d81c90565b64ffffffffff1664ffffffffff03611c5457506000919050565b6000611c5f83611bde565b60405110199392505050565b6000808251604103611ca15760208301516040840151606085015160001a611c958782858561200b565b94509450505050611cd2565b8251604003611cca5760208301516040840151611cbf868383612123565b935093505050611cd2565b506000905060025b9250929050565b6000816004811115611ced57611ced61232f565b03611cf55750565b6001816004811115611d0957611d0961232f565b03611d565760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161048d565b6002816004811115611d6a57611d6a61232f565b03611db75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161048d565b6003816004811115611dcb57611dcb61232f565b03611e3e5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161048d565b6004816004811115611e5257611e5261232f565b03610c1e5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161048d565b600054610100900460ff16611f425760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161048d565b6106e133611596565b600080601f5b600f8160ff161115611fa9576000611f6a8260086127e2565b60ff1685901c9050611f7b81612175565b61ffff16841793508160ff16601014611f9657601084901b93505b50611fa26001826127bf565b9050611f51565b50600f5b60ff8160ff161015612005576000611fc68260086127e2565b60ff1685901c9050611fd781612175565b61ffff16831792508160ff16600014611ff257601083901b92505b50611ffe6001826127bf565b9050611fad565b50915091565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612042575060009050600361211a565b8460ff16601b1415801561205a57508460ff16601c14155b1561206b575060009050600461211a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156120bf573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166121135760006001925092505061211a565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83168161215960ff86901c601b612655565b90506121678782888561200b565b935093505050935093915050565b600061218760048360ff16901c6121a7565b60ff1661ffff919091161760081b61219e826121a7565b60ff1617919050565b600060f08083179060ff821690036121c25750603092915050565b8060ff1660f1036121d65750603192915050565b8060ff1660f2036121ea5750603292915050565b8060ff1660f3036121fe5750603392915050565b8060ff1660f4036122125750603492915050565b8060ff1660f5036122265750603592915050565b8060ff1660f60361223a5750603692915050565b8060ff1660f70361224e5750603792915050565b8060ff1660f8036122625750603892915050565b8060ff1660f9036122765750603992915050565b8060ff1660fa0361228a5750606192915050565b8060ff1660fb0361229e5750606292915050565b8060ff1660fc036122b25750606392915050565b8060ff1660fd036122c65750606492915050565b8060ff1660fe036122da5750606592915050565b8060ff1660ff036122ee5750606692915050565b50919050565b6000806040838503121561230757600080fd5b50508035926020909101359150565b60006020828403121561232857600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110610c1e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b602081016123a28361235e565b91905290565b80610400810183101561110857600080fd5b600080600061044084860312156123d057600080fd5b833592506123e185602086016123a8565b915061042084013590509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261243257600080fd5b813567ffffffffffffffff8082111561244d5761244d6123f2565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715612493576124936123f2565b816040528381528660208588010111156124ac57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600061044084860312156124e257600080fd5b833567ffffffffffffffff8111156124f957600080fd5b61250586828701612421565b9350506123e185602086016123a8565b60006020828403121561252757600080fd5b813567ffffffffffffffff81111561253e57600080fd5b6118db84828501612421565b803573ffffffffffffffffffffffffffffffffffffffff811681146114c157600080fd5b60006020828403121561258057600080fd5b6105ec8261254a565b60008060006060848603121561259e57600080fd5b8335925060208401359150604084013567ffffffffffffffff8111156125c357600080fd5b6125cf86828701612421565b9150509250925092565b600080600080608085870312156125ef57600080fd5b843563ffffffff8116811461260357600080fd5b93506126116020860161254a565b93969395505050506040820135916060013590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561266857612668612626565b500190565b60005b83811015612688578181015183820152602001612670565b838111156106745750506000910152565b600081518084526126b181602086016020860161266d565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600063ffffffff808716835280861660208401525083604083015260806060830152611a7e6080830184612699565b6000825161272481846020870161266d565b9190910192915050565b6020815260006105ec6020830184612699565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036127a1576127a1612626565b5060010190565b6000828210156127ba576127ba612626565b500390565b600060ff821660ff8416808210156127d9576127d9612626565b90039392505050565b600060ff821660ff84168160ff048111821515161561280357612803612626565b029392505050565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201611a7e565b60006bffffffffffffffffffffffff80831681851680830382111561296f5761296f612626565b0194935050505056fea26469706673582212202862e14b7597c5919017b220e9d82a860e41927c9bc40ff704860c45bb65a33164736f6c634300080d0033",
}

// ReplicaABI is the input ABI used to generate the binding from.
// Deprecated: Use ReplicaMetaData.ABI instead.
var ReplicaABI = ReplicaMetaData.ABI

// Deprecated: Use ReplicaMetaData.Sigs instead.
// ReplicaFuncSigs maps the 4-byte function signature to its string representation.
var ReplicaFuncSigs = ReplicaMetaData.Sigs

// ReplicaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReplicaMetaData.Bin instead.
var ReplicaBin = ReplicaMetaData.Bin

// DeployReplica deploys a new Ethereum contract, binding an instance of Replica to it.
func DeployReplica(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32, _processGas *big.Int, _reserveGas *big.Int) (common.Address, *types.Transaction, *Replica, error) {
	parsed, err := ReplicaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReplicaBin), backend, _localDomain, _processGas, _reserveGas)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Replica{ReplicaCaller: ReplicaCaller{contract: contract}, ReplicaTransactor: ReplicaTransactor{contract: contract}, ReplicaFilterer: ReplicaFilterer{contract: contract}}, nil
}

// Replica is an auto generated Go binding around an Ethereum contract.
type Replica struct {
	ReplicaCaller     // Read-only binding to the contract
	ReplicaTransactor // Write-only binding to the contract
	ReplicaFilterer   // Log filterer for contract events
}

// ReplicaCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReplicaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReplicaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReplicaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReplicaSession struct {
	Contract     *Replica          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReplicaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReplicaCallerSession struct {
	Contract *ReplicaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ReplicaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReplicaTransactorSession struct {
	Contract     *ReplicaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReplicaRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReplicaRaw struct {
	Contract *Replica // Generic contract binding to access the raw methods on
}

// ReplicaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReplicaCallerRaw struct {
	Contract *ReplicaCaller // Generic read-only contract binding to access the raw methods on
}

// ReplicaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReplicaTransactorRaw struct {
	Contract *ReplicaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReplica creates a new instance of Replica, bound to a specific deployed contract.
func NewReplica(address common.Address, backend bind.ContractBackend) (*Replica, error) {
	contract, err := bindReplica(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Replica{ReplicaCaller: ReplicaCaller{contract: contract}, ReplicaTransactor: ReplicaTransactor{contract: contract}, ReplicaFilterer: ReplicaFilterer{contract: contract}}, nil
}

// NewReplicaCaller creates a new read-only instance of Replica, bound to a specific deployed contract.
func NewReplicaCaller(address common.Address, caller bind.ContractCaller) (*ReplicaCaller, error) {
	contract, err := bindReplica(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicaCaller{contract: contract}, nil
}

// NewReplicaTransactor creates a new write-only instance of Replica, bound to a specific deployed contract.
func NewReplicaTransactor(address common.Address, transactor bind.ContractTransactor) (*ReplicaTransactor, error) {
	contract, err := bindReplica(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicaTransactor{contract: contract}, nil
}

// NewReplicaFilterer creates a new log filterer instance of Replica, bound to a specific deployed contract.
func NewReplicaFilterer(address common.Address, filterer bind.ContractFilterer) (*ReplicaFilterer, error) {
	contract, err := bindReplica(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReplicaFilterer{contract: contract}, nil
}

// bindReplica binds a generic wrapper to an already deployed contract.
func bindReplica(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReplicaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Replica *ReplicaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Replica.Contract.ReplicaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Replica *ReplicaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Replica.Contract.ReplicaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Replica *ReplicaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Replica.Contract.ReplicaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Replica *ReplicaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Replica.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Replica *ReplicaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Replica.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Replica *ReplicaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Replica.Contract.contract.Transact(opts, method, params...)
}

// PROCESSGAS is a free data retrieval call binding the contract method 0xd88beda2.
//
// Solidity: function PROCESS_GAS() view returns(uint256)
func (_Replica *ReplicaCaller) PROCESSGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "PROCESS_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROCESSGAS is a free data retrieval call binding the contract method 0xd88beda2.
//
// Solidity: function PROCESS_GAS() view returns(uint256)
func (_Replica *ReplicaSession) PROCESSGAS() (*big.Int, error) {
	return _Replica.Contract.PROCESSGAS(&_Replica.CallOpts)
}

// PROCESSGAS is a free data retrieval call binding the contract method 0xd88beda2.
//
// Solidity: function PROCESS_GAS() view returns(uint256)
func (_Replica *ReplicaCallerSession) PROCESSGAS() (*big.Int, error) {
	return _Replica.Contract.PROCESSGAS(&_Replica.CallOpts)
}

// RESERVEGAS is a free data retrieval call binding the contract method 0x25e3beda.
//
// Solidity: function RESERVE_GAS() view returns(uint256)
func (_Replica *ReplicaCaller) RESERVEGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "RESERVE_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RESERVEGAS is a free data retrieval call binding the contract method 0x25e3beda.
//
// Solidity: function RESERVE_GAS() view returns(uint256)
func (_Replica *ReplicaSession) RESERVEGAS() (*big.Int, error) {
	return _Replica.Contract.RESERVEGAS(&_Replica.CallOpts)
}

// RESERVEGAS is a free data retrieval call binding the contract method 0x25e3beda.
//
// Solidity: function RESERVE_GAS() view returns(uint256)
func (_Replica *ReplicaCallerSession) RESERVEGAS() (*big.Int, error) {
	return _Replica.Contract.RESERVEGAS(&_Replica.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Replica *ReplicaCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Replica *ReplicaSession) VERSION() (uint8, error) {
	return _Replica.Contract.VERSION(&_Replica.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Replica *ReplicaCallerSession) VERSION() (uint8, error) {
	return _Replica.Contract.VERSION(&_Replica.CallOpts)
}

// AcceptableRoot is a free data retrieval call binding the contract method 0xa3f81d68.
//
// Solidity: function acceptableRoot(bytes32 _root) view returns(bool)
func (_Replica *ReplicaCaller) AcceptableRoot(opts *bind.CallOpts, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "acceptableRoot", _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptableRoot is a free data retrieval call binding the contract method 0xa3f81d68.
//
// Solidity: function acceptableRoot(bytes32 _root) view returns(bool)
func (_Replica *ReplicaSession) AcceptableRoot(_root [32]byte) (bool, error) {
	return _Replica.Contract.AcceptableRoot(&_Replica.CallOpts, _root)
}

// AcceptableRoot is a free data retrieval call binding the contract method 0xa3f81d68.
//
// Solidity: function acceptableRoot(bytes32 _root) view returns(bool)
func (_Replica *ReplicaCallerSession) AcceptableRoot(_root [32]byte) (bool, error) {
	return _Replica.Contract.AcceptableRoot(&_Replica.CallOpts, _root)
}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_Replica *ReplicaCaller) CommittedRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "committedRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_Replica *ReplicaSession) CommittedRoot() ([32]byte, error) {
	return _Replica.Contract.CommittedRoot(&_Replica.CallOpts)
}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_Replica *ReplicaCallerSession) CommittedRoot() ([32]byte, error) {
	return _Replica.Contract.CommittedRoot(&_Replica.CallOpts)
}

// ConfirmAt is a free data retrieval call binding the contract method 0x71bfb7b8.
//
// Solidity: function confirmAt(bytes32 ) view returns(uint256)
func (_Replica *ReplicaCaller) ConfirmAt(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "confirmAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmAt is a free data retrieval call binding the contract method 0x71bfb7b8.
//
// Solidity: function confirmAt(bytes32 ) view returns(uint256)
func (_Replica *ReplicaSession) ConfirmAt(arg0 [32]byte) (*big.Int, error) {
	return _Replica.Contract.ConfirmAt(&_Replica.CallOpts, arg0)
}

// ConfirmAt is a free data retrieval call binding the contract method 0x71bfb7b8.
//
// Solidity: function confirmAt(bytes32 ) view returns(uint256)
func (_Replica *ReplicaCallerSession) ConfirmAt(arg0 [32]byte) (*big.Int, error) {
	return _Replica.Contract.ConfirmAt(&_Replica.CallOpts, arg0)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_Replica *ReplicaCaller) HomeDomainHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "homeDomainHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_Replica *ReplicaSession) HomeDomainHash() ([32]byte, error) {
	return _Replica.Contract.HomeDomainHash(&_Replica.CallOpts)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_Replica *ReplicaCallerSession) HomeDomainHash() ([32]byte, error) {
	return _Replica.Contract.HomeDomainHash(&_Replica.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Replica *ReplicaCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Replica *ReplicaSession) LocalDomain() (uint32, error) {
	return _Replica.Contract.LocalDomain(&_Replica.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Replica *ReplicaCallerSession) LocalDomain() (uint32, error) {
	return _Replica.Contract.LocalDomain(&_Replica.CallOpts)
}

// Messages is a free data retrieval call binding the contract method 0x2bbd59ca.
//
// Solidity: function messages(bytes32 ) view returns(uint8)
func (_Replica *ReplicaCaller) Messages(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "messages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Messages is a free data retrieval call binding the contract method 0x2bbd59ca.
//
// Solidity: function messages(bytes32 ) view returns(uint8)
func (_Replica *ReplicaSession) Messages(arg0 [32]byte) (uint8, error) {
	return _Replica.Contract.Messages(&_Replica.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x2bbd59ca.
//
// Solidity: function messages(bytes32 ) view returns(uint8)
func (_Replica *ReplicaCallerSession) Messages(arg0 [32]byte) (uint8, error) {
	return _Replica.Contract.Messages(&_Replica.CallOpts, arg0)
}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x39992668.
//
// Solidity: function optimisticSeconds() view returns(uint256)
func (_Replica *ReplicaCaller) OptimisticSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "optimisticSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x39992668.
//
// Solidity: function optimisticSeconds() view returns(uint256)
func (_Replica *ReplicaSession) OptimisticSeconds() (*big.Int, error) {
	return _Replica.Contract.OptimisticSeconds(&_Replica.CallOpts)
}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x39992668.
//
// Solidity: function optimisticSeconds() view returns(uint256)
func (_Replica *ReplicaCallerSession) OptimisticSeconds() (*big.Int, error) {
	return _Replica.Contract.OptimisticSeconds(&_Replica.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Replica *ReplicaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Replica *ReplicaSession) Owner() (common.Address, error) {
	return _Replica.Contract.Owner(&_Replica.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Replica *ReplicaCallerSession) Owner() (common.Address, error) {
	return _Replica.Contract.Owner(&_Replica.CallOpts)
}

// RemoteDomain is a free data retrieval call binding the contract method 0x961681dc.
//
// Solidity: function remoteDomain() view returns(uint32)
func (_Replica *ReplicaCaller) RemoteDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "remoteDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RemoteDomain is a free data retrieval call binding the contract method 0x961681dc.
//
// Solidity: function remoteDomain() view returns(uint32)
func (_Replica *ReplicaSession) RemoteDomain() (uint32, error) {
	return _Replica.Contract.RemoteDomain(&_Replica.CallOpts)
}

// RemoteDomain is a free data retrieval call binding the contract method 0x961681dc.
//
// Solidity: function remoteDomain() view returns(uint32)
func (_Replica *ReplicaCallerSession) RemoteDomain() (uint32, error) {
	return _Replica.Contract.RemoteDomain(&_Replica.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Replica *ReplicaCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Replica *ReplicaSession) State() (uint8, error) {
	return _Replica.Contract.State(&_Replica.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Replica *ReplicaCallerSession) State() (uint8, error) {
	return _Replica.Contract.State(&_Replica.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_Replica *ReplicaCaller) Updater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Replica.contract.Call(opts, &out, "updater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_Replica *ReplicaSession) Updater() (common.Address, error) {
	return _Replica.Contract.Updater(&_Replica.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_Replica *ReplicaCallerSession) Updater() (common.Address, error) {
	return _Replica.Contract.Updater(&_Replica.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xe7e7a7b7.
//
// Solidity: function initialize(uint32 _remoteDomain, address _updater, bytes32 _committedRoot, uint256 _optimisticSeconds) returns()
func (_Replica *ReplicaTransactor) Initialize(opts *bind.TransactOpts, _remoteDomain uint32, _updater common.Address, _committedRoot [32]byte, _optimisticSeconds *big.Int) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "initialize", _remoteDomain, _updater, _committedRoot, _optimisticSeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xe7e7a7b7.
//
// Solidity: function initialize(uint32 _remoteDomain, address _updater, bytes32 _committedRoot, uint256 _optimisticSeconds) returns()
func (_Replica *ReplicaSession) Initialize(_remoteDomain uint32, _updater common.Address, _committedRoot [32]byte, _optimisticSeconds *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.Initialize(&_Replica.TransactOpts, _remoteDomain, _updater, _committedRoot, _optimisticSeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xe7e7a7b7.
//
// Solidity: function initialize(uint32 _remoteDomain, address _updater, bytes32 _committedRoot, uint256 _optimisticSeconds) returns()
func (_Replica *ReplicaTransactorSession) Initialize(_remoteDomain uint32, _updater common.Address, _committedRoot [32]byte, _optimisticSeconds *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.Initialize(&_Replica.TransactOpts, _remoteDomain, _updater, _committedRoot, _optimisticSeconds)
}

// Process is a paid mutator transaction binding the contract method 0x928bc4b2.
//
// Solidity: function process(bytes _message) returns(bool _success)
func (_Replica *ReplicaTransactor) Process(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "process", _message)
}

// Process is a paid mutator transaction binding the contract method 0x928bc4b2.
//
// Solidity: function process(bytes _message) returns(bool _success)
func (_Replica *ReplicaSession) Process(_message []byte) (*types.Transaction, error) {
	return _Replica.Contract.Process(&_Replica.TransactOpts, _message)
}

// Process is a paid mutator transaction binding the contract method 0x928bc4b2.
//
// Solidity: function process(bytes _message) returns(bool _success)
func (_Replica *ReplicaTransactorSession) Process(_message []byte) (*types.Transaction, error) {
	return _Replica.Contract.Process(&_Replica.TransactOpts, _message)
}

// Prove is a paid mutator transaction binding the contract method 0x371d3071.
//
// Solidity: function prove(bytes32 _leaf, bytes32[32] _proof, uint256 _index) returns(bool)
func (_Replica *ReplicaTransactor) Prove(opts *bind.TransactOpts, _leaf [32]byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "prove", _leaf, _proof, _index)
}

// Prove is a paid mutator transaction binding the contract method 0x371d3071.
//
// Solidity: function prove(bytes32 _leaf, bytes32[32] _proof, uint256 _index) returns(bool)
func (_Replica *ReplicaSession) Prove(_leaf [32]byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.Prove(&_Replica.TransactOpts, _leaf, _proof, _index)
}

// Prove is a paid mutator transaction binding the contract method 0x371d3071.
//
// Solidity: function prove(bytes32 _leaf, bytes32[32] _proof, uint256 _index) returns(bool)
func (_Replica *ReplicaTransactorSession) Prove(_leaf [32]byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.Prove(&_Replica.TransactOpts, _leaf, _proof, _index)
}

// ProveAndProcess is a paid mutator transaction binding the contract method 0x6188af0e.
//
// Solidity: function proveAndProcess(bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_Replica *ReplicaTransactor) ProveAndProcess(opts *bind.TransactOpts, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "proveAndProcess", _message, _proof, _index)
}

// ProveAndProcess is a paid mutator transaction binding the contract method 0x6188af0e.
//
// Solidity: function proveAndProcess(bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_Replica *ReplicaSession) ProveAndProcess(_message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.ProveAndProcess(&_Replica.TransactOpts, _message, _proof, _index)
}

// ProveAndProcess is a paid mutator transaction binding the contract method 0x6188af0e.
//
// Solidity: function proveAndProcess(bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_Replica *ReplicaTransactorSession) ProveAndProcess(_message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.ProveAndProcess(&_Replica.TransactOpts, _message, _proof, _index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Replica *ReplicaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Replica *ReplicaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Replica.Contract.RenounceOwnership(&_Replica.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Replica *ReplicaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Replica.Contract.RenounceOwnership(&_Replica.TransactOpts)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x088b5ed3.
//
// Solidity: function setConfirmation(bytes32 _root, uint256 _confirmAt) returns()
func (_Replica *ReplicaTransactor) SetConfirmation(opts *bind.TransactOpts, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "setConfirmation", _root, _confirmAt)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x088b5ed3.
//
// Solidity: function setConfirmation(bytes32 _root, uint256 _confirmAt) returns()
func (_Replica *ReplicaSession) SetConfirmation(_root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.SetConfirmation(&_Replica.TransactOpts, _root, _confirmAt)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x088b5ed3.
//
// Solidity: function setConfirmation(bytes32 _root, uint256 _confirmAt) returns()
func (_Replica *ReplicaTransactorSession) SetConfirmation(_root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.SetConfirmation(&_Replica.TransactOpts, _root, _confirmAt)
}

// SetOptimisticTimeout is a paid mutator transaction binding the contract method 0x885b5e2d.
//
// Solidity: function setOptimisticTimeout(uint256 _optimisticSeconds) returns()
func (_Replica *ReplicaTransactor) SetOptimisticTimeout(opts *bind.TransactOpts, _optimisticSeconds *big.Int) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "setOptimisticTimeout", _optimisticSeconds)
}

// SetOptimisticTimeout is a paid mutator transaction binding the contract method 0x885b5e2d.
//
// Solidity: function setOptimisticTimeout(uint256 _optimisticSeconds) returns()
func (_Replica *ReplicaSession) SetOptimisticTimeout(_optimisticSeconds *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.SetOptimisticTimeout(&_Replica.TransactOpts, _optimisticSeconds)
}

// SetOptimisticTimeout is a paid mutator transaction binding the contract method 0x885b5e2d.
//
// Solidity: function setOptimisticTimeout(uint256 _optimisticSeconds) returns()
func (_Replica *ReplicaTransactorSession) SetOptimisticTimeout(_optimisticSeconds *big.Int) (*types.Transaction, error) {
	return _Replica.Contract.SetOptimisticTimeout(&_Replica.TransactOpts, _optimisticSeconds)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_Replica *ReplicaTransactor) SetUpdater(opts *bind.TransactOpts, _updater common.Address) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "setUpdater", _updater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_Replica *ReplicaSession) SetUpdater(_updater common.Address) (*types.Transaction, error) {
	return _Replica.Contract.SetUpdater(&_Replica.TransactOpts, _updater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_Replica *ReplicaTransactorSession) SetUpdater(_updater common.Address) (*types.Transaction, error) {
	return _Replica.Contract.SetUpdater(&_Replica.TransactOpts, _updater)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Replica *ReplicaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Replica *ReplicaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Replica.Contract.TransferOwnership(&_Replica.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Replica *ReplicaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Replica.Contract.TransferOwnership(&_Replica.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0xb31c01fb.
//
// Solidity: function update(bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns()
func (_Replica *ReplicaTransactor) Update(opts *bind.TransactOpts, _oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Replica.contract.Transact(opts, "update", _oldRoot, _newRoot, _signature)
}

// Update is a paid mutator transaction binding the contract method 0xb31c01fb.
//
// Solidity: function update(bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns()
func (_Replica *ReplicaSession) Update(_oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Replica.Contract.Update(&_Replica.TransactOpts, _oldRoot, _newRoot, _signature)
}

// Update is a paid mutator transaction binding the contract method 0xb31c01fb.
//
// Solidity: function update(bytes32 _oldRoot, bytes32 _newRoot, bytes _signature) returns()
func (_Replica *ReplicaTransactorSession) Update(_oldRoot [32]byte, _newRoot [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Replica.Contract.Update(&_Replica.TransactOpts, _oldRoot, _newRoot, _signature)
}

// ReplicaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Replica contract.
type ReplicaInitializedIterator struct {
	Event *ReplicaInitialized // Event containing the contract specifics and raw log

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
func (it *ReplicaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaInitialized)
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
		it.Event = new(ReplicaInitialized)
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
func (it *ReplicaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaInitialized represents a Initialized event raised by the Replica contract.
type ReplicaInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Replica *ReplicaFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReplicaInitializedIterator, error) {

	logs, sub, err := _Replica.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReplicaInitializedIterator{contract: _Replica.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Replica *ReplicaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReplicaInitialized) (event.Subscription, error) {

	logs, sub, err := _Replica.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaInitialized)
				if err := _Replica.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Replica *ReplicaFilterer) ParseInitialized(log types.Log) (*ReplicaInitialized, error) {
	event := new(ReplicaInitialized)
	if err := _Replica.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaNewUpdaterIterator is returned from FilterNewUpdater and is used to iterate over the raw logs and unpacked data for NewUpdater events raised by the Replica contract.
type ReplicaNewUpdaterIterator struct {
	Event *ReplicaNewUpdater // Event containing the contract specifics and raw log

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
func (it *ReplicaNewUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaNewUpdater)
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
		it.Event = new(ReplicaNewUpdater)
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
func (it *ReplicaNewUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaNewUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaNewUpdater represents a NewUpdater event raised by the Replica contract.
type ReplicaNewUpdater struct {
	OldUpdater common.Address
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUpdater is a free log retrieval operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_Replica *ReplicaFilterer) FilterNewUpdater(opts *bind.FilterOpts) (*ReplicaNewUpdaterIterator, error) {

	logs, sub, err := _Replica.contract.FilterLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return &ReplicaNewUpdaterIterator{contract: _Replica.contract, event: "NewUpdater", logs: logs, sub: sub}, nil
}

// WatchNewUpdater is a free log subscription operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_Replica *ReplicaFilterer) WatchNewUpdater(opts *bind.WatchOpts, sink chan<- *ReplicaNewUpdater) (event.Subscription, error) {

	logs, sub, err := _Replica.contract.WatchLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaNewUpdater)
				if err := _Replica.contract.UnpackLog(event, "NewUpdater", log); err != nil {
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
func (_Replica *ReplicaFilterer) ParseNewUpdater(log types.Log) (*ReplicaNewUpdater, error) {
	event := new(ReplicaNewUpdater)
	if err := _Replica.contract.UnpackLog(event, "NewUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Replica contract.
type ReplicaOwnershipTransferredIterator struct {
	Event *ReplicaOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReplicaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaOwnershipTransferred)
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
		it.Event = new(ReplicaOwnershipTransferred)
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
func (it *ReplicaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaOwnershipTransferred represents a OwnershipTransferred event raised by the Replica contract.
type ReplicaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Replica *ReplicaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReplicaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Replica.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReplicaOwnershipTransferredIterator{contract: _Replica.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Replica *ReplicaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReplicaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Replica.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaOwnershipTransferred)
				if err := _Replica.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Replica *ReplicaFilterer) ParseOwnershipTransferred(log types.Log) (*ReplicaOwnershipTransferred, error) {
	event := new(ReplicaOwnershipTransferred)
	if err := _Replica.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaProcessIterator is returned from FilterProcess and is used to iterate over the raw logs and unpacked data for Process events raised by the Replica contract.
type ReplicaProcessIterator struct {
	Event *ReplicaProcess // Event containing the contract specifics and raw log

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
func (it *ReplicaProcessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaProcess)
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
		it.Event = new(ReplicaProcess)
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
func (it *ReplicaProcessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaProcessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaProcess represents a Process event raised by the Replica contract.
type ReplicaProcess struct {
	MessageHash [32]byte
	Success     bool
	ReturnData  common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProcess is a free log retrieval operation binding the contract event 0xd42de95a9b26f1be134c8ecce389dc4fcfa18753d01661b7b361233569e8fe48.
//
// Solidity: event Process(bytes32 indexed messageHash, bool indexed success, bytes indexed returnData)
func (_Replica *ReplicaFilterer) FilterProcess(opts *bind.FilterOpts, messageHash [][32]byte, success []bool, returnData [][]byte) (*ReplicaProcessIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var returnDataRule []interface{}
	for _, returnDataItem := range returnData {
		returnDataRule = append(returnDataRule, returnDataItem)
	}

	logs, sub, err := _Replica.contract.FilterLogs(opts, "Process", messageHashRule, successRule, returnDataRule)
	if err != nil {
		return nil, err
	}
	return &ReplicaProcessIterator{contract: _Replica.contract, event: "Process", logs: logs, sub: sub}, nil
}

// WatchProcess is a free log subscription operation binding the contract event 0xd42de95a9b26f1be134c8ecce389dc4fcfa18753d01661b7b361233569e8fe48.
//
// Solidity: event Process(bytes32 indexed messageHash, bool indexed success, bytes indexed returnData)
func (_Replica *ReplicaFilterer) WatchProcess(opts *bind.WatchOpts, sink chan<- *ReplicaProcess, messageHash [][32]byte, success []bool, returnData [][]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var returnDataRule []interface{}
	for _, returnDataItem := range returnData {
		returnDataRule = append(returnDataRule, returnDataItem)
	}

	logs, sub, err := _Replica.contract.WatchLogs(opts, "Process", messageHashRule, successRule, returnDataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaProcess)
				if err := _Replica.contract.UnpackLog(event, "Process", log); err != nil {
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

// ParseProcess is a log parse operation binding the contract event 0xd42de95a9b26f1be134c8ecce389dc4fcfa18753d01661b7b361233569e8fe48.
//
// Solidity: event Process(bytes32 indexed messageHash, bool indexed success, bytes indexed returnData)
func (_Replica *ReplicaFilterer) ParseProcess(log types.Log) (*ReplicaProcess, error) {
	event := new(ReplicaProcess)
	if err := _Replica.contract.UnpackLog(event, "Process", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaSetConfirmationIterator is returned from FilterSetConfirmation and is used to iterate over the raw logs and unpacked data for SetConfirmation events raised by the Replica contract.
type ReplicaSetConfirmationIterator struct {
	Event *ReplicaSetConfirmation // Event containing the contract specifics and raw log

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
func (it *ReplicaSetConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaSetConfirmation)
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
		it.Event = new(ReplicaSetConfirmation)
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
func (it *ReplicaSetConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaSetConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaSetConfirmation represents a SetConfirmation event raised by the Replica contract.
type ReplicaSetConfirmation struct {
	Root              [32]byte
	PreviousConfirmAt *big.Int
	NewConfirmAt      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetConfirmation is a free log retrieval operation binding the contract event 0xe906d465281862f8f37a1bdd7d810e50347f06b0b493701cca6ccb8e6a7b2694.
//
// Solidity: event SetConfirmation(bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_Replica *ReplicaFilterer) FilterSetConfirmation(opts *bind.FilterOpts, root [][32]byte) (*ReplicaSetConfirmationIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Replica.contract.FilterLogs(opts, "SetConfirmation", rootRule)
	if err != nil {
		return nil, err
	}
	return &ReplicaSetConfirmationIterator{contract: _Replica.contract, event: "SetConfirmation", logs: logs, sub: sub}, nil
}

// WatchSetConfirmation is a free log subscription operation binding the contract event 0xe906d465281862f8f37a1bdd7d810e50347f06b0b493701cca6ccb8e6a7b2694.
//
// Solidity: event SetConfirmation(bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_Replica *ReplicaFilterer) WatchSetConfirmation(opts *bind.WatchOpts, sink chan<- *ReplicaSetConfirmation, root [][32]byte) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Replica.contract.WatchLogs(opts, "SetConfirmation", rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaSetConfirmation)
				if err := _Replica.contract.UnpackLog(event, "SetConfirmation", log); err != nil {
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

// ParseSetConfirmation is a log parse operation binding the contract event 0xe906d465281862f8f37a1bdd7d810e50347f06b0b493701cca6ccb8e6a7b2694.
//
// Solidity: event SetConfirmation(bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_Replica *ReplicaFilterer) ParseSetConfirmation(log types.Log) (*ReplicaSetConfirmation, error) {
	event := new(ReplicaSetConfirmation)
	if err := _Replica.contract.UnpackLog(event, "SetConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaSetOptimisticTimeoutIterator is returned from FilterSetOptimisticTimeout and is used to iterate over the raw logs and unpacked data for SetOptimisticTimeout events raised by the Replica contract.
type ReplicaSetOptimisticTimeoutIterator struct {
	Event *ReplicaSetOptimisticTimeout // Event containing the contract specifics and raw log

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
func (it *ReplicaSetOptimisticTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaSetOptimisticTimeout)
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
		it.Event = new(ReplicaSetOptimisticTimeout)
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
func (it *ReplicaSetOptimisticTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaSetOptimisticTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaSetOptimisticTimeout represents a SetOptimisticTimeout event raised by the Replica contract.
type ReplicaSetOptimisticTimeout struct {
	Timeout *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetOptimisticTimeout is a free log retrieval operation binding the contract event 0xa8de34d9c26dbb20579bc8c8957813e3168af7c5092d7668390e34471dc4ffb4.
//
// Solidity: event SetOptimisticTimeout(uint256 timeout)
func (_Replica *ReplicaFilterer) FilterSetOptimisticTimeout(opts *bind.FilterOpts) (*ReplicaSetOptimisticTimeoutIterator, error) {

	logs, sub, err := _Replica.contract.FilterLogs(opts, "SetOptimisticTimeout")
	if err != nil {
		return nil, err
	}
	return &ReplicaSetOptimisticTimeoutIterator{contract: _Replica.contract, event: "SetOptimisticTimeout", logs: logs, sub: sub}, nil
}

// WatchSetOptimisticTimeout is a free log subscription operation binding the contract event 0xa8de34d9c26dbb20579bc8c8957813e3168af7c5092d7668390e34471dc4ffb4.
//
// Solidity: event SetOptimisticTimeout(uint256 timeout)
func (_Replica *ReplicaFilterer) WatchSetOptimisticTimeout(opts *bind.WatchOpts, sink chan<- *ReplicaSetOptimisticTimeout) (event.Subscription, error) {

	logs, sub, err := _Replica.contract.WatchLogs(opts, "SetOptimisticTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaSetOptimisticTimeout)
				if err := _Replica.contract.UnpackLog(event, "SetOptimisticTimeout", log); err != nil {
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

// ParseSetOptimisticTimeout is a log parse operation binding the contract event 0xa8de34d9c26dbb20579bc8c8957813e3168af7c5092d7668390e34471dc4ffb4.
//
// Solidity: event SetOptimisticTimeout(uint256 timeout)
func (_Replica *ReplicaFilterer) ParseSetOptimisticTimeout(log types.Log) (*ReplicaSetOptimisticTimeout, error) {
	event := new(ReplicaSetOptimisticTimeout)
	if err := _Replica.contract.UnpackLog(event, "SetOptimisticTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReplicaUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the Replica contract.
type ReplicaUpdateIterator struct {
	Event *ReplicaUpdate // Event containing the contract specifics and raw log

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
func (it *ReplicaUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicaUpdate)
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
		it.Event = new(ReplicaUpdate)
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
func (it *ReplicaUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicaUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicaUpdate represents a Update event raised by the Replica contract.
type ReplicaUpdate struct {
	HomeDomain uint32
	OldRoot    [32]byte
	NewRoot    [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_Replica *ReplicaFilterer) FilterUpdate(opts *bind.FilterOpts, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (*ReplicaUpdateIterator, error) {

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

	logs, sub, err := _Replica.contract.FilterLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return &ReplicaUpdateIterator{contract: _Replica.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_Replica *ReplicaFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *ReplicaUpdate, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Replica.contract.WatchLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicaUpdate)
				if err := _Replica.contract.UnpackLog(event, "Update", log); err != nil {
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
func (_Replica *ReplicaFilterer) ParseUpdate(log types.Log) (*ReplicaUpdate, error) {
	event := new(ReplicaUpdate)
	if err := _Replica.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203cee61e6275b14e669cedbc6e4d399223a6ef5ea03f786fb1141eb0ff3bb5f0a64736f6c634300080d0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208372415ba17311a7cda7ad057b4b282ed8ccb62d6f69a5df49bacafbc81ea9c564736f6c634300080d0033",
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

// SynapseBaseMetaData contains all meta data concerning the SynapseBase contract.
var SynapseBaseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"committedRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumSynapseBase.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"67a6771d": "committedRoot()",
		"45630b1a": "homeDomainHash()",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"c19d93fb": "state()",
		"f2fde38b": "transferOwnership(address)",
		"df034cd0": "updater()",
	},
}

// SynapseBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseBaseMetaData.ABI instead.
var SynapseBaseABI = SynapseBaseMetaData.ABI

// Deprecated: Use SynapseBaseMetaData.Sigs instead.
// SynapseBaseFuncSigs maps the 4-byte function signature to its string representation.
var SynapseBaseFuncSigs = SynapseBaseMetaData.Sigs

// SynapseBase is an auto generated Go binding around an Ethereum contract.
type SynapseBase struct {
	SynapseBaseCaller     // Read-only binding to the contract
	SynapseBaseTransactor // Write-only binding to the contract
	SynapseBaseFilterer   // Log filterer for contract events
}

// SynapseBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseBaseSession struct {
	Contract     *SynapseBase      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseBaseCallerSession struct {
	Contract *SynapseBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SynapseBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseBaseTransactorSession struct {
	Contract     *SynapseBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SynapseBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseBaseRaw struct {
	Contract *SynapseBase // Generic contract binding to access the raw methods on
}

// SynapseBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseBaseCallerRaw struct {
	Contract *SynapseBaseCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseBaseTransactorRaw struct {
	Contract *SynapseBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseBase creates a new instance of SynapseBase, bound to a specific deployed contract.
func NewSynapseBase(address common.Address, backend bind.ContractBackend) (*SynapseBase, error) {
	contract, err := bindSynapseBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseBase{SynapseBaseCaller: SynapseBaseCaller{contract: contract}, SynapseBaseTransactor: SynapseBaseTransactor{contract: contract}, SynapseBaseFilterer: SynapseBaseFilterer{contract: contract}}, nil
}

// NewSynapseBaseCaller creates a new read-only instance of SynapseBase, bound to a specific deployed contract.
func NewSynapseBaseCaller(address common.Address, caller bind.ContractCaller) (*SynapseBaseCaller, error) {
	contract, err := bindSynapseBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseBaseCaller{contract: contract}, nil
}

// NewSynapseBaseTransactor creates a new write-only instance of SynapseBase, bound to a specific deployed contract.
func NewSynapseBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseBaseTransactor, error) {
	contract, err := bindSynapseBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseBaseTransactor{contract: contract}, nil
}

// NewSynapseBaseFilterer creates a new log filterer instance of SynapseBase, bound to a specific deployed contract.
func NewSynapseBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseBaseFilterer, error) {
	contract, err := bindSynapseBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseBaseFilterer{contract: contract}, nil
}

// bindSynapseBase binds a generic wrapper to an already deployed contract.
func bindSynapseBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseBase *SynapseBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseBase.Contract.SynapseBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseBase *SynapseBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBase.Contract.SynapseBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseBase *SynapseBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseBase.Contract.SynapseBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseBase *SynapseBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseBase *SynapseBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseBase *SynapseBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseBase.Contract.contract.Transact(opts, method, params...)
}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_SynapseBase *SynapseBaseCaller) CommittedRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseBase.contract.Call(opts, &out, "committedRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_SynapseBase *SynapseBaseSession) CommittedRoot() ([32]byte, error) {
	return _SynapseBase.Contract.CommittedRoot(&_SynapseBase.CallOpts)
}

// CommittedRoot is a free data retrieval call binding the contract method 0x67a6771d.
//
// Solidity: function committedRoot() view returns(bytes32)
func (_SynapseBase *SynapseBaseCallerSession) CommittedRoot() ([32]byte, error) {
	return _SynapseBase.Contract.CommittedRoot(&_SynapseBase.CallOpts)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_SynapseBase *SynapseBaseCaller) HomeDomainHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynapseBase.contract.Call(opts, &out, "homeDomainHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_SynapseBase *SynapseBaseSession) HomeDomainHash() ([32]byte, error) {
	return _SynapseBase.Contract.HomeDomainHash(&_SynapseBase.CallOpts)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_SynapseBase *SynapseBaseCallerSession) HomeDomainHash() ([32]byte, error) {
	return _SynapseBase.Contract.HomeDomainHash(&_SynapseBase.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseBase *SynapseBaseCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SynapseBase.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseBase *SynapseBaseSession) LocalDomain() (uint32, error) {
	return _SynapseBase.Contract.LocalDomain(&_SynapseBase.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SynapseBase *SynapseBaseCallerSession) LocalDomain() (uint32, error) {
	return _SynapseBase.Contract.LocalDomain(&_SynapseBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseBase *SynapseBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseBase *SynapseBaseSession) Owner() (common.Address, error) {
	return _SynapseBase.Contract.Owner(&_SynapseBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseBase *SynapseBaseCallerSession) Owner() (common.Address, error) {
	return _SynapseBase.Contract.Owner(&_SynapseBase.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_SynapseBase *SynapseBaseCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SynapseBase.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_SynapseBase *SynapseBaseSession) State() (uint8, error) {
	return _SynapseBase.Contract.State(&_SynapseBase.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_SynapseBase *SynapseBaseCallerSession) State() (uint8, error) {
	return _SynapseBase.Contract.State(&_SynapseBase.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_SynapseBase *SynapseBaseCaller) Updater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseBase.contract.Call(opts, &out, "updater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_SynapseBase *SynapseBaseSession) Updater() (common.Address, error) {
	return _SynapseBase.Contract.Updater(&_SynapseBase.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_SynapseBase *SynapseBaseCallerSession) Updater() (common.Address, error) {
	return _SynapseBase.Contract.Updater(&_SynapseBase.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseBase *SynapseBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseBase *SynapseBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseBase.Contract.RenounceOwnership(&_SynapseBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseBase *SynapseBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseBase.Contract.RenounceOwnership(&_SynapseBase.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseBase *SynapseBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynapseBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseBase *SynapseBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseBase.Contract.TransferOwnership(&_SynapseBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseBase *SynapseBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseBase.Contract.TransferOwnership(&_SynapseBase.TransactOpts, newOwner)
}

// SynapseBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SynapseBase contract.
type SynapseBaseInitializedIterator struct {
	Event *SynapseBaseInitialized // Event containing the contract specifics and raw log

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
func (it *SynapseBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBaseInitialized)
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
		it.Event = new(SynapseBaseInitialized)
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
func (it *SynapseBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBaseInitialized represents a Initialized event raised by the SynapseBase contract.
type SynapseBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SynapseBase *SynapseBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*SynapseBaseInitializedIterator, error) {

	logs, sub, err := _SynapseBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SynapseBaseInitializedIterator{contract: _SynapseBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SynapseBase *SynapseBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SynapseBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _SynapseBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBaseInitialized)
				if err := _SynapseBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SynapseBase *SynapseBaseFilterer) ParseInitialized(log types.Log) (*SynapseBaseInitialized, error) {
	event := new(SynapseBaseInitialized)
	if err := _SynapseBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBaseNewUpdaterIterator is returned from FilterNewUpdater and is used to iterate over the raw logs and unpacked data for NewUpdater events raised by the SynapseBase contract.
type SynapseBaseNewUpdaterIterator struct {
	Event *SynapseBaseNewUpdater // Event containing the contract specifics and raw log

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
func (it *SynapseBaseNewUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBaseNewUpdater)
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
		it.Event = new(SynapseBaseNewUpdater)
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
func (it *SynapseBaseNewUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBaseNewUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBaseNewUpdater represents a NewUpdater event raised by the SynapseBase contract.
type SynapseBaseNewUpdater struct {
	OldUpdater common.Address
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUpdater is a free log retrieval operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_SynapseBase *SynapseBaseFilterer) FilterNewUpdater(opts *bind.FilterOpts) (*SynapseBaseNewUpdaterIterator, error) {

	logs, sub, err := _SynapseBase.contract.FilterLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return &SynapseBaseNewUpdaterIterator{contract: _SynapseBase.contract, event: "NewUpdater", logs: logs, sub: sub}, nil
}

// WatchNewUpdater is a free log subscription operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_SynapseBase *SynapseBaseFilterer) WatchNewUpdater(opts *bind.WatchOpts, sink chan<- *SynapseBaseNewUpdater) (event.Subscription, error) {

	logs, sub, err := _SynapseBase.contract.WatchLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBaseNewUpdater)
				if err := _SynapseBase.contract.UnpackLog(event, "NewUpdater", log); err != nil {
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
func (_SynapseBase *SynapseBaseFilterer) ParseNewUpdater(log types.Log) (*SynapseBaseNewUpdater, error) {
	event := new(SynapseBaseNewUpdater)
	if err := _SynapseBase.contract.UnpackLog(event, "NewUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynapseBase contract.
type SynapseBaseOwnershipTransferredIterator struct {
	Event *SynapseBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SynapseBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBaseOwnershipTransferred)
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
		it.Event = new(SynapseBaseOwnershipTransferred)
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
func (it *SynapseBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBaseOwnershipTransferred represents a OwnershipTransferred event raised by the SynapseBase contract.
type SynapseBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseBase *SynapseBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynapseBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBaseOwnershipTransferredIterator{contract: _SynapseBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseBase *SynapseBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynapseBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBaseOwnershipTransferred)
				if err := _SynapseBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SynapseBase *SynapseBaseFilterer) ParseOwnershipTransferred(log types.Log) (*SynapseBaseOwnershipTransferred, error) {
	event := new(SynapseBaseOwnershipTransferred)
	if err := _SynapseBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseBaseUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the SynapseBase contract.
type SynapseBaseUpdateIterator struct {
	Event *SynapseBaseUpdate // Event containing the contract specifics and raw log

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
func (it *SynapseBaseUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseBaseUpdate)
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
		it.Event = new(SynapseBaseUpdate)
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
func (it *SynapseBaseUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseBaseUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseBaseUpdate represents a Update event raised by the SynapseBase contract.
type SynapseBaseUpdate struct {
	HomeDomain uint32
	OldRoot    [32]byte
	NewRoot    [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_SynapseBase *SynapseBaseFilterer) FilterUpdate(opts *bind.FilterOpts, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (*SynapseBaseUpdateIterator, error) {

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

	logs, sub, err := _SynapseBase.contract.FilterLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return &SynapseBaseUpdateIterator{contract: _SynapseBase.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x608828ad904a0c9250c09004ba7226efb08f35a5c815bb3f76b5a8a271cd08b2.
//
// Solidity: event Update(uint32 indexed homeDomain, bytes32 indexed oldRoot, bytes32 indexed newRoot, bytes signature)
func (_SynapseBase *SynapseBaseFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *SynapseBaseUpdate, homeDomain []uint32, oldRoot [][32]byte, newRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SynapseBase.contract.WatchLogs(opts, "Update", homeDomainRule, oldRootRule, newRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseBaseUpdate)
				if err := _SynapseBase.contract.UnpackLog(event, "Update", log); err != nil {
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
func (_SynapseBase *SynapseBaseFilterer) ParseUpdate(log types.Log) (*SynapseBaseUpdate, error) {
	event := new(SynapseBaseUpdate)
	if err := _SynapseBase.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e32d007c1eae4f7987af7a268c15d41972510c5cf7e971fe431389a936aa3dcc64736f6c634300080d0033",
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
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea2646970667358221220b8476f68aaa4a604f39a3aaf48f9a8229390cbe53a75613d74b7cc73fc85eb2664736f6c634300080d0033",
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

// Version0MetaData contains all meta data concerning the Version0 contract.
var Version0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ffa1ad74": "VERSION()",
	},
	Bin: "0x6080604052348015600f57600080fd5b5060808061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063ffa1ad7414602d575b600080fd5b6034600081565b60405160ff909116815260200160405180910390f3fea26469706673582212208178e4dff1287acc1cb64f574366e58a4adab434191732063e1f2ec45b709eb364736f6c634300080d0033",
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
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6113df8061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638f5d90e01161008c578063b9cff16211610066578063b9cff16214610275578063e0e7a913146102ab578063f2fde38b146102be578063f31faefb146102d157600080fd5b80638f5d90e01461022f578063916c3470146102425780639fa92f9d1461025557600080fd5b80636ef0f37f116100c85780636ef0f37f146101cb578063715018a6146101e05780638d3638f4146101e85780638da5cb5b146101f057600080fd5b8063427ebef5146100ef5780635190bc53146101525780635f8b1dba14610190575b600080fd5b61013d6100fd366004611104565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260046020908152604080832063ffffffff8516845290915290205460ff1692915050565b60405190151581526020015b60405180910390f35b61013d61016036600461113d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205463ffffffff16151590565b6101b661019e36600461113d565b60026020526000908152604090205463ffffffff1681565b60405163ffffffff9091168152602001610149565b6101de6101d936600461113d565b6102e4565b005b6101de610397565b6101b6610400565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610149565b6101de61023d36600461113d565b610499565b6101de610250366004611161565b61050c565b60015461020a9073ffffffffffffffffffffffffffffffffffffffff1681565b61020a6102833660046111b1565b60036020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6101de6102b93660046111fd565b610617565b6101de6102cc36600461113d565b610836565b6101de6102df366004611104565b61092f565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103505760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103fe5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b565b600154604080517f8d3638f4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691638d3638f49160048083019260209291908290030181865afa158015610470573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049491906112e8565b905090565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105005760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b61050981610a61565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b73ffffffffffffffffffffffffffffffffffffffff8316600081815260046020908152604080832063ffffffff87168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915581519485529184019190915290917f517de16b526853f481451c5151e87484e1b251ec7d0302efa1019c2ece179c2c910160405180910390a2505050565b63ffffffff831660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16806106905760405162461bcd60e51b815260206004820152600f60248201527f217265706c6963612065786973747300000000000000000000000000000000006044820152606401610347565b8273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663df034cd06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107169190611305565b73ffffffffffffffffffffffffffffffffffffffff16146107795760405162461bcd60e51b815260206004820152601060248201527f2163757272656e742075706461746572000000000000000000000000000000006044820152606401610347565b600061079d8573ffffffffffffffffffffffffffffffffffffffff84168686610b20565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260046020908152604080832063ffffffff8a16845290915290205490915060ff166108265760405162461bcd60e51b815260206004820152600e60248201527f2176616c696420776174636865720000000000000000000000000000000000006044820152606401610347565b61082f82610a61565b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461089d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b73ffffffffffffffffffffffffffffffffffffffff81166109265760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610347565b61050981610c73565b60005473ffffffffffffffffffffffffffffffffffffffff1633146109965760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610347565b61099f82610a61565b73ffffffffffffffffffffffffffffffffffffffff8216600081815260026020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8716908117909155808452600383529281902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016851790555192835290917f8440df9bf8a8542634a9eb196da1735b786ed9aa2fc12b080ac34c5fa81a923491015b60405180910390a25050565b73ffffffffffffffffffffffffffffffffffffffff81166000818152600260208181526040808420805463ffffffff168086526003845282862080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690559486905292825282547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000169092559051928352909182917fce1533133fb359ace801d3176bbad25ace030d714aed35e38a6293c8a60b115b9101610a55565b6000808473ffffffffffffffffffffffffffffffffffffffff166345630b1a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b929190611322565b90506000818786604051602001610be19392919092835260e09190911b7fffffffff00000000000000000000000000000000000000000000000000000000166020830152602482015260440190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815282825280516020918201207f19457468657265756d205369676e6564204d6573736167653a0a33320000000082850152603c8085019190915282518085039091018152605c90930190915281519101209050610c688185610ce8565b979650505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806000610cf78585610d0c565b91509150610d0481610d7a565b509392505050565b6000808251604103610d425760208301516040840151606085015160001a610d3687828585610f66565b94509450505050610d73565b8251604003610d6b5760208301516040840151610d6086838361107e565b935093505050610d73565b506000905060025b9250929050565b6000816004811115610d8e57610d8e61133b565b03610d965750565b6001816004811115610daa57610daa61133b565b03610df75760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610347565b6002816004811115610e0b57610e0b61133b565b03610e585760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610347565b6003816004811115610e6c57610e6c61133b565b03610edf5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610347565b6004816004811115610ef357610ef361133b565b036105095760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610347565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610f9d5750600090506003611075565b8460ff16601b14158015610fb557508460ff16601c14155b15610fc65750600090506004611075565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561101a573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661106e57600060019250925050611075565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316816110b460ff86901c601b61136a565b90506110c287828885610f66565b935093505050935093915050565b73ffffffffffffffffffffffffffffffffffffffff8116811461050957600080fd5b63ffffffff8116811461050957600080fd5b6000806040838503121561111757600080fd5b8235611122816110d0565b91506020830135611132816110f2565b809150509250929050565b60006020828403121561114f57600080fd5b813561115a816110d0565b9392505050565b60008060006060848603121561117657600080fd5b8335611181816110d0565b92506020840135611191816110f2565b9150604084013580151581146111a657600080fd5b809150509250925092565b6000602082840312156111c357600080fd5b813561115a816110f2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561121257600080fd5b833561121d816110f2565b925060208401359150604084013567ffffffffffffffff8082111561124157600080fd5b818601915086601f83011261125557600080fd5b813581811115611267576112676111ce565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156112ad576112ad6111ce565b816040528281528960208487010111156112c657600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b6000602082840312156112fa57600080fd5b815161115a816110f2565b60006020828403121561131757600080fd5b815161115a816110d0565b60006020828403121561133457600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600082198211156113a4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50019056fea26469706673582212202553464f06f090f158c1b5a74b917bd2bf6e0e7f2ec30ac8610dac52c19cfe4a64736f6c634300080d0033",
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
