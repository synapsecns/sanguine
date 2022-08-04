// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package homeharness

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209fc2f9faef3fb94a246517c0df5c36857c54ee531ec785b1e2a8e7467d886fbd64736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122011611066541d386171a021fffbfa00683b2bd43bbd4a7ac45b4c5270eeba03e464736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209c053342380fec451d5f021115572c513d43b81e9452f7d76038f307ad3911d964736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ff9b3eb0991adb4f58c4275dcf0b1e31b504715750a7aba366bd2c467d46690264736f6c634300080d0033",
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

// AuthManagerMetaData contains all meta data concerning the AuthManager contract.
var AuthManagerMetaData = &bind.MetaData{
	ABI: "[]",
}

// AuthManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthManagerMetaData.ABI instead.
var AuthManagerABI = AuthManagerMetaData.ABI

// AuthManager is an auto generated Go binding around an Ethereum contract.
type AuthManager struct {
	AuthManagerCaller     // Read-only binding to the contract
	AuthManagerTransactor // Write-only binding to the contract
	AuthManagerFilterer   // Log filterer for contract events
}

// AuthManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthManagerSession struct {
	Contract     *AuthManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthManagerCallerSession struct {
	Contract *AuthManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AuthManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthManagerTransactorSession struct {
	Contract     *AuthManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AuthManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthManagerRaw struct {
	Contract *AuthManager // Generic contract binding to access the raw methods on
}

// AuthManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthManagerCallerRaw struct {
	Contract *AuthManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AuthManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthManagerTransactorRaw struct {
	Contract *AuthManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthManager creates a new instance of AuthManager, bound to a specific deployed contract.
func NewAuthManager(address common.Address, backend bind.ContractBackend) (*AuthManager, error) {
	contract, err := bindAuthManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthManager{AuthManagerCaller: AuthManagerCaller{contract: contract}, AuthManagerTransactor: AuthManagerTransactor{contract: contract}, AuthManagerFilterer: AuthManagerFilterer{contract: contract}}, nil
}

// NewAuthManagerCaller creates a new read-only instance of AuthManager, bound to a specific deployed contract.
func NewAuthManagerCaller(address common.Address, caller bind.ContractCaller) (*AuthManagerCaller, error) {
	contract, err := bindAuthManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthManagerCaller{contract: contract}, nil
}

// NewAuthManagerTransactor creates a new write-only instance of AuthManager, bound to a specific deployed contract.
func NewAuthManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthManagerTransactor, error) {
	contract, err := bindAuthManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthManagerTransactor{contract: contract}, nil
}

// NewAuthManagerFilterer creates a new log filterer instance of AuthManager, bound to a specific deployed contract.
func NewAuthManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthManagerFilterer, error) {
	contract, err := bindAuthManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthManagerFilterer{contract: contract}, nil
}

// bindAuthManager binds a generic wrapper to an already deployed contract.
func bindAuthManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthManager *AuthManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuthManager.Contract.AuthManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthManager *AuthManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthManager.Contract.AuthManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthManager *AuthManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthManager.Contract.AuthManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthManager *AuthManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuthManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthManager *AuthManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthManager *AuthManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthManager.Contract.contract.Transact(opts, method, params...)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e5d97caf9d27974cddcb30b6d6cbf355faffdccbe58d5ffbb2c493f58bc2398264736f6c634300080d0033",
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

// HeaderMetaData contains all meta data concerning the Header contract.
var HeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206fe09d22d7383b608c334c54048540d82063186c179efc909f95aeaf37de62e164736f6c634300080d0033",
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

// HomeMetaData contains all meta data concerning the Home contract.
var HomeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destinationAndNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tips\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"ImproperAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updaterManager\",\"type\":\"address\"}],\"name\":\"NewUpdaterManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"UpdaterSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_BODY_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historicalRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_attestation\",\"type\":\"bytes\"}],\"name\":\"improperAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUpdaterManager\",\"name\":\"_updaterManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"_systemMessenger\",\"type\":\"address\"}],\"name\":\"setSystemMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updaterManager\",\"type\":\"address\"}],\"name\":\"setUpdaterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumHome.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestUpdate\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemMessenger\",\"outputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updaterManager\",\"outputs\":[{\"internalType\":\"contractIUpdaterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"522ae002": "MAX_MESSAGE_BODY_BYTES()",
		"ffa1ad74": "VERSION()",
		"06661abd": "count()",
		"f7560e40": "dispatch(uint32,bytes32,uint32,bytes,bytes)",
		"7ea97f40": "historicalRoots(uint256)",
		"45630b1a": "homeDomainHash()",
		"88a278ec": "improperAttestation(address,bytes)",
		"c4d66de8": "initialize(address)",
		"8d3638f4": "localDomain()",
		"affed0e0": "nonce()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"ebf0c717": "root()",
		"b7bc563e": "setSystemMessenger(address)",
		"9d54f419": "setUpdater(address)",
		"9776120e": "setUpdaterManager(address)",
		"c19d93fb": "state()",
		"36e104de": "suggestUpdate()",
		"ccbdf9c9": "systemMessenger()",
		"f2fde38b": "transferOwnership(address)",
		"fd54b228": "tree()",
		"df034cd0": "updater()",
		"9df6c8e1": "updaterManager()",
	},
	Bin: "0x60a06040523480156200001157600080fd5b506040516200382138038062003821833981016040819052620000349162000043565b63ffffffff1660805262000072565b6000602082840312156200005657600080fd5b815163ffffffff811681146200006b57600080fd5b9392505050565b60805161377e620000a36000396000818161027e0152818161054701528181610de50152611886015261377e6000f3fe6080604052600436106101805760003560e01c80639df6c8e1116100d6578063df034cd01161007f578063f7560e4011610059578063f7560e40146104a8578063fd54b228146104bb578063ffa1ad74146104d257600080fd5b8063df034cd014610446578063ebf0c71714610473578063f2fde38b1461048857600080fd5b8063c19d93fb116100b0578063c19d93fb146103b5578063c4d66de8146103f9578063ccbdf9c91461041957600080fd5b80639df6c8e114610341578063affed0e014610377578063b7bc563e1461039557600080fd5b80637ea97f40116101385780638da5cb5b116101125780638da5cb5b146102b55780639776120e146103015780639d54f4191461032157600080fd5b80637ea97f401461021c57806388a278ec1461023c5780638d3638f41461026c57600080fd5b806345630b1a1161016957806345630b1a146101da578063522ae002146101ef578063715018a61461020557600080fd5b806306661abd1461018557806336e104de146101a9575b600080fd5b34801561019157600080fd5b506020545b6040519081526020015b60405180910390f35b3480156101b557600080fd5b506101be6104f9565b6040805163ffffffff90931683526020830191909152016101a0565b3480156101e657600080fd5b50610196610540565b3480156101fb57600080fd5b5061019661080081565b34801561021157600080fd5b5061021a610570565b005b34801561022857600080fd5b50610196610237366004612f46565b6105de565b34801561024857600080fd5b5061025c61025736600461305b565b6105ff565b60405190151581526020016101a0565b34801561027857600080fd5b506102a07f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101a0565b3480156102c157600080fd5b5060855473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a0565b34801561030d57600080fd5b5061021a61031c3660046130ab565b61074e565b34801561032d57600080fd5b5061021a61033c3660046130ab565b6107c1565b34801561034d57600080fd5b5061011b546102dc90640100000000900473ffffffffffffffffffffffffffffffffffffffff1681565b34801561038357600080fd5b5061011b546102a09063ffffffff1681565b3480156103a157600080fd5b5061021a6103b03660046130ab565b610881565b3480156103c157600080fd5b5061011b546103ec907801000000000000000000000000000000000000000000000000900460ff1681565b6040516101a091906130f7565b34801561040557600080fd5b5061021a6104143660046130ab565b61092f565b34801561042557600080fd5b5060b8546102dc9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561045257600080fd5b5060b7546102dc9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561047f57600080fd5b50610196610af1565b34801561049457600080fd5b5061021a6104a33660046130ab565b610afd565b61021a6104b636600461314c565b610bf6565b3480156104c757600080fd5b506020546101969081565b3480156104de57600080fd5b506104e7600081565b60405160ff90911681526020016101a0565b6021546000908190801561053b5761051260018261320a565b925060218363ffffffff168154811061052d5761052d613221565b906000526020600020015491505b509091565b600061056b7f0000000000000000000000000000000000000000000000000000000000000000610ef0565b905090565b60855473ffffffffffffffffffffffffffffffffffffffff1633146105dc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b602181815481106105ee57600080fd5b600091825260209091200154905081565b6000600261011b547801000000000000000000000000000000000000000000000000900460ff166002811115610637576106376130c8565b036106845760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064016105d3565b60006106908484610f69565b905060006106a362ffffff198316611072565b905060006106b662ffffff198416611086565b60215490915063ffffffff831610156106ff5760218263ffffffff16815481106106e2576106e2613221565b906000526020600020015481036106ff5760009350505050610748565b61070761109b565b7f287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b86866040516107389291906132c6565b60405180910390a1600193505050505b92915050565b60855473ffffffffffffffffffffffffffffffffffffffff1633146107b55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105d3565b6107be816111b6565b50565b61011b54640100000000900473ffffffffffffffffffffffffffffffffffffffff1633146108315760405162461bcd60e51b815260206004820152600f60248201527f21757064617465724d616e61676572000000000000000000000000000000000060448201526064016105d3565b61083a8161129e565b5061011b80547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000179055565b60855473ffffffffffffffffffffffffffffffffffffffff1633146108e85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105d3565b60b880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600061093b600161131d565b9050801561097057605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610979826111b6565b610a1361011b60049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663df034cd06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109ea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0e91906132f5565b611476565b61011b80547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff1678010000000000000000000000000000000000000000000000001790556021805460018101825560009182527f3a6357012c1a3ae0a17d304c9920310382d968ebcc4b1771f41c6b304205b57001558015610aed57605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b600061056b6000611504565b60855473ffffffffffffffffffffffffffffffffffffffff163314610b645760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105d3565b73ffffffffffffffffffffffffffffffffffffffff8116610bed5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105d3565b6107be81611517565b600261011b547801000000000000000000000000000000000000000000000000900460ff166002811115610c2c57610c2c6130c8565b03610c795760405162461bcd60e51b815260206004820152600c60248201527f6661696c6564207374617465000000000000000000000000000000000000000060448201526064016105d3565b61080081511115610ccc5760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e67000000000000000000000000000000000000000060448201526064016105d3565b34610ce4610cd98461158e565b62ffffff191661159b565b6bffffffffffffffffffffffff1614610d3f5760405162461bcd60e51b815260206004820152600560248201527f217469707300000000000000000000000000000000000000000000000000000060448201526064016105d3565b61011b54610d549063ffffffff166001613312565b61011b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff929092169190911790556000610d95856115fd565b61011b54604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060e090811b821660228401526026830186905293841b811660468301528a841b8116604a830152604e82018a90529288901b909216606e83015280518083036052018152607290920190529091506000610e5f82868661165c565b80516020820120909150610e72816116d7565b61011b5463ffffffff1660208a901b67ffffffff00000000161767ffffffffffffffff166001610ea160205490565b610eab919061320a565b827f718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e7858986604051610edd92919061333a565b60405180910390a4505050505050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e083901b1660208201527f53594e000000000000000000000000000000000000000000000000000000000060248201526000906027015b604051602081830303815290604052805190602001209050919050565b6000610f758282611707565b905060286bffffffffffffffffffffffff601883901c1611610fd95760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e000000000000000000000000000060448201526064016105d3565b61100e83610fec62ffffff19841661172b565b611009610ffe62ffffff198616611740565b62ffffff1916611773565b6117c6565b61102661102062ffffff19831661186e565b84611882565b6107485760405162461bcd60e51b815260206004820152601860248201527f5369676e6572206973206e6f7420616e2075706461746572000000000000000060448201526064016105d3565b600061074862ffffff198316600480611923565b600061074862ffffff19831660086020611953565b61011b805478020000000000000000000000000000000000000000000000007fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff90911617908190556040517f5b3c2cbf00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff6401000000009092049190911690635b3c2cbf90602401600060405180830381600087803b15801561115557600080fd5b505af1158015611169573d6000803e3d6000fd5b505060b75460405133935073ffffffffffffffffffffffffffffffffffffffff90911691507f98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce90600090a3565b73ffffffffffffffffffffffffffffffffffffffff81163b61121a5760405162461bcd60e51b815260206004820152601860248201527f21636f6e747261637420757064617465724d616e61676572000000000000000060448201526064016105d3565b61011b80547fffffffffffffffff0000000000000000000000000000000000000000ffffffff1664010000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf9060200160405180910390a150565b60b7805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a9101610ae4565b605254600090610100900460ff16156113bc578160ff1660011480156113425750303b155b6113b45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105d3565b506000919050565b60525460ff8084169116106114395760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105d3565b50605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b605254610100900460ff166114f35760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105d3565b6114fb611b11565b6107be8161129e565b600061074882611512611b96565b612057565b6085805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610748826002611707565b6000816115b160025b62ffffff1983169061211a565b506115bb8361221b565b6115c484612249565b6115cd8561226a565b6115d68661228b565b6115e0919061335f565b6115ea919061335f565b6115f4919061335f565b91505b50919050565b60007fffffffffffffffffffffffff0000000000000000000000000000000000000000821461162d573392915050565b6116356122ac565b507fffffffffffffffffffffffff0000000000000000000000000000000000000000919050565b825160609060009061167060046002613386565b60ff1661167d91906133af565b9050600084518261168e91906133af565b9050600161169e60046002613386565b60ff1683838989896040516020016116bc97969594939291906133cc565b604051602081830303815290604052925050505b9392505050565b6116e2600082612313565b60216116ee6000611504565b8154600181018355600092835260209092209091015550565b81516000906020840161172264ffffffffff85168284612436565b95945050505050565b600061074862ffffff1983168260288161247d565b6000610748602861176381601886901c6bffffffffffffffffffffffff1661320a565b62ffffff1985169190600061247d565b60606000806117908460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060405191508192506117b58483602001612503565b508181016020016040529052919050565b60006117d762ffffff1984166126a8565b90506117e281612705565b90508373ffffffffffffffffffffffffffffffffffffffff166118058284612740565b73ffffffffffffffffffffffffffffffffffffffff16146118685760405162461bcd60e51b815260206004820152601160248201527f496e76616c6964207369676e617475726500000000000000000000000000000060448201526064016105d3565b50505050565b600061074862ffffff198316826004611923565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff168363ffffffff16146118ff5760405162461bcd60e51b815260206004820152600c60248201527f57726f6e6720646f6d61696e000000000000000000000000000000000000000060448201526064016105d3565b5060b75473ffffffffffffffffffffffffffffffffffffffff908116911614919050565b600061193082602061346a565b61193b906008613386565b60ff16611949858585611953565b901c949350505050565b60008160ff16600003611968575060006116d0565b6119808460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1661199b60ff84168561348d565b1115611a13576119fa6119bc8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166119e28660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16612764565b60405162461bcd60e51b81526004016105d391906134a5565b60208260ff161115611a8d5760405162461bcd60e51b815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e20333220627974657300000000000060648201526084016105d3565b600882026000611aab8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b605254610100900460ff16611b8e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105d3565b6105dc6127d2565b611b9e612f27565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b602081101561211257600182821c8116908190036120be5785826020811061208b5761208b613221565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350612109565b838583602081106120d1576120d1613221565b60200201516040516020016120f0929190918252602082015260400190565b6040516020818303038152906040528051906020012093505b50600101612061565b505092915050565b60006121268383612858565b6122145760006121456121398560d81c90565b64ffffffffff1661287b565b915050600061215a8464ffffffffff1661287b565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e0160405160208183030381529060405290508060405162461bcd60e51b81526004016105d391906134a5565b5090919050565b60008161222860026115a4565b5061223c62ffffff1984166026600c611923565b63ffffffff169392505050565b60008161225660026115a4565b5061223c62ffffff198416601a600c611923565b60008161227760026115a4565b5061223c62ffffff198416600e600c611923565b60008161229860026115a4565b5061223c62ffffff1984166002600c611923565b60b85473ffffffffffffffffffffffffffffffffffffffff1633146105dc5760405162461bcd60e51b815260206004820152601060248201527f2173797374656d4d657373656e6765720000000000000000000000000000000060448201526064016105d3565b602080830154906001906123289060026135d0565b612332919061320a565b81106123805760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c0000000000000000000000000000000060448201526064016105d3565b6001016020830181905560005b602081101561242857816001166001036123bc57828482602081106123b4576123b4613221565b015550505050565b8381602081106123ce576123ce613221565b01546040805160208101929092528101849052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209250600191821c910161238d565b506124316135dc565b505050565b600080612443838561348d565b9050604051811115612453575060005b806000036124685762ffffff199150506116d0565b5050606092831b9190911790911b1760181b90565b6000806124988660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506124b186612965565b846124bc878461348d565b6124c6919061348d565b11156124d95762ffffff199150506124fb565b6124e3858261348d565b90506124f78364ffffffffff168286612436565b9150505b949350505050565b600062ffffff19808416036125805760405162461bcd60e51b815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f657220646572656600000000000000000000000000000000000000000000000060648201526084016105d3565b612589836129ad565b6125fb5760405162461bcd60e51b815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e74657220646572656600000000000000000000000000000000000000000060648201526084016105d3565b60006126158460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061263f8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006040519050848111156126645760206060fd5b8285848460045afa5061269e61267a8760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b9695505050505050565b6000806126c38360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006126ed8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01610f4c565b600080600061274f85856129ea565b9150915061275c81612a58565b509392505050565b606060006127718661287b565b915050600061277f8661287b565b915050600061278d8661287b565b915050600061279b8661287b565b915050838383836040516020016127b5949392919061360b565b604051602081830303815290604052945050505050949350505050565b605254610100900460ff1661284f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105d3565b6105dc33611517565b60008164ffffffffff1661286c8460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff1611156128ee57600061289a826008613386565b60ff1685901c90506128ab81612c44565b61ffff16841793508160ff166010146128c657601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612881565b50600f5b60ff8160ff16101561295f57600061290b826008613386565b60ff1685901c905061291c81612c44565b61ffff16831792508160ff1660001461293757601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016128f2565b50915091565b600061297f8260181c6bffffffffffffffffffffffff1690565b6129978360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b60006129b98260d81c90565b64ffffffffff1664ffffffffff036129d357506000919050565b60006129de83612965565b60405110199392505050565b6000808251604103612a205760208301516040840151606085015160001a612a1487828585612c76565b94509450505050612a51565b8251604003612a495760208301516040840151612a3e868383612d8e565b935093505050612a51565b506000905060025b9250929050565b6000816004811115612a6c57612a6c6130c8565b03612a745750565b6001816004811115612a8857612a886130c8565b03612ad55760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016105d3565b6002816004811115612ae957612ae96130c8565b03612b365760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016105d3565b6003816004811115612b4a57612b4a6130c8565b03612bbd5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016105d3565b6004816004811115612bd157612bd16130c8565b036107be5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016105d3565b6000612c5660048360ff16901c612de0565b60ff1661ffff919091161760081b612c6d82612de0565b60ff1617919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612cad5750600090506003612d85565b8460ff16601b14158015612cc557508460ff16601c14155b15612cd65750600090506004612d85565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612d2a573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612d7e57600060019250925050612d85565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831681612dc460ff86901c601b61348d565b9050612dd287828885612c76565b935093505050935093915050565b600060f08083179060ff82169003612dfb5750603092915050565b8060ff1660f103612e0f5750603192915050565b8060ff1660f203612e235750603292915050565b8060ff1660f303612e375750603392915050565b8060ff1660f403612e4b5750603492915050565b8060ff1660f503612e5f5750603592915050565b8060ff1660f603612e735750603692915050565b8060ff1660f703612e875750603792915050565b8060ff1660f803612e9b5750603892915050565b8060ff1660f903612eaf5750603992915050565b8060ff1660fa03612ec35750606192915050565b8060ff1660fb03612ed75750606292915050565b8060ff1660fc03612eeb5750606392915050565b8060ff1660fd03612eff5750606492915050565b8060ff1660fe03612f135750606592915050565b8060ff1660ff036115f75750606692915050565b6040518061040001604052806020906020820280368337509192915050565b600060208284031215612f5857600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff811681146107be57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112612fc157600080fd5b813567ffffffffffffffff80821115612fdc57612fdc612f81565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561302257613022612f81565b8160405283815286602085880101111561303b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561306e57600080fd5b823561307981612f5f565b9150602083013567ffffffffffffffff81111561309557600080fd5b6130a185828601612fb0565b9150509250929050565b6000602082840312156130bd57600080fd5b81356116d081612f5f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160038310613132577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b803563ffffffff8116811461147157600080fd5b600080600080600060a0868803121561316457600080fd5b61316d86613138565b94506020860135935061318260408701613138565b9250606086013567ffffffffffffffff8082111561319f57600080fd5b6131ab89838a01612fb0565b935060808801359150808211156131c157600080fd5b506131ce88828901612fb0565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561321c5761321c6131db565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b8381101561326b578181015183820152602001613253565b838111156118685750506000910152565b60008151808452613294816020860160208601613250565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006124fb604083018461327c565b60006020828403121561330757600080fd5b81516116d081612f5f565b600063ffffffff808316818516808303821115613331576133316131db565b01949350505050565b60408152600061334d604083018561327c565b8281036020840152611722818561327c565b60006bffffffffffffffffffffffff808316818516808303821115613331576133316131db565b600060ff821660ff84168160ff04811182151516156133a7576133a76131db565b029392505050565b600061ffff808316818516808303821115613331576133316131db565b60007fffff000000000000000000000000000000000000000000000000000000000000808a60f01b168352808960f01b166002840152808860f01b166004840152808760f01b16600684015250845161342c816008850160208901613250565b845190830190613443816008840160208901613250565b8451910190613459816008840160208801613250565b016008019998505050505050505050565b600060ff821660ff841680821015613484576134846131db565b90039392505050565b600082198211156134a0576134a06131db565b500190565b6020815260006116d0602083018461327c565b600181815b8085111561351157817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156134f7576134f76131db565b8085161561350457918102915b93841c93908002906134bd565b509250929050565b60008261352857506001610748565b8161353557506000610748565b816001811461354b576002811461355557613571565b6001915050610748565b60ff841115613566576135666131db565b50506001821b610748565b5060208310610133831016604e8410600b8410161715613594575081810a610748565b61359e83836134b8565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156133a7576133a76131db565b60006116d08383613519565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d820161269e56fea2646970667358221220e13180e8453e418f74764fbefc8999445722ec46f2bb1d23fc30c7c5da0087a764736f6c634300080d0033",
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

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_Home *HomeCaller) HistoricalRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "historicalRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_Home *HomeSession) HistoricalRoots(arg0 *big.Int) ([32]byte, error) {
	return _Home.Contract.HistoricalRoots(&_Home.CallOpts, arg0)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_Home *HomeCallerSession) HistoricalRoots(arg0 *big.Int) ([32]byte, error) {
	return _Home.Contract.HistoricalRoots(&_Home.CallOpts, arg0)
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

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Home *HomeCaller) Nonce(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Home *HomeSession) Nonce() (uint32, error) {
	return _Home.Contract.Nonce(&_Home.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Home *HomeCallerSession) Nonce() (uint32, error) {
	return _Home.Contract.Nonce(&_Home.CallOpts)
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
// Solidity: function suggestUpdate() view returns(uint32 _nonce, bytes32 _root)
func (_Home *HomeCaller) SuggestUpdate(opts *bind.CallOpts) (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "suggestUpdate")

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

// SuggestUpdate is a free data retrieval call binding the contract method 0x36e104de.
//
// Solidity: function suggestUpdate() view returns(uint32 _nonce, bytes32 _root)
func (_Home *HomeSession) SuggestUpdate() (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	return _Home.Contract.SuggestUpdate(&_Home.CallOpts)
}

// SuggestUpdate is a free data retrieval call binding the contract method 0x36e104de.
//
// Solidity: function suggestUpdate() view returns(uint32 _nonce, bytes32 _root)
func (_Home *HomeCallerSession) SuggestUpdate() (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	return _Home.Contract.SuggestUpdate(&_Home.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_Home *HomeCaller) SystemMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Home.contract.Call(opts, &out, "systemMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_Home *HomeSession) SystemMessenger() (common.Address, error) {
	return _Home.Contract.SystemMessenger(&_Home.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_Home *HomeCallerSession) SystemMessenger() (common.Address, error) {
	return _Home.Contract.SystemMessenger(&_Home.CallOpts)
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

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_Home *HomeTransactor) Dispatch(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "dispatch", _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_Home *HomeSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Home.Contract.Dispatch(&_Home.TransactOpts, _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_Home *HomeTransactorSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Home.Contract.Dispatch(&_Home.TransactOpts, _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x88a278ec.
//
// Solidity: function improperAttestation(address _updater, bytes _attestation) returns(bool)
func (_Home *HomeTransactor) ImproperAttestation(opts *bind.TransactOpts, _updater common.Address, _attestation []byte) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "improperAttestation", _updater, _attestation)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x88a278ec.
//
// Solidity: function improperAttestation(address _updater, bytes _attestation) returns(bool)
func (_Home *HomeSession) ImproperAttestation(_updater common.Address, _attestation []byte) (*types.Transaction, error) {
	return _Home.Contract.ImproperAttestation(&_Home.TransactOpts, _updater, _attestation)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x88a278ec.
//
// Solidity: function improperAttestation(address _updater, bytes _attestation) returns(bool)
func (_Home *HomeTransactorSession) ImproperAttestation(_updater common.Address, _attestation []byte) (*types.Transaction, error) {
	return _Home.Contract.ImproperAttestation(&_Home.TransactOpts, _updater, _attestation)
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

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_Home *HomeTransactor) SetSystemMessenger(opts *bind.TransactOpts, _systemMessenger common.Address) (*types.Transaction, error) {
	return _Home.contract.Transact(opts, "setSystemMessenger", _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_Home *HomeSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _Home.Contract.SetSystemMessenger(&_Home.TransactOpts, _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_Home *HomeTransactorSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _Home.Contract.SetSystemMessenger(&_Home.TransactOpts, _systemMessenger)
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
	Tips                []byte
	Message             []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0x718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e785.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes tips, bytes message)
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

// WatchDispatch is a free log subscription operation binding the contract event 0x718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e785.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes tips, bytes message)
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

// ParseDispatch is a log parse operation binding the contract event 0x718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e785.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes tips, bytes message)
func (_Home *HomeFilterer) ParseDispatch(log types.Log) (*HomeDispatch, error) {
	event := new(HomeDispatch)
	if err := _Home.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeImproperAttestationIterator is returned from FilterImproperAttestation and is used to iterate over the raw logs and unpacked data for ImproperAttestation events raised by the Home contract.
type HomeImproperAttestationIterator struct {
	Event *HomeImproperAttestation // Event containing the contract specifics and raw log

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
func (it *HomeImproperAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeImproperAttestation)
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
		it.Event = new(HomeImproperAttestation)
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
func (it *HomeImproperAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeImproperAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeImproperAttestation represents a ImproperAttestation event raised by the Home contract.
type HomeImproperAttestation struct {
	Updater     common.Address
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterImproperAttestation is a free log retrieval operation binding the contract event 0x287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b.
//
// Solidity: event ImproperAttestation(address updater, bytes attestation)
func (_Home *HomeFilterer) FilterImproperAttestation(opts *bind.FilterOpts) (*HomeImproperAttestationIterator, error) {

	logs, sub, err := _Home.contract.FilterLogs(opts, "ImproperAttestation")
	if err != nil {
		return nil, err
	}
	return &HomeImproperAttestationIterator{contract: _Home.contract, event: "ImproperAttestation", logs: logs, sub: sub}, nil
}

// WatchImproperAttestation is a free log subscription operation binding the contract event 0x287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b.
//
// Solidity: event ImproperAttestation(address updater, bytes attestation)
func (_Home *HomeFilterer) WatchImproperAttestation(opts *bind.WatchOpts, sink chan<- *HomeImproperAttestation) (event.Subscription, error) {

	logs, sub, err := _Home.contract.WatchLogs(opts, "ImproperAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeImproperAttestation)
				if err := _Home.contract.UnpackLog(event, "ImproperAttestation", log); err != nil {
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
// Solidity: event ImproperAttestation(address updater, bytes attestation)
func (_Home *HomeFilterer) ParseImproperAttestation(log types.Log) (*HomeImproperAttestation, error) {
	event := new(HomeImproperAttestation)
	if err := _Home.contract.UnpackLog(event, "ImproperAttestation", log); err != nil {
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
	Nonce      uint32
	Root       [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_Home *HomeFilterer) FilterUpdate(opts *bind.FilterOpts, homeDomain []uint32, nonce []uint32, root [][32]byte) (*HomeUpdateIterator, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Home.contract.FilterLogs(opts, "Update", homeDomainRule, nonceRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &HomeUpdateIterator{contract: _Home.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_Home *HomeFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *HomeUpdate, homeDomain []uint32, nonce []uint32, root [][32]byte) (event.Subscription, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Home.contract.WatchLogs(opts, "Update", homeDomainRule, nonceRule, rootRule)
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

// ParseUpdate is a log parse operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
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

// HomeHarnessMetaData contains all meta data concerning the HomeHarness contract.
var HomeHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destinationAndNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tips\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"ImproperAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updaterManager\",\"type\":\"address\"}],\"name\":\"NewUpdaterManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"UpdaterSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_BODY_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"destinationAndNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historicalRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeDomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_attestation\",\"type\":\"bytes\"}],\"name\":\"improperAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUpdaterManager\",\"name\":\"_updaterManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sensitiveValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFailed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"_systemMessenger\",\"type\":\"address\"}],\"name\":\"setSystemMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updater\",\"type\":\"address\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_updaterManager\",\"type\":\"address\"}],\"name\":\"setUpdaterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumHome.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestUpdate\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemMessenger\",\"outputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updaterManager\",\"outputs\":[{\"internalType\":\"contractIUpdaterManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"522ae002": "MAX_MESSAGE_BODY_BYTES()",
		"ffa1ad74": "VERSION()",
		"06661abd": "count()",
		"da180e70": "destinationAndNonce(uint32,uint32)",
		"f7560e40": "dispatch(uint32,bytes32,uint32,bytes,bytes)",
		"7ea97f40": "historicalRoots(uint256)",
		"45630b1a": "homeDomainHash()",
		"88a278ec": "improperAttestation(address,bytes)",
		"c4d66de8": "initialize(address)",
		"8d3638f4": "localDomain()",
		"affed0e0": "nonce()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"ebf0c717": "root()",
		"089d2894": "sensitiveValue()",
		"146901db": "setFailed()",
		"48639d24": "setSensitiveValue(uint256)",
		"b7bc563e": "setSystemMessenger(address)",
		"9d54f419": "setUpdater(address)",
		"9776120e": "setUpdaterManager(address)",
		"c19d93fb": "state()",
		"36e104de": "suggestUpdate()",
		"ccbdf9c9": "systemMessenger()",
		"f2fde38b": "transferOwnership(address)",
		"fd54b228": "tree()",
		"df034cd0": "updater()",
		"9df6c8e1": "updaterManager()",
	},
	Bin: "0x60a06040523480156200001157600080fd5b506040516200393b3803806200393b833981016040819052620000349162000043565b63ffffffff1660805262000072565b6000602082840312156200005657600080fd5b815163ffffffff811681146200006b57600080fd5b9392505050565b608051613898620000a3600039600081816102f60152818161060201528181610ecd01526119d401526138986000f3fe6080604052600436106101ac5760003560e01c80639d54f419116100ec578063da180e701161008a578063f2fde38b11610064578063f2fde38b14610539578063f7560e4014610559578063fd54b2281461056c578063ffa1ad741461058357600080fd5b8063da180e70146104be578063df034cd0146104f7578063ebf0c7171461052457600080fd5b8063b7bc563e116100c6578063b7bc563e1461040d578063c19d93fb1461042d578063c4d66de814610471578063ccbdf9c91461049157600080fd5b80639d54f419146103995780639df6c8e1146103b9578063affed0e0146103ef57600080fd5b8063522ae0021161015957806388a278ec1161013357806388a278ec146102b45780638d3638f4146102e45780638da5cb5b1461032d5780639776120e1461037957600080fd5b8063522ae00214610269578063715018a61461027f5780637ea97f401461029457600080fd5b806336e104de1161018a57806336e104de1461020357806345630b1a1461023457806348639d241461024957600080fd5b806306661abd146101b1578063089d2894146101d5578063146901db146101ec575b600080fd5b3480156101bd57600080fd5b506020545b6040519081526020015b60405180910390f35b3480156101e157600080fd5b506101c261014b5481565b3480156101f857600080fd5b506102016105aa565b005b34801561020f57600080fd5b506102186105b4565b6040805163ffffffff90931683526020830191909152016101cc565b34801561024057600080fd5b506101c26105fb565b34801561025557600080fd5b5061020161026436600461302d565b61062b565b34801561027557600080fd5b506101c261080081565b34801561028b57600080fd5b50610201610639565b3480156102a057600080fd5b506101c26102af36600461302d565b6106a5565b3480156102c057600080fd5b506102d46102cf366004613142565b6106c6565b60405190151581526020016101cc565b3480156102f057600080fd5b506103187f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101cc565b34801561033957600080fd5b5060855473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cc565b34801561038557600080fd5b50610201610394366004613192565b610815565b3480156103a557600080fd5b506102016103b4366004613192565b610888565b3480156103c557600080fd5b5061011b5461035490640100000000900473ffffffffffffffffffffffffffffffffffffffff1681565b3480156103fb57600080fd5b5061011b546103189063ffffffff1681565b34801561041957600080fd5b50610201610428366004613192565b610948565b34801561043957600080fd5b5061011b54610464907801000000000000000000000000000000000000000000000000900460ff1681565b6040516101cc91906131de565b34801561047d57600080fd5b5061020161048c366004613192565b6109f6565b34801561049d57600080fd5b5060b8546103549073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104ca57600080fd5b506104de6104d9366004613233565b610bb8565b60405167ffffffffffffffff90911681526020016101cc565b34801561050357600080fd5b5060b7546103549073ffffffffffffffffffffffffffffffffffffffff1681565b34801561053057600080fd5b506101c2610bd9565b34801561054557600080fd5b50610201610554366004613192565b610be5565b610201610567366004613266565b610cde565b34801561057857600080fd5b506020546101c29081565b34801561058f57600080fd5b50610598600081565b60405160ff90911681526020016101cc565b6105b2610fd8565b565b602154600090819080156105f6576105cd600182613324565b925060218363ffffffff16815481106105e8576105e861333b565b906000526020600020015491505b509091565b60006106267f00000000000000000000000000000000000000000000000000000000000000006110f3565b905090565b61063361116c565b61014b55565b60855473ffffffffffffffffffffffffffffffffffffffff1633146105b25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b602181815481106106b557600080fd5b600091825260209091200154905081565b6000600261011b547801000000000000000000000000000000000000000000000000900460ff1660028111156106fe576106fe6131af565b0361074b5760405162461bcd60e51b815260206004820152600c60248201527f6661696c65642073746174650000000000000000000000000000000000000000604482015260640161069c565b600061075784846111d3565b9050600061076a62ffffff1983166112dc565b9050600061077d62ffffff1984166112f0565b60215490915063ffffffff831610156107c65760218263ffffffff16815481106107a9576107a961333b565b906000526020600020015481036107c6576000935050505061080f565b6107ce610fd8565b7f287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b86866040516107ff9291906133e0565b60405180910390a1600193505050505b92915050565b60855473ffffffffffffffffffffffffffffffffffffffff16331461087c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161069c565b61088581611305565b50565b61011b54640100000000900473ffffffffffffffffffffffffffffffffffffffff1633146108f85760405162461bcd60e51b815260206004820152600f60248201527f21757064617465724d616e616765720000000000000000000000000000000000604482015260640161069c565b610901816113ed565b5061011b80547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000179055565b60855473ffffffffffffffffffffffffffffffffffffffff1633146109af5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161069c565b60b880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000610a02600161146c565b90508015610a3757605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610a4082611305565b610ada61011b60049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663df034cd06040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ab1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad5919061340f565b6115c5565b61011b80547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff1678010000000000000000000000000000000000000000000000001790556021805460018101825560009182527f3a6357012c1a3ae0a17d304c9920310382d968ebcc4b1771f41c6b304205b57001558015610bb457605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b600067ffffffff00000000602084901b1663ffffffff8316175b9392505050565b60006106266000611653565b60855473ffffffffffffffffffffffffffffffffffffffff163314610c4c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161069c565b73ffffffffffffffffffffffffffffffffffffffff8116610cd55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161069c565b61088581611666565b600261011b547801000000000000000000000000000000000000000000000000900460ff166002811115610d1457610d146131af565b03610d615760405162461bcd60e51b815260206004820152600c60248201527f6661696c65642073746174650000000000000000000000000000000000000000604482015260640161069c565b61080081511115610db45760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e670000000000000000000000000000000000000000604482015260640161069c565b34610dcc610dc1846116dd565b62ffffff19166116ea565b6bffffffffffffffffffffffff1614610e275760405162461bcd60e51b815260206004820152600560248201527f2174697073000000000000000000000000000000000000000000000000000000604482015260640161069c565b61011b54610e3c9063ffffffff16600161342c565b61011b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff929092169190911790556000610e7d8561174c565b61011b54604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060e090811b821660228401526026830186905293841b811660468301528a841b8116604a830152604e82018a90529288901b909216606e83015280518083036052018152607290920190529091506000610f478286866117ab565b80516020820120909150610f5a81611825565b61011b5463ffffffff1660208a901b67ffffffff00000000161767ffffffffffffffff166001610f8960205490565b610f939190613324565b827f718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e7858986604051610fc5929190613454565b60405180910390a4505050505050505050565b61011b805478020000000000000000000000000000000000000000000000007fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff90911617908190556040517f5b3c2cbf00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff6401000000009092049190911690635b3c2cbf90602401600060405180830381600087803b15801561109257600080fd5b505af11580156110a6573d6000803e3d6000fd5b505060b75460405133935073ffffffffffffffffffffffffffffffffffffffff90911691507f98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce90600090a3565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e083901b1660208201527f53594e000000000000000000000000000000000000000000000000000000000060248201526000906027015b604051602081830303815290604052805190602001209050919050565b60b85473ffffffffffffffffffffffffffffffffffffffff1633146105b25760405162461bcd60e51b815260206004820152601060248201527f2173797374656d4d657373656e67657200000000000000000000000000000000604482015260640161069c565b60006111df8282611855565b905060286bffffffffffffffffffffffff601883901c16116112435760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e0000000000000000000000000000604482015260640161069c565b6112788361125662ffffff198416611879565b61127361126862ffffff19861661188e565b62ffffff19166118c1565b611914565b61129061128a62ffffff1983166119bc565b846119d0565b61080f5760405162461bcd60e51b815260206004820152601860248201527f5369676e6572206973206e6f7420616e20757064617465720000000000000000604482015260640161069c565b600061080f62ffffff198316600480611a71565b600061080f62ffffff19831660086020611aa1565b73ffffffffffffffffffffffffffffffffffffffff81163b6113695760405162461bcd60e51b815260206004820152601860248201527f21636f6e747261637420757064617465724d616e616765720000000000000000604482015260640161069c565b61011b80547fffffffffffffffff0000000000000000000000000000000000000000ffffffff1664010000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf9060200160405180910390a150565b60b7805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a9101610bab565b605254600090610100900460ff161561150b578160ff1660011480156114915750303b155b6115035760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161069c565b506000919050565b60525460ff8084169116106115885760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161069c565b50605280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b605254610100900460ff166116425760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161069c565b61164a611c5f565b610885816113ed565b600061080f82611661611ce4565b6121a5565b6085805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061080f826002611855565b60008161170060025b62ffffff19831690612268565b5061170a83612369565b61171384612397565b61171c856123b8565b611725866123d9565b61172f9190613479565b6117399190613479565b6117439190613479565b91505b50919050565b60007fffffffffffffffffffffffff0000000000000000000000000000000000000000821461177c573392915050565b61178461116c565b507fffffffffffffffffffffffff0000000000000000000000000000000000000000919050565b82516060906000906117bf600460026134a0565b60ff166117cc91906134c9565b905060008451826117dd91906134c9565b905060016117ed600460026134a0565b60ff16838389898960405160200161180b97969594939291906134e6565b604051602081830303815290604052925050509392505050565b6118306000826123fa565b602161183c6000611653565b8154600181018355600092835260209092209091015550565b81516000906020840161187064ffffffffff8516828461251d565b95945050505050565b600061080f62ffffff19831682602881612564565b600061080f60286118b181601886901c6bffffffffffffffffffffffff16613324565b62ffffff19851691906000612564565b60606000806118de8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061190384836020016125ea565b508181016020016040529052919050565b600061192562ffffff19841661278f565b9050611930816127ec565b90508373ffffffffffffffffffffffffffffffffffffffff166119538284612827565b73ffffffffffffffffffffffffffffffffffffffff16146119b65760405162461bcd60e51b815260206004820152601160248201527f496e76616c6964207369676e6174757265000000000000000000000000000000604482015260640161069c565b50505050565b600061080f62ffffff198316826004611a71565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff168363ffffffff1614611a4d5760405162461bcd60e51b815260206004820152600c60248201527f57726f6e6720646f6d61696e0000000000000000000000000000000000000000604482015260640161069c565b5060b75473ffffffffffffffffffffffffffffffffffffffff908116911614919050565b6000611a7e826020613584565b611a899060086134a0565b60ff16611a97858585611aa1565b901c949350505050565b60008160ff16600003611ab657506000610bd2565b611ace8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16611ae960ff8416856135a7565b1115611b6157611b48611b0a8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16611b308660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff1661284b565b60405162461bcd60e51b815260040161069c91906135bf565b60208260ff161115611bdb5760405162461bcd60e51b815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e203332206279746573000000000000606482015260840161069c565b600882026000611bf98660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b605254610100900460ff16611cdc5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161069c565b6105b26128b9565b611cec61300e565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b602081101561226057600182821c81169081900361220c578582602081106121d9576121d961333b565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350612257565b8385836020811061221f5761221f61333b565b602002015160405160200161223e929190918252602082015260400190565b6040516020818303038152906040528051906020012093505b506001016121af565b505092915050565b6000612274838361293f565b6123625760006122936122878560d81c90565b64ffffffffff16612962565b91505060006122a88464ffffffffff16612962565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e0160405160208183030381529060405290508060405162461bcd60e51b815260040161069c91906135bf565b5090919050565b60008161237660026116f3565b5061238a62ffffff1984166026600c611a71565b63ffffffff169392505050565b6000816123a460026116f3565b5061238a62ffffff198416601a600c611a71565b6000816123c560026116f3565b5061238a62ffffff198416600e600c611a71565b6000816123e660026116f3565b5061238a62ffffff1984166002600c611a71565b6020808301549060019061240f9060026136ea565b6124199190613324565b81106124675760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c00000000000000000000000000000000604482015260640161069c565b6001016020830181905560005b602081101561250f57816001166001036124a3578284826020811061249b5761249b61333b565b015550505050565b8381602081106124b5576124b561333b565b01546040805160208101929092528101849052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209250600191821c9101612474565b506125186136f6565b505050565b60008061252a83856135a7565b905060405181111561253a575060005b8060000361254f5762ffffff19915050610bd2565b5050606092831b9190911790911b1760181b90565b60008061257f8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905061259886612a4c565b846125a387846135a7565b6125ad91906135a7565b11156125c05762ffffff199150506125e2565b6125ca85826135a7565b90506125de8364ffffffffff16828661251d565b9150505b949350505050565b600062ffffff19808416036126675760405162461bcd60e51b815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f6572206465726566000000000000000000000000000000000000000000000000606482015260840161069c565b61267083612a94565b6126e25760405162461bcd60e51b815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e746572206465726566000000000000000000000000000000000000000000606482015260840161069c565b60006126fc8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006127268560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600060405190508481111561274b5760206060fd5b8285848460045afa506127856127618760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b9695505050505050565b6000806127aa8360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006127d48460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c0161114f565b60008060006128368585612ad1565b9150915061284381612b3f565b509392505050565b6060600061285886612962565b915050600061286686612962565b915050600061287486612962565b915050600061288286612962565b9150508383838360405160200161289c9493929190613725565b604051602081830303815290604052945050505050949350505050565b605254610100900460ff166129365760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161069c565b6105b233611666565b60008164ffffffffff166129538460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff1611156129d55760006129818260086134a0565b60ff1685901c905061299281612d2b565b61ffff16841793508160ff166010146129ad57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612968565b50600f5b60ff8160ff161015612a465760006129f28260086134a0565b60ff1685901c9050612a0381612d2b565b61ffff16831792508160ff16600014612a1e57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016129d9565b50915091565b6000612a668260181c6bffffffffffffffffffffffff1690565b612a7e8360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b6000612aa08260d81c90565b64ffffffffff1664ffffffffff03612aba57506000919050565b6000612ac583612a4c565b60405110199392505050565b6000808251604103612b075760208301516040840151606085015160001a612afb87828585612d5d565b94509450505050612b38565b8251604003612b305760208301516040840151612b25868383612e75565b935093505050612b38565b506000905060025b9250929050565b6000816004811115612b5357612b536131af565b03612b5b5750565b6001816004811115612b6f57612b6f6131af565b03612bbc5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161069c565b6002816004811115612bd057612bd06131af565b03612c1d5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161069c565b6003816004811115612c3157612c316131af565b03612ca45760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161069c565b6004816004811115612cb857612cb86131af565b036108855760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161069c565b6000612d3d60048360ff16901c612ec7565b60ff1661ffff919091161760081b612d5482612ec7565b60ff1617919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612d945750600090506003612e6c565b8460ff16601b14158015612dac57508460ff16601c14155b15612dbd5750600090506004612e6c565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612e11573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612e6557600060019250925050612e6c565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831681612eab60ff86901c601b6135a7565b9050612eb987828885612d5d565b935093505050935093915050565b600060f08083179060ff82169003612ee25750603092915050565b8060ff1660f103612ef65750603192915050565b8060ff1660f203612f0a5750603292915050565b8060ff1660f303612f1e5750603392915050565b8060ff1660f403612f325750603492915050565b8060ff1660f503612f465750603592915050565b8060ff1660f603612f5a5750603692915050565b8060ff1660f703612f6e5750603792915050565b8060ff1660f803612f825750603892915050565b8060ff1660f903612f965750603992915050565b8060ff1660fa03612faa5750606192915050565b8060ff1660fb03612fbe5750606292915050565b8060ff1660fc03612fd25750606392915050565b8060ff1660fd03612fe65750606492915050565b8060ff1660fe03612ffa5750606592915050565b8060ff1660ff036117465750606692915050565b6040518061040001604052806020906020820280368337509192915050565b60006020828403121561303f57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461088557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126130a857600080fd5b813567ffffffffffffffff808211156130c3576130c3613068565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561310957613109613068565b8160405283815286602085880101111561312257600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561315557600080fd5b823561316081613046565b9150602083013567ffffffffffffffff81111561317c57600080fd5b61318885828601613097565b9150509250929050565b6000602082840312156131a457600080fd5b8135610bd281613046565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160038310613219577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b803563ffffffff811681146115c057600080fd5b6000806040838503121561324657600080fd5b61324f8361321f565b915061325d6020840161321f565b90509250929050565b600080600080600060a0868803121561327e57600080fd5b6132878661321f565b94506020860135935061329c6040870161321f565b9250606086013567ffffffffffffffff808211156132b957600080fd5b6132c589838a01613097565b935060808801359150808211156132db57600080fd5b506132e888828901613097565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015613336576133366132f5565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b8381101561338557818101518382015260200161336d565b838111156119b65750506000910152565b600081518084526133ae81602086016020860161336a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006125e26040830184613396565b60006020828403121561342157600080fd5b8151610bd281613046565b600063ffffffff80831681851680830382111561344b5761344b6132f5565b01949350505050565b6040815260006134676040830185613396565b82810360208401526118708185613396565b60006bffffffffffffffffffffffff80831681851680830382111561344b5761344b6132f5565b600060ff821660ff84168160ff04811182151516156134c1576134c16132f5565b029392505050565b600061ffff80831681851680830382111561344b5761344b6132f5565b60007fffff000000000000000000000000000000000000000000000000000000000000808a60f01b168352808960f01b166002840152808860f01b166004840152808760f01b16600684015250845161354681600885016020890161336a565b84519083019061355d81600884016020890161336a565b845191019061357381600884016020880161336a565b016008019998505050505050505050565b600060ff821660ff84168082101561359e5761359e6132f5565b90039392505050565b600082198211156135ba576135ba6132f5565b500190565b602081526000610bd26020830184613396565b600181815b8085111561362b57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115613611576136116132f5565b8085161561361e57918102915b93841c93908002906135d7565b509250929050565b6000826136425750600161080f565b8161364f5750600061080f565b8160018114613665576002811461366f5761368b565b600191505061080f565b60ff841115613680576136806132f5565b50506001821b61080f565b5060208310610133831016604e8410600b84101617156136ae575081810a61080f565b6136b883836135d2565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156134c1576134c16132f5565b6000610bd28383613633565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d820161278556fea26469706673582212203679c93c15a4bd1cc45f720d99e5e535960827dc2b376b1d3289028f4592ede964736f6c634300080d0033",
}

// HomeHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use HomeHarnessMetaData.ABI instead.
var HomeHarnessABI = HomeHarnessMetaData.ABI

// Deprecated: Use HomeHarnessMetaData.Sigs instead.
// HomeHarnessFuncSigs maps the 4-byte function signature to its string representation.
var HomeHarnessFuncSigs = HomeHarnessMetaData.Sigs

// HomeHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HomeHarnessMetaData.Bin instead.
var HomeHarnessBin = HomeHarnessMetaData.Bin

// DeployHomeHarness deploys a new Ethereum contract, binding an instance of HomeHarness to it.
func DeployHomeHarness(auth *bind.TransactOpts, backend bind.ContractBackend, _domain uint32) (common.Address, *types.Transaction, *HomeHarness, error) {
	parsed, err := HomeHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HomeHarnessBin), backend, _domain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HomeHarness{HomeHarnessCaller: HomeHarnessCaller{contract: contract}, HomeHarnessTransactor: HomeHarnessTransactor{contract: contract}, HomeHarnessFilterer: HomeHarnessFilterer{contract: contract}}, nil
}

// HomeHarness is an auto generated Go binding around an Ethereum contract.
type HomeHarness struct {
	HomeHarnessCaller     // Read-only binding to the contract
	HomeHarnessTransactor // Write-only binding to the contract
	HomeHarnessFilterer   // Log filterer for contract events
}

// HomeHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type HomeHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HomeHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HomeHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HomeHarnessSession struct {
	Contract     *HomeHarness      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HomeHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HomeHarnessCallerSession struct {
	Contract *HomeHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HomeHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HomeHarnessTransactorSession struct {
	Contract     *HomeHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HomeHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type HomeHarnessRaw struct {
	Contract *HomeHarness // Generic contract binding to access the raw methods on
}

// HomeHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HomeHarnessCallerRaw struct {
	Contract *HomeHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// HomeHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HomeHarnessTransactorRaw struct {
	Contract *HomeHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHomeHarness creates a new instance of HomeHarness, bound to a specific deployed contract.
func NewHomeHarness(address common.Address, backend bind.ContractBackend) (*HomeHarness, error) {
	contract, err := bindHomeHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HomeHarness{HomeHarnessCaller: HomeHarnessCaller{contract: contract}, HomeHarnessTransactor: HomeHarnessTransactor{contract: contract}, HomeHarnessFilterer: HomeHarnessFilterer{contract: contract}}, nil
}

// NewHomeHarnessCaller creates a new read-only instance of HomeHarness, bound to a specific deployed contract.
func NewHomeHarnessCaller(address common.Address, caller bind.ContractCaller) (*HomeHarnessCaller, error) {
	contract, err := bindHomeHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HomeHarnessCaller{contract: contract}, nil
}

// NewHomeHarnessTransactor creates a new write-only instance of HomeHarness, bound to a specific deployed contract.
func NewHomeHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*HomeHarnessTransactor, error) {
	contract, err := bindHomeHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HomeHarnessTransactor{contract: contract}, nil
}

// NewHomeHarnessFilterer creates a new log filterer instance of HomeHarness, bound to a specific deployed contract.
func NewHomeHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*HomeHarnessFilterer, error) {
	contract, err := bindHomeHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HomeHarnessFilterer{contract: contract}, nil
}

// bindHomeHarness binds a generic wrapper to an already deployed contract.
func bindHomeHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HomeHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HomeHarness *HomeHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HomeHarness.Contract.HomeHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HomeHarness *HomeHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HomeHarness.Contract.HomeHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HomeHarness *HomeHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HomeHarness.Contract.HomeHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HomeHarness *HomeHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HomeHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HomeHarness *HomeHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HomeHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HomeHarness *HomeHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HomeHarness.Contract.contract.Transact(opts, method, params...)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_HomeHarness *HomeHarnessCaller) MAXMESSAGEBODYBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "MAX_MESSAGE_BODY_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_HomeHarness *HomeHarnessSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _HomeHarness.Contract.MAXMESSAGEBODYBYTES(&_HomeHarness.CallOpts)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_HomeHarness *HomeHarnessCallerSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _HomeHarness.Contract.MAXMESSAGEBODYBYTES(&_HomeHarness.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_HomeHarness *HomeHarnessCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_HomeHarness *HomeHarnessSession) VERSION() (uint8, error) {
	return _HomeHarness.Contract.VERSION(&_HomeHarness.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_HomeHarness *HomeHarnessCallerSession) VERSION() (uint8, error) {
	return _HomeHarness.Contract.VERSION(&_HomeHarness.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_HomeHarness *HomeHarnessCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_HomeHarness *HomeHarnessSession) Count() (*big.Int, error) {
	return _HomeHarness.Contract.Count(&_HomeHarness.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_HomeHarness *HomeHarnessCallerSession) Count() (*big.Int, error) {
	return _HomeHarness.Contract.Count(&_HomeHarness.CallOpts)
}

// DestinationAndNonce is a free data retrieval call binding the contract method 0xda180e70.
//
// Solidity: function destinationAndNonce(uint32 _destination, uint32 _nonce) pure returns(uint64)
func (_HomeHarness *HomeHarnessCaller) DestinationAndNonce(opts *bind.CallOpts, _destination uint32, _nonce uint32) (uint64, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "destinationAndNonce", _destination, _nonce)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DestinationAndNonce is a free data retrieval call binding the contract method 0xda180e70.
//
// Solidity: function destinationAndNonce(uint32 _destination, uint32 _nonce) pure returns(uint64)
func (_HomeHarness *HomeHarnessSession) DestinationAndNonce(_destination uint32, _nonce uint32) (uint64, error) {
	return _HomeHarness.Contract.DestinationAndNonce(&_HomeHarness.CallOpts, _destination, _nonce)
}

// DestinationAndNonce is a free data retrieval call binding the contract method 0xda180e70.
//
// Solidity: function destinationAndNonce(uint32 _destination, uint32 _nonce) pure returns(uint64)
func (_HomeHarness *HomeHarnessCallerSession) DestinationAndNonce(_destination uint32, _nonce uint32) (uint64, error) {
	return _HomeHarness.Contract.DestinationAndNonce(&_HomeHarness.CallOpts, _destination, _nonce)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_HomeHarness *HomeHarnessCaller) HistoricalRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "historicalRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_HomeHarness *HomeHarnessSession) HistoricalRoots(arg0 *big.Int) ([32]byte, error) {
	return _HomeHarness.Contract.HistoricalRoots(&_HomeHarness.CallOpts, arg0)
}

// HistoricalRoots is a free data retrieval call binding the contract method 0x7ea97f40.
//
// Solidity: function historicalRoots(uint256 ) view returns(bytes32)
func (_HomeHarness *HomeHarnessCallerSession) HistoricalRoots(arg0 *big.Int) ([32]byte, error) {
	return _HomeHarness.Contract.HistoricalRoots(&_HomeHarness.CallOpts, arg0)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_HomeHarness *HomeHarnessCaller) HomeDomainHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "homeDomainHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_HomeHarness *HomeHarnessSession) HomeDomainHash() ([32]byte, error) {
	return _HomeHarness.Contract.HomeDomainHash(&_HomeHarness.CallOpts)
}

// HomeDomainHash is a free data retrieval call binding the contract method 0x45630b1a.
//
// Solidity: function homeDomainHash() view returns(bytes32)
func (_HomeHarness *HomeHarnessCallerSession) HomeDomainHash() ([32]byte, error) {
	return _HomeHarness.Contract.HomeDomainHash(&_HomeHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HomeHarness *HomeHarnessCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HomeHarness *HomeHarnessSession) LocalDomain() (uint32, error) {
	return _HomeHarness.Contract.LocalDomain(&_HomeHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HomeHarness *HomeHarnessCallerSession) LocalDomain() (uint32, error) {
	return _HomeHarness.Contract.LocalDomain(&_HomeHarness.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_HomeHarness *HomeHarnessCaller) Nonce(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_HomeHarness *HomeHarnessSession) Nonce() (uint32, error) {
	return _HomeHarness.Contract.Nonce(&_HomeHarness.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_HomeHarness *HomeHarnessCallerSession) Nonce() (uint32, error) {
	return _HomeHarness.Contract.Nonce(&_HomeHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HomeHarness *HomeHarnessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HomeHarness *HomeHarnessSession) Owner() (common.Address, error) {
	return _HomeHarness.Contract.Owner(&_HomeHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HomeHarness *HomeHarnessCallerSession) Owner() (common.Address, error) {
	return _HomeHarness.Contract.Owner(&_HomeHarness.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_HomeHarness *HomeHarnessCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_HomeHarness *HomeHarnessSession) Root() ([32]byte, error) {
	return _HomeHarness.Contract.Root(&_HomeHarness.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_HomeHarness *HomeHarnessCallerSession) Root() ([32]byte, error) {
	return _HomeHarness.Contract.Root(&_HomeHarness.CallOpts)
}

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_HomeHarness *HomeHarnessCaller) SensitiveValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "sensitiveValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_HomeHarness *HomeHarnessSession) SensitiveValue() (*big.Int, error) {
	return _HomeHarness.Contract.SensitiveValue(&_HomeHarness.CallOpts)
}

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_HomeHarness *HomeHarnessCallerSession) SensitiveValue() (*big.Int, error) {
	return _HomeHarness.Contract.SensitiveValue(&_HomeHarness.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_HomeHarness *HomeHarnessCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_HomeHarness *HomeHarnessSession) State() (uint8, error) {
	return _HomeHarness.Contract.State(&_HomeHarness.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_HomeHarness *HomeHarnessCallerSession) State() (uint8, error) {
	return _HomeHarness.Contract.State(&_HomeHarness.CallOpts)
}

// SuggestUpdate is a free data retrieval call binding the contract method 0x36e104de.
//
// Solidity: function suggestUpdate() view returns(uint32 _nonce, bytes32 _root)
func (_HomeHarness *HomeHarnessCaller) SuggestUpdate(opts *bind.CallOpts) (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "suggestUpdate")

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

// SuggestUpdate is a free data retrieval call binding the contract method 0x36e104de.
//
// Solidity: function suggestUpdate() view returns(uint32 _nonce, bytes32 _root)
func (_HomeHarness *HomeHarnessSession) SuggestUpdate() (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	return _HomeHarness.Contract.SuggestUpdate(&_HomeHarness.CallOpts)
}

// SuggestUpdate is a free data retrieval call binding the contract method 0x36e104de.
//
// Solidity: function suggestUpdate() view returns(uint32 _nonce, bytes32 _root)
func (_HomeHarness *HomeHarnessCallerSession) SuggestUpdate() (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	return _HomeHarness.Contract.SuggestUpdate(&_HomeHarness.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_HomeHarness *HomeHarnessCaller) SystemMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "systemMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_HomeHarness *HomeHarnessSession) SystemMessenger() (common.Address, error) {
	return _HomeHarness.Contract.SystemMessenger(&_HomeHarness.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_HomeHarness *HomeHarnessCallerSession) SystemMessenger() (common.Address, error) {
	return _HomeHarness.Contract.SystemMessenger(&_HomeHarness.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_HomeHarness *HomeHarnessCaller) Tree(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "tree")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_HomeHarness *HomeHarnessSession) Tree() (*big.Int, error) {
	return _HomeHarness.Contract.Tree(&_HomeHarness.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_HomeHarness *HomeHarnessCallerSession) Tree() (*big.Int, error) {
	return _HomeHarness.Contract.Tree(&_HomeHarness.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_HomeHarness *HomeHarnessCaller) Updater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "updater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_HomeHarness *HomeHarnessSession) Updater() (common.Address, error) {
	return _HomeHarness.Contract.Updater(&_HomeHarness.CallOpts)
}

// Updater is a free data retrieval call binding the contract method 0xdf034cd0.
//
// Solidity: function updater() view returns(address)
func (_HomeHarness *HomeHarnessCallerSession) Updater() (common.Address, error) {
	return _HomeHarness.Contract.Updater(&_HomeHarness.CallOpts)
}

// UpdaterManager is a free data retrieval call binding the contract method 0x9df6c8e1.
//
// Solidity: function updaterManager() view returns(address)
func (_HomeHarness *HomeHarnessCaller) UpdaterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HomeHarness.contract.Call(opts, &out, "updaterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpdaterManager is a free data retrieval call binding the contract method 0x9df6c8e1.
//
// Solidity: function updaterManager() view returns(address)
func (_HomeHarness *HomeHarnessSession) UpdaterManager() (common.Address, error) {
	return _HomeHarness.Contract.UpdaterManager(&_HomeHarness.CallOpts)
}

// UpdaterManager is a free data retrieval call binding the contract method 0x9df6c8e1.
//
// Solidity: function updaterManager() view returns(address)
func (_HomeHarness *HomeHarnessCallerSession) UpdaterManager() (common.Address, error) {
	return _HomeHarness.Contract.UpdaterManager(&_HomeHarness.CallOpts)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_HomeHarness *HomeHarnessTransactor) Dispatch(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "dispatch", _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_HomeHarness *HomeHarnessSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _HomeHarness.Contract.Dispatch(&_HomeHarness.TransactOpts, _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
func (_HomeHarness *HomeHarnessTransactorSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _HomeHarness.Contract.Dispatch(&_HomeHarness.TransactOpts, _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x88a278ec.
//
// Solidity: function improperAttestation(address _updater, bytes _attestation) returns(bool)
func (_HomeHarness *HomeHarnessTransactor) ImproperAttestation(opts *bind.TransactOpts, _updater common.Address, _attestation []byte) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "improperAttestation", _updater, _attestation)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x88a278ec.
//
// Solidity: function improperAttestation(address _updater, bytes _attestation) returns(bool)
func (_HomeHarness *HomeHarnessSession) ImproperAttestation(_updater common.Address, _attestation []byte) (*types.Transaction, error) {
	return _HomeHarness.Contract.ImproperAttestation(&_HomeHarness.TransactOpts, _updater, _attestation)
}

// ImproperAttestation is a paid mutator transaction binding the contract method 0x88a278ec.
//
// Solidity: function improperAttestation(address _updater, bytes _attestation) returns(bool)
func (_HomeHarness *HomeHarnessTransactorSession) ImproperAttestation(_updater common.Address, _attestation []byte) (*types.Transaction, error) {
	return _HomeHarness.Contract.ImproperAttestation(&_HomeHarness.TransactOpts, _updater, _attestation)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _updaterManager) returns()
func (_HomeHarness *HomeHarnessTransactor) Initialize(opts *bind.TransactOpts, _updaterManager common.Address) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "initialize", _updaterManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _updaterManager) returns()
func (_HomeHarness *HomeHarnessSession) Initialize(_updaterManager common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.Initialize(&_HomeHarness.TransactOpts, _updaterManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _updaterManager) returns()
func (_HomeHarness *HomeHarnessTransactorSession) Initialize(_updaterManager common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.Initialize(&_HomeHarness.TransactOpts, _updaterManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HomeHarness *HomeHarnessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HomeHarness *HomeHarnessSession) RenounceOwnership() (*types.Transaction, error) {
	return _HomeHarness.Contract.RenounceOwnership(&_HomeHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HomeHarness *HomeHarnessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HomeHarness.Contract.RenounceOwnership(&_HomeHarness.TransactOpts)
}

// SetFailed is a paid mutator transaction binding the contract method 0x146901db.
//
// Solidity: function setFailed() returns()
func (_HomeHarness *HomeHarnessTransactor) SetFailed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "setFailed")
}

// SetFailed is a paid mutator transaction binding the contract method 0x146901db.
//
// Solidity: function setFailed() returns()
func (_HomeHarness *HomeHarnessSession) SetFailed() (*types.Transaction, error) {
	return _HomeHarness.Contract.SetFailed(&_HomeHarness.TransactOpts)
}

// SetFailed is a paid mutator transaction binding the contract method 0x146901db.
//
// Solidity: function setFailed() returns()
func (_HomeHarness *HomeHarnessTransactorSession) SetFailed() (*types.Transaction, error) {
	return _HomeHarness.Contract.SetFailed(&_HomeHarness.TransactOpts)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x48639d24.
//
// Solidity: function setSensitiveValue(uint256 _newValue) returns()
func (_HomeHarness *HomeHarnessTransactor) SetSensitiveValue(opts *bind.TransactOpts, _newValue *big.Int) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "setSensitiveValue", _newValue)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x48639d24.
//
// Solidity: function setSensitiveValue(uint256 _newValue) returns()
func (_HomeHarness *HomeHarnessSession) SetSensitiveValue(_newValue *big.Int) (*types.Transaction, error) {
	return _HomeHarness.Contract.SetSensitiveValue(&_HomeHarness.TransactOpts, _newValue)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x48639d24.
//
// Solidity: function setSensitiveValue(uint256 _newValue) returns()
func (_HomeHarness *HomeHarnessTransactorSession) SetSensitiveValue(_newValue *big.Int) (*types.Transaction, error) {
	return _HomeHarness.Contract.SetSensitiveValue(&_HomeHarness.TransactOpts, _newValue)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_HomeHarness *HomeHarnessTransactor) SetSystemMessenger(opts *bind.TransactOpts, _systemMessenger common.Address) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "setSystemMessenger", _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_HomeHarness *HomeHarnessSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.SetSystemMessenger(&_HomeHarness.TransactOpts, _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_HomeHarness *HomeHarnessTransactorSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.SetSystemMessenger(&_HomeHarness.TransactOpts, _systemMessenger)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_HomeHarness *HomeHarnessTransactor) SetUpdater(opts *bind.TransactOpts, _updater common.Address) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "setUpdater", _updater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_HomeHarness *HomeHarnessSession) SetUpdater(_updater common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.SetUpdater(&_HomeHarness.TransactOpts, _updater)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
//
// Solidity: function setUpdater(address _updater) returns()
func (_HomeHarness *HomeHarnessTransactorSession) SetUpdater(_updater common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.SetUpdater(&_HomeHarness.TransactOpts, _updater)
}

// SetUpdaterManager is a paid mutator transaction binding the contract method 0x9776120e.
//
// Solidity: function setUpdaterManager(address _updaterManager) returns()
func (_HomeHarness *HomeHarnessTransactor) SetUpdaterManager(opts *bind.TransactOpts, _updaterManager common.Address) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "setUpdaterManager", _updaterManager)
}

// SetUpdaterManager is a paid mutator transaction binding the contract method 0x9776120e.
//
// Solidity: function setUpdaterManager(address _updaterManager) returns()
func (_HomeHarness *HomeHarnessSession) SetUpdaterManager(_updaterManager common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.SetUpdaterManager(&_HomeHarness.TransactOpts, _updaterManager)
}

// SetUpdaterManager is a paid mutator transaction binding the contract method 0x9776120e.
//
// Solidity: function setUpdaterManager(address _updaterManager) returns()
func (_HomeHarness *HomeHarnessTransactorSession) SetUpdaterManager(_updaterManager common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.SetUpdaterManager(&_HomeHarness.TransactOpts, _updaterManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HomeHarness *HomeHarnessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HomeHarness.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HomeHarness *HomeHarnessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.TransferOwnership(&_HomeHarness.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HomeHarness *HomeHarnessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HomeHarness.Contract.TransferOwnership(&_HomeHarness.TransactOpts, newOwner)
}

// HomeHarnessDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the HomeHarness contract.
type HomeHarnessDispatchIterator struct {
	Event *HomeHarnessDispatch // Event containing the contract specifics and raw log

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
func (it *HomeHarnessDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeHarnessDispatch)
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
		it.Event = new(HomeHarnessDispatch)
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
func (it *HomeHarnessDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeHarnessDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeHarnessDispatch represents a Dispatch event raised by the HomeHarness contract.
type HomeHarnessDispatch struct {
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
func (_HomeHarness *HomeHarnessFilterer) FilterDispatch(opts *bind.FilterOpts, messageHash [][32]byte, leafIndex []*big.Int, destinationAndNonce []uint64) (*HomeHarnessDispatchIterator, error) {

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

	logs, sub, err := _HomeHarness.contract.FilterLogs(opts, "Dispatch", messageHashRule, leafIndexRule, destinationAndNonceRule)
	if err != nil {
		return nil, err
	}
	return &HomeHarnessDispatchIterator{contract: _HomeHarness.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0x718e547b45036b0526c0cd2f2e3de248b0e8c042c714ecfbee3f5811a5e6e785.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint256 indexed leafIndex, uint64 indexed destinationAndNonce, bytes tips, bytes message)
func (_HomeHarness *HomeHarnessFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *HomeHarnessDispatch, messageHash [][32]byte, leafIndex []*big.Int, destinationAndNonce []uint64) (event.Subscription, error) {

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

	logs, sub, err := _HomeHarness.contract.WatchLogs(opts, "Dispatch", messageHashRule, leafIndexRule, destinationAndNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeHarnessDispatch)
				if err := _HomeHarness.contract.UnpackLog(event, "Dispatch", log); err != nil {
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
func (_HomeHarness *HomeHarnessFilterer) ParseDispatch(log types.Log) (*HomeHarnessDispatch, error) {
	event := new(HomeHarnessDispatch)
	if err := _HomeHarness.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeHarnessImproperAttestationIterator is returned from FilterImproperAttestation and is used to iterate over the raw logs and unpacked data for ImproperAttestation events raised by the HomeHarness contract.
type HomeHarnessImproperAttestationIterator struct {
	Event *HomeHarnessImproperAttestation // Event containing the contract specifics and raw log

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
func (it *HomeHarnessImproperAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeHarnessImproperAttestation)
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
		it.Event = new(HomeHarnessImproperAttestation)
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
func (it *HomeHarnessImproperAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeHarnessImproperAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeHarnessImproperAttestation represents a ImproperAttestation event raised by the HomeHarness contract.
type HomeHarnessImproperAttestation struct {
	Updater     common.Address
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterImproperAttestation is a free log retrieval operation binding the contract event 0x287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b.
//
// Solidity: event ImproperAttestation(address updater, bytes attestation)
func (_HomeHarness *HomeHarnessFilterer) FilterImproperAttestation(opts *bind.FilterOpts) (*HomeHarnessImproperAttestationIterator, error) {

	logs, sub, err := _HomeHarness.contract.FilterLogs(opts, "ImproperAttestation")
	if err != nil {
		return nil, err
	}
	return &HomeHarnessImproperAttestationIterator{contract: _HomeHarness.contract, event: "ImproperAttestation", logs: logs, sub: sub}, nil
}

// WatchImproperAttestation is a free log subscription operation binding the contract event 0x287e2c0e041ca31a0ce7a1ed8b91a7425b2520880947cdbe778c457ca4c48e5b.
//
// Solidity: event ImproperAttestation(address updater, bytes attestation)
func (_HomeHarness *HomeHarnessFilterer) WatchImproperAttestation(opts *bind.WatchOpts, sink chan<- *HomeHarnessImproperAttestation) (event.Subscription, error) {

	logs, sub, err := _HomeHarness.contract.WatchLogs(opts, "ImproperAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeHarnessImproperAttestation)
				if err := _HomeHarness.contract.UnpackLog(event, "ImproperAttestation", log); err != nil {
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
// Solidity: event ImproperAttestation(address updater, bytes attestation)
func (_HomeHarness *HomeHarnessFilterer) ParseImproperAttestation(log types.Log) (*HomeHarnessImproperAttestation, error) {
	event := new(HomeHarnessImproperAttestation)
	if err := _HomeHarness.contract.UnpackLog(event, "ImproperAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeHarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HomeHarness contract.
type HomeHarnessInitializedIterator struct {
	Event *HomeHarnessInitialized // Event containing the contract specifics and raw log

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
func (it *HomeHarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeHarnessInitialized)
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
		it.Event = new(HomeHarnessInitialized)
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
func (it *HomeHarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeHarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeHarnessInitialized represents a Initialized event raised by the HomeHarness contract.
type HomeHarnessInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HomeHarness *HomeHarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*HomeHarnessInitializedIterator, error) {

	logs, sub, err := _HomeHarness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HomeHarnessInitializedIterator{contract: _HomeHarness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HomeHarness *HomeHarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HomeHarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _HomeHarness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeHarnessInitialized)
				if err := _HomeHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_HomeHarness *HomeHarnessFilterer) ParseInitialized(log types.Log) (*HomeHarnessInitialized, error) {
	event := new(HomeHarnessInitialized)
	if err := _HomeHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeHarnessNewUpdaterIterator is returned from FilterNewUpdater and is used to iterate over the raw logs and unpacked data for NewUpdater events raised by the HomeHarness contract.
type HomeHarnessNewUpdaterIterator struct {
	Event *HomeHarnessNewUpdater // Event containing the contract specifics and raw log

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
func (it *HomeHarnessNewUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeHarnessNewUpdater)
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
		it.Event = new(HomeHarnessNewUpdater)
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
func (it *HomeHarnessNewUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeHarnessNewUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeHarnessNewUpdater represents a NewUpdater event raised by the HomeHarness contract.
type HomeHarnessNewUpdater struct {
	OldUpdater common.Address
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUpdater is a free log retrieval operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_HomeHarness *HomeHarnessFilterer) FilterNewUpdater(opts *bind.FilterOpts) (*HomeHarnessNewUpdaterIterator, error) {

	logs, sub, err := _HomeHarness.contract.FilterLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return &HomeHarnessNewUpdaterIterator{contract: _HomeHarness.contract, event: "NewUpdater", logs: logs, sub: sub}, nil
}

// WatchNewUpdater is a free log subscription operation binding the contract event 0x0f20622a7af9e952a6fec654a196f29e04477b5d335772c26902bec35cc9f22a.
//
// Solidity: event NewUpdater(address oldUpdater, address newUpdater)
func (_HomeHarness *HomeHarnessFilterer) WatchNewUpdater(opts *bind.WatchOpts, sink chan<- *HomeHarnessNewUpdater) (event.Subscription, error) {

	logs, sub, err := _HomeHarness.contract.WatchLogs(opts, "NewUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeHarnessNewUpdater)
				if err := _HomeHarness.contract.UnpackLog(event, "NewUpdater", log); err != nil {
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
func (_HomeHarness *HomeHarnessFilterer) ParseNewUpdater(log types.Log) (*HomeHarnessNewUpdater, error) {
	event := new(HomeHarnessNewUpdater)
	if err := _HomeHarness.contract.UnpackLog(event, "NewUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeHarnessNewUpdaterManagerIterator is returned from FilterNewUpdaterManager and is used to iterate over the raw logs and unpacked data for NewUpdaterManager events raised by the HomeHarness contract.
type HomeHarnessNewUpdaterManagerIterator struct {
	Event *HomeHarnessNewUpdaterManager // Event containing the contract specifics and raw log

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
func (it *HomeHarnessNewUpdaterManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeHarnessNewUpdaterManager)
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
		it.Event = new(HomeHarnessNewUpdaterManager)
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
func (it *HomeHarnessNewUpdaterManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeHarnessNewUpdaterManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeHarnessNewUpdaterManager represents a NewUpdaterManager event raised by the HomeHarness contract.
type HomeHarnessNewUpdaterManager struct {
	UpdaterManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewUpdaterManager is a free log retrieval operation binding the contract event 0x958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf.
//
// Solidity: event NewUpdaterManager(address updaterManager)
func (_HomeHarness *HomeHarnessFilterer) FilterNewUpdaterManager(opts *bind.FilterOpts) (*HomeHarnessNewUpdaterManagerIterator, error) {

	logs, sub, err := _HomeHarness.contract.FilterLogs(opts, "NewUpdaterManager")
	if err != nil {
		return nil, err
	}
	return &HomeHarnessNewUpdaterManagerIterator{contract: _HomeHarness.contract, event: "NewUpdaterManager", logs: logs, sub: sub}, nil
}

// WatchNewUpdaterManager is a free log subscription operation binding the contract event 0x958d788fb4c373604cd4c73aa8c592de127d0819b49bb4dc02c8ecd666e965bf.
//
// Solidity: event NewUpdaterManager(address updaterManager)
func (_HomeHarness *HomeHarnessFilterer) WatchNewUpdaterManager(opts *bind.WatchOpts, sink chan<- *HomeHarnessNewUpdaterManager) (event.Subscription, error) {

	logs, sub, err := _HomeHarness.contract.WatchLogs(opts, "NewUpdaterManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeHarnessNewUpdaterManager)
				if err := _HomeHarness.contract.UnpackLog(event, "NewUpdaterManager", log); err != nil {
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
func (_HomeHarness *HomeHarnessFilterer) ParseNewUpdaterManager(log types.Log) (*HomeHarnessNewUpdaterManager, error) {
	event := new(HomeHarnessNewUpdaterManager)
	if err := _HomeHarness.contract.UnpackLog(event, "NewUpdaterManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeHarnessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HomeHarness contract.
type HomeHarnessOwnershipTransferredIterator struct {
	Event *HomeHarnessOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HomeHarnessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeHarnessOwnershipTransferred)
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
		it.Event = new(HomeHarnessOwnershipTransferred)
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
func (it *HomeHarnessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeHarnessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeHarnessOwnershipTransferred represents a OwnershipTransferred event raised by the HomeHarness contract.
type HomeHarnessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HomeHarness *HomeHarnessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HomeHarnessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HomeHarness.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HomeHarnessOwnershipTransferredIterator{contract: _HomeHarness.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HomeHarness *HomeHarnessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HomeHarnessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HomeHarness.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeHarnessOwnershipTransferred)
				if err := _HomeHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HomeHarness *HomeHarnessFilterer) ParseOwnershipTransferred(log types.Log) (*HomeHarnessOwnershipTransferred, error) {
	event := new(HomeHarnessOwnershipTransferred)
	if err := _HomeHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeHarnessUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the HomeHarness contract.
type HomeHarnessUpdateIterator struct {
	Event *HomeHarnessUpdate // Event containing the contract specifics and raw log

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
func (it *HomeHarnessUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeHarnessUpdate)
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
		it.Event = new(HomeHarnessUpdate)
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
func (it *HomeHarnessUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeHarnessUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeHarnessUpdate represents a Update event raised by the HomeHarness contract.
type HomeHarnessUpdate struct {
	HomeDomain uint32
	Nonce      uint32
	Root       [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_HomeHarness *HomeHarnessFilterer) FilterUpdate(opts *bind.FilterOpts, homeDomain []uint32, nonce []uint32, root [][32]byte) (*HomeHarnessUpdateIterator, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _HomeHarness.contract.FilterLogs(opts, "Update", homeDomainRule, nonceRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &HomeHarnessUpdateIterator{contract: _HomeHarness.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_HomeHarness *HomeHarnessFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *HomeHarnessUpdate, homeDomain []uint32, nonce []uint32, root [][32]byte) (event.Subscription, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _HomeHarness.contract.WatchLogs(opts, "Update", homeDomainRule, nonceRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeHarnessUpdate)
				if err := _HomeHarness.contract.UnpackLog(event, "Update", log); err != nil {
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

// ParseUpdate is a log parse operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_HomeHarness *HomeHarnessFilterer) ParseUpdate(log types.Log) (*HomeHarnessUpdate, error) {
	event := new(HomeHarnessUpdate)
	if err := _HomeHarness.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HomeHarnessUpdaterSlashedIterator is returned from FilterUpdaterSlashed and is used to iterate over the raw logs and unpacked data for UpdaterSlashed events raised by the HomeHarness contract.
type HomeHarnessUpdaterSlashedIterator struct {
	Event *HomeHarnessUpdaterSlashed // Event containing the contract specifics and raw log

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
func (it *HomeHarnessUpdaterSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HomeHarnessUpdaterSlashed)
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
		it.Event = new(HomeHarnessUpdaterSlashed)
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
func (it *HomeHarnessUpdaterSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HomeHarnessUpdaterSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HomeHarnessUpdaterSlashed represents a UpdaterSlashed event raised by the HomeHarness contract.
type HomeHarnessUpdaterSlashed struct {
	Updater  common.Address
	Reporter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdaterSlashed is a free log retrieval operation binding the contract event 0x98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce.
//
// Solidity: event UpdaterSlashed(address indexed updater, address indexed reporter)
func (_HomeHarness *HomeHarnessFilterer) FilterUpdaterSlashed(opts *bind.FilterOpts, updater []common.Address, reporter []common.Address) (*HomeHarnessUpdaterSlashedIterator, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _HomeHarness.contract.FilterLogs(opts, "UpdaterSlashed", updaterRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return &HomeHarnessUpdaterSlashedIterator{contract: _HomeHarness.contract, event: "UpdaterSlashed", logs: logs, sub: sub}, nil
}

// WatchUpdaterSlashed is a free log subscription operation binding the contract event 0x98064af315f26d7333ba107ba43a128ec74345f4d4e6f2549840fe092a1c8bce.
//
// Solidity: event UpdaterSlashed(address indexed updater, address indexed reporter)
func (_HomeHarness *HomeHarnessFilterer) WatchUpdaterSlashed(opts *bind.WatchOpts, sink chan<- *HomeHarnessUpdaterSlashed, updater []common.Address, reporter []common.Address) (event.Subscription, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _HomeHarness.contract.WatchLogs(opts, "UpdaterSlashed", updaterRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HomeHarnessUpdaterSlashed)
				if err := _HomeHarness.contract.UnpackLog(event, "UpdaterSlashed", log); err != nil {
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
func (_HomeHarness *HomeHarnessFilterer) ParseUpdaterSlashed(log types.Log) (*HomeHarnessUpdaterSlashed, error) {
	event := new(HomeHarnessUpdaterSlashed)
	if err := _HomeHarness.contract.UnpackLog(event, "UpdaterSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISystemMessengerMetaData contains all meta data concerning the ISystemMessenger contract.
var ISystemMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destDomain\",\"type\":\"uint32\"},{\"internalType\":\"enumISystemMessenger.SystemContracts\",\"name\":\"_recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"sendSystemMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function sendSystemMessage(uint32 _destDomain, uint8 _recipient, bytes _payload) returns()
func (_ISystemMessenger *ISystemMessengerTransactor) SendSystemMessage(opts *bind.TransactOpts, _destDomain uint32, _recipient uint8, _payload []byte) (*types.Transaction, error) {
	return _ISystemMessenger.contract.Transact(opts, "sendSystemMessage", _destDomain, _recipient, _payload)
}

// SendSystemMessage is a paid mutator transaction binding the contract method 0x0d1e27a7.
//
// Solidity: function sendSystemMessage(uint32 _destDomain, uint8 _recipient, bytes _payload) returns()
func (_ISystemMessenger *ISystemMessengerSession) SendSystemMessage(_destDomain uint32, _recipient uint8, _payload []byte) (*types.Transaction, error) {
	return _ISystemMessenger.Contract.SendSystemMessage(&_ISystemMessenger.TransactOpts, _destDomain, _recipient, _payload)
}

// SendSystemMessage is a paid mutator transaction binding the contract method 0x0d1e27a7.
//
// Solidity: function sendSystemMessage(uint32 _destDomain, uint8 _recipient, bytes _payload) returns()
func (_ISystemMessenger *ISystemMessengerTransactorSession) SendSystemMessage(_destDomain uint32, _recipient uint8, _payload []byte) (*types.Transaction, error) {
	return _ISystemMessenger.Contract.SendSystemMessage(&_ISystemMessenger.TransactOpts, _destDomain, _recipient, _payload)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c4288c1f68a87ee3a9b575df9764900d888b03d4d9fabd5cdc409a6247770cc764736f6c634300080d0033",
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
	Bin: "0x608060405234801561001057600080fd5b506106f8806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806306661abd146100515780637ea97f4014610067578063ebf0c7171461007a578063fd54b22814610082575b600080fd5b6020545b60405190815260200160405180910390f35b61005561007536600461067a565b61008c565b6100556100ad565b6020546100559081565b6021818154811061009c57600080fd5b600091825260209091200154905081565b60006100b960006100be565b905090565b60006100d1826100cc6100d7565b610598565b92915050565b6100df61065b565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b602081101561065357600182821c8116908190036105ff578582602081106105cc576105cc610693565b0154604080516020810192909252810185905260600160405160208183030381529060405280519060200120935061064a565b8385836020811061061257610612610693565b6020020151604051602001610631929190918252602082015260400190565b6040516020818303038152906040528051906020012093505b506001016105a2565b505092915050565b6040518061040001604052806020906020820280368337509192915050565b60006020828403121561068c57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea26469706673582212203f41440dedbb02370b286f84700a18356448f4924214a79ae10cefb33dae524764736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200d5bbbe98489fd276117b61389449a22e1ce36a34386d7e610701602775f81e964736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220241b4b2636bac55ba7605d1ab4d2420caca83c931b9d6f23aa426599b236ae0e64736f6c634300080d0033",
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

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204ed28f530a3904e0f2ceefc952499d91c9989ab8a05712fc3f85ae0284d90a3964736f6c634300080d0033",
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

// SystemMessageMetaData contains all meta data concerning the SystemMessage contract.
var SystemMessageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a31cb70514f871b15d34a18a8405ea3a56e7b9ecc7acbfdc606b731decaa752064736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b3bb2f27ef8789863fbf9d163a810f442fda4672f26ba0349baf19b2274d686264736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200f053d0c2d79d4ee08b507be4f4475efaf39164e504c2600f556f3df28a93fdb64736f6c634300080d0033",
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
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea264697066735822122073a7a1a880339f11f867f2e26cd41c4fa9d3d781fee5558ade7c58893f22793564736f6c634300080d0033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newUpdater\",\"type\":\"address\"}],\"name\":\"NewUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"homeDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"Update\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"_systemMessenger\",\"type\":\"address\"}],\"name\":\"setSystemMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemMessenger\",\"outputs\":[{\"internalType\":\"contractISystemMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"b7bc563e": "setSystemMessenger(address)",
		"ccbdf9c9": "systemMessenger()",
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

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_UpdaterStorage *UpdaterStorageCaller) SystemMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpdaterStorage.contract.Call(opts, &out, "systemMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_UpdaterStorage *UpdaterStorageSession) SystemMessenger() (common.Address, error) {
	return _UpdaterStorage.Contract.SystemMessenger(&_UpdaterStorage.CallOpts)
}

// SystemMessenger is a free data retrieval call binding the contract method 0xccbdf9c9.
//
// Solidity: function systemMessenger() view returns(address)
func (_UpdaterStorage *UpdaterStorageCallerSession) SystemMessenger() (common.Address, error) {
	return _UpdaterStorage.Contract.SystemMessenger(&_UpdaterStorage.CallOpts)
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

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_UpdaterStorage *UpdaterStorageTransactor) SetSystemMessenger(opts *bind.TransactOpts, _systemMessenger common.Address) (*types.Transaction, error) {
	return _UpdaterStorage.contract.Transact(opts, "setSystemMessenger", _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_UpdaterStorage *UpdaterStorageSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _UpdaterStorage.Contract.SetSystemMessenger(&_UpdaterStorage.TransactOpts, _systemMessenger)
}

// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
//
// Solidity: function setSystemMessenger(address _systemMessenger) returns()
func (_UpdaterStorage *UpdaterStorageTransactorSession) SetSystemMessenger(_systemMessenger common.Address) (*types.Transaction, error) {
	return _UpdaterStorage.Contract.SetSystemMessenger(&_UpdaterStorage.TransactOpts, _systemMessenger)
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
	Nonce      uint32
	Root       [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_UpdaterStorage *UpdaterStorageFilterer) FilterUpdate(opts *bind.FilterOpts, homeDomain []uint32, nonce []uint32, root [][32]byte) (*UpdaterStorageUpdateIterator, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _UpdaterStorage.contract.FilterLogs(opts, "Update", homeDomainRule, nonceRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &UpdaterStorageUpdateIterator{contract: _UpdaterStorage.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_UpdaterStorage *UpdaterStorageFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *UpdaterStorageUpdate, homeDomain []uint32, nonce []uint32, root [][32]byte) (event.Subscription, error) {

	var homeDomainRule []interface{}
	for _, homeDomainItem := range homeDomain {
		homeDomainRule = append(homeDomainRule, homeDomainItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _UpdaterStorage.contract.WatchLogs(opts, "Update", homeDomainRule, nonceRule, rootRule)
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

// ParseUpdate is a log parse operation binding the contract event 0x3f459c2c4e333807b9c629230cbac6a23dbfd53c030ef9bc6886abb97ada9171.
//
// Solidity: event Update(uint32 indexed homeDomain, uint32 indexed nonce, bytes32 indexed root, bytes signature)
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
	Bin: "0x6080604052348015600f57600080fd5b5060808061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063ffa1ad7414602d575b600080fd5b6034600081565b60405160ff909116815260200160405180910390f3fea264697066735822122043d649d4204e482da3f56ed7eb8af07f60eef1dbf9c0b2b0de6af3540f4575ab64736f6c634300080d0033",
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
