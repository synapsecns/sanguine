// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testmessagebus

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

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122080a8feece8aebd5a29ba816e7ad476af307ac260e3acdf96fefde8ba65c780f564736f6c63430008000033",
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

// ContextChainIdUpgradeableMetaData contains all meta data concerning the ContextChainIdUpgradeable contract.
var ContextChainIdUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextChainIdUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextChainIdUpgradeableMetaData.ABI instead.
var ContextChainIdUpgradeableABI = ContextChainIdUpgradeableMetaData.ABI

// ContextChainIdUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextChainIdUpgradeable struct {
	ContextChainIdUpgradeableCaller     // Read-only binding to the contract
	ContextChainIdUpgradeableTransactor // Write-only binding to the contract
	ContextChainIdUpgradeableFilterer   // Log filterer for contract events
}

// ContextChainIdUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextChainIdUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextChainIdUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextChainIdUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextChainIdUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextChainIdUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextChainIdUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextChainIdUpgradeableSession struct {
	Contract     *ContextChainIdUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContextChainIdUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextChainIdUpgradeableCallerSession struct {
	Contract *ContextChainIdUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContextChainIdUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextChainIdUpgradeableTransactorSession struct {
	Contract     *ContextChainIdUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContextChainIdUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextChainIdUpgradeableRaw struct {
	Contract *ContextChainIdUpgradeable // Generic contract binding to access the raw methods on
}

// ContextChainIdUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextChainIdUpgradeableCallerRaw struct {
	Contract *ContextChainIdUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextChainIdUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextChainIdUpgradeableTransactorRaw struct {
	Contract *ContextChainIdUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextChainIdUpgradeable creates a new instance of ContextChainIdUpgradeable, bound to a specific deployed contract.
func NewContextChainIdUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextChainIdUpgradeable, error) {
	contract, err := bindContextChainIdUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextChainIdUpgradeable{ContextChainIdUpgradeableCaller: ContextChainIdUpgradeableCaller{contract: contract}, ContextChainIdUpgradeableTransactor: ContextChainIdUpgradeableTransactor{contract: contract}, ContextChainIdUpgradeableFilterer: ContextChainIdUpgradeableFilterer{contract: contract}}, nil
}

// NewContextChainIdUpgradeableCaller creates a new read-only instance of ContextChainIdUpgradeable, bound to a specific deployed contract.
func NewContextChainIdUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextChainIdUpgradeableCaller, error) {
	contract, err := bindContextChainIdUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextChainIdUpgradeableCaller{contract: contract}, nil
}

// NewContextChainIdUpgradeableTransactor creates a new write-only instance of ContextChainIdUpgradeable, bound to a specific deployed contract.
func NewContextChainIdUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextChainIdUpgradeableTransactor, error) {
	contract, err := bindContextChainIdUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextChainIdUpgradeableTransactor{contract: contract}, nil
}

// NewContextChainIdUpgradeableFilterer creates a new log filterer instance of ContextChainIdUpgradeable, bound to a specific deployed contract.
func NewContextChainIdUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextChainIdUpgradeableFilterer, error) {
	contract, err := bindContextChainIdUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextChainIdUpgradeableFilterer{contract: contract}, nil
}

// bindContextChainIdUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextChainIdUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextChainIdUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextChainIdUpgradeable *ContextChainIdUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextChainIdUpgradeable.Contract.ContextChainIdUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextChainIdUpgradeable *ContextChainIdUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextChainIdUpgradeable.Contract.ContextChainIdUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextChainIdUpgradeable *ContextChainIdUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextChainIdUpgradeable.Contract.ContextChainIdUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextChainIdUpgradeable *ContextChainIdUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextChainIdUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextChainIdUpgradeable *ContextChainIdUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextChainIdUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextChainIdUpgradeable *ContextChainIdUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextChainIdUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
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

// IAuthVerifierMetaData contains all meta data concerning the IAuthVerifier contract.
var IAuthVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_authData\",\"type\":\"bytes\"}],\"name\":\"msgAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"authenticated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodegroup\",\"type\":\"address\"}],\"name\":\"setNodeGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8b1b3a2d": "msgAuth(bytes)",
		"f6ea2c90": "setNodeGroup(address)",
	},
}

// IAuthVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IAuthVerifierMetaData.ABI instead.
var IAuthVerifierABI = IAuthVerifierMetaData.ABI

// Deprecated: Use IAuthVerifierMetaData.Sigs instead.
// IAuthVerifierFuncSigs maps the 4-byte function signature to its string representation.
var IAuthVerifierFuncSigs = IAuthVerifierMetaData.Sigs

// IAuthVerifier is an auto generated Go binding around an Ethereum contract.
type IAuthVerifier struct {
	IAuthVerifierCaller     // Read-only binding to the contract
	IAuthVerifierTransactor // Write-only binding to the contract
	IAuthVerifierFilterer   // Log filterer for contract events
}

// IAuthVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuthVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuthVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuthVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuthVerifierSession struct {
	Contract     *IAuthVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuthVerifierCallerSession struct {
	Contract *IAuthVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAuthVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuthVerifierTransactorSession struct {
	Contract     *IAuthVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAuthVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuthVerifierRaw struct {
	Contract *IAuthVerifier // Generic contract binding to access the raw methods on
}

// IAuthVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuthVerifierCallerRaw struct {
	Contract *IAuthVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuthVerifierTransactorRaw struct {
	Contract *IAuthVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuthVerifier creates a new instance of IAuthVerifier, bound to a specific deployed contract.
func NewIAuthVerifier(address common.Address, backend bind.ContractBackend) (*IAuthVerifier, error) {
	contract, err := bindIAuthVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuthVerifier{IAuthVerifierCaller: IAuthVerifierCaller{contract: contract}, IAuthVerifierTransactor: IAuthVerifierTransactor{contract: contract}, IAuthVerifierFilterer: IAuthVerifierFilterer{contract: contract}}, nil
}

// NewIAuthVerifierCaller creates a new read-only instance of IAuthVerifier, bound to a specific deployed contract.
func NewIAuthVerifierCaller(address common.Address, caller bind.ContractCaller) (*IAuthVerifierCaller, error) {
	contract, err := bindIAuthVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthVerifierCaller{contract: contract}, nil
}

// NewIAuthVerifierTransactor creates a new write-only instance of IAuthVerifier, bound to a specific deployed contract.
func NewIAuthVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthVerifierTransactor, error) {
	contract, err := bindIAuthVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthVerifierTransactor{contract: contract}, nil
}

// NewIAuthVerifierFilterer creates a new log filterer instance of IAuthVerifier, bound to a specific deployed contract.
func NewIAuthVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthVerifierFilterer, error) {
	contract, err := bindIAuthVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthVerifierFilterer{contract: contract}, nil
}

// bindIAuthVerifier binds a generic wrapper to an already deployed contract.
func bindIAuthVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuthVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuthVerifier *IAuthVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuthVerifier.Contract.IAuthVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuthVerifier *IAuthVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuthVerifier.Contract.IAuthVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuthVerifier *IAuthVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuthVerifier.Contract.IAuthVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuthVerifier *IAuthVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuthVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuthVerifier *IAuthVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuthVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuthVerifier *IAuthVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuthVerifier.Contract.contract.Transact(opts, method, params...)
}

// MsgAuth is a free data retrieval call binding the contract method 0x8b1b3a2d.
//
// Solidity: function msgAuth(bytes _authData) view returns(bool authenticated)
func (_IAuthVerifier *IAuthVerifierCaller) MsgAuth(opts *bind.CallOpts, _authData []byte) (bool, error) {
	var out []interface{}
	err := _IAuthVerifier.contract.Call(opts, &out, "msgAuth", _authData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MsgAuth is a free data retrieval call binding the contract method 0x8b1b3a2d.
//
// Solidity: function msgAuth(bytes _authData) view returns(bool authenticated)
func (_IAuthVerifier *IAuthVerifierSession) MsgAuth(_authData []byte) (bool, error) {
	return _IAuthVerifier.Contract.MsgAuth(&_IAuthVerifier.CallOpts, _authData)
}

// MsgAuth is a free data retrieval call binding the contract method 0x8b1b3a2d.
//
// Solidity: function msgAuth(bytes _authData) view returns(bool authenticated)
func (_IAuthVerifier *IAuthVerifierCallerSession) MsgAuth(_authData []byte) (bool, error) {
	return _IAuthVerifier.Contract.MsgAuth(&_IAuthVerifier.CallOpts, _authData)
}

// SetNodeGroup is a paid mutator transaction binding the contract method 0xf6ea2c90.
//
// Solidity: function setNodeGroup(address _nodegroup) returns()
func (_IAuthVerifier *IAuthVerifierTransactor) SetNodeGroup(opts *bind.TransactOpts, _nodegroup common.Address) (*types.Transaction, error) {
	return _IAuthVerifier.contract.Transact(opts, "setNodeGroup", _nodegroup)
}

// SetNodeGroup is a paid mutator transaction binding the contract method 0xf6ea2c90.
//
// Solidity: function setNodeGroup(address _nodegroup) returns()
func (_IAuthVerifier *IAuthVerifierSession) SetNodeGroup(_nodegroup common.Address) (*types.Transaction, error) {
	return _IAuthVerifier.Contract.SetNodeGroup(&_IAuthVerifier.TransactOpts, _nodegroup)
}

// SetNodeGroup is a paid mutator transaction binding the contract method 0xf6ea2c90.
//
// Solidity: function setNodeGroup(address _nodegroup) returns()
func (_IAuthVerifier *IAuthVerifierTransactorSession) SetNodeGroup(_nodegroup common.Address) (*types.Transaction, error) {
	return _IAuthVerifier.Contract.SetNodeGroup(&_IAuthVerifier.TransactOpts, _nodegroup)
}

// IGasFeePricingMetaData contains all meta data concerning the IGasFeePricing contract.
var IGasFeePricingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"estimateGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasUnitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasTokenPriceRatio\",\"type\":\"uint256\"}],\"name\":\"setCostPerChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"47feadc1": "estimateGasFee(uint256,bytes)",
		"e32192b7": "setCostPerChain(uint256,uint256,uint256)",
	},
}

// IGasFeePricingABI is the input ABI used to generate the binding from.
// Deprecated: Use IGasFeePricingMetaData.ABI instead.
var IGasFeePricingABI = IGasFeePricingMetaData.ABI

// Deprecated: Use IGasFeePricingMetaData.Sigs instead.
// IGasFeePricingFuncSigs maps the 4-byte function signature to its string representation.
var IGasFeePricingFuncSigs = IGasFeePricingMetaData.Sigs

// IGasFeePricing is an auto generated Go binding around an Ethereum contract.
type IGasFeePricing struct {
	IGasFeePricingCaller     // Read-only binding to the contract
	IGasFeePricingTransactor // Write-only binding to the contract
	IGasFeePricingFilterer   // Log filterer for contract events
}

// IGasFeePricingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGasFeePricingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasFeePricingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGasFeePricingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasFeePricingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGasFeePricingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasFeePricingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGasFeePricingSession struct {
	Contract     *IGasFeePricing   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGasFeePricingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGasFeePricingCallerSession struct {
	Contract *IGasFeePricingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IGasFeePricingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGasFeePricingTransactorSession struct {
	Contract     *IGasFeePricingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IGasFeePricingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGasFeePricingRaw struct {
	Contract *IGasFeePricing // Generic contract binding to access the raw methods on
}

// IGasFeePricingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGasFeePricingCallerRaw struct {
	Contract *IGasFeePricingCaller // Generic read-only contract binding to access the raw methods on
}

// IGasFeePricingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGasFeePricingTransactorRaw struct {
	Contract *IGasFeePricingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGasFeePricing creates a new instance of IGasFeePricing, bound to a specific deployed contract.
func NewIGasFeePricing(address common.Address, backend bind.ContractBackend) (*IGasFeePricing, error) {
	contract, err := bindIGasFeePricing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGasFeePricing{IGasFeePricingCaller: IGasFeePricingCaller{contract: contract}, IGasFeePricingTransactor: IGasFeePricingTransactor{contract: contract}, IGasFeePricingFilterer: IGasFeePricingFilterer{contract: contract}}, nil
}

// NewIGasFeePricingCaller creates a new read-only instance of IGasFeePricing, bound to a specific deployed contract.
func NewIGasFeePricingCaller(address common.Address, caller bind.ContractCaller) (*IGasFeePricingCaller, error) {
	contract, err := bindIGasFeePricing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGasFeePricingCaller{contract: contract}, nil
}

// NewIGasFeePricingTransactor creates a new write-only instance of IGasFeePricing, bound to a specific deployed contract.
func NewIGasFeePricingTransactor(address common.Address, transactor bind.ContractTransactor) (*IGasFeePricingTransactor, error) {
	contract, err := bindIGasFeePricing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGasFeePricingTransactor{contract: contract}, nil
}

// NewIGasFeePricingFilterer creates a new log filterer instance of IGasFeePricing, bound to a specific deployed contract.
func NewIGasFeePricingFilterer(address common.Address, filterer bind.ContractFilterer) (*IGasFeePricingFilterer, error) {
	contract, err := bindIGasFeePricing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGasFeePricingFilterer{contract: contract}, nil
}

// bindIGasFeePricing binds a generic wrapper to an already deployed contract.
func bindIGasFeePricing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGasFeePricingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasFeePricing *IGasFeePricingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasFeePricing.Contract.IGasFeePricingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasFeePricing *IGasFeePricingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasFeePricing.Contract.IGasFeePricingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasFeePricing *IGasFeePricingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasFeePricing.Contract.IGasFeePricingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasFeePricing *IGasFeePricingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasFeePricing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasFeePricing *IGasFeePricingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasFeePricing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasFeePricing *IGasFeePricingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasFeePricing.Contract.contract.Transact(opts, method, params...)
}

// EstimateGasFee is a paid mutator transaction binding the contract method 0x47feadc1.
//
// Solidity: function estimateGasFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IGasFeePricing *IGasFeePricingTransactor) EstimateGasFee(opts *bind.TransactOpts, _dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IGasFeePricing.contract.Transact(opts, "estimateGasFee", _dstChainId, _options)
}

// EstimateGasFee is a paid mutator transaction binding the contract method 0x47feadc1.
//
// Solidity: function estimateGasFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IGasFeePricing *IGasFeePricingSession) EstimateGasFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IGasFeePricing.Contract.EstimateGasFee(&_IGasFeePricing.TransactOpts, _dstChainId, _options)
}

// EstimateGasFee is a paid mutator transaction binding the contract method 0x47feadc1.
//
// Solidity: function estimateGasFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IGasFeePricing *IGasFeePricingTransactorSession) EstimateGasFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IGasFeePricing.Contract.EstimateGasFee(&_IGasFeePricing.TransactOpts, _dstChainId, _options)
}

// SetCostPerChain is a paid mutator transaction binding the contract method 0xe32192b7.
//
// Solidity: function setCostPerChain(uint256 _dstChainId, uint256 _gasUnitPrice, uint256 _gasTokenPriceRatio) returns()
func (_IGasFeePricing *IGasFeePricingTransactor) SetCostPerChain(opts *bind.TransactOpts, _dstChainId *big.Int, _gasUnitPrice *big.Int, _gasTokenPriceRatio *big.Int) (*types.Transaction, error) {
	return _IGasFeePricing.contract.Transact(opts, "setCostPerChain", _dstChainId, _gasUnitPrice, _gasTokenPriceRatio)
}

// SetCostPerChain is a paid mutator transaction binding the contract method 0xe32192b7.
//
// Solidity: function setCostPerChain(uint256 _dstChainId, uint256 _gasUnitPrice, uint256 _gasTokenPriceRatio) returns()
func (_IGasFeePricing *IGasFeePricingSession) SetCostPerChain(_dstChainId *big.Int, _gasUnitPrice *big.Int, _gasTokenPriceRatio *big.Int) (*types.Transaction, error) {
	return _IGasFeePricing.Contract.SetCostPerChain(&_IGasFeePricing.TransactOpts, _dstChainId, _gasUnitPrice, _gasTokenPriceRatio)
}

// SetCostPerChain is a paid mutator transaction binding the contract method 0xe32192b7.
//
// Solidity: function setCostPerChain(uint256 _dstChainId, uint256 _gasUnitPrice, uint256 _gasTokenPriceRatio) returns()
func (_IGasFeePricing *IGasFeePricingTransactorSession) SetCostPerChain(_dstChainId *big.Int, _gasUnitPrice *big.Int, _gasTokenPriceRatio *big.Int) (*types.Transaction, error) {
	return _IGasFeePricing.Contract.SetCostPerChain(&_IGasFeePricing.TransactOpts, _dstChainId, _gasUnitPrice, _gasTokenPriceRatio)
}

// ISynMessagingReceiverMetaData contains all meta data concerning the ISynMessagingReceiver contract.
var ISynMessagingReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a6060871": "executeMessage(bytes32,uint256,bytes,address)",
	},
}

// ISynMessagingReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynMessagingReceiverMetaData.ABI instead.
var ISynMessagingReceiverABI = ISynMessagingReceiverMetaData.ABI

// Deprecated: Use ISynMessagingReceiverMetaData.Sigs instead.
// ISynMessagingReceiverFuncSigs maps the 4-byte function signature to its string representation.
var ISynMessagingReceiverFuncSigs = ISynMessagingReceiverMetaData.Sigs

// ISynMessagingReceiver is an auto generated Go binding around an Ethereum contract.
type ISynMessagingReceiver struct {
	ISynMessagingReceiverCaller     // Read-only binding to the contract
	ISynMessagingReceiverTransactor // Write-only binding to the contract
	ISynMessagingReceiverFilterer   // Log filterer for contract events
}

// ISynMessagingReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynMessagingReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynMessagingReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynMessagingReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynMessagingReceiverSession struct {
	Contract     *ISynMessagingReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ISynMessagingReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynMessagingReceiverCallerSession struct {
	Contract *ISynMessagingReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ISynMessagingReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynMessagingReceiverTransactorSession struct {
	Contract     *ISynMessagingReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ISynMessagingReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynMessagingReceiverRaw struct {
	Contract *ISynMessagingReceiver // Generic contract binding to access the raw methods on
}

// ISynMessagingReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynMessagingReceiverCallerRaw struct {
	Contract *ISynMessagingReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ISynMessagingReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynMessagingReceiverTransactorRaw struct {
	Contract *ISynMessagingReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynMessagingReceiver creates a new instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiver(address common.Address, backend bind.ContractBackend) (*ISynMessagingReceiver, error) {
	contract, err := bindISynMessagingReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiver{ISynMessagingReceiverCaller: ISynMessagingReceiverCaller{contract: contract}, ISynMessagingReceiverTransactor: ISynMessagingReceiverTransactor{contract: contract}, ISynMessagingReceiverFilterer: ISynMessagingReceiverFilterer{contract: contract}}, nil
}

// NewISynMessagingReceiverCaller creates a new read-only instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverCaller(address common.Address, caller bind.ContractCaller) (*ISynMessagingReceiverCaller, error) {
	contract, err := bindISynMessagingReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverCaller{contract: contract}, nil
}

// NewISynMessagingReceiverTransactor creates a new write-only instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynMessagingReceiverTransactor, error) {
	contract, err := bindISynMessagingReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverTransactor{contract: contract}, nil
}

// NewISynMessagingReceiverFilterer creates a new log filterer instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynMessagingReceiverFilterer, error) {
	contract, err := bindISynMessagingReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverFilterer{contract: contract}, nil
}

// bindISynMessagingReceiver binds a generic wrapper to an already deployed contract.
func bindISynMessagingReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISynMessagingReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynMessagingReceiver *ISynMessagingReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynMessagingReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.contract.Transact(opts, "executeMessage", _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ExecuteMessage(&_ISynMessagingReceiver.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ExecuteMessage(&_ISynMessagingReceiver.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[]",
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

// MessageBusReceiverUpgradeableMetaData contains all meta data concerning the MessageBusReceiverUpgradeable contract.
var MessageBusReceiverUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CallReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcNonce\",\"type\":\"uint64\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"getExecutedMessage\",\"outputs\":[{\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authVerifier\",\"type\":\"address\"}],\"name\":\"updateAuthVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"},{\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"updateMessageStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c4087335": "authVerifier()",
		"a1b058d8": "executeMessage(uint256,bytes32,address,uint256,uint256,bytes,bytes32)",
		"25b19fa3": "getExecutedMessage(bytes32)",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"a5c0edf3": "updateAuthVerifier(address)",
		"9b11079c": "updateMessageStatus(bytes32,uint8)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610e73806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80639b11079c11610076578063a5c0edf31161005b578063a5c0edf31461012b578063c40873351461013e578063f2fde38b14610146576100a3565b80639b11079c14610105578063a1b058d814610118576100a3565b806325b19fa3146100a85780635c975abb146100d1578063715018a6146100e65780638da5cb5b146100f0575b600080fd5b6100bb6100b6366004610972565b610159565b6040516100c89190610c69565b60405180910390f35b6100d9610171565b6040516100c89190610bd1565b6100ee61017a565b005b6100f8610202565b6040516100c89190610bb0565b6100ee61011336600461098a565b61021e565b6100ee610126366004610a7d565b61030a565b6100ee610139366004610931565b61063d565b6100f8610744565b6100ee610154366004610931565b610760565b60008181526098602052604090205460ff165b919050565b60655460ff1690565b61018261082c565b73ffffffffffffffffffffffffffffffffffffffff166101a0610202565b73ffffffffffffffffffffffffffffffffffffffff16146101f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90610d72565b60405180910390fd5b6102006000610830565b565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b61022661082c565b73ffffffffffffffffffffffffffffffffffffffff16610244610202565b73ffffffffffffffffffffffffffffffffffffffff1614610291576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90610d72565b600082815260986020526040902080548291907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836002811115610301577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055505050565b610312610171565b15610349576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90610d3b565b60008181526098602052604081205460ff166002811115610393577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b146103ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90610da7565b60975460405173ffffffffffffffffffffffffffffffffffffffff90911690638b1b3a2d906103fd903390602001610bb0565b6040516020818303038152906040526040518263ffffffff1660e01b81526004016104289190610c56565b60206040518083038186803b15801561044057600080fd5b505afa158015610454573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104789190610952565b5060008673ffffffffffffffffffffffffffffffffffffffff1663a6060871878a8c8888336040518763ffffffff1660e01b81526004016104bd959493929190610bdc565b600060405180830381600088803b1580156104d757600080fd5b5087f1935050505080156104e9575060015b610566573d808015610517576040519150601f19603f3d011682016040523d82523d6000602084013e61051c565b606091505b507fffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f610547826108a7565b6040516105549190610c56565b60405180910390a1600291505061056a565b5060015b600082815260986020526040902080548291907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018360028111156105da577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055508673ffffffffffffffffffffffffffffffffffffffff16827f04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65838c8960405161062a93929190610c77565b60405180910390a3505050505050505050565b61064561082c565b73ffffffffffffffffffffffffffffffffffffffff16610663610202565b73ffffffffffffffffffffffffffffffffffffffff16146106b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90610d72565b73ffffffffffffffffffffffffffffffffffffffff81166106fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90610d04565b609780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60975473ffffffffffffffffffffffffffffffffffffffff1681565b61076861082c565b73ffffffffffffffffffffffffffffffffffffffff16610786610202565b73ffffffffffffffffffffffffffffffffffffffff16146107d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90610d72565b73ffffffffffffffffffffffffffffffffffffffff8116610820576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90610ca7565b61082981610830565b50565b3390565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606044825110156108ed575060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c79000000602082015261016c565b6004820191508180602001905181019061090791906109bc565b92915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461016c57600080fd5b600060208284031215610942578081fd5b61094b8261090d565b9392505050565b600060208284031215610963578081fd5b8151801515811461094b578182fd5b600060208284031215610983578081fd5b5035919050565b6000806040838503121561099c578081fd5b823591506020830135600381106109b1578182fd5b809150509250929050565b6000602082840312156109cd578081fd5b815167ffffffffffffffff808211156109e4578283fd5b818401915084601f8301126109f7578283fd5b815181811115610a0957610a09610e0e565b60405160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401168201018181108482111715610a4b57610a4b610e0e565b604052818152838201602001871015610a62578485fd5b610a73826020830160208701610dde565b9695505050505050565b60008060008060008060008060e0898b031215610a98578384fd5b8835975060208901359650610aaf60408a0161090d565b9550606089013594506080890135935060a089013567ffffffffffffffff80821115610ad9578485fd5b818b0191508b601f830112610aec578485fd5b813581811115610afa578586fd5b8c6020828501011115610b0b578586fd5b60208301955080945050505060c089013590509295985092959890939650565b60008151808452610b43816020860160208601610dde565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60038110610bac577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b901515815260200190565b600086825285602083015260806040830152836080830152838560a08401378060a0858401015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116830101905073ffffffffffffffffffffffffffffffffffffffff831660608301529695505050505050565b60006020825261094b6020830184610b2b565b602081016109078284610b75565b60608101610c858286610b75565b67ffffffffffffffff8085166020840152808416604084015250949350505050565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201527f6464726573730000000000000000000000000000000000000000000000000000606082015260800190565b6020808252600f908201527f43616e6e6f742073657420746f20300000000000000000000000000000000000604082015260600190565b60208082526010908201527f5061757361626c653a2070617573656400000000000000000000000000000000604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526018908201527f4d65737361676520616c72656164792065786563757465640000000000000000604082015260600190565b60005b83811015610df9578181015183820152602001610de1565b83811115610e08576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea2646970667358221220cee79873fb2a09b77d4bc866527128d03b9dfc2cf18612a3e711d9dd54fce98464736f6c63430008000033",
}

// MessageBusReceiverUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusReceiverUpgradeableMetaData.ABI instead.
var MessageBusReceiverUpgradeableABI = MessageBusReceiverUpgradeableMetaData.ABI

// Deprecated: Use MessageBusReceiverUpgradeableMetaData.Sigs instead.
// MessageBusReceiverUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var MessageBusReceiverUpgradeableFuncSigs = MessageBusReceiverUpgradeableMetaData.Sigs

// MessageBusReceiverUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusReceiverUpgradeableMetaData.Bin instead.
var MessageBusReceiverUpgradeableBin = MessageBusReceiverUpgradeableMetaData.Bin

// DeployMessageBusReceiverUpgradeable deploys a new Ethereum contract, binding an instance of MessageBusReceiverUpgradeable to it.
func DeployMessageBusReceiverUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageBusReceiverUpgradeable, error) {
	parsed, err := MessageBusReceiverUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusReceiverUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusReceiverUpgradeable{MessageBusReceiverUpgradeableCaller: MessageBusReceiverUpgradeableCaller{contract: contract}, MessageBusReceiverUpgradeableTransactor: MessageBusReceiverUpgradeableTransactor{contract: contract}, MessageBusReceiverUpgradeableFilterer: MessageBusReceiverUpgradeableFilterer{contract: contract}}, nil
}

// MessageBusReceiverUpgradeable is an auto generated Go binding around an Ethereum contract.
type MessageBusReceiverUpgradeable struct {
	MessageBusReceiverUpgradeableCaller     // Read-only binding to the contract
	MessageBusReceiverUpgradeableTransactor // Write-only binding to the contract
	MessageBusReceiverUpgradeableFilterer   // Log filterer for contract events
}

// MessageBusReceiverUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusReceiverUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusReceiverUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusReceiverUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusReceiverUpgradeableSession struct {
	Contract     *MessageBusReceiverUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MessageBusReceiverUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusReceiverUpgradeableCallerSession struct {
	Contract *MessageBusReceiverUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// MessageBusReceiverUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusReceiverUpgradeableTransactorSession struct {
	Contract     *MessageBusReceiverUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// MessageBusReceiverUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusReceiverUpgradeableRaw struct {
	Contract *MessageBusReceiverUpgradeable // Generic contract binding to access the raw methods on
}

// MessageBusReceiverUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusReceiverUpgradeableCallerRaw struct {
	Contract *MessageBusReceiverUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusReceiverUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusReceiverUpgradeableTransactorRaw struct {
	Contract *MessageBusReceiverUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusReceiverUpgradeable creates a new instance of MessageBusReceiverUpgradeable, bound to a specific deployed contract.
func NewMessageBusReceiverUpgradeable(address common.Address, backend bind.ContractBackend) (*MessageBusReceiverUpgradeable, error) {
	contract, err := bindMessageBusReceiverUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeable{MessageBusReceiverUpgradeableCaller: MessageBusReceiverUpgradeableCaller{contract: contract}, MessageBusReceiverUpgradeableTransactor: MessageBusReceiverUpgradeableTransactor{contract: contract}, MessageBusReceiverUpgradeableFilterer: MessageBusReceiverUpgradeableFilterer{contract: contract}}, nil
}

// NewMessageBusReceiverUpgradeableCaller creates a new read-only instance of MessageBusReceiverUpgradeable, bound to a specific deployed contract.
func NewMessageBusReceiverUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*MessageBusReceiverUpgradeableCaller, error) {
	contract, err := bindMessageBusReceiverUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeableCaller{contract: contract}, nil
}

// NewMessageBusReceiverUpgradeableTransactor creates a new write-only instance of MessageBusReceiverUpgradeable, bound to a specific deployed contract.
func NewMessageBusReceiverUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusReceiverUpgradeableTransactor, error) {
	contract, err := bindMessageBusReceiverUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeableTransactor{contract: contract}, nil
}

// NewMessageBusReceiverUpgradeableFilterer creates a new log filterer instance of MessageBusReceiverUpgradeable, bound to a specific deployed contract.
func NewMessageBusReceiverUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusReceiverUpgradeableFilterer, error) {
	contract, err := bindMessageBusReceiverUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeableFilterer{contract: contract}, nil
}

