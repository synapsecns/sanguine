// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package origin

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

// AbstractGuardRegistryMetaData contains all meta data concerning the AbstractGuardRegistry contract.
var AbstractGuardRegistryMetaData = &bind.MetaData{
	ABI: "[]",
}

// AbstractGuardRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AbstractGuardRegistryMetaData.ABI instead.
var AbstractGuardRegistryABI = AbstractGuardRegistryMetaData.ABI

// AbstractGuardRegistry is an auto generated Go binding around an Ethereum contract.
type AbstractGuardRegistry struct {
	AbstractGuardRegistryCaller     // Read-only binding to the contract
	AbstractGuardRegistryTransactor // Write-only binding to the contract
	AbstractGuardRegistryFilterer   // Log filterer for contract events
}

// AbstractGuardRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractGuardRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractGuardRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractGuardRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractGuardRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractGuardRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractGuardRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractGuardRegistrySession struct {
	Contract     *AbstractGuardRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AbstractGuardRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractGuardRegistryCallerSession struct {
	Contract *AbstractGuardRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AbstractGuardRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractGuardRegistryTransactorSession struct {
	Contract     *AbstractGuardRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AbstractGuardRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractGuardRegistryRaw struct {
	Contract *AbstractGuardRegistry // Generic contract binding to access the raw methods on
}

// AbstractGuardRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractGuardRegistryCallerRaw struct {
	Contract *AbstractGuardRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AbstractGuardRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractGuardRegistryTransactorRaw struct {
	Contract *AbstractGuardRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractGuardRegistry creates a new instance of AbstractGuardRegistry, bound to a specific deployed contract.
func NewAbstractGuardRegistry(address common.Address, backend bind.ContractBackend) (*AbstractGuardRegistry, error) {
	contract, err := bindAbstractGuardRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractGuardRegistry{AbstractGuardRegistryCaller: AbstractGuardRegistryCaller{contract: contract}, AbstractGuardRegistryTransactor: AbstractGuardRegistryTransactor{contract: contract}, AbstractGuardRegistryFilterer: AbstractGuardRegistryFilterer{contract: contract}}, nil
}

// NewAbstractGuardRegistryCaller creates a new read-only instance of AbstractGuardRegistry, bound to a specific deployed contract.
func NewAbstractGuardRegistryCaller(address common.Address, caller bind.ContractCaller) (*AbstractGuardRegistryCaller, error) {
	contract, err := bindAbstractGuardRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractGuardRegistryCaller{contract: contract}, nil
}

// NewAbstractGuardRegistryTransactor creates a new write-only instance of AbstractGuardRegistry, bound to a specific deployed contract.
func NewAbstractGuardRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AbstractGuardRegistryTransactor, error) {
	contract, err := bindAbstractGuardRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractGuardRegistryTransactor{contract: contract}, nil
}

// NewAbstractGuardRegistryFilterer creates a new log filterer instance of AbstractGuardRegistry, bound to a specific deployed contract.
func NewAbstractGuardRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AbstractGuardRegistryFilterer, error) {
	contract, err := bindAbstractGuardRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractGuardRegistryFilterer{contract: contract}, nil
}

// bindAbstractGuardRegistry binds a generic wrapper to an already deployed contract.
func bindAbstractGuardRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractGuardRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractGuardRegistry *AbstractGuardRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractGuardRegistry.Contract.AbstractGuardRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractGuardRegistry *AbstractGuardRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractGuardRegistry.Contract.AbstractGuardRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractGuardRegistry *AbstractGuardRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractGuardRegistry.Contract.AbstractGuardRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractGuardRegistry *AbstractGuardRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractGuardRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractGuardRegistry *AbstractGuardRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractGuardRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractGuardRegistry *AbstractGuardRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractGuardRegistry.Contract.contract.Transact(opts, method, params...)
}

// AbstractNotaryRegistryMetaData contains all meta data concerning the AbstractNotaryRegistry contract.
var AbstractNotaryRegistryMetaData = &bind.MetaData{
	ABI: "[]",
}

// AbstractNotaryRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AbstractNotaryRegistryMetaData.ABI instead.
var AbstractNotaryRegistryABI = AbstractNotaryRegistryMetaData.ABI

// AbstractNotaryRegistry is an auto generated Go binding around an Ethereum contract.
type AbstractNotaryRegistry struct {
	AbstractNotaryRegistryCaller     // Read-only binding to the contract
	AbstractNotaryRegistryTransactor // Write-only binding to the contract
	AbstractNotaryRegistryFilterer   // Log filterer for contract events
}

// AbstractNotaryRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractNotaryRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractNotaryRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractNotaryRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractNotaryRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractNotaryRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractNotaryRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractNotaryRegistrySession struct {
	Contract     *AbstractNotaryRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AbstractNotaryRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractNotaryRegistryCallerSession struct {
	Contract *AbstractNotaryRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AbstractNotaryRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractNotaryRegistryTransactorSession struct {
	Contract     *AbstractNotaryRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AbstractNotaryRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractNotaryRegistryRaw struct {
	Contract *AbstractNotaryRegistry // Generic contract binding to access the raw methods on
}

// AbstractNotaryRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractNotaryRegistryCallerRaw struct {
	Contract *AbstractNotaryRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AbstractNotaryRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractNotaryRegistryTransactorRaw struct {
	Contract *AbstractNotaryRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractNotaryRegistry creates a new instance of AbstractNotaryRegistry, bound to a specific deployed contract.
func NewAbstractNotaryRegistry(address common.Address, backend bind.ContractBackend) (*AbstractNotaryRegistry, error) {
	contract, err := bindAbstractNotaryRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractNotaryRegistry{AbstractNotaryRegistryCaller: AbstractNotaryRegistryCaller{contract: contract}, AbstractNotaryRegistryTransactor: AbstractNotaryRegistryTransactor{contract: contract}, AbstractNotaryRegistryFilterer: AbstractNotaryRegistryFilterer{contract: contract}}, nil
}

// NewAbstractNotaryRegistryCaller creates a new read-only instance of AbstractNotaryRegistry, bound to a specific deployed contract.
func NewAbstractNotaryRegistryCaller(address common.Address, caller bind.ContractCaller) (*AbstractNotaryRegistryCaller, error) {
	contract, err := bindAbstractNotaryRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractNotaryRegistryCaller{contract: contract}, nil
}

// NewAbstractNotaryRegistryTransactor creates a new write-only instance of AbstractNotaryRegistry, bound to a specific deployed contract.
func NewAbstractNotaryRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AbstractNotaryRegistryTransactor, error) {
	contract, err := bindAbstractNotaryRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractNotaryRegistryTransactor{contract: contract}, nil
}

// NewAbstractNotaryRegistryFilterer creates a new log filterer instance of AbstractNotaryRegistry, bound to a specific deployed contract.
func NewAbstractNotaryRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AbstractNotaryRegistryFilterer, error) {
	contract, err := bindAbstractNotaryRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractNotaryRegistryFilterer{contract: contract}, nil
}

// bindAbstractNotaryRegistry binds a generic wrapper to an already deployed contract.
func bindAbstractNotaryRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractNotaryRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractNotaryRegistry *AbstractNotaryRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractNotaryRegistry.Contract.AbstractNotaryRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractNotaryRegistry *AbstractNotaryRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractNotaryRegistry.Contract.AbstractNotaryRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractNotaryRegistry *AbstractNotaryRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractNotaryRegistry.Contract.AbstractNotaryRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractNotaryRegistry *AbstractNotaryRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractNotaryRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractNotaryRegistry *AbstractNotaryRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractNotaryRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractNotaryRegistry *AbstractNotaryRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractNotaryRegistry.Contract.contract.Transact(opts, method, params...)
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220acecbac834160cfa53c933ee6cb6eebd6efca42086c5f00dc2b524ad5be696e564736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200d20d39259b2e084b4ca136ec55962f145728119357e1655042c1e1458d4410664736f6c634300080d0033",
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

// AttestationMetaData contains all meta data concerning the Attestation contract.
var AttestationMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a9b36290c28acd6b5facd2e217f28285f940360aab04a12cbcc44ab1a418df9664736f6c634300080d0033",
}

// AttestationABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationMetaData.ABI instead.
var AttestationABI = AttestationMetaData.ABI

// AttestationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationMetaData.Bin instead.
var AttestationBin = AttestationMetaData.Bin

// DeployAttestation deploys a new Ethereum contract, binding an instance of Attestation to it.
func DeployAttestation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Attestation, error) {
	parsed, err := AttestationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Attestation{AttestationCaller: AttestationCaller{contract: contract}, AttestationTransactor: AttestationTransactor{contract: contract}, AttestationFilterer: AttestationFilterer{contract: contract}}, nil
}

// Attestation is an auto generated Go binding around an Ethereum contract.
type Attestation struct {
	AttestationCaller     // Read-only binding to the contract
	AttestationTransactor // Write-only binding to the contract
	AttestationFilterer   // Log filterer for contract events
}

// AttestationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationSession struct {
	Contract     *Attestation      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttestationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationCallerSession struct {
	Contract *AttestationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AttestationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationTransactorSession struct {
	Contract     *AttestationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AttestationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationRaw struct {
	Contract *Attestation // Generic contract binding to access the raw methods on
}

// AttestationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationCallerRaw struct {
	Contract *AttestationCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationTransactorRaw struct {
	Contract *AttestationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestation creates a new instance of Attestation, bound to a specific deployed contract.
func NewAttestation(address common.Address, backend bind.ContractBackend) (*Attestation, error) {
	contract, err := bindAttestation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attestation{AttestationCaller: AttestationCaller{contract: contract}, AttestationTransactor: AttestationTransactor{contract: contract}, AttestationFilterer: AttestationFilterer{contract: contract}}, nil
}

// NewAttestationCaller creates a new read-only instance of Attestation, bound to a specific deployed contract.
func NewAttestationCaller(address common.Address, caller bind.ContractCaller) (*AttestationCaller, error) {
	contract, err := bindAttestation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationCaller{contract: contract}, nil
}

// NewAttestationTransactor creates a new write-only instance of Attestation, bound to a specific deployed contract.
func NewAttestationTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationTransactor, error) {
	contract, err := bindAttestation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationTransactor{contract: contract}, nil
}

// NewAttestationFilterer creates a new log filterer instance of Attestation, bound to a specific deployed contract.
func NewAttestationFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationFilterer, error) {
	contract, err := bindAttestation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationFilterer{contract: contract}, nil
}

// bindAttestation binds a generic wrapper to an already deployed contract.
func bindAttestation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attestation *AttestationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attestation.Contract.AttestationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attestation *AttestationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestation.Contract.AttestationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attestation *AttestationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attestation.Contract.AttestationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attestation *AttestationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attestation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attestation *AttestationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attestation *AttestationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attestation.Contract.contract.Transact(opts, method, params...)
}

// AuthMetaData contains all meta data concerning the Auth contract.
var AuthMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220926e13d966b5a834dbb1d9c6b2727c63b9919ca9d955f940f828ece654c13ebe64736f6c634300080d0033",
}

// AuthABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthMetaData.ABI instead.
var AuthABI = AuthMetaData.ABI

// AuthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuthMetaData.Bin instead.
var AuthBin = AuthMetaData.Bin

// DeployAuth deploys a new Ethereum contract, binding an instance of Auth to it.
func DeployAuth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Auth, error) {
	parsed, err := AuthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// Auth is an auto generated Go binding around an Ethereum contract.
type Auth struct {
	AuthCaller     // Read-only binding to the contract
	AuthTransactor // Write-only binding to the contract
	AuthFilterer   // Log filterer for contract events
}

// AuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthSession struct {
	Contract     *Auth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthCallerSession struct {
	Contract *AuthCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthTransactorSession struct {
	Contract     *AuthTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthRaw struct {
	Contract *Auth // Generic contract binding to access the raw methods on
}

// AuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthCallerRaw struct {
	Contract *AuthCaller // Generic read-only contract binding to access the raw methods on
}

// AuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthTransactorRaw struct {
	Contract *AuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuth creates a new instance of Auth, bound to a specific deployed contract.
func NewAuth(address common.Address, backend bind.ContractBackend) (*Auth, error) {
	contract, err := bindAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// NewAuthCaller creates a new read-only instance of Auth, bound to a specific deployed contract.
func NewAuthCaller(address common.Address, caller bind.ContractCaller) (*AuthCaller, error) {
	contract, err := bindAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthCaller{contract: contract}, nil
}

// NewAuthTransactor creates a new write-only instance of Auth, bound to a specific deployed contract.
func NewAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthTransactor, error) {
	contract, err := bindAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthTransactor{contract: contract}, nil
}

// NewAuthFilterer creates a new log filterer instance of Auth, bound to a specific deployed contract.
func NewAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthFilterer, error) {
	contract, err := bindAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthFilterer{contract: contract}, nil
}

// bindAuth binds a generic wrapper to an already deployed contract.
func bindAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.AuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transact(opts, method, params...)
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

// DomainNotaryRegistryMetaData contains all meta data concerning the DomainNotaryRegistry contract.
var DomainNotaryRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_trackedDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"DomainNotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"DomainNotaryRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allNotaries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNotary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notariesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9817e315": "allNotaries()",
		"c07dc7f5": "getNotary(uint256)",
		"8e62e9ef": "notariesAmount()",
	},
	Bin: "0x60a060405234801561001057600080fd5b506040516102e73803806102e783398101604081905261002f9161003d565b63ffffffff1660805261006a565b60006020828403121561004f57600080fd5b815163ffffffff8116811461006357600080fd5b9392505050565b608051610265610082600039600050506102656000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80638e62e9ef146100465780639817e31514610061578063c07dc7f514610076575b600080fd5b61004e6100ae565b6040519081526020015b60405180910390f35b6100696100bf565b604051610058919061018d565b6100896100843660046101e7565b6100cb565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610058565b60006100ba60006100dd565b905090565b60606100ba60006100e7565b60006100d781836100fb565b92915050565b60006100d7825490565b606060006100f483610107565b9392505050565b60006100f48383610163565b60608160000180548060200260200160405190810160405280929190818152602001828054801561015757602002820191906000526020600020905b815481526020019060010190808311610143575b50505050509050919050565b600082600001828154811061017a5761017a610200565b9060005260206000200154905092915050565b6020808252825182820181905260009190848201906040850190845b818110156101db57835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016101a9565b50909695505050505050565b6000602082840312156101f957600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea26469706673582212209a76d5fcaca149db72c7a7fb5ca44248a962520141a701b16c472cdb3248f87f64736f6c634300080d0033",
}

// DomainNotaryRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DomainNotaryRegistryMetaData.ABI instead.
var DomainNotaryRegistryABI = DomainNotaryRegistryMetaData.ABI