// bindMessageBusReceiverUpgradeable binds a generic wrapper to an already deployed contract.
func bindMessageBusReceiverUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusReceiverUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusReceiverUpgradeable.Contract.MessageBusReceiverUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.MessageBusReceiverUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.MessageBusReceiverUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusReceiverUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCaller) AuthVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiverUpgradeable.contract.Call(opts, &out, "authVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) AuthVerifier() (common.Address, error) {
	return _MessageBusReceiverUpgradeable.Contract.AuthVerifier(&_MessageBusReceiverUpgradeable.CallOpts)
}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCallerSession) AuthVerifier() (common.Address, error) {
	return _MessageBusReceiverUpgradeable.Contract.AuthVerifier(&_MessageBusReceiverUpgradeable.CallOpts)
}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCaller) GetExecutedMessage(opts *bind.CallOpts, _messageId [32]byte) (uint8, error) {
	var out []interface{}
	err := _MessageBusReceiverUpgradeable.contract.Call(opts, &out, "getExecutedMessage", _messageId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) GetExecutedMessage(_messageId [32]byte) (uint8, error) {
	return _MessageBusReceiverUpgradeable.Contract.GetExecutedMessage(&_MessageBusReceiverUpgradeable.CallOpts, _messageId)
}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCallerSession) GetExecutedMessage(_messageId [32]byte) (uint8, error) {
	return _MessageBusReceiverUpgradeable.Contract.GetExecutedMessage(&_MessageBusReceiverUpgradeable.CallOpts, _messageId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiverUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) Owner() (common.Address, error) {
	return _MessageBusReceiverUpgradeable.Contract.Owner(&_MessageBusReceiverUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCallerSession) Owner() (common.Address, error) {
	return _MessageBusReceiverUpgradeable.Contract.Owner(&_MessageBusReceiverUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MessageBusReceiverUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) Paused() (bool, error) {
	return _MessageBusReceiverUpgradeable.Contract.Paused(&_MessageBusReceiverUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableCallerSession) Paused() (bool, error) {
	return _MessageBusReceiverUpgradeable.Contract.Paused(&_MessageBusReceiverUpgradeable.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.contract.Transact(opts, "executeMessage", _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.ExecuteMessage(&_MessageBusReceiverUpgradeable.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactorSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.ExecuteMessage(&_MessageBusReceiverUpgradeable.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.RenounceOwnership(&_MessageBusReceiverUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.RenounceOwnership(&_MessageBusReceiverUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.TransferOwnership(&_MessageBusReceiverUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.TransferOwnership(&_MessageBusReceiverUpgradeable.TransactOpts, newOwner)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactor) UpdateAuthVerifier(opts *bind.TransactOpts, _authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.contract.Transact(opts, "updateAuthVerifier", _authVerifier)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) UpdateAuthVerifier(_authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.UpdateAuthVerifier(&_MessageBusReceiverUpgradeable.TransactOpts, _authVerifier)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactorSession) UpdateAuthVerifier(_authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.UpdateAuthVerifier(&_MessageBusReceiverUpgradeable.TransactOpts, _authVerifier)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactor) UpdateMessageStatus(opts *bind.TransactOpts, _messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.contract.Transact(opts, "updateMessageStatus", _messageId, _status)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableSession) UpdateMessageStatus(_messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.UpdateMessageStatus(&_MessageBusReceiverUpgradeable.TransactOpts, _messageId, _status)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableTransactorSession) UpdateMessageStatus(_messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _MessageBusReceiverUpgradeable.Contract.UpdateMessageStatus(&_MessageBusReceiverUpgradeable.TransactOpts, _messageId, _status)
}

// MessageBusReceiverUpgradeableCallRevertedIterator is returned from FilterCallReverted and is used to iterate over the raw logs and unpacked data for CallReverted events raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeableCallRevertedIterator struct {
	Event *MessageBusReceiverUpgradeableCallReverted // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverUpgradeableCallRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverUpgradeableCallReverted)
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
		it.Event = new(MessageBusReceiverUpgradeableCallReverted)
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
func (it *MessageBusReceiverUpgradeableCallRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverUpgradeableCallRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverUpgradeableCallReverted represents a CallReverted event raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeableCallReverted struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallReverted is a free log retrieval operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) FilterCallReverted(opts *bind.FilterOpts) (*MessageBusReceiverUpgradeableCallRevertedIterator, error) {

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.FilterLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeableCallRevertedIterator{contract: _MessageBusReceiverUpgradeable.contract, event: "CallReverted", logs: logs, sub: sub}, nil
}

// WatchCallReverted is a free log subscription operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) WatchCallReverted(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverUpgradeableCallReverted) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.WatchLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverUpgradeableCallReverted)
				if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "CallReverted", log); err != nil {
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

// ParseCallReverted is a log parse operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) ParseCallReverted(log types.Log) (*MessageBusReceiverUpgradeableCallReverted, error) {
	event := new(MessageBusReceiverUpgradeableCallReverted)
	if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "CallReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverUpgradeableExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeableExecutedIterator struct {
	Event *MessageBusReceiverUpgradeableExecuted // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverUpgradeableExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverUpgradeableExecuted)
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
		it.Event = new(MessageBusReceiverUpgradeableExecuted)
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
func (it *MessageBusReceiverUpgradeableExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverUpgradeableExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverUpgradeableExecuted represents a Executed event raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeableExecuted struct {
	MessageId  [32]byte
	Status     uint8
	DstAddress common.Address
	SrcChainId uint64
	SrcNonce   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) FilterExecuted(opts *bind.FilterOpts, messageId [][32]byte, _dstAddress []common.Address) (*MessageBusReceiverUpgradeableExecutedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	var _dstAddressRule []interface{}
	for _, _dstAddressItem := range _dstAddress {
		_dstAddressRule = append(_dstAddressRule, _dstAddressItem)
	}

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.FilterLogs(opts, "Executed", messageIdRule, _dstAddressRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeableExecutedIterator{contract: _MessageBusReceiverUpgradeable.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverUpgradeableExecuted, messageId [][32]byte, _dstAddress []common.Address) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	var _dstAddressRule []interface{}
	for _, _dstAddressItem := range _dstAddress {
		_dstAddressRule = append(_dstAddressRule, _dstAddressItem)
	}

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.WatchLogs(opts, "Executed", messageIdRule, _dstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverUpgradeableExecuted)
				if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) ParseExecuted(log types.Log) (*MessageBusReceiverUpgradeableExecuted, error) {
	event := new(MessageBusReceiverUpgradeableExecuted)
	if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeableOwnershipTransferredIterator struct {
	Event *MessageBusReceiverUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverUpgradeableOwnershipTransferred)
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
		it.Event = new(MessageBusReceiverUpgradeableOwnershipTransferred)
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
func (it *MessageBusReceiverUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusReceiverUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeableOwnershipTransferredIterator{contract: _MessageBusReceiverUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverUpgradeableOwnershipTransferred)
				if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusReceiverUpgradeableOwnershipTransferred, error) {
	event := new(MessageBusReceiverUpgradeableOwnershipTransferred)
	if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeablePausedIterator struct {
	Event *MessageBusReceiverUpgradeablePaused // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverUpgradeablePaused)
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
		it.Event = new(MessageBusReceiverUpgradeablePaused)
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
func (it *MessageBusReceiverUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverUpgradeablePaused represents a Paused event raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*MessageBusReceiverUpgradeablePausedIterator, error) {

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeablePausedIterator{contract: _MessageBusReceiverUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverUpgradeablePaused)
				if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) ParsePaused(log types.Log) (*MessageBusReceiverUpgradeablePaused, error) {
	event := new(MessageBusReceiverUpgradeablePaused)
	if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeableUnpausedIterator struct {
	Event *MessageBusReceiverUpgradeableUnpaused // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverUpgradeableUnpaused)
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
		it.Event = new(MessageBusReceiverUpgradeableUnpaused)
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
func (it *MessageBusReceiverUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverUpgradeableUnpaused represents a Unpaused event raised by the MessageBusReceiverUpgradeable contract.
type MessageBusReceiverUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MessageBusReceiverUpgradeableUnpausedIterator, error) {

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverUpgradeableUnpausedIterator{contract: _MessageBusReceiverUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiverUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverUpgradeableUnpaused)
				if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MessageBusReceiverUpgradeable *MessageBusReceiverUpgradeableFilterer) ParseUnpaused(log types.Log) (*MessageBusReceiverUpgradeableUnpaused, error) {
	event := new(MessageBusReceiverUpgradeableUnpaused)
	if err := _MessageBusReceiverUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderUpgradeableMetaData contains all meta data concerning the MessageBusSenderUpgradeable contract.
var MessageBusSenderUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_srcAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dstAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_srcNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"computeMessageId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasFeePricing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"rescueGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasFeePricing\",\"type\":\"address\"}],\"name\":\"updateGasFeePricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawGasFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f44d57aa": "computeMessageId(address,uint256,bytes32,uint256,uint256,bytes)",
		"5da6d2c4": "estimateFee(uint256,bytes)",
		"9af1d35a": "fees()",
		"aa70fc0e": "gasFeePricing()",
		"affed0e0": "nonce()",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"205e157b": "rescueGas(address)",
		"ac8a4c1b": "sendMessage(bytes32,uint256,bytes,bytes)",
		"72177189": "sendMessage(bytes32,uint256,bytes,bytes,address)",
		"f2fde38b": "transferOwnership(address)",
		"a66dd384": "updateGasFeePricing(address)",
		"d6b457b9": "withdrawGasFees(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611149806100206000396000f3fe6080604052600436106100dd5760003560e01c8063a66dd3841161007f578063affed0e011610059578063affed0e014610203578063d6b457b914610225578063f2fde38b14610245578063f44d57aa14610265576100dd565b8063a66dd384146101bb578063aa70fc0e146101db578063ac8a4c1b146101f0576100dd565b8063715018a6116100bb578063715018a61461015c57806372177189146101715780638da5cb5b146101845780639af1d35a146101a6576100dd565b8063205e157b146100e25780635c975abb146101045780635da6d2c41461012f575b600080fd5b3480156100ee57600080fd5b506101026100fd366004610b4f565b610285565b005b34801561011057600080fd5b5061011961035c565b6040516101269190610e21565b60405180910390f35b34801561013b57600080fd5b5061014f61014a366004610d19565b610365565b6040516101269190610e2c565b34801561016857600080fd5b50610102610457565b61010261017f366004610c6d565b6104d6565b34801561019057600080fd5b506101996104ee565b6040516101269190610dab565b3480156101b257600080fd5b5061014f61050a565b3480156101c757600080fd5b506101026101d6366004610b4f565b610510565b3480156101e757600080fd5b50610199610617565b6101026101fe366004610bed565b610633565b34801561020f57600080fd5b50610218610689565b6040516101269190611056565b34801561023157600080fd5b50610102610240366004610b4f565b6106b1565b34801561025157600080fd5b50610102610260366004610b4f565b610774565b34801561027157600080fd5b5061014f610280366004610b72565b610840565b61028d610882565b73ffffffffffffffffffffffffffffffffffffffff166102ab6104ee565b73ffffffffffffffffffffffffffffffffffffffff1614610301576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610f6e565b60405180910390fd5b600060ca54476103119190611083565b60405190915073ffffffffffffffffffffffffffffffffffffffff83169082156108fc029083906000818181858888f19350505050158015610357573d6000803e3d6000fd5b505050565b60655460ff1690565b60c9546040517f47feadc1000000000000000000000000000000000000000000000000000000008152600091829173ffffffffffffffffffffffffffffffffffffffff909116906347feadc1906103c490889088908890600401611033565b602060405180830381600087803b1580156103de57600080fd5b505af11580156103f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104169190610d01565b90508061044f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610e6c565b949350505050565b61045f610882565b73ffffffffffffffffffffffffffffffffffffffff1661047d6104ee565b73ffffffffffffffffffffffffffffffffffffffff16146104ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610f6e565b6104d46000610886565b565b6104e5878787878787876108fd565b50505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b60ca5481565b610518610882565b73ffffffffffffffffffffffffffffffffffffffff166105366104ee565b73ffffffffffffffffffffffffffffffffffffffff1614610583576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610f6e565b73ffffffffffffffffffffffffffffffffffffffff81166105d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610f00565b60c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60c95473ffffffffffffffffffffffffffffffffffffffff1681565b61063b61035c565b15610672576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610f37565b610681868686868686326108fd565b505050505050565b60c95474010000000000000000000000000000000000000000900467ffffffffffffffff1681565b6106b9610882565b73ffffffffffffffffffffffffffffffffffffffff166106d76104ee565b73ffffffffffffffffffffffffffffffffffffffff1614610724576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610f6e565b60ca5460405173ffffffffffffffffffffffffffffffffffffffff83169082156108fc029083906000818181858888f1935050505015801561076a573d6000803e3d6000fd5b5050600060ca5550565b61077c610882565b73ffffffffffffffffffffffffffffffffffffffff1661079a6104ee565b73ffffffffffffffffffffffffffffffffffffffff16146107e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610f6e565b73ffffffffffffffffffffffffffffffffffffffff8116610834576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610ea3565b61083d81610886565b50565b60008787878787878760405160200161085f9796959493929190610dcc565b604051602081830303815290604052805190602001209050979650505050505050565b3390565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610907610b04565b905080871415610943576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610e35565b6000610950888686610365565b90508034101561098c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f890610fa3565b60006109bd33848c8c60c960149054906101000a900467ffffffffffffffff1667ffffffffffffffff168d8d610840565b905080893373ffffffffffffffffffffffffffffffffffffffff167f864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a868e8d8d60c960149054906101000a900467ffffffffffffffff168e8e8c604051610a2b989796959493929190610fda565b60405180910390a48160ca6000828254610a45919061106b565b909155505060c98054601490610a7c9074010000000000000000000000000000000000000000900467ffffffffffffffff1661109a565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555081341115610af85773ffffffffffffffffffffffffffffffffffffffff84166108fc610ace8434611083565b6040518115909202916000818181858888f19350505050158015610af6573d6000803e3d6000fd5b505b50505050505050505050565b4690565b60008083601f840112610b19578182fd5b50813567ffffffffffffffff811115610b30578182fd5b602083019150836020828501011115610b4857600080fd5b9250929050565b600060208284031215610b60578081fd5b8135610b6b816110f1565b9392505050565b600080600080600080600060c0888a031215610b8c578283fd5b8735610b97816110f1565b96506020880135955060408801359450606088013593506080880135925060a088013567ffffffffffffffff811115610bce578283fd5b610bda8a828b01610b08565b989b979a50959850939692959293505050565b60008060008060008060808789031215610c05578182fd5b8635955060208701359450604087013567ffffffffffffffff80821115610c2a578384fd5b610c368a838b01610b08565b90965094506060890135915080821115610c4e578384fd5b50610c5b89828a01610b08565b979a9699509497509295939492505050565b600080600080600080600060a0888a031215610c87578283fd5b8735965060208801359550604088013567ffffffffffffffff80821115610cac578485fd5b610cb88b838c01610b08565b909750955060608a0135915080821115610cd0578485fd5b50610cdd8a828b01610b08565b9094509250506080880135610cf1816110f1565b8091505092959891949750929550565b600060208284031215610d12578081fd5b5051919050565b600080600060408486031215610d2d578283fd5b83359250602084013567ffffffffffffffff811115610d4a578283fd5b610d5686828701610b08565b9497909650939450505050565b600082845282826020860137806020848601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011685010190509392505050565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b600073ffffffffffffffffffffffffffffffffffffffff8916825287602083015286604083015285606083015284608083015260c060a0830152610e1460c083018486610d63565b9998505050505050505050565b901515815260200190565b90815260200190565b6020808252600f908201527f496e76616c696420636861696e49640000000000000000000000000000000000604082015260600190565b6020808252600b908201527f466565206e6f7420736574000000000000000000000000000000000000000000604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201527f6464726573730000000000000000000000000000000000000000000000000000606082015260800190565b6020808252600f908201527f43616e6e6f742073657420746f20300000000000000000000000000000000000604082015260600190565b60208082526010908201527f5061757361626c653a2070617573656400000000000000000000000000000000604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526014908201527f496e73756666696369656e742067617320666565000000000000000000000000604082015260600190565b600089825288602083015260c06040830152610ffa60c08301888a610d63565b67ffffffffffffffff87166060840152828103608084015261101d818688610d63565b9150508260a08301529998505050505050505050565b60008482526040602083015261104d604083018486610d63565b95945050505050565b67ffffffffffffffff91909116815260200190565b6000821982111561107e5761107e6110c2565b500190565b600082821015611095576110956110c2565b500390565b600067ffffffffffffffff808316818114156110b8576110b86110c2565b6001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461083d57600080fdfea2646970667358221220d9b4cb52b2968bd04d9a99e5cfad1610f75b886346400d381e742051528a901964736f6c63430008000033",
}

// MessageBusSenderUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusSenderUpgradeableMetaData.ABI instead.
var MessageBusSenderUpgradeableABI = MessageBusSenderUpgradeableMetaData.ABI

// Deprecated: Use MessageBusSenderUpgradeableMetaData.Sigs instead.
// MessageBusSenderUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var MessageBusSenderUpgradeableFuncSigs = MessageBusSenderUpgradeableMetaData.Sigs

// MessageBusSenderUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusSenderUpgradeableMetaData.Bin instead.
var MessageBusSenderUpgradeableBin = MessageBusSenderUpgradeableMetaData.Bin

// DeployMessageBusSenderUpgradeable deploys a new Ethereum contract, binding an instance of MessageBusSenderUpgradeable to it.
func DeployMessageBusSenderUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageBusSenderUpgradeable, error) {
	parsed, err := MessageBusSenderUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusSenderUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusSenderUpgradeable{MessageBusSenderUpgradeableCaller: MessageBusSenderUpgradeableCaller{contract: contract}, MessageBusSenderUpgradeableTransactor: MessageBusSenderUpgradeableTransactor{contract: contract}, MessageBusSenderUpgradeableFilterer: MessageBusSenderUpgradeableFilterer{contract: contract}}, nil
}

// MessageBusSenderUpgradeable is an auto generated Go binding around an Ethereum contract.
type MessageBusSenderUpgradeable struct {
	MessageBusSenderUpgradeableCaller     // Read-only binding to the contract
	MessageBusSenderUpgradeableTransactor // Write-only binding to the contract
	MessageBusSenderUpgradeableFilterer   // Log filterer for contract events
}

// MessageBusSenderUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusSenderUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusSenderUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusSenderUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusSenderUpgradeableSession struct {
	Contract     *MessageBusSenderUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MessageBusSenderUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusSenderUpgradeableCallerSession struct {
	Contract *MessageBusSenderUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// MessageBusSenderUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusSenderUpgradeableTransactorSession struct {
	Contract     *MessageBusSenderUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// MessageBusSenderUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusSenderUpgradeableRaw struct {
	Contract *MessageBusSenderUpgradeable // Generic contract binding to access the raw methods on
}

// MessageBusSenderUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusSenderUpgradeableCallerRaw struct {
	Contract *MessageBusSenderUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusSenderUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusSenderUpgradeableTransactorRaw struct {
	Contract *MessageBusSenderUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusSenderUpgradeable creates a new instance of MessageBusSenderUpgradeable, bound to a specific deployed contract.
func NewMessageBusSenderUpgradeable(address common.Address, backend bind.ContractBackend) (*MessageBusSenderUpgradeable, error) {
	contract, err := bindMessageBusSenderUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderUpgradeable{MessageBusSenderUpgradeableCaller: MessageBusSenderUpgradeableCaller{contract: contract}, MessageBusSenderUpgradeableTransactor: MessageBusSenderUpgradeableTransactor{contract: contract}, MessageBusSenderUpgradeableFilterer: MessageBusSenderUpgradeableFilterer{contract: contract}}, nil
}

// NewMessageBusSenderUpgradeableCaller creates a new read-only instance of MessageBusSenderUpgradeable, bound to a specific deployed contract.
func NewMessageBusSenderUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*MessageBusSenderUpgradeableCaller, error) {
	contract, err := bindMessageBusSenderUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderUpgradeableCaller{contract: contract}, nil
}

// NewMessageBusSenderUpgradeableTransactor creates a new write-only instance of MessageBusSenderUpgradeable, bound to a specific deployed contract.
func NewMessageBusSenderUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusSenderUpgradeableTransactor, error) {
	contract, err := bindMessageBusSenderUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderUpgradeableTransactor{contract: contract}, nil
}

// NewMessageBusSenderUpgradeableFilterer creates a new log filterer instance of MessageBusSenderUpgradeable, bound to a specific deployed contract.
func NewMessageBusSenderUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusSenderUpgradeableFilterer, error) {
	contract, err := bindMessageBusSenderUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderUpgradeableFilterer{contract: contract}, nil
}

// bindMessageBusSenderUpgradeable binds a generic wrapper to an already deployed contract.
func bindMessageBusSenderUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusSenderUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusSenderUpgradeable.Contract.MessageBusSenderUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.MessageBusSenderUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.MessageBusSenderUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusSenderUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCaller) ComputeMessageId(opts *bind.CallOpts, _srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageBusSenderUpgradeable.contract.Call(opts, &out, "computeMessageId", _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) ComputeMessageId(_srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	return _MessageBusSenderUpgradeable.Contract.ComputeMessageId(&_MessageBusSenderUpgradeable.CallOpts, _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)
}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCallerSession) ComputeMessageId(_srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	return _MessageBusSenderUpgradeable.Contract.ComputeMessageId(&_MessageBusSenderUpgradeable.CallOpts, _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCaller) Fees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSenderUpgradeable.contract.Call(opts, &out, "fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) Fees() (*big.Int, error) {
	return _MessageBusSenderUpgradeable.Contract.Fees(&_MessageBusSenderUpgradeable.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCallerSession) Fees() (*big.Int, error) {
	return _MessageBusSenderUpgradeable.Contract.Fees(&_MessageBusSenderUpgradeable.CallOpts)
}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCaller) GasFeePricing(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusSenderUpgradeable.contract.Call(opts, &out, "gasFeePricing")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) GasFeePricing() (common.Address, error) {
	return _MessageBusSenderUpgradeable.Contract.GasFeePricing(&_MessageBusSenderUpgradeable.CallOpts)
}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCallerSession) GasFeePricing() (common.Address, error) {
	return _MessageBusSenderUpgradeable.Contract.GasFeePricing(&_MessageBusSenderUpgradeable.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCaller) Nonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MessageBusSenderUpgradeable.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) Nonce() (uint64, error) {
	return _MessageBusSenderUpgradeable.Contract.Nonce(&_MessageBusSenderUpgradeable.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCallerSession) Nonce() (uint64, error) {
	return _MessageBusSenderUpgradeable.Contract.Nonce(&_MessageBusSenderUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusSenderUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) Owner() (common.Address, error) {
	return _MessageBusSenderUpgradeable.Contract.Owner(&_MessageBusSenderUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCallerSession) Owner() (common.Address, error) {
	return _MessageBusSenderUpgradeable.Contract.Owner(&_MessageBusSenderUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MessageBusSenderUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) Paused() (bool, error) {
	return _MessageBusSenderUpgradeable.Contract.Paused(&_MessageBusSenderUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableCallerSession) Paused() (bool, error) {
	return _MessageBusSenderUpgradeable.Contract.Paused(&_MessageBusSenderUpgradeable.CallOpts)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactor) EstimateFee(opts *bind.TransactOpts, _dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.contract.Transact(opts, "estimateFee", _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.EstimateFee(&_MessageBusSenderUpgradeable.TransactOpts, _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.EstimateFee(&_MessageBusSenderUpgradeable.TransactOpts, _dstChainId, _options)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.RenounceOwnership(&_MessageBusSenderUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.RenounceOwnership(&_MessageBusSenderUpgradeable.TransactOpts)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactor) RescueGas(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.contract.Transact(opts, "rescueGas", to)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) RescueGas(to common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.RescueGas(&_MessageBusSenderUpgradeable.TransactOpts, to)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorSession) RescueGas(to common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.RescueGas(&_MessageBusSenderUpgradeable.TransactOpts, to)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactor) SendMessage(opts *bind.TransactOpts, _receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.SendMessage(&_MessageBusSenderUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.SendMessage(&_MessageBusSenderUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactor) SendMessage0(opts *bind.TransactOpts, _receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.contract.Transact(opts, "sendMessage0", _receiver, _dstChainId, _message, _options)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) SendMessage0(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.SendMessage0(&_MessageBusSenderUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorSession) SendMessage0(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.SendMessage0(&_MessageBusSenderUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.TransferOwnership(&_MessageBusSenderUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.TransferOwnership(&_MessageBusSenderUpgradeable.TransactOpts, newOwner)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactor) UpdateGasFeePricing(opts *bind.TransactOpts, _gasFeePricing common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.contract.Transact(opts, "updateGasFeePricing", _gasFeePricing)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) UpdateGasFeePricing(_gasFeePricing common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.UpdateGasFeePricing(&_MessageBusSenderUpgradeable.TransactOpts, _gasFeePricing)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorSession) UpdateGasFeePricing(_gasFeePricing common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.UpdateGasFeePricing(&_MessageBusSenderUpgradeable.TransactOpts, _gasFeePricing)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactor) WithdrawGasFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.contract.Transact(opts, "withdrawGasFees", to)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableSession) WithdrawGasFees(to common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.WithdrawGasFees(&_MessageBusSenderUpgradeable.TransactOpts, to)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableTransactorSession) WithdrawGasFees(to common.Address) (*types.Transaction, error) {
	return _MessageBusSenderUpgradeable.Contract.WithdrawGasFees(&_MessageBusSenderUpgradeable.TransactOpts, to)
}

// MessageBusSenderUpgradeableMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the MessageBusSenderUpgradeable contract.
type MessageBusSenderUpgradeableMessageSentIterator struct {
	Event *MessageBusSenderUpgradeableMessageSent // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderUpgradeableMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderUpgradeableMessageSent)
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
		it.Event = new(MessageBusSenderUpgradeableMessageSent)
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
func (it *MessageBusSenderUpgradeableMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderUpgradeableMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderUpgradeableMessageSent represents a MessageSent event raised by the MessageBusSenderUpgradeable contract.
type MessageBusSenderUpgradeableMessageSent struct {
	Sender     common.Address
	SrcChainID *big.Int
	Receiver   [32]byte
	DstChainId *big.Int
	Message    []byte
	Nonce      uint64
	Options    []byte
	Fee        *big.Int
	MessageId  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) FilterMessageSent(opts *bind.FilterOpts, sender []common.Address, dstChainId []*big.Int, messageId [][32]byte) (*MessageBusSenderUpgradeableMessageSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _MessageBusSenderUpgradeable.contract.FilterLogs(opts, "MessageSent", senderRule, dstChainIdRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderUpgradeableMessageSentIterator{contract: _MessageBusSenderUpgradeable.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MessageBusSenderUpgradeableMessageSent, sender []common.Address, dstChainId []*big.Int, messageId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _MessageBusSenderUpgradeable.contract.WatchLogs(opts, "MessageSent", senderRule, dstChainIdRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderUpgradeableMessageSent)
				if err := _MessageBusSenderUpgradeable.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) ParseMessageSent(log types.Log) (*MessageBusSenderUpgradeableMessageSent, error) {
	event := new(MessageBusSenderUpgradeableMessageSent)
	if err := _MessageBusSenderUpgradeable.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusSenderUpgradeable contract.
type MessageBusSenderUpgradeableOwnershipTransferredIterator struct {
	Event *MessageBusSenderUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderUpgradeableOwnershipTransferred)
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
		it.Event = new(MessageBusSenderUpgradeableOwnershipTransferred)
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
func (it *MessageBusSenderUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusSenderUpgradeable contract.
type MessageBusSenderUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusSenderUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusSenderUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderUpgradeableOwnershipTransferredIterator{contract: _MessageBusSenderUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusSenderUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusSenderUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderUpgradeableOwnershipTransferred)
				if err := _MessageBusSenderUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusSenderUpgradeableOwnershipTransferred, error) {
	event := new(MessageBusSenderUpgradeableOwnershipTransferred)
	if err := _MessageBusSenderUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MessageBusSenderUpgradeable contract.
type MessageBusSenderUpgradeablePausedIterator struct {
	Event *MessageBusSenderUpgradeablePaused // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderUpgradeablePaused)
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
		it.Event = new(MessageBusSenderUpgradeablePaused)
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
func (it *MessageBusSenderUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderUpgradeablePaused represents a Paused event raised by the MessageBusSenderUpgradeable contract.
type MessageBusSenderUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*MessageBusSenderUpgradeablePausedIterator, error) {

	logs, sub, err := _MessageBusSenderUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderUpgradeablePausedIterator{contract: _MessageBusSenderUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MessageBusSenderUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _MessageBusSenderUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderUpgradeablePaused)
				if err := _MessageBusSenderUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) ParsePaused(log types.Log) (*MessageBusSenderUpgradeablePaused, error) {
	event := new(MessageBusSenderUpgradeablePaused)
	if err := _MessageBusSenderUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MessageBusSenderUpgradeable contract.
type MessageBusSenderUpgradeableUnpausedIterator struct {
	Event *MessageBusSenderUpgradeableUnpaused // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderUpgradeableUnpaused)
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
		it.Event = new(MessageBusSenderUpgradeableUnpaused)
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
func (it *MessageBusSenderUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderUpgradeableUnpaused represents a Unpaused event raised by the MessageBusSenderUpgradeable contract.
type MessageBusSenderUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MessageBusSenderUpgradeableUnpausedIterator, error) {

	logs, sub, err := _MessageBusSenderUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderUpgradeableUnpausedIterator{contract: _MessageBusSenderUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MessageBusSenderUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _MessageBusSenderUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderUpgradeableUnpaused)
				if err := _MessageBusSenderUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MessageBusSenderUpgradeable *MessageBusSenderUpgradeableFilterer) ParseUnpaused(log types.Log) (*MessageBusSenderUpgradeableUnpaused, error) {
	event := new(MessageBusSenderUpgradeableUnpaused)
	if err := _MessageBusSenderUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusUpgradeableMetaData contains all meta data concerning the MessageBusUpgradeable contract.
var MessageBusUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CallReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcNonce\",\"type\":\"uint64\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_srcAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dstAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_srcNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"computeMessageId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasFeePricing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"getExecutedMessage\",\"outputs\":[{\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasFeePricing\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_authVerifier\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"rescueGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authVerifier\",\"type\":\"address\"}],\"name\":\"updateAuthVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasFeePricing\",\"type\":\"address\"}],\"name\":\"updateGasFeePricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"},{\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"updateMessageStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawGasFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c4087335": "authVerifier()",
		"f44d57aa": "computeMessageId(address,uint256,bytes32,uint256,uint256,bytes)",
		"5da6d2c4": "estimateFee(uint256,bytes)",
		"a1b058d8": "executeMessage(uint256,bytes32,address,uint256,uint256,bytes,bytes32)",
		"9af1d35a": "fees()",
		"aa70fc0e": "gasFeePricing()",
		"25b19fa3": "getExecutedMessage(bytes32)",
		"485cc955": "initialize(address,address)",
		"affed0e0": "nonce()",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"205e157b": "rescueGas(address)",
		"ac8a4c1b": "sendMessage(bytes32,uint256,bytes,bytes)",
		"72177189": "sendMessage(bytes32,uint256,bytes,bytes,address)",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"a5c0edf3": "updateAuthVerifier(address)",
		"a66dd384": "updateGasFeePricing(address)",
		"9b11079c": "updateMessageStatus(bytes32,uint8)",
		"d6b457b9": "withdrawGasFees(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611e77806100206000396000f3fe6080604052600436106101755760003560e01c80639b11079c116100cb578063ac8a4c1b1161007f578063d6b457b911610059578063d6b457b9146103a9578063f2fde38b146103c9578063f44d57aa146103e957610175565b8063ac8a4c1b1461035f578063affed0e014610372578063c40873351461039457610175565b8063a5c0edf3116100b0578063a5c0edf31461030a578063a66dd3841461032a578063aa70fc0e1461034a57610175565b80639b11079c146102ca578063a1b058d8146102ea57610175565b80635da6d2c41161012d5780638456cb59116101075780638456cb591461027e5780638da5cb5b146102935780639af1d35a146102b557610175565b80635da6d2c414610229578063715018a614610256578063721771891461026b57610175565b80633f4ba83a1161015e5780633f4ba83a146101d2578063485cc955146101e75780635c975abb1461020757610175565b8063205e157b1461017a57806325b19fa31461019c575b600080fd5b34801561018657600080fd5b5061019a6101953660046113fe565b610409565b005b3480156101a857600080fd5b506101bc6101b73660046114f4565b6104c6565b6040516101c9919061199e565b60405180910390f35b3480156101de57600080fd5b5061019a6104de565b3480156101f357600080fd5b5061019a610202366004611421565b610541565b34801561021357600080fd5b5061021c61061d565b6040516101c9919061192f565b34801561023557600080fd5b506102496102443660046117a2565b610626565b6040516101c9919061193a565b34801561026257600080fd5b5061019a6106fe565b61019a6102793660046115b3565b610761565b34801561028a57600080fd5b5061019a610779565b34801561029f57600080fd5b506102a86107da565b6040516101c991906118b9565b3480156102c157600080fd5b506102496107f6565b3480156102d657600080fd5b5061019a6102e536600461150c565b6107fc565b3480156102f657600080fd5b5061019a610305366004611720565b6108b0565b34801561031657600080fd5b5061019a6103253660046113fe565b610b91565b34801561033657600080fd5b5061019a6103453660046113fe565b610c64565b34801561035657600080fd5b506102a8610d37565b61019a61036d366004611533565b610d53565b34801561037e57600080fd5b50610387610d8f565b6040516101c99190611d25565b3480156103a057600080fd5b506102a8610db7565b3480156103b557600080fd5b5061019a6103c43660046113fe565b610dd3565b3480156103d557600080fd5b5061019a6103e43660046113fe565b610e7c565b3480156103f557600080fd5b50610249610404366004611459565b610f14565b610411610f56565b73ffffffffffffffffffffffffffffffffffffffff1661042f6107da565b73ffffffffffffffffffffffffffffffffffffffff161461046b5760405162461bcd60e51b815260040161046290611ba9565b60405180910390fd5b600060ca544761047b9190611d52565b60405190915073ffffffffffffffffffffffffffffffffffffffff83169082156108fc029083906000818181858888f193505050501580156104c1573d6000803e3d6000fd5b505050565b600081815260fb602052604090205460ff165b919050565b6104e6610f56565b73ffffffffffffffffffffffffffffffffffffffff166105046107da565b73ffffffffffffffffffffffffffffffffffffffff16146105375760405162461bcd60e51b815260040161046290611ba9565b61053f610f5a565b565b600054610100900460ff1661055c5760005460ff1615610564565b610564610fc8565b6105805760405162461bcd60e51b815260040161046290611b4c565b600054610100900460ff161580156105c8576000805460ff197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b6105d0610fd9565b6105d8611010565b6105e183611043565b6105ea82610c1d565b80156104c157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055505050565b60655460ff1690565b60c9546040517f47feadc1000000000000000000000000000000000000000000000000000000008152600091829173ffffffffffffffffffffffffffffffffffffffff909116906347feadc19061068590889088908890600401611d02565b602060405180830381600087803b15801561069f57600080fd5b505af11580156106b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d79190611708565b9050806106f65760405162461bcd60e51b815260040161046290611a4a565b949350505050565b610706610f56565b73ffffffffffffffffffffffffffffffffffffffff166107246107da565b73ffffffffffffffffffffffffffffffffffffffff16146107575760405162461bcd60e51b815260040161046290611ba9565b61053f600061106a565b610770878787878787876110e1565b50505050505050565b610781610f56565b73ffffffffffffffffffffffffffffffffffffffff1661079f6107da565b73ffffffffffffffffffffffffffffffffffffffff16146107d25760405162461bcd60e51b815260040161046290611ba9565b61053f6112b4565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b60ca5481565b610804610f56565b73ffffffffffffffffffffffffffffffffffffffff166108226107da565b73ffffffffffffffffffffffffffffffffffffffff16146108555760405162461bcd60e51b815260040161046290611ba9565b600082815260fb60205260409020805482919060ff191660018360028111156108a7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055505050565b6108b861061d565b156108d55760405162461bcd60e51b815260040161046290611b15565b600081815260fb602052604081205460ff16600281111561091f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1461093c5760405162461bcd60e51b815260040161046290611bde565b60fa5460405173ffffffffffffffffffffffffffffffffffffffff90911690638b1b3a2d9061096f9033906020016118b9565b6040516020818303038152906040526040518263ffffffff1660e01b815260040161099a919061198b565b60206040518083038186803b1580156109b257600080fd5b505afa1580156109c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ea91906114d4565b5060008673ffffffffffffffffffffffffffffffffffffffff1663a6060871878a8c8888336040518763ffffffff1660e01b8152600401610a2f959493929190611943565b600060405180830381600088803b158015610a4957600080fd5b5087f193505050508015610a5b575060015b610ad8573d808015610a89576040519150601f19603f3d011682016040523d82523d6000602084013e610a8e565b606091505b507fffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f610ab98261130f565b604051610ac6919061198b565b60405180910390a16002915050610adc565b5060015b600082815260fb60205260409020805482919060ff19166001836002811115610b2e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055508673ffffffffffffffffffffffffffffffffffffffff16827f04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65838c89604051610b7e939291906119ac565b60405180910390a3505050505050505050565b610b99610f56565b73ffffffffffffffffffffffffffffffffffffffff16610bb76107da565b73ffffffffffffffffffffffffffffffffffffffff1614610bea5760405162461bcd60e51b815260040161046290611ba9565b73ffffffffffffffffffffffffffffffffffffffff8116610c1d5760405162461bcd60e51b815260040161046290611ade565b60fa80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610c6c610f56565b73ffffffffffffffffffffffffffffffffffffffff16610c8a6107da565b73ffffffffffffffffffffffffffffffffffffffff1614610cbd5760405162461bcd60e51b815260040161046290611ba9565b73ffffffffffffffffffffffffffffffffffffffff8116610cf05760405162461bcd60e51b815260040161046290611ade565b60c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60c95473ffffffffffffffffffffffffffffffffffffffff1681565b610d5b61061d565b15610d785760405162461bcd60e51b815260040161046290611b15565b610d87868686868686326110e1565b505050505050565b60c95474010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60fa5473ffffffffffffffffffffffffffffffffffffffff1681565b610ddb610f56565b73ffffffffffffffffffffffffffffffffffffffff16610df96107da565b73ffffffffffffffffffffffffffffffffffffffff1614610e2c5760405162461bcd60e51b815260040161046290611ba9565b60ca5460405173ffffffffffffffffffffffffffffffffffffffff83169082156108fc029083906000818181858888f19350505050158015610e72573d6000803e3d6000fd5b5050600060ca5550565b610e84610f56565b73ffffffffffffffffffffffffffffffffffffffff16610ea26107da565b73ffffffffffffffffffffffffffffffffffffffff1614610ed55760405162461bcd60e51b815260040161046290611ba9565b73ffffffffffffffffffffffffffffffffffffffff8116610f085760405162461bcd60e51b815260040161046290611a81565b610f118161106a565b50565b600087878787878787604051602001610f3397969594939291906118da565b604051602081830303815290604052805190602001209050979650505050505050565b3390565b610f6261061d565b610f7e5760405162461bcd60e51b815260040161046290611a13565b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610fb1610f56565b604051610fbe91906118b9565b60405180910390a1565b6000610fd330611375565b15905090565b600054610100900460ff166110005760405162461bcd60e51b815260040161046290611c4c565b61053f61100b610f56565b61106a565b600054610100900460ff166110375760405162461bcd60e51b815260040161046290611c4c565b6065805460ff19169055565b600054610100900460ff16610cf05760405162461bcd60e51b815260040161046290611c4c565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006110eb6113b3565b90508087141561110d5760405162461bcd60e51b8152600401610462906119dc565b600061111a888686610626565b90508034101561113c5760405162461bcd60e51b815260040161046290611c15565b600061116d33848c8c60c960149054906101000a900467ffffffffffffffff1667ffffffffffffffff168d8d610f14565b905080893373ffffffffffffffffffffffffffffffffffffffff167f864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a868e8d8d60c960149054906101000a900467ffffffffffffffff168e8e8c6040516111db989796959493929190611ca9565b60405180910390a48160ca60008282546111f59190611d3a565b909155505060c9805460149061122c9074010000000000000000000000000000000000000000900467ffffffffffffffff16611d99565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550813411156112a85773ffffffffffffffffffffffffffffffffffffffff84166108fc61127e8434611d52565b6040518115909202916000818181858888f193505050501580156112a6573d6000803e3d6000fd5b505b50505050505050505050565b6112bc61061d565b156112d95760405162461bcd60e51b815260040161046290611b15565b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610fb1610f56565b6060604482511015611355575060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c7900000060208201526104d9565b6004820191508180602001905181019061136f9190611647565b92915050565b6000808273ffffffffffffffffffffffffffffffffffffffff16803b806020016040519081016040528181526000908060200190933c511192915050565b4690565b60008083601f8401126113c8578182fd5b50813567ffffffffffffffff8111156113df578182fd5b6020830191508360208285010111156113f757600080fd5b9250929050565b60006020828403121561140f578081fd5b813561141a81611e1f565b9392505050565b60008060408385031215611433578081fd5b823561143e81611e1f565b9150602083013561144e81611e1f565b809150509250929050565b600080600080600080600060c0888a031215611473578283fd5b873561147e81611e1f565b96506020880135955060408801359450606088013593506080880135925060a088013567ffffffffffffffff8111156114b5578283fd5b6114c18a828b016113b7565b989b979a50959850939692959293505050565b6000602082840312156114e5578081fd5b8151801515811461141a578182fd5b600060208284031215611505578081fd5b5035919050565b6000806040838503121561151e578182fd5b8235915060208301356003811061144e578182fd5b6000806000806000806080878903121561154b578182fd5b8635955060208701359450604087013567ffffffffffffffff80821115611570578384fd5b61157c8a838b016113b7565b90965094506060890135915080821115611594578384fd5b506115a189828a016113b7565b979a9699509497509295939492505050565b600080600080600080600060a0888a0312156115cd578283fd5b8735965060208801359550604088013567ffffffffffffffff808211156115f2578485fd5b6115fe8b838c016113b7565b909750955060608a0135915080821115611616578485fd5b506116238a828b016113b7565b909450925050608088013561163781611e1f565b8091505092959891949750929550565b600060208284031215611658578081fd5b815167ffffffffffffffff8082111561166f578283fd5b818401915084601f830112611682578283fd5b81518181111561169457611694611df0565b60405160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011682010181811084821117156116d6576116d6611df0565b6040528181528382016020018710156116ed578485fd5b6116fe826020830160208701611d69565b9695505050505050565b600060208284031215611719578081fd5b5051919050565b60008060008060008060008060e0898b03121561173b578182fd5b8835975060208901359650604089013561175481611e1f565b9550606089013594506080890135935060a089013567ffffffffffffffff81111561177d578283fd5b6117898b828c016113b7565b999c989b50969995989497949560c00135949350505050565b6000806000604084860312156117b6578081fd5b83359250602084013567ffffffffffffffff8111156117d3578182fd5b6117df868287016113b7565b9497909650939450505050565b600082845282826020860137806020848601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011685010190509392505050565b6000815180845261184c816020860160208601611d69565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600381106118b5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b600073ffffffffffffffffffffffffffffffffffffffff8916825287602083015286604083015285606083015284608083015260c060a083015261192260c0830184866117ec565b9998505050505050505050565b901515815260200190565b90815260200190565b6000868252856020830152608060408301526119636080830185876117ec565b905073ffffffffffffffffffffffffffffffffffffffff831660608301529695505050505050565b60006020825261141a6020830184611834565b6020810161136f828461187e565b606081016119ba828661187e565b67ffffffffffffffff8085166020840152808416604084015250949350505050565b6020808252600f908201527f496e76616c696420636861696e49640000000000000000000000000000000000604082015260600190565b60208082526014908201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604082015260600190565b6020808252600b908201527f466565206e6f7420736574000000000000000000000000000000000000000000604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201527f6464726573730000000000000000000000000000000000000000000000000000606082015260800190565b6020808252600f908201527f43616e6e6f742073657420746f20300000000000000000000000000000000000604082015260600190565b60208082526010908201527f5061757361626c653a2070617573656400000000000000000000000000000000604082015260600190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201527f647920696e697469616c697a6564000000000000000000000000000000000000606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526018908201527f4d65737361676520616c72656164792065786563757465640000000000000000604082015260600190565b60208082526014908201527f496e73756666696369656e742067617320666565000000000000000000000000604082015260600190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201527f6e697469616c697a696e67000000000000000000000000000000000000000000606082015260800190565b600089825288602083015260c06040830152611cc960c08301888a6117ec565b67ffffffffffffffff871660608401528281036080840152611cec8186886117ec565b9150508260a08301529998505050505050505050565b600084825260406020830152611d1c6040830184866117ec565b95945050505050565b67ffffffffffffffff91909116815260200190565b60008219821115611d4d57611d4d611dc1565b500190565b600082821015611d6457611d64611dc1565b500390565b60005b83811015611d84578181015183820152602001611d6c565b83811115611d93576000848401525b50505050565b600067ffffffffffffffff80831681811415611db757611db7611dc1565b6001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81168114610f1157600080fdfea2646970667358221220131d19f31d38b0c80dcd99484abda79d51675a25ae643279d41130dcf95e80b264736f6c63430008000033",
}

// MessageBusUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusUpgradeableMetaData.ABI instead.
var MessageBusUpgradeableABI = MessageBusUpgradeableMetaData.ABI

// Deprecated: Use MessageBusUpgradeableMetaData.Sigs instead.
// MessageBusUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var MessageBusUpgradeableFuncSigs = MessageBusUpgradeableMetaData.Sigs

// MessageBusUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusUpgradeableMetaData.Bin instead.
var MessageBusUpgradeableBin = MessageBusUpgradeableMetaData.Bin

// DeployMessageBusUpgradeable deploys a new Ethereum contract, binding an instance of MessageBusUpgradeable to it.
func DeployMessageBusUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageBusUpgradeable, error) {
	parsed, err := MessageBusUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusUpgradeable{MessageBusUpgradeableCaller: MessageBusUpgradeableCaller{contract: contract}, MessageBusUpgradeableTransactor: MessageBusUpgradeableTransactor{contract: contract}, MessageBusUpgradeableFilterer: MessageBusUpgradeableFilterer{contract: contract}}, nil
}

// MessageBusUpgradeable is an auto generated Go binding around an Ethereum contract.
type MessageBusUpgradeable struct {
	MessageBusUpgradeableCaller     // Read-only binding to the contract
	MessageBusUpgradeableTransactor // Write-only binding to the contract
	MessageBusUpgradeableFilterer   // Log filterer for contract events
}

// MessageBusUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusUpgradeableSession struct {
	Contract     *MessageBusUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MessageBusUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusUpgradeableCallerSession struct {
	Contract *MessageBusUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// MessageBusUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusUpgradeableTransactorSession struct {
	Contract     *MessageBusUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// MessageBusUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusUpgradeableRaw struct {
	Contract *MessageBusUpgradeable // Generic contract binding to access the raw methods on
}

// MessageBusUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusUpgradeableCallerRaw struct {
	Contract *MessageBusUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusUpgradeableTransactorRaw struct {
	Contract *MessageBusUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusUpgradeable creates a new instance of MessageBusUpgradeable, bound to a specific deployed contract.
func NewMessageBusUpgradeable(address common.Address, backend bind.ContractBackend) (*MessageBusUpgradeable, error) {
	contract, err := bindMessageBusUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeable{MessageBusUpgradeableCaller: MessageBusUpgradeableCaller{contract: contract}, MessageBusUpgradeableTransactor: MessageBusUpgradeableTransactor{contract: contract}, MessageBusUpgradeableFilterer: MessageBusUpgradeableFilterer{contract: contract}}, nil
}

// NewMessageBusUpgradeableCaller creates a new read-only instance of MessageBusUpgradeable, bound to a specific deployed contract.
func NewMessageBusUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*MessageBusUpgradeableCaller, error) {
	contract, err := bindMessageBusUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeableCaller{contract: contract}, nil
}

// NewMessageBusUpgradeableTransactor creates a new write-only instance of MessageBusUpgradeable, bound to a specific deployed contract.
func NewMessageBusUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusUpgradeableTransactor, error) {
	contract, err := bindMessageBusUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeableTransactor{contract: contract}, nil
}

// NewMessageBusUpgradeableFilterer creates a new log filterer instance of MessageBusUpgradeable, bound to a specific deployed contract.
func NewMessageBusUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusUpgradeableFilterer, error) {
	contract, err := bindMessageBusUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeableFilterer{contract: contract}, nil
}

// bindMessageBusUpgradeable binds a generic wrapper to an already deployed contract.
func bindMessageBusUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusUpgradeable *MessageBusUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusUpgradeable.Contract.MessageBusUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusUpgradeable *MessageBusUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.MessageBusUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusUpgradeable *MessageBusUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.MessageBusUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableCaller) AuthVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusUpgradeable.contract.Call(opts, &out, "authVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) AuthVerifier() (common.Address, error) {
	return _MessageBusUpgradeable.Contract.AuthVerifier(&_MessageBusUpgradeable.CallOpts)
}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerSession) AuthVerifier() (common.Address, error) {
	return _MessageBusUpgradeable.Contract.AuthVerifier(&_MessageBusUpgradeable.CallOpts)
}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_MessageBusUpgradeable *MessageBusUpgradeableCaller) ComputeMessageId(opts *bind.CallOpts, _srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageBusUpgradeable.contract.Call(opts, &out, "computeMessageId", _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) ComputeMessageId(_srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	return _MessageBusUpgradeable.Contract.ComputeMessageId(&_MessageBusUpgradeable.CallOpts, _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)
}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerSession) ComputeMessageId(_srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	return _MessageBusUpgradeable.Contract.ComputeMessageId(&_MessageBusUpgradeable.CallOpts, _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_MessageBusUpgradeable *MessageBusUpgradeableCaller) Fees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusUpgradeable.contract.Call(opts, &out, "fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) Fees() (*big.Int, error) {
	return _MessageBusUpgradeable.Contract.Fees(&_MessageBusUpgradeable.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerSession) Fees() (*big.Int, error) {
	return _MessageBusUpgradeable.Contract.Fees(&_MessageBusUpgradeable.CallOpts)
}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableCaller) GasFeePricing(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusUpgradeable.contract.Call(opts, &out, "gasFeePricing")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) GasFeePricing() (common.Address, error) {
	return _MessageBusUpgradeable.Contract.GasFeePricing(&_MessageBusUpgradeable.CallOpts)
}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerSession) GasFeePricing() (common.Address, error) {
	return _MessageBusUpgradeable.Contract.GasFeePricing(&_MessageBusUpgradeable.CallOpts)
}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_MessageBusUpgradeable *MessageBusUpgradeableCaller) GetExecutedMessage(opts *bind.CallOpts, _messageId [32]byte) (uint8, error) {
	var out []interface{}
	err := _MessageBusUpgradeable.contract.Call(opts, &out, "getExecutedMessage", _messageId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) GetExecutedMessage(_messageId [32]byte) (uint8, error) {
	return _MessageBusUpgradeable.Contract.GetExecutedMessage(&_MessageBusUpgradeable.CallOpts, _messageId)
}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerSession) GetExecutedMessage(_messageId [32]byte) (uint8, error) {
	return _MessageBusUpgradeable.Contract.GetExecutedMessage(&_MessageBusUpgradeable.CallOpts, _messageId)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_MessageBusUpgradeable *MessageBusUpgradeableCaller) Nonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MessageBusUpgradeable.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) Nonce() (uint64, error) {
	return _MessageBusUpgradeable.Contract.Nonce(&_MessageBusUpgradeable.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerSession) Nonce() (uint64, error) {
	return _MessageBusUpgradeable.Contract.Nonce(&_MessageBusUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) Owner() (common.Address, error) {
	return _MessageBusUpgradeable.Contract.Owner(&_MessageBusUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerSession) Owner() (common.Address, error) {
	return _MessageBusUpgradeable.Contract.Owner(&_MessageBusUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusUpgradeable *MessageBusUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MessageBusUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) Paused() (bool, error) {
	return _MessageBusUpgradeable.Contract.Paused(&_MessageBusUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageBusUpgradeable *MessageBusUpgradeableCallerSession) Paused() (bool, error) {
	return _MessageBusUpgradeable.Contract.Paused(&_MessageBusUpgradeable.CallOpts)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) EstimateFee(opts *bind.TransactOpts, _dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "estimateFee", _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.EstimateFee(&_MessageBusUpgradeable.TransactOpts, _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.EstimateFee(&_MessageBusUpgradeable.TransactOpts, _dstChainId, _options)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "executeMessage", _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.ExecuteMessage(&_MessageBusUpgradeable.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.ExecuteMessage(&_MessageBusUpgradeable.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _gasFeePricing, address _authVerifier) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) Initialize(opts *bind.TransactOpts, _gasFeePricing common.Address, _authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "initialize", _gasFeePricing, _authVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _gasFeePricing, address _authVerifier) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) Initialize(_gasFeePricing common.Address, _authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.Initialize(&_MessageBusUpgradeable.TransactOpts, _gasFeePricing, _authVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _gasFeePricing, address _authVerifier) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) Initialize(_gasFeePricing common.Address, _authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.Initialize(&_MessageBusUpgradeable.TransactOpts, _gasFeePricing, _authVerifier)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) Pause() (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.Pause(&_MessageBusUpgradeable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) Pause() (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.Pause(&_MessageBusUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.RenounceOwnership(&_MessageBusUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.RenounceOwnership(&_MessageBusUpgradeable.TransactOpts)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) RescueGas(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "rescueGas", to)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) RescueGas(to common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.RescueGas(&_MessageBusUpgradeable.TransactOpts, to)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) RescueGas(to common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.RescueGas(&_MessageBusUpgradeable.TransactOpts, to)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) SendMessage(opts *bind.TransactOpts, _receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.SendMessage(&_MessageBusUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.SendMessage(&_MessageBusUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) SendMessage0(opts *bind.TransactOpts, _receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "sendMessage0", _receiver, _dstChainId, _message, _options)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) SendMessage0(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.SendMessage0(&_MessageBusUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) SendMessage0(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.SendMessage0(&_MessageBusUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.TransferOwnership(&_MessageBusUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.TransferOwnership(&_MessageBusUpgradeable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) Unpause() (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.Unpause(&_MessageBusUpgradeable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) Unpause() (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.Unpause(&_MessageBusUpgradeable.TransactOpts)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) UpdateAuthVerifier(opts *bind.TransactOpts, _authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "updateAuthVerifier", _authVerifier)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) UpdateAuthVerifier(_authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.UpdateAuthVerifier(&_MessageBusUpgradeable.TransactOpts, _authVerifier)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) UpdateAuthVerifier(_authVerifier common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.UpdateAuthVerifier(&_MessageBusUpgradeable.TransactOpts, _authVerifier)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) UpdateGasFeePricing(opts *bind.TransactOpts, _gasFeePricing common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "updateGasFeePricing", _gasFeePricing)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) UpdateGasFeePricing(_gasFeePricing common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.UpdateGasFeePricing(&_MessageBusUpgradeable.TransactOpts, _gasFeePricing)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) UpdateGasFeePricing(_gasFeePricing common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.UpdateGasFeePricing(&_MessageBusUpgradeable.TransactOpts, _gasFeePricing)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) UpdateMessageStatus(opts *bind.TransactOpts, _messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "updateMessageStatus", _messageId, _status)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) UpdateMessageStatus(_messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.UpdateMessageStatus(&_MessageBusUpgradeable.TransactOpts, _messageId, _status)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) UpdateMessageStatus(_messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.UpdateMessageStatus(&_MessageBusUpgradeable.TransactOpts, _messageId, _status)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactor) WithdrawGasFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.contract.Transact(opts, "withdrawGasFees", to)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableSession) WithdrawGasFees(to common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.WithdrawGasFees(&_MessageBusUpgradeable.TransactOpts, to)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_MessageBusUpgradeable *MessageBusUpgradeableTransactorSession) WithdrawGasFees(to common.Address) (*types.Transaction, error) {
	return _MessageBusUpgradeable.Contract.WithdrawGasFees(&_MessageBusUpgradeable.TransactOpts, to)
}

// MessageBusUpgradeableCallRevertedIterator is returned from FilterCallReverted and is used to iterate over the raw logs and unpacked data for CallReverted events raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableCallRevertedIterator struct {
	Event *MessageBusUpgradeableCallReverted // Event containing the contract specifics and raw log

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
func (it *MessageBusUpgradeableCallRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusUpgradeableCallReverted)
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
		it.Event = new(MessageBusUpgradeableCallReverted)
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
func (it *MessageBusUpgradeableCallRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusUpgradeableCallRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusUpgradeableCallReverted represents a CallReverted event raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableCallReverted struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallReverted is a free log retrieval operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) FilterCallReverted(opts *bind.FilterOpts) (*MessageBusUpgradeableCallRevertedIterator, error) {

	logs, sub, err := _MessageBusUpgradeable.contract.FilterLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeableCallRevertedIterator{contract: _MessageBusUpgradeable.contract, event: "CallReverted", logs: logs, sub: sub}, nil
}

// WatchCallReverted is a free log subscription operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) WatchCallReverted(opts *bind.WatchOpts, sink chan<- *MessageBusUpgradeableCallReverted) (event.Subscription, error) {

	logs, sub, err := _MessageBusUpgradeable.contract.WatchLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusUpgradeableCallReverted)
				if err := _MessageBusUpgradeable.contract.UnpackLog(event, "CallReverted", log); err != nil {
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

// ParseCallReverted is a log parse operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) ParseCallReverted(log types.Log) (*MessageBusUpgradeableCallReverted, error) {
	event := new(MessageBusUpgradeableCallReverted)
	if err := _MessageBusUpgradeable.contract.UnpackLog(event, "CallReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusUpgradeableExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableExecutedIterator struct {
	Event *MessageBusUpgradeableExecuted // Event containing the contract specifics and raw log

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
func (it *MessageBusUpgradeableExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusUpgradeableExecuted)
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
		it.Event = new(MessageBusUpgradeableExecuted)
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
func (it *MessageBusUpgradeableExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusUpgradeableExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusUpgradeableExecuted represents a Executed event raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableExecuted struct {
	MessageId  [32]byte
	Status     uint8
	DstAddress common.Address
	SrcChainId uint64
	SrcNonce   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) FilterExecuted(opts *bind.FilterOpts, messageId [][32]byte, _dstAddress []common.Address) (*MessageBusUpgradeableExecutedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	var _dstAddressRule []interface{}
	for _, _dstAddressItem := range _dstAddress {
		_dstAddressRule = append(_dstAddressRule, _dstAddressItem)
	}

	logs, sub, err := _MessageBusUpgradeable.contract.FilterLogs(opts, "Executed", messageIdRule, _dstAddressRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeableExecutedIterator{contract: _MessageBusUpgradeable.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MessageBusUpgradeableExecuted, messageId [][32]byte, _dstAddress []common.Address) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	var _dstAddressRule []interface{}
	for _, _dstAddressItem := range _dstAddress {
		_dstAddressRule = append(_dstAddressRule, _dstAddressItem)
	}

	logs, sub, err := _MessageBusUpgradeable.contract.WatchLogs(opts, "Executed", messageIdRule, _dstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusUpgradeableExecuted)
				if err := _MessageBusUpgradeable.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) ParseExecuted(log types.Log) (*MessageBusUpgradeableExecuted, error) {
	event := new(MessageBusUpgradeableExecuted)
	if err := _MessageBusUpgradeable.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusUpgradeableMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableMessageSentIterator struct {
	Event *MessageBusUpgradeableMessageSent // Event containing the contract specifics and raw log

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
func (it *MessageBusUpgradeableMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusUpgradeableMessageSent)
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
		it.Event = new(MessageBusUpgradeableMessageSent)
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
func (it *MessageBusUpgradeableMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusUpgradeableMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusUpgradeableMessageSent represents a MessageSent event raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableMessageSent struct {
	Sender     common.Address
	SrcChainID *big.Int
	Receiver   [32]byte
	DstChainId *big.Int
	Message    []byte
	Nonce      uint64
	Options    []byte
	Fee        *big.Int
	MessageId  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) FilterMessageSent(opts *bind.FilterOpts, sender []common.Address, dstChainId []*big.Int, messageId [][32]byte) (*MessageBusUpgradeableMessageSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _MessageBusUpgradeable.contract.FilterLogs(opts, "MessageSent", senderRule, dstChainIdRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeableMessageSentIterator{contract: _MessageBusUpgradeable.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MessageBusUpgradeableMessageSent, sender []common.Address, dstChainId []*big.Int, messageId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _MessageBusUpgradeable.contract.WatchLogs(opts, "MessageSent", senderRule, dstChainIdRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusUpgradeableMessageSent)
				if err := _MessageBusUpgradeable.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) ParseMessageSent(log types.Log) (*MessageBusUpgradeableMessageSent, error) {
	event := new(MessageBusUpgradeableMessageSent)
	if err := _MessageBusUpgradeable.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableOwnershipTransferredIterator struct {
	Event *MessageBusUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusUpgradeableOwnershipTransferred)
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
		it.Event = new(MessageBusUpgradeableOwnershipTransferred)
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
func (it *MessageBusUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeableOwnershipTransferredIterator{contract: _MessageBusUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusUpgradeableOwnershipTransferred)
				if err := _MessageBusUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusUpgradeableOwnershipTransferred, error) {
	event := new(MessageBusUpgradeableOwnershipTransferred)
	if err := _MessageBusUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeablePausedIterator struct {
	Event *MessageBusUpgradeablePaused // Event containing the contract specifics and raw log

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
func (it *MessageBusUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusUpgradeablePaused)
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
		it.Event = new(MessageBusUpgradeablePaused)
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
func (it *MessageBusUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusUpgradeablePaused represents a Paused event raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*MessageBusUpgradeablePausedIterator, error) {

	logs, sub, err := _MessageBusUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeablePausedIterator{contract: _MessageBusUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MessageBusUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _MessageBusUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusUpgradeablePaused)
				if err := _MessageBusUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) ParsePaused(log types.Log) (*MessageBusUpgradeablePaused, error) {
	event := new(MessageBusUpgradeablePaused)
	if err := _MessageBusUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableUnpausedIterator struct {
	Event *MessageBusUpgradeableUnpaused // Event containing the contract specifics and raw log

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
func (it *MessageBusUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusUpgradeableUnpaused)
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
		it.Event = new(MessageBusUpgradeableUnpaused)
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
func (it *MessageBusUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusUpgradeableUnpaused represents a Unpaused event raised by the MessageBusUpgradeable contract.
type MessageBusUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MessageBusUpgradeableUnpausedIterator, error) {

	logs, sub, err := _MessageBusUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MessageBusUpgradeableUnpausedIterator{contract: _MessageBusUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MessageBusUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _MessageBusUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusUpgradeableUnpaused)
				if err := _MessageBusUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MessageBusUpgradeable *MessageBusUpgradeableFilterer) ParseUnpaused(log types.Log) (*MessageBusUpgradeableUnpaused, error) {
	event := new(MessageBusUpgradeableUnpaused)
	if err := _MessageBusUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableUpgradeableMetaData contains all meta data concerning the OwnableUpgradeable contract.
var OwnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// PausableUpgradeableMetaData contains all meta data concerning the PausableUpgradeable contract.
var PausableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c975abb": "paused()",
	},
}

// PausableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableUpgradeableMetaData.ABI instead.
var PausableUpgradeableABI = PausableUpgradeableMetaData.ABI

// Deprecated: Use PausableUpgradeableMetaData.Sigs instead.
// PausableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var PausableUpgradeableFuncSigs = PausableUpgradeableMetaData.Sigs

// PausableUpgradeable is an auto generated Go binding around an Ethereum contract.
type PausableUpgradeable struct {
	PausableUpgradeableCaller     // Read-only binding to the contract
	PausableUpgradeableTransactor // Write-only binding to the contract
	PausableUpgradeableFilterer   // Log filterer for contract events
}

// PausableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableUpgradeableSession struct {
	Contract     *PausableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PausableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableUpgradeableCallerSession struct {
	Contract *PausableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PausableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableUpgradeableTransactorSession struct {
	Contract     *PausableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PausableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableUpgradeableRaw struct {
	Contract *PausableUpgradeable // Generic contract binding to access the raw methods on
}

// PausableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableUpgradeableCallerRaw struct {
	Contract *PausableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableUpgradeableTransactorRaw struct {
	Contract *PausableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableUpgradeable creates a new instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeable(address common.Address, backend bind.ContractBackend) (*PausableUpgradeable, error) {
	contract, err := bindPausableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeable{PausableUpgradeableCaller: PausableUpgradeableCaller{contract: contract}, PausableUpgradeableTransactor: PausableUpgradeableTransactor{contract: contract}, PausableUpgradeableFilterer: PausableUpgradeableFilterer{contract: contract}}, nil
}

// NewPausableUpgradeableCaller creates a new read-only instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*PausableUpgradeableCaller, error) {
	contract, err := bindPausableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableCaller{contract: contract}, nil
}

// NewPausableUpgradeableTransactor creates a new write-only instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableUpgradeableTransactor, error) {
	contract, err := bindPausableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableTransactor{contract: contract}, nil
}

// NewPausableUpgradeableFilterer creates a new log filterer instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableUpgradeableFilterer, error) {
	contract, err := bindPausableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableFilterer{contract: contract}, nil
}

// bindPausableUpgradeable binds a generic wrapper to an already deployed contract.
func bindPausableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeable *PausableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeable.Contract.PausableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeable *PausableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.PausableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeable *PausableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.PausableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeable *PausableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeable *PausableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeable *PausableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PausableUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableSession) Paused() (bool, error) {
	return _PausableUpgradeable.Contract.Paused(&_PausableUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableCallerSession) Paused() (bool, error) {
	return _PausableUpgradeable.Contract.Paused(&_PausableUpgradeable.CallOpts)
}

// PausableUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PausableUpgradeable contract.
type PausableUpgradeablePausedIterator struct {
	Event *PausableUpgradeablePaused // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeablePaused)
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
		it.Event = new(PausableUpgradeablePaused)
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
func (it *PausableUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeablePaused represents a Paused event raised by the PausableUpgradeable contract.
type PausableUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausableUpgradeablePausedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeablePausedIterator{contract: _PausableUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeablePaused)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParsePaused(log types.Log) (*PausableUpgradeablePaused, error) {
	event := new(PausableUpgradeablePaused)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PausableUpgradeable contract.
type PausableUpgradeableUnpausedIterator struct {
	Event *PausableUpgradeableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeableUnpaused)
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
		it.Event = new(PausableUpgradeableUnpaused)
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
func (it *PausableUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeableUnpaused represents a Unpaused event raised by the PausableUpgradeable contract.
type PausableUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUpgradeableUnpausedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableUnpausedIterator{contract: _PausableUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeableUnpaused)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParseUnpaused(log types.Log) (*PausableUpgradeableUnpaused, error) {
	event := new(PausableUpgradeableUnpaused)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessageBusUpgradeableMetaData contains all meta data concerning the TestMessageBusUpgradeable contract.
var TestMessageBusUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CallReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcNonce\",\"type\":\"uint64\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_srcAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dstAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_srcNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"computeMessageId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasFeePricing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"getExecutedMessage\",\"outputs\":[{\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasFeePricing\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_authVerifier\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"rescueGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"testCallReverted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcNonce\",\"type\":\"uint64\"}],\"name\":\"testExecuted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcChainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"testMessageSent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authVerifier\",\"type\":\"address\"}],\"name\":\"updateAuthVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasFeePricing\",\"type\":\"address\"}],\"name\":\"updateGasFeePricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"},{\"internalType\":\"enumMessageBusReceiverUpgradeable.TxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"updateMessageStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawGasFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c4087335": "authVerifier()",
		"f44d57aa": "computeMessageId(address,uint256,bytes32,uint256,uint256,bytes)",
		"5da6d2c4": "estimateFee(uint256,bytes)",
		"a1b058d8": "executeMessage(uint256,bytes32,address,uint256,uint256,bytes,bytes32)",
		"9af1d35a": "fees()",
		"aa70fc0e": "gasFeePricing()",
		"25b19fa3": "getExecutedMessage(bytes32)",
		"485cc955": "initialize(address,address)",
		"affed0e0": "nonce()",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"205e157b": "rescueGas(address)",
		"ac8a4c1b": "sendMessage(bytes32,uint256,bytes,bytes)",
		"72177189": "sendMessage(bytes32,uint256,bytes,bytes,address)",
		"446e9045": "testCallReverted(string)",
		"28cab9af": "testExecuted(bytes32,uint8,address,uint64,uint64)",
		"36d09269": "testMessageSent(address,uint256,bytes32,uint256,bytes,uint64,bytes,uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"a5c0edf3": "updateAuthVerifier(address)",
		"a66dd384": "updateGasFeePricing(address)",
		"9b11079c": "updateMessageStatus(bytes32,uint8)",
		"d6b457b9": "withdrawGasFees(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061219e806100206000396000f3fe6080604052600436106101965760003560e01c80638da5cb5b116100e1578063aa70fc0e1161008a578063c408733511610064578063c408733514610415578063d6b457b91461042a578063f2fde38b1461044a578063f44d57aa1461046a57610196565b8063aa70fc0e146103cb578063ac8a4c1b146103e0578063affed0e0146103f357610196565b8063a1b058d8116100bb578063a1b058d81461036b578063a5c0edf31461038b578063a66dd384146103ab57610196565b80638da5cb5b146103145780639af1d35a146103365780639b11079c1461034b57610196565b8063485cc95511610143578063715018a61161011d578063715018a6146102d757806372177189146102ec5780638456cb59146102ff57610196565b8063485cc955146102685780635c975abb146102885780635da6d2c4146102aa57610196565b806336d092691161017457806336d09269146102135780633f4ba83a14610233578063446e90451461024857610196565b8063205e157b1461019b57806325b19fa3146101bd57806328cab9af146101f3575b600080fd5b3480156101a757600080fd5b506101bb6101b63660046115a8565b61048a565b005b3480156101c957600080fd5b506101dd6101d8366004611764565b610547565b6040516101ea9190611cb1565b60405180910390f35b3480156101ff57600080fd5b506101bb61020e3660046117a7565b61055f565b34801561021f57600080fd5b506101bb61022e366004611603565b6105b9565b34801561023f57600080fd5b506101bb610624565b34801561025457600080fd5b506101bb61026336600461191a565b610687565b34801561027457600080fd5b506101bb6102833660046115cb565b6106c4565b34801561029457600080fd5b5061029d6107a0565b6040516101ea9190611c42565b3480156102b657600080fd5b506102ca6102c5366004611ab5565b6107a9565b6040516101ea9190611c4d565b3480156102e357600080fd5b506101bb610881565b6101bb6102fa366004611886565b6108e4565b34801561030b57600080fd5b506101bb6108fc565b34801561032057600080fd5b5061032961095d565b6040516101ea9190611bcc565b34801561034257600080fd5b506102ca610979565b34801561035757600080fd5b506101bb61036636600461177c565b61097f565b34801561037757600080fd5b506101bb610386366004611a33565b610a33565b34801561039757600080fd5b506101bb6103a63660046115a8565b610d14565b3480156103b757600080fd5b506101bb6103c63660046115a8565b610de7565b3480156103d757600080fd5b50610329610eba565b6101bb6103ee366004611806565b610ed6565b3480156103ff57600080fd5b50610408610f12565b6040516101ea919061204c565b34801561042157600080fd5b50610329610f3a565b34801561043657600080fd5b506101bb6104453660046115a8565b610f56565b34801561045657600080fd5b506101bb6104653660046115a8565b610fff565b34801561047657600080fd5b506102ca6104853660046116c9565b611097565b6104926110d9565b73ffffffffffffffffffffffffffffffffffffffff166104b061095d565b73ffffffffffffffffffffffffffffffffffffffff16146104ec5760405162461bcd60e51b81526004016104e390611ed0565b60405180910390fd5b600060ca54476104fc9190612079565b60405190915073ffffffffffffffffffffffffffffffffffffffff83169082156108fc029083906000818181858888f19350505050158015610542573d6000803e3d6000fd5b505050565b600081815260fb602052604090205460ff165b919050565b8273ffffffffffffffffffffffffffffffffffffffff16857f04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea658685856040516105aa93929190611cbf565b60405180910390a35050505050565b80888c73ffffffffffffffffffffffffffffffffffffffff167f864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a8d8d8c8c8c8c8c8c60405161060f989796959493929190611fd0565b60405180910390a45050505050505050505050565b61062c6110d9565b73ffffffffffffffffffffffffffffffffffffffff1661064a61095d565b73ffffffffffffffffffffffffffffffffffffffff161461067d5760405162461bcd60e51b81526004016104e390611ed0565b6106856110dd565b565b7fffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f82826040516106b8929190611cef565b60405180910390a15050565b600054610100900460ff166106df5760005460ff16156106e7565b6106e761114b565b6107035760405162461bcd60e51b81526004016104e390611e73565b600054610100900460ff1615801561074b576000805460ff197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b61075361115c565b61075b611193565b610764836111c6565b61076d82610da0565b801561054257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055505050565b60655460ff1690565b60c9546040517f47feadc1000000000000000000000000000000000000000000000000000000008152600091829173ffffffffffffffffffffffffffffffffffffffff909116906347feadc19061080890889088908890600401612029565b602060405180830381600087803b15801561082257600080fd5b505af1158015610836573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085a9190611a1b565b9050806108795760405162461bcd60e51b81526004016104e390611d71565b949350505050565b6108896110d9565b73ffffffffffffffffffffffffffffffffffffffff166108a761095d565b73ffffffffffffffffffffffffffffffffffffffff16146108da5760405162461bcd60e51b81526004016104e390611ed0565b61068560006111ed565b6108f387878787878787611264565b50505050505050565b6109046110d9565b73ffffffffffffffffffffffffffffffffffffffff1661092261095d565b73ffffffffffffffffffffffffffffffffffffffff16146109555760405162461bcd60e51b81526004016104e390611ed0565b610685611437565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b60ca5481565b6109876110d9565b73ffffffffffffffffffffffffffffffffffffffff166109a561095d565b73ffffffffffffffffffffffffffffffffffffffff16146109d85760405162461bcd60e51b81526004016104e390611ed0565b600082815260fb60205260409020805482919060ff19166001836002811115610a2a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055505050565b610a3b6107a0565b15610a585760405162461bcd60e51b81526004016104e390611e3c565b600081815260fb602052604081205460ff166002811115610aa2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14610abf5760405162461bcd60e51b81526004016104e390611f05565b60fa5460405173ffffffffffffffffffffffffffffffffffffffff90911690638b1b3a2d90610af2903390602001611bcc565b6040516020818303038152906040526040518263ffffffff1660e01b8152600401610b1d9190611c9e565b60206040518083038186803b158015610b3557600080fd5b505afa158015610b49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6d9190611744565b5060008673ffffffffffffffffffffffffffffffffffffffff1663a6060871878a8c8888336040518763ffffffff1660e01b8152600401610bb2959493929190611c56565b600060405180830381600088803b158015610bcc57600080fd5b5087f193505050508015610bde575060015b610c5b573d808015610c0c576040519150601f19603f3d011682016040523d82523d6000602084013e610c11565b606091505b507fffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f610c3c82611492565b604051610c499190611c9e565b60405180910390a16002915050610c5f565b5060015b600082815260fb60205260409020805482919060ff19166001836002811115610cb1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055508673ffffffffffffffffffffffffffffffffffffffff16827f04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65838c89604051610d0193929190611cbf565b60405180910390a3505050505050505050565b610d1c6110d9565b73ffffffffffffffffffffffffffffffffffffffff16610d3a61095d565b73ffffffffffffffffffffffffffffffffffffffff1614610d6d5760405162461bcd60e51b81526004016104e390611ed0565b73ffffffffffffffffffffffffffffffffffffffff8116610da05760405162461bcd60e51b81526004016104e390611e05565b60fa80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610def6110d9565b73ffffffffffffffffffffffffffffffffffffffff16610e0d61095d565b73ffffffffffffffffffffffffffffffffffffffff1614610e405760405162461bcd60e51b81526004016104e390611ed0565b73ffffffffffffffffffffffffffffffffffffffff8116610e735760405162461bcd60e51b81526004016104e390611e05565b60c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60c95473ffffffffffffffffffffffffffffffffffffffff1681565b610ede6107a0565b15610efb5760405162461bcd60e51b81526004016104e390611e3c565b610f0a86868686868632611264565b505050505050565b60c95474010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60fa5473ffffffffffffffffffffffffffffffffffffffff1681565b610f5e6110d9565b73ffffffffffffffffffffffffffffffffffffffff16610f7c61095d565b73ffffffffffffffffffffffffffffffffffffffff1614610faf5760405162461bcd60e51b81526004016104e390611ed0565b60ca5460405173ffffffffffffffffffffffffffffffffffffffff83169082156108fc029083906000818181858888f19350505050158015610ff5573d6000803e3d6000fd5b5050600060ca5550565b6110076110d9565b73ffffffffffffffffffffffffffffffffffffffff1661102561095d565b73ffffffffffffffffffffffffffffffffffffffff16146110585760405162461bcd60e51b81526004016104e390611ed0565b73ffffffffffffffffffffffffffffffffffffffff811661108b5760405162461bcd60e51b81526004016104e390611da8565b611094816111ed565b50565b6000878787878787876040516020016110b69796959493929190611bed565b604051602081830303815290604052805190602001209050979650505050505050565b3390565b6110e56107a0565b6111015760405162461bcd60e51b81526004016104e390611d3a565b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6111346110d9565b6040516111419190611bcc565b60405180910390a1565b6000611156306114f8565b15905090565b600054610100900460ff166111835760405162461bcd60e51b81526004016104e390611f73565b61068561118e6110d9565b6111ed565b600054610100900460ff166111ba5760405162461bcd60e51b81526004016104e390611f73565b6065805460ff19169055565b600054610100900460ff16610e735760405162461bcd60e51b81526004016104e390611f73565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061126e611536565b9050808714156112905760405162461bcd60e51b81526004016104e390611d03565b600061129d8886866107a9565b9050803410156112bf5760405162461bcd60e51b81526004016104e390611f3c565b60006112f033848c8c60c960149054906101000a900467ffffffffffffffff1667ffffffffffffffff168d8d611097565b905080893373ffffffffffffffffffffffffffffffffffffffff167f864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a868e8d8d60c960149054906101000a900467ffffffffffffffff168e8e8c60405161135e989796959493929190611fd0565b60405180910390a48160ca60008282546113789190612061565b909155505060c980546014906113af9074010000000000000000000000000000000000000000900467ffffffffffffffff166120c0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508134111561142b5773ffffffffffffffffffffffffffffffffffffffff84166108fc6114018434612079565b6040518115909202916000818181858888f19350505050158015611429573d6000803e3d6000fd5b505b50505050505050505050565b61143f6107a0565b1561145c5760405162461bcd60e51b81526004016104e390611e3c565b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586111346110d9565b60606044825110156114d8575060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c79000000602082015261055a565b600482019150818060200190518101906114f2919061195a565b92915050565b6000808273ffffffffffffffffffffffffffffffffffffffff16803b806020016040519081016040528181526000908060200190933c511192915050565b4690565b60008083601f84011261154b578182fd5b50813567ffffffffffffffff811115611562578182fd5b60208301915083602082850101111561157a57600080fd5b9250929050565b80356003811061055a57600080fd5b803567ffffffffffffffff8116811461055a57600080fd5b6000602082840312156115b9578081fd5b81356115c481612146565b9392505050565b600080604083850312156115dd578081fd5b82356115e881612146565b915060208301356115f881612146565b809150509250929050565b60008060008060008060008060008060006101208c8e031215611624578687fd5b61162e8c35612146565b8b359a5060208c0135995060408c0135985060608c0135975067ffffffffffffffff8060808e01351115611660578788fd5b6116708e60808f01358f0161153a565b909850965061168160a08e01611590565b95508060c08e01351115611693578485fd5b506116a48d60c08e01358e0161153a565b9b9e9a9d50989b979a96999598949794969560e0860135956101000135945092505050565b600080600080600080600060c0888a0312156116e3578283fd5b87356116ee81612146565b96506020880135955060408801359450606088013593506080880135925060a088013567ffffffffffffffff811115611725578283fd5b6117318a828b0161153a565b989b979a50959850939692959293505050565b600060208284031215611755578081fd5b815180151581146115c4578182fd5b600060208284031215611775578081fd5b5035919050565b6000806040838503121561178e578182fd5b8235915061179e60208401611581565b90509250929050565b600080600080600060a086880312156117be578081fd5b853594506117ce60208701611581565b935060408601356117de81612146565b92506117ec60608701611590565b91506117fa60808701611590565b90509295509295909350565b6000806000806000806080878903121561181e578384fd5b8635955060208701359450604087013567ffffffffffffffff80821115611843578586fd5b61184f8a838b0161153a565b90965094506060890135915080821115611867578384fd5b5061187489828a0161153a565b979a9699509497509295939492505050565b600080600080600080600060a0888a0312156118a0578081fd5b8735965060208801359550604088013567ffffffffffffffff808211156118c5578283fd5b6118d18b838c0161153a565b909750955060608a01359150808211156118e9578283fd5b506118f68a828b0161153a565b909450925050608088013561190a81612146565b8091505092959891949750929550565b6000806020838503121561192c578182fd5b823567ffffffffffffffff811115611942578283fd5b61194e8582860161153a565b90969095509350505050565b60006020828403121561196b578081fd5b815167ffffffffffffffff80821115611982578283fd5b818401915084601f830112611995578283fd5b8151818111156119a7576119a7612117565b60405160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011682010181811084821117156119e9576119e9612117565b604052818152838201602001871015611a00578485fd5b611a11826020830160208701612090565b9695505050505050565b600060208284031215611a2c578081fd5b5051919050565b60008060008060008060008060e0898b031215611a4e578182fd5b88359750602089013596506040890135611a6781612146565b9550606089013594506080890135935060a089013567ffffffffffffffff811115611a90578283fd5b611a9c8b828c0161153a565b999c989b50969995989497949560c00135949350505050565b600080600060408486031215611ac9578081fd5b83359250602084013567ffffffffffffffff811115611ae6578182fd5b611af28682870161153a565b9497909650939450505050565b600082845282826020860137806020848601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011685010190509392505050565b60008151808452611b5f816020860160208601612090565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60038110611bc8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b600073ffffffffffffffffffffffffffffffffffffffff8916825287602083015286604083015285606083015284608083015260c060a0830152611c3560c083018486611aff565b9998505050505050505050565b901515815260200190565b90815260200190565b600086825285602083015260806040830152611c76608083018587611aff565b905073ffffffffffffffffffffffffffffffffffffffff831660608301529695505050505050565b6000602082526115c46020830184611b47565b602081016114f28284611b91565b60608101611ccd8286611b91565b67ffffffffffffffff8085166020840152808416604084015250949350505050565b600060208252610879602083018486611aff565b6020808252600f908201527f496e76616c696420636861696e49640000000000000000000000000000000000604082015260600190565b60208082526014908201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604082015260600190565b6020808252600b908201527f466565206e6f7420736574000000000000000000000000000000000000000000604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201527f6464726573730000000000000000000000000000000000000000000000000000606082015260800190565b6020808252600f908201527f43616e6e6f742073657420746f20300000000000000000000000000000000000604082015260600190565b60208082526010908201527f5061757361626c653a2070617573656400000000000000000000000000000000604082015260600190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201527f647920696e697469616c697a6564000000000000000000000000000000000000606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60208082526018908201527f4d65737361676520616c72656164792065786563757465640000000000000000604082015260600190565b60208082526014908201527f496e73756666696369656e742067617320666565000000000000000000000000604082015260600190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201527f6e697469616c697a696e67000000000000000000000000000000000000000000606082015260800190565b600089825288602083015260c06040830152611ff060c08301888a611aff565b67ffffffffffffffff871660608401528281036080840152612013818688611aff565b9150508260a08301529998505050505050505050565b600084825260406020830152612043604083018486611aff565b95945050505050565b67ffffffffffffffff91909116815260200190565b60008219821115612074576120746120e8565b500190565b60008282101561208b5761208b6120e8565b500390565b60005b838110156120ab578181015183820152602001612093565b838111156120ba576000848401525b50505050565b600067ffffffffffffffff808316818114156120de576120de6120e8565b6001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461109457600080fdfea2646970667358221220fef43194b018b6863a536385d79656e3665256ba23ebc2bf6c0f1668f830bae764736f6c63430008000033",
}

// TestMessageBusUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMessageBusUpgradeableMetaData.ABI instead.
var TestMessageBusUpgradeableABI = TestMessageBusUpgradeableMetaData.ABI

// Deprecated: Use TestMessageBusUpgradeableMetaData.Sigs instead.
// TestMessageBusUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var TestMessageBusUpgradeableFuncSigs = TestMessageBusUpgradeableMetaData.Sigs

// TestMessageBusUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMessageBusUpgradeableMetaData.Bin instead.
var TestMessageBusUpgradeableBin = TestMessageBusUpgradeableMetaData.Bin

// DeployTestMessageBusUpgradeable deploys a new Ethereum contract, binding an instance of TestMessageBusUpgradeable to it.
func DeployTestMessageBusUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestMessageBusUpgradeable, error) {
	parsed, err := TestMessageBusUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestMessageBusUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestMessageBusUpgradeable{TestMessageBusUpgradeableCaller: TestMessageBusUpgradeableCaller{contract: contract}, TestMessageBusUpgradeableTransactor: TestMessageBusUpgradeableTransactor{contract: contract}, TestMessageBusUpgradeableFilterer: TestMessageBusUpgradeableFilterer{contract: contract}}, nil
}

// TestMessageBusUpgradeable is an auto generated Go binding around an Ethereum contract.
type TestMessageBusUpgradeable struct {
	TestMessageBusUpgradeableCaller     // Read-only binding to the contract
	TestMessageBusUpgradeableTransactor // Write-only binding to the contract
	TestMessageBusUpgradeableFilterer   // Log filterer for contract events
}

// TestMessageBusUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestMessageBusUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessageBusUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestMessageBusUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessageBusUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestMessageBusUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessageBusUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestMessageBusUpgradeableSession struct {
	Contract     *TestMessageBusUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TestMessageBusUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestMessageBusUpgradeableCallerSession struct {
	Contract *TestMessageBusUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// TestMessageBusUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestMessageBusUpgradeableTransactorSession struct {
	Contract     *TestMessageBusUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// TestMessageBusUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestMessageBusUpgradeableRaw struct {
	Contract *TestMessageBusUpgradeable // Generic contract binding to access the raw methods on
}

// TestMessageBusUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestMessageBusUpgradeableCallerRaw struct {
	Contract *TestMessageBusUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// TestMessageBusUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestMessageBusUpgradeableTransactorRaw struct {
	Contract *TestMessageBusUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestMessageBusUpgradeable creates a new instance of TestMessageBusUpgradeable, bound to a specific deployed contract.
func NewTestMessageBusUpgradeable(address common.Address, backend bind.ContractBackend) (*TestMessageBusUpgradeable, error) {
	contract, err := bindTestMessageBusUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeable{TestMessageBusUpgradeableCaller: TestMessageBusUpgradeableCaller{contract: contract}, TestMessageBusUpgradeableTransactor: TestMessageBusUpgradeableTransactor{contract: contract}, TestMessageBusUpgradeableFilterer: TestMessageBusUpgradeableFilterer{contract: contract}}, nil
}

// NewTestMessageBusUpgradeableCaller creates a new read-only instance of TestMessageBusUpgradeable, bound to a specific deployed contract.
func NewTestMessageBusUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*TestMessageBusUpgradeableCaller, error) {
	contract, err := bindTestMessageBusUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeableCaller{contract: contract}, nil
}

// NewTestMessageBusUpgradeableTransactor creates a new write-only instance of TestMessageBusUpgradeable, bound to a specific deployed contract.
func NewTestMessageBusUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*TestMessageBusUpgradeableTransactor, error) {
	contract, err := bindTestMessageBusUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeableTransactor{contract: contract}, nil
}

// NewTestMessageBusUpgradeableFilterer creates a new log filterer instance of TestMessageBusUpgradeable, bound to a specific deployed contract.
func NewTestMessageBusUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*TestMessageBusUpgradeableFilterer, error) {
	contract, err := bindTestMessageBusUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeableFilterer{contract: contract}, nil
}

// bindTestMessageBusUpgradeable binds a generic wrapper to an already deployed contract.
func bindTestMessageBusUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestMessageBusUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestMessageBusUpgradeable.Contract.TestMessageBusUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TestMessageBusUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TestMessageBusUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestMessageBusUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCaller) AuthVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestMessageBusUpgradeable.contract.Call(opts, &out, "authVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) AuthVerifier() (common.Address, error) {
	return _TestMessageBusUpgradeable.Contract.AuthVerifier(&_TestMessageBusUpgradeable.CallOpts)
}

// AuthVerifier is a free data retrieval call binding the contract method 0xc4087335.
//
// Solidity: function authVerifier() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerSession) AuthVerifier() (common.Address, error) {
	return _TestMessageBusUpgradeable.Contract.AuthVerifier(&_TestMessageBusUpgradeable.CallOpts)
}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCaller) ComputeMessageId(opts *bind.CallOpts, _srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	var out []interface{}
	err := _TestMessageBusUpgradeable.contract.Call(opts, &out, "computeMessageId", _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) ComputeMessageId(_srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	return _TestMessageBusUpgradeable.Contract.ComputeMessageId(&_TestMessageBusUpgradeable.CallOpts, _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)
}

// ComputeMessageId is a free data retrieval call binding the contract method 0xf44d57aa.
//
// Solidity: function computeMessageId(address _srcAddress, uint256 _srcChainId, bytes32 _dstAddress, uint256 _dstChainId, uint256 _srcNonce, bytes _message) pure returns(bytes32)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerSession) ComputeMessageId(_srcAddress common.Address, _srcChainId *big.Int, _dstAddress [32]byte, _dstChainId *big.Int, _srcNonce *big.Int, _message []byte) ([32]byte, error) {
	return _TestMessageBusUpgradeable.Contract.ComputeMessageId(&_TestMessageBusUpgradeable.CallOpts, _srcAddress, _srcChainId, _dstAddress, _dstChainId, _srcNonce, _message)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCaller) Fees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestMessageBusUpgradeable.contract.Call(opts, &out, "fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) Fees() (*big.Int, error) {
	return _TestMessageBusUpgradeable.Contract.Fees(&_TestMessageBusUpgradeable.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerSession) Fees() (*big.Int, error) {
	return _TestMessageBusUpgradeable.Contract.Fees(&_TestMessageBusUpgradeable.CallOpts)
}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCaller) GasFeePricing(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestMessageBusUpgradeable.contract.Call(opts, &out, "gasFeePricing")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) GasFeePricing() (common.Address, error) {
	return _TestMessageBusUpgradeable.Contract.GasFeePricing(&_TestMessageBusUpgradeable.CallOpts)
}

// GasFeePricing is a free data retrieval call binding the contract method 0xaa70fc0e.
//
// Solidity: function gasFeePricing() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerSession) GasFeePricing() (common.Address, error) {
	return _TestMessageBusUpgradeable.Contract.GasFeePricing(&_TestMessageBusUpgradeable.CallOpts)
}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCaller) GetExecutedMessage(opts *bind.CallOpts, _messageId [32]byte) (uint8, error) {
	var out []interface{}
	err := _TestMessageBusUpgradeable.contract.Call(opts, &out, "getExecutedMessage", _messageId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) GetExecutedMessage(_messageId [32]byte) (uint8, error) {
	return _TestMessageBusUpgradeable.Contract.GetExecutedMessage(&_TestMessageBusUpgradeable.CallOpts, _messageId)
}

// GetExecutedMessage is a free data retrieval call binding the contract method 0x25b19fa3.
//
// Solidity: function getExecutedMessage(bytes32 _messageId) view returns(uint8)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerSession) GetExecutedMessage(_messageId [32]byte) (uint8, error) {
	return _TestMessageBusUpgradeable.Contract.GetExecutedMessage(&_TestMessageBusUpgradeable.CallOpts, _messageId)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCaller) Nonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TestMessageBusUpgradeable.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) Nonce() (uint64, error) {
	return _TestMessageBusUpgradeable.Contract.Nonce(&_TestMessageBusUpgradeable.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerSession) Nonce() (uint64, error) {
	return _TestMessageBusUpgradeable.Contract.Nonce(&_TestMessageBusUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestMessageBusUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) Owner() (common.Address, error) {
	return _TestMessageBusUpgradeable.Contract.Owner(&_TestMessageBusUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerSession) Owner() (common.Address, error) {
	return _TestMessageBusUpgradeable.Contract.Owner(&_TestMessageBusUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TestMessageBusUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) Paused() (bool, error) {
	return _TestMessageBusUpgradeable.Contract.Paused(&_TestMessageBusUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableCallerSession) Paused() (bool, error) {
	return _TestMessageBusUpgradeable.Contract.Paused(&_TestMessageBusUpgradeable.CallOpts)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) EstimateFee(opts *bind.TransactOpts, _dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "estimateFee", _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.EstimateFee(&_TestMessageBusUpgradeable.TransactOpts, _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.EstimateFee(&_TestMessageBusUpgradeable.TransactOpts, _dstChainId, _options)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "executeMessage", _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.ExecuteMessage(&_TestMessageBusUpgradeable.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa1b058d8.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes32 _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress [32]byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.ExecuteMessage(&_TestMessageBusUpgradeable.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _gasFeePricing, address _authVerifier) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) Initialize(opts *bind.TransactOpts, _gasFeePricing common.Address, _authVerifier common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "initialize", _gasFeePricing, _authVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _gasFeePricing, address _authVerifier) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) Initialize(_gasFeePricing common.Address, _authVerifier common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.Initialize(&_TestMessageBusUpgradeable.TransactOpts, _gasFeePricing, _authVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _gasFeePricing, address _authVerifier) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) Initialize(_gasFeePricing common.Address, _authVerifier common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.Initialize(&_TestMessageBusUpgradeable.TransactOpts, _gasFeePricing, _authVerifier)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) Pause() (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.Pause(&_TestMessageBusUpgradeable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) Pause() (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.Pause(&_TestMessageBusUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.RenounceOwnership(&_TestMessageBusUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.RenounceOwnership(&_TestMessageBusUpgradeable.TransactOpts)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) RescueGas(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "rescueGas", to)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) RescueGas(to common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.RescueGas(&_TestMessageBusUpgradeable.TransactOpts, to)
}

// RescueGas is a paid mutator transaction binding the contract method 0x205e157b.
//
// Solidity: function rescueGas(address to) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) RescueGas(to common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.RescueGas(&_TestMessageBusUpgradeable.TransactOpts, to)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) SendMessage(opts *bind.TransactOpts, _receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.SendMessage(&_TestMessageBusUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x72177189.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options, address _refundAddress) payable returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte, _refundAddress common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.SendMessage(&_TestMessageBusUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options, _refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) SendMessage0(opts *bind.TransactOpts, _receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "sendMessage0", _receiver, _dstChainId, _message, _options)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) SendMessage0(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.SendMessage0(&_TestMessageBusUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) SendMessage0(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.SendMessage0(&_TestMessageBusUpgradeable.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// TestCallReverted is a paid mutator transaction binding the contract method 0x446e9045.
//
// Solidity: function testCallReverted(string reason) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) TestCallReverted(opts *bind.TransactOpts, reason string) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "testCallReverted", reason)
}

// TestCallReverted is a paid mutator transaction binding the contract method 0x446e9045.
//
// Solidity: function testCallReverted(string reason) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) TestCallReverted(reason string) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TestCallReverted(&_TestMessageBusUpgradeable.TransactOpts, reason)
}

// TestCallReverted is a paid mutator transaction binding the contract method 0x446e9045.
//
// Solidity: function testCallReverted(string reason) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) TestCallReverted(reason string) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TestCallReverted(&_TestMessageBusUpgradeable.TransactOpts, reason)
}

// TestExecuted is a paid mutator transaction binding the contract method 0x28cab9af.
//
// Solidity: function testExecuted(bytes32 messageId, uint8 status, address _dstAddress, uint64 srcChainId, uint64 srcNonce) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) TestExecuted(opts *bind.TransactOpts, messageId [32]byte, status uint8, _dstAddress common.Address, srcChainId uint64, srcNonce uint64) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "testExecuted", messageId, status, _dstAddress, srcChainId, srcNonce)
}

// TestExecuted is a paid mutator transaction binding the contract method 0x28cab9af.
//
// Solidity: function testExecuted(bytes32 messageId, uint8 status, address _dstAddress, uint64 srcChainId, uint64 srcNonce) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) TestExecuted(messageId [32]byte, status uint8, _dstAddress common.Address, srcChainId uint64, srcNonce uint64) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TestExecuted(&_TestMessageBusUpgradeable.TransactOpts, messageId, status, _dstAddress, srcChainId, srcNonce)
}

// TestExecuted is a paid mutator transaction binding the contract method 0x28cab9af.
//
// Solidity: function testExecuted(bytes32 messageId, uint8 status, address _dstAddress, uint64 srcChainId, uint64 srcNonce) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) TestExecuted(messageId [32]byte, status uint8, _dstAddress common.Address, srcChainId uint64, srcNonce uint64) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TestExecuted(&_TestMessageBusUpgradeable.TransactOpts, messageId, status, _dstAddress, srcChainId, srcNonce)
}

// TestMessageSent is a paid mutator transaction binding the contract method 0x36d09269.
//
// Solidity: function testMessageSent(address sender, uint256 srcChainID, bytes32 receiver, uint256 dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 messageId) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) TestMessageSent(opts *bind.TransactOpts, sender common.Address, srcChainID *big.Int, receiver [32]byte, dstChainId *big.Int, message []byte, nonce uint64, options []byte, fee *big.Int, messageId [32]byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "testMessageSent", sender, srcChainID, receiver, dstChainId, message, nonce, options, fee, messageId)
}

// TestMessageSent is a paid mutator transaction binding the contract method 0x36d09269.
//
// Solidity: function testMessageSent(address sender, uint256 srcChainID, bytes32 receiver, uint256 dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 messageId) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) TestMessageSent(sender common.Address, srcChainID *big.Int, receiver [32]byte, dstChainId *big.Int, message []byte, nonce uint64, options []byte, fee *big.Int, messageId [32]byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TestMessageSent(&_TestMessageBusUpgradeable.TransactOpts, sender, srcChainID, receiver, dstChainId, message, nonce, options, fee, messageId)
}

// TestMessageSent is a paid mutator transaction binding the contract method 0x36d09269.
//
// Solidity: function testMessageSent(address sender, uint256 srcChainID, bytes32 receiver, uint256 dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 messageId) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) TestMessageSent(sender common.Address, srcChainID *big.Int, receiver [32]byte, dstChainId *big.Int, message []byte, nonce uint64, options []byte, fee *big.Int, messageId [32]byte) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TestMessageSent(&_TestMessageBusUpgradeable.TransactOpts, sender, srcChainID, receiver, dstChainId, message, nonce, options, fee, messageId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TransferOwnership(&_TestMessageBusUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.TransferOwnership(&_TestMessageBusUpgradeable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) Unpause() (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.Unpause(&_TestMessageBusUpgradeable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) Unpause() (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.Unpause(&_TestMessageBusUpgradeable.TransactOpts)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) UpdateAuthVerifier(opts *bind.TransactOpts, _authVerifier common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "updateAuthVerifier", _authVerifier)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) UpdateAuthVerifier(_authVerifier common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.UpdateAuthVerifier(&_TestMessageBusUpgradeable.TransactOpts, _authVerifier)
}

// UpdateAuthVerifier is a paid mutator transaction binding the contract method 0xa5c0edf3.
//
// Solidity: function updateAuthVerifier(address _authVerifier) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) UpdateAuthVerifier(_authVerifier common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.UpdateAuthVerifier(&_TestMessageBusUpgradeable.TransactOpts, _authVerifier)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) UpdateGasFeePricing(opts *bind.TransactOpts, _gasFeePricing common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "updateGasFeePricing", _gasFeePricing)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) UpdateGasFeePricing(_gasFeePricing common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.UpdateGasFeePricing(&_TestMessageBusUpgradeable.TransactOpts, _gasFeePricing)
}

// UpdateGasFeePricing is a paid mutator transaction binding the contract method 0xa66dd384.
//
// Solidity: function updateGasFeePricing(address _gasFeePricing) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) UpdateGasFeePricing(_gasFeePricing common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.UpdateGasFeePricing(&_TestMessageBusUpgradeable.TransactOpts, _gasFeePricing)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) UpdateMessageStatus(opts *bind.TransactOpts, _messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "updateMessageStatus", _messageId, _status)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) UpdateMessageStatus(_messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.UpdateMessageStatus(&_TestMessageBusUpgradeable.TransactOpts, _messageId, _status)
}

// UpdateMessageStatus is a paid mutator transaction binding the contract method 0x9b11079c.
//
// Solidity: function updateMessageStatus(bytes32 _messageId, uint8 _status) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) UpdateMessageStatus(_messageId [32]byte, _status uint8) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.UpdateMessageStatus(&_TestMessageBusUpgradeable.TransactOpts, _messageId, _status)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactor) WithdrawGasFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.contract.Transact(opts, "withdrawGasFees", to)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableSession) WithdrawGasFees(to common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.WithdrawGasFees(&_TestMessageBusUpgradeable.TransactOpts, to)
}

// WithdrawGasFees is a paid mutator transaction binding the contract method 0xd6b457b9.
//
// Solidity: function withdrawGasFees(address to) returns()
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableTransactorSession) WithdrawGasFees(to common.Address) (*types.Transaction, error) {
	return _TestMessageBusUpgradeable.Contract.WithdrawGasFees(&_TestMessageBusUpgradeable.TransactOpts, to)
}

// TestMessageBusUpgradeableCallRevertedIterator is returned from FilterCallReverted and is used to iterate over the raw logs and unpacked data for CallReverted events raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableCallRevertedIterator struct {
	Event *TestMessageBusUpgradeableCallReverted // Event containing the contract specifics and raw log

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
func (it *TestMessageBusUpgradeableCallRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessageBusUpgradeableCallReverted)
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
		it.Event = new(TestMessageBusUpgradeableCallReverted)
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
func (it *TestMessageBusUpgradeableCallRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessageBusUpgradeableCallRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessageBusUpgradeableCallReverted represents a CallReverted event raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableCallReverted struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallReverted is a free log retrieval operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) FilterCallReverted(opts *bind.FilterOpts) (*TestMessageBusUpgradeableCallRevertedIterator, error) {

	logs, sub, err := _TestMessageBusUpgradeable.contract.FilterLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeableCallRevertedIterator{contract: _TestMessageBusUpgradeable.contract, event: "CallReverted", logs: logs, sub: sub}, nil
}

// WatchCallReverted is a free log subscription operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) WatchCallReverted(opts *bind.WatchOpts, sink chan<- *TestMessageBusUpgradeableCallReverted) (event.Subscription, error) {

	logs, sub, err := _TestMessageBusUpgradeable.contract.WatchLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessageBusUpgradeableCallReverted)
				if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "CallReverted", log); err != nil {
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

// ParseCallReverted is a log parse operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) ParseCallReverted(log types.Log) (*TestMessageBusUpgradeableCallReverted, error) {
	event := new(TestMessageBusUpgradeableCallReverted)
	if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "CallReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessageBusUpgradeableExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableExecutedIterator struct {
	Event *TestMessageBusUpgradeableExecuted // Event containing the contract specifics and raw log

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
func (it *TestMessageBusUpgradeableExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessageBusUpgradeableExecuted)
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
		it.Event = new(TestMessageBusUpgradeableExecuted)
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
func (it *TestMessageBusUpgradeableExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessageBusUpgradeableExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessageBusUpgradeableExecuted represents a Executed event raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableExecuted struct {
	MessageId  [32]byte
	Status     uint8
	DstAddress common.Address
	SrcChainId uint64
	SrcNonce   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) FilterExecuted(opts *bind.FilterOpts, messageId [][32]byte, _dstAddress []common.Address) (*TestMessageBusUpgradeableExecutedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	var _dstAddressRule []interface{}
	for _, _dstAddressItem := range _dstAddress {
		_dstAddressRule = append(_dstAddressRule, _dstAddressItem)
	}

	logs, sub, err := _TestMessageBusUpgradeable.contract.FilterLogs(opts, "Executed", messageIdRule, _dstAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeableExecutedIterator{contract: _TestMessageBusUpgradeable.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *TestMessageBusUpgradeableExecuted, messageId [][32]byte, _dstAddress []common.Address) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	var _dstAddressRule []interface{}
	for _, _dstAddressItem := range _dstAddress {
		_dstAddressRule = append(_dstAddressRule, _dstAddressItem)
	}

	logs, sub, err := _TestMessageBusUpgradeable.contract.WatchLogs(opts, "Executed", messageIdRule, _dstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessageBusUpgradeableExecuted)
				if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x04214a849019ea3548afcedee810b5bc1680cfb64e22fdf9051a823f3cdfea65.
//
// Solidity: event Executed(bytes32 indexed messageId, uint8 status, address indexed _dstAddress, uint64 srcChainId, uint64 srcNonce)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) ParseExecuted(log types.Log) (*TestMessageBusUpgradeableExecuted, error) {
	event := new(TestMessageBusUpgradeableExecuted)
	if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessageBusUpgradeableMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableMessageSentIterator struct {
	Event *TestMessageBusUpgradeableMessageSent // Event containing the contract specifics and raw log

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
func (it *TestMessageBusUpgradeableMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessageBusUpgradeableMessageSent)
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
		it.Event = new(TestMessageBusUpgradeableMessageSent)
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
func (it *TestMessageBusUpgradeableMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessageBusUpgradeableMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessageBusUpgradeableMessageSent represents a MessageSent event raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableMessageSent struct {
	Sender     common.Address
	SrcChainID *big.Int
	Receiver   [32]byte
	DstChainId *big.Int
	Message    []byte
	Nonce      uint64
	Options    []byte
	Fee        *big.Int
	MessageId  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) FilterMessageSent(opts *bind.FilterOpts, sender []common.Address, dstChainId []*big.Int, messageId [][32]byte) (*TestMessageBusUpgradeableMessageSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _TestMessageBusUpgradeable.contract.FilterLogs(opts, "MessageSent", senderRule, dstChainIdRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeableMessageSentIterator{contract: _TestMessageBusUpgradeable.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *TestMessageBusUpgradeableMessageSent, sender []common.Address, dstChainId []*big.Int, messageId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _TestMessageBusUpgradeable.contract.WatchLogs(opts, "MessageSent", senderRule, dstChainIdRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessageBusUpgradeableMessageSent)
				if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x864ad5e86ed3626c9517260fbfe1eed395157fd938e459e9fb607a07129cdd2a.
//
// Solidity: event MessageSent(address indexed sender, uint256 srcChainID, bytes32 receiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes options, uint256 fee, bytes32 indexed messageId)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) ParseMessageSent(log types.Log) (*TestMessageBusUpgradeableMessageSent, error) {
	event := new(TestMessageBusUpgradeableMessageSent)
	if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessageBusUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableOwnershipTransferredIterator struct {
	Event *TestMessageBusUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestMessageBusUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessageBusUpgradeableOwnershipTransferred)
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
		it.Event = new(TestMessageBusUpgradeableOwnershipTransferred)
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
func (it *TestMessageBusUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessageBusUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessageBusUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestMessageBusUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestMessageBusUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeableOwnershipTransferredIterator{contract: _TestMessageBusUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestMessageBusUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestMessageBusUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessageBusUpgradeableOwnershipTransferred)
				if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*TestMessageBusUpgradeableOwnershipTransferred, error) {
	event := new(TestMessageBusUpgradeableOwnershipTransferred)
	if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessageBusUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeablePausedIterator struct {
	Event *TestMessageBusUpgradeablePaused // Event containing the contract specifics and raw log

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
func (it *TestMessageBusUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessageBusUpgradeablePaused)
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
		it.Event = new(TestMessageBusUpgradeablePaused)
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
func (it *TestMessageBusUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessageBusUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessageBusUpgradeablePaused represents a Paused event raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*TestMessageBusUpgradeablePausedIterator, error) {

	logs, sub, err := _TestMessageBusUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeablePausedIterator{contract: _TestMessageBusUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TestMessageBusUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _TestMessageBusUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessageBusUpgradeablePaused)
				if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) ParsePaused(log types.Log) (*TestMessageBusUpgradeablePaused, error) {
	event := new(TestMessageBusUpgradeablePaused)
	if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessageBusUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableUnpausedIterator struct {
	Event *TestMessageBusUpgradeableUnpaused // Event containing the contract specifics and raw log

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
func (it *TestMessageBusUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessageBusUpgradeableUnpaused)
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
		it.Event = new(TestMessageBusUpgradeableUnpaused)
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
func (it *TestMessageBusUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessageBusUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessageBusUpgradeableUnpaused represents a Unpaused event raised by the TestMessageBusUpgradeable contract.
type TestMessageBusUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TestMessageBusUpgradeableUnpausedIterator, error) {

	logs, sub, err := _TestMessageBusUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TestMessageBusUpgradeableUnpausedIterator{contract: _TestMessageBusUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TestMessageBusUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _TestMessageBusUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessageBusUpgradeableUnpaused)
				if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TestMessageBusUpgradeable *TestMessageBusUpgradeableFilterer) ParseUnpaused(log types.Log) (*TestMessageBusUpgradeableUnpaused, error) {
	event := new(TestMessageBusUpgradeableUnpaused)
	if err := _TestMessageBusUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