// Deprecated: Use DomainNotaryRegistryMetaData.Sigs instead.
// DomainNotaryRegistryFuncSigs maps the 4-byte function signature to its string representation.
var DomainNotaryRegistryFuncSigs = DomainNotaryRegistryMetaData.Sigs

// DomainNotaryRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DomainNotaryRegistryMetaData.Bin instead.
var DomainNotaryRegistryBin = DomainNotaryRegistryMetaData.Bin

// DeployDomainNotaryRegistry deploys a new Ethereum contract, binding an instance of DomainNotaryRegistry to it.
func DeployDomainNotaryRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _trackedDomain uint32) (common.Address, *types.Transaction, *DomainNotaryRegistry, error) {
	parsed, err := DomainNotaryRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DomainNotaryRegistryBin), backend, _trackedDomain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DomainNotaryRegistry{DomainNotaryRegistryCaller: DomainNotaryRegistryCaller{contract: contract}, DomainNotaryRegistryTransactor: DomainNotaryRegistryTransactor{contract: contract}, DomainNotaryRegistryFilterer: DomainNotaryRegistryFilterer{contract: contract}}, nil
}

// DomainNotaryRegistry is an auto generated Go binding around an Ethereum contract.
type DomainNotaryRegistry struct {
	DomainNotaryRegistryCaller     // Read-only binding to the contract
	DomainNotaryRegistryTransactor // Write-only binding to the contract
	DomainNotaryRegistryFilterer   // Log filterer for contract events
}

// DomainNotaryRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DomainNotaryRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainNotaryRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DomainNotaryRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainNotaryRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DomainNotaryRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainNotaryRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DomainNotaryRegistrySession struct {
	Contract     *DomainNotaryRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DomainNotaryRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DomainNotaryRegistryCallerSession struct {
	Contract *DomainNotaryRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DomainNotaryRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DomainNotaryRegistryTransactorSession struct {
	Contract     *DomainNotaryRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DomainNotaryRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DomainNotaryRegistryRaw struct {
	Contract *DomainNotaryRegistry // Generic contract binding to access the raw methods on
}

// DomainNotaryRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DomainNotaryRegistryCallerRaw struct {
	Contract *DomainNotaryRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DomainNotaryRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DomainNotaryRegistryTransactorRaw struct {
	Contract *DomainNotaryRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDomainNotaryRegistry creates a new instance of DomainNotaryRegistry, bound to a specific deployed contract.
func NewDomainNotaryRegistry(address common.Address, backend bind.ContractBackend) (*DomainNotaryRegistry, error) {
	contract, err := bindDomainNotaryRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DomainNotaryRegistry{DomainNotaryRegistryCaller: DomainNotaryRegistryCaller{contract: contract}, DomainNotaryRegistryTransactor: DomainNotaryRegistryTransactor{contract: contract}, DomainNotaryRegistryFilterer: DomainNotaryRegistryFilterer{contract: contract}}, nil
}

// NewDomainNotaryRegistryCaller creates a new read-only instance of DomainNotaryRegistry, bound to a specific deployed contract.
func NewDomainNotaryRegistryCaller(address common.Address, caller bind.ContractCaller) (*DomainNotaryRegistryCaller, error) {
	contract, err := bindDomainNotaryRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DomainNotaryRegistryCaller{contract: contract}, nil
}

// NewDomainNotaryRegistryTransactor creates a new write-only instance of DomainNotaryRegistry, bound to a specific deployed contract.
func NewDomainNotaryRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DomainNotaryRegistryTransactor, error) {
	contract, err := bindDomainNotaryRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DomainNotaryRegistryTransactor{contract: contract}, nil
}

// NewDomainNotaryRegistryFilterer creates a new log filterer instance of DomainNotaryRegistry, bound to a specific deployed contract.
func NewDomainNotaryRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DomainNotaryRegistryFilterer, error) {
	contract, err := bindDomainNotaryRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DomainNotaryRegistryFilterer{contract: contract}, nil
}

// bindDomainNotaryRegistry binds a generic wrapper to an already deployed contract.
func bindDomainNotaryRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DomainNotaryRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainNotaryRegistry *DomainNotaryRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomainNotaryRegistry.Contract.DomainNotaryRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainNotaryRegistry *DomainNotaryRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainNotaryRegistry.Contract.DomainNotaryRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainNotaryRegistry *DomainNotaryRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainNotaryRegistry.Contract.DomainNotaryRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainNotaryRegistry *DomainNotaryRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomainNotaryRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainNotaryRegistry *DomainNotaryRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainNotaryRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainNotaryRegistry *DomainNotaryRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainNotaryRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllNotaries is a free data retrieval call binding the contract method 0x9817e315.
//
// Solidity: function allNotaries() view returns(address[])
func (_DomainNotaryRegistry *DomainNotaryRegistryCaller) AllNotaries(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DomainNotaryRegistry.contract.Call(opts, &out, "allNotaries")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllNotaries is a free data retrieval call binding the contract method 0x9817e315.
//
// Solidity: function allNotaries() view returns(address[])
func (_DomainNotaryRegistry *DomainNotaryRegistrySession) AllNotaries() ([]common.Address, error) {
	return _DomainNotaryRegistry.Contract.AllNotaries(&_DomainNotaryRegistry.CallOpts)
}

// AllNotaries is a free data retrieval call binding the contract method 0x9817e315.
//
// Solidity: function allNotaries() view returns(address[])
func (_DomainNotaryRegistry *DomainNotaryRegistryCallerSession) AllNotaries() ([]common.Address, error) {
	return _DomainNotaryRegistry.Contract.AllNotaries(&_DomainNotaryRegistry.CallOpts)
}

// GetNotary is a free data retrieval call binding the contract method 0xc07dc7f5.
//
// Solidity: function getNotary(uint256 _index) view returns(address)
func (_DomainNotaryRegistry *DomainNotaryRegistryCaller) GetNotary(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DomainNotaryRegistry.contract.Call(opts, &out, "getNotary", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNotary is a free data retrieval call binding the contract method 0xc07dc7f5.
//
// Solidity: function getNotary(uint256 _index) view returns(address)
func (_DomainNotaryRegistry *DomainNotaryRegistrySession) GetNotary(_index *big.Int) (common.Address, error) {
	return _DomainNotaryRegistry.Contract.GetNotary(&_DomainNotaryRegistry.CallOpts, _index)
}

// GetNotary is a free data retrieval call binding the contract method 0xc07dc7f5.
//
// Solidity: function getNotary(uint256 _index) view returns(address)
func (_DomainNotaryRegistry *DomainNotaryRegistryCallerSession) GetNotary(_index *big.Int) (common.Address, error) {
	return _DomainNotaryRegistry.Contract.GetNotary(&_DomainNotaryRegistry.CallOpts, _index)
}

// NotariesAmount is a free data retrieval call binding the contract method 0x8e62e9ef.
//
// Solidity: function notariesAmount() view returns(uint256)
func (_DomainNotaryRegistry *DomainNotaryRegistryCaller) NotariesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DomainNotaryRegistry.contract.Call(opts, &out, "notariesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NotariesAmount is a free data retrieval call binding the contract method 0x8e62e9ef.
//
// Solidity: function notariesAmount() view returns(uint256)
func (_DomainNotaryRegistry *DomainNotaryRegistrySession) NotariesAmount() (*big.Int, error) {
	return _DomainNotaryRegistry.Contract.NotariesAmount(&_DomainNotaryRegistry.CallOpts)
}

// NotariesAmount is a free data retrieval call binding the contract method 0x8e62e9ef.
//
// Solidity: function notariesAmount() view returns(uint256)
func (_DomainNotaryRegistry *DomainNotaryRegistryCallerSession) NotariesAmount() (*big.Int, error) {
	return _DomainNotaryRegistry.Contract.NotariesAmount(&_DomainNotaryRegistry.CallOpts)
}

// DomainNotaryRegistryDomainNotaryAddedIterator is returned from FilterDomainNotaryAdded and is used to iterate over the raw logs and unpacked data for DomainNotaryAdded events raised by the DomainNotaryRegistry contract.
type DomainNotaryRegistryDomainNotaryAddedIterator struct {
	Event *DomainNotaryRegistryDomainNotaryAdded // Event containing the contract specifics and raw log

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
func (it *DomainNotaryRegistryDomainNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainNotaryRegistryDomainNotaryAdded)
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
		it.Event = new(DomainNotaryRegistryDomainNotaryAdded)
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
func (it *DomainNotaryRegistryDomainNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainNotaryRegistryDomainNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainNotaryRegistryDomainNotaryAdded represents a DomainNotaryAdded event raised by the DomainNotaryRegistry contract.
type DomainNotaryRegistryDomainNotaryAdded struct {
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainNotaryAdded is a free log retrieval operation binding the contract event 0x7ed5310d8818d06ea4a196771a39a73bf55c815addbf7a52ba87c9be409c3dd1.
//
// Solidity: event DomainNotaryAdded(address notary)
func (_DomainNotaryRegistry *DomainNotaryRegistryFilterer) FilterDomainNotaryAdded(opts *bind.FilterOpts) (*DomainNotaryRegistryDomainNotaryAddedIterator, error) {

	logs, sub, err := _DomainNotaryRegistry.contract.FilterLogs(opts, "DomainNotaryAdded")
	if err != nil {
		return nil, err
	}
	return &DomainNotaryRegistryDomainNotaryAddedIterator{contract: _DomainNotaryRegistry.contract, event: "DomainNotaryAdded", logs: logs, sub: sub}, nil
}

// WatchDomainNotaryAdded is a free log subscription operation binding the contract event 0x7ed5310d8818d06ea4a196771a39a73bf55c815addbf7a52ba87c9be409c3dd1.
//
// Solidity: event DomainNotaryAdded(address notary)
func (_DomainNotaryRegistry *DomainNotaryRegistryFilterer) WatchDomainNotaryAdded(opts *bind.WatchOpts, sink chan<- *DomainNotaryRegistryDomainNotaryAdded) (event.Subscription, error) {

	logs, sub, err := _DomainNotaryRegistry.contract.WatchLogs(opts, "DomainNotaryAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainNotaryRegistryDomainNotaryAdded)
				if err := _DomainNotaryRegistry.contract.UnpackLog(event, "DomainNotaryAdded", log); err != nil {
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

// ParseDomainNotaryAdded is a log parse operation binding the contract event 0x7ed5310d8818d06ea4a196771a39a73bf55c815addbf7a52ba87c9be409c3dd1.
//
// Solidity: event DomainNotaryAdded(address notary)
func (_DomainNotaryRegistry *DomainNotaryRegistryFilterer) ParseDomainNotaryAdded(log types.Log) (*DomainNotaryRegistryDomainNotaryAdded, error) {
	event := new(DomainNotaryRegistryDomainNotaryAdded)
	if err := _DomainNotaryRegistry.contract.UnpackLog(event, "DomainNotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainNotaryRegistryDomainNotaryRemovedIterator is returned from FilterDomainNotaryRemoved and is used to iterate over the raw logs and unpacked data for DomainNotaryRemoved events raised by the DomainNotaryRegistry contract.
type DomainNotaryRegistryDomainNotaryRemovedIterator struct {
	Event *DomainNotaryRegistryDomainNotaryRemoved // Event containing the contract specifics and raw log

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
func (it *DomainNotaryRegistryDomainNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainNotaryRegistryDomainNotaryRemoved)
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
		it.Event = new(DomainNotaryRegistryDomainNotaryRemoved)
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
func (it *DomainNotaryRegistryDomainNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainNotaryRegistryDomainNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainNotaryRegistryDomainNotaryRemoved represents a DomainNotaryRemoved event raised by the DomainNotaryRegistry contract.
type DomainNotaryRegistryDomainNotaryRemoved struct {
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainNotaryRemoved is a free log retrieval operation binding the contract event 0xe16811bec5badeb0bade36ad31aab1c20f2997b625833474449f893eeecd3bac.
//
// Solidity: event DomainNotaryRemoved(address notary)
func (_DomainNotaryRegistry *DomainNotaryRegistryFilterer) FilterDomainNotaryRemoved(opts *bind.FilterOpts) (*DomainNotaryRegistryDomainNotaryRemovedIterator, error) {

	logs, sub, err := _DomainNotaryRegistry.contract.FilterLogs(opts, "DomainNotaryRemoved")
	if err != nil {
		return nil, err
	}
	return &DomainNotaryRegistryDomainNotaryRemovedIterator{contract: _DomainNotaryRegistry.contract, event: "DomainNotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchDomainNotaryRemoved is a free log subscription operation binding the contract event 0xe16811bec5badeb0bade36ad31aab1c20f2997b625833474449f893eeecd3bac.
//
// Solidity: event DomainNotaryRemoved(address notary)
func (_DomainNotaryRegistry *DomainNotaryRegistryFilterer) WatchDomainNotaryRemoved(opts *bind.WatchOpts, sink chan<- *DomainNotaryRegistryDomainNotaryRemoved) (event.Subscription, error) {

	logs, sub, err := _DomainNotaryRegistry.contract.WatchLogs(opts, "DomainNotaryRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainNotaryRegistryDomainNotaryRemoved)
				if err := _DomainNotaryRegistry.contract.UnpackLog(event, "DomainNotaryRemoved", log); err != nil {
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

// ParseDomainNotaryRemoved is a log parse operation binding the contract event 0xe16811bec5badeb0bade36ad31aab1c20f2997b625833474449f893eeecd3bac.
//
// Solidity: event DomainNotaryRemoved(address notary)
func (_DomainNotaryRegistry *DomainNotaryRegistryFilterer) ParseDomainNotaryRemoved(log types.Log) (*DomainNotaryRegistryDomainNotaryRemoved, error) {
	event := new(DomainNotaryRegistryDomainNotaryRemoved)
	if err := _DomainNotaryRegistry.contract.UnpackLog(event, "DomainNotaryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f84f9562b5582875b3c867fe26a74c6e2aefa393edc2443853b15354cdd20c1c64736f6c634300080d0033",
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

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cda9d45a50a2ca55dee268087b69dd6bf133c2468739d2a61a9b75e2a1b0baa164736f6c634300080d0033",
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

// GuardRegistryMetaData contains all meta data concerning the GuardRegistry contract.
var GuardRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allGuards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9fe03fa2": "allGuards()",
		"629ddf69": "getGuard(uint256)",
		"246c2449": "guardsAmount()",
	},
	Bin: "0x608060405234801561001057600080fd5b50610265806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063246c244914610046578063629ddf69146100615780639fe03fa214610099575b600080fd5b61004e6100ae565b6040519081526020015b60405180910390f35b61007461006f36600461018d565b6100bf565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610058565b6100a16100d1565b60405161005891906101a6565b60006100ba60006100dd565b905090565b60006100cb81836100e7565b92915050565b60606100ba60006100fa565b60006100cb825490565b60006100f38383610107565b9392505050565b606060006100f383610131565b600082600001828154811061011e5761011e610200565b9060005260206000200154905092915050565b60608160000180548060200260200160405190810160405280929190818152602001828054801561018157602002820191906000526020600020905b81548152602001906001019080831161016d575b50505050509050919050565b60006020828403121561019f57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156101f457835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016101c2565b50909695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220398c56b4bdf21b78e8cb0b2f8ad6ec4a46d4988a6101ac046f4f0ac4307fee2464736f6c634300080d0033",
}

// GuardRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use GuardRegistryMetaData.ABI instead.
var GuardRegistryABI = GuardRegistryMetaData.ABI

// Deprecated: Use GuardRegistryMetaData.Sigs instead.
// GuardRegistryFuncSigs maps the 4-byte function signature to its string representation.
var GuardRegistryFuncSigs = GuardRegistryMetaData.Sigs

// GuardRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GuardRegistryMetaData.Bin instead.
var GuardRegistryBin = GuardRegistryMetaData.Bin

// DeployGuardRegistry deploys a new Ethereum contract, binding an instance of GuardRegistry to it.
func DeployGuardRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GuardRegistry, error) {
	parsed, err := GuardRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GuardRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GuardRegistry{GuardRegistryCaller: GuardRegistryCaller{contract: contract}, GuardRegistryTransactor: GuardRegistryTransactor{contract: contract}, GuardRegistryFilterer: GuardRegistryFilterer{contract: contract}}, nil
}

// GuardRegistry is an auto generated Go binding around an Ethereum contract.
type GuardRegistry struct {
	GuardRegistryCaller     // Read-only binding to the contract
	GuardRegistryTransactor // Write-only binding to the contract
	GuardRegistryFilterer   // Log filterer for contract events
}

// GuardRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuardRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuardRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuardRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuardRegistrySession struct {
	Contract     *GuardRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GuardRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuardRegistryCallerSession struct {
	Contract *GuardRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GuardRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuardRegistryTransactorSession struct {
	Contract     *GuardRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GuardRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuardRegistryRaw struct {
	Contract *GuardRegistry // Generic contract binding to access the raw methods on
}

// GuardRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuardRegistryCallerRaw struct {
	Contract *GuardRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// GuardRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuardRegistryTransactorRaw struct {
	Contract *GuardRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuardRegistry creates a new instance of GuardRegistry, bound to a specific deployed contract.
func NewGuardRegistry(address common.Address, backend bind.ContractBackend) (*GuardRegistry, error) {
	contract, err := bindGuardRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GuardRegistry{GuardRegistryCaller: GuardRegistryCaller{contract: contract}, GuardRegistryTransactor: GuardRegistryTransactor{contract: contract}, GuardRegistryFilterer: GuardRegistryFilterer{contract: contract}}, nil
}

// NewGuardRegistryCaller creates a new read-only instance of GuardRegistry, bound to a specific deployed contract.
func NewGuardRegistryCaller(address common.Address, caller bind.ContractCaller) (*GuardRegistryCaller, error) {
	contract, err := bindGuardRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuardRegistryCaller{contract: contract}, nil
}

// NewGuardRegistryTransactor creates a new write-only instance of GuardRegistry, bound to a specific deployed contract.
func NewGuardRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*GuardRegistryTransactor, error) {
	contract, err := bindGuardRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuardRegistryTransactor{contract: contract}, nil
}

// NewGuardRegistryFilterer creates a new log filterer instance of GuardRegistry, bound to a specific deployed contract.
func NewGuardRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*GuardRegistryFilterer, error) {
	contract, err := bindGuardRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuardRegistryFilterer{contract: contract}, nil
}

// bindGuardRegistry binds a generic wrapper to an already deployed contract.
func bindGuardRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GuardRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardRegistry *GuardRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardRegistry.Contract.GuardRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardRegistry *GuardRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardRegistry.Contract.GuardRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardRegistry *GuardRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardRegistry.Contract.GuardRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardRegistry *GuardRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardRegistry *GuardRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardRegistry *GuardRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_GuardRegistry *GuardRegistryCaller) AllGuards(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GuardRegistry.contract.Call(opts, &out, "allGuards")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_GuardRegistry *GuardRegistrySession) AllGuards() ([]common.Address, error) {
	return _GuardRegistry.Contract.AllGuards(&_GuardRegistry.CallOpts)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_GuardRegistry *GuardRegistryCallerSession) AllGuards() ([]common.Address, error) {
	return _GuardRegistry.Contract.AllGuards(&_GuardRegistry.CallOpts)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_GuardRegistry *GuardRegistryCaller) GetGuard(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GuardRegistry.contract.Call(opts, &out, "getGuard", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_GuardRegistry *GuardRegistrySession) GetGuard(_index *big.Int) (common.Address, error) {
	return _GuardRegistry.Contract.GetGuard(&_GuardRegistry.CallOpts, _index)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_GuardRegistry *GuardRegistryCallerSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _GuardRegistry.Contract.GetGuard(&_GuardRegistry.CallOpts, _index)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_GuardRegistry *GuardRegistryCaller) GuardsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GuardRegistry.contract.Call(opts, &out, "guardsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_GuardRegistry *GuardRegistrySession) GuardsAmount() (*big.Int, error) {
	return _GuardRegistry.Contract.GuardsAmount(&_GuardRegistry.CallOpts)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_GuardRegistry *GuardRegistryCallerSession) GuardsAmount() (*big.Int, error) {
	return _GuardRegistry.Contract.GuardsAmount(&_GuardRegistry.CallOpts)
}

// GuardRegistryGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the GuardRegistry contract.
type GuardRegistryGuardAddedIterator struct {
	Event *GuardRegistryGuardAdded // Event containing the contract specifics and raw log

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
func (it *GuardRegistryGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardRegistryGuardAdded)
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
		it.Event = new(GuardRegistryGuardAdded)
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
func (it *GuardRegistryGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardRegistryGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardRegistryGuardAdded represents a GuardAdded event raised by the GuardRegistry contract.
type GuardRegistryGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_GuardRegistry *GuardRegistryFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*GuardRegistryGuardAddedIterator, error) {

	logs, sub, err := _GuardRegistry.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &GuardRegistryGuardAddedIterator{contract: _GuardRegistry.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_GuardRegistry *GuardRegistryFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *GuardRegistryGuardAdded) (event.Subscription, error) {

	logs, sub, err := _GuardRegistry.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardRegistryGuardAdded)
				if err := _GuardRegistry.contract.UnpackLog(event, "GuardAdded", log); err != nil {
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

// ParseGuardAdded is a log parse operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_GuardRegistry *GuardRegistryFilterer) ParseGuardAdded(log types.Log) (*GuardRegistryGuardAdded, error) {
	event := new(GuardRegistryGuardAdded)
	if err := _GuardRegistry.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardRegistryGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the GuardRegistry contract.
type GuardRegistryGuardRemovedIterator struct {
	Event *GuardRegistryGuardRemoved // Event containing the contract specifics and raw log

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
func (it *GuardRegistryGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardRegistryGuardRemoved)
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
		it.Event = new(GuardRegistryGuardRemoved)
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
func (it *GuardRegistryGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardRegistryGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardRegistryGuardRemoved represents a GuardRemoved event raised by the GuardRegistry contract.
type GuardRegistryGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_GuardRegistry *GuardRegistryFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*GuardRegistryGuardRemovedIterator, error) {

	logs, sub, err := _GuardRegistry.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &GuardRegistryGuardRemovedIterator{contract: _GuardRegistry.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_GuardRegistry *GuardRegistryFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *GuardRegistryGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _GuardRegistry.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardRegistryGuardRemoved)
				if err := _GuardRegistry.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
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

// ParseGuardRemoved is a log parse operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_GuardRegistry *GuardRegistryFilterer) ParseGuardRemoved(log types.Log) (*GuardRegistryGuardRemoved, error) {
	event := new(GuardRegistryGuardRemoved)
	if err := _GuardRegistry.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeaderMetaData contains all meta data concerning the Header contract.
var HeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201ce2d01aef6823acd1fb9d9361200e9a548d4d522e0dbe2a81724ca29ac8d47f64736f6c634300080d0033",
}

// HeaderABI is the input ABI used to generate the binding from.
// Deprecated: Use HeaderMetaData.ABI instead.
var HeaderABI = HeaderMetaData.ABI

// HeaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HeaderMetaData.Bin instead.
var HeaderBin = HeaderMetaData.Bin

// DeployHeader deploys a new Ethereum contract, binding an instance of Header to it.
func DeployHeader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Header, error) {
	parsed, err := HeaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HeaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Header{HeaderCaller: HeaderCaller{contract: contract}, HeaderTransactor: HeaderTransactor{contract: contract}, HeaderFilterer: HeaderFilterer{contract: contract}}, nil
}

// Header is an auto generated Go binding around an Ethereum contract.
type Header struct {
	HeaderCaller     // Read-only binding to the contract
	HeaderTransactor // Write-only binding to the contract
	HeaderFilterer   // Log filterer for contract events
}

// HeaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeaderSession struct {
	Contract     *Header           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeaderCallerSession struct {
	Contract *HeaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HeaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeaderTransactorSession struct {
	Contract     *HeaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeaderRaw struct {
	Contract *Header // Generic contract binding to access the raw methods on
}

// HeaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeaderCallerRaw struct {
	Contract *HeaderCaller // Generic read-only contract binding to access the raw methods on
}

// HeaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeaderTransactorRaw struct {
	Contract *HeaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeader creates a new instance of Header, bound to a specific deployed contract.
func NewHeader(address common.Address, backend bind.ContractBackend) (*Header, error) {
	contract, err := bindHeader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Header{HeaderCaller: HeaderCaller{contract: contract}, HeaderTransactor: HeaderTransactor{contract: contract}, HeaderFilterer: HeaderFilterer{contract: contract}}, nil
}

// NewHeaderCaller creates a new read-only instance of Header, bound to a specific deployed contract.
func NewHeaderCaller(address common.Address, caller bind.ContractCaller) (*HeaderCaller, error) {
	contract, err := bindHeader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderCaller{contract: contract}, nil
}

// NewHeaderTransactor creates a new write-only instance of Header, bound to a specific deployed contract.
func NewHeaderTransactor(address common.Address, transactor bind.ContractTransactor) (*HeaderTransactor, error) {
	contract, err := bindHeader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderTransactor{contract: contract}, nil
}

// NewHeaderFilterer creates a new log filterer instance of Header, bound to a specific deployed contract.
func NewHeaderFilterer(address common.Address, filterer bind.ContractFilterer) (*HeaderFilterer, error) {
	contract, err := bindHeader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeaderFilterer{contract: contract}, nil
}

// bindHeader binds a generic wrapper to an already deployed contract.
func bindHeader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Header *HeaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Header.Contract.HeaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Header *HeaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Header.Contract.HeaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Header *HeaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Header.Contract.HeaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Header *HeaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Header.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Header *HeaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Header.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Header *HeaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Header.Contract.contract.Transact(opts, method, params...)
}

// INotaryManagerMetaData contains all meta data concerning the INotaryManager contract.
var INotaryManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"notary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_reporter\",\"type\":\"address\"}],\"name\":\"slashNotary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9d54c79d": "notary()",
		"bb99e8fa": "slashNotary(address)",
	},
}

// INotaryManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use INotaryManagerMetaData.ABI instead.
var INotaryManagerABI = INotaryManagerMetaData.ABI

// Deprecated: Use INotaryManagerMetaData.Sigs instead.
// INotaryManagerFuncSigs maps the 4-byte function signature to its string representation.
var INotaryManagerFuncSigs = INotaryManagerMetaData.Sigs

// INotaryManager is an auto generated Go binding around an Ethereum contract.
type INotaryManager struct {
	INotaryManagerCaller     // Read-only binding to the contract
	INotaryManagerTransactor // Write-only binding to the contract
	INotaryManagerFilterer   // Log filterer for contract events
}

// INotaryManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type INotaryManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INotaryManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INotaryManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INotaryManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INotaryManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INotaryManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INotaryManagerSession struct {
	Contract     *INotaryManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INotaryManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INotaryManagerCallerSession struct {
	Contract *INotaryManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// INotaryManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INotaryManagerTransactorSession struct {
	Contract     *INotaryManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// INotaryManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type INotaryManagerRaw struct {
	Contract *INotaryManager // Generic contract binding to access the raw methods on
}

// INotaryManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INotaryManagerCallerRaw struct {
	Contract *INotaryManagerCaller // Generic read-only contract binding to access the raw methods on
}

// INotaryManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INotaryManagerTransactorRaw struct {
	Contract *INotaryManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINotaryManager creates a new instance of INotaryManager, bound to a specific deployed contract.
func NewINotaryManager(address common.Address, backend bind.ContractBackend) (*INotaryManager, error) {
	contract, err := bindINotaryManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INotaryManager{INotaryManagerCaller: INotaryManagerCaller{contract: contract}, INotaryManagerTransactor: INotaryManagerTransactor{contract: contract}, INotaryManagerFilterer: INotaryManagerFilterer{contract: contract}}, nil
}

// NewINotaryManagerCaller creates a new read-only instance of INotaryManager, bound to a specific deployed contract.
func NewINotaryManagerCaller(address common.Address, caller bind.ContractCaller) (*INotaryManagerCaller, error) {
	contract, err := bindINotaryManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INotaryManagerCaller{contract: contract}, nil
}

// NewINotaryManagerTransactor creates a new write-only instance of INotaryManager, bound to a specific deployed contract.
func NewINotaryManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*INotaryManagerTransactor, error) {
	contract, err := bindINotaryManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INotaryManagerTransactor{contract: contract}, nil
}

// NewINotaryManagerFilterer creates a new log filterer instance of INotaryManager, bound to a specific deployed contract.
func NewINotaryManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*INotaryManagerFilterer, error) {
	contract, err := bindINotaryManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INotaryManagerFilterer{contract: contract}, nil
}

// bindINotaryManager binds a generic wrapper to an already deployed contract.
func bindINotaryManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INotaryManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INotaryManager *INotaryManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INotaryManager.Contract.INotaryManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INotaryManager *INotaryManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INotaryManager.Contract.INotaryManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INotaryManager *INotaryManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INotaryManager.Contract.INotaryManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INotaryManager *INotaryManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INotaryManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INotaryManager *INotaryManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INotaryManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INotaryManager *INotaryManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INotaryManager.Contract.contract.Transact(opts, method, params...)
}

// Notary is a free data retrieval call binding the contract method 0x9d54c79d.
//
// Solidity: function notary() view returns(address)
func (_INotaryManager *INotaryManagerCaller) Notary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INotaryManager.contract.Call(opts, &out, "notary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Notary is a free data retrieval call binding the contract method 0x9d54c79d.
//
// Solidity: function notary() view returns(address)
func (_INotaryManager *INotaryManagerSession) Notary() (common.Address, error) {
	return _INotaryManager.Contract.Notary(&_INotaryManager.CallOpts)
}

// Notary is a free data retrieval call binding the contract method 0x9d54c79d.
//
// Solidity: function notary() view returns(address)
func (_INotaryManager *INotaryManagerCallerSession) Notary() (common.Address, error) {
	return _INotaryManager.Contract.Notary(&_INotaryManager.CallOpts)
}

// SlashNotary is a paid mutator transaction binding the contract method 0xbb99e8fa.
//
// Solidity: function slashNotary(address _reporter) returns()
func (_INotaryManager *INotaryManagerTransactor) SlashNotary(opts *bind.TransactOpts, _reporter common.Address) (*types.Transaction, error) {
	return _INotaryManager.contract.Transact(opts, "slashNotary", _reporter)
}

// SlashNotary is a paid mutator transaction binding the contract method 0xbb99e8fa.
//
// Solidity: function slashNotary(address _reporter) returns()
func (_INotaryManager *INotaryManagerSession) SlashNotary(_reporter common.Address) (*types.Transaction, error) {
	return _INotaryManager.Contract.SlashNotary(&_INotaryManager.TransactOpts, _reporter)
}

// SlashNotary is a paid mutator transaction binding the contract method 0xbb99e8fa.
//
// Solidity: function slashNotary(address _reporter) returns()
func (_INotaryManager *INotaryManagerTransactorSession) SlashNotary(_reporter common.Address) (*types.Transaction, error) {
	return _INotaryManager.Contract.SlashNotary(&_INotaryManager.TransactOpts, _reporter)
}

// ISystemMessengerMetaData contains all meta data concerning the ISystemMessenger contract.
var ISystemMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"enumISystemMessenger.SystemContracts\",\"name\":\"_recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"sendSystemMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0d1e27a7": "sendSystemMessage(uint32,uint8,bytes)",
	},
}

// ISystemMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ISystemMessengerMetaData.ABI instead.
var ISystemMessengerABI = ISystemMessengerMetaData.ABI

// Deprecated: Use ISystemMessengerMetaData.Sigs instead.
// ISystemMessengerFuncSigs maps the 4-byte function signature to its string representation.
var ISystemMessengerFuncSigs = ISystemMessengerMetaData.Sigs

// ISystemMessenger is an auto generated Go binding around an Ethereum contract.
type ISystemMessenger struct {
	ISystemMessengerCaller     // Read-only binding to the contract
	ISystemMessengerTransactor // Write-only binding to the contract
	ISystemMessengerFilterer   // Log filterer for contract events
}

// ISystemMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemMessengerSession struct {
	Contract     *ISystemMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemMessengerCallerSession struct {
	Contract *ISystemMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ISystemMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemMessengerTransactorSession struct {
	Contract     *ISystemMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ISystemMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemMessengerRaw struct {
	Contract *ISystemMessenger // Generic contract binding to access the raw methods on
}

// ISystemMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemMessengerCallerRaw struct {
	Contract *ISystemMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemMessengerTransactorRaw struct {
	Contract *ISystemMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystemMessenger creates a new instance of ISystemMessenger, bound to a specific deployed contract.
func NewISystemMessenger(address common.Address, backend bind.ContractBackend) (*ISystemMessenger, error) {
	contract, err := bindISystemMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystemMessenger{ISystemMessengerCaller: ISystemMessengerCaller{contract: contract}, ISystemMessengerTransactor: ISystemMessengerTransactor{contract: contract}, ISystemMessengerFilterer: ISystemMessengerFilterer{contract: contract}}, nil
}

// NewISystemMessengerCaller creates a new read-only instance of ISystemMessenger, bound to a specific deployed contract.
func NewISystemMessengerCaller(address common.Address, caller bind.ContractCaller) (*ISystemMessengerCaller, error) {
	contract, err := bindISystemMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemMessengerCaller{contract: contract}, nil
}

// NewISystemMessengerTransactor creates a new write-only instance of ISystemMessenger, bound to a specific deployed contract.
func NewISystemMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemMessengerTransactor, error) {
	contract, err := bindISystemMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemMessengerTransactor{contract: contract}, nil
}

// NewISystemMessengerFilterer creates a new log filterer instance of ISystemMessenger, bound to a specific deployed contract.
func NewISystemMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemMessengerFilterer, error) {
	contract, err := bindISystemMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemMessengerFilterer{contract: contract}, nil
}

// bindISystemMessenger binds a generic wrapper to an already deployed contract.
func bindISystemMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISystemMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemMessenger *ISystemMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemMessenger.Contract.ISystemMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemMessenger *ISystemMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemMessenger.Contract.ISystemMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemMessenger *ISystemMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemMessenger.Contract.ISystemMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemMessenger *ISystemMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemMessenger *ISystemMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemMessenger *ISystemMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemMessenger.Contract.contract.Transact(opts, method, params...)
}

// SendSystemMessage is a paid mutator transaction binding the contract method 0x0d1e27a7.
//
// Solidity: function sendSystemMessage(uint32 _destination, uint8 _recipient, bytes _payload) returns()
func (_ISystemMessenger *ISystemMessengerTransactor) SendSystemMessage(opts *bind.TransactOpts, _destination uint32, _recipient uint8, _payload []byte) (*types.Transaction, error) {
	return _ISystemMessenger.contract.Transact(opts, "sendSystemMessage", _destination, _recipient, _payload)
}

// SendSystemMessage is a paid mutator transaction binding the contract method 0x0d1e27a7.
//
// Solidity: function sendSystemMessage(uint32 _destination, uint8 _recipient, bytes _payload) returns()
func (_ISystemMessenger *ISystemMessengerSession) SendSystemMessage(_destination uint32, _recipient uint8, _payload []byte) (*types.Transaction, error) {
	return _ISystemMessenger.Contract.SendSystemMessage(&_ISystemMessenger.TransactOpts, _destination, _recipient, _payload)
}

// SendSystemMessage is a paid mutator transaction binding the contract method 0x0d1e27a7.
//
// Solidity: function sendSystemMessage(uint32 _destination, uint8 _recipient, bytes _payload) returns()
func (_ISystemMessenger *ISystemMessengerTransactorSession) SendSystemMessage(_destination uint32, _recipient uint8, _payload []byte) (*types.Transaction, error) {
	return _ISystemMessenger.Contract.SendSystemMessage(&_ISystemMessenger.TransactOpts, _destination, _recipient, _payload)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122094bd546c873b9d2a1d6c76b995db26cab2fd74cec557c0a7467c71c9b096c45564736f6c634300080d0033",
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
	ABI: "[{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historicalRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"06661abd": "count()",
		"7ea97f40": "historicalRoots(uint256)",
		"ebf0c717": "root()",
		"fd54b228": "tree()",
	},
	Bin: "0x608060405234801561001057600080fd5b506106f8806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806306661abd146100515780637ea97f4014610067578063ebf0c7171461007a578063fd54b22814610082575b600080fd5b6020545b60405190815260200160405180910390f35b61005561007536600461067a565b61008c565b6100556100ad565b6020546100559081565b6021818154811061009c57600080fd5b600091825260209091200154905081565b60006100b960006100be565b905090565b60006100d1826100cc6100d7565b610598565b92915050565b6100df61065b565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b602081101561065357600182821c8116908190036105ff578582602081106105cc576105cc610693565b0154604080516020810192909252810185905260600160405160208183030381529060405280519060200120935061064a565b8385836020811061061257610612610693565b6020020151604051602001610631929190918252602082015260400190565b6040516020818303038152906040528051906020012093505b506001016105a2565b505092915050565b6040518061040001604052806020906020820280368337509192915050565b60006020828403121561068c57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220d4fbb46c28a38c9627fbbe76143e9b53be0eaf4215219eea958e4e2382b8efdc64736f6c634300080d0033",
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

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_MerkleTreeManager *MerkleTreeManagerCaller) HistoricalRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MerkleTreeManager.contract.Call(opts, &out, "historicalRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_MerkleTreeManager *MerkleTreeManagerSession) HistoricalRoots(arg0 *big.Int) ([32]byte, error) {
	return _MerkleTreeManager.Contract.HistoricalRoots(&_MerkleTreeManager.CallOpts, arg0)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_MerkleTreeManager *MerkleTreeManagerCallerSession) HistoricalRoots(arg0 *big.Int) ([32]byte, error) {
	return _MerkleTreeManager.Contract.HistoricalRoots(&_MerkleTreeManager.CallOpts, arg0)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e753058fb05496bace094b1bd190a43afb1cc157d9b1d45476d3b4deb72c78e664736f6c634300080d0033",
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

// OriginMetaData contains all meta data concerning the Origin contract.
var OriginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destinationAndNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tips\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"DomainNotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"DomainNotaryRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"ImproperAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notaryManager\",\"type\":\"address\"}],\"name\":\"NewNotaryManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"NotarySlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_BODY_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allGuards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allNotaries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNotary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historicalRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attestation\",\"type\":\"bytes\"}],\"name\":\"improperAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINotaryManager\",\"name\":\"_notaryManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notariesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notaryManager\",\"outputs\":[{\"internalType\":\"contractINotaryManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"setNotary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_notaryManager\",\"type\":\"address\"}],\"name\":\"setNotaryManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"_systemMessenger\",\"type\":\"address\"}],\"name\":\"setSystemMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumOrigin.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestAttestation\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemMessenger\",\"outputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"522ae002": "MAX_MESSAGE_BODY_BYTES()",
		"ffa1ad74": "VERSION()",
		"9fe03fa2": "allGuards()",
		"9817e315": "allNotaries()",
		"06661abd": "count()",
		"f7560e40": "dispatch(uint32,bytes32,uint32,bytes,bytes)",
		"629ddf69": "getGuard(uint256)",
		"c07dc7f5": "getNotary(uint256)",
		"246c2449": "guardsAmount()",
		"7ea97f40": "historicalRoots(uint256)",
		"0afe7f90": "improperAttestation(bytes)",
		"c4d66de8": "initialize(address)",
		"8d3638f4": "localDomain()",
		"affed0e0": "nonce()",
		"8e62e9ef": "notariesAmount()",
		"f85b597e": "notaryManager()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"ebf0c717": "root()",
		"a394a0e6": "setNotary(address)",
		"a340abc1": "setNotaryManager(address)",
		"b7bc563e": "setSystemMessenger(address)",
		"c19d93fb": "state()",
		"524787d0": "suggestAttestation()",
		"ccbdf9c9": "systemMessenger()",
		"f2fde38b": "transferOwnership(address)",
		"fd54b228": "tree()",
	},
	Bin: "0x60c06040523480156200001157600080fd5b506040516200387038038062003870833981016040819052620000349162000048565b63ffffffff16608081905260a05262000077565b6000602082840312156200005b57600080fd5b815163ffffffff811681146200007057600080fd5b9392505050565b60805160a0516137cc620000a460003960006118cb0152600081816102ef0152610ec901526137cc6000f3fe6080604052600436106101ac5760003560e01c8063a340abc1116100ec578063ccbdf9c91161008a578063f7560e4011610064578063f7560e4014610501578063f85b597e14610514578063fd54b2281461054a578063ffa1ad741461056157600080fd5b8063ccbdf9c91461049f578063ebf0c717146104cc578063f2fde38b146104e157600080fd5b8063b7bc563e116100c6578063b7bc563e146103fb578063c07dc7f51461041b578063c19d93fb1461043b578063c4d66de81461047f57600080fd5b8063a340abc11461039d578063a394a0e6146103bd578063affed0e0146103dd57600080fd5b8063715018a6116101595780638da5cb5b116101335780638da5cb5b146103265780638e62e9ef146103515780639817e315146103665780639fe03fa21461038857600080fd5b8063715018a6146102a65780637ea97f40146102bd5780638d3638f4146102dd57600080fd5b8063522ae0021161018a578063522ae0021461021a578063524787d014610230578063629ddf691461026157600080fd5b806306661abd146101b15780630afe7f90146101d5578063246c244914610205575b600080fd5b3480156101bd57600080fd5b506020545b6040519081526020015b60405180910390f35b3480156101e157600080fd5b506101f56101f036600461302b565b610588565b60405190151581526020016101cc565b34801561021157600080fd5b506101c26106df565b34801561022657600080fd5b506101c261080081565b34801561023c57600080fd5b506102456106f0565b6040805163ffffffff90931683526020830191909152016101cc565b34801561026d57600080fd5b5061028161027c366004613060565b610737565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cc565b3480156102b257600080fd5b506102bb61074a565b005b3480156102c957600080fd5b506101c26102d8366004613060565b6107b3565b3480156102e957600080fd5b506103117f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101cc565b34801561033257600080fd5b5060855473ffffffffffffffffffffffffffffffffffffffff16610281565b34801561035d57600080fd5b506101c26107d4565b34801561037257600080fd5b5061037b6107e0565b6040516101cc9190613079565b34801561039457600080fd5b5061037b6107ec565b3480156103a957600080fd5b506102bb6103b83660046130f5565b6107f8565b3480156103c957600080fd5b506102bb6103d83660046130f5565b61086b565b3480156103e957600080fd5b5061011e546103119063ffffffff1681565b34801561040757600080fd5b506102bb6104163660046130f5565b61092c565b34801561042757600080fd5b50610281610436366004613060565b6109da565b34801561044757600080fd5b5061011e54610472907801000000000000000000000000000000000000000000000000900460ff1681565b6040516101cc9190613141565b34801561048b57600080fd5b506102bb61049a3660046130f5565b6109e7565b3480156104ab57600080fd5b5060b7546102819073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104d857600080fd5b506101c2610bd5565b3480156104ed57600080fd5b506102bb6104fc3660046130f5565b610be1565b6102bb61050f366004613196565b610cda565b34801561052057600080fd5b5061011e5461028190640100000000900473ffffffffffffffffffffffffffffffffffffffff1681565b34801561055657600080fd5b506020546101c29081565b34801561056d57600080fd5b50610576600081565b60405160ff90911681526020016101cc565b6000600261011e547801000000000000000000000000000000000000000000000000900460ff1660028111156105c0576105c0613112565b036106125760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064015b60405180910390fd5b60008061061e84610fd4565b9150915060006106338262ffffff19166110e4565b9050600061064662ffffff1984166110f8565b60215490915063ffffffff8316101561068f5760218263ffffffff168154811061067257610672613225565b9060005260206000200154810361068f5750600095945050505050565b6106988461110d565b7f287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b84876040516106c99291906132ce565b60405180910390a160019450505050505b919050565b60006106eb60eb611225565b905090565b602154600090819080156107325761070960018261332c565b925060218363ffffffff168154811061072457610724613225565b906000526020600020015491505b509091565b600061074460eb8361122f565b92915050565b60855473ffffffffffffffffffffffffffffffffffffffff1633146107b15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610609565b565b602181815481106107c357600080fd5b600091825260209091200154905081565b60006106eb60b8611225565b60606106eb60b8611242565b60606106eb60eb611242565b60855473ffffffffffffffffffffffffffffffffffffffff16331461085f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610609565b6108688161124f565b50565b61011e54640100000000900473ffffffffffffffffffffffffffffffffffffffff1633146108db5760405162461bcd60e51b815260206004820152600e60248201527f216e6f746172794d616e616765720000000000000000000000000000000000006044820152606401610609565b6108e481611337565b505061011e80547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000179055565b60855473ffffffffffffffffffffffffffffffffffffffff1633146109935760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610609565b60b780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600061074460b88361122f565b60006109f3600161139a565b90508015610a2857605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610a306114ee565b610a398261124f565b610ad361011e60049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639d54c79d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aaa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ace9190613343565b611337565b5061011e80547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000179055602180546001810182556000919091527f27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d7577f3a6357012c1a3ae0a17d304c9920310382d968ebcc4b1771f41c6b304205b570909101558015610bd157605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b60006106eb6000611573565b60855473ffffffffffffffffffffffffffffffffffffffff163314610c485760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610609565b73ffffffffffffffffffffffffffffffffffffffff8116610cd15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610609565b61086881611586565b600261011e547801000000000000000000000000000000000000000000000000900460ff166002811115610d1057610d10613112565b03610d5d5760405162461bcd60e51b815260206004820152600c60248201527f6661696c656420737461746500000000000000000000000000000000000000006044820152606401610609565b61080081511115610db05760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e6700000000000000000000000000000000000000006044820152606401610609565b34610dc8610dbd846115fd565b62ffffff191661160a565b6bffffffffffffffffffffffff1614610e235760405162461bcd60e51b815260206004820152600560248201527f21746970730000000000000000000000000000000000000000000000000000006044820152606401610609565b61011e54610e389063ffffffff166001613360565b61011e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff929092169190911790556000610e798561166c565b61011e54604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060e090811b821660228401526026830186905293841b811660468301528a841b8116604a830152604e82018a90529288901b909216606e83015280518083036052018152607290920190529091506000610f438286866116cb565b80516020820120909150610f5681611745565b61011e5463ffffffff1660208a901b67ffffffff00000000161767ffffffffffffffff166001610f8560205490565b610f8f919061332c565b827f718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e7858986604051610fc1929190613388565b60405180910390a4505050505050505050565b600080610fe18382611775565b905060286bffffffffffffffffffffffff601883901c16116110455760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610609565b61107961105762ffffff198316611799565b61107461106962ffffff1985166117ae565b62ffffff19166117e1565b611834565b915061109361108d62ffffff1983166118b3565b836118c7565b6110df5760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061206e6f74617279000000000000000000006044820152606401610609565b915091565b600061074462ffffff19831660048061194f565b600061074462ffffff1983166008602061197f565b61011e805478020000000000000000000000000000000000000000000000007fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff90911617908190556040517fbb99e8fa00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff640100000000909204919091169063bb99e8fa90602401600060405180830381600087803b1580156111c757600080fd5b505af11580156111db573d6000803e3d6000fd5b505060405133925073ffffffffffffffffffffffffffffffffffffffff841691507f9ad95700c0e79cb6384afc527c4e16b94a9c6e39ea2ba3824ad6f856bdc4a40190600090a350565b6000610744825490565b600061123b8383611b3d565b9392505050565b6060600061123b83611b67565b73ffffffffffffffffffffffffffffffffffffffff81163b6112b35760405162461bcd60e51b815260206004820152601760248201527f21636f6e7472616374206e6f746172794d616e616765720000000000000000006044820152606401610609565b61011e80547fffffffffffffffff0000000000000000000000000000000000000000ffffffff1664010000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527fe3befd3a32a53f50ff7d1421555fbd40e5ead3a7ed75417db43a23faffe093169060200160405180910390a150565b600061134460b883611bc3565b905080156106da5760405173ffffffffffffffffffffffffffffffffffffffff831681527f7ed5310d8818d06ea4a196771a39a73bf55c815addbf7a52ba87c9be409c3dd19060200160405180910390a1919050565b605254600090610100900460ff1615611439578160ff1660011480156113bf5750303b155b6114315760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610609565b506000919050565b60525460ff8084169116106114b65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610609565b50605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b605254610100900460ff1661156b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610609565b6107b1611be5565b600061074482611581611c6b565b61212c565b6085805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610744826002611775565b60008161162060025b62ffffff198316906121ef565b5061162a836122f0565b6116338461231e565b61163c8561233f565b61164586612360565b61164f91906133ad565b61165991906133ad565b61166391906133ad565b91505b50919050565b60007fffffffffffffffffffffffff0000000000000000000000000000000000000000821461169c573392915050565b6116a4612381565b507fffffffffffffffffffffffff0000000000000000000000000000000000000000919050565b82516060906000906116df600460026133d4565b60ff166116ec91906133fd565b905060008451826116fd91906133fd565b9050600161170d600460026133d4565b60ff16838389898960405160200161172b979695949392919061341a565b604051602081830303815290604052925050509392505050565b6117506000826123e8565b602161175c6000611573565b8154600181018355600092835260209092209091015550565b81516000906020840161179064ffffffffff8516828461250b565b95945050505050565b600061074462ffffff19831682602881612552565b600061074460286117d181601886901c6bffffffffffffffffffffffff1661332c565b62ffffff19851691906000612552565b60606000806117fe8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061182384836020016125d6565b508181016020016040529052919050565b60008061184662ffffff198516612771565b905061189f816040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506118ab81846127ce565b949350505050565b600061074462ffffff19831682600461194f565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff168363ffffffff16146119445760405162461bcd60e51b815260206004820152600c60248201527f57726f6e6720646f6d61696e00000000000000000000000000000000000000006044820152606401610609565b61123b60b8836127f2565b600061195c8260206134b8565b6119679060086133d4565b60ff1661197585858561197f565b901c949350505050565b60008160ff166000036119945750600061123b565b6119ac8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166119c760ff8416856134db565b1115611a3f57611a266119e88560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16611a0e8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16612821565b60405162461bcd60e51b815260040161060991906134f3565b60208260ff161115611ab95760405162461bcd60e51b815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e2033322062797465730000000000006064820152608401610609565b600882026000611ad78660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b6000826000018281548110611b5457611b54613225565b9060005260206000200154905092915050565b606081600001805480602002602001604051908101604052809291908181526020018280548015611bb757602002820191906000526020600020905b815481526020019060010190808311611ba3575b50505050509050919050565b600061123b8373ffffffffffffffffffffffffffffffffffffffff841661288f565b605254610100900460ff16611c625760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610609565b6107b133611586565b611c73612f32565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b60208110156121e757600182821c8116908190036121935785826020811061216057612160613225565b015460408051602081019290925281018590526060016040516020818303038152906040528051906020012093506121de565b838583602081106121a6576121a6613225565b60200201516040516020016121c5929190918252602082015260400190565b6040516020818303038152906040528051906020012093505b50600101612136565b505092915050565b60006121fb83836128de565b6122e957600061221a61220e8560d81c90565b64ffffffffff16612901565b915050600061222f8464ffffffffff16612901565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e0160405160208183030381529060405290508060405162461bcd60e51b815260040161060991906134f3565b5090919050565b6000816122fd6002611613565b5061231162ffffff1984166026600c61194f565b63ffffffff169392505050565b60008161232b6002611613565b5061231162ffffff198416601a600c61194f565b60008161234c6002611613565b5061231162ffffff198416600e600c61194f565b60008161236d6002611613565b5061231162ffffff1984166002600c61194f565b60b75473ffffffffffffffffffffffffffffffffffffffff1633146107b15760405162461bcd60e51b815260206004820152601060248201527f2173797374656d4d657373656e676572000000000000000000000000000000006044820152606401610609565b602080830154906001906123fd90600261361e565b612407919061332c565b81106124555760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c000000000000000000000000000000006044820152606401610609565b6001016020830181905560005b60208110156124fd5781600116600103612491578284826020811061248957612489613225565b015550505050565b8381602081106124a3576124a3613225565b01546040805160208101929092528101849052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209250600191821c9101612462565b5061250661362a565b505050565b60008061251883856134db565b9050604051811115612528575060005b8060000361253d5762ffffff1991505061123b565b5050606092831b9190911790911b1760181b90565b60008061256d8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050612586866129eb565b8461259187846134db565b61259b91906134db565b11156125ae5762ffffff199150506118ab565b6125b885826134db565b90506125cc8364ffffffffff16828661250b565b9695505050505050565b600062ffffff19808416036126535760405162461bcd60e51b815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f65722064657265660000000000000000000000000000000000000000000000006064820152608401610609565b61265c83612a33565b6126ce5760405162461bcd60e51b815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e7465722064657265660000000000000000000000000000000000000000006064820152608401610609565b60006126e88460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006127128560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006040519050848111156127375760206060fd5b8285848460045afa506125cc61274d8760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b60008061278c8360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006127b68460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b60008060006127dd8585612a70565b915091506127ea81612ab5565b509392505050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054151561123b565b6060600061282e86612901565b915050600061283c86612901565b915050600061284a86612901565b915050600061285886612901565b915050838383836040516020016128729493929190613659565b604051602081830303815290604052945050505050949350505050565b60008181526001830160205260408120546128d657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610744565b506000610744565b60008164ffffffffff166128f28460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff1611156129745760006129208260086133d4565b60ff1685901c905061293181612ca1565b61ffff16841793508160ff1660101461294c57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612907565b50600f5b60ff8160ff1610156129e55760006129918260086133d4565b60ff1685901c90506129a281612ca1565b61ffff16831792508160ff166000146129bd57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612978565b50915091565b6000612a058260181c6bffffffffffffffffffffffff1690565b612a1d8360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b6000612a3f8260d81c90565b64ffffffffff1664ffffffffff03612a5957506000919050565b6000612a64836129eb565b60405110199392505050565b6000808251604103612aa65760208301516040840151606085015160001a612a9a87828585612cd3565b94509450505050612aae565b506000905060025b9250929050565b6000816004811115612ac957612ac9613112565b03612ad15750565b6001816004811115612ae557612ae5613112565b03612b325760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610609565b6002816004811115612b4657612b46613112565b03612b935760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610609565b6003816004811115612ba757612ba7613112565b03612c1a5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610609565b6004816004811115612c2e57612c2e613112565b036108685760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610609565b6000612cb360048360ff16901c612deb565b60ff1661ffff919091161760081b612cca82612deb565b60ff1617919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612d0a5750600090506003612de2565b8460ff16601b14158015612d2257508460ff16601c14155b15612d335750600090506004612de2565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612d87573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612ddb57600060019250925050612de2565b9150600090505b94509492505050565b600060f08083179060ff82169003612e065750603092915050565b8060ff1660f103612e1a5750603192915050565b8060ff1660f203612e2e5750603292915050565b8060ff1660f303612e425750603392915050565b8060ff1660f403612e565750603492915050565b8060ff1660f503612e6a5750603592915050565b8060ff1660f603612e7e5750603692915050565b8060ff1660f703612e925750603792915050565b8060ff1660f803612ea65750603892915050565b8060ff1660f903612eba5750603992915050565b8060ff1660fa03612ece5750606192915050565b8060ff1660fb03612ee25750606292915050565b8060ff1660fc03612ef65750606392915050565b8060ff1660fd03612f0a5750606492915050565b8060ff1660fe03612f1e5750606592915050565b8060ff1660ff036116665750606692915050565b6040518061040001604052806020906020820280368337509192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112612f9157600080fd5b813567ffffffffffffffff80821115612fac57612fac612f51565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715612ff257612ff2612f51565b8160405283815286602085880101111561300b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561303d57600080fd5b813567ffffffffffffffff81111561305457600080fd5b6118ab84828501612f80565b60006020828403121561307257600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156130c757835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613095565b50909695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461086857600080fd5b60006020828403121561310757600080fd5b813561123b816130d3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b602081016003831061317c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b803563ffffffff811681146106da57600080fd5b600080600080600060a086880312156131ae57600080fd5b6131b786613182565b9450602086013593506131cc60408701613182565b9250606086013567ffffffffffffffff808211156131e957600080fd5b6131f589838a01612f80565b9350608088013591508082111561320b57600080fd5b5061321888828901612f80565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b8381101561326f578181015183820152602001613257565b8381111561327e576000848401525b50505050565b6000815180845261329c816020860160208601613254565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006118ab6040830184613284565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561333e5761333e6132fd565b500390565b60006020828403121561335557600080fd5b815161123b816130d3565b600063ffffffff80831681851680830382111561337f5761337f6132fd565b01949350505050565b60408152600061339b6040830185613284565b82810360208401526117908185613284565b60006bffffffffffffffffffffffff80831681851680830382111561337f5761337f6132fd565b600060ff821660ff84168160ff04811182151516156133f5576133f56132fd565b029392505050565b600061ffff80831681851680830382111561337f5761337f6132fd565b60007fffff000000000000000000000000000000000000000000000000000000000000808a60f01b168352808960f01b166002840152808860f01b166004840152808760f01b16600684015250845161347a816008850160208901613254565b845190830190613491816008840160208901613254565b84519101906134a7816008840160208801613254565b016008019998505050505050505050565b600060ff821660ff8416808210156134d2576134d26132fd565b90039392505050565b600082198211156134ee576134ee6132fd565b500190565b60208152600061123b6020830184613284565b600181815b8085111561355f57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115613545576135456132fd565b8085161561355257918102915b93841c939080029061350b565b509250929050565b60008261357657506001610744565b8161358357506000610744565b816001811461359957600281146135a3576135bf565b6001915050610744565b60ff8411156135b4576135b46132fd565b50506001821b610744565b5060208310610133831016604e8410600b84101617156135e2575081810a610744565b6135ec8383613506565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156133f5576133f56132fd565b600061123b8383613567565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d82016125cc56fea2646970667358221220787548b469938009dd905354731829a34ad5324a2fa0aafb8907e53d941ee32964736f6c634300080d0033",
}

// OriginABI is the input ABI used to generate the binding from.
// Deprecated: Use OriginMetaData.ABI instead.
var OriginABI = OriginMetaData.ABI

// Deprecated: Use OriginMetaData.Sigs instead.
// OriginFuncSigs maps the 4-byte function signature to its string representation.
var OriginFuncSigs = OriginMetaData.Sigs

// OriginBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OriginMetaData.Bin instead.
var OriginBin = OriginMetaData.Bin

// DeployOrigin deploys a new Ethereum contract, binding an instance of Origin to it.
func DeployOrigin(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32) (common.Address, *types.Transaction, *Origin, error) {
	parsed, err := OriginMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OriginBin), backend, _localDomain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Origin{OriginCaller: OriginCaller{contract: contract}, OriginTransactor: OriginTransactor{contract: contract}, OriginFilterer: OriginFilterer{contract: contract}}, nil
}

// Origin is an auto generated Go binding around an Ethereum contract.
type Origin struct {
	OriginCaller     // Read-only binding to the contract
	OriginTransactor // Write-only binding to the contract
	OriginFilterer   // Log filterer for contract events
}

// OriginCaller is an auto generated read-only Go binding around an Ethereum contract.
type OriginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OriginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OriginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OriginSession struct {
	Contract     *Origin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OriginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OriginCallerSession struct {
	Contract *OriginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OriginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OriginTransactorSession struct {
	Contract     *OriginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OriginRaw is an auto generated low-level Go binding around an Ethereum contract.
type OriginRaw struct {
	Contract *Origin // Generic contract binding to access the raw methods on
}

// OriginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OriginCallerRaw struct {
	Contract *OriginCaller // Generic read-only contract binding to access the raw methods on
}

// OriginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OriginTransactorRaw struct {
	Contract *OriginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrigin creates a new instance of Origin, bound to a specific deployed contract.
func NewOrigin(address common.Address, backend bind.ContractBackend) (*Origin, error) {
	contract, err := bindOrigin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Origin{OriginCaller: OriginCaller{contract: contract}, OriginTransactor: OriginTransactor{contract: contract}, OriginFilterer: OriginFilterer{contract: contract}}, nil
}

// NewOriginCaller creates a new read-only instance of Origin, bound to a specific deployed contract.
func NewOriginCaller(address common.Address, caller bind.ContractCaller) (*OriginCaller, error) {
	contract, err := bindOrigin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OriginCaller{contract: contract}, nil
}

// NewOriginTransactor creates a new write-only instance of Origin, bound to a specific deployed contract.
func NewOriginTransactor(address common.Address, transactor bind.ContractTransactor) (*OriginTransactor, error) {
	contract, err := bindOrigin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OriginTransactor{contract: contract}, nil
}

// NewOriginFilterer creates a new log filterer instance of Origin, bound to a specific deployed contract.
func NewOriginFilterer(address common.Address, filterer bind.ContractFilterer) (*OriginFilterer, error) {
	contract, err := bindOrigin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OriginFilterer{contract: contract}, nil
}

// bindOrigin binds a generic wrapper to an already deployed contract.
func bindOrigin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OriginABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Origin *OriginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Origin.Contract.OriginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Origin *OriginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Origin.Contract.OriginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Origin *OriginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Origin.Contract.OriginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Origin *OriginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Origin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Origin *OriginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Origin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Origin *OriginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Origin.Contract.contract.Transact(opts, method, params...)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Origin *OriginCaller) MAXMESSAGEBODYBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "MAX_MESSAGE_BODY_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Origin *OriginSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _Origin.Contract.MAXMESSAGEBODYBYTES(&_Origin.CallOpts)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Origin *OriginCallerSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _Origin.Contract.MAXMESSAGEBODYBYTES(&_Origin.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Origin *OriginCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Origin *OriginSession) VERSION() (uint8, error) {
	return _Origin.Contract.VERSION(&_Origin.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Origin *OriginCallerSession) VERSION() (uint8, error) {
	return _Origin.Contract.VERSION(&_Origin.CallOpts)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_Origin *OriginCaller) AllGuards(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "allGuards")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_Origin *OriginSession) AllGuards() ([]common.Address, error) {
	return _Origin.Contract.AllGuards(&_Origin.CallOpts)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_Origin *OriginCallerSession) AllGuards() ([]common.Address, error) {
	return _Origin.Contract.AllGuards(&_Origin.CallOpts)
}

// AllNotaries is a free data retrieval call binding the contract method 0x9817e315.
//
// Solidity: function allNotaries() view returns(address[])
func (_Origin *OriginCaller) AllNotaries(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "allNotaries")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllNotaries is a free data retrieval call binding the contract method 0x9817e315.
//
// Solidity: function allNotaries() view returns(address[])
func (_Origin *OriginSession) AllNotaries() ([]common.Address, error) {
	return _Origin.Contract.AllNotaries(&_Origin.CallOpts)
}

// AllNotaries is a free data retrieval call binding the contract method 0x9817e315.
//
// Solidity: function allNotaries() view returns(address[])
func (_Origin *OriginCallerSession) AllNotaries() ([]common.Address, error) {
	return _Origin.Contract.AllNotaries(&_Origin.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Origin *OriginCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Origin *OriginSession) Count() (*big.Int, error) {
	return _Origin.Contract.Count(&_Origin.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Origin *OriginCallerSession) Count() (*big.Int, error) {
	return _Origin.Contract.Count(&_Origin.CallOpts)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_Origin *OriginCaller) GetGuard(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "getGuard", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_Origin *OriginSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _Origin.Contract.GetGuard(&_Origin.CallOpts, _index)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_Origin *OriginCallerSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _Origin.Contract.GetGuard(&_Origin.CallOpts, _index)
}

// GetNotary is a free data retrieval call binding the contract method 0xc07dc7f5.
//
// Solidity: function getNotary(uint256 _index) view returns(address)
func (_Origin *OriginCaller) GetNotary(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "getNotary", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNotary is a free data retrieval call binding the contract method 0xc07dc7f5.
//
// Solidity: function getNotary(uint256 _index) view returns(address)
func (_Origin *OriginSession) GetNotary(_index *big.Int) (common.Address, error) {
	return _Origin.Contract.GetNotary(&_Origin.CallOpts, _index)
}

// GetNotary is a free data retrieval call binding the contract method 0xc07dc7f5.
//
// Solidity: function getNotary(uint256 _index) view returns(address)
func (_Origin *OriginCallerSession) GetNotary(_index *big.Int) (common.Address, error) {
	return _Origin.Contract.GetNotary(&_Origin.CallOpts, _index)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_Origin *OriginCaller) GuardsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "guardsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_Origin *OriginSession) GuardsAmount() (*big.Int, error) {
	return _Origin.Contract.GuardsAmount(&_Origin.CallOpts)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_Origin *OriginCallerSession) GuardsAmount() (*big.Int, error) {
	return _Origin.Contract.GuardsAmount(&_Origin.CallOpts)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_Origin *OriginCaller) HistoricalRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "historicalRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_Origin *OriginSession) HistoricalRoots(arg0 *big.Int) ([32]byte, error) {
	return _Origin.Contract.HistoricalRoots(&_Origin.CallOpts, arg0)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_Origin *OriginCallerSession) HistoricalRoots(arg0 *big.Int) ([32]byte, error) {
	return _Origin.Contract.HistoricalRoots(&_Origin.CallOpts, arg0)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Origin *OriginCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Origin *OriginSession) LocalDomain() (uint32, error) {
	return _Origin.Contract.LocalDomain(&_Origin.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Origin *OriginCallerSession) LocalDomain() (uint32, error) {
	return _Origin.Contract.LocalDomain(&_Origin.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Origin *OriginCaller) Nonce(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Origin *OriginSession) Nonce() (uint32, error) {
	return _Origin.Contract.Nonce(&_Origin.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Origin *OriginCallerSession) Nonce() (uint32, error) {
	return _Origin.Contract.Nonce(&_Origin.CallOpts)
}

// NotariesAmount is a free data retrieval call binding the contract method 0x8e62e9ef.
//
// Solidity: function notariesAmount() view returns(uint256)
func (_Origin *OriginCaller) NotariesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "notariesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NotariesAmount is a free data retrieval call binding the contract method 0x8e62e9ef.
//
// Solidity: function notariesAmount() view returns(uint256)
func (_Origin *OriginSession) NotariesAmount() (*big.Int, error) {
	return _Origin.Contract.NotariesAmount(&_Origin.CallOpts)
}

// NotariesAmount is a free data retrieval call binding the contract method 0x8e62e9ef.
//
// Solidity: function notariesAmount() view returns(uint256)
func (_Origin *OriginCallerSession) NotariesAmount() (*big.Int, error) {
	return _Origin.Contract.NotariesAmount(&_Origin.CallOpts)
}

// NotaryManager is a free data retrieval call binding the contract method 0xf85b597e.
//
// Solidity: function notaryManager() view returns(address)
func (_Origin *OriginCaller) NotaryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "notaryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NotaryManager is a free data retrieval call binding the contract method 0xf85b597e.
//
// Solidity: function notaryManager() view returns(address)
func (_Origin *OriginSession) NotaryManager() (common.Address, error) {
	return _Origin.Contract.NotaryManager(&_Origin.CallOpts)
}

// NotaryManager is a free data retrieval call binding the contract method 0xf85b597e.
//
// Solidity: function notaryManager() view returns(address)
func (_Origin *OriginCallerSession) NotaryManager() (common.Address, error) {
	return _Origin.Contract.NotaryManager(&_Origin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Origin *OriginCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Origin *OriginSession) Owner() (common.Address, error) {
	return _Origin.Contract.Owner(&_Origin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Origin *OriginCallerSession) Owner() (common.Address, error) {
	return _Origin.Contract.Owner(&_Origin.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Origin *OriginCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Origin *OriginSession) Root() ([32]byte, error) {
	return _Origin.Contract.Root(&_Origin.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Origin *OriginCallerSession) Root() ([32]byte, error) {
	return _Origin.Contract.Root(&_Origin.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Origin *OriginCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Origin *OriginSession) State() (uint8, error) {
	return _Origin.Contract.State(&_Origin.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Origin *OriginCallerSession) State() (uint8, error) {
	return _Origin.Contract.State(&_Origin.CallOpts)
}

// SuggestAttestation is a free data retrieval call binding the contract method 0x524787d0.
//
// Solidity: function suggestAttestation() view returns(uint32 _nonce, bytes32 _root)
func (_Origin *OriginCaller) SuggestAttestation(opts *bind.CallOpts) (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "suggestAttestation")

	outstruct := new(struct {
		Nonce uint32
		Root  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Root = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SuggestAttestation is a free data retrieval call binding the contract method 0x524787d0.
//
// Solidity: function suggestAttestation() view returns(uint32 _nonce, bytes32 _root)
func (_Origin *OriginSession) SuggestAttestation() (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	return _Origin.Contract.SuggestAttestation(&_Origin.CallOpts)
}

// SuggestAttestation is a free data retrieval call binding the contract method 0x524787d0.
//
// Solidity: function suggestAttestation() view returns(uint32 _nonce, bytes32 _root)
func (_Origin *OriginCallerSession) SuggestAttestation() (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	return _Origin.Contract.SuggestAttestation(&_Origin.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_Origin *OriginCaller) SystemMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "systemMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_Origin *OriginSession) SystemMessenger() (common.Address, error) {
	return _Origin.Contract.SystemMessenger(&_Origin.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_Origin *OriginCallerSession) SystemMessenger() (common.Address, error) {
	return _Origin.Contract.SystemMessenger(&_Origin.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Origin *OriginCaller) Tree(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "tree")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Origin *OriginSession) Tree() (*big.Int, error) {
	return _Origin.Contract.Tree(&_Origin.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Origin *OriginCallerSession) Tree() (*big.Int, error) {
	return _Origin.Contract.Tree(&_Origin.CallOpts)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_Origin *OriginTransactor) Dispatch(opts *bind.TransactOpts, _destination uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "dispatch", _destination, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_Origin *OriginSession) Dispatch(_destination uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Origin.Contract.Dispatch(&_Origin.TransactOpts, _destination, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_Origin *OriginTransactorSession) Dispatch(_destination uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Origin.Contract.Dispatch(&_Origin.TransactOpts, _destination, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x0afe7f90.
//
// Solidity: function improperAttestation(bytes _attestation) returns(bool)
func (_Origin *OriginTransactor) ImproperAttestation(opts *bind.TransactOpts, _attestation []byte) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "improperAttestation", _attestation)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x0afe7f90.
//
// Solidity: function improperAttestation(bytes _attestation) returns(bool)
func (_Origin *OriginSession) ImproperAttestation(_attestation []byte) (*types.Transaction, error) {
	return _Origin.Contract.ImproperAttestation(&_Origin.TransactOpts, _attestation)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x0afe7f90.
//
// Solidity: function improperAttestation(bytes _attestation) returns(bool)
func (_Origin *OriginTransactorSession) ImproperAttestation(_attestation []byte) (*types.Transaction, error) {
	return _Origin.Contract.ImproperAttestation(&_Origin.TransactOpts, _attestation)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _notaryManager) returns()
func (_Origin *OriginTransactor) Initialize(opts *bind.TransactOpts, _notaryManager common.Address) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "initialize", _notaryManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _notaryManager) returns()
func (_Origin *OriginSession) Initialize(_notaryManager common.Address) (*types.Transaction, error) {
	return _Origin.Contract.Initialize(&_Origin.TransactOpts, _notaryManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _notaryManager) returns()
func (_Origin *OriginTransactorSession) Initialize(_notaryManager common.Address) (*types.Transaction, error) {
	return _Origin.Contract.Initialize(&_Origin.TransactOpts, _notaryManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Origin *OriginTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Origin *OriginSession) RenounceOwnership() (*types.Transaction, error) {
	return _Origin.Contract.RenounceOwnership(&_Origin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Origin *OriginTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Origin.Contract.RenounceOwnership(&_Origin.TransactOpts)
}

// SetNotary is a paid mutator transaction binding the contract method 0xa394a0e6.
//
// Solidity: function setNotary(address _notary) returns()
func (_Origin *OriginTransactor) SetNotary(opts *bind.TransactOpts, _notary common.Address) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "setNotary", _notary)
}

// SetNotary is a paid mutator transaction binding the contract method 0xa394a0e6.
//
// Solidity: function setNotary(address _notary) returns()
func (_Origin *OriginSession) SetNotary(_notary common.Address) (*types.Transaction, error) {
	return _Origin.Contract.SetNotary(&_Origin.TransactOpts, _notary)
}

// SetNotary is a paid mutator transaction binding the contract method 0xa394a0e6.
//
// Solidity: function setNotary(address _notary) returns()
func (_Origin *OriginTransactorSession) SetNotary(_notary common.Address) (*types.Transaction, error) {
	return _Origin.Contract.SetNotary(&_Origin.TransactOpts, _notary)
}

// SetNotaryManager is a paid mutator transaction binding the contract method 0xa340abc1.
//
// Solidity: function setNotaryManager(address _notaryManager) returns()
func (_Origin *OriginTransactor) SetNotaryManager(opts *bind.TransactOpts, _notaryManager common.Address) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "setNotaryManager", _notaryManager)
}

// SetNotaryManager is a paid mutator transaction binding the contract method 0xa340abc1.
//
// Solidity: function setNotaryManager(address _notaryManager) returns()
func (_Origin *OriginSession) SetNotaryManager(_notaryManager common.Address) (*types.Transaction, error) {
	return _Origin.Contract.SetNotaryManager(&_Origin.TransactOpts, _notaryManager)
}

// SetNotaryManager is a paid mutator transaction binding the contract method 0xa340abc1.
//
// Solidity: function setNotaryManager(address _notaryManager) returns()
func (_Origin *OriginTransactorSession) SetNotaryManager(_notaryManager common.Address) (*types.Transaction, error) {
	return _Origin.Contract.SetNotaryManager(&_Origin.TransactOpts, _notaryManager)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_Origin *OriginTransactor) SetSystemMessenger(opts *bind.TransactOpts, _systemMessenger common.Address) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "setSystemMessenger", _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_Origin *OriginSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _Origin.Contract.SetSystemMessenger(&_Origin.TransactOpts, _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_Origin *OriginTransactorSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _Origin.Contract.SetSystemMessenger(&_Origin.TransactOpts, _systemMessenger)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Origin *OriginTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Origin *OriginSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Origin.Contract.TransferOwnership(&_Origin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Origin *OriginTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Origin.Contract.TransferOwnership(&_Origin.TransactOpts, newOwner)
}

// OriginDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the Origin contract.
type OriginDispatchIterator struct {
	Event *OriginDispatch // Event containing the contract specifics and raw log

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
func (it *OriginDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginDispatch)
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
		it.Event = new(OriginDispatch)
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
func (it *OriginDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginDispatch represents a Dispatch event raised by the Origin contract.
type OriginDispatch struct {
	MessageHash         [32]byte
	LeafIndex           *big.Int
	DestinationAndNonce uint64
	Tips                []byte
	Message             []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0x718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e785.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes tips, bytes message)
func (_Origin *OriginFilterer) FilterDispatch(opts *bind.FilterOpts, messageHash [][32]byte, leafIndex []*big.Int, destinationAndNonce []uint64) (*OriginDispatchIterator, error) {

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

	logs, sub, err := _Origin.contract.FilterLogs(opts, "Dispatch", messageHashRule, leafIndexRule, destinationAndNonceRule)
	if err != nil {
		return nil, err
	}
	return &OriginDispatchIterator{contract: _Origin.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0x718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e785.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes tips, bytes message)
func (_Origin *OriginFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *OriginDispatch, messageHash [][32]byte, leafIndex []*big.Int, destinationAndNonce []uint64) (event.Subscription, error) {

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

	logs, sub, err := _Origin.contract.WatchLogs(opts, "Dispatch", messageHashRule, leafIndexRule, destinationAndNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginDispatch)
				if err := _Origin.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0x718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e785.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes tips, bytes message)
func (_Origin *OriginFilterer) ParseDispatch(log types.Log) (*OriginDispatch, error) {
	event := new(OriginDispatch)
	if err := _Origin.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginDomainNotaryAddedIterator is returned from FilterDomainNotaryAdded and is used to iterate over the raw logs and unpacked data for DomainNotaryAdded events raised by the Origin contract.
type OriginDomainNotaryAddedIterator struct {
	Event *OriginDomainNotaryAdded // Event containing the contract specifics and raw log

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
func (it *OriginDomainNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginDomainNotaryAdded)
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
		it.Event = new(OriginDomainNotaryAdded)
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
func (it *OriginDomainNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginDomainNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginDomainNotaryAdded represents a DomainNotaryAdded event raised by the Origin contract.
type OriginDomainNotaryAdded struct {
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainNotaryAdded is a free log retrieval operation binding the contract event 0x7ed5310d8818d06ea4a196771a39a73bf55c815addbf7a52ba87c9be409c3dd1.
//
// Solidity: event DomainNotaryAdded(address notary)
func (_Origin *OriginFilterer) FilterDomainNotaryAdded(opts *bind.FilterOpts) (*OriginDomainNotaryAddedIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "DomainNotaryAdded")
	if err != nil {
		return nil, err
	}
	return &OriginDomainNotaryAddedIterator{contract: _Origin.contract, event: "DomainNotaryAdded", logs: logs, sub: sub}, nil
}

// WatchDomainNotaryAdded is a free log subscription operation binding the contract event 0x7ed5310d8818d06ea4a196771a39a73bf55c815addbf7a52ba87c9be409c3dd1.
//
// Solidity: event DomainNotaryAdded(address notary)
func (_Origin *OriginFilterer) WatchDomainNotaryAdded(opts *bind.WatchOpts, sink chan<- *OriginDomainNotaryAdded) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "DomainNotaryAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginDomainNotaryAdded)
				if err := _Origin.contract.UnpackLog(event, "DomainNotaryAdded", log); err != nil {
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

// ParseDomainNotaryAdded is a log parse operation binding the contract event 0x7ed5310d8818d06ea4a196771a39a73bf55c815addbf7a52ba87c9be409c3dd1.
//
// Solidity: event DomainNotaryAdded(address notary)
func (_Origin *OriginFilterer) ParseDomainNotaryAdded(log types.Log) (*OriginDomainNotaryAdded, error) {
	event := new(OriginDomainNotaryAdded)
	if err := _Origin.contract.UnpackLog(event, "DomainNotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginDomainNotaryRemovedIterator is returned from FilterDomainNotaryRemoved and is used to iterate over the raw logs and unpacked data for DomainNotaryRemoved events raised by the Origin contract.
type OriginDomainNotaryRemovedIterator struct {
	Event *OriginDomainNotaryRemoved // Event containing the contract specifics and raw log

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
func (it *OriginDomainNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginDomainNotaryRemoved)
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
		it.Event = new(OriginDomainNotaryRemoved)
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
func (it *OriginDomainNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginDomainNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginDomainNotaryRemoved represents a DomainNotaryRemoved event raised by the Origin contract.
type OriginDomainNotaryRemoved struct {
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainNotaryRemoved is a free log retrieval operation binding the contract event 0xe16811bec5badeb0bade36ad31aab1c20f2997b625833474449f893eeecd3bac.
//
// Solidity: event DomainNotaryRemoved(address notary)
func (_Origin *OriginFilterer) FilterDomainNotaryRemoved(opts *bind.FilterOpts) (*OriginDomainNotaryRemovedIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "DomainNotaryRemoved")
	if err != nil {
		return nil, err
	}
	return &OriginDomainNotaryRemovedIterator{contract: _Origin.contract, event: "DomainNotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchDomainNotaryRemoved is a free log subscription operation binding the contract event 0xe16811bec5badeb0bade36ad31aab1c20f2997b625833474449f893eeecd3bac.
//
// Solidity: event DomainNotaryRemoved(address notary)
func (_Origin *OriginFilterer) WatchDomainNotaryRemoved(opts *bind.WatchOpts, sink chan<- *OriginDomainNotaryRemoved) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "DomainNotaryRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginDomainNotaryRemoved)
				if err := _Origin.contract.UnpackLog(event, "DomainNotaryRemoved", log); err != nil {
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

// ParseDomainNotaryRemoved is a log parse operation binding the contract event 0xe16811bec5badeb0bade36ad31aab1c20f2997b625833474449f893eeecd3bac.
//
// Solidity: event DomainNotaryRemoved(address notary)
func (_Origin *OriginFilterer) ParseDomainNotaryRemoved(log types.Log) (*OriginDomainNotaryRemoved, error) {
	event := new(OriginDomainNotaryRemoved)
	if err := _Origin.contract.UnpackLog(event, "DomainNotaryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the Origin contract.
type OriginGuardAddedIterator struct {
	Event *OriginGuardAdded // Event containing the contract specifics and raw log

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
func (it *OriginGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginGuardAdded)
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
		it.Event = new(OriginGuardAdded)
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
func (it *OriginGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginGuardAdded represents a GuardAdded event raised by the Origin contract.
type OriginGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_Origin *OriginFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*OriginGuardAddedIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &OriginGuardAddedIterator{contract: _Origin.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_Origin *OriginFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *OriginGuardAdded) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginGuardAdded)
				if err := _Origin.contract.UnpackLog(event, "GuardAdded", log); err != nil {
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

// ParseGuardAdded is a log parse operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_Origin *OriginFilterer) ParseGuardAdded(log types.Log) (*OriginGuardAdded, error) {
	event := new(OriginGuardAdded)
	if err := _Origin.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the Origin contract.
type OriginGuardRemovedIterator struct {
	Event *OriginGuardRemoved // Event containing the contract specifics and raw log

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
func (it *OriginGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginGuardRemoved)
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
		it.Event = new(OriginGuardRemoved)
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
func (it *OriginGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginGuardRemoved represents a GuardRemoved event raised by the Origin contract.
type OriginGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_Origin *OriginFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*OriginGuardRemovedIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &OriginGuardRemovedIterator{contract: _Origin.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_Origin *OriginFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *OriginGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginGuardRemoved)
				if err := _Origin.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
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

// ParseGuardRemoved is a log parse operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_Origin *OriginFilterer) ParseGuardRemoved(log types.Log) (*OriginGuardRemoved, error) {
	event := new(OriginGuardRemoved)
	if err := _Origin.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginImproperAttestationIterator is returned from FilterImproperAttestation and is used to iterate over the raw logs and unpacked data for ImproperAttestation events raised by the Origin contract.
type OriginImproperAttestationIterator struct {
	Event *OriginImproperAttestation // Event containing the contract specifics and raw log

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
func (it *OriginImproperAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginImproperAttestation)
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
		it.Event = new(OriginImproperAttestation)
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
func (it *OriginImproperAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginImproperAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginImproperAttestation represents a ImproperAttestation event raised by the Origin contract.
type OriginImproperAttestation struct {
	Notary      common.Address
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterImproperAttestation is a free log retrieval operation binding the contract event 0x287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b.
//
// Solidity: event ImproperAttestation(address notary, bytes attestation)
func (_Origin *OriginFilterer) FilterImproperAttestation(opts *bind.FilterOpts) (*OriginImproperAttestationIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "ImproperAttestation")
	if err != nil {
		return nil, err
	}
	return &OriginImproperAttestationIterator{contract: _Origin.contract, event: "ImproperAttestation", logs: logs, sub: sub}, nil
}

// WatchImproperAttestation is a free log subscription operation binding the contract event 0x287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b.
//
// Solidity: event ImproperAttestation(address notary, bytes attestation)
func (_Origin *OriginFilterer) WatchImproperAttestation(opts *bind.WatchOpts, sink chan<- *OriginImproperAttestation) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "ImproperAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginImproperAttestation)
				if err := _Origin.contract.UnpackLog(event, "ImproperAttestation", log); err != nil {
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

// ParseImproperAttestation is a log parse operation binding the contract event 0x287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b.
//
// Solidity: event ImproperAttestation(address notary, bytes attestation)
func (_Origin *OriginFilterer) ParseImproperAttestation(log types.Log) (*OriginImproperAttestation, error) {
	event := new(OriginImproperAttestation)
	if err := _Origin.contract.UnpackLog(event, "ImproperAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Origin contract.
type OriginInitializedIterator struct {
	Event *OriginInitialized // Event containing the contract specifics and raw log

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
func (it *OriginInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginInitialized)
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
		it.Event = new(OriginInitialized)
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
func (it *OriginInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginInitialized represents a Initialized event raised by the Origin contract.
type OriginInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Origin *OriginFilterer) FilterInitialized(opts *bind.FilterOpts) (*OriginInitializedIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OriginInitializedIterator{contract: _Origin.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Origin *OriginFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OriginInitialized) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginInitialized)
				if err := _Origin.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Origin *OriginFilterer) ParseInitialized(log types.Log) (*OriginInitialized, error) {
	event := new(OriginInitialized)
	if err := _Origin.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginNewNotaryManagerIterator is returned from FilterNewNotaryManager and is used to iterate over the raw logs and unpacked data for NewNotaryManager events raised by the Origin contract.
type OriginNewNotaryManagerIterator struct {
	Event *OriginNewNotaryManager // Event containing the contract specifics and raw log

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
func (it *OriginNewNotaryManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginNewNotaryManager)
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
		it.Event = new(OriginNewNotaryManager)
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
func (it *OriginNewNotaryManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginNewNotaryManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginNewNotaryManager represents a NewNotaryManager event raised by the Origin contract.
type OriginNewNotaryManager struct {
	NotaryManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewNotaryManager is a free log retrieval operation binding the contract event 0xe3befd3a32a53f50ff7d1421555fbd40e5ead3a7ed75417db43a23faffe09316.
//
// Solidity: event NewNotaryManager(address notaryManager)
func (_Origin *OriginFilterer) FilterNewNotaryManager(opts *bind.FilterOpts) (*OriginNewNotaryManagerIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "NewNotaryManager")
	if err != nil {
		return nil, err
	}
	return &OriginNewNotaryManagerIterator{contract: _Origin.contract, event: "NewNotaryManager", logs: logs, sub: sub}, nil
}

// WatchNewNotaryManager is a free log subscription operation binding the contract event 0xe3befd3a32a53f50ff7d1421555fbd40e5ead3a7ed75417db43a23faffe09316.
//
// Solidity: event NewNotaryManager(address notaryManager)
func (_Origin *OriginFilterer) WatchNewNotaryManager(opts *bind.WatchOpts, sink chan<- *OriginNewNotaryManager) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "NewNotaryManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginNewNotaryManager)
				if err := _Origin.contract.UnpackLog(event, "NewNotaryManager", log); err != nil {
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

// ParseNewNotaryManager is a log parse operation binding the contract event 0xe3befd3a32a53f50ff7d1421555fbd40e5ead3a7ed75417db43a23faffe09316.
//
// Solidity: event NewNotaryManager(address notaryManager)
func (_Origin *OriginFilterer) ParseNewNotaryManager(log types.Log) (*OriginNewNotaryManager, error) {
	event := new(OriginNewNotaryManager)
	if err := _Origin.contract.UnpackLog(event, "NewNotaryManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginNotarySlashedIterator is returned from FilterNotarySlashed and is used to iterate over the raw logs and unpacked data for NotarySlashed events raised by the Origin contract.
type OriginNotarySlashedIterator struct {
	Event *OriginNotarySlashed // Event containing the contract specifics and raw log

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
func (it *OriginNotarySlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginNotarySlashed)
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
		it.Event = new(OriginNotarySlashed)
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
func (it *OriginNotarySlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginNotarySlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginNotarySlashed represents a NotarySlashed event raised by the Origin contract.
type OriginNotarySlashed struct {
	Notary   common.Address
	Reporter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNotarySlashed is a free log retrieval operation binding the contract event 0x9ad95700c0e79cb6384afc527c4e16b94a9c6e39ea2ba3824ad6f856bdc4a401.
//
// Solidity: event NotarySlashed(address indexed notary, address indexed reporter)
func (_Origin *OriginFilterer) FilterNotarySlashed(opts *bind.FilterOpts, notary []common.Address, reporter []common.Address) (*OriginNotarySlashedIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "NotarySlashed", notaryRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return &OriginNotarySlashedIterator{contract: _Origin.contract, event: "NotarySlashed", logs: logs, sub: sub}, nil
}

// WatchNotarySlashed is a free log subscription operation binding the contract event 0x9ad95700c0e79cb6384afc527c4e16b94a9c6e39ea2ba3824ad6f856bdc4a401.
//
// Solidity: event NotarySlashed(address indexed notary, address indexed reporter)
func (_Origin *OriginFilterer) WatchNotarySlashed(opts *bind.WatchOpts, sink chan<- *OriginNotarySlashed, notary []common.Address, reporter []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "NotarySlashed", notaryRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginNotarySlashed)
				if err := _Origin.contract.UnpackLog(event, "NotarySlashed", log); err != nil {
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

// ParseNotarySlashed is a log parse operation binding the contract event 0x9ad95700c0e79cb6384afc527c4e16b94a9c6e39ea2ba3824ad6f856bdc4a401.
//
// Solidity: event NotarySlashed(address indexed notary, address indexed reporter)
func (_Origin *OriginFilterer) ParseNotarySlashed(log types.Log) (*OriginNotarySlashed, error) {
	event := new(OriginNotarySlashed)
	if err := _Origin.contract.UnpackLog(event, "NotarySlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Origin contract.
type OriginOwnershipTransferredIterator struct {
	Event *OriginOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OriginOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginOwnershipTransferred)
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
		it.Event = new(OriginOwnershipTransferred)
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
func (it *OriginOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginOwnershipTransferred represents a OwnershipTransferred event raised by the Origin contract.
type OriginOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Origin *OriginFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OriginOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OriginOwnershipTransferredIterator{contract: _Origin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Origin *OriginFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OriginOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginOwnershipTransferred)
				if err := _Origin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Origin *OriginFilterer) ParseOwnershipTransferred(log types.Log) (*OriginOwnershipTransferred, error) {
	event := new(OriginOwnershipTransferred)
	if err := _Origin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e9a77e211943312c1f6233ac8cd472f7602007725a4e88c454b631749172774264736f6c634300080d0033",
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

// SystemContractMetaData contains all meta data concerning the SystemContract contract.
var SystemContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"_systemMessenger\",\"type\":\"address\"}],\"name\":\"setSystemMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemMessenger\",\"outputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"b7bc563e": "setSystemMessenger(address)",
		"ccbdf9c9": "systemMessenger()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// SystemContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemContractMetaData.ABI instead.
var SystemContractABI = SystemContractMetaData.ABI

// Deprecated: Use SystemContractMetaData.Sigs instead.
// SystemContractFuncSigs maps the 4-byte function signature to its string representation.
var SystemContractFuncSigs = SystemContractMetaData.Sigs

// SystemContract is an auto generated Go binding around an Ethereum contract.
type SystemContract struct {
	SystemContractCaller     // Read-only binding to the contract
	SystemContractTransactor // Write-only binding to the contract
	SystemContractFilterer   // Log filterer for contract events
}

// SystemContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemContractSession struct {
	Contract     *SystemContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemContractCallerSession struct {
	Contract *SystemContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SystemContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemContractTransactorSession struct {
	Contract     *SystemContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SystemContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemContractRaw struct {
	Contract *SystemContract // Generic contract binding to access the raw methods on
}

// SystemContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemContractCallerRaw struct {
	Contract *SystemContractCaller // Generic read-only contract binding to access the raw methods on
}

// SystemContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemContractTransactorRaw struct {
	Contract *SystemContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemContract creates a new instance of SystemContract, bound to a specific deployed contract.
func NewSystemContract(address common.Address, backend bind.ContractBackend) (*SystemContract, error) {
	contract, err := bindSystemContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemContract{SystemContractCaller: SystemContractCaller{contract: contract}, SystemContractTransactor: SystemContractTransactor{contract: contract}, SystemContractFilterer: SystemContractFilterer{contract: contract}}, nil
}

// NewSystemContractCaller creates a new read-only instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractCaller(address common.Address, caller bind.ContractCaller) (*SystemContractCaller, error) {
	contract, err := bindSystemContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractCaller{contract: contract}, nil
}

// NewSystemContractTransactor creates a new write-only instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemContractTransactor, error) {
	contract, err := bindSystemContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractTransactor{contract: contract}, nil
}

// NewSystemContractFilterer creates a new log filterer instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemContractFilterer, error) {
	contract, err := bindSystemContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemContractFilterer{contract: contract}, nil
}

// bindSystemContract binds a generic wrapper to an already deployed contract.
func bindSystemContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContract *SystemContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContract.Contract.SystemContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContract *SystemContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContract.Contract.SystemContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContract *SystemContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContract.Contract.SystemContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContract *SystemContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContract *SystemContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContract *SystemContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContract.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContract *SystemContractCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContract *SystemContractSession) LocalDomain() (uint32, error) {
	return _SystemContract.Contract.LocalDomain(&_SystemContract.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContract *SystemContractCallerSession) LocalDomain() (uint32, error) {
	return _SystemContract.Contract.LocalDomain(&_SystemContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContract *SystemContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContract *SystemContractSession) Owner() (common.Address, error) {
	return _SystemContract.Contract.Owner(&_SystemContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContract *SystemContractCallerSession) Owner() (common.Address, error) {
	return _SystemContract.Contract.Owner(&_SystemContract.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_SystemContract *SystemContractCaller) SystemMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "systemMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_SystemContract *SystemContractSession) SystemMessenger() (common.Address, error) {
	return _SystemContract.Contract.SystemMessenger(&_SystemContract.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_SystemContract *SystemContractCallerSession) SystemMessenger() (common.Address, error) {
	return _SystemContract.Contract.SystemMessenger(&_SystemContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContract *SystemContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContract *SystemContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemContract.Contract.RenounceOwnership(&_SystemContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContract *SystemContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemContract.Contract.RenounceOwnership(&_SystemContract.TransactOpts)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_SystemContract *SystemContractTransactor) SetSystemMessenger(opts *bind.TransactOpts, _systemMessenger common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setSystemMessenger", _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_SystemContract *SystemContractSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetSystemMessenger(&_SystemContract.TransactOpts, _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_SystemContract *SystemContractTransactorSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetSystemMessenger(&_SystemContract.TransactOpts, _systemMessenger)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContract *SystemContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContract *SystemContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.TransferOwnership(&_SystemContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContract *SystemContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.TransferOwnership(&_SystemContract.TransactOpts, newOwner)
}

// SystemContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemContract contract.
type SystemContractInitializedIterator struct {
	Event *SystemContractInitialized // Event containing the contract specifics and raw log

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
func (it *SystemContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractInitialized)
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
		it.Event = new(SystemContractInitialized)
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
func (it *SystemContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractInitialized represents a Initialized event raised by the SystemContract contract.
type SystemContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemContract *SystemContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemContractInitializedIterator, error) {

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemContractInitializedIterator{contract: _SystemContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemContract *SystemContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemContractInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractInitialized)
				if err := _SystemContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SystemContract *SystemContractFilterer) ParseInitialized(log types.Log) (*SystemContractInitialized, error) {
	event := new(SystemContractInitialized)
	if err := _SystemContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemContract contract.
type SystemContractOwnershipTransferredIterator struct {
	Event *SystemContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractOwnershipTransferred)
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
		it.Event = new(SystemContractOwnershipTransferred)
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
func (it *SystemContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractOwnershipTransferred represents a OwnershipTransferred event raised by the SystemContract contract.
type SystemContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemContract *SystemContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemContractOwnershipTransferredIterator{contract: _SystemContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemContract *SystemContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractOwnershipTransferred)
				if err := _SystemContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SystemContract *SystemContractFilterer) ParseOwnershipTransferred(log types.Log) (*SystemContractOwnershipTransferred, error) {
	event := new(SystemContractOwnershipTransferred)
	if err := _SystemContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemMessageMetaData contains all meta data concerning the SystemMessage contract.
var SystemMessageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122027529421adade73122a75517a71c500924777e887f10b28ba7917bdd1edacf4d64736f6c634300080d0033",
}

// SystemMessageABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemMessageMetaData.ABI instead.
var SystemMessageABI = SystemMessageMetaData.ABI

// SystemMessageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemMessageMetaData.Bin instead.
var SystemMessageBin = SystemMessageMetaData.Bin

// DeploySystemMessage deploys a new Ethereum contract, binding an instance of SystemMessage to it.
func DeploySystemMessage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemMessage, error) {
	parsed, err := SystemMessageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemMessageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemMessage{SystemMessageCaller: SystemMessageCaller{contract: contract}, SystemMessageTransactor: SystemMessageTransactor{contract: contract}, SystemMessageFilterer: SystemMessageFilterer{contract: contract}}, nil
}

// SystemMessage is an auto generated Go binding around an Ethereum contract.
type SystemMessage struct {
	SystemMessageCaller     // Read-only binding to the contract
	SystemMessageTransactor // Write-only binding to the contract
	SystemMessageFilterer   // Log filterer for contract events
}

// SystemMessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemMessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemMessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemMessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemMessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemMessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemMessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemMessageSession struct {
	Contract     *SystemMessage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemMessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemMessageCallerSession struct {
	Contract *SystemMessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SystemMessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemMessageTransactorSession struct {
	Contract     *SystemMessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SystemMessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemMessageRaw struct {
	Contract *SystemMessage // Generic contract binding to access the raw methods on
}

// SystemMessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemMessageCallerRaw struct {
	Contract *SystemMessageCaller // Generic read-only contract binding to access the raw methods on
}

// SystemMessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemMessageTransactorRaw struct {
	Contract *SystemMessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemMessage creates a new instance of SystemMessage, bound to a specific deployed contract.
func NewSystemMessage(address common.Address, backend bind.ContractBackend) (*SystemMessage, error) {
	contract, err := bindSystemMessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemMessage{SystemMessageCaller: SystemMessageCaller{contract: contract}, SystemMessageTransactor: SystemMessageTransactor{contract: contract}, SystemMessageFilterer: SystemMessageFilterer{contract: contract}}, nil
}

// NewSystemMessageCaller creates a new read-only instance of SystemMessage, bound to a specific deployed contract.
func NewSystemMessageCaller(address common.Address, caller bind.ContractCaller) (*SystemMessageCaller, error) {
	contract, err := bindSystemMessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemMessageCaller{contract: contract}, nil
}

// NewSystemMessageTransactor creates a new write-only instance of SystemMessage, bound to a specific deployed contract.
func NewSystemMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemMessageTransactor, error) {
	contract, err := bindSystemMessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemMessageTransactor{contract: contract}, nil
}

// NewSystemMessageFilterer creates a new log filterer instance of SystemMessage, bound to a specific deployed contract.
func NewSystemMessageFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemMessageFilterer, error) {
	contract, err := bindSystemMessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemMessageFilterer{contract: contract}, nil
}

// bindSystemMessage binds a generic wrapper to an already deployed contract.
func bindSystemMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemMessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemMessage *SystemMessageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemMessage.Contract.SystemMessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemMessage *SystemMessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemMessage.Contract.SystemMessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemMessage *SystemMessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemMessage.Contract.SystemMessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemMessage *SystemMessageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemMessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemMessage *SystemMessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemMessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemMessage *SystemMessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemMessage.Contract.contract.Transact(opts, method, params...)
}

// TipsMetaData contains all meta data concerning the Tips contract.
var TipsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206227045dbe3377c6d79ee1e984b8319b303b1a0b6680b27c18f8c8768e05217564736f6c634300080d0033",
}

// TipsABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsMetaData.ABI instead.
var TipsABI = TipsMetaData.ABI

// TipsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TipsMetaData.Bin instead.
var TipsBin = TipsMetaData.Bin

// DeployTips deploys a new Ethereum contract, binding an instance of Tips to it.
func DeployTips(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tips, error) {
	parsed, err := TipsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TipsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tips{TipsCaller: TipsCaller{contract: contract}, TipsTransactor: TipsTransactor{contract: contract}, TipsFilterer: TipsFilterer{contract: contract}}, nil
}

// Tips is an auto generated Go binding around an Ethereum contract.
type Tips struct {
	TipsCaller     // Read-only binding to the contract
	TipsTransactor // Write-only binding to the contract
	TipsFilterer   // Log filterer for contract events
}

// TipsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsSession struct {
	Contract     *Tips             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsCallerSession struct {
	Contract *TipsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TipsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsTransactorSession struct {
	Contract     *TipsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsRaw struct {
	Contract *Tips // Generic contract binding to access the raw methods on
}

// TipsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsCallerRaw struct {
	Contract *TipsCaller // Generic read-only contract binding to access the raw methods on
}

// TipsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsTransactorRaw struct {
	Contract *TipsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTips creates a new instance of Tips, bound to a specific deployed contract.
func NewTips(address common.Address, backend bind.ContractBackend) (*Tips, error) {
	contract, err := bindTips(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tips{TipsCaller: TipsCaller{contract: contract}, TipsTransactor: TipsTransactor{contract: contract}, TipsFilterer: TipsFilterer{contract: contract}}, nil
}

// NewTipsCaller creates a new read-only instance of Tips, bound to a specific deployed contract.
func NewTipsCaller(address common.Address, caller bind.ContractCaller) (*TipsCaller, error) {
	contract, err := bindTips(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsCaller{contract: contract}, nil
}

// NewTipsTransactor creates a new write-only instance of Tips, bound to a specific deployed contract.
func NewTipsTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsTransactor, error) {
	contract, err := bindTips(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsTransactor{contract: contract}, nil
}

// NewTipsFilterer creates a new log filterer instance of Tips, bound to a specific deployed contract.
func NewTipsFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsFilterer, error) {
	contract, err := bindTips(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsFilterer{contract: contract}, nil
}

// bindTips binds a generic wrapper to an already deployed contract.
func bindTips(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tips *TipsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tips.Contract.TipsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tips *TipsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tips.Contract.TipsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tips *TipsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tips.Contract.TipsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tips *TipsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tips.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tips *TipsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tips.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tips *TipsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tips.Contract.contract.Transact(opts, method, params...)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b58fc32853740185ba7e18a8262de25da9210bf28611b32aa2e55ce748a48f6d64736f6c634300080d0033",
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
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea264697066735822122044a5bac384ad11d548d580063c9eea73e813613af34a179640e68b0b09b5822564736f6c634300080d0033",
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
	Bin: "0x6080604052348015600f57600080fd5b5060808061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063ffa1ad7414602d575b600080fd5b6034600081565b60405160ff909116815260200160405180910390f3fea26469706673582212200b8f27c2013655339fd256280a9828da11490670e24a1df2042c9558e2a6ffe764736f6c634300080d0033",
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
